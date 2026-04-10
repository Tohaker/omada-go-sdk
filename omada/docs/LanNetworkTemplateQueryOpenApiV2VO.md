# LanNetworkTemplateQueryOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlRule** | Pointer to **bool** | Show AccessControlRule is enabled or not | [optional] 
**AllLan** | Pointer to **bool** | When Internet pre-config is closed or Internet pre-config is Universal, allLAN is \&quot;true\&quot;; after adopting gateway, allLAN is \&quot;false\&quot;. | [optional] 
**Application** | Pointer to **int32** | Effective device type should be a value as follows: 0: Gateway and Switch; 1: Switch | [optional] 
**DhcpGuard** | Pointer to [**DhcpServersSetting**](DhcpServersSetting.md) |  | [optional] 
**DhcpL2RelayEnable** | Pointer to **bool** | The switch of DHCP L2 relay | [optional] 
**DhcpSettingsVO** | Pointer to [**DhcpSettingsTemplateOpenApiVO**](DhcpSettingsTemplateOpenApiVO.md) |  | [optional] 
**Dhcpv6Guard** | Pointer to [**Dhcpv6ServersSetting**](Dhcpv6ServersSetting.md) |  | [optional] 
**Domain** | Pointer to **string** | The domain of this network | [optional] 
**GatewaySubnet** | Pointer to **string** | When purpose is interface, gateway subnet is needed. Format: IP/Mask | [optional] 
**Id** | Pointer to **string** | LAN network ID | [optional] 
**IgmpSnoopEnable** | **bool** | Enable IGMP snooping | 
**InterfaceIds** | Pointer to **[]string** | Gateway LAN port IDs (acquired from \&quot;Check WAN/LAN status\&quot;) | [optional] 
**LanNeworkIpv6Config** | Pointer to [**LanNetworkIpv6ConfigTemplateOpenApiVO**](LanNetworkIpv6ConfigTemplateOpenApiVO.md) |  | [optional] 
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

### NewLanNetworkTemplateQueryOpenApiV2VO

`func NewLanNetworkTemplateQueryOpenApiV2VO(igmpSnoopEnable bool, name string, purpose int32, ) *LanNetworkTemplateQueryOpenApiV2VO`

NewLanNetworkTemplateQueryOpenApiV2VO instantiates a new LanNetworkTemplateQueryOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkTemplateQueryOpenApiV2VOWithDefaults

`func NewLanNetworkTemplateQueryOpenApiV2VOWithDefaults() *LanNetworkTemplateQueryOpenApiV2VO`

NewLanNetworkTemplateQueryOpenApiV2VOWithDefaults instantiates a new LanNetworkTemplateQueryOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlRule

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetAccessControlRule() bool`

GetAccessControlRule returns the AccessControlRule field if non-nil, zero value otherwise.

### GetAccessControlRuleOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetAccessControlRuleOk() (*bool, bool)`

GetAccessControlRuleOk returns a tuple with the AccessControlRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlRule

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetAccessControlRule(v bool)`

SetAccessControlRule sets AccessControlRule field to given value.

### HasAccessControlRule

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasAccessControlRule() bool`

HasAccessControlRule returns a boolean if a field has been set.

### GetAllLan

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetAllLan() bool`

GetAllLan returns the AllLan field if non-nil, zero value otherwise.

### GetAllLanOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetAllLanOk() (*bool, bool)`

GetAllLanOk returns a tuple with the AllLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLan

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetAllLan(v bool)`

SetAllLan sets AllLan field to given value.

### HasAllLan

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasAllLan() bool`

HasAllLan returns a boolean if a field has been set.

### GetApplication

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetApplication() int32`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetApplicationOk() (*int32, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetApplication(v int32)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDhcpGuard

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDhcpGuard() DhcpServersSetting`

GetDhcpGuard returns the DhcpGuard field if non-nil, zero value otherwise.

### GetDhcpGuardOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDhcpGuardOk() (*DhcpServersSetting, bool)`

GetDhcpGuardOk returns a tuple with the DhcpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpGuard

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetDhcpGuard(v DhcpServersSetting)`

SetDhcpGuard sets DhcpGuard field to given value.

### HasDhcpGuard

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasDhcpGuard() bool`

HasDhcpGuard returns a boolean if a field has been set.

### GetDhcpL2RelayEnable

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDhcpL2RelayEnable() bool`

GetDhcpL2RelayEnable returns the DhcpL2RelayEnable field if non-nil, zero value otherwise.

### GetDhcpL2RelayEnableOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDhcpL2RelayEnableOk() (*bool, bool)`

GetDhcpL2RelayEnableOk returns a tuple with the DhcpL2RelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelayEnable

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetDhcpL2RelayEnable(v bool)`

SetDhcpL2RelayEnable sets DhcpL2RelayEnable field to given value.

### HasDhcpL2RelayEnable

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasDhcpL2RelayEnable() bool`

HasDhcpL2RelayEnable returns a boolean if a field has been set.

### GetDhcpSettingsVO

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDhcpSettingsVO() DhcpSettingsTemplateOpenApiVO`

GetDhcpSettingsVO returns the DhcpSettingsVO field if non-nil, zero value otherwise.

### GetDhcpSettingsVOOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDhcpSettingsVOOk() (*DhcpSettingsTemplateOpenApiVO, bool)`

GetDhcpSettingsVOOk returns a tuple with the DhcpSettingsVO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSettingsVO

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetDhcpSettingsVO(v DhcpSettingsTemplateOpenApiVO)`

SetDhcpSettingsVO sets DhcpSettingsVO field to given value.

### HasDhcpSettingsVO

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasDhcpSettingsVO() bool`

HasDhcpSettingsVO returns a boolean if a field has been set.

### GetDhcpv6Guard

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDhcpv6Guard() Dhcpv6ServersSetting`

GetDhcpv6Guard returns the Dhcpv6Guard field if non-nil, zero value otherwise.

### GetDhcpv6GuardOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDhcpv6GuardOk() (*Dhcpv6ServersSetting, bool)`

GetDhcpv6GuardOk returns a tuple with the Dhcpv6Guard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Guard

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetDhcpv6Guard(v Dhcpv6ServersSetting)`

SetDhcpv6Guard sets Dhcpv6Guard field to given value.

### HasDhcpv6Guard

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasDhcpv6Guard() bool`

HasDhcpv6Guard returns a boolean if a field has been set.

### GetDomain

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetId

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgmpSnoopEnable

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetIgmpSnoopEnable() bool`

GetIgmpSnoopEnable returns the IgmpSnoopEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopEnableOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetIgmpSnoopEnableOk() (*bool, bool)`

GetIgmpSnoopEnableOk returns a tuple with the IgmpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopEnable

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetIgmpSnoopEnable(v bool)`

SetIgmpSnoopEnable sets IgmpSnoopEnable field to given value.


### GetInterfaceIds

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetInterfaceIds() []string`

GetInterfaceIds returns the InterfaceIds field if non-nil, zero value otherwise.

### GetInterfaceIdsOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetInterfaceIdsOk() (*[]string, bool)`

GetInterfaceIdsOk returns a tuple with the InterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIds

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetInterfaceIds(v []string)`

SetInterfaceIds sets InterfaceIds field to given value.

### HasInterfaceIds

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasInterfaceIds() bool`

HasInterfaceIds returns a boolean if a field has been set.

### GetLanNeworkIpv6Config

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetLanNeworkIpv6Config() LanNetworkIpv6ConfigTemplateOpenApiVO`

GetLanNeworkIpv6Config returns the LanNeworkIpv6Config field if non-nil, zero value otherwise.

### GetLanNeworkIpv6ConfigOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetLanNeworkIpv6ConfigOk() (*LanNetworkIpv6ConfigTemplateOpenApiVO, bool)`

GetLanNeworkIpv6ConfigOk returns a tuple with the LanNeworkIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNeworkIpv6Config

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetLanNeworkIpv6Config(v LanNetworkIpv6ConfigTemplateOpenApiVO)`

SetLanNeworkIpv6Config sets LanNeworkIpv6Config field to given value.

### HasLanNeworkIpv6Config

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasLanNeworkIpv6Config() bool`

HasLanNeworkIpv6Config returns a boolean if a field has been set.

### GetMldSnoopEnable

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetMldSnoopEnable() bool`

GetMldSnoopEnable returns the MldSnoopEnable field if non-nil, zero value otherwise.

### GetMldSnoopEnableOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetMldSnoopEnableOk() (*bool, bool)`

GetMldSnoopEnableOk returns a tuple with the MldSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldSnoopEnable

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetMldSnoopEnable(v bool)`

SetMldSnoopEnable sets MldSnoopEnable field to given value.

### HasMldSnoopEnable

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasMldSnoopEnable() bool`

HasMldSnoopEnable returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetName(v string)`

SetName sets Name field to given value.


### GetPortal

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPortal() bool`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPortalOk() (*bool, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetPortal(v bool)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetPortalId

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetPortalName

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPortalName() string`

GetPortalName returns the PortalName field if non-nil, zero value otherwise.

### GetPortalNameOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPortalNameOk() (*string, bool)`

GetPortalNameOk returns a tuple with the PortalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalName

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetPortalName(v string)`

SetPortalName sets PortalName field to given value.

### HasPortalName

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasPortalName() bool`

HasPortalName returns a boolean if a field has been set.

### GetPrimary

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPurpose

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPurpose() int32`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetPurposeOk() (*int32, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetPurpose(v int32)`

SetPurpose sets Purpose field to given value.


### GetRateLimit

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetRateLimit() bool`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetRateLimitOk() (*bool, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetRateLimit(v bool)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetVlans

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *LanNetworkTemplateQueryOpenApiV2VO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *LanNetworkTemplateQueryOpenApiV2VO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *LanNetworkTemplateQueryOpenApiV2VO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


