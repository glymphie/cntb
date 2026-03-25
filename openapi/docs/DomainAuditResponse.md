# DomainAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]DomainAuditResponseData**](DomainAuditResponseData.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewDomainAuditResponse

`func NewDomainAuditResponse(pagination PaginationMeta, data []DomainAuditResponseData, links Links, ) *DomainAuditResponse`

NewDomainAuditResponse instantiates a new DomainAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAuditResponseWithDefaults

`func NewDomainAuditResponseWithDefaults() *DomainAuditResponse`

NewDomainAuditResponseWithDefaults instantiates a new DomainAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *DomainAuditResponse) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainAuditResponse) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainAuditResponse) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *DomainAuditResponse) GetData() []DomainAuditResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainAuditResponse) GetDataOk() (*[]DomainAuditResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainAuditResponse) SetData(v []DomainAuditResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *DomainAuditResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainAuditResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainAuditResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


