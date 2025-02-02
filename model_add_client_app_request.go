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

// checks if the AddClientAppRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddClientAppRequest{}

// AddClientAppRequest struct for AddClientAppRequest
type AddClientAppRequest struct {
	// The ID of the OpenId Connect app to allow access through.
	AppId *int32 `json:"app_id,omitempty"`
	// An array of Scope IDs that represent scopes the app can request
	Scopes []int32 `json:"scopes,omitempty"`
}

// NewAddClientAppRequest instantiates a new AddClientAppRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddClientAppRequest() *AddClientAppRequest {
	this := AddClientAppRequest{}
	return &this
}

// NewAddClientAppRequestWithDefaults instantiates a new AddClientAppRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddClientAppRequestWithDefaults() *AddClientAppRequest {
	this := AddClientAppRequest{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AddClientAppRequest) GetAppId() int32 {
	if o == nil || isNil(o.AppId) {
		var ret int32
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddClientAppRequest) GetAppIdOk() (*int32, bool) {
	if o == nil || isNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AddClientAppRequest) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given int32 and assigns it to the AppId field.
func (o *AddClientAppRequest) SetAppId(v int32) {
	o.AppId = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AddClientAppRequest) GetScopes() []int32 {
	if o == nil || isNil(o.Scopes) {
		var ret []int32
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddClientAppRequest) GetScopesOk() ([]int32, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AddClientAppRequest) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []int32 and assigns it to the Scopes field.
func (o *AddClientAppRequest) SetScopes(v []int32) {
	o.Scopes = v
}

func (o AddClientAppRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddClientAppRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableAddClientAppRequest struct {
	value *AddClientAppRequest
	isSet bool
}

func (v NullableAddClientAppRequest) Get() *AddClientAppRequest {
	return v.value
}

func (v *NullableAddClientAppRequest) Set(val *AddClientAppRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddClientAppRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddClientAppRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddClientAppRequest(val *AddClientAppRequest) *NullableAddClientAppRequest {
	return &NullableAddClientAppRequest{value: val, isSet: true}
}

func (v NullableAddClientAppRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddClientAppRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


