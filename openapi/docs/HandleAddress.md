# HandleAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street** | **string** | Street | 
**StreetNumber** | **string** | Street Number | 
**City** | **string** | City | 
**Country** | **string** | Country | 
**ZipCode** | **string** | ZipCode | 
**Siret** | Pointer to **string** | Siret | [optional] 
**Region** | Pointer to **string** | Region | [optional] 

## Methods

### NewHandleAddress

`func NewHandleAddress(street string, streetNumber string, city string, country string, zipCode string, ) *HandleAddress`

NewHandleAddress instantiates a new HandleAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandleAddressWithDefaults

`func NewHandleAddressWithDefaults() *HandleAddress`

NewHandleAddressWithDefaults instantiates a new HandleAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet

`func (o *HandleAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *HandleAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *HandleAddress) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetStreetNumber

`func (o *HandleAddress) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *HandleAddress) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *HandleAddress) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.


### GetCity

`func (o *HandleAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *HandleAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *HandleAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *HandleAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *HandleAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *HandleAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetZipCode

`func (o *HandleAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *HandleAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *HandleAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetSiret

`func (o *HandleAddress) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *HandleAddress) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *HandleAddress) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *HandleAddress) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetRegion

`func (o *HandleAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandleAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandleAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandleAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


