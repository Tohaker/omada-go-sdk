# DeviceCliVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | Pointer to **string** | Device mac | [optional] 
**DeviceModel** | Pointer to **string** | Model of device | [optional] 
**DeviceName** | Pointer to **string** | Device name | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device | [optional] 
**VariableMap** | Pointer to **map[string]string** | The values of different user-defined variables. | [optional] 

## Methods

### NewDeviceCliVO

`func NewDeviceCliVO() *DeviceCliVO`

NewDeviceCliVO instantiates a new DeviceCliVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCliVOWithDefaults

`func NewDeviceCliVOWithDefaults() *DeviceCliVO`

NewDeviceCliVOWithDefaults instantiates a new DeviceCliVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *DeviceCliVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *DeviceCliVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *DeviceCliVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *DeviceCliVO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceModel

`func (o *DeviceCliVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *DeviceCliVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *DeviceCliVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *DeviceCliVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceName

`func (o *DeviceCliVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceCliVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceCliVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceCliVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceCliVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceCliVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceCliVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceCliVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetVariableMap

`func (o *DeviceCliVO) GetVariableMap() map[string]string`

GetVariableMap returns the VariableMap field if non-nil, zero value otherwise.

### GetVariableMapOk

`func (o *DeviceCliVO) GetVariableMapOk() (*map[string]string, bool)`

GetVariableMapOk returns a tuple with the VariableMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMap

`func (o *DeviceCliVO) SetVariableMap(v map[string]string)`

SetVariableMap sets VariableMap field to given value.

### HasVariableMap

`func (o *DeviceCliVO) HasVariableMap() bool`

HasVariableMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


