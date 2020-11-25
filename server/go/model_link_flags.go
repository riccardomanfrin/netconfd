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

// LinkFlags the model 'LinkFlags'
type LinkFlags string

// List of link_flags
const (
	BROADCAST LinkFlags = "BROADCAST"
	MULTICAST LinkFlags = "MULTICAST"
	SLAVE LinkFlags = "SLAVE"
	UP LinkFlags = "UP"
	LOWER_UP LinkFlags = "LOWER_UP"
)

func (v *LinkFlags) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkFlags(value)
	for _, existing := range []LinkFlags{ "BROADCAST", "MULTICAST", "SLAVE", "UP", "LOWER_UP",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkFlags", value)
}

// Ptr returns reference to link_flags value
func (v LinkFlags) Ptr() *LinkFlags {
	return &v
}

type NullableLinkFlags struct {
	value *LinkFlags
	isSet bool
}

func (v NullableLinkFlags) Get() *LinkFlags {
	return v.value
}

func (v *NullableLinkFlags) Set(val *LinkFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkFlags(val *LinkFlags) *NullableLinkFlags {
	return &NullableLinkFlags{value: val, isSet: true}
}

func (v NullableLinkFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
