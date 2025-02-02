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

// checks if the AssignUsersToPrivilegeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignUsersToPrivilegeRequest{}

// AssignUsersToPrivilegeRequest struct for AssignUsersToPrivilegeRequest
type AssignUsersToPrivilegeRequest struct {
	Users []int32 `json:"users,omitempty"`
}

// NewAssignUsersToPrivilegeRequest instantiates a new AssignUsersToPrivilegeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignUsersToPrivilegeRequest() *AssignUsersToPrivilegeRequest {
	this := AssignUsersToPrivilegeRequest{}
	return &this
}

// NewAssignUsersToPrivilegeRequestWithDefaults instantiates a new AssignUsersToPrivilegeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignUsersToPrivilegeRequestWithDefaults() *AssignUsersToPrivilegeRequest {
	this := AssignUsersToPrivilegeRequest{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AssignUsersToPrivilegeRequest) GetUsers() []int32 {
	if o == nil || isNil(o.Users) {
		var ret []int32
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignUsersToPrivilegeRequest) GetUsersOk() ([]int32, bool) {
	if o == nil || isNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AssignUsersToPrivilegeRequest) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *AssignUsersToPrivilegeRequest) SetUsers(v []int32) {
	o.Users = v
}

func (o AssignUsersToPrivilegeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignUsersToPrivilegeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableAssignUsersToPrivilegeRequest struct {
	value *AssignUsersToPrivilegeRequest
	isSet bool
}

func (v NullableAssignUsersToPrivilegeRequest) Get() *AssignUsersToPrivilegeRequest {
	return v.value
}

func (v *NullableAssignUsersToPrivilegeRequest) Set(val *AssignUsersToPrivilegeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignUsersToPrivilegeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignUsersToPrivilegeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignUsersToPrivilegeRequest(val *AssignUsersToPrivilegeRequest) *NullableAssignUsersToPrivilegeRequest {
	return &NullableAssignUsersToPrivilegeRequest{value: val, isSet: true}
}

func (v NullableAssignUsersToPrivilegeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignUsersToPrivilegeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


