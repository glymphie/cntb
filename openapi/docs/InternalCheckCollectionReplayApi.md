# \InternalCheckCollectionReplayApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplayCheckCollection**](InternalCheckCollectionReplayApi.md#ReplayCheckCollection) | **Post** /internal/v1/troubleshooting/check-collections/replays | Replay changes for Check



## ReplayCheckCollection

> ReplayResponse ReplayCheckCollection(ctx).XRequestId(xRequestId).CheckCollectionsReplayRequest(checkCollectionsReplayRequest).XTraceId(xTraceId).Execute()

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
    checkCollectionsReplayRequest := *openapiclient.NewCheckCollectionsReplayRequest("cntb", "DE-123") // CheckCollectionsReplayRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalCheckCollectionReplayApi.ReplayCheckCollection(context.Background()).XRequestId(xRequestId).CheckCollectionsReplayRequest(checkCollectionsReplayRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalCheckCollectionReplayApi.ReplayCheckCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplayCheckCollection`: ReplayResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalCheckCollectionReplayApi.ReplayCheckCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplayCheckCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **checkCollectionsReplayRequest** | [**CheckCollectionsReplayRequest**](CheckCollectionsReplayRequest.md) |  | 
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

