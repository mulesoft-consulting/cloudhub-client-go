# Replicas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the replica in RTF | [optional] 
**State** | Pointer to **string** | The current state of the replica | [optional] 
**DeploymentLocation** | Pointer to **string** | The node id in which the replica is deployed. | [optional] 
**CurrentDeploymentVersion** | Pointer to **string** | The version deployed in the replica. | [optional] 
**Reason** | Pointer to **string** | In case of an error, it should provide information about the root cause. | [optional] 

## Methods

### NewReplicas

`func NewReplicas() *Replicas`

NewReplicas instantiates a new Replicas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicasWithDefaults

`func NewReplicasWithDefaults() *Replicas`

NewReplicasWithDefaults instantiates a new Replicas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Replicas) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Replicas) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Replicas) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Replicas) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *Replicas) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Replicas) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Replicas) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Replicas) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDeploymentLocation

`func (o *Replicas) GetDeploymentLocation() string`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *Replicas) GetDeploymentLocationOk() (*string, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *Replicas) SetDeploymentLocation(v string)`

SetDeploymentLocation sets DeploymentLocation field to given value.

### HasDeploymentLocation

`func (o *Replicas) HasDeploymentLocation() bool`

HasDeploymentLocation returns a boolean if a field has been set.

### GetCurrentDeploymentVersion

`func (o *Replicas) GetCurrentDeploymentVersion() string`

GetCurrentDeploymentVersion returns the CurrentDeploymentVersion field if non-nil, zero value otherwise.

### GetCurrentDeploymentVersionOk

`func (o *Replicas) GetCurrentDeploymentVersionOk() (*string, bool)`

GetCurrentDeploymentVersionOk returns a tuple with the CurrentDeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeploymentVersion

`func (o *Replicas) SetCurrentDeploymentVersion(v string)`

SetCurrentDeploymentVersion sets CurrentDeploymentVersion field to given value.

### HasCurrentDeploymentVersion

`func (o *Replicas) HasCurrentDeploymentVersion() bool`

HasCurrentDeploymentVersion returns a boolean if a field has been set.

### GetReason

`func (o *Replicas) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Replicas) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Replicas) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Replicas) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


