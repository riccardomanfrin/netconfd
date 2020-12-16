package nc

import (
	"crypto/md5"
	"fmt"
	"net"

	"github.com/riccardomanfrin/netlink"
	"gitlab.lan.athonet.com/core/netconfd/logger"
)

// ModelDefault This is equivalent to 0.0.0.0/0 or ::/0
type ModelDefault string

// Scope scope of the object (link or global)
type Scope string

// List of scope
const (
	LINK   Scope = "link"
	GLOBAL Scope = "global"
)

// Route IP L3 Ruote entry
type Route struct {
	ID      RouteID  `json:"id"`
	Dst     CIDRAddr `json:"dst,omitempty"`
	Gateway net.IP   `json:"gateway,omitempty"`
	// Interface name
	Dev      LinkID `json:"dev,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Metric   int32  `json:"metric,omitempty"`
	Scope    Scope  `json:"scope,omitempty"`
	Prefsrc  net.IP `json:"prefsrc,omitempty"`
	// Route flags
	Flags *[]string `json:"flags,omitempty"`
}

func routeParse(route netlink.Route) (Route, error) {
	ncroute := Route{}
	if route.Dst == nil {
		ncroute.Dst.SetIP(net.IPv4(0, 0, 0, 0))
		ncroute.Dst.SetPrefixLen(0)
	} else {
		ncroute.Dst.SetNet(*route.Dst)
	}
	l, err := netlink.LinkByIndex(route.LinkIndex)
	if err != nil {
		return ncroute, err
	}
	ncroute.Dev = LinkID(l.Attrs().Name)
	ncroute.Gateway = route.Gw
	ncroute.Protocol = route.Protocol.String()
	ncroute.Prefsrc = route.Src
	ncroute.Metric = int32(route.Priority)
	ncroute.Scope = Scope(route.Scope.String())
	ncroute.ID = RouteIDGet(ncroute)
	return ncroute, nil
}

func routeFormat(route Route) (netlink.Route, error) {
	nlroute := netlink.Route{}
	dst := route.Dst.ToIPNet()
	nlroute.Dst = &dst
	nlroute.Gw = route.Gateway
	nlroute.Priority = int(route.Metric)
	if route.Dev.IsValid() {
		l, err := LinkGet(route.Dev)
		if err != nil {
			return nlroute, NewRouteLinkDeviceNotFoundError(route.ID, route.Dev)
		}
		nlroute.LinkIndex = int(l.Ifindex)
	}
	return nlroute, nil
}

//RouteID identifies a route via MD5 of its content
type RouteID string

func RouteIDGet(route Route) RouteID {
	md := md5.New()
	md.Sum([]byte(route.Gateway.String()))
	md.Sum([]byte(route.Dev))
	data := md.Sum([]byte(route.Dst.String()))
	return RouteID(fmt.Sprintf("%x", md5.Sum(data)))
}

//RoutesGet returns the array of routes
func RoutesGet() ([]Route, error) {
	routes, err := netlink.RouteList(nil, netlink.FAMILY_ALL)
	if err != nil {
		return nil, mapNetlinkError(err, nil)
	}
	ncroutes := make([]Route, len(routes))
	for i, r := range routes {
		ncroutes[i], err = routeParse(r)
		if err != nil {
			return ncroutes, err
		}
	}
	return ncroutes, nil
}

//RouteGet Returns the list of existing link layer devices on the machine
func RouteGet(_routeID RouteID) (Route, error) {
	routes, err := RoutesGet()
	if err != nil {
		return Route{}, err
	}
	for _, r := range routes {
		if RouteIDGet(r) == _routeID {
			return r, nil
		}
	}
	return Route{}, NewRouteByIDNotFoundError(_routeID)
}

//RouteDelete deletes a route by ID
func RouteDelete(routeid RouteID) error {
	routes, err := netlink.RouteList(nil, netlink.FAMILY_ALL)
	if err != nil {
		return mapNetlinkError(err, nil)
	}

	for _, r := range routes {
		route, err := routeParse(r)
		if err != nil {
			return mapNetlinkError(err, nil)
		}
		if routeid == RouteIDGet(route) {
			return mapNetlinkError(netlink.RouteDel(&r), nil)
		}
	}
	return NewRouteByIDNotFoundError(routeid)
}

//RoutesConfigure configures the whole set of links to manage in the correct sequential order
//for example some of the link properties require other links to be established already or
//to have the link down/up etc..
//This function tries to wipe out every type of conflicting in place configuration such as
//existing links whose ifname LinkID collides with the ones being created.
func RoutesConfigure(routes []Route) error {
	for _, r := range routes {
		err := RouteDelete(RouteIDGet(r))
		if err != nil {
			if _, ok := err.(*NotFoundError); ok != true {
				logger.Log.Warning(err.Error())
			}
		}
		if _, err := RouteCreate(r); err != nil {
			/* Some routes just cannot be erased
			 * so we accept the fact that they are there already */
			if _, ok := err.(*ConflictError); ok != true {
				return err
			}
		}
	}
	return nil
}

//RoutesDelete deletes all routes
func RoutesDelete() error {
	routes, err := RoutesGet()
	if err != nil {
		return err
	}
	for _, r := range routes {
		err = RouteDelete(RouteIDGet(r))
		if err != nil {
			return err
		}
	}
	return nil
}

//RouteCreate create and add a new route
func RouteCreate(route Route) (RouteID, error) {
	routeid := RouteIDGet(route)
	route.ID = routeid
	routes, err := RoutesGet()
	if err != nil {
		return routeid, err
	}
	for _, r := range routes {
		if r.ID == routeid {
			return routeid, NewRouteExistsConflictError(routeid)
		}
	}

	nlroute, err := routeFormat(route)
	if err != nil {
		return routeid, err
	}

	err = netlink.RouteAdd(&nlroute)
	if err != nil {
		return routeid, mapNetlinkError(err, route)
	}

	return routeid, nil
}
