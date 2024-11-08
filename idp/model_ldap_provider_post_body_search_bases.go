/*
Identity Provider Management API

Description of Identity Provider API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// checks if the LdapProviderPostBodySearchBases type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapProviderPostBodySearchBases{}

// LdapProviderPostBodySearchBases struct for LdapProviderPostBodySearchBases
type LdapProviderPostBodySearchBases struct {
	Group string `json:"group"`
	User string `json:"user"`
}

// NewLdapProviderPostBodySearchBases instantiates a new LdapProviderPostBodySearchBases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProviderPostBodySearchBases(group string, user string) *LdapProviderPostBodySearchBases {
	this := LdapProviderPostBodySearchBases{}
	this.Group = group
	this.User = user
	return &this
}

// NewLdapProviderPostBodySearchBasesWithDefaults instantiates a new LdapProviderPostBodySearchBases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderPostBodySearchBasesWithDefaults() *LdapProviderPostBodySearchBases {
	this := LdapProviderPostBodySearchBases{}
	return &this
}

// GetGroup returns the Group field value
func (o *LdapProviderPostBodySearchBases) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *LdapProviderPostBodySearchBases) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *LdapProviderPostBodySearchBases) SetGroup(v string) {
	o.Group = v
}

// GetUser returns the User field value
func (o *LdapProviderPostBodySearchBases) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *LdapProviderPostBodySearchBases) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *LdapProviderPostBodySearchBases) SetUser(v string) {
	o.User = v
}

func (o LdapProviderPostBodySearchBases) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapProviderPostBodySearchBases) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableLdapProviderPostBodySearchBases struct {
	value *LdapProviderPostBodySearchBases
	isSet bool
}

func (v NullableLdapProviderPostBodySearchBases) Get() *LdapProviderPostBodySearchBases {
	return v.value
}

func (v *NullableLdapProviderPostBodySearchBases) Set(val *LdapProviderPostBodySearchBases) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProviderPostBodySearchBases) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProviderPostBodySearchBases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProviderPostBodySearchBases(val *LdapProviderPostBodySearchBases) *NullableLdapProviderPostBodySearchBases {
	return &NullableLdapProviderPostBodySearchBases{value: val, isSet: true}
}

func (v NullableLdapProviderPostBodySearchBases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProviderPostBodySearchBases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

