# FileServerGlobalRestoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | **string** | Saving directory path for backup file. Parameter [filePath] should be 1 - 128 ASCII characters. | 
**ServerConfig** | [**FileServerOpenApiVO**](FileServerOpenApiVO.md) |  | 
**SkipDevice** | **bool** | Whether skip import devices. | 

## Methods

### NewFileServerGlobalRestoreVO

`func NewFileServerGlobalRestoreVO(filePath string, serverConfig FileServerOpenApiVO, skipDevice bool, ) *FileServerGlobalRestoreVO`

NewFileServerGlobalRestoreVO instantiates a new FileServerGlobalRestoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileServerGlobalRestoreVOWithDefaults

`func NewFileServerGlobalRestoreVOWithDefaults() *FileServerGlobalRestoreVO`

NewFileServerGlobalRestoreVOWithDefaults instantiates a new FileServerGlobalRestoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *FileServerGlobalRestoreVO) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileServerGlobalRestoreVO) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileServerGlobalRestoreVO) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetServerConfig

`func (o *FileServerGlobalRestoreVO) GetServerConfig() FileServerOpenApiVO`

GetServerConfig returns the ServerConfig field if non-nil, zero value otherwise.

### GetServerConfigOk

`func (o *FileServerGlobalRestoreVO) GetServerConfigOk() (*FileServerOpenApiVO, bool)`

GetServerConfigOk returns a tuple with the ServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfig

`func (o *FileServerGlobalRestoreVO) SetServerConfig(v FileServerOpenApiVO)`

SetServerConfig sets ServerConfig field to given value.


### GetSkipDevice

`func (o *FileServerGlobalRestoreVO) GetSkipDevice() bool`

GetSkipDevice returns the SkipDevice field if non-nil, zero value otherwise.

### GetSkipDeviceOk

`func (o *FileServerGlobalRestoreVO) GetSkipDeviceOk() (*bool, bool)`

GetSkipDeviceOk returns a tuple with the SkipDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDevice

`func (o *FileServerGlobalRestoreVO) SetSkipDevice(v bool)`

SetSkipDevice sets SkipDevice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


