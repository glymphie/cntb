# RemediesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**RemedyTemplateId** | **float32** | Remedy Template for this remedy | 
**RemedyCollectionId** | Pointer to **float32** | Remedy Collection for this remedy | [optional] 
**OrgId** | **string** | Id of your organization, if unknown please contact us | 
**AccountId** | **string** | Account Id | 

## Methods

### NewRemediesCreateRequest

`func NewRemediesCreateRequest(objectType string, objectId string, remedyTemplateId float32, orgId string, accountId string, ) *RemediesCreateRequest`

NewRemediesCreateRequest instantiates a new RemediesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediesCreateRequestWithDefaults

`func NewRemediesCreateRequestWithDefaults() *RemediesCreateRequest`

NewRemediesCreateRequestWithDefaults instantiates a new RemediesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *RemediesCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RemediesCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RemediesCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *RemediesCreateRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RemediesCreateRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RemediesCreateRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetRemedyTemplateId

`func (o *RemediesCreateRequest) GetRemedyTemplateId() float32`

GetRemedyTemplateId returns the RemedyTemplateId field if non-nil, zero value otherwise.

### GetRemedyTemplateIdOk

`func (o *RemediesCreateRequest) GetRemedyTemplateIdOk() (*float32, bool)`

GetRemedyTemplateIdOk returns a tuple with the RemedyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyTemplateId

`func (o *RemediesCreateRequest) SetRemedyTemplateId(v float32)`

SetRemedyTemplateId sets RemedyTemplateId field to given value.


### GetRemedyCollectionId

`func (o *RemediesCreateRequest) GetRemedyCollectionId() float32`

GetRemedyCollectionId returns the RemedyCollectionId field if non-nil, zero value otherwise.

### GetRemedyCollectionIdOk

`func (o *RemediesCreateRequest) GetRemedyCollectionIdOk() (*float32, bool)`

GetRemedyCollectionIdOk returns a tuple with the RemedyCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyCollectionId

`func (o *RemediesCreateRequest) SetRemedyCollectionId(v float32)`

SetRemedyCollectionId sets RemedyCollectionId field to given value.

### HasRemedyCollectionId

`func (o *RemediesCreateRequest) HasRemedyCollectionId() bool`

HasRemedyCollectionId returns a boolean if a field has been set.

### GetOrgId

`func (o *RemediesCreateRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RemediesCreateRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RemediesCreateRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *RemediesCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RemediesCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RemediesCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


