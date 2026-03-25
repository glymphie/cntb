# \InternalCheckCollectionsApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCheckCollection**](InternalCheckCollectionsApi.md#CancelCheckCollection) | **Patch** /internal/v1/troubleshooting/check-collections/{orgId}/{checkCollectionId} | Cancel check collection
[**GetCheckCollection**](InternalCheckCollectionsApi.md#GetCheckCollection) | **Get** /internal/v1/troubleshooting/check-collections/{orgId}/{checkCollectionId} | Get check collection
[**ListCheckCollections**](InternalCheckCollectionsApi.md#ListCheckCollections) | **Get** /internal/v1/troubleshooting/check-collections | List check collections
[**StartCheckCollection**](InternalCheckCollectionsApi.md#StartCheckCollection) | **Post** /internal/v1/troubleshooting/check-collections | Start check collection



## CancelCheckCollection

> CheckCollectionsGetResponse CancelCheckCollection(ctx, checkCollectionId, orgId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()

Cancel check collection



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
    checkCollectionId := float32(12345) // float32 | Check collection's id
    orgId := "cntb" // string | Org ID
    cancelRequest := *openapiclient.NewCancelRequest() // CancelRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalCheckCollectionsApi.CancelCheckCollection(context.Background(), checkCollectionId, orgId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalCheckCollectionsApi.CancelCheckCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelCheckCollection`: CheckCollectionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalCheckCollectionsApi.CancelCheckCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkCollectionId** | **float32** | Check collection&#39;s id | 
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCheckCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **cancelRequest** | [**CancelRequest**](CancelRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CheckCollectionsGetResponse**](CheckCollectionsGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckCollection

> CheckCollectionsGetResponse GetCheckCollection(ctx, checkCollectionId, orgId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get check collection



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
    checkCollectionId := float32(12345) // float32 | Check collection's id
    orgId := "cntb" // string | Org ID
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalCheckCollectionsApi.GetCheckCollection(context.Background(), checkCollectionId, orgId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalCheckCollectionsApi.GetCheckCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckCollection`: CheckCollectionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalCheckCollectionsApi.GetCheckCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkCollectionId** | **float32** | Check collection&#39;s id | 
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CheckCollectionsGetResponse**](CheckCollectionsGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCheckCollections

> CheckCollectionsListResponse ListCheckCollections(ctx).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).CheckCollectionTemplateId(checkCollectionTemplateId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).AccountId(accountId).Execute()

List check collections



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
    checkCollectionTemplateId := float32(12345) // float32 | Check Collection Template for this check collection (optional)
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
    resp, r, err := api_client.InternalCheckCollectionsApi.ListCheckCollections(context.Background()).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).CheckCollectionTemplateId(checkCollectionTemplateId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalCheckCollectionsApi.ListCheckCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCheckCollections`: CheckCollectionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalCheckCollectionsApi.ListCheckCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCheckCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **orgIds** | **[]string** | Org IDs | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **objectType** | **string** | Object type to be handled | 
 **objectId** | **string** | ID of the object, to be handled | 
 **checkCollectionTemplateId** | **float32** | Check Collection Template for this check collection | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **creationStartTime** | **time.Time** | Start of search time range for created date | 
 **creationEndTime** | **time.Time** | End of search time range for created date | 
 **modificationStartTime** | **time.Time** | Start of search time range for modified date | 
 **modificationEndTime** | **time.Time** | End of search time range for modified date | 
 **accountId** | **string** | Filter by account ID | 

### Return type

[**CheckCollectionsListResponse**](CheckCollectionsListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCheckCollection

> CheckCollectionsGetResponse StartCheckCollection(ctx).XRequestId(xRequestId).CheckCollectionCreateRequest(checkCollectionCreateRequest).XTraceId(xTraceId).Execute()

Start check collection



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
    checkCollectionCreateRequest := *openapiclient.NewCheckCollectionCreateRequest("vserver", "4711", float32(12345), "cntb", "DE-123") // CheckCollectionCreateRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalCheckCollectionsApi.StartCheckCollection(context.Background()).XRequestId(xRequestId).CheckCollectionCreateRequest(checkCollectionCreateRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalCheckCollectionsApi.StartCheckCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartCheckCollection`: CheckCollectionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalCheckCollectionsApi.StartCheckCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartCheckCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **checkCollectionCreateRequest** | [**CheckCollectionCreateRequest**](CheckCollectionCreateRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CheckCollectionsGetResponse**](CheckCollectionsGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

