# BatchFullChannelDetectApListOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceInfoList** | Pointer to [**[]ApInfoOpenApiVO**](ApInfoOpenApiVO.md) | A list of information for AP devices. | [optional] 
**ScanApNum** | Pointer to **int32** | The number of AP devices participating in batch full channel detection. | [optional] 

## Methods

### NewBatchFullChannelDetectApListOpenApiVO

`func NewBatchFullChannelDetectApListOpenApiVO() *BatchFullChannelDetectApListOpenApiVO`

NewBatchFullChannelDetectApListOpenApiVO instantiates a new BatchFullChannelDetectApListOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchFullChannelDetectApListOpenApiVOWithDefaults

`func NewBatchFullChannelDetectApListOpenApiVOWithDefaults() *BatchFullChannelDetectApListOpenApiVO`

NewBatchFullChannelDetectApListOpenApiVOWithDefaults instantiates a new BatchFullChannelDetectApListOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceInfoList

`func (o *BatchFullChannelDetectApListOpenApiVO) GetDeviceInfoList() []ApInfoOpenApiVO`

GetDeviceInfoList returns the DeviceInfoList field if non-nil, zero value otherwise.

### GetDeviceInfoListOk

`func (o *BatchFullChannelDetectApListOpenApiVO) GetDeviceInfoListOk() (*[]ApInfoOpenApiVO, bool)`

GetDeviceInfoListOk returns a tuple with the DeviceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfoList

`func (o *BatchFullChannelDetectApListOpenApiVO) SetDeviceInfoList(v []ApInfoOpenApiVO)`

SetDeviceInfoList sets DeviceInfoList field to given value.

### HasDeviceInfoList

`func (o *BatchFullChannelDetectApListOpenApiVO) HasDeviceInfoList() bool`

HasDeviceInfoList returns a boolean if a field has been set.

### GetScanApNum

`func (o *BatchFullChannelDetectApListOpenApiVO) GetScanApNum() int32`

GetScanApNum returns the ScanApNum field if non-nil, zero value otherwise.

### GetScanApNumOk

`func (o *BatchFullChannelDetectApListOpenApiVO) GetScanApNumOk() (*int32, bool)`

GetScanApNumOk returns a tuple with the ScanApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanApNum

`func (o *BatchFullChannelDetectApListOpenApiVO) SetScanApNum(v int32)`

SetScanApNum sets ScanApNum field to given value.

### HasScanApNum

`func (o *BatchFullChannelDetectApListOpenApiVO) HasScanApNum() bool`

HasScanApNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


