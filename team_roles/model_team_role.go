/*
 * Team Roles API
 *
 * Description of the Team Roles API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package team_roles

import (
	"encoding/json"
)

// TeamRole struct for TeamRole
type TeamRole struct {
	Name *string `json:"name,omitempty"`
	RoleId *string `json:"role_id,omitempty"`
	ContextParams *ContextParams `json:"context_params,omitempty"`
}

// NewTeamRole instantiates a new TeamRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamRole() *TeamRole {
	this := TeamRole{}
	return &this
}

// NewTeamRoleWithDefaults instantiates a new TeamRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamRoleWithDefaults() *TeamRole {
	this := TeamRole{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TeamRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRole) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TeamRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TeamRole) SetName(v string) {
	o.Name = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *TeamRole) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRole) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *TeamRole) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *TeamRole) SetRoleId(v string) {
	o.RoleId = &v
}

// GetContextParams returns the ContextParams field value if set, zero value otherwise.
func (o *TeamRole) GetContextParams() ContextParams {
	if o == nil || o.ContextParams == nil {
		var ret ContextParams
		return ret
	}
	return *o.ContextParams
}

// GetContextParamsOk returns a tuple with the ContextParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRole) GetContextParamsOk() (*ContextParams, bool) {
	if o == nil || o.ContextParams == nil {
		return nil, false
	}
	return o.ContextParams, true
}

// HasContextParams returns a boolean if a field has been set.
func (o *TeamRole) HasContextParams() bool {
	if o != nil && o.ContextParams != nil {
		return true
	}

	return false
}

// SetContextParams gets a reference to the given ContextParams and assigns it to the ContextParams field.
func (o *TeamRole) SetContextParams(v ContextParams) {
	o.ContextParams = &v
}

func (o TeamRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RoleId != nil {
		toSerialize["role_id"] = o.RoleId
	}
	if o.ContextParams != nil {
		toSerialize["context_params"] = o.ContextParams
	}
	return json.Marshal(toSerialize)
}

type NullableTeamRole struct {
	value *TeamRole
	isSet bool
}

func (v NullableTeamRole) Get() *TeamRole {
	return v.value
}

func (v *NullableTeamRole) Set(val *TeamRole) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamRole) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamRole(val *TeamRole) *NullableTeamRole {
	return &NullableTeamRole{value: val, isSet: true}
}

func (v NullableTeamRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

