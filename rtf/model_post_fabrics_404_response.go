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

// checks if the PostFabrics404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostFabrics404Response{}

// PostFabrics404Response struct for PostFabrics404Response
type PostFabrics404Response struct {
	Name *string `json:"name,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewPostFabrics404Response instantiates a new PostFabrics404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostFabrics404Response() *PostFabrics404Response {
	this := PostFabrics404Response{}
	return &this
}

// NewPostFabrics404ResponseWithDefaults instantiates a new PostFabrics404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostFabrics404ResponseWithDefaults() *PostFabrics404Response {
	this := PostFabrics404Response{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostFabrics404Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFabrics404Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostFabrics404Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostFabrics404Response) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostFabrics404Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFabrics404Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostFabrics404Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostFabrics404Response) SetMessage(v string) {
	o.Message = &v
}

func (o PostFabrics404Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostFabrics404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePostFabrics404Response struct {
	value *PostFabrics404Response
	isSet bool
}

func (v NullablePostFabrics404Response) Get() *PostFabrics404Response {
	return v.value
}

func (v *NullablePostFabrics404Response) Set(val *PostFabrics404Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostFabrics404Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostFabrics404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostFabrics404Response(val *PostFabrics404Response) *NullablePostFabrics404Response {
	return &NullablePostFabrics404Response{value: val, isSet: true}
}

func (v NullablePostFabrics404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostFabrics404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

