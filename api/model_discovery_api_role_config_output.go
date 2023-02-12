/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DiscoveryAPIRoleConfigOutput struct for DiscoveryAPIRoleConfigOutput
type DiscoveryAPIRoleConfigOutput struct {
	Config DiscoveryRoleConfig `json:"config"`
}

// NewDiscoveryAPIRoleConfigOutput instantiates a new DiscoveryAPIRoleConfigOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryAPIRoleConfigOutput(config DiscoveryRoleConfig) *DiscoveryAPIRoleConfigOutput {
	this := DiscoveryAPIRoleConfigOutput{}
	this.Config = config
	return &this
}

// NewDiscoveryAPIRoleConfigOutputWithDefaults instantiates a new DiscoveryAPIRoleConfigOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryAPIRoleConfigOutputWithDefaults() *DiscoveryAPIRoleConfigOutput {
	this := DiscoveryAPIRoleConfigOutput{}
	return &this
}

// GetConfig returns the Config field value
func (o *DiscoveryAPIRoleConfigOutput) GetConfig() DiscoveryRoleConfig {
	if o == nil {
		var ret DiscoveryRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPIRoleConfigOutput) GetConfigOk() (*DiscoveryRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *DiscoveryAPIRoleConfigOutput) SetConfig(v DiscoveryRoleConfig) {
	o.Config = v
}

func (o DiscoveryAPIRoleConfigOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableDiscoveryAPIRoleConfigOutput struct {
	value *DiscoveryAPIRoleConfigOutput
	isSet bool
}

func (v NullableDiscoveryAPIRoleConfigOutput) Get() *DiscoveryAPIRoleConfigOutput {
	return v.value
}

func (v *NullableDiscoveryAPIRoleConfigOutput) Set(val *DiscoveryAPIRoleConfigOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryAPIRoleConfigOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryAPIRoleConfigOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryAPIRoleConfigOutput(val *DiscoveryAPIRoleConfigOutput) *NullableDiscoveryAPIRoleConfigOutput {
	return &NullableDiscoveryAPIRoleConfigOutput{value: val, isSet: true}
}

func (v NullableDiscoveryAPIRoleConfigOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryAPIRoleConfigOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
