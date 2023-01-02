/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApiAPIToolTracerouteOutputHop struct for ApiAPIToolTracerouteOutputHop
type ApiAPIToolTracerouteOutputHop struct {
	Address     *string `json:"address,omitempty"`
	ElapsedTime *int32  `json:"elapsedTime,omitempty"`
	Success     *bool   `json:"success,omitempty"`
}

// NewApiAPIToolTracerouteOutputHop instantiates a new ApiAPIToolTracerouteOutputHop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIToolTracerouteOutputHop() *ApiAPIToolTracerouteOutputHop {
	this := ApiAPIToolTracerouteOutputHop{}
	return &this
}

// NewApiAPIToolTracerouteOutputHopWithDefaults instantiates a new ApiAPIToolTracerouteOutputHop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIToolTracerouteOutputHopWithDefaults() *ApiAPIToolTracerouteOutputHop {
	this := ApiAPIToolTracerouteOutputHop{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ApiAPIToolTracerouteOutputHop) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolTracerouteOutputHop) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ApiAPIToolTracerouteOutputHop) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ApiAPIToolTracerouteOutputHop) SetAddress(v string) {
	o.Address = &v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *ApiAPIToolTracerouteOutputHop) GetElapsedTime() int32 {
	if o == nil || o.ElapsedTime == nil {
		var ret int32
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolTracerouteOutputHop) GetElapsedTimeOk() (*int32, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *ApiAPIToolTracerouteOutputHop) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int32 and assigns it to the ElapsedTime field.
func (o *ApiAPIToolTracerouteOutputHop) SetElapsedTime(v int32) {
	o.ElapsedTime = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiAPIToolTracerouteOutputHop) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolTracerouteOutputHop) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiAPIToolTracerouteOutputHop) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiAPIToolTracerouteOutputHop) SetSuccess(v bool) {
	o.Success = &v
}

func (o ApiAPIToolTracerouteOutputHop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.ElapsedTime != nil {
		toSerialize["elapsedTime"] = o.ElapsedTime
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableApiAPIToolTracerouteOutputHop struct {
	value *ApiAPIToolTracerouteOutputHop
	isSet bool
}

func (v NullableApiAPIToolTracerouteOutputHop) Get() *ApiAPIToolTracerouteOutputHop {
	return v.value
}

func (v *NullableApiAPIToolTracerouteOutputHop) Set(val *ApiAPIToolTracerouteOutputHop) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIToolTracerouteOutputHop) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIToolTracerouteOutputHop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIToolTracerouteOutputHop(val *ApiAPIToolTracerouteOutputHop) *NullableApiAPIToolTracerouteOutputHop {
	return &NullableApiAPIToolTracerouteOutputHop{value: val, isSet: true}
}

func (v NullableApiAPIToolTracerouteOutputHop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIToolTracerouteOutputHop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
