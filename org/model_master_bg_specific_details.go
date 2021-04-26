/*
 * Organization API
 *
 * Description of the Organization API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// MasterBGSpecificDetails struct for MasterBGSpecificDetails
type MasterBGSpecificDetails struct {
	// An explanation about the purpose of this instance.
	SessionTimeout int32 `json:"sessionTimeout"`
	// An explanation about the purpose of this instance.
	Subscription map[string]map[string]interface{} `json:"subscription"`
}

// NewMasterBGSpecificDetails instantiates a new MasterBGSpecificDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterBGSpecificDetails(sessionTimeout int32, subscription map[string]map[string]interface{}) *MasterBGSpecificDetails {
	this := MasterBGSpecificDetails{}
	this.SessionTimeout = sessionTimeout
	this.Subscription = subscription
	return &this
}

// NewMasterBGSpecificDetailsWithDefaults instantiates a new MasterBGSpecificDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterBGSpecificDetailsWithDefaults() *MasterBGSpecificDetails {
	this := MasterBGSpecificDetails{}
	var sessionTimeout int32 = 0
	this.SessionTimeout = sessionTimeout
	return &this
}

// GetSessionTimeout returns the SessionTimeout field value
func (o *MasterBGSpecificDetails) GetSessionTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SessionTimeout
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value
// and a boolean to check if the value has been set.
func (o *MasterBGSpecificDetails) GetSessionTimeoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SessionTimeout, true
}

// SetSessionTimeout sets field value
func (o *MasterBGSpecificDetails) SetSessionTimeout(v int32) {
	o.SessionTimeout = v
}

// GetSubscription returns the Subscription field value
func (o *MasterBGSpecificDetails) GetSubscription() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *MasterBGSpecificDetails) GetSubscriptionOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *MasterBGSpecificDetails) SetSubscription(v map[string]map[string]interface{}) {
	o.Subscription = v
}

func (o MasterBGSpecificDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sessionTimeout"] = o.SessionTimeout
	}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableMasterBGSpecificDetails struct {
	value *MasterBGSpecificDetails
	isSet bool
}

func (v NullableMasterBGSpecificDetails) Get() *MasterBGSpecificDetails {
	return v.value
}

func (v *NullableMasterBGSpecificDetails) Set(val *MasterBGSpecificDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterBGSpecificDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterBGSpecificDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterBGSpecificDetails(val *MasterBGSpecificDetails) *NullableMasterBGSpecificDetails {
	return &NullableMasterBGSpecificDetails{value: val, isSet: true}
}

func (v NullableMasterBGSpecificDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterBGSpecificDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

