/*
 * Anypoint MQ Exchange specfication
 *
 * Anypoint MQ Exchange API specification
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ame

import (
	"encoding/json"
)

// ExchangeBody struct for ExchangeBody
type ExchangeBody struct {
	Encrypted *bool `json:"encrypted,omitempty"`
	Type *string `json:"type,omitempty"`
	ExchangeId *string `json:"exchangeId,omitempty"`
}

// NewExchangeBody instantiates a new ExchangeBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeBody() *ExchangeBody {
	this := ExchangeBody{}
	return &this
}

// NewExchangeBodyWithDefaults instantiates a new ExchangeBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeBodyWithDefaults() *ExchangeBody {
	this := ExchangeBody{}
	return &this
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *ExchangeBody) GetEncrypted() bool {
	if o == nil || o.Encrypted == nil {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeBody) GetEncryptedOk() (*bool, bool) {
	if o == nil || o.Encrypted == nil {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *ExchangeBody) HasEncrypted() bool {
	if o != nil && o.Encrypted != nil {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *ExchangeBody) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExchangeBody) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeBody) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExchangeBody) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExchangeBody) SetType(v string) {
	o.Type = &v
}

// GetExchangeId returns the ExchangeId field value if set, zero value otherwise.
func (o *ExchangeBody) GetExchangeId() string {
	if o == nil || o.ExchangeId == nil {
		var ret string
		return ret
	}
	return *o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeBody) GetExchangeIdOk() (*string, bool) {
	if o == nil || o.ExchangeId == nil {
		return nil, false
	}
	return o.ExchangeId, true
}

// HasExchangeId returns a boolean if a field has been set.
func (o *ExchangeBody) HasExchangeId() bool {
	if o != nil && o.ExchangeId != nil {
		return true
	}

	return false
}

// SetExchangeId gets a reference to the given string and assigns it to the ExchangeId field.
func (o *ExchangeBody) SetExchangeId(v string) {
	o.ExchangeId = &v
}

func (o ExchangeBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Encrypted != nil {
		toSerialize["encrypted"] = o.Encrypted
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ExchangeId != nil {
		toSerialize["exchangeId"] = o.ExchangeId
	}
	return json.Marshal(toSerialize)
}

type NullableExchangeBody struct {
	value *ExchangeBody
	isSet bool
}

func (v NullableExchangeBody) Get() *ExchangeBody {
	return v.value
}

func (v *NullableExchangeBody) Set(val *ExchangeBody) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeBody) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeBody(val *ExchangeBody) *NullableExchangeBody {
	return &NullableExchangeBody{value: val, isSet: true}
}

func (v NullableExchangeBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


