# OswPortSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpDetectEnable** | Pointer to **bool** | Indicates whether arp detect is enabled | [optional] 
**BandCtrl** | Pointer to [**OswBandCtrlVO**](OswBandCtrlVO.md) |  | [optional] 
**BandWidthCtrlType** | Pointer to **int32** | BandWidthCtrlType should be a value as follows: 0: Off; 1: Rate Limit; 2: Storming Control | [optional] 
**DhcpL2RelaySettings** | Pointer to [**OswPortDhcpL2RelayVO**](OswPortDhcpL2RelayVO.md) |  | [optional] 
**DhcpSnoopEnable** | Pointer to **bool** | Indicates whether dhcp snoop is enabled | [optional] 
**Disable** | Pointer to **bool** | Indicates whether to disable | [optional] 
**Dot1pPriority** | Pointer to **int32** | Dot1p Priority | [optional] 
**Dot1x** | Pointer to **int32** | Dot1x should be a value as follows: 0: Force unauthorized; 1: Force authorized; 2: Auto | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**EeeEnable** | Pointer to **bool** | Indicates whether EEE is enabled | [optional] 
**FastLeaveEnable** | Pointer to **bool** | Indicates whether igmpSnooping fastLeave is enabled | [optional] 
**FecLinkPeerApplyEnable** | Pointer to **bool** | Indicates whether the FEC Mode is synchronously applied to the link peer port. | [optional] 
**FecMode** | Pointer to **int32** | FEC mode should be a value as follows: 1: Off; 2: RS528; 3: RS544; 4: Auto; 5: Base-R | [optional] 
**FlowControlEnable** | Pointer to **bool** | Indicates whether flow control is enabled | [optional] 
**IgmpFastLeaveEnable** | Pointer to **bool** | Indicates whether igmp fast leave is enabled | [optional] 
**IgmpSnoopingEnable** | Pointer to **bool** | Indicates whether IGMP Snooping is enabled | [optional] 
**Impbs** | Pointer to [**[]ImpbVO**](ImpbVO.md) | Impbs | [optional] 
**LagSetting** | Pointer to [**OswLagBasicVO**](OswLagBasicVO.md) |  | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed should be a value as follows: 0: auto; 1: 10M; 2: 100M; 3: 1000M; 4: 2.5G; 5: 10G | [optional] 
**LldpMedEnable** | Pointer to **bool** | Indicates whether LLDP-MED is enabled | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | Indicates whether loopbackDetect port based is enabled | [optional] 
**LoopbackDetectVlanBasedEnable** | Pointer to **bool** | Indicates whether loopbackDetect vlan based is enabled | [optional] 
**MirroredLags** | Pointer to **[]int32** | Mirrored Lags | [optional] 
**MirroredPorts** | Pointer to **[]int32** | Mirrored Ports | [optional] 
**MldFastLeaveEnable** | Pointer to **bool** | Indicates whether mld fast leave is enabled | [optional] 
**Name** | Pointer to **string** | Port or Lag Name | [optional] 
**NativeBridgeVlan** | Pointer to **int32** | Native Network Bridge Vlan. | [optional] 
**NativeNetworkId** | Pointer to **string** | Native Network ID, Native Network cannot be selected from Tagged Networks or Untagged Networks. | [optional] 
**NetworkTagsSetting** | Pointer to **int32** | Network Tags Setting should be a value as follows: 0: Allow All; 1: Block All; 2: Custom | [optional] 
**Operation** | Pointer to **string** | Operation should be a value as follows: \&quot;switching\&quot; or \&quot;mirroring\&quot; or \&quot;aggregating\&quot; | [optional] 
**Poe** | Pointer to **int32** | Poe should be a value as follows: 0: Off; 1: 802.3at/af | [optional] 
**PortAlertEnable** | Pointer to **bool** | Indicates whether port alert is enabled | [optional] 
**PortIsolationEnable** | Pointer to **bool** | Indicates whether port isolation is enabled | [optional] 
**ProfileId** | Pointer to **string** | Profile ID | [optional] 
**ProfileOverrideEnable** | Pointer to **bool** | Indicates whether to enable Profile Override | [optional] 
**QosQueueEnable** | Pointer to **bool** | Indicates whether the ES device port has enabled the Qos scheduling queue configuration | [optional] 
**QueueId** | Pointer to **int32** | ES Qos scheduling queue ID | [optional] 
**SpanningTreeEnable** | Pointer to **bool** | Indicates whether SpanningTree is enabled | [optional] 
**SpanningTreeSetting** | Pointer to [**SpanningTreeSettingVO**](SpanningTreeSettingVO.md) |  | [optional] 
**StormCtrl** | Pointer to [**OswStormCtrlVO**](OswStormCtrlVO.md) |  | [optional] 
**TagBridgeVlanMap** | Pointer to **map[string][]int32** | Tag Network Bridge Vlan Map | [optional] 
**TagIds** | Pointer to **[]string** | Tag ID List | [optional] 
**TagNetworkIds** | Pointer to **[]string** | Tag Network IDs | [optional] 
**TopoNotifyEnable** | Pointer to **bool** | Indicates whether Topology Notify is enabled | [optional] 
**TrustMode** | Pointer to **int32** | TrustMode should be a value as follows: 0: Untrusted; 1: Trust 802.1p; 2: Trust DSCP | [optional] 
**UntagBridgeVlanMap** | Pointer to **map[string][]int32** | Untag Network Bridge Vlan Map | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag Network IDs | [optional] 
**VoiceBridgeVlan** | Pointer to **int32** | Voice Network Bridge Vlan | [optional] 
**VoiceDscp** | Pointer to **int32** | Voice DSCP | [optional] 
**VoiceDscpEnable** | Pointer to **bool** | Indicates whether voice DSCP is enabled | [optional] 
**VoiceNetworkEnable** | Pointer to **bool** | Indicates whether voice network is enabled | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewOswPortSettingVO

`func NewOswPortSettingVO() *OswPortSettingVO`

NewOswPortSettingVO instantiates a new OswPortSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortSettingVOWithDefaults

`func NewOswPortSettingVOWithDefaults() *OswPortSettingVO`

NewOswPortSettingVOWithDefaults instantiates a new OswPortSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpDetectEnable

`func (o *OswPortSettingVO) GetArpDetectEnable() bool`

GetArpDetectEnable returns the ArpDetectEnable field if non-nil, zero value otherwise.

### GetArpDetectEnableOk

`func (o *OswPortSettingVO) GetArpDetectEnableOk() (*bool, bool)`

GetArpDetectEnableOk returns a tuple with the ArpDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpDetectEnable

`func (o *OswPortSettingVO) SetArpDetectEnable(v bool)`

SetArpDetectEnable sets ArpDetectEnable field to given value.

### HasArpDetectEnable

`func (o *OswPortSettingVO) HasArpDetectEnable() bool`

HasArpDetectEnable returns a boolean if a field has been set.

### GetBandCtrl

`func (o *OswPortSettingVO) GetBandCtrl() OswBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *OswPortSettingVO) GetBandCtrlOk() (*OswBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *OswPortSettingVO) SetBandCtrl(v OswBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *OswPortSettingVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *OswPortSettingVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *OswPortSettingVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *OswPortSettingVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *OswPortSettingVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetDhcpL2RelaySettings

`func (o *OswPortSettingVO) GetDhcpL2RelaySettings() OswPortDhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *OswPortSettingVO) GetDhcpL2RelaySettingsOk() (*OswPortDhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *OswPortSettingVO) SetDhcpL2RelaySettings(v OswPortDhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *OswPortSettingVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDhcpSnoopEnable

`func (o *OswPortSettingVO) GetDhcpSnoopEnable() bool`

GetDhcpSnoopEnable returns the DhcpSnoopEnable field if non-nil, zero value otherwise.

### GetDhcpSnoopEnableOk

`func (o *OswPortSettingVO) GetDhcpSnoopEnableOk() (*bool, bool)`

GetDhcpSnoopEnableOk returns a tuple with the DhcpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnoopEnable

`func (o *OswPortSettingVO) SetDhcpSnoopEnable(v bool)`

SetDhcpSnoopEnable sets DhcpSnoopEnable field to given value.

### HasDhcpSnoopEnable

`func (o *OswPortSettingVO) HasDhcpSnoopEnable() bool`

HasDhcpSnoopEnable returns a boolean if a field has been set.

### GetDisable

`func (o *OswPortSettingVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswPortSettingVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswPortSettingVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswPortSettingVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *OswPortSettingVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *OswPortSettingVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *OswPortSettingVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *OswPortSettingVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDot1x

`func (o *OswPortSettingVO) GetDot1x() int32`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *OswPortSettingVO) GetDot1xOk() (*int32, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *OswPortSettingVO) SetDot1x(v int32)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *OswPortSettingVO) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetDuplex

`func (o *OswPortSettingVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswPortSettingVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswPortSettingVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswPortSettingVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEeeEnable

`func (o *OswPortSettingVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *OswPortSettingVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *OswPortSettingVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *OswPortSettingVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *OswPortSettingVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *OswPortSettingVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *OswPortSettingVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *OswPortSettingVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFecLinkPeerApplyEnable

`func (o *OswPortSettingVO) GetFecLinkPeerApplyEnable() bool`

GetFecLinkPeerApplyEnable returns the FecLinkPeerApplyEnable field if non-nil, zero value otherwise.

### GetFecLinkPeerApplyEnableOk

`func (o *OswPortSettingVO) GetFecLinkPeerApplyEnableOk() (*bool, bool)`

GetFecLinkPeerApplyEnableOk returns a tuple with the FecLinkPeerApplyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecLinkPeerApplyEnable

`func (o *OswPortSettingVO) SetFecLinkPeerApplyEnable(v bool)`

SetFecLinkPeerApplyEnable sets FecLinkPeerApplyEnable field to given value.

### HasFecLinkPeerApplyEnable

`func (o *OswPortSettingVO) HasFecLinkPeerApplyEnable() bool`

HasFecLinkPeerApplyEnable returns a boolean if a field has been set.

### GetFecMode

`func (o *OswPortSettingVO) GetFecMode() int32`

GetFecMode returns the FecMode field if non-nil, zero value otherwise.

### GetFecModeOk

`func (o *OswPortSettingVO) GetFecModeOk() (*int32, bool)`

GetFecModeOk returns a tuple with the FecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecMode

`func (o *OswPortSettingVO) SetFecMode(v int32)`

SetFecMode sets FecMode field to given value.

### HasFecMode

`func (o *OswPortSettingVO) HasFecMode() bool`

HasFecMode returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *OswPortSettingVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *OswPortSettingVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *OswPortSettingVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *OswPortSettingVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *OswPortSettingVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *OswPortSettingVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *OswPortSettingVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *OswPortSettingVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetIgmpSnoopingEnable

`func (o *OswPortSettingVO) GetIgmpSnoopingEnable() bool`

GetIgmpSnoopingEnable returns the IgmpSnoopingEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnableOk

`func (o *OswPortSettingVO) GetIgmpSnoopingEnableOk() (*bool, bool)`

GetIgmpSnoopingEnableOk returns a tuple with the IgmpSnoopingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnable

`func (o *OswPortSettingVO) SetIgmpSnoopingEnable(v bool)`

SetIgmpSnoopingEnable sets IgmpSnoopingEnable field to given value.

### HasIgmpSnoopingEnable

`func (o *OswPortSettingVO) HasIgmpSnoopingEnable() bool`

HasIgmpSnoopingEnable returns a boolean if a field has been set.

### GetImpbs

`func (o *OswPortSettingVO) GetImpbs() []ImpbVO`

GetImpbs returns the Impbs field if non-nil, zero value otherwise.

### GetImpbsOk

`func (o *OswPortSettingVO) GetImpbsOk() (*[]ImpbVO, bool)`

GetImpbsOk returns a tuple with the Impbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpbs

`func (o *OswPortSettingVO) SetImpbs(v []ImpbVO)`

SetImpbs sets Impbs field to given value.

### HasImpbs

`func (o *OswPortSettingVO) HasImpbs() bool`

HasImpbs returns a boolean if a field has been set.

### GetLagSetting

`func (o *OswPortSettingVO) GetLagSetting() OswLagBasicVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *OswPortSettingVO) GetLagSettingOk() (*OswLagBasicVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *OswPortSettingVO) SetLagSetting(v OswLagBasicVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *OswPortSettingVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswPortSettingVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswPortSettingVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswPortSettingVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswPortSettingVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLldpMedEnable

`func (o *OswPortSettingVO) GetLldpMedEnable() bool`

GetLldpMedEnable returns the LldpMedEnable field if non-nil, zero value otherwise.

### GetLldpMedEnableOk

`func (o *OswPortSettingVO) GetLldpMedEnableOk() (*bool, bool)`

GetLldpMedEnableOk returns a tuple with the LldpMedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMedEnable

`func (o *OswPortSettingVO) SetLldpMedEnable(v bool)`

SetLldpMedEnable sets LldpMedEnable field to given value.

### HasLldpMedEnable

`func (o *OswPortSettingVO) HasLldpMedEnable() bool`

HasLldpMedEnable returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswPortSettingVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswPortSettingVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswPortSettingVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswPortSettingVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackDetectVlanBasedEnable

`func (o *OswPortSettingVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *OswPortSettingVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *OswPortSettingVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *OswPortSettingVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMirroredLags

`func (o *OswPortSettingVO) GetMirroredLags() []int32`

GetMirroredLags returns the MirroredLags field if non-nil, zero value otherwise.

### GetMirroredLagsOk

`func (o *OswPortSettingVO) GetMirroredLagsOk() (*[]int32, bool)`

GetMirroredLagsOk returns a tuple with the MirroredLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredLags

`func (o *OswPortSettingVO) SetMirroredLags(v []int32)`

SetMirroredLags sets MirroredLags field to given value.

### HasMirroredLags

`func (o *OswPortSettingVO) HasMirroredLags() bool`

HasMirroredLags returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *OswPortSettingVO) GetMirroredPorts() []int32`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *OswPortSettingVO) GetMirroredPortsOk() (*[]int32, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *OswPortSettingVO) SetMirroredPorts(v []int32)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *OswPortSettingVO) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *OswPortSettingVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *OswPortSettingVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *OswPortSettingVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *OswPortSettingVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *OswPortSettingVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswPortSettingVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswPortSettingVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswPortSettingVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *OswPortSettingVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *OswPortSettingVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *OswPortSettingVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *OswPortSettingVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *OswPortSettingVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *OswPortSettingVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *OswPortSettingVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *OswPortSettingVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *OswPortSettingVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *OswPortSettingVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *OswPortSettingVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *OswPortSettingVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetOperation

`func (o *OswPortSettingVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OswPortSettingVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OswPortSettingVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *OswPortSettingVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPoe

`func (o *OswPortSettingVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OswPortSettingVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OswPortSettingVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OswPortSettingVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPortAlertEnable

`func (o *OswPortSettingVO) GetPortAlertEnable() bool`

GetPortAlertEnable returns the PortAlertEnable field if non-nil, zero value otherwise.

### GetPortAlertEnableOk

`func (o *OswPortSettingVO) GetPortAlertEnableOk() (*bool, bool)`

GetPortAlertEnableOk returns a tuple with the PortAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAlertEnable

`func (o *OswPortSettingVO) SetPortAlertEnable(v bool)`

SetPortAlertEnable sets PortAlertEnable field to given value.

### HasPortAlertEnable

`func (o *OswPortSettingVO) HasPortAlertEnable() bool`

HasPortAlertEnable returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *OswPortSettingVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *OswPortSettingVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *OswPortSettingVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *OswPortSettingVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetProfileId

`func (o *OswPortSettingVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OswPortSettingVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OswPortSettingVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *OswPortSettingVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *OswPortSettingVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *OswPortSettingVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *OswPortSettingVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *OswPortSettingVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetQosQueueEnable

`func (o *OswPortSettingVO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *OswPortSettingVO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *OswPortSettingVO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *OswPortSettingVO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQueueId

`func (o *OswPortSettingVO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *OswPortSettingVO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *OswPortSettingVO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *OswPortSettingVO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetSpanningTreeEnable

`func (o *OswPortSettingVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *OswPortSettingVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *OswPortSettingVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.

### HasSpanningTreeEnable

`func (o *OswPortSettingVO) HasSpanningTreeEnable() bool`

HasSpanningTreeEnable returns a boolean if a field has been set.

### GetSpanningTreeSetting

`func (o *OswPortSettingVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *OswPortSettingVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *OswPortSettingVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *OswPortSettingVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetStormCtrl

`func (o *OswPortSettingVO) GetStormCtrl() OswStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *OswPortSettingVO) GetStormCtrlOk() (*OswStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *OswPortSettingVO) SetStormCtrl(v OswStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *OswPortSettingVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetTagBridgeVlanMap

`func (o *OswPortSettingVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *OswPortSettingVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *OswPortSettingVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *OswPortSettingVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagIds

`func (o *OswPortSettingVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswPortSettingVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswPortSettingVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswPortSettingVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *OswPortSettingVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *OswPortSettingVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *OswPortSettingVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *OswPortSettingVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTopoNotifyEnable

`func (o *OswPortSettingVO) GetTopoNotifyEnable() bool`

GetTopoNotifyEnable returns the TopoNotifyEnable field if non-nil, zero value otherwise.

### GetTopoNotifyEnableOk

`func (o *OswPortSettingVO) GetTopoNotifyEnableOk() (*bool, bool)`

GetTopoNotifyEnableOk returns a tuple with the TopoNotifyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopoNotifyEnable

`func (o *OswPortSettingVO) SetTopoNotifyEnable(v bool)`

SetTopoNotifyEnable sets TopoNotifyEnable field to given value.

### HasTopoNotifyEnable

`func (o *OswPortSettingVO) HasTopoNotifyEnable() bool`

HasTopoNotifyEnable returns a boolean if a field has been set.

### GetTrustMode

`func (o *OswPortSettingVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *OswPortSettingVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *OswPortSettingVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *OswPortSettingVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetUntagBridgeVlanMap

`func (o *OswPortSettingVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *OswPortSettingVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *OswPortSettingVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *OswPortSettingVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *OswPortSettingVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *OswPortSettingVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *OswPortSettingVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *OswPortSettingVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *OswPortSettingVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *OswPortSettingVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *OswPortSettingVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *OswPortSettingVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceDscp

`func (o *OswPortSettingVO) GetVoiceDscp() int32`

GetVoiceDscp returns the VoiceDscp field if non-nil, zero value otherwise.

### GetVoiceDscpOk

`func (o *OswPortSettingVO) GetVoiceDscpOk() (*int32, bool)`

GetVoiceDscpOk returns a tuple with the VoiceDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscp

`func (o *OswPortSettingVO) SetVoiceDscp(v int32)`

SetVoiceDscp sets VoiceDscp field to given value.

### HasVoiceDscp

`func (o *OswPortSettingVO) HasVoiceDscp() bool`

HasVoiceDscp returns a boolean if a field has been set.

### GetVoiceDscpEnable

`func (o *OswPortSettingVO) GetVoiceDscpEnable() bool`

GetVoiceDscpEnable returns the VoiceDscpEnable field if non-nil, zero value otherwise.

### GetVoiceDscpEnableOk

`func (o *OswPortSettingVO) GetVoiceDscpEnableOk() (*bool, bool)`

GetVoiceDscpEnableOk returns a tuple with the VoiceDscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpEnable

`func (o *OswPortSettingVO) SetVoiceDscpEnable(v bool)`

SetVoiceDscpEnable sets VoiceDscpEnable field to given value.

### HasVoiceDscpEnable

`func (o *OswPortSettingVO) HasVoiceDscpEnable() bool`

HasVoiceDscpEnable returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *OswPortSettingVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *OswPortSettingVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *OswPortSettingVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *OswPortSettingVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *OswPortSettingVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *OswPortSettingVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *OswPortSettingVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *OswPortSettingVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


