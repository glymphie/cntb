# \InternalCheckReplayApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplayCheck**](InternalCheckReplayApi.md#ReplayCheck) | **Post** /internal/v1/troubleshooting/checks/replays | Replay changes for Check



## ReplayCheck

> ReplayResponse ReplayCheck(ctx).XRequestId(xRequestId).ChecksReplayRequest(checksReplayRequest).XTraceId(xTraceId).Execute()

Replay changes for Check



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    checksReplayRequest := *openapiclient.NewChecksReplayRequest("cntb", "DE-123") // ChecksReplayRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalCheckReplayApi.ReplayCheck(context.Background()).XRequestId(xRequestId).ChecksReplayRequest(checksReplayRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalCheckReplayApi.ReplayCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplayCheck`: ReplayResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalCheckReplayApi.ReplayCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplayCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **checksReplayRequest** | [**ChecksReplayRequest**](ChecksReplayRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ReplayResponse**](ReplayResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

