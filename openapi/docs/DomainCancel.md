# DomainCancel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Domain** | **string** | Domain name | 
**Sld** | **string** | Domain SLD | 
**Tld** | **string** | Domain TLD | 
**CancelDate** | **time.Time** | The cancel date of domain | 

## Methods

### NewDomainCancel

`func NewDomainCancel(tenantId string, customerId string, domain string, sld string, tld string, cancelDate time.Time, ) *DomainCancel`

NewDomainCancel instantiates a new DomainCancel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCancelWithDefaults

`func NewDomainCancelWithDefaults() *DomainCancel`

NewDomainCancelWithDefaults instantiates a new DomainCancel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DomainCancel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DomainCancel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DomainCancel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DomainCancel) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DomainCancel) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DomainCancel) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDomain

`func (o *DomainCancel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainCancel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainCancel) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetSld

`func (o *DomainCancel) GetSld() string`

GetSld returns the Sld field if non-nil, zero value otherwise.

### GetSldOk

`func (o *DomainCancel) GetSldOk() (*string, bool)`

GetSldOk returns a tuple with the Sld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSld

`func (o *DomainCancel) SetSld(v string)`

SetSld sets Sld field to given value.


### GetTld

`func (o *DomainCancel) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *DomainCancel) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *DomainCancel) SetTld(v string)`

SetTld sets Tld field to given value.


### GetCancelDate

`func (o *DomainCancel) GetCancelDate() time.Time`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *DomainCancel) GetCancelDateOk() (*time.Time, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *DomainCancel) SetCancelDate(v time.Time)`

SetCancelDate sets CancelDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


