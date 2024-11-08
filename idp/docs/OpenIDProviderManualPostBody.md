# OpenIDProviderManualPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ArcNamespace** | Pointer to **string** |  | [optional] 
**Type** | [**OpenIDProviderManualPostBodyType**](OpenIDProviderManualPostBodyType.md) |  | 
**OidcProvider** | [**OpenIDProviderManualPostBodyOidcProvider**](OpenIDProviderManualPostBodyOidcProvider.md) |  | 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 
**LoginDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenIDProviderManualPostBody

`func NewOpenIDProviderManualPostBody(name string, type_ OpenIDProviderManualPostBodyType, oidcProvider OpenIDProviderManualPostBodyOidcProvider, ) *OpenIDProviderManualPostBody`

NewOpenIDProviderManualPostBody instantiates a new OpenIDProviderManualPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderManualPostBodyWithDefaults

`func NewOpenIDProviderManualPostBodyWithDefaults() *OpenIDProviderManualPostBody`

NewOpenIDProviderManualPostBodyWithDefaults instantiates a new OpenIDProviderManualPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OpenIDProviderManualPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIDProviderManualPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIDProviderManualPostBody) SetName(v string)`

SetName sets Name field to given value.


### GetArcNamespace

`func (o *OpenIDProviderManualPostBody) GetArcNamespace() string`

GetArcNamespace returns the ArcNamespace field if non-nil, zero value otherwise.

### GetArcNamespaceOk

`func (o *OpenIDProviderManualPostBody) GetArcNamespaceOk() (*string, bool)`

GetArcNamespaceOk returns a tuple with the ArcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcNamespace

`func (o *OpenIDProviderManualPostBody) SetArcNamespace(v string)`

SetArcNamespace sets ArcNamespace field to given value.

### HasArcNamespace

`func (o *OpenIDProviderManualPostBody) HasArcNamespace() bool`

HasArcNamespace returns a boolean if a field has been set.

### GetType

`func (o *OpenIDProviderManualPostBody) GetType() OpenIDProviderManualPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenIDProviderManualPostBody) GetTypeOk() (*OpenIDProviderManualPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenIDProviderManualPostBody) SetType(v OpenIDProviderManualPostBodyType)`

SetType sets Type field to given value.


### GetOidcProvider

`func (o *OpenIDProviderManualPostBody) GetOidcProvider() OpenIDProviderManualPostBodyOidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *OpenIDProviderManualPostBody) GetOidcProviderOk() (*OpenIDProviderManualPostBodyOidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *OpenIDProviderManualPostBody) SetOidcProvider(v OpenIDProviderManualPostBodyOidcProvider)`

SetOidcProvider sets OidcProvider field to given value.


### GetAllowUntrustedCertificates

`func (o *OpenIDProviderManualPostBody) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *OpenIDProviderManualPostBody) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *OpenIDProviderManualPostBody) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *OpenIDProviderManualPostBody) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.

### GetLoginDisabled

`func (o *OpenIDProviderManualPostBody) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *OpenIDProviderManualPostBody) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *OpenIDProviderManualPostBody) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *OpenIDProviderManualPostBody) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


