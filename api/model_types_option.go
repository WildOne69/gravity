/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TypesOption struct for TypesOption
type TypesOption struct {
	Tag     NullableInt32  `json:"tag,omitempty"`
	TagName *string        `json:"tagName,omitempty"`
	Value   NullableString `json:"value,omitempty"`
	Value64 []string       `json:"value64,omitempty"`
}

// NewTypesOption instantiates a new TypesOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesOption() *TypesOption {
	this := TypesOption{}
	return &this
}

// NewTypesOptionWithDefaults instantiates a new TypesOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesOptionWithDefaults() *TypesOption {
	this := TypesOption{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypesOption) GetTag() int32 {
	if o == nil || o.Tag.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypesOption) GetTagOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *TypesOption) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableInt32 and assigns it to the Tag field.
func (o *TypesOption) SetTag(v int32) {
	o.Tag.Set(&v)
}

// SetTagNil sets the value for Tag to be an explicit nil
func (o *TypesOption) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *TypesOption) UnsetTag() {
	o.Tag.Unset()
}

// GetTagName returns the TagName field value if set, zero value otherwise.
func (o *TypesOption) GetTagName() string {
	if o == nil || o.TagName == nil {
		var ret string
		return ret
	}
	return *o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOption) GetTagNameOk() (*string, bool) {
	if o == nil || o.TagName == nil {
		return nil, false
	}
	return o.TagName, true
}

// HasTagName returns a boolean if a field has been set.
func (o *TypesOption) HasTagName() bool {
	if o != nil && o.TagName != nil {
		return true
	}

	return false
}

// SetTagName gets a reference to the given string and assigns it to the TagName field.
func (o *TypesOption) SetTagName(v string) {
	o.TagName = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypesOption) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypesOption) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *TypesOption) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *TypesOption) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *TypesOption) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *TypesOption) UnsetValue() {
	o.Value.Unset()
}

// GetValue64 returns the Value64 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypesOption) GetValue64() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Value64
}

// GetValue64Ok returns a tuple with the Value64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypesOption) GetValue64Ok() ([]string, bool) {
	if o == nil || o.Value64 == nil {
		return nil, false
	}
	return o.Value64, true
}

// HasValue64 returns a boolean if a field has been set.
func (o *TypesOption) HasValue64() bool {
	if o != nil && o.Value64 != nil {
		return true
	}

	return false
}

// SetValue64 gets a reference to the given []string and assigns it to the Value64 field.
func (o *TypesOption) SetValue64(v []string) {
	o.Value64 = v
}

func (o TypesOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	if o.TagName != nil {
		toSerialize["tagName"] = o.TagName
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Value64 != nil {
		toSerialize["value64"] = o.Value64
	}
	return json.Marshal(toSerialize)
}

type NullableTypesOption struct {
	value *TypesOption
	isSet bool
}

func (v NullableTypesOption) Get() *TypesOption {
	return v.value
}

func (v *NullableTypesOption) Set(val *TypesOption) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesOption) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesOption(val *TypesOption) *NullableTypesOption {
	return &NullableTypesOption{value: val, isSet: true}
}

func (v NullableTypesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
