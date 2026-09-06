# OswNetworkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpRelay** | Pointer to [**OswDhcpRelayVO**](OswDhcpRelayVO.md) |  | [optional] 
**DhcpServer** | Pointer to [**OswDhcpServerVO**](OswDhcpServerVO.md) |  | [optional] 
**Id** | Pointer to **string** | Network ID | [optional] 
**Ip** | Pointer to [**OswIpSettingVO**](OswIpSettingVO.md) |  | [optional] 
**Ipv6** | Pointer to [**OswIpv6SettingVO**](OswIpv6SettingVO.md) |  | [optional] 
**Ipv6Enable** | Pointer to **bool** | Enable IPV6 or not. | [optional] 
**Mode** | **int32** | DHCP mode. 0: None, mode 1: DHCP Server, mode 2: DHCP Relay. | 
**Mtu** | Pointer to **int32** | MTU, MTU value should be less than or equal to the jumbo value. | [optional] 
**Mvlan** | Pointer to **bool** | Indicate the vlan is management vlan or not. | [optional] 
**Name** | Pointer to **string** | Switch network name. | [optional] 
**Status** | Pointer to **int32** | Enable status of the network vlan. 0: disable; 1: enable. | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID. | [optional] 
**VrfId** | Pointer to **string** | VRF ID | [optional] 

## Methods

### NewOswNetworkVO

`func NewOswNetworkVO(mode int32, ) *OswNetworkVO`

NewOswNetworkVO instantiates a new OswNetworkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswNetworkVOWithDefaults

`func NewOswNetworkVOWithDefaults() *OswNetworkVO`

NewOswNetworkVOWithDefaults instantiates a new OswNetworkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpRelay

`func (o *OswNetworkVO) GetDhcpRelay() OswDhcpRelayVO`

GetDhcpRelay returns the DhcpRelay field if non-nil, zero value otherwise.

### GetDhcpRelayOk

`func (o *OswNetworkVO) GetDhcpRelayOk() (*OswDhcpRelayVO, bool)`

GetDhcpRelayOk returns a tuple with the DhcpRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelay

`func (o *OswNetworkVO) SetDhcpRelay(v OswDhcpRelayVO)`

SetDhcpRelay sets DhcpRelay field to given value.

### HasDhcpRelay

`func (o *OswNetworkVO) HasDhcpRelay() bool`

HasDhcpRelay returns a boolean if a field has been set.

### GetDhcpServer

`func (o *OswNetworkVO) GetDhcpServer() OswDhcpServerVO`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *OswNetworkVO) GetDhcpServerOk() (*OswDhcpServerVO, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *OswNetworkVO) SetDhcpServer(v OswDhcpServerVO)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *OswNetworkVO) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetId

`func (o *OswNetworkVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswNetworkVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswNetworkVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswNetworkVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *OswNetworkVO) GetIp() OswIpSettingVO`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswNetworkVO) GetIpOk() (*OswIpSettingVO, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswNetworkVO) SetIp(v OswIpSettingVO)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswNetworkVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *OswNetworkVO) GetIpv6() OswIpv6SettingVO`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *OswNetworkVO) GetIpv6Ok() (*OswIpv6SettingVO, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *OswNetworkVO) SetIpv6(v OswIpv6SettingVO)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *OswNetworkVO) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetIpv6Enable

`func (o *OswNetworkVO) GetIpv6Enable() bool`

GetIpv6Enable returns the Ipv6Enable field if non-nil, zero value otherwise.

### GetIpv6EnableOk

`func (o *OswNetworkVO) GetIpv6EnableOk() (*bool, bool)`

GetIpv6EnableOk returns a tuple with the Ipv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enable

`func (o *OswNetworkVO) SetIpv6Enable(v bool)`

SetIpv6Enable sets Ipv6Enable field to given value.

### HasIpv6Enable

`func (o *OswNetworkVO) HasIpv6Enable() bool`

HasIpv6Enable returns a boolean if a field has been set.

### GetMode

`func (o *OswNetworkVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OswNetworkVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OswNetworkVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetMtu

`func (o *OswNetworkVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *OswNetworkVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *OswNetworkVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *OswNetworkVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetMvlan

`func (o *OswNetworkVO) GetMvlan() bool`

GetMvlan returns the Mvlan field if non-nil, zero value otherwise.

### GetMvlanOk

`func (o *OswNetworkVO) GetMvlanOk() (*bool, bool)`

GetMvlanOk returns a tuple with the Mvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlan

`func (o *OswNetworkVO) SetMvlan(v bool)`

SetMvlan sets Mvlan field to given value.

### HasMvlan

`func (o *OswNetworkVO) HasMvlan() bool`

HasMvlan returns a boolean if a field has been set.

### GetName

`func (o *OswNetworkVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswNetworkVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswNetworkVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswNetworkVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OswNetworkVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswNetworkVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswNetworkVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswNetworkVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlan

`func (o *OswNetworkVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OswNetworkVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OswNetworkVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *OswNetworkVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVrfId

`func (o *OswNetworkVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *OswNetworkVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *OswNetworkVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *OswNetworkVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


