# DomainCancelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainCancel**](DomainCancel.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewDomainCancelResponse

`func NewDomainCancelResponse(data []DomainCancel, links SelfLinks, ) *DomainCancelResponse`

NewDomainCancelResponse instantiates a new DomainCancelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCancelResponseWithDefaults

`func NewDomainCancelResponseWithDefaults() *DomainCancelResponse`

NewDomainCancelResponseWithDefaults instantiates a new DomainCancelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainCancelResponse) GetData() []DomainCancel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainCancelResponse) GetDataOk() (*[]DomainCancel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainCancelResponse) SetData(v []DomainCancel)`

SetData sets Data field to given value.


### GetLinks

`func (o *DomainCancelResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainCancelResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainCancelResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


