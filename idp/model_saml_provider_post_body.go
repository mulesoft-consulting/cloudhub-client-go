/*
Identity Provider Management API

Description of Identity Provider API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// checks if the SamlProviderPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlProviderPostBody{}

// SamlProviderPostBody struct for SamlProviderPostBody
type SamlProviderPostBody struct {
	ArcNamespace *string `json:"arc_namespace,omitempty"`
	Name string `json:"name"`
	Type SamlProviderPostBodyType `json:"type"`
	ServiceProvider SamlProviderPostBodyServiceProvider `json:"service_provider"`
	Saml SamlProviderPostBodySaml `json:"saml"`
	LoginDisabled *bool `json:"login_disabled,omitempty"`
}

// NewSamlProviderPostBody instantiates a new SamlProviderPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlProviderPostBody(name string, type_ SamlProviderPostBodyType, serviceProvider SamlProviderPostBodyServiceProvider, saml SamlProviderPostBodySaml) *SamlProviderPostBody {
	this := SamlProviderPostBody{}
	this.Name = name
	this.Type = type_
	this.ServiceProvider = serviceProvider
	this.Saml = saml
	return &this
}

// NewSamlProviderPostBodyWithDefaults instantiates a new SamlProviderPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlProviderPostBodyWithDefaults() *SamlProviderPostBody {
	this := SamlProviderPostBody{}
	return &this
}

// GetArcNamespace returns the ArcNamespace field value if set, zero value otherwise.
func (o *SamlProviderPostBody) GetArcNamespace() string {
	if o == nil || IsNil(o.ArcNamespace) {
		var ret string
		return ret
	}
	return *o.ArcNamespace
}

// GetArcNamespaceOk returns a tuple with the ArcNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPostBody) GetArcNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.ArcNamespace) {
		return nil, false
	}
	return o.ArcNamespace, true
}

// HasArcNamespace returns a boolean if a field has been set.
func (o *SamlProviderPostBody) HasArcNamespace() bool {
	if o != nil && !IsNil(o.ArcNamespace) {
		return true
	}

	return false
}

// SetArcNamespace gets a reference to the given string and assigns it to the ArcNamespace field.
func (o *SamlProviderPostBody) SetArcNamespace(v string) {
	o.ArcNamespace = &v
}

// GetName returns the Name field value
func (o *SamlProviderPostBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SamlProviderPostBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SamlProviderPostBody) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *SamlProviderPostBody) GetType() SamlProviderPostBodyType {
	if o == nil {
		var ret SamlProviderPostBodyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SamlProviderPostBody) GetTypeOk() (*SamlProviderPostBodyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SamlProviderPostBody) SetType(v SamlProviderPostBodyType) {
	o.Type = v
}

// GetServiceProvider returns the ServiceProvider field value
func (o *SamlProviderPostBody) GetServiceProvider() SamlProviderPostBodyServiceProvider {
	if o == nil {
		var ret SamlProviderPostBodyServiceProvider
		return ret
	}

	return o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value
// and a boolean to check if the value has been set.
func (o *SamlProviderPostBody) GetServiceProviderOk() (*SamlProviderPostBodyServiceProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceProvider, true
}

// SetServiceProvider sets field value
func (o *SamlProviderPostBody) SetServiceProvider(v SamlProviderPostBodyServiceProvider) {
	o.ServiceProvider = v
}

// GetSaml returns the Saml field value
func (o *SamlProviderPostBody) GetSaml() SamlProviderPostBodySaml {
	if o == nil {
		var ret SamlProviderPostBodySaml
		return ret
	}

	return o.Saml
}

// GetSamlOk returns a tuple with the Saml field value
// and a boolean to check if the value has been set.
func (o *SamlProviderPostBody) GetSamlOk() (*SamlProviderPostBodySaml, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Saml, true
}

// SetSaml sets field value
func (o *SamlProviderPostBody) SetSaml(v SamlProviderPostBodySaml) {
	o.Saml = v
}

// GetLoginDisabled returns the LoginDisabled field value if set, zero value otherwise.
func (o *SamlProviderPostBody) GetLoginDisabled() bool {
	if o == nil || IsNil(o.LoginDisabled) {
		var ret bool
		return ret
	}
	return *o.LoginDisabled
}

// GetLoginDisabledOk returns a tuple with the LoginDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPostBody) GetLoginDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LoginDisabled) {
		return nil, false
	}
	return o.LoginDisabled, true
}

// HasLoginDisabled returns a boolean if a field has been set.
func (o *SamlProviderPostBody) HasLoginDisabled() bool {
	if o != nil && !IsNil(o.LoginDisabled) {
		return true
	}

	return false
}

// SetLoginDisabled gets a reference to the given bool and assigns it to the LoginDisabled field.
func (o *SamlProviderPostBody) SetLoginDisabled(v bool) {
	o.LoginDisabled = &v
}

func (o SamlProviderPostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlProviderPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArcNamespace) {
		toSerialize["arc_namespace"] = o.ArcNamespace
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["service_provider"] = o.ServiceProvider
	toSerialize["saml"] = o.Saml
	if !IsNil(o.LoginDisabled) {
		toSerialize["login_disabled"] = o.LoginDisabled
	}
	return toSerialize, nil
}

type NullableSamlProviderPostBody struct {
	value *SamlProviderPostBody
	isSet bool
}

func (v NullableSamlProviderPostBody) Get() *SamlProviderPostBody {
	return v.value
}

func (v *NullableSamlProviderPostBody) Set(val *SamlProviderPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlProviderPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlProviderPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlProviderPostBody(val *SamlProviderPostBody) *NullableSamlProviderPostBody {
	return &NullableSamlProviderPostBody{value: val, isSet: true}
}

func (v NullableSamlProviderPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlProviderPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

