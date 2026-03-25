# HandleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**HandleId** | **string** | Handle ID | 
**HandleType** | **string** | Handle Type | 
**IsDefault** | **bool** | Flag if the handle is default or not | 
**FirstName** | **string** | Handle first name | 
**LastName** | **string** | Handle last name | 
**Organization** | Pointer to **string** | The organization of the handle | [optional] 
**Email** | **string** | Handle email | 
**Gender** | **string** | Handle gender | 
**BirthInfo** | Pointer to [**HandleBirthInfo**](HandleBirthInfo.md) | The birth info of the handle | [optional] 
**Address** | [**HandleAddress**](HandleAddress.md) | Address details for handle | 
**Phone** | [**HandlePhone**](HandlePhone.md) | Handle phone | 
**Fax** | Pointer to [**HandlePhone**](HandlePhone.md) | Handle fax | [optional] 

## Methods

### NewHandleResponse

`func NewHandleResponse(tenantId string, customerId string, handleId string, handleType string, isDefault bool, firstName string, lastName string, email string, gender string, address HandleAddress, phone HandlePhone, ) *HandleResponse`

NewHandleResponse instantiates a new HandleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandleResponseWithDefaults

`func NewHandleResponseWithDefaults() *HandleResponse`

NewHandleResponseWithDefaults instantiates a new HandleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *HandleResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HandleResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HandleResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *HandleResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *HandleResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *HandleResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetHandleId

`func (o *HandleResponse) GetHandleId() string`

GetHandleId returns the HandleId field if non-nil, zero value otherwise.

### GetHandleIdOk

`func (o *HandleResponse) GetHandleIdOk() (*string, bool)`

GetHandleIdOk returns a tuple with the HandleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleId

`func (o *HandleResponse) SetHandleId(v string)`

SetHandleId sets HandleId field to given value.


### GetHandleType

`func (o *HandleResponse) GetHandleType() string`

GetHandleType returns the HandleType field if non-nil, zero value otherwise.

### GetHandleTypeOk

`func (o *HandleResponse) GetHandleTypeOk() (*string, bool)`

GetHandleTypeOk returns a tuple with the HandleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleType

`func (o *HandleResponse) SetHandleType(v string)`

SetHandleType sets HandleType field to given value.


### GetIsDefault

`func (o *HandleResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *HandleResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *HandleResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetFirstName

`func (o *HandleResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *HandleResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *HandleResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *HandleResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *HandleResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *HandleResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetOrganization

`func (o *HandleResponse) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HandleResponse) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HandleResponse) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HandleResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEmail

`func (o *HandleResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *HandleResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *HandleResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGender

`func (o *HandleResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *HandleResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *HandleResponse) SetGender(v string)`

SetGender sets Gender field to given value.


### GetBirthInfo

`func (o *HandleResponse) GetBirthInfo() HandleBirthInfo`

GetBirthInfo returns the BirthInfo field if non-nil, zero value otherwise.

### GetBirthInfoOk

`func (o *HandleResponse) GetBirthInfoOk() (*HandleBirthInfo, bool)`

GetBirthInfoOk returns a tuple with the BirthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthInfo

`func (o *HandleResponse) SetBirthInfo(v HandleBirthInfo)`

SetBirthInfo sets BirthInfo field to given value.

### HasBirthInfo

`func (o *HandleResponse) HasBirthInfo() bool`

HasBirthInfo returns a boolean if a field has been set.

### GetAddress

`func (o *HandleResponse) GetAddress() HandleAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HandleResponse) GetAddressOk() (*HandleAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HandleResponse) SetAddress(v HandleAddress)`

SetAddress sets Address field to given value.


### GetPhone

`func (o *HandleResponse) GetPhone() HandlePhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *HandleResponse) GetPhoneOk() (*HandlePhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *HandleResponse) SetPhone(v HandlePhone)`

SetPhone sets Phone field to given value.


### GetFax

`func (o *HandleResponse) GetFax() HandlePhone`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *HandleResponse) GetFaxOk() (*HandlePhone, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *HandleResponse) SetFax(v HandlePhone)`

SetFax sets Fax field to given value.

### HasFax

`func (o *HandleResponse) HasFax() bool`

HasFax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


