/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MonitoringRoleConfig struct for MonitoringRoleConfig
type MonitoringRoleConfig struct {
	Port *int32 `json:"port,omitempty"`
}

// NewMonitoringRoleConfig instantiates a new MonitoringRoleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringRoleConfig() *MonitoringRoleConfig {
	this := MonitoringRoleConfig{}
	return &this
}

// NewMonitoringRoleConfigWithDefaults instantiates a new MonitoringRoleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringRoleConfigWithDefaults() *MonitoringRoleConfig {
	this := MonitoringRoleConfig{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *MonitoringRoleConfig) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringRoleConfig) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *MonitoringRoleConfig) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *MonitoringRoleConfig) SetPort(v int32) {
	o.Port = &v
}

func (o MonitoringRoleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringRoleConfig struct {
	value *MonitoringRoleConfig
	isSet bool
}

func (v NullableMonitoringRoleConfig) Get() *MonitoringRoleConfig {
	return v.value
}

func (v *NullableMonitoringRoleConfig) Set(val *MonitoringRoleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringRoleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringRoleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringRoleConfig(val *MonitoringRoleConfig) *NullableMonitoringRoleConfig {
	return &NullableMonitoringRoleConfig{value: val, isSet: true}
}

func (v NullableMonitoringRoleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringRoleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
