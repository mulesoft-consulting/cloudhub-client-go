/*
 * Connected App API
 *
 * Description of the Connected App API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connected_app

import (
	"encoding/json"
)

// ConnectedAppPatchExtAllOf struct for ConnectedAppPatchExtAllOf
type ConnectedAppPatchExtAllOf struct {
	ClientId *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewConnectedAppPatchExtAllOf instantiates a new ConnectedAppPatchExtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedAppPatchExtAllOf() *ConnectedAppPatchExtAllOf {
	this := ConnectedAppPatchExtAllOf{}
	return &this
}

// NewConnectedAppPatchExtAllOfWithDefaults instantiates a new ConnectedAppPatchExtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedAppPatchExtAllOfWithDefaults() *ConnectedAppPatchExtAllOf {
	this := ConnectedAppPatchExtAllOf{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ConnectedAppPatchExtAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppPatchExtAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ConnectedAppPatchExtAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ConnectedAppPatchExtAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ConnectedAppPatchExtAllOf) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppPatchExtAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ConnectedAppPatchExtAllOf) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ConnectedAppPatchExtAllOf) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ConnectedAppPatchExtAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppPatchExtAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ConnectedAppPatchExtAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ConnectedAppPatchExtAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ConnectedAppPatchExtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableConnectedAppPatchExtAllOf struct {
	value *ConnectedAppPatchExtAllOf
	isSet bool
}

func (v NullableConnectedAppPatchExtAllOf) Get() *ConnectedAppPatchExtAllOf {
	return v.value
}

func (v *NullableConnectedAppPatchExtAllOf) Set(val *ConnectedAppPatchExtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedAppPatchExtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedAppPatchExtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedAppPatchExtAllOf(val *ConnectedAppPatchExtAllOf) *NullableConnectedAppPatchExtAllOf {
	return &NullableConnectedAppPatchExtAllOf{value: val, isSet: true}
}

func (v NullableConnectedAppPatchExtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedAppPatchExtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


