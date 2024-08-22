# Http

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inbound** | Pointer to [**HttpInbound**](HttpInbound.md) |  | [optional] 

## Methods

### NewHttp

`func NewHttp() *Http`

NewHttp instantiates a new Http object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpWithDefaults

`func NewHttpWithDefaults() *Http`

NewHttpWithDefaults instantiates a new Http object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbound

`func (o *Http) GetInbound() HttpInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *Http) GetInboundOk() (*HttpInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *Http) SetInbound(v HttpInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *Http) HasInbound() bool`

HasInbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


