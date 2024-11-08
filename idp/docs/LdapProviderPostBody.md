# LdapProviderPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ArcNamespace** | Pointer to **string** |  | [optional] 
**Type** | [**LdapProviderPostBodyType**](LdapProviderPostBodyType.md) |  | 
**Connection** | [**LdapProviderPostBodyConnection**](LdapProviderPostBodyConnection.md) |  | 
**SearchBases** | [**LdapProviderPostBodySearchBases**](LdapProviderPostBodySearchBases.md) |  | 
**Dns** | [**LdapProviderPostBodySearchBases**](LdapProviderPostBodySearchBases.md) |  | 
**Filters** | [**LdapProviderPostBodyFilters**](LdapProviderPostBodyFilters.md) |  | 
**UserMapping** | [**LdapProviderPostBodyUserMapping**](LdapProviderPostBodyUserMapping.md) |  | 
**GroupMapping** | [**LdapProviderPostBodyGroupMapping**](LdapProviderPostBodyGroupMapping.md) |  | 

## Methods

### NewLdapProviderPostBody

`func NewLdapProviderPostBody(name string, type_ LdapProviderPostBodyType, connection LdapProviderPostBodyConnection, searchBases LdapProviderPostBodySearchBases, dns LdapProviderPostBodySearchBases, filters LdapProviderPostBodyFilters, userMapping LdapProviderPostBodyUserMapping, groupMapping LdapProviderPostBodyGroupMapping, ) *LdapProviderPostBody`

NewLdapProviderPostBody instantiates a new LdapProviderPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapProviderPostBodyWithDefaults

`func NewLdapProviderPostBodyWithDefaults() *LdapProviderPostBody`

NewLdapProviderPostBodyWithDefaults instantiates a new LdapProviderPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LdapProviderPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapProviderPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapProviderPostBody) SetName(v string)`

SetName sets Name field to given value.


### GetArcNamespace

`func (o *LdapProviderPostBody) GetArcNamespace() string`

GetArcNamespace returns the ArcNamespace field if non-nil, zero value otherwise.

### GetArcNamespaceOk

`func (o *LdapProviderPostBody) GetArcNamespaceOk() (*string, bool)`

GetArcNamespaceOk returns a tuple with the ArcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcNamespace

`func (o *LdapProviderPostBody) SetArcNamespace(v string)`

SetArcNamespace sets ArcNamespace field to given value.

### HasArcNamespace

`func (o *LdapProviderPostBody) HasArcNamespace() bool`

HasArcNamespace returns a boolean if a field has been set.

### GetType

`func (o *LdapProviderPostBody) GetType() LdapProviderPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LdapProviderPostBody) GetTypeOk() (*LdapProviderPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LdapProviderPostBody) SetType(v LdapProviderPostBodyType)`

SetType sets Type field to given value.


### GetConnection

`func (o *LdapProviderPostBody) GetConnection() LdapProviderPostBodyConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *LdapProviderPostBody) GetConnectionOk() (*LdapProviderPostBodyConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *LdapProviderPostBody) SetConnection(v LdapProviderPostBodyConnection)`

SetConnection sets Connection field to given value.


### GetSearchBases

`func (o *LdapProviderPostBody) GetSearchBases() LdapProviderPostBodySearchBases`

GetSearchBases returns the SearchBases field if non-nil, zero value otherwise.

### GetSearchBasesOk

`func (o *LdapProviderPostBody) GetSearchBasesOk() (*LdapProviderPostBodySearchBases, bool)`

GetSearchBasesOk returns a tuple with the SearchBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBases

`func (o *LdapProviderPostBody) SetSearchBases(v LdapProviderPostBodySearchBases)`

SetSearchBases sets SearchBases field to given value.


### GetDns

`func (o *LdapProviderPostBody) GetDns() LdapProviderPostBodySearchBases`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *LdapProviderPostBody) GetDnsOk() (*LdapProviderPostBodySearchBases, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *LdapProviderPostBody) SetDns(v LdapProviderPostBodySearchBases)`

SetDns sets Dns field to given value.


### GetFilters

`func (o *LdapProviderPostBody) GetFilters() LdapProviderPostBodyFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LdapProviderPostBody) GetFiltersOk() (*LdapProviderPostBodyFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LdapProviderPostBody) SetFilters(v LdapProviderPostBodyFilters)`

SetFilters sets Filters field to given value.


### GetUserMapping

`func (o *LdapProviderPostBody) GetUserMapping() LdapProviderPostBodyUserMapping`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *LdapProviderPostBody) GetUserMappingOk() (*LdapProviderPostBodyUserMapping, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *LdapProviderPostBody) SetUserMapping(v LdapProviderPostBodyUserMapping)`

SetUserMapping sets UserMapping field to given value.


### GetGroupMapping

`func (o *LdapProviderPostBody) GetGroupMapping() LdapProviderPostBodyGroupMapping`

GetGroupMapping returns the GroupMapping field if non-nil, zero value otherwise.

### GetGroupMappingOk

`func (o *LdapProviderPostBody) GetGroupMappingOk() (*LdapProviderPostBodyGroupMapping, bool)`

GetGroupMappingOk returns a tuple with the GroupMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMapping

`func (o *LdapProviderPostBody) SetGroupMapping(v LdapProviderPostBodyGroupMapping)`

SetGroupMapping sets GroupMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


