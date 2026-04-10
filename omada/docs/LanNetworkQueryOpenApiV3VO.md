# LanNetworkQueryOpenApiV3VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site ID | [optional] 
**AccessControlRule** | Pointer to **bool** | Show AccessControlRule is enabled or not | [optional] 
**AllLan** | Pointer to **bool** | When Internet pre-config is closed or Internet pre-config is Universal, allLAN is \&quot;true\&quot;; after adopting gateway, allLAN is \&quot;false\&quot;. | [optional] 
**Application** | Pointer to **int32** | Effective device type should be a value as follows: 0: Gateway and Switch; 1: Switch | [optional] 
**ArpDetectionEnable** | Pointer to **bool** | Enable arp detection. Only valid when deviceType is 1 and gateway supports this feature. | [optional] 
**DeviceMac** | Pointer to **string** | DHCP Server Device mac. Only valid when deviceType is 1 or 2. When deviceType is 1, deviceMac can be empty when there is no gateway in the site. | [optional] 
**DeviceType** | **int32** | DHCP Server Device type. It should be a value as follows: 0:External Device 1:gateway 2:switch 3:none | 
**DhcpGuard** | Pointer to [**DhcpServersSetting**](DhcpServersSetting.md) |  | [optional] 
**DhcpL2RelayEnable** | Pointer to **bool** | The switch of DHCP L2 relay | [optional] 
**DhcpRelay** | Pointer to [**OswDhcpRelayOpenApiVO**](OswDhcpRelayOpenApiVO.md) |  | [optional] 
**DhcpServer** | Pointer to [**OswDhcpServerOpenApiVO**](OswDhcpServerOpenApiVO.md) |  | [optional] 
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
**InterfaceIds** | Pointer to **[]string** | Gateway LAN port IDs (acquired from \&quot;Check WAN/LAN status\&quot;) | [optional] 
**Ip** | Pointer to [**OswIpSettingBriefOpenapiVO**](OswIpSettingBriefOpenapiVO.md) |  | [optional] 
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
**Purpose** | **int32** | LAN network purpose should be a value as follows: 0: VLAN; 1: interface | 
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

### NewLanNetworkQueryOpenApiV3VO

`func NewLanNetworkQueryOpenApiV3VO(deviceType int32, igmpSnoopEnable bool, name string, purpose int32, ) *LanNetworkQueryOpenApiV3VO`

NewLanNetworkQueryOpenApiV3VO instantiates a new LanNetworkQueryOpenApiV3VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkQueryOpenApiV3VOWithDefaults

`func NewLanNetworkQueryOpenApiV3VOWithDefaults() *LanNetworkQueryOpenApiV3VO`

NewLanNetworkQueryOpenApiV3VOWithDefaults instantiates a new LanNetworkQueryOpenApiV3VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *LanNetworkQueryOpenApiV3VO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *LanNetworkQueryOpenApiV3VO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *LanNetworkQueryOpenApiV3VO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *LanNetworkQueryOpenApiV3VO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccessControlRule

`func (o *LanNetworkQueryOpenApiV3VO) GetAccessControlRule() bool`

GetAccessControlRule returns the AccessControlRule field if non-nil, zero value otherwise.

### GetAccessControlRuleOk

`func (o *LanNetworkQueryOpenApiV3VO) GetAccessControlRuleOk() (*bool, bool)`

GetAccessControlRuleOk returns a tuple with the AccessControlRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlRule

`func (o *LanNetworkQueryOpenApiV3VO) SetAccessControlRule(v bool)`

SetAccessControlRule sets AccessControlRule field to given value.

### HasAccessControlRule

`func (o *LanNetworkQueryOpenApiV3VO) HasAccessControlRule() bool`

HasAccessControlRule returns a boolean if a field has been set.

### GetAllLan

`func (o *LanNetworkQueryOpenApiV3VO) GetAllLan() bool`

GetAllLan returns the AllLan field if non-nil, zero value otherwise.

### GetAllLanOk

`func (o *LanNetworkQueryOpenApiV3VO) GetAllLanOk() (*bool, bool)`

GetAllLanOk returns a tuple with the AllLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLan

`func (o *LanNetworkQueryOpenApiV3VO) SetAllLan(v bool)`

SetAllLan sets AllLan field to given value.

### HasAllLan

`func (o *LanNetworkQueryOpenApiV3VO) HasAllLan() bool`

HasAllLan returns a boolean if a field has been set.

### GetApplication

`func (o *LanNetworkQueryOpenApiV3VO) GetApplication() int32`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *LanNetworkQueryOpenApiV3VO) GetApplicationOk() (*int32, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *LanNetworkQueryOpenApiV3VO) SetApplication(v int32)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *LanNetworkQueryOpenApiV3VO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetArpDetectionEnable

`func (o *LanNetworkQueryOpenApiV3VO) GetArpDetectionEnable() bool`

GetArpDetectionEnable returns the ArpDetectionEnable field if non-nil, zero value otherwise.

### GetArpDetectionEnableOk

`func (o *LanNetworkQueryOpenApiV3VO) GetArpDetectionEnableOk() (*bool, bool)`

GetArpDetectionEnableOk returns a tuple with the ArpDetectionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpDetectionEnable

`func (o *LanNetworkQueryOpenApiV3VO) SetArpDetectionEnable(v bool)`

SetArpDetectionEnable sets ArpDetectionEnable field to given value.

### HasArpDetectionEnable

`func (o *LanNetworkQueryOpenApiV3VO) HasArpDetectionEnable() bool`

HasArpDetectionEnable returns a boolean if a field has been set.

### GetDeviceMac

`func (o *LanNetworkQueryOpenApiV3VO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *LanNetworkQueryOpenApiV3VO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *LanNetworkQueryOpenApiV3VO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceType

`func (o *LanNetworkQueryOpenApiV3VO) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *LanNetworkQueryOpenApiV3VO) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.


### GetDhcpGuard

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpGuard() DhcpServersSetting`

GetDhcpGuard returns the DhcpGuard field if non-nil, zero value otherwise.

### GetDhcpGuardOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpGuardOk() (*DhcpServersSetting, bool)`

GetDhcpGuardOk returns a tuple with the DhcpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpGuard

`func (o *LanNetworkQueryOpenApiV3VO) SetDhcpGuard(v DhcpServersSetting)`

SetDhcpGuard sets DhcpGuard field to given value.

### HasDhcpGuard

`func (o *LanNetworkQueryOpenApiV3VO) HasDhcpGuard() bool`

HasDhcpGuard returns a boolean if a field has been set.

### GetDhcpL2RelayEnable

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpL2RelayEnable() bool`

GetDhcpL2RelayEnable returns the DhcpL2RelayEnable field if non-nil, zero value otherwise.

### GetDhcpL2RelayEnableOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpL2RelayEnableOk() (*bool, bool)`

GetDhcpL2RelayEnableOk returns a tuple with the DhcpL2RelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelayEnable

`func (o *LanNetworkQueryOpenApiV3VO) SetDhcpL2RelayEnable(v bool)`

SetDhcpL2RelayEnable sets DhcpL2RelayEnable field to given value.

### HasDhcpL2RelayEnable

`func (o *LanNetworkQueryOpenApiV3VO) HasDhcpL2RelayEnable() bool`

HasDhcpL2RelayEnable returns a boolean if a field has been set.

### GetDhcpRelay

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpRelay() OswDhcpRelayOpenApiVO`

GetDhcpRelay returns the DhcpRelay field if non-nil, zero value otherwise.

### GetDhcpRelayOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpRelayOk() (*OswDhcpRelayOpenApiVO, bool)`

GetDhcpRelayOk returns a tuple with the DhcpRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelay

`func (o *LanNetworkQueryOpenApiV3VO) SetDhcpRelay(v OswDhcpRelayOpenApiVO)`

SetDhcpRelay sets DhcpRelay field to given value.

### HasDhcpRelay

`func (o *LanNetworkQueryOpenApiV3VO) HasDhcpRelay() bool`

HasDhcpRelay returns a boolean if a field has been set.

### GetDhcpServer

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpServer() OswDhcpServerOpenApiVO`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpServerOk() (*OswDhcpServerOpenApiVO, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *LanNetworkQueryOpenApiV3VO) SetDhcpServer(v OswDhcpServerOpenApiVO)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *LanNetworkQueryOpenApiV3VO) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpServerNum

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpServerNum() int32`

GetDhcpServerNum returns the DhcpServerNum field if non-nil, zero value otherwise.

### GetDhcpServerNumOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpServerNumOk() (*int32, bool)`

GetDhcpServerNumOk returns a tuple with the DhcpServerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerNum

`func (o *LanNetworkQueryOpenApiV3VO) SetDhcpServerNum(v int32)`

SetDhcpServerNum sets DhcpServerNum field to given value.

### HasDhcpServerNum

`func (o *LanNetworkQueryOpenApiV3VO) HasDhcpServerNum() bool`

HasDhcpServerNum returns a boolean if a field has been set.

### GetDhcpSettings

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpSettings() DhcpSettingInfo`

GetDhcpSettings returns the DhcpSettings field if non-nil, zero value otherwise.

### GetDhcpSettingsOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpSettingsOk() (*DhcpSettingInfo, bool)`

GetDhcpSettingsOk returns a tuple with the DhcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSettings

`func (o *LanNetworkQueryOpenApiV3VO) SetDhcpSettings(v DhcpSettingInfo)`

SetDhcpSettings sets DhcpSettings field to given value.

### HasDhcpSettings

`func (o *LanNetworkQueryOpenApiV3VO) HasDhcpSettings() bool`

HasDhcpSettings returns a boolean if a field has been set.

### GetDhcpv6Guard

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpv6Guard() Dhcpv6ServersSetting`

GetDhcpv6Guard returns the Dhcpv6Guard field if non-nil, zero value otherwise.

### GetDhcpv6GuardOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDhcpv6GuardOk() (*Dhcpv6ServersSetting, bool)`

GetDhcpv6GuardOk returns a tuple with the Dhcpv6Guard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Guard

`func (o *LanNetworkQueryOpenApiV3VO) SetDhcpv6Guard(v Dhcpv6ServersSetting)`

SetDhcpv6Guard sets Dhcpv6Guard field to given value.

### HasDhcpv6Guard

`func (o *LanNetworkQueryOpenApiV3VO) HasDhcpv6Guard() bool`

HasDhcpv6Guard returns a boolean if a field has been set.

### GetDomain

`func (o *LanNetworkQueryOpenApiV3VO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanNetworkQueryOpenApiV3VO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanNetworkQueryOpenApiV3VO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LanNetworkQueryOpenApiV3VO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExistArpDetection

`func (o *LanNetworkQueryOpenApiV3VO) GetExistArpDetection() bool`

GetExistArpDetection returns the ExistArpDetection field if non-nil, zero value otherwise.

### GetExistArpDetectionOk

`func (o *LanNetworkQueryOpenApiV3VO) GetExistArpDetectionOk() (*bool, bool)`

GetExistArpDetectionOk returns a tuple with the ExistArpDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistArpDetection

`func (o *LanNetworkQueryOpenApiV3VO) SetExistArpDetection(v bool)`

SetExistArpDetection sets ExistArpDetection field to given value.

### HasExistArpDetection

`func (o *LanNetworkQueryOpenApiV3VO) HasExistArpDetection() bool`

HasExistArpDetection returns a boolean if a field has been set.

### GetExistCustomDhcpOption

`func (o *LanNetworkQueryOpenApiV3VO) GetExistCustomDhcpOption() bool`

GetExistCustomDhcpOption returns the ExistCustomDhcpOption field if non-nil, zero value otherwise.

### GetExistCustomDhcpOptionOk

`func (o *LanNetworkQueryOpenApiV3VO) GetExistCustomDhcpOptionOk() (*bool, bool)`

GetExistCustomDhcpOptionOk returns a tuple with the ExistCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCustomDhcpOption

`func (o *LanNetworkQueryOpenApiV3VO) SetExistCustomDhcpOption(v bool)`

SetExistCustomDhcpOption sets ExistCustomDhcpOption field to given value.

### HasExistCustomDhcpOption

`func (o *LanNetworkQueryOpenApiV3VO) HasExistCustomDhcpOption() bool`

HasExistCustomDhcpOption returns a boolean if a field has been set.

### GetExistDhcpNextServer

`func (o *LanNetworkQueryOpenApiV3VO) GetExistDhcpNextServer() bool`

GetExistDhcpNextServer returns the ExistDhcpNextServer field if non-nil, zero value otherwise.

### GetExistDhcpNextServerOk

`func (o *LanNetworkQueryOpenApiV3VO) GetExistDhcpNextServerOk() (*bool, bool)`

GetExistDhcpNextServerOk returns a tuple with the ExistDhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistDhcpNextServer

`func (o *LanNetworkQueryOpenApiV3VO) SetExistDhcpNextServer(v bool)`

SetExistDhcpNextServer sets ExistDhcpNextServer field to given value.

### HasExistDhcpNextServer

`func (o *LanNetworkQueryOpenApiV3VO) HasExistDhcpNextServer() bool`

HasExistDhcpNextServer returns a boolean if a field has been set.

### GetExistMultiVlan

`func (o *LanNetworkQueryOpenApiV3VO) GetExistMultiVlan() bool`

GetExistMultiVlan returns the ExistMultiVlan field if non-nil, zero value otherwise.

### GetExistMultiVlanOk

`func (o *LanNetworkQueryOpenApiV3VO) GetExistMultiVlanOk() (*bool, bool)`

GetExistMultiVlanOk returns a tuple with the ExistMultiVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistMultiVlan

`func (o *LanNetworkQueryOpenApiV3VO) SetExistMultiVlan(v bool)`

SetExistMultiVlan sets ExistMultiVlan field to given value.

### HasExistMultiVlan

`func (o *LanNetworkQueryOpenApiV3VO) HasExistMultiVlan() bool`

HasExistMultiVlan returns a boolean if a field has been set.

### GetExistNetworkIsolation

`func (o *LanNetworkQueryOpenApiV3VO) GetExistNetworkIsolation() bool`

GetExistNetworkIsolation returns the ExistNetworkIsolation field if non-nil, zero value otherwise.

### GetExistNetworkIsolationOk

`func (o *LanNetworkQueryOpenApiV3VO) GetExistNetworkIsolationOk() (*bool, bool)`

GetExistNetworkIsolationOk returns a tuple with the ExistNetworkIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistNetworkIsolation

`func (o *LanNetworkQueryOpenApiV3VO) SetExistNetworkIsolation(v bool)`

SetExistNetworkIsolation sets ExistNetworkIsolation field to given value.

### HasExistNetworkIsolation

`func (o *LanNetworkQueryOpenApiV3VO) HasExistNetworkIsolation() bool`

HasExistNetworkIsolation returns a boolean if a field has been set.

### GetExistRA

`func (o *LanNetworkQueryOpenApiV3VO) GetExistRA() bool`

GetExistRA returns the ExistRA field if non-nil, zero value otherwise.

### GetExistRAOk

`func (o *LanNetworkQueryOpenApiV3VO) GetExistRAOk() (*bool, bool)`

GetExistRAOk returns a tuple with the ExistRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistRA

`func (o *LanNetworkQueryOpenApiV3VO) SetExistRA(v bool)`

SetExistRA sets ExistRA field to given value.

### HasExistRA

`func (o *LanNetworkQueryOpenApiV3VO) HasExistRA() bool`

HasExistRA returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *LanNetworkQueryOpenApiV3VO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *LanNetworkQueryOpenApiV3VO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *LanNetworkQueryOpenApiV3VO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *LanNetworkQueryOpenApiV3VO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *LanNetworkQueryOpenApiV3VO) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *LanNetworkQueryOpenApiV3VO) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *LanNetworkQueryOpenApiV3VO) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *LanNetworkQueryOpenApiV3VO) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetId

`func (o *LanNetworkQueryOpenApiV3VO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanNetworkQueryOpenApiV3VO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanNetworkQueryOpenApiV3VO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanNetworkQueryOpenApiV3VO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgmpSnoopEnable

`func (o *LanNetworkQueryOpenApiV3VO) GetIgmpSnoopEnable() bool`

GetIgmpSnoopEnable returns the IgmpSnoopEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopEnableOk

`func (o *LanNetworkQueryOpenApiV3VO) GetIgmpSnoopEnableOk() (*bool, bool)`

GetIgmpSnoopEnableOk returns a tuple with the IgmpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopEnable

`func (o *LanNetworkQueryOpenApiV3VO) SetIgmpSnoopEnable(v bool)`

SetIgmpSnoopEnable sets IgmpSnoopEnable field to given value.


### GetInterfaceIds

`func (o *LanNetworkQueryOpenApiV3VO) GetInterfaceIds() []string`

GetInterfaceIds returns the InterfaceIds field if non-nil, zero value otherwise.

### GetInterfaceIdsOk

`func (o *LanNetworkQueryOpenApiV3VO) GetInterfaceIdsOk() (*[]string, bool)`

GetInterfaceIdsOk returns a tuple with the InterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIds

`func (o *LanNetworkQueryOpenApiV3VO) SetInterfaceIds(v []string)`

SetInterfaceIds sets InterfaceIds field to given value.

### HasInterfaceIds

`func (o *LanNetworkQueryOpenApiV3VO) HasInterfaceIds() bool`

HasInterfaceIds returns a boolean if a field has been set.

### GetIp

`func (o *LanNetworkQueryOpenApiV3VO) GetIp() OswIpSettingBriefOpenapiVO`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LanNetworkQueryOpenApiV3VO) GetIpOk() (*OswIpSettingBriefOpenapiVO, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LanNetworkQueryOpenApiV3VO) SetIp(v OswIpSettingBriefOpenapiVO)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LanNetworkQueryOpenApiV3VO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsolation

`func (o *LanNetworkQueryOpenApiV3VO) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *LanNetworkQueryOpenApiV3VO) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *LanNetworkQueryOpenApiV3VO) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *LanNetworkQueryOpenApiV3VO) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetLanNeworkIpv6Config

`func (o *LanNetworkQueryOpenApiV3VO) GetLanNeworkIpv6Config() LanNetworkIPV6Config`

GetLanNeworkIpv6Config returns the LanNeworkIpv6Config field if non-nil, zero value otherwise.

### GetLanNeworkIpv6ConfigOk

`func (o *LanNetworkQueryOpenApiV3VO) GetLanNeworkIpv6ConfigOk() (*LanNetworkIPV6Config, bool)`

GetLanNeworkIpv6ConfigOk returns a tuple with the LanNeworkIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNeworkIpv6Config

`func (o *LanNetworkQueryOpenApiV3VO) SetLanNeworkIpv6Config(v LanNetworkIPV6Config)`

SetLanNeworkIpv6Config sets LanNeworkIpv6Config field to given value.

### HasLanNeworkIpv6Config

`func (o *LanNetworkQueryOpenApiV3VO) HasLanNeworkIpv6Config() bool`

HasLanNeworkIpv6Config returns a boolean if a field has been set.

### GetMldSnoopEnable

`func (o *LanNetworkQueryOpenApiV3VO) GetMldSnoopEnable() bool`

GetMldSnoopEnable returns the MldSnoopEnable field if non-nil, zero value otherwise.

### GetMldSnoopEnableOk

`func (o *LanNetworkQueryOpenApiV3VO) GetMldSnoopEnableOk() (*bool, bool)`

GetMldSnoopEnableOk returns a tuple with the MldSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldSnoopEnable

`func (o *LanNetworkQueryOpenApiV3VO) SetMldSnoopEnable(v bool)`

SetMldSnoopEnable sets MldSnoopEnable field to given value.

### HasMldSnoopEnable

`func (o *LanNetworkQueryOpenApiV3VO) HasMldSnoopEnable() bool`

HasMldSnoopEnable returns a boolean if a field has been set.

### GetMode

`func (o *LanNetworkQueryOpenApiV3VO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *LanNetworkQueryOpenApiV3VO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *LanNetworkQueryOpenApiV3VO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *LanNetworkQueryOpenApiV3VO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkQueryOpenApiV3VO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkQueryOpenApiV3VO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkQueryOpenApiV3VO) SetName(v string)`

SetName sets Name field to given value.


### GetOrigName

`func (o *LanNetworkQueryOpenApiV3VO) GetOrigName() string`

GetOrigName returns the OrigName field if non-nil, zero value otherwise.

### GetOrigNameOk

`func (o *LanNetworkQueryOpenApiV3VO) GetOrigNameOk() (*string, bool)`

GetOrigNameOk returns a tuple with the OrigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigName

`func (o *LanNetworkQueryOpenApiV3VO) SetOrigName(v string)`

SetOrigName sets OrigName field to given value.

### HasOrigName

`func (o *LanNetworkQueryOpenApiV3VO) HasOrigName() bool`

HasOrigName returns a boolean if a field has been set.

### GetPortal

`func (o *LanNetworkQueryOpenApiV3VO) GetPortal() bool`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *LanNetworkQueryOpenApiV3VO) GetPortalOk() (*bool, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *LanNetworkQueryOpenApiV3VO) SetPortal(v bool)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *LanNetworkQueryOpenApiV3VO) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetPortalId

`func (o *LanNetworkQueryOpenApiV3VO) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *LanNetworkQueryOpenApiV3VO) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *LanNetworkQueryOpenApiV3VO) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *LanNetworkQueryOpenApiV3VO) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetPortalName

`func (o *LanNetworkQueryOpenApiV3VO) GetPortalName() string`

GetPortalName returns the PortalName field if non-nil, zero value otherwise.

### GetPortalNameOk

`func (o *LanNetworkQueryOpenApiV3VO) GetPortalNameOk() (*string, bool)`

GetPortalNameOk returns a tuple with the PortalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalName

`func (o *LanNetworkQueryOpenApiV3VO) SetPortalName(v string)`

SetPortalName sets PortalName field to given value.

### HasPortalName

`func (o *LanNetworkQueryOpenApiV3VO) HasPortalName() bool`

HasPortalName returns a boolean if a field has been set.

### GetPrimary

`func (o *LanNetworkQueryOpenApiV3VO) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *LanNetworkQueryOpenApiV3VO) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *LanNetworkQueryOpenApiV3VO) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *LanNetworkQueryOpenApiV3VO) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPurpose

`func (o *LanNetworkQueryOpenApiV3VO) GetPurpose() int32`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *LanNetworkQueryOpenApiV3VO) GetPurposeOk() (*int32, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *LanNetworkQueryOpenApiV3VO) SetPurpose(v int32)`

SetPurpose sets Purpose field to given value.


### GetQosQueueEnable

`func (o *LanNetworkQueryOpenApiV3VO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *LanNetworkQueryOpenApiV3VO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *LanNetworkQueryOpenApiV3VO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *LanNetworkQueryOpenApiV3VO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQueueId

`func (o *LanNetworkQueryOpenApiV3VO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *LanNetworkQueryOpenApiV3VO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *LanNetworkQueryOpenApiV3VO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *LanNetworkQueryOpenApiV3VO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetRateLimit

`func (o *LanNetworkQueryOpenApiV3VO) GetRateLimit() bool`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *LanNetworkQueryOpenApiV3VO) GetRateLimitOk() (*bool, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *LanNetworkQueryOpenApiV3VO) SetRateLimit(v bool)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *LanNetworkQueryOpenApiV3VO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetResource

`func (o *LanNetworkQueryOpenApiV3VO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *LanNetworkQueryOpenApiV3VO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *LanNetworkQueryOpenApiV3VO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *LanNetworkQueryOpenApiV3VO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStackId

`func (o *LanNetworkQueryOpenApiV3VO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *LanNetworkQueryOpenApiV3VO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *LanNetworkQueryOpenApiV3VO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *LanNetworkQueryOpenApiV3VO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetState

`func (o *LanNetworkQueryOpenApiV3VO) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LanNetworkQueryOpenApiV3VO) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LanNetworkQueryOpenApiV3VO) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *LanNetworkQueryOpenApiV3VO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubnetOverride

`func (o *LanNetworkQueryOpenApiV3VO) GetSubnetOverride() bool`

GetSubnetOverride returns the SubnetOverride field if non-nil, zero value otherwise.

### GetSubnetOverrideOk

`func (o *LanNetworkQueryOpenApiV3VO) GetSubnetOverrideOk() (*bool, bool)`

GetSubnetOverrideOk returns a tuple with the SubnetOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOverride

`func (o *LanNetworkQueryOpenApiV3VO) SetSubnetOverride(v bool)`

SetSubnetOverride sets SubnetOverride field to given value.

### HasSubnetOverride

`func (o *LanNetworkQueryOpenApiV3VO) HasSubnetOverride() bool`

HasSubnetOverride returns a boolean if a field has been set.

### GetSubnetOverrideEnable

`func (o *LanNetworkQueryOpenApiV3VO) GetSubnetOverrideEnable() bool`

GetSubnetOverrideEnable returns the SubnetOverrideEnable field if non-nil, zero value otherwise.

### GetSubnetOverrideEnableOk

`func (o *LanNetworkQueryOpenApiV3VO) GetSubnetOverrideEnableOk() (*bool, bool)`

GetSubnetOverrideEnableOk returns a tuple with the SubnetOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOverrideEnable

`func (o *LanNetworkQueryOpenApiV3VO) SetSubnetOverrideEnable(v bool)`

SetSubnetOverrideEnable sets SubnetOverrideEnable field to given value.

### HasSubnetOverrideEnable

`func (o *LanNetworkQueryOpenApiV3VO) HasSubnetOverrideEnable() bool`

HasSubnetOverrideEnable returns a boolean if a field has been set.

### GetTotalIpNum

`func (o *LanNetworkQueryOpenApiV3VO) GetTotalIpNum() int64`

GetTotalIpNum returns the TotalIpNum field if non-nil, zero value otherwise.

### GetTotalIpNumOk

`func (o *LanNetworkQueryOpenApiV3VO) GetTotalIpNumOk() (*int64, bool)`

GetTotalIpNumOk returns a tuple with the TotalIpNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIpNum

`func (o *LanNetworkQueryOpenApiV3VO) SetTotalIpNum(v int64)`

SetTotalIpNum sets TotalIpNum field to given value.

### HasTotalIpNum

`func (o *LanNetworkQueryOpenApiV3VO) HasTotalIpNum() bool`

HasTotalIpNum returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkQueryOpenApiV3VO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkQueryOpenApiV3VO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkQueryOpenApiV3VO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkQueryOpenApiV3VO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkQueryOpenApiV3VO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkQueryOpenApiV3VO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkQueryOpenApiV3VO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkQueryOpenApiV3VO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetVlans

`func (o *LanNetworkQueryOpenApiV3VO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *LanNetworkQueryOpenApiV3VO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *LanNetworkQueryOpenApiV3VO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *LanNetworkQueryOpenApiV3VO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetVrfId

`func (o *LanNetworkQueryOpenApiV3VO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *LanNetworkQueryOpenApiV3VO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *LanNetworkQueryOpenApiV3VO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *LanNetworkQueryOpenApiV3VO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


