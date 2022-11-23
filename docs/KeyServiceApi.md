# \KeyServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyServiceGetKey**](KeyServiceApi.md#KeyServiceGetKey) | **Get** /v1/auth/key | 



## KeyServiceGetKey

> KeyGetKeyResponse KeyServiceGetKey(ctx).Execute()



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
    resp, r, err := apiClient.KeyServiceApi.KeyServiceGetKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyServiceApi.KeyServiceGetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeyServiceGetKey`: KeyGetKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyServiceApi.KeyServiceGetKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKeyServiceGetKeyRequest struct via the builder pattern


### Return type

[**KeyGetKeyResponse**](KeyGetKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

