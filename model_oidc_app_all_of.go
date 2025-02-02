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

// checks if the OidcAppAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OidcAppAllOf{}

// OidcAppAllOf struct for OidcAppAllOf
type OidcAppAllOf struct {
	Configuration ConfigurationOidc `json:"configuration"`
	Sso *SsoOidc `json:"sso,omitempty"`
}

// NewOidcAppAllOf instantiates a new OidcAppAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcAppAllOf(configuration ConfigurationOidc) *OidcAppAllOf {
	this := OidcAppAllOf{}
	this.Configuration = configuration
	return &this
}

// NewOidcAppAllOfWithDefaults instantiates a new OidcAppAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcAppAllOfWithDefaults() *OidcAppAllOf {
	this := OidcAppAllOf{}
	return &this
}

// GetConfiguration returns the Configuration field value
func (o *OidcAppAllOf) GetConfiguration() ConfigurationOidc {
	if o == nil {
		var ret ConfigurationOidc
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *OidcAppAllOf) GetConfigurationOk() (*ConfigurationOidc, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *OidcAppAllOf) SetConfiguration(v ConfigurationOidc) {
	o.Configuration = v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *OidcAppAllOf) GetSso() SsoOidc {
	if o == nil || isNil(o.Sso) {
		var ret SsoOidc
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcAppAllOf) GetSsoOk() (*SsoOidc, bool) {
	if o == nil || isNil(o.Sso) {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *OidcAppAllOf) HasSso() bool {
	if o != nil && !isNil(o.Sso) {
		return true
	}

	return false
}

// SetSso gets a reference to the given SsoOidc and assigns it to the Sso field.
func (o *OidcAppAllOf) SetSso(v SsoOidc) {
	o.Sso = &v
}

func (o OidcAppAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcAppAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configuration"] = o.Configuration
	if !isNil(o.Sso) {
		toSerialize["sso"] = o.Sso
	}
	return toSerialize, nil
}

type NullableOidcAppAllOf struct {
	value *OidcAppAllOf
	isSet bool
}

func (v NullableOidcAppAllOf) Get() *OidcAppAllOf {
	return v.value
}

func (v *NullableOidcAppAllOf) Set(val *OidcAppAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcAppAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcAppAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcAppAllOf(val *OidcAppAllOf) *NullableOidcAppAllOf {
	return &NullableOidcAppAllOf{value: val, isSet: true}
}

func (v NullableOidcAppAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcAppAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


