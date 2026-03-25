# HandlePatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Handle email | 
**Gender** | **string** | Handle gender | 
**BirthInfo** | Pointer to [**HandleBirthInfo**](HandleBirthInfo.md) | The birth info of the handle | [optional] 
**Address** | [**HandleAddress**](HandleAddress.md) | Address details for handle | 
**Phone** | [**HandlePhone**](HandlePhone.md) | Handle phone | 
**Fax** | Pointer to [**HandlePhone**](HandlePhone.md) | Handle fax | [optional] 

## Methods

### NewHandlePatchRequest

`func NewHandlePatchRequest(email string, gender string, address HandleAddress, phone HandlePhone, ) *HandlePatchRequest`

NewHandlePatchRequest instantiates a new HandlePatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlePatchRequestWithDefaults

`func NewHandlePatchRequestWithDefaults() *HandlePatchRequest`

NewHandlePatchRequestWithDefaults instantiates a new HandlePatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *HandlePatchRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *HandlePatchRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *HandlePatchRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGender

`func (o *HandlePatchRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *HandlePatchRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *HandlePatchRequest) SetGender(v string)`

SetGender sets Gender field to given value.


### GetBirthInfo

`func (o *HandlePatchRequest) GetBirthInfo() HandleBirthInfo`

GetBirthInfo returns the BirthInfo field if non-nil, zero value otherwise.

### GetBirthInfoOk

`func (o *HandlePatchRequest) GetBirthInfoOk() (*HandleBirthInfo, bool)`

GetBirthInfoOk returns a tuple with the BirthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthInfo

`func (o *HandlePatchRequest) SetBirthInfo(v HandleBirthInfo)`

SetBirthInfo sets BirthInfo field to given value.

### HasBirthInfo

`func (o *HandlePatchRequest) HasBirthInfo() bool`

HasBirthInfo returns a boolean if a field has been set.

### GetAddress

`func (o *HandlePatchRequest) GetAddress() HandleAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HandlePatchRequest) GetAddressOk() (*HandleAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HandlePatchRequest) SetAddress(v HandleAddress)`

SetAddress sets Address field to given value.


### GetPhone

`func (o *HandlePatchRequest) GetPhone() HandlePhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *HandlePatchRequest) GetPhoneOk() (*HandlePhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *HandlePatchRequest) SetPhone(v HandlePhone)`

SetPhone sets Phone field to given value.


### GetFax

`func (o *HandlePatchRequest) GetFax() HandlePhone`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *HandlePatchRequest) GetFaxOk() (*HandlePhone, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *HandlePatchRequest) SetFax(v HandlePhone)`

SetFax sets Fax field to given value.

### HasFax

`func (o *HandlePatchRequest) HasFax() bool`

HasFax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


