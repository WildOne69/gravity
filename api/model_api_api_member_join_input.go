/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApiAPIMemberJoinInput struct for ApiAPIMemberJoinInput
type ApiAPIMemberJoinInput struct {
	Peer *string `json:"peer,omitempty"`
}

// NewApiAPIMemberJoinInput instantiates a new ApiAPIMemberJoinInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIMemberJoinInput() *ApiAPIMemberJoinInput {
	this := ApiAPIMemberJoinInput{}
	return &this
}

// NewApiAPIMemberJoinInputWithDefaults instantiates a new ApiAPIMemberJoinInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIMemberJoinInputWithDefaults() *ApiAPIMemberJoinInput {
	this := ApiAPIMemberJoinInput{}
	return &this
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *ApiAPIMemberJoinInput) GetPeer() string {
	if o == nil || o.Peer == nil {
		var ret string
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIMemberJoinInput) GetPeerOk() (*string, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *ApiAPIMemberJoinInput) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given string and assigns it to the Peer field.
func (o *ApiAPIMemberJoinInput) SetPeer(v string) {
	o.Peer = &v
}

func (o ApiAPIMemberJoinInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Peer != nil {
		toSerialize["peer"] = o.Peer
	}
	return json.Marshal(toSerialize)
}

type NullableApiAPIMemberJoinInput struct {
	value *ApiAPIMemberJoinInput
	isSet bool
}

func (v NullableApiAPIMemberJoinInput) Get() *ApiAPIMemberJoinInput {
	return v.value
}

func (v *NullableApiAPIMemberJoinInput) Set(val *ApiAPIMemberJoinInput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIMemberJoinInput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIMemberJoinInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIMemberJoinInput(val *ApiAPIMemberJoinInput) *NullableApiAPIMemberJoinInput {
	return &NullableApiAPIMemberJoinInput{value: val, isSet: true}
}

func (v NullableApiAPIMemberJoinInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIMemberJoinInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}