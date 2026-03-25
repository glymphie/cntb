# HandleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandleType** | **string** | The type of the handle | 
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

### NewHandleCreateRequest

`func NewHandleCreateRequest(handleType string, firstName string, lastName string, email string, gender string, address HandleAddress, phone HandlePhone, ) *HandleCreateRequest`

NewHandleCreateRequest instantiates a new HandleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandleCreateRequestWithDefaults

`func NewHandleCreateRequestWithDefaults() *HandleCreateRequest`

NewHandleCreateRequestWithDefaults instantiates a new HandleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandleType

`func (o *HandleCreateRequest) GetHandleType() string`

GetHandleType returns the HandleType field if non-nil, zero value otherwise.

### GetHandleTypeOk

`func (o *HandleCreateRequest) GetHandleTypeOk() (*string, bool)`

GetHandleTypeOk returns a tuple with the HandleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleType

`func (o *HandleCreateRequest) SetHandleType(v string)`

SetHandleType sets HandleType field to given value.


### GetFirstName

`func (o *HandleCreateRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *HandleCreateRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *HandleCreateRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *HandleCreateRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *HandleCreateRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *HandleCreateRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetOrganization

`func (o *HandleCreateRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HandleCreateRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HandleCreateRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HandleCreateRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEmail

`func (o *HandleCreateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *HandleCreateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *HandleCreateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGender

`func (o *HandleCreateRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *HandleCreateRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *HandleCreateRequest) SetGender(v string)`

SetGender sets Gender field to given value.


### GetBirthInfo

`func (o *HandleCreateRequest) GetBirthInfo() HandleBirthInfo`

GetBirthInfo returns the BirthInfo field if non-nil, zero value otherwise.

### GetBirthInfoOk

`func (o *HandleCreateRequest) GetBirthInfoOk() (*HandleBirthInfo, bool)`

GetBirthInfoOk returns a tuple with the BirthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthInfo

`func (o *HandleCreateRequest) SetBirthInfo(v HandleBirthInfo)`

SetBirthInfo sets BirthInfo field to given value.

### HasBirthInfo

`func (o *HandleCreateRequest) HasBirthInfo() bool`

HasBirthInfo returns a boolean if a field has been set.

### GetAddress

`func (o *HandleCreateRequest) GetAddress() HandleAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HandleCreateRequest) GetAddressOk() (*HandleAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HandleCreateRequest) SetAddress(v HandleAddress)`

SetAddress sets Address field to given value.


### GetPhone

`func (o *HandleCreateRequest) GetPhone() HandlePhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *HandleCreateRequest) GetPhoneOk() (*HandlePhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *HandleCreateRequest) SetPhone(v HandlePhone)`

SetPhone sets Phone field to given value.


### GetFax

`func (o *HandleCreateRequest) GetFax() HandlePhone`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *HandleCreateRequest) GetFaxOk() (*HandlePhone, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *HandleCreateRequest) SetFax(v HandlePhone)`

SetFax sets Fax field to given value.

### HasFax

`func (o *HandleCreateRequest) HasFax() bool`

HasFax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


