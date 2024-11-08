# SamlProviderPostBodySaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** |  | 
**Issuer** | **string** |  | 
**PublicKey** | **[]string** |  | 
**ClaimsMapping** | Pointer to [**SamlProviderPostBodySamlClaimsMapping**](SamlProviderPostBodySamlClaimsMapping.md) |  | [optional] 
**IdpInitiatedSsoEnabled** | Pointer to **bool** |  | [optional] 
**SpInitiatedSsoEnabled** | Pointer to **bool** |  | [optional] 
**UseComposerAcsUrl** | Pointer to **bool** |  | [optional] 
**RequireEncryptedSamlAssertions** | Pointer to **bool** |  | [optional] 

## Methods

### NewSamlProviderPostBodySaml

`func NewSamlProviderPostBodySaml(audience string, issuer string, publicKey []string, ) *SamlProviderPostBodySaml`

NewSamlProviderPostBodySaml instantiates a new SamlProviderPostBodySaml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlProviderPostBodySamlWithDefaults

`func NewSamlProviderPostBodySamlWithDefaults() *SamlProviderPostBodySaml`

NewSamlProviderPostBodySamlWithDefaults instantiates a new SamlProviderPostBodySaml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *SamlProviderPostBodySaml) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SamlProviderPostBodySaml) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SamlProviderPostBodySaml) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetIssuer

`func (o *SamlProviderPostBodySaml) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SamlProviderPostBodySaml) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SamlProviderPostBodySaml) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetPublicKey

`func (o *SamlProviderPostBodySaml) GetPublicKey() []string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SamlProviderPostBodySaml) GetPublicKeyOk() (*[]string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SamlProviderPostBodySaml) SetPublicKey(v []string)`

SetPublicKey sets PublicKey field to given value.


### GetClaimsMapping

`func (o *SamlProviderPostBodySaml) GetClaimsMapping() SamlProviderPostBodySamlClaimsMapping`

GetClaimsMapping returns the ClaimsMapping field if non-nil, zero value otherwise.

### GetClaimsMappingOk

`func (o *SamlProviderPostBodySaml) GetClaimsMappingOk() (*SamlProviderPostBodySamlClaimsMapping, bool)`

GetClaimsMappingOk returns a tuple with the ClaimsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMapping

`func (o *SamlProviderPostBodySaml) SetClaimsMapping(v SamlProviderPostBodySamlClaimsMapping)`

SetClaimsMapping sets ClaimsMapping field to given value.

### HasClaimsMapping

`func (o *SamlProviderPostBodySaml) HasClaimsMapping() bool`

HasClaimsMapping returns a boolean if a field has been set.

### GetIdpInitiatedSsoEnabled

`func (o *SamlProviderPostBodySaml) GetIdpInitiatedSsoEnabled() bool`

GetIdpInitiatedSsoEnabled returns the IdpInitiatedSsoEnabled field if non-nil, zero value otherwise.

### GetIdpInitiatedSsoEnabledOk

`func (o *SamlProviderPostBodySaml) GetIdpInitiatedSsoEnabledOk() (*bool, bool)`

GetIdpInitiatedSsoEnabledOk returns a tuple with the IdpInitiatedSsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInitiatedSsoEnabled

`func (o *SamlProviderPostBodySaml) SetIdpInitiatedSsoEnabled(v bool)`

SetIdpInitiatedSsoEnabled sets IdpInitiatedSsoEnabled field to given value.

### HasIdpInitiatedSsoEnabled

`func (o *SamlProviderPostBodySaml) HasIdpInitiatedSsoEnabled() bool`

HasIdpInitiatedSsoEnabled returns a boolean if a field has been set.

### GetSpInitiatedSsoEnabled

`func (o *SamlProviderPostBodySaml) GetSpInitiatedSsoEnabled() bool`

GetSpInitiatedSsoEnabled returns the SpInitiatedSsoEnabled field if non-nil, zero value otherwise.

### GetSpInitiatedSsoEnabledOk

`func (o *SamlProviderPostBodySaml) GetSpInitiatedSsoEnabledOk() (*bool, bool)`

GetSpInitiatedSsoEnabledOk returns a tuple with the SpInitiatedSsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiatedSsoEnabled

`func (o *SamlProviderPostBodySaml) SetSpInitiatedSsoEnabled(v bool)`

SetSpInitiatedSsoEnabled sets SpInitiatedSsoEnabled field to given value.

### HasSpInitiatedSsoEnabled

`func (o *SamlProviderPostBodySaml) HasSpInitiatedSsoEnabled() bool`

HasSpInitiatedSsoEnabled returns a boolean if a field has been set.

### GetUseComposerAcsUrl

`func (o *SamlProviderPostBodySaml) GetUseComposerAcsUrl() bool`

GetUseComposerAcsUrl returns the UseComposerAcsUrl field if non-nil, zero value otherwise.

### GetUseComposerAcsUrlOk

`func (o *SamlProviderPostBodySaml) GetUseComposerAcsUrlOk() (*bool, bool)`

GetUseComposerAcsUrlOk returns a tuple with the UseComposerAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseComposerAcsUrl

`func (o *SamlProviderPostBodySaml) SetUseComposerAcsUrl(v bool)`

SetUseComposerAcsUrl sets UseComposerAcsUrl field to given value.

### HasUseComposerAcsUrl

`func (o *SamlProviderPostBodySaml) HasUseComposerAcsUrl() bool`

HasUseComposerAcsUrl returns a boolean if a field has been set.

### GetRequireEncryptedSamlAssertions

`func (o *SamlProviderPostBodySaml) GetRequireEncryptedSamlAssertions() bool`

GetRequireEncryptedSamlAssertions returns the RequireEncryptedSamlAssertions field if non-nil, zero value otherwise.

### GetRequireEncryptedSamlAssertionsOk

`func (o *SamlProviderPostBodySaml) GetRequireEncryptedSamlAssertionsOk() (*bool, bool)`

GetRequireEncryptedSamlAssertionsOk returns a tuple with the RequireEncryptedSamlAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEncryptedSamlAssertions

`func (o *SamlProviderPostBodySaml) SetRequireEncryptedSamlAssertions(v bool)`

SetRequireEncryptedSamlAssertions sets RequireEncryptedSamlAssertions field to given value.

### HasRequireEncryptedSamlAssertions

`func (o *SamlProviderPostBodySaml) HasRequireEncryptedSamlAssertions() bool`

HasRequireEncryptedSamlAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


