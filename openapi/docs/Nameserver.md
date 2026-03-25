# Nameserver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **[]string** | Nameservers | 
**IpV4** | Pointer to **[]string** | IPv4 of nameserver | [optional] 
**IpV6** | Pointer to **[]string** | IPv6 of nameserver | [optional] 

## Methods

### NewNameserver

`func NewNameserver(hostname []string, ) *Nameserver`

NewNameserver instantiates a new Nameserver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameserverWithDefaults

`func NewNameserverWithDefaults() *Nameserver`

NewNameserverWithDefaults instantiates a new Nameserver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *Nameserver) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Nameserver) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Nameserver) SetHostname(v []string)`

SetHostname sets Hostname field to given value.


### GetIpV4

`func (o *Nameserver) GetIpV4() []string`

GetIpV4 returns the IpV4 field if non-nil, zero value otherwise.

### GetIpV4Ok

`func (o *Nameserver) GetIpV4Ok() (*[]string, bool)`

GetIpV4Ok returns a tuple with the IpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4

`func (o *Nameserver) SetIpV4(v []string)`

SetIpV4 sets IpV4 field to given value.

### HasIpV4

`func (o *Nameserver) HasIpV4() bool`

HasIpV4 returns a boolean if a field has been set.

### GetIpV6

`func (o *Nameserver) GetIpV6() []string`

GetIpV6 returns the IpV6 field if non-nil, zero value otherwise.

### GetIpV6Ok

`func (o *Nameserver) GetIpV6Ok() (*[]string, bool)`

GetIpV6Ok returns a tuple with the IpV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6

`func (o *Nameserver) SetIpV6(v []string)`

SetIpV6 sets IpV6 field to given value.

### HasIpV6

`func (o *Nameserver) HasIpV6() bool`

HasIpV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


