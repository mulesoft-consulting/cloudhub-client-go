# OpenIDProviderGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OidcProvider** | [**OpenIDProviderGetOidcProvider**](OpenIDProviderGetOidcProvider.md) |  | 
**Name** | **string** |  | 
**OrgId** | **string** |  | 
**ServiceProvider** | [**OpenIDProviderGetServiceProvider**](OpenIDProviderGetServiceProvider.md) |  | 
**LoginDisabled** | Pointer to **bool** |  | [optional] 
**AllowUntrustedCertificates** | **bool** |  | 
**Type** | [**OpenIDProviderManualPostBodyType**](OpenIDProviderManualPostBodyType.md) |  | 
**ProviderId** | **string** |  | 

## Methods

### NewOpenIDProviderGet

`func NewOpenIDProviderGet(oidcProvider OpenIDProviderGetOidcProvider, name string, orgId string, serviceProvider OpenIDProviderGetServiceProvider, allowUntrustedCertificates bool, type_ OpenIDProviderManualPostBodyType, providerId string, ) *OpenIDProviderGet`

NewOpenIDProviderGet instantiates a new OpenIDProviderGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderGetWithDefaults

`func NewOpenIDProviderGetWithDefaults() *OpenIDProviderGet`

NewOpenIDProviderGetWithDefaults instantiates a new OpenIDProviderGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidcProvider

`func (o *OpenIDProviderGet) GetOidcProvider() OpenIDProviderGetOidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *OpenIDProviderGet) GetOidcProviderOk() (*OpenIDProviderGetOidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *OpenIDProviderGet) SetOidcProvider(v OpenIDProviderGetOidcProvider)`

SetOidcProvider sets OidcProvider field to given value.


### GetName

`func (o *OpenIDProviderGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIDProviderGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIDProviderGet) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *OpenIDProviderGet) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OpenIDProviderGet) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OpenIDProviderGet) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetServiceProvider

`func (o *OpenIDProviderGet) GetServiceProvider() OpenIDProviderGetServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *OpenIDProviderGet) GetServiceProviderOk() (*OpenIDProviderGetServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *OpenIDProviderGet) SetServiceProvider(v OpenIDProviderGetServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.


### GetLoginDisabled

`func (o *OpenIDProviderGet) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *OpenIDProviderGet) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *OpenIDProviderGet) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *OpenIDProviderGet) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.

### GetAllowUntrustedCertificates

`func (o *OpenIDProviderGet) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *OpenIDProviderGet) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *OpenIDProviderGet) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.


### GetType

`func (o *OpenIDProviderGet) GetType() OpenIDProviderManualPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenIDProviderGet) GetTypeOk() (*OpenIDProviderManualPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenIDProviderGet) SetType(v OpenIDProviderManualPostBodyType)`

SetType sets Type field to given value.


### GetProviderId

`func (o *OpenIDProviderGet) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *OpenIDProviderGet) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *OpenIDProviderGet) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


