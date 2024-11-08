# SamlProviderPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcNamespace** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Type** | [**SamlProviderPostBodyType**](SamlProviderPostBodyType.md) |  | 
**ServiceProvider** | [**SamlProviderPostBodyServiceProvider**](SamlProviderPostBodyServiceProvider.md) |  | 
**Saml** | [**SamlProviderPostBodySaml**](SamlProviderPostBodySaml.md) |  | 
**LoginDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSamlProviderPostBody

`func NewSamlProviderPostBody(name string, type_ SamlProviderPostBodyType, serviceProvider SamlProviderPostBodyServiceProvider, saml SamlProviderPostBodySaml, ) *SamlProviderPostBody`

NewSamlProviderPostBody instantiates a new SamlProviderPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlProviderPostBodyWithDefaults

`func NewSamlProviderPostBodyWithDefaults() *SamlProviderPostBody`

NewSamlProviderPostBodyWithDefaults instantiates a new SamlProviderPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcNamespace

`func (o *SamlProviderPostBody) GetArcNamespace() string`

GetArcNamespace returns the ArcNamespace field if non-nil, zero value otherwise.

### GetArcNamespaceOk

`func (o *SamlProviderPostBody) GetArcNamespaceOk() (*string, bool)`

GetArcNamespaceOk returns a tuple with the ArcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcNamespace

`func (o *SamlProviderPostBody) SetArcNamespace(v string)`

SetArcNamespace sets ArcNamespace field to given value.

### HasArcNamespace

`func (o *SamlProviderPostBody) HasArcNamespace() bool`

HasArcNamespace returns a boolean if a field has been set.

### GetName

`func (o *SamlProviderPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlProviderPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlProviderPostBody) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SamlProviderPostBody) GetType() SamlProviderPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SamlProviderPostBody) GetTypeOk() (*SamlProviderPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SamlProviderPostBody) SetType(v SamlProviderPostBodyType)`

SetType sets Type field to given value.


### GetServiceProvider

`func (o *SamlProviderPostBody) GetServiceProvider() SamlProviderPostBodyServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *SamlProviderPostBody) GetServiceProviderOk() (*SamlProviderPostBodyServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *SamlProviderPostBody) SetServiceProvider(v SamlProviderPostBodyServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.


### GetSaml

`func (o *SamlProviderPostBody) GetSaml() SamlProviderPostBodySaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *SamlProviderPostBody) GetSamlOk() (*SamlProviderPostBodySaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *SamlProviderPostBody) SetSaml(v SamlProviderPostBodySaml)`

SetSaml sets Saml field to given value.


### GetLoginDisabled

`func (o *SamlProviderPostBody) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *SamlProviderPostBody) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *SamlProviderPostBody) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *SamlProviderPostBody) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


