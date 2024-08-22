# HttpInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicUrl** | Pointer to **string** |  | [optional] 
**PathRewrite** | Pointer to **string** | This field is not supported on deployments to Shared Spaces. Requests containing values other than null will be rejected. | [optional] 
**LastMileSecurity** | Pointer to **bool** |  | [optional] 
**ForwardSslSession** | Pointer to **bool** |  | [optional] 
**InternalUrl** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpInbound

`func NewHttpInbound() *HttpInbound`

NewHttpInbound instantiates a new HttpInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpInboundWithDefaults

`func NewHttpInboundWithDefaults() *HttpInbound`

NewHttpInboundWithDefaults instantiates a new HttpInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicUrl

`func (o *HttpInbound) GetPublicUrl() string`

GetPublicUrl returns the PublicUrl field if non-nil, zero value otherwise.

### GetPublicUrlOk

`func (o *HttpInbound) GetPublicUrlOk() (*string, bool)`

GetPublicUrlOk returns a tuple with the PublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUrl

`func (o *HttpInbound) SetPublicUrl(v string)`

SetPublicUrl sets PublicUrl field to given value.

### HasPublicUrl

`func (o *HttpInbound) HasPublicUrl() bool`

HasPublicUrl returns a boolean if a field has been set.

### GetPathRewrite

`func (o *HttpInbound) GetPathRewrite() string`

GetPathRewrite returns the PathRewrite field if non-nil, zero value otherwise.

### GetPathRewriteOk

`func (o *HttpInbound) GetPathRewriteOk() (*string, bool)`

GetPathRewriteOk returns a tuple with the PathRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathRewrite

`func (o *HttpInbound) SetPathRewrite(v string)`

SetPathRewrite sets PathRewrite field to given value.

### HasPathRewrite

`func (o *HttpInbound) HasPathRewrite() bool`

HasPathRewrite returns a boolean if a field has been set.

### GetLastMileSecurity

`func (o *HttpInbound) GetLastMileSecurity() bool`

GetLastMileSecurity returns the LastMileSecurity field if non-nil, zero value otherwise.

### GetLastMileSecurityOk

`func (o *HttpInbound) GetLastMileSecurityOk() (*bool, bool)`

GetLastMileSecurityOk returns a tuple with the LastMileSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMileSecurity

`func (o *HttpInbound) SetLastMileSecurity(v bool)`

SetLastMileSecurity sets LastMileSecurity field to given value.

### HasLastMileSecurity

`func (o *HttpInbound) HasLastMileSecurity() bool`

HasLastMileSecurity returns a boolean if a field has been set.

### GetForwardSslSession

`func (o *HttpInbound) GetForwardSslSession() bool`

GetForwardSslSession returns the ForwardSslSession field if non-nil, zero value otherwise.

### GetForwardSslSessionOk

`func (o *HttpInbound) GetForwardSslSessionOk() (*bool, bool)`

GetForwardSslSessionOk returns a tuple with the ForwardSslSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardSslSession

`func (o *HttpInbound) SetForwardSslSession(v bool)`

SetForwardSslSession sets ForwardSslSession field to given value.

### HasForwardSslSession

`func (o *HttpInbound) HasForwardSslSession() bool`

HasForwardSslSession returns a boolean if a field has been set.

### GetInternalUrl

`func (o *HttpInbound) GetInternalUrl() string`

GetInternalUrl returns the InternalUrl field if non-nil, zero value otherwise.

### GetInternalUrlOk

`func (o *HttpInbound) GetInternalUrlOk() (*string, bool)`

GetInternalUrlOk returns a tuple with the InternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalUrl

`func (o *HttpInbound) SetInternalUrl(v string)`

SetInternalUrl sets InternalUrl field to given value.

### HasInternalUrl

`func (o *HttpInbound) HasInternalUrl() bool`

HasInternalUrl returns a boolean if a field has been set.

### GetUniqueId

`func (o *HttpInbound) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *HttpInbound) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *HttpInbound) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *HttpInbound) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


