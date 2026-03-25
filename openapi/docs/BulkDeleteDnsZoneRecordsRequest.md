# BulkDeleteDnsZoneRecordsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordIds** | **[]int64** | List of zone record ids to delete | 

## Methods

### NewBulkDeleteDnsZoneRecordsRequest

`func NewBulkDeleteDnsZoneRecordsRequest(recordIds []int64, ) *BulkDeleteDnsZoneRecordsRequest`

NewBulkDeleteDnsZoneRecordsRequest instantiates a new BulkDeleteDnsZoneRecordsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteDnsZoneRecordsRequestWithDefaults

`func NewBulkDeleteDnsZoneRecordsRequestWithDefaults() *BulkDeleteDnsZoneRecordsRequest`

NewBulkDeleteDnsZoneRecordsRequestWithDefaults instantiates a new BulkDeleteDnsZoneRecordsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordIds

`func (o *BulkDeleteDnsZoneRecordsRequest) GetRecordIds() []int64`

GetRecordIds returns the RecordIds field if non-nil, zero value otherwise.

### GetRecordIdsOk

`func (o *BulkDeleteDnsZoneRecordsRequest) GetRecordIdsOk() (*[]int64, bool)`

GetRecordIdsOk returns a tuple with the RecordIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIds

`func (o *BulkDeleteDnsZoneRecordsRequest) SetRecordIds(v []int64)`

SetRecordIds sets RecordIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


