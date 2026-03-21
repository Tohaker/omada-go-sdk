# OswNetworkOpenApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpRelay** | Pointer to [**OswDhcpRelayOpenApiVO**](OswDhcpRelayOpenApiVO.md) |  | [optional] 
**DhcpServer** | Pointer to [**OswDhcpServerOpenApiVO**](OswDhcpServerOpenApiVO.md) |  | [optional] 
**Id** | **string** | Network ID | 
**Ip** | Pointer to [**OswIpSettingOpenApiVO**](OswIpSettingOpenApiVO.md) |  | [optional] 
**Ipv6** | Pointer to [**OswIpv6SettingOpenApiVO**](OswIpv6SettingOpenApiVO.md) |  | [optional] 
**Ipv6Enable** | Pointer to **bool** | Enable IPV6 or not. | [optional] 
**Mode** | **int32** | DHCP mode. 0: None, mode 1: DHCP Server, mode 2: DHCP Relay | 
**Mvlan** | **bool** | Indicate the vlan is management vlan or not. | 
**Name** | Pointer to **string** | Switch network name. | [optional] 
**Status** | Pointer to **bool** | Enable status of the network vlan. | [optional] 
**Vlan** | **int32** | VLAN ID. | 
**VrfId** | Pointer to **string** | VRF ID | [optional] 

## Methods

### NewOswNetworkOpenApi

`func NewOswNetworkOpenApi(id string, mode int32, mvlan bool, vlan int32, ) *OswNetworkOpenApi`

NewOswNetworkOpenApi instantiates a new OswNetworkOpenApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswNetworkOpenApiWithDefaults

`func NewOswNetworkOpenApiWithDefaults() *OswNetworkOpenApi`

NewOswNetworkOpenApiWithDefaults instantiates a new OswNetworkOpenApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpRelay

`func (o *OswNetworkOpenApi) GetDhcpRelay() OswDhcpRelayOpenApiVO`

GetDhcpRelay returns the DhcpRelay field if non-nil, zero value otherwise.

### GetDhcpRelayOk

`func (o *OswNetworkOpenApi) GetDhcpRelayOk() (*OswDhcpRelayOpenApiVO, bool)`

GetDhcpRelayOk returns a tuple with the DhcpRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelay

`func (o *OswNetworkOpenApi) SetDhcpRelay(v OswDhcpRelayOpenApiVO)`

SetDhcpRelay sets DhcpRelay field to given value.

### HasDhcpRelay

`func (o *OswNetworkOpenApi) HasDhcpRelay() bool`

HasDhcpRelay returns a boolean if a field has been set.

### GetDhcpServer

`func (o *OswNetworkOpenApi) GetDhcpServer() OswDhcpServerOpenApiVO`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *OswNetworkOpenApi) GetDhcpServerOk() (*OswDhcpServerOpenApiVO, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *OswNetworkOpenApi) SetDhcpServer(v OswDhcpServerOpenApiVO)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *OswNetworkOpenApi) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetId

`func (o *OswNetworkOpenApi) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswNetworkOpenApi) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswNetworkOpenApi) SetId(v string)`

SetId sets Id field to given value.


### GetIp

`func (o *OswNetworkOpenApi) GetIp() OswIpSettingOpenApiVO`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswNetworkOpenApi) GetIpOk() (*OswIpSettingOpenApiVO, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswNetworkOpenApi) SetIp(v OswIpSettingOpenApiVO)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswNetworkOpenApi) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *OswNetworkOpenApi) GetIpv6() OswIpv6SettingOpenApiVO`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *OswNetworkOpenApi) GetIpv6Ok() (*OswIpv6SettingOpenApiVO, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *OswNetworkOpenApi) SetIpv6(v OswIpv6SettingOpenApiVO)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *OswNetworkOpenApi) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetIpv6Enable

`func (o *OswNetworkOpenApi) GetIpv6Enable() bool`

GetIpv6Enable returns the Ipv6Enable field if non-nil, zero value otherwise.

### GetIpv6EnableOk

`func (o *OswNetworkOpenApi) GetIpv6EnableOk() (*bool, bool)`

GetIpv6EnableOk returns a tuple with the Ipv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enable

`func (o *OswNetworkOpenApi) SetIpv6Enable(v bool)`

SetIpv6Enable sets Ipv6Enable field to given value.

### HasIpv6Enable

`func (o *OswNetworkOpenApi) HasIpv6Enable() bool`

HasIpv6Enable returns a boolean if a field has been set.

### GetMode

`func (o *OswNetworkOpenApi) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OswNetworkOpenApi) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OswNetworkOpenApi) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetMvlan

`func (o *OswNetworkOpenApi) GetMvlan() bool`

GetMvlan returns the Mvlan field if non-nil, zero value otherwise.

### GetMvlanOk

`func (o *OswNetworkOpenApi) GetMvlanOk() (*bool, bool)`

GetMvlanOk returns a tuple with the Mvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlan

`func (o *OswNetworkOpenApi) SetMvlan(v bool)`

SetMvlan sets Mvlan field to given value.


### GetName

`func (o *OswNetworkOpenApi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswNetworkOpenApi) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswNetworkOpenApi) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswNetworkOpenApi) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OswNetworkOpenApi) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswNetworkOpenApi) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswNetworkOpenApi) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswNetworkOpenApi) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlan

`func (o *OswNetworkOpenApi) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OswNetworkOpenApi) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OswNetworkOpenApi) SetVlan(v int32)`

SetVlan sets Vlan field to given value.


### GetVrfId

`func (o *OswNetworkOpenApi) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *OswNetworkOpenApi) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *OswNetworkOpenApi) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *OswNetworkOpenApi) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


