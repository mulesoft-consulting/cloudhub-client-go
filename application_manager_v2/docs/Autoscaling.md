# Autoscaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enables or disables the Autoscaling feature. The possible values are: true or false.  | [optional] 
**MinReplicas** | Pointer to **int32** | Set the minimum amount of replicas for your deployment. The minimum accepted value is 1. The maximum is 3.  | [optional] 
**MaxReplicas** | Pointer to **int32** | Set the maximum amount of replicas your application can scale to. The minimum accepted value is 2. The maximum is 32.  | [optional] 

## Methods

### NewAutoscaling

`func NewAutoscaling() *Autoscaling`

NewAutoscaling instantiates a new Autoscaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalingWithDefaults

`func NewAutoscalingWithDefaults() *Autoscaling`

NewAutoscalingWithDefaults instantiates a new Autoscaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Autoscaling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Autoscaling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Autoscaling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Autoscaling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinReplicas

`func (o *Autoscaling) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *Autoscaling) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *Autoscaling) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *Autoscaling) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *Autoscaling) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *Autoscaling) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *Autoscaling) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *Autoscaling) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


