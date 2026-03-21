# BackupFileResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupTime** | Pointer to **int64** | Backup time(ms). | [optional] 
**FileName** | Pointer to **string** | File name of backup file. Parameter [fileName] should be 1 - 128 ASCII characters. | [optional] 
**Size** | Pointer to **int64** | Size of backup file(Byte). | [optional] 

## Methods

### NewBackupFileResultVO

`func NewBackupFileResultVO() *BackupFileResultVO`

NewBackupFileResultVO instantiates a new BackupFileResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupFileResultVOWithDefaults

`func NewBackupFileResultVOWithDefaults() *BackupFileResultVO`

NewBackupFileResultVOWithDefaults instantiates a new BackupFileResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupTime

`func (o *BackupFileResultVO) GetBackupTime() int64`

GetBackupTime returns the BackupTime field if non-nil, zero value otherwise.

### GetBackupTimeOk

`func (o *BackupFileResultVO) GetBackupTimeOk() (*int64, bool)`

GetBackupTimeOk returns a tuple with the BackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTime

`func (o *BackupFileResultVO) SetBackupTime(v int64)`

SetBackupTime sets BackupTime field to given value.

### HasBackupTime

`func (o *BackupFileResultVO) HasBackupTime() bool`

HasBackupTime returns a boolean if a field has been set.

### GetFileName

`func (o *BackupFileResultVO) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BackupFileResultVO) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BackupFileResultVO) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BackupFileResultVO) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetSize

`func (o *BackupFileResultVO) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BackupFileResultVO) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BackupFileResultVO) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BackupFileResultVO) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


