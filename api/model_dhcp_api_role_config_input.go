/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DhcpAPIRoleConfigInput struct for DhcpAPIRoleConfigInput
type DhcpAPIRoleConfigInput struct {
	Config DhcpRoleConfig `json:"config"`
}

// NewDhcpAPIRoleConfigInput instantiates a new DhcpAPIRoleConfigInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpAPIRoleConfigInput(config DhcpRoleConfig) *DhcpAPIRoleConfigInput {
	this := DhcpAPIRoleConfigInput{}
	this.Config = config
	return &this
}

// NewDhcpAPIRoleConfigInputWithDefaults instantiates a new DhcpAPIRoleConfigInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpAPIRoleConfigInputWithDefaults() *DhcpAPIRoleConfigInput {
	this := DhcpAPIRoleConfigInput{}
	return &this
}

// GetConfig returns the Config field value
func (o *DhcpAPIRoleConfigInput) GetConfig() DhcpRoleConfig {
	if o == nil {
		var ret DhcpRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *DhcpAPIRoleConfigInput) GetConfigOk() (*DhcpRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *DhcpAPIRoleConfigInput) SetConfig(v DhcpRoleConfig) {
	o.Config = v
}

func (o DhcpAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpAPIRoleConfigInput struct {
	value *DhcpAPIRoleConfigInput
	isSet bool
}

func (v NullableDhcpAPIRoleConfigInput) Get() *DhcpAPIRoleConfigInput {
	return v.value
}

func (v *NullableDhcpAPIRoleConfigInput) Set(val *DhcpAPIRoleConfigInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpAPIRoleConfigInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpAPIRoleConfigInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpAPIRoleConfigInput(val *DhcpAPIRoleConfigInput) *NullableDhcpAPIRoleConfigInput {
	return &NullableDhcpAPIRoleConfigInput{value: val, isSet: true}
}

func (v NullableDhcpAPIRoleConfigInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpAPIRoleConfigInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
