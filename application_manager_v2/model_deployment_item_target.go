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

// checks if the DeploymentItemTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentItemTarget{}

// DeploymentItemTarget struct for DeploymentItemTarget
type DeploymentItemTarget struct {
	Provider *string `json:"provider,omitempty"`
	TargetId *string `json:"targetId,omitempty"`
}

// NewDeploymentItemTarget instantiates a new DeploymentItemTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentItemTarget() *DeploymentItemTarget {
	this := DeploymentItemTarget{}
	return &this
}

// NewDeploymentItemTargetWithDefaults instantiates a new DeploymentItemTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentItemTargetWithDefaults() *DeploymentItemTarget {
	this := DeploymentItemTarget{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *DeploymentItemTarget) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentItemTarget) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *DeploymentItemTarget) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *DeploymentItemTarget) SetProvider(v string) {
	o.Provider = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *DeploymentItemTarget) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentItemTarget) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *DeploymentItemTarget) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *DeploymentItemTarget) SetTargetId(v string) {
	o.TargetId = &v
}

func (o DeploymentItemTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentItemTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	return toSerialize, nil
}

type NullableDeploymentItemTarget struct {
	value *DeploymentItemTarget
	isSet bool
}

func (v NullableDeploymentItemTarget) Get() *DeploymentItemTarget {
	return v.value
}

func (v *NullableDeploymentItemTarget) Set(val *DeploymentItemTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentItemTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentItemTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentItemTarget(val *DeploymentItemTarget) *NullableDeploymentItemTarget {
	return &NullableDeploymentItemTarget{value: val, isSet: true}
}

func (v NullableDeploymentItemTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentItemTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


