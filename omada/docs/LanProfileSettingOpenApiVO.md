# LanProfileSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandCtrl** | Pointer to [**BandCtrlVO**](BandCtrlVO.md) |  | [optional] 
**BandWidthCtrlType** | **int32** | BandWidthCtrlType should be a value as follows: 0: off, 1: rate limit, 2: storming control | 
**DhcpL2RelaySettings** | Pointer to [**DhcpL2RelayVO**](DhcpL2RelayVO.md) |  | [optional] 
**Dot1pPriority** | Pointer to **int32** | 802.1p Priority | [optional] 
**Dot1x** | **int32** | Dot1x should be a value as follows: 0: force unauthorized, 1: force authorized, 2:auto | 
**EeeEnable** | Pointer to **bool** | EEE enable status | [optional] 
**EsEnable** | Pointer to **bool** | ES enable status | [optional] 
**EsEnableTaggedNetworkIds** | Pointer to **[]string** | ES enable tagged network ids | [optional] 
**EsTaggedModified** | Pointer to **bool** | Indicates whether ES Tagged is modified | [optional] 
**FastLeaveEnable** | Pointer to **bool** | IGMP Snooping fast leave enable status | [optional] 
**FlowControlEnable** | Pointer to **bool** | FlowControl enable status | [optional] 
**IgmpFastLeaveEnable** | Pointer to **bool** | Indicates whether igmp fast leave is enabled | [optional] 
**LldpMedEnable** | **bool** | LLDP-MED enable status | 
**LoopbackDetectEnable** | **bool** | LoopbackDetect enable status | 
**LoopbackDetectVlanBasedEnable** | Pointer to **bool** | LoopbackDetectVLANBased enable status | [optional] 
**MldFastLeaveEnable** | Pointer to **bool** | Indicates whether mld fast leave is enabled | [optional] 
**Name** | **string** | Name should contain 1 to 128 characters. | 
**NativeBridgeVlan** | Pointer to **int32** | Native Network Bridge Vlan. | [optional] 
**NativeNetworkId** | Pointer to **string** | Native Network ID, Native Network cannot be selected from Tagged Networks or Untagged Networks. | [optional] 
**NetworkTagsSetting** | Pointer to **int32** | network Tags configuration mode should be a value as follows: 0:allow all, 1:block all, 2:custom | [optional] 
**Poe** | **int32** | PoE should be a value as follows: 0: on, 1: off, 2: \&quot;do not modify\&quot; | 
**PortIsolationEnable** | **bool** | Port-isolation enable status | 
**SpanningTreeEnable** | **bool** | SpanningTree enable status | 
**SpanningTreeSetting** | Pointer to [**SpanningTreeSettingVO**](SpanningTreeSettingVO.md) |  | [optional] 
**StormCtrl** | Pointer to [**StormCtrlVO**](StormCtrlVO.md) |  | [optional] 
**TagBridgeVlanMap** | Pointer to **map[string][]int32** | Tag Network Bridge Vlan Map | [optional] 
**TagNetworkIds** | Pointer to **[]string** | Tag Network IDs | [optional] 
**TrustMode** | Pointer to **int32** | Trust mode | [optional] 
**UntagBridgeVlanMap** | Pointer to **map[string][]int32** | Untag Network Bridge Vlan Map | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag Network IDs | [optional] 
**VlanConfigEnable** | Pointer to **bool** | Vlan configuration enable status | [optional] 
**VoiceBridgeVlan** | Pointer to **int32** | Voice Network Bridge Vlan | [optional] 
**VoiceDscp** | Pointer to **int32** | Voice DSCP | [optional] 
**VoiceDscpEnable** | Pointer to **bool** | Indicates whether voice DSCP is enabled | [optional] 
**VoiceNetworkEnable** | Pointer to **bool** | Indicates whether voice network is enabled | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewLanProfileSettingOpenApiVO

`func NewLanProfileSettingOpenApiVO(bandWidthCtrlType int32, dot1x int32, lldpMedEnable bool, loopbackDetectEnable bool, name string, poe int32, portIsolationEnable bool, spanningTreeEnable bool, ) *LanProfileSettingOpenApiVO`

NewLanProfileSettingOpenApiVO instantiates a new LanProfileSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanProfileSettingOpenApiVOWithDefaults

`func NewLanProfileSettingOpenApiVOWithDefaults() *LanProfileSettingOpenApiVO`

NewLanProfileSettingOpenApiVOWithDefaults instantiates a new LanProfileSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandCtrl

`func (o *LanProfileSettingOpenApiVO) GetBandCtrl() BandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *LanProfileSettingOpenApiVO) GetBandCtrlOk() (*BandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *LanProfileSettingOpenApiVO) SetBandCtrl(v BandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *LanProfileSettingOpenApiVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *LanProfileSettingOpenApiVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *LanProfileSettingOpenApiVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *LanProfileSettingOpenApiVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.


### GetDhcpL2RelaySettings

`func (o *LanProfileSettingOpenApiVO) GetDhcpL2RelaySettings() DhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *LanProfileSettingOpenApiVO) GetDhcpL2RelaySettingsOk() (*DhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *LanProfileSettingOpenApiVO) SetDhcpL2RelaySettings(v DhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *LanProfileSettingOpenApiVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *LanProfileSettingOpenApiVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *LanProfileSettingOpenApiVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *LanProfileSettingOpenApiVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *LanProfileSettingOpenApiVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDot1x

`func (o *LanProfileSettingOpenApiVO) GetDot1x() int32`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *LanProfileSettingOpenApiVO) GetDot1xOk() (*int32, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *LanProfileSettingOpenApiVO) SetDot1x(v int32)`

SetDot1x sets Dot1x field to given value.


### GetEeeEnable

`func (o *LanProfileSettingOpenApiVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *LanProfileSettingOpenApiVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *LanProfileSettingOpenApiVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *LanProfileSettingOpenApiVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetEsEnable

`func (o *LanProfileSettingOpenApiVO) GetEsEnable() bool`

GetEsEnable returns the EsEnable field if non-nil, zero value otherwise.

### GetEsEnableOk

`func (o *LanProfileSettingOpenApiVO) GetEsEnableOk() (*bool, bool)`

GetEsEnableOk returns a tuple with the EsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsEnable

`func (o *LanProfileSettingOpenApiVO) SetEsEnable(v bool)`

SetEsEnable sets EsEnable field to given value.

### HasEsEnable

`func (o *LanProfileSettingOpenApiVO) HasEsEnable() bool`

HasEsEnable returns a boolean if a field has been set.

### GetEsEnableTaggedNetworkIds

`func (o *LanProfileSettingOpenApiVO) GetEsEnableTaggedNetworkIds() []string`

GetEsEnableTaggedNetworkIds returns the EsEnableTaggedNetworkIds field if non-nil, zero value otherwise.

### GetEsEnableTaggedNetworkIdsOk

`func (o *LanProfileSettingOpenApiVO) GetEsEnableTaggedNetworkIdsOk() (*[]string, bool)`

GetEsEnableTaggedNetworkIdsOk returns a tuple with the EsEnableTaggedNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsEnableTaggedNetworkIds

`func (o *LanProfileSettingOpenApiVO) SetEsEnableTaggedNetworkIds(v []string)`

SetEsEnableTaggedNetworkIds sets EsEnableTaggedNetworkIds field to given value.

### HasEsEnableTaggedNetworkIds

`func (o *LanProfileSettingOpenApiVO) HasEsEnableTaggedNetworkIds() bool`

HasEsEnableTaggedNetworkIds returns a boolean if a field has been set.

### GetEsTaggedModified

`func (o *LanProfileSettingOpenApiVO) GetEsTaggedModified() bool`

GetEsTaggedModified returns the EsTaggedModified field if non-nil, zero value otherwise.

### GetEsTaggedModifiedOk

`func (o *LanProfileSettingOpenApiVO) GetEsTaggedModifiedOk() (*bool, bool)`

GetEsTaggedModifiedOk returns a tuple with the EsTaggedModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsTaggedModified

`func (o *LanProfileSettingOpenApiVO) SetEsTaggedModified(v bool)`

SetEsTaggedModified sets EsTaggedModified field to given value.

### HasEsTaggedModified

`func (o *LanProfileSettingOpenApiVO) HasEsTaggedModified() bool`

HasEsTaggedModified returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *LanProfileSettingOpenApiVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *LanProfileSettingOpenApiVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *LanProfileSettingOpenApiVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *LanProfileSettingOpenApiVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *LanProfileSettingOpenApiVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *LanProfileSettingOpenApiVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetLldpMedEnable

`func (o *LanProfileSettingOpenApiVO) GetLldpMedEnable() bool`

GetLldpMedEnable returns the LldpMedEnable field if non-nil, zero value otherwise.

### GetLldpMedEnableOk

`func (o *LanProfileSettingOpenApiVO) GetLldpMedEnableOk() (*bool, bool)`

GetLldpMedEnableOk returns a tuple with the LldpMedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMedEnable

`func (o *LanProfileSettingOpenApiVO) SetLldpMedEnable(v bool)`

SetLldpMedEnable sets LldpMedEnable field to given value.


### GetLoopbackDetectEnable

`func (o *LanProfileSettingOpenApiVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *LanProfileSettingOpenApiVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *LanProfileSettingOpenApiVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.


### GetLoopbackDetectVlanBasedEnable

`func (o *LanProfileSettingOpenApiVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *LanProfileSettingOpenApiVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *LanProfileSettingOpenApiVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *LanProfileSettingOpenApiVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *LanProfileSettingOpenApiVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *LanProfileSettingOpenApiVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *LanProfileSettingOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanProfileSettingOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanProfileSettingOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNativeBridgeVlan

`func (o *LanProfileSettingOpenApiVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *LanProfileSettingOpenApiVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *LanProfileSettingOpenApiVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *LanProfileSettingOpenApiVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *LanProfileSettingOpenApiVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *LanProfileSettingOpenApiVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *LanProfileSettingOpenApiVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *LanProfileSettingOpenApiVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *LanProfileSettingOpenApiVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *LanProfileSettingOpenApiVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *LanProfileSettingOpenApiVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *LanProfileSettingOpenApiVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetPoe

`func (o *LanProfileSettingOpenApiVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *LanProfileSettingOpenApiVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *LanProfileSettingOpenApiVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.


### GetPortIsolationEnable

`func (o *LanProfileSettingOpenApiVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *LanProfileSettingOpenApiVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *LanProfileSettingOpenApiVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.


### GetSpanningTreeEnable

`func (o *LanProfileSettingOpenApiVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *LanProfileSettingOpenApiVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *LanProfileSettingOpenApiVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.


### GetSpanningTreeSetting

`func (o *LanProfileSettingOpenApiVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *LanProfileSettingOpenApiVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *LanProfileSettingOpenApiVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *LanProfileSettingOpenApiVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetStormCtrl

`func (o *LanProfileSettingOpenApiVO) GetStormCtrl() StormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *LanProfileSettingOpenApiVO) GetStormCtrlOk() (*StormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *LanProfileSettingOpenApiVO) SetStormCtrl(v StormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *LanProfileSettingOpenApiVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetTagBridgeVlanMap

`func (o *LanProfileSettingOpenApiVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *LanProfileSettingOpenApiVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *LanProfileSettingOpenApiVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *LanProfileSettingOpenApiVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *LanProfileSettingOpenApiVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *LanProfileSettingOpenApiVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *LanProfileSettingOpenApiVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *LanProfileSettingOpenApiVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTrustMode

`func (o *LanProfileSettingOpenApiVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *LanProfileSettingOpenApiVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *LanProfileSettingOpenApiVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *LanProfileSettingOpenApiVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetUntagBridgeVlanMap

`func (o *LanProfileSettingOpenApiVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *LanProfileSettingOpenApiVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *LanProfileSettingOpenApiVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *LanProfileSettingOpenApiVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *LanProfileSettingOpenApiVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *LanProfileSettingOpenApiVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *LanProfileSettingOpenApiVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *LanProfileSettingOpenApiVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVlanConfigEnable

`func (o *LanProfileSettingOpenApiVO) GetVlanConfigEnable() bool`

GetVlanConfigEnable returns the VlanConfigEnable field if non-nil, zero value otherwise.

### GetVlanConfigEnableOk

`func (o *LanProfileSettingOpenApiVO) GetVlanConfigEnableOk() (*bool, bool)`

GetVlanConfigEnableOk returns a tuple with the VlanConfigEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfigEnable

`func (o *LanProfileSettingOpenApiVO) SetVlanConfigEnable(v bool)`

SetVlanConfigEnable sets VlanConfigEnable field to given value.

### HasVlanConfigEnable

`func (o *LanProfileSettingOpenApiVO) HasVlanConfigEnable() bool`

HasVlanConfigEnable returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *LanProfileSettingOpenApiVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *LanProfileSettingOpenApiVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *LanProfileSettingOpenApiVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *LanProfileSettingOpenApiVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceDscp

`func (o *LanProfileSettingOpenApiVO) GetVoiceDscp() int32`

GetVoiceDscp returns the VoiceDscp field if non-nil, zero value otherwise.

### GetVoiceDscpOk

`func (o *LanProfileSettingOpenApiVO) GetVoiceDscpOk() (*int32, bool)`

GetVoiceDscpOk returns a tuple with the VoiceDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscp

`func (o *LanProfileSettingOpenApiVO) SetVoiceDscp(v int32)`

SetVoiceDscp sets VoiceDscp field to given value.

### HasVoiceDscp

`func (o *LanProfileSettingOpenApiVO) HasVoiceDscp() bool`

HasVoiceDscp returns a boolean if a field has been set.

### GetVoiceDscpEnable

`func (o *LanProfileSettingOpenApiVO) GetVoiceDscpEnable() bool`

GetVoiceDscpEnable returns the VoiceDscpEnable field if non-nil, zero value otherwise.

### GetVoiceDscpEnableOk

`func (o *LanProfileSettingOpenApiVO) GetVoiceDscpEnableOk() (*bool, bool)`

GetVoiceDscpEnableOk returns a tuple with the VoiceDscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpEnable

`func (o *LanProfileSettingOpenApiVO) SetVoiceDscpEnable(v bool)`

SetVoiceDscpEnable sets VoiceDscpEnable field to given value.

### HasVoiceDscpEnable

`func (o *LanProfileSettingOpenApiVO) HasVoiceDscpEnable() bool`

HasVoiceDscpEnable returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *LanProfileSettingOpenApiVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *LanProfileSettingOpenApiVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *LanProfileSettingOpenApiVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *LanProfileSettingOpenApiVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *LanProfileSettingOpenApiVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *LanProfileSettingOpenApiVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *LanProfileSettingOpenApiVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *LanProfileSettingOpenApiVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


