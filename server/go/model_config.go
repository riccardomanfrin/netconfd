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

// Config struct for Config
type Config struct {
	Global  *Global  `json:"global,omitempty"`
	Network *Network `json:"network,omitempty"`
}

// NewConfig instantiates a new Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfig() *Config {
	this := Config{}
	return &this
}

// NewConfigWithDefaults instantiates a new Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigWithDefaults() *Config {
	this := Config{}
	return &this
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *Config) GetGlobal() Global {
	if o == nil || o.Global == nil {
		var ret Global
		return ret
	}
	return *o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Config) GetGlobalOk() (*Global, bool) {
	if o == nil || o.Global == nil {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *Config) HasGlobal() bool {
	if o != nil && o.Global != nil {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given Global and assigns it to the Global field.
func (o *Config) SetGlobal(v Global) {
	o.Global = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *Config) GetNetwork() Network {
	if o == nil || o.Network == nil {
		var ret Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Config) GetNetworkOk() (*Network, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *Config) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given Network and assigns it to the Network field.
func (o *Config) SetNetwork(v Network) {
	o.Network = &v
}

func (o Config) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Global != nil {
		toSerialize["global"] = o.Global
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableConfig struct {
	value *Config
	isSet bool
}

func (v NullableConfig) Get() *Config {
	return v.value
}

func (v *NullableConfig) Set(val *Config) {
	v.value = val
	v.isSet = true
}

func (v NullableConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfig(val *Config) *NullableConfig {
	return &NullableConfig{value: val, isSet: true}
}

func (v NullableConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
