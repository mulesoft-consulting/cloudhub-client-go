# MuleAgentSchedulingService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** |  | [optional] 
**Schedulers** | Pointer to [**[]Scheduler**](Scheduler.md) |  | [optional] 

## Methods

### NewMuleAgentSchedulingService

`func NewMuleAgentSchedulingService() *MuleAgentSchedulingService`

NewMuleAgentSchedulingService instantiates a new MuleAgentSchedulingService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuleAgentSchedulingServiceWithDefaults

`func NewMuleAgentSchedulingServiceWithDefaults() *MuleAgentSchedulingService`

NewMuleAgentSchedulingServiceWithDefaults instantiates a new MuleAgentSchedulingService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *MuleAgentSchedulingService) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *MuleAgentSchedulingService) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *MuleAgentSchedulingService) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *MuleAgentSchedulingService) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetSchedulers

`func (o *MuleAgentSchedulingService) GetSchedulers() []Scheduler`

GetSchedulers returns the Schedulers field if non-nil, zero value otherwise.

### GetSchedulersOk

`func (o *MuleAgentSchedulingService) GetSchedulersOk() (*[]Scheduler, bool)`

GetSchedulersOk returns a tuple with the Schedulers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulers

`func (o *MuleAgentSchedulingService) SetSchedulers(v []Scheduler)`

SetSchedulers sets Schedulers field to given value.

### HasSchedulers

`func (o *MuleAgentSchedulingService) HasSchedulers() bool`

HasSchedulers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


