# AuthRefreshTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthRefreshTokenResponse

`func NewAuthRefreshTokenResponse() *AuthRefreshTokenResponse`

NewAuthRefreshTokenResponse instantiates a new AuthRefreshTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRefreshTokenResponseWithDefaults

`func NewAuthRefreshTokenResponseWithDefaults() *AuthRefreshTokenResponse`

NewAuthRefreshTokenResponseWithDefaults instantiates a new AuthRefreshTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AuthRefreshTokenResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthRefreshTokenResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthRefreshTokenResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthRefreshTokenResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetError

`func (o *AuthRefreshTokenResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthRefreshTokenResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthRefreshTokenResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AuthRefreshTokenResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


