package nc

import (
	"github.com/vishvananda/netlink"
	"gitlab.lan.athonet.com/riccardo.manfrin/netconfd/logger"
)

// LinkAddrInfo struct for LinkAddrInfo
type LinkAddrInfo struct {
	Local CIDRAddr `json:"local,omitempty"`
	//Prefixlen int32    `json:"prefixlen,omitempty"`
	//Broadcast CIDRAddr `json:"broadcast,omitempty"`
}

// LinkLinkinfoInfoSlaveData Info about slave state/config
type LinkLinkinfoInfoSlaveData struct {
	// State of the link:   * `ACTIVE` - Link is actively used   * `BACKUP` - Link is used for failover
	State string `json:"state,omitempty"`
	// MII Status:   * `UP`    * `DOWN`
	MiiStatus string `json:"mii_status,omitempty"`
	// Number of link failures
	LinkFailureCount int32 `json:"link_failure_count,omitempty"`
	// Hardware address
	PermHwaddr string `json:"perm_hwaddr,omitempty"`
	// Queue Identifier
	QueueId int16 `json:"queue_id,omitempty"`
}

//LinkLinkinfo definition
type LinkLinkinfo struct {
	// Type of link layer interface. Supported Types:   * `dummy` - Dummy link type interface for binding intenal services   * `bridge` - Link layer virtual switch type interface   * `bond` - Bond type interface letting two interfaces be seen as one   * `vlan` - Virtual LAN (TAG ID based) interface   * `veth` - Virtual ethernet (with virtual MAC and IP address)   * `macvlan` - Direct virtual eth interface connected to the physical interface,      with owned mac address   * `ipvlan` - Direct virtual eth interface connected to the physical interface.     Physical interface MAC address is reused. L2 type directly connects the lan to      the host phyisical device. L3 type adds a routing layer in between.
	InfoKind string `json:"info_kind,omitempty"`
	// FILL ME
	InfoSlaveKind string                    `json:"info_slave_kind,omitempty"`
	InfoSlaveData LinkLinkinfoInfoSlaveData `json:"info_slave_data,omitempty"`
}

//LinkID ifname copy identifier
type LinkID string

//Link definition
type Link struct {
	LinkID LinkID
	// Inteface index ID
	Ifindex int32 `json:"ifindex,omitempty"`
	// Interface name
	Ifname string `json:"ifname"`
	// Maximum Transfer Unit value
	Mtu int32 `json:"mtu,omitempty"`
	// In case the interface is part of a bond or bridge, specifies the bond/bridge interface it belongs to.
	Master   string         `json:"master,omitempty"`
	Linkinfo LinkLinkinfo   `json:"linkinfo,omitempty"`
	LinkType string         `json:"link_type"`
	Address  string         `json:"address,omitempty"`
	AddrInfo []LinkAddrInfo `json:"addr_info,omitempty"`
}

func linkParse(link netlink.Link) Link {
	nclink := Link{}
	la := link.Attrs()
	nclink.LinkID = LinkID(la.Name)
	nclink.Ifname = la.Name
	nclink.Mtu = int32(la.MTU)
	nclink.Linkinfo.InfoKind = link.Type()
	nclink.LinkType = la.EncapType
	addrs, err := netlink.AddrList(link, netlink.FAMILY_ALL)
	if la.Slave != nil {
		nclink.Linkinfo.InfoSlaveKind = la.Slave.SlaveType()
		switch la.Slave.(type) {
		case *netlink.BondSlave:
			{
				bondslave := la.Slave.(*netlink.BondSlave)
				nclink.Linkinfo.InfoSlaveData.State = bondslave.State.String()
				nclink.Linkinfo.InfoSlaveData.MiiStatus = bondslave.MiiStatus.String()
				nclink.Linkinfo.InfoSlaveData.PermHwaddr = bondslave.PermHardwareAddr.String()
				nclink.Linkinfo.InfoSlaveData.QueueId = bondslave.QueueId
				nclink.Linkinfo.InfoSlaveData.LinkFailureCount = bondslave.LinkFailureCount
			}
		default:
			{
				logger.Log.Warning("Unsupported type of slave/master type interface")
			}
		}
	}

	if err == nil {
		nclink.AddrInfo = make([]LinkAddrInfo, len(addrs))
		for i, a := range addrs {
			nclink.AddrInfo[i].Local.Parse(a.IPNet.String())
		}
	}
	return nclink
}

//LinksGet Returns the list of existing link layer devices on the machine
func LinksGet() ([]Link, error) {
	links, err := netlink.LinkList()
	if err != nil {
		return nil, err
	}
	nclinks := make([]Link, len(links))
	for i, l := range links {
		nclinks[i] = linkParse(l)
	}
	return nclinks, nil
}

//LinkGet Returns the list of existing link layer devices on the machine
func LinkGet(LinkID LinkID) (Link, error) {
	nclink := Link{}
	link, err := netlink.LinkByName(string(LinkID))
	if err == nil {
		nclink = linkParse(link)
	}
	return nclink, err
}

// LinkCreate creates a link layer interface
// Link types (or kind):
// $> ip link help type
// ...
// TYPE := { vlan | veth | vcan | vxcan | dummy | ifb | macvlan | macvtap |
//	bridge | bond | team | ipoib | ip6tnl | ipip | sit | vxlan |
//	gre | gretap | erspan | ip6gre | ip6gretap | ip6erspan |
//	vti | nlmon | team_slave | bond_slave | ipvlan | geneve |
//	bridge_slave | vrf | macsec }
func LinkCreate(link Link) error {
	var err error = nil
	ifname := link.Ifname
	kind := link.Linkinfo.InfoKind

	l, _ := netlink.LinkByName(ifname)
	if l != nil {
		return NewLinkExistsConflictError(LinkID(ifname))
	}

	switch kind {
	case "dummy":
		{
			err = LinkDummyCreate(ifname)
		}
	case "bond":
		{
			err = LinkBondCreate(ifname)
		}
	case "bridge":
		{
			err = LinkBridgeCreate(ifname)
		}
	default:
		err = NewUnknownLinkKindError(kind)
	}
	if err != nil {
		logger.Log.Warning(err)
	}
	return err
}

//LinkDelete deletes a link layer interface
func LinkDelete(ifname string) error {
	return nil
}

//LinkDummyCreate Creates a new dummy link
func LinkDummyCreate(ifname string) error {
	attrs := netlink.NewLinkAttrs()
	attrs.Name = ifname
	link := &netlink.Dummy{
		LinkAttrs: attrs,
	}
	return netlink.LinkAdd(link)
}

//LinkBondCreate Creates a new bond link
func LinkBondCreate(ifname string) error {

	return nil
}

//LinkBridgeCreate Creates a new dummy link
func LinkBridgeCreate(ifname string) error {

	return nil
}
