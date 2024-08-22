# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** | The cloud provider the target belongs to. | [optional] 
**TargetId** | Pointer to **string** | The unique identifier of the target. | [optional] 
**DeploymentSettings** | Pointer to [**DeploymentSettings**](DeploymentSettings.md) |  | [optional] 
**Replicas** | Pointer to **int32** | The number of replicas of the deployment. | [optional] 

## Methods

### NewTarget

`func NewTarget() *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *Target) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Target) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Target) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Target) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetTargetId

`func (o *Target) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Target) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Target) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Target) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetDeploymentSettings

`func (o *Target) GetDeploymentSettings() DeploymentSettings`

GetDeploymentSettings returns the DeploymentSettings field if non-nil, zero value otherwise.

### GetDeploymentSettingsOk

`func (o *Target) GetDeploymentSettingsOk() (*DeploymentSettings, bool)`

GetDeploymentSettingsOk returns a tuple with the DeploymentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSettings

`func (o *Target) SetDeploymentSettings(v DeploymentSettings)`

SetDeploymentSettings sets DeploymentSettings field to given value.

### HasDeploymentSettings

`func (o *Target) HasDeploymentSettings() bool`

HasDeploymentSettings returns a boolean if a field has been set.

### GetReplicas

`func (o *Target) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Target) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Target) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *Target) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


