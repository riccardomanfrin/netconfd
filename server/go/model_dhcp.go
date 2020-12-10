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

// Dhcp DHCP link context to enable. When an object of this kind is specified, the DHCP protocol daemon is enabled on the  defined interface if it exists.  
type Dhcp struct {
	// Interface name 
	Ifname string `json:"ifname"`
	// Required to patch config and remove  DHCP control of a link interface  (defaulted to true if not present) 
	Enabled *bool `json:"enabled,omitempty"`
}

// NewDhcp instantiates a new Dhcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcp(ifname string, ) *Dhcp {
	this := Dhcp{}
	this.Ifname = ifname
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewDhcpWithDefaults instantiates a new Dhcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpWithDefaults() *Dhcp {
	this := Dhcp{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetIfname returns the Ifname field value
func (o *Dhcp) GetIfname() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Ifname
}

// GetIfnameOk returns a tuple with the Ifname field value
// and a boolean to check if the value has been set.
func (o *Dhcp) GetIfnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ifname, true
}

// SetIfname sets field value
func (o *Dhcp) SetIfname(v string) {
	o.Ifname = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Dhcp) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dhcp) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Dhcp) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Dhcp) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o Dhcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ifname"] = o.Ifname
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableDhcp struct {
	value *Dhcp
	isSet bool
}

func (v NullableDhcp) Get() *Dhcp {
	return v.value
}

func (v *NullableDhcp) Set(val *Dhcp) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcp) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcp(val *Dhcp) *NullableDhcp {
	return &NullableDhcp{value: val, isSet: true}
}

func (v NullableDhcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


