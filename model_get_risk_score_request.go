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

// checks if the GetRiskScoreRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRiskScoreRequest{}

// GetRiskScoreRequest struct for GetRiskScoreRequest
type GetRiskScoreRequest struct {
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
}

// NewGetRiskScoreRequest instantiates a new GetRiskScoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRiskScoreRequest(ip string, userAgent string, user RiskUser) *GetRiskScoreRequest {
	this := GetRiskScoreRequest{}
	this.Ip = ip
	this.UserAgent = userAgent
	this.User = user
	return &this
}

// NewGetRiskScoreRequestWithDefaults instantiates a new GetRiskScoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRiskScoreRequestWithDefaults() *GetRiskScoreRequest {
	this := GetRiskScoreRequest{}
	return &this
}

// GetIp returns the Ip field value
func (o *GetRiskScoreRequest) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *GetRiskScoreRequest) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *GetRiskScoreRequest) SetIp(v string) {
	o.Ip = v
}

// GetUserAgent returns the UserAgent field value
func (o *GetRiskScoreRequest) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *GetRiskScoreRequest) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *GetRiskScoreRequest) SetUserAgent(v string) {
	o.UserAgent = v
}

// GetUser returns the User field value
func (o *GetRiskScoreRequest) GetUser() RiskUser {
	if o == nil {
		var ret RiskUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GetRiskScoreRequest) GetUserOk() (*RiskUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GetRiskScoreRequest) SetUser(v RiskUser) {
	o.User = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GetRiskScoreRequest) GetSource() Source {
	if o == nil || isNil(o.Source) {
		var ret Source
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRiskScoreRequest) GetSourceOk() (*Source, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GetRiskScoreRequest) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Source and assigns it to the Source field.
func (o *GetRiskScoreRequest) SetSource(v Source) {
	o.Source = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *GetRiskScoreRequest) GetSession() Session {
	if o == nil || isNil(o.Session) {
		var ret Session
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRiskScoreRequest) GetSessionOk() (*Session, bool) {
	if o == nil || isNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *GetRiskScoreRequest) HasSession() bool {
	if o != nil && !isNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given Session and assigns it to the Session field.
func (o *GetRiskScoreRequest) SetSession(v Session) {
	o.Session = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *GetRiskScoreRequest) GetDevice() RiskDevice {
	if o == nil || isNil(o.Device) {
		var ret RiskDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRiskScoreRequest) GetDeviceOk() (*RiskDevice, bool) {
	if o == nil || isNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *GetRiskScoreRequest) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given RiskDevice and assigns it to the Device field.
func (o *GetRiskScoreRequest) SetDevice(v RiskDevice) {
	o.Device = &v
}

// GetFp returns the Fp field value if set, zero value otherwise.
func (o *GetRiskScoreRequest) GetFp() string {
	if o == nil || isNil(o.Fp) {
		var ret string
		return ret
	}
	return *o.Fp
}

// GetFpOk returns a tuple with the Fp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRiskScoreRequest) GetFpOk() (*string, bool) {
	if o == nil || isNil(o.Fp) {
		return nil, false
	}
	return o.Fp, true
}

// HasFp returns a boolean if a field has been set.
func (o *GetRiskScoreRequest) HasFp() bool {
	if o != nil && !isNil(o.Fp) {
		return true
	}

	return false
}

// SetFp gets a reference to the given string and assigns it to the Fp field.
func (o *GetRiskScoreRequest) SetFp(v string) {
	o.Fp = &v
}

func (o GetRiskScoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRiskScoreRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableGetRiskScoreRequest struct {
	value *GetRiskScoreRequest
	isSet bool
}

func (v NullableGetRiskScoreRequest) Get() *GetRiskScoreRequest {
	return v.value
}

func (v *NullableGetRiskScoreRequest) Set(val *GetRiskScoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRiskScoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRiskScoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRiskScoreRequest(val *GetRiskScoreRequest) *NullableGetRiskScoreRequest {
	return &NullableGetRiskScoreRequest{value: val, isSet: true}
}

func (v NullableGetRiskScoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRiskScoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


