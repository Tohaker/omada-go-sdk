# StackPortSettingVO

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
**UntagBridgeVlanMap** | Pointer to **map[string][]int32** | Untag Network Bridge Vlan Map | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag Network IDs | [optional] 
**VoiceBridgeVlan** | Pointer to **int32** | Voice Network Bridge Vlan | [optional] 
**VoiceDscp** | Pointer to **int32** | Voice DSCP | [optional] 
**VoiceDscpEnable** | Pointer to **bool** | Indicates whether voice DSCP is enabled | [optional] 
**VoiceNetworkEnable** | Pointer to **bool** | Indicates whether voice network is enabled | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewStackPortSettingVO

`func NewStackPortSettingVO() *StackPortSettingVO`

NewStackPortSettingVO instantiates a new StackPortSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackPortSettingVOWithDefaults

`func NewStackPortSettingVOWithDefaults() *StackPortSettingVO`

NewStackPortSettingVOWithDefaults instantiates a new StackPortSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpDetectEnable

`func (o *StackPortSettingVO) GetArpDetectEnable() bool`

GetArpDetectEnable returns the ArpDetectEnable field if non-nil, zero value otherwise.

### GetArpDetectEnableOk

`func (o *StackPortSettingVO) GetArpDetectEnableOk() (*bool, bool)`

GetArpDetectEnableOk returns a tuple with the ArpDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpDetectEnable

`func (o *StackPortSettingVO) SetArpDetectEnable(v bool)`

SetArpDetectEnable sets ArpDetectEnable field to given value.

### HasArpDetectEnable

`func (o *StackPortSettingVO) HasArpDetectEnable() bool`

HasArpDetectEnable returns a boolean if a field has been set.

### GetBandCtrl

`func (o *StackPortSettingVO) GetBandCtrl() OswBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *StackPortSettingVO) GetBandCtrlOk() (*OswBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *StackPortSettingVO) SetBandCtrl(v OswBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *StackPortSettingVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *StackPortSettingVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *StackPortSettingVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *StackPortSettingVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *StackPortSettingVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetDhcpL2RelaySettings

`func (o *StackPortSettingVO) GetDhcpL2RelaySettings() OswPortDhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *StackPortSettingVO) GetDhcpL2RelaySettingsOk() (*OswPortDhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *StackPortSettingVO) SetDhcpL2RelaySettings(v OswPortDhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *StackPortSettingVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDhcpSnoopEnable

`func (o *StackPortSettingVO) GetDhcpSnoopEnable() bool`

GetDhcpSnoopEnable returns the DhcpSnoopEnable field if non-nil, zero value otherwise.

### GetDhcpSnoopEnableOk

`func (o *StackPortSettingVO) GetDhcpSnoopEnableOk() (*bool, bool)`

GetDhcpSnoopEnableOk returns a tuple with the DhcpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnoopEnable

`func (o *StackPortSettingVO) SetDhcpSnoopEnable(v bool)`

SetDhcpSnoopEnable sets DhcpSnoopEnable field to given value.

### HasDhcpSnoopEnable

`func (o *StackPortSettingVO) HasDhcpSnoopEnable() bool`

HasDhcpSnoopEnable returns a boolean if a field has been set.

### GetDisable

`func (o *StackPortSettingVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *StackPortSettingVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *StackPortSettingVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *StackPortSettingVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *StackPortSettingVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *StackPortSettingVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *StackPortSettingVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *StackPortSettingVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDot1x

`func (o *StackPortSettingVO) GetDot1x() int32`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *StackPortSettingVO) GetDot1xOk() (*int32, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *StackPortSettingVO) SetDot1x(v int32)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *StackPortSettingVO) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetDuplex

`func (o *StackPortSettingVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *StackPortSettingVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *StackPortSettingVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *StackPortSettingVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEeeEnable

`func (o *StackPortSettingVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *StackPortSettingVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *StackPortSettingVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *StackPortSettingVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *StackPortSettingVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *StackPortSettingVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *StackPortSettingVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *StackPortSettingVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFecLinkPeerApplyEnable

`func (o *StackPortSettingVO) GetFecLinkPeerApplyEnable() bool`

GetFecLinkPeerApplyEnable returns the FecLinkPeerApplyEnable field if non-nil, zero value otherwise.

### GetFecLinkPeerApplyEnableOk

`func (o *StackPortSettingVO) GetFecLinkPeerApplyEnableOk() (*bool, bool)`

GetFecLinkPeerApplyEnableOk returns a tuple with the FecLinkPeerApplyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecLinkPeerApplyEnable

`func (o *StackPortSettingVO) SetFecLinkPeerApplyEnable(v bool)`

SetFecLinkPeerApplyEnable sets FecLinkPeerApplyEnable field to given value.

### HasFecLinkPeerApplyEnable

`func (o *StackPortSettingVO) HasFecLinkPeerApplyEnable() bool`

HasFecLinkPeerApplyEnable returns a boolean if a field has been set.

### GetFecMode

`func (o *StackPortSettingVO) GetFecMode() int32`

GetFecMode returns the FecMode field if non-nil, zero value otherwise.

### GetFecModeOk

`func (o *StackPortSettingVO) GetFecModeOk() (*int32, bool)`

GetFecModeOk returns a tuple with the FecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecMode

`func (o *StackPortSettingVO) SetFecMode(v int32)`

SetFecMode sets FecMode field to given value.

### HasFecMode

`func (o *StackPortSettingVO) HasFecMode() bool`

HasFecMode returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *StackPortSettingVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *StackPortSettingVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *StackPortSettingVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *StackPortSettingVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *StackPortSettingVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *StackPortSettingVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *StackPortSettingVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *StackPortSettingVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetIgmpSnoopingEnable

`func (o *StackPortSettingVO) GetIgmpSnoopingEnable() bool`

GetIgmpSnoopingEnable returns the IgmpSnoopingEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnableOk

`func (o *StackPortSettingVO) GetIgmpSnoopingEnableOk() (*bool, bool)`

GetIgmpSnoopingEnableOk returns a tuple with the IgmpSnoopingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnable

`func (o *StackPortSettingVO) SetIgmpSnoopingEnable(v bool)`

SetIgmpSnoopingEnable sets IgmpSnoopingEnable field to given value.

### HasIgmpSnoopingEnable

`func (o *StackPortSettingVO) HasIgmpSnoopingEnable() bool`

HasIgmpSnoopingEnable returns a boolean if a field has been set.

### GetImpbs

`func (o *StackPortSettingVO) GetImpbs() []ImpbVO`

GetImpbs returns the Impbs field if non-nil, zero value otherwise.

### GetImpbsOk

`func (o *StackPortSettingVO) GetImpbsOk() (*[]ImpbVO, bool)`

GetImpbsOk returns a tuple with the Impbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpbs

`func (o *StackPortSettingVO) SetImpbs(v []ImpbVO)`

SetImpbs sets Impbs field to given value.

### HasImpbs

`func (o *StackPortSettingVO) HasImpbs() bool`

HasImpbs returns a boolean if a field has been set.

### GetLagSetting

`func (o *StackPortSettingVO) GetLagSetting() OswLagBasicVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *StackPortSettingVO) GetLagSettingOk() (*OswLagBasicVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *StackPortSettingVO) SetLagSetting(v OswLagBasicVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *StackPortSettingVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *StackPortSettingVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *StackPortSettingVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *StackPortSettingVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *StackPortSettingVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLldpMedEnable

`func (o *StackPortSettingVO) GetLldpMedEnable() bool`

GetLldpMedEnable returns the LldpMedEnable field if non-nil, zero value otherwise.

### GetLldpMedEnableOk

`func (o *StackPortSettingVO) GetLldpMedEnableOk() (*bool, bool)`

GetLldpMedEnableOk returns a tuple with the LldpMedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMedEnable

`func (o *StackPortSettingVO) SetLldpMedEnable(v bool)`

SetLldpMedEnable sets LldpMedEnable field to given value.

### HasLldpMedEnable

`func (o *StackPortSettingVO) HasLldpMedEnable() bool`

HasLldpMedEnable returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *StackPortSettingVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *StackPortSettingVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *StackPortSettingVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *StackPortSettingVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackDetectVlanBasedEnable

`func (o *StackPortSettingVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *StackPortSettingVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *StackPortSettingVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *StackPortSettingVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMac

`func (o *StackPortSettingVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *StackPortSettingVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *StackPortSettingVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *StackPortSettingVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMirroredLags

`func (o *StackPortSettingVO) GetMirroredLags() []int32`

GetMirroredLags returns the MirroredLags field if non-nil, zero value otherwise.

### GetMirroredLagsOk

`func (o *StackPortSettingVO) GetMirroredLagsOk() (*[]int32, bool)`

GetMirroredLagsOk returns a tuple with the MirroredLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredLags

`func (o *StackPortSettingVO) SetMirroredLags(v []int32)`

SetMirroredLags sets MirroredLags field to given value.

### HasMirroredLags

`func (o *StackPortSettingVO) HasMirroredLags() bool`

HasMirroredLags returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *StackPortSettingVO) GetMirroredPorts() []int32`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *StackPortSettingVO) GetMirroredPortsOk() (*[]int32, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *StackPortSettingVO) SetMirroredPorts(v []int32)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *StackPortSettingVO) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *StackPortSettingVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *StackPortSettingVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *StackPortSettingVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *StackPortSettingVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *StackPortSettingVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackPortSettingVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackPortSettingVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StackPortSettingVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *StackPortSettingVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *StackPortSettingVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *StackPortSettingVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *StackPortSettingVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *StackPortSettingVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *StackPortSettingVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *StackPortSettingVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *StackPortSettingVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *StackPortSettingVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *StackPortSettingVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *StackPortSettingVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *StackPortSettingVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetOperation

`func (o *StackPortSettingVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *StackPortSettingVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *StackPortSettingVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *StackPortSettingVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPoe

`func (o *StackPortSettingVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *StackPortSettingVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *StackPortSettingVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *StackPortSettingVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPortAlertEnable

`func (o *StackPortSettingVO) GetPortAlertEnable() bool`

GetPortAlertEnable returns the PortAlertEnable field if non-nil, zero value otherwise.

### GetPortAlertEnableOk

`func (o *StackPortSettingVO) GetPortAlertEnableOk() (*bool, bool)`

GetPortAlertEnableOk returns a tuple with the PortAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAlertEnable

`func (o *StackPortSettingVO) SetPortAlertEnable(v bool)`

SetPortAlertEnable sets PortAlertEnable field to given value.

### HasPortAlertEnable

`func (o *StackPortSettingVO) HasPortAlertEnable() bool`

HasPortAlertEnable returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *StackPortSettingVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *StackPortSettingVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *StackPortSettingVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *StackPortSettingVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetProfileId

`func (o *StackPortSettingVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *StackPortSettingVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *StackPortSettingVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *StackPortSettingVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *StackPortSettingVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *StackPortSettingVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *StackPortSettingVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *StackPortSettingVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetQosQueueEnable

`func (o *StackPortSettingVO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *StackPortSettingVO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *StackPortSettingVO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *StackPortSettingVO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQueueId

`func (o *StackPortSettingVO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *StackPortSettingVO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *StackPortSettingVO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *StackPortSettingVO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetSpanningTreeEnable

`func (o *StackPortSettingVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *StackPortSettingVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *StackPortSettingVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.

### HasSpanningTreeEnable

`func (o *StackPortSettingVO) HasSpanningTreeEnable() bool`

HasSpanningTreeEnable returns a boolean if a field has been set.

### GetSpanningTreeSetting

`func (o *StackPortSettingVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *StackPortSettingVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *StackPortSettingVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *StackPortSettingVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetStackSetting

`func (o *StackPortSettingVO) GetStackSetting() OswPortStackSettingVO`

GetStackSetting returns the StackSetting field if non-nil, zero value otherwise.

### GetStackSettingOk

`func (o *StackPortSettingVO) GetStackSettingOk() (*OswPortStackSettingVO, bool)`

GetStackSettingOk returns a tuple with the StackSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackSetting

`func (o *StackPortSettingVO) SetStackSetting(v OswPortStackSettingVO)`

SetStackSetting sets StackSetting field to given value.

### HasStackSetting

`func (o *StackPortSettingVO) HasStackSetting() bool`

HasStackSetting returns a boolean if a field has been set.

### GetStormCtrl

`func (o *StackPortSettingVO) GetStormCtrl() OswStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *StackPortSettingVO) GetStormCtrlOk() (*OswStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *StackPortSettingVO) SetStormCtrl(v OswStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *StackPortSettingVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetTagBridgeVlanMap

`func (o *StackPortSettingVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *StackPortSettingVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *StackPortSettingVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *StackPortSettingVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagIds

`func (o *StackPortSettingVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *StackPortSettingVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *StackPortSettingVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *StackPortSettingVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *StackPortSettingVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *StackPortSettingVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *StackPortSettingVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *StackPortSettingVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTopoNotifyEnable

`func (o *StackPortSettingVO) GetTopoNotifyEnable() bool`

GetTopoNotifyEnable returns the TopoNotifyEnable field if non-nil, zero value otherwise.

### GetTopoNotifyEnableOk

`func (o *StackPortSettingVO) GetTopoNotifyEnableOk() (*bool, bool)`

GetTopoNotifyEnableOk returns a tuple with the TopoNotifyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopoNotifyEnable

`func (o *StackPortSettingVO) SetTopoNotifyEnable(v bool)`

SetTopoNotifyEnable sets TopoNotifyEnable field to given value.

### HasTopoNotifyEnable

`func (o *StackPortSettingVO) HasTopoNotifyEnable() bool`

HasTopoNotifyEnable returns a boolean if a field has been set.

### GetTrustMode

`func (o *StackPortSettingVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *StackPortSettingVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *StackPortSettingVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *StackPortSettingVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetUntagBridgeVlanMap

`func (o *StackPortSettingVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *StackPortSettingVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *StackPortSettingVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *StackPortSettingVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *StackPortSettingVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *StackPortSettingVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *StackPortSettingVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *StackPortSettingVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *StackPortSettingVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *StackPortSettingVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *StackPortSettingVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *StackPortSettingVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceDscp

`func (o *StackPortSettingVO) GetVoiceDscp() int32`

GetVoiceDscp returns the VoiceDscp field if non-nil, zero value otherwise.

### GetVoiceDscpOk

`func (o *StackPortSettingVO) GetVoiceDscpOk() (*int32, bool)`

GetVoiceDscpOk returns a tuple with the VoiceDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscp

`func (o *StackPortSettingVO) SetVoiceDscp(v int32)`

SetVoiceDscp sets VoiceDscp field to given value.

### HasVoiceDscp

`func (o *StackPortSettingVO) HasVoiceDscp() bool`

HasVoiceDscp returns a boolean if a field has been set.

### GetVoiceDscpEnable

`func (o *StackPortSettingVO) GetVoiceDscpEnable() bool`

GetVoiceDscpEnable returns the VoiceDscpEnable field if non-nil, zero value otherwise.

### GetVoiceDscpEnableOk

`func (o *StackPortSettingVO) GetVoiceDscpEnableOk() (*bool, bool)`

GetVoiceDscpEnableOk returns a tuple with the VoiceDscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpEnable

`func (o *StackPortSettingVO) SetVoiceDscpEnable(v bool)`

SetVoiceDscpEnable sets VoiceDscpEnable field to given value.

### HasVoiceDscpEnable

`func (o *StackPortSettingVO) HasVoiceDscpEnable() bool`

HasVoiceDscpEnable returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *StackPortSettingVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *StackPortSettingVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *StackPortSettingVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *StackPortSettingVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *StackPortSettingVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *StackPortSettingVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *StackPortSettingVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *StackPortSettingVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


