# LdapProviderGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserMapping** | [**LdapProviderPostBodyUserMapping**](LdapProviderPostBodyUserMapping.md) |  | 
**Filters** | [**LdapProviderPostBodyFilters**](LdapProviderPostBodyFilters.md) |  | 
**Name** | **string** |  | 
**OrgId** | **string** |  | 
**Dns** | [**LdapProviderPostBodySearchBases**](LdapProviderPostBodySearchBases.md) |  | 
**GroupMapping** | [**LdapProviderPostBodyGroupMapping**](LdapProviderPostBodyGroupMapping.md) |  | 
**Connection** | [**LdapProviderPostBodyConnection**](LdapProviderPostBodyConnection.md) |  | 
**ArcNamespace** | Pointer to **string** |  | [optional] 
**ProviderId** | **string** |  | 
**Type** | [**LdapProviderPostBodyType**](LdapProviderPostBodyType.md) |  | 
**SearchBases** | [**LdapProviderPostBodySearchBases**](LdapProviderPostBodySearchBases.md) |  | 

## Methods

### NewLdapProviderGet

`func NewLdapProviderGet(userMapping LdapProviderPostBodyUserMapping, filters LdapProviderPostBodyFilters, name string, orgId string, dns LdapProviderPostBodySearchBases, groupMapping LdapProviderPostBodyGroupMapping, connection LdapProviderPostBodyConnection, providerId string, type_ LdapProviderPostBodyType, searchBases LdapProviderPostBodySearchBases, ) *LdapProviderGet`

NewLdapProviderGet instantiates a new LdapProviderGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapProviderGetWithDefaults

`func NewLdapProviderGetWithDefaults() *LdapProviderGet`

NewLdapProviderGetWithDefaults instantiates a new LdapProviderGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserMapping

`func (o *LdapProviderGet) GetUserMapping() LdapProviderPostBodyUserMapping`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *LdapProviderGet) GetUserMappingOk() (*LdapProviderPostBodyUserMapping, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *LdapProviderGet) SetUserMapping(v LdapProviderPostBodyUserMapping)`

SetUserMapping sets UserMapping field to given value.


### GetFilters

`func (o *LdapProviderGet) GetFilters() LdapProviderPostBodyFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LdapProviderGet) GetFiltersOk() (*LdapProviderPostBodyFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LdapProviderGet) SetFilters(v LdapProviderPostBodyFilters)`

SetFilters sets Filters field to given value.


### GetName

`func (o *LdapProviderGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapProviderGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapProviderGet) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *LdapProviderGet) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *LdapProviderGet) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *LdapProviderGet) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetDns

`func (o *LdapProviderGet) GetDns() LdapProviderPostBodySearchBases`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *LdapProviderGet) GetDnsOk() (*LdapProviderPostBodySearchBases, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *LdapProviderGet) SetDns(v LdapProviderPostBodySearchBases)`

SetDns sets Dns field to given value.


### GetGroupMapping

`func (o *LdapProviderGet) GetGroupMapping() LdapProviderPostBodyGroupMapping`

GetGroupMapping returns the GroupMapping field if non-nil, zero value otherwise.

### GetGroupMappingOk

`func (o *LdapProviderGet) GetGroupMappingOk() (*LdapProviderPostBodyGroupMapping, bool)`

GetGroupMappingOk returns a tuple with the GroupMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMapping

`func (o *LdapProviderGet) SetGroupMapping(v LdapProviderPostBodyGroupMapping)`

SetGroupMapping sets GroupMapping field to given value.


### GetConnection

`func (o *LdapProviderGet) GetConnection() LdapProviderPostBodyConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *LdapProviderGet) GetConnectionOk() (*LdapProviderPostBodyConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *LdapProviderGet) SetConnection(v LdapProviderPostBodyConnection)`

SetConnection sets Connection field to given value.


### GetArcNamespace

`func (o *LdapProviderGet) GetArcNamespace() string`

GetArcNamespace returns the ArcNamespace field if non-nil, zero value otherwise.

### GetArcNamespaceOk

`func (o *LdapProviderGet) GetArcNamespaceOk() (*string, bool)`

GetArcNamespaceOk returns a tuple with the ArcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcNamespace

`func (o *LdapProviderGet) SetArcNamespace(v string)`

SetArcNamespace sets ArcNamespace field to given value.

### HasArcNamespace

`func (o *LdapProviderGet) HasArcNamespace() bool`

HasArcNamespace returns a boolean if a field has been set.

### GetProviderId

`func (o *LdapProviderGet) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *LdapProviderGet) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *LdapProviderGet) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetType

`func (o *LdapProviderGet) GetType() LdapProviderPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LdapProviderGet) GetTypeOk() (*LdapProviderPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LdapProviderGet) SetType(v LdapProviderPostBodyType)`

SetType sets Type field to given value.


### GetSearchBases

`func (o *LdapProviderGet) GetSearchBases() LdapProviderPostBodySearchBases`

GetSearchBases returns the SearchBases field if non-nil, zero value otherwise.

### GetSearchBasesOk

`func (o *LdapProviderGet) GetSearchBasesOk() (*LdapProviderPostBodySearchBases, bool)`

GetSearchBasesOk returns a tuple with the SearchBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBases

`func (o *LdapProviderGet) SetSearchBases(v LdapProviderPostBodySearchBases)`

SetSearchBases sets SearchBases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


