# OpenIDProviderPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**SamlProviderPatchType**](SamlProviderPatchType.md) |  | [optional] 
**OidcProvider** | Pointer to [**OpenIDProviderPatchOidcProvider**](OpenIDProviderPatchOidcProvider.md) |  | [optional] 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 
**LoginDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenIDProviderPatch

`func NewOpenIDProviderPatch() *OpenIDProviderPatch`

NewOpenIDProviderPatch instantiates a new OpenIDProviderPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderPatchWithDefaults

`func NewOpenIDProviderPatchWithDefaults() *OpenIDProviderPatch`

NewOpenIDProviderPatchWithDefaults instantiates a new OpenIDProviderPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OpenIDProviderPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIDProviderPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIDProviderPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenIDProviderPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *OpenIDProviderPatch) GetType() SamlProviderPatchType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenIDProviderPatch) GetTypeOk() (*SamlProviderPatchType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenIDProviderPatch) SetType(v SamlProviderPatchType)`

SetType sets Type field to given value.

### HasType

`func (o *OpenIDProviderPatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOidcProvider

`func (o *OpenIDProviderPatch) GetOidcProvider() OpenIDProviderPatchOidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *OpenIDProviderPatch) GetOidcProviderOk() (*OpenIDProviderPatchOidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *OpenIDProviderPatch) SetOidcProvider(v OpenIDProviderPatchOidcProvider)`

SetOidcProvider sets OidcProvider field to given value.

### HasOidcProvider

`func (o *OpenIDProviderPatch) HasOidcProvider() bool`

HasOidcProvider returns a boolean if a field has been set.

### GetAllowUntrustedCertificates

`func (o *OpenIDProviderPatch) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *OpenIDProviderPatch) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *OpenIDProviderPatch) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *OpenIDProviderPatch) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.

### GetLoginDisabled

`func (o *OpenIDProviderPatch) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *OpenIDProviderPatch) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *OpenIDProviderPatch) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *OpenIDProviderPatch) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


