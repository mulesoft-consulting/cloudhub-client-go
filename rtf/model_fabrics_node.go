/*
Runtime Fabrics API

Runtime Fabrics API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtf

import (
	"encoding/json"
)

// checks if the FabricsNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricsNode{}

// FabricsNode struct for FabricsNode
type FabricsNode struct {
	Uid *string `json:"uid,omitempty"`
	Name *string `json:"name,omitempty"`
	KubeletVersion *string `json:"kubeletVersion,omitempty"`
	DockerVersion *string `json:"dockerVersion,omitempty"`
	Role *string `json:"role,omitempty"`
	Status *Status `json:"status,omitempty"`
	Capacity *Capacity `json:"capacity,omitempty"`
	AllocatedRequestCapacity *AllocatedRequestCapacity `json:"allocatedRequestCapacity,omitempty"`
	AllocatedLimitCapacity *AllocatedLimitCapacity `json:"allocatedLimitCapacity,omitempty"`
}

// NewFabricsNode instantiates a new FabricsNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricsNode() *FabricsNode {
	this := FabricsNode{}
	return &this
}

// NewFabricsNodeWithDefaults instantiates a new FabricsNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricsNodeWithDefaults() *FabricsNode {
	this := FabricsNode{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *FabricsNode) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *FabricsNode) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *FabricsNode) SetUid(v string) {
	o.Uid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricsNode) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricsNode) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricsNode) SetName(v string) {
	o.Name = &v
}

// GetKubeletVersion returns the KubeletVersion field value if set, zero value otherwise.
func (o *FabricsNode) GetKubeletVersion() string {
	if o == nil || IsNil(o.KubeletVersion) {
		var ret string
		return ret
	}
	return *o.KubeletVersion
}

// GetKubeletVersionOk returns a tuple with the KubeletVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetKubeletVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KubeletVersion) {
		return nil, false
	}
	return o.KubeletVersion, true
}

// HasKubeletVersion returns a boolean if a field has been set.
func (o *FabricsNode) HasKubeletVersion() bool {
	if o != nil && !IsNil(o.KubeletVersion) {
		return true
	}

	return false
}

// SetKubeletVersion gets a reference to the given string and assigns it to the KubeletVersion field.
func (o *FabricsNode) SetKubeletVersion(v string) {
	o.KubeletVersion = &v
}

// GetDockerVersion returns the DockerVersion field value if set, zero value otherwise.
func (o *FabricsNode) GetDockerVersion() string {
	if o == nil || IsNil(o.DockerVersion) {
		var ret string
		return ret
	}
	return *o.DockerVersion
}

// GetDockerVersionOk returns a tuple with the DockerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetDockerVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DockerVersion) {
		return nil, false
	}
	return o.DockerVersion, true
}

// HasDockerVersion returns a boolean if a field has been set.
func (o *FabricsNode) HasDockerVersion() bool {
	if o != nil && !IsNil(o.DockerVersion) {
		return true
	}

	return false
}

// SetDockerVersion gets a reference to the given string and assigns it to the DockerVersion field.
func (o *FabricsNode) SetDockerVersion(v string) {
	o.DockerVersion = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *FabricsNode) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *FabricsNode) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *FabricsNode) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FabricsNode) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FabricsNode) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *FabricsNode) SetStatus(v Status) {
	o.Status = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *FabricsNode) GetCapacity() Capacity {
	if o == nil || IsNil(o.Capacity) {
		var ret Capacity
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetCapacityOk() (*Capacity, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *FabricsNode) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given Capacity and assigns it to the Capacity field.
func (o *FabricsNode) SetCapacity(v Capacity) {
	o.Capacity = &v
}

// GetAllocatedRequestCapacity returns the AllocatedRequestCapacity field value if set, zero value otherwise.
func (o *FabricsNode) GetAllocatedRequestCapacity() AllocatedRequestCapacity {
	if o == nil || IsNil(o.AllocatedRequestCapacity) {
		var ret AllocatedRequestCapacity
		return ret
	}
	return *o.AllocatedRequestCapacity
}

// GetAllocatedRequestCapacityOk returns a tuple with the AllocatedRequestCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetAllocatedRequestCapacityOk() (*AllocatedRequestCapacity, bool) {
	if o == nil || IsNil(o.AllocatedRequestCapacity) {
		return nil, false
	}
	return o.AllocatedRequestCapacity, true
}

// HasAllocatedRequestCapacity returns a boolean if a field has been set.
func (o *FabricsNode) HasAllocatedRequestCapacity() bool {
	if o != nil && !IsNil(o.AllocatedRequestCapacity) {
		return true
	}

	return false
}

// SetAllocatedRequestCapacity gets a reference to the given AllocatedRequestCapacity and assigns it to the AllocatedRequestCapacity field.
func (o *FabricsNode) SetAllocatedRequestCapacity(v AllocatedRequestCapacity) {
	o.AllocatedRequestCapacity = &v
}

// GetAllocatedLimitCapacity returns the AllocatedLimitCapacity field value if set, zero value otherwise.
func (o *FabricsNode) GetAllocatedLimitCapacity() AllocatedLimitCapacity {
	if o == nil || IsNil(o.AllocatedLimitCapacity) {
		var ret AllocatedLimitCapacity
		return ret
	}
	return *o.AllocatedLimitCapacity
}

// GetAllocatedLimitCapacityOk returns a tuple with the AllocatedLimitCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsNode) GetAllocatedLimitCapacityOk() (*AllocatedLimitCapacity, bool) {
	if o == nil || IsNil(o.AllocatedLimitCapacity) {
		return nil, false
	}
	return o.AllocatedLimitCapacity, true
}

// HasAllocatedLimitCapacity returns a boolean if a field has been set.
func (o *FabricsNode) HasAllocatedLimitCapacity() bool {
	if o != nil && !IsNil(o.AllocatedLimitCapacity) {
		return true
	}

	return false
}

// SetAllocatedLimitCapacity gets a reference to the given AllocatedLimitCapacity and assigns it to the AllocatedLimitCapacity field.
func (o *FabricsNode) SetAllocatedLimitCapacity(v AllocatedLimitCapacity) {
	o.AllocatedLimitCapacity = &v
}

func (o FabricsNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricsNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.KubeletVersion) {
		toSerialize["kubeletVersion"] = o.KubeletVersion
	}
	if !IsNil(o.DockerVersion) {
		toSerialize["dockerVersion"] = o.DockerVersion
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !IsNil(o.AllocatedRequestCapacity) {
		toSerialize["allocatedRequestCapacity"] = o.AllocatedRequestCapacity
	}
	if !IsNil(o.AllocatedLimitCapacity) {
		toSerialize["allocatedLimitCapacity"] = o.AllocatedLimitCapacity
	}
	return toSerialize, nil
}

type NullableFabricsNode struct {
	value *FabricsNode
	isSet bool
}

func (v NullableFabricsNode) Get() *FabricsNode {
	return v.value
}

func (v *NullableFabricsNode) Set(val *FabricsNode) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricsNode) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricsNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricsNode(val *FabricsNode) *NullableFabricsNode {
	return &NullableFabricsNode{value: val, isSet: true}
}

func (v NullableFabricsNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricsNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


