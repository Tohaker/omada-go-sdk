# VirtualWanIpv4SettingInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Dhcp** | Pointer to [**VirtualWanIpv4DhcpOpenApiVO**](VirtualWanIpv4DhcpOpenApiVO.md) |  | [optional] 
**Ipv4Pppoe** | Pointer to [**VirtualWanIpv4PppoeOpenApiVO**](VirtualWanIpv4PppoeOpenApiVO.md) |  | [optional] 
**Ipv4Static** | Pointer to [**VirtualWanIpv4StaticOpenApiVO**](VirtualWanIpv4StaticOpenApiVO.md) |  | [optional] 
**Proto** | Pointer to **string** | Virtual WAN IPv4 connection type. | [optional] 
**ProtoType** | Pointer to **int32** | Virtual WAN IPv4 proto type, 0:static; 1:DHCP; 2:PPPoE. | [optional] 
**QosTag** | Pointer to **int32** | Qos Tag. | [optional] 
**QosTagEnable** | Pointer to **bool** | Whether Qos Tag is enable. | [optional] 
**SupportInternetVlan** | Pointer to **bool** | Whether support internet vlan. | [optional] 
**SupportQosTagEnable** | Pointer to **bool** | Whether support Qos Tag enable. | [optional] 
**VlanId** | Pointer to **int32** | Vlan ID. | [optional] 
**VlanPriority** | Pointer to **int32** | Vlan Priority. | [optional] 

## Methods

### NewVirtualWanIpv4SettingInfoOpenApiVO

`func NewVirtualWanIpv4SettingInfoOpenApiVO() *VirtualWanIpv4SettingInfoOpenApiVO`

NewVirtualWanIpv4SettingInfoOpenApiVO instantiates a new VirtualWanIpv4SettingInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanIpv4SettingInfoOpenApiVOWithDefaults

`func NewVirtualWanIpv4SettingInfoOpenApiVOWithDefaults() *VirtualWanIpv4SettingInfoOpenApiVO`

NewVirtualWanIpv4SettingInfoOpenApiVOWithDefaults instantiates a new VirtualWanIpv4SettingInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Dhcp

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetIpv4Dhcp() VirtualWanIpv4DhcpOpenApiVO`

GetIpv4Dhcp returns the Ipv4Dhcp field if non-nil, zero value otherwise.

### GetIpv4DhcpOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetIpv4DhcpOk() (*VirtualWanIpv4DhcpOpenApiVO, bool)`

GetIpv4DhcpOk returns a tuple with the Ipv4Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Dhcp

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetIpv4Dhcp(v VirtualWanIpv4DhcpOpenApiVO)`

SetIpv4Dhcp sets Ipv4Dhcp field to given value.

### HasIpv4Dhcp

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasIpv4Dhcp() bool`

HasIpv4Dhcp returns a boolean if a field has been set.

### GetIpv4Pppoe

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetIpv4Pppoe() VirtualWanIpv4PppoeOpenApiVO`

GetIpv4Pppoe returns the Ipv4Pppoe field if non-nil, zero value otherwise.

### GetIpv4PppoeOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetIpv4PppoeOk() (*VirtualWanIpv4PppoeOpenApiVO, bool)`

GetIpv4PppoeOk returns a tuple with the Ipv4Pppoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Pppoe

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetIpv4Pppoe(v VirtualWanIpv4PppoeOpenApiVO)`

SetIpv4Pppoe sets Ipv4Pppoe field to given value.

### HasIpv4Pppoe

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasIpv4Pppoe() bool`

HasIpv4Pppoe returns a boolean if a field has been set.

### GetIpv4Static

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetIpv4Static() VirtualWanIpv4StaticOpenApiVO`

GetIpv4Static returns the Ipv4Static field if non-nil, zero value otherwise.

### GetIpv4StaticOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetIpv4StaticOk() (*VirtualWanIpv4StaticOpenApiVO, bool)`

GetIpv4StaticOk returns a tuple with the Ipv4Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Static

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetIpv4Static(v VirtualWanIpv4StaticOpenApiVO)`

SetIpv4Static sets Ipv4Static field to given value.

### HasIpv4Static

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasIpv4Static() bool`

HasIpv4Static returns a boolean if a field has been set.

### GetProto

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetProtoType

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetProtoType() int32`

GetProtoType returns the ProtoType field if non-nil, zero value otherwise.

### GetProtoTypeOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetProtoTypeOk() (*int32, bool)`

GetProtoTypeOk returns a tuple with the ProtoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoType

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetProtoType(v int32)`

SetProtoType sets ProtoType field to given value.

### HasProtoType

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasProtoType() bool`

HasProtoType returns a boolean if a field has been set.

### GetQosTag

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetQosTag() int32`

GetQosTag returns the QosTag field if non-nil, zero value otherwise.

### GetQosTagOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetQosTagOk() (*int32, bool)`

GetQosTagOk returns a tuple with the QosTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosTag

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetQosTag(v int32)`

SetQosTag sets QosTag field to given value.

### HasQosTag

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasQosTag() bool`

HasQosTag returns a boolean if a field has been set.

### GetQosTagEnable

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetQosTagEnable() bool`

GetQosTagEnable returns the QosTagEnable field if non-nil, zero value otherwise.

### GetQosTagEnableOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetQosTagEnableOk() (*bool, bool)`

GetQosTagEnableOk returns a tuple with the QosTagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosTagEnable

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetQosTagEnable(v bool)`

SetQosTagEnable sets QosTagEnable field to given value.

### HasQosTagEnable

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasQosTagEnable() bool`

HasQosTagEnable returns a boolean if a field has been set.

### GetSupportInternetVlan

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetSupportInternetVlan() bool`

GetSupportInternetVlan returns the SupportInternetVlan field if non-nil, zero value otherwise.

### GetSupportInternetVlanOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetSupportInternetVlanOk() (*bool, bool)`

GetSupportInternetVlanOk returns a tuple with the SupportInternetVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportInternetVlan

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetSupportInternetVlan(v bool)`

SetSupportInternetVlan sets SupportInternetVlan field to given value.

### HasSupportInternetVlan

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasSupportInternetVlan() bool`

HasSupportInternetVlan returns a boolean if a field has been set.

### GetSupportQosTagEnable

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetSupportQosTagEnable() bool`

GetSupportQosTagEnable returns the SupportQosTagEnable field if non-nil, zero value otherwise.

### GetSupportQosTagEnableOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetSupportQosTagEnableOk() (*bool, bool)`

GetSupportQosTagEnableOk returns a tuple with the SupportQosTagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportQosTagEnable

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetSupportQosTagEnable(v bool)`

SetSupportQosTagEnable sets SupportQosTagEnable field to given value.

### HasSupportQosTagEnable

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasSupportQosTagEnable() bool`

HasSupportQosTagEnable returns a boolean if a field has been set.

### GetVlanId

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanPriority

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetVlanPriority() int32`

GetVlanPriority returns the VlanPriority field if non-nil, zero value otherwise.

### GetVlanPriorityOk

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) GetVlanPriorityOk() (*int32, bool)`

GetVlanPriorityOk returns a tuple with the VlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPriority

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) SetVlanPriority(v int32)`

SetVlanPriority sets VlanPriority field to given value.

### HasVlanPriority

`func (o *VirtualWanIpv4SettingInfoOpenApiVO) HasVlanPriority() bool`

HasVlanPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


