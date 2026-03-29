# OswStackMemberPortVO

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
**StackId** | Pointer to **string** | Stack ID | [optional] 
**StackPortsGroupIndex** | Pointer to **int32** | Number of the stacking port aggregation group to join | [optional] 
**StackSetting** | Pointer to [**OswPortStackSettingVO**](OswPortStackSettingVO.md) |  | [optional] 
**StackSupport** | Pointer to **bool** | Stack Support | [optional] 
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

### NewOswStackMemberPortVO

`func NewOswStackMemberPortVO() *OswStackMemberPortVO`

NewOswStackMemberPortVO instantiates a new OswStackMemberPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackMemberPortVOWithDefaults

`func NewOswStackMemberPortVOWithDefaults() *OswStackMemberPortVO`

NewOswStackMemberPortVOWithDefaults instantiates a new OswStackMemberPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswStackMemberPortVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswStackMemberPortVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswStackMemberPortVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswStackMemberPortVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAggregatingPorts

`func (o *OswStackMemberPortVO) GetAggregatingPorts() []int32`

GetAggregatingPorts returns the AggregatingPorts field if non-nil, zero value otherwise.

### GetAggregatingPortsOk

`func (o *OswStackMemberPortVO) GetAggregatingPortsOk() (*[]int32, bool)`

GetAggregatingPortsOk returns a tuple with the AggregatingPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatingPorts

`func (o *OswStackMemberPortVO) SetAggregatingPorts(v []int32)`

SetAggregatingPorts sets AggregatingPorts field to given value.

### HasAggregatingPorts

`func (o *OswStackMemberPortVO) HasAggregatingPorts() bool`

HasAggregatingPorts returns a boolean if a field has been set.

### GetAllAggregatingPorts

`func (o *OswStackMemberPortVO) GetAllAggregatingPorts() []int32`

GetAllAggregatingPorts returns the AllAggregatingPorts field if non-nil, zero value otherwise.

### GetAllAggregatingPortsOk

`func (o *OswStackMemberPortVO) GetAllAggregatingPortsOk() (*[]int32, bool)`

GetAllAggregatingPortsOk returns a tuple with the AllAggregatingPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAggregatingPorts

`func (o *OswStackMemberPortVO) SetAllAggregatingPorts(v []int32)`

SetAllAggregatingPorts sets AllAggregatingPorts field to given value.

### HasAllAggregatingPorts

`func (o *OswStackMemberPortVO) HasAllAggregatingPorts() bool`

HasAllAggregatingPorts returns a boolean if a field has been set.

### GetAllAggregatingStPorts

`func (o *OswStackMemberPortVO) GetAllAggregatingStPorts() []OswStandPortVO`

GetAllAggregatingStPorts returns the AllAggregatingStPorts field if non-nil, zero value otherwise.

### GetAllAggregatingStPortsOk

`func (o *OswStackMemberPortVO) GetAllAggregatingStPortsOk() (*[]OswStandPortVO, bool)`

GetAllAggregatingStPortsOk returns a tuple with the AllAggregatingStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAggregatingStPorts

`func (o *OswStackMemberPortVO) SetAllAggregatingStPorts(v []OswStandPortVO)`

SetAllAggregatingStPorts sets AllAggregatingStPorts field to given value.

### HasAllAggregatingStPorts

`func (o *OswStackMemberPortVO) HasAllAggregatingStPorts() bool`

HasAllAggregatingStPorts returns a boolean if a field has been set.

### GetAllMirroredPorts

`func (o *OswStackMemberPortVO) GetAllMirroredPorts() []int32`

GetAllMirroredPorts returns the AllMirroredPorts field if non-nil, zero value otherwise.

### GetAllMirroredPortsOk

`func (o *OswStackMemberPortVO) GetAllMirroredPortsOk() (*[]int32, bool)`

GetAllMirroredPortsOk returns a tuple with the AllMirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroredPorts

`func (o *OswStackMemberPortVO) SetAllMirroredPorts(v []int32)`

SetAllMirroredPorts sets AllMirroredPorts field to given value.

### HasAllMirroredPorts

`func (o *OswStackMemberPortVO) HasAllMirroredPorts() bool`

HasAllMirroredPorts returns a boolean if a field has been set.

### GetAllMirroredStPorts

`func (o *OswStackMemberPortVO) GetAllMirroredStPorts() []OswStandPortVO`

GetAllMirroredStPorts returns the AllMirroredStPorts field if non-nil, zero value otherwise.

### GetAllMirroredStPortsOk

`func (o *OswStackMemberPortVO) GetAllMirroredStPortsOk() (*[]OswStandPortVO, bool)`

GetAllMirroredStPortsOk returns a tuple with the AllMirroredStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroredStPorts

`func (o *OswStackMemberPortVO) SetAllMirroredStPorts(v []OswStandPortVO)`

SetAllMirroredStPorts sets AllMirroredStPorts field to given value.

### HasAllMirroredStPorts

`func (o *OswStackMemberPortVO) HasAllMirroredStPorts() bool`

HasAllMirroredStPorts returns a boolean if a field has been set.

### GetAllMirroringPorts

`func (o *OswStackMemberPortVO) GetAllMirroringPorts() []int32`

GetAllMirroringPorts returns the AllMirroringPorts field if non-nil, zero value otherwise.

### GetAllMirroringPortsOk

`func (o *OswStackMemberPortVO) GetAllMirroringPortsOk() (*[]int32, bool)`

GetAllMirroringPortsOk returns a tuple with the AllMirroringPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroringPorts

`func (o *OswStackMemberPortVO) SetAllMirroringPorts(v []int32)`

SetAllMirroringPorts sets AllMirroringPorts field to given value.

### HasAllMirroringPorts

`func (o *OswStackMemberPortVO) HasAllMirroringPorts() bool`

HasAllMirroringPorts returns a boolean if a field has been set.

### GetAllMirroringStPorts

`func (o *OswStackMemberPortVO) GetAllMirroringStPorts() []OswStandPortVO`

GetAllMirroringStPorts returns the AllMirroringStPorts field if non-nil, zero value otherwise.

### GetAllMirroringStPortsOk

`func (o *OswStackMemberPortVO) GetAllMirroringStPortsOk() (*[]OswStandPortVO, bool)`

GetAllMirroringStPortsOk returns a tuple with the AllMirroringStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroringStPorts

`func (o *OswStackMemberPortVO) SetAllMirroringStPorts(v []OswStandPortVO)`

SetAllMirroringStPorts sets AllMirroringStPorts field to given value.

### HasAllMirroringStPorts

`func (o *OswStackMemberPortVO) HasAllMirroringStPorts() bool`

HasAllMirroringStPorts returns a boolean if a field has been set.

### GetAllMlagDadPorts

`func (o *OswStackMemberPortVO) GetAllMlagDadPorts() []int32`

GetAllMlagDadPorts returns the AllMlagDadPorts field if non-nil, zero value otherwise.

### GetAllMlagDadPortsOk

`func (o *OswStackMemberPortVO) GetAllMlagDadPortsOk() (*[]int32, bool)`

GetAllMlagDadPortsOk returns a tuple with the AllMlagDadPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMlagDadPorts

`func (o *OswStackMemberPortVO) SetAllMlagDadPorts(v []int32)`

SetAllMlagDadPorts sets AllMlagDadPorts field to given value.

### HasAllMlagDadPorts

`func (o *OswStackMemberPortVO) HasAllMlagDadPorts() bool`

HasAllMlagDadPorts returns a boolean if a field has been set.

### GetAllMlagPeerLinkPorts

`func (o *OswStackMemberPortVO) GetAllMlagPeerLinkPorts() []int32`

GetAllMlagPeerLinkPorts returns the AllMlagPeerLinkPorts field if non-nil, zero value otherwise.

### GetAllMlagPeerLinkPortsOk

`func (o *OswStackMemberPortVO) GetAllMlagPeerLinkPortsOk() (*[]int32, bool)`

GetAllMlagPeerLinkPortsOk returns a tuple with the AllMlagPeerLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMlagPeerLinkPorts

`func (o *OswStackMemberPortVO) SetAllMlagPeerLinkPorts(v []int32)`

SetAllMlagPeerLinkPorts sets AllMlagPeerLinkPorts field to given value.

### HasAllMlagPeerLinkPorts

`func (o *OswStackMemberPortVO) HasAllMlagPeerLinkPorts() bool`

HasAllMlagPeerLinkPorts returns a boolean if a field has been set.

### GetBandCtrl

`func (o *OswStackMemberPortVO) GetBandCtrl() OswBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *OswStackMemberPortVO) GetBandCtrlOk() (*OswBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *OswStackMemberPortVO) SetBandCtrl(v OswBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *OswStackMemberPortVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *OswStackMemberPortVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *OswStackMemberPortVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *OswStackMemberPortVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *OswStackMemberPortVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetConfigMlagDad

`func (o *OswStackMemberPortVO) GetConfigMlagDad() bool`

GetConfigMlagDad returns the ConfigMlagDad field if non-nil, zero value otherwise.

### GetConfigMlagDadOk

`func (o *OswStackMemberPortVO) GetConfigMlagDadOk() (*bool, bool)`

GetConfigMlagDadOk returns a tuple with the ConfigMlagDad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagDad

`func (o *OswStackMemberPortVO) SetConfigMlagDad(v bool)`

SetConfigMlagDad sets ConfigMlagDad field to given value.

### HasConfigMlagDad

`func (o *OswStackMemberPortVO) HasConfigMlagDad() bool`

HasConfigMlagDad returns a boolean if a field has been set.

### GetConfigMlagPeerLink

`func (o *OswStackMemberPortVO) GetConfigMlagPeerLink() bool`

GetConfigMlagPeerLink returns the ConfigMlagPeerLink field if non-nil, zero value otherwise.

### GetConfigMlagPeerLinkOk

`func (o *OswStackMemberPortVO) GetConfigMlagPeerLinkOk() (*bool, bool)`

GetConfigMlagPeerLinkOk returns a tuple with the ConfigMlagPeerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagPeerLink

`func (o *OswStackMemberPortVO) SetConfigMlagPeerLink(v bool)`

SetConfigMlagPeerLink sets ConfigMlagPeerLink field to given value.

### HasConfigMlagPeerLink

`func (o *OswStackMemberPortVO) HasConfigMlagPeerLink() bool`

HasConfigMlagPeerLink returns a boolean if a field has been set.

### GetConfigStack

`func (o *OswStackMemberPortVO) GetConfigStack() bool`

GetConfigStack returns the ConfigStack field if non-nil, zero value otherwise.

### GetConfigStackOk

`func (o *OswStackMemberPortVO) GetConfigStackOk() (*bool, bool)`

GetConfigStackOk returns a tuple with the ConfigStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStack

`func (o *OswStackMemberPortVO) SetConfigStack(v bool)`

SetConfigStack sets ConfigStack field to given value.

### HasConfigStack

`func (o *OswStackMemberPortVO) HasConfigStack() bool`

HasConfigStack returns a boolean if a field has been set.

### GetDhcpL2RelaySettings

`func (o *OswStackMemberPortVO) GetDhcpL2RelaySettings() OswPortDhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *OswStackMemberPortVO) GetDhcpL2RelaySettingsOk() (*OswPortDhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *OswStackMemberPortVO) SetDhcpL2RelaySettings(v OswPortDhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *OswStackMemberPortVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDisable

`func (o *OswStackMemberPortVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswStackMemberPortVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswStackMemberPortVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswStackMemberPortVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *OswStackMemberPortVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *OswStackMemberPortVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *OswStackMemberPortVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *OswStackMemberPortVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDot1x

`func (o *OswStackMemberPortVO) GetDot1x() int32`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *OswStackMemberPortVO) GetDot1xOk() (*int32, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *OswStackMemberPortVO) SetDot1x(v int32)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *OswStackMemberPortVO) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetDuplex

`func (o *OswStackMemberPortVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswStackMemberPortVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswStackMemberPortVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswStackMemberPortVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEeeEnable

`func (o *OswStackMemberPortVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *OswStackMemberPortVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *OswStackMemberPortVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *OswStackMemberPortVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetEsEnableAllProfileCanAdd

`func (o *OswStackMemberPortVO) GetEsEnableAllProfileCanAdd() bool`

GetEsEnableAllProfileCanAdd returns the EsEnableAllProfileCanAdd field if non-nil, zero value otherwise.

### GetEsEnableAllProfileCanAddOk

`func (o *OswStackMemberPortVO) GetEsEnableAllProfileCanAddOk() (*bool, bool)`

GetEsEnableAllProfileCanAddOk returns a tuple with the EsEnableAllProfileCanAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsEnableAllProfileCanAdd

`func (o *OswStackMemberPortVO) SetEsEnableAllProfileCanAdd(v bool)`

SetEsEnableAllProfileCanAdd sets EsEnableAllProfileCanAdd field to given value.

### HasEsEnableAllProfileCanAdd

`func (o *OswStackMemberPortVO) HasEsEnableAllProfileCanAdd() bool`

HasEsEnableAllProfileCanAdd returns a boolean if a field has been set.

### GetEsQosSupport

`func (o *OswStackMemberPortVO) GetEsQosSupport() bool`

GetEsQosSupport returns the EsQosSupport field if non-nil, zero value otherwise.

### GetEsQosSupportOk

`func (o *OswStackMemberPortVO) GetEsQosSupportOk() (*bool, bool)`

GetEsQosSupportOk returns a tuple with the EsQosSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsQosSupport

`func (o *OswStackMemberPortVO) SetEsQosSupport(v bool)`

SetEsQosSupport sets EsQosSupport field to given value.

### HasEsQosSupport

`func (o *OswStackMemberPortVO) HasEsQosSupport() bool`

HasEsQosSupport returns a boolean if a field has been set.

### GetExtendModeEnable

`func (o *OswStackMemberPortVO) GetExtendModeEnable() bool`

GetExtendModeEnable returns the ExtendModeEnable field if non-nil, zero value otherwise.

### GetExtendModeEnableOk

`func (o *OswStackMemberPortVO) GetExtendModeEnableOk() (*bool, bool)`

GetExtendModeEnableOk returns a tuple with the ExtendModeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendModeEnable

`func (o *OswStackMemberPortVO) SetExtendModeEnable(v bool)`

SetExtendModeEnable sets ExtendModeEnable field to given value.

### HasExtendModeEnable

`func (o *OswStackMemberPortVO) HasExtendModeEnable() bool`

HasExtendModeEnable returns a boolean if a field has been set.

### GetExtendModeSupport

`func (o *OswStackMemberPortVO) GetExtendModeSupport() bool`

GetExtendModeSupport returns the ExtendModeSupport field if non-nil, zero value otherwise.

### GetExtendModeSupportOk

`func (o *OswStackMemberPortVO) GetExtendModeSupportOk() (*bool, bool)`

GetExtendModeSupportOk returns a tuple with the ExtendModeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendModeSupport

`func (o *OswStackMemberPortVO) SetExtendModeSupport(v bool)`

SetExtendModeSupport sets ExtendModeSupport field to given value.

### HasExtendModeSupport

`func (o *OswStackMemberPortVO) HasExtendModeSupport() bool`

HasExtendModeSupport returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *OswStackMemberPortVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *OswStackMemberPortVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *OswStackMemberPortVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *OswStackMemberPortVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFecCap

`func (o *OswStackMemberPortVO) GetFecCap() []OswFecCapVO`

GetFecCap returns the FecCap field if non-nil, zero value otherwise.

### GetFecCapOk

`func (o *OswStackMemberPortVO) GetFecCapOk() (*[]OswFecCapVO, bool)`

GetFecCapOk returns a tuple with the FecCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecCap

`func (o *OswStackMemberPortVO) SetFecCap(v []OswFecCapVO)`

SetFecCap sets FecCap field to given value.

### HasFecCap

`func (o *OswStackMemberPortVO) HasFecCap() bool`

HasFecCap returns a boolean if a field has been set.

### GetFecLinkPeerApplyEnable

`func (o *OswStackMemberPortVO) GetFecLinkPeerApplyEnable() bool`

GetFecLinkPeerApplyEnable returns the FecLinkPeerApplyEnable field if non-nil, zero value otherwise.

### GetFecLinkPeerApplyEnableOk

`func (o *OswStackMemberPortVO) GetFecLinkPeerApplyEnableOk() (*bool, bool)`

GetFecLinkPeerApplyEnableOk returns a tuple with the FecLinkPeerApplyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecLinkPeerApplyEnable

`func (o *OswStackMemberPortVO) SetFecLinkPeerApplyEnable(v bool)`

SetFecLinkPeerApplyEnable sets FecLinkPeerApplyEnable field to given value.

### HasFecLinkPeerApplyEnable

`func (o *OswStackMemberPortVO) HasFecLinkPeerApplyEnable() bool`

HasFecLinkPeerApplyEnable returns a boolean if a field has been set.

### GetFecMode

`func (o *OswStackMemberPortVO) GetFecMode() int32`

GetFecMode returns the FecMode field if non-nil, zero value otherwise.

### GetFecModeOk

`func (o *OswStackMemberPortVO) GetFecModeOk() (*int32, bool)`

GetFecModeOk returns a tuple with the FecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecMode

`func (o *OswStackMemberPortVO) SetFecMode(v int32)`

SetFecMode sets FecMode field to given value.

### HasFecMode

`func (o *OswStackMemberPortVO) HasFecMode() bool`

HasFecMode returns a boolean if a field has been set.

### GetFecSupport

`func (o *OswStackMemberPortVO) GetFecSupport() bool`

GetFecSupport returns the FecSupport field if non-nil, zero value otherwise.

### GetFecSupportOk

`func (o *OswStackMemberPortVO) GetFecSupportOk() (*bool, bool)`

GetFecSupportOk returns a tuple with the FecSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecSupport

`func (o *OswStackMemberPortVO) SetFecSupport(v bool)`

SetFecSupport sets FecSupport field to given value.

### HasFecSupport

`func (o *OswStackMemberPortVO) HasFecSupport() bool`

HasFecSupport returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *OswStackMemberPortVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *OswStackMemberPortVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *OswStackMemberPortVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *OswStackMemberPortVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetId

`func (o *OswStackMemberPortVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswStackMemberPortVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswStackMemberPortVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswStackMemberPortVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *OswStackMemberPortVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *OswStackMemberPortVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *OswStackMemberPortVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *OswStackMemberPortVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetIgmpSnoopingEnable

`func (o *OswStackMemberPortVO) GetIgmpSnoopingEnable() bool`

GetIgmpSnoopingEnable returns the IgmpSnoopingEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnableOk

`func (o *OswStackMemberPortVO) GetIgmpSnoopingEnableOk() (*bool, bool)`

GetIgmpSnoopingEnableOk returns a tuple with the IgmpSnoopingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnable

`func (o *OswStackMemberPortVO) SetIgmpSnoopingEnable(v bool)`

SetIgmpSnoopingEnable sets IgmpSnoopingEnable field to given value.

### HasIgmpSnoopingEnable

`func (o *OswStackMemberPortVO) HasIgmpSnoopingEnable() bool`

HasIgmpSnoopingEnable returns a boolean if a field has been set.

### GetLagSetting

`func (o *OswStackMemberPortVO) GetLagSetting() OswLagVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *OswStackMemberPortVO) GetLagSettingOk() (*OswLagVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *OswStackMemberPortVO) SetLagSetting(v OswLagVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *OswStackMemberPortVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetLagStatus

`func (o *OswStackMemberPortVO) GetLagStatus() OswLagStatusVO`

GetLagStatus returns the LagStatus field if non-nil, zero value otherwise.

### GetLagStatusOk

`func (o *OswStackMemberPortVO) GetLagStatusOk() (*OswLagStatusVO, bool)`

GetLagStatusOk returns a tuple with the LagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagStatus

`func (o *OswStackMemberPortVO) SetLagStatus(v OswLagStatusVO)`

SetLagStatus sets LagStatus field to given value.

### HasLagStatus

`func (o *OswStackMemberPortVO) HasLagStatus() bool`

HasLagStatus returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswStackMemberPortVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswStackMemberPortVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswStackMemberPortVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswStackMemberPortVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLldpMedEnable

`func (o *OswStackMemberPortVO) GetLldpMedEnable() bool`

GetLldpMedEnable returns the LldpMedEnable field if non-nil, zero value otherwise.

### GetLldpMedEnableOk

`func (o *OswStackMemberPortVO) GetLldpMedEnableOk() (*bool, bool)`

GetLldpMedEnableOk returns a tuple with the LldpMedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMedEnable

`func (o *OswStackMemberPortVO) SetLldpMedEnable(v bool)`

SetLldpMedEnable sets LldpMedEnable field to given value.

### HasLldpMedEnable

`func (o *OswStackMemberPortVO) HasLldpMedEnable() bool`

HasLldpMedEnable returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswStackMemberPortVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswStackMemberPortVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswStackMemberPortVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswStackMemberPortVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswStackMemberPortVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswStackMemberPortVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswStackMemberPortVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswStackMemberPortVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackDetectVlanBasedEnable

`func (o *OswStackMemberPortVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *OswStackMemberPortVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *OswStackMemberPortVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *OswStackMemberPortVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMadUsed

`func (o *OswStackMemberPortVO) GetMadUsed() bool`

GetMadUsed returns the MadUsed field if non-nil, zero value otherwise.

### GetMadUsedOk

`func (o *OswStackMemberPortVO) GetMadUsedOk() (*bool, bool)`

GetMadUsedOk returns a tuple with the MadUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadUsed

`func (o *OswStackMemberPortVO) SetMadUsed(v bool)`

SetMadUsed sets MadUsed field to given value.

### HasMadUsed

`func (o *OswStackMemberPortVO) HasMadUsed() bool`

HasMadUsed returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *OswStackMemberPortVO) GetMaxSpeed() int32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *OswStackMemberPortVO) GetMaxSpeedOk() (*int32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *OswStackMemberPortVO) SetMaxSpeed(v int32)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *OswStackMemberPortVO) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetMirroredLags

`func (o *OswStackMemberPortVO) GetMirroredLags() []MirroredLag`

GetMirroredLags returns the MirroredLags field if non-nil, zero value otherwise.

### GetMirroredLagsOk

`func (o *OswStackMemberPortVO) GetMirroredLagsOk() (*[]MirroredLag, bool)`

GetMirroredLagsOk returns a tuple with the MirroredLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredLags

`func (o *OswStackMemberPortVO) SetMirroredLags(v []MirroredLag)`

SetMirroredLags sets MirroredLags field to given value.

### HasMirroredLags

`func (o *OswStackMemberPortVO) HasMirroredLags() bool`

HasMirroredLags returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *OswStackMemberPortVO) GetMirroredPorts() []MirroredPort`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *OswStackMemberPortVO) GetMirroredPortsOk() (*[]MirroredPort, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *OswStackMemberPortVO) SetMirroredPorts(v []MirroredPort)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *OswStackMemberPortVO) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetMlagPeerAllPortsConfigInfo

`func (o *OswStackMemberPortVO) GetMlagPeerAllPortsConfigInfo() OswMlagPeerAllPortsConfigInfoVO`

GetMlagPeerAllPortsConfigInfo returns the MlagPeerAllPortsConfigInfo field if non-nil, zero value otherwise.

### GetMlagPeerAllPortsConfigInfoOk

`func (o *OswStackMemberPortVO) GetMlagPeerAllPortsConfigInfoOk() (*OswMlagPeerAllPortsConfigInfoVO, bool)`

GetMlagPeerAllPortsConfigInfoOk returns a tuple with the MlagPeerAllPortsConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerAllPortsConfigInfo

`func (o *OswStackMemberPortVO) SetMlagPeerAllPortsConfigInfo(v OswMlagPeerAllPortsConfigInfoVO)`

SetMlagPeerAllPortsConfigInfo sets MlagPeerAllPortsConfigInfo field to given value.

### HasMlagPeerAllPortsConfigInfo

`func (o *OswStackMemberPortVO) HasMlagPeerAllPortsConfigInfo() bool`

HasMlagPeerAllPortsConfigInfo returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *OswStackMemberPortVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *OswStackMemberPortVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *OswStackMemberPortVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *OswStackMemberPortVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *OswStackMemberPortVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackMemberPortVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackMemberPortVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStackMemberPortVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *OswStackMemberPortVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *OswStackMemberPortVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *OswStackMemberPortVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *OswStackMemberPortVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *OswStackMemberPortVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *OswStackMemberPortVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *OswStackMemberPortVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *OswStackMemberPortVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkConflict

`func (o *OswStackMemberPortVO) GetNetworkConflict() bool`

GetNetworkConflict returns the NetworkConflict field if non-nil, zero value otherwise.

### GetNetworkConflictOk

`func (o *OswStackMemberPortVO) GetNetworkConflictOk() (*bool, bool)`

GetNetworkConflictOk returns a tuple with the NetworkConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConflict

`func (o *OswStackMemberPortVO) SetNetworkConflict(v bool)`

SetNetworkConflict sets NetworkConflict field to given value.

### HasNetworkConflict

`func (o *OswStackMemberPortVO) HasNetworkConflict() bool`

HasNetworkConflict returns a boolean if a field has been set.

### GetNetworkMode

`func (o *OswStackMemberPortVO) GetNetworkMode() int32`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *OswStackMemberPortVO) GetNetworkModeOk() (*int32, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *OswStackMemberPortVO) SetNetworkMode(v int32)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *OswStackMemberPortVO) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *OswStackMemberPortVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *OswStackMemberPortVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *OswStackMemberPortVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *OswStackMemberPortVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetOperation

`func (o *OswStackMemberPortVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OswStackMemberPortVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OswStackMemberPortVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *OswStackMemberPortVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPoe

`func (o *OswStackMemberPortVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OswStackMemberPortVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OswStackMemberPortVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OswStackMemberPortVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPoeDisplayType

`func (o *OswStackMemberPortVO) GetPoeDisplayType() int32`

GetPoeDisplayType returns the PoeDisplayType field if non-nil, zero value otherwise.

### GetPoeDisplayTypeOk

`func (o *OswStackMemberPortVO) GetPoeDisplayTypeOk() (*int32, bool)`

GetPoeDisplayTypeOk returns a tuple with the PoeDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisplayType

`func (o *OswStackMemberPortVO) SetPoeDisplayType(v int32)`

SetPoeDisplayType sets PoeDisplayType field to given value.

### HasPoeDisplayType

`func (o *OswStackMemberPortVO) HasPoeDisplayType() bool`

HasPoeDisplayType returns a boolean if a field has been set.

### GetPort

`func (o *OswStackMemberPortVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswStackMemberPortVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswStackMemberPortVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswStackMemberPortVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortAlertEnable

`func (o *OswStackMemberPortVO) GetPortAlertEnable() bool`

GetPortAlertEnable returns the PortAlertEnable field if non-nil, zero value otherwise.

### GetPortAlertEnableOk

`func (o *OswStackMemberPortVO) GetPortAlertEnableOk() (*bool, bool)`

GetPortAlertEnableOk returns a tuple with the PortAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAlertEnable

`func (o *OswStackMemberPortVO) SetPortAlertEnable(v bool)`

SetPortAlertEnable sets PortAlertEnable field to given value.

### HasPortAlertEnable

`func (o *OswStackMemberPortVO) HasPortAlertEnable() bool`

HasPortAlertEnable returns a boolean if a field has been set.

### GetPortCap

`func (o *OswStackMemberPortVO) GetPortCap() []OswLinkCapVO`

GetPortCap returns the PortCap field if non-nil, zero value otherwise.

### GetPortCapOk

`func (o *OswStackMemberPortVO) GetPortCapOk() (*[]OswLinkCapVO, bool)`

GetPortCapOk returns a tuple with the PortCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCap

`func (o *OswStackMemberPortVO) SetPortCap(v []OswLinkCapVO)`

SetPortCap sets PortCap field to given value.

### HasPortCap

`func (o *OswStackMemberPortVO) HasPortCap() bool`

HasPortCap returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *OswStackMemberPortVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *OswStackMemberPortVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *OswStackMemberPortVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *OswStackMemberPortVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetPortSpeedCap

`func (o *OswStackMemberPortVO) GetPortSpeedCap() []OswPortSpeedCapVO`

GetPortSpeedCap returns the PortSpeedCap field if non-nil, zero value otherwise.

### GetPortSpeedCapOk

`func (o *OswStackMemberPortVO) GetPortSpeedCapOk() (*[]OswPortSpeedCapVO, bool)`

GetPortSpeedCapOk returns a tuple with the PortSpeedCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeedCap

`func (o *OswStackMemberPortVO) SetPortSpeedCap(v []OswPortSpeedCapVO)`

SetPortSpeedCap sets PortSpeedCap field to given value.

### HasPortSpeedCap

`func (o *OswStackMemberPortVO) HasPortSpeedCap() bool`

HasPortSpeedCap returns a boolean if a field has been set.

### GetPortStatus

`func (o *OswStackMemberPortVO) GetPortStatus() OswPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *OswStackMemberPortVO) GetPortStatusOk() (*OswPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *OswStackMemberPortVO) SetPortStatus(v OswPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *OswStackMemberPortVO) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetProfileId

`func (o *OswStackMemberPortVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OswStackMemberPortVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OswStackMemberPortVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *OswStackMemberPortVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *OswStackMemberPortVO) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *OswStackMemberPortVO) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *OswStackMemberPortVO) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *OswStackMemberPortVO) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *OswStackMemberPortVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *OswStackMemberPortVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *OswStackMemberPortVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *OswStackMemberPortVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetQosQueueEnable

`func (o *OswStackMemberPortVO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *OswStackMemberPortVO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *OswStackMemberPortVO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *OswStackMemberPortVO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQosSupport

`func (o *OswStackMemberPortVO) GetQosSupport() bool`

GetQosSupport returns the QosSupport field if non-nil, zero value otherwise.

### GetQosSupportOk

`func (o *OswStackMemberPortVO) GetQosSupportOk() (*bool, bool)`

GetQosSupportOk returns a tuple with the QosSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSupport

`func (o *OswStackMemberPortVO) SetQosSupport(v bool)`

SetQosSupport sets QosSupport field to given value.

### HasQosSupport

`func (o *OswStackMemberPortVO) HasQosSupport() bool`

HasQosSupport returns a boolean if a field has been set.

### GetQueueId

`func (o *OswStackMemberPortVO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *OswStackMemberPortVO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *OswStackMemberPortVO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *OswStackMemberPortVO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetResource

`func (o *OswStackMemberPortVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswStackMemberPortVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswStackMemberPortVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswStackMemberPortVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSpanningTreeEnable

`func (o *OswStackMemberPortVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *OswStackMemberPortVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *OswStackMemberPortVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.

### HasSpanningTreeEnable

`func (o *OswStackMemberPortVO) HasSpanningTreeEnable() bool`

HasSpanningTreeEnable returns a boolean if a field has been set.

### GetSpanningTreeSetting

`func (o *OswStackMemberPortVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *OswStackMemberPortVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *OswStackMemberPortVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *OswStackMemberPortVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetSpeed

`func (o *OswStackMemberPortVO) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OswStackMemberPortVO) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OswStackMemberPortVO) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *OswStackMemberPortVO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStackId

`func (o *OswStackMemberPortVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswStackMemberPortVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswStackMemberPortVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswStackMemberPortVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackPortsGroupIndex

`func (o *OswStackMemberPortVO) GetStackPortsGroupIndex() int32`

GetStackPortsGroupIndex returns the StackPortsGroupIndex field if non-nil, zero value otherwise.

### GetStackPortsGroupIndexOk

`func (o *OswStackMemberPortVO) GetStackPortsGroupIndexOk() (*int32, bool)`

GetStackPortsGroupIndexOk returns a tuple with the StackPortsGroupIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortsGroupIndex

`func (o *OswStackMemberPortVO) SetStackPortsGroupIndex(v int32)`

SetStackPortsGroupIndex sets StackPortsGroupIndex field to given value.

### HasStackPortsGroupIndex

`func (o *OswStackMemberPortVO) HasStackPortsGroupIndex() bool`

HasStackPortsGroupIndex returns a boolean if a field has been set.

### GetStackSetting

`func (o *OswStackMemberPortVO) GetStackSetting() OswPortStackSettingVO`

GetStackSetting returns the StackSetting field if non-nil, zero value otherwise.

### GetStackSettingOk

`func (o *OswStackMemberPortVO) GetStackSettingOk() (*OswPortStackSettingVO, bool)`

GetStackSettingOk returns a tuple with the StackSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackSetting

`func (o *OswStackMemberPortVO) SetStackSetting(v OswPortStackSettingVO)`

SetStackSetting sets StackSetting field to given value.

### HasStackSetting

`func (o *OswStackMemberPortVO) HasStackSetting() bool`

HasStackSetting returns a boolean if a field has been set.

### GetStackSupport

`func (o *OswStackMemberPortVO) GetStackSupport() bool`

GetStackSupport returns the StackSupport field if non-nil, zero value otherwise.

### GetStackSupportOk

`func (o *OswStackMemberPortVO) GetStackSupportOk() (*bool, bool)`

GetStackSupportOk returns a tuple with the StackSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackSupport

`func (o *OswStackMemberPortVO) SetStackSupport(v bool)`

SetStackSupport sets StackSupport field to given value.

### HasStackSupport

`func (o *OswStackMemberPortVO) HasStackSupport() bool`

HasStackSupport returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswStackMemberPortVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswStackMemberPortVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswStackMemberPortVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswStackMemberPortVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetStormCtrl

`func (o *OswStackMemberPortVO) GetStormCtrl() OswStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *OswStackMemberPortVO) GetStormCtrlOk() (*OswStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *OswStackMemberPortVO) SetStormCtrl(v OswStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *OswStackMemberPortVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetSupportLocate

`func (o *OswStackMemberPortVO) GetSupportLocate() bool`

GetSupportLocate returns the SupportLocate field if non-nil, zero value otherwise.

### GetSupportLocateOk

`func (o *OswStackMemberPortVO) GetSupportLocateOk() (*bool, bool)`

GetSupportLocateOk returns a tuple with the SupportLocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocate

`func (o *OswStackMemberPortVO) SetSupportLocate(v bool)`

SetSupportLocate sets SupportLocate field to given value.

### HasSupportLocate

`func (o *OswStackMemberPortVO) HasSupportLocate() bool`

HasSupportLocate returns a boolean if a field has been set.

### GetSupportPoe

`func (o *OswStackMemberPortVO) GetSupportPoe() bool`

GetSupportPoe returns the SupportPoe field if non-nil, zero value otherwise.

### GetSupportPoeOk

`func (o *OswStackMemberPortVO) GetSupportPoeOk() (*bool, bool)`

GetSupportPoeOk returns a tuple with the SupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPoe

`func (o *OswStackMemberPortVO) SetSupportPoe(v bool)`

SetSupportPoe sets SupportPoe field to given value.

### HasSupportPoe

`func (o *OswStackMemberPortVO) HasSupportPoe() bool`

HasSupportPoe returns a boolean if a field has been set.

### GetSwitchId

`func (o *OswStackMemberPortVO) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *OswStackMemberPortVO) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *OswStackMemberPortVO) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *OswStackMemberPortVO) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetSwitchMac

`func (o *OswStackMemberPortVO) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *OswStackMemberPortVO) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *OswStackMemberPortVO) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.

### HasSwitchMac

`func (o *OswStackMemberPortVO) HasSwitchMac() bool`

HasSwitchMac returns a boolean if a field has been set.

### GetTagBridgeVlanMap

`func (o *OswStackMemberPortVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *OswStackMemberPortVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *OswStackMemberPortVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *OswStackMemberPortVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagIds

`func (o *OswStackMemberPortVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswStackMemberPortVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswStackMemberPortVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswStackMemberPortVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagName

`func (o *OswStackMemberPortVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *OswStackMemberPortVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *OswStackMemberPortVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *OswStackMemberPortVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *OswStackMemberPortVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *OswStackMemberPortVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *OswStackMemberPortVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *OswStackMemberPortVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTopoNotifyEnable

`func (o *OswStackMemberPortVO) GetTopoNotifyEnable() bool`

GetTopoNotifyEnable returns the TopoNotifyEnable field if non-nil, zero value otherwise.

### GetTopoNotifyEnableOk

`func (o *OswStackMemberPortVO) GetTopoNotifyEnableOk() (*bool, bool)`

GetTopoNotifyEnableOk returns a tuple with the TopoNotifyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopoNotifyEnable

`func (o *OswStackMemberPortVO) SetTopoNotifyEnable(v bool)`

SetTopoNotifyEnable sets TopoNotifyEnable field to given value.

### HasTopoNotifyEnable

`func (o *OswStackMemberPortVO) HasTopoNotifyEnable() bool`

HasTopoNotifyEnable returns a boolean if a field has been set.

### GetTrustMode

`func (o *OswStackMemberPortVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *OswStackMemberPortVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *OswStackMemberPortVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *OswStackMemberPortVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetType

`func (o *OswStackMemberPortVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswStackMemberPortVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswStackMemberPortVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswStackMemberPortVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUntagBridgeVlanMap

`func (o *OswStackMemberPortVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *OswStackMemberPortVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *OswStackMemberPortVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *OswStackMemberPortVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *OswStackMemberPortVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *OswStackMemberPortVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *OswStackMemberPortVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *OswStackMemberPortVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *OswStackMemberPortVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *OswStackMemberPortVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *OswStackMemberPortVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *OswStackMemberPortVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceDscp

`func (o *OswStackMemberPortVO) GetVoiceDscp() int32`

GetVoiceDscp returns the VoiceDscp field if non-nil, zero value otherwise.

### GetVoiceDscpOk

`func (o *OswStackMemberPortVO) GetVoiceDscpOk() (*int32, bool)`

GetVoiceDscpOk returns a tuple with the VoiceDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscp

`func (o *OswStackMemberPortVO) SetVoiceDscp(v int32)`

SetVoiceDscp sets VoiceDscp field to given value.

### HasVoiceDscp

`func (o *OswStackMemberPortVO) HasVoiceDscp() bool`

HasVoiceDscp returns a boolean if a field has been set.

### GetVoiceDscpEnable

`func (o *OswStackMemberPortVO) GetVoiceDscpEnable() bool`

GetVoiceDscpEnable returns the VoiceDscpEnable field if non-nil, zero value otherwise.

### GetVoiceDscpEnableOk

`func (o *OswStackMemberPortVO) GetVoiceDscpEnableOk() (*bool, bool)`

GetVoiceDscpEnableOk returns a tuple with the VoiceDscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpEnable

`func (o *OswStackMemberPortVO) SetVoiceDscpEnable(v bool)`

SetVoiceDscpEnable sets VoiceDscpEnable field to given value.

### HasVoiceDscpEnable

`func (o *OswStackMemberPortVO) HasVoiceDscpEnable() bool`

HasVoiceDscpEnable returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *OswStackMemberPortVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *OswStackMemberPortVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *OswStackMemberPortVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *OswStackMemberPortVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *OswStackMemberPortVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *OswStackMemberPortVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *OswStackMemberPortVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *OswStackMemberPortVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


