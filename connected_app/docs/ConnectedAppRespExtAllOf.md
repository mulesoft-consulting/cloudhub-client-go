# ConnectedAppRespExtAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewConnectedAppRespExtAllOf

`func NewConnectedAppRespExtAllOf() *ConnectedAppRespExtAllOf`

NewConnectedAppRespExtAllOf instantiates a new ConnectedAppRespExtAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedAppRespExtAllOfWithDefaults

`func NewConnectedAppRespExtAllOfWithDefaults() *ConnectedAppRespExtAllOf`

NewConnectedAppRespExtAllOfWithDefaults instantiates a new ConnectedAppRespExtAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ConnectedAppRespExtAllOf) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConnectedAppRespExtAllOf) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConnectedAppRespExtAllOf) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ConnectedAppRespExtAllOf) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetClientId

`func (o *ConnectedAppRespExtAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConnectedAppRespExtAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConnectedAppRespExtAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ConnectedAppRespExtAllOf) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ConnectedAppRespExtAllOf) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ConnectedAppRespExtAllOf) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ConnectedAppRespExtAllOf) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ConnectedAppRespExtAllOf) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetOwnerOrgId

`func (o *ConnectedAppRespExtAllOf) GetOwnerOrgId() string`

GetOwnerOrgId returns the OwnerOrgId field if non-nil, zero value otherwise.

### GetOwnerOrgIdOk

`func (o *ConnectedAppRespExtAllOf) GetOwnerOrgIdOk() (*string, bool)`

GetOwnerOrgIdOk returns a tuple with the OwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrgId

`func (o *ConnectedAppRespExtAllOf) SetOwnerOrgId(v string)`

SetOwnerOrgId sets OwnerOrgId field to given value.

### HasOwnerOrgId

`func (o *ConnectedAppRespExtAllOf) HasOwnerOrgId() bool`

HasOwnerOrgId returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *ConnectedAppRespExtAllOf) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *ConnectedAppRespExtAllOf) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *ConnectedAppRespExtAllOf) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *ConnectedAppRespExtAllOf) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetEnabled

`func (o *ConnectedAppRespExtAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectedAppRespExtAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectedAppRespExtAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConnectedAppRespExtAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPolicyUri

`func (o *ConnectedAppRespExtAllOf) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *ConnectedAppRespExtAllOf) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *ConnectedAppRespExtAllOf) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *ConnectedAppRespExtAllOf) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetTosUri

`func (o *ConnectedAppRespExtAllOf) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *ConnectedAppRespExtAllOf) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *ConnectedAppRespExtAllOf) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *ConnectedAppRespExtAllOf) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetCertExpiry

`func (o *ConnectedAppRespExtAllOf) GetCertExpiry() string`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *ConnectedAppRespExtAllOf) GetCertExpiryOk() (*string, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *ConnectedAppRespExtAllOf) SetCertExpiry(v string)`

SetCertExpiry sets CertExpiry field to given value.

### HasCertExpiry

`func (o *ConnectedAppRespExtAllOf) HasCertExpiry() bool`

HasCertExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


