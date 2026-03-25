# DomainHandles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | Domain&#39;s owner handle | 
**Admin** | **string** | Domain&#39;s admin handle | 
**Tech** | **string** | Domain&#39;s tech handle | 
**Zone** | **string** | Domain&#39;s zone handle | 

## Methods

### NewDomainHandles

`func NewDomainHandles(owner string, admin string, tech string, zone string, ) *DomainHandles`

NewDomainHandles instantiates a new DomainHandles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainHandlesWithDefaults

`func NewDomainHandlesWithDefaults() *DomainHandles`

NewDomainHandlesWithDefaults instantiates a new DomainHandles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *DomainHandles) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DomainHandles) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DomainHandles) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetAdmin

`func (o *DomainHandles) GetAdmin() string`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *DomainHandles) GetAdminOk() (*string, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *DomainHandles) SetAdmin(v string)`

SetAdmin sets Admin field to given value.


### GetTech

`func (o *DomainHandles) GetTech() string`

GetTech returns the Tech field if non-nil, zero value otherwise.

### GetTechOk

`func (o *DomainHandles) GetTechOk() (*string, bool)`

GetTechOk returns a tuple with the Tech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTech

`func (o *DomainHandles) SetTech(v string)`

SetTech sets Tech field to given value.


### GetZone

`func (o *DomainHandles) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *DomainHandles) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *DomainHandles) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


