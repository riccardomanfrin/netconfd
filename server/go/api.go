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
	"context"
	"net/http"
)



// NetworkApiRouter defines the required methods for binding the api requests to a responses for the NetworkApi
// The NetworkApiRouter implementation should parse necessary information from the http request, 
// pass the data to a NetworkApiServicer to perform the required actions, then write the service results to the http response.
type NetworkApiRouter interface { 
	ConfigGet(http.ResponseWriter, *http.Request)
	ConfigLinkCreate(http.ResponseWriter, *http.Request)
	ConfigLinkDel(http.ResponseWriter, *http.Request)
	ConfigLinkGet(http.ResponseWriter, *http.Request)
	ConfigNetNSCreate(http.ResponseWriter, *http.Request)
	ConfigNetNSDel(http.ResponseWriter, *http.Request)
	ConfigNetNSGet(http.ResponseWriter, *http.Request)
	ConfigRouteCreate(http.ResponseWriter, *http.Request)
	ConfigRouteDel(http.ResponseWriter, *http.Request)
	ConfigRouteGet(http.ResponseWriter, *http.Request)
	ConfigRuleCreate(http.ResponseWriter, *http.Request)
	ConfigRuleDel(http.ResponseWriter, *http.Request)
	ConfigRuleGet(http.ResponseWriter, *http.Request)
	ConfigSet(http.ResponseWriter, *http.Request)
	ConfigVRFCreate(http.ResponseWriter, *http.Request)
	ConfigVRFDel(http.ResponseWriter, *http.Request)
	ConfigVRFGet(http.ResponseWriter, *http.Request)
}


// NetworkApiServicer defines the api actions for the NetworkApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type NetworkApiServicer interface { 
	ConfigGet(context.Context) (ImplResponse, error)
	ConfigLinkCreate(context.Context, Link) (ImplResponse, error)
	ConfigLinkDel(context.Context, string) (ImplResponse, error)
	ConfigLinkGet(context.Context, string) (ImplResponse, error)
	ConfigNetNSCreate(context.Context, Netns) (ImplResponse, error)
	ConfigNetNSDel(context.Context, string) (ImplResponse, error)
	ConfigNetNSGet(context.Context, string) (ImplResponse, error)
	ConfigRouteCreate(context.Context, Route) (ImplResponse, error)
	ConfigRouteDel(context.Context, string) (ImplResponse, error)
	ConfigRouteGet(context.Context, string) (ImplResponse, error)
	ConfigRuleCreate(context.Context, map[string]interface{}) (ImplResponse, error)
	ConfigRuleDel(context.Context, string) (ImplResponse, error)
	ConfigRuleGet(context.Context, string) (ImplResponse, error)
	ConfigSet(context.Context, Config) (ImplResponse, error)
	ConfigVRFCreate(context.Context, map[string]interface{}) (ImplResponse, error)
	ConfigVRFDel(context.Context, string) (ImplResponse, error)
	ConfigVRFGet(context.Context, string) (ImplResponse, error)
}
