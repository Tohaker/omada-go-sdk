# SelectPortBindingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignIpDeviceMac** | Pointer to **string** | Assign Ip Device Mac. | [optional] 
**AssignIpDeviceType** | **int32** | Assign Ip Device Type. It should be a value as follows: 1: gateway 2: switch 0: third-party | 
**AssignIpStackId** | Pointer to **string** | Assign Ip Stack Id, only valid when assigned ip device is a stack. | [optional] 
**DeviceList** | Pointer to [**[]PortBindingVO**](PortBindingVO.md) | Device List | [optional] 
**FlowControlEnable** | Pointer to **bool** | Enable Flow Control, only valid when creating network. | [optional] 
**InternetPorts** | Pointer to **[]string** | Internet Ports | [optional] 
**PortIsolationEnable** | Pointer to **bool** | Enable Port Isolation, only valid when creating network. | [optional] 
**TagIds** | Pointer to **[]string** | Tag ID List | [optional] 
**Vlan** | Pointer to **int32** | vlan, only valid when vlanType is 0 (single vlan) | [optional] 
**VlanType** | **int32** | Network type, It should be a value as follows : 0:single vlan 1:multi vlan | 
**Vlans** | Pointer to **string** | vlans, only valid when vlanType is 1 (multi vlan). VLAN format: 200, 1-100. | [optional] 

## Methods

### NewSelectPortBindingVO

`func NewSelectPortBindingVO(assignIpDeviceType int32, vlanType int32, ) *SelectPortBindingVO`

NewSelectPortBindingVO instantiates a new SelectPortBindingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectPortBindingVOWithDefaults

`func NewSelectPortBindingVOWithDefaults() *SelectPortBindingVO`

NewSelectPortBindingVOWithDefaults instantiates a new SelectPortBindingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignIpDeviceMac

`func (o *SelectPortBindingVO) GetAssignIpDeviceMac() string`

GetAssignIpDeviceMac returns the AssignIpDeviceMac field if non-nil, zero value otherwise.

### GetAssignIpDeviceMacOk

`func (o *SelectPortBindingVO) GetAssignIpDeviceMacOk() (*string, bool)`

GetAssignIpDeviceMacOk returns a tuple with the AssignIpDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignIpDeviceMac

`func (o *SelectPortBindingVO) SetAssignIpDeviceMac(v string)`

SetAssignIpDeviceMac sets AssignIpDeviceMac field to given value.

### HasAssignIpDeviceMac

`func (o *SelectPortBindingVO) HasAssignIpDeviceMac() bool`

HasAssignIpDeviceMac returns a boolean if a field has been set.

### GetAssignIpDeviceType

`func (o *SelectPortBindingVO) GetAssignIpDeviceType() int32`

GetAssignIpDeviceType returns the AssignIpDeviceType field if non-nil, zero value otherwise.

### GetAssignIpDeviceTypeOk

`func (o *SelectPortBindingVO) GetAssignIpDeviceTypeOk() (*int32, bool)`

GetAssignIpDeviceTypeOk returns a tuple with the AssignIpDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignIpDeviceType

`func (o *SelectPortBindingVO) SetAssignIpDeviceType(v int32)`

SetAssignIpDeviceType sets AssignIpDeviceType field to given value.


### GetAssignIpStackId

`func (o *SelectPortBindingVO) GetAssignIpStackId() string`

GetAssignIpStackId returns the AssignIpStackId field if non-nil, zero value otherwise.

### GetAssignIpStackIdOk

`func (o *SelectPortBindingVO) GetAssignIpStackIdOk() (*string, bool)`

GetAssignIpStackIdOk returns a tuple with the AssignIpStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignIpStackId

`func (o *SelectPortBindingVO) SetAssignIpStackId(v string)`

SetAssignIpStackId sets AssignIpStackId field to given value.

### HasAssignIpStackId

`func (o *SelectPortBindingVO) HasAssignIpStackId() bool`

HasAssignIpStackId returns a boolean if a field has been set.

### GetDeviceList

`func (o *SelectPortBindingVO) GetDeviceList() []PortBindingVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *SelectPortBindingVO) GetDeviceListOk() (*[]PortBindingVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *SelectPortBindingVO) SetDeviceList(v []PortBindingVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *SelectPortBindingVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *SelectPortBindingVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *SelectPortBindingVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *SelectPortBindingVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *SelectPortBindingVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetInternetPorts

`func (o *SelectPortBindingVO) GetInternetPorts() []string`

GetInternetPorts returns the InternetPorts field if non-nil, zero value otherwise.

### GetInternetPortsOk

`func (o *SelectPortBindingVO) GetInternetPortsOk() (*[]string, bool)`

GetInternetPortsOk returns a tuple with the InternetPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetPorts

`func (o *SelectPortBindingVO) SetInternetPorts(v []string)`

SetInternetPorts sets InternetPorts field to given value.

### HasInternetPorts

`func (o *SelectPortBindingVO) HasInternetPorts() bool`

HasInternetPorts returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *SelectPortBindingVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *SelectPortBindingVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *SelectPortBindingVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *SelectPortBindingVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetTagIds

`func (o *SelectPortBindingVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *SelectPortBindingVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *SelectPortBindingVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *SelectPortBindingVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetVlan

`func (o *SelectPortBindingVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *SelectPortBindingVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *SelectPortBindingVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *SelectPortBindingVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *SelectPortBindingVO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *SelectPortBindingVO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *SelectPortBindingVO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.


### GetVlans

`func (o *SelectPortBindingVO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *SelectPortBindingVO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *SelectPortBindingVO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *SelectPortBindingVO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


