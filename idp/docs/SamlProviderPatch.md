# SamlProviderPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**SamlProviderPatchType**](SamlProviderPatchType.md) |  | [optional] 
**ServiceProvider** | Pointer to [**SamlProviderPatchServiceProvider**](SamlProviderPatchServiceProvider.md) |  | [optional] 
**Saml** | Pointer to [**SamlProviderPatchSaml**](SamlProviderPatchSaml.md) |  | [optional] 
**LoginDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSamlProviderPatch

`func NewSamlProviderPatch() *SamlProviderPatch`

NewSamlProviderPatch instantiates a new SamlProviderPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlProviderPatchWithDefaults

`func NewSamlProviderPatchWithDefaults() *SamlProviderPatch`

NewSamlProviderPatchWithDefaults instantiates a new SamlProviderPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SamlProviderPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlProviderPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlProviderPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlProviderPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SamlProviderPatch) GetType() SamlProviderPatchType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SamlProviderPatch) GetTypeOk() (*SamlProviderPatchType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SamlProviderPatch) SetType(v SamlProviderPatchType)`

SetType sets Type field to given value.

### HasType

`func (o *SamlProviderPatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServiceProvider

`func (o *SamlProviderPatch) GetServiceProvider() SamlProviderPatchServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *SamlProviderPatch) GetServiceProviderOk() (*SamlProviderPatchServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *SamlProviderPatch) SetServiceProvider(v SamlProviderPatchServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *SamlProviderPatch) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.

### GetSaml

`func (o *SamlProviderPatch) GetSaml() SamlProviderPatchSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *SamlProviderPatch) GetSamlOk() (*SamlProviderPatchSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *SamlProviderPatch) SetSaml(v SamlProviderPatchSaml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *SamlProviderPatch) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetLoginDisabled

`func (o *SamlProviderPatch) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *SamlProviderPatch) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *SamlProviderPatch) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *SamlProviderPatch) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


