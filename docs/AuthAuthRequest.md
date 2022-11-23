# AuthAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**SaveMe** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthAuthRequest

`func NewAuthAuthRequest() *AuthAuthRequest`

NewAuthAuthRequest instantiates a new AuthAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAuthRequestWithDefaults

`func NewAuthAuthRequestWithDefaults() *AuthAuthRequest`

NewAuthAuthRequestWithDefaults instantiates a new AuthAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *AuthAuthRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AuthAuthRequest) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AuthAuthRequest) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *AuthAuthRequest) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetPassword

`func (o *AuthAuthRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthAuthRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthAuthRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthAuthRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCode

`func (o *AuthAuthRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthAuthRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthAuthRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthAuthRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSaveMe

`func (o *AuthAuthRequest) GetSaveMe() bool`

GetSaveMe returns the SaveMe field if non-nil, zero value otherwise.

### GetSaveMeOk

`func (o *AuthAuthRequest) GetSaveMeOk() (*bool, bool)`

GetSaveMeOk returns a tuple with the SaveMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveMe

`func (o *AuthAuthRequest) SetSaveMe(v bool)`

SetSaveMe sets SaveMe field to given value.

### HasSaveMe

`func (o *AuthAuthRequest) HasSaveMe() bool`

HasSaveMe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


