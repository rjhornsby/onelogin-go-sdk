/*
OneLogin API

OpenAPI Specification for OneLogin

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onelogin

import (
	"encoding/json"
)

// checks if the RemoveUserRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveUserRoleRequest{}

// RemoveUserRoleRequest struct for RemoveUserRoleRequest
type RemoveUserRoleRequest struct {
	RoleIdArray []RemoveUserRoleRequestRoleIdArrayInner `json:"role_id_array"`
}

// NewRemoveUserRoleRequest instantiates a new RemoveUserRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveUserRoleRequest(roleIdArray []RemoveUserRoleRequestRoleIdArrayInner) *RemoveUserRoleRequest {
	this := RemoveUserRoleRequest{}
	this.RoleIdArray = roleIdArray
	return &this
}

// NewRemoveUserRoleRequestWithDefaults instantiates a new RemoveUserRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveUserRoleRequestWithDefaults() *RemoveUserRoleRequest {
	this := RemoveUserRoleRequest{}
	return &this
}

// GetRoleIdArray returns the RoleIdArray field value
func (o *RemoveUserRoleRequest) GetRoleIdArray() []RemoveUserRoleRequestRoleIdArrayInner {
	if o == nil {
		var ret []RemoveUserRoleRequestRoleIdArrayInner
		return ret
	}

	return o.RoleIdArray
}

// GetRoleIdArrayOk returns a tuple with the RoleIdArray field value
// and a boolean to check if the value has been set.
func (o *RemoveUserRoleRequest) GetRoleIdArrayOk() ([]RemoveUserRoleRequestRoleIdArrayInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIdArray, true
}

// SetRoleIdArray sets field value
func (o *RemoveUserRoleRequest) SetRoleIdArray(v []RemoveUserRoleRequestRoleIdArrayInner) {
	o.RoleIdArray = v
}

func (o RemoveUserRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveUserRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role_id_array"] = o.RoleIdArray
	return toSerialize, nil
}

type NullableRemoveUserRoleRequest struct {
	value *RemoveUserRoleRequest
	isSet bool
}

func (v NullableRemoveUserRoleRequest) Get() *RemoveUserRoleRequest {
	return v.value
}

func (v *NullableRemoveUserRoleRequest) Set(val *RemoveUserRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveUserRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveUserRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveUserRoleRequest(val *RemoveUserRoleRequest) *NullableRemoveUserRoleRequest {
	return &NullableRemoveUserRoleRequest{value: val, isSet: true}
}

func (v NullableRemoveUserRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveUserRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


