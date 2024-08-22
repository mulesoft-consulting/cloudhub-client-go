# FabricsNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**KubeletVersion** | Pointer to **string** |  | [optional] 
**DockerVersion** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Capacity** | Pointer to [**Capacity**](Capacity.md) |  | [optional] 
**AllocatedRequestCapacity** | Pointer to [**AllocatedRequestCapacity**](AllocatedRequestCapacity.md) |  | [optional] 
**AllocatedLimitCapacity** | Pointer to [**AllocatedLimitCapacity**](AllocatedLimitCapacity.md) |  | [optional] 

## Methods

### NewFabricsNode

`func NewFabricsNode() *FabricsNode`

NewFabricsNode instantiates a new FabricsNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricsNodeWithDefaults

`func NewFabricsNodeWithDefaults() *FabricsNode`

NewFabricsNodeWithDefaults instantiates a new FabricsNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *FabricsNode) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *FabricsNode) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *FabricsNode) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *FabricsNode) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *FabricsNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricsNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricsNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricsNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKubeletVersion

`func (o *FabricsNode) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *FabricsNode) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *FabricsNode) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.

### HasKubeletVersion

`func (o *FabricsNode) HasKubeletVersion() bool`

HasKubeletVersion returns a boolean if a field has been set.

### GetDockerVersion

`func (o *FabricsNode) GetDockerVersion() string`

GetDockerVersion returns the DockerVersion field if non-nil, zero value otherwise.

### GetDockerVersionOk

`func (o *FabricsNode) GetDockerVersionOk() (*string, bool)`

GetDockerVersionOk returns a tuple with the DockerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerVersion

`func (o *FabricsNode) SetDockerVersion(v string)`

SetDockerVersion sets DockerVersion field to given value.

### HasDockerVersion

`func (o *FabricsNode) HasDockerVersion() bool`

HasDockerVersion returns a boolean if a field has been set.

### GetRole

`func (o *FabricsNode) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FabricsNode) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FabricsNode) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *FabricsNode) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *FabricsNode) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FabricsNode) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FabricsNode) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FabricsNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCapacity

`func (o *FabricsNode) GetCapacity() Capacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *FabricsNode) GetCapacityOk() (*Capacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *FabricsNode) SetCapacity(v Capacity)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *FabricsNode) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetAllocatedRequestCapacity

`func (o *FabricsNode) GetAllocatedRequestCapacity() AllocatedRequestCapacity`

GetAllocatedRequestCapacity returns the AllocatedRequestCapacity field if non-nil, zero value otherwise.

### GetAllocatedRequestCapacityOk

`func (o *FabricsNode) GetAllocatedRequestCapacityOk() (*AllocatedRequestCapacity, bool)`

GetAllocatedRequestCapacityOk returns a tuple with the AllocatedRequestCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedRequestCapacity

`func (o *FabricsNode) SetAllocatedRequestCapacity(v AllocatedRequestCapacity)`

SetAllocatedRequestCapacity sets AllocatedRequestCapacity field to given value.

### HasAllocatedRequestCapacity

`func (o *FabricsNode) HasAllocatedRequestCapacity() bool`

HasAllocatedRequestCapacity returns a boolean if a field has been set.

### GetAllocatedLimitCapacity

`func (o *FabricsNode) GetAllocatedLimitCapacity() AllocatedLimitCapacity`

GetAllocatedLimitCapacity returns the AllocatedLimitCapacity field if non-nil, zero value otherwise.

### GetAllocatedLimitCapacityOk

`func (o *FabricsNode) GetAllocatedLimitCapacityOk() (*AllocatedLimitCapacity, bool)`

GetAllocatedLimitCapacityOk returns a tuple with the AllocatedLimitCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedLimitCapacity

`func (o *FabricsNode) SetAllocatedLimitCapacity(v AllocatedLimitCapacity)`

SetAllocatedLimitCapacity sets AllocatedLimitCapacity field to given value.

### HasAllocatedLimitCapacity

`func (o *FabricsNode) HasAllocatedLimitCapacity() bool`

HasAllocatedLimitCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


