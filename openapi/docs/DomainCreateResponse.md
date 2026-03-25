# DomainCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainResponse**](DomainResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewDomainCreateResponse

`func NewDomainCreateResponse(data []DomainResponse, links SelfLinks, ) *DomainCreateResponse`

NewDomainCreateResponse instantiates a new DomainCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCreateResponseWithDefaults

`func NewDomainCreateResponseWithDefaults() *DomainCreateResponse`

NewDomainCreateResponseWithDefaults instantiates a new DomainCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainCreateResponse) GetData() []DomainResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainCreateResponse) GetDataOk() (*[]DomainResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainCreateResponse) SetData(v []DomainResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *DomainCreateResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainCreateResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainCreateResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


