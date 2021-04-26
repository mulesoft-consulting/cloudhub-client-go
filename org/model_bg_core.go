/*
 * Organization API
 *
 * Description of the Organization API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// BGCore struct for BGCore
type BGCore struct {
	// An explanation about the purpose of this instance.
	ClientId string `json:"clientId"`
	// An explanation about the purpose of this instance.
	CreatedAt string `json:"createdAt"`
	// An explanation about the purpose of this instance.
	Domain string `json:"domain"`
	// An explanation about the purpose of this instance.
	Entitlements map[string]map[string]interface{} `json:"entitlements"`
	// An explanation about the purpose of this instance.
	Environments []AnyOfmap `json:"environments"`
	// An explanation about the purpose of this instance.
	Id string `json:"id"`
	// An explanation about the purpose of this instance.
	IdproviderId string `json:"idprovider_id"`
	// An explanation about the purpose of this instance.
	IsAutomaticAdminPromotionExempt bool `json:"isAutomaticAdminPromotionExempt"`
	// An explanation about the purpose of this instance.
	IsFederated bool `json:"isFederated"`
	// An explanation about the purpose of this instance.
	IsMaster bool `json:"isMaster"`
	// An explanation about the purpose of this instance.
	MfaRequired string `json:"mfaRequired"`
	// An explanation about the purpose of this instance.
	Name string `json:"name"`
	// An explanation about the purpose of this instance.
	OwnerId string `json:"ownerId"`
	// An explanation about the purpose of this instance.
	ParentOrganizationIds []string `json:"parentOrganizationIds"`
	// An explanation about the purpose of this instance.
	Properties map[string]map[string]interface{} `json:"properties"`
	// An explanation about the purpose of this instance.
	SubOrganizationIds []AnyOfstring `json:"subOrganizationIds"`
	// An explanation about the purpose of this instance.
	TenantOrganizationIds []string `json:"tenantOrganizationIds"`
	// An explanation about the purpose of this instance.
	UpdatedAt string `json:"updatedAt"`
}

// NewBGCore instantiates a new BGCore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBGCore(clientId string, createdAt string, domain string, entitlements map[string]map[string]interface{}, environments []AnyOfmap, id string, idproviderId string, isAutomaticAdminPromotionExempt bool, isFederated bool, isMaster bool, mfaRequired string, name string, ownerId string, parentOrganizationIds []string, properties map[string]map[string]interface{}, subOrganizationIds []AnyOfstring, tenantOrganizationIds []string, updatedAt string) *BGCore {
	this := BGCore{}
	this.ClientId = clientId
	this.CreatedAt = createdAt
	this.Domain = domain
	this.Entitlements = entitlements
	this.Environments = environments
	this.Id = id
	this.IdproviderId = idproviderId
	this.IsAutomaticAdminPromotionExempt = isAutomaticAdminPromotionExempt
	this.IsFederated = isFederated
	this.IsMaster = isMaster
	this.MfaRequired = mfaRequired
	this.Name = name
	this.OwnerId = ownerId
	this.ParentOrganizationIds = parentOrganizationIds
	this.Properties = properties
	this.SubOrganizationIds = subOrganizationIds
	this.TenantOrganizationIds = tenantOrganizationIds
	this.UpdatedAt = updatedAt
	return &this
}

// NewBGCoreWithDefaults instantiates a new BGCore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBGCoreWithDefaults() *BGCore {
	this := BGCore{}
	var clientId string = ""
	this.ClientId = clientId
	var createdAt string = ""
	this.CreatedAt = createdAt
	var domain string = ""
	this.Domain = domain
	var id string = ""
	this.Id = id
	var idproviderId string = ""
	this.IdproviderId = idproviderId
	var isAutomaticAdminPromotionExempt bool = false
	this.IsAutomaticAdminPromotionExempt = isAutomaticAdminPromotionExempt
	var isFederated bool = false
	this.IsFederated = isFederated
	var isMaster bool = false
	this.IsMaster = isMaster
	var mfaRequired string = ""
	this.MfaRequired = mfaRequired
	var name string = ""
	this.Name = name
	var ownerId string = ""
	this.OwnerId = ownerId
	var updatedAt string = ""
	this.UpdatedAt = updatedAt
	return &this
}

// GetClientId returns the ClientId field value
func (o *BGCore) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *BGCore) SetClientId(v string) {
	o.ClientId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BGCore) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BGCore) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDomain returns the Domain field value
func (o *BGCore) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *BGCore) SetDomain(v string) {
	o.Domain = v
}

// GetEntitlements returns the Entitlements field value
func (o *BGCore) GetEntitlements() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetEntitlementsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entitlements, true
}

// SetEntitlements sets field value
func (o *BGCore) SetEntitlements(v map[string]map[string]interface{}) {
	o.Entitlements = v
}

// GetEnvironments returns the Environments field value
func (o *BGCore) GetEnvironments() []AnyOfmap {
	if o == nil {
		var ret []AnyOfmap
		return ret
	}

	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetEnvironmentsOk() (*[]AnyOfmap, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environments, true
}

// SetEnvironments sets field value
func (o *BGCore) SetEnvironments(v []AnyOfmap) {
	o.Environments = v
}

// GetId returns the Id field value
func (o *BGCore) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BGCore) SetId(v string) {
	o.Id = v
}

// GetIdproviderId returns the IdproviderId field value
func (o *BGCore) GetIdproviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdproviderId
}

// GetIdproviderIdOk returns a tuple with the IdproviderId field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetIdproviderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdproviderId, true
}

// SetIdproviderId sets field value
func (o *BGCore) SetIdproviderId(v string) {
	o.IdproviderId = v
}

// GetIsAutomaticAdminPromotionExempt returns the IsAutomaticAdminPromotionExempt field value
func (o *BGCore) GetIsAutomaticAdminPromotionExempt() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutomaticAdminPromotionExempt
}

// GetIsAutomaticAdminPromotionExemptOk returns a tuple with the IsAutomaticAdminPromotionExempt field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetIsAutomaticAdminPromotionExemptOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsAutomaticAdminPromotionExempt, true
}

// SetIsAutomaticAdminPromotionExempt sets field value
func (o *BGCore) SetIsAutomaticAdminPromotionExempt(v bool) {
	o.IsAutomaticAdminPromotionExempt = v
}

// GetIsFederated returns the IsFederated field value
func (o *BGCore) GetIsFederated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFederated
}

// GetIsFederatedOk returns a tuple with the IsFederated field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetIsFederatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsFederated, true
}

// SetIsFederated sets field value
func (o *BGCore) SetIsFederated(v bool) {
	o.IsFederated = v
}

// GetIsMaster returns the IsMaster field value
func (o *BGCore) GetIsMaster() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMaster
}

// GetIsMasterOk returns a tuple with the IsMaster field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetIsMasterOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsMaster, true
}

// SetIsMaster sets field value
func (o *BGCore) SetIsMaster(v bool) {
	o.IsMaster = v
}

// GetMfaRequired returns the MfaRequired field value
func (o *BGCore) GetMfaRequired() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MfaRequired
}

// GetMfaRequiredOk returns a tuple with the MfaRequired field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetMfaRequiredOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MfaRequired, true
}

// SetMfaRequired sets field value
func (o *BGCore) SetMfaRequired(v string) {
	o.MfaRequired = v
}

// GetName returns the Name field value
func (o *BGCore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BGCore) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *BGCore) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *BGCore) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetParentOrganizationIds returns the ParentOrganizationIds field value
func (o *BGCore) GetParentOrganizationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ParentOrganizationIds
}

// GetParentOrganizationIdsOk returns a tuple with the ParentOrganizationIds field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetParentOrganizationIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParentOrganizationIds, true
}

// SetParentOrganizationIds sets field value
func (o *BGCore) SetParentOrganizationIds(v []string) {
	o.ParentOrganizationIds = v
}

// GetProperties returns the Properties field value
func (o *BGCore) GetProperties() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetPropertiesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *BGCore) SetProperties(v map[string]map[string]interface{}) {
	o.Properties = v
}

// GetSubOrganizationIds returns the SubOrganizationIds field value
func (o *BGCore) GetSubOrganizationIds() []AnyOfstring {
	if o == nil {
		var ret []AnyOfstring
		return ret
	}

	return o.SubOrganizationIds
}

// GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetSubOrganizationIdsOk() (*[]AnyOfstring, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubOrganizationIds, true
}

// SetSubOrganizationIds sets field value
func (o *BGCore) SetSubOrganizationIds(v []AnyOfstring) {
	o.SubOrganizationIds = v
}

// GetTenantOrganizationIds returns the TenantOrganizationIds field value
func (o *BGCore) GetTenantOrganizationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TenantOrganizationIds
}

// GetTenantOrganizationIdsOk returns a tuple with the TenantOrganizationIds field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetTenantOrganizationIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantOrganizationIds, true
}

// SetTenantOrganizationIds sets field value
func (o *BGCore) SetTenantOrganizationIds(v []string) {
	o.TenantOrganizationIds = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BGCore) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BGCore) GetUpdatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BGCore) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o BGCore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["entitlements"] = o.Entitlements
	}
	if true {
		toSerialize["environments"] = o.Environments
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["idprovider_id"] = o.IdproviderId
	}
	if true {
		toSerialize["isAutomaticAdminPromotionExempt"] = o.IsAutomaticAdminPromotionExempt
	}
	if true {
		toSerialize["isFederated"] = o.IsFederated
	}
	if true {
		toSerialize["isMaster"] = o.IsMaster
	}
	if true {
		toSerialize["mfaRequired"] = o.MfaRequired
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ownerId"] = o.OwnerId
	}
	if true {
		toSerialize["parentOrganizationIds"] = o.ParentOrganizationIds
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["subOrganizationIds"] = o.SubOrganizationIds
	}
	if true {
		toSerialize["tenantOrganizationIds"] = o.TenantOrganizationIds
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableBGCore struct {
	value *BGCore
	isSet bool
}

func (v NullableBGCore) Get() *BGCore {
	return v.value
}

func (v *NullableBGCore) Set(val *BGCore) {
	v.value = val
	v.isSet = true
}

func (v NullableBGCore) IsSet() bool {
	return v.isSet
}

func (v *NullableBGCore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBGCore(val *BGCore) *NullableBGCore {
	return &NullableBGCore{value: val, isSet: true}
}

func (v NullableBGCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBGCore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

