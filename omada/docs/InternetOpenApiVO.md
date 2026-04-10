# InternetOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Is Internet enable. | [optional] 
**EnableModified** | Pointer to **bool** | Whether it enable modified. | [optional] 
**GatewayMac** | Pointer to **string** | Gateway MAC. | [optional] 
**Interval** | Pointer to **int32** | Time taken to check the connection of wan port. | [optional] 
**LteIntervalTips** | Pointer to **bool** | Whether to display the interval prompt for LTE ports. | [optional] 
**LteWanSettings** | Pointer to [**[]LteWanPortSettingOpenApiVO**](LteWanPortSettingOpenApiVO.md) | A list of LTE wan setting. | [optional] 
**NetworkNames** | Pointer to **[]string** | A list of the name of the LanNetwork entry that was invalidated on the frontend. | [optional] 
**OmadacId** | Pointer to **string** | Omadac ID. | [optional] 
**OsgPortInfo** | Pointer to [**OsgPortInfoOpenApiVO**](OsgPortInfoOpenApiVO.md) |  | [optional] 
**PortOnlineStatus** | Pointer to [**PortOnlineStatusOpenApiVO**](PortOnlineStatusOpenApiVO.md) |  | [optional] 
**PortUuids** | Pointer to **[]string** | The list of wan port UUID. | [optional] 
**Resource** | Pointer to **int32** | Resource of data. | [optional] 
**SiteId** | **string** | Site ID. | 
**SupportAllWan** | Pointer to **bool** | Whether the device supports all wan. | [optional] 
**SupportCustomInterval** | Pointer to **bool** | Whether the device supports Custom Interval. | [optional] 
**SupportDhcpOptions** | Pointer to **bool** | Whether WAN port dhcpOptions can be customized. | [optional] 
**SupportDiscreteWan** | Pointer to **bool** | Whether the device supports Discrete Wan. | [optional] 
**SupportDsLite** | Pointer to **bool** | Whether the WAN port supports configuring DS-Lite. | [optional] 
**SupportDsl** | Pointer to **bool** | Whether it supports DSL. | [optional] 
**SupportDualSim** | Pointer to **int32** | Whether it supports dual SIM single standby: 0: not support, 1: support. | [optional] 
**SupportIpv6NonAddress** | Pointer to **bool** | Whether the IPv6 configuration supports Non-Address. | [optional] 
**SupportLte** | Pointer to **bool** | Whether to configure LTE Wan. | [optional] 
**SupportMapE** | Pointer to **bool** | Whether the WAN port supports configuring MAP-E. | [optional] 
**SupportMaxWanNum** | Pointer to **int32** | The maximum number of wans supported. | [optional] 
**SupportMssClamping** | Pointer to **bool** | Whether MSS clamping is supported for PPPoE, L2TP, and PPTP. | [optional] 
**SupportNetworkIsolation** | Pointer to **bool** | Whether it supports isolate network. | [optional] 
**SupportPppoeMru** | Pointer to **bool** | Whether to configure PPPoE MRUs. | [optional] 
**SupportReduceUsbRfi** | Pointer to **bool** | Whether the device supports Reduce USB 3.0 Interference Reduction. | [optional] 
**SupportTimingMode** | Pointer to **bool** | Whether the Load Balance Failover supports Timing mode. | [optional] 
**SupportUsbDhcpOptions** | Pointer to **bool** | Whether USB port dhcpOptions can be customized. | [optional] 
**SupportVirtualWan** | Pointer to **bool** | Whether it supports virtual wan. | [optional] 
**SupportWanMultipleIp** | Pointer to **bool** | Whether to configure multiple IP addresses for wan ports. | [optional] 
**Unit** | Pointer to **int32** | Unit. | [optional] 
**UnmatchedWans** | Pointer to **[]string** | A list of the name of the WAN port that was invalidated on the frontend. | [optional] 
**UsbLteSettings** | Pointer to [**[]UsbLteSettingOpenApiVO**](UsbLteSettingOpenApiVO.md) | A list of USB LTE setting. | [optional] 
**WanLoadBalance** | Pointer to [**WanLoadBalanceOpenApiVO**](WanLoadBalanceOpenApiVO.md) |  | [optional] 
**WanPortSettings** | Pointer to [**[]WanPortSettingOpenApiVO**](WanPortSettingOpenApiVO.md) | A list of wan port setting. | [optional] 

## Methods

### NewInternetOpenApiVO

`func NewInternetOpenApiVO(siteId string, ) *InternetOpenApiVO`

NewInternetOpenApiVO instantiates a new InternetOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetOpenApiVOWithDefaults

`func NewInternetOpenApiVOWithDefaults() *InternetOpenApiVO`

NewInternetOpenApiVOWithDefaults instantiates a new InternetOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *InternetOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *InternetOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *InternetOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *InternetOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableModified

`func (o *InternetOpenApiVO) GetEnableModified() bool`

GetEnableModified returns the EnableModified field if non-nil, zero value otherwise.

### GetEnableModifiedOk

`func (o *InternetOpenApiVO) GetEnableModifiedOk() (*bool, bool)`

GetEnableModifiedOk returns a tuple with the EnableModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableModified

`func (o *InternetOpenApiVO) SetEnableModified(v bool)`

SetEnableModified sets EnableModified field to given value.

### HasEnableModified

`func (o *InternetOpenApiVO) HasEnableModified() bool`

HasEnableModified returns a boolean if a field has been set.

### GetGatewayMac

`func (o *InternetOpenApiVO) GetGatewayMac() string`

GetGatewayMac returns the GatewayMac field if non-nil, zero value otherwise.

### GetGatewayMacOk

`func (o *InternetOpenApiVO) GetGatewayMacOk() (*string, bool)`

GetGatewayMacOk returns a tuple with the GatewayMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMac

`func (o *InternetOpenApiVO) SetGatewayMac(v string)`

SetGatewayMac sets GatewayMac field to given value.

### HasGatewayMac

`func (o *InternetOpenApiVO) HasGatewayMac() bool`

HasGatewayMac returns a boolean if a field has been set.

### GetInterval

`func (o *InternetOpenApiVO) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InternetOpenApiVO) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InternetOpenApiVO) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InternetOpenApiVO) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLteIntervalTips

`func (o *InternetOpenApiVO) GetLteIntervalTips() bool`

GetLteIntervalTips returns the LteIntervalTips field if non-nil, zero value otherwise.

### GetLteIntervalTipsOk

`func (o *InternetOpenApiVO) GetLteIntervalTipsOk() (*bool, bool)`

GetLteIntervalTipsOk returns a tuple with the LteIntervalTips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteIntervalTips

`func (o *InternetOpenApiVO) SetLteIntervalTips(v bool)`

SetLteIntervalTips sets LteIntervalTips field to given value.

### HasLteIntervalTips

`func (o *InternetOpenApiVO) HasLteIntervalTips() bool`

HasLteIntervalTips returns a boolean if a field has been set.

### GetLteWanSettings

`func (o *InternetOpenApiVO) GetLteWanSettings() []LteWanPortSettingOpenApiVO`

GetLteWanSettings returns the LteWanSettings field if non-nil, zero value otherwise.

### GetLteWanSettingsOk

`func (o *InternetOpenApiVO) GetLteWanSettingsOk() (*[]LteWanPortSettingOpenApiVO, bool)`

GetLteWanSettingsOk returns a tuple with the LteWanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteWanSettings

`func (o *InternetOpenApiVO) SetLteWanSettings(v []LteWanPortSettingOpenApiVO)`

SetLteWanSettings sets LteWanSettings field to given value.

### HasLteWanSettings

`func (o *InternetOpenApiVO) HasLteWanSettings() bool`

HasLteWanSettings returns a boolean if a field has been set.

### GetNetworkNames

`func (o *InternetOpenApiVO) GetNetworkNames() []string`

GetNetworkNames returns the NetworkNames field if non-nil, zero value otherwise.

### GetNetworkNamesOk

`func (o *InternetOpenApiVO) GetNetworkNamesOk() (*[]string, bool)`

GetNetworkNamesOk returns a tuple with the NetworkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNames

`func (o *InternetOpenApiVO) SetNetworkNames(v []string)`

SetNetworkNames sets NetworkNames field to given value.

### HasNetworkNames

`func (o *InternetOpenApiVO) HasNetworkNames() bool`

HasNetworkNames returns a boolean if a field has been set.

### GetOmadacId

`func (o *InternetOpenApiVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *InternetOpenApiVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *InternetOpenApiVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *InternetOpenApiVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOsgPortInfo

`func (o *InternetOpenApiVO) GetOsgPortInfo() OsgPortInfoOpenApiVO`

GetOsgPortInfo returns the OsgPortInfo field if non-nil, zero value otherwise.

### GetOsgPortInfoOk

`func (o *InternetOpenApiVO) GetOsgPortInfoOk() (*OsgPortInfoOpenApiVO, bool)`

GetOsgPortInfoOk returns a tuple with the OsgPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsgPortInfo

`func (o *InternetOpenApiVO) SetOsgPortInfo(v OsgPortInfoOpenApiVO)`

SetOsgPortInfo sets OsgPortInfo field to given value.

### HasOsgPortInfo

`func (o *InternetOpenApiVO) HasOsgPortInfo() bool`

HasOsgPortInfo returns a boolean if a field has been set.

### GetPortOnlineStatus

`func (o *InternetOpenApiVO) GetPortOnlineStatus() PortOnlineStatusOpenApiVO`

GetPortOnlineStatus returns the PortOnlineStatus field if non-nil, zero value otherwise.

### GetPortOnlineStatusOk

`func (o *InternetOpenApiVO) GetPortOnlineStatusOk() (*PortOnlineStatusOpenApiVO, bool)`

GetPortOnlineStatusOk returns a tuple with the PortOnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOnlineStatus

`func (o *InternetOpenApiVO) SetPortOnlineStatus(v PortOnlineStatusOpenApiVO)`

SetPortOnlineStatus sets PortOnlineStatus field to given value.

### HasPortOnlineStatus

`func (o *InternetOpenApiVO) HasPortOnlineStatus() bool`

HasPortOnlineStatus returns a boolean if a field has been set.

### GetPortUuids

`func (o *InternetOpenApiVO) GetPortUuids() []string`

GetPortUuids returns the PortUuids field if non-nil, zero value otherwise.

### GetPortUuidsOk

`func (o *InternetOpenApiVO) GetPortUuidsOk() (*[]string, bool)`

GetPortUuidsOk returns a tuple with the PortUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuids

`func (o *InternetOpenApiVO) SetPortUuids(v []string)`

SetPortUuids sets PortUuids field to given value.

### HasPortUuids

`func (o *InternetOpenApiVO) HasPortUuids() bool`

HasPortUuids returns a boolean if a field has been set.

### GetResource

`func (o *InternetOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *InternetOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *InternetOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *InternetOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSiteId

`func (o *InternetOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InternetOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InternetOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSupportAllWan

`func (o *InternetOpenApiVO) GetSupportAllWan() bool`

GetSupportAllWan returns the SupportAllWan field if non-nil, zero value otherwise.

### GetSupportAllWanOk

`func (o *InternetOpenApiVO) GetSupportAllWanOk() (*bool, bool)`

GetSupportAllWanOk returns a tuple with the SupportAllWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAllWan

`func (o *InternetOpenApiVO) SetSupportAllWan(v bool)`

SetSupportAllWan sets SupportAllWan field to given value.

### HasSupportAllWan

`func (o *InternetOpenApiVO) HasSupportAllWan() bool`

HasSupportAllWan returns a boolean if a field has been set.

### GetSupportCustomInterval

`func (o *InternetOpenApiVO) GetSupportCustomInterval() bool`

GetSupportCustomInterval returns the SupportCustomInterval field if non-nil, zero value otherwise.

### GetSupportCustomIntervalOk

`func (o *InternetOpenApiVO) GetSupportCustomIntervalOk() (*bool, bool)`

GetSupportCustomIntervalOk returns a tuple with the SupportCustomInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomInterval

`func (o *InternetOpenApiVO) SetSupportCustomInterval(v bool)`

SetSupportCustomInterval sets SupportCustomInterval field to given value.

### HasSupportCustomInterval

`func (o *InternetOpenApiVO) HasSupportCustomInterval() bool`

HasSupportCustomInterval returns a boolean if a field has been set.

### GetSupportDhcpOptions

`func (o *InternetOpenApiVO) GetSupportDhcpOptions() bool`

GetSupportDhcpOptions returns the SupportDhcpOptions field if non-nil, zero value otherwise.

### GetSupportDhcpOptionsOk

`func (o *InternetOpenApiVO) GetSupportDhcpOptionsOk() (*bool, bool)`

GetSupportDhcpOptionsOk returns a tuple with the SupportDhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpOptions

`func (o *InternetOpenApiVO) SetSupportDhcpOptions(v bool)`

SetSupportDhcpOptions sets SupportDhcpOptions field to given value.

### HasSupportDhcpOptions

`func (o *InternetOpenApiVO) HasSupportDhcpOptions() bool`

HasSupportDhcpOptions returns a boolean if a field has been set.

### GetSupportDiscreteWan

`func (o *InternetOpenApiVO) GetSupportDiscreteWan() bool`

GetSupportDiscreteWan returns the SupportDiscreteWan field if non-nil, zero value otherwise.

### GetSupportDiscreteWanOk

`func (o *InternetOpenApiVO) GetSupportDiscreteWanOk() (*bool, bool)`

GetSupportDiscreteWanOk returns a tuple with the SupportDiscreteWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDiscreteWan

`func (o *InternetOpenApiVO) SetSupportDiscreteWan(v bool)`

SetSupportDiscreteWan sets SupportDiscreteWan field to given value.

### HasSupportDiscreteWan

`func (o *InternetOpenApiVO) HasSupportDiscreteWan() bool`

HasSupportDiscreteWan returns a boolean if a field has been set.

### GetSupportDsLite

`func (o *InternetOpenApiVO) GetSupportDsLite() bool`

GetSupportDsLite returns the SupportDsLite field if non-nil, zero value otherwise.

### GetSupportDsLiteOk

`func (o *InternetOpenApiVO) GetSupportDsLiteOk() (*bool, bool)`

GetSupportDsLiteOk returns a tuple with the SupportDsLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDsLite

`func (o *InternetOpenApiVO) SetSupportDsLite(v bool)`

SetSupportDsLite sets SupportDsLite field to given value.

### HasSupportDsLite

`func (o *InternetOpenApiVO) HasSupportDsLite() bool`

HasSupportDsLite returns a boolean if a field has been set.

### GetSupportDsl

`func (o *InternetOpenApiVO) GetSupportDsl() bool`

GetSupportDsl returns the SupportDsl field if non-nil, zero value otherwise.

### GetSupportDslOk

`func (o *InternetOpenApiVO) GetSupportDslOk() (*bool, bool)`

GetSupportDslOk returns a tuple with the SupportDsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDsl

`func (o *InternetOpenApiVO) SetSupportDsl(v bool)`

SetSupportDsl sets SupportDsl field to given value.

### HasSupportDsl

`func (o *InternetOpenApiVO) HasSupportDsl() bool`

HasSupportDsl returns a boolean if a field has been set.

### GetSupportDualSim

`func (o *InternetOpenApiVO) GetSupportDualSim() int32`

GetSupportDualSim returns the SupportDualSim field if non-nil, zero value otherwise.

### GetSupportDualSimOk

`func (o *InternetOpenApiVO) GetSupportDualSimOk() (*int32, bool)`

GetSupportDualSimOk returns a tuple with the SupportDualSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDualSim

`func (o *InternetOpenApiVO) SetSupportDualSim(v int32)`

SetSupportDualSim sets SupportDualSim field to given value.

### HasSupportDualSim

`func (o *InternetOpenApiVO) HasSupportDualSim() bool`

HasSupportDualSim returns a boolean if a field has been set.

### GetSupportIpv6NonAddress

`func (o *InternetOpenApiVO) GetSupportIpv6NonAddress() bool`

GetSupportIpv6NonAddress returns the SupportIpv6NonAddress field if non-nil, zero value otherwise.

### GetSupportIpv6NonAddressOk

`func (o *InternetOpenApiVO) GetSupportIpv6NonAddressOk() (*bool, bool)`

GetSupportIpv6NonAddressOk returns a tuple with the SupportIpv6NonAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpv6NonAddress

`func (o *InternetOpenApiVO) SetSupportIpv6NonAddress(v bool)`

SetSupportIpv6NonAddress sets SupportIpv6NonAddress field to given value.

### HasSupportIpv6NonAddress

`func (o *InternetOpenApiVO) HasSupportIpv6NonAddress() bool`

HasSupportIpv6NonAddress returns a boolean if a field has been set.

### GetSupportLte

`func (o *InternetOpenApiVO) GetSupportLte() bool`

GetSupportLte returns the SupportLte field if non-nil, zero value otherwise.

### GetSupportLteOk

`func (o *InternetOpenApiVO) GetSupportLteOk() (*bool, bool)`

GetSupportLteOk returns a tuple with the SupportLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLte

`func (o *InternetOpenApiVO) SetSupportLte(v bool)`

SetSupportLte sets SupportLte field to given value.

### HasSupportLte

`func (o *InternetOpenApiVO) HasSupportLte() bool`

HasSupportLte returns a boolean if a field has been set.

### GetSupportMapE

`func (o *InternetOpenApiVO) GetSupportMapE() bool`

GetSupportMapE returns the SupportMapE field if non-nil, zero value otherwise.

### GetSupportMapEOk

`func (o *InternetOpenApiVO) GetSupportMapEOk() (*bool, bool)`

GetSupportMapEOk returns a tuple with the SupportMapE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMapE

`func (o *InternetOpenApiVO) SetSupportMapE(v bool)`

SetSupportMapE sets SupportMapE field to given value.

### HasSupportMapE

`func (o *InternetOpenApiVO) HasSupportMapE() bool`

HasSupportMapE returns a boolean if a field has been set.

### GetSupportMaxWanNum

`func (o *InternetOpenApiVO) GetSupportMaxWanNum() int32`

GetSupportMaxWanNum returns the SupportMaxWanNum field if non-nil, zero value otherwise.

### GetSupportMaxWanNumOk

`func (o *InternetOpenApiVO) GetSupportMaxWanNumOk() (*int32, bool)`

GetSupportMaxWanNumOk returns a tuple with the SupportMaxWanNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMaxWanNum

`func (o *InternetOpenApiVO) SetSupportMaxWanNum(v int32)`

SetSupportMaxWanNum sets SupportMaxWanNum field to given value.

### HasSupportMaxWanNum

`func (o *InternetOpenApiVO) HasSupportMaxWanNum() bool`

HasSupportMaxWanNum returns a boolean if a field has been set.

### GetSupportMssClamping

`func (o *InternetOpenApiVO) GetSupportMssClamping() bool`

GetSupportMssClamping returns the SupportMssClamping field if non-nil, zero value otherwise.

### GetSupportMssClampingOk

`func (o *InternetOpenApiVO) GetSupportMssClampingOk() (*bool, bool)`

GetSupportMssClampingOk returns a tuple with the SupportMssClamping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMssClamping

`func (o *InternetOpenApiVO) SetSupportMssClamping(v bool)`

SetSupportMssClamping sets SupportMssClamping field to given value.

### HasSupportMssClamping

`func (o *InternetOpenApiVO) HasSupportMssClamping() bool`

HasSupportMssClamping returns a boolean if a field has been set.

### GetSupportNetworkIsolation

`func (o *InternetOpenApiVO) GetSupportNetworkIsolation() bool`

GetSupportNetworkIsolation returns the SupportNetworkIsolation field if non-nil, zero value otherwise.

### GetSupportNetworkIsolationOk

`func (o *InternetOpenApiVO) GetSupportNetworkIsolationOk() (*bool, bool)`

GetSupportNetworkIsolationOk returns a tuple with the SupportNetworkIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportNetworkIsolation

`func (o *InternetOpenApiVO) SetSupportNetworkIsolation(v bool)`

SetSupportNetworkIsolation sets SupportNetworkIsolation field to given value.

### HasSupportNetworkIsolation

`func (o *InternetOpenApiVO) HasSupportNetworkIsolation() bool`

HasSupportNetworkIsolation returns a boolean if a field has been set.

### GetSupportPppoeMru

`func (o *InternetOpenApiVO) GetSupportPppoeMru() bool`

GetSupportPppoeMru returns the SupportPppoeMru field if non-nil, zero value otherwise.

### GetSupportPppoeMruOk

`func (o *InternetOpenApiVO) GetSupportPppoeMruOk() (*bool, bool)`

GetSupportPppoeMruOk returns a tuple with the SupportPppoeMru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPppoeMru

`func (o *InternetOpenApiVO) SetSupportPppoeMru(v bool)`

SetSupportPppoeMru sets SupportPppoeMru field to given value.

### HasSupportPppoeMru

`func (o *InternetOpenApiVO) HasSupportPppoeMru() bool`

HasSupportPppoeMru returns a boolean if a field has been set.

### GetSupportReduceUsbRfi

`func (o *InternetOpenApiVO) GetSupportReduceUsbRfi() bool`

GetSupportReduceUsbRfi returns the SupportReduceUsbRfi field if non-nil, zero value otherwise.

### GetSupportReduceUsbRfiOk

`func (o *InternetOpenApiVO) GetSupportReduceUsbRfiOk() (*bool, bool)`

GetSupportReduceUsbRfiOk returns a tuple with the SupportReduceUsbRfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportReduceUsbRfi

`func (o *InternetOpenApiVO) SetSupportReduceUsbRfi(v bool)`

SetSupportReduceUsbRfi sets SupportReduceUsbRfi field to given value.

### HasSupportReduceUsbRfi

`func (o *InternetOpenApiVO) HasSupportReduceUsbRfi() bool`

HasSupportReduceUsbRfi returns a boolean if a field has been set.

### GetSupportTimingMode

`func (o *InternetOpenApiVO) GetSupportTimingMode() bool`

GetSupportTimingMode returns the SupportTimingMode field if non-nil, zero value otherwise.

### GetSupportTimingModeOk

`func (o *InternetOpenApiVO) GetSupportTimingModeOk() (*bool, bool)`

GetSupportTimingModeOk returns a tuple with the SupportTimingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTimingMode

`func (o *InternetOpenApiVO) SetSupportTimingMode(v bool)`

SetSupportTimingMode sets SupportTimingMode field to given value.

### HasSupportTimingMode

`func (o *InternetOpenApiVO) HasSupportTimingMode() bool`

HasSupportTimingMode returns a boolean if a field has been set.

### GetSupportUsbDhcpOptions

`func (o *InternetOpenApiVO) GetSupportUsbDhcpOptions() bool`

GetSupportUsbDhcpOptions returns the SupportUsbDhcpOptions field if non-nil, zero value otherwise.

### GetSupportUsbDhcpOptionsOk

`func (o *InternetOpenApiVO) GetSupportUsbDhcpOptionsOk() (*bool, bool)`

GetSupportUsbDhcpOptionsOk returns a tuple with the SupportUsbDhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUsbDhcpOptions

`func (o *InternetOpenApiVO) SetSupportUsbDhcpOptions(v bool)`

SetSupportUsbDhcpOptions sets SupportUsbDhcpOptions field to given value.

### HasSupportUsbDhcpOptions

`func (o *InternetOpenApiVO) HasSupportUsbDhcpOptions() bool`

HasSupportUsbDhcpOptions returns a boolean if a field has been set.

### GetSupportVirtualWan

`func (o *InternetOpenApiVO) GetSupportVirtualWan() bool`

GetSupportVirtualWan returns the SupportVirtualWan field if non-nil, zero value otherwise.

### GetSupportVirtualWanOk

`func (o *InternetOpenApiVO) GetSupportVirtualWanOk() (*bool, bool)`

GetSupportVirtualWanOk returns a tuple with the SupportVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVirtualWan

`func (o *InternetOpenApiVO) SetSupportVirtualWan(v bool)`

SetSupportVirtualWan sets SupportVirtualWan field to given value.

### HasSupportVirtualWan

`func (o *InternetOpenApiVO) HasSupportVirtualWan() bool`

HasSupportVirtualWan returns a boolean if a field has been set.

### GetSupportWanMultipleIp

`func (o *InternetOpenApiVO) GetSupportWanMultipleIp() bool`

GetSupportWanMultipleIp returns the SupportWanMultipleIp field if non-nil, zero value otherwise.

### GetSupportWanMultipleIpOk

`func (o *InternetOpenApiVO) GetSupportWanMultipleIpOk() (*bool, bool)`

GetSupportWanMultipleIpOk returns a tuple with the SupportWanMultipleIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWanMultipleIp

`func (o *InternetOpenApiVO) SetSupportWanMultipleIp(v bool)`

SetSupportWanMultipleIp sets SupportWanMultipleIp field to given value.

### HasSupportWanMultipleIp

`func (o *InternetOpenApiVO) HasSupportWanMultipleIp() bool`

HasSupportWanMultipleIp returns a boolean if a field has been set.

### GetUnit

`func (o *InternetOpenApiVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InternetOpenApiVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InternetOpenApiVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InternetOpenApiVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnmatchedWans

`func (o *InternetOpenApiVO) GetUnmatchedWans() []string`

GetUnmatchedWans returns the UnmatchedWans field if non-nil, zero value otherwise.

### GetUnmatchedWansOk

`func (o *InternetOpenApiVO) GetUnmatchedWansOk() (*[]string, bool)`

GetUnmatchedWansOk returns a tuple with the UnmatchedWans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedWans

`func (o *InternetOpenApiVO) SetUnmatchedWans(v []string)`

SetUnmatchedWans sets UnmatchedWans field to given value.

### HasUnmatchedWans

`func (o *InternetOpenApiVO) HasUnmatchedWans() bool`

HasUnmatchedWans returns a boolean if a field has been set.

### GetUsbLteSettings

`func (o *InternetOpenApiVO) GetUsbLteSettings() []UsbLteSettingOpenApiVO`

GetUsbLteSettings returns the UsbLteSettings field if non-nil, zero value otherwise.

### GetUsbLteSettingsOk

`func (o *InternetOpenApiVO) GetUsbLteSettingsOk() (*[]UsbLteSettingOpenApiVO, bool)`

GetUsbLteSettingsOk returns a tuple with the UsbLteSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbLteSettings

`func (o *InternetOpenApiVO) SetUsbLteSettings(v []UsbLteSettingOpenApiVO)`

SetUsbLteSettings sets UsbLteSettings field to given value.

### HasUsbLteSettings

`func (o *InternetOpenApiVO) HasUsbLteSettings() bool`

HasUsbLteSettings returns a boolean if a field has been set.

### GetWanLoadBalance

`func (o *InternetOpenApiVO) GetWanLoadBalance() WanLoadBalanceOpenApiVO`

GetWanLoadBalance returns the WanLoadBalance field if non-nil, zero value otherwise.

### GetWanLoadBalanceOk

`func (o *InternetOpenApiVO) GetWanLoadBalanceOk() (*WanLoadBalanceOpenApiVO, bool)`

GetWanLoadBalanceOk returns a tuple with the WanLoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanLoadBalance

`func (o *InternetOpenApiVO) SetWanLoadBalance(v WanLoadBalanceOpenApiVO)`

SetWanLoadBalance sets WanLoadBalance field to given value.

### HasWanLoadBalance

`func (o *InternetOpenApiVO) HasWanLoadBalance() bool`

HasWanLoadBalance returns a boolean if a field has been set.

### GetWanPortSettings

`func (o *InternetOpenApiVO) GetWanPortSettings() []WanPortSettingOpenApiVO`

GetWanPortSettings returns the WanPortSettings field if non-nil, zero value otherwise.

### GetWanPortSettingsOk

`func (o *InternetOpenApiVO) GetWanPortSettingsOk() (*[]WanPortSettingOpenApiVO, bool)`

GetWanPortSettingsOk returns a tuple with the WanPortSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortSettings

`func (o *InternetOpenApiVO) SetWanPortSettings(v []WanPortSettingOpenApiVO)`

SetWanPortSettings sets WanPortSettings field to given value.

### HasWanPortSettings

`func (o *InternetOpenApiVO) HasWanPortSettings() bool`

HasWanPortSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


