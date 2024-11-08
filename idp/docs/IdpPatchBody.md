# IdpPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**SamlProviderPatchType**](SamlProviderPatchType.md) |  | [optional] 
**ServiceProvider** | Pointer to [**SamlProviderPatchServiceProvider**](SamlProviderPatchServiceProvider.md) |  | [optional] 
**Saml** | Pointer to [**SamlProviderPatchSaml**](SamlProviderPatchSaml.md) |  | [optional] 
**LoginDisabled** | Pointer to **bool** |  | [optional] 
**Connection** | Pointer to [**LdapProviderPatchConnection**](LdapProviderPatchConnection.md) |  | [optional] 
**SearchBases** | Pointer to [**LdapProviderPatchSearchBases**](LdapProviderPatchSearchBases.md) |  | [optional] 
**Dns** | Pointer to [**LdapProviderPatchSearchBases**](LdapProviderPatchSearchBases.md) |  | [optional] 
**Filters** | Pointer to [**LdapProviderPatchFilters**](LdapProviderPatchFilters.md) |  | [optional] 
**UserMapping** | Pointer to [**LdapProviderPatchUserMapping**](LdapProviderPatchUserMapping.md) |  | [optional] 
**GroupMapping** | Pointer to [**LdapProviderPatchGroupMapping**](LdapProviderPatchGroupMapping.md) |  | [optional] 
**OidcProvider** | Pointer to [**OpenIDProviderPatchOidcProvider**](OpenIDProviderPatchOidcProvider.md) |  | [optional] 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 

## Methods

### NewIdpPatchBody

`func NewIdpPatchBody() *IdpPatchBody`

NewIdpPatchBody instantiates a new IdpPatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPatchBodyWithDefaults

`func NewIdpPatchBodyWithDefaults() *IdpPatchBody`

NewIdpPatchBodyWithDefaults instantiates a new IdpPatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdpPatchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpPatchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpPatchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdpPatchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IdpPatchBody) GetType() SamlProviderPatchType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpPatchBody) GetTypeOk() (*SamlProviderPatchType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpPatchBody) SetType(v SamlProviderPatchType)`

SetType sets Type field to given value.

### HasType

`func (o *IdpPatchBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServiceProvider

`func (o *IdpPatchBody) GetServiceProvider() SamlProviderPatchServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *IdpPatchBody) GetServiceProviderOk() (*SamlProviderPatchServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *IdpPatchBody) SetServiceProvider(v SamlProviderPatchServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *IdpPatchBody) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.

### GetSaml

`func (o *IdpPatchBody) GetSaml() SamlProviderPatchSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *IdpPatchBody) GetSamlOk() (*SamlProviderPatchSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *IdpPatchBody) SetSaml(v SamlProviderPatchSaml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *IdpPatchBody) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetLoginDisabled

`func (o *IdpPatchBody) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *IdpPatchBody) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *IdpPatchBody) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *IdpPatchBody) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.

### GetConnection

`func (o *IdpPatchBody) GetConnection() LdapProviderPatchConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *IdpPatchBody) GetConnectionOk() (*LdapProviderPatchConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *IdpPatchBody) SetConnection(v LdapProviderPatchConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *IdpPatchBody) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetSearchBases

`func (o *IdpPatchBody) GetSearchBases() LdapProviderPatchSearchBases`

GetSearchBases returns the SearchBases field if non-nil, zero value otherwise.

### GetSearchBasesOk

`func (o *IdpPatchBody) GetSearchBasesOk() (*LdapProviderPatchSearchBases, bool)`

GetSearchBasesOk returns a tuple with the SearchBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBases

`func (o *IdpPatchBody) SetSearchBases(v LdapProviderPatchSearchBases)`

SetSearchBases sets SearchBases field to given value.

### HasSearchBases

`func (o *IdpPatchBody) HasSearchBases() bool`

HasSearchBases returns a boolean if a field has been set.

### GetDns

`func (o *IdpPatchBody) GetDns() LdapProviderPatchSearchBases`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *IdpPatchBody) GetDnsOk() (*LdapProviderPatchSearchBases, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *IdpPatchBody) SetDns(v LdapProviderPatchSearchBases)`

SetDns sets Dns field to given value.

### HasDns

`func (o *IdpPatchBody) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetFilters

`func (o *IdpPatchBody) GetFilters() LdapProviderPatchFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *IdpPatchBody) GetFiltersOk() (*LdapProviderPatchFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *IdpPatchBody) SetFilters(v LdapProviderPatchFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *IdpPatchBody) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUserMapping

`func (o *IdpPatchBody) GetUserMapping() LdapProviderPatchUserMapping`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *IdpPatchBody) GetUserMappingOk() (*LdapProviderPatchUserMapping, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *IdpPatchBody) SetUserMapping(v LdapProviderPatchUserMapping)`

SetUserMapping sets UserMapping field to given value.

### HasUserMapping

`func (o *IdpPatchBody) HasUserMapping() bool`

HasUserMapping returns a boolean if a field has been set.

### GetGroupMapping

`func (o *IdpPatchBody) GetGroupMapping() LdapProviderPatchGroupMapping`

GetGroupMapping returns the GroupMapping field if non-nil, zero value otherwise.

### GetGroupMappingOk

`func (o *IdpPatchBody) GetGroupMappingOk() (*LdapProviderPatchGroupMapping, bool)`

GetGroupMappingOk returns a tuple with the GroupMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMapping

`func (o *IdpPatchBody) SetGroupMapping(v LdapProviderPatchGroupMapping)`

SetGroupMapping sets GroupMapping field to given value.

### HasGroupMapping

`func (o *IdpPatchBody) HasGroupMapping() bool`

HasGroupMapping returns a boolean if a field has been set.

### GetOidcProvider

`func (o *IdpPatchBody) GetOidcProvider() OpenIDProviderPatchOidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *IdpPatchBody) GetOidcProviderOk() (*OpenIDProviderPatchOidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *IdpPatchBody) SetOidcProvider(v OpenIDProviderPatchOidcProvider)`

SetOidcProvider sets OidcProvider field to given value.

### HasOidcProvider

`func (o *IdpPatchBody) HasOidcProvider() bool`

HasOidcProvider returns a boolean if a field has been set.

### GetAllowUntrustedCertificates

`func (o *IdpPatchBody) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *IdpPatchBody) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *IdpPatchBody) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *IdpPatchBody) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


