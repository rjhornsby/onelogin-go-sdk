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

// checks if the SetUserStateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetUserStateRequest{}

// SetUserStateRequest struct for SetUserStateRequest
type SetUserStateRequest struct {
	// Set to the state value. Valid values include:   - 0 : Unapproved   - 1 : Approved   - 2 : Rejected   - 3 : Unlicensed
	State int32 `json:"state"`
}

// NewSetUserStateRequest instantiates a new SetUserStateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetUserStateRequest(state int32) *SetUserStateRequest {
	this := SetUserStateRequest{}
	this.State = state
	return &this
}

// NewSetUserStateRequestWithDefaults instantiates a new SetUserStateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetUserStateRequestWithDefaults() *SetUserStateRequest {
	this := SetUserStateRequest{}
	return &this
}

// GetState returns the State field value
func (o *SetUserStateRequest) GetState() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SetUserStateRequest) GetStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SetUserStateRequest) SetState(v int32) {
	o.State = v
}

func (o SetUserStateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetUserStateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableSetUserStateRequest struct {
	value *SetUserStateRequest
	isSet bool
}

func (v NullableSetUserStateRequest) Get() *SetUserStateRequest {
	return v.value
}

func (v *NullableSetUserStateRequest) Set(val *SetUserStateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetUserStateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetUserStateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetUserStateRequest(val *SetUserStateRequest) *NullableSetUserStateRequest {
	return &NullableSetUserStateRequest{value: val, isSet: true}
}

func (v NullableSetUserStateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetUserStateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


