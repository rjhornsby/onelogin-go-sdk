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

// checks if the SamlAppAllOfParametersSamlUsername type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlAppAllOfParametersSamlUsername{}

// SamlAppAllOfParametersSamlUsername struct for SamlAppAllOfParametersSamlUsername
type SamlAppAllOfParametersSamlUsername struct {
	UserAttributeMappings string `json:"user_attribute_mappings"`
}

// NewSamlAppAllOfParametersSamlUsername instantiates a new SamlAppAllOfParametersSamlUsername object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAppAllOfParametersSamlUsername(userAttributeMappings string) *SamlAppAllOfParametersSamlUsername {
	this := SamlAppAllOfParametersSamlUsername{}
	this.UserAttributeMappings = userAttributeMappings
	return &this
}

// NewSamlAppAllOfParametersSamlUsernameWithDefaults instantiates a new SamlAppAllOfParametersSamlUsername object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAppAllOfParametersSamlUsernameWithDefaults() *SamlAppAllOfParametersSamlUsername {
	this := SamlAppAllOfParametersSamlUsername{}
	return &this
}

// GetUserAttributeMappings returns the UserAttributeMappings field value
func (o *SamlAppAllOfParametersSamlUsername) GetUserAttributeMappings() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAttributeMappings
}

// GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field value
// and a boolean to check if the value has been set.
func (o *SamlAppAllOfParametersSamlUsername) GetUserAttributeMappingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAttributeMappings, true
}

// SetUserAttributeMappings sets field value
func (o *SamlAppAllOfParametersSamlUsername) SetUserAttributeMappings(v string) {
	o.UserAttributeMappings = v
}

func (o SamlAppAllOfParametersSamlUsername) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlAppAllOfParametersSamlUsername) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_attribute_mappings"] = o.UserAttributeMappings
	return toSerialize, nil
}

type NullableSamlAppAllOfParametersSamlUsername struct {
	value *SamlAppAllOfParametersSamlUsername
	isSet bool
}

func (v NullableSamlAppAllOfParametersSamlUsername) Get() *SamlAppAllOfParametersSamlUsername {
	return v.value
}

func (v *NullableSamlAppAllOfParametersSamlUsername) Set(val *SamlAppAllOfParametersSamlUsername) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAppAllOfParametersSamlUsername) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAppAllOfParametersSamlUsername) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAppAllOfParametersSamlUsername(val *SamlAppAllOfParametersSamlUsername) *NullableSamlAppAllOfParametersSamlUsername {
	return &NullableSamlAppAllOfParametersSamlUsername{value: val, isSet: true}
}

func (v NullableSamlAppAllOfParametersSamlUsername) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAppAllOfParametersSamlUsername) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


