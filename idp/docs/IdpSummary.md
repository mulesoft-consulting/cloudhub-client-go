# IdpSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **string** |  | 
**OrgId** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**IdpSummaryType**](IdpSummaryType.md) |  | 

## Methods

### NewIdpSummary

`func NewIdpSummary(providerId string, orgId string, name string, type_ IdpSummaryType, ) *IdpSummary`

NewIdpSummary instantiates a new IdpSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpSummaryWithDefaults

`func NewIdpSummaryWithDefaults() *IdpSummary`

NewIdpSummaryWithDefaults instantiates a new IdpSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *IdpSummary) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *IdpSummary) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *IdpSummary) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetOrgId

`func (o *IdpSummary) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *IdpSummary) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *IdpSummary) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetName

`func (o *IdpSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpSummary) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IdpSummary) GetType() IdpSummaryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpSummary) GetTypeOk() (*IdpSummaryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpSummary) SetType(v IdpSummaryType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


