# LdapProviderPatchFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupsByUsername** | Pointer to **string** |  | [optional] 
**UserByUsername** | Pointer to **string** |  | [optional] 

## Methods

### NewLdapProviderPatchFilters

`func NewLdapProviderPatchFilters() *LdapProviderPatchFilters`

NewLdapProviderPatchFilters instantiates a new LdapProviderPatchFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapProviderPatchFiltersWithDefaults

`func NewLdapProviderPatchFiltersWithDefaults() *LdapProviderPatchFilters`

NewLdapProviderPatchFiltersWithDefaults instantiates a new LdapProviderPatchFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupsByUsername

`func (o *LdapProviderPatchFilters) GetGroupsByUsername() string`

GetGroupsByUsername returns the GroupsByUsername field if non-nil, zero value otherwise.

### GetGroupsByUsernameOk

`func (o *LdapProviderPatchFilters) GetGroupsByUsernameOk() (*string, bool)`

GetGroupsByUsernameOk returns a tuple with the GroupsByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsByUsername

`func (o *LdapProviderPatchFilters) SetGroupsByUsername(v string)`

SetGroupsByUsername sets GroupsByUsername field to given value.

### HasGroupsByUsername

`func (o *LdapProviderPatchFilters) HasGroupsByUsername() bool`

HasGroupsByUsername returns a boolean if a field has been set.

### GetUserByUsername

`func (o *LdapProviderPatchFilters) GetUserByUsername() string`

GetUserByUsername returns the UserByUsername field if non-nil, zero value otherwise.

### GetUserByUsernameOk

`func (o *LdapProviderPatchFilters) GetUserByUsernameOk() (*string, bool)`

GetUserByUsernameOk returns a tuple with the UserByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserByUsername

`func (o *LdapProviderPatchFilters) SetUserByUsername(v string)`

SetUserByUsername sets UserByUsername field to given value.

### HasUserByUsername

`func (o *LdapProviderPatchFilters) HasUserByUsername() bool`

HasUserByUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


