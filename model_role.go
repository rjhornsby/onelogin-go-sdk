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

// checks if the Role type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Role{}

// Role struct for Role
type Role struct {
	// Role ID
	Id *int32 `json:"id,omitempty"`
	// The name of the role.
	Name string `json:"name"`
	// A list of app IDs that will be assigned to the role.
	Apps []int32 `json:"apps,omitempty"`
	// A list of user IDs to assign to the role.
	Users []int32 `json:"users,omitempty"`
	// A list of user IDs to assign as role administrators.
	Admins []int32 `json:"admins,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole(name string) *Role {
	this := Role{}
	this.Name = name
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Role) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Role) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Role) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Role) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Role) SetName(v string) {
	o.Name = v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *Role) GetApps() []int32 {
	if o == nil || isNil(o.Apps) {
		var ret []int32
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetAppsOk() ([]int32, bool) {
	if o == nil || isNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *Role) HasApps() bool {
	if o != nil && !isNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []int32 and assigns it to the Apps field.
func (o *Role) SetApps(v []int32) {
	o.Apps = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Role) GetUsers() []int32 {
	if o == nil || isNil(o.Users) {
		var ret []int32
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetUsersOk() ([]int32, bool) {
	if o == nil || isNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Role) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *Role) SetUsers(v []int32) {
	o.Users = v
}

// GetAdmins returns the Admins field value if set, zero value otherwise.
func (o *Role) GetAdmins() []int32 {
	if o == nil || isNil(o.Admins) {
		var ret []int32
		return ret
	}
	return o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetAdminsOk() ([]int32, bool) {
	if o == nil || isNil(o.Admins) {
		return nil, false
	}
	return o.Admins, true
}

// HasAdmins returns a boolean if a field has been set.
func (o *Role) HasAdmins() bool {
	if o != nil && !isNil(o.Admins) {
		return true
	}

	return false
}

// SetAdmins gets a reference to the given []int32 and assigns it to the Admins field.
func (o *Role) SetAdmins(v []int32) {
	o.Admins = v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Role) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	if !isNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !isNil(o.Admins) {
		toSerialize["admins"] = o.Admins
	}
	return toSerialize, nil
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


