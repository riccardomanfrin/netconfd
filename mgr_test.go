package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.lan.athonet.com/riccardo.manfrin/netconfd/nc"
	oas "gitlab.lan.athonet.com/riccardo.manfrin/netconfd/server/go"
)

func parseSampleConfig(sampleConfig string) oas.Config {
	var config oas.Config
	json.Unmarshal([]byte(sampleConfig), &config)
	return config
}

func genSampleConfig() oas.Config {
	sampleConfig := `{
		"global": {},
		"host_network": {
		  "links": [
			{
			  "ifname": "bond0",
			  "link_type": "ether",
			  "linkinfo": {
				"info_kind": "bond",
				"info_data": {
				  "mode": "active-backup"
				}
			  }
			},
			{
			  "ifname": "dummy0",
			  "link_type": "ether",
			  "linkinfo": {
				"info_kind": "dummy",
				"info_slave_data": {
				  "state": "BACKUP"
				}
			  },
			  "master": "bond0"
			},
			{
			  "ifname": "dummy1",
			  "link_type": "ether",
			  "linkinfo": {
				"info_kind": "dummy",
				"info_slave_data": {
				  "state": "ACTIVE"
				}
			  },
			  "master": "bond0"
			}
		  ],
		  "routes": []
		}
	  }`
	return parseSampleConfig(sampleConfig)
}

func newConfigSetReq(config oas.Config) *http.Request {
	reqbody, _ := json.Marshal(config)
	iobody := bytes.NewReader(reqbody)
	req, _ := http.NewRequest("PUT", "/api/1/config", iobody)
	req.Header.Add("Content-Type", "application/json")
	return req
}
func newConfigGetReq() *http.Request {
	req, _ := http.NewRequest("GET", "/api/1/config", nil)
	req.Header.Add("Content-Type", "application/json")
	return req
}

var m *Manager = NewManager()

func checkResponse(t *testing.T, rr *httptest.ResponseRecorder, httpStatusCode int, ncErrorCode nc.ErrorCode, ncreason string) {
	if status := rr.Code; status != httpStatusCode {
		t.Errorf("HTTP Status code mismatch: got %v want %v",
			status,
			http.StatusBadRequest)
	}
	var genericError nc.GenericError
	err := json.Unmarshal(rr.Body.Bytes(), &genericError)
	if err != nil {
		t.Errorf("Err Unmarshal failure")
	}
	if ncErrorCode != nc.RESERVED {
		if genericError.Code != ncErrorCode {
			t.Errorf("Err Code mismatch: got %v, want %v",
				genericError.Code,
				nc.SEMANTIC)
		}
	}
	if ncreason != "" {
		if genericError.Reason != ncreason {
			t.Errorf("Err Reason mismatch: got %v, want %v",
				genericError.Reason,
				ncreason)
		}
	}
}

func runConfigSet(config oas.Config) *httptest.ResponseRecorder {
	req := newConfigSetReq(config)
	rr := httptest.NewRecorder()
	m.ServeHTTP(rr, req)
	return rr
}

func runConfigGet() oas.Config {
	req := newConfigGetReq()
	rr := httptest.NewRecorder()
	m.ServeHTTP(rr, req)
	return parseSampleConfig(string(rr.Body.Bytes()))
}

/* Tests are divided by OK and EC where
 * - OK are checks on a correct action
 * - EC are checks on faulty behavior/requests (edge cases)
 */

//Test001 - EC-001 Active-Backup Bond Without ActiveSlave
func Test001(t *testing.T) {
	c := genSampleConfig()
	*(*c.HostNetwork.Links)[2].Linkinfo.InfoSlaveData.State = "BACKUP"
	rr := runConfigSet(c)
	checkResponse(t, rr, http.StatusBadRequest, nc.SEMANTIC, "Active Slave Iface not found for Active-Backup type bond bond0")
}

//Test002 - EC-002 Active-Backup Bond With Multiple Active Slaves
func Test002(t *testing.T) {
	c := genSampleConfig()
	*(*c.HostNetwork.Links)[1].Linkinfo.InfoSlaveData.State = "ACTIVE"
	rr := runConfigSet(c)
	checkResponse(t, rr, http.StatusBadRequest, nc.SEMANTIC, "Multiple Active Slave Ifaces found for Active-Backup type bond bond0")
}

//Test003 - EC-003 Non Active-Backup Bond With Backup Slave
func Test003(t *testing.T) {
	c := genSampleConfig()
	*(*c.HostNetwork.Links)[0].Linkinfo.InfoData.Mode = "balance-rr"
	rr := runConfigSet(c)
	checkResponse(t, rr, http.StatusBadRequest, nc.SEMANTIC, "Backup Slave Iface dummy0 found for non Active-Backup type bond bond0")
}

//Test004 - OK-005 Bond params check
func Test004(t *testing.T) {
	cset := genSampleConfig()
	rr := runConfigSet(cset)
	checkResponse(t, rr, http.StatusOK, nc.RESERVED, "")
	cget := runConfigGet()
	if cget != cget {
		t.Errorf("Figa")
	}
}
