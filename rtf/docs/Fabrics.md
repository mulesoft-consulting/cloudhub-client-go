# Fabrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**VendorMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DesiredVersion** | Pointer to **string** |  | [optional] 
**AvailableUpgradeVersion** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Upgrade** | Pointer to [**FabricsUpgrade**](FabricsUpgrade.md) |  | [optional] 
**Nodes** | Pointer to [**[]FabricsNode**](FabricsNode.md) |  | [optional] 
**ActivationData** | Pointer to **string** |  | [optional] 
**SecondsSinceHeartbeat** | Pointer to **int32** |  | [optional] 
**KubernetesVersion** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**LicenseExpiryDate** | Pointer to **int64** |  | [optional] 
**IsManaged** | Pointer to **bool** |  | [optional] 
**IsHelmManaged** | Pointer to **bool** |  | [optional] 
**AppScopedLogForwarding** | Pointer to **bool** |  | [optional] 
**ClusterConfigurationLevel** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**Features**](Features.md) |  | [optional] 
**Ingress** | Pointer to [**Ingress**](Ingress.md) |  | [optional] 

## Methods

### NewFabrics

`func NewFabrics() *Fabrics`

NewFabrics instantiates a new Fabrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricsWithDefaults

`func NewFabricsWithDefaults() *Fabrics`

NewFabricsWithDefaults instantiates a new Fabrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Fabrics) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fabrics) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fabrics) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Fabrics) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Fabrics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Fabrics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Fabrics) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Fabrics) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *Fabrics) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Fabrics) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Fabrics) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Fabrics) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVendor

`func (o *Fabrics) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Fabrics) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Fabrics) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Fabrics) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorMetadata

`func (o *Fabrics) GetVendorMetadata() map[string]interface{}`

GetVendorMetadata returns the VendorMetadata field if non-nil, zero value otherwise.

### GetVendorMetadataOk

`func (o *Fabrics) GetVendorMetadataOk() (*map[string]interface{}, bool)`

GetVendorMetadataOk returns a tuple with the VendorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorMetadata

`func (o *Fabrics) SetVendorMetadata(v map[string]interface{})`

SetVendorMetadata sets VendorMetadata field to given value.

### HasVendorMetadata

`func (o *Fabrics) HasVendorMetadata() bool`

HasVendorMetadata returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Fabrics) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Fabrics) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Fabrics) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Fabrics) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetVersion

`func (o *Fabrics) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Fabrics) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Fabrics) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Fabrics) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *Fabrics) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Fabrics) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Fabrics) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Fabrics) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDesiredVersion

`func (o *Fabrics) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *Fabrics) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *Fabrics) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *Fabrics) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.

### GetAvailableUpgradeVersion

`func (o *Fabrics) GetAvailableUpgradeVersion() string`

GetAvailableUpgradeVersion returns the AvailableUpgradeVersion field if non-nil, zero value otherwise.

### GetAvailableUpgradeVersionOk

`func (o *Fabrics) GetAvailableUpgradeVersionOk() (*string, bool)`

GetAvailableUpgradeVersionOk returns a tuple with the AvailableUpgradeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUpgradeVersion

`func (o *Fabrics) SetAvailableUpgradeVersion(v string)`

SetAvailableUpgradeVersion sets AvailableUpgradeVersion field to given value.

### HasAvailableUpgradeVersion

`func (o *Fabrics) HasAvailableUpgradeVersion() bool`

HasAvailableUpgradeVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Fabrics) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Fabrics) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Fabrics) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Fabrics) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpgrade

`func (o *Fabrics) GetUpgrade() FabricsUpgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *Fabrics) GetUpgradeOk() (*FabricsUpgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *Fabrics) SetUpgrade(v FabricsUpgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *Fabrics) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetNodes

`func (o *Fabrics) GetNodes() []FabricsNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Fabrics) GetNodesOk() (*[]FabricsNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Fabrics) SetNodes(v []FabricsNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Fabrics) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetActivationData

`func (o *Fabrics) GetActivationData() string`

GetActivationData returns the ActivationData field if non-nil, zero value otherwise.

### GetActivationDataOk

`func (o *Fabrics) GetActivationDataOk() (*string, bool)`

GetActivationDataOk returns a tuple with the ActivationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationData

`func (o *Fabrics) SetActivationData(v string)`

SetActivationData sets ActivationData field to given value.

### HasActivationData

`func (o *Fabrics) HasActivationData() bool`

HasActivationData returns a boolean if a field has been set.

### GetSecondsSinceHeartbeat

`func (o *Fabrics) GetSecondsSinceHeartbeat() int32`

GetSecondsSinceHeartbeat returns the SecondsSinceHeartbeat field if non-nil, zero value otherwise.

### GetSecondsSinceHeartbeatOk

`func (o *Fabrics) GetSecondsSinceHeartbeatOk() (*int32, bool)`

GetSecondsSinceHeartbeatOk returns a tuple with the SecondsSinceHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsSinceHeartbeat

`func (o *Fabrics) SetSecondsSinceHeartbeat(v int32)`

SetSecondsSinceHeartbeat sets SecondsSinceHeartbeat field to given value.

### HasSecondsSinceHeartbeat

`func (o *Fabrics) HasSecondsSinceHeartbeat() bool`

HasSecondsSinceHeartbeat returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *Fabrics) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *Fabrics) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *Fabrics) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *Fabrics) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetNamespace

`func (o *Fabrics) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Fabrics) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Fabrics) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Fabrics) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetLicenseExpiryDate

`func (o *Fabrics) GetLicenseExpiryDate() int64`

GetLicenseExpiryDate returns the LicenseExpiryDate field if non-nil, zero value otherwise.

### GetLicenseExpiryDateOk

`func (o *Fabrics) GetLicenseExpiryDateOk() (*int64, bool)`

GetLicenseExpiryDateOk returns a tuple with the LicenseExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiryDate

`func (o *Fabrics) SetLicenseExpiryDate(v int64)`

SetLicenseExpiryDate sets LicenseExpiryDate field to given value.

### HasLicenseExpiryDate

`func (o *Fabrics) HasLicenseExpiryDate() bool`

HasLicenseExpiryDate returns a boolean if a field has been set.

### GetIsManaged

`func (o *Fabrics) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *Fabrics) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *Fabrics) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *Fabrics) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetIsHelmManaged

`func (o *Fabrics) GetIsHelmManaged() bool`

GetIsHelmManaged returns the IsHelmManaged field if non-nil, zero value otherwise.

### GetIsHelmManagedOk

`func (o *Fabrics) GetIsHelmManagedOk() (*bool, bool)`

GetIsHelmManagedOk returns a tuple with the IsHelmManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHelmManaged

`func (o *Fabrics) SetIsHelmManaged(v bool)`

SetIsHelmManaged sets IsHelmManaged field to given value.

### HasIsHelmManaged

`func (o *Fabrics) HasIsHelmManaged() bool`

HasIsHelmManaged returns a boolean if a field has been set.

### GetAppScopedLogForwarding

`func (o *Fabrics) GetAppScopedLogForwarding() bool`

GetAppScopedLogForwarding returns the AppScopedLogForwarding field if non-nil, zero value otherwise.

### GetAppScopedLogForwardingOk

`func (o *Fabrics) GetAppScopedLogForwardingOk() (*bool, bool)`

GetAppScopedLogForwardingOk returns a tuple with the AppScopedLogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppScopedLogForwarding

`func (o *Fabrics) SetAppScopedLogForwarding(v bool)`

SetAppScopedLogForwarding sets AppScopedLogForwarding field to given value.

### HasAppScopedLogForwarding

`func (o *Fabrics) HasAppScopedLogForwarding() bool`

HasAppScopedLogForwarding returns a boolean if a field has been set.

### GetClusterConfigurationLevel

`func (o *Fabrics) GetClusterConfigurationLevel() string`

GetClusterConfigurationLevel returns the ClusterConfigurationLevel field if non-nil, zero value otherwise.

### GetClusterConfigurationLevelOk

`func (o *Fabrics) GetClusterConfigurationLevelOk() (*string, bool)`

GetClusterConfigurationLevelOk returns a tuple with the ClusterConfigurationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConfigurationLevel

`func (o *Fabrics) SetClusterConfigurationLevel(v string)`

SetClusterConfigurationLevel sets ClusterConfigurationLevel field to given value.

### HasClusterConfigurationLevel

`func (o *Fabrics) HasClusterConfigurationLevel() bool`

HasClusterConfigurationLevel returns a boolean if a field has been set.

### GetFeatures

`func (o *Fabrics) GetFeatures() Features`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Fabrics) GetFeaturesOk() (*Features, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Fabrics) SetFeatures(v Features)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Fabrics) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetIngress

`func (o *Fabrics) GetIngress() Ingress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *Fabrics) GetIngressOk() (*Ingress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *Fabrics) SetIngress(v Ingress)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *Fabrics) HasIngress() bool`

HasIngress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


