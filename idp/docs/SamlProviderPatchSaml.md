# SamlProviderPatchSaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **[]string** |  | [optional] 
**ClaimsMapping** | Pointer to [**SamlProviderPostBodySamlClaimsMapping**](SamlProviderPostBodySamlClaimsMapping.md) |  | [optional] 
**IdpInitiatedSsoEnabled** | Pointer to **bool** |  | [optional] 
**SpInitiatedSsoEnabled** | Pointer to **bool** |  | [optional] 
**UseComposerAcsUrl** | Pointer to **bool** |  | [optional] 
**RequireEncryptedSamlAssertions** | Pointer to **bool** |  | [optional] 

## Methods

### NewSamlProviderPatchSaml

`func NewSamlProviderPatchSaml() *SamlProviderPatchSaml`

NewSamlProviderPatchSaml instantiates a new SamlProviderPatchSaml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlProviderPatchSamlWithDefaults

`func NewSamlProviderPatchSamlWithDefaults() *SamlProviderPatchSaml`

NewSamlProviderPatchSamlWithDefaults instantiates a new SamlProviderPatchSaml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *SamlProviderPatchSaml) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SamlProviderPatchSaml) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SamlProviderPatchSaml) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SamlProviderPatchSaml) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *SamlProviderPatchSaml) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SamlProviderPatchSaml) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SamlProviderPatchSaml) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SamlProviderPatchSaml) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetPublicKey

`func (o *SamlProviderPatchSaml) GetPublicKey() []string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SamlProviderPatchSaml) GetPublicKeyOk() (*[]string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SamlProviderPatchSaml) SetPublicKey(v []string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SamlProviderPatchSaml) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetClaimsMapping

`func (o *SamlProviderPatchSaml) GetClaimsMapping() SamlProviderPostBodySamlClaimsMapping`

GetClaimsMapping returns the ClaimsMapping field if non-nil, zero value otherwise.

### GetClaimsMappingOk

`func (o *SamlProviderPatchSaml) GetClaimsMappingOk() (*SamlProviderPostBodySamlClaimsMapping, bool)`

GetClaimsMappingOk returns a tuple with the ClaimsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMapping

`func (o *SamlProviderPatchSaml) SetClaimsMapping(v SamlProviderPostBodySamlClaimsMapping)`

SetClaimsMapping sets ClaimsMapping field to given value.

### HasClaimsMapping

`func (o *SamlProviderPatchSaml) HasClaimsMapping() bool`

HasClaimsMapping returns a boolean if a field has been set.

### GetIdpInitiatedSsoEnabled

`func (o *SamlProviderPatchSaml) GetIdpInitiatedSsoEnabled() bool`

GetIdpInitiatedSsoEnabled returns the IdpInitiatedSsoEnabled field if non-nil, zero value otherwise.

### GetIdpInitiatedSsoEnabledOk

`func (o *SamlProviderPatchSaml) GetIdpInitiatedSsoEnabledOk() (*bool, bool)`

GetIdpInitiatedSsoEnabledOk returns a tuple with the IdpInitiatedSsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInitiatedSsoEnabled

`func (o *SamlProviderPatchSaml) SetIdpInitiatedSsoEnabled(v bool)`

SetIdpInitiatedSsoEnabled sets IdpInitiatedSsoEnabled field to given value.

### HasIdpInitiatedSsoEnabled

`func (o *SamlProviderPatchSaml) HasIdpInitiatedSsoEnabled() bool`

HasIdpInitiatedSsoEnabled returns a boolean if a field has been set.

### GetSpInitiatedSsoEnabled

`func (o *SamlProviderPatchSaml) GetSpInitiatedSsoEnabled() bool`

GetSpInitiatedSsoEnabled returns the SpInitiatedSsoEnabled field if non-nil, zero value otherwise.

### GetSpInitiatedSsoEnabledOk

`func (o *SamlProviderPatchSaml) GetSpInitiatedSsoEnabledOk() (*bool, bool)`

GetSpInitiatedSsoEnabledOk returns a tuple with the SpInitiatedSsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiatedSsoEnabled

`func (o *SamlProviderPatchSaml) SetSpInitiatedSsoEnabled(v bool)`

SetSpInitiatedSsoEnabled sets SpInitiatedSsoEnabled field to given value.

### HasSpInitiatedSsoEnabled

`func (o *SamlProviderPatchSaml) HasSpInitiatedSsoEnabled() bool`

HasSpInitiatedSsoEnabled returns a boolean if a field has been set.

### GetUseComposerAcsUrl

`func (o *SamlProviderPatchSaml) GetUseComposerAcsUrl() bool`

GetUseComposerAcsUrl returns the UseComposerAcsUrl field if non-nil, zero value otherwise.

### GetUseComposerAcsUrlOk

`func (o *SamlProviderPatchSaml) GetUseComposerAcsUrlOk() (*bool, bool)`

GetUseComposerAcsUrlOk returns a tuple with the UseComposerAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseComposerAcsUrl

`func (o *SamlProviderPatchSaml) SetUseComposerAcsUrl(v bool)`

SetUseComposerAcsUrl sets UseComposerAcsUrl field to given value.

### HasUseComposerAcsUrl

`func (o *SamlProviderPatchSaml) HasUseComposerAcsUrl() bool`

HasUseComposerAcsUrl returns a boolean if a field has been set.

### GetRequireEncryptedSamlAssertions

`func (o *SamlProviderPatchSaml) GetRequireEncryptedSamlAssertions() bool`

GetRequireEncryptedSamlAssertions returns the RequireEncryptedSamlAssertions field if non-nil, zero value otherwise.

### GetRequireEncryptedSamlAssertionsOk

`func (o *SamlProviderPatchSaml) GetRequireEncryptedSamlAssertionsOk() (*bool, bool)`

GetRequireEncryptedSamlAssertionsOk returns a tuple with the RequireEncryptedSamlAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEncryptedSamlAssertions

`func (o *SamlProviderPatchSaml) SetRequireEncryptedSamlAssertions(v bool)`

SetRequireEncryptedSamlAssertions sets RequireEncryptedSamlAssertions field to given value.

### HasRequireEncryptedSamlAssertions

`func (o *SamlProviderPatchSaml) HasRequireEncryptedSamlAssertions() bool`

HasRequireEncryptedSamlAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


