# OpenIDProviderDynamicPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ArcNamespace** | Pointer to **string** |  | [optional] 
**Type** | [**OpenIDProviderManualPostBodyType**](OpenIDProviderManualPostBodyType.md) |  | 
**OidcProvider** | [**OpenIDProviderDynamicPostBodyOidcProvider**](OpenIDProviderDynamicPostBodyOidcProvider.md) |  | 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 
**LoginDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenIDProviderDynamicPostBody

`func NewOpenIDProviderDynamicPostBody(name string, type_ OpenIDProviderManualPostBodyType, oidcProvider OpenIDProviderDynamicPostBodyOidcProvider, ) *OpenIDProviderDynamicPostBody`

NewOpenIDProviderDynamicPostBody instantiates a new OpenIDProviderDynamicPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderDynamicPostBodyWithDefaults

`func NewOpenIDProviderDynamicPostBodyWithDefaults() *OpenIDProviderDynamicPostBody`

NewOpenIDProviderDynamicPostBodyWithDefaults instantiates a new OpenIDProviderDynamicPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OpenIDProviderDynamicPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIDProviderDynamicPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIDProviderDynamicPostBody) SetName(v string)`

SetName sets Name field to given value.


### GetArcNamespace

`func (o *OpenIDProviderDynamicPostBody) GetArcNamespace() string`

GetArcNamespace returns the ArcNamespace field if non-nil, zero value otherwise.

### GetArcNamespaceOk

`func (o *OpenIDProviderDynamicPostBody) GetArcNamespaceOk() (*string, bool)`

GetArcNamespaceOk returns a tuple with the ArcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcNamespace

`func (o *OpenIDProviderDynamicPostBody) SetArcNamespace(v string)`

SetArcNamespace sets ArcNamespace field to given value.

### HasArcNamespace

`func (o *OpenIDProviderDynamicPostBody) HasArcNamespace() bool`

HasArcNamespace returns a boolean if a field has been set.

### GetType

`func (o *OpenIDProviderDynamicPostBody) GetType() OpenIDProviderManualPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenIDProviderDynamicPostBody) GetTypeOk() (*OpenIDProviderManualPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenIDProviderDynamicPostBody) SetType(v OpenIDProviderManualPostBodyType)`

SetType sets Type field to given value.


### GetOidcProvider

`func (o *OpenIDProviderDynamicPostBody) GetOidcProvider() OpenIDProviderDynamicPostBodyOidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *OpenIDProviderDynamicPostBody) GetOidcProviderOk() (*OpenIDProviderDynamicPostBodyOidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *OpenIDProviderDynamicPostBody) SetOidcProvider(v OpenIDProviderDynamicPostBodyOidcProvider)`

SetOidcProvider sets OidcProvider field to given value.


### GetAllowUntrustedCertificates

`func (o *OpenIDProviderDynamicPostBody) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *OpenIDProviderDynamicPostBody) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *OpenIDProviderDynamicPostBody) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *OpenIDProviderDynamicPostBody) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.

### GetLoginDisabled

`func (o *OpenIDProviderDynamicPostBody) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *OpenIDProviderDynamicPostBody) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *OpenIDProviderDynamicPostBody) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *OpenIDProviderDynamicPostBody) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


