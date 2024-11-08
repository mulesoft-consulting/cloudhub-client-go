# ConnectedAppRespExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**PublicKeys** | Pointer to **[]string** | Application public key (PEM format). Used to validate JWT authorization grants. | [optional] 
**ClientUri** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** | the organization where the connected app is created | [optional] 
**ClientId** | Pointer to **string** | connected app client id | [optional] 
**ClientSecret** | Pointer to **string** | connected app generated secret | [optional] 
**OwnerOrgId** | Pointer to **string** | connected app owner organization id | [optional] 
**OwnerUserId** | Pointer to **string** | connected app owner user id | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**PolicyUri** | Pointer to **string** |  | [optional] 
**TosUri** | Pointer to **string** |  | [optional] 
**CertExpiry** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectedAppRespExt

`func NewConnectedAppRespExt() *ConnectedAppRespExt`

NewConnectedAppRespExt instantiates a new ConnectedAppRespExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedAppRespExtWithDefaults

`func NewConnectedAppRespExtWithDefaults() *ConnectedAppRespExt`

NewConnectedAppRespExtWithDefaults instantiates a new ConnectedAppRespExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ConnectedAppRespExt) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ConnectedAppRespExt) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ConnectedAppRespExt) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ConnectedAppRespExt) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetGrantTypes

`func (o *ConnectedAppRespExt) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ConnectedAppRespExt) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ConnectedAppRespExt) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *ConnectedAppRespExt) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ConnectedAppRespExt) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ConnectedAppRespExt) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ConnectedAppRespExt) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ConnectedAppRespExt) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *ConnectedAppRespExt) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ConnectedAppRespExt) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ConnectedAppRespExt) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ConnectedAppRespExt) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPublicKeys

`func (o *ConnectedAppRespExt) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *ConnectedAppRespExt) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *ConnectedAppRespExt) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *ConnectedAppRespExt) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.

### GetClientUri

`func (o *ConnectedAppRespExt) GetClientUri() string`

GetClientUri returns the ClientUri field if non-nil, zero value otherwise.

### GetClientUriOk

`func (o *ConnectedAppRespExt) GetClientUriOk() (*string, bool)`

GetClientUriOk returns a tuple with the ClientUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUri

`func (o *ConnectedAppRespExt) SetClientUri(v string)`

SetClientUri sets ClientUri field to given value.

### HasClientUri

`func (o *ConnectedAppRespExt) HasClientUri() bool`

HasClientUri returns a boolean if a field has been set.

### GetAudience

`func (o *ConnectedAppRespExt) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ConnectedAppRespExt) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ConnectedAppRespExt) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *ConnectedAppRespExt) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetOrgId

`func (o *ConnectedAppRespExt) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConnectedAppRespExt) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConnectedAppRespExt) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ConnectedAppRespExt) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetClientId

`func (o *ConnectedAppRespExt) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConnectedAppRespExt) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConnectedAppRespExt) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ConnectedAppRespExt) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ConnectedAppRespExt) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ConnectedAppRespExt) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ConnectedAppRespExt) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ConnectedAppRespExt) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetOwnerOrgId

`func (o *ConnectedAppRespExt) GetOwnerOrgId() string`

GetOwnerOrgId returns the OwnerOrgId field if non-nil, zero value otherwise.

### GetOwnerOrgIdOk

`func (o *ConnectedAppRespExt) GetOwnerOrgIdOk() (*string, bool)`

GetOwnerOrgIdOk returns a tuple with the OwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrgId

`func (o *ConnectedAppRespExt) SetOwnerOrgId(v string)`

SetOwnerOrgId sets OwnerOrgId field to given value.

### HasOwnerOrgId

`func (o *ConnectedAppRespExt) HasOwnerOrgId() bool`

HasOwnerOrgId returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *ConnectedAppRespExt) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *ConnectedAppRespExt) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *ConnectedAppRespExt) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *ConnectedAppRespExt) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetEnabled

`func (o *ConnectedAppRespExt) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectedAppRespExt) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectedAppRespExt) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConnectedAppRespExt) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPolicyUri

`func (o *ConnectedAppRespExt) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *ConnectedAppRespExt) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *ConnectedAppRespExt) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *ConnectedAppRespExt) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetTosUri

`func (o *ConnectedAppRespExt) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *ConnectedAppRespExt) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *ConnectedAppRespExt) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *ConnectedAppRespExt) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetCertExpiry

`func (o *ConnectedAppRespExt) GetCertExpiry() string`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *ConnectedAppRespExt) GetCertExpiryOk() (*string, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *ConnectedAppRespExt) SetCertExpiry(v string)`

SetCertExpiry sets CertExpiry field to given value.

### HasCertExpiry

`func (o *ConnectedAppRespExt) HasCertExpiry() bool`

HasCertExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


