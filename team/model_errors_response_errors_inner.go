/*
Team API

Description of the Team API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package team

import (
	"encoding/json"
)

// checks if the ErrorsResponseErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorsResponseErrorsInner{}

// ErrorsResponseErrorsInner struct for ErrorsResponseErrorsInner
type ErrorsResponseErrorsInner struct {
	Type *string `json:"type,omitempty"`
	DataPath *string `json:"dataPath,omitempty"`
	Keyword *string `json:"keyword,omitempty"`
	Schema *string `json:"schema,omitempty"`
	Data *string `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewErrorsResponseErrorsInner instantiates a new ErrorsResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsResponseErrorsInner() *ErrorsResponseErrorsInner {
	this := ErrorsResponseErrorsInner{}
	return &this
}

// NewErrorsResponseErrorsInnerWithDefaults instantiates a new ErrorsResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsResponseErrorsInnerWithDefaults() *ErrorsResponseErrorsInner {
	this := ErrorsResponseErrorsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorsResponseErrorsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponseErrorsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorsResponseErrorsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorsResponseErrorsInner) SetType(v string) {
	o.Type = &v
}

// GetDataPath returns the DataPath field value if set, zero value otherwise.
func (o *ErrorsResponseErrorsInner) GetDataPath() string {
	if o == nil || IsNil(o.DataPath) {
		var ret string
		return ret
	}
	return *o.DataPath
}

// GetDataPathOk returns a tuple with the DataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponseErrorsInner) GetDataPathOk() (*string, bool) {
	if o == nil || IsNil(o.DataPath) {
		return nil, false
	}
	return o.DataPath, true
}

// HasDataPath returns a boolean if a field has been set.
func (o *ErrorsResponseErrorsInner) HasDataPath() bool {
	if o != nil && !IsNil(o.DataPath) {
		return true
	}

	return false
}

// SetDataPath gets a reference to the given string and assigns it to the DataPath field.
func (o *ErrorsResponseErrorsInner) SetDataPath(v string) {
	o.DataPath = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *ErrorsResponseErrorsInner) GetKeyword() string {
	if o == nil || IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponseErrorsInner) GetKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *ErrorsResponseErrorsInner) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *ErrorsResponseErrorsInner) SetKeyword(v string) {
	o.Keyword = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ErrorsResponseErrorsInner) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponseErrorsInner) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ErrorsResponseErrorsInner) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ErrorsResponseErrorsInner) SetSchema(v string) {
	o.Schema = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ErrorsResponseErrorsInner) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponseErrorsInner) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ErrorsResponseErrorsInner) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ErrorsResponseErrorsInner) SetData(v string) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorsResponseErrorsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponseErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorsResponseErrorsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorsResponseErrorsInner) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorsResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorsResponseErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DataPath) {
		toSerialize["dataPath"] = o.DataPath
	}
	if !IsNil(o.Keyword) {
		toSerialize["keyword"] = o.Keyword
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableErrorsResponseErrorsInner struct {
	value *ErrorsResponseErrorsInner
	isSet bool
}

func (v NullableErrorsResponseErrorsInner) Get() *ErrorsResponseErrorsInner {
	return v.value
}

func (v *NullableErrorsResponseErrorsInner) Set(val *ErrorsResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsResponseErrorsInner(val *ErrorsResponseErrorsInner) *NullableErrorsResponseErrorsInner {
	return &NullableErrorsResponseErrorsInner{value: val, isSet: true}
}

func (v NullableErrorsResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


