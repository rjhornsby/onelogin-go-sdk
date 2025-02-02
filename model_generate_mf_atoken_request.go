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

// checks if the GenerateMFAtokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateMFAtokenRequest{}

// GenerateMFAtokenRequest struct for GenerateMFAtokenRequest
type GenerateMFAtokenRequest struct {
	// Set the duration of the token in seconds. Token expiration defaults to 259200 seconds = 72 hours. 72 hours is the max value.
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	// Defines if the token is reusable multiple times within the expiry window. Value defaults to false. If set to true, token can be used multiple times, until it expires.
	Reusable *bool `json:"reusable,omitempty"`
}

// NewGenerateMFAtokenRequest instantiates a new GenerateMFAtokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateMFAtokenRequest() *GenerateMFAtokenRequest {
	this := GenerateMFAtokenRequest{}
	var reusable bool = false
	this.Reusable = &reusable
	return &this
}

// NewGenerateMFAtokenRequestWithDefaults instantiates a new GenerateMFAtokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateMFAtokenRequestWithDefaults() *GenerateMFAtokenRequest {
	this := GenerateMFAtokenRequest{}
	var reusable bool = false
	this.Reusable = &reusable
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *GenerateMFAtokenRequest) GetExpiresIn() int32 {
	if o == nil || isNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMFAtokenRequest) GetExpiresInOk() (*int32, bool) {
	if o == nil || isNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *GenerateMFAtokenRequest) HasExpiresIn() bool {
	if o != nil && !isNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *GenerateMFAtokenRequest) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetReusable returns the Reusable field value if set, zero value otherwise.
func (o *GenerateMFAtokenRequest) GetReusable() bool {
	if o == nil || isNil(o.Reusable) {
		var ret bool
		return ret
	}
	return *o.Reusable
}

// GetReusableOk returns a tuple with the Reusable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMFAtokenRequest) GetReusableOk() (*bool, bool) {
	if o == nil || isNil(o.Reusable) {
		return nil, false
	}
	return o.Reusable, true
}

// HasReusable returns a boolean if a field has been set.
func (o *GenerateMFAtokenRequest) HasReusable() bool {
	if o != nil && !isNil(o.Reusable) {
		return true
	}

	return false
}

// SetReusable gets a reference to the given bool and assigns it to the Reusable field.
func (o *GenerateMFAtokenRequest) SetReusable(v bool) {
	o.Reusable = &v
}

func (o GenerateMFAtokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateMFAtokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !isNil(o.Reusable) {
		toSerialize["reusable"] = o.Reusable
	}
	return toSerialize, nil
}

type NullableGenerateMFAtokenRequest struct {
	value *GenerateMFAtokenRequest
	isSet bool
}

func (v NullableGenerateMFAtokenRequest) Get() *GenerateMFAtokenRequest {
	return v.value
}

func (v *NullableGenerateMFAtokenRequest) Set(val *GenerateMFAtokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateMFAtokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateMFAtokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateMFAtokenRequest(val *GenerateMFAtokenRequest) *NullableGenerateMFAtokenRequest {
	return &NullableGenerateMFAtokenRequest{value: val, isSet: true}
}

func (v NullableGenerateMFAtokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateMFAtokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


