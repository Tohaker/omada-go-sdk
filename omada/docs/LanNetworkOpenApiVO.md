# LanNetworkOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllLan** | Pointer to **bool** | When Internet pre-config is closed or Internet pre-config is Universal, allLAN is \&quot;true\&quot;; after adopting gateway, allLAN is \&quot;false\&quot;. | [optional] 
**Application** | Pointer to **int32** | Effective device type, 0: Gateway, Switch; 1: Switch | [optional] 
**DhcpGuard** | Pointer to [**DhcpServersSetting**](DhcpServersSetting.md) |  | [optional] 
**DhcpL2RelayEnable** | Pointer to **bool** | The switch of DHCP L2 relay | [optional] 
**DhcpSettingsVO** | Pointer to [**DhcpSettings**](DhcpSettings.md) |  | [optional] 
**Dhcpv6Guard** | Pointer to [**Dhcpv6ServersSetting**](Dhcpv6ServersSetting.md) |  | [optional] 
**Domain** | Pointer to **string** | The domain of this network | [optional] 
**GatewaySubnet** | Pointer to **string** | When purpose is interface, gateway subnet is needed. Format: IP/Mask | [optional] 
**IgmpSnoopEnable** | **bool** | Enable IGMP snooping | 
**InterfaceIds** | Pointer to **[]string** | Gateway LAN port IDs, which are acquired from \&quot;Check WAN/LAN status\&quot; | [optional] 
**Isolation** | Pointer to **bool** | Whether isolate this network. | [optional] 
**LanNeworkIpv6Config** | Pointer to [**LanNetworkIPV6Config**](LanNetworkIPV6Config.md) |  | [optional] 
**MldSnoopEnable** | Pointer to **bool** | Enable MLD snooping | [optional] 
**Name** | **string** | LAN network name should contain 1 to 128 characters. | 
**Purpose** | **int32** | LAN network purpose, 0: VLAN, 1: interface | 
**Vlan** | Pointer to **int32** | When purpose is \&quot;VLAN\&quot; or purpose is \&quot;interface\&quot; and VLANType is 0, create VLAN, which ranges from 1 to 4090. | [optional] 
**VlanType** | Pointer to **int32** | When purpose is interface, VLANType is valid. 0: Single; 1: Multiple | [optional] 
**Vlans** | Pointer to **string** | When purpose is \&quot;interface\&quot; and VLANType is 1, batch create VLANs. VLAN format: 200, 1-100. | [optional] 

## Methods

### NewLanNetworkOpenApiVO

`func NewLanNetworkOpenApiVO(igmpSnoopEnable bool, name string, purpose int32, ) *LanNetworkOpenApiVO`

NewLanNetworkOpenApiVO instantiates a new LanNetworkOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkOpenApiVOWithDefaults

`func NewLanNetworkOpenApiVOWithDefaults() *LanNetworkOpenApiVO`

NewLanNetworkOpenApiVOWithDefaults instantiates a new LanNetworkOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllLan

`func (o *LanNetworkOpenApiVO) GetAllLan() bool`

GetAllLan returns the AllLan field if non-nil, zero value otherwise.

### GetAllLanOk

`func (o *LanNetworkOpenApiVO) GetAllLanOk() (*bool, bool)`

GetAllLanOk returns a tuple with the AllLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLan

`func (o *LanNetworkOpenApiVO) SetAllLan(v bool)`

SetAllLan sets AllLan field to given value.

### HasAllLan

`func (o *LanNetworkOpenApiVO) HasAllLan() bool`

HasAllLan returns a boolean if a field has been set.

### GetApplication

`func (o *LanNetworkOpenApiVO) GetApplication() int32`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *LanNetworkOpenApiVO) GetApplicationOk() (*int32, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *LanNetworkOpenApiVO) SetApplication(v int32)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *LanNetworkOpenApiVO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDhcpGuard

`func (o *LanNetworkOpenApiVO) GetDhcpGuard() DhcpServersSetting`

GetDhcpGuard returns the DhcpGuard field if non-nil, zero value otherwise.

### GetDhcpGuardOk

`func (o *LanNetworkOpenApiVO) GetDhcpGuardOk() (*DhcpServersSetting, bool)`

GetDhcpGuardOk returns a tuple with the DhcpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpGuard

`func (o *LanNetworkOpenApiVO) SetDhcpGuard(v DhcpServersSetting)`

SetDhcpGuard sets DhcpGuard field to given value.

### HasDhcpGuard

`func (o *LanNetworkOpenApiVO) HasDhcpGuard() bool`

HasDhcpGuard returns a boolean if a field has been set.

### GetDhcpL2RelayEnable

`func (o *LanNetworkOpenApiVO) GetDhcpL2RelayEnable() bool`

GetDhcpL2RelayEnable returns the DhcpL2RelayEnable field if non-nil, zero value otherwise.

### GetDhcpL2RelayEnableOk

`func (o *LanNetworkOpenApiVO) GetDhcpL2RelayEnableOk() (*bool, bool)`

GetDhcpL2RelayEnableOk returns a tuple with the DhcpL2RelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelayEnable

`func (o *LanNetworkOpenApiVO) SetDhcpL2RelayEnable(v bool)`

SetDhcpL2RelayEnable sets DhcpL2RelayEnable field to given value.

### HasDhcpL2RelayEnable

`func (o *LanNetworkOpenApiVO) HasDhcpL2RelayEnable() bool`

HasDhcpL2RelayEnable returns a boolean if a field has been set.

### GetDhcpSettingsVO

`func (o *LanNetworkOpenApiVO) GetDhcpSettingsVO() DhcpSettings`

GetDhcpSettingsVO returns the DhcpSettingsVO field if non-nil, zero value otherwise.

### GetDhcpSettingsVOOk

`func (o *LanNetworkOpenApiVO) GetDhcpSettingsVOOk() (*DhcpSettings, bool)`

GetDhcpSettingsVOOk returns a tuple with the DhcpSettingsVO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSettingsVO

`func (o *LanNetworkOpenApiVO) SetDhcpSettingsVO(v DhcpSettings)`

SetDhcpSettingsVO sets DhcpSettingsVO field to given value.

### HasDhcpSettingsVO

`func (o *LanNetworkOpenApiVO) HasDhcpSettingsVO() bool`

HasDhcpSettingsVO returns a boolean if a field has been set.

### GetDhcpv6Guard

`func (o *LanNetworkOpenApiVO) GetDhcpv6Guard() Dhcpv6ServersSetting`

GetDhcpv6Guard returns the Dhcpv6Guard field if non-nil, zero value otherwise.

### GetDhcpv6GuardOk

`func (o *LanNetworkOpenApiVO) GetDhcpv6GuardOk() (*Dhcpv6ServersSetting, bool)`

GetDhcpv6GuardOk returns a tuple with the Dhcpv6Guard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Guard

`func (o *LanNetworkOpenApiVO) SetDhcpv6Guard(v Dhcpv6ServersSetting)`

SetDhcpv6Guard sets Dhcpv6Guard field to given value.

### HasDhcpv6Guard

`func (o *LanNetworkOpenApiVO) HasDhcpv6Guard() bool`

HasDhcpv6Guard returns a boolean if a field has been set.

### GetDomain

`func (o *LanNetworkOpenApiVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanNetworkOpenApiVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanNetworkOpenApiVO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LanNetworkOpenApiVO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *LanNetworkOpenApiVO) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *LanNetworkOpenApiVO) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *LanNetworkOpenApiVO) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *LanNetworkOpenApiVO) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetIgmpSnoopEnable

`func (o *LanNetworkOpenApiVO) GetIgmpSnoopEnable() bool`

GetIgmpSnoopEnable returns the IgmpSnoopEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopEnableOk

`func (o *LanNetworkOpenApiVO) GetIgmpSnoopEnableOk() (*bool, bool)`

GetIgmpSnoopEnableOk returns a tuple with the IgmpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopEnable

`func (o *LanNetworkOpenApiVO) SetIgmpSnoopEnable(v bool)`

SetIgmpSnoopEnable sets IgmpSnoopEnable field to given value.


### GetInterfaceIds

`func (o *LanNetworkOpenApiVO) GetInterfaceIds() []string`

GetInterfaceIds returns the InterfaceIds field if non-nil, zero value otherwise.

### GetInterfaceIdsOk

`func (o *LanNetworkOpenApiVO) GetInterfaceIdsOk() (*[]string, bool)`

GetInterfaceIdsOk returns a tuple with the InterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIds

`func (o *LanNetworkOpenApiVO) SetInterfaceIds(v []string)`

SetInterfaceIds sets InterfaceIds field to given value.

### HasInterfaceIds

`func (o *LanNetworkOpenApiVO) HasInterfaceIds() bool`

HasInterfaceIds returns a boolean if a field has been set.

### GetIsolation

`func (o *LanNetworkOpenApiVO) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *LanNetworkOpenApiVO) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *LanNetworkOpenApiVO) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *LanNetworkOpenApiVO) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetLanNeworkIpv6Config

`func (o *LanNetworkOpenApiVO) GetLanNeworkIpv6Config() LanNetworkIPV6Config`

GetLanNeworkIpv6Config returns the LanNeworkIpv6Config field if non-nil, zero value otherwise.

### GetLanNeworkIpv6ConfigOk

`func (o *LanNetworkOpenApiVO) GetLanNeworkIpv6ConfigOk() (*LanNetworkIPV6Config, bool)`

GetLanNeworkIpv6ConfigOk returns a tuple with the LanNeworkIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNeworkIpv6Config

`func (o *LanNetworkOpenApiVO) SetLanNeworkIpv6Config(v LanNetworkIPV6Config)`

SetLanNeworkIpv6Config sets LanNeworkIpv6Config field to given value.

### HasLanNeworkIpv6Config

`func (o *LanNetworkOpenApiVO) HasLanNeworkIpv6Config() bool`

HasLanNeworkIpv6Config returns a boolean if a field has been set.

### GetMldSnoopEnable

`func (o *LanNetworkOpenApiVO) GetMldSnoopEnable() bool`

GetMldSnoopEnable returns the MldSnoopEnable field if non-nil, zero value otherwise.

### GetMldSnoopEnableOk

`func (o *LanNetworkOpenApiVO) GetMldSnoopEnableOk() (*bool, bool)`

GetMldSnoopEnableOk returns a tuple with the MldSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldSnoopEnable

`func (o *LanNetworkOpenApiVO) SetMldSnoopEnable(v bool)`

SetMldSnoopEnable sets MldSnoopEnable field to given value.

### HasMldSnoopEnable

`func (o *LanNetworkOpenApiVO) HasMldSnoopEnable() bool`

HasMldSnoopEnable returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPurpose

`func (o *LanNetworkOpenApiVO) GetPurpose() int32`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *LanNetworkOpenApiVO) GetPurposeOk() (*int32, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *LanNetworkOpenApiVO) SetPurpose(v int32)`

SetPurpose sets Purpose field to given value.


### GetVlan

`func (o *LanNetworkOpenApiVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkOpenApiVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkOpenApiVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkOpenApiVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkOpenApiVO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkOpenApiVO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkOpenApiVO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkOpenApiVO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetVlans

`func (o *LanNetworkOpenApiVO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *LanNetworkOpenApiVO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *LanNetworkOpenApiVO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *LanNetworkOpenApiVO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


