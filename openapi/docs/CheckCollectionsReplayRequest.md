# CheckCollectionsReplayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | Id of your organization, if unknown please contact us | 
**AccountId** | **string** | Account Id | 
**CreationStartTime** | Pointer to **time.Time** | Earliest creation date of changes to replay | [optional] 
**CreationEndTime** | Pointer to **time.Time** | Latest creation date of changes to replay | [optional] 
**Rate** | Pointer to **float32** | Message publishing frequency. How many messages per second get published. Default: 20 | [optional] 
**CheckCollectionIds** | Pointer to **[]float32** | Check collection&#39;s id | [optional] 

## Methods

### NewCheckCollectionsReplayRequest

`func NewCheckCollectionsReplayRequest(orgId string, accountId string, ) *CheckCollectionsReplayRequest`

NewCheckCollectionsReplayRequest instantiates a new CheckCollectionsReplayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCollectionsReplayRequestWithDefaults

`func NewCheckCollectionsReplayRequestWithDefaults() *CheckCollectionsReplayRequest`

NewCheckCollectionsReplayRequestWithDefaults instantiates a new CheckCollectionsReplayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *CheckCollectionsReplayRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CheckCollectionsReplayRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CheckCollectionsReplayRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *CheckCollectionsReplayRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckCollectionsReplayRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckCollectionsReplayRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreationStartTime

`func (o *CheckCollectionsReplayRequest) GetCreationStartTime() time.Time`

GetCreationStartTime returns the CreationStartTime field if non-nil, zero value otherwise.

### GetCreationStartTimeOk

`func (o *CheckCollectionsReplayRequest) GetCreationStartTimeOk() (*time.Time, bool)`

GetCreationStartTimeOk returns a tuple with the CreationStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStartTime

`func (o *CheckCollectionsReplayRequest) SetCreationStartTime(v time.Time)`

SetCreationStartTime sets CreationStartTime field to given value.

### HasCreationStartTime

`func (o *CheckCollectionsReplayRequest) HasCreationStartTime() bool`

HasCreationStartTime returns a boolean if a field has been set.

### GetCreationEndTime

`func (o *CheckCollectionsReplayRequest) GetCreationEndTime() time.Time`

GetCreationEndTime returns the CreationEndTime field if non-nil, zero value otherwise.

### GetCreationEndTimeOk

`func (o *CheckCollectionsReplayRequest) GetCreationEndTimeOk() (*time.Time, bool)`

GetCreationEndTimeOk returns a tuple with the CreationEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationEndTime

`func (o *CheckCollectionsReplayRequest) SetCreationEndTime(v time.Time)`

SetCreationEndTime sets CreationEndTime field to given value.

### HasCreationEndTime

`func (o *CheckCollectionsReplayRequest) HasCreationEndTime() bool`

HasCreationEndTime returns a boolean if a field has been set.

### GetRate

`func (o *CheckCollectionsReplayRequest) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CheckCollectionsReplayRequest) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CheckCollectionsReplayRequest) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *CheckCollectionsReplayRequest) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetCheckCollectionIds

`func (o *CheckCollectionsReplayRequest) GetCheckCollectionIds() []float32`

GetCheckCollectionIds returns the CheckCollectionIds field if non-nil, zero value otherwise.

### GetCheckCollectionIdsOk

`func (o *CheckCollectionsReplayRequest) GetCheckCollectionIdsOk() (*[]float32, bool)`

GetCheckCollectionIdsOk returns a tuple with the CheckCollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionIds

`func (o *CheckCollectionsReplayRequest) SetCheckCollectionIds(v []float32)`

SetCheckCollectionIds sets CheckCollectionIds field to given value.

### HasCheckCollectionIds

`func (o *CheckCollectionsReplayRequest) HasCheckCollectionIds() bool`

HasCheckCollectionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


