# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the application. | [optional] 
**DesiredState** | Pointer to **NullableString** | The desired state of the application. | [optional] 
**Ref** | Pointer to [**Ref**](Ref.md) |  | [optional] 
**Configuration** | Pointer to [**AppConfiguration**](AppConfiguration.md) |  | [optional] 
**ResourceAssets** | Pointer to **map[string]interface{}** |  | [optional] 
**VCores** | Pointer to **float32** | The application allocated virtual cores. | [optional] 
**Integrations** | Pointer to [**ApplicationIntegrations**](ApplicationIntegrations.md) |  | [optional] 
**ObjectStoreV2Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Application) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Application) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Application) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Application) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDesiredState

`func (o *Application) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *Application) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *Application) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *Application) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### SetDesiredStateNil

`func (o *Application) SetDesiredStateNil(b bool)`

 SetDesiredStateNil sets the value for DesiredState to be an explicit nil

### UnsetDesiredState
`func (o *Application) UnsetDesiredState()`

UnsetDesiredState ensures that no value is present for DesiredState, not even an explicit nil
### GetRef

`func (o *Application) GetRef() Ref`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Application) GetRefOk() (*Ref, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Application) SetRef(v Ref)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Application) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *Application) GetConfiguration() AppConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Application) GetConfigurationOk() (*AppConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Application) SetConfiguration(v AppConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Application) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetResourceAssets

`func (o *Application) GetResourceAssets() map[string]interface{}`

GetResourceAssets returns the ResourceAssets field if non-nil, zero value otherwise.

### GetResourceAssetsOk

`func (o *Application) GetResourceAssetsOk() (*map[string]interface{}, bool)`

GetResourceAssetsOk returns a tuple with the ResourceAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAssets

`func (o *Application) SetResourceAssets(v map[string]interface{})`

SetResourceAssets sets ResourceAssets field to given value.

### HasResourceAssets

`func (o *Application) HasResourceAssets() bool`

HasResourceAssets returns a boolean if a field has been set.

### GetVCores

`func (o *Application) GetVCores() float32`

GetVCores returns the VCores field if non-nil, zero value otherwise.

### GetVCoresOk

`func (o *Application) GetVCoresOk() (*float32, bool)`

GetVCoresOk returns a tuple with the VCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCores

`func (o *Application) SetVCores(v float32)`

SetVCores sets VCores field to given value.

### HasVCores

`func (o *Application) HasVCores() bool`

HasVCores returns a boolean if a field has been set.

### GetIntegrations

`func (o *Application) GetIntegrations() ApplicationIntegrations`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *Application) GetIntegrationsOk() (*ApplicationIntegrations, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *Application) SetIntegrations(v ApplicationIntegrations)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *Application) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetObjectStoreV2Enabled

`func (o *Application) GetObjectStoreV2Enabled() bool`

GetObjectStoreV2Enabled returns the ObjectStoreV2Enabled field if non-nil, zero value otherwise.

### GetObjectStoreV2EnabledOk

`func (o *Application) GetObjectStoreV2EnabledOk() (*bool, bool)`

GetObjectStoreV2EnabledOk returns a tuple with the ObjectStoreV2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreV2Enabled

`func (o *Application) SetObjectStoreV2Enabled(v bool)`

SetObjectStoreV2Enabled sets ObjectStoreV2Enabled field to given value.

### HasObjectStoreV2Enabled

`func (o *Application) HasObjectStoreV2Enabled() bool`

HasObjectStoreV2Enabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


