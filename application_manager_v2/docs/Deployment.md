# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**LastModifiedDate** | Pointer to **int64** |  | [optional] 
**DesiredVersion** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to [**[]Replicas**](Replicas.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**Target** | Pointer to [**Target**](Target.md) |  | [optional] 
**LastSuccessfulVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Deployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Deployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Deployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Deployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Deployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Deployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationDate

`func (o *Deployment) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Deployment) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Deployment) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Deployment) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *Deployment) GetLastModifiedDate() int64`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *Deployment) GetLastModifiedDateOk() (*int64, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *Deployment) SetLastModifiedDate(v int64)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *Deployment) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetDesiredVersion

`func (o *Deployment) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *Deployment) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *Deployment) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *Deployment) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.

### GetReplicas

`func (o *Deployment) GetReplicas() []Replicas`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Deployment) GetReplicasOk() (*[]Replicas, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Deployment) SetReplicas(v []Replicas)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *Deployment) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetStatus

`func (o *Deployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Deployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApplication

`func (o *Deployment) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Deployment) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Deployment) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Deployment) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetTarget

`func (o *Deployment) GetTarget() Target`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Deployment) GetTargetOk() (*Target, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Deployment) SetTarget(v Target)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Deployment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetLastSuccessfulVersion

`func (o *Deployment) GetLastSuccessfulVersion() string`

GetLastSuccessfulVersion returns the LastSuccessfulVersion field if non-nil, zero value otherwise.

### GetLastSuccessfulVersionOk

`func (o *Deployment) GetLastSuccessfulVersionOk() (*string, bool)`

GetLastSuccessfulVersionOk returns a tuple with the LastSuccessfulVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulVersion

`func (o *Deployment) SetLastSuccessfulVersion(v string)`

SetLastSuccessfulVersion sets LastSuccessfulVersion field to given value.

### HasLastSuccessfulVersion

`func (o *Deployment) HasLastSuccessfulVersion() bool`

HasLastSuccessfulVersion returns a boolean if a field has been set.

### SetLastSuccessfulVersionNil

`func (o *Deployment) SetLastSuccessfulVersionNil(b bool)`

 SetLastSuccessfulVersionNil sets the value for LastSuccessfulVersion to be an explicit nil

### UnsetLastSuccessfulVersion
`func (o *Deployment) UnsetLastSuccessfulVersion()`

UnsetLastSuccessfulVersion ensures that no value is present for LastSuccessfulVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


