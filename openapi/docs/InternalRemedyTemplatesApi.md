# \InternalRemedyTemplatesApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemedyTemplate**](InternalRemedyTemplatesApi.md#GetRemedyTemplate) | **Get** /internal/v1/troubleshooting/remedy-templates/{orgId}/{remedyTemplateId} | Get remedy
[**ListRemedyTemplates**](InternalRemedyTemplatesApi.md#ListRemedyTemplates) | **Get** /internal/v1/troubleshooting/remedy-templates | List remedy templates



## GetRemedyTemplate

> RemedyTemplatesGetResponse GetRemedyTemplate(ctx, orgId, remedyTemplateId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

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
    orgId := "cntb" // string | Org ID
    remedyTemplateId := float32(12345) // float32 | Remedy template's id
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalRemedyTemplatesApi.GetRemedyTemplate(context.Background(), orgId, remedyTemplateId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalRemedyTemplatesApi.GetRemedyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemedyTemplate`: RemedyTemplatesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalRemedyTemplatesApi.GetRemedyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**remedyTemplateId** | **float32** | Remedy template&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemedyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**RemedyTemplatesGetResponse**](RemedyTemplatesGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemedyTemplates

> RemedyTemplatesListResponse ListRemedyTemplates(ctx).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).AccountId(accountId).Internal(internal).ObjectType(objectType).CollectorClass(collectorClass).RemedyClass(remedyClass).Execute()

List remedy templates



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
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    creationStartTime := time.Now() // time.Time | Start of search time range for created date (optional)
    creationEndTime := time.Now() // time.Time | End of search time range for created date (optional)
    modificationStartTime := time.Now() // time.Time | Start of search time range for modified date (optional)
    modificationEndTime := time.Now() // time.Time | End of search time range for modified date (optional)
    accountId := "DE-123" // string | Filter by account ID (optional)
    internal := false // bool | Is remedy only internal (not shown to the customer) (optional)
    objectType := "vserver" // string | Object type for which the remedy template can be used (optional)
    collectorClass := "InstanceCollector.ts" // string | Class used to collect the required information for the remedy (optional)
    remedyClass := "PingRemedy.ts" // string | Class used to perform the remedy (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalRemedyTemplatesApi.ListRemedyTemplates(context.Background()).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).AccountId(accountId).Internal(internal).ObjectType(objectType).CollectorClass(collectorClass).RemedyClass(remedyClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalRemedyTemplatesApi.ListRemedyTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRemedyTemplates`: RemedyTemplatesListResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalRemedyTemplatesApi.ListRemedyTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRemedyTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **orgIds** | **[]string** | Org IDs | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **creationStartTime** | **time.Time** | Start of search time range for created date | 
 **creationEndTime** | **time.Time** | End of search time range for created date | 
 **modificationStartTime** | **time.Time** | Start of search time range for modified date | 
 **modificationEndTime** | **time.Time** | End of search time range for modified date | 
 **accountId** | **string** | Filter by account ID | 
 **internal** | **bool** | Is remedy only internal (not shown to the customer) | 
 **objectType** | **string** | Object type for which the remedy template can be used | 
 **collectorClass** | **string** | Class used to collect the required information for the remedy | 
 **remedyClass** | **string** | Class used to perform the remedy | 

### Return type

[**RemedyTemplatesListResponse**](RemedyTemplatesListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

