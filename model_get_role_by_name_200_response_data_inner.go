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

// checks if the GetRoleByName200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoleByName200ResponseDataInner{}

// GetRoleByName200ResponseDataInner struct for GetRoleByName200ResponseDataInner
type GetRoleByName200ResponseDataInner struct {
	// Role ID
	Id *int32 `json:"id,omitempty"`
	// Role Name
	Name *string `json:"name,omitempty"`
}

// NewGetRoleByName200ResponseDataInner instantiates a new GetRoleByName200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoleByName200ResponseDataInner() *GetRoleByName200ResponseDataInner {
	this := GetRoleByName200ResponseDataInner{}
	return &this
}

// NewGetRoleByName200ResponseDataInnerWithDefaults instantiates a new GetRoleByName200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoleByName200ResponseDataInnerWithDefaults() *GetRoleByName200ResponseDataInner {
	this := GetRoleByName200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetRoleByName200ResponseDataInner) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleByName200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetRoleByName200ResponseDataInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetRoleByName200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetRoleByName200ResponseDataInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleByName200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetRoleByName200ResponseDataInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetRoleByName200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

func (o GetRoleByName200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoleByName200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetRoleByName200ResponseDataInner struct {
	value *GetRoleByName200ResponseDataInner
	isSet bool
}

func (v NullableGetRoleByName200ResponseDataInner) Get() *GetRoleByName200ResponseDataInner {
	return v.value
}

func (v *NullableGetRoleByName200ResponseDataInner) Set(val *GetRoleByName200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleByName200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleByName200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleByName200ResponseDataInner(val *GetRoleByName200ResponseDataInner) *NullableGetRoleByName200ResponseDataInner {
	return &NullableGetRoleByName200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetRoleByName200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleByName200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


