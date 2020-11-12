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

// LinkLinkinfo Additional link info attributes 
type LinkLinkinfo struct {
	// Type of link layer interface. Supported Types:   * `device`- Physical device   * `dummy` - Dummy link type interface for binding intenal services   * `bridge` - Link layer virtual switch type interface   * `bond` - Bond type interface letting two interfaces be seen as one   * `vlan` - Virtual LAN (TAG ID based) interface   * `veth` - Virtual ethernet (with virtual MAC and IP address)   * `macvlan` - Direct virtual eth interface connected to the physical interface,      with owned mac address   * `ipvlan` - Direct virtual eth interface connected to the physical interface.     Physical interface MAC address is reused. L2 type directly connects the lan to      the host phyisical device. L3 type adds a routing layer in between. 
	InfoKind *string `json:"info_kind,omitempty"`
	InfoData *LinkLinkinfoInfoData `json:"info_data,omitempty"`
}

// NewLinkLinkinfo instantiates a new LinkLinkinfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkLinkinfo() *LinkLinkinfo {
	this := LinkLinkinfo{}
	return &this
}

// NewLinkLinkinfoWithDefaults instantiates a new LinkLinkinfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkLinkinfoWithDefaults() *LinkLinkinfo {
	this := LinkLinkinfo{}
	return &this
}

// GetInfoKind returns the InfoKind field value if set, zero value otherwise.
func (o *LinkLinkinfo) GetInfoKind() string {
	if o == nil || o.InfoKind == nil {
		var ret string
		return ret
	}
	return *o.InfoKind
}

// GetInfoKindOk returns a tuple with the InfoKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfo) GetInfoKindOk() (*string, bool) {
	if o == nil || o.InfoKind == nil {
		return nil, false
	}
	return o.InfoKind, true
}

// HasInfoKind returns a boolean if a field has been set.
func (o *LinkLinkinfo) HasInfoKind() bool {
	if o != nil && o.InfoKind != nil {
		return true
	}

	return false
}

// SetInfoKind gets a reference to the given string and assigns it to the InfoKind field.
func (o *LinkLinkinfo) SetInfoKind(v string) {
	o.InfoKind = &v
}

// GetInfoData returns the InfoData field value if set, zero value otherwise.
func (o *LinkLinkinfo) GetInfoData() LinkLinkinfoInfoData {
	if o == nil || o.InfoData == nil {
		var ret LinkLinkinfoInfoData
		return ret
	}
	return *o.InfoData
}

// GetInfoDataOk returns a tuple with the InfoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfo) GetInfoDataOk() (*LinkLinkinfoInfoData, bool) {
	if o == nil || o.InfoData == nil {
		return nil, false
	}
	return o.InfoData, true
}

// HasInfoData returns a boolean if a field has been set.
func (o *LinkLinkinfo) HasInfoData() bool {
	if o != nil && o.InfoData != nil {
		return true
	}

	return false
}

// SetInfoData gets a reference to the given LinkLinkinfoInfoData and assigns it to the InfoData field.
func (o *LinkLinkinfo) SetInfoData(v LinkLinkinfoInfoData) {
	o.InfoData = &v
}

func (o LinkLinkinfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InfoKind != nil {
		toSerialize["info_kind"] = o.InfoKind
	}
	if o.InfoData != nil {
		toSerialize["info_data"] = o.InfoData
	}
	return json.Marshal(toSerialize)
}

type NullableLinkLinkinfo struct {
	value *LinkLinkinfo
	isSet bool
}

func (v NullableLinkLinkinfo) Get() *LinkLinkinfo {
	return v.value
}

func (v *NullableLinkLinkinfo) Set(val *LinkLinkinfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkLinkinfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkLinkinfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkLinkinfo(val *LinkLinkinfo) *NullableLinkLinkinfo {
	return &NullableLinkLinkinfo{value: val, isSet: true}
}

func (v NullableLinkLinkinfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkLinkinfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


