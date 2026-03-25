# RemedyTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **time.Time** | Creation date | 
**ModifiedDate** | **time.Time** | Modify date | 
**OrgId** | **string** | Org id | 
**AccountId** | **string** | Account id | 
**RemedyTemplateId** | **float32** | Remedy template&#39;s id | 
**Name** | **string** | Name of the remedy template | 
**Description** | **string** | Description for the remedy template | 
**Internal** | **bool** | Is remedy only internal (not shown to the customer) | 
**ObjectType** | **string** | Object type for which the remedy template can be used | 
**CollectorClass** | **string** | Class used to collect the required information for the remedy | 
**RemedyClass** | **string** | Class used to perform the remedy | 
**Requirements** | **map[string]interface{}** | Requirements for remedy (reboot, reinstall, ...) | 
**CheckTemplateIds** | **[]string** | Check Template IDs that are related to this remedy | 

## Methods

### NewRemedyTemplateResponse

`func NewRemedyTemplateResponse(createdDate time.Time, modifiedDate time.Time, orgId string, accountId string, remedyTemplateId float32, name string, description string, internal bool, objectType string, collectorClass string, remedyClass string, requirements map[string]interface{}, checkTemplateIds []string, ) *RemedyTemplateResponse`

NewRemedyTemplateResponse instantiates a new RemedyTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemedyTemplateResponseWithDefaults

`func NewRemedyTemplateResponseWithDefaults() *RemedyTemplateResponse`

NewRemedyTemplateResponseWithDefaults instantiates a new RemedyTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *RemedyTemplateResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *RemedyTemplateResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *RemedyTemplateResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *RemedyTemplateResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *RemedyTemplateResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *RemedyTemplateResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetOrgId

`func (o *RemedyTemplateResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RemedyTemplateResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RemedyTemplateResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *RemedyTemplateResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RemedyTemplateResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RemedyTemplateResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetRemedyTemplateId

`func (o *RemedyTemplateResponse) GetRemedyTemplateId() float32`

GetRemedyTemplateId returns the RemedyTemplateId field if non-nil, zero value otherwise.

### GetRemedyTemplateIdOk

`func (o *RemedyTemplateResponse) GetRemedyTemplateIdOk() (*float32, bool)`

GetRemedyTemplateIdOk returns a tuple with the RemedyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyTemplateId

`func (o *RemedyTemplateResponse) SetRemedyTemplateId(v float32)`

SetRemedyTemplateId sets RemedyTemplateId field to given value.


### GetName

`func (o *RemedyTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemedyTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemedyTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RemedyTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemedyTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemedyTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInternal

`func (o *RemedyTemplateResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *RemedyTemplateResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *RemedyTemplateResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetObjectType

`func (o *RemedyTemplateResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RemedyTemplateResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RemedyTemplateResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectorClass

`func (o *RemedyTemplateResponse) GetCollectorClass() string`

GetCollectorClass returns the CollectorClass field if non-nil, zero value otherwise.

### GetCollectorClassOk

`func (o *RemedyTemplateResponse) GetCollectorClassOk() (*string, bool)`

GetCollectorClassOk returns a tuple with the CollectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorClass

`func (o *RemedyTemplateResponse) SetCollectorClass(v string)`

SetCollectorClass sets CollectorClass field to given value.


### GetRemedyClass

`func (o *RemedyTemplateResponse) GetRemedyClass() string`

GetRemedyClass returns the RemedyClass field if non-nil, zero value otherwise.

### GetRemedyClassOk

`func (o *RemedyTemplateResponse) GetRemedyClassOk() (*string, bool)`

GetRemedyClassOk returns a tuple with the RemedyClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyClass

`func (o *RemedyTemplateResponse) SetRemedyClass(v string)`

SetRemedyClass sets RemedyClass field to given value.


### GetRequirements

`func (o *RemedyTemplateResponse) GetRequirements() map[string]interface{}`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *RemedyTemplateResponse) GetRequirementsOk() (*map[string]interface{}, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *RemedyTemplateResponse) SetRequirements(v map[string]interface{})`

SetRequirements sets Requirements field to given value.


### GetCheckTemplateIds

`func (o *RemedyTemplateResponse) GetCheckTemplateIds() []string`

GetCheckTemplateIds returns the CheckTemplateIds field if non-nil, zero value otherwise.

### GetCheckTemplateIdsOk

`func (o *RemedyTemplateResponse) GetCheckTemplateIdsOk() (*[]string, bool)`

GetCheckTemplateIdsOk returns a tuple with the CheckTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplateIds

`func (o *RemedyTemplateResponse) SetCheckTemplateIds(v []string)`

SetCheckTemplateIds sets CheckTemplateIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


