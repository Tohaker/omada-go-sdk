# LanNetworkQueryOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlRule** | Pointer to **bool** | Show AccessControlRule is enabled or not | [optional] 
**AllLan** | Pointer to **bool** | When Internet pre-config is closed or Internet pre-config is Universal, allLAN is \&quot;true\&quot;; after adopting gateway, allLAN is \&quot;false\&quot;. | [optional] 
**Application** | Pointer to **int32** | Effective device type should be a value as follows: 0: Gateway and Switch; 1: Switch | [optional] 
**DhcpGuard** | Pointer to [**DhcpServersSetting**](DhcpServersSetting.md) |  | [optional] 
**DhcpL2RelayEnable** | Pointer to **bool** | The switch of DHCP L2 relay | [optional] 
**DhcpSettingsVO** | Pointer to [**DhcpSettingInfo**](DhcpSettingInfo.md) |  | [optional] 
**Dhcpv6Guard** | Pointer to [**Dhcpv6ServersSetting**](Dhcpv6ServersSetting.md) |  | [optional] 
**Domain** | Pointer to **string** | The domain of this network | [optional] 
**ExistArpDetection** | Pointer to **bool** | Whether Arp Detection is configured. | [optional] 
**ExistCustomDhcpOption** | Pointer to **bool** | Whether custom DHCP Options has been configured. | [optional] 
**ExistDhcpNextServer** | Pointer to **bool** | Whether DHCP Next Server has been configured. | [optional] 
**ExistMultiVlan** | Pointer to **bool** | Whether VLAN Type is Multiple. | [optional] 
**ExistRA** | Pointer to **bool** | Whether RA has been configured. | [optional] 
**GatewaySubnet** | Pointer to **string** | When purpose is interface, gateway subnet is needed. Format: IP/Mask | [optional] 
**Id** | Pointer to **string** | LAN network ID | [optional] 
**IgmpSnoopEnable** | **bool** | Enable IGMP snooping | 
**InterfaceIds** | Pointer to **[]string** | Gateway LAN port IDs (acquired from \&quot;Check WAN/LAN status\&quot;) | [optional] 
**LanNeworkIpv6Config** | Pointer to [**LanNetworkIPV6Config**](LanNetworkIPV6Config.md) |  | [optional] 
**MldSnoopEnable** | Pointer to **bool** | Enable MLD snooping | [optional] 
**Name** | **string** | LAN network name should contain 1 to 128 characters. | 
**Portal** | Pointer to **bool** | Show portal is enabled or not | [optional] 
**PortalId** | Pointer to **string** | Show portal ID | [optional] 
**PortalName** | Pointer to **string** | Show related portal name | [optional] 
**Primary** | Pointer to **bool** | Primary | [optional] 
**Purpose** | **int32** | LAN network purpose should be a value as follows: 0: VLAN; 1: interface | 
**RateLimit** | Pointer to **bool** | Show RateLimit is enabled or not | [optional] 
**Vlan** | Pointer to **int32** | Only Valid when vlanType is 0. Vlan should be within the range of 1-4094. | [optional] 
**VlanType** | Pointer to **int32** | When purpose is interface, VLANType should be a value as follows: 0: Single; 1: Multiple | [optional] 
**Vlans** | Pointer to **string** | When purpose is interface and VLANType is 1, batch create VLANs. VLAN format: 200, 1-100. | [optional] 

## Methods

### NewLanNetworkQueryOpenApiV2VO

`func NewLanNetworkQueryOpenApiV2VO(igmpSnoopEnable bool, name string, purpose int32, ) *LanNetworkQueryOpenApiV2VO`

NewLanNetworkQueryOpenApiV2VO instantiates a new LanNetworkQueryOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkQueryOpenApiV2VOWithDefaults

`func NewLanNetworkQueryOpenApiV2VOWithDefaults() *LanNetworkQueryOpenApiV2VO`

NewLanNetworkQueryOpenApiV2VOWithDefaults instantiates a new LanNetworkQueryOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlRule

`func (o *LanNetworkQueryOpenApiV2VO) GetAccessControlRule() bool`

GetAccessControlRule returns the AccessControlRule field if non-nil, zero value otherwise.

### GetAccessControlRuleOk

`func (o *LanNetworkQueryOpenApiV2VO) GetAccessControlRuleOk() (*bool, bool)`

GetAccessControlRuleOk returns a tuple with the AccessControlRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlRule

`func (o *LanNetworkQueryOpenApiV2VO) SetAccessControlRule(v bool)`

SetAccessControlRule sets AccessControlRule field to given value.

### HasAccessControlRule

`func (o *LanNetworkQueryOpenApiV2VO) HasAccessControlRule() bool`

HasAccessControlRule returns a boolean if a field has been set.

### GetAllLan

`func (o *LanNetworkQueryOpenApiV2VO) GetAllLan() bool`

GetAllLan returns the AllLan field if non-nil, zero value otherwise.

### GetAllLanOk

`func (o *LanNetworkQueryOpenApiV2VO) GetAllLanOk() (*bool, bool)`

GetAllLanOk returns a tuple with the AllLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLan

`func (o *LanNetworkQueryOpenApiV2VO) SetAllLan(v bool)`

SetAllLan sets AllLan field to given value.

### HasAllLan

`func (o *LanNetworkQueryOpenApiV2VO) HasAllLan() bool`

HasAllLan returns a boolean if a field has been set.

### GetApplication

`func (o *LanNetworkQueryOpenApiV2VO) GetApplication() int32`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *LanNetworkQueryOpenApiV2VO) GetApplicationOk() (*int32, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *LanNetworkQueryOpenApiV2VO) SetApplication(v int32)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *LanNetworkQueryOpenApiV2VO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDhcpGuard

`func (o *LanNetworkQueryOpenApiV2VO) GetDhcpGuard() DhcpServersSetting`

GetDhcpGuard returns the DhcpGuard field if non-nil, zero value otherwise.

### GetDhcpGuardOk

`func (o *LanNetworkQueryOpenApiV2VO) GetDhcpGuardOk() (*DhcpServersSetting, bool)`

GetDhcpGuardOk returns a tuple with the DhcpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpGuard

`func (o *LanNetworkQueryOpenApiV2VO) SetDhcpGuard(v DhcpServersSetting)`

SetDhcpGuard sets DhcpGuard field to given value.

### HasDhcpGuard

`func (o *LanNetworkQueryOpenApiV2VO) HasDhcpGuard() bool`

HasDhcpGuard returns a boolean if a field has been set.

### GetDhcpL2RelayEnable

`func (o *LanNetworkQueryOpenApiV2VO) GetDhcpL2RelayEnable() bool`

GetDhcpL2RelayEnable returns the DhcpL2RelayEnable field if non-nil, zero value otherwise.

### GetDhcpL2RelayEnableOk

`func (o *LanNetworkQueryOpenApiV2VO) GetDhcpL2RelayEnableOk() (*bool, bool)`

GetDhcpL2RelayEnableOk returns a tuple with the DhcpL2RelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelayEnable

`func (o *LanNetworkQueryOpenApiV2VO) SetDhcpL2RelayEnable(v bool)`

SetDhcpL2RelayEnable sets DhcpL2RelayEnable field to given value.

### HasDhcpL2RelayEnable

`func (o *LanNetworkQueryOpenApiV2VO) HasDhcpL2RelayEnable() bool`

HasDhcpL2RelayEnable returns a boolean if a field has been set.

### GetDhcpSettingsVO

`func (o *LanNetworkQueryOpenApiV2VO) GetDhcpSettingsVO() DhcpSettingInfo`

GetDhcpSettingsVO returns the DhcpSettingsVO field if non-nil, zero value otherwise.

### GetDhcpSettingsVOOk

`func (o *LanNetworkQueryOpenApiV2VO) GetDhcpSettingsVOOk() (*DhcpSettingInfo, bool)`

GetDhcpSettingsVOOk returns a tuple with the DhcpSettingsVO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSettingsVO

`func (o *LanNetworkQueryOpenApiV2VO) SetDhcpSettingsVO(v DhcpSettingInfo)`

SetDhcpSettingsVO sets DhcpSettingsVO field to given value.

### HasDhcpSettingsVO

`func (o *LanNetworkQueryOpenApiV2VO) HasDhcpSettingsVO() bool`

HasDhcpSettingsVO returns a boolean if a field has been set.

### GetDhcpv6Guard

`func (o *LanNetworkQueryOpenApiV2VO) GetDhcpv6Guard() Dhcpv6ServersSetting`

GetDhcpv6Guard returns the Dhcpv6Guard field if non-nil, zero value otherwise.

### GetDhcpv6GuardOk

`func (o *LanNetworkQueryOpenApiV2VO) GetDhcpv6GuardOk() (*Dhcpv6ServersSetting, bool)`

GetDhcpv6GuardOk returns a tuple with the Dhcpv6Guard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Guard

`func (o *LanNetworkQueryOpenApiV2VO) SetDhcpv6Guard(v Dhcpv6ServersSetting)`

SetDhcpv6Guard sets Dhcpv6Guard field to given value.

### HasDhcpv6Guard

`func (o *LanNetworkQueryOpenApiV2VO) HasDhcpv6Guard() bool`

HasDhcpv6Guard returns a boolean if a field has been set.

### GetDomain

`func (o *LanNetworkQueryOpenApiV2VO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanNetworkQueryOpenApiV2VO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanNetworkQueryOpenApiV2VO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LanNetworkQueryOpenApiV2VO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExistArpDetection

`func (o *LanNetworkQueryOpenApiV2VO) GetExistArpDetection() bool`

GetExistArpDetection returns the ExistArpDetection field if non-nil, zero value otherwise.

### GetExistArpDetectionOk

`func (o *LanNetworkQueryOpenApiV2VO) GetExistArpDetectionOk() (*bool, bool)`

GetExistArpDetectionOk returns a tuple with the ExistArpDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistArpDetection

`func (o *LanNetworkQueryOpenApiV2VO) SetExistArpDetection(v bool)`

SetExistArpDetection sets ExistArpDetection field to given value.

### HasExistArpDetection

`func (o *LanNetworkQueryOpenApiV2VO) HasExistArpDetection() bool`

HasExistArpDetection returns a boolean if a field has been set.

### GetExistCustomDhcpOption

`func (o *LanNetworkQueryOpenApiV2VO) GetExistCustomDhcpOption() bool`

GetExistCustomDhcpOption returns the ExistCustomDhcpOption field if non-nil, zero value otherwise.

### GetExistCustomDhcpOptionOk

`func (o *LanNetworkQueryOpenApiV2VO) GetExistCustomDhcpOptionOk() (*bool, bool)`

GetExistCustomDhcpOptionOk returns a tuple with the ExistCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCustomDhcpOption

`func (o *LanNetworkQueryOpenApiV2VO) SetExistCustomDhcpOption(v bool)`

SetExistCustomDhcpOption sets ExistCustomDhcpOption field to given value.

### HasExistCustomDhcpOption

`func (o *LanNetworkQueryOpenApiV2VO) HasExistCustomDhcpOption() bool`

HasExistCustomDhcpOption returns a boolean if a field has been set.

### GetExistDhcpNextServer

`func (o *LanNetworkQueryOpenApiV2VO) GetExistDhcpNextServer() bool`

GetExistDhcpNextServer returns the ExistDhcpNextServer field if non-nil, zero value otherwise.

### GetExistDhcpNextServerOk

`func (o *LanNetworkQueryOpenApiV2VO) GetExistDhcpNextServerOk() (*bool, bool)`

GetExistDhcpNextServerOk returns a tuple with the ExistDhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistDhcpNextServer

`func (o *LanNetworkQueryOpenApiV2VO) SetExistDhcpNextServer(v bool)`

SetExistDhcpNextServer sets ExistDhcpNextServer field to given value.

### HasExistDhcpNextServer

`func (o *LanNetworkQueryOpenApiV2VO) HasExistDhcpNextServer() bool`

HasExistDhcpNextServer returns a boolean if a field has been set.

### GetExistMultiVlan

`func (o *LanNetworkQueryOpenApiV2VO) GetExistMultiVlan() bool`

GetExistMultiVlan returns the ExistMultiVlan field if non-nil, zero value otherwise.

### GetExistMultiVlanOk

`func (o *LanNetworkQueryOpenApiV2VO) GetExistMultiVlanOk() (*bool, bool)`

GetExistMultiVlanOk returns a tuple with the ExistMultiVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistMultiVlan

`func (o *LanNetworkQueryOpenApiV2VO) SetExistMultiVlan(v bool)`

SetExistMultiVlan sets ExistMultiVlan field to given value.

### HasExistMultiVlan

`func (o *LanNetworkQueryOpenApiV2VO) HasExistMultiVlan() bool`

HasExistMultiVlan returns a boolean if a field has been set.

### GetExistRA

`func (o *LanNetworkQueryOpenApiV2VO) GetExistRA() bool`

GetExistRA returns the ExistRA field if non-nil, zero value otherwise.

### GetExistRAOk

`func (o *LanNetworkQueryOpenApiV2VO) GetExistRAOk() (*bool, bool)`

GetExistRAOk returns a tuple with the ExistRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistRA

`func (o *LanNetworkQueryOpenApiV2VO) SetExistRA(v bool)`

SetExistRA sets ExistRA field to given value.

### HasExistRA

`func (o *LanNetworkQueryOpenApiV2VO) HasExistRA() bool`

HasExistRA returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *LanNetworkQueryOpenApiV2VO) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *LanNetworkQueryOpenApiV2VO) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *LanNetworkQueryOpenApiV2VO) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *LanNetworkQueryOpenApiV2VO) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetId

`func (o *LanNetworkQueryOpenApiV2VO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanNetworkQueryOpenApiV2VO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanNetworkQueryOpenApiV2VO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanNetworkQueryOpenApiV2VO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgmpSnoopEnable

`func (o *LanNetworkQueryOpenApiV2VO) GetIgmpSnoopEnable() bool`

GetIgmpSnoopEnable returns the IgmpSnoopEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopEnableOk

`func (o *LanNetworkQueryOpenApiV2VO) GetIgmpSnoopEnableOk() (*bool, bool)`

GetIgmpSnoopEnableOk returns a tuple with the IgmpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopEnable

`func (o *LanNetworkQueryOpenApiV2VO) SetIgmpSnoopEnable(v bool)`

SetIgmpSnoopEnable sets IgmpSnoopEnable field to given value.


### GetInterfaceIds

`func (o *LanNetworkQueryOpenApiV2VO) GetInterfaceIds() []string`

GetInterfaceIds returns the InterfaceIds field if non-nil, zero value otherwise.

### GetInterfaceIdsOk

`func (o *LanNetworkQueryOpenApiV2VO) GetInterfaceIdsOk() (*[]string, bool)`

GetInterfaceIdsOk returns a tuple with the InterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIds

`func (o *LanNetworkQueryOpenApiV2VO) SetInterfaceIds(v []string)`

SetInterfaceIds sets InterfaceIds field to given value.

### HasInterfaceIds

`func (o *LanNetworkQueryOpenApiV2VO) HasInterfaceIds() bool`

HasInterfaceIds returns a boolean if a field has been set.

### GetLanNeworkIpv6Config

`func (o *LanNetworkQueryOpenApiV2VO) GetLanNeworkIpv6Config() LanNetworkIPV6Config`

GetLanNeworkIpv6Config returns the LanNeworkIpv6Config field if non-nil, zero value otherwise.

### GetLanNeworkIpv6ConfigOk

`func (o *LanNetworkQueryOpenApiV2VO) GetLanNeworkIpv6ConfigOk() (*LanNetworkIPV6Config, bool)`

GetLanNeworkIpv6ConfigOk returns a tuple with the LanNeworkIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNeworkIpv6Config

`func (o *LanNetworkQueryOpenApiV2VO) SetLanNeworkIpv6Config(v LanNetworkIPV6Config)`

SetLanNeworkIpv6Config sets LanNeworkIpv6Config field to given value.

### HasLanNeworkIpv6Config

`func (o *LanNetworkQueryOpenApiV2VO) HasLanNeworkIpv6Config() bool`

HasLanNeworkIpv6Config returns a boolean if a field has been set.

### GetMldSnoopEnable

`func (o *LanNetworkQueryOpenApiV2VO) GetMldSnoopEnable() bool`

GetMldSnoopEnable returns the MldSnoopEnable field if non-nil, zero value otherwise.

### GetMldSnoopEnableOk

`func (o *LanNetworkQueryOpenApiV2VO) GetMldSnoopEnableOk() (*bool, bool)`

GetMldSnoopEnableOk returns a tuple with the MldSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldSnoopEnable

`func (o *LanNetworkQueryOpenApiV2VO) SetMldSnoopEnable(v bool)`

SetMldSnoopEnable sets MldSnoopEnable field to given value.

### HasMldSnoopEnable

`func (o *LanNetworkQueryOpenApiV2VO) HasMldSnoopEnable() bool`

HasMldSnoopEnable returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkQueryOpenApiV2VO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkQueryOpenApiV2VO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkQueryOpenApiV2VO) SetName(v string)`

SetName sets Name field to given value.


### GetPortal

`func (o *LanNetworkQueryOpenApiV2VO) GetPortal() bool`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *LanNetworkQueryOpenApiV2VO) GetPortalOk() (*bool, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *LanNetworkQueryOpenApiV2VO) SetPortal(v bool)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *LanNetworkQueryOpenApiV2VO) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetPortalId

`func (o *LanNetworkQueryOpenApiV2VO) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *LanNetworkQueryOpenApiV2VO) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *LanNetworkQueryOpenApiV2VO) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *LanNetworkQueryOpenApiV2VO) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetPortalName

`func (o *LanNetworkQueryOpenApiV2VO) GetPortalName() string`

GetPortalName returns the PortalName field if non-nil, zero value otherwise.

### GetPortalNameOk

`func (o *LanNetworkQueryOpenApiV2VO) GetPortalNameOk() (*string, bool)`

GetPortalNameOk returns a tuple with the PortalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalName

`func (o *LanNetworkQueryOpenApiV2VO) SetPortalName(v string)`

SetPortalName sets PortalName field to given value.

### HasPortalName

`func (o *LanNetworkQueryOpenApiV2VO) HasPortalName() bool`

HasPortalName returns a boolean if a field has been set.

### GetPrimary

`func (o *LanNetworkQueryOpenApiV2VO) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *LanNetworkQueryOpenApiV2VO) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *LanNetworkQueryOpenApiV2VO) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *LanNetworkQueryOpenApiV2VO) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPurpose

`func (o *LanNetworkQueryOpenApiV2VO) GetPurpose() int32`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *LanNetworkQueryOpenApiV2VO) GetPurposeOk() (*int32, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *LanNetworkQueryOpenApiV2VO) SetPurpose(v int32)`

SetPurpose sets Purpose field to given value.


### GetRateLimit

`func (o *LanNetworkQueryOpenApiV2VO) GetRateLimit() bool`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *LanNetworkQueryOpenApiV2VO) GetRateLimitOk() (*bool, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *LanNetworkQueryOpenApiV2VO) SetRateLimit(v bool)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *LanNetworkQueryOpenApiV2VO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkQueryOpenApiV2VO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkQueryOpenApiV2VO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkQueryOpenApiV2VO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkQueryOpenApiV2VO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkQueryOpenApiV2VO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkQueryOpenApiV2VO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkQueryOpenApiV2VO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkQueryOpenApiV2VO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetVlans

`func (o *LanNetworkQueryOpenApiV2VO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *LanNetworkQueryOpenApiV2VO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *LanNetworkQueryOpenApiV2VO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *LanNetworkQueryOpenApiV2VO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


