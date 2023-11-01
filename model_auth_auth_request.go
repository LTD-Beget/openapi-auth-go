/*
API Аутентификации

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiAuth

import (
	"encoding/json"
)

// AuthAuthRequest struct for AuthAuthRequest
type AuthAuthRequest struct {
	Login *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
	Code *string `json:"code,omitempty"`
	SaveMe *bool `json:"saveMe,omitempty"`
}

// NewAuthAuthRequest instantiates a new AuthAuthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAuthRequest() *AuthAuthRequest {
	this := AuthAuthRequest{}
	return &this
}

// NewAuthAuthRequestWithDefaults instantiates a new AuthAuthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAuthRequestWithDefaults() *AuthAuthRequest {
	this := AuthAuthRequest{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *AuthAuthRequest) GetLogin() string {
	if o == nil || isNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAuthRequest) GetLoginOk() (*string, bool) {
	if o == nil || isNil(o.Login) {
    return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *AuthAuthRequest) HasLogin() bool {
	if o != nil && !isNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *AuthAuthRequest) SetLogin(v string) {
	o.Login = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AuthAuthRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAuthRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthAuthRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AuthAuthRequest) SetPassword(v string) {
	o.Password = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AuthAuthRequest) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAuthRequest) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AuthAuthRequest) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AuthAuthRequest) SetCode(v string) {
	o.Code = &v
}

// GetSaveMe returns the SaveMe field value if set, zero value otherwise.
func (o *AuthAuthRequest) GetSaveMe() bool {
	if o == nil || isNil(o.SaveMe) {
		var ret bool
		return ret
	}
	return *o.SaveMe
}

// GetSaveMeOk returns a tuple with the SaveMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAuthRequest) GetSaveMeOk() (*bool, bool) {
	if o == nil || isNil(o.SaveMe) {
    return nil, false
	}
	return o.SaveMe, true
}

// HasSaveMe returns a boolean if a field has been set.
func (o *AuthAuthRequest) HasSaveMe() bool {
	if o != nil && !isNil(o.SaveMe) {
		return true
	}

	return false
}

// SetSaveMe gets a reference to the given bool and assigns it to the SaveMe field.
func (o *AuthAuthRequest) SetSaveMe(v bool) {
	o.SaveMe = &v
}

func (o AuthAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.SaveMe) {
		toSerialize["saveMe"] = o.SaveMe
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAuthRequest struct {
	value *AuthAuthRequest
	isSet bool
}

func (v NullableAuthAuthRequest) Get() *AuthAuthRequest {
	return v.value
}

func (v *NullableAuthAuthRequest) Set(val *AuthAuthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAuthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAuthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAuthRequest(val *AuthAuthRequest) *NullableAuthAuthRequest {
	return &NullableAuthAuthRequest{value: val, isSet: true}
}

func (v NullableAuthAuthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAuthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


