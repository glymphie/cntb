# CreatePtrRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ptr** | **string** | PTR Record name | 
**Ip** | **string** | IP Address | 
**Ttl** | **int64** | Time to live for the PTR record in seconds | 

## Methods

### NewCreatePtrRecordRequest

`func NewCreatePtrRecordRequest(ptr string, ip string, ttl int64, ) *CreatePtrRecordRequest`

NewCreatePtrRecordRequest instantiates a new CreatePtrRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePtrRecordRequestWithDefaults

`func NewCreatePtrRecordRequestWithDefaults() *CreatePtrRecordRequest`

NewCreatePtrRecordRequestWithDefaults instantiates a new CreatePtrRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPtr

`func (o *CreatePtrRecordRequest) GetPtr() string`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *CreatePtrRecordRequest) GetPtrOk() (*string, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *CreatePtrRecordRequest) SetPtr(v string)`

SetPtr sets Ptr field to given value.


### GetIp

`func (o *CreatePtrRecordRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreatePtrRecordRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreatePtrRecordRequest) SetIp(v string)`

SetIp sets Ip field to given value.


### GetTtl

`func (o *CreatePtrRecordRequest) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreatePtrRecordRequest) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreatePtrRecordRequest) SetTtl(v int64)`

SetTtl sets Ttl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


