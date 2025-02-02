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

// checks if the TrackRiskEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackRiskEventRequest{}

// TrackRiskEventRequest struct for TrackRiskEventRequest
type TrackRiskEventRequest struct {
	// Verbs are used to distinguish between different types of events.
	Verb string `json:"verb"`
	// The IP address of the User's request.
	Ip string `json:"ip"`
	// The user agent of the User's request.
	UserAgent string `json:"user_agent"`
	User RiskUser `json:"user"`
	Source *Source `json:"source,omitempty"`
	Session *Session `json:"session,omitempty"`
	Device *RiskDevice `json:"device,omitempty"`
	// Set to the value of the __tdli_fp cookie.
	Fp *string `json:"fp,omitempty"`
	// Date and time of the event in IS08601 format. Useful for preloading old events. Defaults to date time this API request is received.
	Published *string `json:"published,omitempty"`
}

// NewTrackRiskEventRequest instantiates a new TrackRiskEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackRiskEventRequest(verb string, ip string, userAgent string, user RiskUser) *TrackRiskEventRequest {
	this := TrackRiskEventRequest{}
	this.Verb = verb
	this.Ip = ip
	this.UserAgent = userAgent
	this.User = user
	return &this
}

// NewTrackRiskEventRequestWithDefaults instantiates a new TrackRiskEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackRiskEventRequestWithDefaults() *TrackRiskEventRequest {
	this := TrackRiskEventRequest{}
	return &this
}

// GetVerb returns the Verb field value
func (o *TrackRiskEventRequest) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *TrackRiskEventRequest) SetVerb(v string) {
	o.Verb = v
}

// GetIp returns the Ip field value
func (o *TrackRiskEventRequest) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *TrackRiskEventRequest) SetIp(v string) {
	o.Ip = v
}

// GetUserAgent returns the UserAgent field value
func (o *TrackRiskEventRequest) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *TrackRiskEventRequest) SetUserAgent(v string) {
	o.UserAgent = v
}

// GetUser returns the User field value
func (o *TrackRiskEventRequest) GetUser() RiskUser {
	if o == nil {
		var ret RiskUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetUserOk() (*RiskUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TrackRiskEventRequest) SetUser(v RiskUser) {
	o.User = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TrackRiskEventRequest) GetSource() Source {
	if o == nil || isNil(o.Source) {
		var ret Source
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetSourceOk() (*Source, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TrackRiskEventRequest) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Source and assigns it to the Source field.
func (o *TrackRiskEventRequest) SetSource(v Source) {
	o.Source = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *TrackRiskEventRequest) GetSession() Session {
	if o == nil || isNil(o.Session) {
		var ret Session
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetSessionOk() (*Session, bool) {
	if o == nil || isNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *TrackRiskEventRequest) HasSession() bool {
	if o != nil && !isNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given Session and assigns it to the Session field.
func (o *TrackRiskEventRequest) SetSession(v Session) {
	o.Session = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *TrackRiskEventRequest) GetDevice() RiskDevice {
	if o == nil || isNil(o.Device) {
		var ret RiskDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetDeviceOk() (*RiskDevice, bool) {
	if o == nil || isNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *TrackRiskEventRequest) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given RiskDevice and assigns it to the Device field.
func (o *TrackRiskEventRequest) SetDevice(v RiskDevice) {
	o.Device = &v
}

// GetFp returns the Fp field value if set, zero value otherwise.
func (o *TrackRiskEventRequest) GetFp() string {
	if o == nil || isNil(o.Fp) {
		var ret string
		return ret
	}
	return *o.Fp
}

// GetFpOk returns a tuple with the Fp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetFpOk() (*string, bool) {
	if o == nil || isNil(o.Fp) {
		return nil, false
	}
	return o.Fp, true
}

// HasFp returns a boolean if a field has been set.
func (o *TrackRiskEventRequest) HasFp() bool {
	if o != nil && !isNil(o.Fp) {
		return true
	}

	return false
}

// SetFp gets a reference to the given string and assigns it to the Fp field.
func (o *TrackRiskEventRequest) SetFp(v string) {
	o.Fp = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *TrackRiskEventRequest) GetPublished() string {
	if o == nil || isNil(o.Published) {
		var ret string
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackRiskEventRequest) GetPublishedOk() (*string, bool) {
	if o == nil || isNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *TrackRiskEventRequest) HasPublished() bool {
	if o != nil && !isNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given string and assigns it to the Published field.
func (o *TrackRiskEventRequest) SetPublished(v string) {
	o.Published = &v
}

func (o TrackRiskEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackRiskEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verb"] = o.Verb
	toSerialize["ip"] = o.Ip
	toSerialize["user_agent"] = o.UserAgent
	toSerialize["user"] = o.User
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Session) {
		toSerialize["session"] = o.Session
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Fp) {
		toSerialize["fp"] = o.Fp
	}
	if !isNil(o.Published) {
		toSerialize["published"] = o.Published
	}
	return toSerialize, nil
}

type NullableTrackRiskEventRequest struct {
	value *TrackRiskEventRequest
	isSet bool
}

func (v NullableTrackRiskEventRequest) Get() *TrackRiskEventRequest {
	return v.value
}

func (v *NullableTrackRiskEventRequest) Set(val *TrackRiskEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackRiskEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackRiskEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackRiskEventRequest(val *TrackRiskEventRequest) *NullableTrackRiskEventRequest {
	return &NullableTrackRiskEventRequest{value: val, isSet: true}
}

func (v NullableTrackRiskEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackRiskEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


