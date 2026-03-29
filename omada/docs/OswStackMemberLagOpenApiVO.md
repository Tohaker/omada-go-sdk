# OswStackMemberLagOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandCtrl** | Pointer to [**OswBandCtrlVO**](OswBandCtrlVO.md) |  | [optional] 
**BandWidthCtrlType** | Pointer to **int32** | BandWidthCtrlType should be a value as follows: 0: Off; 1: Rate Limit; 2: Storming Control | [optional] 
**DhcpL2RelaySettings** | Pointer to [**OswPortDhcpL2RelayVO**](OswPortDhcpL2RelayVO.md) |  | [optional] 
**Dot1pPriority** | Pointer to **int32** | Dot1p Priority | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**EeeEnable** | Pointer to **bool** | Indicates whether EEE is enabled | [optional] 
**FastLeaveEnable** | Pointer to **bool** | Indicates whether igmpSnooping fastLeave is enabled | [optional] 
**FlowControlEnable** | Pointer to **bool** | Indicates whether flow control is enabled | [optional] 
**IgmpFastLeaveEnable** | Pointer to **bool** | Indicates whether igmp fast leave is enabled | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed should be a value as follows: 0: auto; 1: 10M; 2: 100M; 3: 1000M; 4: 2.5G; 5: 10G | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | Indicates whether loopbackDetect port based is enabled | [optional] 
**LoopbackDetectVlanBasedEnable** | Pointer to **bool** | Indicates whether loopbackDetect vlan based is enabled | [optional] 
**MldFastLeaveEnable** | Pointer to **bool** | Indicates whether mld fast leave is enabled | [optional] 
**Name** | Pointer to **string** | Lag Name | [optional] 
**PortIsolationEnable** | Pointer to **bool** | Indicates whether port isolation is enabled | [optional] 
**ProfileId** | Pointer to **string** | Profile ID | [optional] 
**ProfileOverrideEnable** | **bool** | Indicates whether to enable Profile Override | 
**SpanningTreeEnable** | Pointer to **bool** | Indicates whether SpanningTree is enabled | [optional] 
**SpanningTreeSetting** | Pointer to [**SpanningTreeSettingVO**](SpanningTreeSettingVO.md) |  | [optional] 
**StandardPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | LAG Standard ports | [optional] 
**StormCtrl** | Pointer to [**OswStormCtrlVO**](OswStormCtrlVO.md) |  | [optional] 
**TrustMode** | Pointer to **int32** | TrustMode should be a value as follows: 0: Untrusted; 1: Trust 802.1p; 2: Trust DSCP | [optional] 

## Methods

### NewOswStackMemberLagOpenApiVO

`func NewOswStackMemberLagOpenApiVO(profileOverrideEnable bool, ) *OswStackMemberLagOpenApiVO`

NewOswStackMemberLagOpenApiVO instantiates a new OswStackMemberLagOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackMemberLagOpenApiVOWithDefaults

`func NewOswStackMemberLagOpenApiVOWithDefaults() *OswStackMemberLagOpenApiVO`

NewOswStackMemberLagOpenApiVOWithDefaults instantiates a new OswStackMemberLagOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandCtrl

`func (o *OswStackMemberLagOpenApiVO) GetBandCtrl() OswBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *OswStackMemberLagOpenApiVO) GetBandCtrlOk() (*OswBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *OswStackMemberLagOpenApiVO) SetBandCtrl(v OswBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *OswStackMemberLagOpenApiVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *OswStackMemberLagOpenApiVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *OswStackMemberLagOpenApiVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *OswStackMemberLagOpenApiVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *OswStackMemberLagOpenApiVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetDhcpL2RelaySettings

`func (o *OswStackMemberLagOpenApiVO) GetDhcpL2RelaySettings() OswPortDhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *OswStackMemberLagOpenApiVO) GetDhcpL2RelaySettingsOk() (*OswPortDhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *OswStackMemberLagOpenApiVO) SetDhcpL2RelaySettings(v OswPortDhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *OswStackMemberLagOpenApiVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *OswStackMemberLagOpenApiVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *OswStackMemberLagOpenApiVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *OswStackMemberLagOpenApiVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *OswStackMemberLagOpenApiVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDuplex

`func (o *OswStackMemberLagOpenApiVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswStackMemberLagOpenApiVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswStackMemberLagOpenApiVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswStackMemberLagOpenApiVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEeeEnable

`func (o *OswStackMemberLagOpenApiVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *OswStackMemberLagOpenApiVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *OswStackMemberLagOpenApiVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *OswStackMemberLagOpenApiVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *OswStackMemberLagOpenApiVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *OswStackMemberLagOpenApiVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswStackMemberLagOpenApiVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswStackMemberLagOpenApiVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswStackMemberLagOpenApiVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswStackMemberLagOpenApiVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswStackMemberLagOpenApiVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswStackMemberLagOpenApiVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswStackMemberLagOpenApiVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackDetectVlanBasedEnable

`func (o *OswStackMemberLagOpenApiVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *OswStackMemberLagOpenApiVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *OswStackMemberLagOpenApiVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *OswStackMemberLagOpenApiVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *OswStackMemberLagOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackMemberLagOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackMemberLagOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStackMemberLagOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *OswStackMemberLagOpenApiVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *OswStackMemberLagOpenApiVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *OswStackMemberLagOpenApiVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetProfileId

`func (o *OswStackMemberLagOpenApiVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OswStackMemberLagOpenApiVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OswStackMemberLagOpenApiVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *OswStackMemberLagOpenApiVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *OswStackMemberLagOpenApiVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *OswStackMemberLagOpenApiVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.


### GetSpanningTreeEnable

`func (o *OswStackMemberLagOpenApiVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *OswStackMemberLagOpenApiVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *OswStackMemberLagOpenApiVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.

### HasSpanningTreeEnable

`func (o *OswStackMemberLagOpenApiVO) HasSpanningTreeEnable() bool`

HasSpanningTreeEnable returns a boolean if a field has been set.

### GetSpanningTreeSetting

`func (o *OswStackMemberLagOpenApiVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *OswStackMemberLagOpenApiVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *OswStackMemberLagOpenApiVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *OswStackMemberLagOpenApiVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetStandardPorts

`func (o *OswStackMemberLagOpenApiVO) GetStandardPorts() []OswStandPortVO`

GetStandardPorts returns the StandardPorts field if non-nil, zero value otherwise.

### GetStandardPortsOk

`func (o *OswStackMemberLagOpenApiVO) GetStandardPortsOk() (*[]OswStandPortVO, bool)`

GetStandardPortsOk returns a tuple with the StandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPorts

`func (o *OswStackMemberLagOpenApiVO) SetStandardPorts(v []OswStandPortVO)`

SetStandardPorts sets StandardPorts field to given value.

### HasStandardPorts

`func (o *OswStackMemberLagOpenApiVO) HasStandardPorts() bool`

HasStandardPorts returns a boolean if a field has been set.

### GetStormCtrl

`func (o *OswStackMemberLagOpenApiVO) GetStormCtrl() OswStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *OswStackMemberLagOpenApiVO) GetStormCtrlOk() (*OswStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *OswStackMemberLagOpenApiVO) SetStormCtrl(v OswStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *OswStackMemberLagOpenApiVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetTrustMode

`func (o *OswStackMemberLagOpenApiVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *OswStackMemberLagOpenApiVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *OswStackMemberLagOpenApiVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *OswStackMemberLagOpenApiVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


