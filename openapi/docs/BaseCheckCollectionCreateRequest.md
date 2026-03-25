# BaseCheckCollectionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**CheckCollectionTemplateId** | **float32** | Check Template for this check collection | 

## Methods

### NewBaseCheckCollectionCreateRequest

`func NewBaseCheckCollectionCreateRequest(objectType string, objectId string, checkCollectionTemplateId float32, ) *BaseCheckCollectionCreateRequest`

NewBaseCheckCollectionCreateRequest instantiates a new BaseCheckCollectionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCheckCollectionCreateRequestWithDefaults

`func NewBaseCheckCollectionCreateRequestWithDefaults() *BaseCheckCollectionCreateRequest`

NewBaseCheckCollectionCreateRequestWithDefaults instantiates a new BaseCheckCollectionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *BaseCheckCollectionCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BaseCheckCollectionCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BaseCheckCollectionCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *BaseCheckCollectionCreateRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BaseCheckCollectionCreateRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BaseCheckCollectionCreateRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetCheckCollectionTemplateId

`func (o *BaseCheckCollectionCreateRequest) GetCheckCollectionTemplateId() float32`

GetCheckCollectionTemplateId returns the CheckCollectionTemplateId field if non-nil, zero value otherwise.

### GetCheckCollectionTemplateIdOk

`func (o *BaseCheckCollectionCreateRequest) GetCheckCollectionTemplateIdOk() (*float32, bool)`

GetCheckCollectionTemplateIdOk returns a tuple with the CheckCollectionTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionTemplateId

`func (o *BaseCheckCollectionCreateRequest) SetCheckCollectionTemplateId(v float32)`

SetCheckCollectionTemplateId sets CheckCollectionTemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


