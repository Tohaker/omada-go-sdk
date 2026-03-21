# SiteImportOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | **string** | File path of site backup config file. | 
**SiteName** | **string** | Target site name. It should contain 1 to 64 characters. | 
**SkipDevice** | Pointer to **bool** | Whether skip device info(if true: skip import device; if false: import device in config file; default for false | [optional] 

## Methods

### NewSiteImportOpenApiVO

`func NewSiteImportOpenApiVO(filePath string, siteName string, ) *SiteImportOpenApiVO`

NewSiteImportOpenApiVO instantiates a new SiteImportOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteImportOpenApiVOWithDefaults

`func NewSiteImportOpenApiVOWithDefaults() *SiteImportOpenApiVO`

NewSiteImportOpenApiVOWithDefaults instantiates a new SiteImportOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *SiteImportOpenApiVO) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SiteImportOpenApiVO) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SiteImportOpenApiVO) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetSiteName

`func (o *SiteImportOpenApiVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *SiteImportOpenApiVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *SiteImportOpenApiVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetSkipDevice

`func (o *SiteImportOpenApiVO) GetSkipDevice() bool`

GetSkipDevice returns the SkipDevice field if non-nil, zero value otherwise.

### GetSkipDeviceOk

`func (o *SiteImportOpenApiVO) GetSkipDeviceOk() (*bool, bool)`

GetSkipDeviceOk returns a tuple with the SkipDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDevice

`func (o *SiteImportOpenApiVO) SetSkipDevice(v bool)`

SetSkipDevice sets SkipDevice field to given value.

### HasSkipDevice

`func (o *SiteImportOpenApiVO) HasSkipDevice() bool`

HasSkipDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


