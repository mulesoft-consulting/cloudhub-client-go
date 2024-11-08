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

// checks if the LdapProviderPostBodyUserMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapProviderPostBodyUserMapping{}

// LdapProviderPostBodyUserMapping struct for LdapProviderPostBodyUserMapping
type LdapProviderPostBodyUserMapping struct {
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	Id string `json:"id"`
	LastName string `json:"lastName"`
	Username string `json:"username"`
}

// NewLdapProviderPostBodyUserMapping instantiates a new LdapProviderPostBodyUserMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProviderPostBodyUserMapping(email string, firstName string, id string, lastName string, username string) *LdapProviderPostBodyUserMapping {
	this := LdapProviderPostBodyUserMapping{}
	this.Email = email
	this.FirstName = firstName
	this.Id = id
	this.LastName = lastName
	this.Username = username
	return &this
}

// NewLdapProviderPostBodyUserMappingWithDefaults instantiates a new LdapProviderPostBodyUserMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderPostBodyUserMappingWithDefaults() *LdapProviderPostBodyUserMapping {
	this := LdapProviderPostBodyUserMapping{}
	return &this
}

// GetEmail returns the Email field value
func (o *LdapProviderPostBodyUserMapping) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *LdapProviderPostBodyUserMapping) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *LdapProviderPostBodyUserMapping) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *LdapProviderPostBodyUserMapping) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *LdapProviderPostBodyUserMapping) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *LdapProviderPostBodyUserMapping) SetFirstName(v string) {
	o.FirstName = v
}

// GetId returns the Id field value
func (o *LdapProviderPostBodyUserMapping) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LdapProviderPostBodyUserMapping) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LdapProviderPostBodyUserMapping) SetId(v string) {
	o.Id = v
}

// GetLastName returns the LastName field value
func (o *LdapProviderPostBodyUserMapping) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *LdapProviderPostBodyUserMapping) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *LdapProviderPostBodyUserMapping) SetLastName(v string) {
	o.LastName = v
}

// GetUsername returns the Username field value
func (o *LdapProviderPostBodyUserMapping) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *LdapProviderPostBodyUserMapping) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *LdapProviderPostBodyUserMapping) SetUsername(v string) {
	o.Username = v
}

func (o LdapProviderPostBodyUserMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapProviderPostBodyUserMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["firstName"] = o.FirstName
	toSerialize["id"] = o.Id
	toSerialize["lastName"] = o.LastName
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableLdapProviderPostBodyUserMapping struct {
	value *LdapProviderPostBodyUserMapping
	isSet bool
}

func (v NullableLdapProviderPostBodyUserMapping) Get() *LdapProviderPostBodyUserMapping {
	return v.value
}

func (v *NullableLdapProviderPostBodyUserMapping) Set(val *LdapProviderPostBodyUserMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProviderPostBodyUserMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProviderPostBodyUserMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProviderPostBodyUserMapping(val *LdapProviderPostBodyUserMapping) *NullableLdapProviderPostBodyUserMapping {
	return &NullableLdapProviderPostBodyUserMapping{value: val, isSet: true}
}

func (v NullableLdapProviderPostBodyUserMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProviderPostBodyUserMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

