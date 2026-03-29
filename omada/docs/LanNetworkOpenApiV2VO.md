# LanNetworkOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllLan** | Pointer to **bool** | When Internet pre-config is closed or Internet pre-config is Universal, allLAN is \&quot;true\&quot;; after adopting gateway, allLAN is \&quot;false\&quot;. | [optional] 
**Application** | Pointer to **int32** | Effective device type should be a value as follows: 0: Gateway and Switch; 1: Switch | [optional] 
**DhcpGuard** | Pointer to [**DhcpServersSetting**](DhcpServersSetting.md) |  | [optional] 
**DhcpL2RelayEnable** | Pointer to **bool** | The switch of DHCP L2 relay | [optional] 
**DhcpSettingsVO** | Pointer to [**DhcpSettingConfig**](DhcpSettingConfig.md) |  | [optional] 
**Dhcpv6Guard** | Pointer to [**Dhcpv6ServersSetting**](Dhcpv6ServersSetting.md) |  | [optional] 
**Domain** | Pointer to **string** | The domain of this network | [optional] 
**GatewaySubnet** | Pointer to **string** | When purpose is interface, gateway subnet is needed. Format: IP/Mask | [optional] 
**IgmpSnoopEnable** | **bool** | Enable IGMP snooping | 
**InterfaceIds** | Pointer to **[]string** | Gateway LAN port IDs, which are acquired from \&quot;Check WAN/LAN status\&quot; | [optional] 
**LanNetworkIpv6Config** | Pointer to [**LanNetworkIPV6Config**](LanNetworkIPV6Config.md) |  | [optional] 
**MldSnoopEnable** | Pointer to **bool** | Enable MLD snooping | [optional] 
**Name** | **string** | LAN network name should contain 1 to 128 characters. | 
**Purpose** | **int32** | LAN network purpose should be a value as follows: 0: VLAN; 1: interface | 
**QosQueueEnable** | Pointer to **bool** | The switch of QoS queue. | [optional] 
**QueueId** | Pointer to **int32** | QoS queue Id. | [optional] 
**SubnetOverrideEnable** | Pointer to **bool** | The switch status of DHCP Settings Overrides. | [optional] 
**Vlan** | Pointer to **int32** | When purpose is \&quot;VLAN\&quot; or purpose is \&quot;interface\&quot; and VLANType is 0, vlan should be within the range of 1-4090.When purpose is \&quot;VLAN\&quot; and VLANType is 0 and \&quot;application\&quot; is 1, vlan could be within the range of 1-4094. | [optional] 
**VlanType** | Pointer to **int32** | When purpose is interface, VLANType should be a value as follows: 0: Single; 1: Multiple | [optional] 
**Vlans** | Pointer to **string** | When purpose is \&quot;interface\&quot; and VLANType is 1, batch create VLANs. VLAN format: 200, 1-100. | [optional] 

## Methods

### NewLanNetworkOpenApiV2VO

`func NewLanNetworkOpenApiV2VO(igmpSnoopEnable bool, name string, purpose int32, ) *LanNetworkOpenApiV2VO`

NewLanNetworkOpenApiV2VO instantiates a new LanNetworkOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkOpenApiV2VOWithDefaults

`func NewLanNetworkOpenApiV2VOWithDefaults() *LanNetworkOpenApiV2VO`

NewLanNetworkOpenApiV2VOWithDefaults instantiates a new LanNetworkOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllLan

`func (o *LanNetworkOpenApiV2VO) GetAllLan() bool`

GetAllLan returns the AllLan field if non-nil, zero value otherwise.

### GetAllLanOk

`func (o *LanNetworkOpenApiV2VO) GetAllLanOk() (*bool, bool)`

GetAllLanOk returns a tuple with the AllLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLan

`func (o *LanNetworkOpenApiV2VO) SetAllLan(v bool)`

SetAllLan sets AllLan field to given value.

### HasAllLan

`func (o *LanNetworkOpenApiV2VO) HasAllLan() bool`

HasAllLan returns a boolean if a field has been set.

### GetApplication

`func (o *LanNetworkOpenApiV2VO) GetApplication() int32`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *LanNetworkOpenApiV2VO) GetApplicationOk() (*int32, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *LanNetworkOpenApiV2VO) SetApplication(v int32)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *LanNetworkOpenApiV2VO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDhcpGuard

`func (o *LanNetworkOpenApiV2VO) GetDhcpGuard() DhcpServersSetting`

GetDhcpGuard returns the DhcpGuard field if non-nil, zero value otherwise.

### GetDhcpGuardOk

`func (o *LanNetworkOpenApiV2VO) GetDhcpGuardOk() (*DhcpServersSetting, bool)`

GetDhcpGuardOk returns a tuple with the DhcpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpGuard

`func (o *LanNetworkOpenApiV2VO) SetDhcpGuard(v DhcpServersSetting)`

SetDhcpGuard sets DhcpGuard field to given value.

### HasDhcpGuard

`func (o *LanNetworkOpenApiV2VO) HasDhcpGuard() bool`

HasDhcpGuard returns a boolean if a field has been set.

### GetDhcpL2RelayEnable

`func (o *LanNetworkOpenApiV2VO) GetDhcpL2RelayEnable() bool`

GetDhcpL2RelayEnable returns the DhcpL2RelayEnable field if non-nil, zero value otherwise.

### GetDhcpL2RelayEnableOk

`func (o *LanNetworkOpenApiV2VO) GetDhcpL2RelayEnableOk() (*bool, bool)`

GetDhcpL2RelayEnableOk returns a tuple with the DhcpL2RelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelayEnable

`func (o *LanNetworkOpenApiV2VO) SetDhcpL2RelayEnable(v bool)`

SetDhcpL2RelayEnable sets DhcpL2RelayEnable field to given value.

### HasDhcpL2RelayEnable

`func (o *LanNetworkOpenApiV2VO) HasDhcpL2RelayEnable() bool`

HasDhcpL2RelayEnable returns a boolean if a field has been set.

### GetDhcpSettingsVO

`func (o *LanNetworkOpenApiV2VO) GetDhcpSettingsVO() DhcpSettingConfig`

GetDhcpSettingsVO returns the DhcpSettingsVO field if non-nil, zero value otherwise.

### GetDhcpSettingsVOOk

`func (o *LanNetworkOpenApiV2VO) GetDhcpSettingsVOOk() (*DhcpSettingConfig, bool)`

GetDhcpSettingsVOOk returns a tuple with the DhcpSettingsVO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSettingsVO

`func (o *LanNetworkOpenApiV2VO) SetDhcpSettingsVO(v DhcpSettingConfig)`

SetDhcpSettingsVO sets DhcpSettingsVO field to given value.

### HasDhcpSettingsVO

`func (o *LanNetworkOpenApiV2VO) HasDhcpSettingsVO() bool`

HasDhcpSettingsVO returns a boolean if a field has been set.

### GetDhcpv6Guard

`func (o *LanNetworkOpenApiV2VO) GetDhcpv6Guard() Dhcpv6ServersSetting`

GetDhcpv6Guard returns the Dhcpv6Guard field if non-nil, zero value otherwise.

### GetDhcpv6GuardOk

`func (o *LanNetworkOpenApiV2VO) GetDhcpv6GuardOk() (*Dhcpv6ServersSetting, bool)`

GetDhcpv6GuardOk returns a tuple with the Dhcpv6Guard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Guard

`func (o *LanNetworkOpenApiV2VO) SetDhcpv6Guard(v Dhcpv6ServersSetting)`

SetDhcpv6Guard sets Dhcpv6Guard field to given value.

### HasDhcpv6Guard

`func (o *LanNetworkOpenApiV2VO) HasDhcpv6Guard() bool`

HasDhcpv6Guard returns a boolean if a field has been set.

### GetDomain

`func (o *LanNetworkOpenApiV2VO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanNetworkOpenApiV2VO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanNetworkOpenApiV2VO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LanNetworkOpenApiV2VO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *LanNetworkOpenApiV2VO) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *LanNetworkOpenApiV2VO) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *LanNetworkOpenApiV2VO) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *LanNetworkOpenApiV2VO) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetIgmpSnoopEnable

`func (o *LanNetworkOpenApiV2VO) GetIgmpSnoopEnable() bool`

GetIgmpSnoopEnable returns the IgmpSnoopEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopEnableOk

`func (o *LanNetworkOpenApiV2VO) GetIgmpSnoopEnableOk() (*bool, bool)`

GetIgmpSnoopEnableOk returns a tuple with the IgmpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopEnable

`func (o *LanNetworkOpenApiV2VO) SetIgmpSnoopEnable(v bool)`

SetIgmpSnoopEnable sets IgmpSnoopEnable field to given value.


### GetInterfaceIds

`func (o *LanNetworkOpenApiV2VO) GetInterfaceIds() []string`

GetInterfaceIds returns the InterfaceIds field if non-nil, zero value otherwise.

### GetInterfaceIdsOk

`func (o *LanNetworkOpenApiV2VO) GetInterfaceIdsOk() (*[]string, bool)`

GetInterfaceIdsOk returns a tuple with the InterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIds

`func (o *LanNetworkOpenApiV2VO) SetInterfaceIds(v []string)`

SetInterfaceIds sets InterfaceIds field to given value.

### HasInterfaceIds

`func (o *LanNetworkOpenApiV2VO) HasInterfaceIds() bool`

HasInterfaceIds returns a boolean if a field has been set.

### GetLanNetworkIpv6Config

`func (o *LanNetworkOpenApiV2VO) GetLanNetworkIpv6Config() LanNetworkIPV6Config`

GetLanNetworkIpv6Config returns the LanNetworkIpv6Config field if non-nil, zero value otherwise.

### GetLanNetworkIpv6ConfigOk

`func (o *LanNetworkOpenApiV2VO) GetLanNetworkIpv6ConfigOk() (*LanNetworkIPV6Config, bool)`

GetLanNetworkIpv6ConfigOk returns a tuple with the LanNetworkIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkIpv6Config

`func (o *LanNetworkOpenApiV2VO) SetLanNetworkIpv6Config(v LanNetworkIPV6Config)`

SetLanNetworkIpv6Config sets LanNetworkIpv6Config field to given value.

### HasLanNetworkIpv6Config

`func (o *LanNetworkOpenApiV2VO) HasLanNetworkIpv6Config() bool`

HasLanNetworkIpv6Config returns a boolean if a field has been set.

### GetMldSnoopEnable

`func (o *LanNetworkOpenApiV2VO) GetMldSnoopEnable() bool`

GetMldSnoopEnable returns the MldSnoopEnable field if non-nil, zero value otherwise.

### GetMldSnoopEnableOk

`func (o *LanNetworkOpenApiV2VO) GetMldSnoopEnableOk() (*bool, bool)`

GetMldSnoopEnableOk returns a tuple with the MldSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldSnoopEnable

`func (o *LanNetworkOpenApiV2VO) SetMldSnoopEnable(v bool)`

SetMldSnoopEnable sets MldSnoopEnable field to given value.

### HasMldSnoopEnable

`func (o *LanNetworkOpenApiV2VO) HasMldSnoopEnable() bool`

HasMldSnoopEnable returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkOpenApiV2VO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkOpenApiV2VO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkOpenApiV2VO) SetName(v string)`

SetName sets Name field to given value.


### GetPurpose

`func (o *LanNetworkOpenApiV2VO) GetPurpose() int32`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *LanNetworkOpenApiV2VO) GetPurposeOk() (*int32, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *LanNetworkOpenApiV2VO) SetPurpose(v int32)`

SetPurpose sets Purpose field to given value.


### GetQosQueueEnable

`func (o *LanNetworkOpenApiV2VO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *LanNetworkOpenApiV2VO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *LanNetworkOpenApiV2VO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *LanNetworkOpenApiV2VO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQueueId

`func (o *LanNetworkOpenApiV2VO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *LanNetworkOpenApiV2VO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *LanNetworkOpenApiV2VO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *LanNetworkOpenApiV2VO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetSubnetOverrideEnable

`func (o *LanNetworkOpenApiV2VO) GetSubnetOverrideEnable() bool`

GetSubnetOverrideEnable returns the SubnetOverrideEnable field if non-nil, zero value otherwise.

### GetSubnetOverrideEnableOk

`func (o *LanNetworkOpenApiV2VO) GetSubnetOverrideEnableOk() (*bool, bool)`

GetSubnetOverrideEnableOk returns a tuple with the SubnetOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOverrideEnable

`func (o *LanNetworkOpenApiV2VO) SetSubnetOverrideEnable(v bool)`

SetSubnetOverrideEnable sets SubnetOverrideEnable field to given value.

### HasSubnetOverrideEnable

`func (o *LanNetworkOpenApiV2VO) HasSubnetOverrideEnable() bool`

HasSubnetOverrideEnable returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkOpenApiV2VO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkOpenApiV2VO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkOpenApiV2VO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkOpenApiV2VO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkOpenApiV2VO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkOpenApiV2VO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkOpenApiV2VO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkOpenApiV2VO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetVlans

`func (o *LanNetworkOpenApiV2VO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *LanNetworkOpenApiV2VO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *LanNetworkOpenApiV2VO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *LanNetworkOpenApiV2VO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


