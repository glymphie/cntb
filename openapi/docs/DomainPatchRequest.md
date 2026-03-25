# DomainPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nameservers** | Pointer to [**[]Nameserver**](Nameserver.md) | Nameservers | [optional] 
**Handles** | Pointer to [**DomainHandles**](DomainHandles.md) | The handles of the domain | [optional] 

## Methods

### NewDomainPatchRequest

`func NewDomainPatchRequest() *DomainPatchRequest`

NewDomainPatchRequest instantiates a new DomainPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainPatchRequestWithDefaults

`func NewDomainPatchRequestWithDefaults() *DomainPatchRequest`

NewDomainPatchRequestWithDefaults instantiates a new DomainPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameservers

`func (o *DomainPatchRequest) GetNameservers() []Nameserver`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DomainPatchRequest) GetNameserversOk() (*[]Nameserver, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DomainPatchRequest) SetNameservers(v []Nameserver)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *DomainPatchRequest) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetHandles

`func (o *DomainPatchRequest) GetHandles() DomainHandles`

GetHandles returns the Handles field if non-nil, zero value otherwise.

### GetHandlesOk

`func (o *DomainPatchRequest) GetHandlesOk() (*DomainHandles, bool)`

GetHandlesOk returns a tuple with the Handles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandles

`func (o *DomainPatchRequest) SetHandles(v DomainHandles)`

SetHandles sets Handles field to given value.

### HasHandles

`func (o *DomainPatchRequest) HasHandles() bool`

HasHandles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


