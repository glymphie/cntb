# ExtCheckCollectionsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ExtCheckCollectionResponse**](ExtCheckCollectionResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewExtCheckCollectionsGetResponse

`func NewExtCheckCollectionsGetResponse(data []ExtCheckCollectionResponse, links SelfLinks, ) *ExtCheckCollectionsGetResponse`

NewExtCheckCollectionsGetResponse instantiates a new ExtCheckCollectionsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtCheckCollectionsGetResponseWithDefaults

`func NewExtCheckCollectionsGetResponseWithDefaults() *ExtCheckCollectionsGetResponse`

NewExtCheckCollectionsGetResponseWithDefaults instantiates a new ExtCheckCollectionsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExtCheckCollectionsGetResponse) GetData() []ExtCheckCollectionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExtCheckCollectionsGetResponse) GetDataOk() (*[]ExtCheckCollectionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExtCheckCollectionsGetResponse) SetData(v []ExtCheckCollectionResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ExtCheckCollectionsGetResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtCheckCollectionsGetResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtCheckCollectionsGetResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


