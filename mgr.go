// Copyright (c) 2021, Athonet S.r.l. All rights reserved.
// riccardo.manfrin@athonet.com

package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"time"

	oas "github.com/riccardomanfrin/netconfd/server/go"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gorilla/mux"
	"github.com/riccardomanfrin/netconfd/comm"
	"github.com/riccardomanfrin/netconfd/logger"
	"github.com/riccardomanfrin/netconfd/nc"
)

//RouteHandler is the handler for a certain route OperationID
type RouteHandler func(body interface{}) (interface{}, error)

//Manager of the machine takes HTTP requests, perform actions and give results back [blocking]
type Manager struct {
	Conf           *oas.Config
	HTTPServer     *http.Server
	ServeMux       *http.ServeMux
	SyncHTTPServer *http.Server
	SyncServeMux   *http.ServeMux
	client         http.Client
	m2MEndpointURL string
	openapi        *openapi3.Swagger
	router         routers.Router
	routeHandlers  *mux.Router
}

//PushContent describes the push content
type PushContent struct {
	FilePath       string `json:"filepath"`
	B64FileContent string `json:"b64filecontent"`
}

//File encodes the sync state of a file between local and remote
type File struct {
	Name      string `json:"name"`
	LocalMd5  string `json:"local_md5"`
	RemoteMd5 string `json:"remote_md5"`
	Aligned   bool   `json:"aligned"`
}

//FilesState encodes the state of all monitored files among remote and local
type FilesState struct {
	Files []File `json:"files"`
}

func urlMatch(in string, expected string) bool {
	return strings.Contains(in, expected)
}

func fail(err int, log string) {
	logger.Log.Error(log)
	os.Exit(err)
}

func httpError(w http.ResponseWriter, e error, code int) {
	logger.Log.Warning(e.Error())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	io.WriteString(w, comm.ToJSONString(e))
}

func httpOk(w http.ResponseWriter) {
	w.WriteHeader(200)
}

//LoadConfig loads the configuration
func (m *Manager) LoadConfig(conffile *string) error {
	err := m.Conf.LoadConfig(conffile)
	if err != nil {
		logger.Log.Warning(err.Error())
		logger.Log.Info("Trying to apply a failsafe config...")
		m.Conf.Global = oas.NewGlobal()
		m.Conf.Global.Mgmt = oas.NewGlobalMgmtWithDefaults()
		if err = m.patchFailSafeConfig(); err != nil {
			logger.Log.Warning(err.Error())
		}
	}

	loglev := "INF"

	if (m.Conf.Global != nil) && (m.Conf.Global.LogLev != nil) {
		loglev = *m.Conf.Global.LogLev
	}

	logger.LoggerSetLevel(loglev)

	host := "0.0.0.0"
	port := os.Getenv("PORT")
	m.Conf.Global.Mgmt.Host = &host
	if port != "" {
		p, _ := strconv.ParseInt(port, 10, 0)
		pint32 := int32(p)
		m.Conf.Global.Mgmt.Port = &pint32
	}
	if m.Conf.Global != nil && m.Conf.Global.Mgmt != nil {
		if m.Conf.Global.Mgmt.Host != nil {
			host = *m.Conf.Global.Mgmt.Host
		}
		if m.Conf.Global.Mgmt.Port != nil {
			port = strconv.FormatUint(uint64(*m.Conf.Global.Mgmt.Port), 10)
		}
	}

	logger.Log.Notice("Starting mgmt service on " + host + ":" + port)
	logger.Log.Notice("Log level set to " + loglev)

	m.HTTPServer = &http.Server{
		Addr:           host + ":" + port,
		Handler:        m.ServeMux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return nil
}

func (m *Manager) patchConfig(Conf *oas.Config) error {
	reqbody, err := json.Marshal(Conf)
	if err != nil {
		return err
	}
	logger.Log.Notice(fmt.Sprintf("Config: %v", string(reqbody)))
	iobody := bytes.NewReader(reqbody)
	req, _ := http.NewRequest("PATCH", "/api/1/mgmt/config", iobody)
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	m.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		return fmt.Errorf("Failed to patch-apply config - Error %v - Network Not configured", rr)
	}
	return nil
}

func (m *Manager) enforceUnmanaged() error {
	logger.Log.Info("Loading Unmanaged resources config")
	c := oas.Config{
		Network: &oas.Network{
			Unmanaged: m.Conf.Network.Unmanaged,
		},
	}
	return m.patchConfig(&c)
}
func (m *Manager) patchInitialConfig() error {
	logger.Log.Info("Patching with persisted initial config...")
	return m.patchConfig(m.Conf)
}

func (m *Manager) patchFailSafeConfig() error {
	logger.Log.Info("Patching with failsafe config...")
	failSafeConfig := `
	{
		"links": [
			{
				"flags": [
					"up"
				],
				"ifindex": 1,
				"ifname": "lo",
				"link_type": "loopback"
			},
			{
				"flags": [
					"up"
				],
				"ifname": "eth0",
				"link_type": "ether",
				"linkinfo": {
					"info_kind": "device"
				}
			}
		],
		"dhcp": [
			{
				"ifname": "eth0"
			}
		]
	}`
	logger.LoggerSetLevel("DBG")
	network := oas.Network{}
	if err := json.Unmarshal([]byte(failSafeConfig), &network); err != nil {
		return err
	}
	m.Conf.Network = &network
	if err := m.patchConfig(m.Conf); err != nil {
		return err
	}

	return nil
}

func passThrough(c context.Context, input *openapi3filter.AuthenticationInput) error {
	return nil
}

func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	// Find route
	route, pathParams, _ := m.router.FindRoute(r)

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    r,
		PathParams: pathParams,
		Route:      route,
		Options: &openapi3filter.Options{
			AuthenticationFunc: passThrough,
		},
	}
	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		err = &comm.ErrorString{S: fmt.Sprintf("HTTP API validation error: %v", err.Error())}
		httpError(w, err, http.StatusBadRequest)
		return
	}

	wrw := &wrapperResponseWriter{
		ResponseWriter: w,
		buf:            &bytes.Buffer{},
	}

	m.routeHandlers.ServeHTTP(wrw, r)

	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 wrw.Status,
		Header:                 w.Header(),
	}

	responseValidationInput.SetBodyBytes(wrw.buf.Bytes())

	// Validate response.
	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		panic(err)
	}

	if (wrw.Status != http.StatusOK) && (wrw.Status != http.StatusCreated) {
		logger.Log.Warning("HTTP Error", wrw.Status, ":", string(wrw.buf.Bytes()))
	} else {
		logger.Log.Debug("HTTP Ok", wrw.Status, ":", string(wrw.buf.Bytes()))
	}
}
func checkIP(ipstr string) error {
	if net.ParseIP(ipstr) == nil {
		return nc.NewInvalidIPAddressError(ipstr)
	}
	return nil
}

//go:embed swaggerui/*
var swaggeruiFS embed.FS

//go:embed swaggerui/openapi.yaml
var openAPIYAML []byte

//NewManager creates new manager
func NewManager() *Manager {
	openapi, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(openAPIYAML)
	if openapi == nil && err != nil {
		panic(err)
	}

	//Opt in IPv4 and IPv6 validation
	openapi3.DefineIPv4Format()
	openapi3.DefineIPv6Format()
	openapi3.DefineStringFormatCallback("cidr", nc.CIDRAddrValidate)
	openapi3.DefineStringFormatCallback("ip", checkIP)

	router, err := gorillamux.NewRouter(openapi)
	if err != nil {
		log.Fatal(err)
	}
	serveMux := http.NewServeMux()

	networkApiService := oas.NewNetworkApiService()
	networkApiController := oas.NewNetworkApiController(networkApiService)
	systemApiService := oas.NewSystemApiService()
	systemAPIController := oas.NewSystemApiController(systemApiService)
	m := &Manager{
		openapi:       openapi,
		router:        router,
		ServeMux:      serveMux,
		routeHandlers: oas.NewRouter(networkApiController, systemAPIController),
	}
	sysApiService, _ := systemApiService.(*oas.SystemApiService)
	m.Conf = &sysApiService.Conf
	serveMux.Handle("/", m)

	serveMux.Handle("/swaggerui/", http.FileServer(http.FS(swaggeruiFS)))
	nc.InitErrorsLogsTracing()
	return m
}

//ListenAndServe activates HTTP server
func (m *Manager) ListenAndServe() {
	err := m.HTTPServer.ListenAndServe()
	logger.Fatal(err)
}

func (m *Manager) ifaceRorder() error {
	return nc.LinksVMReorder()
}

//Start activates manager
func (m *Manager) Start() {
	if *skipbootconfig == false {
		if err := m.ifaceRorder(); err != nil {
			logger.Log.Warning(fmt.Sprintf("Failed to reorder interfaces: %v", err.Error()))
		}
		if err := m.patchInitialConfig(); err != nil {
			logger.Log.Warning(err.Error())
			if err := m.patchFailSafeConfig(); err != nil {
				logger.Log.Warning(err.Error())
			}
		}
	} else {
		logger.Log.Info("Skipped boot configuration enforcement")
		if err := m.enforceUnmanaged(); err != nil {
			logger.Log.Warning(fmt.Sprintf("Failed to load Unmanaged resources: %v", err.Error()))
		}
	}
	logger.Log.Notice("Started netConfD manager...")

	go m.ListenAndServe()
}

type wrapperResponseWriter struct {
	http.ResponseWriter
	buf    *bytes.Buffer
	Status int
}

func (wrw *wrapperResponseWriter) Write(p []byte) (int, error) {
	wrw.ResponseWriter.Write(p)
	return wrw.buf.Write(p)
}

//WriteHeader wrapper to record status
func (wrw *wrapperResponseWriter) WriteHeader(status int) {
	wrw.Status = status
	wrw.ResponseWriter.WriteHeader(status)
}
