/*
 * Connected App API
 *
 * Description of the Connected App API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connected_app

import (
	"encoding/json"
)

// ScopeCore struct for ScopeCore
type ScopeCore struct {
	Scope *string `json:"scope,omitempty"`
	ContextParams *ContextParams `json:"context_params,omitempty"`
}

// NewScopeCore instantiates a new ScopeCore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeCore() *ScopeCore {
	this := ScopeCore{}
	return &this
}

// NewScopeCoreWithDefaults instantiates a new ScopeCore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeCoreWithDefaults() *ScopeCore {
	this := ScopeCore{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ScopeCore) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeCore) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ScopeCore) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ScopeCore) SetScope(v string) {
	o.Scope = &v
}

// GetContextParams returns the ContextParams field value if set, zero value otherwise.
func (o *ScopeCore) GetContextParams() ContextParams {
	if o == nil || o.ContextParams == nil {
		var ret ContextParams
		return ret
	}
	return *o.ContextParams
}

// GetContextParamsOk returns a tuple with the ContextParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeCore) GetContextParamsOk() (*ContextParams, bool) {
	if o == nil || o.ContextParams == nil {
		return nil, false
	}
	return o.ContextParams, true
}

// HasContextParams returns a boolean if a field has been set.
func (o *ScopeCore) HasContextParams() bool {
	if o != nil && o.ContextParams != nil {
		return true
	}

	return false
}

// SetContextParams gets a reference to the given ContextParams and assigns it to the ContextParams field.
func (o *ScopeCore) SetContextParams(v ContextParams) {
	o.ContextParams = &v
}

func (o ScopeCore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.ContextParams != nil {
		toSerialize["context_params"] = o.ContextParams
	}
	return json.Marshal(toSerialize)
}

type NullableScopeCore struct {
	value *ScopeCore
	isSet bool
}

func (v NullableScopeCore) Get() *ScopeCore {
	return v.value
}

func (v *NullableScopeCore) Set(val *ScopeCore) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeCore) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeCore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeCore(val *ScopeCore) *NullableScopeCore {
	return &NullableScopeCore{value: val, isSet: true}
}

func (v NullableScopeCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeCore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


