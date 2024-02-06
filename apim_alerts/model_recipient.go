/*
Alert API Manager

Alert for API Manager

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_alerts

import (
	"encoding/json"
)

// checks if the Recipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recipient{}

// Recipient struct for Recipient
type Recipient struct {
	Type string `json:"type"`
	Value string `json:"value"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

// NewRecipient instantiates a new Recipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipient(type_ string, value string, firstName string, lastName string) *Recipient {
	this := Recipient{}
	this.Type = type_
	this.Value = value
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewRecipientWithDefaults instantiates a new Recipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientWithDefaults() *Recipient {
	this := Recipient{}
	return &this
}

// GetType returns the Type field value
func (o *Recipient) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Recipient) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Recipient) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *Recipient) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Recipient) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Recipient) SetValue(v string) {
	o.Value = v
}

// GetFirstName returns the FirstName field value
func (o *Recipient) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *Recipient) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *Recipient) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *Recipient) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *Recipient) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *Recipient) SetLastName(v string) {
	o.LastName = v
}

func (o Recipient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	return toSerialize, nil
}

type NullableRecipient struct {
	value *Recipient
	isSet bool
}

func (v NullableRecipient) Get() *Recipient {
	return v.value
}

func (v *NullableRecipient) Set(val *Recipient) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipient(val *Recipient) *NullableRecipient {
	return &NullableRecipient{value: val, isSet: true}
}

func (v NullableRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


