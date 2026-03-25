# \HandlesApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHandle**](HandlesApi.md#CreateHandle) | **Post** /v1/domains/handles | Create specific handle
[**ListHandles**](HandlesApi.md#ListHandles) | **Get** /v1/domains/handles | List all handles
[**RemoveHandle**](HandlesApi.md#RemoveHandle) | **Delete** /v1/domains/handles/{handleId} | Remove specific handle
[**RetrieveHandle**](HandlesApi.md#RetrieveHandle) | **Get** /v1/domains/handles/{handleId} | Get specific handle
[**SetDefaultHandle**](HandlesApi.md#SetDefaultHandle) | **Patch** /v1/domains/handles/{handleId}/default | Set default handle
[**UpdateHandle**](HandlesApi.md#UpdateHandle) | **Put** /v1/domains/handles/{handleId} | Update specific handle



## CreateHandle

> HandleCreateResponse CreateHandle(ctx).XRequestId(xRequestId).HandleCreateRequest(handleCreateRequest).XTraceId(xTraceId).Execute()

Create specific handle



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
    handleCreateRequest := *openapiclient.NewHandleCreateRequest("organization", "John", "Doe", "john.doe@test.com", "male", *openapiclient.NewHandleAddress("My Street", "12", "Munich", "DE", "12345"), *openapiclient.NewHandlePhone("+40", "123456789")) // HandleCreateRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HandlesApi.CreateHandle(context.Background()).XRequestId(xRequestId).HandleCreateRequest(handleCreateRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HandlesApi.CreateHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHandle`: HandleCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `HandlesApi.CreateHandle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **handleCreateRequest** | [**HandleCreateRequest**](HandleCreateRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**HandleCreateResponse**](HandleCreateResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHandles

> HandleListResponse ListHandles(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).ShowDefaults(showDefaults).Search(search).Countries(countries).HandleType(handleType).FirstName(firstName).LastName(lastName).Execute()

List all handles



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
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    name := "MyHandle" // string | Filter as substring match for handle name. (optional)
    showDefaults := true // bool | Filter handles to show or not the public handles (optional)
    search := "hello" // string | full text search on handles on handleid, organization name, handlename  (optional)
    countries := "DE,US" // string | list of country codes to filter handles that are available in these countries (comma separated) (optional)
    handleType := "person" // string | Filter handles by type, e.g. person, organization. (optional)
    firstName := "FirstName" // string | Filter handles by first name. (optional)
    lastName := "LastName" // string | Filter handles by last name. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HandlesApi.ListHandles(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).ShowDefaults(showDefaults).Search(search).Countries(countries).HandleType(handleType).FirstName(firstName).LastName(lastName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HandlesApi.ListHandles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHandles`: HandleListResponse
    fmt.Fprintf(os.Stdout, "Response from `HandlesApi.ListHandles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHandlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **name** | **string** | Filter as substring match for handle name. | 
 **showDefaults** | **bool** | Filter handles to show or not the public handles | 
 **search** | **string** | full text search on handles on handleid, organization name, handlename  | 
 **countries** | **string** | list of country codes to filter handles that are available in these countries (comma separated) | 
 **handleType** | **string** | Filter handles by type, e.g. person, organization. | 
 **firstName** | **string** | Filter handles by first name. | 
 **lastName** | **string** | Filter handles by last name. | 

### Return type

[**HandleListResponse**](HandleListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveHandle

> RemoveHandle(ctx, handleId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Remove specific handle



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
    handleId := "CA123O1" // string | The identifier of the handle
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HandlesApi.RemoveHandle(context.Background(), handleId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HandlesApi.RemoveHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handleId** | **string** | The identifier of the handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveHandle

> HandleFindResponse RetrieveHandle(ctx, handleId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific handle



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
    handleId := "CA123O1" // string | The identifier of the handle
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HandlesApi.RetrieveHandle(context.Background(), handleId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HandlesApi.RetrieveHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveHandle`: HandleFindResponse
    fmt.Fprintf(os.Stdout, "Response from `HandlesApi.RetrieveHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handleId** | **string** | The identifier of the handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**HandleFindResponse**](HandleFindResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultHandle

> SetDefaultHandleResponse SetDefaultHandle(ctx, handleId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Set default handle



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
    handleId := "CA123O1" // string | The identifier of the handle
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HandlesApi.SetDefaultHandle(context.Background(), handleId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HandlesApi.SetDefaultHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDefaultHandle`: SetDefaultHandleResponse
    fmt.Fprintf(os.Stdout, "Response from `HandlesApi.SetDefaultHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handleId** | **string** | The identifier of the handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**SetDefaultHandleResponse**](SetDefaultHandleResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHandle

> HandlePatchResponse UpdateHandle(ctx, handleId).XRequestId(xRequestId).HandlePatchRequest(handlePatchRequest).XTraceId(xTraceId).Execute()

Update specific handle



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
    handleId := "CA123O1" // string | The identifier of the handle
    handlePatchRequest := *openapiclient.NewHandlePatchRequest("john.doe@test.com", "male", *openapiclient.NewHandleAddress("My Street", "12", "Munich", "DE", "12345"), *openapiclient.NewHandlePhone("+40", "123456789")) // HandlePatchRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HandlesApi.UpdateHandle(context.Background(), handleId).XRequestId(xRequestId).HandlePatchRequest(handlePatchRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HandlesApi.UpdateHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHandle`: HandlePatchResponse
    fmt.Fprintf(os.Stdout, "Response from `HandlesApi.UpdateHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handleId** | **string** | The identifier of the handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **handlePatchRequest** | [**HandlePatchRequest**](HandlePatchRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**HandlePatchResponse**](HandlePatchResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

