# DeviceOuiModeQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** | Device MAC. E.g. AA-BB-CC-DD-11-22 . | 
**DeviceName** | Pointer to **string** | Device Name | [optional] 
**LagInfo** | Pointer to [**[]LagInfoOpenApiVO**](LagInfoOpenApiVO.md) | Switch lag info. | [optional] 
**LagList** | Pointer to **[]int32** | Configured Switch lag. | [optional] 
**PortList** | Pointer to **[]int32** | Configured Switch port. | [optional] 

## Methods

### NewDeviceOuiModeQueryOpenApiVO

`func NewDeviceOuiModeQueryOpenApiVO(deviceMac string, ) *DeviceOuiModeQueryOpenApiVO`

NewDeviceOuiModeQueryOpenApiVO instantiates a new DeviceOuiModeQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOuiModeQueryOpenApiVOWithDefaults

`func NewDeviceOuiModeQueryOpenApiVOWithDefaults() *DeviceOuiModeQueryOpenApiVO`

NewDeviceOuiModeQueryOpenApiVOWithDefaults instantiates a new DeviceOuiModeQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *DeviceOuiModeQueryOpenApiVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *DeviceOuiModeQueryOpenApiVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *DeviceOuiModeQueryOpenApiVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetDeviceName

`func (o *DeviceOuiModeQueryOpenApiVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceOuiModeQueryOpenApiVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceOuiModeQueryOpenApiVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceOuiModeQueryOpenApiVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetLagInfo

`func (o *DeviceOuiModeQueryOpenApiVO) GetLagInfo() []LagInfoOpenApiVO`

GetLagInfo returns the LagInfo field if non-nil, zero value otherwise.

### GetLagInfoOk

`func (o *DeviceOuiModeQueryOpenApiVO) GetLagInfoOk() (*[]LagInfoOpenApiVO, bool)`

GetLagInfoOk returns a tuple with the LagInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagInfo

`func (o *DeviceOuiModeQueryOpenApiVO) SetLagInfo(v []LagInfoOpenApiVO)`

SetLagInfo sets LagInfo field to given value.

### HasLagInfo

`func (o *DeviceOuiModeQueryOpenApiVO) HasLagInfo() bool`

HasLagInfo returns a boolean if a field has been set.

### GetLagList

`func (o *DeviceOuiModeQueryOpenApiVO) GetLagList() []int32`

GetLagList returns the LagList field if non-nil, zero value otherwise.

### GetLagListOk

`func (o *DeviceOuiModeQueryOpenApiVO) GetLagListOk() (*[]int32, bool)`

GetLagListOk returns a tuple with the LagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagList

`func (o *DeviceOuiModeQueryOpenApiVO) SetLagList(v []int32)`

SetLagList sets LagList field to given value.

### HasLagList

`func (o *DeviceOuiModeQueryOpenApiVO) HasLagList() bool`

HasLagList returns a boolean if a field has been set.

### GetPortList

`func (o *DeviceOuiModeQueryOpenApiVO) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *DeviceOuiModeQueryOpenApiVO) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *DeviceOuiModeQueryOpenApiVO) SetPortList(v []int32)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *DeviceOuiModeQueryOpenApiVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


