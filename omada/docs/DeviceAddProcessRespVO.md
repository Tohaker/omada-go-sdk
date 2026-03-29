# DeviceAddProcessRespVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]SnAddResultVO**](SnAddResultVO.md) | Devices add result | [optional] 
**ProcessStatus** | Pointer to **int32** | It should be a value as follows: 0: init; 1: doing; 2: done | [optional] 

## Methods

### NewDeviceAddProcessRespVO

`func NewDeviceAddProcessRespVO() *DeviceAddProcessRespVO`

NewDeviceAddProcessRespVO instantiates a new DeviceAddProcessRespVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAddProcessRespVOWithDefaults

`func NewDeviceAddProcessRespVOWithDefaults() *DeviceAddProcessRespVO`

NewDeviceAddProcessRespVOWithDefaults instantiates a new DeviceAddProcessRespVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *DeviceAddProcessRespVO) GetDevices() []SnAddResultVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DeviceAddProcessRespVO) GetDevicesOk() (*[]SnAddResultVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DeviceAddProcessRespVO) SetDevices(v []SnAddResultVO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *DeviceAddProcessRespVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetProcessStatus

`func (o *DeviceAddProcessRespVO) GetProcessStatus() int32`

GetProcessStatus returns the ProcessStatus field if non-nil, zero value otherwise.

### GetProcessStatusOk

`func (o *DeviceAddProcessRespVO) GetProcessStatusOk() (*int32, bool)`

GetProcessStatusOk returns a tuple with the ProcessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessStatus

`func (o *DeviceAddProcessRespVO) SetProcessStatus(v int32)`

SetProcessStatus sets ProcessStatus field to given value.

### HasProcessStatus

`func (o *DeviceAddProcessRespVO) HasProcessStatus() bool`

HasProcessStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


