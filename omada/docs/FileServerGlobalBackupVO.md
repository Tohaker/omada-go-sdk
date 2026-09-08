# FileServerGlobalBackupVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | **string** | Saving directory path for backup file. Parameter [filePath] should be 1 - 128 ASCII characters. | 
**RetainAuthRecord** | Pointer to **bool** | Whether need retain auth record. | [optional] 
**RetainFirmwareLog** | Pointer to **bool** | Whether need retain firmware log. | [optional] 
**RetainUser** | Pointer to **bool** | Whether need retain user info. | [optional] 
**Retention** | Pointer to **int32** | Backup data retention, values are as follows: 0 :no limit others:7，30，60，90，180，365. | [optional] 
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


### GetRetainAuthRecord

`func (o *FileServerGlobalBackupVO) GetRetainAuthRecord() bool`

GetRetainAuthRecord returns the RetainAuthRecord field if non-nil, zero value otherwise.

### GetRetainAuthRecordOk

`func (o *FileServerGlobalBackupVO) GetRetainAuthRecordOk() (*bool, bool)`

GetRetainAuthRecordOk returns a tuple with the RetainAuthRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAuthRecord

`func (o *FileServerGlobalBackupVO) SetRetainAuthRecord(v bool)`

SetRetainAuthRecord sets RetainAuthRecord field to given value.

### HasRetainAuthRecord

`func (o *FileServerGlobalBackupVO) HasRetainAuthRecord() bool`

HasRetainAuthRecord returns a boolean if a field has been set.

### GetRetainFirmwareLog

`func (o *FileServerGlobalBackupVO) GetRetainFirmwareLog() bool`

GetRetainFirmwareLog returns the RetainFirmwareLog field if non-nil, zero value otherwise.

### GetRetainFirmwareLogOk

`func (o *FileServerGlobalBackupVO) GetRetainFirmwareLogOk() (*bool, bool)`

GetRetainFirmwareLogOk returns a tuple with the RetainFirmwareLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFirmwareLog

`func (o *FileServerGlobalBackupVO) SetRetainFirmwareLog(v bool)`

SetRetainFirmwareLog sets RetainFirmwareLog field to given value.

### HasRetainFirmwareLog

`func (o *FileServerGlobalBackupVO) HasRetainFirmwareLog() bool`

HasRetainFirmwareLog returns a boolean if a field has been set.

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

### GetRetention

`func (o *FileServerGlobalBackupVO) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *FileServerGlobalBackupVO) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *FileServerGlobalBackupVO) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *FileServerGlobalBackupVO) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

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


