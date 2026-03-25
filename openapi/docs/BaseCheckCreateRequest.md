# BaseCheckCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**CheckTemplateId** | **float32** | Check Template for this check | 
**CheckCollectionId** | Pointer to **float32** | Check Collection for this check | [optional] 

## Methods

### NewBaseCheckCreateRequest

`func NewBaseCheckCreateRequest(objectType string, objectId string, checkTemplateId float32, ) *BaseCheckCreateRequest`

NewBaseCheckCreateRequest instantiates a new BaseCheckCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCheckCreateRequestWithDefaults

`func NewBaseCheckCreateRequestWithDefaults() *BaseCheckCreateRequest`

NewBaseCheckCreateRequestWithDefaults instantiates a new BaseCheckCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *BaseCheckCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BaseCheckCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BaseCheckCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *BaseCheckCreateRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BaseCheckCreateRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BaseCheckCreateRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetCheckTemplateId

`func (o *BaseCheckCreateRequest) GetCheckTemplateId() float32`

GetCheckTemplateId returns the CheckTemplateId field if non-nil, zero value otherwise.

### GetCheckTemplateIdOk

`func (o *BaseCheckCreateRequest) GetCheckTemplateIdOk() (*float32, bool)`

GetCheckTemplateIdOk returns a tuple with the CheckTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplateId

`func (o *BaseCheckCreateRequest) SetCheckTemplateId(v float32)`

SetCheckTemplateId sets CheckTemplateId field to given value.


### GetCheckCollectionId

`func (o *BaseCheckCreateRequest) GetCheckCollectionId() float32`

GetCheckCollectionId returns the CheckCollectionId field if non-nil, zero value otherwise.

### GetCheckCollectionIdOk

`func (o *BaseCheckCreateRequest) GetCheckCollectionIdOk() (*float32, bool)`

GetCheckCollectionIdOk returns a tuple with the CheckCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionId

`func (o *BaseCheckCreateRequest) SetCheckCollectionId(v float32)`

SetCheckCollectionId sets CheckCollectionId field to given value.

### HasCheckCollectionId

`func (o *BaseCheckCreateRequest) HasCheckCollectionId() bool`

HasCheckCollectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


