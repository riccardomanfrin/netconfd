/*
 * netConfD API
 *
 * Network Configurator service
 *
 * API version: 0.2.0
 * Contact: support@athonet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"

	"gitlab.lan.athonet.com/core/netconfd/nc"
)

// NetworkApiService is a service that implents the logic for the NetworkApiServicer
// This service should implement the business logic for every endpoint for the NetworkApi API.
// Include any external packages or services that will be required by this service.
type NetworkApiService struct {
}

// NewNetworkApiService creates a default api service
func NewNetworkApiService() NetworkApiServicer {
	return &NetworkApiService{}
}

// ConfigDHCPCreate - Create DHCP
func (s *NetworkApiService) ConfigDHCPCreate(ctx context.Context, dhcp Dhcp) (ImplResponse, error) {
	ncdhcp, err := ncDhcpFormat(dhcp)
	if err != nil {
		return PostErrorResponse(err, nil)
	}
	err = nc.DHCPCreate(ncdhcp)
	return PostErrorResponse(err, nil)
}

// ConfigDHCPDel - Delete DHCP
func (s *NetworkApiService) ConfigDHCPDel(ctx context.Context, ifname string) (ImplResponse, error) {
	err := nc.DHCPDelete(nc.LinkID(ifname))
	return DeleteErrorResponse(err, nil)
}

// ConfigDHCPGet - Get DHCP
func (s *NetworkApiService) ConfigDHCPGet(ctx context.Context, ifname string) (ImplResponse, error) {
	ncdhcp, err := nc.DHCPGet(nc.LinkID(ifname))
	if err != nil {
		return GetErrorResponse(err, nil)
	}
	dhcp := ncDhcpParse(ncdhcp)
	return GetErrorResponse(err, dhcp)
}

// ConfigDHCPsGet - Get All DHCP
func (s *NetworkApiService) ConfigDHCPsGet(ctx context.Context) (ImplResponse, error) {
	dhcps, err := dhcpsGet()
	return GetErrorResponse(err, dhcps)
}

// ConfigDNSCreate - Create DNS
func (s *NetworkApiService) ConfigDNSCreate(ctx context.Context, dns Dns) (ImplResponse, error) {
	ncdns, err := ncDnsFormat(dns)
	if err != nil {
		return PostErrorResponse(err, nil)
	}
	err = nc.DNSCreate(ncdns)
	return PostErrorResponse(err, nil)
}

// ConfigDNSDel - Delete DNS
func (s *NetworkApiService) ConfigDNSDel(ctx context.Context, dnsid Dnsid) (ImplResponse, error) {
	err := nc.DNSDelete(nc.DnsID(dnsid))
	return DeleteErrorResponse(err, nil)
}

// ConfigDNSGet - Get DNS
func (s *NetworkApiService) ConfigDNSGet(ctx context.Context, dnsid Dnsid) (ImplResponse, error) {
	nclink, err := nc.DnsGet(nc.DnsID(dnsid))
	if err != nil {
		return GetErrorResponse(err, nil)
	}

	return GetErrorResponse(err, ncDnsParse(nclink))
}

// ConfigDNSsGet - Get All DNS config
func (s *NetworkApiService) ConfigDNSsGet(ctx context.Context) (ImplResponse, error) {
	dnss, err := dnssGet()
	return GetErrorResponse(err, dnss)
}

// ConfigLinkCreate - Create New Link
func (s *NetworkApiService) ConfigLinkCreate(ctx context.Context, link Link) (ImplResponse, error) {
	nclink, err := ncLinkFormat(link)
	if err != nil {
		return PostErrorResponse(err, nil)
	}
	if nclink.Master != "" {
		/* You cannot enslave a link if it is UP */
		err = nc.LinkCreateDown(nclink)
		if err != nil {
			return PostErrorResponse(err, nil)
		}

		nc.LinkSetMaster(nclink.Ifname, nclink.Master)

		if nclink.Flags.HaveFlag(nc.LinkFlag("up")) {
			err = nc.LinkSetUp(nclink.Ifname)
			if err != nil {
				return PostErrorResponse(err, nil)
			}
		}
		if nclink.Mtu > 0 {
			err = nc.LinkSetMTU(nclink.Ifname, int(nclink.Mtu))
			if err != nil {
				return PostErrorResponse(err, nil)
			}
		}
	} else {
		err = nc.LinkCreate(nclink)
		if err != nil {
			return PostErrorResponse(err, nil)
		}
		if nclink.Mtu > 0 {
			err = nc.LinkSetMTU(nclink.Ifname, int(nclink.Mtu))
			if err != nil {
				return PostErrorResponse(err, nil)
			}
		}
	}

	return PostErrorResponse(err, nil)
}

// ConfigLinkDel - Brings down and delete a link layer interface
func (s *NetworkApiService) ConfigLinkDel(ctx context.Context, ifname string) (ImplResponse, error) {
	err := nc.LinkDelete(nc.LinkID(ifname))
	return DeleteErrorResponse(err, nil)
}

// ConfigLinkGet - Retrieve link layer interface information
func (s *NetworkApiService) ConfigLinkGet(ctx context.Context, ifname string) (ImplResponse, error) {
	nclink, err := nc.LinkGet(nc.LinkID(ifname))
	if err != nil {
		return GetErrorResponse(err, nil)
	}

	return GetErrorResponse(err, ncLinkParse(nclink))
}

// ConfigLinksGet - Get all link layer interfaces
func (s *NetworkApiService) ConfigLinksGet(ctx context.Context) (ImplResponse, error) {
	links, err := linksGet()

	return GetErrorResponse(err, links)
}

// ConfigRouteCreate - Configures a route
func (s *NetworkApiService) ConfigRouteCreate(ctx context.Context, route Route) (ImplResponse, error) {
	ncroute, err := ncRouteFormat(route)
	if err != nil {
		return PostErrorResponse(err, nil)
	}
	routeid, err := nc.RouteCreate(ncroute)
	if err != nil {
		return PostErrorResponse(err, nil)
	}
	return PostErrorResponse(err, routeid)
}

// ConfigRouteDel - Brings down and delete an L3 IP route
func (s *NetworkApiService) ConfigRouteDel(ctx context.Context, routeid string) (ImplResponse, error) {
	err := nc.RouteDelete(nc.RouteID(routeid))
	return DeleteErrorResponse(err, nil)
}

// ConfigRouteGet - Get a L3 route details
func (s *NetworkApiService) ConfigRouteGet(ctx context.Context, routeid string) (ImplResponse, error) {
	ncroute, err := nc.RouteGet(nc.RouteID(routeid))
	if err != nil {
		return GetErrorResponse(err, nil)
	}
	return GetErrorResponse(err, ncRouteParse(ncroute))
}

// ConfigRoutesGet - Get all routing table routes
func (s *NetworkApiService) ConfigRoutesGet(ctx context.Context) (ImplResponse, error) {
	routes, err := routesGet()
	return GetErrorResponse(err, routes)
}

// ConfigUnmanagedCreate - Create Unmanaged
func (s *NetworkApiService) ConfigUnmanagedCreate(ctx context.Context, unmanaged Unmanaged) (ImplResponse, error) {
	ncunmanaged, err := ncUnmanagedFormat(unmanaged)
	if err != nil {
		return PostErrorResponse(err, nil)
	}
	err = nc.UnmanagedCreate(ncunmanaged)
	return PostErrorResponse(err, nil)
}

// ConfigUnmanagedDel - Delete Unmanaged
func (s *NetworkApiService) ConfigUnmanagedDel(ctx context.Context, id string) (ImplResponse, error) {
	err := nc.UnmanagedDelete(nc.UnmanagedID(id))
	return DeleteErrorResponse(err, nil)
}

// ConfigUnmanagedGet - Get Unmanaged
func (s *NetworkApiService) ConfigUnmanagedGet(ctx context.Context, id string) (ImplResponse, error) {
	ncunmanaged, err := nc.UnmanagedGet(nc.UnmanagedID(id))
	if err != nil {
		return GetErrorResponse(err, nil)
	}

	return GetErrorResponse(err, ncUnmanagedParse(ncunmanaged))
}

// ConfigUnmanagedListGet - Get All Unmanaged
func (s *NetworkApiService) ConfigUnmanagedListGet(ctx context.Context) (ImplResponse, error) {
	links, err := unmanagedListGet()

	return GetErrorResponse(err, links)
}
