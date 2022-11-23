# \AuthServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthServiceAuth**](AuthServiceApi.md#AuthServiceAuth) | **Post** /v1/auth | 
[**AuthServiceLogout**](AuthServiceApi.md#AuthServiceLogout) | **Post** /v1/auth/logout | 
[**AuthServiceRefreshToken**](AuthServiceApi.md#AuthServiceRefreshToken) | **Post** /v1/auth/refresh | 
[**AuthServiceSwitchUser**](AuthServiceApi.md#AuthServiceSwitchUser) | **Post** /v1/auth/switch | 



## AuthServiceAuth

> AuthAuthResponse AuthServiceAuth(ctx).AuthAuthRequest(authAuthRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authAuthRequest := *openapiclient.NewAuthAuthRequest() // AuthAuthRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthServiceApi.AuthServiceAuth(context.Background()).AuthAuthRequest(authAuthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceApi.AuthServiceAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthServiceAuth`: AuthAuthResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthServiceApi.AuthServiceAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authAuthRequest** | [**AuthAuthRequest**](AuthAuthRequest.md) |  | 

### Return type

[**AuthAuthResponse**](AuthAuthResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceLogout

> map[string]interface{} AuthServiceLogout(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthServiceApi.AuthServiceLogout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceApi.AuthServiceLogout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthServiceLogout`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthServiceApi.AuthServiceLogout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceLogoutRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceRefreshToken

> AuthRefreshTokenResponse AuthServiceRefreshToken(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthServiceApi.AuthServiceRefreshToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceApi.AuthServiceRefreshToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthServiceRefreshToken`: AuthRefreshTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthServiceApi.AuthServiceRefreshToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceRefreshTokenRequest struct via the builder pattern


### Return type

[**AuthRefreshTokenResponse**](AuthRefreshTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceSwitchUser

> AuthSwitchUserResponse AuthServiceSwitchUser(ctx).AuthSwitchUserRequest(authSwitchUserRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authSwitchUserRequest := *openapiclient.NewAuthSwitchUserRequest() // AuthSwitchUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthServiceApi.AuthServiceSwitchUser(context.Background()).AuthSwitchUserRequest(authSwitchUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceApi.AuthServiceSwitchUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthServiceSwitchUser`: AuthSwitchUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthServiceApi.AuthServiceSwitchUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceSwitchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authSwitchUserRequest** | [**AuthSwitchUserRequest**](AuthSwitchUserRequest.md) |  | 

### Return type

[**AuthSwitchUserResponse**](AuthSwitchUserResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

