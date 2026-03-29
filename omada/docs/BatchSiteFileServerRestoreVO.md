# BatchSiteFileServerRestoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerConfig** | [**FileServerOpenApiVO**](FileServerOpenApiVO.md) |  | 
**SiteInfos** | [**[]FileServerSiteRestoreVO**](FileServerSiteRestoreVO.md) | Sites to restore. | 

## Methods

### NewBatchSiteFileServerRestoreVO

`func NewBatchSiteFileServerRestoreVO(serverConfig FileServerOpenApiVO, siteInfos []FileServerSiteRestoreVO, ) *BatchSiteFileServerRestoreVO`

NewBatchSiteFileServerRestoreVO instantiates a new BatchSiteFileServerRestoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSiteFileServerRestoreVOWithDefaults

`func NewBatchSiteFileServerRestoreVOWithDefaults() *BatchSiteFileServerRestoreVO`

NewBatchSiteFileServerRestoreVOWithDefaults instantiates a new BatchSiteFileServerRestoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerConfig

`func (o *BatchSiteFileServerRestoreVO) GetServerConfig() FileServerOpenApiVO`

GetServerConfig returns the ServerConfig field if non-nil, zero value otherwise.

### GetServerConfigOk

`func (o *BatchSiteFileServerRestoreVO) GetServerConfigOk() (*FileServerOpenApiVO, bool)`

GetServerConfigOk returns a tuple with the ServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfig

`func (o *BatchSiteFileServerRestoreVO) SetServerConfig(v FileServerOpenApiVO)`

SetServerConfig sets ServerConfig field to given value.


### GetSiteInfos

`func (o *BatchSiteFileServerRestoreVO) GetSiteInfos() []FileServerSiteRestoreVO`

GetSiteInfos returns the SiteInfos field if non-nil, zero value otherwise.

### GetSiteInfosOk

`func (o *BatchSiteFileServerRestoreVO) GetSiteInfosOk() (*[]FileServerSiteRestoreVO, bool)`

GetSiteInfosOk returns a tuple with the SiteInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteInfos

`func (o *BatchSiteFileServerRestoreVO) SetSiteInfos(v []FileServerSiteRestoreVO)`

SetSiteInfos sets SiteInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


