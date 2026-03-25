# \InternalRemediesApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRemedy**](InternalRemediesApi.md#CancelRemedy) | **Patch** /internal/v1/troubleshooting/remedies/{orgId}/{remedyId} | Cancel remedy
[**GetRemedy**](InternalRemediesApi.md#GetRemedy) | **Get** /internal/v1/troubleshooting/remedies/{orgId}/{remedyId} | Get remedy
[**ListRemedies**](InternalRemediesApi.md#ListRemedies) | **Get** /internal/v1/troubleshooting/remedies | List remedy
[**StartRemedy**](InternalRemediesApi.md#StartRemedy) | **Post** /internal/v1/troubleshooting/remedies | Start remedy



## CancelRemedy

> RemediesGetResponse CancelRemedy(ctx, remedyId, orgId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()

Cancel remedy



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
    remedyId := float32(12345) // float32 | Remedy's id
    orgId := "cntb" // string | Org ID
    cancelRequest := *openapiclient.NewCancelRequest() // CancelRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalRemediesApi.CancelRemedy(context.Background(), remedyId, orgId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalRemediesApi.CancelRemedy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelRemedy`: RemediesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalRemediesApi.CancelRemedy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remedyId** | **float32** | Remedy&#39;s id | 
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRemedyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **cancelRequest** | [**CancelRequest**](CancelRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**RemediesGetResponse**](RemediesGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemedy

> RemediesGetResponse GetRemedy(ctx, remedyId, orgId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get remedy



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
    remedyId := float32(12345) // float32 | Remedy's id
    orgId := "cntb" // string | Org ID
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalRemediesApi.GetRemedy(context.Background(), remedyId, orgId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalRemediesApi.GetRemedy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemedy`: RemediesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalRemediesApi.GetRemedy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remedyId** | **float32** | Remedy&#39;s id | 
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemedyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**RemediesGetResponse**](RemediesGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemedies

> RemediesListResponse ListRemedies(ctx).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).Status(status).RemedyCollectionId(remedyCollectionId).RemedyTemplateId(remedyTemplateId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).AccountId(accountId).Execute()

List remedy



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
    remedyCollectionId := float32(12345) // float32 | ID of remedy collection if started in scope of a collection (optional)
    remedyTemplateId := float32(12345) // float32 | Remedy Template for this check (optional)
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
    resp, r, err := api_client.InternalRemediesApi.ListRemedies(context.Background()).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).Status(status).RemedyCollectionId(remedyCollectionId).RemedyTemplateId(remedyTemplateId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalRemediesApi.ListRemedies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRemedies`: RemediesListResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalRemediesApi.ListRemedies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRemediesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **orgIds** | **[]string** | Org IDs | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **objectType** | **string** | Object type to be handled | 
 **objectId** | **string** | ID of the object, to be handled | 
 **status** | **string** | Status of the handle | 
 **remedyCollectionId** | **float32** | ID of remedy collection if started in scope of a collection | 
 **remedyTemplateId** | **float32** | Remedy Template for this check | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **creationStartTime** | **time.Time** | Start of search time range for created date | 
 **creationEndTime** | **time.Time** | End of search time range for created date | 
 **modificationStartTime** | **time.Time** | Start of search time range for modified date | 
 **modificationEndTime** | **time.Time** | End of search time range for modified date | 
 **accountId** | **string** | Filter by account ID | 

### Return type

[**RemediesListResponse**](RemediesListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRemedy

> RemediesGetResponse StartRemedy(ctx).XRequestId(xRequestId).RemediesCreateRequest(remediesCreateRequest).XTraceId(xTraceId).Execute()

Start remedy



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
    remediesCreateRequest := *openapiclient.NewRemediesCreateRequest("vserver", "4711", float32(12345), "cntb", "DE-123") // RemediesCreateRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalRemediesApi.StartRemedy(context.Background()).XRequestId(xRequestId).RemediesCreateRequest(remediesCreateRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalRemediesApi.StartRemedy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartRemedy`: RemediesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalRemediesApi.StartRemedy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartRemedyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **remediesCreateRequest** | [**RemediesCreateRequest**](RemediesCreateRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**RemediesGetResponse**](RemediesGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

