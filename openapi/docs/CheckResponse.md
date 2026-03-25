# CheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the handle | 
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**CheckId** | **float32** | Check&#39;s id | 
**CheckCollectionId** | **float32** | ID of check collection if started in scope of a collection | 
**CheckTemplateId** | **float32** | Check Template for this check | 
**Name** | **string** | Name of this check template | 
**Note** | **string** | Note to be shown to the customer | 
**DurationMs** | **float32** | Duration of the check in milliseconds | 
**CreatedDate** | **time.Time** | Creation date | 
**ModifiedDate** | **time.Time** | Modify date | 
**OrgId** | **string** | Org id | 
**AccountId** | **string** | Account id | 
**InternalNote** | **string** | Note which is shown only internally to the agent | 
**Log** | **string** | Detailed log of the check execution | 

## Methods

### NewCheckResponse

`func NewCheckResponse(status string, objectType string, objectId string, checkId float32, checkCollectionId float32, checkTemplateId float32, name string, note string, durationMs float32, createdDate time.Time, modifiedDate time.Time, orgId string, accountId string, internalNote string, log string, ) *CheckResponse`

NewCheckResponse instantiates a new CheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckResponseWithDefaults

`func NewCheckResponseWithDefaults() *CheckResponse`

NewCheckResponseWithDefaults instantiates a new CheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CheckResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObjectType

`func (o *CheckResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CheckResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CheckResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *CheckResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CheckResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CheckResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetCheckId

`func (o *CheckResponse) GetCheckId() float32`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *CheckResponse) GetCheckIdOk() (*float32, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *CheckResponse) SetCheckId(v float32)`

SetCheckId sets CheckId field to given value.


### GetCheckCollectionId

`func (o *CheckResponse) GetCheckCollectionId() float32`

GetCheckCollectionId returns the CheckCollectionId field if non-nil, zero value otherwise.

### GetCheckCollectionIdOk

`func (o *CheckResponse) GetCheckCollectionIdOk() (*float32, bool)`

GetCheckCollectionIdOk returns a tuple with the CheckCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionId

`func (o *CheckResponse) SetCheckCollectionId(v float32)`

SetCheckCollectionId sets CheckCollectionId field to given value.


### GetCheckTemplateId

`func (o *CheckResponse) GetCheckTemplateId() float32`

GetCheckTemplateId returns the CheckTemplateId field if non-nil, zero value otherwise.

### GetCheckTemplateIdOk

`func (o *CheckResponse) GetCheckTemplateIdOk() (*float32, bool)`

GetCheckTemplateIdOk returns a tuple with the CheckTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplateId

`func (o *CheckResponse) SetCheckTemplateId(v float32)`

SetCheckTemplateId sets CheckTemplateId field to given value.


### GetName

`func (o *CheckResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *CheckResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CheckResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CheckResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetDurationMs

`func (o *CheckResponse) GetDurationMs() float32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *CheckResponse) GetDurationMsOk() (*float32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *CheckResponse) SetDurationMs(v float32)`

SetDurationMs sets DurationMs field to given value.


### GetCreatedDate

`func (o *CheckResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CheckResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CheckResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *CheckResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *CheckResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *CheckResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetOrgId

`func (o *CheckResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CheckResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CheckResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *CheckResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetInternalNote

`func (o *CheckResponse) GetInternalNote() string`

GetInternalNote returns the InternalNote field if non-nil, zero value otherwise.

### GetInternalNoteOk

`func (o *CheckResponse) GetInternalNoteOk() (*string, bool)`

GetInternalNoteOk returns a tuple with the InternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNote

`func (o *CheckResponse) SetInternalNote(v string)`

SetInternalNote sets InternalNote field to given value.


### GetLog

`func (o *CheckResponse) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *CheckResponse) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *CheckResponse) SetLog(v string)`

SetLog sets Log field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


