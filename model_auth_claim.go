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

// checks if the AuthClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthClaim{}

// AuthClaim struct for AuthClaim
type AuthClaim struct {
	// The attribute name for the claim when its returned in an Access Token
	Name string `json:"name"`
	// A user attribute to map values from For custom attributes prefix the name of the attribute with `custom_attribute_`. e.g. To get the value for custom attribute `employee_id` use `custom_attribute_employee_id`.
	UserAttributeMappings *string `json:"user_attribute_mappings,omitempty"`
	// When `user_attribute_mappings` is set to `_macro_` this macro will be used to assign the parameter value.
	UserAttributeMacros *string `json:"user_attribute_macros,omitempty"`
}

// NewAuthClaim instantiates a new AuthClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthClaim(name string) *AuthClaim {
	this := AuthClaim{}
	this.Name = name
	return &this
}

// NewAuthClaimWithDefaults instantiates a new AuthClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthClaimWithDefaults() *AuthClaim {
	this := AuthClaim{}
	return &this
}

// GetName returns the Name field value
func (o *AuthClaim) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthClaim) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthClaim) SetName(v string) {
	o.Name = v
}

// GetUserAttributeMappings returns the UserAttributeMappings field value if set, zero value otherwise.
func (o *AuthClaim) GetUserAttributeMappings() string {
	if o == nil || isNil(o.UserAttributeMappings) {
		var ret string
		return ret
	}
	return *o.UserAttributeMappings
}

// GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthClaim) GetUserAttributeMappingsOk() (*string, bool) {
	if o == nil || isNil(o.UserAttributeMappings) {
		return nil, false
	}
	return o.UserAttributeMappings, true
}

// HasUserAttributeMappings returns a boolean if a field has been set.
func (o *AuthClaim) HasUserAttributeMappings() bool {
	if o != nil && !isNil(o.UserAttributeMappings) {
		return true
	}

	return false
}

// SetUserAttributeMappings gets a reference to the given string and assigns it to the UserAttributeMappings field.
func (o *AuthClaim) SetUserAttributeMappings(v string) {
	o.UserAttributeMappings = &v
}

// GetUserAttributeMacros returns the UserAttributeMacros field value if set, zero value otherwise.
func (o *AuthClaim) GetUserAttributeMacros() string {
	if o == nil || isNil(o.UserAttributeMacros) {
		var ret string
		return ret
	}
	return *o.UserAttributeMacros
}

// GetUserAttributeMacrosOk returns a tuple with the UserAttributeMacros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthClaim) GetUserAttributeMacrosOk() (*string, bool) {
	if o == nil || isNil(o.UserAttributeMacros) {
		return nil, false
	}
	return o.UserAttributeMacros, true
}

// HasUserAttributeMacros returns a boolean if a field has been set.
func (o *AuthClaim) HasUserAttributeMacros() bool {
	if o != nil && !isNil(o.UserAttributeMacros) {
		return true
	}

	return false
}

// SetUserAttributeMacros gets a reference to the given string and assigns it to the UserAttributeMacros field.
func (o *AuthClaim) SetUserAttributeMacros(v string) {
	o.UserAttributeMacros = &v
}

func (o AuthClaim) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthClaim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.UserAttributeMappings) {
		toSerialize["user_attribute_mappings"] = o.UserAttributeMappings
	}
	if !isNil(o.UserAttributeMacros) {
		toSerialize["user_attribute_macros"] = o.UserAttributeMacros
	}
	return toSerialize, nil
}

type NullableAuthClaim struct {
	value *AuthClaim
	isSet bool
}

func (v NullableAuthClaim) Get() *AuthClaim {
	return v.value
}

func (v *NullableAuthClaim) Set(val *AuthClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthClaim(val *AuthClaim) *NullableAuthClaim {
	return &NullableAuthClaim{value: val, isSet: true}
}

func (v NullableAuthClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


