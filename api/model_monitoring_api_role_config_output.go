/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MonitoringAPIRoleConfigOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringAPIRoleConfigOutput{}

// MonitoringAPIRoleConfigOutput struct for MonitoringAPIRoleConfigOutput
type MonitoringAPIRoleConfigOutput struct {
	Config MonitoringRoleConfig `json:"config"`
}

// NewMonitoringAPIRoleConfigOutput instantiates a new MonitoringAPIRoleConfigOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAPIRoleConfigOutput(config MonitoringRoleConfig) *MonitoringAPIRoleConfigOutput {
	this := MonitoringAPIRoleConfigOutput{}
	this.Config = config
	return &this
}

// NewMonitoringAPIRoleConfigOutputWithDefaults instantiates a new MonitoringAPIRoleConfigOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAPIRoleConfigOutputWithDefaults() *MonitoringAPIRoleConfigOutput {
	this := MonitoringAPIRoleConfigOutput{}
	return &this
}

// GetConfig returns the Config field value
func (o *MonitoringAPIRoleConfigOutput) GetConfig() MonitoringRoleConfig {
	if o == nil {
		var ret MonitoringRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *MonitoringAPIRoleConfigOutput) GetConfigOk() (*MonitoringRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *MonitoringAPIRoleConfigOutput) SetConfig(v MonitoringRoleConfig) {
	o.Config = v
}

func (o MonitoringAPIRoleConfigOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringAPIRoleConfigOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

type NullableMonitoringAPIRoleConfigOutput struct {
	value *MonitoringAPIRoleConfigOutput
	isSet bool
}

func (v NullableMonitoringAPIRoleConfigOutput) Get() *MonitoringAPIRoleConfigOutput {
	return v.value
}

func (v *NullableMonitoringAPIRoleConfigOutput) Set(val *MonitoringAPIRoleConfigOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAPIRoleConfigOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAPIRoleConfigOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAPIRoleConfigOutput(val *MonitoringAPIRoleConfigOutput) *NullableMonitoringAPIRoleConfigOutput {
	return &NullableMonitoringAPIRoleConfigOutput{value: val, isSet: true}
}

func (v NullableMonitoringAPIRoleConfigOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAPIRoleConfigOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
