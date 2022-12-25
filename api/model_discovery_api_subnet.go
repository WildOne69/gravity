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

// DiscoveryAPISubnet struct for DiscoveryAPISubnet
type DiscoveryAPISubnet struct {
	DiscoveryTTL int32  `json:"discoveryTTL"`
	DnsResolver  string `json:"dnsResolver"`
	Name         string `json:"name"`
	SubnetCidr   string `json:"subnetCidr"`
}

// NewDiscoveryAPISubnet instantiates a new DiscoveryAPISubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryAPISubnet(discoveryTTL int32, dnsResolver string, name string, subnetCidr string) *DiscoveryAPISubnet {
	this := DiscoveryAPISubnet{}
	this.DiscoveryTTL = discoveryTTL
	this.DnsResolver = dnsResolver
	this.Name = name
	this.SubnetCidr = subnetCidr
	return &this
}

// NewDiscoveryAPISubnetWithDefaults instantiates a new DiscoveryAPISubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryAPISubnetWithDefaults() *DiscoveryAPISubnet {
	this := DiscoveryAPISubnet{}
	return &this
}

// GetDiscoveryTTL returns the DiscoveryTTL field value
func (o *DiscoveryAPISubnet) GetDiscoveryTTL() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiscoveryTTL
}

// GetDiscoveryTTLOk returns a tuple with the DiscoveryTTL field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPISubnet) GetDiscoveryTTLOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoveryTTL, true
}

// SetDiscoveryTTL sets field value
func (o *DiscoveryAPISubnet) SetDiscoveryTTL(v int32) {
	o.DiscoveryTTL = v
}

// GetDnsResolver returns the DnsResolver field value
func (o *DiscoveryAPISubnet) GetDnsResolver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DnsResolver
}

// GetDnsResolverOk returns a tuple with the DnsResolver field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPISubnet) GetDnsResolverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsResolver, true
}

// SetDnsResolver sets field value
func (o *DiscoveryAPISubnet) SetDnsResolver(v string) {
	o.DnsResolver = v
}

// GetName returns the Name field value
func (o *DiscoveryAPISubnet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPISubnet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DiscoveryAPISubnet) SetName(v string) {
	o.Name = v
}

// GetSubnetCidr returns the SubnetCidr field value
func (o *DiscoveryAPISubnet) GetSubnetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetCidr
}

// GetSubnetCidrOk returns a tuple with the SubnetCidr field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAPISubnet) GetSubnetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetCidr, true
}

// SetSubnetCidr sets field value
func (o *DiscoveryAPISubnet) SetSubnetCidr(v string) {
	o.SubnetCidr = v
}

func (o DiscoveryAPISubnet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["discoveryTTL"] = o.DiscoveryTTL
	}
	if true {
		toSerialize["dnsResolver"] = o.DnsResolver
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["subnetCidr"] = o.SubnetCidr
	}
	return json.Marshal(toSerialize)
}

type NullableDiscoveryAPISubnet struct {
	value *DiscoveryAPISubnet
	isSet bool
}

func (v NullableDiscoveryAPISubnet) Get() *DiscoveryAPISubnet {
	return v.value
}

func (v *NullableDiscoveryAPISubnet) Set(val *DiscoveryAPISubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryAPISubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryAPISubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryAPISubnet(val *DiscoveryAPISubnet) *NullableDiscoveryAPISubnet {
	return &NullableDiscoveryAPISubnet{value: val, isSet: true}
}

func (v NullableDiscoveryAPISubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryAPISubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}