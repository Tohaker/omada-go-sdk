# DeviceOuiModeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** | Device MAC. E.g. AA-BB-CC-DD-11-22 . When \&quot;oldFirmwareDevice\&quot; is true, deivce should configure in only one OUI Based VLAN rule. | 
**LagList** | Pointer to **[]int32** | Device lag list. | [optional] 
**PortList** | Pointer to **[]int32** | Device port list. | [optional] 

## Methods

### NewDeviceOuiModeOpenApiVO

`func NewDeviceOuiModeOpenApiVO(deviceMac string, ) *DeviceOuiModeOpenApiVO`

NewDeviceOuiModeOpenApiVO instantiates a new DeviceOuiModeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOuiModeOpenApiVOWithDefaults

`func NewDeviceOuiModeOpenApiVOWithDefaults() *DeviceOuiModeOpenApiVO`

NewDeviceOuiModeOpenApiVOWithDefaults instantiates a new DeviceOuiModeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *DeviceOuiModeOpenApiVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *DeviceOuiModeOpenApiVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *DeviceOuiModeOpenApiVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetLagList

`func (o *DeviceOuiModeOpenApiVO) GetLagList() []int32`

GetLagList returns the LagList field if non-nil, zero value otherwise.

### GetLagListOk

`func (o *DeviceOuiModeOpenApiVO) GetLagListOk() (*[]int32, bool)`

GetLagListOk returns a tuple with the LagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagList

`func (o *DeviceOuiModeOpenApiVO) SetLagList(v []int32)`

SetLagList sets LagList field to given value.

### HasLagList

`func (o *DeviceOuiModeOpenApiVO) HasLagList() bool`

HasLagList returns a boolean if a field has been set.

### GetPortList

`func (o *DeviceOuiModeOpenApiVO) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *DeviceOuiModeOpenApiVO) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *DeviceOuiModeOpenApiVO) SetPortList(v []int32)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *DeviceOuiModeOpenApiVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


