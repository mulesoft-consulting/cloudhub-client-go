/*
Secret Group TLS Context API

Secret Group TLS Context API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_tlscontext

import (
	"encoding/json"
)

// checks if the TlsContextDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TlsContextDetails{}

// TlsContextDetails struct for TlsContextDetails
type TlsContextDetails struct {
	// The target application
	Target *string `json:"target,omitempty"`
	Name *string `json:"name,omitempty"`
	ExpirationDate *string `json:"expirationDate,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Keystore *SecretPath `json:"keystore,omitempty"`
	Truststore *SecretPath `json:"truststore,omitempty"`
	// This flag is to enable client authentication
	EnableMutualAuthentication *bool `json:"enableMutualAuthentication,omitempty"`
	MutualAuthentication *MutualAuthentication `json:"mutualAuthentication,omitempty"`
	AcceptableCipherSuites *AcceptableCipherSuites `json:"acceptableCipherSuites,omitempty"`
	AcceptableTlsVersions *AcceptableTlsVersions `json:"acceptableTlsVersions,omitempty"`
	// List of enabled cipher suites for Mule target
	CipherSuites []string `json:"cipherSuites,omitempty"`
	// Setting this flag to true indicates that certificate validation should not be enforced, i.e. the truststore, even though set, is ignored at runtime.
	Insecure *bool `json:"insecure,omitempty"`
	// Minimum TLS version supported.
	MinTlsVersion *string `json:"minTlsVersion,omitempty"`
	// Maximum TLS version supported.
	MaxTlsVersion *string `json:"maxTlsVersion,omitempty"`
	// supported HTTP versions in the most-to-least preferred order. At least one version must be specified.
	AlpnProtocols []string `json:"alpnProtocols,omitempty"`
	InboundSettings *TlsContextDetailsInboundSettings `json:"inboundSettings,omitempty"`
	OutboundSettings *TlsContextDetailsOutboundSettings `json:"outboundSettings,omitempty"`
}

// NewTlsContextDetails instantiates a new TlsContextDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsContextDetails() *TlsContextDetails {
	this := TlsContextDetails{}
	var insecure bool = false
	this.Insecure = &insecure
	return &this
}

// NewTlsContextDetailsWithDefaults instantiates a new TlsContextDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsContextDetailsWithDefaults() *TlsContextDetails {
	this := TlsContextDetails{}
	var insecure bool = false
	this.Insecure = &insecure
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *TlsContextDetails) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *TlsContextDetails) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *TlsContextDetails) SetTarget(v string) {
	o.Target = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TlsContextDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TlsContextDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TlsContextDetails) SetName(v string) {
	o.Name = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *TlsContextDetails) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *TlsContextDetails) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *TlsContextDetails) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TlsContextDetails) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TlsContextDetails) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *TlsContextDetails) SetMeta(v Meta) {
	o.Meta = &v
}

// GetKeystore returns the Keystore field value if set, zero value otherwise.
func (o *TlsContextDetails) GetKeystore() SecretPath {
	if o == nil || IsNil(o.Keystore) {
		var ret SecretPath
		return ret
	}
	return *o.Keystore
}

// GetKeystoreOk returns a tuple with the Keystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetKeystoreOk() (*SecretPath, bool) {
	if o == nil || IsNil(o.Keystore) {
		return nil, false
	}
	return o.Keystore, true
}

// HasKeystore returns a boolean if a field has been set.
func (o *TlsContextDetails) HasKeystore() bool {
	if o != nil && !IsNil(o.Keystore) {
		return true
	}

	return false
}

// SetKeystore gets a reference to the given SecretPath and assigns it to the Keystore field.
func (o *TlsContextDetails) SetKeystore(v SecretPath) {
	o.Keystore = &v
}

// GetTruststore returns the Truststore field value if set, zero value otherwise.
func (o *TlsContextDetails) GetTruststore() SecretPath {
	if o == nil || IsNil(o.Truststore) {
		var ret SecretPath
		return ret
	}
	return *o.Truststore
}

// GetTruststoreOk returns a tuple with the Truststore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetTruststoreOk() (*SecretPath, bool) {
	if o == nil || IsNil(o.Truststore) {
		return nil, false
	}
	return o.Truststore, true
}

// HasTruststore returns a boolean if a field has been set.
func (o *TlsContextDetails) HasTruststore() bool {
	if o != nil && !IsNil(o.Truststore) {
		return true
	}

	return false
}

// SetTruststore gets a reference to the given SecretPath and assigns it to the Truststore field.
func (o *TlsContextDetails) SetTruststore(v SecretPath) {
	o.Truststore = &v
}

// GetEnableMutualAuthentication returns the EnableMutualAuthentication field value if set, zero value otherwise.
func (o *TlsContextDetails) GetEnableMutualAuthentication() bool {
	if o == nil || IsNil(o.EnableMutualAuthentication) {
		var ret bool
		return ret
	}
	return *o.EnableMutualAuthentication
}

// GetEnableMutualAuthenticationOk returns a tuple with the EnableMutualAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetEnableMutualAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMutualAuthentication) {
		return nil, false
	}
	return o.EnableMutualAuthentication, true
}

// HasEnableMutualAuthentication returns a boolean if a field has been set.
func (o *TlsContextDetails) HasEnableMutualAuthentication() bool {
	if o != nil && !IsNil(o.EnableMutualAuthentication) {
		return true
	}

	return false
}

// SetEnableMutualAuthentication gets a reference to the given bool and assigns it to the EnableMutualAuthentication field.
func (o *TlsContextDetails) SetEnableMutualAuthentication(v bool) {
	o.EnableMutualAuthentication = &v
}

// GetMutualAuthentication returns the MutualAuthentication field value if set, zero value otherwise.
func (o *TlsContextDetails) GetMutualAuthentication() MutualAuthentication {
	if o == nil || IsNil(o.MutualAuthentication) {
		var ret MutualAuthentication
		return ret
	}
	return *o.MutualAuthentication
}

// GetMutualAuthenticationOk returns a tuple with the MutualAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetMutualAuthenticationOk() (*MutualAuthentication, bool) {
	if o == nil || IsNil(o.MutualAuthentication) {
		return nil, false
	}
	return o.MutualAuthentication, true
}

// HasMutualAuthentication returns a boolean if a field has been set.
func (o *TlsContextDetails) HasMutualAuthentication() bool {
	if o != nil && !IsNil(o.MutualAuthentication) {
		return true
	}

	return false
}

// SetMutualAuthentication gets a reference to the given MutualAuthentication and assigns it to the MutualAuthentication field.
func (o *TlsContextDetails) SetMutualAuthentication(v MutualAuthentication) {
	o.MutualAuthentication = &v
}

// GetAcceptableCipherSuites returns the AcceptableCipherSuites field value if set, zero value otherwise.
func (o *TlsContextDetails) GetAcceptableCipherSuites() AcceptableCipherSuites {
	if o == nil || IsNil(o.AcceptableCipherSuites) {
		var ret AcceptableCipherSuites
		return ret
	}
	return *o.AcceptableCipherSuites
}

// GetAcceptableCipherSuitesOk returns a tuple with the AcceptableCipherSuites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetAcceptableCipherSuitesOk() (*AcceptableCipherSuites, bool) {
	if o == nil || IsNil(o.AcceptableCipherSuites) {
		return nil, false
	}
	return o.AcceptableCipherSuites, true
}

// HasAcceptableCipherSuites returns a boolean if a field has been set.
func (o *TlsContextDetails) HasAcceptableCipherSuites() bool {
	if o != nil && !IsNil(o.AcceptableCipherSuites) {
		return true
	}

	return false
}

// SetAcceptableCipherSuites gets a reference to the given AcceptableCipherSuites and assigns it to the AcceptableCipherSuites field.
func (o *TlsContextDetails) SetAcceptableCipherSuites(v AcceptableCipherSuites) {
	o.AcceptableCipherSuites = &v
}

// GetAcceptableTlsVersions returns the AcceptableTlsVersions field value if set, zero value otherwise.
func (o *TlsContextDetails) GetAcceptableTlsVersions() AcceptableTlsVersions {
	if o == nil || IsNil(o.AcceptableTlsVersions) {
		var ret AcceptableTlsVersions
		return ret
	}
	return *o.AcceptableTlsVersions
}

// GetAcceptableTlsVersionsOk returns a tuple with the AcceptableTlsVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetAcceptableTlsVersionsOk() (*AcceptableTlsVersions, bool) {
	if o == nil || IsNil(o.AcceptableTlsVersions) {
		return nil, false
	}
	return o.AcceptableTlsVersions, true
}

// HasAcceptableTlsVersions returns a boolean if a field has been set.
func (o *TlsContextDetails) HasAcceptableTlsVersions() bool {
	if o != nil && !IsNil(o.AcceptableTlsVersions) {
		return true
	}

	return false
}

// SetAcceptableTlsVersions gets a reference to the given AcceptableTlsVersions and assigns it to the AcceptableTlsVersions field.
func (o *TlsContextDetails) SetAcceptableTlsVersions(v AcceptableTlsVersions) {
	o.AcceptableTlsVersions = &v
}

// GetCipherSuites returns the CipherSuites field value if set, zero value otherwise.
func (o *TlsContextDetails) GetCipherSuites() []string {
	if o == nil || IsNil(o.CipherSuites) {
		var ret []string
		return ret
	}
	return o.CipherSuites
}

// GetCipherSuitesOk returns a tuple with the CipherSuites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetCipherSuitesOk() ([]string, bool) {
	if o == nil || IsNil(o.CipherSuites) {
		return nil, false
	}
	return o.CipherSuites, true
}

// HasCipherSuites returns a boolean if a field has been set.
func (o *TlsContextDetails) HasCipherSuites() bool {
	if o != nil && !IsNil(o.CipherSuites) {
		return true
	}

	return false
}

// SetCipherSuites gets a reference to the given []string and assigns it to the CipherSuites field.
func (o *TlsContextDetails) SetCipherSuites(v []string) {
	o.CipherSuites = v
}

// GetInsecure returns the Insecure field value if set, zero value otherwise.
func (o *TlsContextDetails) GetInsecure() bool {
	if o == nil || IsNil(o.Insecure) {
		var ret bool
		return ret
	}
	return *o.Insecure
}

// GetInsecureOk returns a tuple with the Insecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetInsecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Insecure) {
		return nil, false
	}
	return o.Insecure, true
}

// HasInsecure returns a boolean if a field has been set.
func (o *TlsContextDetails) HasInsecure() bool {
	if o != nil && !IsNil(o.Insecure) {
		return true
	}

	return false
}

// SetInsecure gets a reference to the given bool and assigns it to the Insecure field.
func (o *TlsContextDetails) SetInsecure(v bool) {
	o.Insecure = &v
}

// GetMinTlsVersion returns the MinTlsVersion field value if set, zero value otherwise.
func (o *TlsContextDetails) GetMinTlsVersion() string {
	if o == nil || IsNil(o.MinTlsVersion) {
		var ret string
		return ret
	}
	return *o.MinTlsVersion
}

// GetMinTlsVersionOk returns a tuple with the MinTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetMinTlsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinTlsVersion) {
		return nil, false
	}
	return o.MinTlsVersion, true
}

// HasMinTlsVersion returns a boolean if a field has been set.
func (o *TlsContextDetails) HasMinTlsVersion() bool {
	if o != nil && !IsNil(o.MinTlsVersion) {
		return true
	}

	return false
}

// SetMinTlsVersion gets a reference to the given string and assigns it to the MinTlsVersion field.
func (o *TlsContextDetails) SetMinTlsVersion(v string) {
	o.MinTlsVersion = &v
}

// GetMaxTlsVersion returns the MaxTlsVersion field value if set, zero value otherwise.
func (o *TlsContextDetails) GetMaxTlsVersion() string {
	if o == nil || IsNil(o.MaxTlsVersion) {
		var ret string
		return ret
	}
	return *o.MaxTlsVersion
}

// GetMaxTlsVersionOk returns a tuple with the MaxTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetMaxTlsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MaxTlsVersion) {
		return nil, false
	}
	return o.MaxTlsVersion, true
}

// HasMaxTlsVersion returns a boolean if a field has been set.
func (o *TlsContextDetails) HasMaxTlsVersion() bool {
	if o != nil && !IsNil(o.MaxTlsVersion) {
		return true
	}

	return false
}

// SetMaxTlsVersion gets a reference to the given string and assigns it to the MaxTlsVersion field.
func (o *TlsContextDetails) SetMaxTlsVersion(v string) {
	o.MaxTlsVersion = &v
}

// GetAlpnProtocols returns the AlpnProtocols field value if set, zero value otherwise.
func (o *TlsContextDetails) GetAlpnProtocols() []string {
	if o == nil || IsNil(o.AlpnProtocols) {
		var ret []string
		return ret
	}
	return o.AlpnProtocols
}

// GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetAlpnProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.AlpnProtocols) {
		return nil, false
	}
	return o.AlpnProtocols, true
}

// HasAlpnProtocols returns a boolean if a field has been set.
func (o *TlsContextDetails) HasAlpnProtocols() bool {
	if o != nil && !IsNil(o.AlpnProtocols) {
		return true
	}

	return false
}

// SetAlpnProtocols gets a reference to the given []string and assigns it to the AlpnProtocols field.
func (o *TlsContextDetails) SetAlpnProtocols(v []string) {
	o.AlpnProtocols = v
}

// GetInboundSettings returns the InboundSettings field value if set, zero value otherwise.
func (o *TlsContextDetails) GetInboundSettings() TlsContextDetailsInboundSettings {
	if o == nil || IsNil(o.InboundSettings) {
		var ret TlsContextDetailsInboundSettings
		return ret
	}
	return *o.InboundSettings
}

// GetInboundSettingsOk returns a tuple with the InboundSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetInboundSettingsOk() (*TlsContextDetailsInboundSettings, bool) {
	if o == nil || IsNil(o.InboundSettings) {
		return nil, false
	}
	return o.InboundSettings, true
}

// HasInboundSettings returns a boolean if a field has been set.
func (o *TlsContextDetails) HasInboundSettings() bool {
	if o != nil && !IsNil(o.InboundSettings) {
		return true
	}

	return false
}

// SetInboundSettings gets a reference to the given TlsContextDetailsInboundSettings and assigns it to the InboundSettings field.
func (o *TlsContextDetails) SetInboundSettings(v TlsContextDetailsInboundSettings) {
	o.InboundSettings = &v
}

// GetOutboundSettings returns the OutboundSettings field value if set, zero value otherwise.
func (o *TlsContextDetails) GetOutboundSettings() TlsContextDetailsOutboundSettings {
	if o == nil || IsNil(o.OutboundSettings) {
		var ret TlsContextDetailsOutboundSettings
		return ret
	}
	return *o.OutboundSettings
}

// GetOutboundSettingsOk returns a tuple with the OutboundSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContextDetails) GetOutboundSettingsOk() (*TlsContextDetailsOutboundSettings, bool) {
	if o == nil || IsNil(o.OutboundSettings) {
		return nil, false
	}
	return o.OutboundSettings, true
}

// HasOutboundSettings returns a boolean if a field has been set.
func (o *TlsContextDetails) HasOutboundSettings() bool {
	if o != nil && !IsNil(o.OutboundSettings) {
		return true
	}

	return false
}

// SetOutboundSettings gets a reference to the given TlsContextDetailsOutboundSettings and assigns it to the OutboundSettings field.
func (o *TlsContextDetails) SetOutboundSettings(v TlsContextDetailsOutboundSettings) {
	o.OutboundSettings = &v
}

func (o TlsContextDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TlsContextDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Keystore) {
		toSerialize["keystore"] = o.Keystore
	}
	if !IsNil(o.Truststore) {
		toSerialize["truststore"] = o.Truststore
	}
	if !IsNil(o.EnableMutualAuthentication) {
		toSerialize["enableMutualAuthentication"] = o.EnableMutualAuthentication
	}
	if !IsNil(o.MutualAuthentication) {
		toSerialize["mutualAuthentication"] = o.MutualAuthentication
	}
	if !IsNil(o.AcceptableCipherSuites) {
		toSerialize["acceptableCipherSuites"] = o.AcceptableCipherSuites
	}
	if !IsNil(o.AcceptableTlsVersions) {
		toSerialize["acceptableTlsVersions"] = o.AcceptableTlsVersions
	}
	if !IsNil(o.CipherSuites) {
		toSerialize["cipherSuites"] = o.CipherSuites
	}
	if !IsNil(o.Insecure) {
		toSerialize["insecure"] = o.Insecure
	}
	if !IsNil(o.MinTlsVersion) {
		toSerialize["minTlsVersion"] = o.MinTlsVersion
	}
	if !IsNil(o.MaxTlsVersion) {
		toSerialize["maxTlsVersion"] = o.MaxTlsVersion
	}
	if !IsNil(o.AlpnProtocols) {
		toSerialize["alpnProtocols"] = o.AlpnProtocols
	}
	if !IsNil(o.InboundSettings) {
		toSerialize["inboundSettings"] = o.InboundSettings
	}
	if !IsNil(o.OutboundSettings) {
		toSerialize["outboundSettings"] = o.OutboundSettings
	}
	return toSerialize, nil
}

type NullableTlsContextDetails struct {
	value *TlsContextDetails
	isSet bool
}

func (v NullableTlsContextDetails) Get() *TlsContextDetails {
	return v.value
}

func (v *NullableTlsContextDetails) Set(val *TlsContextDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTlsContextDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTlsContextDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTlsContextDetails(val *TlsContextDetails) *NullableTlsContextDetails {
	return &NullableTlsContextDetails{value: val, isSet: true}
}

func (v NullableTlsContextDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTlsContextDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

