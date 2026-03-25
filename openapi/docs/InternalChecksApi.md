# \InternalChecksApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCheck**](InternalChecksApi.md#CancelCheck) | **Patch** /internal/v1/troubleshooting/checks/{orgId}/{checkId} | Cancel check
[**GetCheck**](InternalChecksApi.md#GetCheck) | **Get** /internal/v1/troubleshooting/checks/{orgId}/{checkId} | Get check
[**ListChecks**](InternalChecksApi.md#ListChecks) | **Get** /internal/v1/troubleshooting/checks | List check
[**StartCheck**](InternalChecksApi.md#StartCheck) | **Post** /internal/v1/troubleshooting/checks | Start check



## CancelCheck

> ChecksGetResponse CancelCheck(ctx, checkId, orgId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()

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
    orgId := "cntb" // string | Org ID
    cancelRequest := *openapiclient.NewCancelRequest() // CancelRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalChecksApi.CancelCheck(context.Background(), checkId, orgId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalChecksApi.CancelCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelCheck`: ChecksGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalChecksApi.CancelCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **float32** | Check&#39;s id | 
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **cancelRequest** | [**CancelRequest**](CancelRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ChecksGetResponse**](ChecksGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheck

> ChecksGetResponse GetCheck(ctx, checkId, orgId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

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
    orgId := "cntb" // string | Org ID
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalChecksApi.GetCheck(context.Background(), checkId, orgId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalChecksApi.GetCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheck`: ChecksGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalChecksApi.GetCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **float32** | Check&#39;s id | 
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ChecksGetResponse**](ChecksGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChecks

> ChecksListResponse ListChecks(ctx).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).Status(status).CheckCollectionId(checkCollectionId).CheckTemplateId(checkTemplateId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).AccountId(accountId).Execute()

List check



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    orgIds := []string{"Inner_example"} // []string | Org IDs
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    objectType := "vserver" // string | Object type to be handled (optional)
    objectId := "4711" // string | ID of the object, to be handled (optional)
    status := "failed" // string | Status of the handle (optional)
    checkCollectionId := float32(12345) // float32 | ID of check collection if started in scope of a collection (optional)
    checkTemplateId := float32(12345) // float32 | Check Template for this check (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    creationStartTime := time.Now() // time.Time | Start of search time range for created date (optional)
    creationEndTime := time.Now() // time.Time | End of search time range for created date (optional)
    modificationStartTime := time.Now() // time.Time | Start of search time range for modified date (optional)
    modificationEndTime := time.Now() // time.Time | End of search time range for modified date (optional)
    accountId := "DE-123" // string | Filter by account ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalChecksApi.ListChecks(context.Background()).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).Status(status).CheckCollectionId(checkCollectionId).CheckTemplateId(checkTemplateId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalChecksApi.ListChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChecks`: ChecksListResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalChecksApi.ListChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **orgIds** | **[]string** | Org IDs | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **objectType** | **string** | Object type to be handled | 
 **objectId** | **string** | ID of the object, to be handled | 
 **status** | **string** | Status of the handle | 
 **checkCollectionId** | **float32** | ID of check collection if started in scope of a collection | 
 **checkTemplateId** | **float32** | Check Template for this check | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **creationStartTime** | **time.Time** | Start of search time range for created date | 
 **creationEndTime** | **time.Time** | End of search time range for created date | 
 **modificationStartTime** | **time.Time** | Start of search time range for modified date | 
 **modificationEndTime** | **time.Time** | End of search time range for modified date | 
 **accountId** | **string** | Filter by account ID | 

### Return type

[**ChecksListResponse**](ChecksListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCheck

> ChecksGetResponse StartCheck(ctx).XRequestId(xRequestId).CheckCreateRequest(checkCreateRequest).XTraceId(xTraceId).Execute()

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
    checkCreateRequest := *openapiclient.NewCheckCreateRequest("vserver", "4711", float32(12345), "cntb", "DE-123") // CheckCreateRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalChecksApi.StartCheck(context.Background()).XRequestId(xRequestId).CheckCreateRequest(checkCreateRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalChecksApi.StartCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartCheck`: ChecksGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalChecksApi.StartCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **checkCreateRequest** | [**CheckCreateRequest**](CheckCreateRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ChecksGetResponse**](ChecksGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

