/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DnsAPIRecordsPutInput struct for DnsAPIRecordsPutInput
type DnsAPIRecordsPutInput struct {
	Data         string `json:"data"`
	MxPreference *int32 `json:"mxPreference,omitempty"`
	SrvPort      *int32 `json:"srvPort,omitempty"`
	SrvPriority  *int32 `json:"srvPriority,omitempty"`
	SrvWeight    *int32 `json:"srvWeight,omitempty"`
	Type         string `json:"type"`
}

// NewDnsAPIRecordsPutInput instantiates a new DnsAPIRecordsPutInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsAPIRecordsPutInput(data string, type_ string) *DnsAPIRecordsPutInput {
	this := DnsAPIRecordsPutInput{}
	this.Data = data
	this.Type = type_
	return &this
}

// NewDnsAPIRecordsPutInputWithDefaults instantiates a new DnsAPIRecordsPutInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsAPIRecordsPutInputWithDefaults() *DnsAPIRecordsPutInput {
	this := DnsAPIRecordsPutInput{}
	return &this
}

// GetData returns the Data field value
func (o *DnsAPIRecordsPutInput) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRecordsPutInput) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DnsAPIRecordsPutInput) SetData(v string) {
	o.Data = v
}

// GetMxPreference returns the MxPreference field value if set, zero value otherwise.
func (o *DnsAPIRecordsPutInput) GetMxPreference() int32 {
	if o == nil || o.MxPreference == nil {
		var ret int32
		return ret
	}
	return *o.MxPreference
}

// GetMxPreferenceOk returns a tuple with the MxPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAPIRecordsPutInput) GetMxPreferenceOk() (*int32, bool) {
	if o == nil || o.MxPreference == nil {
		return nil, false
	}
	return o.MxPreference, true
}

// HasMxPreference returns a boolean if a field has been set.
func (o *DnsAPIRecordsPutInput) HasMxPreference() bool {
	if o != nil && o.MxPreference != nil {
		return true
	}

	return false
}

// SetMxPreference gets a reference to the given int32 and assigns it to the MxPreference field.
func (o *DnsAPIRecordsPutInput) SetMxPreference(v int32) {
	o.MxPreference = &v
}

// GetSrvPort returns the SrvPort field value if set, zero value otherwise.
func (o *DnsAPIRecordsPutInput) GetSrvPort() int32 {
	if o == nil || o.SrvPort == nil {
		var ret int32
		return ret
	}
	return *o.SrvPort
}

// GetSrvPortOk returns a tuple with the SrvPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAPIRecordsPutInput) GetSrvPortOk() (*int32, bool) {
	if o == nil || o.SrvPort == nil {
		return nil, false
	}
	return o.SrvPort, true
}

// HasSrvPort returns a boolean if a field has been set.
func (o *DnsAPIRecordsPutInput) HasSrvPort() bool {
	if o != nil && o.SrvPort != nil {
		return true
	}

	return false
}

// SetSrvPort gets a reference to the given int32 and assigns it to the SrvPort field.
func (o *DnsAPIRecordsPutInput) SetSrvPort(v int32) {
	o.SrvPort = &v
}

// GetSrvPriority returns the SrvPriority field value if set, zero value otherwise.
func (o *DnsAPIRecordsPutInput) GetSrvPriority() int32 {
	if o == nil || o.SrvPriority == nil {
		var ret int32
		return ret
	}
	return *o.SrvPriority
}

// GetSrvPriorityOk returns a tuple with the SrvPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAPIRecordsPutInput) GetSrvPriorityOk() (*int32, bool) {
	if o == nil || o.SrvPriority == nil {
		return nil, false
	}
	return o.SrvPriority, true
}

// HasSrvPriority returns a boolean if a field has been set.
func (o *DnsAPIRecordsPutInput) HasSrvPriority() bool {
	if o != nil && o.SrvPriority != nil {
		return true
	}

	return false
}

// SetSrvPriority gets a reference to the given int32 and assigns it to the SrvPriority field.
func (o *DnsAPIRecordsPutInput) SetSrvPriority(v int32) {
	o.SrvPriority = &v
}

// GetSrvWeight returns the SrvWeight field value if set, zero value otherwise.
func (o *DnsAPIRecordsPutInput) GetSrvWeight() int32 {
	if o == nil || o.SrvWeight == nil {
		var ret int32
		return ret
	}
	return *o.SrvWeight
}

// GetSrvWeightOk returns a tuple with the SrvWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAPIRecordsPutInput) GetSrvWeightOk() (*int32, bool) {
	if o == nil || o.SrvWeight == nil {
		return nil, false
	}
	return o.SrvWeight, true
}

// HasSrvWeight returns a boolean if a field has been set.
func (o *DnsAPIRecordsPutInput) HasSrvWeight() bool {
	if o != nil && o.SrvWeight != nil {
		return true
	}

	return false
}

// SetSrvWeight gets a reference to the given int32 and assigns it to the SrvWeight field.
func (o *DnsAPIRecordsPutInput) SetSrvWeight(v int32) {
	o.SrvWeight = &v
}

// GetType returns the Type field value
func (o *DnsAPIRecordsPutInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DnsAPIRecordsPutInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DnsAPIRecordsPutInput) SetType(v string) {
	o.Type = v
}

func (o DnsAPIRecordsPutInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.MxPreference != nil {
		toSerialize["mxPreference"] = o.MxPreference
	}
	if o.SrvPort != nil {
		toSerialize["srvPort"] = o.SrvPort
	}
	if o.SrvPriority != nil {
		toSerialize["srvPriority"] = o.SrvPriority
	}
	if o.SrvWeight != nil {
		toSerialize["srvWeight"] = o.SrvWeight
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDnsAPIRecordsPutInput struct {
	value *DnsAPIRecordsPutInput
	isSet bool
}

func (v NullableDnsAPIRecordsPutInput) Get() *DnsAPIRecordsPutInput {
	return v.value
}

func (v *NullableDnsAPIRecordsPutInput) Set(val *DnsAPIRecordsPutInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsAPIRecordsPutInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsAPIRecordsPutInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsAPIRecordsPutInput(val *DnsAPIRecordsPutInput) *NullableDnsAPIRecordsPutInput {
	return &NullableDnsAPIRecordsPutInput{value: val, isSet: true}
}

func (v NullableDnsAPIRecordsPutInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsAPIRecordsPutInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
