# RemediesReplayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | Id of your organization, if unknown please contact us | 
**AccountId** | **string** | Account Id | 
**CreationStartTime** | Pointer to **time.Time** | Earliest creation date of changes to replay | [optional] 
**CreationEndTime** | Pointer to **time.Time** | Latest creation date of changes to replay | [optional] 
**Rate** | Pointer to **float32** | Message publishing frequency. How many messages per second get published. Default: 20 | [optional] 
**RemedyIds** | Pointer to **[]float32** | Remedy&#39;s id | [optional] 

## Methods

### NewRemediesReplayRequest

`func NewRemediesReplayRequest(orgId string, accountId string, ) *RemediesReplayRequest`

NewRemediesReplayRequest instantiates a new RemediesReplayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediesReplayRequestWithDefaults

`func NewRemediesReplayRequestWithDefaults() *RemediesReplayRequest`

NewRemediesReplayRequestWithDefaults instantiates a new RemediesReplayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *RemediesReplayRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RemediesReplayRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RemediesReplayRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *RemediesReplayRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RemediesReplayRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RemediesReplayRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreationStartTime

`func (o *RemediesReplayRequest) GetCreationStartTime() time.Time`

GetCreationStartTime returns the CreationStartTime field if non-nil, zero value otherwise.

### GetCreationStartTimeOk

`func (o *RemediesReplayRequest) GetCreationStartTimeOk() (*time.Time, bool)`

GetCreationStartTimeOk returns a tuple with the CreationStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStartTime

`func (o *RemediesReplayRequest) SetCreationStartTime(v time.Time)`

SetCreationStartTime sets CreationStartTime field to given value.

### HasCreationStartTime

`func (o *RemediesReplayRequest) HasCreationStartTime() bool`

HasCreationStartTime returns a boolean if a field has been set.

### GetCreationEndTime

`func (o *RemediesReplayRequest) GetCreationEndTime() time.Time`

GetCreationEndTime returns the CreationEndTime field if non-nil, zero value otherwise.

### GetCreationEndTimeOk

`func (o *RemediesReplayRequest) GetCreationEndTimeOk() (*time.Time, bool)`

GetCreationEndTimeOk returns a tuple with the CreationEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationEndTime

`func (o *RemediesReplayRequest) SetCreationEndTime(v time.Time)`

SetCreationEndTime sets CreationEndTime field to given value.

### HasCreationEndTime

`func (o *RemediesReplayRequest) HasCreationEndTime() bool`

HasCreationEndTime returns a boolean if a field has been set.

### GetRate

`func (o *RemediesReplayRequest) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *RemediesReplayRequest) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *RemediesReplayRequest) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *RemediesReplayRequest) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRemedyIds

`func (o *RemediesReplayRequest) GetRemedyIds() []float32`

GetRemedyIds returns the RemedyIds field if non-nil, zero value otherwise.

### GetRemedyIdsOk

`func (o *RemediesReplayRequest) GetRemedyIdsOk() (*[]float32, bool)`

GetRemedyIdsOk returns a tuple with the RemedyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyIds

`func (o *RemediesReplayRequest) SetRemedyIds(v []float32)`

SetRemedyIds sets RemedyIds field to given value.

### HasRemedyIds

`func (o *RemediesReplayRequest) HasRemedyIds() bool`

HasRemedyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


