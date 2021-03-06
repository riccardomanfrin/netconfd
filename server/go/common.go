package openapi

import (
	"net"

	"github.com/riccardomanfrin/netconfd/logger"
	"github.com/riccardomanfrin/netconfd/nc"
)

func dnssGet() ([]Dns, error) {
	var dnss []Dns
	ncdnss, err := nc.DNSsGet()
	if err == nil {
		dnss = make([]Dns, len(ncdnss))
		for i, l := range ncdnss {
			dnss[i] = ncDnsParse(l)
		}
	}
	return dnss, err
}

func linksGet() ([]Link, error) {
	var links []Link
	nclinks, err := nc.LinksGet()
	if err == nil {
		links = make([]Link, len(nclinks))
		for i, l := range nclinks {
			links[i] = ncLinkParse(l)
		}
	}
	return links, err
}

func dhcpsGet() ([]Dhcp, error) {
	var dhcps []Dhcp
	ncdhcps, err := nc.DHCPsGet()
	if err == nil {
		dhcps = make([]Dhcp, len(ncdhcps))
		for i, d := range ncdhcps {
			dhcps[i] = ncDhcpParse(d)
		}
	}
	return dhcps, err
}

func unmanagedListGet() ([]Unmanaged, error) {
	var umgmts []Unmanaged
	ncumgmts, err := nc.UnmanagedListGet()
	if err == nil {
		umgmts = make([]Unmanaged, len(ncumgmts))
		for i, u := range ncumgmts {
			umgmts[i] = ncUnmanagedParse(u)
		}
	}
	return umgmts, err
}

func ncUnmanagedFormat(u Unmanaged) (nc.Unmanaged, error) {
	d := nc.Unmanaged{
		Type: nc.Type(u.GetType()),
		ID:   nc.UnmanagedID(u.GetId()),
	}
	return d, nil
}

func ncUnmanagedParse(ncunmanaged nc.Unmanaged) Unmanaged {
	d := Unmanaged{}
	d.SetId(string(ncunmanaged.ID))
	d.SetType(string(ncunmanaged.Type))
	return d
}

func ncDhcpFormat(dhcp Dhcp) (nc.Dhcp, error) {
	d := nc.Dhcp{
		Ifname: nc.LinkID(dhcp.Ifname),
	}
	return d, nil
}

func ncDhcpParse(ncdhcp nc.Dhcp) Dhcp {
	d := Dhcp{
		Ifname: string(ncdhcp.Ifname),
	}
	return d
}

func ncRouteFormat(route Route) (nc.Route, error) {
	ncroute := nc.Route{}
	ncroute.Dst = route.Dst
	if route.Gateway != nil {
		ncroute.Gateway = *route.Gateway
	}
	if route.Prefsrc != nil {
		ncroute.Prefsrc = *route.Prefsrc
	}
	if route.Dev != nil {
		ncroute.Dev = nc.LinkID(*route.Dev)
	}
	if route.Metric != nil {
		ncroute.Metric = *route.Metric
	}
	if route.Table != nil {
		ncroute.Table = *route.Table
	}
	return ncroute, nil
}

func ncRuleFormat(rule Rule) (nc.Rule, error) {
	//TODO
	ncrule := nc.Rule{
		Priority: rule.GetPriority(),
		Family:   rule.GetFamily(),
		Table:    rule.GetTable(),
		Mark:     rule.GetMark(),
		Mask:     rule.GetMask(),

		TunID: rule.GetTunid(),
		Goto:  rule.GetGoto(),

		Flow:              rule.GetFlow(),
		IifName:           rule.GetIif(),
		OifName:           rule.GetOif(),
		SuppressIfgroup:   rule.GetSuppressIfgroup(),
		SuppressPrefixlen: rule.GetSuppressPrefixlen(),
		Invert:            rule.GetNot(),
	}

	if dport, ok := rule.GetDportOk(); ok {
		ncrule.Dport = &nc.PortRange{
			Start: uint16(*dport),
			End:   uint16(*dport),
		}
	} else {
		if dportstart, ok := rule.GetDportStartOk(); ok {
			if dportend, ok := rule.GetDportEndOk(); ok {
				ncrule.Dport = &nc.PortRange{
					Start: uint16(*dportstart),
					End:   uint16(*dportend),
				}
			}
		}
	}
	if sport, ok := rule.GetSportOk(); ok {
		ncrule.Sport = &nc.PortRange{
			Start: uint16(*sport),
			End:   uint16(*sport),
		}
	} else {
		if sportstart, ok := rule.GetSportStartOk(); ok {
			if sportend, ok := rule.GetSportEndOk(); ok {
				ncrule.Sport = &nc.PortRange{
					Start: uint16(*sportstart),
					End:   uint16(*sportend),
				}
			}
		}
	}
	if src, ok := rule.GetSrcOk(); ok {
		snet, err := nc.CIDRAddrLoad(*src, rule.GetSrclen())
		if err != nil {
			return ncrule, err
		}
		ipnet := snet.ToIPNet()
		ncrule.Src = &ipnet
	}
	if dst, ok := rule.GetDstOk(); ok {
		dnet, err := nc.CIDRAddrLoad(*dst, rule.GetDstlen())
		if err != nil {
			return ncrule, err
		}
		ipnet := dnet.ToIPNet()
		ncrule.Dst = &ipnet
	}
	if tos, ok := rule.GetTosOk(); ok {
		ncrule.Tos = uint(*tos)
	}

	return ncrule, nil
}

func ncRuleParse(ncrule nc.Rule) Rule {
	//TODO

	r := Rule{
		Priority:          &ncrule.Priority,
		Family:            &ncrule.Family,
		Table:             &ncrule.Table,
		Mark:              &ncrule.Mark,
		Mask:              &ncrule.Mask,
		Tunid:             &ncrule.TunID,
		Goto:              &ncrule.Goto,
		Flow:              &ncrule.Flow,
		Iif:               &ncrule.IifName,
		Oif:               &ncrule.OifName,
		SuppressIfgroup:   &ncrule.SuppressIfgroup,
		SuppressPrefixlen: &ncrule.SuppressPrefixlen,
		Not:               &ncrule.Invert,
	}
	if ncrule.Tos != 0 {
		tos := int(ncrule.Tos)
		r.Tos = &tos
	}
	if ncrule.Src != nil {
		s := ncrule.Src.String()
		r.Src = &s
	}
	if ncrule.Dst != nil {
		d := ncrule.Dst.String()
		r.Dst = &d
	}

	if ncrule.Dport != nil {
		if ncrule.Dport.IsSingle() {
			p := int(ncrule.Dport.End)
			r.Dport = &p
		} else {
			s := int(ncrule.Dport.Start)
			e := int(ncrule.Dport.End)
			r.DportStart = &s
			r.DportEnd = &e
		}
	}
	if ncrule.Sport != nil {
		if ncrule.Sport.IsSingle() {
			p := int(ncrule.Sport.End)
			r.Sport = &p
		} else {
			s := int(ncrule.Sport.Start)
			e := int(ncrule.Sport.End)
			r.SportStart = &s
			r.SportEnd = &e
		}
	}

	return r
}

func rulesGet() ([]Rule, error) {
	var rules []Rule
	ncrules, err := nc.RulesGet()
	if err == nil {
		rules = make([]Rule, len(ncrules))
		for i, r := range ncrules {
			rules[i] = ncRuleParse(r)
		}
	}
	return rules, err
}

func ncRouteParse(ncroute nc.Route) Route {
	var route Route
	id := string(ncroute.ID)
	route.Id = &id
	prefsrc := ncroute.Prefsrc.String()
	if prefsrc != "" {
		route.SetPrefsrc(prefsrc)
	}
	dst := ncroute.Dst.String()
	if dst != "" {
		route.SetDst(dst)
	}
	gw := ncroute.Gateway.String()
	if gw != "" {
		route.SetGateway(gw)
	}
	route.SetDev(string(ncroute.Dev))
	route.SetProtocol(ncroute.Protocol)
	route.SetMetric(ncroute.Metric)
	route.SetScope(Scope(ncroute.Scope))
	route.SetTable(ncroute.Table)
	return route
}

func routesGet() ([]Route, error) {
	var routes []Route
	ncroutes, err := nc.RoutesGet()
	if err == nil {
		routes = make([]Route, len(ncroutes))
		for i, r := range ncroutes {
			routes[i] = ncRouteParse(r)
		}
	}
	return routes, err
}

func ncDnsParse(ncdns nc.Dns) Dns {
	ns := ncdns.Nameserver.String()
	prio := Dnsid(ncdns.Id)
	return Dns{
		Nameserver: ns,
		Id:         prio,
	}
}

func ncDnsFormat(dns Dns) (nc.Dns, error) {
	d := nc.Dns{
		Nameserver: net.ParseIP(dns.GetNameserver()),
		Id:         nc.DnsID(dns.GetId()),
	}
	return d, nil
}

func ncLinkParse(nclink nc.Link) Link {
	link := Link{
		Ifname:    string(nclink.Ifname),
		Ifindex:   &nclink.Ifindex,
		LinkType:  nclink.LinkType,
		Operstate: &nclink.Operstate,
	}

	link.Mtu = &nclink.Mtu

	flagsLen := len(nclink.Flags)
	if flagsLen > 0 {
		lfs := make([]LinkFlag, flagsLen)
		link.Flags = &lfs
		for i, lf := range nclink.Flags {
			(*link.Flags)[i] = LinkFlag(lf)
		}
	}

	if len(nclink.AddrInfo) > 0 {
		lai := make([]LinkAddrInfo, len(nclink.AddrInfo))
		for i, a := range nclink.AddrInfo {
			lai[i].Local = a.Local.ToIPNet().IP
			lai[i].Prefixlen = int32(a.Local.PrefixLen())
			if a.Address != nil {
				lai[i].Address = a.Address
			}
		}
		link.AddrInfo = &lai
	}

	lli := DiscriminatedLinkInfo{}
	if nclink.Linkinfo.InfoKind != "" {
		lli.InfoKind = &nclink.Linkinfo.InfoKind
		link.Linkinfo = &lli
	}

	switch nclink.Linkinfo.InfoKind {
	case "vlan":
		{
			did := DiscriminatedInfoData{}
			did.SetProtocol(nclink.Linkinfo.InfoData.Protocol)
			did.SetId(nclink.Linkinfo.InfoData.Id)
			lli.InfoData = &did
			parentLink := string(nclink.Link)
			link.Link = &parentLink
		}
	case "gre":
		{
			did := DiscriminatedInfoData{}
			did.SetLocal(nclink.Linkinfo.InfoData.Local.String())
			did.SetRemote(nclink.Linkinfo.InfoData.Remote.String())
			lli.InfoData = &did
		}
	case "bond":
		{
			did := DiscriminatedInfoData{}
			did.SetMode(nclink.Linkinfo.InfoData.Mode)
			did.SetMiimon(nclink.Linkinfo.InfoData.Miimon)
			did.SetDowndelay(nclink.Linkinfo.InfoData.Downdelay)
			did.SetUpdelay(nclink.Linkinfo.InfoData.Updelay)
			did.SetPeerNotifyDelay(nclink.Linkinfo.InfoData.PeerNotifyDelay)
			did.SetUseCarrier(nclink.Linkinfo.InfoData.UseCarrier)
			did.SetArpInterval(nclink.Linkinfo.InfoData.ArpInterval)
			did.SetArpValidate(nclink.Linkinfo.InfoData.ArpValidate)
			did.SetLpInterval(nclink.Linkinfo.InfoData.LpInterval)
			did.SetArpAllTargets(nclink.Linkinfo.InfoData.ArpAllTargets)
			did.SetPacketsPerSlave(nclink.Linkinfo.InfoData.PacketsPerSlave)
			did.SetFailOverMac(nclink.Linkinfo.InfoData.FailOverMac)
			did.SetXmitHashPolicy(nclink.Linkinfo.InfoData.XmitHashPolicy)
			did.SetResendIgmp(nclink.Linkinfo.InfoData.ResendIgmp)
			did.SetMinLinks(nclink.Linkinfo.InfoData.MinLinks)
			did.SetArpInterval(nclink.Linkinfo.InfoData.ArpInterval)
			did.SetPrimaryReselect(nclink.Linkinfo.InfoData.PrimaryReselect)
			did.SetTlbDynamicLb(nclink.Linkinfo.InfoData.TlbDynamicLb)
			did.SetAdSelect(nclink.Linkinfo.InfoData.AdSelect)
			did.SetAdLacpRate(nclink.Linkinfo.InfoData.AdLacpRate)
			did.SetMode(nclink.Linkinfo.InfoData.Mode)
			did.SetAllSlavesActive(nclink.Linkinfo.InfoData.AllSlavesActive)
			did.SetUseCarrier(nclink.Linkinfo.InfoData.UseCarrier)
			lli.InfoData = &did
		}
	case "device":
	case "bridge":
	case "dummy":
	case "ppp":
	case "tun":
	case "tap":
	case "vrf":
		{
			did := DiscriminatedInfoData{}
			did.SetTable(*nclink.Linkinfo.InfoData.Table)
			lli.InfoData = &did
		}
	default:
		{
			logger.Log.Warning("Unknown Link Kind : %v", nclink.Linkinfo.InfoKind)
		}
	}

	if nclink.Master != "" {
		master := string(nclink.Master)
		if nclink.Linkinfo.InfoSlaveKind != "" {
			isd := LinkLinkinfoInfoSlaveData{}
			lli.InfoSlaveData = &isd
			icisd := &nclink.Linkinfo.InfoSlaveData
			lli.SetInfoSlaveKind(nclink.Linkinfo.InfoSlaveKind)
			switch nclink.Linkinfo.InfoSlaveKind {
			case "bond":
				{
					isd.SetState(icisd.State)
					isd.SetLinkFailureCount(int32(icisd.LinkFailureCount))
					isd.SetMiiStatus(icisd.MiiStatus)
					isd.SetPermHwaddr(icisd.PermHwaddr)
				}
			case "vrf":
				{
					isd.SetTable(*icisd.Table)
				}
			}
		}
		link.SetMaster(master)

	}

	return link
}

func ncLinkFormat(link Link) (nc.Link, error) {

	nclink := nc.Link{
		Ifname:   nc.LinkID(link.GetIfname()),
		Linkinfo: nc.LinkLinkinfo{},
		Mtu:      link.GetMtu(),
		LinkType: link.GetLinkType(),
		Master:   nc.LinkID(link.GetMaster()),
	}

	if link.Link != nil {
		nclink.Link = nc.LinkID(*link.Link)
	}

	if link.Flags != nil {
		flagsLen := len(*link.Flags)
		if flagsLen > 0 {
			nclink.Flags = make([]nc.LinkFlag, flagsLen)
			for i, lf := range *link.Flags {
				nclink.Flags[i] = nc.LinkFlag(lf)
			}
		}
	}

	li := link.GetLinkinfo()

	if li.InfoData != nil {
		nclink.Linkinfo.InfoData = nc.LinkLinkinfoInfoData{
			Mode:            li.InfoData.GetMode(),
			Miimon:          li.InfoData.GetMiimon(),
			Downdelay:       li.InfoData.GetDowndelay(),
			Updelay:         li.InfoData.GetUpdelay(),
			PeerNotifyDelay: li.InfoData.GetPeerNotifyDelay(),
			UseCarrier:      li.InfoData.GetUseCarrier(),
			ArpInterval:     li.InfoData.GetArpInterval(),
			ArpValidate:     li.InfoData.GetArpValidate(),
			LpInterval:      li.InfoData.GetLpInterval(),
			ArpAllTargets:   li.InfoData.GetArpAllTargets(),
			PacketsPerSlave: li.InfoData.GetPacketsPerSlave(),
			FailOverMac:     li.InfoData.GetFailOverMac(),
			XmitHashPolicy:  li.InfoData.GetXmitHashPolicy(),
			ResendIgmp:      li.InfoData.GetResendIgmp(),
			MinLinks:        li.InfoData.GetMinLinks(),
			PrimaryReselect: li.InfoData.GetPrimaryReselect(),
			TlbDynamicLb:    li.InfoData.GetTlbDynamicLb(),
			AdSelect:        li.InfoData.GetAdSelect(),
			AdLacpRate:      li.InfoData.GetAdLacpRate(),
			AllSlavesActive: li.InfoData.GetAllSlavesActive(),
			Protocol:        li.InfoData.GetProtocol(),
			Id:              li.InfoData.GetId(),
			Local:           net.ParseIP(li.InfoData.GetLocal()),
			Remote:          net.ParseIP(li.InfoData.GetRemote()),
			Table:           li.InfoData.Table,
		}

	}
	if li.InfoKind != nil {
		nclink.Linkinfo.InfoKind = *li.InfoKind
	}

	if li.InfoSlaveKind != nil {
		nclink.Linkinfo.InfoSlaveKind = *li.InfoSlaveKind
	}

	if li.InfoSlaveData != nil {
		isd := &nclink.Linkinfo.InfoSlaveData
		isd.State = li.InfoSlaveData.GetState()
		isd.Table = li.InfoSlaveData.Table
	}
	if link.AddrInfo != nil {
		nclink.AddrInfo = make([]nc.LinkAddrInfo, len(*link.AddrInfo))
		for i, addr := range *link.AddrInfo {
			var cidrnet nc.CIDRAddr
			cidrnet.SetIP(addr.Local)
			err := cidrnet.SetPrefixLen(int(addr.Prefixlen))
			if err != nil {
				return nclink, err
			}
			lai := nc.LinkAddrInfo{
				Local:   cidrnet,
				Address: addr.Address,
			}
			nclink.AddrInfo[i] = lai
		}
	}
	return nclink, nil
}

func ncNetFormat(config Config) (nc.Network, error) {
	network := nc.Network{}
	if config.Network != nil {
		if config.Network.Links != nil {
			network.Links = make([]nc.Link, len(*config.Network.Links))
			for i, l := range *config.Network.Links {
				var err error
				network.Links[i], err = ncLinkFormat(l)
				if err != nil {
					return network, err
				}
			}

		}
		if config.Network.Routes != nil {
			network.Routes = make([]nc.Route, len(*config.Network.Routes))
			for i, l := range *config.Network.Routes {
				network.Routes[i], _ = ncRouteFormat(l)
			}
		}
		if config.Network.Dhcp != nil {
			network.Dhcp = make([]nc.Dhcp, len(*config.Network.Dhcp))
			for i, d := range *config.Network.Dhcp {
				network.Dhcp[i], _ = ncDhcpFormat(d)
			}
		}
		if config.Network.Dns != nil {
			network.Dnss = make([]nc.Dns, len(*config.Network.Dns))
			for i, s := range *config.Network.Dns {
				network.Dnss[i], _ = ncDnsFormat(s)
			}
		}
		if config.Network.Unmanaged != nil {
			network.Unmanaged = make([]nc.Unmanaged, len(*config.Network.Unmanaged))
			for i, s := range *config.Network.Unmanaged {
				network.Unmanaged[i], _ = ncUnmanagedFormat(s)
			}
		}
	}
	return network, nil
}

func ncNetParse(net nc.Network) Network {
	links := make([]Link, len(net.Links))
	routes := make([]Route, len(net.Routes))
	dhcps := make([]Dhcp, len(net.Dhcp))
	dnss := make([]Dns, len(net.Dnss))
	rules := make([]Rule, len(net.Rules))
	unmanaged := make([]Unmanaged, len(net.Unmanaged))
	for i, l := range net.Links {
		links[i] = ncLinkParse(l)
	}
	for i, r := range net.Routes {
		routes[i] = ncRouteParse(r)
	}
	for i, d := range net.Dhcp {
		dhcps[i] = ncDhcpParse(d)
	}
	for i, s := range net.Dnss {
		dnss[i] = ncDnsParse(s)
	}
	for i, u := range net.Unmanaged {
		unmanaged[i] = ncUnmanagedParse(u)
	}
	for i, r := range net.Rules {
		rules[i] = ncRuleParse(r)
	}
	return Network{
		Links:     &links,
		Routes:    &routes,
		Dhcp:      &dhcps,
		Dns:       &dnss,
		Unmanaged: &unmanaged,
		Rules:     &rules,
	}
}
