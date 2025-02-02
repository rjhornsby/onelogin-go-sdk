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

// checks if the HookStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HookStatus{}

// HookStatus struct for HookStatus
type HookStatus struct {
	// responses status nam
	Name *string `json:"name,omitempty"`
	// your operation was successful
	Message *string `json:"message,omitempty"`
}

// NewHookStatus instantiates a new HookStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHookStatus() *HookStatus {
	this := HookStatus{}
	return &this
}

// NewHookStatusWithDefaults instantiates a new HookStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookStatusWithDefaults() *HookStatus {
	this := HookStatus{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HookStatus) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookStatus) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HookStatus) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HookStatus) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HookStatus) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookStatus) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HookStatus) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HookStatus) SetMessage(v string) {
	o.Message = &v
}

func (o HookStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HookStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableHookStatus struct {
	value *HookStatus
	isSet bool
}

func (v NullableHookStatus) Get() *HookStatus {
	return v.value
}

func (v *NullableHookStatus) Set(val *HookStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHookStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHookStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHookStatus(val *HookStatus) *NullableHookStatus {
	return &NullableHookStatus{value: val, isSet: true}
}

func (v NullableHookStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHookStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


