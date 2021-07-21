/*
 * Alert API Manager
 *
 * Alert for API Manager
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alert_apim

import (
	"encoding/json"
)

// Alert struct for Alert
type Alert struct {
	ApiAlertsVersion string `json:"apiAlertsVersion"`
	Name string `json:"name"`
	Type string `json:"type"`
	Enabled bool `json:"enabled"`
	Severity string `json:"severity"`
	Recipients []Recipient `json:"recipients"`
	Condition Condition `json:"condition"`
	Period Period `json:"period"`
}

// NewAlert instantiates a new Alert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlert(apiAlertsVersion string, name string, type_ string, enabled bool, severity string, recipients []Recipient, condition Condition, period Period) *Alert {
	this := Alert{}
	this.ApiAlertsVersion = apiAlertsVersion
	this.Name = name
	this.Type = type_
	this.Enabled = enabled
	this.Severity = severity
	this.Recipients = recipients
	this.Condition = condition
	this.Period = period
	return &this
}

// NewAlertWithDefaults instantiates a new Alert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertWithDefaults() *Alert {
	this := Alert{}
	return &this
}

// GetApiAlertsVersion returns the ApiAlertsVersion field value
func (o *Alert) GetApiAlertsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiAlertsVersion
}

// GetApiAlertsVersionOk returns a tuple with the ApiAlertsVersion field value
// and a boolean to check if the value has been set.
func (o *Alert) GetApiAlertsVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiAlertsVersion, true
}

// SetApiAlertsVersion sets field value
func (o *Alert) SetApiAlertsVersion(v string) {
	o.ApiAlertsVersion = v
}

// GetName returns the Name field value
func (o *Alert) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Alert) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Alert) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Alert) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Alert) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Alert) SetType(v string) {
	o.Type = v
}

// GetEnabled returns the Enabled field value
func (o *Alert) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Alert) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Alert) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSeverity returns the Severity field value
func (o *Alert) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Alert) GetSeverityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Alert) SetSeverity(v string) {
	o.Severity = v
}

// GetRecipients returns the Recipients field value
func (o *Alert) GetRecipients() []Recipient {
	if o == nil {
		var ret []Recipient
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *Alert) GetRecipientsOk() (*[]Recipient, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *Alert) SetRecipients(v []Recipient) {
	o.Recipients = v
}

// GetCondition returns the Condition field value
func (o *Alert) GetCondition() Condition {
	if o == nil {
		var ret Condition
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *Alert) GetConditionOk() (*Condition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *Alert) SetCondition(v Condition) {
	o.Condition = v
}

// GetPeriod returns the Period field value
func (o *Alert) GetPeriod() Period {
	if o == nil {
		var ret Period
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *Alert) GetPeriodOk() (*Period, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *Alert) SetPeriod(v Period) {
	o.Period = v
}

func (o Alert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiAlertsVersion"] = o.ApiAlertsVersion
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["condition"] = o.Condition
	}
	if true {
		toSerialize["period"] = o.Period
	}
	return json.Marshal(toSerialize)
}

type NullableAlert struct {
	value *Alert
	isSet bool
}

func (v NullableAlert) Get() *Alert {
	return v.value
}

func (v *NullableAlert) Set(val *Alert) {
	v.value = val
	v.isSet = true
}

func (v NullableAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlert(val *Alert) *NullableAlert {
	return &NullableAlert{value: val, isSet: true}
}

func (v NullableAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


