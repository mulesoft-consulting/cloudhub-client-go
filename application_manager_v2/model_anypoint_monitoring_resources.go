/*
Deployment

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Runtime Fabric and CloudHub 2.0 targets only. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v2

import (
	"encoding/json"
)

// checks if the AnypointMonitoringResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnypointMonitoringResources{}

// AnypointMonitoringResources struct for AnypointMonitoringResources
type AnypointMonitoringResources struct {
	Cpu *AnypointMonitoringResourcesCpu `json:"cpu,omitempty"`
	Memory *AnypointMonitoringResourcesCpu `json:"memory,omitempty"`
}

// NewAnypointMonitoringResources instantiates a new AnypointMonitoringResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnypointMonitoringResources() *AnypointMonitoringResources {
	this := AnypointMonitoringResources{}
	return &this
}

// NewAnypointMonitoringResourcesWithDefaults instantiates a new AnypointMonitoringResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnypointMonitoringResourcesWithDefaults() *AnypointMonitoringResources {
	this := AnypointMonitoringResources{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *AnypointMonitoringResources) GetCpu() AnypointMonitoringResourcesCpu {
	if o == nil || IsNil(o.Cpu) {
		var ret AnypointMonitoringResourcesCpu
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnypointMonitoringResources) GetCpuOk() (*AnypointMonitoringResourcesCpu, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *AnypointMonitoringResources) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given AnypointMonitoringResourcesCpu and assigns it to the Cpu field.
func (o *AnypointMonitoringResources) SetCpu(v AnypointMonitoringResourcesCpu) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *AnypointMonitoringResources) GetMemory() AnypointMonitoringResourcesCpu {
	if o == nil || IsNil(o.Memory) {
		var ret AnypointMonitoringResourcesCpu
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnypointMonitoringResources) GetMemoryOk() (*AnypointMonitoringResourcesCpu, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *AnypointMonitoringResources) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given AnypointMonitoringResourcesCpu and assigns it to the Memory field.
func (o *AnypointMonitoringResources) SetMemory(v AnypointMonitoringResourcesCpu) {
	o.Memory = &v
}

func (o AnypointMonitoringResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnypointMonitoringResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableAnypointMonitoringResources struct {
	value *AnypointMonitoringResources
	isSet bool
}

func (v NullableAnypointMonitoringResources) Get() *AnypointMonitoringResources {
	return v.value
}

func (v *NullableAnypointMonitoringResources) Set(val *AnypointMonitoringResources) {
	v.value = val
	v.isSet = true
}

func (v NullableAnypointMonitoringResources) IsSet() bool {
	return v.isSet
}

func (v *NullableAnypointMonitoringResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnypointMonitoringResources(val *AnypointMonitoringResources) *NullableAnypointMonitoringResources {
	return &NullableAnypointMonitoringResources{value: val, isSet: true}
}

func (v NullableAnypointMonitoringResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnypointMonitoringResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


