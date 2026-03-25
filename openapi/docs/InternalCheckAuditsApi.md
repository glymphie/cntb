# \InternalCheckAuditsApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveChecksAuditsList**](InternalCheckAuditsApi.md#RetrieveChecksAuditsList) | **Get** /internal/v1/troubleshooting/checks/audits | List history about your Data (audit)



## RetrieveChecksAuditsList

> ChecksAuditListResponse RetrieveChecksAuditsList(ctx).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).AccountId(accountId).RequestId(requestId).ForeignChangedBy(foreignChangedBy).ChangedBy(changedBy).CheckId(checkId).Execute()

List history about your Data (audit)



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
    accountId := "DE-123" // string | Filter by account ID (optional)
    requestId := "D5FD9FAF-58C0-4406-8F46-F449B8E4FEC3" // string | The requestId of the API call which led to the change. (optional)
    foreignChangedBy := "23cbb6d6-cb11-4330-bdff-7bb791df2e23" // string | Foreign uerId of the user which led to the change. (optional)
    changedBy := "23cbb6d6-cb11-4330-bdff-7bb791df2e23" // string | UserId of the user which led to the change. (optional)
    checkId := float32(12345) // float32 | Check's id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalCheckAuditsApi.RetrieveChecksAuditsList(context.Background()).XRequestId(xRequestId).OrgIds(orgIds).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).AccountId(accountId).RequestId(requestId).ForeignChangedBy(foreignChangedBy).ChangedBy(changedBy).CheckId(checkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalCheckAuditsApi.RetrieveChecksAuditsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveChecksAuditsList`: ChecksAuditListResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalCheckAuditsApi.RetrieveChecksAuditsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveChecksAuditsListRequest struct via the builder pattern


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
 **accountId** | **string** | Filter by account ID | 
 **requestId** | **string** | The requestId of the API call which led to the change. | 
 **foreignChangedBy** | **string** | Foreign uerId of the user which led to the change. | 
 **changedBy** | **string** | UserId of the user which led to the change. | 
 **checkId** | **float32** | Check&#39;s id | 

### Return type

[**ChecksAuditListResponse**](ChecksAuditListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

