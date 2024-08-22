# DeploymentResponsePaged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]DeploymentItem**](DeploymentItem.md) |  | [optional] 

## Methods

### NewDeploymentResponsePaged

`func NewDeploymentResponsePaged() *DeploymentResponsePaged`

NewDeploymentResponsePaged instantiates a new DeploymentResponsePaged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentResponsePagedWithDefaults

`func NewDeploymentResponsePagedWithDefaults() *DeploymentResponsePaged`

NewDeploymentResponsePagedWithDefaults instantiates a new DeploymentResponsePaged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *DeploymentResponsePaged) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DeploymentResponsePaged) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DeploymentResponsePaged) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DeploymentResponsePaged) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *DeploymentResponsePaged) GetItems() []DeploymentItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeploymentResponsePaged) GetItemsOk() (*[]DeploymentItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeploymentResponsePaged) SetItems(v []DeploymentItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *DeploymentResponsePaged) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


