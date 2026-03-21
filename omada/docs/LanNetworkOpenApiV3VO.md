# LanNetworkOpenApiV3VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpDetectionEnable** | Pointer to **bool** | Enable arp detection. Only valid when deviceType is 1 and gateway supports this feature. | [optional] 
**DeviceMac** | Pointer to **string** | DHCP Server Device mac. Only valid when deviceType is 1 or 2. When deviceType is 1, deviceMac can be empty when there is no gateway in the site. | [optional] 
**DeviceType** | **int32** | DHCP Server Device type. It should be a value as follows: 0:External Device 1:gateway 2:switch 3:none | 
**DhcpGuard** | Pointer to [**DhcpServersSetting**](DhcpServersSetting.md) |  | [optional] 
**DhcpL2RelayEnable** | Pointer to **bool** | The switch of DHCP L2 relay | [optional] 
**DhcpRelay** | Pointer to [**OswDhcpRelayOpenApiVO**](OswDhcpRelayOpenApiVO.md) |  | [optional] 
**DhcpServer** | Pointer to [**OswDhcpServerOpenApiVO**](OswDhcpServerOpenApiVO.md) |  | [optional] 
**DhcpSettings** | Pointer to [**DhcpSettingConfig**](DhcpSettingConfig.md) |  | [optional] 
**Dhcpv6Guard** | Pointer to [**Dhcpv6ServersSetting**](Dhcpv6ServersSetting.md) |  | [optional] 
**Domain** | Pointer to **string** | The domain of this network | [optional] 
**GatewaySubnet** | Pointer to **string** | When deviceType is 1, gateway subnet is needed. Format: IP/Mask | [optional] 
**IgmpSnoopEnable** | **bool** | Enable IGMP snooping | 
**Ip** | Pointer to [**OswIpSettingBriefOpenapiVO**](OswIpSettingBriefOpenapiVO.md) |  | [optional] 
**Isolation** | Pointer to **bool** | Whether network isolated. | [optional] 
**LanNetworkIpv6Config** | Pointer to [**LanNetworkIPV6Config**](LanNetworkIPV6Config.md) |  | [optional] 
**MldSnoopEnable** | Pointer to **bool** | Enable MLD snooping | [optional] 
**Mode** | Pointer to **int32** | DHCP mode. 0: None 1: DHCP Server 2: DHCP Relay. Only valid when deviceType is 2. | [optional] 
**Name** | **string** | LAN network name should contain 1 to 128 characters. | 
**QosQueueEnable** | Pointer to **bool** | The switch of QoS queue. | [optional] 
**QueueId** | Pointer to **int32** | QoS queue Id. | [optional] 
**StackId** | Pointer to **string** | DHCP Server Device stackId. Only valid when deviceType is 2 and the selected device is stack | [optional] 
**SubnetOverrideEnable** | Pointer to **bool** | The switch status of DHCP Settings Overrides. | [optional] 
**Vlan** | Pointer to **int32** | Only Valid when vlanType is 0. When deviceType is 1, vlan should be within the range of 1-4090.When deviceType is 0 , 2 or 3, vlan could be within the range of 1-4094. | [optional] 
**VlanType** | Pointer to **int32** | VLANType should be a value as follows: 0: Single; 1: Multiple | [optional] 
**Vlans** | Pointer to **string** | Only valid when vlanType is 1 and device type is 0 , 1 or 3. When device type is 0 or 3, batch create single vlan, when deviceType is 1 , create bridge vlan. VLAN format: 200, 1-100. | [optional] 
**VrfId** | Pointer to **string** | VRF ID | [optional] 

## Methods

### NewLanNetworkOpenApiV3VO

`func NewLanNetworkOpenApiV3VO(deviceType int32, igmpSnoopEnable bool, name string, ) *LanNetworkOpenApiV3VO`

NewLanNetworkOpenApiV3VO instantiates a new LanNetworkOpenApiV3VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkOpenApiV3VOWithDefaults

`func NewLanNetworkOpenApiV3VOWithDefaults() *LanNetworkOpenApiV3VO`

NewLanNetworkOpenApiV3VOWithDefaults instantiates a new LanNetworkOpenApiV3VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpDetectionEnable

`func (o *LanNetworkOpenApiV3VO) GetArpDetectionEnable() bool`

GetArpDetectionEnable returns the ArpDetectionEnable field if non-nil, zero value otherwise.

### GetArpDetectionEnableOk

`func (o *LanNetworkOpenApiV3VO) GetArpDetectionEnableOk() (*bool, bool)`

GetArpDetectionEnableOk returns a tuple with the ArpDetectionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpDetectionEnable

`func (o *LanNetworkOpenApiV3VO) SetArpDetectionEnable(v bool)`

SetArpDetectionEnable sets ArpDetectionEnable field to given value.

### HasArpDetectionEnable

`func (o *LanNetworkOpenApiV3VO) HasArpDetectionEnable() bool`

HasArpDetectionEnable returns a boolean if a field has been set.

### GetDeviceMac

`func (o *LanNetworkOpenApiV3VO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *LanNetworkOpenApiV3VO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *LanNetworkOpenApiV3VO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *LanNetworkOpenApiV3VO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceType

`func (o *LanNetworkOpenApiV3VO) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *LanNetworkOpenApiV3VO) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *LanNetworkOpenApiV3VO) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.


### GetDhcpGuard

`func (o *LanNetworkOpenApiV3VO) GetDhcpGuard() DhcpServersSetting`

GetDhcpGuard returns the DhcpGuard field if non-nil, zero value otherwise.

### GetDhcpGuardOk

`func (o *LanNetworkOpenApiV3VO) GetDhcpGuardOk() (*DhcpServersSetting, bool)`

GetDhcpGuardOk returns a tuple with the DhcpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpGuard

`func (o *LanNetworkOpenApiV3VO) SetDhcpGuard(v DhcpServersSetting)`

SetDhcpGuard sets DhcpGuard field to given value.

### HasDhcpGuard

`func (o *LanNetworkOpenApiV3VO) HasDhcpGuard() bool`

HasDhcpGuard returns a boolean if a field has been set.

### GetDhcpL2RelayEnable

`func (o *LanNetworkOpenApiV3VO) GetDhcpL2RelayEnable() bool`

GetDhcpL2RelayEnable returns the DhcpL2RelayEnable field if non-nil, zero value otherwise.

### GetDhcpL2RelayEnableOk

`func (o *LanNetworkOpenApiV3VO) GetDhcpL2RelayEnableOk() (*bool, bool)`

GetDhcpL2RelayEnableOk returns a tuple with the DhcpL2RelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelayEnable

`func (o *LanNetworkOpenApiV3VO) SetDhcpL2RelayEnable(v bool)`

SetDhcpL2RelayEnable sets DhcpL2RelayEnable field to given value.

### HasDhcpL2RelayEnable

`func (o *LanNetworkOpenApiV3VO) HasDhcpL2RelayEnable() bool`

HasDhcpL2RelayEnable returns a boolean if a field has been set.

### GetDhcpRelay

`func (o *LanNetworkOpenApiV3VO) GetDhcpRelay() OswDhcpRelayOpenApiVO`

GetDhcpRelay returns the DhcpRelay field if non-nil, zero value otherwise.

### GetDhcpRelayOk

`func (o *LanNetworkOpenApiV3VO) GetDhcpRelayOk() (*OswDhcpRelayOpenApiVO, bool)`

GetDhcpRelayOk returns a tuple with the DhcpRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelay

`func (o *LanNetworkOpenApiV3VO) SetDhcpRelay(v OswDhcpRelayOpenApiVO)`

SetDhcpRelay sets DhcpRelay field to given value.

### HasDhcpRelay

`func (o *LanNetworkOpenApiV3VO) HasDhcpRelay() bool`

HasDhcpRelay returns a boolean if a field has been set.

### GetDhcpServer

`func (o *LanNetworkOpenApiV3VO) GetDhcpServer() OswDhcpServerOpenApiVO`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *LanNetworkOpenApiV3VO) GetDhcpServerOk() (*OswDhcpServerOpenApiVO, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *LanNetworkOpenApiV3VO) SetDhcpServer(v OswDhcpServerOpenApiVO)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *LanNetworkOpenApiV3VO) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpSettings

`func (o *LanNetworkOpenApiV3VO) GetDhcpSettings() DhcpSettingConfig`

GetDhcpSettings returns the DhcpSettings field if non-nil, zero value otherwise.

### GetDhcpSettingsOk

`func (o *LanNetworkOpenApiV3VO) GetDhcpSettingsOk() (*DhcpSettingConfig, bool)`

GetDhcpSettingsOk returns a tuple with the DhcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSettings

`func (o *LanNetworkOpenApiV3VO) SetDhcpSettings(v DhcpSettingConfig)`

SetDhcpSettings sets DhcpSettings field to given value.

### HasDhcpSettings

`func (o *LanNetworkOpenApiV3VO) HasDhcpSettings() bool`

HasDhcpSettings returns a boolean if a field has been set.

### GetDhcpv6Guard

`func (o *LanNetworkOpenApiV3VO) GetDhcpv6Guard() Dhcpv6ServersSetting`

GetDhcpv6Guard returns the Dhcpv6Guard field if non-nil, zero value otherwise.

### GetDhcpv6GuardOk

`func (o *LanNetworkOpenApiV3VO) GetDhcpv6GuardOk() (*Dhcpv6ServersSetting, bool)`

GetDhcpv6GuardOk returns a tuple with the Dhcpv6Guard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Guard

`func (o *LanNetworkOpenApiV3VO) SetDhcpv6Guard(v Dhcpv6ServersSetting)`

SetDhcpv6Guard sets Dhcpv6Guard field to given value.

### HasDhcpv6Guard

`func (o *LanNetworkOpenApiV3VO) HasDhcpv6Guard() bool`

HasDhcpv6Guard returns a boolean if a field has been set.

### GetDomain

`func (o *LanNetworkOpenApiV3VO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanNetworkOpenApiV3VO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanNetworkOpenApiV3VO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LanNetworkOpenApiV3VO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *LanNetworkOpenApiV3VO) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *LanNetworkOpenApiV3VO) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *LanNetworkOpenApiV3VO) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *LanNetworkOpenApiV3VO) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetIgmpSnoopEnable

`func (o *LanNetworkOpenApiV3VO) GetIgmpSnoopEnable() bool`

GetIgmpSnoopEnable returns the IgmpSnoopEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopEnableOk

`func (o *LanNetworkOpenApiV3VO) GetIgmpSnoopEnableOk() (*bool, bool)`

GetIgmpSnoopEnableOk returns a tuple with the IgmpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopEnable

`func (o *LanNetworkOpenApiV3VO) SetIgmpSnoopEnable(v bool)`

SetIgmpSnoopEnable sets IgmpSnoopEnable field to given value.


### GetIp

`func (o *LanNetworkOpenApiV3VO) GetIp() OswIpSettingBriefOpenapiVO`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LanNetworkOpenApiV3VO) GetIpOk() (*OswIpSettingBriefOpenapiVO, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LanNetworkOpenApiV3VO) SetIp(v OswIpSettingBriefOpenapiVO)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LanNetworkOpenApiV3VO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsolation

`func (o *LanNetworkOpenApiV3VO) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *LanNetworkOpenApiV3VO) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *LanNetworkOpenApiV3VO) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *LanNetworkOpenApiV3VO) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetLanNetworkIpv6Config

`func (o *LanNetworkOpenApiV3VO) GetLanNetworkIpv6Config() LanNetworkIPV6Config`

GetLanNetworkIpv6Config returns the LanNetworkIpv6Config field if non-nil, zero value otherwise.

### GetLanNetworkIpv6ConfigOk

`func (o *LanNetworkOpenApiV3VO) GetLanNetworkIpv6ConfigOk() (*LanNetworkIPV6Config, bool)`

GetLanNetworkIpv6ConfigOk returns a tuple with the LanNetworkIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkIpv6Config

`func (o *LanNetworkOpenApiV3VO) SetLanNetworkIpv6Config(v LanNetworkIPV6Config)`

SetLanNetworkIpv6Config sets LanNetworkIpv6Config field to given value.

### HasLanNetworkIpv6Config

`func (o *LanNetworkOpenApiV3VO) HasLanNetworkIpv6Config() bool`

HasLanNetworkIpv6Config returns a boolean if a field has been set.

### GetMldSnoopEnable

`func (o *LanNetworkOpenApiV3VO) GetMldSnoopEnable() bool`

GetMldSnoopEnable returns the MldSnoopEnable field if non-nil, zero value otherwise.

### GetMldSnoopEnableOk

`func (o *LanNetworkOpenApiV3VO) GetMldSnoopEnableOk() (*bool, bool)`

GetMldSnoopEnableOk returns a tuple with the MldSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldSnoopEnable

`func (o *LanNetworkOpenApiV3VO) SetMldSnoopEnable(v bool)`

SetMldSnoopEnable sets MldSnoopEnable field to given value.

### HasMldSnoopEnable

`func (o *LanNetworkOpenApiV3VO) HasMldSnoopEnable() bool`

HasMldSnoopEnable returns a boolean if a field has been set.

### GetMode

`func (o *LanNetworkOpenApiV3VO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *LanNetworkOpenApiV3VO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *LanNetworkOpenApiV3VO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *LanNetworkOpenApiV3VO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkOpenApiV3VO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkOpenApiV3VO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkOpenApiV3VO) SetName(v string)`

SetName sets Name field to given value.


### GetQosQueueEnable

`func (o *LanNetworkOpenApiV3VO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *LanNetworkOpenApiV3VO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *LanNetworkOpenApiV3VO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *LanNetworkOpenApiV3VO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQueueId

`func (o *LanNetworkOpenApiV3VO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *LanNetworkOpenApiV3VO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *LanNetworkOpenApiV3VO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *LanNetworkOpenApiV3VO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetStackId

`func (o *LanNetworkOpenApiV3VO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *LanNetworkOpenApiV3VO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *LanNetworkOpenApiV3VO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *LanNetworkOpenApiV3VO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetSubnetOverrideEnable

`func (o *LanNetworkOpenApiV3VO) GetSubnetOverrideEnable() bool`

GetSubnetOverrideEnable returns the SubnetOverrideEnable field if non-nil, zero value otherwise.

### GetSubnetOverrideEnableOk

`func (o *LanNetworkOpenApiV3VO) GetSubnetOverrideEnableOk() (*bool, bool)`

GetSubnetOverrideEnableOk returns a tuple with the SubnetOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOverrideEnable

`func (o *LanNetworkOpenApiV3VO) SetSubnetOverrideEnable(v bool)`

SetSubnetOverrideEnable sets SubnetOverrideEnable field to given value.

### HasSubnetOverrideEnable

`func (o *LanNetworkOpenApiV3VO) HasSubnetOverrideEnable() bool`

HasSubnetOverrideEnable returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkOpenApiV3VO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkOpenApiV3VO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkOpenApiV3VO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkOpenApiV3VO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkOpenApiV3VO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkOpenApiV3VO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkOpenApiV3VO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkOpenApiV3VO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetVlans

`func (o *LanNetworkOpenApiV3VO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *LanNetworkOpenApiV3VO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *LanNetworkOpenApiV3VO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *LanNetworkOpenApiV3VO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetVrfId

`func (o *LanNetworkOpenApiV3VO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *LanNetworkOpenApiV3VO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *LanNetworkOpenApiV3VO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *LanNetworkOpenApiV3VO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


