# AccountsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]AccountOrgResponse**](AccountOrgResponse.md) |  | 
**Links** | **map[string]interface{}** |  | 

## Methods

### NewAccountsListResponse

`func NewAccountsListResponse(pagination PaginationMeta, data []AccountOrgResponse, links map[string]interface{}, ) *AccountsListResponse`

NewAccountsListResponse instantiates a new AccountsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsListResponseWithDefaults

`func NewAccountsListResponseWithDefaults() *AccountsListResponse`

NewAccountsListResponseWithDefaults instantiates a new AccountsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AccountsListResponse) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AccountsListResponse) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AccountsListResponse) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *AccountsListResponse) GetData() []AccountOrgResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountsListResponse) GetDataOk() (*[]AccountOrgResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountsListResponse) SetData(v []AccountOrgResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *AccountsListResponse) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccountsListResponse) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccountsListResponse) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


