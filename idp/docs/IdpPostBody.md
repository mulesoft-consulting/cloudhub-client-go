# IdpPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcNamespace** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Type** | [**OpenIDProviderManualPostBodyType**](OpenIDProviderManualPostBodyType.md) |  | 
**ServiceProvider** | [**SamlProviderPostBodyServiceProvider**](SamlProviderPostBodyServiceProvider.md) |  | 
**Saml** | [**SamlProviderPostBodySaml**](SamlProviderPostBodySaml.md) |  | 
**LoginDisabled** | Pointer to **bool** |  | [optional] 
**Connection** | [**LdapProviderPostBodyConnection**](LdapProviderPostBodyConnection.md) |  | 
**SearchBases** | [**LdapProviderPostBodySearchBases**](LdapProviderPostBodySearchBases.md) |  | 
**Dns** | [**LdapProviderPostBodySearchBases**](LdapProviderPostBodySearchBases.md) |  | 
**Filters** | [**LdapProviderPostBodyFilters**](LdapProviderPostBodyFilters.md) |  | 
**UserMapping** | [**LdapProviderPostBodyUserMapping**](LdapProviderPostBodyUserMapping.md) |  | 
**GroupMapping** | [**LdapProviderPostBodyGroupMapping**](LdapProviderPostBodyGroupMapping.md) |  | 
**OidcProvider** | [**OpenIDProviderDynamicPostBodyOidcProvider**](OpenIDProviderDynamicPostBodyOidcProvider.md) |  | 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 

## Methods

### NewIdpPostBody

`func NewIdpPostBody(name string, type_ OpenIDProviderManualPostBodyType, serviceProvider SamlProviderPostBodyServiceProvider, saml SamlProviderPostBodySaml, connection LdapProviderPostBodyConnection, searchBases LdapProviderPostBodySearchBases, dns LdapProviderPostBodySearchBases, filters LdapProviderPostBodyFilters, userMapping LdapProviderPostBodyUserMapping, groupMapping LdapProviderPostBodyGroupMapping, oidcProvider OpenIDProviderDynamicPostBodyOidcProvider, ) *IdpPostBody`

NewIdpPostBody instantiates a new IdpPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPostBodyWithDefaults

`func NewIdpPostBodyWithDefaults() *IdpPostBody`

NewIdpPostBodyWithDefaults instantiates a new IdpPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcNamespace

`func (o *IdpPostBody) GetArcNamespace() string`

GetArcNamespace returns the ArcNamespace field if non-nil, zero value otherwise.

### GetArcNamespaceOk

`func (o *IdpPostBody) GetArcNamespaceOk() (*string, bool)`

GetArcNamespaceOk returns a tuple with the ArcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcNamespace

`func (o *IdpPostBody) SetArcNamespace(v string)`

SetArcNamespace sets ArcNamespace field to given value.

### HasArcNamespace

`func (o *IdpPostBody) HasArcNamespace() bool`

HasArcNamespace returns a boolean if a field has been set.

### GetName

`func (o *IdpPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpPostBody) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IdpPostBody) GetType() OpenIDProviderManualPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpPostBody) GetTypeOk() (*OpenIDProviderManualPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpPostBody) SetType(v OpenIDProviderManualPostBodyType)`

SetType sets Type field to given value.


### GetServiceProvider

`func (o *IdpPostBody) GetServiceProvider() SamlProviderPostBodyServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *IdpPostBody) GetServiceProviderOk() (*SamlProviderPostBodyServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *IdpPostBody) SetServiceProvider(v SamlProviderPostBodyServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.


### GetSaml

`func (o *IdpPostBody) GetSaml() SamlProviderPostBodySaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *IdpPostBody) GetSamlOk() (*SamlProviderPostBodySaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *IdpPostBody) SetSaml(v SamlProviderPostBodySaml)`

SetSaml sets Saml field to given value.


### GetLoginDisabled

`func (o *IdpPostBody) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *IdpPostBody) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *IdpPostBody) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *IdpPostBody) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.

### GetConnection

`func (o *IdpPostBody) GetConnection() LdapProviderPostBodyConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *IdpPostBody) GetConnectionOk() (*LdapProviderPostBodyConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *IdpPostBody) SetConnection(v LdapProviderPostBodyConnection)`

SetConnection sets Connection field to given value.


### GetSearchBases

`func (o *IdpPostBody) GetSearchBases() LdapProviderPostBodySearchBases`

GetSearchBases returns the SearchBases field if non-nil, zero value otherwise.

### GetSearchBasesOk

`func (o *IdpPostBody) GetSearchBasesOk() (*LdapProviderPostBodySearchBases, bool)`

GetSearchBasesOk returns a tuple with the SearchBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBases

`func (o *IdpPostBody) SetSearchBases(v LdapProviderPostBodySearchBases)`

SetSearchBases sets SearchBases field to given value.


### GetDns

`func (o *IdpPostBody) GetDns() LdapProviderPostBodySearchBases`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *IdpPostBody) GetDnsOk() (*LdapProviderPostBodySearchBases, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *IdpPostBody) SetDns(v LdapProviderPostBodySearchBases)`

SetDns sets Dns field to given value.


### GetFilters

`func (o *IdpPostBody) GetFilters() LdapProviderPostBodyFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *IdpPostBody) GetFiltersOk() (*LdapProviderPostBodyFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *IdpPostBody) SetFilters(v LdapProviderPostBodyFilters)`

SetFilters sets Filters field to given value.


### GetUserMapping

`func (o *IdpPostBody) GetUserMapping() LdapProviderPostBodyUserMapping`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *IdpPostBody) GetUserMappingOk() (*LdapProviderPostBodyUserMapping, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *IdpPostBody) SetUserMapping(v LdapProviderPostBodyUserMapping)`

SetUserMapping sets UserMapping field to given value.


### GetGroupMapping

`func (o *IdpPostBody) GetGroupMapping() LdapProviderPostBodyGroupMapping`

GetGroupMapping returns the GroupMapping field if non-nil, zero value otherwise.

### GetGroupMappingOk

`func (o *IdpPostBody) GetGroupMappingOk() (*LdapProviderPostBodyGroupMapping, bool)`

GetGroupMappingOk returns a tuple with the GroupMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMapping

`func (o *IdpPostBody) SetGroupMapping(v LdapProviderPostBodyGroupMapping)`

SetGroupMapping sets GroupMapping field to given value.


### GetOidcProvider

`func (o *IdpPostBody) GetOidcProvider() OpenIDProviderDynamicPostBodyOidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *IdpPostBody) GetOidcProviderOk() (*OpenIDProviderDynamicPostBodyOidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *IdpPostBody) SetOidcProvider(v OpenIDProviderDynamicPostBodyOidcProvider)`

SetOidcProvider sets OidcProvider field to given value.


### GetAllowUntrustedCertificates

`func (o *IdpPostBody) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *IdpPostBody) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *IdpPostBody) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *IdpPostBody) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


