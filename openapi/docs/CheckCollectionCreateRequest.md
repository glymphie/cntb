# CheckCollectionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**CheckCollectionTemplateId** | **float32** | Check Template for this check collection | 
**OrgId** | **string** | Id of your organization, if unknown please contact us | 
**AccountId** | **string** | Account Id | 

## Methods

### NewCheckCollectionCreateRequest

`func NewCheckCollectionCreateRequest(objectType string, objectId string, checkCollectionTemplateId float32, orgId string, accountId string, ) *CheckCollectionCreateRequest`

NewCheckCollectionCreateRequest instantiates a new CheckCollectionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCollectionCreateRequestWithDefaults

`func NewCheckCollectionCreateRequestWithDefaults() *CheckCollectionCreateRequest`

NewCheckCollectionCreateRequestWithDefaults instantiates a new CheckCollectionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *CheckCollectionCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CheckCollectionCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CheckCollectionCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *CheckCollectionCreateRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CheckCollectionCreateRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CheckCollectionCreateRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetCheckCollectionTemplateId

`func (o *CheckCollectionCreateRequest) GetCheckCollectionTemplateId() float32`

GetCheckCollectionTemplateId returns the CheckCollectionTemplateId field if non-nil, zero value otherwise.

### GetCheckCollectionTemplateIdOk

`func (o *CheckCollectionCreateRequest) GetCheckCollectionTemplateIdOk() (*float32, bool)`

GetCheckCollectionTemplateIdOk returns a tuple with the CheckCollectionTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionTemplateId

`func (o *CheckCollectionCreateRequest) SetCheckCollectionTemplateId(v float32)`

SetCheckCollectionTemplateId sets CheckCollectionTemplateId field to given value.


### GetOrgId

`func (o *CheckCollectionCreateRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CheckCollectionCreateRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CheckCollectionCreateRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *CheckCollectionCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckCollectionCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckCollectionCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


