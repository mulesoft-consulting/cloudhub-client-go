# DeploymentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | deployment id | [optional] 
**Name** | Pointer to **string** | the name of deployment | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**LastModifiedDate** | Pointer to **int64** |  | [optional] 
**Target** | Pointer to [**DeploymentItemTarget**](DeploymentItemTarget.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**DeploymentItemApplication**](DeploymentItemApplication.md) |  | [optional] 
**CurrentRuntimeVersion** | Pointer to **string** |  | [optional] 
**LastSuccessfulRuntimeVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentItem

`func NewDeploymentItem() *DeploymentItem`

NewDeploymentItem instantiates a new DeploymentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentItemWithDefaults

`func NewDeploymentItemWithDefaults() *DeploymentItem`

NewDeploymentItemWithDefaults instantiates a new DeploymentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeploymentItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationDate

`func (o *DeploymentItem) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *DeploymentItem) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *DeploymentItem) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *DeploymentItem) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *DeploymentItem) GetLastModifiedDate() int64`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *DeploymentItem) GetLastModifiedDateOk() (*int64, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *DeploymentItem) SetLastModifiedDate(v int64)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *DeploymentItem) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetTarget

`func (o *DeploymentItem) GetTarget() DeploymentItemTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DeploymentItem) GetTargetOk() (*DeploymentItemTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DeploymentItem) SetTarget(v DeploymentItemTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DeploymentItem) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApplication

`func (o *DeploymentItem) GetApplication() DeploymentItemApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DeploymentItem) GetApplicationOk() (*DeploymentItemApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DeploymentItem) SetApplication(v DeploymentItemApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *DeploymentItem) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCurrentRuntimeVersion

`func (o *DeploymentItem) GetCurrentRuntimeVersion() string`

GetCurrentRuntimeVersion returns the CurrentRuntimeVersion field if non-nil, zero value otherwise.

### GetCurrentRuntimeVersionOk

`func (o *DeploymentItem) GetCurrentRuntimeVersionOk() (*string, bool)`

GetCurrentRuntimeVersionOk returns a tuple with the CurrentRuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRuntimeVersion

`func (o *DeploymentItem) SetCurrentRuntimeVersion(v string)`

SetCurrentRuntimeVersion sets CurrentRuntimeVersion field to given value.

### HasCurrentRuntimeVersion

`func (o *DeploymentItem) HasCurrentRuntimeVersion() bool`

HasCurrentRuntimeVersion returns a boolean if a field has been set.

### GetLastSuccessfulRuntimeVersion

`func (o *DeploymentItem) GetLastSuccessfulRuntimeVersion() string`

GetLastSuccessfulRuntimeVersion returns the LastSuccessfulRuntimeVersion field if non-nil, zero value otherwise.

### GetLastSuccessfulRuntimeVersionOk

`func (o *DeploymentItem) GetLastSuccessfulRuntimeVersionOk() (*string, bool)`

GetLastSuccessfulRuntimeVersionOk returns a tuple with the LastSuccessfulRuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulRuntimeVersion

`func (o *DeploymentItem) SetLastSuccessfulRuntimeVersion(v string)`

SetLastSuccessfulRuntimeVersion sets LastSuccessfulRuntimeVersion field to given value.

### HasLastSuccessfulRuntimeVersion

`func (o *DeploymentItem) HasLastSuccessfulRuntimeVersion() bool`

HasLastSuccessfulRuntimeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


