# CheckCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**CheckTemplateId** | **float32** | Check Template for this check | 
**CheckCollectionId** | Pointer to **float32** | Check Collection for this check | [optional] 
**OrgId** | **string** | Id of your organization, if unknown please contact us | 
**AccountId** | **string** | Account Id | 

## Methods

### NewCheckCreateRequest

`func NewCheckCreateRequest(objectType string, objectId string, checkTemplateId float32, orgId string, accountId string, ) *CheckCreateRequest`

NewCheckCreateRequest instantiates a new CheckCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCreateRequestWithDefaults

`func NewCheckCreateRequestWithDefaults() *CheckCreateRequest`

NewCheckCreateRequestWithDefaults instantiates a new CheckCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *CheckCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CheckCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CheckCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *CheckCreateRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CheckCreateRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CheckCreateRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetCheckTemplateId

`func (o *CheckCreateRequest) GetCheckTemplateId() float32`

GetCheckTemplateId returns the CheckTemplateId field if non-nil, zero value otherwise.

### GetCheckTemplateIdOk

`func (o *CheckCreateRequest) GetCheckTemplateIdOk() (*float32, bool)`

GetCheckTemplateIdOk returns a tuple with the CheckTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplateId

`func (o *CheckCreateRequest) SetCheckTemplateId(v float32)`

SetCheckTemplateId sets CheckTemplateId field to given value.


### GetCheckCollectionId

`func (o *CheckCreateRequest) GetCheckCollectionId() float32`

GetCheckCollectionId returns the CheckCollectionId field if non-nil, zero value otherwise.

### GetCheckCollectionIdOk

`func (o *CheckCreateRequest) GetCheckCollectionIdOk() (*float32, bool)`

GetCheckCollectionIdOk returns a tuple with the CheckCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionId

`func (o *CheckCreateRequest) SetCheckCollectionId(v float32)`

SetCheckCollectionId sets CheckCollectionId field to given value.

### HasCheckCollectionId

`func (o *CheckCreateRequest) HasCheckCollectionId() bool`

HasCheckCollectionId returns a boolean if a field has been set.

### GetOrgId

`func (o *CheckCreateRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CheckCreateRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CheckCreateRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *CheckCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


