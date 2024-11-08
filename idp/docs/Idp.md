# Idp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OrgId** | **string** |  | 
**Saml** | [**SamlProviderPostBodySaml**](SamlProviderPostBodySaml.md) |  | 
**ServiceProvider** | [**OpenIDProviderGetServiceProvider**](OpenIDProviderGetServiceProvider.md) |  | 
**LoginDisabled** | Pointer to **bool** |  | [optional] 
**ArcNamespace** | Pointer to **string** |  | [optional] 
**ProviderId** | **string** |  | 
**Type** | [**OpenIDProviderManualPostBodyType**](OpenIDProviderManualPostBodyType.md) |  | 
**UserMapping** | [**LdapProviderPostBodyUserMapping**](LdapProviderPostBodyUserMapping.md) |  | 
**Filters** | [**LdapProviderPostBodyFilters**](LdapProviderPostBodyFilters.md) |  | 
**Dns** | [**LdapProviderPostBodySearchBases**](LdapProviderPostBodySearchBases.md) |  | 
**GroupMapping** | [**LdapProviderPostBodyGroupMapping**](LdapProviderPostBodyGroupMapping.md) |  | 
**Connection** | [**LdapProviderPostBodyConnection**](LdapProviderPostBodyConnection.md) |  | 
**SearchBases** | [**LdapProviderPostBodySearchBases**](LdapProviderPostBodySearchBases.md) |  | 
**OidcProvider** | [**OpenIDProviderGetOidcProvider**](OpenIDProviderGetOidcProvider.md) |  | 
**AllowUntrustedCertificates** | **bool** |  | 

## Methods

### NewIdp

`func NewIdp(name string, orgId string, saml SamlProviderPostBodySaml, serviceProvider OpenIDProviderGetServiceProvider, providerId string, type_ OpenIDProviderManualPostBodyType, userMapping LdapProviderPostBodyUserMapping, filters LdapProviderPostBodyFilters, dns LdapProviderPostBodySearchBases, groupMapping LdapProviderPostBodyGroupMapping, connection LdapProviderPostBodyConnection, searchBases LdapProviderPostBodySearchBases, oidcProvider OpenIDProviderGetOidcProvider, allowUntrustedCertificates bool, ) *Idp`

NewIdp instantiates a new Idp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpWithDefaults

`func NewIdpWithDefaults() *Idp`

NewIdpWithDefaults instantiates a new Idp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Idp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Idp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Idp) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *Idp) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Idp) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Idp) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSaml

`func (o *Idp) GetSaml() SamlProviderPostBodySaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *Idp) GetSamlOk() (*SamlProviderPostBodySaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *Idp) SetSaml(v SamlProviderPostBodySaml)`

SetSaml sets Saml field to given value.


### GetServiceProvider

`func (o *Idp) GetServiceProvider() OpenIDProviderGetServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *Idp) GetServiceProviderOk() (*OpenIDProviderGetServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *Idp) SetServiceProvider(v OpenIDProviderGetServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.


### GetLoginDisabled

`func (o *Idp) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *Idp) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *Idp) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *Idp) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.

### GetArcNamespace

`func (o *Idp) GetArcNamespace() string`

GetArcNamespace returns the ArcNamespace field if non-nil, zero value otherwise.

### GetArcNamespaceOk

`func (o *Idp) GetArcNamespaceOk() (*string, bool)`

GetArcNamespaceOk returns a tuple with the ArcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcNamespace

`func (o *Idp) SetArcNamespace(v string)`

SetArcNamespace sets ArcNamespace field to given value.

### HasArcNamespace

`func (o *Idp) HasArcNamespace() bool`

HasArcNamespace returns a boolean if a field has been set.

### GetProviderId

`func (o *Idp) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Idp) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Idp) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetType

`func (o *Idp) GetType() OpenIDProviderManualPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Idp) GetTypeOk() (*OpenIDProviderManualPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Idp) SetType(v OpenIDProviderManualPostBodyType)`

SetType sets Type field to given value.


### GetUserMapping

`func (o *Idp) GetUserMapping() LdapProviderPostBodyUserMapping`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *Idp) GetUserMappingOk() (*LdapProviderPostBodyUserMapping, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *Idp) SetUserMapping(v LdapProviderPostBodyUserMapping)`

SetUserMapping sets UserMapping field to given value.


### GetFilters

`func (o *Idp) GetFilters() LdapProviderPostBodyFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Idp) GetFiltersOk() (*LdapProviderPostBodyFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Idp) SetFilters(v LdapProviderPostBodyFilters)`

SetFilters sets Filters field to given value.


### GetDns

`func (o *Idp) GetDns() LdapProviderPostBodySearchBases`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *Idp) GetDnsOk() (*LdapProviderPostBodySearchBases, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *Idp) SetDns(v LdapProviderPostBodySearchBases)`

SetDns sets Dns field to given value.


### GetGroupMapping

`func (o *Idp) GetGroupMapping() LdapProviderPostBodyGroupMapping`

GetGroupMapping returns the GroupMapping field if non-nil, zero value otherwise.

### GetGroupMappingOk

`func (o *Idp) GetGroupMappingOk() (*LdapProviderPostBodyGroupMapping, bool)`

GetGroupMappingOk returns a tuple with the GroupMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMapping

`func (o *Idp) SetGroupMapping(v LdapProviderPostBodyGroupMapping)`

SetGroupMapping sets GroupMapping field to given value.


### GetConnection

`func (o *Idp) GetConnection() LdapProviderPostBodyConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *Idp) GetConnectionOk() (*LdapProviderPostBodyConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *Idp) SetConnection(v LdapProviderPostBodyConnection)`

SetConnection sets Connection field to given value.


### GetSearchBases

`func (o *Idp) GetSearchBases() LdapProviderPostBodySearchBases`

GetSearchBases returns the SearchBases field if non-nil, zero value otherwise.

### GetSearchBasesOk

`func (o *Idp) GetSearchBasesOk() (*LdapProviderPostBodySearchBases, bool)`

GetSearchBasesOk returns a tuple with the SearchBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBases

`func (o *Idp) SetSearchBases(v LdapProviderPostBodySearchBases)`

SetSearchBases sets SearchBases field to given value.


### GetOidcProvider

`func (o *Idp) GetOidcProvider() OpenIDProviderGetOidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *Idp) GetOidcProviderOk() (*OpenIDProviderGetOidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *Idp) SetOidcProvider(v OpenIDProviderGetOidcProvider)`

SetOidcProvider sets OidcProvider field to given value.


### GetAllowUntrustedCertificates

`func (o *Idp) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *Idp) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *Idp) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


