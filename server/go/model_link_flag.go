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
	"fmt"
)

// LinkFlag Flags of the interface  Supported types:   * `broadcast` - Request support for broadcast   * `multicast` - Request support for multicast   * `loopback` - Specify interface as loopback type   * `pointtopoint` - Request support for point-to-point   * `up` - Request link UP state 
type LinkFlag string

// List of link_flag
const (
	BROADCAST LinkFlag = "broadcast"
	MULTICAST LinkFlag = "multicast"
	LOOPBACK LinkFlag = "loopback"
	UP LinkFlag = "up"
	POINTTOPOINT LinkFlag = "pointtopoint"
)

func (v *LinkFlag) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkFlag(value)
	for _, existing := range []LinkFlag{ "broadcast", "multicast", "loopback", "up", "pointtopoint",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkFlag", value)
}

// Ptr returns reference to link_flag value
func (v LinkFlag) Ptr() *LinkFlag {
	return &v
}

type NullableLinkFlag struct {
	value *LinkFlag
	isSet bool
}

func (v NullableLinkFlag) Get() *LinkFlag {
	return v.value
}

func (v *NullableLinkFlag) Set(val *LinkFlag) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkFlag(val *LinkFlag) *NullableLinkFlag {
	return &NullableLinkFlag{value: val, isSet: true}
}

func (v NullableLinkFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

