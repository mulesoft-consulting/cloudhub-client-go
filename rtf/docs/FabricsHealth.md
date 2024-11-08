# FabricsHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterMonitoring** | Pointer to [**FabricsHealthStatus**](FabricsHealthStatus.md) |  | [optional] 
**ManageDeployments** | Pointer to [**FabricsHealthStatus**](FabricsHealthStatus.md) |  | [optional] 
**LoadBalancing** | Pointer to [**FabricsHealthStatus**](FabricsHealthStatus.md) |  | [optional] 
**AnypointMonitoring** | Pointer to [**FabricsHealthStatus**](FabricsHealthStatus.md) |  | [optional] 
**ExternalLogForwarding** | Pointer to [**FabricsHealthStatus**](FabricsHealthStatus.md) |  | [optional] 
**Appliance** | Pointer to [**FabricsHealthStatus**](FabricsHealthStatus.md) |  | [optional] 
**Infrastructure** | Pointer to [**FabricsHealthStatus**](FabricsHealthStatus.md) |  | [optional] 
**PersistentGateway** | Pointer to [**FabricsHealthStatus**](FabricsHealthStatus.md) |  | [optional] 

## Methods

### NewFabricsHealth

`func NewFabricsHealth() *FabricsHealth`

NewFabricsHealth instantiates a new FabricsHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricsHealthWithDefaults

`func NewFabricsHealthWithDefaults() *FabricsHealth`

NewFabricsHealthWithDefaults instantiates a new FabricsHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterMonitoring

`func (o *FabricsHealth) GetClusterMonitoring() FabricsHealthStatus`

GetClusterMonitoring returns the ClusterMonitoring field if non-nil, zero value otherwise.

### GetClusterMonitoringOk

`func (o *FabricsHealth) GetClusterMonitoringOk() (*FabricsHealthStatus, bool)`

GetClusterMonitoringOk returns a tuple with the ClusterMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMonitoring

`func (o *FabricsHealth) SetClusterMonitoring(v FabricsHealthStatus)`

SetClusterMonitoring sets ClusterMonitoring field to given value.

### HasClusterMonitoring

`func (o *FabricsHealth) HasClusterMonitoring() bool`

HasClusterMonitoring returns a boolean if a field has been set.

### GetManageDeployments

`func (o *FabricsHealth) GetManageDeployments() FabricsHealthStatus`

GetManageDeployments returns the ManageDeployments field if non-nil, zero value otherwise.

### GetManageDeploymentsOk

`func (o *FabricsHealth) GetManageDeploymentsOk() (*FabricsHealthStatus, bool)`

GetManageDeploymentsOk returns a tuple with the ManageDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageDeployments

`func (o *FabricsHealth) SetManageDeployments(v FabricsHealthStatus)`

SetManageDeployments sets ManageDeployments field to given value.

### HasManageDeployments

`func (o *FabricsHealth) HasManageDeployments() bool`

HasManageDeployments returns a boolean if a field has been set.

### GetLoadBalancing

`func (o *FabricsHealth) GetLoadBalancing() FabricsHealthStatus`

GetLoadBalancing returns the LoadBalancing field if non-nil, zero value otherwise.

### GetLoadBalancingOk

`func (o *FabricsHealth) GetLoadBalancingOk() (*FabricsHealthStatus, bool)`

GetLoadBalancingOk returns a tuple with the LoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancing

`func (o *FabricsHealth) SetLoadBalancing(v FabricsHealthStatus)`

SetLoadBalancing sets LoadBalancing field to given value.

### HasLoadBalancing

`func (o *FabricsHealth) HasLoadBalancing() bool`

HasLoadBalancing returns a boolean if a field has been set.

### GetAnypointMonitoring

`func (o *FabricsHealth) GetAnypointMonitoring() FabricsHealthStatus`

GetAnypointMonitoring returns the AnypointMonitoring field if non-nil, zero value otherwise.

### GetAnypointMonitoringOk

`func (o *FabricsHealth) GetAnypointMonitoringOk() (*FabricsHealthStatus, bool)`

GetAnypointMonitoringOk returns a tuple with the AnypointMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnypointMonitoring

`func (o *FabricsHealth) SetAnypointMonitoring(v FabricsHealthStatus)`

SetAnypointMonitoring sets AnypointMonitoring field to given value.

### HasAnypointMonitoring

`func (o *FabricsHealth) HasAnypointMonitoring() bool`

HasAnypointMonitoring returns a boolean if a field has been set.

### GetExternalLogForwarding

`func (o *FabricsHealth) GetExternalLogForwarding() FabricsHealthStatus`

GetExternalLogForwarding returns the ExternalLogForwarding field if non-nil, zero value otherwise.

### GetExternalLogForwardingOk

`func (o *FabricsHealth) GetExternalLogForwardingOk() (*FabricsHealthStatus, bool)`

GetExternalLogForwardingOk returns a tuple with the ExternalLogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogForwarding

`func (o *FabricsHealth) SetExternalLogForwarding(v FabricsHealthStatus)`

SetExternalLogForwarding sets ExternalLogForwarding field to given value.

### HasExternalLogForwarding

`func (o *FabricsHealth) HasExternalLogForwarding() bool`

HasExternalLogForwarding returns a boolean if a field has been set.

### GetAppliance

`func (o *FabricsHealth) GetAppliance() FabricsHealthStatus`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *FabricsHealth) GetApplianceOk() (*FabricsHealthStatus, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *FabricsHealth) SetAppliance(v FabricsHealthStatus)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *FabricsHealth) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetInfrastructure

`func (o *FabricsHealth) GetInfrastructure() FabricsHealthStatus`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *FabricsHealth) GetInfrastructureOk() (*FabricsHealthStatus, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *FabricsHealth) SetInfrastructure(v FabricsHealthStatus)`

SetInfrastructure sets Infrastructure field to given value.

### HasInfrastructure

`func (o *FabricsHealth) HasInfrastructure() bool`

HasInfrastructure returns a boolean if a field has been set.

### GetPersistentGateway

`func (o *FabricsHealth) GetPersistentGateway() FabricsHealthStatus`

GetPersistentGateway returns the PersistentGateway field if non-nil, zero value otherwise.

### GetPersistentGatewayOk

`func (o *FabricsHealth) GetPersistentGatewayOk() (*FabricsHealthStatus, bool)`

GetPersistentGatewayOk returns a tuple with the PersistentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGateway

`func (o *FabricsHealth) SetPersistentGateway(v FabricsHealthStatus)`

SetPersistentGateway sets PersistentGateway field to given value.

### HasPersistentGateway

`func (o *FabricsHealth) HasPersistentGateway() bool`

HasPersistentGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


