# BulkDeleteResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**FailedIds** | **[]int64** | Failed zone record IDs | 

## Methods

### NewBulkDeleteResultResponse

`func NewBulkDeleteResultResponse(tenantId string, customerId string, failedIds []int64, ) *BulkDeleteResultResponse`

NewBulkDeleteResultResponse instantiates a new BulkDeleteResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteResultResponseWithDefaults

`func NewBulkDeleteResultResponseWithDefaults() *BulkDeleteResultResponse`

NewBulkDeleteResultResponseWithDefaults instantiates a new BulkDeleteResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *BulkDeleteResultResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BulkDeleteResultResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BulkDeleteResultResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *BulkDeleteResultResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *BulkDeleteResultResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *BulkDeleteResultResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetFailedIds

`func (o *BulkDeleteResultResponse) GetFailedIds() []int64`

GetFailedIds returns the FailedIds field if non-nil, zero value otherwise.

### GetFailedIdsOk

`func (o *BulkDeleteResultResponse) GetFailedIdsOk() (*[]int64, bool)`

GetFailedIdsOk returns a tuple with the FailedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedIds

`func (o *BulkDeleteResultResponse) SetFailedIds(v []int64)`

SetFailedIds sets FailedIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


