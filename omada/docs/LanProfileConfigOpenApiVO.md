# LanProfileConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandCtrl** | Pointer to [**BandCtrlVO**](BandCtrlVO.md) |  | [optional] 
**BandWidthCtrlType** | **int32** | BandWidthCtrlType should be a value as follows: 0: off, 1: rate limit, 2: storming control | 
**DhcpL2RelaySettings** | Pointer to [**DhcpL2RelayVO**](DhcpL2RelayVO.md) |  | [optional] 
**Dot1pPriority** | Pointer to **int32** | 802.1p Priority | [optional] 
**Dot1x** | **int32** | Dot1x should be a value as follows: 0: force unauthorized, 1: force authorized, 2:auto | 
**EeeEnable** | Pointer to **bool** | EEE enable status | [optional] 
**FastLeaveEnable** | Pointer to **bool** | IGMP Snooping fast leave enable status | [optional] 
**FlowControlEnable** | Pointer to **bool** | FlowControl enable status | [optional] 
**IgmpFastLeaveEnable** | Pointer to **bool** | Indicates whether igmp fast leave is enabled | [optional] 
**LldpMedEnable** | **bool** | LLDP-MED enable status | 
**LoopbackDetectEnable** | **bool** | LoopbackDetect enable status | 
**LoopbackDetectVlanBasedEnable** | Pointer to **bool** | LoopbackDetectVLANBased enable status | [optional] 
**MldFastLeaveEnable** | Pointer to **bool** | Indicates whether mld fast leave is enabled | [optional] 
**Name** | **string** | Name should contain 1 to 128 characters. | 
**NativeNetworkId** | **string** | Native network ID, Native Network cannot be selected from Tagged Networks or Untagged Networks. | 
**Poe** | **int32** | PoE should be a value as follows: 0: on, 1: off, 2: \&quot;do not modify\&quot; | 
**PortIsolationEnable** | **bool** | Port-isolation enable status | 
**SpanningTreeEnable** | **bool** | SpanningTree enable status | 
**SpanningTreeSetting** | Pointer to [**SpanningTreeSettingVO**](SpanningTreeSettingVO.md) |  | [optional] 
**StormCtrl** | Pointer to [**StormCtrlVO**](StormCtrlVO.md) |  | [optional] 
**TagNetworkIds** | Pointer to **[]string** | Tag network IDs | [optional] 
**TrustMode** | Pointer to **int32** | Trust mode | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag network IDs | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewLanProfileConfigOpenApiVO

`func NewLanProfileConfigOpenApiVO(bandWidthCtrlType int32, dot1x int32, lldpMedEnable bool, loopbackDetectEnable bool, name string, nativeNetworkId string, poe int32, portIsolationEnable bool, spanningTreeEnable bool, ) *LanProfileConfigOpenApiVO`

NewLanProfileConfigOpenApiVO instantiates a new LanProfileConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanProfileConfigOpenApiVOWithDefaults

`func NewLanProfileConfigOpenApiVOWithDefaults() *LanProfileConfigOpenApiVO`

NewLanProfileConfigOpenApiVOWithDefaults instantiates a new LanProfileConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandCtrl

`func (o *LanProfileConfigOpenApiVO) GetBandCtrl() BandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *LanProfileConfigOpenApiVO) GetBandCtrlOk() (*BandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *LanProfileConfigOpenApiVO) SetBandCtrl(v BandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *LanProfileConfigOpenApiVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *LanProfileConfigOpenApiVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *LanProfileConfigOpenApiVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *LanProfileConfigOpenApiVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.


### GetDhcpL2RelaySettings

`func (o *LanProfileConfigOpenApiVO) GetDhcpL2RelaySettings() DhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *LanProfileConfigOpenApiVO) GetDhcpL2RelaySettingsOk() (*DhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *LanProfileConfigOpenApiVO) SetDhcpL2RelaySettings(v DhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *LanProfileConfigOpenApiVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *LanProfileConfigOpenApiVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *LanProfileConfigOpenApiVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *LanProfileConfigOpenApiVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *LanProfileConfigOpenApiVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDot1x

`func (o *LanProfileConfigOpenApiVO) GetDot1x() int32`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *LanProfileConfigOpenApiVO) GetDot1xOk() (*int32, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *LanProfileConfigOpenApiVO) SetDot1x(v int32)`

SetDot1x sets Dot1x field to given value.


### GetEeeEnable

`func (o *LanProfileConfigOpenApiVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *LanProfileConfigOpenApiVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *LanProfileConfigOpenApiVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *LanProfileConfigOpenApiVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *LanProfileConfigOpenApiVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *LanProfileConfigOpenApiVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *LanProfileConfigOpenApiVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *LanProfileConfigOpenApiVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *LanProfileConfigOpenApiVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *LanProfileConfigOpenApiVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetLldpMedEnable

`func (o *LanProfileConfigOpenApiVO) GetLldpMedEnable() bool`

GetLldpMedEnable returns the LldpMedEnable field if non-nil, zero value otherwise.

### GetLldpMedEnableOk

`func (o *LanProfileConfigOpenApiVO) GetLldpMedEnableOk() (*bool, bool)`

GetLldpMedEnableOk returns a tuple with the LldpMedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMedEnable

`func (o *LanProfileConfigOpenApiVO) SetLldpMedEnable(v bool)`

SetLldpMedEnable sets LldpMedEnable field to given value.


### GetLoopbackDetectEnable

`func (o *LanProfileConfigOpenApiVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *LanProfileConfigOpenApiVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *LanProfileConfigOpenApiVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.


### GetLoopbackDetectVlanBasedEnable

`func (o *LanProfileConfigOpenApiVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *LanProfileConfigOpenApiVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *LanProfileConfigOpenApiVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *LanProfileConfigOpenApiVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *LanProfileConfigOpenApiVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *LanProfileConfigOpenApiVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *LanProfileConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanProfileConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanProfileConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNativeNetworkId

`func (o *LanProfileConfigOpenApiVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *LanProfileConfigOpenApiVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *LanProfileConfigOpenApiVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.


### GetPoe

`func (o *LanProfileConfigOpenApiVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *LanProfileConfigOpenApiVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *LanProfileConfigOpenApiVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.


### GetPortIsolationEnable

`func (o *LanProfileConfigOpenApiVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *LanProfileConfigOpenApiVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *LanProfileConfigOpenApiVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.


### GetSpanningTreeEnable

`func (o *LanProfileConfigOpenApiVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *LanProfileConfigOpenApiVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *LanProfileConfigOpenApiVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.


### GetSpanningTreeSetting

`func (o *LanProfileConfigOpenApiVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *LanProfileConfigOpenApiVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *LanProfileConfigOpenApiVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *LanProfileConfigOpenApiVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetStormCtrl

`func (o *LanProfileConfigOpenApiVO) GetStormCtrl() StormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *LanProfileConfigOpenApiVO) GetStormCtrlOk() (*StormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *LanProfileConfigOpenApiVO) SetStormCtrl(v StormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *LanProfileConfigOpenApiVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *LanProfileConfigOpenApiVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *LanProfileConfigOpenApiVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *LanProfileConfigOpenApiVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *LanProfileConfigOpenApiVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTrustMode

`func (o *LanProfileConfigOpenApiVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *LanProfileConfigOpenApiVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *LanProfileConfigOpenApiVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *LanProfileConfigOpenApiVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *LanProfileConfigOpenApiVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *LanProfileConfigOpenApiVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *LanProfileConfigOpenApiVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *LanProfileConfigOpenApiVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *LanProfileConfigOpenApiVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *LanProfileConfigOpenApiVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *LanProfileConfigOpenApiVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *LanProfileConfigOpenApiVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


