/*
API Manager Upstream API

API Manager Upstream API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_upstream

import (
	"encoding/json"
)

// checks if the UpstreamTlsContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpstreamTlsContext{}

// UpstreamTlsContext struct for UpstreamTlsContext
type UpstreamTlsContext struct {
	SecretGroupId *string `json:"secretGroupId,omitempty"`
	TlsContextId *string `json:"tlsContextId,omitempty"`
	Audit *Audit `json:"audit,omitempty"`
}

// NewUpstreamTlsContext instantiates a new UpstreamTlsContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpstreamTlsContext() *UpstreamTlsContext {
	this := UpstreamTlsContext{}
	return &this
}

// NewUpstreamTlsContextWithDefaults instantiates a new UpstreamTlsContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpstreamTlsContextWithDefaults() *UpstreamTlsContext {
	this := UpstreamTlsContext{}
	return &this
}

// GetSecretGroupId returns the SecretGroupId field value if set, zero value otherwise.
func (o *UpstreamTlsContext) GetSecretGroupId() string {
	if o == nil || IsNil(o.SecretGroupId) {
		var ret string
		return ret
	}
	return *o.SecretGroupId
}

// GetSecretGroupIdOk returns a tuple with the SecretGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamTlsContext) GetSecretGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecretGroupId) {
		return nil, false
	}
	return o.SecretGroupId, true
}

// HasSecretGroupId returns a boolean if a field has been set.
func (o *UpstreamTlsContext) HasSecretGroupId() bool {
	if o != nil && !IsNil(o.SecretGroupId) {
		return true
	}

	return false
}

// SetSecretGroupId gets a reference to the given string and assigns it to the SecretGroupId field.
func (o *UpstreamTlsContext) SetSecretGroupId(v string) {
	o.SecretGroupId = &v
}

// GetTlsContextId returns the TlsContextId field value if set, zero value otherwise.
func (o *UpstreamTlsContext) GetTlsContextId() string {
	if o == nil || IsNil(o.TlsContextId) {
		var ret string
		return ret
	}
	return *o.TlsContextId
}

// GetTlsContextIdOk returns a tuple with the TlsContextId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamTlsContext) GetTlsContextIdOk() (*string, bool) {
	if o == nil || IsNil(o.TlsContextId) {
		return nil, false
	}
	return o.TlsContextId, true
}

// HasTlsContextId returns a boolean if a field has been set.
func (o *UpstreamTlsContext) HasTlsContextId() bool {
	if o != nil && !IsNil(o.TlsContextId) {
		return true
	}

	return false
}

// SetTlsContextId gets a reference to the given string and assigns it to the TlsContextId field.
func (o *UpstreamTlsContext) SetTlsContextId(v string) {
	o.TlsContextId = &v
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *UpstreamTlsContext) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamTlsContext) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *UpstreamTlsContext) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *UpstreamTlsContext) SetAudit(v Audit) {
	o.Audit = &v
}

func (o UpstreamTlsContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpstreamTlsContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecretGroupId) {
		toSerialize["secretGroupId"] = o.SecretGroupId
	}
	if !IsNil(o.TlsContextId) {
		toSerialize["tlsContextId"] = o.TlsContextId
	}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	return toSerialize, nil
}

type NullableUpstreamTlsContext struct {
	value *UpstreamTlsContext
	isSet bool
}

func (v NullableUpstreamTlsContext) Get() *UpstreamTlsContext {
	return v.value
}

func (v *NullableUpstreamTlsContext) Set(val *UpstreamTlsContext) {
	v.value = val
	v.isSet = true
}

func (v NullableUpstreamTlsContext) IsSet() bool {
	return v.isSet
}

func (v *NullableUpstreamTlsContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpstreamTlsContext(val *UpstreamTlsContext) *NullableUpstreamTlsContext {
	return &NullableUpstreamTlsContext{value: val, isSet: true}
}

func (v NullableUpstreamTlsContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpstreamTlsContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

