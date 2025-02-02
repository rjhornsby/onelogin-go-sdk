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

// checks if the GetMFAFactors200ResponseDataAuthFactorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMFAFactors200ResponseDataAuthFactorsInner{}

// GetMFAFactors200ResponseDataAuthFactorsInner struct for GetMFAFactors200ResponseDataAuthFactorsInner
type GetMFAFactors200ResponseDataAuthFactorsInner struct {
	// Official authentication factor name, as it appears to administrators in OneLogin.
	Name *string `json:"name,omitempty"`
	// Identifier for the factor which will be used for user enrollment
	FactorId *int32 `json:"factor_id,omitempty"`
}

// NewGetMFAFactors200ResponseDataAuthFactorsInner instantiates a new GetMFAFactors200ResponseDataAuthFactorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMFAFactors200ResponseDataAuthFactorsInner() *GetMFAFactors200ResponseDataAuthFactorsInner {
	this := GetMFAFactors200ResponseDataAuthFactorsInner{}
	return &this
}

// NewGetMFAFactors200ResponseDataAuthFactorsInnerWithDefaults instantiates a new GetMFAFactors200ResponseDataAuthFactorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMFAFactors200ResponseDataAuthFactorsInnerWithDefaults() *GetMFAFactors200ResponseDataAuthFactorsInner {
	this := GetMFAFactors200ResponseDataAuthFactorsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetMFAFactors200ResponseDataAuthFactorsInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMFAFactors200ResponseDataAuthFactorsInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetMFAFactors200ResponseDataAuthFactorsInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetMFAFactors200ResponseDataAuthFactorsInner) SetName(v string) {
	o.Name = &v
}

// GetFactorId returns the FactorId field value if set, zero value otherwise.
func (o *GetMFAFactors200ResponseDataAuthFactorsInner) GetFactorId() int32 {
	if o == nil || isNil(o.FactorId) {
		var ret int32
		return ret
	}
	return *o.FactorId
}

// GetFactorIdOk returns a tuple with the FactorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMFAFactors200ResponseDataAuthFactorsInner) GetFactorIdOk() (*int32, bool) {
	if o == nil || isNil(o.FactorId) {
		return nil, false
	}
	return o.FactorId, true
}

// HasFactorId returns a boolean if a field has been set.
func (o *GetMFAFactors200ResponseDataAuthFactorsInner) HasFactorId() bool {
	if o != nil && !isNil(o.FactorId) {
		return true
	}

	return false
}

// SetFactorId gets a reference to the given int32 and assigns it to the FactorId field.
func (o *GetMFAFactors200ResponseDataAuthFactorsInner) SetFactorId(v int32) {
	o.FactorId = &v
}

func (o GetMFAFactors200ResponseDataAuthFactorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMFAFactors200ResponseDataAuthFactorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.FactorId) {
		toSerialize["factor_id"] = o.FactorId
	}
	return toSerialize, nil
}

type NullableGetMFAFactors200ResponseDataAuthFactorsInner struct {
	value *GetMFAFactors200ResponseDataAuthFactorsInner
	isSet bool
}

func (v NullableGetMFAFactors200ResponseDataAuthFactorsInner) Get() *GetMFAFactors200ResponseDataAuthFactorsInner {
	return v.value
}

func (v *NullableGetMFAFactors200ResponseDataAuthFactorsInner) Set(val *GetMFAFactors200ResponseDataAuthFactorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMFAFactors200ResponseDataAuthFactorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMFAFactors200ResponseDataAuthFactorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMFAFactors200ResponseDataAuthFactorsInner(val *GetMFAFactors200ResponseDataAuthFactorsInner) *NullableGetMFAFactors200ResponseDataAuthFactorsInner {
	return &NullableGetMFAFactors200ResponseDataAuthFactorsInner{value: val, isSet: true}
}

func (v NullableGetMFAFactors200ResponseDataAuthFactorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMFAFactors200ResponseDataAuthFactorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


