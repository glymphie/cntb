# \ChecksApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelExtCheck**](ChecksApi.md#CancelExtCheck) | **Patch** /v1/troubleshooting/checks/{checkId} | Cancel check
[**GetExtCheck**](ChecksApi.md#GetExtCheck) | **Get** /v1/troubleshooting/checks/{checkId} | Get check
[**ListExtChecks**](ChecksApi.md#ListExtChecks) | **Get** /v1/troubleshooting/checks | List check
[**StartExtCheck**](ChecksApi.md#StartExtCheck) | **Post** /v1/troubleshooting/checks | Start check



## CancelExtCheck

> ExtChecksGetResponse CancelExtCheck(ctx, checkId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()

Cancel check



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
    checkId := float32(12345) // float32 | Check's id
    cancelRequest := *openapiclient.NewCancelRequest() // CancelRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.CancelExtCheck(context.Background(), checkId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.CancelExtCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelExtCheck`: ExtChecksGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.CancelExtCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **float32** | Check&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelExtCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **cancelRequest** | [**CancelRequest**](CancelRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ExtChecksGetResponse**](ExtChecksGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtCheck

> ExtChecksGetResponse GetExtCheck(ctx, checkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get check



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
    checkId := float32(12345) // float32 | Check's id
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.GetExtCheck(context.Background(), checkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.GetExtCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtCheck`: ExtChecksGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.GetExtCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **float32** | Check&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ExtChecksGetResponse**](ExtChecksGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtChecks

> ExtChecksListResponse ListExtChecks(ctx).XRequestId(xRequestId).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).Status(status).CheckCollectionId(checkCollectionId).CheckTemplateId(checkTemplateId).Execute()

List check



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
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    objectType := "vserver" // string | Object type to be handled (optional)
    objectId := "4711" // string | ID of the object, to be handled (optional)
    status := "failed" // string | Status of the handle (optional)
    checkCollectionId := float32(12345) // float32 | ID of check collection if started in scope of a collection (optional)
    checkTemplateId := float32(12345) // float32 | Check Template for this check (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.ListExtChecks(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).Status(status).CheckCollectionId(checkCollectionId).CheckTemplateId(checkTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ListExtChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtChecks`: ExtChecksListResponse
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ListExtChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **objectType** | **string** | Object type to be handled | 
 **objectId** | **string** | ID of the object, to be handled | 
 **status** | **string** | Status of the handle | 
 **checkCollectionId** | **float32** | ID of check collection if started in scope of a collection | 
 **checkTemplateId** | **float32** | Check Template for this check | 

### Return type

[**ExtChecksListResponse**](ExtChecksListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartExtCheck

> ExtChecksGetResponse StartExtCheck(ctx).XRequestId(xRequestId).BaseCheckCreateRequest(baseCheckCreateRequest).XTraceId(xTraceId).Execute()

Start check



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
    baseCheckCreateRequest := *openapiclient.NewBaseCheckCreateRequest("vserver", "4711", float32(12345)) // BaseCheckCreateRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.StartExtCheck(context.Background()).XRequestId(xRequestId).BaseCheckCreateRequest(baseCheckCreateRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.StartExtCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartExtCheck`: ExtChecksGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.StartExtCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartExtCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **baseCheckCreateRequest** | [**BaseCheckCreateRequest**](BaseCheckCreateRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ExtChecksGetResponse**](ExtChecksGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

