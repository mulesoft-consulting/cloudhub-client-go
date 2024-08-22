/*
Runtime Fabrics API

Runtime Fabrics API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtf

import (
	"encoding/json"
)

// checks if the ErrorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorsResponse{}

// ErrorsResponse struct for ErrorsResponse
type ErrorsResponse struct {
	Errors []ErrorsResponseErrorsInner `json:"errors,omitempty"`
}

// NewErrorsResponse instantiates a new ErrorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsResponse() *ErrorsResponse {
	this := ErrorsResponse{}
	return &this
}

// NewErrorsResponseWithDefaults instantiates a new ErrorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsResponseWithDefaults() *ErrorsResponse {
	this := ErrorsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ErrorsResponse) GetErrors() []ErrorsResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []ErrorsResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponse) GetErrorsOk() ([]ErrorsResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ErrorsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorsResponseErrorsInner and assigns it to the Errors field.
func (o *ErrorsResponse) SetErrors(v []ErrorsResponseErrorsInner) {
	o.Errors = v
}

func (o ErrorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableErrorsResponse struct {
	value *ErrorsResponse
	isSet bool
}

func (v NullableErrorsResponse) Get() *ErrorsResponse {
	return v.value
}

func (v *NullableErrorsResponse) Set(val *ErrorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsResponse(val *ErrorsResponse) *NullableErrorsResponse {
	return &NullableErrorsResponse{value: val, isSet: true}
}

func (v NullableErrorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


