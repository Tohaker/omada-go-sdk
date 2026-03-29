# FileServerSiteRestoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | **string** | Full path of Backup file, example: /backup/omada_backup.cfg. Parameter [filePath] should be 1 - 128 ASCII characters. | 
**SiteId** | **string** | Site ID to restore. | 

## Methods

### NewFileServerSiteRestoreVO

`func NewFileServerSiteRestoreVO(filePath string, siteId string, ) *FileServerSiteRestoreVO`

NewFileServerSiteRestoreVO instantiates a new FileServerSiteRestoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileServerSiteRestoreVOWithDefaults

`func NewFileServerSiteRestoreVOWithDefaults() *FileServerSiteRestoreVO`

NewFileServerSiteRestoreVOWithDefaults instantiates a new FileServerSiteRestoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *FileServerSiteRestoreVO) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileServerSiteRestoreVO) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileServerSiteRestoreVO) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetSiteId

`func (o *FileServerSiteRestoreVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *FileServerSiteRestoreVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *FileServerSiteRestoreVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


