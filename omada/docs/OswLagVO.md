# OswLagVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site | [optional] 
**AllAggregatingPorts** | Pointer to **[]int32** | All aggregating ports of the current Switch | [optional] 
**AllMirroredPorts** | Pointer to **[]int32** | All mirrored ports on the Switch | [optional] 
**AllMirroringPorts** | Pointer to **[]int32** | All mirroring ports of the current Switch | [optional] 
**AllMlagDadPorts** | Pointer to **[]int32** |  | [optional] 
**AllMlagPeerLinkPorts** | Pointer to **[]int32** |  | [optional] 
**BandCtrl** | Pointer to [**OswBandCtrlVO**](OswBandCtrlVO.md) |  | [optional] 
**BandWidthCtrlType** | Pointer to **int32** | BandWidthCtrlType should be a value as follows: 0: Off; 1: Rate Limit; 2: Storming Control | [optional] 
**DhcpL2RelaySettings** | Pointer to [**OswPortDhcpL2RelayVO**](OswPortDhcpL2RelayVO.md) |  | [optional] 
**Disable** | Pointer to **bool** | Indicates whether to disable | [optional] 
**Dot1pPriority** | Pointer to **int32** | Dot1p Priority | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**EeeEnable** | Pointer to **bool** | Indicates whether EEE is enabled | [optional] 
**EsEnableAllProfileCanAdd** | Pointer to **bool** | Indicates whether es enable of all profile can continue to add VLANs | [optional] 
**EsQosSupport** | Pointer to **bool** | Indicates whether the ES device port supports modification of QoS configuration | [optional] 
**FastLeaveEnable** | Pointer to **bool** | Indicates whether igmpSnooping fastLeave is enabled | [optional] 
**FlowControlEnable** | Pointer to **bool** | Indicates whether flow control is enabled | [optional] 
**Id** | Pointer to **string** | ID | [optional] 
**IgmpFastLeaveEnable** | Pointer to **bool** | Indicates whether igmp fast leave is enabled | [optional] 
**IgmpSnoopingEnable** | Pointer to **bool** | Indicates whether IGMP Snooping is enabled | [optional] 
**LagId** | Pointer to **int32** | Lag ID | [optional] 
**LagStatus** | Pointer to [**OswLagStatusVO**](OswLagStatusVO.md) |  | [optional] 
**LagType** | Pointer to **int32** | Lag Type should be a value as follows: 1: STATIC; 2: LACP; 3: LACP ACTIVE; 4: LACP PASSIVE | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed should be a value as follows: 0: auto; 1: 10M; 2: 100M; 3: 1000M; 4: 2.5G; 5: 10G | [optional] 
**LocateEnable** | Pointer to **bool** | Whether locate function is enabled | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | Indicates whether loopbackDetect port based is enabled | [optional] 
**LoopbackDetectVlanBasedEnable** | Pointer to **bool** | Indicates whether loopbackDetect vlan based is enabled | [optional] 
**MlagEnable** | Pointer to **bool** |  | [optional] 
**MlagName** | Pointer to **string** |  | [optional] 
**MlagPeerAllPortsConfigInfo** | Pointer to [**OswMlagPeerAllPortsConfigInfoVO**](OswMlagPeerAllPortsConfigInfoVO.md) |  | [optional] 
**MlagPeerSetting** | Pointer to [**OswMlagPeerSettingVO**](OswMlagPeerSettingVO.md) |  | [optional] 
**MldFastLeaveEnable** | Pointer to **bool** | Indicates whether mld fast leave is enabled | [optional] 
**Name** | Pointer to **string** | Lag Name | [optional] 
**NativeBridgeVlan** | Pointer to **int32** | Native Network Bridge Vlan. | [optional] 
**NativeNetworkId** | Pointer to **string** | Native Network ID, Native Network cannot be selected from Tagged Networks or Untagged Networks. | [optional] 
**NetworkConflict** | Pointer to **bool** | Indicates whether the VLAN configuration activated on the port is inconsistent with the VLAN configuration in the Profile. | [optional] 
**NetworkMode** | Pointer to **int32** | Network Mode should be a value as follows: 0: Trunk, 1: Access | [optional] 
**NetworkTagsSetting** | Pointer to **int32** | Network Tags Setting should be a value as follows: 0: Allow All; 1: Block All; 2: Custom | [optional] 
**PortAlertEnable** | Pointer to **bool** | Indicates whether port alert is enabled | [optional] 
**PortIsolationEnable** | Pointer to **bool** | Indicates whether port isolation is enabled | [optional] 
**Ports** | Pointer to **[]int32** | Lag Ports | [optional] 
**ProfileId** | Pointer to **string** | Profile ID | [optional] 
**ProfileName** | Pointer to **string** | Lan Profile Name | [optional] 
**ProfileOverrideEnable** | Pointer to **bool** | Indicates whether to enable Profile Override | [optional] 
**QosQueueEnable** | Pointer to **bool** | Indicates whether the ES device port has enabled the Qos scheduling queue configuration | [optional] 
**QosSupport** | Pointer to **bool** | Indicates whether qos support is enabled | [optional] 
**QueueId** | Pointer to **int32** | ES Qos scheduling queue ID | [optional] 
**Resource** | Pointer to **int32** | Data Source. Resource should be a value as follows: 0: new created; 1: from template; 2: override | [optional] 
**SpanningTreeEnable** | Pointer to **bool** | Indicates whether SpanningTree is enabled | [optional] 
**SpanningTreeSetting** | Pointer to [**SpanningTreeSettingVO**](SpanningTreeSettingVO.md) |  | [optional] 
**StPorts** | Pointer to **[]string** | Lag StPorts | [optional] 
**StormCtrl** | Pointer to [**OswStormCtrlVO**](OswStormCtrlVO.md) |  | [optional] 
**SwitchId** | Pointer to **string** | Switch ID | [optional] 
**SwitchMac** | Pointer to **string** | Switch MAC Address | [optional] 
**TagBridgeVlanMap** | Pointer to **map[string][]int32** | Tag Network Bridge Vlan Map | [optional] 
**TagIds** | Pointer to **[]string** | Lag label ID List | [optional] 
**TagName** | Pointer to **string** | Lag label Name | [optional] 
**TagNetworkIds** | Pointer to **[]string** | Tag Network IDs | [optional] 
**TrustMode** | Pointer to **int32** | TrustMode should be a value as follows: 0: Untrusted; 1: Trust 802.1p; 2: Trust DSCP | [optional] 
**UntagBridgeVlanMap** | Pointer to **map[string][]int32** | Untag Network Bridge Vlan Map | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag Network IDs | [optional] 
**VoiceBridgeVlan** | Pointer to **int32** | Voice Network Bridge Vlan | [optional] 
**VoiceDscp** | Pointer to **int32** | Voice DSCP | [optional] 
**VoiceDscpEnable** | Pointer to **bool** | Voice DSCP enable status | [optional] 
**VoiceNetworkEnable** | Pointer to **bool** | Voice network enable status | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewOswLagVO

`func NewOswLagVO() *OswLagVO`

NewOswLagVO instantiates a new OswLagVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswLagVOWithDefaults

`func NewOswLagVOWithDefaults() *OswLagVO`

NewOswLagVOWithDefaults instantiates a new OswLagVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswLagVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswLagVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswLagVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswLagVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAllAggregatingPorts

`func (o *OswLagVO) GetAllAggregatingPorts() []int32`

GetAllAggregatingPorts returns the AllAggregatingPorts field if non-nil, zero value otherwise.

### GetAllAggregatingPortsOk

`func (o *OswLagVO) GetAllAggregatingPortsOk() (*[]int32, bool)`

GetAllAggregatingPortsOk returns a tuple with the AllAggregatingPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAggregatingPorts

`func (o *OswLagVO) SetAllAggregatingPorts(v []int32)`

SetAllAggregatingPorts sets AllAggregatingPorts field to given value.

### HasAllAggregatingPorts

`func (o *OswLagVO) HasAllAggregatingPorts() bool`

HasAllAggregatingPorts returns a boolean if a field has been set.

### GetAllMirroredPorts

`func (o *OswLagVO) GetAllMirroredPorts() []int32`

GetAllMirroredPorts returns the AllMirroredPorts field if non-nil, zero value otherwise.

### GetAllMirroredPortsOk

`func (o *OswLagVO) GetAllMirroredPortsOk() (*[]int32, bool)`

GetAllMirroredPortsOk returns a tuple with the AllMirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroredPorts

`func (o *OswLagVO) SetAllMirroredPorts(v []int32)`

SetAllMirroredPorts sets AllMirroredPorts field to given value.

### HasAllMirroredPorts

`func (o *OswLagVO) HasAllMirroredPorts() bool`

HasAllMirroredPorts returns a boolean if a field has been set.

### GetAllMirroringPorts

`func (o *OswLagVO) GetAllMirroringPorts() []int32`

GetAllMirroringPorts returns the AllMirroringPorts field if non-nil, zero value otherwise.

### GetAllMirroringPortsOk

`func (o *OswLagVO) GetAllMirroringPortsOk() (*[]int32, bool)`

GetAllMirroringPortsOk returns a tuple with the AllMirroringPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMirroringPorts

`func (o *OswLagVO) SetAllMirroringPorts(v []int32)`

SetAllMirroringPorts sets AllMirroringPorts field to given value.

### HasAllMirroringPorts

`func (o *OswLagVO) HasAllMirroringPorts() bool`

HasAllMirroringPorts returns a boolean if a field has been set.

### GetAllMlagDadPorts

`func (o *OswLagVO) GetAllMlagDadPorts() []int32`

GetAllMlagDadPorts returns the AllMlagDadPorts field if non-nil, zero value otherwise.

### GetAllMlagDadPortsOk

`func (o *OswLagVO) GetAllMlagDadPortsOk() (*[]int32, bool)`

GetAllMlagDadPortsOk returns a tuple with the AllMlagDadPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMlagDadPorts

`func (o *OswLagVO) SetAllMlagDadPorts(v []int32)`

SetAllMlagDadPorts sets AllMlagDadPorts field to given value.

### HasAllMlagDadPorts

`func (o *OswLagVO) HasAllMlagDadPorts() bool`

HasAllMlagDadPorts returns a boolean if a field has been set.

### GetAllMlagPeerLinkPorts

`func (o *OswLagVO) GetAllMlagPeerLinkPorts() []int32`

GetAllMlagPeerLinkPorts returns the AllMlagPeerLinkPorts field if non-nil, zero value otherwise.

### GetAllMlagPeerLinkPortsOk

`func (o *OswLagVO) GetAllMlagPeerLinkPortsOk() (*[]int32, bool)`

GetAllMlagPeerLinkPortsOk returns a tuple with the AllMlagPeerLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMlagPeerLinkPorts

`func (o *OswLagVO) SetAllMlagPeerLinkPorts(v []int32)`

SetAllMlagPeerLinkPorts sets AllMlagPeerLinkPorts field to given value.

### HasAllMlagPeerLinkPorts

`func (o *OswLagVO) HasAllMlagPeerLinkPorts() bool`

HasAllMlagPeerLinkPorts returns a boolean if a field has been set.

### GetBandCtrl

`func (o *OswLagVO) GetBandCtrl() OswBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *OswLagVO) GetBandCtrlOk() (*OswBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *OswLagVO) SetBandCtrl(v OswBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *OswLagVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *OswLagVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *OswLagVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *OswLagVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *OswLagVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetDhcpL2RelaySettings

`func (o *OswLagVO) GetDhcpL2RelaySettings() OswPortDhcpL2RelayVO`

GetDhcpL2RelaySettings returns the DhcpL2RelaySettings field if non-nil, zero value otherwise.

### GetDhcpL2RelaySettingsOk

`func (o *OswLagVO) GetDhcpL2RelaySettingsOk() (*OswPortDhcpL2RelayVO, bool)`

GetDhcpL2RelaySettingsOk returns a tuple with the DhcpL2RelaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySettings

`func (o *OswLagVO) SetDhcpL2RelaySettings(v OswPortDhcpL2RelayVO)`

SetDhcpL2RelaySettings sets DhcpL2RelaySettings field to given value.

### HasDhcpL2RelaySettings

`func (o *OswLagVO) HasDhcpL2RelaySettings() bool`

HasDhcpL2RelaySettings returns a boolean if a field has been set.

### GetDisable

`func (o *OswLagVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswLagVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswLagVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswLagVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDot1pPriority

`func (o *OswLagVO) GetDot1pPriority() int32`

GetDot1pPriority returns the Dot1pPriority field if non-nil, zero value otherwise.

### GetDot1pPriorityOk

`func (o *OswLagVO) GetDot1pPriorityOk() (*int32, bool)`

GetDot1pPriorityOk returns a tuple with the Dot1pPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1pPriority

`func (o *OswLagVO) SetDot1pPriority(v int32)`

SetDot1pPriority sets Dot1pPriority field to given value.

### HasDot1pPriority

`func (o *OswLagVO) HasDot1pPriority() bool`

HasDot1pPriority returns a boolean if a field has been set.

### GetDuplex

`func (o *OswLagVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswLagVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswLagVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswLagVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEeeEnable

`func (o *OswLagVO) GetEeeEnable() bool`

GetEeeEnable returns the EeeEnable field if non-nil, zero value otherwise.

### GetEeeEnableOk

`func (o *OswLagVO) GetEeeEnableOk() (*bool, bool)`

GetEeeEnableOk returns a tuple with the EeeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeEnable

`func (o *OswLagVO) SetEeeEnable(v bool)`

SetEeeEnable sets EeeEnable field to given value.

### HasEeeEnable

`func (o *OswLagVO) HasEeeEnable() bool`

HasEeeEnable returns a boolean if a field has been set.

### GetEsEnableAllProfileCanAdd

`func (o *OswLagVO) GetEsEnableAllProfileCanAdd() bool`

GetEsEnableAllProfileCanAdd returns the EsEnableAllProfileCanAdd field if non-nil, zero value otherwise.

### GetEsEnableAllProfileCanAddOk

`func (o *OswLagVO) GetEsEnableAllProfileCanAddOk() (*bool, bool)`

GetEsEnableAllProfileCanAddOk returns a tuple with the EsEnableAllProfileCanAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsEnableAllProfileCanAdd

`func (o *OswLagVO) SetEsEnableAllProfileCanAdd(v bool)`

SetEsEnableAllProfileCanAdd sets EsEnableAllProfileCanAdd field to given value.

### HasEsEnableAllProfileCanAdd

`func (o *OswLagVO) HasEsEnableAllProfileCanAdd() bool`

HasEsEnableAllProfileCanAdd returns a boolean if a field has been set.

### GetEsQosSupport

`func (o *OswLagVO) GetEsQosSupport() bool`

GetEsQosSupport returns the EsQosSupport field if non-nil, zero value otherwise.

### GetEsQosSupportOk

`func (o *OswLagVO) GetEsQosSupportOk() (*bool, bool)`

GetEsQosSupportOk returns a tuple with the EsQosSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsQosSupport

`func (o *OswLagVO) SetEsQosSupport(v bool)`

SetEsQosSupport sets EsQosSupport field to given value.

### HasEsQosSupport

`func (o *OswLagVO) HasEsQosSupport() bool`

HasEsQosSupport returns a boolean if a field has been set.

### GetFastLeaveEnable

`func (o *OswLagVO) GetFastLeaveEnable() bool`

GetFastLeaveEnable returns the FastLeaveEnable field if non-nil, zero value otherwise.

### GetFastLeaveEnableOk

`func (o *OswLagVO) GetFastLeaveEnableOk() (*bool, bool)`

GetFastLeaveEnableOk returns a tuple with the FastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastLeaveEnable

`func (o *OswLagVO) SetFastLeaveEnable(v bool)`

SetFastLeaveEnable sets FastLeaveEnable field to given value.

### HasFastLeaveEnable

`func (o *OswLagVO) HasFastLeaveEnable() bool`

HasFastLeaveEnable returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *OswLagVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *OswLagVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *OswLagVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *OswLagVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetId

`func (o *OswLagVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswLagVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswLagVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswLagVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgmpFastLeaveEnable

`func (o *OswLagVO) GetIgmpFastLeaveEnable() bool`

GetIgmpFastLeaveEnable returns the IgmpFastLeaveEnable field if non-nil, zero value otherwise.

### GetIgmpFastLeaveEnableOk

`func (o *OswLagVO) GetIgmpFastLeaveEnableOk() (*bool, bool)`

GetIgmpFastLeaveEnableOk returns a tuple with the IgmpFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveEnable

`func (o *OswLagVO) SetIgmpFastLeaveEnable(v bool)`

SetIgmpFastLeaveEnable sets IgmpFastLeaveEnable field to given value.

### HasIgmpFastLeaveEnable

`func (o *OswLagVO) HasIgmpFastLeaveEnable() bool`

HasIgmpFastLeaveEnable returns a boolean if a field has been set.

### GetIgmpSnoopingEnable

`func (o *OswLagVO) GetIgmpSnoopingEnable() bool`

GetIgmpSnoopingEnable returns the IgmpSnoopingEnable field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnableOk

`func (o *OswLagVO) GetIgmpSnoopingEnableOk() (*bool, bool)`

GetIgmpSnoopingEnableOk returns a tuple with the IgmpSnoopingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnable

`func (o *OswLagVO) SetIgmpSnoopingEnable(v bool)`

SetIgmpSnoopingEnable sets IgmpSnoopingEnable field to given value.

### HasIgmpSnoopingEnable

`func (o *OswLagVO) HasIgmpSnoopingEnable() bool`

HasIgmpSnoopingEnable returns a boolean if a field has been set.

### GetLagId

`func (o *OswLagVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswLagVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswLagVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OswLagVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLagStatus

`func (o *OswLagVO) GetLagStatus() OswLagStatusVO`

GetLagStatus returns the LagStatus field if non-nil, zero value otherwise.

### GetLagStatusOk

`func (o *OswLagVO) GetLagStatusOk() (*OswLagStatusVO, bool)`

GetLagStatusOk returns a tuple with the LagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagStatus

`func (o *OswLagVO) SetLagStatus(v OswLagStatusVO)`

SetLagStatus sets LagStatus field to given value.

### HasLagStatus

`func (o *OswLagVO) HasLagStatus() bool`

HasLagStatus returns a boolean if a field has been set.

### GetLagType

`func (o *OswLagVO) GetLagType() int32`

GetLagType returns the LagType field if non-nil, zero value otherwise.

### GetLagTypeOk

`func (o *OswLagVO) GetLagTypeOk() (*int32, bool)`

GetLagTypeOk returns a tuple with the LagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagType

`func (o *OswLagVO) SetLagType(v int32)`

SetLagType sets LagType field to given value.

### HasLagType

`func (o *OswLagVO) HasLagType() bool`

HasLagType returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswLagVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswLagVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswLagVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswLagVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswLagVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswLagVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswLagVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswLagVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswLagVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswLagVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswLagVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswLagVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackDetectVlanBasedEnable

`func (o *OswLagVO) GetLoopbackDetectVlanBasedEnable() bool`

GetLoopbackDetectVlanBasedEnable returns the LoopbackDetectVlanBasedEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectVlanBasedEnableOk

`func (o *OswLagVO) GetLoopbackDetectVlanBasedEnableOk() (*bool, bool)`

GetLoopbackDetectVlanBasedEnableOk returns a tuple with the LoopbackDetectVlanBasedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectVlanBasedEnable

`func (o *OswLagVO) SetLoopbackDetectVlanBasedEnable(v bool)`

SetLoopbackDetectVlanBasedEnable sets LoopbackDetectVlanBasedEnable field to given value.

### HasLoopbackDetectVlanBasedEnable

`func (o *OswLagVO) HasLoopbackDetectVlanBasedEnable() bool`

HasLoopbackDetectVlanBasedEnable returns a boolean if a field has been set.

### GetMlagEnable

`func (o *OswLagVO) GetMlagEnable() bool`

GetMlagEnable returns the MlagEnable field if non-nil, zero value otherwise.

### GetMlagEnableOk

`func (o *OswLagVO) GetMlagEnableOk() (*bool, bool)`

GetMlagEnableOk returns a tuple with the MlagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagEnable

`func (o *OswLagVO) SetMlagEnable(v bool)`

SetMlagEnable sets MlagEnable field to given value.

### HasMlagEnable

`func (o *OswLagVO) HasMlagEnable() bool`

HasMlagEnable returns a boolean if a field has been set.

### GetMlagName

`func (o *OswLagVO) GetMlagName() string`

GetMlagName returns the MlagName field if non-nil, zero value otherwise.

### GetMlagNameOk

`func (o *OswLagVO) GetMlagNameOk() (*string, bool)`

GetMlagNameOk returns a tuple with the MlagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagName

`func (o *OswLagVO) SetMlagName(v string)`

SetMlagName sets MlagName field to given value.

### HasMlagName

`func (o *OswLagVO) HasMlagName() bool`

HasMlagName returns a boolean if a field has been set.

### GetMlagPeerAllPortsConfigInfo

`func (o *OswLagVO) GetMlagPeerAllPortsConfigInfo() OswMlagPeerAllPortsConfigInfoVO`

GetMlagPeerAllPortsConfigInfo returns the MlagPeerAllPortsConfigInfo field if non-nil, zero value otherwise.

### GetMlagPeerAllPortsConfigInfoOk

`func (o *OswLagVO) GetMlagPeerAllPortsConfigInfoOk() (*OswMlagPeerAllPortsConfigInfoVO, bool)`

GetMlagPeerAllPortsConfigInfoOk returns a tuple with the MlagPeerAllPortsConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerAllPortsConfigInfo

`func (o *OswLagVO) SetMlagPeerAllPortsConfigInfo(v OswMlagPeerAllPortsConfigInfoVO)`

SetMlagPeerAllPortsConfigInfo sets MlagPeerAllPortsConfigInfo field to given value.

### HasMlagPeerAllPortsConfigInfo

`func (o *OswLagVO) HasMlagPeerAllPortsConfigInfo() bool`

HasMlagPeerAllPortsConfigInfo returns a boolean if a field has been set.

### GetMlagPeerSetting

`func (o *OswLagVO) GetMlagPeerSetting() OswMlagPeerSettingVO`

GetMlagPeerSetting returns the MlagPeerSetting field if non-nil, zero value otherwise.

### GetMlagPeerSettingOk

`func (o *OswLagVO) GetMlagPeerSettingOk() (*OswMlagPeerSettingVO, bool)`

GetMlagPeerSettingOk returns a tuple with the MlagPeerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerSetting

`func (o *OswLagVO) SetMlagPeerSetting(v OswMlagPeerSettingVO)`

SetMlagPeerSetting sets MlagPeerSetting field to given value.

### HasMlagPeerSetting

`func (o *OswLagVO) HasMlagPeerSetting() bool`

HasMlagPeerSetting returns a boolean if a field has been set.

### GetMldFastLeaveEnable

`func (o *OswLagVO) GetMldFastLeaveEnable() bool`

GetMldFastLeaveEnable returns the MldFastLeaveEnable field if non-nil, zero value otherwise.

### GetMldFastLeaveEnableOk

`func (o *OswLagVO) GetMldFastLeaveEnableOk() (*bool, bool)`

GetMldFastLeaveEnableOk returns a tuple with the MldFastLeaveEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldFastLeaveEnable

`func (o *OswLagVO) SetMldFastLeaveEnable(v bool)`

SetMldFastLeaveEnable sets MldFastLeaveEnable field to given value.

### HasMldFastLeaveEnable

`func (o *OswLagVO) HasMldFastLeaveEnable() bool`

HasMldFastLeaveEnable returns a boolean if a field has been set.

### GetName

`func (o *OswLagVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswLagVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswLagVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswLagVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *OswLagVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *OswLagVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *OswLagVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *OswLagVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *OswLagVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *OswLagVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *OswLagVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *OswLagVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkConflict

`func (o *OswLagVO) GetNetworkConflict() bool`

GetNetworkConflict returns the NetworkConflict field if non-nil, zero value otherwise.

### GetNetworkConflictOk

`func (o *OswLagVO) GetNetworkConflictOk() (*bool, bool)`

GetNetworkConflictOk returns a tuple with the NetworkConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConflict

`func (o *OswLagVO) SetNetworkConflict(v bool)`

SetNetworkConflict sets NetworkConflict field to given value.

### HasNetworkConflict

`func (o *OswLagVO) HasNetworkConflict() bool`

HasNetworkConflict returns a boolean if a field has been set.

### GetNetworkMode

`func (o *OswLagVO) GetNetworkMode() int32`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *OswLagVO) GetNetworkModeOk() (*int32, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *OswLagVO) SetNetworkMode(v int32)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *OswLagVO) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *OswLagVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *OswLagVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *OswLagVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *OswLagVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetPortAlertEnable

`func (o *OswLagVO) GetPortAlertEnable() bool`

GetPortAlertEnable returns the PortAlertEnable field if non-nil, zero value otherwise.

### GetPortAlertEnableOk

`func (o *OswLagVO) GetPortAlertEnableOk() (*bool, bool)`

GetPortAlertEnableOk returns a tuple with the PortAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAlertEnable

`func (o *OswLagVO) SetPortAlertEnable(v bool)`

SetPortAlertEnable sets PortAlertEnable field to given value.

### HasPortAlertEnable

`func (o *OswLagVO) HasPortAlertEnable() bool`

HasPortAlertEnable returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *OswLagVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *OswLagVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *OswLagVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *OswLagVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetPorts

`func (o *OswLagVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswLagVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswLagVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswLagVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProfileId

`func (o *OswLagVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OswLagVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OswLagVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *OswLagVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *OswLagVO) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *OswLagVO) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *OswLagVO) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *OswLagVO) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *OswLagVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *OswLagVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *OswLagVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *OswLagVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetQosQueueEnable

`func (o *OswLagVO) GetQosQueueEnable() bool`

GetQosQueueEnable returns the QosQueueEnable field if non-nil, zero value otherwise.

### GetQosQueueEnableOk

`func (o *OswLagVO) GetQosQueueEnableOk() (*bool, bool)`

GetQosQueueEnableOk returns a tuple with the QosQueueEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosQueueEnable

`func (o *OswLagVO) SetQosQueueEnable(v bool)`

SetQosQueueEnable sets QosQueueEnable field to given value.

### HasQosQueueEnable

`func (o *OswLagVO) HasQosQueueEnable() bool`

HasQosQueueEnable returns a boolean if a field has been set.

### GetQosSupport

`func (o *OswLagVO) GetQosSupport() bool`

GetQosSupport returns the QosSupport field if non-nil, zero value otherwise.

### GetQosSupportOk

`func (o *OswLagVO) GetQosSupportOk() (*bool, bool)`

GetQosSupportOk returns a tuple with the QosSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSupport

`func (o *OswLagVO) SetQosSupport(v bool)`

SetQosSupport sets QosSupport field to given value.

### HasQosSupport

`func (o *OswLagVO) HasQosSupport() bool`

HasQosSupport returns a boolean if a field has been set.

### GetQueueId

`func (o *OswLagVO) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *OswLagVO) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *OswLagVO) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *OswLagVO) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetResource

`func (o *OswLagVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswLagVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswLagVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswLagVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSpanningTreeEnable

`func (o *OswLagVO) GetSpanningTreeEnable() bool`

GetSpanningTreeEnable returns the SpanningTreeEnable field if non-nil, zero value otherwise.

### GetSpanningTreeEnableOk

`func (o *OswLagVO) GetSpanningTreeEnableOk() (*bool, bool)`

GetSpanningTreeEnableOk returns a tuple with the SpanningTreeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeEnable

`func (o *OswLagVO) SetSpanningTreeEnable(v bool)`

SetSpanningTreeEnable sets SpanningTreeEnable field to given value.

### HasSpanningTreeEnable

`func (o *OswLagVO) HasSpanningTreeEnable() bool`

HasSpanningTreeEnable returns a boolean if a field has been set.

### GetSpanningTreeSetting

`func (o *OswLagVO) GetSpanningTreeSetting() SpanningTreeSettingVO`

GetSpanningTreeSetting returns the SpanningTreeSetting field if non-nil, zero value otherwise.

### GetSpanningTreeSettingOk

`func (o *OswLagVO) GetSpanningTreeSettingOk() (*SpanningTreeSettingVO, bool)`

GetSpanningTreeSettingOk returns a tuple with the SpanningTreeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeSetting

`func (o *OswLagVO) SetSpanningTreeSetting(v SpanningTreeSettingVO)`

SetSpanningTreeSetting sets SpanningTreeSetting field to given value.

### HasSpanningTreeSetting

`func (o *OswLagVO) HasSpanningTreeSetting() bool`

HasSpanningTreeSetting returns a boolean if a field has been set.

### GetStPorts

`func (o *OswLagVO) GetStPorts() []string`

GetStPorts returns the StPorts field if non-nil, zero value otherwise.

### GetStPortsOk

`func (o *OswLagVO) GetStPortsOk() (*[]string, bool)`

GetStPortsOk returns a tuple with the StPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStPorts

`func (o *OswLagVO) SetStPorts(v []string)`

SetStPorts sets StPorts field to given value.

### HasStPorts

`func (o *OswLagVO) HasStPorts() bool`

HasStPorts returns a boolean if a field has been set.

### GetStormCtrl

`func (o *OswLagVO) GetStormCtrl() OswStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *OswLagVO) GetStormCtrlOk() (*OswStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *OswLagVO) SetStormCtrl(v OswStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *OswLagVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetSwitchId

`func (o *OswLagVO) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *OswLagVO) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *OswLagVO) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *OswLagVO) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetSwitchMac

`func (o *OswLagVO) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *OswLagVO) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *OswLagVO) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.

### HasSwitchMac

`func (o *OswLagVO) HasSwitchMac() bool`

HasSwitchMac returns a boolean if a field has been set.

### GetTagBridgeVlanMap

`func (o *OswLagVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *OswLagVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *OswLagVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *OswLagVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagIds

`func (o *OswLagVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswLagVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswLagVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswLagVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagName

`func (o *OswLagVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *OswLagVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *OswLagVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *OswLagVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *OswLagVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *OswLagVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *OswLagVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *OswLagVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetTrustMode

`func (o *OswLagVO) GetTrustMode() int32`

GetTrustMode returns the TrustMode field if non-nil, zero value otherwise.

### GetTrustModeOk

`func (o *OswLagVO) GetTrustModeOk() (*int32, bool)`

GetTrustModeOk returns a tuple with the TrustMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustMode

`func (o *OswLagVO) SetTrustMode(v int32)`

SetTrustMode sets TrustMode field to given value.

### HasTrustMode

`func (o *OswLagVO) HasTrustMode() bool`

HasTrustMode returns a boolean if a field has been set.

### GetUntagBridgeVlanMap

`func (o *OswLagVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *OswLagVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *OswLagVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *OswLagVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *OswLagVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *OswLagVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *OswLagVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *OswLagVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *OswLagVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *OswLagVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *OswLagVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *OswLagVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceDscp

`func (o *OswLagVO) GetVoiceDscp() int32`

GetVoiceDscp returns the VoiceDscp field if non-nil, zero value otherwise.

### GetVoiceDscpOk

`func (o *OswLagVO) GetVoiceDscpOk() (*int32, bool)`

GetVoiceDscpOk returns a tuple with the VoiceDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscp

`func (o *OswLagVO) SetVoiceDscp(v int32)`

SetVoiceDscp sets VoiceDscp field to given value.

### HasVoiceDscp

`func (o *OswLagVO) HasVoiceDscp() bool`

HasVoiceDscp returns a boolean if a field has been set.

### GetVoiceDscpEnable

`func (o *OswLagVO) GetVoiceDscpEnable() bool`

GetVoiceDscpEnable returns the VoiceDscpEnable field if non-nil, zero value otherwise.

### GetVoiceDscpEnableOk

`func (o *OswLagVO) GetVoiceDscpEnableOk() (*bool, bool)`

GetVoiceDscpEnableOk returns a tuple with the VoiceDscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpEnable

`func (o *OswLagVO) SetVoiceDscpEnable(v bool)`

SetVoiceDscpEnable sets VoiceDscpEnable field to given value.

### HasVoiceDscpEnable

`func (o *OswLagVO) HasVoiceDscpEnable() bool`

HasVoiceDscpEnable returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *OswLagVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *OswLagVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *OswLagVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *OswLagVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *OswLagVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *OswLagVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *OswLagVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *OswLagVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


