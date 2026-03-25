# BaseRemedyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**RemedyTemplateId** | **float32** | Remedy Template for this remedy | 
**RemedyCollectionId** | Pointer to **float32** | Remedy Collection for this remedy | [optional] 

## Methods

### NewBaseRemedyCreateRequest

`func NewBaseRemedyCreateRequest(objectType string, objectId string, remedyTemplateId float32, ) *BaseRemedyCreateRequest`

NewBaseRemedyCreateRequest instantiates a new BaseRemedyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseRemedyCreateRequestWithDefaults

`func NewBaseRemedyCreateRequestWithDefaults() *BaseRemedyCreateRequest`

NewBaseRemedyCreateRequestWithDefaults instantiates a new BaseRemedyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *BaseRemedyCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BaseRemedyCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BaseRemedyCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *BaseRemedyCreateRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BaseRemedyCreateRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BaseRemedyCreateRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetRemedyTemplateId

`func (o *BaseRemedyCreateRequest) GetRemedyTemplateId() float32`

GetRemedyTemplateId returns the RemedyTemplateId field if non-nil, zero value otherwise.

### GetRemedyTemplateIdOk

`func (o *BaseRemedyCreateRequest) GetRemedyTemplateIdOk() (*float32, bool)`

GetRemedyTemplateIdOk returns a tuple with the RemedyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyTemplateId

`func (o *BaseRemedyCreateRequest) SetRemedyTemplateId(v float32)`

SetRemedyTemplateId sets RemedyTemplateId field to given value.


### GetRemedyCollectionId

`func (o *BaseRemedyCreateRequest) GetRemedyCollectionId() float32`

GetRemedyCollectionId returns the RemedyCollectionId field if non-nil, zero value otherwise.

### GetRemedyCollectionIdOk

`func (o *BaseRemedyCreateRequest) GetRemedyCollectionIdOk() (*float32, bool)`

GetRemedyCollectionIdOk returns a tuple with the RemedyCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyCollectionId

`func (o *BaseRemedyCreateRequest) SetRemedyCollectionId(v float32)`

SetRemedyCollectionId sets RemedyCollectionId field to given value.

### HasRemedyCollectionId

`func (o *BaseRemedyCreateRequest) HasRemedyCollectionId() bool`

HasRemedyCollectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


