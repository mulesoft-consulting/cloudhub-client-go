# LdapProviderPostBodyConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binddn** | **string** |  | 
**ConnectTimeoutSeconds** | **float32** |  | 
**Host** | **string** |  | 
**OperationTimeoutMs** | **float32** |  | 
**Password** | **string** |  | 
**Port** | **float32** |  | 
**Validatecert** | Pointer to **bool** |  | [optional] 

## Methods

### NewLdapProviderPostBodyConnection

`func NewLdapProviderPostBodyConnection(binddn string, connectTimeoutSeconds float32, host string, operationTimeoutMs float32, password string, port float32, ) *LdapProviderPostBodyConnection`

NewLdapProviderPostBodyConnection instantiates a new LdapProviderPostBodyConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapProviderPostBodyConnectionWithDefaults

`func NewLdapProviderPostBodyConnectionWithDefaults() *LdapProviderPostBodyConnection`

NewLdapProviderPostBodyConnectionWithDefaults instantiates a new LdapProviderPostBodyConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinddn

`func (o *LdapProviderPostBodyConnection) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *LdapProviderPostBodyConnection) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *LdapProviderPostBodyConnection) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.


### GetConnectTimeoutSeconds

`func (o *LdapProviderPostBodyConnection) GetConnectTimeoutSeconds() float32`

GetConnectTimeoutSeconds returns the ConnectTimeoutSeconds field if non-nil, zero value otherwise.

### GetConnectTimeoutSecondsOk

`func (o *LdapProviderPostBodyConnection) GetConnectTimeoutSecondsOk() (*float32, bool)`

GetConnectTimeoutSecondsOk returns a tuple with the ConnectTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeoutSeconds

`func (o *LdapProviderPostBodyConnection) SetConnectTimeoutSeconds(v float32)`

SetConnectTimeoutSeconds sets ConnectTimeoutSeconds field to given value.


### GetHost

`func (o *LdapProviderPostBodyConnection) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LdapProviderPostBodyConnection) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LdapProviderPostBodyConnection) SetHost(v string)`

SetHost sets Host field to given value.


### GetOperationTimeoutMs

`func (o *LdapProviderPostBodyConnection) GetOperationTimeoutMs() float32`

GetOperationTimeoutMs returns the OperationTimeoutMs field if non-nil, zero value otherwise.

### GetOperationTimeoutMsOk

`func (o *LdapProviderPostBodyConnection) GetOperationTimeoutMsOk() (*float32, bool)`

GetOperationTimeoutMsOk returns a tuple with the OperationTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTimeoutMs

`func (o *LdapProviderPostBodyConnection) SetOperationTimeoutMs(v float32)`

SetOperationTimeoutMs sets OperationTimeoutMs field to given value.


### GetPassword

`func (o *LdapProviderPostBodyConnection) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LdapProviderPostBodyConnection) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LdapProviderPostBodyConnection) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPort

`func (o *LdapProviderPostBodyConnection) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LdapProviderPostBodyConnection) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LdapProviderPostBodyConnection) SetPort(v float32)`

SetPort sets Port field to given value.


### GetValidatecert

`func (o *LdapProviderPostBodyConnection) GetValidatecert() bool`

GetValidatecert returns the Validatecert field if non-nil, zero value otherwise.

### GetValidatecertOk

`func (o *LdapProviderPostBodyConnection) GetValidatecertOk() (*bool, bool)`

GetValidatecertOk returns a tuple with the Validatecert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatecert

`func (o *LdapProviderPostBodyConnection) SetValidatecert(v bool)`

SetValidatecert sets Validatecert field to given value.

### HasValidatecert

`func (o *LdapProviderPostBodyConnection) HasValidatecert() bool`

HasValidatecert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


