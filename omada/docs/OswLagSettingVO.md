# OswLagSettingVO

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
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**EeeEnable** | Pointer to **bool** | Indicates whether EEE is enabled | [optional] 
**FastLeaveEnable** | Pointer to **bool** | Indicates whether igmpSnooping fastLeave is enabled | [optional] 
**FecLinkPeerApplyEnable** | Pointer to **bool** | Indicates whether the FEC Mode is synchronously applied to the link peer port. | [optional] 
**FecMode** | Pointer to **int32** | FEC mode should be a value as follows: 1: Off; 2: RS528; 3: RS544; 4: Auto; 5: Base-R | [optional] 
**FlowControlEnable** | Pointer to **bool** | Indicates whether flow control is enabled | [optional] 
**IgmpFastLeaveEnable** | Pointer to **bool** | Indicates whether igmp fast leave is enabled | [optional] 
**IgmpSnoopingEnable** | Pointer to **bool** | Indicates whether IGMP Snooping is enabled | [optional] 
**Impbs** | Pointer to [**[]ImpbVO**](ImpbVO.md) | Impbs | [optional] 
**LagId** | Pointer to **int32** | Lag ID | [optional] 
**LagType** | Pointer to **int32** | Lag Type should be a value as follows: 1: STATIC; 2: LACP; 3: LACP ACTIVE; 4: LACP PASSIVE | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed should be a value as follows: 0: auto; 1: 10M; 2: 100M; 3: 1000M; 4: 2.5G; 5: 10G | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | Indicates whether loopbackDetect port based is enabled | [optional] 
**LoopbackDetectVlanBasedEnable** | Pointer to **bool** | Indicates whether loopbackDetect vlan based is enabled | [optional] 
**MlagEnable** | Pointer to **bool** | Indicates whether M-LAG port is enabled | [optional] 
**MlagName** | Pointer to **string** | M-LAG group Name | [optional] 
**MlagPeerSetting** | Pointer to [**OswMlagPeerSettingVO**](OswMlagPeerSettingVO.md) |  | [optional] 
**MldFastLeaveEnable** | Pointer to **bool** | Indicates whether mld fast leave is enabled | [optional] 
**Name** | Pointer to **string** | Port or Lag Name | [optional] 
**NativeBridgeVlan** | Pointer to **int32** | Native Network Bridge Vlan. | [optional] 
**NativeNetworkId** | Pointer to **string** | Native Network ID, Native Network cannot be selected from Tagged Networks or Untagged Networks. | [optional] 
**NetworkTagsSetting** | Pointer to **int32** | Network Tags Setting should be a value as follows: 0: Allow All; 1: Block All; 2: Custom | [optional] 
**PortAlertEnable** | Pointer to **bool** | Indicates whether port alert is enabled | [optional] 
**PortIsolationEnable** | Pointer to **bool** | Indicates whether port isolation is enabled | [optional] 
**Ports** | Pointer to **[]int32** | Lag Ports | [optional] 
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
**TrustMode** | Pointer to **int32** | TrustMode should be a value as follows: 0: Untrusted; 1: Trust 802.1p; 2: Trust DSCP | [optional] 
**UntagBridgeVlanMap** | Pointer to **map[string][]int32** | Untag Network Bridge Vlan Map | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag Network IDs | [optional] 
**VoiceBridgeVlan** | Pointer to **int32** | Voice Network Bridge Vlan | [optional] 
**VoiceDscp** | Pointer to **int32** | Voice DSCP | [optional] 
**VoiceDscpEnable** | Pointer to **bool** | Indicates whether voice DSCP is enabled | [optional] 
**VoiceNetworkEnable** | Pointer to **bool** | Indicates whether voice network is enabled | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewOswLagSettingVO

`func NewOswLagSettingVO() *OswLagSettingVO`

NewOswLagSettingVO instantiates a new OswLagSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswLagSettingVOWithDefaults

`func NewOswLagSettingVOWithDefaults() *OswLagSettingVO`

NewOswLagSettingVOWithDefaults instantiates a new OswLagSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpDetectEnable

`func (o *OswLagSettingVO) GetArpDetectEnable() bool`

GetArpDetectEnable returns the ArpDetectEnable field if non-nil, zero value otherwise.

### GetArpDetectEnableOk

`func (o *OswLagSettingVO) GetArpDetectEnableOk() (*bool, bool)`

GetArpDetectEnableOk returns a tuple with the ArpDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpDetectEnable

`func (o *OswLagSettingVO) SetArpDetectEnable(v bool)`

SetArpDetectEnable sets ArpDetectEnable field to given value.

### HasArpDetectEnable

`func (o *OswLagSettingVO) HasArpDetectEnable() bool`

HasArpDetectEnable returns a boolean if a field has been set.

### GetBandCtrl

`func (o *OswLagSettingVO) GetBandCtrl() OswBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *OswLagSettingVO) GetBandCtrlOk() (*OswBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *OswLagSettingVO) SetBandCtrl(v OswBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *OswLagSettingVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *OswLagSettingVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *OswLagSettingVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *OswLagSettingVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *OswLagSettingVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetDhcpL2RelaySettings

`func (o *OswLagSettingVO) GetDhcpL2RelaySettings() OswPortDhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *OswLagSettingVO) GetDhcpL2RelaySettingsOk() (*OswPortDhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *OswLagSettingVO) SetDhcpL2RelaySettings(v OswPortDhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *OswLagSettingVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDhcpSnoopEnable

`func (o *OswLagSettingVO) GetDhcpSnoopEnable() bool`

GetDhcpSnoopEnable returns the DhcpSnoopEnable field if non-nil, zero value otherwise.

### GetDhcpSnoopEnableOk

`func (o *OswLagSettingVO) GetDhcpSnoopEnableOk() (*bool, bool)`

GetDhcpSnoopEnableOk returns a tuple with the DhcpSnoopEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnoopEnable

`func (o *OswLagSettingVO) SetDhcpSnoopEnable(v bool)`

SetDhcpSnoopEnable sets DhcpSnoopEnable field to given value.

### HasDhcpSnoopEnable

`func (o *OswLagSettingVO) HasDhcpSnoopEnable() bool`

HasDhcpSnoopEnable returns a boolean if a field has been set.

### GetDisable

`func (o *OswLagSettingVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswLagSettingVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswLagSettingVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswLagSettingVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *OswLagSettingVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *OswLagSettingVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *OswLagSettingVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *OswLagSettingVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDuplex

`func (o *OswLagSettingVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswLagSettingVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswLagSettingVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswLagSettingVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEeeEnable

`func (o *OswLagSettingVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *OswLagSettingVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *OswLagSettingVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *OswLagSettingVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *OswLagSettingVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *OswLagSettingVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *OswLagSettingVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *OswLagSettingVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFecLinkPeerApplyEnable

`func (o *OswLagSettingVO) GetFecLinkPeerApplyEnable() bool`

GetFecLinkPeerApplyEnable returns the FecLinkPeerApplyEnable field if non-nil, zero value otherwise.

### GetFecLinkPeerApplyEnableOk

`func (o *OswLagSettingVO) GetFecLinkPeerApplyEnableOk() (*bool, bool)`

GetFecLinkPeerApplyEnableOk returns a tuple with the FecLinkPeerApplyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecLinkPeerApplyEnable

`func (o *OswLagSettingVO) SetFecLinkPeerApplyEnable(v bool)`

SetFecLinkPeerApplyEnable sets FecLinkPeerApplyEnable field to given value.

### HasFecLinkPeerApplyEnable

`func (o *OswLagSettingVO) HasFecLinkPeerApplyEnable() bool`

HasFecLinkPeerApplyEnable returns a boolean if a field has been set.

### GetFecMode

`func (o *OswLagSettingVO) GetFecMode() int32`

GetFecMode returns the FecMode field if non-nil, zero value otherwise.

### GetFecModeOk

`func (o *OswLagSettingVO) GetFecModeOk() (*int32, bool)`

GetFecModeOk returns a tuple with the FecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecMode

`func (o *OswLagSettingVO) SetFecMode(v int32)`

SetFecMode sets FecMode field to given value.

### HasFecMode

`func (o *OswLagSettingVO) HasFecMode() bool`

HasFecMode returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *OswLagSettingVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *OswLagSettingVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *OswLagSettingVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *OswLagSettingVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *OswLagSettingVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *OswLagSettingVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *OswLagSettingVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *OswLagSettingVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetIgmpSnoopingEnable

`func (o *OswLagSettingVO) GetIgmpSnoopingEnable() bool`

GetIgmpSnoopingEnable returns the IgmpSnoopingEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnableOk

`func (o *OswLagSettingVO) GetIgmpSnoopingEnableOk() (*bool, bool)`

GetIgmpSnoopingEnableOk returns a tuple with the IgmpSnoopingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnable

`func (o *OswLagSettingVO) SetIgmpSnoopingEnable(v bool)`

SetIgmpSnoopingEnable sets IgmpSnoopingEnable field to given value.

### HasIgmpSnoopingEnable

`func (o *OswLagSettingVO) HasIgmpSnoopingEnable() bool`

HasIgmpSnoopingEnable returns a boolean if a field has been set.

### GetImpbs

`func (o *OswLagSettingVO) GetImpbs() []ImpbVO`

GetImpbs returns the Impbs field if non-nil, zero value otherwise.

### GetImpbsOk

`func (o *OswLagSettingVO) GetImpbsOk() (*[]ImpbVO, bool)`

GetImpbsOk returns a tuple with the Impbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpbs

`func (o *OswLagSettingVO) SetImpbs(v []ImpbVO)`

SetImpbs sets Impbs field to given value.

### HasImpbs

`func (o *OswLagSettingVO) HasImpbs() bool`

HasImpbs returns a boolean if a field has been set.

### GetLagId

`func (o *OswLagSettingVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswLagSettingVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswLagSettingVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OswLagSettingVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLagType

`func (o *OswLagSettingVO) GetLagType() int32`

GetLagType returns the LagType field if non-nil, zero value otherwise.

### GetLagTypeOk

`func (o *OswLagSettingVO) GetLagTypeOk() (*int32, bool)`

GetLagTypeOk returns a tuple with the LagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagType

`func (o *OswLagSettingVO) SetLagType(v int32)`

SetLagType sets LagType field to given value.

### HasLagType

`func (o *OswLagSettingVO) HasLagType() bool`

HasLagType returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswLagSettingVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswLagSettingVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswLagSettingVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswLagSettingVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswLagSettingVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswLagSettingVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswLagSettingVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswLagSettingVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackDetectVlanBasedEnable

`func (o *OswLagSettingVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *OswLagSettingVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *OswLagSettingVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *OswLagSettingVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMlagEnable

`func (o *OswLagSettingVO) GetMlagEnable() bool`

GetMlagEnable returns the MlagEnable field if non-nil, zero value otherwise.

### GetMlagEnableOk

`func (o *OswLagSettingVO) GetMlagEnableOk() (*bool, bool)`

GetMlagEnableOk returns a tuple with the MlagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagEnable

`func (o *OswLagSettingVO) SetMlagEnable(v bool)`

SetMlagEnable sets MlagEnable field to given value.

### HasMlagEnable

`func (o *OswLagSettingVO) HasMlagEnable() bool`

HasMlagEnable returns a boolean if a field has been set.

### GetMlagName

`func (o *OswLagSettingVO) GetMlagName() string`

GetMlagName returns the MlagName field if non-nil, zero value otherwise.

### GetMlagNameOk

`func (o *OswLagSettingVO) GetMlagNameOk() (*string, bool)`

GetMlagNameOk returns a tuple with the MlagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagName

`func (o *OswLagSettingVO) SetMlagName(v string)`

SetMlagName sets MlagName field to given value.

### HasMlagName

`func (o *OswLagSettingVO) HasMlagName() bool`

HasMlagName returns a boolean if a field has been set.

### GetMlagPeerSetting

`func (o *OswLagSettingVO) GetMlagPeerSetting() OswMlagPeerSettingVO`

GetMlagPeerSetting returns the MlagPeerSetting field if non-nil, zero value otherwise.

### GetMlagPeerSettingOk

`func (o *OswLagSettingVO) GetMlagPeerSettingOk() (*OswMlagPeerSettingVO, bool)`

GetMlagPeerSettingOk returns a tuple with the MlagPeerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerSetting

`func (o *OswLagSettingVO) SetMlagPeerSetting(v OswMlagPeerSettingVO)`

SetMlagPeerSetting sets MlagPeerSetting field to given value.

### HasMlagPeerSetting

`func (o *OswLagSettingVO) HasMlagPeerSetting() bool`

HasMlagPeerSetting returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *OswLagSettingVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *OswLagSettingVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *OswLagSettingVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *OswLagSettingVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *OswLagSettingVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswLagSettingVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswLagSettingVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswLagSettingVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *OswLagSettingVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *OswLagSettingVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *OswLagSettingVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *OswLagSettingVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *OswLagSettingVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *OswLagSettingVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *OswLagSettingVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *OswLagSettingVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *OswLagSettingVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *OswLagSettingVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *OswLagSettingVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *OswLagSettingVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetPortAlertEnable

`func (o *OswLagSettingVO) GetPortAlertEnable() bool`

GetPortAlertEnable returns the PortAlertEnable field if non-nil, zero value otherwise.

### GetPortAlertEnableOk

`func (o *OswLagSettingVO) GetPortAlertEnableOk() (*bool, bool)`

GetPortAlertEnableOk returns a tuple with the PortAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAlertEnable

`func (o *OswLagSettingVO) SetPortAlertEnable(v bool)`

SetPortAlertEnable sets PortAlertEnable field to given value.

### HasPortAlertEnable

`func (o *OswLagSettingVO) HasPortAlertEnable() bool`

HasPortAlertEnable returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *OswLagSettingVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *OswLagSettingVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *OswLagSettingVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *OswLagSettingVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetPorts

`func (o *OswLagSettingVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswLagSettingVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswLagSettingVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswLagSettingVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProfileId

`func (o *OswLagSettingVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OswLagSettingVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OswLagSettingVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *OswLagSettingVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *OswLagSettingVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *OswLagSettingVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *OswLagSettingVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *OswLagSettingVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetQosQueueEnable

`func (o *OswLagSettingVO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *OswLagSettingVO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *OswLagSettingVO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *OswLagSettingVO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQueueId

`func (o *OswLagSettingVO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *OswLagSettingVO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *OswLagSettingVO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *OswLagSettingVO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetSpanningTreeEnable

`func (o *OswLagSettingVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *OswLagSettingVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *OswLagSettingVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.

### HasSpanningTreeEnable

`func (o *OswLagSettingVO) HasSpanningTreeEnable() bool`

HasSpanningTreeEnable returns a boolean if a field has been set.

### GetSpanningTreeSetting

`func (o *OswLagSettingVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *OswLagSettingVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *OswLagSettingVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *OswLagSettingVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetStormCtrl

`func (o *OswLagSettingVO) GetStormCtrl() OswStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *OswLagSettingVO) GetStormCtrlOk() (*OswStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *OswLagSettingVO) SetStormCtrl(v OswStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *OswLagSettingVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetTagBridgeVlanMap

`func (o *OswLagSettingVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *OswLagSettingVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *OswLagSettingVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *OswLagSettingVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagIds

`func (o *OswLagSettingVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswLagSettingVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswLagSettingVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswLagSettingVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *OswLagSettingVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *OswLagSettingVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *OswLagSettingVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *OswLagSettingVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTrustMode

`func (o *OswLagSettingVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *OswLagSettingVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *OswLagSettingVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *OswLagSettingVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetUntagBridgeVlanMap

`func (o *OswLagSettingVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *OswLagSettingVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *OswLagSettingVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *OswLagSettingVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *OswLagSettingVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *OswLagSettingVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *OswLagSettingVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *OswLagSettingVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *OswLagSettingVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *OswLagSettingVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *OswLagSettingVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *OswLagSettingVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceDscp

`func (o *OswLagSettingVO) GetVoiceDscp() int32`

GetVoiceDscp returns the VoiceDscp field if non-nil, zero value otherwise.

### GetVoiceDscpOk

`func (o *OswLagSettingVO) GetVoiceDscpOk() (*int32, bool)`

GetVoiceDscpOk returns a tuple with the VoiceDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscp

`func (o *OswLagSettingVO) SetVoiceDscp(v int32)`

SetVoiceDscp sets VoiceDscp field to given value.

### HasVoiceDscp

`func (o *OswLagSettingVO) HasVoiceDscp() bool`

HasVoiceDscp returns a boolean if a field has been set.

### GetVoiceDscpEnable

`func (o *OswLagSettingVO) GetVoiceDscpEnable() bool`

GetVoiceDscpEnable returns the VoiceDscpEnable field if non-nil, zero value otherwise.

### GetVoiceDscpEnableOk

`func (o *OswLagSettingVO) GetVoiceDscpEnableOk() (*bool, bool)`

GetVoiceDscpEnableOk returns a tuple with the VoiceDscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpEnable

`func (o *OswLagSettingVO) SetVoiceDscpEnable(v bool)`

SetVoiceDscpEnable sets VoiceDscpEnable field to given value.

### HasVoiceDscpEnable

`func (o *OswLagSettingVO) HasVoiceDscpEnable() bool`

HasVoiceDscpEnable returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *OswLagSettingVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *OswLagSettingVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *OswLagSettingVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *OswLagSettingVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *OswLagSettingVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *OswLagSettingVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *OswLagSettingVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *OswLagSettingVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


