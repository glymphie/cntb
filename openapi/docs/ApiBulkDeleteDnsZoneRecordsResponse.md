# ApiBulkDeleteDnsZoneRecordsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BulkDeleteResultResponse**](BulkDeleteResultResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewApiBulkDeleteDnsZoneRecordsResponse

`func NewApiBulkDeleteDnsZoneRecordsResponse(data []BulkDeleteResultResponse, links SelfLinks, ) *ApiBulkDeleteDnsZoneRecordsResponse`

NewApiBulkDeleteDnsZoneRecordsResponse instantiates a new ApiBulkDeleteDnsZoneRecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBulkDeleteDnsZoneRecordsResponseWithDefaults

`func NewApiBulkDeleteDnsZoneRecordsResponseWithDefaults() *ApiBulkDeleteDnsZoneRecordsResponse`

NewApiBulkDeleteDnsZoneRecordsResponseWithDefaults instantiates a new ApiBulkDeleteDnsZoneRecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiBulkDeleteDnsZoneRecordsResponse) GetData() []BulkDeleteResultResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiBulkDeleteDnsZoneRecordsResponse) GetDataOk() (*[]BulkDeleteResultResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiBulkDeleteDnsZoneRecordsResponse) SetData(v []BulkDeleteResultResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ApiBulkDeleteDnsZoneRecordsResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiBulkDeleteDnsZoneRecordsResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiBulkDeleteDnsZoneRecordsResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


