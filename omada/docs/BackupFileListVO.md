# BackupFileListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileList** | Pointer to [**[]BackupFileResultVO**](BackupFileResultVO.md) | File list of backup files. | [optional] 

## Methods

### NewBackupFileListVO

`func NewBackupFileListVO() *BackupFileListVO`

NewBackupFileListVO instantiates a new BackupFileListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupFileListVOWithDefaults

`func NewBackupFileListVOWithDefaults() *BackupFileListVO`

NewBackupFileListVOWithDefaults instantiates a new BackupFileListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileList

`func (o *BackupFileListVO) GetFileList() []BackupFileResultVO`

GetFileList returns the FileList field if non-nil, zero value otherwise.

### GetFileListOk

`func (o *BackupFileListVO) GetFileListOk() (*[]BackupFileResultVO, bool)`

GetFileListOk returns a tuple with the FileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileList

`func (o *BackupFileListVO) SetFileList(v []BackupFileResultVO)`

SetFileList sets FileList field to given value.

### HasFileList

`func (o *BackupFileListVO) HasFileList() bool`

HasFileList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


