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

// checks if the GetUserRoles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserRoles200Response{}

// GetUserRoles200Response struct for GetUserRoles200Response
type GetUserRoles200Response struct {
	Status *Error `json:"status,omitempty"`
	// List of Role IDs that are assigned to the User
	Data []int32 `json:"data,omitempty"`
}

// NewGetUserRoles200Response instantiates a new GetUserRoles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserRoles200Response() *GetUserRoles200Response {
	this := GetUserRoles200Response{}
	return &this
}

// NewGetUserRoles200ResponseWithDefaults instantiates a new GetUserRoles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserRoles200ResponseWithDefaults() *GetUserRoles200Response {
	this := GetUserRoles200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetUserRoles200Response) GetStatus() Error {
	if o == nil || isNil(o.Status) {
		var ret Error
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRoles200Response) GetStatusOk() (*Error, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetUserRoles200Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Error and assigns it to the Status field.
func (o *GetUserRoles200Response) SetStatus(v Error) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetUserRoles200Response) GetData() []int32 {
	if o == nil || isNil(o.Data) {
		var ret []int32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRoles200Response) GetDataOk() ([]int32, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetUserRoles200Response) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []int32 and assigns it to the Data field.
func (o *GetUserRoles200Response) SetData(v []int32) {
	o.Data = v
}

func (o GetUserRoles200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserRoles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetUserRoles200Response struct {
	value *GetUserRoles200Response
	isSet bool
}

func (v NullableGetUserRoles200Response) Get() *GetUserRoles200Response {
	return v.value
}

func (v *NullableGetUserRoles200Response) Set(val *GetUserRoles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserRoles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserRoles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserRoles200Response(val *GetUserRoles200Response) *NullableGetUserRoles200Response {
	return &NullableGetUserRoles200Response{value: val, isSet: true}
}

func (v NullableGetUserRoles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserRoles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


