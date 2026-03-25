# ChecksReplayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | Id of your organization, if unknown please contact us | 
**AccountId** | **string** | Account Id | 
**CreationStartTime** | Pointer to **time.Time** | Earliest creation date of changes to replay | [optional] 
**CreationEndTime** | Pointer to **time.Time** | Latest creation date of changes to replay | [optional] 
**Rate** | Pointer to **float32** | Message publishing frequency. How many messages per second get published. Default: 20 | [optional] 
**CheckIds** | Pointer to **[]float32** | Check&#39;s id | [optional] 

## Methods

### NewChecksReplayRequest

`func NewChecksReplayRequest(orgId string, accountId string, ) *ChecksReplayRequest`

NewChecksReplayRequest instantiates a new ChecksReplayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksReplayRequestWithDefaults

`func NewChecksReplayRequestWithDefaults() *ChecksReplayRequest`

NewChecksReplayRequestWithDefaults instantiates a new ChecksReplayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ChecksReplayRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ChecksReplayRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ChecksReplayRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *ChecksReplayRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ChecksReplayRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ChecksReplayRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreationStartTime

`func (o *ChecksReplayRequest) GetCreationStartTime() time.Time`

GetCreationStartTime returns the CreationStartTime field if non-nil, zero value otherwise.

### GetCreationStartTimeOk

`func (o *ChecksReplayRequest) GetCreationStartTimeOk() (*time.Time, bool)`

GetCreationStartTimeOk returns a tuple with the CreationStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStartTime

`func (o *ChecksReplayRequest) SetCreationStartTime(v time.Time)`

SetCreationStartTime sets CreationStartTime field to given value.

### HasCreationStartTime

`func (o *ChecksReplayRequest) HasCreationStartTime() bool`

HasCreationStartTime returns a boolean if a field has been set.

### GetCreationEndTime

`func (o *ChecksReplayRequest) GetCreationEndTime() time.Time`

GetCreationEndTime returns the CreationEndTime field if non-nil, zero value otherwise.

### GetCreationEndTimeOk

`func (o *ChecksReplayRequest) GetCreationEndTimeOk() (*time.Time, bool)`

GetCreationEndTimeOk returns a tuple with the CreationEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationEndTime

`func (o *ChecksReplayRequest) SetCreationEndTime(v time.Time)`

SetCreationEndTime sets CreationEndTime field to given value.

### HasCreationEndTime

`func (o *ChecksReplayRequest) HasCreationEndTime() bool`

HasCreationEndTime returns a boolean if a field has been set.

### GetRate

`func (o *ChecksReplayRequest) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ChecksReplayRequest) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ChecksReplayRequest) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ChecksReplayRequest) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetCheckIds

`func (o *ChecksReplayRequest) GetCheckIds() []float32`

GetCheckIds returns the CheckIds field if non-nil, zero value otherwise.

### GetCheckIdsOk

`func (o *ChecksReplayRequest) GetCheckIdsOk() (*[]float32, bool)`

GetCheckIdsOk returns a tuple with the CheckIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIds

`func (o *ChecksReplayRequest) SetCheckIds(v []float32)`

SetCheckIds sets CheckIds field to given value.

### HasCheckIds

`func (o *ChecksReplayRequest) HasCheckIds() bool`

HasCheckIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


