# DeploymentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clustered** | Pointer to **bool** |  | [optional] 
**EnforceDeployingReplicasAcrossNodes** | Pointer to **bool** |  | [optional] 
**Http** | Pointer to [**Http**](Http.md) |  | [optional] 
**Jvm** | Pointer to [**Jvm**](Jvm.md) |  | [optional] 
**Runtime** | Pointer to [**Runtime**](Runtime.md) |  | [optional] 
**Autoscaling** | Pointer to [**Autoscaling**](Autoscaling.md) |  | [optional] 
**UpdateStrategy** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**Resources**](Resources.md) |  | [optional] 
**LastMileSecurity** | Pointer to **bool** |  | [optional] 
**DisableAmLogForwarding** | Pointer to **bool** |  | [optional] 
**PersistentObjectStore** | Pointer to **bool** |  | [optional] 
**AnypointMonitoringScope** | Pointer to **string** |  | [optional] 
**Sidecars** | Pointer to [**Sidecars**](Sidecars.md) |  | [optional] 
**ForwardSslSession** | Pointer to **bool** |  | [optional] 
**DisableExternalLogForwarding** | Pointer to **bool** |  | [optional] 
**TracingEnabled** | Pointer to **bool** | Exposes tracing information in OpenTelemetry standard from Mule Runtime into all APM solutions. Supported only on:   * Mule Runtime versions 4.6 and above.   * RTF agent 2.6.33 and above. Learn more at https://docs.mulesoft.com/cloudhub-2/ch2-deploy-private-space#enable-trace-data-collection  | [optional] 
**GenerateDefaultPublicUrl** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeploymentSettings

`func NewDeploymentSettings() *DeploymentSettings`

NewDeploymentSettings instantiates a new DeploymentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentSettingsWithDefaults

`func NewDeploymentSettingsWithDefaults() *DeploymentSettings`

NewDeploymentSettingsWithDefaults instantiates a new DeploymentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClustered

`func (o *DeploymentSettings) GetClustered() bool`

GetClustered returns the Clustered field if non-nil, zero value otherwise.

### GetClusteredOk

`func (o *DeploymentSettings) GetClusteredOk() (*bool, bool)`

GetClusteredOk returns a tuple with the Clustered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClustered

`func (o *DeploymentSettings) SetClustered(v bool)`

SetClustered sets Clustered field to given value.

### HasClustered

`func (o *DeploymentSettings) HasClustered() bool`

HasClustered returns a boolean if a field has been set.

### GetEnforceDeployingReplicasAcrossNodes

`func (o *DeploymentSettings) GetEnforceDeployingReplicasAcrossNodes() bool`

GetEnforceDeployingReplicasAcrossNodes returns the EnforceDeployingReplicasAcrossNodes field if non-nil, zero value otherwise.

### GetEnforceDeployingReplicasAcrossNodesOk

`func (o *DeploymentSettings) GetEnforceDeployingReplicasAcrossNodesOk() (*bool, bool)`

GetEnforceDeployingReplicasAcrossNodesOk returns a tuple with the EnforceDeployingReplicasAcrossNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDeployingReplicasAcrossNodes

`func (o *DeploymentSettings) SetEnforceDeployingReplicasAcrossNodes(v bool)`

SetEnforceDeployingReplicasAcrossNodes sets EnforceDeployingReplicasAcrossNodes field to given value.

### HasEnforceDeployingReplicasAcrossNodes

`func (o *DeploymentSettings) HasEnforceDeployingReplicasAcrossNodes() bool`

HasEnforceDeployingReplicasAcrossNodes returns a boolean if a field has been set.

### GetHttp

`func (o *DeploymentSettings) GetHttp() Http`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *DeploymentSettings) GetHttpOk() (*Http, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *DeploymentSettings) SetHttp(v Http)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *DeploymentSettings) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetJvm

`func (o *DeploymentSettings) GetJvm() Jvm`

GetJvm returns the Jvm field if non-nil, zero value otherwise.

### GetJvmOk

`func (o *DeploymentSettings) GetJvmOk() (*Jvm, bool)`

GetJvmOk returns a tuple with the Jvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJvm

`func (o *DeploymentSettings) SetJvm(v Jvm)`

SetJvm sets Jvm field to given value.

### HasJvm

`func (o *DeploymentSettings) HasJvm() bool`

HasJvm returns a boolean if a field has been set.

### GetRuntime

`func (o *DeploymentSettings) GetRuntime() Runtime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *DeploymentSettings) GetRuntimeOk() (*Runtime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *DeploymentSettings) SetRuntime(v Runtime)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *DeploymentSettings) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetAutoscaling

`func (o *DeploymentSettings) GetAutoscaling() Autoscaling`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *DeploymentSettings) GetAutoscalingOk() (*Autoscaling, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *DeploymentSettings) SetAutoscaling(v Autoscaling)`

SetAutoscaling sets Autoscaling field to given value.

### HasAutoscaling

`func (o *DeploymentSettings) HasAutoscaling() bool`

HasAutoscaling returns a boolean if a field has been set.

### GetUpdateStrategy

`func (o *DeploymentSettings) GetUpdateStrategy() string`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *DeploymentSettings) GetUpdateStrategyOk() (*string, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *DeploymentSettings) SetUpdateStrategy(v string)`

SetUpdateStrategy sets UpdateStrategy field to given value.

### HasUpdateStrategy

`func (o *DeploymentSettings) HasUpdateStrategy() bool`

HasUpdateStrategy returns a boolean if a field has been set.

### GetResources

`func (o *DeploymentSettings) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DeploymentSettings) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DeploymentSettings) SetResources(v Resources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DeploymentSettings) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetLastMileSecurity

`func (o *DeploymentSettings) GetLastMileSecurity() bool`

GetLastMileSecurity returns the LastMileSecurity field if non-nil, zero value otherwise.

### GetLastMileSecurityOk

`func (o *DeploymentSettings) GetLastMileSecurityOk() (*bool, bool)`

GetLastMileSecurityOk returns a tuple with the LastMileSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMileSecurity

`func (o *DeploymentSettings) SetLastMileSecurity(v bool)`

SetLastMileSecurity sets LastMileSecurity field to given value.

### HasLastMileSecurity

`func (o *DeploymentSettings) HasLastMileSecurity() bool`

HasLastMileSecurity returns a boolean if a field has been set.

### GetDisableAmLogForwarding

`func (o *DeploymentSettings) GetDisableAmLogForwarding() bool`

GetDisableAmLogForwarding returns the DisableAmLogForwarding field if non-nil, zero value otherwise.

### GetDisableAmLogForwardingOk

`func (o *DeploymentSettings) GetDisableAmLogForwardingOk() (*bool, bool)`

GetDisableAmLogForwardingOk returns a tuple with the DisableAmLogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAmLogForwarding

`func (o *DeploymentSettings) SetDisableAmLogForwarding(v bool)`

SetDisableAmLogForwarding sets DisableAmLogForwarding field to given value.

### HasDisableAmLogForwarding

`func (o *DeploymentSettings) HasDisableAmLogForwarding() bool`

HasDisableAmLogForwarding returns a boolean if a field has been set.

### GetPersistentObjectStore

`func (o *DeploymentSettings) GetPersistentObjectStore() bool`

GetPersistentObjectStore returns the PersistentObjectStore field if non-nil, zero value otherwise.

### GetPersistentObjectStoreOk

`func (o *DeploymentSettings) GetPersistentObjectStoreOk() (*bool, bool)`

GetPersistentObjectStoreOk returns a tuple with the PersistentObjectStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentObjectStore

`func (o *DeploymentSettings) SetPersistentObjectStore(v bool)`

SetPersistentObjectStore sets PersistentObjectStore field to given value.

### HasPersistentObjectStore

`func (o *DeploymentSettings) HasPersistentObjectStore() bool`

HasPersistentObjectStore returns a boolean if a field has been set.

### GetAnypointMonitoringScope

`func (o *DeploymentSettings) GetAnypointMonitoringScope() string`

GetAnypointMonitoringScope returns the AnypointMonitoringScope field if non-nil, zero value otherwise.

### GetAnypointMonitoringScopeOk

`func (o *DeploymentSettings) GetAnypointMonitoringScopeOk() (*string, bool)`

GetAnypointMonitoringScopeOk returns a tuple with the AnypointMonitoringScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnypointMonitoringScope

`func (o *DeploymentSettings) SetAnypointMonitoringScope(v string)`

SetAnypointMonitoringScope sets AnypointMonitoringScope field to given value.

### HasAnypointMonitoringScope

`func (o *DeploymentSettings) HasAnypointMonitoringScope() bool`

HasAnypointMonitoringScope returns a boolean if a field has been set.

### GetSidecars

`func (o *DeploymentSettings) GetSidecars() Sidecars`

GetSidecars returns the Sidecars field if non-nil, zero value otherwise.

### GetSidecarsOk

`func (o *DeploymentSettings) GetSidecarsOk() (*Sidecars, bool)`

GetSidecarsOk returns a tuple with the Sidecars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidecars

`func (o *DeploymentSettings) SetSidecars(v Sidecars)`

SetSidecars sets Sidecars field to given value.

### HasSidecars

`func (o *DeploymentSettings) HasSidecars() bool`

HasSidecars returns a boolean if a field has been set.

### GetForwardSslSession

`func (o *DeploymentSettings) GetForwardSslSession() bool`

GetForwardSslSession returns the ForwardSslSession field if non-nil, zero value otherwise.

### GetForwardSslSessionOk

`func (o *DeploymentSettings) GetForwardSslSessionOk() (*bool, bool)`

GetForwardSslSessionOk returns a tuple with the ForwardSslSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardSslSession

`func (o *DeploymentSettings) SetForwardSslSession(v bool)`

SetForwardSslSession sets ForwardSslSession field to given value.

### HasForwardSslSession

`func (o *DeploymentSettings) HasForwardSslSession() bool`

HasForwardSslSession returns a boolean if a field has been set.

### GetDisableExternalLogForwarding

`func (o *DeploymentSettings) GetDisableExternalLogForwarding() bool`

GetDisableExternalLogForwarding returns the DisableExternalLogForwarding field if non-nil, zero value otherwise.

### GetDisableExternalLogForwardingOk

`func (o *DeploymentSettings) GetDisableExternalLogForwardingOk() (*bool, bool)`

GetDisableExternalLogForwardingOk returns a tuple with the DisableExternalLogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableExternalLogForwarding

`func (o *DeploymentSettings) SetDisableExternalLogForwarding(v bool)`

SetDisableExternalLogForwarding sets DisableExternalLogForwarding field to given value.

### HasDisableExternalLogForwarding

`func (o *DeploymentSettings) HasDisableExternalLogForwarding() bool`

HasDisableExternalLogForwarding returns a boolean if a field has been set.

### GetTracingEnabled

`func (o *DeploymentSettings) GetTracingEnabled() bool`

GetTracingEnabled returns the TracingEnabled field if non-nil, zero value otherwise.

### GetTracingEnabledOk

`func (o *DeploymentSettings) GetTracingEnabledOk() (*bool, bool)`

GetTracingEnabledOk returns a tuple with the TracingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracingEnabled

`func (o *DeploymentSettings) SetTracingEnabled(v bool)`

SetTracingEnabled sets TracingEnabled field to given value.

### HasTracingEnabled

`func (o *DeploymentSettings) HasTracingEnabled() bool`

HasTracingEnabled returns a boolean if a field has been set.

### GetGenerateDefaultPublicUrl

`func (o *DeploymentSettings) GetGenerateDefaultPublicUrl() bool`

GetGenerateDefaultPublicUrl returns the GenerateDefaultPublicUrl field if non-nil, zero value otherwise.

### GetGenerateDefaultPublicUrlOk

`func (o *DeploymentSettings) GetGenerateDefaultPublicUrlOk() (*bool, bool)`

GetGenerateDefaultPublicUrlOk returns a tuple with the GenerateDefaultPublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateDefaultPublicUrl

`func (o *DeploymentSettings) SetGenerateDefaultPublicUrl(v bool)`

SetGenerateDefaultPublicUrl sets GenerateDefaultPublicUrl field to given value.

### HasGenerateDefaultPublicUrl

`func (o *DeploymentSettings) HasGenerateDefaultPublicUrl() bool`

HasGenerateDefaultPublicUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


