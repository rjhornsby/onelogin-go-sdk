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

// checks if the GenericAppProvisioning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericAppProvisioning{}

// GenericAppProvisioning Indicates if provisioning is enabled for this app.
type GenericAppProvisioning struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGenericAppProvisioning instantiates a new GenericAppProvisioning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericAppProvisioning() *GenericAppProvisioning {
	this := GenericAppProvisioning{}
	return &this
}

// NewGenericAppProvisioningWithDefaults instantiates a new GenericAppProvisioning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericAppProvisioningWithDefaults() *GenericAppProvisioning {
	this := GenericAppProvisioning{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GenericAppProvisioning) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericAppProvisioning) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GenericAppProvisioning) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GenericAppProvisioning) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GenericAppProvisioning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericAppProvisioning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGenericAppProvisioning struct {
	value *GenericAppProvisioning
	isSet bool
}

func (v NullableGenericAppProvisioning) Get() *GenericAppProvisioning {
	return v.value
}

func (v *NullableGenericAppProvisioning) Set(val *GenericAppProvisioning) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericAppProvisioning) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericAppProvisioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericAppProvisioning(val *GenericAppProvisioning) *NullableGenericAppProvisioning {
	return &NullableGenericAppProvisioning{value: val, isSet: true}
}

func (v NullableGenericAppProvisioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericAppProvisioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


