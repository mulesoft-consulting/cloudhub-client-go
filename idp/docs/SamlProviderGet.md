# SamlProviderGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OrgId** | **string** |  | 
**Saml** | [**SamlProviderPostBodySaml**](SamlProviderPostBodySaml.md) |  | 
**ServiceProvider** | [**SamlProviderPostBodyServiceProvider**](SamlProviderPostBodyServiceProvider.md) |  | 
**LoginDisabled** | Pointer to **bool** |  | [optional] 
**ArcNamespace** | Pointer to **string** |  | [optional] 
**ProviderId** | **string** |  | 
**Type** | [**SamlProviderPostBodyType**](SamlProviderPostBodyType.md) |  | 

## Methods

### NewSamlProviderGet

`func NewSamlProviderGet(name string, orgId string, saml SamlProviderPostBodySaml, serviceProvider SamlProviderPostBodyServiceProvider, providerId string, type_ SamlProviderPostBodyType, ) *SamlProviderGet`

NewSamlProviderGet instantiates a new SamlProviderGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlProviderGetWithDefaults

`func NewSamlProviderGetWithDefaults() *SamlProviderGet`

NewSamlProviderGetWithDefaults instantiates a new SamlProviderGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SamlProviderGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlProviderGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlProviderGet) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *SamlProviderGet) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SamlProviderGet) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SamlProviderGet) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSaml

`func (o *SamlProviderGet) GetSaml() SamlProviderPostBodySaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *SamlProviderGet) GetSamlOk() (*SamlProviderPostBodySaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *SamlProviderGet) SetSaml(v SamlProviderPostBodySaml)`

SetSaml sets Saml field to given value.


### GetServiceProvider

`func (o *SamlProviderGet) GetServiceProvider() SamlProviderPostBodyServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *SamlProviderGet) GetServiceProviderOk() (*SamlProviderPostBodyServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *SamlProviderGet) SetServiceProvider(v SamlProviderPostBodyServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.


### GetLoginDisabled

`func (o *SamlProviderGet) GetLoginDisabled() bool`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *SamlProviderGet) GetLoginDisabledOk() (*bool, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *SamlProviderGet) SetLoginDisabled(v bool)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *SamlProviderGet) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.

### GetArcNamespace

`func (o *SamlProviderGet) GetArcNamespace() string`

GetArcNamespace returns the ArcNamespace field if non-nil, zero value otherwise.

### GetArcNamespaceOk

`func (o *SamlProviderGet) GetArcNamespaceOk() (*string, bool)`

GetArcNamespaceOk returns a tuple with the ArcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcNamespace

`func (o *SamlProviderGet) SetArcNamespace(v string)`

SetArcNamespace sets ArcNamespace field to given value.

### HasArcNamespace

`func (o *SamlProviderGet) HasArcNamespace() bool`

HasArcNamespace returns a boolean if a field has been set.

### GetProviderId

`func (o *SamlProviderGet) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *SamlProviderGet) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *SamlProviderGet) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetType

`func (o *SamlProviderGet) GetType() SamlProviderPostBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SamlProviderGet) GetTypeOk() (*SamlProviderPostBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SamlProviderGet) SetType(v SamlProviderPostBodyType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


