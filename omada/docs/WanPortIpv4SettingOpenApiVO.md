# WanPortIpv4SettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Dhcp** | Pointer to [**Ipv4DhcpOpenApiVO**](Ipv4DhcpOpenApiVO.md) |  | [optional] 
**Ipv4Ipoa** | Pointer to [**Ipv4IpoaOpenApiVO**](Ipv4IpoaOpenApiVO.md) |  | [optional] 
**Ipv4L2tp** | Pointer to [**Ipv4L2tpOpenApiVO**](Ipv4L2tpOpenApiVO.md) |  | [optional] 
**Ipv4Pppoa** | Pointer to [**Ipv4PppoaOpenApiVO**](Ipv4PppoaOpenApiVO.md) |  | [optional] 
**Ipv4Pppoe** | Pointer to [**Ipv4PppoeOpenApiVO**](Ipv4PppoeOpenApiVO.md) |  | [optional] 
**Ipv4Pptp** | Pointer to [**Ipv4PptpOpenApiVO**](Ipv4PptpOpenApiVO.md) |  | [optional] 
**Ipv4Static** | Pointer to [**Ipv4StaticOpenApiVO**](Ipv4StaticOpenApiVO.md) |  | [optional] 
**ProtoType** | **int32** | IPv4 connection type should be a value as follows: 0:static; 1:DHCP; 2:PPPoE; 3:L2TP; 4:PPTP. | 
**QosTag** | Pointer to **bool** | 802.1Q Tag. It takes effect when [vlanId] is not 0. | [optional] 
**VlanId** | **int32** | VLAN ID should be within the range of 0–4094, 0 means disable. | 
**VlanPriority** | Pointer to **int32** | Vlan Priority. It takes effect when [vlanId] is not 0, and it should be within the range of 0–7. | [optional] 

## Methods

### NewWanPortIpv4SettingOpenApiVO

`func NewWanPortIpv4SettingOpenApiVO(protoType int32, vlanId int32, ) *WanPortIpv4SettingOpenApiVO`

NewWanPortIpv4SettingOpenApiVO instantiates a new WanPortIpv4SettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanPortIpv4SettingOpenApiVOWithDefaults

`func NewWanPortIpv4SettingOpenApiVOWithDefaults() *WanPortIpv4SettingOpenApiVO`

NewWanPortIpv4SettingOpenApiVOWithDefaults instantiates a new WanPortIpv4SettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Dhcp

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4Dhcp() Ipv4DhcpOpenApiVO`

GetIpv4Dhcp returns the Ipv4Dhcp field if non-nil, zero value otherwise.

### GetIpv4DhcpOk

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4DhcpOk() (*Ipv4DhcpOpenApiVO, bool)`

GetIpv4DhcpOk returns a tuple with the Ipv4Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Dhcp

`func (o *WanPortIpv4SettingOpenApiVO) SetIpv4Dhcp(v Ipv4DhcpOpenApiVO)`

SetIpv4Dhcp sets Ipv4Dhcp field to given value.

### HasIpv4Dhcp

`func (o *WanPortIpv4SettingOpenApiVO) HasIpv4Dhcp() bool`

HasIpv4Dhcp returns a boolean if a field has been set.

### GetIpv4Ipoa

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4Ipoa() Ipv4IpoaOpenApiVO`

GetIpv4Ipoa returns the Ipv4Ipoa field if non-nil, zero value otherwise.

### GetIpv4IpoaOk

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4IpoaOk() (*Ipv4IpoaOpenApiVO, bool)`

GetIpv4IpoaOk returns a tuple with the Ipv4Ipoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Ipoa

`func (o *WanPortIpv4SettingOpenApiVO) SetIpv4Ipoa(v Ipv4IpoaOpenApiVO)`

SetIpv4Ipoa sets Ipv4Ipoa field to given value.

### HasIpv4Ipoa

`func (o *WanPortIpv4SettingOpenApiVO) HasIpv4Ipoa() bool`

HasIpv4Ipoa returns a boolean if a field has been set.

### GetIpv4L2tp

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4L2tp() Ipv4L2tpOpenApiVO`

GetIpv4L2tp returns the Ipv4L2tp field if non-nil, zero value otherwise.

### GetIpv4L2tpOk

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4L2tpOk() (*Ipv4L2tpOpenApiVO, bool)`

GetIpv4L2tpOk returns a tuple with the Ipv4L2tp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4L2tp

`func (o *WanPortIpv4SettingOpenApiVO) SetIpv4L2tp(v Ipv4L2tpOpenApiVO)`

SetIpv4L2tp sets Ipv4L2tp field to given value.

### HasIpv4L2tp

`func (o *WanPortIpv4SettingOpenApiVO) HasIpv4L2tp() bool`

HasIpv4L2tp returns a boolean if a field has been set.

### GetIpv4Pppoa

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4Pppoa() Ipv4PppoaOpenApiVO`

GetIpv4Pppoa returns the Ipv4Pppoa field if non-nil, zero value otherwise.

### GetIpv4PppoaOk

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4PppoaOk() (*Ipv4PppoaOpenApiVO, bool)`

GetIpv4PppoaOk returns a tuple with the Ipv4Pppoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Pppoa

`func (o *WanPortIpv4SettingOpenApiVO) SetIpv4Pppoa(v Ipv4PppoaOpenApiVO)`

SetIpv4Pppoa sets Ipv4Pppoa field to given value.

### HasIpv4Pppoa

`func (o *WanPortIpv4SettingOpenApiVO) HasIpv4Pppoa() bool`

HasIpv4Pppoa returns a boolean if a field has been set.

### GetIpv4Pppoe

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4Pppoe() Ipv4PppoeOpenApiVO`

GetIpv4Pppoe returns the Ipv4Pppoe field if non-nil, zero value otherwise.

### GetIpv4PppoeOk

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4PppoeOk() (*Ipv4PppoeOpenApiVO, bool)`

GetIpv4PppoeOk returns a tuple with the Ipv4Pppoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Pppoe

`func (o *WanPortIpv4SettingOpenApiVO) SetIpv4Pppoe(v Ipv4PppoeOpenApiVO)`

SetIpv4Pppoe sets Ipv4Pppoe field to given value.

### HasIpv4Pppoe

`func (o *WanPortIpv4SettingOpenApiVO) HasIpv4Pppoe() bool`

HasIpv4Pppoe returns a boolean if a field has been set.

### GetIpv4Pptp

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4Pptp() Ipv4PptpOpenApiVO`

GetIpv4Pptp returns the Ipv4Pptp field if non-nil, zero value otherwise.

### GetIpv4PptpOk

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4PptpOk() (*Ipv4PptpOpenApiVO, bool)`

GetIpv4PptpOk returns a tuple with the Ipv4Pptp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Pptp

`func (o *WanPortIpv4SettingOpenApiVO) SetIpv4Pptp(v Ipv4PptpOpenApiVO)`

SetIpv4Pptp sets Ipv4Pptp field to given value.

### HasIpv4Pptp

`func (o *WanPortIpv4SettingOpenApiVO) HasIpv4Pptp() bool`

HasIpv4Pptp returns a boolean if a field has been set.

### GetIpv4Static

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4Static() Ipv4StaticOpenApiVO`

GetIpv4Static returns the Ipv4Static field if non-nil, zero value otherwise.

### GetIpv4StaticOk

`func (o *WanPortIpv4SettingOpenApiVO) GetIpv4StaticOk() (*Ipv4StaticOpenApiVO, bool)`

GetIpv4StaticOk returns a tuple with the Ipv4Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Static

`func (o *WanPortIpv4SettingOpenApiVO) SetIpv4Static(v Ipv4StaticOpenApiVO)`

SetIpv4Static sets Ipv4Static field to given value.

### HasIpv4Static

`func (o *WanPortIpv4SettingOpenApiVO) HasIpv4Static() bool`

HasIpv4Static returns a boolean if a field has been set.

### GetProtoType

`func (o *WanPortIpv4SettingOpenApiVO) GetProtoType() int32`

GetProtoType returns the ProtoType field if non-nil, zero value otherwise.

### GetProtoTypeOk

`func (o *WanPortIpv4SettingOpenApiVO) GetProtoTypeOk() (*int32, bool)`

GetProtoTypeOk returns a tuple with the ProtoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoType

`func (o *WanPortIpv4SettingOpenApiVO) SetProtoType(v int32)`

SetProtoType sets ProtoType field to given value.


### GetQosTag

`func (o *WanPortIpv4SettingOpenApiVO) GetQosTag() bool`

GetQosTag returns the QosTag field if non-nil, zero value otherwise.

### GetQosTagOk

`func (o *WanPortIpv4SettingOpenApiVO) GetQosTagOk() (*bool, bool)`

GetQosTagOk returns a tuple with the QosTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosTag

`func (o *WanPortIpv4SettingOpenApiVO) SetQosTag(v bool)`

SetQosTag sets QosTag field to given value.

### HasQosTag

`func (o *WanPortIpv4SettingOpenApiVO) HasQosTag() bool`

HasQosTag returns a boolean if a field has been set.

### GetVlanId

`func (o *WanPortIpv4SettingOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *WanPortIpv4SettingOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *WanPortIpv4SettingOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetVlanPriority

`func (o *WanPortIpv4SettingOpenApiVO) GetVlanPriority() int32`

GetVlanPriority returns the VlanPriority field if non-nil, zero value otherwise.

### GetVlanPriorityOk

`func (o *WanPortIpv4SettingOpenApiVO) GetVlanPriorityOk() (*int32, bool)`

GetVlanPriorityOk returns a tuple with the VlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPriority

`func (o *WanPortIpv4SettingOpenApiVO) SetVlanPriority(v int32)`

SetVlanPriority sets VlanPriority field to given value.

### HasVlanPriority

`func (o *WanPortIpv4SettingOpenApiVO) HasVlanPriority() bool`

HasVlanPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


