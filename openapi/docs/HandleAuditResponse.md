# HandleAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]HandleAuditResponseData**](HandleAuditResponseData.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewHandleAuditResponse

`func NewHandleAuditResponse(pagination PaginationMeta, data []HandleAuditResponseData, links Links, ) *HandleAuditResponse`

NewHandleAuditResponse instantiates a new HandleAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandleAuditResponseWithDefaults

`func NewHandleAuditResponseWithDefaults() *HandleAuditResponse`

NewHandleAuditResponseWithDefaults instantiates a new HandleAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *HandleAuditResponse) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *HandleAuditResponse) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *HandleAuditResponse) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *HandleAuditResponse) GetData() []HandleAuditResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HandleAuditResponse) GetDataOk() (*[]HandleAuditResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HandleAuditResponse) SetData(v []HandleAuditResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *HandleAuditResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HandleAuditResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HandleAuditResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


