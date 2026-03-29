# BatchStackPortSettingVO

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
**Mac** | Pointer to **string** | Mac address of the port. This parameter can be null when batch modify stack ports, but not null when modify a single stack port | [optional] 
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
**StackSetting** | Pointer to [**OswPortStackSettingVO**](OswPortStackSettingVO.md) |  | [optional] 
**StormCtrl** | Pointer to [**OswStormCtrlVO**](OswStormCtrlVO.md) |  | [optional] 
**TagBridgeVlanMap** | Pointer to **map[string][]int32** | Tag Network Bridge Vlan Map | [optional] 
**TagIds** | Pointer to **[]string** | Tag ID List | [optional] 
**TagNetworkIds** | Pointer to **[]string** | Tag Network IDs | [optional] 
**TopoNotifyEnable** | Pointer to **bool** | Indicates whether Topology Notify is enabled | [optional] 
**TrustMode** | Pointer to **int32** | TrustMode should be a value as follows: 0: Untrusted; 1: Trust 802.1p; 2: Trust DSCP | [optional] 
**UnitList** | [**[]OswStackUnitVO**](OswStackUnitVO.md) |  | 
**UntagBridgeVlanMap** | Pointer to **map[string][]int32** | Untag Network Bridge Vlan Map | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag Network IDs | [optional] 
**VoiceBridgeVlan** | Pointer to **int32** | Voice Network Bridge Vlan | [optional] 
**VoiceDscp** | Pointer to **int32** | Voice DSCP | [optional] 
**VoiceDscpEnable** | Pointer to **bool** | Indicates whether voice DSCP is enabled | [optional] 
**VoiceNetworkEnable** | Pointer to **bool** | Indicates whether voice network is enabled | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewBatchStackPortSettingVO

`func NewBatchStackPortSettingVO(unitList []OswStackUnitVO, ) *BatchStackPortSettingVO`

NewBatchStackPortSettingVO instantiates a new BatchStackPortSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchStackPortSettingVOWithDefaults

`func NewBatchStackPortSettingVOWithDefaults() *BatchStackPortSettingVO`

NewBatchStackPortSettingVOWithDefaults instantiates a new BatchStackPortSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpDetectEnable

`func (o *BatchStackPortSettingVO) GetArpDetectEnable() bool`

GetArpDetectEnable returns the ArpDetectEnable field if non-nil, zero value otherwise.

### GetArpDetectEnableOk

`func (o *BatchStackPortSettingVO) GetArpDetectEnableOk() (*bool, bool)`

GetArpDetectEnableOk returns a tuple with the ArpDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpDetectEnable

`func (o *BatchStackPortSettingVO) SetArpDetectEnable(v bool)`

SetArpDetectEnable sets ArpDetectEnable field to given value.

### HasArpDetectEnable

`func (o *BatchStackPortSettingVO) HasArpDetectEnable() bool`

HasArpDetectEnable returns a boolean if a field has been set.

### GetBandCtrl

`func (o *BatchStackPortSettingVO) GetBandCtrl() OswBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *BatchStackPortSettingVO) GetBandCtrlOk() (*OswBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *BatchStackPortSettingVO) SetBandCtrl(v OswBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *BatchStackPortSettingVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *BatchStackPortSettingVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *BatchStackPortSettingVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *BatchStackPortSettingVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *BatchStackPortSettingVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetDhcpL2RelaySettings

`func (o *BatchStackPortSettingVO) GetDhcpL2RelaySettings() OswPortDhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *BatchStackPortSettingVO) GetDhcpL2RelaySettingsOk() (*OswPortDhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *BatchStackPortSettingVO) SetDhcpL2RelaySettings(v OswPortDhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *BatchStackPortSettingVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDhcpSnoopEnable

`func (o *BatchStackPortSettingVO) GetDhcpSnoopEnable() bool`

GetDhcpSnoopEnable returns the DhcpSnoopEnable field if non-nil, zero value otherwise.

### GetDhcpSnoopEnableOk

`func (o *BatchStackPortSettingVO) GetDhcpSnoopEnableOk() (*bool, bool)`

GetDhcpSnoopEnableOk returns a tuple with the DhcpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnoopEnable

`func (o *BatchStackPortSettingVO) SetDhcpSnoopEnable(v bool)`

SetDhcpSnoopEnable sets DhcpSnoopEnable field to given value.

### HasDhcpSnoopEnable

`func (o *BatchStackPortSettingVO) HasDhcpSnoopEnable() bool`

HasDhcpSnoopEnable returns a boolean if a field has been set.

### GetDisable

`func (o *BatchStackPortSettingVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *BatchStackPortSettingVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *BatchStackPortSettingVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *BatchStackPortSettingVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *BatchStackPortSettingVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *BatchStackPortSettingVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *BatchStackPortSettingVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *BatchStackPortSettingVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDot1x

`func (o *BatchStackPortSettingVO) GetDot1x() int32`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *BatchStackPortSettingVO) GetDot1xOk() (*int32, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *BatchStackPortSettingVO) SetDot1x(v int32)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *BatchStackPortSettingVO) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetDuplex

`func (o *BatchStackPortSettingVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *BatchStackPortSettingVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *BatchStackPortSettingVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *BatchStackPortSettingVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEeeEnable

`func (o *BatchStackPortSettingVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *BatchStackPortSettingVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *BatchStackPortSettingVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *BatchStackPortSettingVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *BatchStackPortSettingVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *BatchStackPortSettingVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *BatchStackPortSettingVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *BatchStackPortSettingVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFecLinkPeerApplyEnable

`func (o *BatchStackPortSettingVO) GetFecLinkPeerApplyEnable() bool`

GetFecLinkPeerApplyEnable returns the FecLinkPeerApplyEnable field if non-nil, zero value otherwise.

### GetFecLinkPeerApplyEnableOk

`func (o *BatchStackPortSettingVO) GetFecLinkPeerApplyEnableOk() (*bool, bool)`

GetFecLinkPeerApplyEnableOk returns a tuple with the FecLinkPeerApplyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecLinkPeerApplyEnable

`func (o *BatchStackPortSettingVO) SetFecLinkPeerApplyEnable(v bool)`

SetFecLinkPeerApplyEnable sets FecLinkPeerApplyEnable field to given value.

### HasFecLinkPeerApplyEnable

`func (o *BatchStackPortSettingVO) HasFecLinkPeerApplyEnable() bool`

HasFecLinkPeerApplyEnable returns a boolean if a field has been set.

### GetFecMode

`func (o *BatchStackPortSettingVO) GetFecMode() int32`

GetFecMode returns the FecMode field if non-nil, zero value otherwise.

### GetFecModeOk

`func (o *BatchStackPortSettingVO) GetFecModeOk() (*int32, bool)`

GetFecModeOk returns a tuple with the FecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecMode

`func (o *BatchStackPortSettingVO) SetFecMode(v int32)`

SetFecMode sets FecMode field to given value.

### HasFecMode

`func (o *BatchStackPortSettingVO) HasFecMode() bool`

HasFecMode returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *BatchStackPortSettingVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *BatchStackPortSettingVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *BatchStackPortSettingVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *BatchStackPortSettingVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *BatchStackPortSettingVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *BatchStackPortSettingVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *BatchStackPortSettingVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *BatchStackPortSettingVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetIgmpSnoopingEnable

`func (o *BatchStackPortSettingVO) GetIgmpSnoopingEnable() bool`

GetIgmpSnoopingEnable returns the IgmpSnoopingEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnableOk

`func (o *BatchStackPortSettingVO) GetIgmpSnoopingEnableOk() (*bool, bool)`

GetIgmpSnoopingEnableOk returns a tuple with the IgmpSnoopingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnable

`func (o *BatchStackPortSettingVO) SetIgmpSnoopingEnable(v bool)`

SetIgmpSnoopingEnable sets IgmpSnoopingEnable field to given value.

### HasIgmpSnoopingEnable

`func (o *BatchStackPortSettingVO) HasIgmpSnoopingEnable() bool`

HasIgmpSnoopingEnable returns a boolean if a field has been set.

### GetImpbs

`func (o *BatchStackPortSettingVO) GetImpbs() []ImpbVO`

GetImpbs returns the Impbs field if non-nil, zero value otherwise.

### GetImpbsOk

`func (o *BatchStackPortSettingVO) GetImpbsOk() (*[]ImpbVO, bool)`

GetImpbsOk returns a tuple with the Impbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpbs

`func (o *BatchStackPortSettingVO) SetImpbs(v []ImpbVO)`

SetImpbs sets Impbs field to given value.

### HasImpbs

`func (o *BatchStackPortSettingVO) HasImpbs() bool`

HasImpbs returns a boolean if a field has been set.

### GetLagSetting

`func (o *BatchStackPortSettingVO) GetLagSetting() OswLagBasicVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *BatchStackPortSettingVO) GetLagSettingOk() (*OswLagBasicVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *BatchStackPortSettingVO) SetLagSetting(v OswLagBasicVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *BatchStackPortSettingVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *BatchStackPortSettingVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *BatchStackPortSettingVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *BatchStackPortSettingVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *BatchStackPortSettingVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLldpMedEnable

`func (o *BatchStackPortSettingVO) GetLldpMedEnable() bool`

GetLldpMedEnable returns the LldpMedEnable field if non-nil, zero value otherwise.

### GetLldpMedEnableOk

`func (o *BatchStackPortSettingVO) GetLldpMedEnableOk() (*bool, bool)`

GetLldpMedEnableOk returns a tuple with the LldpMedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMedEnable

`func (o *BatchStackPortSettingVO) SetLldpMedEnable(v bool)`

SetLldpMedEnable sets LldpMedEnable field to given value.

### HasLldpMedEnable

`func (o *BatchStackPortSettingVO) HasLldpMedEnable() bool`

HasLldpMedEnable returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *BatchStackPortSettingVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *BatchStackPortSettingVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *BatchStackPortSettingVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *BatchStackPortSettingVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackDetectVlanBasedEnable

`func (o *BatchStackPortSettingVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *BatchStackPortSettingVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *BatchStackPortSettingVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *BatchStackPortSettingVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMac

`func (o *BatchStackPortSettingVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *BatchStackPortSettingVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *BatchStackPortSettingVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *BatchStackPortSettingVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMirroredLags

`func (o *BatchStackPortSettingVO) GetMirroredLags() []int32`

GetMirroredLags returns the MirroredLags field if non-nil, zero value otherwise.

### GetMirroredLagsOk

`func (o *BatchStackPortSettingVO) GetMirroredLagsOk() (*[]int32, bool)`

GetMirroredLagsOk returns a tuple with the MirroredLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredLags

`func (o *BatchStackPortSettingVO) SetMirroredLags(v []int32)`

SetMirroredLags sets MirroredLags field to given value.

### HasMirroredLags

`func (o *BatchStackPortSettingVO) HasMirroredLags() bool`

HasMirroredLags returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *BatchStackPortSettingVO) GetMirroredPorts() []int32`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *BatchStackPortSettingVO) GetMirroredPortsOk() (*[]int32, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *BatchStackPortSettingVO) SetMirroredPorts(v []int32)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *BatchStackPortSettingVO) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *BatchStackPortSettingVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *BatchStackPortSettingVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *BatchStackPortSettingVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *BatchStackPortSettingVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *BatchStackPortSettingVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchStackPortSettingVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchStackPortSettingVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BatchStackPortSettingVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *BatchStackPortSettingVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *BatchStackPortSettingVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *BatchStackPortSettingVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *BatchStackPortSettingVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *BatchStackPortSettingVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *BatchStackPortSettingVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *BatchStackPortSettingVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *BatchStackPortSettingVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *BatchStackPortSettingVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *BatchStackPortSettingVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *BatchStackPortSettingVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *BatchStackPortSettingVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetOperation

`func (o *BatchStackPortSettingVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BatchStackPortSettingVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BatchStackPortSettingVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BatchStackPortSettingVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPoe

`func (o *BatchStackPortSettingVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *BatchStackPortSettingVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *BatchStackPortSettingVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *BatchStackPortSettingVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPortAlertEnable

`func (o *BatchStackPortSettingVO) GetPortAlertEnable() bool`

GetPortAlertEnable returns the PortAlertEnable field if non-nil, zero value otherwise.

### GetPortAlertEnableOk

`func (o *BatchStackPortSettingVO) GetPortAlertEnableOk() (*bool, bool)`

GetPortAlertEnableOk returns a tuple with the PortAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAlertEnable

`func (o *BatchStackPortSettingVO) SetPortAlertEnable(v bool)`

SetPortAlertEnable sets PortAlertEnable field to given value.

### HasPortAlertEnable

`func (o *BatchStackPortSettingVO) HasPortAlertEnable() bool`

HasPortAlertEnable returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *BatchStackPortSettingVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *BatchStackPortSettingVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *BatchStackPortSettingVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *BatchStackPortSettingVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetProfileId

`func (o *BatchStackPortSettingVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *BatchStackPortSettingVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *BatchStackPortSettingVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *BatchStackPortSettingVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *BatchStackPortSettingVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *BatchStackPortSettingVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *BatchStackPortSettingVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *BatchStackPortSettingVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetQosQueueEnable

`func (o *BatchStackPortSettingVO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *BatchStackPortSettingVO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *BatchStackPortSettingVO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *BatchStackPortSettingVO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQueueId

`func (o *BatchStackPortSettingVO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *BatchStackPortSettingVO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *BatchStackPortSettingVO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *BatchStackPortSettingVO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetSpanningTreeEnable

`func (o *BatchStackPortSettingVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *BatchStackPortSettingVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *BatchStackPortSettingVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.

### HasSpanningTreeEnable

`func (o *BatchStackPortSettingVO) HasSpanningTreeEnable() bool`

HasSpanningTreeEnable returns a boolean if a field has been set.

### GetSpanningTreeSetting

`func (o *BatchStackPortSettingVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *BatchStackPortSettingVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *BatchStackPortSettingVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *BatchStackPortSettingVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetStackSetting

`func (o *BatchStackPortSettingVO) GetStackSetting() OswPortStackSettingVO`

GetStackSetting returns the StackSetting field if non-nil, zero value otherwise.

### GetStackSettingOk

`func (o *BatchStackPortSettingVO) GetStackSettingOk() (*OswPortStackSettingVO, bool)`

GetStackSettingOk returns a tuple with the StackSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackSetting

`func (o *BatchStackPortSettingVO) SetStackSetting(v OswPortStackSettingVO)`

SetStackSetting sets StackSetting field to given value.

### HasStackSetting

`func (o *BatchStackPortSettingVO) HasStackSetting() bool`

HasStackSetting returns a boolean if a field has been set.

### GetStormCtrl

`func (o *BatchStackPortSettingVO) GetStormCtrl() OswStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *BatchStackPortSettingVO) GetStormCtrlOk() (*OswStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *BatchStackPortSettingVO) SetStormCtrl(v OswStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *BatchStackPortSettingVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetTagBridgeVlanMap

`func (o *BatchStackPortSettingVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *BatchStackPortSettingVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *BatchStackPortSettingVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *BatchStackPortSettingVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagIds

`func (o *BatchStackPortSettingVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *BatchStackPortSettingVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *BatchStackPortSettingVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *BatchStackPortSettingVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *BatchStackPortSettingVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *BatchStackPortSettingVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *BatchStackPortSettingVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *BatchStackPortSettingVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTopoNotifyEnable

`func (o *BatchStackPortSettingVO) GetTopoNotifyEnable() bool`

GetTopoNotifyEnable returns the TopoNotifyEnable field if non-nil, zero value otherwise.

### GetTopoNotifyEnableOk

`func (o *BatchStackPortSettingVO) GetTopoNotifyEnableOk() (*bool, bool)`

GetTopoNotifyEnableOk returns a tuple with the TopoNotifyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopoNotifyEnable

`func (o *BatchStackPortSettingVO) SetTopoNotifyEnable(v bool)`

SetTopoNotifyEnable sets TopoNotifyEnable field to given value.

### HasTopoNotifyEnable

`func (o *BatchStackPortSettingVO) HasTopoNotifyEnable() bool`

HasTopoNotifyEnable returns a boolean if a field has been set.

### GetTrustMode

`func (o *BatchStackPortSettingVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *BatchStackPortSettingVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *BatchStackPortSettingVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *BatchStackPortSettingVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetUnitList

`func (o *BatchStackPortSettingVO) GetUnitList() []OswStackUnitVO`

GetUnitList returns the UnitList field if non-nil, zero value otherwise.

### GetUnitListOk

`func (o *BatchStackPortSettingVO) GetUnitListOk() (*[]OswStackUnitVO, bool)`

GetUnitListOk returns a tuple with the UnitList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitList

`func (o *BatchStackPortSettingVO) SetUnitList(v []OswStackUnitVO)`

SetUnitList sets UnitList field to given value.


### GetUntagBridgeVlanMap

`func (o *BatchStackPortSettingVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *BatchStackPortSettingVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *BatchStackPortSettingVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *BatchStackPortSettingVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *BatchStackPortSettingVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *BatchStackPortSettingVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *BatchStackPortSettingVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *BatchStackPortSettingVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *BatchStackPortSettingVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *BatchStackPortSettingVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *BatchStackPortSettingVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *BatchStackPortSettingVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceDscp

`func (o *BatchStackPortSettingVO) GetVoiceDscp() int32`

GetVoiceDscp returns the VoiceDscp field if non-nil, zero value otherwise.

### GetVoiceDscpOk

`func (o *BatchStackPortSettingVO) GetVoiceDscpOk() (*int32, bool)`

GetVoiceDscpOk returns a tuple with the VoiceDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscp

`func (o *BatchStackPortSettingVO) SetVoiceDscp(v int32)`

SetVoiceDscp sets VoiceDscp field to given value.

### HasVoiceDscp

`func (o *BatchStackPortSettingVO) HasVoiceDscp() bool`

HasVoiceDscp returns a boolean if a field has been set.

### GetVoiceDscpEnable

`func (o *BatchStackPortSettingVO) GetVoiceDscpEnable() bool`

GetVoiceDscpEnable returns the VoiceDscpEnable field if non-nil, zero value otherwise.

### GetVoiceDscpEnableOk

`func (o *BatchStackPortSettingVO) GetVoiceDscpEnableOk() (*bool, bool)`

GetVoiceDscpEnableOk returns a tuple with the VoiceDscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpEnable

`func (o *BatchStackPortSettingVO) SetVoiceDscpEnable(v bool)`

SetVoiceDscpEnable sets VoiceDscpEnable field to given value.

### HasVoiceDscpEnable

`func (o *BatchStackPortSettingVO) HasVoiceDscpEnable() bool`

HasVoiceDscpEnable returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *BatchStackPortSettingVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *BatchStackPortSettingVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *BatchStackPortSettingVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *BatchStackPortSettingVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *BatchStackPortSettingVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *BatchStackPortSettingVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *BatchStackPortSettingVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *BatchStackPortSettingVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


