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

// LinkLinkinfoInfoSlaveData Info about slave state/config
type LinkLinkinfoInfoSlaveData struct {
	// State of the link:   * `ACTIVE` - Link is actively used   * `BACKUP` - Link is used for failover 
	State *string `json:"state,omitempty"`
	// MII Status:   * `UP`    * `DOWN` 
	MiiStatus *string `json:"mii_status,omitempty"`
	// Number of link failures 
	LinkFailureCount *int32 `json:"link_failure_count,omitempty"`
	// MAC L2 interface HW address
	PermHwaddr *string `json:"perm_hwaddr,omitempty"`
	// Queue Identifier 
	QueueId *int32 `json:"queue_id,omitempty"`
}

// NewLinkLinkinfoInfoSlaveData instantiates a new LinkLinkinfoInfoSlaveData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkLinkinfoInfoSlaveData() *LinkLinkinfoInfoSlaveData {
	this := LinkLinkinfoInfoSlaveData{}
	return &this
}

// NewLinkLinkinfoInfoSlaveDataWithDefaults instantiates a new LinkLinkinfoInfoSlaveData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkLinkinfoInfoSlaveDataWithDefaults() *LinkLinkinfoInfoSlaveData {
	this := LinkLinkinfoInfoSlaveData{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoSlaveData) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoSlaveData) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoSlaveData) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LinkLinkinfoInfoSlaveData) SetState(v string) {
	o.State = &v
}

// GetMiiStatus returns the MiiStatus field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoSlaveData) GetMiiStatus() string {
	if o == nil || o.MiiStatus == nil {
		var ret string
		return ret
	}
	return *o.MiiStatus
}

// GetMiiStatusOk returns a tuple with the MiiStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoSlaveData) GetMiiStatusOk() (*string, bool) {
	if o == nil || o.MiiStatus == nil {
		return nil, false
	}
	return o.MiiStatus, true
}

// HasMiiStatus returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoSlaveData) HasMiiStatus() bool {
	if o != nil && o.MiiStatus != nil {
		return true
	}

	return false
}

// SetMiiStatus gets a reference to the given string and assigns it to the MiiStatus field.
func (o *LinkLinkinfoInfoSlaveData) SetMiiStatus(v string) {
	o.MiiStatus = &v
}

// GetLinkFailureCount returns the LinkFailureCount field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoSlaveData) GetLinkFailureCount() int32 {
	if o == nil || o.LinkFailureCount == nil {
		var ret int32
		return ret
	}
	return *o.LinkFailureCount
}

// GetLinkFailureCountOk returns a tuple with the LinkFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoSlaveData) GetLinkFailureCountOk() (*int32, bool) {
	if o == nil || o.LinkFailureCount == nil {
		return nil, false
	}
	return o.LinkFailureCount, true
}

// HasLinkFailureCount returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoSlaveData) HasLinkFailureCount() bool {
	if o != nil && o.LinkFailureCount != nil {
		return true
	}

	return false
}

// SetLinkFailureCount gets a reference to the given int32 and assigns it to the LinkFailureCount field.
func (o *LinkLinkinfoInfoSlaveData) SetLinkFailureCount(v int32) {
	o.LinkFailureCount = &v
}

// GetPermHwaddr returns the PermHwaddr field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoSlaveData) GetPermHwaddr() string {
	if o == nil || o.PermHwaddr == nil {
		var ret string
		return ret
	}
	return *o.PermHwaddr
}

// GetPermHwaddrOk returns a tuple with the PermHwaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoSlaveData) GetPermHwaddrOk() (*string, bool) {
	if o == nil || o.PermHwaddr == nil {
		return nil, false
	}
	return o.PermHwaddr, true
}

// HasPermHwaddr returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoSlaveData) HasPermHwaddr() bool {
	if o != nil && o.PermHwaddr != nil {
		return true
	}

	return false
}

// SetPermHwaddr gets a reference to the given string and assigns it to the PermHwaddr field.
func (o *LinkLinkinfoInfoSlaveData) SetPermHwaddr(v string) {
	o.PermHwaddr = &v
}

// GetQueueId returns the QueueId field value if set, zero value otherwise.
func (o *LinkLinkinfoInfoSlaveData) GetQueueId() int32 {
	if o == nil || o.QueueId == nil {
		var ret int32
		return ret
	}
	return *o.QueueId
}

// GetQueueIdOk returns a tuple with the QueueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkLinkinfoInfoSlaveData) GetQueueIdOk() (*int32, bool) {
	if o == nil || o.QueueId == nil {
		return nil, false
	}
	return o.QueueId, true
}

// HasQueueId returns a boolean if a field has been set.
func (o *LinkLinkinfoInfoSlaveData) HasQueueId() bool {
	if o != nil && o.QueueId != nil {
		return true
	}

	return false
}

// SetQueueId gets a reference to the given int32 and assigns it to the QueueId field.
func (o *LinkLinkinfoInfoSlaveData) SetQueueId(v int32) {
	o.QueueId = &v
}

func (o LinkLinkinfoInfoSlaveData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.MiiStatus != nil {
		toSerialize["mii_status"] = o.MiiStatus
	}
	if o.LinkFailureCount != nil {
		toSerialize["link_failure_count"] = o.LinkFailureCount
	}
	if o.PermHwaddr != nil {
		toSerialize["perm_hwaddr"] = o.PermHwaddr
	}
	if o.QueueId != nil {
		toSerialize["queue_id"] = o.QueueId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkLinkinfoInfoSlaveData struct {
	value *LinkLinkinfoInfoSlaveData
	isSet bool
}

func (v NullableLinkLinkinfoInfoSlaveData) Get() *LinkLinkinfoInfoSlaveData {
	return v.value
}

func (v *NullableLinkLinkinfoInfoSlaveData) Set(val *LinkLinkinfoInfoSlaveData) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkLinkinfoInfoSlaveData) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkLinkinfoInfoSlaveData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkLinkinfoInfoSlaveData(val *LinkLinkinfoInfoSlaveData) *NullableLinkLinkinfoInfoSlaveData {
	return &NullableLinkLinkinfoInfoSlaveData{value: val, isSet: true}
}

func (v NullableLinkLinkinfoInfoSlaveData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkLinkinfoInfoSlaveData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

