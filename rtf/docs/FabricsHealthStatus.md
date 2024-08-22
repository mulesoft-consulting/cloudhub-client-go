# FabricsHealthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Healthy** | Pointer to **bool** | True is the component is healthy. | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Probes** | Pointer to **string** | Probes collected for this health check. Only applicable for Appliance probes. | [optional] 
**FailedProbes** | Pointer to [**[]FabricsHealthStatusFailedProbesInner**](FabricsHealthStatusFailedProbesInner.md) | Probe failures attributing to the result of this health check. | [optional] 

## Methods

### NewFabricsHealthStatus

`func NewFabricsHealthStatus() *FabricsHealthStatus`

NewFabricsHealthStatus instantiates a new FabricsHealthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricsHealthStatusWithDefaults

`func NewFabricsHealthStatusWithDefaults() *FabricsHealthStatus`

NewFabricsHealthStatusWithDefaults instantiates a new FabricsHealthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthy

`func (o *FabricsHealthStatus) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *FabricsHealthStatus) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *FabricsHealthStatus) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *FabricsHealthStatus) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FabricsHealthStatus) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FabricsHealthStatus) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FabricsHealthStatus) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FabricsHealthStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProbes

`func (o *FabricsHealthStatus) GetProbes() string`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *FabricsHealthStatus) GetProbesOk() (*string, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *FabricsHealthStatus) SetProbes(v string)`

SetProbes sets Probes field to given value.

### HasProbes

`func (o *FabricsHealthStatus) HasProbes() bool`

HasProbes returns a boolean if a field has been set.

### GetFailedProbes

`func (o *FabricsHealthStatus) GetFailedProbes() []FabricsHealthStatusFailedProbesInner`

GetFailedProbes returns the FailedProbes field if non-nil, zero value otherwise.

### GetFailedProbesOk

`func (o *FabricsHealthStatus) GetFailedProbesOk() (*[]FabricsHealthStatusFailedProbesInner, bool)`

GetFailedProbesOk returns a tuple with the FailedProbes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedProbes

`func (o *FabricsHealthStatus) SetFailedProbes(v []FabricsHealthStatusFailedProbesInner)`

SetFailedProbes sets FailedProbes field to given value.

### HasFailedProbes

`func (o *FabricsHealthStatus) HasFailedProbes() bool`

HasFailedProbes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


