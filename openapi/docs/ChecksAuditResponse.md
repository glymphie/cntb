# ChecksAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **time.Time** | Creation date | 
**ModifiedDate** | **time.Time** | Modify date | 
**OrgId** | **string** | Org id | 
**AccountId** | **string** | Account id | 
**AuditId** | **int64** | The ID of the audit entry. | 
**Action** | **string** | Type of the action. | 
**ForeignChangedBy** | **string** | Id of a foreign user (given on the api request via header) who performed the change | 
**ForeignUsername** | **string** | Name of the foreign user (given on the api request via header) which led to the change. | 
**ChangedBy** | **string** | Id of user who performed the change | 
**Username** | **string** | Name of the user which led to the change. | 
**RequestId** | **string** | The requestId of the API call which led to the change. | 
**TraceId** | **string** | The traceId of the API call which led to the change. | 
**Changes** | [**Changes**](Changes.md) | List of changed properties | 
**CheckId** | **float32** | Check&#39;s id | 

## Methods

### NewChecksAuditResponse

`func NewChecksAuditResponse(createdDate time.Time, modifiedDate time.Time, orgId string, accountId string, auditId int64, action string, foreignChangedBy string, foreignUsername string, changedBy string, username string, requestId string, traceId string, changes Changes, checkId float32, ) *ChecksAuditResponse`

NewChecksAuditResponse instantiates a new ChecksAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksAuditResponseWithDefaults

`func NewChecksAuditResponseWithDefaults() *ChecksAuditResponse`

NewChecksAuditResponseWithDefaults instantiates a new ChecksAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *ChecksAuditResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ChecksAuditResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ChecksAuditResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ChecksAuditResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ChecksAuditResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ChecksAuditResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetOrgId

`func (o *ChecksAuditResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ChecksAuditResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ChecksAuditResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *ChecksAuditResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ChecksAuditResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ChecksAuditResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAuditId

`func (o *ChecksAuditResponse) GetAuditId() int64`

GetAuditId returns the AuditId field if non-nil, zero value otherwise.

### GetAuditIdOk

`func (o *ChecksAuditResponse) GetAuditIdOk() (*int64, bool)`

GetAuditIdOk returns a tuple with the AuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditId

`func (o *ChecksAuditResponse) SetAuditId(v int64)`

SetAuditId sets AuditId field to given value.


### GetAction

`func (o *ChecksAuditResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ChecksAuditResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ChecksAuditResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetForeignChangedBy

`func (o *ChecksAuditResponse) GetForeignChangedBy() string`

GetForeignChangedBy returns the ForeignChangedBy field if non-nil, zero value otherwise.

### GetForeignChangedByOk

`func (o *ChecksAuditResponse) GetForeignChangedByOk() (*string, bool)`

GetForeignChangedByOk returns a tuple with the ForeignChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignChangedBy

`func (o *ChecksAuditResponse) SetForeignChangedBy(v string)`

SetForeignChangedBy sets ForeignChangedBy field to given value.


### GetForeignUsername

`func (o *ChecksAuditResponse) GetForeignUsername() string`

GetForeignUsername returns the ForeignUsername field if non-nil, zero value otherwise.

### GetForeignUsernameOk

`func (o *ChecksAuditResponse) GetForeignUsernameOk() (*string, bool)`

GetForeignUsernameOk returns a tuple with the ForeignUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignUsername

`func (o *ChecksAuditResponse) SetForeignUsername(v string)`

SetForeignUsername sets ForeignUsername field to given value.


### GetChangedBy

`func (o *ChecksAuditResponse) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *ChecksAuditResponse) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *ChecksAuditResponse) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.


### GetUsername

`func (o *ChecksAuditResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ChecksAuditResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ChecksAuditResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRequestId

`func (o *ChecksAuditResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ChecksAuditResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ChecksAuditResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTraceId

`func (o *ChecksAuditResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ChecksAuditResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ChecksAuditResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetChanges

`func (o *ChecksAuditResponse) GetChanges() Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ChecksAuditResponse) GetChangesOk() (*Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ChecksAuditResponse) SetChanges(v Changes)`

SetChanges sets Changes field to given value.


### GetCheckId

`func (o *ChecksAuditResponse) GetCheckId() float32`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *ChecksAuditResponse) GetCheckIdOk() (*float32, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *ChecksAuditResponse) SetCheckId(v float32)`

SetCheckId sets CheckId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


