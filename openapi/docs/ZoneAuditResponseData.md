# ZoneAuditResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the audit entry. | 
**Action** | **string** | Type of the action. | 
**Timestamp** | **time.Time** | When the change took place. | 
**TenantId** | **string** | Customer tenant id | 
**CustomerId** | **string** | Customer ID | 
**ChangedBy** | **string** | Id of user who performed the change | 
**Username** | **string** | Name of the user which led to the change. | 
**RequestId** | **string** | The requestId of the API call which led to the change. | 
**TraceId** | **string** | The traceId of the API call which led to the change. | 
**Name** | **string** | DNS Zone name | 
**Changes** | Pointer to **map[string]interface{}** | List of actual changes. | [optional] 

## Methods

### NewZoneAuditResponseData

`func NewZoneAuditResponseData(id int64, action string, timestamp time.Time, tenantId string, customerId string, changedBy string, username string, requestId string, traceId string, name string, ) *ZoneAuditResponseData`

NewZoneAuditResponseData instantiates a new ZoneAuditResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuditResponseDataWithDefaults

`func NewZoneAuditResponseDataWithDefaults() *ZoneAuditResponseData`

NewZoneAuditResponseDataWithDefaults instantiates a new ZoneAuditResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneAuditResponseData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneAuditResponseData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneAuditResponseData) SetId(v int64)`

SetId sets Id field to given value.


### GetAction

`func (o *ZoneAuditResponseData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ZoneAuditResponseData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ZoneAuditResponseData) SetAction(v string)`

SetAction sets Action field to given value.


### GetTimestamp

`func (o *ZoneAuditResponseData) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ZoneAuditResponseData) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ZoneAuditResponseData) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTenantId

`func (o *ZoneAuditResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ZoneAuditResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ZoneAuditResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ZoneAuditResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ZoneAuditResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ZoneAuditResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetChangedBy

`func (o *ZoneAuditResponseData) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *ZoneAuditResponseData) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *ZoneAuditResponseData) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.


### GetUsername

`func (o *ZoneAuditResponseData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ZoneAuditResponseData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ZoneAuditResponseData) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRequestId

`func (o *ZoneAuditResponseData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ZoneAuditResponseData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ZoneAuditResponseData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTraceId

`func (o *ZoneAuditResponseData) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ZoneAuditResponseData) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ZoneAuditResponseData) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetName

`func (o *ZoneAuditResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneAuditResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneAuditResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetChanges

`func (o *ZoneAuditResponseData) GetChanges() map[string]interface{}`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ZoneAuditResponseData) GetChangesOk() (*map[string]interface{}, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ZoneAuditResponseData) SetChanges(v map[string]interface{})`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *ZoneAuditResponseData) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


