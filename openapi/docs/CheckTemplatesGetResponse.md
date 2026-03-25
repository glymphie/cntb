# CheckTemplatesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CheckTemplateResponse**](CheckTemplateResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewCheckTemplatesGetResponse

`func NewCheckTemplatesGetResponse(data []CheckTemplateResponse, links SelfLinks, ) *CheckTemplatesGetResponse`

NewCheckTemplatesGetResponse instantiates a new CheckTemplatesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTemplatesGetResponseWithDefaults

`func NewCheckTemplatesGetResponseWithDefaults() *CheckTemplatesGetResponse`

NewCheckTemplatesGetResponseWithDefaults instantiates a new CheckTemplatesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CheckTemplatesGetResponse) GetData() []CheckTemplateResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckTemplatesGetResponse) GetDataOk() (*[]CheckTemplateResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CheckTemplatesGetResponse) SetData(v []CheckTemplateResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *CheckTemplatesGetResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CheckTemplatesGetResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CheckTemplatesGetResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


