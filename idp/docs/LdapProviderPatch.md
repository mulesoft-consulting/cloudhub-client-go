# LdapProviderPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**SamlProviderPatchType**](SamlProviderPatchType.md) |  | [optional] 
**Connection** | Pointer to [**LdapProviderPatchConnection**](LdapProviderPatchConnection.md) |  | [optional] 
**SearchBases** | Pointer to [**LdapProviderPatchSearchBases**](LdapProviderPatchSearchBases.md) |  | [optional] 
**Dns** | Pointer to [**LdapProviderPatchSearchBases**](LdapProviderPatchSearchBases.md) |  | [optional] 
**Filters** | Pointer to [**LdapProviderPatchFilters**](LdapProviderPatchFilters.md) |  | [optional] 
**UserMapping** | Pointer to [**LdapProviderPatchUserMapping**](LdapProviderPatchUserMapping.md) |  | [optional] 
**GroupMapping** | Pointer to [**LdapProviderPatchGroupMapping**](LdapProviderPatchGroupMapping.md) |  | [optional] 

## Methods

### NewLdapProviderPatch

`func NewLdapProviderPatch() *LdapProviderPatch`

NewLdapProviderPatch instantiates a new LdapProviderPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapProviderPatchWithDefaults

`func NewLdapProviderPatchWithDefaults() *LdapProviderPatch`

NewLdapProviderPatchWithDefaults instantiates a new LdapProviderPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LdapProviderPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapProviderPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapProviderPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapProviderPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *LdapProviderPatch) GetType() SamlProviderPatchType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LdapProviderPatch) GetTypeOk() (*SamlProviderPatchType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LdapProviderPatch) SetType(v SamlProviderPatchType)`

SetType sets Type field to given value.

### HasType

`func (o *LdapProviderPatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConnection

`func (o *LdapProviderPatch) GetConnection() LdapProviderPatchConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *LdapProviderPatch) GetConnectionOk() (*LdapProviderPatchConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *LdapProviderPatch) SetConnection(v LdapProviderPatchConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *LdapProviderPatch) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetSearchBases

`func (o *LdapProviderPatch) GetSearchBases() LdapProviderPatchSearchBases`

GetSearchBases returns the SearchBases field if non-nil, zero value otherwise.

### GetSearchBasesOk

`func (o *LdapProviderPatch) GetSearchBasesOk() (*LdapProviderPatchSearchBases, bool)`

GetSearchBasesOk returns a tuple with the SearchBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBases

`func (o *LdapProviderPatch) SetSearchBases(v LdapProviderPatchSearchBases)`

SetSearchBases sets SearchBases field to given value.

### HasSearchBases

`func (o *LdapProviderPatch) HasSearchBases() bool`

HasSearchBases returns a boolean if a field has been set.

### GetDns

`func (o *LdapProviderPatch) GetDns() LdapProviderPatchSearchBases`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *LdapProviderPatch) GetDnsOk() (*LdapProviderPatchSearchBases, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *LdapProviderPatch) SetDns(v LdapProviderPatchSearchBases)`

SetDns sets Dns field to given value.

### HasDns

`func (o *LdapProviderPatch) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetFilters

`func (o *LdapProviderPatch) GetFilters() LdapProviderPatchFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LdapProviderPatch) GetFiltersOk() (*LdapProviderPatchFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LdapProviderPatch) SetFilters(v LdapProviderPatchFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *LdapProviderPatch) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUserMapping

`func (o *LdapProviderPatch) GetUserMapping() LdapProviderPatchUserMapping`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *LdapProviderPatch) GetUserMappingOk() (*LdapProviderPatchUserMapping, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *LdapProviderPatch) SetUserMapping(v LdapProviderPatchUserMapping)`

SetUserMapping sets UserMapping field to given value.

### HasUserMapping

`func (o *LdapProviderPatch) HasUserMapping() bool`

HasUserMapping returns a boolean if a field has been set.

### GetGroupMapping

`func (o *LdapProviderPatch) GetGroupMapping() LdapProviderPatchGroupMapping`

GetGroupMapping returns the GroupMapping field if non-nil, zero value otherwise.

### GetGroupMappingOk

`func (o *LdapProviderPatch) GetGroupMappingOk() (*LdapProviderPatchGroupMapping, bool)`

GetGroupMappingOk returns a tuple with the GroupMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMapping

`func (o *LdapProviderPatch) SetGroupMapping(v LdapProviderPatchGroupMapping)`

SetGroupMapping sets GroupMapping field to given value.

### HasGroupMapping

`func (o *LdapProviderPatch) HasGroupMapping() bool`

HasGroupMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


