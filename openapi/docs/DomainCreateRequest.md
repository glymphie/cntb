# DomainCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain name | 
**AuthCode** | Pointer to **string** | The domain auth code | [optional] 
**Handles** | [**DomainHandles**](DomainHandles.md) | The handles of the domain | 
**Nameservers** | [**[]Nameserver**](Nameserver.md) | Nameservers | 
**ResourceType** | Pointer to **string** | The identifier of the resource type | [optional] 
**ResourceId** | Pointer to **string** | The identifier of the resource id | [optional] 

## Methods

### NewDomainCreateRequest

`func NewDomainCreateRequest(domain string, handles DomainHandles, nameservers []Nameserver, ) *DomainCreateRequest`

NewDomainCreateRequest instantiates a new DomainCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCreateRequestWithDefaults

`func NewDomainCreateRequestWithDefaults() *DomainCreateRequest`

NewDomainCreateRequestWithDefaults instantiates a new DomainCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainCreateRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainCreateRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainCreateRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetAuthCode

`func (o *DomainCreateRequest) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *DomainCreateRequest) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *DomainCreateRequest) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *DomainCreateRequest) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetHandles

`func (o *DomainCreateRequest) GetHandles() DomainHandles`

GetHandles returns the Handles field if non-nil, zero value otherwise.

### GetHandlesOk

`func (o *DomainCreateRequest) GetHandlesOk() (*DomainHandles, bool)`

GetHandlesOk returns a tuple with the Handles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandles

`func (o *DomainCreateRequest) SetHandles(v DomainHandles)`

SetHandles sets Handles field to given value.


### GetNameservers

`func (o *DomainCreateRequest) GetNameservers() []Nameserver`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DomainCreateRequest) GetNameserversOk() (*[]Nameserver, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DomainCreateRequest) SetNameservers(v []Nameserver)`

SetNameservers sets Nameservers field to given value.


### GetResourceType

`func (o *DomainCreateRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DomainCreateRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DomainCreateRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DomainCreateRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceId

`func (o *DomainCreateRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DomainCreateRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DomainCreateRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *DomainCreateRequest) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


