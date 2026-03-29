# FileServerGlobalBackupVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | **string** | Saving directory path for backup file. Parameter [filePath] should be 1 - 128 ASCII characters. | 
**RetainUser** | Pointer to **bool** | Whether need retain user info. | [optional] 
**ServerConfig** | [**FileServerOpenApiVO**](FileServerOpenApiVO.md) |  | 

## Methods

### NewFileServerGlobalBackupVO

`func NewFileServerGlobalBackupVO(filePath string, serverConfig FileServerOpenApiVO, ) *FileServerGlobalBackupVO`

NewFileServerGlobalBackupVO instantiates a new FileServerGlobalBackupVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileServerGlobalBackupVOWithDefaults

`func NewFileServerGlobalBackupVOWithDefaults() *FileServerGlobalBackupVO`

NewFileServerGlobalBackupVOWithDefaults instantiates a new FileServerGlobalBackupVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *FileServerGlobalBackupVO) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileServerGlobalBackupVO) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileServerGlobalBackupVO) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetRetainUser

`func (o *FileServerGlobalBackupVO) GetRetainUser() bool`

GetRetainUser returns the RetainUser field if non-nil, zero value otherwise.

### GetRetainUserOk

`func (o *FileServerGlobalBackupVO) GetRetainUserOk() (*bool, bool)`

GetRetainUserOk returns a tuple with the RetainUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUser

`func (o *FileServerGlobalBackupVO) SetRetainUser(v bool)`

SetRetainUser sets RetainUser field to given value.

### HasRetainUser

`func (o *FileServerGlobalBackupVO) HasRetainUser() bool`

HasRetainUser returns a boolean if a field has been set.

### GetServerConfig

`func (o *FileServerGlobalBackupVO) GetServerConfig() FileServerOpenApiVO`

GetServerConfig returns the ServerConfig field if non-nil, zero value otherwise.

### GetServerConfigOk

`func (o *FileServerGlobalBackupVO) GetServerConfigOk() (*FileServerOpenApiVO, bool)`

GetServerConfigOk returns a tuple with the ServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfig

`func (o *FileServerGlobalBackupVO) SetServerConfig(v FileServerOpenApiVO)`

SetServerConfig sets ServerConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


