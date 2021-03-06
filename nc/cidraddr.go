// Copyright (c) 2021, Athonet S.r.l. All rights reserved.
// riccardo.manfrin@athonet.com

package nc

import (
	"net"
	"strconv"
)

// CIDRAddr is an address and a network mask (According to RFC 4632 and RFC 4291).
// Additionally to a net.IPNet, it allows for specifying
// further than the netmask bits. Those are intended to define
// an addresses within the IP network being defined along with.
// E.g. : 10.1.2.3/24 -> 10.1.2.3 in network 10.1.2.0/24
type CIDRAddr struct {
	ip   net.IP
	mask int
}

//NewCIDRAddr creates new CIDR address. If network is unspecified it is assumed to be /32 for ipv4 and /128 for ipv6
func NewCIDRAddr(addr string) CIDRAddr {
	a := CIDRAddr{}
	a.ParseCIDRNetStr(addr)
	return a
}

func CIDRAddrLoad(ip string, len int) (CIDRAddr, error) {
	a := CIDRAddr{}
	a.ParseIP(ip)
	err := a.SetPrefixLen(len)
	return a, err
}

//IsValid returns true if address is set and valid
func (a *CIDRAddr) IsValid() bool {
	return a.ip != nil && len(a.ip) > 0
}

//ParseIP parses the IP address
func (a *CIDRAddr) ParseIP(ip string) {
	a.ip = net.ParseIP(ip)
}

//SetIP parses the IP address
func (a *CIDRAddr) SetIP(ip net.IP) {
	a.ip = ip
	if ip == nil {
		a.mask = 0
	} else if a.IsV4() {
		a.mask = 32
	} else {
		a.mask = 128
	}
}

//SetNet parses the IP address
func (a *CIDRAddr) SetNet(ipnet net.IPNet) {
	a.ip = ipnet.IP
	ones, _ := ipnet.Mask.Size()
	a.mask = ones
}

//SetPrefixLen translates an IP network prefix length into a CIDRAddr mask
func (a *CIDRAddr) SetPrefixLen(len int) error {
	if a.IsV4() {
		if len > 32 || len < 0 {
			return NewInvalidPrefixLenForIPv4AddrError(len)
		}
	} else {
		if len > 128 || len < 0 {
			return NewInvalidPrefixLenForIPv6AddrError(len)
		}
	}
	a.mask = len
	return nil
}

//ParseIPNet translates an IP network into a CIDRAddr
func (a *CIDRAddr) ParseIPNet(ip net.IPNet) {
	a.ParseCIDRNetStr(ip.String())
}

//ToIPNet returns an IP network (the non network part is zeroed out)
func (a *CIDRAddr) ToIPNet() net.IPNet {
	if a.IsV4() {
		ipMask := net.CIDRMask(a.mask, 32)
		return net.IPNet{IP: a.ip, Mask: ipMask}
	}
	ipMask := net.CIDRMask(a.mask, 128)
	return net.IPNet{IP: a.ip, Mask: ipMask}
}

//ParseCIDRNetStr loads a CIDR network from a string. If network is unspecified it is assumed to be /32 for ipv4 and /128 for ipv6
func (a *CIDRAddr) ParseCIDRNetStr(straddr string) error {
	var e error
	var ipnet *net.IPNet

	if straddr == "default" {
		straddr = "0.0.0.0/0"
	}

	a.ip, ipnet, e = net.ParseCIDR(straddr)
	if e != nil {
		a.ip = net.ParseIP(straddr)
		if a.ip == nil {
			return NewInvalidIPAddressError(straddr)
		}
		if a.IsV4() {
			a.mask = 32
		} else {
			a.mask = 128
		}

	} else {
		a.ip = ipnet.IP
		a.mask, _ = ipnet.Mask.Size()
	}
	return nil
}

//IsV4 tells if the address is V4
func (a *CIDRAddr) IsV4() bool {
	return a.ip.To4() != nil
}

func (a *CIDRAddr) String() string {

	if !a.IsValid() {
		return ""
	}
	if a.mask == 32 && a.IsV4() {
		return a.ip.String()
	} else if a.mask == 128 && !a.IsV4() {
		return a.ip.String()
	} else {
		return a.ip.String() + "/" + strconv.Itoa(a.mask)
	}
}

//Netmask returns the netmask (e.g. 255.255.255.0) of a CIDR address/network
func (a *CIDRAddr) Netmask() string {
	if a.IsV4() {
		return net.CIDRMask(a.mask, 32).String()
	}
	return net.CIDRMask(a.mask, 128).String()
}

//Address returns the address (e.g. 255.255.255.0) of a CIDR address/network
func (a *CIDRAddr) Address() string {
	return a.ip.String()
}

//PrefixLen returns the length of the network prefix
func (a *CIDRAddr) PrefixLen() int {
	return a.mask
}

//CIDRAddrValidate validates a string as being or not a CIDR addr
func CIDRAddrValidate(cidraddr string) error {
	var a CIDRAddr
	return a.ParseCIDRNetStr(cidraddr)
}

//UnmarshalJSON implements unmarshalling
func (a *CIDRAddr) UnmarshalJSON(data []byte) error {
	unquotedStr, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	return a.ParseCIDRNetStr(unquotedStr)
}

//MarshalJSON JSON Marshaller
func (a CIDRAddr) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(a.String())), nil
}
