# Scheduler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**FlowName** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TimeUnit** | Pointer to **string** |  | [optional] 
**Frequency** | Pointer to **string** |  | [optional] 
**StartDelay** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 

## Methods

### NewScheduler

`func NewScheduler() *Scheduler`

NewScheduler instantiates a new Scheduler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerWithDefaults

`func NewSchedulerWithDefaults() *Scheduler`

NewSchedulerWithDefaults instantiates a new Scheduler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Scheduler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Scheduler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Scheduler) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Scheduler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Scheduler) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Scheduler) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Scheduler) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Scheduler) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlowName

`func (o *Scheduler) GetFlowName() string`

GetFlowName returns the FlowName field if non-nil, zero value otherwise.

### GetFlowNameOk

`func (o *Scheduler) GetFlowNameOk() (*string, bool)`

GetFlowNameOk returns a tuple with the FlowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowName

`func (o *Scheduler) SetFlowName(v string)`

SetFlowName sets FlowName field to given value.

### HasFlowName

`func (o *Scheduler) HasFlowName() bool`

HasFlowName returns a boolean if a field has been set.

### GetEnabled

`func (o *Scheduler) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Scheduler) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Scheduler) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Scheduler) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeUnit

`func (o *Scheduler) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *Scheduler) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *Scheduler) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *Scheduler) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.

### GetFrequency

`func (o *Scheduler) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Scheduler) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Scheduler) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *Scheduler) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetStartDelay

`func (o *Scheduler) GetStartDelay() string`

GetStartDelay returns the StartDelay field if non-nil, zero value otherwise.

### GetStartDelayOk

`func (o *Scheduler) GetStartDelayOk() (*string, bool)`

GetStartDelayOk returns a tuple with the StartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDelay

`func (o *Scheduler) SetStartDelay(v string)`

SetStartDelay sets StartDelay field to given value.

### HasStartDelay

`func (o *Scheduler) HasStartDelay() bool`

HasStartDelay returns a boolean if a field has been set.

### GetExpression

`func (o *Scheduler) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Scheduler) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Scheduler) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Scheduler) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetTimeZone

`func (o *Scheduler) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Scheduler) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Scheduler) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Scheduler) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


