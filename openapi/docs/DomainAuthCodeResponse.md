# DomainAuthCodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Domain** | **string** | Domain name | 
**DomainDetails** | [**DomainDetails**](DomainDetails.md) | Domain Details | 
**Status** | **string** | Domain Status | 
**Nameservers** | **[]string** | Nameservers | 
**Handles** | [**DomainHandles**](DomainHandles.md) | The handles of the domain | 
**RegistrationDate** | **time.Time** | The registration date of domain | 
**RenewalDate** | **time.Time** | The renewal date of domain | 
**TerminationDate** | **time.Time** | The termination date of domain | 
**CancelDate** | **time.Time** | The cancel date of domain | 
**DnssecKeys** | **[]string** | DNSSEC keys | 
**TransferOutConfirmation** | **bool** | Transfer out confirmation | 
**AuthCode** | **string** | Your auth code of the domain | 
**AuthCodeChanged** | [**ChangedAuthCode**](ChangedAuthCode.md) | Details if the auth code has been changed | 

## Methods

### NewDomainAuthCodeResponse

`func NewDomainAuthCodeResponse(tenantId string, customerId string, domain string, domainDetails DomainDetails, status string, nameservers []string, handles DomainHandles, registrationDate time.Time, renewalDate time.Time, terminationDate time.Time, cancelDate time.Time, dnssecKeys []string, transferOutConfirmation bool, authCode string, authCodeChanged ChangedAuthCode, ) *DomainAuthCodeResponse`

NewDomainAuthCodeResponse instantiates a new DomainAuthCodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAuthCodeResponseWithDefaults

`func NewDomainAuthCodeResponseWithDefaults() *DomainAuthCodeResponse`

NewDomainAuthCodeResponseWithDefaults instantiates a new DomainAuthCodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DomainAuthCodeResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DomainAuthCodeResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DomainAuthCodeResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DomainAuthCodeResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DomainAuthCodeResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DomainAuthCodeResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDomain

`func (o *DomainAuthCodeResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainAuthCodeResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainAuthCodeResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDomainDetails

`func (o *DomainAuthCodeResponse) GetDomainDetails() DomainDetails`

GetDomainDetails returns the DomainDetails field if non-nil, zero value otherwise.

### GetDomainDetailsOk

`func (o *DomainAuthCodeResponse) GetDomainDetailsOk() (*DomainDetails, bool)`

GetDomainDetailsOk returns a tuple with the DomainDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDetails

`func (o *DomainAuthCodeResponse) SetDomainDetails(v DomainDetails)`

SetDomainDetails sets DomainDetails field to given value.


### GetStatus

`func (o *DomainAuthCodeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainAuthCodeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainAuthCodeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNameservers

`func (o *DomainAuthCodeResponse) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DomainAuthCodeResponse) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DomainAuthCodeResponse) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.


### GetHandles

`func (o *DomainAuthCodeResponse) GetHandles() DomainHandles`

GetHandles returns the Handles field if non-nil, zero value otherwise.

### GetHandlesOk

`func (o *DomainAuthCodeResponse) GetHandlesOk() (*DomainHandles, bool)`

GetHandlesOk returns a tuple with the Handles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandles

`func (o *DomainAuthCodeResponse) SetHandles(v DomainHandles)`

SetHandles sets Handles field to given value.


### GetRegistrationDate

`func (o *DomainAuthCodeResponse) GetRegistrationDate() time.Time`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *DomainAuthCodeResponse) GetRegistrationDateOk() (*time.Time, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *DomainAuthCodeResponse) SetRegistrationDate(v time.Time)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetRenewalDate

`func (o *DomainAuthCodeResponse) GetRenewalDate() time.Time`

GetRenewalDate returns the RenewalDate field if non-nil, zero value otherwise.

### GetRenewalDateOk

`func (o *DomainAuthCodeResponse) GetRenewalDateOk() (*time.Time, bool)`

GetRenewalDateOk returns a tuple with the RenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalDate

`func (o *DomainAuthCodeResponse) SetRenewalDate(v time.Time)`

SetRenewalDate sets RenewalDate field to given value.


### GetTerminationDate

`func (o *DomainAuthCodeResponse) GetTerminationDate() time.Time`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *DomainAuthCodeResponse) GetTerminationDateOk() (*time.Time, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *DomainAuthCodeResponse) SetTerminationDate(v time.Time)`

SetTerminationDate sets TerminationDate field to given value.


### GetCancelDate

`func (o *DomainAuthCodeResponse) GetCancelDate() time.Time`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *DomainAuthCodeResponse) GetCancelDateOk() (*time.Time, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *DomainAuthCodeResponse) SetCancelDate(v time.Time)`

SetCancelDate sets CancelDate field to given value.


### GetDnssecKeys

`func (o *DomainAuthCodeResponse) GetDnssecKeys() []string`

GetDnssecKeys returns the DnssecKeys field if non-nil, zero value otherwise.

### GetDnssecKeysOk

`func (o *DomainAuthCodeResponse) GetDnssecKeysOk() (*[]string, bool)`

GetDnssecKeysOk returns a tuple with the DnssecKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecKeys

`func (o *DomainAuthCodeResponse) SetDnssecKeys(v []string)`

SetDnssecKeys sets DnssecKeys field to given value.


### GetTransferOutConfirmation

`func (o *DomainAuthCodeResponse) GetTransferOutConfirmation() bool`

GetTransferOutConfirmation returns the TransferOutConfirmation field if non-nil, zero value otherwise.

### GetTransferOutConfirmationOk

`func (o *DomainAuthCodeResponse) GetTransferOutConfirmationOk() (*bool, bool)`

GetTransferOutConfirmationOk returns a tuple with the TransferOutConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferOutConfirmation

`func (o *DomainAuthCodeResponse) SetTransferOutConfirmation(v bool)`

SetTransferOutConfirmation sets TransferOutConfirmation field to given value.


### GetAuthCode

`func (o *DomainAuthCodeResponse) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *DomainAuthCodeResponse) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *DomainAuthCodeResponse) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetAuthCodeChanged

`func (o *DomainAuthCodeResponse) GetAuthCodeChanged() ChangedAuthCode`

GetAuthCodeChanged returns the AuthCodeChanged field if non-nil, zero value otherwise.

### GetAuthCodeChangedOk

`func (o *DomainAuthCodeResponse) GetAuthCodeChangedOk() (*ChangedAuthCode, bool)`

GetAuthCodeChangedOk returns a tuple with the AuthCodeChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCodeChanged

`func (o *DomainAuthCodeResponse) SetAuthCodeChanged(v ChangedAuthCode)`

SetAuthCodeChanged sets AuthCodeChanged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


