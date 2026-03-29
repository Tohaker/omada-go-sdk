# DeviceExportCliVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** | Device mac | 
**VariableMap** | Pointer to **map[string]string** | The values of different user-defined variables. | [optional] 

## Methods

### NewDeviceExportCliVO

`func NewDeviceExportCliVO(deviceMac string, ) *DeviceExportCliVO`

NewDeviceExportCliVO instantiates a new DeviceExportCliVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceExportCliVOWithDefaults

`func NewDeviceExportCliVOWithDefaults() *DeviceExportCliVO`

NewDeviceExportCliVOWithDefaults instantiates a new DeviceExportCliVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *DeviceExportCliVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *DeviceExportCliVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *DeviceExportCliVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetVariableMap

`func (o *DeviceExportCliVO) GetVariableMap() map[string]string`

GetVariableMap returns the VariableMap field if non-nil, zero value otherwise.

### GetVariableMapOk

`func (o *DeviceExportCliVO) GetVariableMapOk() (*map[string]string, bool)`

GetVariableMapOk returns a tuple with the VariableMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMap

`func (o *DeviceExportCliVO) SetVariableMap(v map[string]string)`

SetVariableMap sets VariableMap field to given value.

### HasVariableMap

`func (o *DeviceExportCliVO) HasVariableMap() bool`

HasVariableMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


