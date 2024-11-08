/*
Deployment

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Runtime Fabric and CloudHub 2.0 targets only. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v2

import (
	"encoding/json"
)

// checks if the MuleAgentAppPropService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MuleAgentAppPropService{}

// MuleAgentAppPropService struct for MuleAgentAppPropService
type MuleAgentAppPropService struct {
	ApplicationName *string `json:"applicationName,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	SecureProperties map[string]interface{} `json:"secureProperties,omitempty"`
}

// NewMuleAgentAppPropService instantiates a new MuleAgentAppPropService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMuleAgentAppPropService() *MuleAgentAppPropService {
	this := MuleAgentAppPropService{}
	return &this
}

// NewMuleAgentAppPropServiceWithDefaults instantiates a new MuleAgentAppPropService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMuleAgentAppPropServiceWithDefaults() *MuleAgentAppPropService {
	this := MuleAgentAppPropService{}
	return &this
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *MuleAgentAppPropService) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuleAgentAppPropService) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *MuleAgentAppPropService) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *MuleAgentAppPropService) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MuleAgentAppPropService) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuleAgentAppPropService) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MuleAgentAppPropService) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *MuleAgentAppPropService) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetSecureProperties returns the SecureProperties field value if set, zero value otherwise.
func (o *MuleAgentAppPropService) GetSecureProperties() map[string]interface{} {
	if o == nil || IsNil(o.SecureProperties) {
		var ret map[string]interface{}
		return ret
	}
	return o.SecureProperties
}

// GetSecurePropertiesOk returns a tuple with the SecureProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuleAgentAppPropService) GetSecurePropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SecureProperties) {
		return map[string]interface{}{}, false
	}
	return o.SecureProperties, true
}

// HasSecureProperties returns a boolean if a field has been set.
func (o *MuleAgentAppPropService) HasSecureProperties() bool {
	if o != nil && !IsNil(o.SecureProperties) {
		return true
	}

	return false
}

// SetSecureProperties gets a reference to the given map[string]interface{} and assigns it to the SecureProperties field.
func (o *MuleAgentAppPropService) SetSecureProperties(v map[string]interface{}) {
	o.SecureProperties = v
}

func (o MuleAgentAppPropService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MuleAgentAppPropService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationName) {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.SecureProperties) {
		toSerialize["secureProperties"] = o.SecureProperties
	}
	return toSerialize, nil
}

type NullableMuleAgentAppPropService struct {
	value *MuleAgentAppPropService
	isSet bool
}

func (v NullableMuleAgentAppPropService) Get() *MuleAgentAppPropService {
	return v.value
}

func (v *NullableMuleAgentAppPropService) Set(val *MuleAgentAppPropService) {
	v.value = val
	v.isSet = true
}

func (v NullableMuleAgentAppPropService) IsSet() bool {
	return v.isSet
}

func (v *NullableMuleAgentAppPropService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMuleAgentAppPropService(val *MuleAgentAppPropService) *NullableMuleAgentAppPropService {
	return &NullableMuleAgentAppPropService{value: val, isSet: true}
}

func (v NullableMuleAgentAppPropService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMuleAgentAppPropService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

