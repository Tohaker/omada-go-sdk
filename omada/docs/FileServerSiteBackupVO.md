# FileServerSiteBackupVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | **string** | Saving path of directory for all backup config files, example: /backup. Parameter [filePath] should be 1 - 128 ASCII characters. | 
**ServerConfig** | [**FileServerOpenApiVO**](FileServerOpenApiVO.md) |  | 
**SiteIds** | **[]string** | Site ID list to backup, up to 300 entries are allowed for the site ID list. | 

## Methods

### NewFileServerSiteBackupVO

`func NewFileServerSiteBackupVO(filePath string, serverConfig FileServerOpenApiVO, siteIds []string, ) *FileServerSiteBackupVO`

NewFileServerSiteBackupVO instantiates a new FileServerSiteBackupVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileServerSiteBackupVOWithDefaults

`func NewFileServerSiteBackupVOWithDefaults() *FileServerSiteBackupVO`

NewFileServerSiteBackupVOWithDefaults instantiates a new FileServerSiteBackupVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *FileServerSiteBackupVO) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileServerSiteBackupVO) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileServerSiteBackupVO) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetServerConfig

`func (o *FileServerSiteBackupVO) GetServerConfig() FileServerOpenApiVO`

GetServerConfig returns the ServerConfig field if non-nil, zero value otherwise.

### GetServerConfigOk

`func (o *FileServerSiteBackupVO) GetServerConfigOk() (*FileServerOpenApiVO, bool)`

GetServerConfigOk returns a tuple with the ServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfig

`func (o *FileServerSiteBackupVO) SetServerConfig(v FileServerOpenApiVO)`

SetServerConfig sets ServerConfig field to given value.


### GetSiteIds

`func (o *FileServerSiteBackupVO) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *FileServerSiteBackupVO) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *FileServerSiteBackupVO) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


