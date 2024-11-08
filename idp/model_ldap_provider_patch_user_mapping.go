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

// checks if the LdapProviderPatchUserMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapProviderPatchUserMapping{}

// LdapProviderPatchUserMapping struct for LdapProviderPatchUserMapping
type LdapProviderPatchUserMapping struct {
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewLdapProviderPatchUserMapping instantiates a new LdapProviderPatchUserMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProviderPatchUserMapping() *LdapProviderPatchUserMapping {
	this := LdapProviderPatchUserMapping{}
	return &this
}

// NewLdapProviderPatchUserMappingWithDefaults instantiates a new LdapProviderPatchUserMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderPatchUserMappingWithDefaults() *LdapProviderPatchUserMapping {
	this := LdapProviderPatchUserMapping{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LdapProviderPatchUserMapping) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchUserMapping) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LdapProviderPatchUserMapping) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LdapProviderPatchUserMapping) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *LdapProviderPatchUserMapping) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchUserMapping) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *LdapProviderPatchUserMapping) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *LdapProviderPatchUserMapping) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LdapProviderPatchUserMapping) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchUserMapping) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LdapProviderPatchUserMapping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LdapProviderPatchUserMapping) SetId(v string) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *LdapProviderPatchUserMapping) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchUserMapping) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *LdapProviderPatchUserMapping) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *LdapProviderPatchUserMapping) SetLastName(v string) {
	o.LastName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *LdapProviderPatchUserMapping) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchUserMapping) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *LdapProviderPatchUserMapping) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *LdapProviderPatchUserMapping) SetUsername(v string) {
	o.Username = &v
}

func (o LdapProviderPatchUserMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapProviderPatchUserMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableLdapProviderPatchUserMapping struct {
	value *LdapProviderPatchUserMapping
	isSet bool
}

func (v NullableLdapProviderPatchUserMapping) Get() *LdapProviderPatchUserMapping {
	return v.value
}

func (v *NullableLdapProviderPatchUserMapping) Set(val *LdapProviderPatchUserMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProviderPatchUserMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProviderPatchUserMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProviderPatchUserMapping(val *LdapProviderPatchUserMapping) *NullableLdapProviderPatchUserMapping {
	return &NullableLdapProviderPatchUserMapping{value: val, isSet: true}
}

func (v NullableLdapProviderPatchUserMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProviderPatchUserMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


