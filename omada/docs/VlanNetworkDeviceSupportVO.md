# VlanNetworkDeviceSupportVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device mac | [optional] 
**PortList** | Pointer to [**[]VlanNetworkDevicePortSupportVO**](VlanNetworkDevicePortSupportVO.md) | Port list | [optional] 
**StackId** | Pointer to **string** | Stack ID, only valid when the device is stack | [optional] 
**Type** | Pointer to **string** | Device type.It should be a value as follows: gateway, switch, ap, olt | [optional] 

## Methods

### NewVlanNetworkDeviceSupportVO

`func NewVlanNetworkDeviceSupportVO() *VlanNetworkDeviceSupportVO`

NewVlanNetworkDeviceSupportVO instantiates a new VlanNetworkDeviceSupportVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkDeviceSupportVOWithDefaults

`func NewVlanNetworkDeviceSupportVOWithDefaults() *VlanNetworkDeviceSupportVO`

NewVlanNetworkDeviceSupportVOWithDefaults instantiates a new VlanNetworkDeviceSupportVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *VlanNetworkDeviceSupportVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VlanNetworkDeviceSupportVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VlanNetworkDeviceSupportVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VlanNetworkDeviceSupportVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPortList

`func (o *VlanNetworkDeviceSupportVO) GetPortList() []VlanNetworkDevicePortSupportVO`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *VlanNetworkDeviceSupportVO) GetPortListOk() (*[]VlanNetworkDevicePortSupportVO, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *VlanNetworkDeviceSupportVO) SetPortList(v []VlanNetworkDevicePortSupportVO)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *VlanNetworkDeviceSupportVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetStackId

`func (o *VlanNetworkDeviceSupportVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *VlanNetworkDeviceSupportVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *VlanNetworkDeviceSupportVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *VlanNetworkDeviceSupportVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetType

`func (o *VlanNetworkDeviceSupportVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VlanNetworkDeviceSupportVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VlanNetworkDeviceSupportVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VlanNetworkDeviceSupportVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


