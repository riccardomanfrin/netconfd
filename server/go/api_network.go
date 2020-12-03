/*
 * netConfD API
 *
 * Network Configurator service
 *
 * API version: 0.1.0
 * Contact: support@athonet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A NetworkApiController binds http requests to an api service and writes the service results to the http response
type NetworkApiController struct {
	service NetworkApiServicer
}

// NewNetworkApiController creates a default api controller
func NewNetworkApiController(s NetworkApiServicer) Router {
	return &NetworkApiController{ service: s }
}

// Routes returns all of the api route for the NetworkApiController
func (c *NetworkApiController) Routes() Routes {
	return Routes{ 
		{
			"ConfigLinkCreate",
			strings.ToUpper("Post"),
			"/api/1/links",
			c.ConfigLinkCreate,
		},
		{
			"ConfigLinkDel",
			strings.ToUpper("Delete"),
			"/api/1/links/{ifname}",
			c.ConfigLinkDel,
		},
		{
			"ConfigLinkGet",
			strings.ToUpper("Get"),
			"/api/1/links/{ifname}",
			c.ConfigLinkGet,
		},
		{
			"ConfigLinksGet",
			strings.ToUpper("Get"),
			"/api/1/links",
			c.ConfigLinksGet,
		},
		{
			"ConfigNFTableCreate",
			strings.ToUpper("Post"),
			"/api/1/nftables",
			c.ConfigNFTableCreate,
		},
		{
			"ConfigNFTableDel",
			strings.ToUpper("Delete"),
			"/api/1/ntables/{nftableid}",
			c.ConfigNFTableDel,
		},
		{
			"ConfigNFTableGet",
			strings.ToUpper("Get"),
			"/api/1/ntables/{nftableid}",
			c.ConfigNFTableGet,
		},
		{
			"ConfigNFTablesGet",
			strings.ToUpper("Get"),
			"/api/1/nftables",
			c.ConfigNFTablesGet,
		},
		{
			"ConfigNetNSCreate",
			strings.ToUpper("Post"),
			"/api/1/netns",
			c.ConfigNetNSCreate,
		},
		{
			"ConfigNetNSDel",
			strings.ToUpper("Delete"),
			"/api/1/netns/{netnsid}",
			c.ConfigNetNSDel,
		},
		{
			"ConfigNetNSGet",
			strings.ToUpper("Get"),
			"/api/1/netns/{netnsid}",
			c.ConfigNetNSGet,
		},
		{
			"ConfigNetNSsGet",
			strings.ToUpper("Get"),
			"/api/1/netns",
			c.ConfigNetNSsGet,
		},
		{
			"ConfigRouteCreate",
			strings.ToUpper("Post"),
			"/api/1/routes",
			c.ConfigRouteCreate,
		},
		{
			"ConfigRouteDel",
			strings.ToUpper("Delete"),
			"/api/1/routes/{routeid}",
			c.ConfigRouteDel,
		},
		{
			"ConfigRouteGet",
			strings.ToUpper("Get"),
			"/api/1/routes/{routeid}",
			c.ConfigRouteGet,
		},
		{
			"ConfigRoutesGet",
			strings.ToUpper("Get"),
			"/api/1/routes",
			c.ConfigRoutesGet,
		},
		{
			"ConfigRuleCreate",
			strings.ToUpper("Post"),
			"/api/1/rules",
			c.ConfigRuleCreate,
		},
		{
			"ConfigRuleDel",
			strings.ToUpper("Delete"),
			"/api/1/rules/{ruleid}",
			c.ConfigRuleDel,
		},
		{
			"ConfigRuleGet",
			strings.ToUpper("Get"),
			"/api/1/rules/{ruleid}",
			c.ConfigRuleGet,
		},
		{
			"ConfigRulesGet",
			strings.ToUpper("Get"),
			"/api/1/rules",
			c.ConfigRulesGet,
		},
		{
			"ConfigVRFCreate",
			strings.ToUpper("Post"),
			"/api/1/vrfs",
			c.ConfigVRFCreate,
		},
		{
			"ConfigVRFDel",
			strings.ToUpper("Delete"),
			"/api/1/vrfs/{vrfid}",
			c.ConfigVRFDel,
		},
		{
			"ConfigVRFGet",
			strings.ToUpper("Get"),
			"/api/1/vrfs/{vrfid}",
			c.ConfigVRFGet,
		},
		{
			"ConfigVRFsGet",
			strings.ToUpper("Get"),
			"/api/1/vrfs",
			c.ConfigVRFsGet,
		},
	}
}

// ConfigLinkCreate - Configures and brings up a link layer interface 
func (c *NetworkApiController) ConfigLinkCreate(w http.ResponseWriter, r *http.Request) { 
	link := &Link{}
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigLinkCreate(r.Context(), *link)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigLinkDel - Brings down and delete a link layer interface 
func (c *NetworkApiController) ConfigLinkDel(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	ifname := params["ifname"]
	result, err := c.service.ConfigLinkDel(r.Context(), ifname)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigLinkGet - Retrieve link layer interface information 
func (c *NetworkApiController) ConfigLinkGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	ifname := params["ifname"]
	result, err := c.service.ConfigLinkGet(r.Context(), ifname)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigLinksGet - Get all link layer interfaces 
func (c *NetworkApiController) ConfigLinksGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.ConfigLinksGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigNFTableCreate - Configures an new NFTable 
func (c *NetworkApiController) ConfigNFTableCreate(w http.ResponseWriter, r *http.Request) { 
	body := &map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigNFTableCreate(r.Context(), *body)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigNFTableDel - Removes a NFTable 
func (c *NetworkApiController) ConfigNFTableDel(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	nftableid, err := parseInt32Parameter(params["nftableid"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigNFTableDel(r.Context(), nftableid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigNFTableGet - Get a NFTable 
func (c *NetworkApiController) ConfigNFTableGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	nftableid, err := parseInt32Parameter(params["nftableid"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigNFTableGet(r.Context(), nftableid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigNFTablesGet - Get the list all NFTables 
func (c *NetworkApiController) ConfigNFTablesGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.ConfigNFTablesGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigNetNSCreate - Configures an new Network Namespace 
func (c *NetworkApiController) ConfigNetNSCreate(w http.ResponseWriter, r *http.Request) { 
	netns := &Netns{}
	if err := json.NewDecoder(r.Body).Decode(&netns); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigNetNSCreate(r.Context(), *netns)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigNetNSDel - Removes an IP Rule 
func (c *NetworkApiController) ConfigNetNSDel(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	netnsid := params["netnsid"]
	result, err := c.service.ConfigNetNSDel(r.Context(), netnsid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigNetNSGet - Get a network namespace 
func (c *NetworkApiController) ConfigNetNSGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	netnsid := params["netnsid"]
	result, err := c.service.ConfigNetNSGet(r.Context(), netnsid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigNetNSsGet - Get the list all network namespaces 
func (c *NetworkApiController) ConfigNetNSsGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.ConfigNetNSsGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigRouteCreate - Configures a route 
func (c *NetworkApiController) ConfigRouteCreate(w http.ResponseWriter, r *http.Request) { 
	route := &Route{}
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigRouteCreate(r.Context(), *route)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigRouteDel - Brings down and delete an L3 IP route 
func (c *NetworkApiController) ConfigRouteDel(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	routeid, err := parseInt32Parameter(params["routeid"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigRouteDel(r.Context(), routeid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigRouteGet - Get a L3 route details 
func (c *NetworkApiController) ConfigRouteGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	routeid := params["routeid"]
	result, err := c.service.ConfigRouteGet(r.Context(), routeid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigRoutesGet - Get all routing table routes 
func (c *NetworkApiController) ConfigRoutesGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.ConfigRoutesGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigRuleCreate - Configures an IP rule 
func (c *NetworkApiController) ConfigRuleCreate(w http.ResponseWriter, r *http.Request) { 
	body := &map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigRuleCreate(r.Context(), *body)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigRuleDel - Removes an IP Rule 
func (c *NetworkApiController) ConfigRuleDel(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	ruleid, err := parseInt32Parameter(params["ruleid"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigRuleDel(r.Context(), ruleid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigRuleGet - Get an IP rule details 
func (c *NetworkApiController) ConfigRuleGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	ruleid, err := parseInt32Parameter(params["ruleid"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigRuleGet(r.Context(), ruleid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigRulesGet - Get all ip rules list 
func (c *NetworkApiController) ConfigRulesGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.ConfigRulesGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigVRFCreate - Configures an new VRF 
func (c *NetworkApiController) ConfigVRFCreate(w http.ResponseWriter, r *http.Request) { 
	body := &map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigVRFCreate(r.Context(), *body)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigVRFDel - Removes a VRF 
func (c *NetworkApiController) ConfigVRFDel(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	vrfid, err := parseInt32Parameter(params["vrfid"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigVRFDel(r.Context(), vrfid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigVRFGet - Get a VRF 
func (c *NetworkApiController) ConfigVRFGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	vrfid, err := parseInt32Parameter(params["vrfid"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ConfigVRFGet(r.Context(), vrfid)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ConfigVRFsGet - Get the list all VRFs 
func (c *NetworkApiController) ConfigVRFsGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.ConfigVRFsGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err, &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
