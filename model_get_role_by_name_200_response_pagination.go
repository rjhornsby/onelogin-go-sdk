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

// checks if the GetRoleByName200ResponsePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoleByName200ResponsePagination{}

// GetRoleByName200ResponsePagination struct for GetRoleByName200ResponsePagination
type GetRoleByName200ResponsePagination struct {
	AfterCursor *string `json:"after_cursor,omitempty"`
	BeforeCursor *string `json:"before_cursor,omitempty"`
	NextLink *string `json:"next_link,omitempty"`
	PreviousLink *string `json:"previous_link,omitempty"`
}

// NewGetRoleByName200ResponsePagination instantiates a new GetRoleByName200ResponsePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoleByName200ResponsePagination() *GetRoleByName200ResponsePagination {
	this := GetRoleByName200ResponsePagination{}
	return &this
}

// NewGetRoleByName200ResponsePaginationWithDefaults instantiates a new GetRoleByName200ResponsePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoleByName200ResponsePaginationWithDefaults() *GetRoleByName200ResponsePagination {
	this := GetRoleByName200ResponsePagination{}
	return &this
}

// GetAfterCursor returns the AfterCursor field value if set, zero value otherwise.
func (o *GetRoleByName200ResponsePagination) GetAfterCursor() string {
	if o == nil || isNil(o.AfterCursor) {
		var ret string
		return ret
	}
	return *o.AfterCursor
}

// GetAfterCursorOk returns a tuple with the AfterCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleByName200ResponsePagination) GetAfterCursorOk() (*string, bool) {
	if o == nil || isNil(o.AfterCursor) {
		return nil, false
	}
	return o.AfterCursor, true
}

// HasAfterCursor returns a boolean if a field has been set.
func (o *GetRoleByName200ResponsePagination) HasAfterCursor() bool {
	if o != nil && !isNil(o.AfterCursor) {
		return true
	}

	return false
}

// SetAfterCursor gets a reference to the given string and assigns it to the AfterCursor field.
func (o *GetRoleByName200ResponsePagination) SetAfterCursor(v string) {
	o.AfterCursor = &v
}

// GetBeforeCursor returns the BeforeCursor field value if set, zero value otherwise.
func (o *GetRoleByName200ResponsePagination) GetBeforeCursor() string {
	if o == nil || isNil(o.BeforeCursor) {
		var ret string
		return ret
	}
	return *o.BeforeCursor
}

// GetBeforeCursorOk returns a tuple with the BeforeCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleByName200ResponsePagination) GetBeforeCursorOk() (*string, bool) {
	if o == nil || isNil(o.BeforeCursor) {
		return nil, false
	}
	return o.BeforeCursor, true
}

// HasBeforeCursor returns a boolean if a field has been set.
func (o *GetRoleByName200ResponsePagination) HasBeforeCursor() bool {
	if o != nil && !isNil(o.BeforeCursor) {
		return true
	}

	return false
}

// SetBeforeCursor gets a reference to the given string and assigns it to the BeforeCursor field.
func (o *GetRoleByName200ResponsePagination) SetBeforeCursor(v string) {
	o.BeforeCursor = &v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *GetRoleByName200ResponsePagination) GetNextLink() string {
	if o == nil || isNil(o.NextLink) {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleByName200ResponsePagination) GetNextLinkOk() (*string, bool) {
	if o == nil || isNil(o.NextLink) {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *GetRoleByName200ResponsePagination) HasNextLink() bool {
	if o != nil && !isNil(o.NextLink) {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *GetRoleByName200ResponsePagination) SetNextLink(v string) {
	o.NextLink = &v
}

// GetPreviousLink returns the PreviousLink field value if set, zero value otherwise.
func (o *GetRoleByName200ResponsePagination) GetPreviousLink() string {
	if o == nil || isNil(o.PreviousLink) {
		var ret string
		return ret
	}
	return *o.PreviousLink
}

// GetPreviousLinkOk returns a tuple with the PreviousLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleByName200ResponsePagination) GetPreviousLinkOk() (*string, bool) {
	if o == nil || isNil(o.PreviousLink) {
		return nil, false
	}
	return o.PreviousLink, true
}

// HasPreviousLink returns a boolean if a field has been set.
func (o *GetRoleByName200ResponsePagination) HasPreviousLink() bool {
	if o != nil && !isNil(o.PreviousLink) {
		return true
	}

	return false
}

// SetPreviousLink gets a reference to the given string and assigns it to the PreviousLink field.
func (o *GetRoleByName200ResponsePagination) SetPreviousLink(v string) {
	o.PreviousLink = &v
}

func (o GetRoleByName200ResponsePagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoleByName200ResponsePagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AfterCursor) {
		toSerialize["after_cursor"] = o.AfterCursor
	}
	if !isNil(o.BeforeCursor) {
		toSerialize["before_cursor"] = o.BeforeCursor
	}
	if !isNil(o.NextLink) {
		toSerialize["next_link"] = o.NextLink
	}
	if !isNil(o.PreviousLink) {
		toSerialize["previous_link"] = o.PreviousLink
	}
	return toSerialize, nil
}

type NullableGetRoleByName200ResponsePagination struct {
	value *GetRoleByName200ResponsePagination
	isSet bool
}

func (v NullableGetRoleByName200ResponsePagination) Get() *GetRoleByName200ResponsePagination {
	return v.value
}

func (v *NullableGetRoleByName200ResponsePagination) Set(val *GetRoleByName200ResponsePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleByName200ResponsePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleByName200ResponsePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleByName200ResponsePagination(val *GetRoleByName200ResponsePagination) *NullableGetRoleByName200ResponsePagination {
	return &NullableGetRoleByName200ResponsePagination{value: val, isSet: true}
}

func (v NullableGetRoleByName200ResponsePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleByName200ResponsePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


