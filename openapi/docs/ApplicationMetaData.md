# ApplicationMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | **[]string** |  | 
**LogoUrl** | **string** |  | 
**InitialUsername** | **string** |  | 
**DocumentationUrls** | **[]string** |  | 
**RequiresPasswordForInstall** | **bool** |  | 

## Methods

### NewApplicationMetaData

`func NewApplicationMetaData(urls []string, logoUrl string, initialUsername string, documentationUrls []string, requiresPasswordForInstall bool, ) *ApplicationMetaData`

NewApplicationMetaData instantiates a new ApplicationMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationMetaDataWithDefaults

`func NewApplicationMetaDataWithDefaults() *ApplicationMetaData`

NewApplicationMetaDataWithDefaults instantiates a new ApplicationMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *ApplicationMetaData) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *ApplicationMetaData) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *ApplicationMetaData) SetUrls(v []string)`

SetUrls sets Urls field to given value.


### GetLogoUrl

`func (o *ApplicationMetaData) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ApplicationMetaData) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ApplicationMetaData) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetInitialUsername

`func (o *ApplicationMetaData) GetInitialUsername() string`

GetInitialUsername returns the InitialUsername field if non-nil, zero value otherwise.

### GetInitialUsernameOk

`func (o *ApplicationMetaData) GetInitialUsernameOk() (*string, bool)`

GetInitialUsernameOk returns a tuple with the InitialUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUsername

`func (o *ApplicationMetaData) SetInitialUsername(v string)`

SetInitialUsername sets InitialUsername field to given value.


### GetDocumentationUrls

`func (o *ApplicationMetaData) GetDocumentationUrls() []string`

GetDocumentationUrls returns the DocumentationUrls field if non-nil, zero value otherwise.

### GetDocumentationUrlsOk

`func (o *ApplicationMetaData) GetDocumentationUrlsOk() (*[]string, bool)`

GetDocumentationUrlsOk returns a tuple with the DocumentationUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrls

`func (o *ApplicationMetaData) SetDocumentationUrls(v []string)`

SetDocumentationUrls sets DocumentationUrls field to given value.


### GetRequiresPasswordForInstall

`func (o *ApplicationMetaData) GetRequiresPasswordForInstall() bool`

GetRequiresPasswordForInstall returns the RequiresPasswordForInstall field if non-nil, zero value otherwise.

### GetRequiresPasswordForInstallOk

`func (o *ApplicationMetaData) GetRequiresPasswordForInstallOk() (*bool, bool)`

GetRequiresPasswordForInstallOk returns a tuple with the RequiresPasswordForInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPasswordForInstall

`func (o *ApplicationMetaData) SetRequiresPasswordForInstall(v bool)`

SetRequiresPasswordForInstall sets RequiresPasswordForInstall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


