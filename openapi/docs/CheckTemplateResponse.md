# CheckTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **time.Time** | Creation date | 
**ModifiedDate** | **time.Time** | Modify date | 
**OrgId** | **string** | Org id | 
**AccountId** | **string** | Account id | 
**CheckTemplateId** | **float32** | Check template&#39;s id | 
**Name** | **string** | Name of the check template | 
**Description** | **string** | Description for the check template | 
**Internal** | **bool** | Is check only internal (not shown to the customer) | 
**ObjectType** | **string** | Object type for which the check template can be used | 
**CollectorClass** | **string** | Class used to collect the required information for the check | 
**CheckClass** | **string** | Class used to perform the check | 
**RemedyTemplateIds** | **[]string** | Remedy Template IDs that are related to this remedy | 

## Methods

### NewCheckTemplateResponse

`func NewCheckTemplateResponse(createdDate time.Time, modifiedDate time.Time, orgId string, accountId string, checkTemplateId float32, name string, description string, internal bool, objectType string, collectorClass string, checkClass string, remedyTemplateIds []string, ) *CheckTemplateResponse`

NewCheckTemplateResponse instantiates a new CheckTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTemplateResponseWithDefaults

`func NewCheckTemplateResponseWithDefaults() *CheckTemplateResponse`

NewCheckTemplateResponseWithDefaults instantiates a new CheckTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *CheckTemplateResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CheckTemplateResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CheckTemplateResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *CheckTemplateResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *CheckTemplateResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *CheckTemplateResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetOrgId

`func (o *CheckTemplateResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CheckTemplateResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CheckTemplateResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *CheckTemplateResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckTemplateResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckTemplateResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCheckTemplateId

`func (o *CheckTemplateResponse) GetCheckTemplateId() float32`

GetCheckTemplateId returns the CheckTemplateId field if non-nil, zero value otherwise.

### GetCheckTemplateIdOk

`func (o *CheckTemplateResponse) GetCheckTemplateIdOk() (*float32, bool)`

GetCheckTemplateIdOk returns a tuple with the CheckTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplateId

`func (o *CheckTemplateResponse) SetCheckTemplateId(v float32)`

SetCheckTemplateId sets CheckTemplateId field to given value.


### GetName

`func (o *CheckTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CheckTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInternal

`func (o *CheckTemplateResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CheckTemplateResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CheckTemplateResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetObjectType

`func (o *CheckTemplateResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CheckTemplateResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CheckTemplateResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectorClass

`func (o *CheckTemplateResponse) GetCollectorClass() string`

GetCollectorClass returns the CollectorClass field if non-nil, zero value otherwise.

### GetCollectorClassOk

`func (o *CheckTemplateResponse) GetCollectorClassOk() (*string, bool)`

GetCollectorClassOk returns a tuple with the CollectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorClass

`func (o *CheckTemplateResponse) SetCollectorClass(v string)`

SetCollectorClass sets CollectorClass field to given value.


### GetCheckClass

`func (o *CheckTemplateResponse) GetCheckClass() string`

GetCheckClass returns the CheckClass field if non-nil, zero value otherwise.

### GetCheckClassOk

`func (o *CheckTemplateResponse) GetCheckClassOk() (*string, bool)`

GetCheckClassOk returns a tuple with the CheckClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckClass

`func (o *CheckTemplateResponse) SetCheckClass(v string)`

SetCheckClass sets CheckClass field to given value.


### GetRemedyTemplateIds

`func (o *CheckTemplateResponse) GetRemedyTemplateIds() []string`

GetRemedyTemplateIds returns the RemedyTemplateIds field if non-nil, zero value otherwise.

### GetRemedyTemplateIdsOk

`func (o *CheckTemplateResponse) GetRemedyTemplateIdsOk() (*[]string, bool)`

GetRemedyTemplateIdsOk returns a tuple with the RemedyTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyTemplateIds

`func (o *CheckTemplateResponse) SetRemedyTemplateIds(v []string)`

SetRemedyTemplateIds sets RemedyTemplateIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


