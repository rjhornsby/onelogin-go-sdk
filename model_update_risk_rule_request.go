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

// checks if the UpdateRiskRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRiskRuleRequest{}

// UpdateRiskRuleRequest struct for UpdateRiskRuleRequest
type UpdateRiskRuleRequest struct {
	// The ID of the Rule to Update
	Id *string `json:"id,omitempty"`
}

// NewUpdateRiskRuleRequest instantiates a new UpdateRiskRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRiskRuleRequest() *UpdateRiskRuleRequest {
	this := UpdateRiskRuleRequest{}
	return &this
}

// NewUpdateRiskRuleRequestWithDefaults instantiates a new UpdateRiskRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRiskRuleRequestWithDefaults() *UpdateRiskRuleRequest {
	this := UpdateRiskRuleRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateRiskRuleRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRiskRuleRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateRiskRuleRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateRiskRuleRequest) SetId(v string) {
	o.Id = &v
}

func (o UpdateRiskRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRiskRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableUpdateRiskRuleRequest struct {
	value *UpdateRiskRuleRequest
	isSet bool
}

func (v NullableUpdateRiskRuleRequest) Get() *UpdateRiskRuleRequest {
	return v.value
}

func (v *NullableUpdateRiskRuleRequest) Set(val *UpdateRiskRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRiskRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRiskRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRiskRuleRequest(val *UpdateRiskRuleRequest) *NullableUpdateRiskRuleRequest {
	return &NullableUpdateRiskRuleRequest{value: val, isSet: true}
}

func (v NullableUpdateRiskRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRiskRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


