# UpdateDnsZoneRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | **float32** | TTL | 
**Prio** | **float32** | Prio | 
**Type** | **string** | DNS record type | 
**Data** | **string** | Data | 
**Port** | Pointer to **float32** | Port | [optional] 
**Weight** | Pointer to **float32** | Weight | [optional] 
**Flag** | Pointer to **float32** | Flag | [optional] 
**Tag** | Pointer to **string** | Tag | [optional] 

## Methods

### NewUpdateDnsZoneRecordRequest

`func NewUpdateDnsZoneRecordRequest(ttl float32, prio float32, type_ string, data string, ) *UpdateDnsZoneRecordRequest`

NewUpdateDnsZoneRecordRequest instantiates a new UpdateDnsZoneRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDnsZoneRecordRequestWithDefaults

`func NewUpdateDnsZoneRecordRequestWithDefaults() *UpdateDnsZoneRecordRequest`

NewUpdateDnsZoneRecordRequestWithDefaults instantiates a new UpdateDnsZoneRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *UpdateDnsZoneRecordRequest) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UpdateDnsZoneRecordRequest) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UpdateDnsZoneRecordRequest) SetTtl(v float32)`

SetTtl sets Ttl field to given value.


### GetPrio

`func (o *UpdateDnsZoneRecordRequest) GetPrio() float32`

GetPrio returns the Prio field if non-nil, zero value otherwise.

### GetPrioOk

`func (o *UpdateDnsZoneRecordRequest) GetPrioOk() (*float32, bool)`

GetPrioOk returns a tuple with the Prio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrio

`func (o *UpdateDnsZoneRecordRequest) SetPrio(v float32)`

SetPrio sets Prio field to given value.


### GetType

`func (o *UpdateDnsZoneRecordRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateDnsZoneRecordRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateDnsZoneRecordRequest) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *UpdateDnsZoneRecordRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateDnsZoneRecordRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateDnsZoneRecordRequest) SetData(v string)`

SetData sets Data field to given value.


### GetPort

`func (o *UpdateDnsZoneRecordRequest) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateDnsZoneRecordRequest) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateDnsZoneRecordRequest) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateDnsZoneRecordRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetWeight

`func (o *UpdateDnsZoneRecordRequest) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateDnsZoneRecordRequest) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateDnsZoneRecordRequest) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *UpdateDnsZoneRecordRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetFlag

`func (o *UpdateDnsZoneRecordRequest) GetFlag() float32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *UpdateDnsZoneRecordRequest) GetFlagOk() (*float32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *UpdateDnsZoneRecordRequest) SetFlag(v float32)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *UpdateDnsZoneRecordRequest) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTag

`func (o *UpdateDnsZoneRecordRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UpdateDnsZoneRecordRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UpdateDnsZoneRecordRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *UpdateDnsZoneRecordRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


