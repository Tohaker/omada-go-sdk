# AutoBackupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Occurrence** | Pointer to [**BaseScheduleTimeVO**](BaseScheduleTimeVO.md) |  | [optional] 
**AvailablePaths** | Pointer to **[]string** | Available paths for backup schedule | [optional] 
**DataSheets** | Pointer to **[]string** | Data to backup, values are as follows: knowClient, omadaLog, auditLog. | [optional] 
**Enable** | **bool** | Whether enable backup schedule task | 
**FileServerConfig** | Pointer to [**FileServerOpenApiVO**](FileServerOpenApiVO.md) |  | [optional] 
**MaxNumberOfFile** | Pointer to **int32** | Max number of schedule backup files | [optional] 
**RetainAuthRecord** | Pointer to **bool** | Whether retain auth record when backup. | [optional] 
**RetainFirmwareLog** | Pointer to **bool** | Whether retain firmware log when backup. | [optional] 
**RetainSetting** | Pointer to **bool** | Whether retain setting when backup. | [optional] 
**RetainUser** | Pointer to **bool** | Whether retain user when backup. | [optional] 
**Retention** | Pointer to **int32** | Retention setting for data, values are as follows: -1 retention settings only, 0 backup all data, only for software controller, 7 retention backup data for 7 days, 30 retention backup data for 30 days, 60 retention backup data for 60 days, 90 retention backup data for 90 days, 180 retention backup data for 180 days, 365 retention backup data for 365 days | [optional] 
**SavingPath** | Pointer to **string** | Saving path for backup schedule | [optional] 
**StorageType** | Pointer to **int32** | Storage type, values are as follows: 0 for file server, 1 for cloud server(only for Cloud based controller), 2 for local disk, null for default path | [optional] 

## Methods

### NewAutoBackupOpenApiVO

`func NewAutoBackupOpenApiVO(enable bool, ) *AutoBackupOpenApiVO`

NewAutoBackupOpenApiVO instantiates a new AutoBackupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoBackupOpenApiVOWithDefaults

`func NewAutoBackupOpenApiVOWithDefaults() *AutoBackupOpenApiVO`

NewAutoBackupOpenApiVOWithDefaults instantiates a new AutoBackupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurrence

`func (o *AutoBackupOpenApiVO) GetOccurrence() BaseScheduleTimeVO`

GetOccurrence returns the Occurrence field if non-nil, zero value otherwise.

### GetOccurrenceOk

`func (o *AutoBackupOpenApiVO) GetOccurrenceOk() (*BaseScheduleTimeVO, bool)`

GetOccurrenceOk returns a tuple with the Occurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrence

`func (o *AutoBackupOpenApiVO) SetOccurrence(v BaseScheduleTimeVO)`

SetOccurrence sets Occurrence field to given value.

### HasOccurrence

`func (o *AutoBackupOpenApiVO) HasOccurrence() bool`

HasOccurrence returns a boolean if a field has been set.

### GetAvailablePaths

`func (o *AutoBackupOpenApiVO) GetAvailablePaths() []string`

GetAvailablePaths returns the AvailablePaths field if non-nil, zero value otherwise.

### GetAvailablePathsOk

`func (o *AutoBackupOpenApiVO) GetAvailablePathsOk() (*[]string, bool)`

GetAvailablePathsOk returns a tuple with the AvailablePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaths

`func (o *AutoBackupOpenApiVO) SetAvailablePaths(v []string)`

SetAvailablePaths sets AvailablePaths field to given value.

### HasAvailablePaths

`func (o *AutoBackupOpenApiVO) HasAvailablePaths() bool`

HasAvailablePaths returns a boolean if a field has been set.

### GetDataSheets

`func (o *AutoBackupOpenApiVO) GetDataSheets() []string`

GetDataSheets returns the DataSheets field if non-nil, zero value otherwise.

### GetDataSheetsOk

`func (o *AutoBackupOpenApiVO) GetDataSheetsOk() (*[]string, bool)`

GetDataSheetsOk returns a tuple with the DataSheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSheets

`func (o *AutoBackupOpenApiVO) SetDataSheets(v []string)`

SetDataSheets sets DataSheets field to given value.

### HasDataSheets

`func (o *AutoBackupOpenApiVO) HasDataSheets() bool`

HasDataSheets returns a boolean if a field has been set.

### GetEnable

`func (o *AutoBackupOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AutoBackupOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AutoBackupOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetFileServerConfig

`func (o *AutoBackupOpenApiVO) GetFileServerConfig() FileServerOpenApiVO`

GetFileServerConfig returns the FileServerConfig field if non-nil, zero value otherwise.

### GetFileServerConfigOk

`func (o *AutoBackupOpenApiVO) GetFileServerConfigOk() (*FileServerOpenApiVO, bool)`

GetFileServerConfigOk returns a tuple with the FileServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServerConfig

`func (o *AutoBackupOpenApiVO) SetFileServerConfig(v FileServerOpenApiVO)`

SetFileServerConfig sets FileServerConfig field to given value.

### HasFileServerConfig

`func (o *AutoBackupOpenApiVO) HasFileServerConfig() bool`

HasFileServerConfig returns a boolean if a field has been set.

### GetMaxNumberOfFile

`func (o *AutoBackupOpenApiVO) GetMaxNumberOfFile() int32`

GetMaxNumberOfFile returns the MaxNumberOfFile field if non-nil, zero value otherwise.

### GetMaxNumberOfFileOk

`func (o *AutoBackupOpenApiVO) GetMaxNumberOfFileOk() (*int32, bool)`

GetMaxNumberOfFileOk returns a tuple with the MaxNumberOfFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfFile

`func (o *AutoBackupOpenApiVO) SetMaxNumberOfFile(v int32)`

SetMaxNumberOfFile sets MaxNumberOfFile field to given value.

### HasMaxNumberOfFile

`func (o *AutoBackupOpenApiVO) HasMaxNumberOfFile() bool`

HasMaxNumberOfFile returns a boolean if a field has been set.

### GetRetainAuthRecord

`func (o *AutoBackupOpenApiVO) GetRetainAuthRecord() bool`

GetRetainAuthRecord returns the RetainAuthRecord field if non-nil, zero value otherwise.

### GetRetainAuthRecordOk

`func (o *AutoBackupOpenApiVO) GetRetainAuthRecordOk() (*bool, bool)`

GetRetainAuthRecordOk returns a tuple with the RetainAuthRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAuthRecord

`func (o *AutoBackupOpenApiVO) SetRetainAuthRecord(v bool)`

SetRetainAuthRecord sets RetainAuthRecord field to given value.

### HasRetainAuthRecord

`func (o *AutoBackupOpenApiVO) HasRetainAuthRecord() bool`

HasRetainAuthRecord returns a boolean if a field has been set.

### GetRetainFirmwareLog

`func (o *AutoBackupOpenApiVO) GetRetainFirmwareLog() bool`

GetRetainFirmwareLog returns the RetainFirmwareLog field if non-nil, zero value otherwise.

### GetRetainFirmwareLogOk

`func (o *AutoBackupOpenApiVO) GetRetainFirmwareLogOk() (*bool, bool)`

GetRetainFirmwareLogOk returns a tuple with the RetainFirmwareLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFirmwareLog

`func (o *AutoBackupOpenApiVO) SetRetainFirmwareLog(v bool)`

SetRetainFirmwareLog sets RetainFirmwareLog field to given value.

### HasRetainFirmwareLog

`func (o *AutoBackupOpenApiVO) HasRetainFirmwareLog() bool`

HasRetainFirmwareLog returns a boolean if a field has been set.

### GetRetainSetting

`func (o *AutoBackupOpenApiVO) GetRetainSetting() bool`

GetRetainSetting returns the RetainSetting field if non-nil, zero value otherwise.

### GetRetainSettingOk

`func (o *AutoBackupOpenApiVO) GetRetainSettingOk() (*bool, bool)`

GetRetainSettingOk returns a tuple with the RetainSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainSetting

`func (o *AutoBackupOpenApiVO) SetRetainSetting(v bool)`

SetRetainSetting sets RetainSetting field to given value.

### HasRetainSetting

`func (o *AutoBackupOpenApiVO) HasRetainSetting() bool`

HasRetainSetting returns a boolean if a field has been set.

### GetRetainUser

`func (o *AutoBackupOpenApiVO) GetRetainUser() bool`

GetRetainUser returns the RetainUser field if non-nil, zero value otherwise.

### GetRetainUserOk

`func (o *AutoBackupOpenApiVO) GetRetainUserOk() (*bool, bool)`

GetRetainUserOk returns a tuple with the RetainUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUser

`func (o *AutoBackupOpenApiVO) SetRetainUser(v bool)`

SetRetainUser sets RetainUser field to given value.

### HasRetainUser

`func (o *AutoBackupOpenApiVO) HasRetainUser() bool`

HasRetainUser returns a boolean if a field has been set.

### GetRetention

`func (o *AutoBackupOpenApiVO) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *AutoBackupOpenApiVO) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *AutoBackupOpenApiVO) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *AutoBackupOpenApiVO) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSavingPath

`func (o *AutoBackupOpenApiVO) GetSavingPath() string`

GetSavingPath returns the SavingPath field if non-nil, zero value otherwise.

### GetSavingPathOk

`func (o *AutoBackupOpenApiVO) GetSavingPathOk() (*string, bool)`

GetSavingPathOk returns a tuple with the SavingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingPath

`func (o *AutoBackupOpenApiVO) SetSavingPath(v string)`

SetSavingPath sets SavingPath field to given value.

### HasSavingPath

`func (o *AutoBackupOpenApiVO) HasSavingPath() bool`

HasSavingPath returns a boolean if a field has been set.

### GetStorageType

`func (o *AutoBackupOpenApiVO) GetStorageType() int32`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AutoBackupOpenApiVO) GetStorageTypeOk() (*int32, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AutoBackupOpenApiVO) SetStorageType(v int32)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *AutoBackupOpenApiVO) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


