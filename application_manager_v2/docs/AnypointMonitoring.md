# AnypointMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**AnypointMonitoringResources**](AnypointMonitoringResources.md) |  | [optional] 

## Methods

### NewAnypointMonitoring

`func NewAnypointMonitoring() *AnypointMonitoring`

NewAnypointMonitoring instantiates a new AnypointMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnypointMonitoringWithDefaults

`func NewAnypointMonitoringWithDefaults() *AnypointMonitoring`

NewAnypointMonitoringWithDefaults instantiates a new AnypointMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *AnypointMonitoring) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AnypointMonitoring) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AnypointMonitoring) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AnypointMonitoring) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetResources

`func (o *AnypointMonitoring) GetResources() AnypointMonitoringResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AnypointMonitoring) GetResourcesOk() (*AnypointMonitoringResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AnypointMonitoring) SetResources(v AnypointMonitoringResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AnypointMonitoring) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


