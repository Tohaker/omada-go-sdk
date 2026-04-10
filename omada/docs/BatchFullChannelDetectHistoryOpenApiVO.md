# BatchFullChannelDetectHistoryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceInfoList** | Pointer to [**[]ApInfoOpenApiVO**](ApInfoOpenApiVO.md) | Device information list | [optional] 
**EnableInterference** | Pointer to **bool** | Whether to enable interference detect | [optional] 
**HistoryId** | Pointer to **string** | Id of batch full channel detect history. | [optional] 
**LastSeen** | Pointer to **int64** | Last full channel detect time. | [optional] 
**ScanApNum** | Pointer to **int32** | The total number of APs participating in the full channel detect. | [optional] 
**Status** | Pointer to **int32** | full channel detect status. 0: finish, 1: scanning. | [optional] 

## Methods

### NewBatchFullChannelDetectHistoryOpenApiVO

`func NewBatchFullChannelDetectHistoryOpenApiVO() *BatchFullChannelDetectHistoryOpenApiVO`

NewBatchFullChannelDetectHistoryOpenApiVO instantiates a new BatchFullChannelDetectHistoryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchFullChannelDetectHistoryOpenApiVOWithDefaults

`func NewBatchFullChannelDetectHistoryOpenApiVOWithDefaults() *BatchFullChannelDetectHistoryOpenApiVO`

NewBatchFullChannelDetectHistoryOpenApiVOWithDefaults instantiates a new BatchFullChannelDetectHistoryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceInfoList

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetDeviceInfoList() []ApInfoOpenApiVO`

GetDeviceInfoList returns the DeviceInfoList field if non-nil, zero value otherwise.

### GetDeviceInfoListOk

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetDeviceInfoListOk() (*[]ApInfoOpenApiVO, bool)`

GetDeviceInfoListOk returns a tuple with the DeviceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfoList

`func (o *BatchFullChannelDetectHistoryOpenApiVO) SetDeviceInfoList(v []ApInfoOpenApiVO)`

SetDeviceInfoList sets DeviceInfoList field to given value.

### HasDeviceInfoList

`func (o *BatchFullChannelDetectHistoryOpenApiVO) HasDeviceInfoList() bool`

HasDeviceInfoList returns a boolean if a field has been set.

### GetEnableInterference

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetEnableInterference() bool`

GetEnableInterference returns the EnableInterference field if non-nil, zero value otherwise.

### GetEnableInterferenceOk

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetEnableInterferenceOk() (*bool, bool)`

GetEnableInterferenceOk returns a tuple with the EnableInterference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterference

`func (o *BatchFullChannelDetectHistoryOpenApiVO) SetEnableInterference(v bool)`

SetEnableInterference sets EnableInterference field to given value.

### HasEnableInterference

`func (o *BatchFullChannelDetectHistoryOpenApiVO) HasEnableInterference() bool`

HasEnableInterference returns a boolean if a field has been set.

### GetHistoryId

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetHistoryId() string`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetHistoryIdOk() (*string, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *BatchFullChannelDetectHistoryOpenApiVO) SetHistoryId(v string)`

SetHistoryId sets HistoryId field to given value.

### HasHistoryId

`func (o *BatchFullChannelDetectHistoryOpenApiVO) HasHistoryId() bool`

HasHistoryId returns a boolean if a field has been set.

### GetLastSeen

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *BatchFullChannelDetectHistoryOpenApiVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *BatchFullChannelDetectHistoryOpenApiVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetScanApNum

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetScanApNum() int32`

GetScanApNum returns the ScanApNum field if non-nil, zero value otherwise.

### GetScanApNumOk

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetScanApNumOk() (*int32, bool)`

GetScanApNumOk returns a tuple with the ScanApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanApNum

`func (o *BatchFullChannelDetectHistoryOpenApiVO) SetScanApNum(v int32)`

SetScanApNum sets ScanApNum field to given value.

### HasScanApNum

`func (o *BatchFullChannelDetectHistoryOpenApiVO) HasScanApNum() bool`

HasScanApNum returns a boolean if a field has been set.

### GetStatus

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchFullChannelDetectHistoryOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchFullChannelDetectHistoryOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


