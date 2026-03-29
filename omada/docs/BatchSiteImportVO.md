# BatchSiteImportVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileServerConfig** | [**FileServerOpenApiVO**](FileServerOpenApiVO.md) |  | 
**SiteImportConfigList** | [**[]SiteImportOpenApiVO**](SiteImportOpenApiVO.md) | Site import config list. max size 300. | 

## Methods

### NewBatchSiteImportVO

`func NewBatchSiteImportVO(fileServerConfig FileServerOpenApiVO, siteImportConfigList []SiteImportOpenApiVO, ) *BatchSiteImportVO`

NewBatchSiteImportVO instantiates a new BatchSiteImportVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSiteImportVOWithDefaults

`func NewBatchSiteImportVOWithDefaults() *BatchSiteImportVO`

NewBatchSiteImportVOWithDefaults instantiates a new BatchSiteImportVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileServerConfig

`func (o *BatchSiteImportVO) GetFileServerConfig() FileServerOpenApiVO`

GetFileServerConfig returns the FileServerConfig field if non-nil, zero value otherwise.

### GetFileServerConfigOk

`func (o *BatchSiteImportVO) GetFileServerConfigOk() (*FileServerOpenApiVO, bool)`

GetFileServerConfigOk returns a tuple with the FileServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServerConfig

`func (o *BatchSiteImportVO) SetFileServerConfig(v FileServerOpenApiVO)`

SetFileServerConfig sets FileServerConfig field to given value.


### GetSiteImportConfigList

`func (o *BatchSiteImportVO) GetSiteImportConfigList() []SiteImportOpenApiVO`

GetSiteImportConfigList returns the SiteImportConfigList field if non-nil, zero value otherwise.

### GetSiteImportConfigListOk

`func (o *BatchSiteImportVO) GetSiteImportConfigListOk() (*[]SiteImportOpenApiVO, bool)`

GetSiteImportConfigListOk returns a tuple with the SiteImportConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteImportConfigList

`func (o *BatchSiteImportVO) SetSiteImportConfigList(v []SiteImportOpenApiVO)`

SetSiteImportConfigList sets SiteImportConfigList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


