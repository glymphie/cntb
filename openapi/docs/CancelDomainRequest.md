# CancelDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason for cancelling an domain | [optional] 
**ReasonText** | Pointer to **string** | Reason Text when &#x60;Other&#x60; reason got selected while cancelling an domain | [optional] 
**CancelDate** | Pointer to **time.Time** | Date of cancellation | [optional] 

## Methods

### NewCancelDomainRequest

`func NewCancelDomainRequest() *CancelDomainRequest`

NewCancelDomainRequest instantiates a new CancelDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelDomainRequestWithDefaults

`func NewCancelDomainRequestWithDefaults() *CancelDomainRequest`

NewCancelDomainRequestWithDefaults instantiates a new CancelDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CancelDomainRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelDomainRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelDomainRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CancelDomainRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonText

`func (o *CancelDomainRequest) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *CancelDomainRequest) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *CancelDomainRequest) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *CancelDomainRequest) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### GetCancelDate

`func (o *CancelDomainRequest) GetCancelDate() time.Time`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *CancelDomainRequest) GetCancelDateOk() (*time.Time, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *CancelDomainRequest) SetCancelDate(v time.Time)`

SetCancelDate sets CancelDate field to given value.

### HasCancelDate

`func (o *CancelDomainRequest) HasCancelDate() bool`

HasCancelDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


