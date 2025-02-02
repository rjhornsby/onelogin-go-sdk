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

// checks if the TokenClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenClaim{}

// TokenClaim struct for TokenClaim
type TokenClaim struct {
	// The unique ID of the claim.
	Id *int32 `json:"id,omitempty"`
	// The UI label for the claims.
	Label *string `json:"label,omitempty"`
	// A user attribute to map values from.
	UserAttributeMappings *string `json:"user_attribute_mappings,omitempty"`
	// When `user_attribute_mappings` is set to `_macro_` this macro will be used to assign the claims value.
	UserAttributeMacros *string `json:"user_attribute_macros,omitempty"`
	// The type of transformation to perform on multi valued attributes.
	AttributeTransformations *string `json:"attribute_transformations,omitempty"`
	// not used
	SkipIfBlank *bool `json:"skip_if_blank,omitempty"`
	// Relates to Rules/Entitlements. Not supported yet.
	Values []string `json:"values,omitempty"`
	// Relates to Rules/Entitlements. Not supported yet.
	DefaultValues *string `json:"default_values,omitempty"`
	// Relates to Rules/Entitlements. Not supported yet.
	ProvisionedEntitlements *bool `json:"provisioned_entitlements,omitempty"`
}

// NewTokenClaim instantiates a new TokenClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenClaim() *TokenClaim {
	this := TokenClaim{}
	return &this
}

// NewTokenClaimWithDefaults instantiates a new TokenClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenClaimWithDefaults() *TokenClaim {
	this := TokenClaim{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenClaim) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenClaim) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TokenClaim) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *TokenClaim) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *TokenClaim) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *TokenClaim) SetLabel(v string) {
	o.Label = &v
}

// GetUserAttributeMappings returns the UserAttributeMappings field value if set, zero value otherwise.
func (o *TokenClaim) GetUserAttributeMappings() string {
	if o == nil || isNil(o.UserAttributeMappings) {
		var ret string
		return ret
	}
	return *o.UserAttributeMappings
}

// GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetUserAttributeMappingsOk() (*string, bool) {
	if o == nil || isNil(o.UserAttributeMappings) {
		return nil, false
	}
	return o.UserAttributeMappings, true
}

// HasUserAttributeMappings returns a boolean if a field has been set.
func (o *TokenClaim) HasUserAttributeMappings() bool {
	if o != nil && !isNil(o.UserAttributeMappings) {
		return true
	}

	return false
}

// SetUserAttributeMappings gets a reference to the given string and assigns it to the UserAttributeMappings field.
func (o *TokenClaim) SetUserAttributeMappings(v string) {
	o.UserAttributeMappings = &v
}

// GetUserAttributeMacros returns the UserAttributeMacros field value if set, zero value otherwise.
func (o *TokenClaim) GetUserAttributeMacros() string {
	if o == nil || isNil(o.UserAttributeMacros) {
		var ret string
		return ret
	}
	return *o.UserAttributeMacros
}

// GetUserAttributeMacrosOk returns a tuple with the UserAttributeMacros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetUserAttributeMacrosOk() (*string, bool) {
	if o == nil || isNil(o.UserAttributeMacros) {
		return nil, false
	}
	return o.UserAttributeMacros, true
}

// HasUserAttributeMacros returns a boolean if a field has been set.
func (o *TokenClaim) HasUserAttributeMacros() bool {
	if o != nil && !isNil(o.UserAttributeMacros) {
		return true
	}

	return false
}

// SetUserAttributeMacros gets a reference to the given string and assigns it to the UserAttributeMacros field.
func (o *TokenClaim) SetUserAttributeMacros(v string) {
	o.UserAttributeMacros = &v
}

// GetAttributeTransformations returns the AttributeTransformations field value if set, zero value otherwise.
func (o *TokenClaim) GetAttributeTransformations() string {
	if o == nil || isNil(o.AttributeTransformations) {
		var ret string
		return ret
	}
	return *o.AttributeTransformations
}

// GetAttributeTransformationsOk returns a tuple with the AttributeTransformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetAttributeTransformationsOk() (*string, bool) {
	if o == nil || isNil(o.AttributeTransformations) {
		return nil, false
	}
	return o.AttributeTransformations, true
}

// HasAttributeTransformations returns a boolean if a field has been set.
func (o *TokenClaim) HasAttributeTransformations() bool {
	if o != nil && !isNil(o.AttributeTransformations) {
		return true
	}

	return false
}

// SetAttributeTransformations gets a reference to the given string and assigns it to the AttributeTransformations field.
func (o *TokenClaim) SetAttributeTransformations(v string) {
	o.AttributeTransformations = &v
}

// GetSkipIfBlank returns the SkipIfBlank field value if set, zero value otherwise.
func (o *TokenClaim) GetSkipIfBlank() bool {
	if o == nil || isNil(o.SkipIfBlank) {
		var ret bool
		return ret
	}
	return *o.SkipIfBlank
}

// GetSkipIfBlankOk returns a tuple with the SkipIfBlank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetSkipIfBlankOk() (*bool, bool) {
	if o == nil || isNil(o.SkipIfBlank) {
		return nil, false
	}
	return o.SkipIfBlank, true
}

// HasSkipIfBlank returns a boolean if a field has been set.
func (o *TokenClaim) HasSkipIfBlank() bool {
	if o != nil && !isNil(o.SkipIfBlank) {
		return true
	}

	return false
}

// SetSkipIfBlank gets a reference to the given bool and assigns it to the SkipIfBlank field.
func (o *TokenClaim) SetSkipIfBlank(v bool) {
	o.SkipIfBlank = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *TokenClaim) GetValues() []string {
	if o == nil || isNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetValuesOk() ([]string, bool) {
	if o == nil || isNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *TokenClaim) HasValues() bool {
	if o != nil && !isNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *TokenClaim) SetValues(v []string) {
	o.Values = v
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *TokenClaim) GetDefaultValues() string {
	if o == nil || isNil(o.DefaultValues) {
		var ret string
		return ret
	}
	return *o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetDefaultValuesOk() (*string, bool) {
	if o == nil || isNil(o.DefaultValues) {
		return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *TokenClaim) HasDefaultValues() bool {
	if o != nil && !isNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given string and assigns it to the DefaultValues field.
func (o *TokenClaim) SetDefaultValues(v string) {
	o.DefaultValues = &v
}

// GetProvisionedEntitlements returns the ProvisionedEntitlements field value if set, zero value otherwise.
func (o *TokenClaim) GetProvisionedEntitlements() bool {
	if o == nil || isNil(o.ProvisionedEntitlements) {
		var ret bool
		return ret
	}
	return *o.ProvisionedEntitlements
}

// GetProvisionedEntitlementsOk returns a tuple with the ProvisionedEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaim) GetProvisionedEntitlementsOk() (*bool, bool) {
	if o == nil || isNil(o.ProvisionedEntitlements) {
		return nil, false
	}
	return o.ProvisionedEntitlements, true
}

// HasProvisionedEntitlements returns a boolean if a field has been set.
func (o *TokenClaim) HasProvisionedEntitlements() bool {
	if o != nil && !isNil(o.ProvisionedEntitlements) {
		return true
	}

	return false
}

// SetProvisionedEntitlements gets a reference to the given bool and assigns it to the ProvisionedEntitlements field.
func (o *TokenClaim) SetProvisionedEntitlements(v bool) {
	o.ProvisionedEntitlements = &v
}

func (o TokenClaim) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenClaim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.UserAttributeMappings) {
		toSerialize["user_attribute_mappings"] = o.UserAttributeMappings
	}
	if !isNil(o.UserAttributeMacros) {
		toSerialize["user_attribute_macros"] = o.UserAttributeMacros
	}
	if !isNil(o.AttributeTransformations) {
		toSerialize["attribute_transformations"] = o.AttributeTransformations
	}
	if !isNil(o.SkipIfBlank) {
		toSerialize["skip_if_blank"] = o.SkipIfBlank
	}
	if !isNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !isNil(o.DefaultValues) {
		toSerialize["default_values"] = o.DefaultValues
	}
	if !isNil(o.ProvisionedEntitlements) {
		toSerialize["provisioned_entitlements"] = o.ProvisionedEntitlements
	}
	return toSerialize, nil
}

type NullableTokenClaim struct {
	value *TokenClaim
	isSet bool
}

func (v NullableTokenClaim) Get() *TokenClaim {
	return v.value
}

func (v *NullableTokenClaim) Set(val *TokenClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenClaim(val *TokenClaim) *NullableTokenClaim {
	return &NullableTokenClaim{value: val, isSet: true}
}

func (v NullableTokenClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


