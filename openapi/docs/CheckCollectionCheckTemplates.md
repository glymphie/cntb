# CheckCollectionCheckTemplates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckId** | **float32** | Id of the check | 
**CheckTemplateId** | **float32** | Id of the check template | 
**RunConcurrent** | **bool** | Can this check template be run in parallel with other checks | 
**IgnoreErrors** | **bool** | Will errors be ignored when running this check template | 
**CheckTemplates** | Pointer to [**[]CheckCollectionCheckTemplates**](CheckCollectionCheckTemplates.md) | Nested check templates | [optional] 

## Methods

### NewCheckCollectionCheckTemplates

`func NewCheckCollectionCheckTemplates(checkId float32, checkTemplateId float32, runConcurrent bool, ignoreErrors bool, ) *CheckCollectionCheckTemplates`

NewCheckCollectionCheckTemplates instantiates a new CheckCollectionCheckTemplates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCollectionCheckTemplatesWithDefaults

`func NewCheckCollectionCheckTemplatesWithDefaults() *CheckCollectionCheckTemplates`

NewCheckCollectionCheckTemplatesWithDefaults instantiates a new CheckCollectionCheckTemplates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckId

`func (o *CheckCollectionCheckTemplates) GetCheckId() float32`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *CheckCollectionCheckTemplates) GetCheckIdOk() (*float32, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *CheckCollectionCheckTemplates) SetCheckId(v float32)`

SetCheckId sets CheckId field to given value.


### GetCheckTemplateId

`func (o *CheckCollectionCheckTemplates) GetCheckTemplateId() float32`

GetCheckTemplateId returns the CheckTemplateId field if non-nil, zero value otherwise.

### GetCheckTemplateIdOk

`func (o *CheckCollectionCheckTemplates) GetCheckTemplateIdOk() (*float32, bool)`

GetCheckTemplateIdOk returns a tuple with the CheckTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplateId

`func (o *CheckCollectionCheckTemplates) SetCheckTemplateId(v float32)`

SetCheckTemplateId sets CheckTemplateId field to given value.


### GetRunConcurrent

`func (o *CheckCollectionCheckTemplates) GetRunConcurrent() bool`

GetRunConcurrent returns the RunConcurrent field if non-nil, zero value otherwise.

### GetRunConcurrentOk

`func (o *CheckCollectionCheckTemplates) GetRunConcurrentOk() (*bool, bool)`

GetRunConcurrentOk returns a tuple with the RunConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunConcurrent

`func (o *CheckCollectionCheckTemplates) SetRunConcurrent(v bool)`

SetRunConcurrent sets RunConcurrent field to given value.


### GetIgnoreErrors

`func (o *CheckCollectionCheckTemplates) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *CheckCollectionCheckTemplates) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *CheckCollectionCheckTemplates) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.


### GetCheckTemplates

`func (o *CheckCollectionCheckTemplates) GetCheckTemplates() []CheckCollectionCheckTemplates`

GetCheckTemplates returns the CheckTemplates field if non-nil, zero value otherwise.

### GetCheckTemplatesOk

`func (o *CheckCollectionCheckTemplates) GetCheckTemplatesOk() (*[]CheckCollectionCheckTemplates, bool)`

GetCheckTemplatesOk returns a tuple with the CheckTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplates

`func (o *CheckCollectionCheckTemplates) SetCheckTemplates(v []CheckCollectionCheckTemplates)`

SetCheckTemplates sets CheckTemplates field to given value.

### HasCheckTemplates

`func (o *CheckCollectionCheckTemplates) HasCheckTemplates() bool`

HasCheckTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


