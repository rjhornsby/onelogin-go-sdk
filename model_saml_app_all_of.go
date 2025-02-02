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

// checks if the SamlAppAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlAppAllOf{}

// SamlAppAllOf struct for SamlAppAllOf
type SamlAppAllOf struct {
	Configuration ConfigurationSaml `json:"configuration"`
	Sso *SsoSaml `json:"sso,omitempty"`
	Parameters SamlAppAllOfParameters `json:"parameters"`
}

// NewSamlAppAllOf instantiates a new SamlAppAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAppAllOf(configuration ConfigurationSaml, parameters SamlAppAllOfParameters) *SamlAppAllOf {
	this := SamlAppAllOf{}
	this.Configuration = configuration
	this.Parameters = parameters
	return &this
}

// NewSamlAppAllOfWithDefaults instantiates a new SamlAppAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAppAllOfWithDefaults() *SamlAppAllOf {
	this := SamlAppAllOf{}
	return &this
}

// GetConfiguration returns the Configuration field value
func (o *SamlAppAllOf) GetConfiguration() ConfigurationSaml {
	if o == nil {
		var ret ConfigurationSaml
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *SamlAppAllOf) GetConfigurationOk() (*ConfigurationSaml, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *SamlAppAllOf) SetConfiguration(v ConfigurationSaml) {
	o.Configuration = v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *SamlAppAllOf) GetSso() SsoSaml {
	if o == nil || isNil(o.Sso) {
		var ret SsoSaml
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAppAllOf) GetSsoOk() (*SsoSaml, bool) {
	if o == nil || isNil(o.Sso) {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *SamlAppAllOf) HasSso() bool {
	if o != nil && !isNil(o.Sso) {
		return true
	}

	return false
}

// SetSso gets a reference to the given SsoSaml and assigns it to the Sso field.
func (o *SamlAppAllOf) SetSso(v SsoSaml) {
	o.Sso = &v
}

// GetParameters returns the Parameters field value
func (o *SamlAppAllOf) GetParameters() SamlAppAllOfParameters {
	if o == nil {
		var ret SamlAppAllOfParameters
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *SamlAppAllOf) GetParametersOk() (*SamlAppAllOfParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *SamlAppAllOf) SetParameters(v SamlAppAllOfParameters) {
	o.Parameters = v
}

func (o SamlAppAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlAppAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configuration"] = o.Configuration
	if !isNil(o.Sso) {
		toSerialize["sso"] = o.Sso
	}
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

type NullableSamlAppAllOf struct {
	value *SamlAppAllOf
	isSet bool
}

func (v NullableSamlAppAllOf) Get() *SamlAppAllOf {
	return v.value
}

func (v *NullableSamlAppAllOf) Set(val *SamlAppAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAppAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAppAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAppAllOf(val *SamlAppAllOf) *NullableSamlAppAllOf {
	return &NullableSamlAppAllOf{value: val, isSet: true}
}

func (v NullableSamlAppAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAppAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


