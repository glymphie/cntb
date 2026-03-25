# ChecksGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CheckResponse**](CheckResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewChecksGetResponse

`func NewChecksGetResponse(data []CheckResponse, links SelfLinks, ) *ChecksGetResponse`

NewChecksGetResponse instantiates a new ChecksGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksGetResponseWithDefaults

`func NewChecksGetResponseWithDefaults() *ChecksGetResponse`

NewChecksGetResponseWithDefaults instantiates a new ChecksGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChecksGetResponse) GetData() []CheckResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChecksGetResponse) GetDataOk() (*[]CheckResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChecksGetResponse) SetData(v []CheckResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ChecksGetResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ChecksGetResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ChecksGetResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


