/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// DeploymentsPodInfo struct for DeploymentsPodInfo
type DeploymentsPodInfo struct {
	// The IP of this pod.
	Ip *string `json:"ip,omitempty"`
	// Name is the Hostname of this pod.
	Name *string `json:"name,omitempty"`
	// Node hosting this pod.
	WorkerNode *string `json:"workerNode,omitempty"`
}

// NewDeploymentsPodInfo instantiates a new DeploymentsPodInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentsPodInfo() *DeploymentsPodInfo {
	this := DeploymentsPodInfo{}
	return &this
}

// NewDeploymentsPodInfoWithDefaults instantiates a new DeploymentsPodInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentsPodInfoWithDefaults() *DeploymentsPodInfo {
	this := DeploymentsPodInfo{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *DeploymentsPodInfo) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsPodInfo) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *DeploymentsPodInfo) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *DeploymentsPodInfo) SetIp(v string) {
	o.Ip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentsPodInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsPodInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentsPodInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentsPodInfo) SetName(v string) {
	o.Name = &v
}

// GetWorkerNode returns the WorkerNode field value if set, zero value otherwise.
func (o *DeploymentsPodInfo) GetWorkerNode() string {
	if o == nil || o.WorkerNode == nil {
		var ret string
		return ret
	}
	return *o.WorkerNode
}

// GetWorkerNodeOk returns a tuple with the WorkerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentsPodInfo) GetWorkerNodeOk() (*string, bool) {
	if o == nil || o.WorkerNode == nil {
		return nil, false
	}
	return o.WorkerNode, true
}

// HasWorkerNode returns a boolean if a field has been set.
func (o *DeploymentsPodInfo) HasWorkerNode() bool {
	if o != nil && o.WorkerNode != nil {
		return true
	}

	return false
}

// SetWorkerNode gets a reference to the given string and assigns it to the WorkerNode field.
func (o *DeploymentsPodInfo) SetWorkerNode(v string) {
	o.WorkerNode = &v
}

func (o DeploymentsPodInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.WorkerNode != nil {
		toSerialize["workerNode"] = o.WorkerNode
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentsPodInfo struct {
	value *DeploymentsPodInfo
	isSet bool
}

func (v NullableDeploymentsPodInfo) Get() *DeploymentsPodInfo {
	return v.value
}

func (v *NullableDeploymentsPodInfo) Set(val *DeploymentsPodInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentsPodInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentsPodInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentsPodInfo(val *DeploymentsPodInfo) *NullableDeploymentsPodInfo {
	return &NullableDeploymentsPodInfo{value: val, isSet: true}
}

func (v NullableDeploymentsPodInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentsPodInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

