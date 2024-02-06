/*
Secret Group Keystore API

Secret Group Keystore API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_keystore

import (
	"encoding/json"
)

// checks if the CertificateValidity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateValidity{}

// CertificateValidity Details about validity period of this certificate
type CertificateValidity struct {
	NotBefore *string `json:"notBefore,omitempty"`
	NotAfter *string `json:"notAfter,omitempty"`
}

// NewCertificateValidity instantiates a new CertificateValidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateValidity() *CertificateValidity {
	this := CertificateValidity{}
	return &this
}

// NewCertificateValidityWithDefaults instantiates a new CertificateValidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateValidityWithDefaults() *CertificateValidity {
	this := CertificateValidity{}
	return &this
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *CertificateValidity) GetNotBefore() string {
	if o == nil || IsNil(o.NotBefore) {
		var ret string
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateValidity) GetNotBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *CertificateValidity) HasNotBefore() bool {
	if o != nil && !IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given string and assigns it to the NotBefore field.
func (o *CertificateValidity) SetNotBefore(v string) {
	o.NotBefore = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *CertificateValidity) GetNotAfter() string {
	if o == nil || IsNil(o.NotAfter) {
		var ret string
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateValidity) GetNotAfterOk() (*string, bool) {
	if o == nil || IsNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *CertificateValidity) HasNotAfter() bool {
	if o != nil && !IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given string and assigns it to the NotAfter field.
func (o *CertificateValidity) SetNotAfter(v string) {
	o.NotAfter = &v
}

func (o CertificateValidity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateValidity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotBefore) {
		toSerialize["notBefore"] = o.NotBefore
	}
	if !IsNil(o.NotAfter) {
		toSerialize["notAfter"] = o.NotAfter
	}
	return toSerialize, nil
}

type NullableCertificateValidity struct {
	value *CertificateValidity
	isSet bool
}

func (v NullableCertificateValidity) Get() *CertificateValidity {
	return v.value
}

func (v *NullableCertificateValidity) Set(val *CertificateValidity) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateValidity) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateValidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateValidity(val *CertificateValidity) *NullableCertificateValidity {
	return &NullableCertificateValidity{value: val, isSet: true}
}

func (v NullableCertificateValidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateValidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


