# LanNetworkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site ID | [optional] 
**AccessControlRule** | Pointer to **bool** | Show AccessControlRule is enabled or not | [optional] 
**AllLan** | Pointer to **bool** | When Internet pre-config is closed or Internet pre-config is Universal, allLAN is \&quot;true\&quot;; after adopting gateway, allLAN is \&quot;false\&quot;. | [optional] 
**Application** | Pointer to **int32** | Effective device type should be a value as follows: 0: Gateway and Switch; 1: Switch | [optional] 
**ArpDetectionEnable** | Pointer to **bool** | Enable arp detection. Only valid when deviceType is 1 and gateway supports this feature. | [optional] 
**DeviceMac** | Pointer to **string** | DHCP Server Device mac. Only valid when deviceType is 1 or 2. When deviceType is 1, deviceMac can be empty when there is no gateway in the site. | [optional] 
**DeviceType** | Pointer to **int32** | DHCP Server Device type. It should be a value as follows: 0:External Device 1:gateway 2:switch 3:none | [optional] 
**DhcpGuard** | Pointer to [**DhcpServersSetting**](DhcpServersSetting.md) |  | [optional] 
**DhcpL2RelayEnable** | Pointer to **bool** | The switch of DHCP L2 relay | [optional] 
**DhcpRelay** | Pointer to [**OswDhcpRelayVO**](OswDhcpRelayVO.md) |  | [optional] 
**DhcpServer** | Pointer to [**OswDhcpServerVO**](OswDhcpServerVO.md) |  | [optional] 
**DhcpServerNum** | Pointer to **int32** | The number of dhcp server devices in effect, Only valid when vlanType is 0. | [optional] 
**DhcpSettings** | Pointer to [**DhcpSettingInfo**](DhcpSettingInfo.md) |  | [optional] 
**Dhcpv6Guard** | Pointer to [**Dhcpv6ServersSetting**](Dhcpv6ServersSetting.md) |  | [optional] 
**Domain** | Pointer to **string** | The domain of this network | [optional] 
**ExistArpDetection** | Pointer to **bool** | Whether Arp Detection is configured. | [optional] 
**ExistCustomDhcpOption** | Pointer to **bool** | Whether custom DHCP Options has been configured. | [optional] 
**ExistDhcpNextServer** | Pointer to **bool** | Whether DHCP Next Server has been configured. | [optional] 
**ExistMultiVlan** | Pointer to **bool** | Whether VLAN Type is Multiple. | [optional] 
**ExistNetworkIsolation** | Pointer to **bool** | Whether Network Isolation is configured. | [optional] 
**ExistRA** | Pointer to **bool** | Whether RA has been configured. | [optional] 
**FastLeaveEnable** | Pointer to **bool** | IGMP Snooping fast leave enable status | [optional] 
**GatewaySubnet** | Pointer to **string** | When purpose is interface, gateway subnet is needed. Format: IP/Mask | [optional] 
**Id** | Pointer to **string** | LAN network ID | [optional] 
**IgmpSnoopEnable** | **bool** | Enable IGMP snooping | 
**Interface** | Pointer to **bool** |  | [optional] 
**InterfaceIds** | Pointer to **[]string** | Gateway LAN port IDs (acquired from \&quot;Check WAN/LAN status\&quot;) | [optional] 
**Ip** | Pointer to [**OswIpSettingVO**](OswIpSettingVO.md) |  | [optional] 
**Isolation** | Pointer to **bool** | Whether network isolated. | [optional] 
**LanNeworkIpv6Config** | Pointer to [**LanNetworkIPV6Config**](LanNetworkIPV6Config.md) |  | [optional] 
**MldSnoopEnable** | Pointer to **bool** | Enable MLD snooping | [optional] 
**Mode** | Pointer to **int32** | DHCP mode. 0: None 1: DHCP Server 2: DHCP Relay. Only valid when deviceType is 2. | [optional] 
**Name** | **string** | LAN network name should contain 1 to 128 characters. | 
**OrigName** | Pointer to **string** | Original name | [optional] 
**Portal** | Pointer to **bool** | Show portal is enabled or not | [optional] 
**PortalId** | Pointer to **string** | Show portal ID | [optional] 
**PortalName** | Pointer to **string** | Show related portal name | [optional] 
**Primary** | Pointer to **bool** | Primary | [optional] 
**Purpose** | **string** | LAN network purpose should be as follows: vlan or interface | 
**QosQueueEnable** | Pointer to **bool** | The switch of QoS queue. | [optional] 
**QueueId** | Pointer to **int32** | QoS queue Id. | [optional] 
**RateLimit** | Pointer to **bool** | Show RateLimit is enabled or not | [optional] 
**Resource** | Pointer to **int32** | Resource is a value as follows: 0: new created; 1: from template; 2: override | [optional] 
**StackId** | Pointer to **string** | DHCP Server Device stackId. Only valid when deviceType is 2 and the selected device is stack | [optional] 
**State** | Pointer to **int32** | Network delivering state. It should be a value as follows: 0: not in delivering 1:delivering 2. deliver done | [optional] 
**SubnetOverride** | Pointer to **bool** | Subnet override | [optional] 
**SubnetOverrideEnable** | Pointer to **bool** | Subnet override enable status | [optional] 
**TotalIpNum** | Pointer to **int64** | Total ip num | [optional] 
**Vlan** | Pointer to **int32** | Only Valid when vlanType is 0. Vlan should be within the range of 1-4094. | [optional] 
**VlanType** | Pointer to **int32** | When purpose is interface, VLANType should be a value as follows: 0: Single; 1: Multiple | [optional] 
**Vlans** | Pointer to **string** | When purpose is interface and VLANType is 1, batch create VLANs. VLAN format: 200, 1-100. | [optional] 
**VrfId** | Pointer to **string** | VRF ID | [optional] 

## Methods

### NewLanNetworkVO

`func NewLanNetworkVO(igmpSnoopEnable bool, name string, purpose string, ) *LanNetworkVO`

NewLanNetworkVO instantiates a new LanNetworkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkVOWithDefaults

`func NewLanNetworkVOWithDefaults() *LanNetworkVO`

NewLanNetworkVOWithDefaults instantiates a new LanNetworkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *LanNetworkVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *LanNetworkVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *LanNetworkVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *LanNetworkVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccessControlRule

`func (o *LanNetworkVO) GetAccessControlRule() bool`

GetAccessControlRule returns the AccessControlRule field if non-nil, zero value otherwise.

### GetAccessControlRuleOk

`func (o *LanNetworkVO) GetAccessControlRuleOk() (*bool, bool)`

GetAccessControlRuleOk returns a tuple with the AccessControlRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlRule

`func (o *LanNetworkVO) SetAccessControlRule(v bool)`

SetAccessControlRule sets AccessControlRule field to given value.

### HasAccessControlRule

`func (o *LanNetworkVO) HasAccessControlRule() bool`

HasAccessControlRule returns a boolean if a field has been set.

### GetAllLan

`func (o *LanNetworkVO) GetAllLan() bool`

GetAllLan returns the AllLan field if non-nil, zero value otherwise.

### GetAllLanOk

`func (o *LanNetworkVO) GetAllLanOk() (*bool, bool)`

GetAllLanOk returns a tuple with the AllLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLan

`func (o *LanNetworkVO) SetAllLan(v bool)`

SetAllLan sets AllLan field to given value.

### HasAllLan

`func (o *LanNetworkVO) HasAllLan() bool`

HasAllLan returns a boolean if a field has been set.

### GetApplication

`func (o *LanNetworkVO) GetApplication() int32`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *LanNetworkVO) GetApplicationOk() (*int32, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *LanNetworkVO) SetApplication(v int32)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *LanNetworkVO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetArpDetectionEnable

`func (o *LanNetworkVO) GetArpDetectionEnable() bool`

GetArpDetectionEnable returns the ArpDetectionEnable field if non-nil, zero value otherwise.

### GetArpDetectionEnableOk

`func (o *LanNetworkVO) GetArpDetectionEnableOk() (*bool, bool)`

GetArpDetectionEnableOk returns a tuple with the ArpDetectionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpDetectionEnable

`func (o *LanNetworkVO) SetArpDetectionEnable(v bool)`

SetArpDetectionEnable sets ArpDetectionEnable field to given value.

### HasArpDetectionEnable

`func (o *LanNetworkVO) HasArpDetectionEnable() bool`

HasArpDetectionEnable returns a boolean if a field has been set.

### GetDeviceMac

`func (o *LanNetworkVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *LanNetworkVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *LanNetworkVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *LanNetworkVO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceType

`func (o *LanNetworkVO) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *LanNetworkVO) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *LanNetworkVO) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *LanNetworkVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDhcpGuard

`func (o *LanNetworkVO) GetDhcpGuard() DhcpServersSetting`

GetDhcpGuard returns the DhcpGuard field if non-nil, zero value otherwise.

### GetDhcpGuardOk

`func (o *LanNetworkVO) GetDhcpGuardOk() (*DhcpServersSetting, bool)`

GetDhcpGuardOk returns a tuple with the DhcpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpGuard

`func (o *LanNetworkVO) SetDhcpGuard(v DhcpServersSetting)`

SetDhcpGuard sets DhcpGuard field to given value.

### HasDhcpGuard

`func (o *LanNetworkVO) HasDhcpGuard() bool`

HasDhcpGuard returns a boolean if a field has been set.

### GetDhcpL2RelayEnable

`func (o *LanNetworkVO) GetDhcpL2RelayEnable() bool`

GetDhcpL2RelayEnable returns the DhcpL2RelayEnable field if non-nil, zero value otherwise.

### GetDhcpL2RelayEnableOk

`func (o *LanNetworkVO) GetDhcpL2RelayEnableOk() (*bool, bool)`

GetDhcpL2RelayEnableOk returns a tuple with the DhcpL2RelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelayEnable

`func (o *LanNetworkVO) SetDhcpL2RelayEnable(v bool)`

SetDhcpL2RelayEnable sets DhcpL2RelayEnable field to given value.

### HasDhcpL2RelayEnable

`func (o *LanNetworkVO) HasDhcpL2RelayEnable() bool`

HasDhcpL2RelayEnable returns a boolean if a field has been set.

### GetDhcpRelay

`func (o *LanNetworkVO) GetDhcpRelay() OswDhcpRelayVO`

GetDhcpRelay returns the DhcpRelay field if non-nil, zero value otherwise.

### GetDhcpRelayOk

`func (o *LanNetworkVO) GetDhcpRelayOk() (*OswDhcpRelayVO, bool)`

GetDhcpRelayOk returns a tuple with the DhcpRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelay

`func (o *LanNetworkVO) SetDhcpRelay(v OswDhcpRelayVO)`

SetDhcpRelay sets DhcpRelay field to given value.

### HasDhcpRelay

`func (o *LanNetworkVO) HasDhcpRelay() bool`

HasDhcpRelay returns a boolean if a field has been set.

### GetDhcpServer

`func (o *LanNetworkVO) GetDhcpServer() OswDhcpServerVO`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *LanNetworkVO) GetDhcpServerOk() (*OswDhcpServerVO, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *LanNetworkVO) SetDhcpServer(v OswDhcpServerVO)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *LanNetworkVO) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpServerNum

`func (o *LanNetworkVO) GetDhcpServerNum() int32`

GetDhcpServerNum returns the DhcpServerNum field if non-nil, zero value otherwise.

### GetDhcpServerNumOk

`func (o *LanNetworkVO) GetDhcpServerNumOk() (*int32, bool)`

GetDhcpServerNumOk returns a tuple with the DhcpServerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerNum

`func (o *LanNetworkVO) SetDhcpServerNum(v int32)`

SetDhcpServerNum sets DhcpServerNum field to given value.

### HasDhcpServerNum

`func (o *LanNetworkVO) HasDhcpServerNum() bool`

HasDhcpServerNum returns a boolean if a field has been set.

### GetDhcpSettings

`func (o *LanNetworkVO) GetDhcpSettings() DhcpSettingInfo`

GetDhcpSettings returns the DhcpSettings field if non-nil, zero value otherwise.

### GetDhcpSettingsOk

`func (o *LanNetworkVO) GetDhcpSettingsOk() (*DhcpSettingInfo, bool)`

GetDhcpSettingsOk returns a tuple with the DhcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSettings

`func (o *LanNetworkVO) SetDhcpSettings(v DhcpSettingInfo)`

SetDhcpSettings sets DhcpSettings field to given value.

### HasDhcpSettings

`func (o *LanNetworkVO) HasDhcpSettings() bool`

HasDhcpSettings returns a boolean if a field has been set.

### GetDhcpv6Guard

`func (o *LanNetworkVO) GetDhcpv6Guard() Dhcpv6ServersSetting`

GetDhcpv6Guard returns the Dhcpv6Guard field if non-nil, zero value otherwise.

### GetDhcpv6GuardOk

`func (o *LanNetworkVO) GetDhcpv6GuardOk() (*Dhcpv6ServersSetting, bool)`

GetDhcpv6GuardOk returns a tuple with the Dhcpv6Guard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Guard

`func (o *LanNetworkVO) SetDhcpv6Guard(v Dhcpv6ServersSetting)`

SetDhcpv6Guard sets Dhcpv6Guard field to given value.

### HasDhcpv6Guard

`func (o *LanNetworkVO) HasDhcpv6Guard() bool`

HasDhcpv6Guard returns a boolean if a field has been set.

### GetDomain

`func (o *LanNetworkVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanNetworkVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanNetworkVO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LanNetworkVO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExistArpDetection

`func (o *LanNetworkVO) GetExistArpDetection() bool`

GetExistArpDetection returns the ExistArpDetection field if non-nil, zero value otherwise.

### GetExistArpDetectionOk

`func (o *LanNetworkVO) GetExistArpDetectionOk() (*bool, bool)`

GetExistArpDetectionOk returns a tuple with the ExistArpDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistArpDetection

`func (o *LanNetworkVO) SetExistArpDetection(v bool)`

SetExistArpDetection sets ExistArpDetection field to given value.

### HasExistArpDetection

`func (o *LanNetworkVO) HasExistArpDetection() bool`

HasExistArpDetection returns a boolean if a field has been set.

### GetExistCustomDhcpOption

`func (o *LanNetworkVO) GetExistCustomDhcpOption() bool`

GetExistCustomDhcpOption returns the ExistCustomDhcpOption field if non-nil, zero value otherwise.

### GetExistCustomDhcpOptionOk

`func (o *LanNetworkVO) GetExistCustomDhcpOptionOk() (*bool, bool)`

GetExistCustomDhcpOptionOk returns a tuple with the ExistCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCustomDhcpOption

`func (o *LanNetworkVO) SetExistCustomDhcpOption(v bool)`

SetExistCustomDhcpOption sets ExistCustomDhcpOption field to given value.

### HasExistCustomDhcpOption

`func (o *LanNetworkVO) HasExistCustomDhcpOption() bool`

HasExistCustomDhcpOption returns a boolean if a field has been set.

### GetExistDhcpNextServer

`func (o *LanNetworkVO) GetExistDhcpNextServer() bool`

GetExistDhcpNextServer returns the ExistDhcpNextServer field if non-nil, zero value otherwise.

### GetExistDhcpNextServerOk

`func (o *LanNetworkVO) GetExistDhcpNextServerOk() (*bool, bool)`

GetExistDhcpNextServerOk returns a tuple with the ExistDhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistDhcpNextServer

`func (o *LanNetworkVO) SetExistDhcpNextServer(v bool)`

SetExistDhcpNextServer sets ExistDhcpNextServer field to given value.

### HasExistDhcpNextServer

`func (o *LanNetworkVO) HasExistDhcpNextServer() bool`

HasExistDhcpNextServer returns a boolean if a field has been set.

### GetExistMultiVlan

`func (o *LanNetworkVO) GetExistMultiVlan() bool`

GetExistMultiVlan returns the ExistMultiVlan field if non-nil, zero value otherwise.

### GetExistMultiVlanOk

`func (o *LanNetworkVO) GetExistMultiVlanOk() (*bool, bool)`

GetExistMultiVlanOk returns a tuple with the ExistMultiVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistMultiVlan

`func (o *LanNetworkVO) SetExistMultiVlan(v bool)`

SetExistMultiVlan sets ExistMultiVlan field to given value.

### HasExistMultiVlan

`func (o *LanNetworkVO) HasExistMultiVlan() bool`

HasExistMultiVlan returns a boolean if a field has been set.

### GetExistNetworkIsolation

`func (o *LanNetworkVO) GetExistNetworkIsolation() bool`

GetExistNetworkIsolation returns the ExistNetworkIsolation field if non-nil, zero value otherwise.

### GetExistNetworkIsolationOk

`func (o *LanNetworkVO) GetExistNetworkIsolationOk() (*bool, bool)`

GetExistNetworkIsolationOk returns a tuple with the ExistNetworkIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistNetworkIsolation

`func (o *LanNetworkVO) SetExistNetworkIsolation(v bool)`

SetExistNetworkIsolation sets ExistNetworkIsolation field to given value.

### HasExistNetworkIsolation

`func (o *LanNetworkVO) HasExistNetworkIsolation() bool`

HasExistNetworkIsolation returns a boolean if a field has been set.

### GetExistRA

`func (o *LanNetworkVO) GetExistRA() bool`

GetExistRA returns the ExistRA field if non-nil, zero value otherwise.

### GetExistRAOk

`func (o *LanNetworkVO) GetExistRAOk() (*bool, bool)`

GetExistRAOk returns a tuple with the ExistRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistRA

`func (o *LanNetworkVO) SetExistRA(v bool)`

SetExistRA sets ExistRA field to given value.

### HasExistRA

`func (o *LanNetworkVO) HasExistRA() bool`

HasExistRA returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *LanNetworkVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *LanNetworkVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *LanNetworkVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *LanNetworkVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *LanNetworkVO) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *LanNetworkVO) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *LanNetworkVO) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *LanNetworkVO) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetId

`func (o *LanNetworkVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanNetworkVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanNetworkVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanNetworkVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgmpSnoopEnable

`func (o *LanNetworkVO) GetIgmpSnoopEnable() bool`

GetIgmpSnoopEnable returns the IgmpSnoopEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopEnableOk

`func (o *LanNetworkVO) GetIgmpSnoopEnableOk() (*bool, bool)`

GetIgmpSnoopEnableOk returns a tuple with the IgmpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopEnable

`func (o *LanNetworkVO) SetIgmpSnoopEnable(v bool)`

SetIgmpSnoopEnable sets IgmpSnoopEnable field to given value.


### GetInterface

`func (o *LanNetworkVO) GetInterface() bool`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *LanNetworkVO) GetInterfaceOk() (*bool, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *LanNetworkVO) SetInterface(v bool)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *LanNetworkVO) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetInterfaceIds

`func (o *LanNetworkVO) GetInterfaceIds() []string`

GetInterfaceIds returns the InterfaceIds field if non-nil, zero value otherwise.

### GetInterfaceIdsOk

`func (o *LanNetworkVO) GetInterfaceIdsOk() (*[]string, bool)`

GetInterfaceIdsOk returns a tuple with the InterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIds

`func (o *LanNetworkVO) SetInterfaceIds(v []string)`

SetInterfaceIds sets InterfaceIds field to given value.

### HasInterfaceIds

`func (o *LanNetworkVO) HasInterfaceIds() bool`

HasInterfaceIds returns a boolean if a field has been set.

### GetIp

`func (o *LanNetworkVO) GetIp() OswIpSettingVO`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LanNetworkVO) GetIpOk() (*OswIpSettingVO, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LanNetworkVO) SetIp(v OswIpSettingVO)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LanNetworkVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsolation

`func (o *LanNetworkVO) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *LanNetworkVO) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *LanNetworkVO) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *LanNetworkVO) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetLanNeworkIpv6Config

`func (o *LanNetworkVO) GetLanNeworkIpv6Config() LanNetworkIPV6Config`

GetLanNeworkIpv6Config returns the LanNeworkIpv6Config field if non-nil, zero value otherwise.

### GetLanNeworkIpv6ConfigOk

`func (o *LanNetworkVO) GetLanNeworkIpv6ConfigOk() (*LanNetworkIPV6Config, bool)`

GetLanNeworkIpv6ConfigOk returns a tuple with the LanNeworkIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNeworkIpv6Config

`func (o *LanNetworkVO) SetLanNeworkIpv6Config(v LanNetworkIPV6Config)`

SetLanNeworkIpv6Config sets LanNeworkIpv6Config field to given value.

### HasLanNeworkIpv6Config

`func (o *LanNetworkVO) HasLanNeworkIpv6Config() bool`

HasLanNeworkIpv6Config returns a boolean if a field has been set.

### GetMldSnoopEnable

`func (o *LanNetworkVO) GetMldSnoopEnable() bool`

GetMldSnoopEnable returns the MldSnoopEnable field if non-nil, zero value otherwise.

### GetMldSnoopEnableOk

`func (o *LanNetworkVO) GetMldSnoopEnableOk() (*bool, bool)`

GetMldSnoopEnableOk returns a tuple with the MldSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldSnoopEnable

`func (o *LanNetworkVO) SetMldSnoopEnable(v bool)`

SetMldSnoopEnable sets MldSnoopEnable field to given value.

### HasMldSnoopEnable

`func (o *LanNetworkVO) HasMldSnoopEnable() bool`

HasMldSnoopEnable returns a boolean if a field has been set.

### GetMode

`func (o *LanNetworkVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *LanNetworkVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *LanNetworkVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *LanNetworkVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkVO) SetName(v string)`

SetName sets Name field to given value.


### GetOrigName

`func (o *LanNetworkVO) GetOrigName() string`

GetOrigName returns the OrigName field if non-nil, zero value otherwise.

### GetOrigNameOk

`func (o *LanNetworkVO) GetOrigNameOk() (*string, bool)`

GetOrigNameOk returns a tuple with the OrigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigName

`func (o *LanNetworkVO) SetOrigName(v string)`

SetOrigName sets OrigName field to given value.

### HasOrigName

`func (o *LanNetworkVO) HasOrigName() bool`

HasOrigName returns a boolean if a field has been set.

### GetPortal

`func (o *LanNetworkVO) GetPortal() bool`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *LanNetworkVO) GetPortalOk() (*bool, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *LanNetworkVO) SetPortal(v bool)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *LanNetworkVO) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetPortalId

`func (o *LanNetworkVO) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *LanNetworkVO) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *LanNetworkVO) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *LanNetworkVO) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetPortalName

`func (o *LanNetworkVO) GetPortalName() string`

GetPortalName returns the PortalName field if non-nil, zero value otherwise.

### GetPortalNameOk

`func (o *LanNetworkVO) GetPortalNameOk() (*string, bool)`

GetPortalNameOk returns a tuple with the PortalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalName

`func (o *LanNetworkVO) SetPortalName(v string)`

SetPortalName sets PortalName field to given value.

### HasPortalName

`func (o *LanNetworkVO) HasPortalName() bool`

HasPortalName returns a boolean if a field has been set.

### GetPrimary

`func (o *LanNetworkVO) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *LanNetworkVO) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *LanNetworkVO) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *LanNetworkVO) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPurpose

`func (o *LanNetworkVO) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *LanNetworkVO) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *LanNetworkVO) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetQosQueueEnable

`func (o *LanNetworkVO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *LanNetworkVO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *LanNetworkVO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *LanNetworkVO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQueueId

`func (o *LanNetworkVO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *LanNetworkVO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *LanNetworkVO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *LanNetworkVO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetRateLimit

`func (o *LanNetworkVO) GetRateLimit() bool`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *LanNetworkVO) GetRateLimitOk() (*bool, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *LanNetworkVO) SetRateLimit(v bool)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *LanNetworkVO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetResource

`func (o *LanNetworkVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *LanNetworkVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *LanNetworkVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *LanNetworkVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStackId

`func (o *LanNetworkVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *LanNetworkVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *LanNetworkVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *LanNetworkVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetState

`func (o *LanNetworkVO) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LanNetworkVO) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LanNetworkVO) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *LanNetworkVO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubnetOverride

`func (o *LanNetworkVO) GetSubnetOverride() bool`

GetSubnetOverride returns the SubnetOverride field if non-nil, zero value otherwise.

### GetSubnetOverrideOk

`func (o *LanNetworkVO) GetSubnetOverrideOk() (*bool, bool)`

GetSubnetOverrideOk returns a tuple with the SubnetOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOverride

`func (o *LanNetworkVO) SetSubnetOverride(v bool)`

SetSubnetOverride sets SubnetOverride field to given value.

### HasSubnetOverride

`func (o *LanNetworkVO) HasSubnetOverride() bool`

HasSubnetOverride returns a boolean if a field has been set.

### GetSubnetOverrideEnable

`func (o *LanNetworkVO) GetSubnetOverrideEnable() bool`

GetSubnetOverrideEnable returns the SubnetOverrideEnable field if non-nil, zero value otherwise.

### GetSubnetOverrideEnableOk

`func (o *LanNetworkVO) GetSubnetOverrideEnableOk() (*bool, bool)`

GetSubnetOverrideEnableOk returns a tuple with the SubnetOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOverrideEnable

`func (o *LanNetworkVO) SetSubnetOverrideEnable(v bool)`

SetSubnetOverrideEnable sets SubnetOverrideEnable field to given value.

### HasSubnetOverrideEnable

`func (o *LanNetworkVO) HasSubnetOverrideEnable() bool`

HasSubnetOverrideEnable returns a boolean if a field has been set.

### GetTotalIpNum

`func (o *LanNetworkVO) GetTotalIpNum() int64`

GetTotalIpNum returns the TotalIpNum field if non-nil, zero value otherwise.

### GetTotalIpNumOk

`func (o *LanNetworkVO) GetTotalIpNumOk() (*int64, bool)`

GetTotalIpNumOk returns a tuple with the TotalIpNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIpNum

`func (o *LanNetworkVO) SetTotalIpNum(v int64)`

SetTotalIpNum sets TotalIpNum field to given value.

### HasTotalIpNum

`func (o *LanNetworkVO) HasTotalIpNum() bool`

HasTotalIpNum returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkVO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkVO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkVO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkVO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetVlans

`func (o *LanNetworkVO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *LanNetworkVO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *LanNetworkVO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *LanNetworkVO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetVrfId

`func (o *LanNetworkVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *LanNetworkVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *LanNetworkVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *LanNetworkVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


