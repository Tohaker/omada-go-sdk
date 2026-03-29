# BatchBindDeviceResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedDevices** | Pointer to [**[]DeviceBindResultOpenApiVO**](DeviceBindResultOpenApiVO.md) | Devices with failed operation. | [optional] 
**SuccessedDevices** | Pointer to [**[]DeviceBindResultOpenApiVO**](DeviceBindResultOpenApiVO.md) | Devices with successful operation. | [optional] 

## Methods

### NewBatchBindDeviceResultOpenApiVO

`func NewBatchBindDeviceResultOpenApiVO() *BatchBindDeviceResultOpenApiVO`

NewBatchBindDeviceResultOpenApiVO instantiates a new BatchBindDeviceResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchBindDeviceResultOpenApiVOWithDefaults

`func NewBatchBindDeviceResultOpenApiVOWithDefaults() *BatchBindDeviceResultOpenApiVO`

NewBatchBindDeviceResultOpenApiVOWithDefaults instantiates a new BatchBindDeviceResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedDevices

`func (o *BatchBindDeviceResultOpenApiVO) GetFailedDevices() []DeviceBindResultOpenApiVO`

GetFailedDevices returns the FailedDevices field if non-nil, zero value otherwise.

### GetFailedDevicesOk

`func (o *BatchBindDeviceResultOpenApiVO) GetFailedDevicesOk() (*[]DeviceBindResultOpenApiVO, bool)`

GetFailedDevicesOk returns a tuple with the FailedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDevices

`func (o *BatchBindDeviceResultOpenApiVO) SetFailedDevices(v []DeviceBindResultOpenApiVO)`

SetFailedDevices sets FailedDevices field to given value.

### HasFailedDevices

`func (o *BatchBindDeviceResultOpenApiVO) HasFailedDevices() bool`

HasFailedDevices returns a boolean if a field has been set.

### GetSuccessedDevices

`func (o *BatchBindDeviceResultOpenApiVO) GetSuccessedDevices() []DeviceBindResultOpenApiVO`

GetSuccessedDevices returns the SuccessedDevices field if non-nil, zero value otherwise.

### GetSuccessedDevicesOk

`func (o *BatchBindDeviceResultOpenApiVO) GetSuccessedDevicesOk() (*[]DeviceBindResultOpenApiVO, bool)`

GetSuccessedDevicesOk returns a tuple with the SuccessedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessedDevices

`func (o *BatchBindDeviceResultOpenApiVO) SetSuccessedDevices(v []DeviceBindResultOpenApiVO)`

SetSuccessedDevices sets SuccessedDevices field to given value.

### HasSuccessedDevices

`func (o *BatchBindDeviceResultOpenApiVO) HasSuccessedDevices() bool`

HasSuccessedDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


