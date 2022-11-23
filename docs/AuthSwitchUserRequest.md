# AuthSwitchUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthSwitchUserRequest

`func NewAuthSwitchUserRequest() *AuthSwitchUserRequest`

NewAuthSwitchUserRequest instantiates a new AuthSwitchUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSwitchUserRequestWithDefaults

`func NewAuthSwitchUserRequestWithDefaults() *AuthSwitchUserRequest`

NewAuthSwitchUserRequestWithDefaults instantiates a new AuthSwitchUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *AuthSwitchUserRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AuthSwitchUserRequest) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AuthSwitchUserRequest) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *AuthSwitchUserRequest) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetPassword

`func (o *AuthSwitchUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthSwitchUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthSwitchUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthSwitchUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCode

`func (o *AuthSwitchUserRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthSwitchUserRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthSwitchUserRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthSwitchUserRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


