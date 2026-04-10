# VlanNetworkDeviceSupportInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicePortSupportVlan** | Pointer to [**map[string]VlanNetworkDeviceSupportVO**](VlanNetworkDeviceSupportVO.md) | The map key is device mac. The map value is the device info | [optional] 

## Methods

### NewVlanNetworkDeviceSupportInfoVO

`func NewVlanNetworkDeviceSupportInfoVO() *VlanNetworkDeviceSupportInfoVO`

NewVlanNetworkDeviceSupportInfoVO instantiates a new VlanNetworkDeviceSupportInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkDeviceSupportInfoVOWithDefaults

`func NewVlanNetworkDeviceSupportInfoVOWithDefaults() *VlanNetworkDeviceSupportInfoVO`

NewVlanNetworkDeviceSupportInfoVOWithDefaults instantiates a new VlanNetworkDeviceSupportInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicePortSupportVlan

`func (o *VlanNetworkDeviceSupportInfoVO) GetDevicePortSupportVlan() map[string]VlanNetworkDeviceSupportVO`

GetDevicePortSupportVlan returns the DevicePortSupportVlan field if non-nil, zero value otherwise.

### GetDevicePortSupportVlanOk

`func (o *VlanNetworkDeviceSupportInfoVO) GetDevicePortSupportVlanOk() (*map[string]VlanNetworkDeviceSupportVO, bool)`

GetDevicePortSupportVlanOk returns a tuple with the DevicePortSupportVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortSupportVlan

`func (o *VlanNetworkDeviceSupportInfoVO) SetDevicePortSupportVlan(v map[string]VlanNetworkDeviceSupportVO)`

SetDevicePortSupportVlan sets DevicePortSupportVlan field to given value.

### HasDevicePortSupportVlan

`func (o *VlanNetworkDeviceSupportInfoVO) HasDevicePortSupportVlan() bool`

HasDevicePortSupportVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


