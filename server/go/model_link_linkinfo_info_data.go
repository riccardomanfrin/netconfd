/*
 * netConfD API
 *
 * Network Configurator service
 *
 * API version: 0.1.0
 * Contact: support@athonet.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LinkLinkinfoInfoData Additional information on the link
type LinkLinkinfoInfoData struct {
	// Bonding modes. Supported Modes:   * `balance-rr` - Round-robin: Transmit network packets in sequential order from the first available network interface (NIC) slave through the last. This mode provides load balancing and fault tolerance.   * `active-backup` - Active-backup: Only one NIC slave in the bond is active. A different slave becomes active if, and only if, the active slave fails. The single logical bonded interface's MAC address is externally visible on only one NIC (port) to avoid distortion in the network switch. This mode provides fault tolerance.   * `balance-xor` - XOR: Transmit network packets based on a hash of the packet's source and destination. The default algorithm only considers MAC addresses (layer2). Newer versions allow selection of additional policies based on IP addresses (layer2+3) and TCP/UDP port numbers (layer3+4). This selects the same NIC slave for each destination MAC address, IP address, or IP address and port combination, respectively. This mode provides load balancing and fault tolerance.   * `broadcast` - Broadcast: Transmit network packets on all slave network interfaces. This mode provides fault tolerance.   * `802.3ad` - IEEE 802.3ad Dynamic link aggregation: Creates aggregation groups that share the same speed and duplex settings. Utilizes all slave network interfaces in the active aggregator group according to the 802.3ad specification. This mode is similar to the XOR mode above and supports the same balancing policies. The link is set up dynamically between two LACP-supporting peers.   * `balance-tlb` - Adaptive transmit load balancing: Linux bonding driver mode that does not require any special network-switch support. The outgoing network packet traffic is distributed according to the current load (computed relative to the speed) on each network interface slave. Incoming traffic is received by one currently designated slave network interface. If this receiving slave fails, another slave takes over the MAC address of the failed receiving slave.   * `balance-alb` - Adaptive load balancing: includes balance-tlb plus receive load balancing (rlb) for IPV4 traffic, and does not require any special network switch support. The receive load balancing is achieved by ARP negotiation. The bonding driver intercepts the ARP Replies sent by the local system on their way out and overwrites the source hardware address with the unique hardware address of one of the NIC slaves in the single logical bonded interface such that different network-peers use different MAC addresses for their network packet traffic. 
	Mode *string `json:"mode,omitempty"`
	// Specifies the MII link monitoring frequency in milliseconds.  The default value is 0, and this will disable the MII monitor 
	Miimon *int32 `json:"miimon,omitempty"`
	// Specifies the time, in milliseconds, to wait before enabling a slave after a  link recovery has been detected. The updelay value must be a multiple of the miimon value 
	Updelay *int32 `json:"updelay,omitempty"`
	// Specifies the time, in milliseconds, to wait before disabling a slave after a  link failure has been detected. The downdelay value must be a multiple of the miimon value. 
	Downdelay *int32 `json:"downdelay,omitempty"`
	// Specify the delay, in milliseconds, between each peer notification (gratuitous ARP and unsolicited IPv6 Neighbor Advertisement) when they are issued after a failover event. This delay should be a multiple of the link monitor interval (arp_interval or miimon, whichever is active). The default value is 0 which means to match the value of the link monitor interval. 
	PeerNotifyDelay *int32 `json:"peer_notify_delay,omitempty"`
	// Specifies whether or not miimon should use MII or ETHTOOL ioctls vs. netif_carrier_ok() to determine the link status. The MII or ETHTOOL ioctls are less efficient and utilize a deprecated calling sequence within the kernel.  The netif_carrier_ok() relies on the device driver to maintain its state with netif_carrier_on/off; at this writing, most, but not all, device drivers support this facility. If bonding insists that the link is up when it should not be, it may be that your network device driver does not support netif_carrier_on/off.  The default state for netif_carrier is \"carrier on,\" so if a driver does not support netif_carrier, it will appear as if the link is always up.  In this case, setting use_carrier to 0 will cause bonding to revert to the MII / ETHTOOL ioctl method to determine the link state. A value of 1 enables the use of netif_carrier_ok(), a value of 0 will use the deprecated MII / ETHTOOL ioctls.  The default value is 1.
	UseCarrier *int32 `json:"use_carrier,omitempty"`
	// Specifies the ARP link monitoring frequency in milliseconds. The ARP monitor works by periodically checking the slave devices to determine whether they have sent or received traffic recently (the precise criteria depends upon the bonding mode, and the state of the slave).  Regular traffic is generated via ARP probes issued for the addresses specified by the arp_ip_target option. This behavior can be modified by the arp_validate option, below. If ARP monitoring is used in an etherchannel compatible mode (modes 0 and 2), the switch should be configured in a mode that evenly distributes packets across all links. If the switch is configured to distribute the packets in an XOR fashion, all replies from the ARP targets will be received on the same link which could cause the other team members to fail.  ARP monitoring should not be used in conjunction with miimon.  A value of 0 disables ARP monitoring.  The default value is 0. 
	ArpInterval *int32 `json:"arp_interval,omitempty"`
	// Specifies whether or not ARP probes and replies should be validated in any mode that supports arp monitoring, or whether non-ARP traffic should be filtered (disregarded) for link monitoring purposes. Possible values are: * `none` - or 0 No validation or filtering is performed. * `active` - or 1 Validation is performed only for the active slave. * `backup` - or 2 Validation is performed only for backup slaves. * `all` - or 3 Validation is performed for all slaves. * `filter` - or 4 Filtering is applied to all slaves. No validation is performed. * `filter_active` - or 5 Filtering is applied to all slaves, validation is performed only for the active slave. * `filter_backup` - or 6 Filtering is applied to all slaves, validation is performed only for backup slaves. 
	ArpValidate *string `json:"arp_validate,omitempty"`
	// Specifies the quantity of arp_ip_targets that must be reachable in order for the ARP monitor to consider a slave as being up. This option affects only active-backup mode for slaves with arp_validation enabled. Possible values are: * `any` - or 0   consider the slave up only when any of the arp_ip_targets   is reachable  * `all` - or 1   consider the slave up only when all of the arp_ip_targets   are reachable 
	ArpAllTargets *string `json:"arp_all_targets,omitempty"`
	// Specifies the reselection policy for the primary slave.  This affects how the primary slave is chosen to become the active slave when failure of the active slave or recovery of the primary slave occurs.  This option is designed to prevent flip-flopping between the primary slave and other slaves.  Possible values are:    * `always` - or 0 (default)     The primary slave becomes the active slave whenever it     comes back up.   * `better` - or 1     The primary slave becomes the active slave when it comes     back up, if the speed and duplex of the primary slave is     better than the speed and duplex of the current active     slave.   * `failure` - or 2     The primary slave becomes the active slave only if the     current active slave fails and the primary slave is up.  The primary_reselect setting is ignored in two cases:    * If no slaves are active, the first slave to recover is     made the active slave.    * When initially enslaved, the primary slave is always made     the active slave.  Changing the primary_reselect policy via sysfs will cause an immediate selection of the best active slave according to the new policy.  This may or may not result in a change of the active slave, depending upon the circumstances. This option was added for bonding version 3.6.0. 
	PrimaryReselect *string `json:"primary_reselect,omitempty"`
	// Specifies whether active-backup mode should set all slaves to the same MAC address at enslavement (the traditional behavior), or, when enabled, perform special handling of the bond's MAC address in accordance with the selected policy. The default policy is none, unless the first slave cannot change its MAC address, in which case the active policy is selected by default. This option may be modified via sysfs only when no slaves are present in the bond. This option was added in bonding version 3.2.0.  The \"follow\" policy was added in bonding version 3.3.0. Possible values are:   * `none` - or 0   This setting disables fail_over_mac, and causes   bonding to set all slaves of an active-backup bond to   the same MAC address at enslavement time.  This is the   default.   * `active` - or 1   The \"active\" fail_over_mac policy indicates that the   MAC address of the bond should always be the MAC   address of the currently active slave.  The MAC   address of the slaves is not changed; instead, the MAC   address of the bond changes during a failover.   This policy is useful for devices that cannot ever   alter their MAC address, or for devices that refuse   incoming broadcasts with their own source MAC (which   interferes with the ARP monitor).   The down side of this policy is that every device on   the network must be updated via gratuitous ARP,   vs. just updating a switch or set of switches (which   often takes place for any traffic, not just ARP   traffic, if the switch snoops incoming traffic to   update its tables) for the traditional method.  If the   gratuitous ARP is lost, communication may be   disrupted.   When this policy is used in conjunction with the mii   monitor, devices which assert link up prior to being   able to actually transmit and receive are particularly   susceptible to loss of the gratuitous ARP, and an   appropriate updelay setting may be required.   * `follow` - or 2   The \"follow\" fail_over_mac policy causes the MAC   address of the bond to be selected normally (normally   the MAC address of the first slave added to the bond).   However, the second and subsequent slaves are not set   to this MAC address while they are in a backup role; a   slave is programmed with the bond's MAC address at   failover time (and the formerly active slave receives   the newly active slave's MAC address).   This policy is useful for multiport devices that   either become confused or incur a performance penalty   when multiple ports are programmed with the same MAC   address. 
	FailOverMac *string `json:"fail_over_mac,omitempty"`
	// Hash policy to route packets on different bond interfaces.  Supported Modes:   * `layer2` - Hash is made on L2 metadata (default)   * `layer2+3` - Hash is made on L2 and L3 metadata   * `layer3+4` - Hash is made on L3 and L4 metadata 
	XmitHashPolicy *string `json:"xmit_hash_policy,omitempty"`
	// Specifies the number of IGMP membership reports to be issued after a failover event. One membership report is issued immediately after the failover, subsequent packets are sent in each 200ms interval.  The valid range is 0 - 255; the default value is 1. A value of 0 prevents the IGMP membership report from being issued in response to the failover event.  This option is useful for bonding modes balance-rr (0), active-backup (1), balance-tlb (5) and balance-alb (6), in which a failover can switch the IGMP traffic from one slave to another.  Therefore a fresh IGMP report must be issued to cause the switch to forward the incoming IGMP traffic over the newly selected slave.  This option was added for bonding version 3.7.0. 
	ResendIgmp *int32 `json:"resend_igmp,omitempty"`
	// Specifies that duplicate frames (received on inactive ports) should be dropped (0) or delivered (1).  Normally, bonding will drop duplicate frames (received on inactive ports), which is desirable for most users. But there are some times it is nice to allow duplicate frames to be delivered.  The default value is 0 (drop duplicate frames received on inactive ports). 
	AllSlavesActive *int32 `json:"all_slaves_active,omitempty"`
	// Specifies the minimum number of links that must be active before asserting carrier. It is similar to the Cisco EtherChannel min-links feature. This allows setting the minimum number of member ports that must be up (link-up state) before marking the bond device as up (carrier on). This is useful for situations where higher level services such as clustering want to ensure a minimum number of low bandwidth links are active before switchover. This option only affect 802.3ad mode.  The default value is 0. This will cause carrier to be asserted (for 802.3ad mode) whenever there is an active aggregator, regardless of the number of available links in that aggregator. Note that, because an aggregator cannot be active without at least one available link, setting this option to 0 or to 1 has the exact same effect. 
	MinLinks *int32 `json:"min_links,omitempty"`
	// Specifies the number of seconds between instances where the bonding driver sends learning packets to each slaves peer switch.  The valid range is 1 - 0x7fffffff; the default value is 1. This Option has effect only in balance-tlb and balance-alb modes. 
	LpInterval *int32 `json:"lp_interval,omitempty"`
	// Specify the number of packets to transmit through a slave before moving to the next one. When set to 0 then a slave is chosen at random.  The valid range is 0 - 65535; the default value is 1. This option has effect only in balance-rr mode. 
	PacketsPerSlave *int32 `json:"packets_per_slave,omitempty"`
	// Rate at which LACP control packets are sent to an LACP-supported interface Supported Modes:   * `slow` - LACP Slow Rate (less bandwidth, default)   * `fast` - LACP Fast Rate (faster fault detection) 
	AdLacpRate *string `json:"ad_lacp_rate,omitempty"`
	// Specifies the 802.3ad aggregation selection logic to use.  The possible values and their effects are:   * `stable` - or 0     The active aggregator is chosen by largest aggregate     bandwidth.     Reselection of the active aggregator occurs only when all     slaves of the active aggregator are down or the active     aggregator has no slaves.     This is the default value.   * `bandwidth` or 1     The active aggregator is chosen by largest aggregate     bandwidth.  Reselection occurs if:     - A slave is added to or removed from the bond     - Any slave's link state changes     - Any slave's 802.3ad association state changes     - The bond's administrative state changes to up   * `count` - or 2     The active aggregator is chosen by the largest number of     ports (slaves).  Reselection occurs as described under the     \"bandwidth\" setting, above.      The bandwidth and count selection policies permit failover of 802.3ad aggregations when partial failure of the active aggregator occurs.  This keeps the aggregator with the highest availability (either in bandwidth or in number of ports) active at all times. This option was added in bonding version 3.4.0. 
	AdSelect *string `json:"ad_select,omitempty"`
	// Specifies if dynamic shuffling of flows is enabled in tlb mode. The value has no effect on any other modes.  The default behavior of tlb mode is to shuffle active flows across slaves based on the load in that interval. This gives nice lb characteristics but can cause packet reordering. If re-ordering is a concern use this variable to disable flow shuffling and rely on load balancing provided solely by the hash distribution. xmit-hash-policy can be used to select the appropriate hashing for the setup.  The sysfs entry can be used to change the setting per bond device and the initial value is derived from the module parameter. The sysfs entry is allowed to be changed only if the bond device is down.  The default value is \"1\" that enables flow shuffling while value \"0\" disables it. This option was added in bonding driver 3.7.1 
	TlbDynamicLb *int32 `json:"tlb_dynamic_lb,omitempty"`
}

// NewLinkLinkinfoInfoData instantiates a new LinkLinkinfoInfoData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkLinkinfoInfoData() *LinkLinkinfoInfoData {
	this := LinkLinkinfoInfoData{}
	return &this
}

// NewLinkLinkinfoInfoDataWithDefaults instantiates a new LinkLinkinfoInfoData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkLinkinfoInfoDataWithDefaults() *LinkLinkinfoInfoData {
	this := LinkLinkinfoInfoData{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoData) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *LinkLinkinfoInfoData) SetMode(v string) {
	o.Mode = &v
}

// GetMiimon returns the Miimon field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetMiimon() int32 {
	if o == nil || o.Miimon == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.Miimon
}

// GetMiimonOk returns a tuple with the Miimon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetMiimonOk() (*int32, bool) {
	if o == nil || o.Miimon == nil {
		return nil, false
	}
	return o.Miimon, true
}

// HasMiimon returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasMiimon() bool {
	if o != nil && o.Miimon != nil {
		return true
	}

	return false
}

// SetMiimon gets a reference to the given int32 and assigns it to the Miimon field.
func (o *LinkLinkinfoInfoData) SetMiimon(v int32) {
	o.Miimon = &v
}

// GetUpdelay returns the Updelay field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetUpdelay() int32 {
	if o == nil || o.Updelay == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.Updelay
}

// GetUpdelayOk returns a tuple with the Updelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetUpdelayOk() (*int32, bool) {
	if o == nil || o.Updelay == nil {
		return nil, false
	}
	return o.Updelay, true
}

// HasUpdelay returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasUpdelay() bool {
	if o != nil && o.Updelay != nil {
		return true
	}

	return false
}

// SetUpdelay gets a reference to the given int32 and assigns it to the Updelay field.
func (o *LinkLinkinfoInfoData) SetUpdelay(v int32) {
	o.Updelay = &v
}

// GetDowndelay returns the Downdelay field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetDowndelay() int32 {
	if o == nil || o.Downdelay == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.Downdelay
}

// GetDowndelayOk returns a tuple with the Downdelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetDowndelayOk() (*int32, bool) {
	if o == nil || o.Downdelay == nil {
		return nil, false
	}
	return o.Downdelay, true
}

// HasDowndelay returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasDowndelay() bool {
	if o != nil && o.Downdelay != nil {
		return true
	}

	return false
}

// SetDowndelay gets a reference to the given int32 and assigns it to the Downdelay field.
func (o *LinkLinkinfoInfoData) SetDowndelay(v int32) {
	o.Downdelay = &v
}

// GetPeerNotifyDelay returns the PeerNotifyDelay field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetPeerNotifyDelay() int32 {
	if o == nil || o.PeerNotifyDelay == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.PeerNotifyDelay
}

// GetPeerNotifyDelayOk returns a tuple with the PeerNotifyDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetPeerNotifyDelayOk() (*int32, bool) {
	if o == nil || o.PeerNotifyDelay == nil {
		return nil, false
	}
	return o.PeerNotifyDelay, true
}

// HasPeerNotifyDelay returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasPeerNotifyDelay() bool {
	if o != nil && o.PeerNotifyDelay != nil {
		return true
	}

	return false
}

// SetPeerNotifyDelay gets a reference to the given int32 and assigns it to the PeerNotifyDelay field.
func (o *LinkLinkinfoInfoData) SetPeerNotifyDelay(v int32) {
	o.PeerNotifyDelay = &v
}

// GetUseCarrier returns the UseCarrier field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetUseCarrier() int32 {
	if o == nil || o.UseCarrier == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.UseCarrier
}

// GetUseCarrierOk returns a tuple with the UseCarrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetUseCarrierOk() (*int32, bool) {
	if o == nil || o.UseCarrier == nil {
		return nil, false
	}
	return o.UseCarrier, true
}

// HasUseCarrier returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasUseCarrier() bool {
	if o != nil && o.UseCarrier != nil {
		return true
	}

	return false
}

// SetUseCarrier gets a reference to the given int32 and assigns it to the UseCarrier field.
func (o *LinkLinkinfoInfoData) SetUseCarrier(v int32) {
	o.UseCarrier = &v
}

// GetArpInterval returns the ArpInterval field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetArpInterval() int32 {
	if o == nil || o.ArpInterval == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.ArpInterval
}

// GetArpIntervalOk returns a tuple with the ArpInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetArpIntervalOk() (*int32, bool) {
	if o == nil || o.ArpInterval == nil {
		return nil, false
	}
	return o.ArpInterval, true
}

// HasArpInterval returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasArpInterval() bool {
	if o != nil && o.ArpInterval != nil {
		return true
	}

	return false
}

// SetArpInterval gets a reference to the given int32 and assigns it to the ArpInterval field.
func (o *LinkLinkinfoInfoData) SetArpInterval(v int32) {
	o.ArpInterval = &v
}

// GetArpValidate returns the ArpValidate field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoData) GetArpValidate() string {
	if o == nil || o.ArpValidate == nil {
		var ret string
		return ret
	}
	return *o.ArpValidate
}

// GetArpValidateOk returns a tuple with the ArpValidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetArpValidateOk() (*string, bool) {
	if o == nil || o.ArpValidate == nil {
		return nil, false
	}
	return o.ArpValidate, true
}

// HasArpValidate returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasArpValidate() bool {
	if o != nil && o.ArpValidate != nil {
		return true
	}

	return false
}

// SetArpValidate gets a reference to the given string and assigns it to the ArpValidate field.
func (o *LinkLinkinfoInfoData) SetArpValidate(v string) {
	o.ArpValidate = &v
}

// GetArpAllTargets returns the ArpAllTargets field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoData) GetArpAllTargets() string {
	if o == nil || o.ArpAllTargets == nil {
		var ret string
		return ret
	}
	return *o.ArpAllTargets
}

// GetArpAllTargetsOk returns a tuple with the ArpAllTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetArpAllTargetsOk() (*string, bool) {
	if o == nil || o.ArpAllTargets == nil {
		return nil, false
	}
	return o.ArpAllTargets, true
}

// HasArpAllTargets returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasArpAllTargets() bool {
	if o != nil && o.ArpAllTargets != nil {
		return true
	}

	return false
}

// SetArpAllTargets gets a reference to the given string and assigns it to the ArpAllTargets field.
func (o *LinkLinkinfoInfoData) SetArpAllTargets(v string) {
	o.ArpAllTargets = &v
}

// GetPrimaryReselect returns the PrimaryReselect field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoData) GetPrimaryReselect() string {
	if o == nil || o.PrimaryReselect == nil {
		var ret string
		return ret
	}
	return *o.PrimaryReselect
}

// GetPrimaryReselectOk returns a tuple with the PrimaryReselect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetPrimaryReselectOk() (*string, bool) {
	if o == nil || o.PrimaryReselect == nil {
		return nil, false
	}
	return o.PrimaryReselect, true
}

// HasPrimaryReselect returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasPrimaryReselect() bool {
	if o != nil && o.PrimaryReselect != nil {
		return true
	}

	return false
}

// SetPrimaryReselect gets a reference to the given string and assigns it to the PrimaryReselect field.
func (o *LinkLinkinfoInfoData) SetPrimaryReselect(v string) {
	o.PrimaryReselect = &v
}

// GetFailOverMac returns the FailOverMac field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoData) GetFailOverMac() string {
	if o == nil || o.FailOverMac == nil {
		var ret string
		return ret
	}
	return *o.FailOverMac
}

// GetFailOverMacOk returns a tuple with the FailOverMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetFailOverMacOk() (*string, bool) {
	if o == nil || o.FailOverMac == nil {
		return nil, false
	}
	return o.FailOverMac, true
}

// HasFailOverMac returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasFailOverMac() bool {
	if o != nil && o.FailOverMac != nil {
		return true
	}

	return false
}

// SetFailOverMac gets a reference to the given string and assigns it to the FailOverMac field.
func (o *LinkLinkinfoInfoData) SetFailOverMac(v string) {
	o.FailOverMac = &v
}

// GetXmitHashPolicy returns the XmitHashPolicy field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoData) GetXmitHashPolicy() string {
	if o == nil || o.XmitHashPolicy == nil {
		var ret string
		return ret
	}
	return *o.XmitHashPolicy
}

// GetXmitHashPolicyOk returns a tuple with the XmitHashPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetXmitHashPolicyOk() (*string, bool) {
	if o == nil || o.XmitHashPolicy == nil {
		return nil, false
	}
	return o.XmitHashPolicy, true
}

// HasXmitHashPolicy returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasXmitHashPolicy() bool {
	if o != nil && o.XmitHashPolicy != nil {
		return true
	}

	return false
}

// SetXmitHashPolicy gets a reference to the given string and assigns it to the XmitHashPolicy field.
func (o *LinkLinkinfoInfoData) SetXmitHashPolicy(v string) {
	o.XmitHashPolicy = &v
}

// GetResendIgmp returns the ResendIgmp field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetResendIgmp() int32 {
	if o == nil || o.ResendIgmp == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.ResendIgmp
}

// GetResendIgmpOk returns a tuple with the ResendIgmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetResendIgmpOk() (*int32, bool) {
	if o == nil || o.ResendIgmp == nil {
		return nil, false
	}
	return o.ResendIgmp, true
}

// HasResendIgmp returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasResendIgmp() bool {
	if o != nil && o.ResendIgmp != nil {
		return true
	}

	return false
}

// SetResendIgmp gets a reference to the given int32 and assigns it to the ResendIgmp field.
func (o *LinkLinkinfoInfoData) SetResendIgmp(v int32) {
	o.ResendIgmp = &v
}

// GetAllSlavesActive returns the AllSlavesActive field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetAllSlavesActive() int32 {
	if o == nil || o.AllSlavesActive == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.AllSlavesActive
}

// GetAllSlavesActiveOk returns a tuple with the AllSlavesActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetAllSlavesActiveOk() (*int32, bool) {
	if o == nil || o.AllSlavesActive == nil {
		return nil, false
	}
	return o.AllSlavesActive, true
}

// HasAllSlavesActive returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasAllSlavesActive() bool {
	if o != nil && o.AllSlavesActive != nil {
		return true
	}

	return false
}

// SetAllSlavesActive gets a reference to the given int32 and assigns it to the AllSlavesActive field.
func (o *LinkLinkinfoInfoData) SetAllSlavesActive(v int32) {
	o.AllSlavesActive = &v
}

// GetMinLinks returns the MinLinks field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetMinLinks() int32 {
	if o == nil || o.MinLinks == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.MinLinks
}

// GetMinLinksOk returns a tuple with the MinLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetMinLinksOk() (*int32, bool) {
	if o == nil || o.MinLinks == nil {
		return nil, false
	}
	return o.MinLinks, true
}

// HasMinLinks returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasMinLinks() bool {
	if o != nil && o.MinLinks != nil {
		return true
	}

	return false
}

// SetMinLinks gets a reference to the given int32 and assigns it to the MinLinks field.
func (o *LinkLinkinfoInfoData) SetMinLinks(v int32) {
	o.MinLinks = &v
}

// GetLpInterval returns the LpInterval field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetLpInterval() int32 {
	if o == nil || o.LpInterval == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.LpInterval
}

// GetLpIntervalOk returns a tuple with the LpInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetLpIntervalOk() (*int32, bool) {
	if o == nil || o.LpInterval == nil {
		return nil, false
	}
	return o.LpInterval, true
}

// HasLpInterval returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasLpInterval() bool {
	if o != nil && o.LpInterval != nil {
		return true
	}

	return false
}

// SetLpInterval gets a reference to the given int32 and assigns it to the LpInterval field.
func (o *LinkLinkinfoInfoData) SetLpInterval(v int32) {
	o.LpInterval = &v
}

// GetPacketsPerSlave returns the PacketsPerSlave field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetPacketsPerSlave() int32 {
	if o == nil || o.PacketsPerSlave == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.PacketsPerSlave
}

// GetPacketsPerSlaveOk returns a tuple with the PacketsPerSlave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetPacketsPerSlaveOk() (*int32, bool) {
	if o == nil || o.PacketsPerSlave == nil {
		return nil, false
	}
	return o.PacketsPerSlave, true
}

// HasPacketsPerSlave returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasPacketsPerSlave() bool {
	if o != nil && o.PacketsPerSlave != nil {
		return true
	}

	return false
}

// SetPacketsPerSlave gets a reference to the given int32 and assigns it to the PacketsPerSlave field.
func (o *LinkLinkinfoInfoData) SetPacketsPerSlave(v int32) {
	o.PacketsPerSlave = &v
}

// GetAdLacpRate returns the AdLacpRate field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoData) GetAdLacpRate() string {
	if o == nil || o.AdLacpRate == nil {
		var ret string
		return ret
	}
	return *o.AdLacpRate
}

// GetAdLacpRateOk returns a tuple with the AdLacpRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetAdLacpRateOk() (*string, bool) {
	if o == nil || o.AdLacpRate == nil {
		return nil, false
	}
	return o.AdLacpRate, true
}

// HasAdLacpRate returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasAdLacpRate() bool {
	if o != nil && o.AdLacpRate != nil {
		return true
	}

	return false
}

// SetAdLacpRate gets a reference to the given string and assigns it to the AdLacpRate field.
func (o *LinkLinkinfoInfoData) SetAdLacpRate(v string) {
	o.AdLacpRate = &v
}

// GetAdSelect returns the AdSelect field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoData) GetAdSelect() string {
	if o == nil || o.AdSelect == nil {
		var ret string
		return ret
	}
	return *o.AdSelect
}

// GetAdSelectOk returns a tuple with the AdSelect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetAdSelectOk() (*string, bool) {
	if o == nil || o.AdSelect == nil {
		return nil, false
	}
	return o.AdSelect, true
}

// HasAdSelect returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasAdSelect() bool {
	if o != nil && o.AdSelect != nil {
		return true
	}

	return false
}

// SetAdSelect gets a reference to the given string and assigns it to the AdSelect field.
func (o *LinkLinkinfoInfoData) SetAdSelect(v string) {
	o.AdSelect = &v
}

// GetTlbDynamicLb returns the TlbDynamicLb field value if set, -1 otherwise.
func (o *LinkLinkinfoInfoData) GetTlbDynamicLb() int32 {
	if o == nil || o.TlbDynamicLb == nil {
		var ret int32 = -1 
		return ret
	}
	return *o.TlbDynamicLb
}

// GetTlbDynamicLbOk returns a tuple with the TlbDynamicLb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoData) GetTlbDynamicLbOk() (*int32, bool) {
	if o == nil || o.TlbDynamicLb == nil {
		return nil, false
	}
	return o.TlbDynamicLb, true
}

// HasTlbDynamicLb returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoData) HasTlbDynamicLb() bool {
	if o != nil && o.TlbDynamicLb != nil {
		return true
	}

	return false
}

// SetTlbDynamicLb gets a reference to the given int32 and assigns it to the TlbDynamicLb field.
func (o *LinkLinkinfoInfoData) SetTlbDynamicLb(v int32) {
	o.TlbDynamicLb = &v
}

func (o LinkLinkinfoInfoData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Miimon != nil {
		toSerialize["miimon"] = o.Miimon
	}
	if o.Updelay != nil {
		toSerialize["updelay"] = o.Updelay
	}
	if o.Downdelay != nil {
		toSerialize["downdelay"] = o.Downdelay
	}
	if o.PeerNotifyDelay != nil {
		toSerialize["peer_notify_delay"] = o.PeerNotifyDelay
	}
	if o.UseCarrier != nil {
		toSerialize["use_carrier"] = o.UseCarrier
	}
	if o.ArpInterval != nil {
		toSerialize["arp_interval"] = o.ArpInterval
	}
	if o.ArpValidate != nil {
		toSerialize["arp_validate"] = o.ArpValidate
	}
	if o.ArpAllTargets != nil {
		toSerialize["arp_all_targets"] = o.ArpAllTargets
	}
	if o.PrimaryReselect != nil {
		toSerialize["primary_reselect"] = o.PrimaryReselect
	}
	if o.FailOverMac != nil {
		toSerialize["fail_over_mac"] = o.FailOverMac
	}
	if o.XmitHashPolicy != nil {
		toSerialize["xmit_hash_policy"] = o.XmitHashPolicy
	}
	if o.ResendIgmp != nil {
		toSerialize["resend_igmp"] = o.ResendIgmp
	}
	if o.AllSlavesActive != nil {
		toSerialize["all_slaves_active"] = o.AllSlavesActive
	}
	if o.MinLinks != nil {
		toSerialize["min_links"] = o.MinLinks
	}
	if o.LpInterval != nil {
		toSerialize["lp_interval"] = o.LpInterval
	}
	if o.PacketsPerSlave != nil {
		toSerialize["packets_per_slave"] = o.PacketsPerSlave
	}
	if o.AdLacpRate != nil {
		toSerialize["ad_lacp_rate"] = o.AdLacpRate
	}
	if o.AdSelect != nil {
		toSerialize["ad_select"] = o.AdSelect
	}
	if o.TlbDynamicLb != nil {
		toSerialize["tlb_dynamic_lb"] = o.TlbDynamicLb
	}
	return json.Marshal(toSerialize)
}

type NullableLinkLinkinfoInfoData struct {
	value *LinkLinkinfoInfoData
	isSet bool
}

func (v NullableLinkLinkinfoInfoData) Get() *LinkLinkinfoInfoData {
	return v.value
}

func (v *NullableLinkLinkinfoInfoData) Set(val *LinkLinkinfoInfoData) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkLinkinfoInfoData) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkLinkinfoInfoData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkLinkinfoInfoData(val *LinkLinkinfoInfoData) *NullableLinkLinkinfoInfoData {
	return &NullableLinkLinkinfoInfoData{value: val, isSet: true}
}

func (v NullableLinkLinkinfoInfoData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkLinkinfoInfoData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


