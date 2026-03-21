# OswPortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site to which the port belongs | [optional] 
**AggregatingPorts** | Pointer to **[]int32** | Aggregated ports (including this port), non-null only in aggregating mode | [optional] 
**AllAggregatingPorts** | Pointer to **[]int32** | All ports in the aggregating state on the Switch | [optional] 
**AllAggregatingStPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | All aggregating ports of the current Stack | [optional] 
**AllMirroredPorts** | Pointer to **[]int32** | All mirrored ports on the Switch | [optional] 
**AllMirroredStPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | All mirrored ports of the current Stack | [optional] 
**AllMirroringPorts** | Pointer to **[]int32** | All ports in the mirroring state on the Switch | [optional] 
**AllMirroringStPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | All mirroring ports of the current Stack | [optional] 
**AllMlagDadPorts** | Pointer to **[]int32** |  | [optional] 
**AllMlagPeerLinkPorts** | Pointer to **[]int32** |  | [optional] 
**BandCtrl** | Pointer to [**OswBandCtrlVO**](OswBandCtrlVO.md) |  | [optional] 
**BandWidthCtrlType** | Pointer to **int32** | BandWidth Control should be a value as follows: 0: Off; 1: Rate Limit; 2: Storming Control | [optional] 
**ConfigMlagDad** | Pointer to **bool** |  | [optional] 
**ConfigMlagPeerLink** | Pointer to **bool** |  | [optional] 
**ConfigStack** | Pointer to **bool** | Indicates whether the current port is configured as a stack port (joined a stack aggregation group) | [optional] 
**DhcpL2RelaySettings** | Pointer to [**OswPortDhcpL2RelayVO**](OswPortDhcpL2RelayVO.md) |  | [optional] 
**Disable** | Pointer to **bool** | Indicates whether to disable the port | [optional] 
**Dot1pPriority** | Pointer to **int32** | Dot1p Priority | [optional] 
**Dot1x** | Pointer to **int32** | Dot1x should be a value as follows: 0: Force unauthorized; 1: Force authorized; 2: Auto | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**EeeEnable** | Pointer to **bool** | Indicates whether EEE is enabled | [optional] 
**EsEnableAllProfileCanAdd** | Pointer to **bool** | Indicates whether es enable of all profile can continue to add VLANs | [optional] 
**EsQosSupport** | Pointer to **bool** | Indicates whether the ES device port supports modification of QoS configuration | [optional] 
**ExtendModeEnable** | Pointer to **bool** | Indicates whether extendMode is enabled | [optional] 
**ExtendModeSupport** | Pointer to **bool** |  | [optional] 
**FastLeaveEnable** | Pointer to **bool** | Indicates whether igmpSnooping fastLeave is enabled | [optional] 
**FecCap** | Pointer to [**[]OswFecCapVO**](OswFecCapVO.md) |  | [optional] 
**FecLinkPeerApplyEnable** | Pointer to **bool** |  | [optional] 
**FecMode** | Pointer to **int32** |  | [optional] 
**FecSupport** | Pointer to **bool** |  | [optional] 
**FlowControlEnable** | Pointer to **bool** | Indicates whether flow control is enabled | [optional] 
**Id** | Pointer to **string** | ID | [optional] 
**IgmpFastLeaveEnable** | Pointer to **bool** |  | [optional] 
**IgmpSnoopingEnable** | Pointer to **bool** | Indicates whether IGMP Snooping is enabled | [optional] 
**LagSetting** | Pointer to [**OswLagVO**](OswLagVO.md) |  | [optional] 
**LagStatus** | Pointer to [**OswLagStatusVO**](OswLagStatusVO.md) |  | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed should be a value as follows: 0: auto; 1: 10M; 2: 100M; 3: 1000M; 4: 10G | [optional] 
**LldpMedEnable** | Pointer to **bool** | Indicates whether LLDP MED is enabled | [optional] 
**LocateEnable** | Pointer to **bool** | Whether locate function is enabled | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | Indicates whether loopbackDetect port based is enabled | [optional] 
**LoopbackDetectVlanBasedEnable** | Pointer to **bool** | Indicates whether loopbackDetect vlan based is enabled | [optional] 
**MadUsed** | Pointer to **bool** | Mad Used | [optional] 
**MaxSpeed** | Pointer to **int32** | MaxSpeed should be a value as follows: 1: 10Mbps; 2: 100Mbps; 3: 1000Mbps; 4: 2.5Gbps; 5: 10Gbps | [optional] 
**MirroredLags** | Pointer to [**[]MirroredLag**](MirroredLag.md) | Monitored LAG | [optional] 
**MirroredPorts** | Pointer to [**[]MirroredPort**](MirroredPort.md) | Monitored Port | [optional] 
**MlagPeerAllPortsConfigInfo** | Pointer to [**OswMlagPeerAllPortsConfigInfoVO**](OswMlagPeerAllPortsConfigInfoVO.md) |  | [optional] 
**MldFastLeaveEnable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | Port Name | [optional] 
**NativeBridgeVlan** | Pointer to **int32** | Native Network Bridge Vlan. | [optional] 
**NativeNetworkId** | Pointer to **string** | Native Network ID, Native Network cannot be selected from Tagged Networks or Untagged Networks. | [optional] 
**NetworkConflict** | Pointer to **bool** | Indicates whether the VLAN configuration activated on the port is inconsistent with the VLAN configuration in the Profile. | [optional] 
**NetworkMode** | Pointer to **int32** | Network Mode should be a value as follows: 0: Trunk, 1: Access | [optional] 
**NetworkTagsSetting** | Pointer to **int32** | Network Tags Setting should be a value as follows: 0: Allow All; 1: Block All; 2: Custom | [optional] 
**Operation** | Pointer to **string** | Operation should be a value as follows: SWITCHING; MIRRORING; AGGREGATING | [optional] 
**Poe** | Pointer to **int32** | PoE switch should be a value as follows: 0: Off; 1: 802.3at/af | [optional] 
**PoeDisplayType** | Pointer to **int32** | PoeDisplayType should be a value as follows: -1: Not Support POE; 0: Support POE; 1: POE(4W); 2: POE(7W); 3: POE(15.4W); 4: POE+(30W); 5: POE++(45W); 6: POE++(60W); 7: POE++(75W); 8: POE++(90W); 9: POE++(100W). | [optional] 
**Port** | Pointer to **int32** | Port | [optional] 
**PortAlertEnable** | Pointer to **bool** | Indicates whether Port Alert is enabled | [optional] 
**PortCap** | Pointer to [**[]OswLinkCapVO**](OswLinkCapVO.md) | Port Capability | [optional] 
**PortIsolationEnable** | Pointer to **bool** | Indicates whether port isolation is enabled | [optional] 
**PortSpeedCap** | Pointer to [**[]OswPortSpeedCapVO**](OswPortSpeedCapVO.md) | Port Speed Capability | [optional] 
**PortStatus** | Pointer to [**OswPortStatusVO**](OswPortStatusVO.md) |  | [optional] 
**ProfileId** | Pointer to **string** | Lan Profile ID | [optional] 
**ProfileName** | Pointer to **string** | Lan Profile Name | [optional] 
**ProfileOverrideEnable** | Pointer to **bool** | Indicates whether to enable Profile Override | [optional] 
**QosQueueEnable** | Pointer to **bool** | Indicates whether the ES device port has enabled the Qos scheduling queue configuration | [optional] 
**QosSupport** | Pointer to **bool** | Indicates whether the device port supports modifying qos configuration | [optional] 
**QueueId** | Pointer to **int32** | ES Qos scheduling queue ID | [optional] 
**Resource** | Pointer to **int32** | Data Source. Resource should be a value as follows: 0: new created; 1: from template; 2: override | [optional] 
**SpanningTreeEnable** | Pointer to **bool** | Indicates whether SpanningTree is enabled | [optional] 
**SpanningTreeSetting** | Pointer to [**SpanningTreeSettingVO**](SpanningTreeSettingVO.md) |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**StackPortsGroupIndex** | Pointer to **int32** | Number of the stacking port aggregation group to join | [optional] 
**StandardPort** | Pointer to [**OswStandPortVO**](OswStandPortVO.md) |  | [optional] 
**StormCtrl** | Pointer to [**OswStormCtrlVO**](OswStormCtrlVO.md) |  | [optional] 
**SupportLocate** | Pointer to **bool** | Whether the port support locate. | [optional] 
**SupportPoe** | Pointer to **bool** | Indicates whether PoE is supported | [optional] 
**SwitchId** | Pointer to **string** | Switch ID to which the port belongs | [optional] 
**SwitchMac** | Pointer to **string** | Switch Mac to which the port belongs | [optional] 
**TagBridgeVlanMap** | Pointer to **map[string][]int32** | Tag Network Bridge Vlan Map | [optional] 
**TagIds** | Pointer to **[]string** | Port label ID List | [optional] 
**TagName** | Pointer to **string** | Port label Name | [optional] 
**TagNetworkIds** | Pointer to **[]string** | Tag Network IDs | [optional] 
**TopoNotifyEnable** | Pointer to **bool** | Indicates whether to enable topology change notification | [optional] 
**TrustMode** | Pointer to **int32** | TrustMode should be a value as follows: 0: Untrusted; 1: Trust 802.1p; 2: Trust DSCP | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 1: Copper; 2: Combo; 3: SFP | [optional] 
**UntagBridgeVlanMap** | Pointer to **map[string][]int32** | Untag Network Bridge Vlan Map | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag Network IDs | [optional] 
**VoiceBridgeVlan** | Pointer to **int32** | Voice Network Bridge Vlan | [optional] 
**VoiceDscp** | Pointer to **int32** | Voice DSCP | [optional] 
**VoiceDscpEnable** | Pointer to **bool** | Voice DSCP enable status | [optional] 
**VoiceNetworkEnable** | Pointer to **bool** | Voice network enable status | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewOswPortVO

`func NewOswPortVO() *OswPortVO`

NewOswPortVO instantiates a new OswPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortVOWithDefaults

`func NewOswPortVOWithDefaults() *OswPortVO`

NewOswPortVOWithDefaults instantiates a new OswPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswPortVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswPortVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswPortVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswPortVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAggregatingPorts

`func (o *OswPortVO) GetAggregatingPorts() []int32`

GetAggregatingPorts returns the AggregatingPorts field if non-nil, zero value otherwise.

### GetAggregatingPortsOk

`func (o *OswPortVO) GetAggregatingPortsOk() (*[]int32, bool)`

GetAggregatingPortsOk returns a tuple with the AggregatingPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatingPorts

`func (o *OswPortVO) SetAggregatingPorts(v []int32)`

SetAggregatingPorts sets AggregatingPorts field to given value.

### HasAggregatingPorts

`func (o *OswPortVO) HasAggregatingPorts() bool`

HasAggregatingPorts returns a boolean if a field has been set.

### GetAllAggregatingPorts

`func (o *OswPortVO) GetAllAggregatingPorts() []int32`

GetAllAggregatingPorts returns the AllAggregatingPorts field if non-nil, zero value otherwise.

### GetAllAggregatingPortsOk

`func (o *OswPortVO) GetAllAggregatingPortsOk() (*[]int32, bool)`

GetAllAggregatingPortsOk returns a tuple with the AllAggregatingPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAggregatingPorts

`func (o *OswPortVO) SetAllAggregatingPorts(v []int32)`

SetAllAggregatingPorts sets AllAggregatingPorts field to given value.

### HasAllAggregatingPorts

`func (o *OswPortVO) HasAllAggregatingPorts() bool`

HasAllAggregatingPorts returns a boolean if a field has been set.

### GetAllAggregatingStPorts

`func (o *OswPortVO) GetAllAggregatingStPorts() []OswStandPortVO`

GetAllAggregatingStPorts returns the AllAggregatingStPorts field if non-nil, zero value otherwise.

### GetAllAggregatingStPortsOk

`func (o *OswPortVO) GetAllAggregatingStPortsOk() (*[]OswStandPortVO, bool)`

GetAllAggregatingStPortsOk returns a tuple with the AllAggregatingStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAggregatingStPorts

`func (o *OswPortVO) SetAllAggregatingStPorts(v []OswStandPortVO)`

SetAllAggregatingStPorts sets AllAggregatingStPorts field to given value.

### HasAllAggregatingStPorts

`func (o *OswPortVO) HasAllAggregatingStPorts() bool`

HasAllAggregatingStPorts returns a boolean if a field has been set.

### GetAllMirroredPorts

`func (o *OswPortVO) GetAllMirroredPorts() []int32`

GetAllMirroredPorts returns the AllMirroredPorts field if non-nil, zero value otherwise.

### GetAllMirroredPortsOk

`func (o *OswPortVO) GetAllMirroredPortsOk() (*[]int32, bool)`

GetAllMirroredPortsOk returns a tuple with the AllMirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroredPorts

`func (o *OswPortVO) SetAllMirroredPorts(v []int32)`

SetAllMirroredPorts sets AllMirroredPorts field to given value.

### HasAllMirroredPorts

`func (o *OswPortVO) HasAllMirroredPorts() bool`

HasAllMirroredPorts returns a boolean if a field has been set.

### GetAllMirroredStPorts

`func (o *OswPortVO) GetAllMirroredStPorts() []OswStandPortVO`

GetAllMirroredStPorts returns the AllMirroredStPorts field if non-nil, zero value otherwise.

### GetAllMirroredStPortsOk

`func (o *OswPortVO) GetAllMirroredStPortsOk() (*[]OswStandPortVO, bool)`

GetAllMirroredStPortsOk returns a tuple with the AllMirroredStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroredStPorts

`func (o *OswPortVO) SetAllMirroredStPorts(v []OswStandPortVO)`

SetAllMirroredStPorts sets AllMirroredStPorts field to given value.

### HasAllMirroredStPorts

`func (o *OswPortVO) HasAllMirroredStPorts() bool`

HasAllMirroredStPorts returns a boolean if a field has been set.

### GetAllMirroringPorts

`func (o *OswPortVO) GetAllMirroringPorts() []int32`

GetAllMirroringPorts returns the AllMirroringPorts field if non-nil, zero value otherwise.

### GetAllMirroringPortsOk

`func (o *OswPortVO) GetAllMirroringPortsOk() (*[]int32, bool)`

GetAllMirroringPortsOk returns a tuple with the AllMirroringPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroringPorts

`func (o *OswPortVO) SetAllMirroringPorts(v []int32)`

SetAllMirroringPorts sets AllMirroringPorts field to given value.

### HasAllMirroringPorts

`func (o *OswPortVO) HasAllMirroringPorts() bool`

HasAllMirroringPorts returns a boolean if a field has been set.

### GetAllMirroringStPorts

`func (o *OswPortVO) GetAllMirroringStPorts() []OswStandPortVO`

GetAllMirroringStPorts returns the AllMirroringStPorts field if non-nil, zero value otherwise.

### GetAllMirroringStPortsOk

`func (o *OswPortVO) GetAllMirroringStPortsOk() (*[]OswStandPortVO, bool)`

GetAllMirroringStPortsOk returns a tuple with the AllMirroringStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroringStPorts

`func (o *OswPortVO) SetAllMirroringStPorts(v []OswStandPortVO)`

SetAllMirroringStPorts sets AllMirroringStPorts field to given value.

### HasAllMirroringStPorts

`func (o *OswPortVO) HasAllMirroringStPorts() bool`

HasAllMirroringStPorts returns a boolean if a field has been set.

### GetAllMlagDadPorts

`func (o *OswPortVO) GetAllMlagDadPorts() []int32`

GetAllMlagDadPorts returns the AllMlagDadPorts field if non-nil, zero value otherwise.

### GetAllMlagDadPortsOk

`func (o *OswPortVO) GetAllMlagDadPortsOk() (*[]int32, bool)`

GetAllMlagDadPortsOk returns a tuple with the AllMlagDadPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMlagDadPorts

`func (o *OswPortVO) SetAllMlagDadPorts(v []int32)`

SetAllMlagDadPorts sets AllMlagDadPorts field to given value.

### HasAllMlagDadPorts

`func (o *OswPortVO) HasAllMlagDadPorts() bool`

HasAllMlagDadPorts returns a boolean if a field has been set.

### GetAllMlagPeerLinkPorts

`func (o *OswPortVO) GetAllMlagPeerLinkPorts() []int32`

GetAllMlagPeerLinkPorts returns the AllMlagPeerLinkPorts field if non-nil, zero value otherwise.

### GetAllMlagPeerLinkPortsOk

`func (o *OswPortVO) GetAllMlagPeerLinkPortsOk() (*[]int32, bool)`

GetAllMlagPeerLinkPortsOk returns a tuple with the AllMlagPeerLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMlagPeerLinkPorts

`func (o *OswPortVO) SetAllMlagPeerLinkPorts(v []int32)`

SetAllMlagPeerLinkPorts sets AllMlagPeerLinkPorts field to given value.

### HasAllMlagPeerLinkPorts

`func (o *OswPortVO) HasAllMlagPeerLinkPorts() bool`

HasAllMlagPeerLinkPorts returns a boolean if a field has been set.

### GetBandCtrl

`func (o *OswPortVO) GetBandCtrl() OswBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *OswPortVO) GetBandCtrlOk() (*OswBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *OswPortVO) SetBandCtrl(v OswBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *OswPortVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *OswPortVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *OswPortVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *OswPortVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *OswPortVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetConfigMlagDad

`func (o *OswPortVO) GetConfigMlagDad() bool`

GetConfigMlagDad returns the ConfigMlagDad field if non-nil, zero value otherwise.

### GetConfigMlagDadOk

`func (o *OswPortVO) GetConfigMlagDadOk() (*bool, bool)`

GetConfigMlagDadOk returns a tuple with the ConfigMlagDad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagDad

`func (o *OswPortVO) SetConfigMlagDad(v bool)`

SetConfigMlagDad sets ConfigMlagDad field to given value.

### HasConfigMlagDad

`func (o *OswPortVO) HasConfigMlagDad() bool`

HasConfigMlagDad returns a boolean if a field has been set.

### GetConfigMlagPeerLink

`func (o *OswPortVO) GetConfigMlagPeerLink() bool`

GetConfigMlagPeerLink returns the ConfigMlagPeerLink field if non-nil, zero value otherwise.

### GetConfigMlagPeerLinkOk

`func (o *OswPortVO) GetConfigMlagPeerLinkOk() (*bool, bool)`

GetConfigMlagPeerLinkOk returns a tuple with the ConfigMlagPeerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagPeerLink

`func (o *OswPortVO) SetConfigMlagPeerLink(v bool)`

SetConfigMlagPeerLink sets ConfigMlagPeerLink field to given value.

### HasConfigMlagPeerLink

`func (o *OswPortVO) HasConfigMlagPeerLink() bool`

HasConfigMlagPeerLink returns a boolean if a field has been set.

### GetConfigStack

`func (o *OswPortVO) GetConfigStack() bool`

GetConfigStack returns the ConfigStack field if non-nil, zero value otherwise.

### GetConfigStackOk

`func (o *OswPortVO) GetConfigStackOk() (*bool, bool)`

GetConfigStackOk returns a tuple with the ConfigStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStack

`func (o *OswPortVO) SetConfigStack(v bool)`

SetConfigStack sets ConfigStack field to given value.

### HasConfigStack

`func (o *OswPortVO) HasConfigStack() bool`

HasConfigStack returns a boolean if a field has been set.

### GetDhcpL2RelaySettings

`func (o *OswPortVO) GetDhcpL2RelaySettings() OswPortDhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *OswPortVO) GetDhcpL2RelaySettingsOk() (*OswPortDhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *OswPortVO) SetDhcpL2RelaySettings(v OswPortDhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *OswPortVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDisable

`func (o *OswPortVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswPortVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswPortVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswPortVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *OswPortVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *OswPortVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *OswPortVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *OswPortVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDot1x

`func (o *OswPortVO) GetDot1x() int32`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *OswPortVO) GetDot1xOk() (*int32, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *OswPortVO) SetDot1x(v int32)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *OswPortVO) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetDuplex

`func (o *OswPortVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswPortVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswPortVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswPortVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEeeEnable

`func (o *OswPortVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *OswPortVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *OswPortVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *OswPortVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetEsEnableAllProfileCanAdd

`func (o *OswPortVO) GetEsEnableAllProfileCanAdd() bool`

GetEsEnableAllProfileCanAdd returns the EsEnableAllProfileCanAdd field if non-nil, zero value otherwise.

### GetEsEnableAllProfileCanAddOk

`func (o *OswPortVO) GetEsEnableAllProfileCanAddOk() (*bool, bool)`

GetEsEnableAllProfileCanAddOk returns a tuple with the EsEnableAllProfileCanAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsEnableAllProfileCanAdd

`func (o *OswPortVO) SetEsEnableAllProfileCanAdd(v bool)`

SetEsEnableAllProfileCanAdd sets EsEnableAllProfileCanAdd field to given value.

### HasEsEnableAllProfileCanAdd

`func (o *OswPortVO) HasEsEnableAllProfileCanAdd() bool`

HasEsEnableAllProfileCanAdd returns a boolean if a field has been set.

### GetEsQosSupport

`func (o *OswPortVO) GetEsQosSupport() bool`

GetEsQosSupport returns the EsQosSupport field if non-nil, zero value otherwise.

### GetEsQosSupportOk

`func (o *OswPortVO) GetEsQosSupportOk() (*bool, bool)`

GetEsQosSupportOk returns a tuple with the EsQosSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsQosSupport

`func (o *OswPortVO) SetEsQosSupport(v bool)`

SetEsQosSupport sets EsQosSupport field to given value.

### HasEsQosSupport

`func (o *OswPortVO) HasEsQosSupport() bool`

HasEsQosSupport returns a boolean if a field has been set.

### GetExtendModeEnable

`func (o *OswPortVO) GetExtendModeEnable() bool`

GetExtendModeEnable returns the ExtendModeEnable field if non-nil, zero value otherwise.

### GetExtendModeEnableOk

`func (o *OswPortVO) GetExtendModeEnableOk() (*bool, bool)`

GetExtendModeEnableOk returns a tuple with the ExtendModeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendModeEnable

`func (o *OswPortVO) SetExtendModeEnable(v bool)`

SetExtendModeEnable sets ExtendModeEnable field to given value.

### HasExtendModeEnable

`func (o *OswPortVO) HasExtendModeEnable() bool`

HasExtendModeEnable returns a boolean if a field has been set.

### GetExtendModeSupport

`func (o *OswPortVO) GetExtendModeSupport() bool`

GetExtendModeSupport returns the ExtendModeSupport field if non-nil, zero value otherwise.

### GetExtendModeSupportOk

`func (o *OswPortVO) GetExtendModeSupportOk() (*bool, bool)`

GetExtendModeSupportOk returns a tuple with the ExtendModeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendModeSupport

`func (o *OswPortVO) SetExtendModeSupport(v bool)`

SetExtendModeSupport sets ExtendModeSupport field to given value.

### HasExtendModeSupport

`func (o *OswPortVO) HasExtendModeSupport() bool`

HasExtendModeSupport returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *OswPortVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *OswPortVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *OswPortVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *OswPortVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFecCap

`func (o *OswPortVO) GetFecCap() []OswFecCapVO`

GetFecCap returns the FecCap field if non-nil, zero value otherwise.

### GetFecCapOk

`func (o *OswPortVO) GetFecCapOk() (*[]OswFecCapVO, bool)`

GetFecCapOk returns a tuple with the FecCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecCap

`func (o *OswPortVO) SetFecCap(v []OswFecCapVO)`

SetFecCap sets FecCap field to given value.

### HasFecCap

`func (o *OswPortVO) HasFecCap() bool`

HasFecCap returns a boolean if a field has been set.

### GetFecLinkPeerApplyEnable

`func (o *OswPortVO) GetFecLinkPeerApplyEnable() bool`

GetFecLinkPeerApplyEnable returns the FecLinkPeerApplyEnable field if non-nil, zero value otherwise.

### GetFecLinkPeerApplyEnableOk

`func (o *OswPortVO) GetFecLinkPeerApplyEnableOk() (*bool, bool)`

GetFecLinkPeerApplyEnableOk returns a tuple with the FecLinkPeerApplyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecLinkPeerApplyEnable

`func (o *OswPortVO) SetFecLinkPeerApplyEnable(v bool)`

SetFecLinkPeerApplyEnable sets FecLinkPeerApplyEnable field to given value.

### HasFecLinkPeerApplyEnable

`func (o *OswPortVO) HasFecLinkPeerApplyEnable() bool`

HasFecLinkPeerApplyEnable returns a boolean if a field has been set.

### GetFecMode

`func (o *OswPortVO) GetFecMode() int32`

GetFecMode returns the FecMode field if non-nil, zero value otherwise.

### GetFecModeOk

`func (o *OswPortVO) GetFecModeOk() (*int32, bool)`

GetFecModeOk returns a tuple with the FecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecMode

`func (o *OswPortVO) SetFecMode(v int32)`

SetFecMode sets FecMode field to given value.

### HasFecMode

`func (o *OswPortVO) HasFecMode() bool`

HasFecMode returns a boolean if a field has been set.

### GetFecSupport

`func (o *OswPortVO) GetFecSupport() bool`

GetFecSupport returns the FecSupport field if non-nil, zero value otherwise.

### GetFecSupportOk

`func (o *OswPortVO) GetFecSupportOk() (*bool, bool)`

GetFecSupportOk returns a tuple with the FecSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecSupport

`func (o *OswPortVO) SetFecSupport(v bool)`

SetFecSupport sets FecSupport field to given value.

### HasFecSupport

`func (o *OswPortVO) HasFecSupport() bool`

HasFecSupport returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *OswPortVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *OswPortVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *OswPortVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *OswPortVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetId

`func (o *OswPortVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswPortVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswPortVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswPortVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *OswPortVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *OswPortVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *OswPortVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *OswPortVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetIgmpSnoopingEnable

`func (o *OswPortVO) GetIgmpSnoopingEnable() bool`

GetIgmpSnoopingEnable returns the IgmpSnoopingEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnableOk

`func (o *OswPortVO) GetIgmpSnoopingEnableOk() (*bool, bool)`

GetIgmpSnoopingEnableOk returns a tuple with the IgmpSnoopingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnable

`func (o *OswPortVO) SetIgmpSnoopingEnable(v bool)`

SetIgmpSnoopingEnable sets IgmpSnoopingEnable field to given value.

### HasIgmpSnoopingEnable

`func (o *OswPortVO) HasIgmpSnoopingEnable() bool`

HasIgmpSnoopingEnable returns a boolean if a field has been set.

### GetLagSetting

`func (o *OswPortVO) GetLagSetting() OswLagVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *OswPortVO) GetLagSettingOk() (*OswLagVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *OswPortVO) SetLagSetting(v OswLagVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *OswPortVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetLagStatus

`func (o *OswPortVO) GetLagStatus() OswLagStatusVO`

GetLagStatus returns the LagStatus field if non-nil, zero value otherwise.

### GetLagStatusOk

`func (o *OswPortVO) GetLagStatusOk() (*OswLagStatusVO, bool)`

GetLagStatusOk returns a tuple with the LagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagStatus

`func (o *OswPortVO) SetLagStatus(v OswLagStatusVO)`

SetLagStatus sets LagStatus field to given value.

### HasLagStatus

`func (o *OswPortVO) HasLagStatus() bool`

HasLagStatus returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswPortVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswPortVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswPortVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswPortVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLldpMedEnable

`func (o *OswPortVO) GetLldpMedEnable() bool`

GetLldpMedEnable returns the LldpMedEnable field if non-nil, zero value otherwise.

### GetLldpMedEnableOk

`func (o *OswPortVO) GetLldpMedEnableOk() (*bool, bool)`

GetLldpMedEnableOk returns a tuple with the LldpMedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMedEnable

`func (o *OswPortVO) SetLldpMedEnable(v bool)`

SetLldpMedEnable sets LldpMedEnable field to given value.

### HasLldpMedEnable

`func (o *OswPortVO) HasLldpMedEnable() bool`

HasLldpMedEnable returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswPortVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswPortVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswPortVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswPortVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswPortVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswPortVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswPortVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswPortVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackDetectVlanBasedEnable

`func (o *OswPortVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *OswPortVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *OswPortVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *OswPortVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMadUsed

`func (o *OswPortVO) GetMadUsed() bool`

GetMadUsed returns the MadUsed field if non-nil, zero value otherwise.

### GetMadUsedOk

`func (o *OswPortVO) GetMadUsedOk() (*bool, bool)`

GetMadUsedOk returns a tuple with the MadUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadUsed

`func (o *OswPortVO) SetMadUsed(v bool)`

SetMadUsed sets MadUsed field to given value.

### HasMadUsed

`func (o *OswPortVO) HasMadUsed() bool`

HasMadUsed returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *OswPortVO) GetMaxSpeed() int32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *OswPortVO) GetMaxSpeedOk() (*int32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *OswPortVO) SetMaxSpeed(v int32)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *OswPortVO) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetMirroredLags

`func (o *OswPortVO) GetMirroredLags() []MirroredLag`

GetMirroredLags returns the MirroredLags field if non-nil, zero value otherwise.

### GetMirroredLagsOk

`func (o *OswPortVO) GetMirroredLagsOk() (*[]MirroredLag, bool)`

GetMirroredLagsOk returns a tuple with the MirroredLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredLags

`func (o *OswPortVO) SetMirroredLags(v []MirroredLag)`

SetMirroredLags sets MirroredLags field to given value.

### HasMirroredLags

`func (o *OswPortVO) HasMirroredLags() bool`

HasMirroredLags returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *OswPortVO) GetMirroredPorts() []MirroredPort`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *OswPortVO) GetMirroredPortsOk() (*[]MirroredPort, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *OswPortVO) SetMirroredPorts(v []MirroredPort)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *OswPortVO) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetMlagPeerAllPortsConfigInfo

`func (o *OswPortVO) GetMlagPeerAllPortsConfigInfo() OswMlagPeerAllPortsConfigInfoVO`

GetMlagPeerAllPortsConfigInfo returns the MlagPeerAllPortsConfigInfo field if non-nil, zero value otherwise.

### GetMlagPeerAllPortsConfigInfoOk

`func (o *OswPortVO) GetMlagPeerAllPortsConfigInfoOk() (*OswMlagPeerAllPortsConfigInfoVO, bool)`

GetMlagPeerAllPortsConfigInfoOk returns a tuple with the MlagPeerAllPortsConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerAllPortsConfigInfo

`func (o *OswPortVO) SetMlagPeerAllPortsConfigInfo(v OswMlagPeerAllPortsConfigInfoVO)`

SetMlagPeerAllPortsConfigInfo sets MlagPeerAllPortsConfigInfo field to given value.

### HasMlagPeerAllPortsConfigInfo

`func (o *OswPortVO) HasMlagPeerAllPortsConfigInfo() bool`

HasMlagPeerAllPortsConfigInfo returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *OswPortVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *OswPortVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *OswPortVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *OswPortVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *OswPortVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswPortVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswPortVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswPortVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *OswPortVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *OswPortVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *OswPortVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *OswPortVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *OswPortVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *OswPortVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *OswPortVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *OswPortVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkConflict

`func (o *OswPortVO) GetNetworkConflict() bool`

GetNetworkConflict returns the NetworkConflict field if non-nil, zero value otherwise.

### GetNetworkConflictOk

`func (o *OswPortVO) GetNetworkConflictOk() (*bool, bool)`

GetNetworkConflictOk returns a tuple with the NetworkConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConflict

`func (o *OswPortVO) SetNetworkConflict(v bool)`

SetNetworkConflict sets NetworkConflict field to given value.

### HasNetworkConflict

`func (o *OswPortVO) HasNetworkConflict() bool`

HasNetworkConflict returns a boolean if a field has been set.

### GetNetworkMode

`func (o *OswPortVO) GetNetworkMode() int32`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *OswPortVO) GetNetworkModeOk() (*int32, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *OswPortVO) SetNetworkMode(v int32)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *OswPortVO) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *OswPortVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *OswPortVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *OswPortVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *OswPortVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetOperation

`func (o *OswPortVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OswPortVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OswPortVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *OswPortVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPoe

`func (o *OswPortVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OswPortVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OswPortVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OswPortVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPoeDisplayType

`func (o *OswPortVO) GetPoeDisplayType() int32`

GetPoeDisplayType returns the PoeDisplayType field if non-nil, zero value otherwise.

### GetPoeDisplayTypeOk

`func (o *OswPortVO) GetPoeDisplayTypeOk() (*int32, bool)`

GetPoeDisplayTypeOk returns a tuple with the PoeDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisplayType

`func (o *OswPortVO) SetPoeDisplayType(v int32)`

SetPoeDisplayType sets PoeDisplayType field to given value.

### HasPoeDisplayType

`func (o *OswPortVO) HasPoeDisplayType() bool`

HasPoeDisplayType returns a boolean if a field has been set.

### GetPort

`func (o *OswPortVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswPortVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswPortVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswPortVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortAlertEnable

`func (o *OswPortVO) GetPortAlertEnable() bool`

GetPortAlertEnable returns the PortAlertEnable field if non-nil, zero value otherwise.

### GetPortAlertEnableOk

`func (o *OswPortVO) GetPortAlertEnableOk() (*bool, bool)`

GetPortAlertEnableOk returns a tuple with the PortAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAlertEnable

`func (o *OswPortVO) SetPortAlertEnable(v bool)`

SetPortAlertEnable sets PortAlertEnable field to given value.

### HasPortAlertEnable

`func (o *OswPortVO) HasPortAlertEnable() bool`

HasPortAlertEnable returns a boolean if a field has been set.

### GetPortCap

`func (o *OswPortVO) GetPortCap() []OswLinkCapVO`

GetPortCap returns the PortCap field if non-nil, zero value otherwise.

### GetPortCapOk

`func (o *OswPortVO) GetPortCapOk() (*[]OswLinkCapVO, bool)`

GetPortCapOk returns a tuple with the PortCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCap

`func (o *OswPortVO) SetPortCap(v []OswLinkCapVO)`

SetPortCap sets PortCap field to given value.

### HasPortCap

`func (o *OswPortVO) HasPortCap() bool`

HasPortCap returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *OswPortVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *OswPortVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *OswPortVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *OswPortVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetPortSpeedCap

`func (o *OswPortVO) GetPortSpeedCap() []OswPortSpeedCapVO`

GetPortSpeedCap returns the PortSpeedCap field if non-nil, zero value otherwise.

### GetPortSpeedCapOk

`func (o *OswPortVO) GetPortSpeedCapOk() (*[]OswPortSpeedCapVO, bool)`

GetPortSpeedCapOk returns a tuple with the PortSpeedCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeedCap

`func (o *OswPortVO) SetPortSpeedCap(v []OswPortSpeedCapVO)`

SetPortSpeedCap sets PortSpeedCap field to given value.

### HasPortSpeedCap

`func (o *OswPortVO) HasPortSpeedCap() bool`

HasPortSpeedCap returns a boolean if a field has been set.

### GetPortStatus

`func (o *OswPortVO) GetPortStatus() OswPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *OswPortVO) GetPortStatusOk() (*OswPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *OswPortVO) SetPortStatus(v OswPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *OswPortVO) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetProfileId

`func (o *OswPortVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OswPortVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OswPortVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *OswPortVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *OswPortVO) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *OswPortVO) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *OswPortVO) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *OswPortVO) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *OswPortVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *OswPortVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *OswPortVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *OswPortVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetQosQueueEnable

`func (o *OswPortVO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *OswPortVO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *OswPortVO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *OswPortVO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQosSupport

`func (o *OswPortVO) GetQosSupport() bool`

GetQosSupport returns the QosSupport field if non-nil, zero value otherwise.

### GetQosSupportOk

`func (o *OswPortVO) GetQosSupportOk() (*bool, bool)`

GetQosSupportOk returns a tuple with the QosSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSupport

`func (o *OswPortVO) SetQosSupport(v bool)`

SetQosSupport sets QosSupport field to given value.

### HasQosSupport

`func (o *OswPortVO) HasQosSupport() bool`

HasQosSupport returns a boolean if a field has been set.

### GetQueueId

`func (o *OswPortVO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *OswPortVO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *OswPortVO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *OswPortVO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetResource

`func (o *OswPortVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswPortVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswPortVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswPortVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSpanningTreeEnable

`func (o *OswPortVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *OswPortVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *OswPortVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.

### HasSpanningTreeEnable

`func (o *OswPortVO) HasSpanningTreeEnable() bool`

HasSpanningTreeEnable returns a boolean if a field has been set.

### GetSpanningTreeSetting

`func (o *OswPortVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *OswPortVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *OswPortVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *OswPortVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetSpeed

`func (o *OswPortVO) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OswPortVO) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OswPortVO) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *OswPortVO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStackPortsGroupIndex

`func (o *OswPortVO) GetStackPortsGroupIndex() int32`

GetStackPortsGroupIndex returns the StackPortsGroupIndex field if non-nil, zero value otherwise.

### GetStackPortsGroupIndexOk

`func (o *OswPortVO) GetStackPortsGroupIndexOk() (*int32, bool)`

GetStackPortsGroupIndexOk returns a tuple with the StackPortsGroupIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortsGroupIndex

`func (o *OswPortVO) SetStackPortsGroupIndex(v int32)`

SetStackPortsGroupIndex sets StackPortsGroupIndex field to given value.

### HasStackPortsGroupIndex

`func (o *OswPortVO) HasStackPortsGroupIndex() bool`

HasStackPortsGroupIndex returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswPortVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswPortVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswPortVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswPortVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetStormCtrl

`func (o *OswPortVO) GetStormCtrl() OswStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *OswPortVO) GetStormCtrlOk() (*OswStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *OswPortVO) SetStormCtrl(v OswStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *OswPortVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetSupportLocate

`func (o *OswPortVO) GetSupportLocate() bool`

GetSupportLocate returns the SupportLocate field if non-nil, zero value otherwise.

### GetSupportLocateOk

`func (o *OswPortVO) GetSupportLocateOk() (*bool, bool)`

GetSupportLocateOk returns a tuple with the SupportLocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocate

`func (o *OswPortVO) SetSupportLocate(v bool)`

SetSupportLocate sets SupportLocate field to given value.

### HasSupportLocate

`func (o *OswPortVO) HasSupportLocate() bool`

HasSupportLocate returns a boolean if a field has been set.

### GetSupportPoe

`func (o *OswPortVO) GetSupportPoe() bool`

GetSupportPoe returns the SupportPoe field if non-nil, zero value otherwise.

### GetSupportPoeOk

`func (o *OswPortVO) GetSupportPoeOk() (*bool, bool)`

GetSupportPoeOk returns a tuple with the SupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPoe

`func (o *OswPortVO) SetSupportPoe(v bool)`

SetSupportPoe sets SupportPoe field to given value.

### HasSupportPoe

`func (o *OswPortVO) HasSupportPoe() bool`

HasSupportPoe returns a boolean if a field has been set.

### GetSwitchId

`func (o *OswPortVO) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *OswPortVO) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *OswPortVO) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *OswPortVO) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetSwitchMac

`func (o *OswPortVO) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *OswPortVO) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *OswPortVO) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.

### HasSwitchMac

`func (o *OswPortVO) HasSwitchMac() bool`

HasSwitchMac returns a boolean if a field has been set.

### GetTagBridgeVlanMap

`func (o *OswPortVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *OswPortVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *OswPortVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *OswPortVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagIds

`func (o *OswPortVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswPortVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswPortVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswPortVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagName

`func (o *OswPortVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *OswPortVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *OswPortVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *OswPortVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *OswPortVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *OswPortVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *OswPortVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *OswPortVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTopoNotifyEnable

`func (o *OswPortVO) GetTopoNotifyEnable() bool`

GetTopoNotifyEnable returns the TopoNotifyEnable field if non-nil, zero value otherwise.

### GetTopoNotifyEnableOk

`func (o *OswPortVO) GetTopoNotifyEnableOk() (*bool, bool)`

GetTopoNotifyEnableOk returns a tuple with the TopoNotifyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopoNotifyEnable

`func (o *OswPortVO) SetTopoNotifyEnable(v bool)`

SetTopoNotifyEnable sets TopoNotifyEnable field to given value.

### HasTopoNotifyEnable

`func (o *OswPortVO) HasTopoNotifyEnable() bool`

HasTopoNotifyEnable returns a boolean if a field has been set.

### GetTrustMode

`func (o *OswPortVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *OswPortVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *OswPortVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *OswPortVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetType

`func (o *OswPortVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswPortVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswPortVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswPortVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUntagBridgeVlanMap

`func (o *OswPortVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *OswPortVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *OswPortVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *OswPortVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *OswPortVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *OswPortVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *OswPortVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *OswPortVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *OswPortVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *OswPortVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *OswPortVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *OswPortVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceDscp

`func (o *OswPortVO) GetVoiceDscp() int32`

GetVoiceDscp returns the VoiceDscp field if non-nil, zero value otherwise.

### GetVoiceDscpOk

`func (o *OswPortVO) GetVoiceDscpOk() (*int32, bool)`

GetVoiceDscpOk returns a tuple with the VoiceDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscp

`func (o *OswPortVO) SetVoiceDscp(v int32)`

SetVoiceDscp sets VoiceDscp field to given value.

### HasVoiceDscp

`func (o *OswPortVO) HasVoiceDscp() bool`

HasVoiceDscp returns a boolean if a field has been set.

### GetVoiceDscpEnable

`func (o *OswPortVO) GetVoiceDscpEnable() bool`

GetVoiceDscpEnable returns the VoiceDscpEnable field if non-nil, zero value otherwise.

### GetVoiceDscpEnableOk

`func (o *OswPortVO) GetVoiceDscpEnableOk() (*bool, bool)`

GetVoiceDscpEnableOk returns a tuple with the VoiceDscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpEnable

`func (o *OswPortVO) SetVoiceDscpEnable(v bool)`

SetVoiceDscpEnable sets VoiceDscpEnable field to given value.

### HasVoiceDscpEnable

`func (o *OswPortVO) HasVoiceDscpEnable() bool`

HasVoiceDscpEnable returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *OswPortVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *OswPortVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *OswPortVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *OswPortVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *OswPortVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *OswPortVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *OswPortVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *OswPortVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


