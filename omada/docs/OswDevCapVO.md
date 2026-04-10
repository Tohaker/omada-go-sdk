# OswDevCapVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpTableSupport** | Pointer to **bool** |  | [optional] 
**CableTestSupport** | Pointer to **bool** | Cable Test Support | [optional] 
**CliSupport** | Pointer to **bool** | CLI Support | [optional] 
**CustomDhcpOptionSupport** | Pointer to **bool** |  | [optional] 
**DdmSupport** | Pointer to **bool** |  | [optional] 
**DhcpL2RelaySupport** | Pointer to **bool** | Dhcp L2 Relay Support | [optional] 
**DhcpRangeSupport** | Pointer to **bool** |  | [optional] 
**DhcpRelayIfNum** | Pointer to **int32** |  | [optional] 
**DhcpReservationSupport** | Pointer to **bool** |  | [optional] 
**DhcpServerPoolNum** | Pointer to **int32** |  | [optional] 
**DnsLoopUpSupport** | Pointer to **bool** |  | [optional] 
**DomainPingSupport** | Pointer to **bool** |  | [optional] 
**DomainTraceRouteSupport** | Pointer to **bool** |  | [optional] 
**EeeSupport** | Pointer to **bool** | eee Support | [optional] 
**EsQosSupport** | Pointer to **bool** | Agile Series Switch QoS Support | [optional] 
**FecSupport** | Pointer to **bool** | FEC Support | [optional] 
**FlowControlSupport** | Pointer to **bool** | FlowControl Support | [optional] 
**GponSupport** | Pointer to **bool** |  | [optional] 
**IgmpFastLeaveSupport** | Pointer to **bool** | Igmp FastLeve Support | [optional] 
**IpcDetectSupport** | Pointer to **bool** |  | [optional] 
**Ipv4StaticRouteNum** | Pointer to **int32** |  | [optional] 
**Ipv6StaticRouteNum** | Pointer to **int32** |  | [optional] 
**JumboOddSupport** | Pointer to **bool** | Jumbo Odd Support | [optional] 
**JumboSupport** | Pointer to **bool** | Jumbo Support | [optional] 
**LagCap** | Pointer to [**LagCapVO**](LagCapVO.md) |  | [optional] 
**LocatePortSupport** | Pointer to **bool** | Locate port Support | [optional] 
**LoopbackInterfaceSupport** | Pointer to **bool** | Loopback Interface Support | [optional] 
**LoopbackVlanBasedSupport** | Pointer to **bool** | Loopback Detect Vlan Based Support | [optional] 
**MaxLagMember** | Pointer to **int32** | Max Lag Member | [optional] 
**MaxLagNum** | Pointer to **int32** | Max Lag Num | [optional] 
**MaxMirrorGroup** | Pointer to **int32** | Max Mirror Group | [optional] 
**MaxMirroredPort** | Pointer to **int32** | Max Mirrored Port | [optional] 
**MaxRelayServerNum** | Pointer to **int32** | Max DHCP relay server num | [optional] 
**MaxStackGroupNumber** | Pointer to **int32** | Max Stack Group Number | [optional] 
**MaxStackUnitNumber** | Pointer to **int32** | Max Stack Unit Number | [optional] 
**MirrorSupport** | Pointer to **bool** |  | [optional] 
**MlagGroupId** | Pointer to **int32** |  | [optional] 
**MlagVersion** | Pointer to **string** |  | [optional] 
**MstpInsNum** | Pointer to **int32** | Mstp Instance Num | [optional] 
**NeedFullSync** | Pointer to **bool** | Need full sync | [optional] 
**NetworkCheckSupport** | Pointer to **bool** |  | [optional] 
**OspfDeadIntervalSupport** | Pointer to **bool** | OSPF Dead Interval Support  | [optional] 
**OspfSupport** | Pointer to **bool** | OSPF Support | [optional] 
**OuiBasedVlanSupport** | Pointer to **bool** | Oui Based Vlan Support | [optional] 
**PacketCaptureSupport** | Pointer to **bool** |  | [optional] 
**PingSupport** | Pointer to **bool** |  | [optional] 
**PoePortNum** | Pointer to **int32** | Poe Port Num | [optional] 
**PoeSupport** | Pointer to **bool** | Poe Support | [optional] 
**QosForVlanSupport** | Pointer to **bool** | Qos for Vlan Support | [optional] 
**QosSupport** | Pointer to **bool** | QoS Support | [optional] 
**SfpBeginNum** | Pointer to **int32** | SFP Begin Num | [optional] 
**SfpNum** | Pointer to **int32** | SFP Num | [optional] 
**StackPortCap** | Pointer to **map[string][]string** | Stack Port Capability | [optional] 
**StaticRouteNum** | Pointer to **int32** |  | [optional] 
**StkVer** | Pointer to **string** | Stack Version | [optional] 
**StkableGroupId** | Pointer to **int32** | Stackable Group ID | [optional] 
**StormRateModeSupport** | Pointer to **bool** | StormControl RateMode Support | [optional] 
**StpExtendSupport** | Pointer to **bool** | STP Extend Support | [optional] 
**SupportArpDetect** | Pointer to **bool** | Whether support arp detect | [optional] 
**SupportBt** | Pointer to **bool** | Support Poe Bt | [optional] 
**SupportClearCounters** | Pointer to **bool** | Whether the device supports clear counters | [optional] 
**SupportConfigSync** | Pointer to **bool** | Whether the device supports config sync | [optional] 
**SupportDhcpSnoop** | Pointer to **bool** | Whether support dhcp snoop | [optional] 
**SupportEsHealth** | Pointer to **bool** |  | [optional] 
**SupportEsMulticast** | Pointer to **bool** | Whether Agile Series Switch support multicast | [optional] 
**SupportFan** | Pointer to **bool** | Support Fan | [optional] 
**SupportGetDhcpClientTable** | Pointer to **bool** |  | [optional] 
**SupportGetDhcpServerInfo** | Pointer to **bool** |  | [optional] 
**SupportGetOspfNeighborTable** | Pointer to **bool** |  | [optional] 
**SupportGetVlanIfInfo** | Pointer to **bool** |  | [optional] 
**SupportImpb** | Pointer to **bool** | Whether support impb | [optional] 
**SupportMad** | Pointer to **bool** |  | [optional] 
**SupportMlag** | Pointer to **bool** |  | [optional] 
**SupportMulticast** | Pointer to **bool** | Whether support multicast | [optional] 
**SupportRelayMultiServer** | Pointer to **bool** | Whether the device supports DHCP relay multi Server | [optional] 
**SupportRunningConfig** | Pointer to **bool** | Whether the device supports showing running config | [optional] 
**SupportSdm** | Pointer to **bool** |  | [optional] 
**SupportVrf** | Pointer to **bool** |  | [optional] 
**TerminalSupport** | Pointer to **bool** |  | [optional] 
**TracerouteSupport** | Pointer to **bool** |  | [optional] 
**UplinkSupport** | Pointer to **bool** |  | [optional] 
**VlanIfNum** | Pointer to **int32** | Vlan Interface Num | [optional] 
**VoiceDscpSupport** | Pointer to **bool** |  | [optional] 
**VrfNum** | Pointer to **int32** |  | [optional] 
**VrrpSupport** | Pointer to **bool** | Vrrp Support | [optional] 

## Methods

### NewOswDevCapVO

`func NewOswDevCapVO() *OswDevCapVO`

NewOswDevCapVO instantiates a new OswDevCapVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDevCapVOWithDefaults

`func NewOswDevCapVOWithDefaults() *OswDevCapVO`

NewOswDevCapVOWithDefaults instantiates a new OswDevCapVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpTableSupport

`func (o *OswDevCapVO) GetArpTableSupport() bool`

GetArpTableSupport returns the ArpTableSupport field if non-nil, zero value otherwise.

### GetArpTableSupportOk

`func (o *OswDevCapVO) GetArpTableSupportOk() (*bool, bool)`

GetArpTableSupportOk returns a tuple with the ArpTableSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpTableSupport

`func (o *OswDevCapVO) SetArpTableSupport(v bool)`

SetArpTableSupport sets ArpTableSupport field to given value.

### HasArpTableSupport

`func (o *OswDevCapVO) HasArpTableSupport() bool`

HasArpTableSupport returns a boolean if a field has been set.

### GetCableTestSupport

`func (o *OswDevCapVO) GetCableTestSupport() bool`

GetCableTestSupport returns the CableTestSupport field if non-nil, zero value otherwise.

### GetCableTestSupportOk

`func (o *OswDevCapVO) GetCableTestSupportOk() (*bool, bool)`

GetCableTestSupportOk returns a tuple with the CableTestSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableTestSupport

`func (o *OswDevCapVO) SetCableTestSupport(v bool)`

SetCableTestSupport sets CableTestSupport field to given value.

### HasCableTestSupport

`func (o *OswDevCapVO) HasCableTestSupport() bool`

HasCableTestSupport returns a boolean if a field has been set.

### GetCliSupport

`func (o *OswDevCapVO) GetCliSupport() bool`

GetCliSupport returns the CliSupport field if non-nil, zero value otherwise.

### GetCliSupportOk

`func (o *OswDevCapVO) GetCliSupportOk() (*bool, bool)`

GetCliSupportOk returns a tuple with the CliSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliSupport

`func (o *OswDevCapVO) SetCliSupport(v bool)`

SetCliSupport sets CliSupport field to given value.

### HasCliSupport

`func (o *OswDevCapVO) HasCliSupport() bool`

HasCliSupport returns a boolean if a field has been set.

### GetCustomDhcpOptionSupport

`func (o *OswDevCapVO) GetCustomDhcpOptionSupport() bool`

GetCustomDhcpOptionSupport returns the CustomDhcpOptionSupport field if non-nil, zero value otherwise.

### GetCustomDhcpOptionSupportOk

`func (o *OswDevCapVO) GetCustomDhcpOptionSupportOk() (*bool, bool)`

GetCustomDhcpOptionSupportOk returns a tuple with the CustomDhcpOptionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDhcpOptionSupport

`func (o *OswDevCapVO) SetCustomDhcpOptionSupport(v bool)`

SetCustomDhcpOptionSupport sets CustomDhcpOptionSupport field to given value.

### HasCustomDhcpOptionSupport

`func (o *OswDevCapVO) HasCustomDhcpOptionSupport() bool`

HasCustomDhcpOptionSupport returns a boolean if a field has been set.

### GetDdmSupport

`func (o *OswDevCapVO) GetDdmSupport() bool`

GetDdmSupport returns the DdmSupport field if non-nil, zero value otherwise.

### GetDdmSupportOk

`func (o *OswDevCapVO) GetDdmSupportOk() (*bool, bool)`

GetDdmSupportOk returns a tuple with the DdmSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdmSupport

`func (o *OswDevCapVO) SetDdmSupport(v bool)`

SetDdmSupport sets DdmSupport field to given value.

### HasDdmSupport

`func (o *OswDevCapVO) HasDdmSupport() bool`

HasDdmSupport returns a boolean if a field has been set.

### GetDhcpL2RelaySupport

`func (o *OswDevCapVO) GetDhcpL2RelaySupport() bool`

GetDhcpL2RelaySupport returns the DhcpL2RelaySupport field if non-nil, zero value otherwise.

### GetDhcpL2RelaySupportOk

`func (o *OswDevCapVO) GetDhcpL2RelaySupportOk() (*bool, bool)`

GetDhcpL2RelaySupportOk returns a tuple with the DhcpL2RelaySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpL2RelaySupport

`func (o *OswDevCapVO) SetDhcpL2RelaySupport(v bool)`

SetDhcpL2RelaySupport sets DhcpL2RelaySupport field to given value.

### HasDhcpL2RelaySupport

`func (o *OswDevCapVO) HasDhcpL2RelaySupport() bool`

HasDhcpL2RelaySupport returns a boolean if a field has been set.

### GetDhcpRangeSupport

`func (o *OswDevCapVO) GetDhcpRangeSupport() bool`

GetDhcpRangeSupport returns the DhcpRangeSupport field if non-nil, zero value otherwise.

### GetDhcpRangeSupportOk

`func (o *OswDevCapVO) GetDhcpRangeSupportOk() (*bool, bool)`

GetDhcpRangeSupportOk returns a tuple with the DhcpRangeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRangeSupport

`func (o *OswDevCapVO) SetDhcpRangeSupport(v bool)`

SetDhcpRangeSupport sets DhcpRangeSupport field to given value.

### HasDhcpRangeSupport

`func (o *OswDevCapVO) HasDhcpRangeSupport() bool`

HasDhcpRangeSupport returns a boolean if a field has been set.

### GetDhcpRelayIfNum

`func (o *OswDevCapVO) GetDhcpRelayIfNum() int32`

GetDhcpRelayIfNum returns the DhcpRelayIfNum field if non-nil, zero value otherwise.

### GetDhcpRelayIfNumOk

`func (o *OswDevCapVO) GetDhcpRelayIfNumOk() (*int32, bool)`

GetDhcpRelayIfNumOk returns a tuple with the DhcpRelayIfNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayIfNum

`func (o *OswDevCapVO) SetDhcpRelayIfNum(v int32)`

SetDhcpRelayIfNum sets DhcpRelayIfNum field to given value.

### HasDhcpRelayIfNum

`func (o *OswDevCapVO) HasDhcpRelayIfNum() bool`

HasDhcpRelayIfNum returns a boolean if a field has been set.

### GetDhcpReservationSupport

`func (o *OswDevCapVO) GetDhcpReservationSupport() bool`

GetDhcpReservationSupport returns the DhcpReservationSupport field if non-nil, zero value otherwise.

### GetDhcpReservationSupportOk

`func (o *OswDevCapVO) GetDhcpReservationSupportOk() (*bool, bool)`

GetDhcpReservationSupportOk returns a tuple with the DhcpReservationSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpReservationSupport

`func (o *OswDevCapVO) SetDhcpReservationSupport(v bool)`

SetDhcpReservationSupport sets DhcpReservationSupport field to given value.

### HasDhcpReservationSupport

`func (o *OswDevCapVO) HasDhcpReservationSupport() bool`

HasDhcpReservationSupport returns a boolean if a field has been set.

### GetDhcpServerPoolNum

`func (o *OswDevCapVO) GetDhcpServerPoolNum() int32`

GetDhcpServerPoolNum returns the DhcpServerPoolNum field if non-nil, zero value otherwise.

### GetDhcpServerPoolNumOk

`func (o *OswDevCapVO) GetDhcpServerPoolNumOk() (*int32, bool)`

GetDhcpServerPoolNumOk returns a tuple with the DhcpServerPoolNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerPoolNum

`func (o *OswDevCapVO) SetDhcpServerPoolNum(v int32)`

SetDhcpServerPoolNum sets DhcpServerPoolNum field to given value.

### HasDhcpServerPoolNum

`func (o *OswDevCapVO) HasDhcpServerPoolNum() bool`

HasDhcpServerPoolNum returns a boolean if a field has been set.

### GetDnsLoopUpSupport

`func (o *OswDevCapVO) GetDnsLoopUpSupport() bool`

GetDnsLoopUpSupport returns the DnsLoopUpSupport field if non-nil, zero value otherwise.

### GetDnsLoopUpSupportOk

`func (o *OswDevCapVO) GetDnsLoopUpSupportOk() (*bool, bool)`

GetDnsLoopUpSupportOk returns a tuple with the DnsLoopUpSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLoopUpSupport

`func (o *OswDevCapVO) SetDnsLoopUpSupport(v bool)`

SetDnsLoopUpSupport sets DnsLoopUpSupport field to given value.

### HasDnsLoopUpSupport

`func (o *OswDevCapVO) HasDnsLoopUpSupport() bool`

HasDnsLoopUpSupport returns a boolean if a field has been set.

### GetDomainPingSupport

`func (o *OswDevCapVO) GetDomainPingSupport() bool`

GetDomainPingSupport returns the DomainPingSupport field if non-nil, zero value otherwise.

### GetDomainPingSupportOk

`func (o *OswDevCapVO) GetDomainPingSupportOk() (*bool, bool)`

GetDomainPingSupportOk returns a tuple with the DomainPingSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPingSupport

`func (o *OswDevCapVO) SetDomainPingSupport(v bool)`

SetDomainPingSupport sets DomainPingSupport field to given value.

### HasDomainPingSupport

`func (o *OswDevCapVO) HasDomainPingSupport() bool`

HasDomainPingSupport returns a boolean if a field has been set.

### GetDomainTraceRouteSupport

`func (o *OswDevCapVO) GetDomainTraceRouteSupport() bool`

GetDomainTraceRouteSupport returns the DomainTraceRouteSupport field if non-nil, zero value otherwise.

### GetDomainTraceRouteSupportOk

`func (o *OswDevCapVO) GetDomainTraceRouteSupportOk() (*bool, bool)`

GetDomainTraceRouteSupportOk returns a tuple with the DomainTraceRouteSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainTraceRouteSupport

`func (o *OswDevCapVO) SetDomainTraceRouteSupport(v bool)`

SetDomainTraceRouteSupport sets DomainTraceRouteSupport field to given value.

### HasDomainTraceRouteSupport

`func (o *OswDevCapVO) HasDomainTraceRouteSupport() bool`

HasDomainTraceRouteSupport returns a boolean if a field has been set.

### GetEeeSupport

`func (o *OswDevCapVO) GetEeeSupport() bool`

GetEeeSupport returns the EeeSupport field if non-nil, zero value otherwise.

### GetEeeSupportOk

`func (o *OswDevCapVO) GetEeeSupportOk() (*bool, bool)`

GetEeeSupportOk returns a tuple with the EeeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeeSupport

`func (o *OswDevCapVO) SetEeeSupport(v bool)`

SetEeeSupport sets EeeSupport field to given value.

### HasEeeSupport

`func (o *OswDevCapVO) HasEeeSupport() bool`

HasEeeSupport returns a boolean if a field has been set.

### GetEsQosSupport

`func (o *OswDevCapVO) GetEsQosSupport() bool`

GetEsQosSupport returns the EsQosSupport field if non-nil, zero value otherwise.

### GetEsQosSupportOk

`func (o *OswDevCapVO) GetEsQosSupportOk() (*bool, bool)`

GetEsQosSupportOk returns a tuple with the EsQosSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsQosSupport

`func (o *OswDevCapVO) SetEsQosSupport(v bool)`

SetEsQosSupport sets EsQosSupport field to given value.

### HasEsQosSupport

`func (o *OswDevCapVO) HasEsQosSupport() bool`

HasEsQosSupport returns a boolean if a field has been set.

### GetFecSupport

`func (o *OswDevCapVO) GetFecSupport() bool`

GetFecSupport returns the FecSupport field if non-nil, zero value otherwise.

### GetFecSupportOk

`func (o *OswDevCapVO) GetFecSupportOk() (*bool, bool)`

GetFecSupportOk returns a tuple with the FecSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecSupport

`func (o *OswDevCapVO) SetFecSupport(v bool)`

SetFecSupport sets FecSupport field to given value.

### HasFecSupport

`func (o *OswDevCapVO) HasFecSupport() bool`

HasFecSupport returns a boolean if a field has been set.

### GetFlowControlSupport

`func (o *OswDevCapVO) GetFlowControlSupport() bool`

GetFlowControlSupport returns the FlowControlSupport field if non-nil, zero value otherwise.

### GetFlowControlSupportOk

`func (o *OswDevCapVO) GetFlowControlSupportOk() (*bool, bool)`

GetFlowControlSupportOk returns a tuple with the FlowControlSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlSupport

`func (o *OswDevCapVO) SetFlowControlSupport(v bool)`

SetFlowControlSupport sets FlowControlSupport field to given value.

### HasFlowControlSupport

`func (o *OswDevCapVO) HasFlowControlSupport() bool`

HasFlowControlSupport returns a boolean if a field has been set.

### GetGponSupport

`func (o *OswDevCapVO) GetGponSupport() bool`

GetGponSupport returns the GponSupport field if non-nil, zero value otherwise.

### GetGponSupportOk

`func (o *OswDevCapVO) GetGponSupportOk() (*bool, bool)`

GetGponSupportOk returns a tuple with the GponSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGponSupport

`func (o *OswDevCapVO) SetGponSupport(v bool)`

SetGponSupport sets GponSupport field to given value.

### HasGponSupport

`func (o *OswDevCapVO) HasGponSupport() bool`

HasGponSupport returns a boolean if a field has been set.

### GetIgmpFastLeaveSupport

`func (o *OswDevCapVO) GetIgmpFastLeaveSupport() bool`

GetIgmpFastLeaveSupport returns the IgmpFastLeaveSupport field if non-nil, zero value otherwise.

### GetIgmpFastLeaveSupportOk

`func (o *OswDevCapVO) GetIgmpFastLeaveSupportOk() (*bool, bool)`

GetIgmpFastLeaveSupportOk returns a tuple with the IgmpFastLeaveSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpFastLeaveSupport

`func (o *OswDevCapVO) SetIgmpFastLeaveSupport(v bool)`

SetIgmpFastLeaveSupport sets IgmpFastLeaveSupport field to given value.

### HasIgmpFastLeaveSupport

`func (o *OswDevCapVO) HasIgmpFastLeaveSupport() bool`

HasIgmpFastLeaveSupport returns a boolean if a field has been set.

### GetIpcDetectSupport

`func (o *OswDevCapVO) GetIpcDetectSupport() bool`

GetIpcDetectSupport returns the IpcDetectSupport field if non-nil, zero value otherwise.

### GetIpcDetectSupportOk

`func (o *OswDevCapVO) GetIpcDetectSupportOk() (*bool, bool)`

GetIpcDetectSupportOk returns a tuple with the IpcDetectSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpcDetectSupport

`func (o *OswDevCapVO) SetIpcDetectSupport(v bool)`

SetIpcDetectSupport sets IpcDetectSupport field to given value.

### HasIpcDetectSupport

`func (o *OswDevCapVO) HasIpcDetectSupport() bool`

HasIpcDetectSupport returns a boolean if a field has been set.

### GetIpv4StaticRouteNum

`func (o *OswDevCapVO) GetIpv4StaticRouteNum() int32`

GetIpv4StaticRouteNum returns the Ipv4StaticRouteNum field if non-nil, zero value otherwise.

### GetIpv4StaticRouteNumOk

`func (o *OswDevCapVO) GetIpv4StaticRouteNumOk() (*int32, bool)`

GetIpv4StaticRouteNumOk returns a tuple with the Ipv4StaticRouteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4StaticRouteNum

`func (o *OswDevCapVO) SetIpv4StaticRouteNum(v int32)`

SetIpv4StaticRouteNum sets Ipv4StaticRouteNum field to given value.

### HasIpv4StaticRouteNum

`func (o *OswDevCapVO) HasIpv4StaticRouteNum() bool`

HasIpv4StaticRouteNum returns a boolean if a field has been set.

### GetIpv6StaticRouteNum

`func (o *OswDevCapVO) GetIpv6StaticRouteNum() int32`

GetIpv6StaticRouteNum returns the Ipv6StaticRouteNum field if non-nil, zero value otherwise.

### GetIpv6StaticRouteNumOk

`func (o *OswDevCapVO) GetIpv6StaticRouteNumOk() (*int32, bool)`

GetIpv6StaticRouteNumOk returns a tuple with the Ipv6StaticRouteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6StaticRouteNum

`func (o *OswDevCapVO) SetIpv6StaticRouteNum(v int32)`

SetIpv6StaticRouteNum sets Ipv6StaticRouteNum field to given value.

### HasIpv6StaticRouteNum

`func (o *OswDevCapVO) HasIpv6StaticRouteNum() bool`

HasIpv6StaticRouteNum returns a boolean if a field has been set.

### GetJumboOddSupport

`func (o *OswDevCapVO) GetJumboOddSupport() bool`

GetJumboOddSupport returns the JumboOddSupport field if non-nil, zero value otherwise.

### GetJumboOddSupportOk

`func (o *OswDevCapVO) GetJumboOddSupportOk() (*bool, bool)`

GetJumboOddSupportOk returns a tuple with the JumboOddSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboOddSupport

`func (o *OswDevCapVO) SetJumboOddSupport(v bool)`

SetJumboOddSupport sets JumboOddSupport field to given value.

### HasJumboOddSupport

`func (o *OswDevCapVO) HasJumboOddSupport() bool`

HasJumboOddSupport returns a boolean if a field has been set.

### GetJumboSupport

`func (o *OswDevCapVO) GetJumboSupport() bool`

GetJumboSupport returns the JumboSupport field if non-nil, zero value otherwise.

### GetJumboSupportOk

`func (o *OswDevCapVO) GetJumboSupportOk() (*bool, bool)`

GetJumboSupportOk returns a tuple with the JumboSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboSupport

`func (o *OswDevCapVO) SetJumboSupport(v bool)`

SetJumboSupport sets JumboSupport field to given value.

### HasJumboSupport

`func (o *OswDevCapVO) HasJumboSupport() bool`

HasJumboSupport returns a boolean if a field has been set.

### GetLagCap

`func (o *OswDevCapVO) GetLagCap() LagCapVO`

GetLagCap returns the LagCap field if non-nil, zero value otherwise.

### GetLagCapOk

`func (o *OswDevCapVO) GetLagCapOk() (*LagCapVO, bool)`

GetLagCapOk returns a tuple with the LagCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagCap

`func (o *OswDevCapVO) SetLagCap(v LagCapVO)`

SetLagCap sets LagCap field to given value.

### HasLagCap

`func (o *OswDevCapVO) HasLagCap() bool`

HasLagCap returns a boolean if a field has been set.

### GetLocatePortSupport

`func (o *OswDevCapVO) GetLocatePortSupport() bool`

GetLocatePortSupport returns the LocatePortSupport field if non-nil, zero value otherwise.

### GetLocatePortSupportOk

`func (o *OswDevCapVO) GetLocatePortSupportOk() (*bool, bool)`

GetLocatePortSupportOk returns a tuple with the LocatePortSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatePortSupport

`func (o *OswDevCapVO) SetLocatePortSupport(v bool)`

SetLocatePortSupport sets LocatePortSupport field to given value.

### HasLocatePortSupport

`func (o *OswDevCapVO) HasLocatePortSupport() bool`

HasLocatePortSupport returns a boolean if a field has been set.

### GetLoopbackInterfaceSupport

`func (o *OswDevCapVO) GetLoopbackInterfaceSupport() bool`

GetLoopbackInterfaceSupport returns the LoopbackInterfaceSupport field if non-nil, zero value otherwise.

### GetLoopbackInterfaceSupportOk

`func (o *OswDevCapVO) GetLoopbackInterfaceSupportOk() (*bool, bool)`

GetLoopbackInterfaceSupportOk returns a tuple with the LoopbackInterfaceSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackInterfaceSupport

`func (o *OswDevCapVO) SetLoopbackInterfaceSupport(v bool)`

SetLoopbackInterfaceSupport sets LoopbackInterfaceSupport field to given value.

### HasLoopbackInterfaceSupport

`func (o *OswDevCapVO) HasLoopbackInterfaceSupport() bool`

HasLoopbackInterfaceSupport returns a boolean if a field has been set.

### GetLoopbackVlanBasedSupport

`func (o *OswDevCapVO) GetLoopbackVlanBasedSupport() bool`

GetLoopbackVlanBasedSupport returns the LoopbackVlanBasedSupport field if non-nil, zero value otherwise.

### GetLoopbackVlanBasedSupportOk

`func (o *OswDevCapVO) GetLoopbackVlanBasedSupportOk() (*bool, bool)`

GetLoopbackVlanBasedSupportOk returns a tuple with the LoopbackVlanBasedSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackVlanBasedSupport

`func (o *OswDevCapVO) SetLoopbackVlanBasedSupport(v bool)`

SetLoopbackVlanBasedSupport sets LoopbackVlanBasedSupport field to given value.

### HasLoopbackVlanBasedSupport

`func (o *OswDevCapVO) HasLoopbackVlanBasedSupport() bool`

HasLoopbackVlanBasedSupport returns a boolean if a field has been set.

### GetMaxLagMember

`func (o *OswDevCapVO) GetMaxLagMember() int32`

GetMaxLagMember returns the MaxLagMember field if non-nil, zero value otherwise.

### GetMaxLagMemberOk

`func (o *OswDevCapVO) GetMaxLagMemberOk() (*int32, bool)`

GetMaxLagMemberOk returns a tuple with the MaxLagMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagMember

`func (o *OswDevCapVO) SetMaxLagMember(v int32)`

SetMaxLagMember sets MaxLagMember field to given value.

### HasMaxLagMember

`func (o *OswDevCapVO) HasMaxLagMember() bool`

HasMaxLagMember returns a boolean if a field has been set.

### GetMaxLagNum

`func (o *OswDevCapVO) GetMaxLagNum() int32`

GetMaxLagNum returns the MaxLagNum field if non-nil, zero value otherwise.

### GetMaxLagNumOk

`func (o *OswDevCapVO) GetMaxLagNumOk() (*int32, bool)`

GetMaxLagNumOk returns a tuple with the MaxLagNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLagNum

`func (o *OswDevCapVO) SetMaxLagNum(v int32)`

SetMaxLagNum sets MaxLagNum field to given value.

### HasMaxLagNum

`func (o *OswDevCapVO) HasMaxLagNum() bool`

HasMaxLagNum returns a boolean if a field has been set.

### GetMaxMirrorGroup

`func (o *OswDevCapVO) GetMaxMirrorGroup() int32`

GetMaxMirrorGroup returns the MaxMirrorGroup field if non-nil, zero value otherwise.

### GetMaxMirrorGroupOk

`func (o *OswDevCapVO) GetMaxMirrorGroupOk() (*int32, bool)`

GetMaxMirrorGroupOk returns a tuple with the MaxMirrorGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMirrorGroup

`func (o *OswDevCapVO) SetMaxMirrorGroup(v int32)`

SetMaxMirrorGroup sets MaxMirrorGroup field to given value.

### HasMaxMirrorGroup

`func (o *OswDevCapVO) HasMaxMirrorGroup() bool`

HasMaxMirrorGroup returns a boolean if a field has been set.

### GetMaxMirroredPort

`func (o *OswDevCapVO) GetMaxMirroredPort() int32`

GetMaxMirroredPort returns the MaxMirroredPort field if non-nil, zero value otherwise.

### GetMaxMirroredPortOk

`func (o *OswDevCapVO) GetMaxMirroredPortOk() (*int32, bool)`

GetMaxMirroredPortOk returns a tuple with the MaxMirroredPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMirroredPort

`func (o *OswDevCapVO) SetMaxMirroredPort(v int32)`

SetMaxMirroredPort sets MaxMirroredPort field to given value.

### HasMaxMirroredPort

`func (o *OswDevCapVO) HasMaxMirroredPort() bool`

HasMaxMirroredPort returns a boolean if a field has been set.

### GetMaxRelayServerNum

`func (o *OswDevCapVO) GetMaxRelayServerNum() int32`

GetMaxRelayServerNum returns the MaxRelayServerNum field if non-nil, zero value otherwise.

### GetMaxRelayServerNumOk

`func (o *OswDevCapVO) GetMaxRelayServerNumOk() (*int32, bool)`

GetMaxRelayServerNumOk returns a tuple with the MaxRelayServerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRelayServerNum

`func (o *OswDevCapVO) SetMaxRelayServerNum(v int32)`

SetMaxRelayServerNum sets MaxRelayServerNum field to given value.

### HasMaxRelayServerNum

`func (o *OswDevCapVO) HasMaxRelayServerNum() bool`

HasMaxRelayServerNum returns a boolean if a field has been set.

### GetMaxStackGroupNumber

`func (o *OswDevCapVO) GetMaxStackGroupNumber() int32`

GetMaxStackGroupNumber returns the MaxStackGroupNumber field if non-nil, zero value otherwise.

### GetMaxStackGroupNumberOk

`func (o *OswDevCapVO) GetMaxStackGroupNumberOk() (*int32, bool)`

GetMaxStackGroupNumberOk returns a tuple with the MaxStackGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStackGroupNumber

`func (o *OswDevCapVO) SetMaxStackGroupNumber(v int32)`

SetMaxStackGroupNumber sets MaxStackGroupNumber field to given value.

### HasMaxStackGroupNumber

`func (o *OswDevCapVO) HasMaxStackGroupNumber() bool`

HasMaxStackGroupNumber returns a boolean if a field has been set.

### GetMaxStackUnitNumber

`func (o *OswDevCapVO) GetMaxStackUnitNumber() int32`

GetMaxStackUnitNumber returns the MaxStackUnitNumber field if non-nil, zero value otherwise.

### GetMaxStackUnitNumberOk

`func (o *OswDevCapVO) GetMaxStackUnitNumberOk() (*int32, bool)`

GetMaxStackUnitNumberOk returns a tuple with the MaxStackUnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStackUnitNumber

`func (o *OswDevCapVO) SetMaxStackUnitNumber(v int32)`

SetMaxStackUnitNumber sets MaxStackUnitNumber field to given value.

### HasMaxStackUnitNumber

`func (o *OswDevCapVO) HasMaxStackUnitNumber() bool`

HasMaxStackUnitNumber returns a boolean if a field has been set.

### GetMirrorSupport

`func (o *OswDevCapVO) GetMirrorSupport() bool`

GetMirrorSupport returns the MirrorSupport field if non-nil, zero value otherwise.

### GetMirrorSupportOk

`func (o *OswDevCapVO) GetMirrorSupportOk() (*bool, bool)`

GetMirrorSupportOk returns a tuple with the MirrorSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorSupport

`func (o *OswDevCapVO) SetMirrorSupport(v bool)`

SetMirrorSupport sets MirrorSupport field to given value.

### HasMirrorSupport

`func (o *OswDevCapVO) HasMirrorSupport() bool`

HasMirrorSupport returns a boolean if a field has been set.

### GetMlagGroupId

`func (o *OswDevCapVO) GetMlagGroupId() int32`

GetMlagGroupId returns the MlagGroupId field if non-nil, zero value otherwise.

### GetMlagGroupIdOk

`func (o *OswDevCapVO) GetMlagGroupIdOk() (*int32, bool)`

GetMlagGroupIdOk returns a tuple with the MlagGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagGroupId

`func (o *OswDevCapVO) SetMlagGroupId(v int32)`

SetMlagGroupId sets MlagGroupId field to given value.

### HasMlagGroupId

`func (o *OswDevCapVO) HasMlagGroupId() bool`

HasMlagGroupId returns a boolean if a field has been set.

### GetMlagVersion

`func (o *OswDevCapVO) GetMlagVersion() string`

GetMlagVersion returns the MlagVersion field if non-nil, zero value otherwise.

### GetMlagVersionOk

`func (o *OswDevCapVO) GetMlagVersionOk() (*string, bool)`

GetMlagVersionOk returns a tuple with the MlagVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagVersion

`func (o *OswDevCapVO) SetMlagVersion(v string)`

SetMlagVersion sets MlagVersion field to given value.

### HasMlagVersion

`func (o *OswDevCapVO) HasMlagVersion() bool`

HasMlagVersion returns a boolean if a field has been set.

### GetMstpInsNum

`func (o *OswDevCapVO) GetMstpInsNum() int32`

GetMstpInsNum returns the MstpInsNum field if non-nil, zero value otherwise.

### GetMstpInsNumOk

`func (o *OswDevCapVO) GetMstpInsNumOk() (*int32, bool)`

GetMstpInsNumOk returns a tuple with the MstpInsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstpInsNum

`func (o *OswDevCapVO) SetMstpInsNum(v int32)`

SetMstpInsNum sets MstpInsNum field to given value.

### HasMstpInsNum

`func (o *OswDevCapVO) HasMstpInsNum() bool`

HasMstpInsNum returns a boolean if a field has been set.

### GetNeedFullSync

`func (o *OswDevCapVO) GetNeedFullSync() bool`

GetNeedFullSync returns the NeedFullSync field if non-nil, zero value otherwise.

### GetNeedFullSyncOk

`func (o *OswDevCapVO) GetNeedFullSyncOk() (*bool, bool)`

GetNeedFullSyncOk returns a tuple with the NeedFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedFullSync

`func (o *OswDevCapVO) SetNeedFullSync(v bool)`

SetNeedFullSync sets NeedFullSync field to given value.

### HasNeedFullSync

`func (o *OswDevCapVO) HasNeedFullSync() bool`

HasNeedFullSync returns a boolean if a field has been set.

### GetNetworkCheckSupport

`func (o *OswDevCapVO) GetNetworkCheckSupport() bool`

GetNetworkCheckSupport returns the NetworkCheckSupport field if non-nil, zero value otherwise.

### GetNetworkCheckSupportOk

`func (o *OswDevCapVO) GetNetworkCheckSupportOk() (*bool, bool)`

GetNetworkCheckSupportOk returns a tuple with the NetworkCheckSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCheckSupport

`func (o *OswDevCapVO) SetNetworkCheckSupport(v bool)`

SetNetworkCheckSupport sets NetworkCheckSupport field to given value.

### HasNetworkCheckSupport

`func (o *OswDevCapVO) HasNetworkCheckSupport() bool`

HasNetworkCheckSupport returns a boolean if a field has been set.

### GetOspfDeadIntervalSupport

`func (o *OswDevCapVO) GetOspfDeadIntervalSupport() bool`

GetOspfDeadIntervalSupport returns the OspfDeadIntervalSupport field if non-nil, zero value otherwise.

### GetOspfDeadIntervalSupportOk

`func (o *OswDevCapVO) GetOspfDeadIntervalSupportOk() (*bool, bool)`

GetOspfDeadIntervalSupportOk returns a tuple with the OspfDeadIntervalSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfDeadIntervalSupport

`func (o *OswDevCapVO) SetOspfDeadIntervalSupport(v bool)`

SetOspfDeadIntervalSupport sets OspfDeadIntervalSupport field to given value.

### HasOspfDeadIntervalSupport

`func (o *OswDevCapVO) HasOspfDeadIntervalSupport() bool`

HasOspfDeadIntervalSupport returns a boolean if a field has been set.

### GetOspfSupport

`func (o *OswDevCapVO) GetOspfSupport() bool`

GetOspfSupport returns the OspfSupport field if non-nil, zero value otherwise.

### GetOspfSupportOk

`func (o *OswDevCapVO) GetOspfSupportOk() (*bool, bool)`

GetOspfSupportOk returns a tuple with the OspfSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSupport

`func (o *OswDevCapVO) SetOspfSupport(v bool)`

SetOspfSupport sets OspfSupport field to given value.

### HasOspfSupport

`func (o *OswDevCapVO) HasOspfSupport() bool`

HasOspfSupport returns a boolean if a field has been set.

### GetOuiBasedVlanSupport

`func (o *OswDevCapVO) GetOuiBasedVlanSupport() bool`

GetOuiBasedVlanSupport returns the OuiBasedVlanSupport field if non-nil, zero value otherwise.

### GetOuiBasedVlanSupportOk

`func (o *OswDevCapVO) GetOuiBasedVlanSupportOk() (*bool, bool)`

GetOuiBasedVlanSupportOk returns a tuple with the OuiBasedVlanSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiBasedVlanSupport

`func (o *OswDevCapVO) SetOuiBasedVlanSupport(v bool)`

SetOuiBasedVlanSupport sets OuiBasedVlanSupport field to given value.

### HasOuiBasedVlanSupport

`func (o *OswDevCapVO) HasOuiBasedVlanSupport() bool`

HasOuiBasedVlanSupport returns a boolean if a field has been set.

### GetPacketCaptureSupport

`func (o *OswDevCapVO) GetPacketCaptureSupport() bool`

GetPacketCaptureSupport returns the PacketCaptureSupport field if non-nil, zero value otherwise.

### GetPacketCaptureSupportOk

`func (o *OswDevCapVO) GetPacketCaptureSupportOk() (*bool, bool)`

GetPacketCaptureSupportOk returns a tuple with the PacketCaptureSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketCaptureSupport

`func (o *OswDevCapVO) SetPacketCaptureSupport(v bool)`

SetPacketCaptureSupport sets PacketCaptureSupport field to given value.

### HasPacketCaptureSupport

`func (o *OswDevCapVO) HasPacketCaptureSupport() bool`

HasPacketCaptureSupport returns a boolean if a field has been set.

### GetPingSupport

`func (o *OswDevCapVO) GetPingSupport() bool`

GetPingSupport returns the PingSupport field if non-nil, zero value otherwise.

### GetPingSupportOk

`func (o *OswDevCapVO) GetPingSupportOk() (*bool, bool)`

GetPingSupportOk returns a tuple with the PingSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSupport

`func (o *OswDevCapVO) SetPingSupport(v bool)`

SetPingSupport sets PingSupport field to given value.

### HasPingSupport

`func (o *OswDevCapVO) HasPingSupport() bool`

HasPingSupport returns a boolean if a field has been set.

### GetPoePortNum

`func (o *OswDevCapVO) GetPoePortNum() int32`

GetPoePortNum returns the PoePortNum field if non-nil, zero value otherwise.

### GetPoePortNumOk

`func (o *OswDevCapVO) GetPoePortNumOk() (*int32, bool)`

GetPoePortNumOk returns a tuple with the PoePortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePortNum

`func (o *OswDevCapVO) SetPoePortNum(v int32)`

SetPoePortNum sets PoePortNum field to given value.

### HasPoePortNum

`func (o *OswDevCapVO) HasPoePortNum() bool`

HasPoePortNum returns a boolean if a field has been set.

### GetPoeSupport

`func (o *OswDevCapVO) GetPoeSupport() bool`

GetPoeSupport returns the PoeSupport field if non-nil, zero value otherwise.

### GetPoeSupportOk

`func (o *OswDevCapVO) GetPoeSupportOk() (*bool, bool)`

GetPoeSupportOk returns a tuple with the PoeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeSupport

`func (o *OswDevCapVO) SetPoeSupport(v bool)`

SetPoeSupport sets PoeSupport field to given value.

### HasPoeSupport

`func (o *OswDevCapVO) HasPoeSupport() bool`

HasPoeSupport returns a boolean if a field has been set.

### GetQosForVlanSupport

`func (o *OswDevCapVO) GetQosForVlanSupport() bool`

GetQosForVlanSupport returns the QosForVlanSupport field if non-nil, zero value otherwise.

### GetQosForVlanSupportOk

`func (o *OswDevCapVO) GetQosForVlanSupportOk() (*bool, bool)`

GetQosForVlanSupportOk returns a tuple with the QosForVlanSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosForVlanSupport

`func (o *OswDevCapVO) SetQosForVlanSupport(v bool)`

SetQosForVlanSupport sets QosForVlanSupport field to given value.

### HasQosForVlanSupport

`func (o *OswDevCapVO) HasQosForVlanSupport() bool`

HasQosForVlanSupport returns a boolean if a field has been set.

### GetQosSupport

`func (o *OswDevCapVO) GetQosSupport() bool`

GetQosSupport returns the QosSupport field if non-nil, zero value otherwise.

### GetQosSupportOk

`func (o *OswDevCapVO) GetQosSupportOk() (*bool, bool)`

GetQosSupportOk returns a tuple with the QosSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSupport

`func (o *OswDevCapVO) SetQosSupport(v bool)`

SetQosSupport sets QosSupport field to given value.

### HasQosSupport

`func (o *OswDevCapVO) HasQosSupport() bool`

HasQosSupport returns a boolean if a field has been set.

### GetSfpBeginNum

`func (o *OswDevCapVO) GetSfpBeginNum() int32`

GetSfpBeginNum returns the SfpBeginNum field if non-nil, zero value otherwise.

### GetSfpBeginNumOk

`func (o *OswDevCapVO) GetSfpBeginNumOk() (*int32, bool)`

GetSfpBeginNumOk returns a tuple with the SfpBeginNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpBeginNum

`func (o *OswDevCapVO) SetSfpBeginNum(v int32)`

SetSfpBeginNum sets SfpBeginNum field to given value.

### HasSfpBeginNum

`func (o *OswDevCapVO) HasSfpBeginNum() bool`

HasSfpBeginNum returns a boolean if a field has been set.

### GetSfpNum

`func (o *OswDevCapVO) GetSfpNum() int32`

GetSfpNum returns the SfpNum field if non-nil, zero value otherwise.

### GetSfpNumOk

`func (o *OswDevCapVO) GetSfpNumOk() (*int32, bool)`

GetSfpNumOk returns a tuple with the SfpNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpNum

`func (o *OswDevCapVO) SetSfpNum(v int32)`

SetSfpNum sets SfpNum field to given value.

### HasSfpNum

`func (o *OswDevCapVO) HasSfpNum() bool`

HasSfpNum returns a boolean if a field has been set.

### GetStackPortCap

`func (o *OswDevCapVO) GetStackPortCap() map[string][]string`

GetStackPortCap returns the StackPortCap field if non-nil, zero value otherwise.

### GetStackPortCapOk

`func (o *OswDevCapVO) GetStackPortCapOk() (*map[string][]string, bool)`

GetStackPortCapOk returns a tuple with the StackPortCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortCap

`func (o *OswDevCapVO) SetStackPortCap(v map[string][]string)`

SetStackPortCap sets StackPortCap field to given value.

### HasStackPortCap

`func (o *OswDevCapVO) HasStackPortCap() bool`

HasStackPortCap returns a boolean if a field has been set.

### GetStaticRouteNum

`func (o *OswDevCapVO) GetStaticRouteNum() int32`

GetStaticRouteNum returns the StaticRouteNum field if non-nil, zero value otherwise.

### GetStaticRouteNumOk

`func (o *OswDevCapVO) GetStaticRouteNumOk() (*int32, bool)`

GetStaticRouteNumOk returns a tuple with the StaticRouteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRouteNum

`func (o *OswDevCapVO) SetStaticRouteNum(v int32)`

SetStaticRouteNum sets StaticRouteNum field to given value.

### HasStaticRouteNum

`func (o *OswDevCapVO) HasStaticRouteNum() bool`

HasStaticRouteNum returns a boolean if a field has been set.

### GetStkVer

`func (o *OswDevCapVO) GetStkVer() string`

GetStkVer returns the StkVer field if non-nil, zero value otherwise.

### GetStkVerOk

`func (o *OswDevCapVO) GetStkVerOk() (*string, bool)`

GetStkVerOk returns a tuple with the StkVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkVer

`func (o *OswDevCapVO) SetStkVer(v string)`

SetStkVer sets StkVer field to given value.

### HasStkVer

`func (o *OswDevCapVO) HasStkVer() bool`

HasStkVer returns a boolean if a field has been set.

### GetStkableGroupId

`func (o *OswDevCapVO) GetStkableGroupId() int32`

GetStkableGroupId returns the StkableGroupId field if non-nil, zero value otherwise.

### GetStkableGroupIdOk

`func (o *OswDevCapVO) GetStkableGroupIdOk() (*int32, bool)`

GetStkableGroupIdOk returns a tuple with the StkableGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkableGroupId

`func (o *OswDevCapVO) SetStkableGroupId(v int32)`

SetStkableGroupId sets StkableGroupId field to given value.

### HasStkableGroupId

`func (o *OswDevCapVO) HasStkableGroupId() bool`

HasStkableGroupId returns a boolean if a field has been set.

### GetStormRateModeSupport

`func (o *OswDevCapVO) GetStormRateModeSupport() bool`

GetStormRateModeSupport returns the StormRateModeSupport field if non-nil, zero value otherwise.

### GetStormRateModeSupportOk

`func (o *OswDevCapVO) GetStormRateModeSupportOk() (*bool, bool)`

GetStormRateModeSupportOk returns a tuple with the StormRateModeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormRateModeSupport

`func (o *OswDevCapVO) SetStormRateModeSupport(v bool)`

SetStormRateModeSupport sets StormRateModeSupport field to given value.

### HasStormRateModeSupport

`func (o *OswDevCapVO) HasStormRateModeSupport() bool`

HasStormRateModeSupport returns a boolean if a field has been set.

### GetStpExtendSupport

`func (o *OswDevCapVO) GetStpExtendSupport() bool`

GetStpExtendSupport returns the StpExtendSupport field if non-nil, zero value otherwise.

### GetStpExtendSupportOk

`func (o *OswDevCapVO) GetStpExtendSupportOk() (*bool, bool)`

GetStpExtendSupportOk returns a tuple with the StpExtendSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpExtendSupport

`func (o *OswDevCapVO) SetStpExtendSupport(v bool)`

SetStpExtendSupport sets StpExtendSupport field to given value.

### HasStpExtendSupport

`func (o *OswDevCapVO) HasStpExtendSupport() bool`

HasStpExtendSupport returns a boolean if a field has been set.

### GetSupportArpDetect

`func (o *OswDevCapVO) GetSupportArpDetect() bool`

GetSupportArpDetect returns the SupportArpDetect field if non-nil, zero value otherwise.

### GetSupportArpDetectOk

`func (o *OswDevCapVO) GetSupportArpDetectOk() (*bool, bool)`

GetSupportArpDetectOk returns a tuple with the SupportArpDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportArpDetect

`func (o *OswDevCapVO) SetSupportArpDetect(v bool)`

SetSupportArpDetect sets SupportArpDetect field to given value.

### HasSupportArpDetect

`func (o *OswDevCapVO) HasSupportArpDetect() bool`

HasSupportArpDetect returns a boolean if a field has been set.

### GetSupportBt

`func (o *OswDevCapVO) GetSupportBt() bool`

GetSupportBt returns the SupportBt field if non-nil, zero value otherwise.

### GetSupportBtOk

`func (o *OswDevCapVO) GetSupportBtOk() (*bool, bool)`

GetSupportBtOk returns a tuple with the SupportBt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBt

`func (o *OswDevCapVO) SetSupportBt(v bool)`

SetSupportBt sets SupportBt field to given value.

### HasSupportBt

`func (o *OswDevCapVO) HasSupportBt() bool`

HasSupportBt returns a boolean if a field has been set.

### GetSupportClearCounters

`func (o *OswDevCapVO) GetSupportClearCounters() bool`

GetSupportClearCounters returns the SupportClearCounters field if non-nil, zero value otherwise.

### GetSupportClearCountersOk

`func (o *OswDevCapVO) GetSupportClearCountersOk() (*bool, bool)`

GetSupportClearCountersOk returns a tuple with the SupportClearCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportClearCounters

`func (o *OswDevCapVO) SetSupportClearCounters(v bool)`

SetSupportClearCounters sets SupportClearCounters field to given value.

### HasSupportClearCounters

`func (o *OswDevCapVO) HasSupportClearCounters() bool`

HasSupportClearCounters returns a boolean if a field has been set.

### GetSupportConfigSync

`func (o *OswDevCapVO) GetSupportConfigSync() bool`

GetSupportConfigSync returns the SupportConfigSync field if non-nil, zero value otherwise.

### GetSupportConfigSyncOk

`func (o *OswDevCapVO) GetSupportConfigSyncOk() (*bool, bool)`

GetSupportConfigSyncOk returns a tuple with the SupportConfigSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportConfigSync

`func (o *OswDevCapVO) SetSupportConfigSync(v bool)`

SetSupportConfigSync sets SupportConfigSync field to given value.

### HasSupportConfigSync

`func (o *OswDevCapVO) HasSupportConfigSync() bool`

HasSupportConfigSync returns a boolean if a field has been set.

### GetSupportDhcpSnoop

`func (o *OswDevCapVO) GetSupportDhcpSnoop() bool`

GetSupportDhcpSnoop returns the SupportDhcpSnoop field if non-nil, zero value otherwise.

### GetSupportDhcpSnoopOk

`func (o *OswDevCapVO) GetSupportDhcpSnoopOk() (*bool, bool)`

GetSupportDhcpSnoopOk returns a tuple with the SupportDhcpSnoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpSnoop

`func (o *OswDevCapVO) SetSupportDhcpSnoop(v bool)`

SetSupportDhcpSnoop sets SupportDhcpSnoop field to given value.

### HasSupportDhcpSnoop

`func (o *OswDevCapVO) HasSupportDhcpSnoop() bool`

HasSupportDhcpSnoop returns a boolean if a field has been set.

### GetSupportEsHealth

`func (o *OswDevCapVO) GetSupportEsHealth() bool`

GetSupportEsHealth returns the SupportEsHealth field if non-nil, zero value otherwise.

### GetSupportEsHealthOk

`func (o *OswDevCapVO) GetSupportEsHealthOk() (*bool, bool)`

GetSupportEsHealthOk returns a tuple with the SupportEsHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEsHealth

`func (o *OswDevCapVO) SetSupportEsHealth(v bool)`

SetSupportEsHealth sets SupportEsHealth field to given value.

### HasSupportEsHealth

`func (o *OswDevCapVO) HasSupportEsHealth() bool`

HasSupportEsHealth returns a boolean if a field has been set.

### GetSupportEsMulticast

`func (o *OswDevCapVO) GetSupportEsMulticast() bool`

GetSupportEsMulticast returns the SupportEsMulticast field if non-nil, zero value otherwise.

### GetSupportEsMulticastOk

`func (o *OswDevCapVO) GetSupportEsMulticastOk() (*bool, bool)`

GetSupportEsMulticastOk returns a tuple with the SupportEsMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEsMulticast

`func (o *OswDevCapVO) SetSupportEsMulticast(v bool)`

SetSupportEsMulticast sets SupportEsMulticast field to given value.

### HasSupportEsMulticast

`func (o *OswDevCapVO) HasSupportEsMulticast() bool`

HasSupportEsMulticast returns a boolean if a field has been set.

### GetSupportFan

`func (o *OswDevCapVO) GetSupportFan() bool`

GetSupportFan returns the SupportFan field if non-nil, zero value otherwise.

### GetSupportFanOk

`func (o *OswDevCapVO) GetSupportFanOk() (*bool, bool)`

GetSupportFanOk returns a tuple with the SupportFan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportFan

`func (o *OswDevCapVO) SetSupportFan(v bool)`

SetSupportFan sets SupportFan field to given value.

### HasSupportFan

`func (o *OswDevCapVO) HasSupportFan() bool`

HasSupportFan returns a boolean if a field has been set.

### GetSupportGetDhcpClientTable

`func (o *OswDevCapVO) GetSupportGetDhcpClientTable() bool`

GetSupportGetDhcpClientTable returns the SupportGetDhcpClientTable field if non-nil, zero value otherwise.

### GetSupportGetDhcpClientTableOk

`func (o *OswDevCapVO) GetSupportGetDhcpClientTableOk() (*bool, bool)`

GetSupportGetDhcpClientTableOk returns a tuple with the SupportGetDhcpClientTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGetDhcpClientTable

`func (o *OswDevCapVO) SetSupportGetDhcpClientTable(v bool)`

SetSupportGetDhcpClientTable sets SupportGetDhcpClientTable field to given value.

### HasSupportGetDhcpClientTable

`func (o *OswDevCapVO) HasSupportGetDhcpClientTable() bool`

HasSupportGetDhcpClientTable returns a boolean if a field has been set.

### GetSupportGetDhcpServerInfo

`func (o *OswDevCapVO) GetSupportGetDhcpServerInfo() bool`

GetSupportGetDhcpServerInfo returns the SupportGetDhcpServerInfo field if non-nil, zero value otherwise.

### GetSupportGetDhcpServerInfoOk

`func (o *OswDevCapVO) GetSupportGetDhcpServerInfoOk() (*bool, bool)`

GetSupportGetDhcpServerInfoOk returns a tuple with the SupportGetDhcpServerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGetDhcpServerInfo

`func (o *OswDevCapVO) SetSupportGetDhcpServerInfo(v bool)`

SetSupportGetDhcpServerInfo sets SupportGetDhcpServerInfo field to given value.

### HasSupportGetDhcpServerInfo

`func (o *OswDevCapVO) HasSupportGetDhcpServerInfo() bool`

HasSupportGetDhcpServerInfo returns a boolean if a field has been set.

### GetSupportGetOspfNeighborTable

`func (o *OswDevCapVO) GetSupportGetOspfNeighborTable() bool`

GetSupportGetOspfNeighborTable returns the SupportGetOspfNeighborTable field if non-nil, zero value otherwise.

### GetSupportGetOspfNeighborTableOk

`func (o *OswDevCapVO) GetSupportGetOspfNeighborTableOk() (*bool, bool)`

GetSupportGetOspfNeighborTableOk returns a tuple with the SupportGetOspfNeighborTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGetOspfNeighborTable

`func (o *OswDevCapVO) SetSupportGetOspfNeighborTable(v bool)`

SetSupportGetOspfNeighborTable sets SupportGetOspfNeighborTable field to given value.

### HasSupportGetOspfNeighborTable

`func (o *OswDevCapVO) HasSupportGetOspfNeighborTable() bool`

HasSupportGetOspfNeighborTable returns a boolean if a field has been set.

### GetSupportGetVlanIfInfo

`func (o *OswDevCapVO) GetSupportGetVlanIfInfo() bool`

GetSupportGetVlanIfInfo returns the SupportGetVlanIfInfo field if non-nil, zero value otherwise.

### GetSupportGetVlanIfInfoOk

`func (o *OswDevCapVO) GetSupportGetVlanIfInfoOk() (*bool, bool)`

GetSupportGetVlanIfInfoOk returns a tuple with the SupportGetVlanIfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGetVlanIfInfo

`func (o *OswDevCapVO) SetSupportGetVlanIfInfo(v bool)`

SetSupportGetVlanIfInfo sets SupportGetVlanIfInfo field to given value.

### HasSupportGetVlanIfInfo

`func (o *OswDevCapVO) HasSupportGetVlanIfInfo() bool`

HasSupportGetVlanIfInfo returns a boolean if a field has been set.

### GetSupportImpb

`func (o *OswDevCapVO) GetSupportImpb() bool`

GetSupportImpb returns the SupportImpb field if non-nil, zero value otherwise.

### GetSupportImpbOk

`func (o *OswDevCapVO) GetSupportImpbOk() (*bool, bool)`

GetSupportImpbOk returns a tuple with the SupportImpb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportImpb

`func (o *OswDevCapVO) SetSupportImpb(v bool)`

SetSupportImpb sets SupportImpb field to given value.

### HasSupportImpb

`func (o *OswDevCapVO) HasSupportImpb() bool`

HasSupportImpb returns a boolean if a field has been set.

### GetSupportMad

`func (o *OswDevCapVO) GetSupportMad() bool`

GetSupportMad returns the SupportMad field if non-nil, zero value otherwise.

### GetSupportMadOk

`func (o *OswDevCapVO) GetSupportMadOk() (*bool, bool)`

GetSupportMadOk returns a tuple with the SupportMad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMad

`func (o *OswDevCapVO) SetSupportMad(v bool)`

SetSupportMad sets SupportMad field to given value.

### HasSupportMad

`func (o *OswDevCapVO) HasSupportMad() bool`

HasSupportMad returns a boolean if a field has been set.

### GetSupportMlag

`func (o *OswDevCapVO) GetSupportMlag() bool`

GetSupportMlag returns the SupportMlag field if non-nil, zero value otherwise.

### GetSupportMlagOk

`func (o *OswDevCapVO) GetSupportMlagOk() (*bool, bool)`

GetSupportMlagOk returns a tuple with the SupportMlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMlag

`func (o *OswDevCapVO) SetSupportMlag(v bool)`

SetSupportMlag sets SupportMlag field to given value.

### HasSupportMlag

`func (o *OswDevCapVO) HasSupportMlag() bool`

HasSupportMlag returns a boolean if a field has been set.

### GetSupportMulticast

`func (o *OswDevCapVO) GetSupportMulticast() bool`

GetSupportMulticast returns the SupportMulticast field if non-nil, zero value otherwise.

### GetSupportMulticastOk

`func (o *OswDevCapVO) GetSupportMulticastOk() (*bool, bool)`

GetSupportMulticastOk returns a tuple with the SupportMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMulticast

`func (o *OswDevCapVO) SetSupportMulticast(v bool)`

SetSupportMulticast sets SupportMulticast field to given value.

### HasSupportMulticast

`func (o *OswDevCapVO) HasSupportMulticast() bool`

HasSupportMulticast returns a boolean if a field has been set.

### GetSupportRelayMultiServer

`func (o *OswDevCapVO) GetSupportRelayMultiServer() bool`

GetSupportRelayMultiServer returns the SupportRelayMultiServer field if non-nil, zero value otherwise.

### GetSupportRelayMultiServerOk

`func (o *OswDevCapVO) GetSupportRelayMultiServerOk() (*bool, bool)`

GetSupportRelayMultiServerOk returns a tuple with the SupportRelayMultiServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRelayMultiServer

`func (o *OswDevCapVO) SetSupportRelayMultiServer(v bool)`

SetSupportRelayMultiServer sets SupportRelayMultiServer field to given value.

### HasSupportRelayMultiServer

`func (o *OswDevCapVO) HasSupportRelayMultiServer() bool`

HasSupportRelayMultiServer returns a boolean if a field has been set.

### GetSupportRunningConfig

`func (o *OswDevCapVO) GetSupportRunningConfig() bool`

GetSupportRunningConfig returns the SupportRunningConfig field if non-nil, zero value otherwise.

### GetSupportRunningConfigOk

`func (o *OswDevCapVO) GetSupportRunningConfigOk() (*bool, bool)`

GetSupportRunningConfigOk returns a tuple with the SupportRunningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRunningConfig

`func (o *OswDevCapVO) SetSupportRunningConfig(v bool)`

SetSupportRunningConfig sets SupportRunningConfig field to given value.

### HasSupportRunningConfig

`func (o *OswDevCapVO) HasSupportRunningConfig() bool`

HasSupportRunningConfig returns a boolean if a field has been set.

### GetSupportSdm

`func (o *OswDevCapVO) GetSupportSdm() bool`

GetSupportSdm returns the SupportSdm field if non-nil, zero value otherwise.

### GetSupportSdmOk

`func (o *OswDevCapVO) GetSupportSdmOk() (*bool, bool)`

GetSupportSdmOk returns a tuple with the SupportSdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSdm

`func (o *OswDevCapVO) SetSupportSdm(v bool)`

SetSupportSdm sets SupportSdm field to given value.

### HasSupportSdm

`func (o *OswDevCapVO) HasSupportSdm() bool`

HasSupportSdm returns a boolean if a field has been set.

### GetSupportVrf

`func (o *OswDevCapVO) GetSupportVrf() bool`

GetSupportVrf returns the SupportVrf field if non-nil, zero value otherwise.

### GetSupportVrfOk

`func (o *OswDevCapVO) GetSupportVrfOk() (*bool, bool)`

GetSupportVrfOk returns a tuple with the SupportVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVrf

`func (o *OswDevCapVO) SetSupportVrf(v bool)`

SetSupportVrf sets SupportVrf field to given value.

### HasSupportVrf

`func (o *OswDevCapVO) HasSupportVrf() bool`

HasSupportVrf returns a boolean if a field has been set.

### GetTerminalSupport

`func (o *OswDevCapVO) GetTerminalSupport() bool`

GetTerminalSupport returns the TerminalSupport field if non-nil, zero value otherwise.

### GetTerminalSupportOk

`func (o *OswDevCapVO) GetTerminalSupportOk() (*bool, bool)`

GetTerminalSupportOk returns a tuple with the TerminalSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalSupport

`func (o *OswDevCapVO) SetTerminalSupport(v bool)`

SetTerminalSupport sets TerminalSupport field to given value.

### HasTerminalSupport

`func (o *OswDevCapVO) HasTerminalSupport() bool`

HasTerminalSupport returns a boolean if a field has been set.

### GetTracerouteSupport

`func (o *OswDevCapVO) GetTracerouteSupport() bool`

GetTracerouteSupport returns the TracerouteSupport field if non-nil, zero value otherwise.

### GetTracerouteSupportOk

`func (o *OswDevCapVO) GetTracerouteSupportOk() (*bool, bool)`

GetTracerouteSupportOk returns a tuple with the TracerouteSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracerouteSupport

`func (o *OswDevCapVO) SetTracerouteSupport(v bool)`

SetTracerouteSupport sets TracerouteSupport field to given value.

### HasTracerouteSupport

`func (o *OswDevCapVO) HasTracerouteSupport() bool`

HasTracerouteSupport returns a boolean if a field has been set.

### GetUplinkSupport

`func (o *OswDevCapVO) GetUplinkSupport() bool`

GetUplinkSupport returns the UplinkSupport field if non-nil, zero value otherwise.

### GetUplinkSupportOk

`func (o *OswDevCapVO) GetUplinkSupportOk() (*bool, bool)`

GetUplinkSupportOk returns a tuple with the UplinkSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSupport

`func (o *OswDevCapVO) SetUplinkSupport(v bool)`

SetUplinkSupport sets UplinkSupport field to given value.

### HasUplinkSupport

`func (o *OswDevCapVO) HasUplinkSupport() bool`

HasUplinkSupport returns a boolean if a field has been set.

### GetVlanIfNum

`func (o *OswDevCapVO) GetVlanIfNum() int32`

GetVlanIfNum returns the VlanIfNum field if non-nil, zero value otherwise.

### GetVlanIfNumOk

`func (o *OswDevCapVO) GetVlanIfNumOk() (*int32, bool)`

GetVlanIfNumOk returns a tuple with the VlanIfNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIfNum

`func (o *OswDevCapVO) SetVlanIfNum(v int32)`

SetVlanIfNum sets VlanIfNum field to given value.

### HasVlanIfNum

`func (o *OswDevCapVO) HasVlanIfNum() bool`

HasVlanIfNum returns a boolean if a field has been set.

### GetVoiceDscpSupport

`func (o *OswDevCapVO) GetVoiceDscpSupport() bool`

GetVoiceDscpSupport returns the VoiceDscpSupport field if non-nil, zero value otherwise.

### GetVoiceDscpSupportOk

`func (o *OswDevCapVO) GetVoiceDscpSupportOk() (*bool, bool)`

GetVoiceDscpSupportOk returns a tuple with the VoiceDscpSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDscpSupport

`func (o *OswDevCapVO) SetVoiceDscpSupport(v bool)`

SetVoiceDscpSupport sets VoiceDscpSupport field to given value.

### HasVoiceDscpSupport

`func (o *OswDevCapVO) HasVoiceDscpSupport() bool`

HasVoiceDscpSupport returns a boolean if a field has been set.

### GetVrfNum

`func (o *OswDevCapVO) GetVrfNum() int32`

GetVrfNum returns the VrfNum field if non-nil, zero value otherwise.

### GetVrfNumOk

`func (o *OswDevCapVO) GetVrfNumOk() (*int32, bool)`

GetVrfNumOk returns a tuple with the VrfNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfNum

`func (o *OswDevCapVO) SetVrfNum(v int32)`

SetVrfNum sets VrfNum field to given value.

### HasVrfNum

`func (o *OswDevCapVO) HasVrfNum() bool`

HasVrfNum returns a boolean if a field has been set.

### GetVrrpSupport

`func (o *OswDevCapVO) GetVrrpSupport() bool`

GetVrrpSupport returns the VrrpSupport field if non-nil, zero value otherwise.

### GetVrrpSupportOk

`func (o *OswDevCapVO) GetVrrpSupportOk() (*bool, bool)`

GetVrrpSupportOk returns a tuple with the VrrpSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpSupport

`func (o *OswDevCapVO) SetVrrpSupport(v bool)`

SetVrrpSupport sets VrrpSupport field to given value.

### HasVrrpSupport

`func (o *OswDevCapVO) HasVrrpSupport() bool`

HasVrrpSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


