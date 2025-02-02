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

// checks if the UpdateClientAppRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateClientAppRequest{}

// UpdateClientAppRequest struct for UpdateClientAppRequest
type UpdateClientAppRequest struct {
	// An array of Scope IDs the scopes the app can request
	Scopes []int32 `json:"scopes,omitempty"`
}

// NewUpdateClientAppRequest instantiates a new UpdateClientAppRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClientAppRequest() *UpdateClientAppRequest {
	this := UpdateClientAppRequest{}
	return &this
}

// NewUpdateClientAppRequestWithDefaults instantiates a new UpdateClientAppRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClientAppRequestWithDefaults() *UpdateClientAppRequest {
	this := UpdateClientAppRequest{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *UpdateClientAppRequest) GetScopes() []int32 {
	if o == nil || isNil(o.Scopes) {
		var ret []int32
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClientAppRequest) GetScopesOk() ([]int32, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *UpdateClientAppRequest) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []int32 and assigns it to the Scopes field.
func (o *UpdateClientAppRequest) SetScopes(v []int32) {
	o.Scopes = v
}

func (o UpdateClientAppRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateClientAppRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableUpdateClientAppRequest struct {
	value *UpdateClientAppRequest
	isSet bool
}

func (v NullableUpdateClientAppRequest) Get() *UpdateClientAppRequest {
	return v.value
}

func (v *NullableUpdateClientAppRequest) Set(val *UpdateClientAppRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateClientAppRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateClientAppRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateClientAppRequest(val *UpdateClientAppRequest) *NullableUpdateClientAppRequest {
	return &NullableUpdateClientAppRequest{value: val, isSet: true}
}

func (v NullableUpdateClientAppRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateClientAppRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


