# VirtualWanIpv4SettingConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Dhcp** | Pointer to [**VirtualWanIpv4DhcpOpenApiVO**](VirtualWanIpv4DhcpOpenApiVO.md) |  | [optional] 
**Ipv4Ipoa** | Pointer to [**VirtualWanIpv4IpoaOpenApiVO**](VirtualWanIpv4IpoaOpenApiVO.md) |  | [optional] 
**Ipv4Pppoa** | Pointer to [**VirtualWanIpv4PppoaOpenApiVO**](VirtualWanIpv4PppoaOpenApiVO.md) |  | [optional] 
**Ipv4Pppoe** | Pointer to [**VirtualWanIpv4PppoeOpenApiVO**](VirtualWanIpv4PppoeOpenApiVO.md) |  | [optional] 
**Ipv4Static** | Pointer to [**VirtualWanIpv4StaticOpenApiVO**](VirtualWanIpv4StaticOpenApiVO.md) |  | [optional] 
**Proto** | **string** | Virtual WAN IPv4 proto type, use static, dhcp, pppoe. | 
**QosTag** | Pointer to **int32** | Qos Tag. Parameter [qosTag] should between 0 and 7. | [optional] 
**QosTagEnable** | Pointer to **bool** | Whether to enable 802.1Q Tag. | [optional] 
**VlanId** | **int32** | Vlan ID. Parameter [vlanId] should between 1 and 4094. | 
**VlanPriority** | Pointer to **int32** | Vlan Priority. It takes effect when [vlanId] is not 0, and it should be within the range of 0–7. | [optional] 

## Methods

### NewVirtualWanIpv4SettingConfigOpenApiVO

`func NewVirtualWanIpv4SettingConfigOpenApiVO(proto string, vlanId int32, ) *VirtualWanIpv4SettingConfigOpenApiVO`

NewVirtualWanIpv4SettingConfigOpenApiVO instantiates a new VirtualWanIpv4SettingConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanIpv4SettingConfigOpenApiVOWithDefaults

`func NewVirtualWanIpv4SettingConfigOpenApiVOWithDefaults() *VirtualWanIpv4SettingConfigOpenApiVO`

NewVirtualWanIpv4SettingConfigOpenApiVOWithDefaults instantiates a new VirtualWanIpv4SettingConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Dhcp

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4Dhcp() VirtualWanIpv4DhcpOpenApiVO`

GetIpv4Dhcp returns the Ipv4Dhcp field if non-nil, zero value otherwise.

### GetIpv4DhcpOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4DhcpOk() (*VirtualWanIpv4DhcpOpenApiVO, bool)`

GetIpv4DhcpOk returns a tuple with the Ipv4Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Dhcp

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetIpv4Dhcp(v VirtualWanIpv4DhcpOpenApiVO)`

SetIpv4Dhcp sets Ipv4Dhcp field to given value.

### HasIpv4Dhcp

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) HasIpv4Dhcp() bool`

HasIpv4Dhcp returns a boolean if a field has been set.

### GetIpv4Ipoa

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4Ipoa() VirtualWanIpv4IpoaOpenApiVO`

GetIpv4Ipoa returns the Ipv4Ipoa field if non-nil, zero value otherwise.

### GetIpv4IpoaOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4IpoaOk() (*VirtualWanIpv4IpoaOpenApiVO, bool)`

GetIpv4IpoaOk returns a tuple with the Ipv4Ipoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Ipoa

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetIpv4Ipoa(v VirtualWanIpv4IpoaOpenApiVO)`

SetIpv4Ipoa sets Ipv4Ipoa field to given value.

### HasIpv4Ipoa

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) HasIpv4Ipoa() bool`

HasIpv4Ipoa returns a boolean if a field has been set.

### GetIpv4Pppoa

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4Pppoa() VirtualWanIpv4PppoaOpenApiVO`

GetIpv4Pppoa returns the Ipv4Pppoa field if non-nil, zero value otherwise.

### GetIpv4PppoaOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4PppoaOk() (*VirtualWanIpv4PppoaOpenApiVO, bool)`

GetIpv4PppoaOk returns a tuple with the Ipv4Pppoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Pppoa

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetIpv4Pppoa(v VirtualWanIpv4PppoaOpenApiVO)`

SetIpv4Pppoa sets Ipv4Pppoa field to given value.

### HasIpv4Pppoa

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) HasIpv4Pppoa() bool`

HasIpv4Pppoa returns a boolean if a field has been set.

### GetIpv4Pppoe

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4Pppoe() VirtualWanIpv4PppoeOpenApiVO`

GetIpv4Pppoe returns the Ipv4Pppoe field if non-nil, zero value otherwise.

### GetIpv4PppoeOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4PppoeOk() (*VirtualWanIpv4PppoeOpenApiVO, bool)`

GetIpv4PppoeOk returns a tuple with the Ipv4Pppoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Pppoe

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetIpv4Pppoe(v VirtualWanIpv4PppoeOpenApiVO)`

SetIpv4Pppoe sets Ipv4Pppoe field to given value.

### HasIpv4Pppoe

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) HasIpv4Pppoe() bool`

HasIpv4Pppoe returns a boolean if a field has been set.

### GetIpv4Static

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4Static() VirtualWanIpv4StaticOpenApiVO`

GetIpv4Static returns the Ipv4Static field if non-nil, zero value otherwise.

### GetIpv4StaticOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetIpv4StaticOk() (*VirtualWanIpv4StaticOpenApiVO, bool)`

GetIpv4StaticOk returns a tuple with the Ipv4Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Static

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetIpv4Static(v VirtualWanIpv4StaticOpenApiVO)`

SetIpv4Static sets Ipv4Static field to given value.

### HasIpv4Static

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) HasIpv4Static() bool`

HasIpv4Static returns a boolean if a field has been set.

### GetProto

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetProto(v string)`

SetProto sets Proto field to given value.


### GetQosTag

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetQosTag() int32`

GetQosTag returns the QosTag field if non-nil, zero value otherwise.

### GetQosTagOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetQosTagOk() (*int32, bool)`

GetQosTagOk returns a tuple with the QosTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosTag

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetQosTag(v int32)`

SetQosTag sets QosTag field to given value.

### HasQosTag

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) HasQosTag() bool`

HasQosTag returns a boolean if a field has been set.

### GetQosTagEnable

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetQosTagEnable() bool`

GetQosTagEnable returns the QosTagEnable field if non-nil, zero value otherwise.

### GetQosTagEnableOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetQosTagEnableOk() (*bool, bool)`

GetQosTagEnableOk returns a tuple with the QosTagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosTagEnable

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetQosTagEnable(v bool)`

SetQosTagEnable sets QosTagEnable field to given value.

### HasQosTagEnable

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) HasQosTagEnable() bool`

HasQosTagEnable returns a boolean if a field has been set.

### GetVlanId

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetVlanPriority

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetVlanPriority() int32`

GetVlanPriority returns the VlanPriority field if non-nil, zero value otherwise.

### GetVlanPriorityOk

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) GetVlanPriorityOk() (*int32, bool)`

GetVlanPriorityOk returns a tuple with the VlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPriority

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) SetVlanPriority(v int32)`

SetVlanPriority sets VlanPriority field to given value.

### HasVlanPriority

`func (o *VirtualWanIpv4SettingConfigOpenApiVO) HasVlanPriority() bool`

HasVlanPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


