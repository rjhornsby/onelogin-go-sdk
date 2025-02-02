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

// checks if the GetEvents200ResponsePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEvents200ResponsePagination{}

// GetEvents200ResponsePagination struct for GetEvents200ResponsePagination
type GetEvents200ResponsePagination struct {
	BeforeCursor *string `json:"before_cursor,omitempty"`
	AfterCursor *string `json:"after_cursor,omitempty"`
	PreviousLink *string `json:"previous_link,omitempty"`
	NextLink *string `json:"next_link,omitempty"`
}

// NewGetEvents200ResponsePagination instantiates a new GetEvents200ResponsePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEvents200ResponsePagination() *GetEvents200ResponsePagination {
	this := GetEvents200ResponsePagination{}
	return &this
}

// NewGetEvents200ResponsePaginationWithDefaults instantiates a new GetEvents200ResponsePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEvents200ResponsePaginationWithDefaults() *GetEvents200ResponsePagination {
	this := GetEvents200ResponsePagination{}
	return &this
}

// GetBeforeCursor returns the BeforeCursor field value if set, zero value otherwise.
func (o *GetEvents200ResponsePagination) GetBeforeCursor() string {
	if o == nil || isNil(o.BeforeCursor) {
		var ret string
		return ret
	}
	return *o.BeforeCursor
}

// GetBeforeCursorOk returns a tuple with the BeforeCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponsePagination) GetBeforeCursorOk() (*string, bool) {
	if o == nil || isNil(o.BeforeCursor) {
		return nil, false
	}
	return o.BeforeCursor, true
}

// HasBeforeCursor returns a boolean if a field has been set.
func (o *GetEvents200ResponsePagination) HasBeforeCursor() bool {
	if o != nil && !isNil(o.BeforeCursor) {
		return true
	}

	return false
}

// SetBeforeCursor gets a reference to the given string and assigns it to the BeforeCursor field.
func (o *GetEvents200ResponsePagination) SetBeforeCursor(v string) {
	o.BeforeCursor = &v
}

// GetAfterCursor returns the AfterCursor field value if set, zero value otherwise.
func (o *GetEvents200ResponsePagination) GetAfterCursor() string {
	if o == nil || isNil(o.AfterCursor) {
		var ret string
		return ret
	}
	return *o.AfterCursor
}

// GetAfterCursorOk returns a tuple with the AfterCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponsePagination) GetAfterCursorOk() (*string, bool) {
	if o == nil || isNil(o.AfterCursor) {
		return nil, false
	}
	return o.AfterCursor, true
}

// HasAfterCursor returns a boolean if a field has been set.
func (o *GetEvents200ResponsePagination) HasAfterCursor() bool {
	if o != nil && !isNil(o.AfterCursor) {
		return true
	}

	return false
}

// SetAfterCursor gets a reference to the given string and assigns it to the AfterCursor field.
func (o *GetEvents200ResponsePagination) SetAfterCursor(v string) {
	o.AfterCursor = &v
}

// GetPreviousLink returns the PreviousLink field value if set, zero value otherwise.
func (o *GetEvents200ResponsePagination) GetPreviousLink() string {
	if o == nil || isNil(o.PreviousLink) {
		var ret string
		return ret
	}
	return *o.PreviousLink
}

// GetPreviousLinkOk returns a tuple with the PreviousLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponsePagination) GetPreviousLinkOk() (*string, bool) {
	if o == nil || isNil(o.PreviousLink) {
		return nil, false
	}
	return o.PreviousLink, true
}

// HasPreviousLink returns a boolean if a field has been set.
func (o *GetEvents200ResponsePagination) HasPreviousLink() bool {
	if o != nil && !isNil(o.PreviousLink) {
		return true
	}

	return false
}

// SetPreviousLink gets a reference to the given string and assigns it to the PreviousLink field.
func (o *GetEvents200ResponsePagination) SetPreviousLink(v string) {
	o.PreviousLink = &v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *GetEvents200ResponsePagination) GetNextLink() string {
	if o == nil || isNil(o.NextLink) {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponsePagination) GetNextLinkOk() (*string, bool) {
	if o == nil || isNil(o.NextLink) {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *GetEvents200ResponsePagination) HasNextLink() bool {
	if o != nil && !isNil(o.NextLink) {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *GetEvents200ResponsePagination) SetNextLink(v string) {
	o.NextLink = &v
}

func (o GetEvents200ResponsePagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEvents200ResponsePagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BeforeCursor) {
		toSerialize["before_cursor"] = o.BeforeCursor
	}
	if !isNil(o.AfterCursor) {
		toSerialize["after_cursor"] = o.AfterCursor
	}
	if !isNil(o.PreviousLink) {
		toSerialize["previous_link"] = o.PreviousLink
	}
	if !isNil(o.NextLink) {
		toSerialize["next_link"] = o.NextLink
	}
	return toSerialize, nil
}

type NullableGetEvents200ResponsePagination struct {
	value *GetEvents200ResponsePagination
	isSet bool
}

func (v NullableGetEvents200ResponsePagination) Get() *GetEvents200ResponsePagination {
	return v.value
}

func (v *NullableGetEvents200ResponsePagination) Set(val *GetEvents200ResponsePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEvents200ResponsePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEvents200ResponsePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEvents200ResponsePagination(val *GetEvents200ResponsePagination) *NullableGetEvents200ResponsePagination {
	return &NullableGetEvents200ResponsePagination{value: val, isSet: true}
}

func (v NullableGetEvents200ResponsePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEvents200ResponsePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


