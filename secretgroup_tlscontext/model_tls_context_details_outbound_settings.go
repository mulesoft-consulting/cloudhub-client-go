/*
Secret Group TLS Context API

Secret Group TLS Context API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_tlscontext

import (
	"encoding/json"
)

// checks if the TlsContextDetailsOutboundSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TlsContextDetailsOutboundSettings{}

// TlsContextDetailsOutboundSettings Properties that are applicable only when the TLS context is used to secure outbound traffic.
type TlsContextDetailsOutboundSettings struct {
	// flag that indicates whether the server certificate validation must be skipped.
	SkipServerCertValidation *bool `json:"skipServerCertValidation,omitempty"`
}

// NewTlsContextDetailsOutboundSettings instantiates a new TlsContextDetailsOutboundSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsContextDetailsOutboundSettings() *TlsContextDetailsOutboundSettings {
	this := TlsContextDetailsOutboundSettings{}
	return &this
}

// NewTlsContextDetailsOutboundSettingsWithDefaults instantiates a new TlsContextDetailsOutboundSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsContextDetailsOutboundSettingsWithDefaults() *TlsContextDetailsOutboundSettings {
	this := TlsContextDetailsOutboundSettings{}
	return &this
}

// GetSkipServerCertValidation returns the SkipServerCertValidation field value if set, zero value otherwise.
func (o *TlsContextDetailsOutboundSettings) GetSkipServerCertValidation() bool {
	if o == nil || IsNil(o.SkipServerCertValidation) {
		var ret bool
		return ret
	}
	return *o.SkipServerCertValidation
}

// GetSkipServerCertValidationOk returns a tuple with the SkipServerCertValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetailsOutboundSettings) GetSkipServerCertValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipServerCertValidation) {
		return nil, false
	}
	return o.SkipServerCertValidation, true
}

// HasSkipServerCertValidation returns a boolean if a field has been set.
func (o *TlsContextDetailsOutboundSettings) HasSkipServerCertValidation() bool {
	if o != nil && !IsNil(o.SkipServerCertValidation) {
		return true
	}

	return false
}

// SetSkipServerCertValidation gets a reference to the given bool and assigns it to the SkipServerCertValidation field.
func (o *TlsContextDetailsOutboundSettings) SetSkipServerCertValidation(v bool) {
	o.SkipServerCertValidation = &v
}

func (o TlsContextDetailsOutboundSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TlsContextDetailsOutboundSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkipServerCertValidation) {
		toSerialize["skipServerCertValidation"] = o.SkipServerCertValidation
	}
	return toSerialize, nil
}

type NullableTlsContextDetailsOutboundSettings struct {
	value *TlsContextDetailsOutboundSettings
	isSet bool
}

func (v NullableTlsContextDetailsOutboundSettings) Get() *TlsContextDetailsOutboundSettings {
	return v.value
}

func (v *NullableTlsContextDetailsOutboundSettings) Set(val *TlsContextDetailsOutboundSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTlsContextDetailsOutboundSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTlsContextDetailsOutboundSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTlsContextDetailsOutboundSettings(val *TlsContextDetailsOutboundSettings) *NullableTlsContextDetailsOutboundSettings {
	return &NullableTlsContextDetailsOutboundSettings{value: val, isSet: true}
}

func (v NullableTlsContextDetailsOutboundSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTlsContextDetailsOutboundSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

