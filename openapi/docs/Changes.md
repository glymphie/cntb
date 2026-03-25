# Changes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prev** | **map[string]interface{}** | Previous values of changed properties | 
**New** | **map[string]interface{}** | New values of changed properties | 

## Methods

### NewChanges

`func NewChanges(prev map[string]interface{}, new map[string]interface{}, ) *Changes`

NewChanges instantiates a new Changes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangesWithDefaults

`func NewChangesWithDefaults() *Changes`

NewChangesWithDefaults instantiates a new Changes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrev

`func (o *Changes) GetPrev() map[string]interface{}`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *Changes) GetPrevOk() (*map[string]interface{}, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *Changes) SetPrev(v map[string]interface{})`

SetPrev sets Prev field to given value.


### GetNew

`func (o *Changes) GetNew() map[string]interface{}`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *Changes) GetNewOk() (*map[string]interface{}, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *Changes) SetNew(v map[string]interface{})`

SetNew sets New field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


