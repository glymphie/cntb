# ChangedAuthCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changed** | Pointer to **bool** | Flag that indicates if the auth code got changed | [optional] 
**Date** | Pointer to **time.Time** | The date when auth code got changed | [optional] 

## Methods

### NewChangedAuthCode

`func NewChangedAuthCode() *ChangedAuthCode`

NewChangedAuthCode instantiates a new ChangedAuthCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangedAuthCodeWithDefaults

`func NewChangedAuthCodeWithDefaults() *ChangedAuthCode`

NewChangedAuthCodeWithDefaults instantiates a new ChangedAuthCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanged

`func (o *ChangedAuthCode) GetChanged() bool`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *ChangedAuthCode) GetChangedOk() (*bool, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *ChangedAuthCode) SetChanged(v bool)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *ChangedAuthCode) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetDate

`func (o *ChangedAuthCode) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ChangedAuthCode) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ChangedAuthCode) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ChangedAuthCode) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


