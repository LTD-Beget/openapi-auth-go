# AuthAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthAuthResponse

`func NewAuthAuthResponse() *AuthAuthResponse`

NewAuthAuthResponse instantiates a new AuthAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAuthResponseWithDefaults

`func NewAuthAuthResponseWithDefaults() *AuthAuthResponse`

NewAuthAuthResponseWithDefaults instantiates a new AuthAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AuthAuthResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthAuthResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthAuthResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthAuthResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetError

`func (o *AuthAuthResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthAuthResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthAuthResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AuthAuthResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


