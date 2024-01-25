/*
Flex Gateway API

Description of the Flex Gateway API

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexgateway

import (
	"encoding/json"
)

// checks if the FlexGatewayRegistrationToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlexGatewayRegistrationToken{}

// FlexGatewayRegistrationToken struct for FlexGatewayRegistrationToken
type FlexGatewayRegistrationToken struct {
	RegistrationToken *string `json:"registrationToken,omitempty"`
}

// NewFlexGatewayRegistrationToken instantiates a new FlexGatewayRegistrationToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexGatewayRegistrationToken() *FlexGatewayRegistrationToken {
	this := FlexGatewayRegistrationToken{}
	return &this
}

// NewFlexGatewayRegistrationTokenWithDefaults instantiates a new FlexGatewayRegistrationToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexGatewayRegistrationTokenWithDefaults() *FlexGatewayRegistrationToken {
	this := FlexGatewayRegistrationToken{}
	return &this
}

// GetRegistrationToken returns the RegistrationToken field value if set, zero value otherwise.
func (o *FlexGatewayRegistrationToken) GetRegistrationToken() string {
	if o == nil || IsNil(o.RegistrationToken) {
		var ret string
		return ret
	}
	return *o.RegistrationToken
}

// GetRegistrationTokenOk returns a tuple with the RegistrationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayRegistrationToken) GetRegistrationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationToken) {
		return nil, false
	}
	return o.RegistrationToken, true
}

// HasRegistrationToken returns a boolean if a field has been set.
func (o *FlexGatewayRegistrationToken) HasRegistrationToken() bool {
	if o != nil && !IsNil(o.RegistrationToken) {
		return true
	}

	return false
}

// SetRegistrationToken gets a reference to the given string and assigns it to the RegistrationToken field.
func (o *FlexGatewayRegistrationToken) SetRegistrationToken(v string) {
	o.RegistrationToken = &v
}

func (o FlexGatewayRegistrationToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlexGatewayRegistrationToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegistrationToken) {
		toSerialize["registrationToken"] = o.RegistrationToken
	}
	return toSerialize, nil
}

type NullableFlexGatewayRegistrationToken struct {
	value *FlexGatewayRegistrationToken
	isSet bool
}

func (v NullableFlexGatewayRegistrationToken) Get() *FlexGatewayRegistrationToken {
	return v.value
}

func (v *NullableFlexGatewayRegistrationToken) Set(val *FlexGatewayRegistrationToken) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexGatewayRegistrationToken) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexGatewayRegistrationToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexGatewayRegistrationToken(val *FlexGatewayRegistrationToken) *NullableFlexGatewayRegistrationToken {
	return &NullableFlexGatewayRegistrationToken{value: val, isSet: true}
}

func (v NullableFlexGatewayRegistrationToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexGatewayRegistrationToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

