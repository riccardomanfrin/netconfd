/*
 * netConfD API
 *
 * Network Configurator service
 *
 * API version: 0.3.0
 * Contact: support@athonet.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Dns Name server for DNS resolution
type Dns struct {
	Id Dnsid `json:"__id"`
	// The DNS server ip address to send DNS queries to
	Nameserver string `json:"nameserver"`
}

// NewDns instantiates a new Dns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDns(id Dnsid, nameserver string, ) *Dns {
	this := Dns{}
	this.Id = id
	this.Nameserver = nameserver
	return &this
}

// NewDnsWithDefaults instantiates a new Dns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsWithDefaults() *Dns {
	this := Dns{}
	var id Dnsid = PRIMARY
	this.Id = id
	return &this
}

// GetId returns the Id field value
func (o *Dns) GetId() Dnsid {
	if o == nil  {
		var ret Dnsid
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Dns) GetIdOk() (*Dnsid, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Dns) SetId(v Dnsid) {
	o.Id = v
}

// GetNameserver returns the Nameserver field value
func (o *Dns) GetNameserver() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Nameserver
}

// GetNameserverOk returns a tuple with the Nameserver field value
// and a boolean to check if the value has been set.
func (o *Dns) GetNameserverOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nameserver, true
}

// SetNameserver sets field value
func (o *Dns) SetNameserver(v string) {
	o.Nameserver = v
}

func (o Dns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["__id"] = o.Id
	}
	if true {
		toSerialize["nameserver"] = o.Nameserver
	}
	return json.Marshal(toSerialize)
}

type NullableDns struct {
	value *Dns
	isSet bool
}

func (v NullableDns) Get() *Dns {
	return v.value
}

func (v *NullableDns) Set(val *Dns) {
	v.value = val
	v.isSet = true
}

func (v NullableDns) IsSet() bool {
	return v.isSet
}

func (v *NullableDns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDns(val *Dns) *NullableDns {
	return &NullableDns{value: val, isSet: true}
}

func (v NullableDns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


