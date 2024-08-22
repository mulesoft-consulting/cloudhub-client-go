# DeploymentRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**Target**](Target.md) |  | [optional] 
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**DesiredVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentRequestBody

`func NewDeploymentRequestBody() *DeploymentRequestBody`

NewDeploymentRequestBody instantiates a new DeploymentRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRequestBodyWithDefaults

`func NewDeploymentRequestBodyWithDefaults() *DeploymentRequestBody`

NewDeploymentRequestBodyWithDefaults instantiates a new DeploymentRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *DeploymentRequestBody) GetTarget() Target`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DeploymentRequestBody) GetTargetOk() (*Target, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DeploymentRequestBody) SetTarget(v Target)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DeploymentRequestBody) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetApplication

`func (o *DeploymentRequestBody) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DeploymentRequestBody) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DeploymentRequestBody) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *DeploymentRequestBody) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDesiredVersion

`func (o *DeploymentRequestBody) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *DeploymentRequestBody) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *DeploymentRequestBody) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *DeploymentRequestBody) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


