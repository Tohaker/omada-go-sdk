# OsgPortConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailablePvidNames** | Pointer to [**[]OsgPvIdNameVO**](OsgPvIdNameVO.md) |  | [optional] 
**AvailablePvids** | Pointer to **[]int32** |  | [optional] 
**BandCtrl** | Pointer to [**OsgPortBandCtrlVO**](OsgPortBandCtrlVO.md) |  | [optional] 
**BandWidthCtrlType** | Pointer to **int32** |  | [optional] 
**DslSettings** | Pointer to [**DslSettingsVO**](DslSettingsVO.md) |  | [optional] 
**Duplex** | Pointer to **int32** |  | [optional] 
**FlowControl** | Pointer to **bool** |  | [optional] 
**LinkSpeed** | Pointer to **int32** |  | [optional] 
**LoopbackControl** | Pointer to **int32** |  | [optional] 
**MirrorEnable** | Pointer to **bool** |  | [optional] 
**MirrorMode** | Pointer to **int32** |  | [optional] 
**MirroredPorts** | Pointer to **[]int32** |  | [optional] 
**PoeMode** | Pointer to **int32** |  | [optional] 
**Port** | **int32** |  | 
**PortCap** | Pointer to [**[]OsgLinkVO**](OsgLinkVO.md) |  | [optional] 
**PortIsolationEnable** | Pointer to **bool** |  | [optional] 
**PortList** | Pointer to **[]int32** |  | [optional] 
**PortStat** | Pointer to [**OsgPortStatVO**](OsgPortStatVO.md) |  | [optional] 
**PortSupportPoe** | Pointer to **bool** |  | [optional] 
**Pvid** | Pointer to **int32** |  | [optional] 
**Resource** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StormCtrl** | Pointer to [**OsgPortStormCtrlVO**](OsgPortStormCtrlVO.md) |  | [optional] 
**SupportBandWidthCtrl** | Pointer to **bool** |  | [optional] 
**SupportFlowControl** | Pointer to **bool** |  | [optional] 
**SupportLoopbackControl** | Pointer to **bool** |  | [optional] 
**SupportMirror** | Pointer to **bool** |  | [optional] 
**SupportPortControl** | Pointer to **bool** |  | [optional] 
**SupportPortIsolation** | Pointer to **bool** |  | [optional] 

## Methods

### NewOsgPortConfigVO

`func NewOsgPortConfigVO(port int32, ) *OsgPortConfigVO`

NewOsgPortConfigVO instantiates a new OsgPortConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgPortConfigVOWithDefaults

`func NewOsgPortConfigVOWithDefaults() *OsgPortConfigVO`

NewOsgPortConfigVOWithDefaults instantiates a new OsgPortConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailablePvidNames

`func (o *OsgPortConfigVO) GetAvailablePvidNames() []OsgPvIdNameVO`

GetAvailablePvidNames returns the AvailablePvidNames field if non-nil, zero value otherwise.

### GetAvailablePvidNamesOk

`func (o *OsgPortConfigVO) GetAvailablePvidNamesOk() (*[]OsgPvIdNameVO, bool)`

GetAvailablePvidNamesOk returns a tuple with the AvailablePvidNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePvidNames

`func (o *OsgPortConfigVO) SetAvailablePvidNames(v []OsgPvIdNameVO)`

SetAvailablePvidNames sets AvailablePvidNames field to given value.

### HasAvailablePvidNames

`func (o *OsgPortConfigVO) HasAvailablePvidNames() bool`

HasAvailablePvidNames returns a boolean if a field has been set.

### GetAvailablePvids

`func (o *OsgPortConfigVO) GetAvailablePvids() []int32`

GetAvailablePvids returns the AvailablePvids field if non-nil, zero value otherwise.

### GetAvailablePvidsOk

`func (o *OsgPortConfigVO) GetAvailablePvidsOk() (*[]int32, bool)`

GetAvailablePvidsOk returns a tuple with the AvailablePvids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePvids

`func (o *OsgPortConfigVO) SetAvailablePvids(v []int32)`

SetAvailablePvids sets AvailablePvids field to given value.

### HasAvailablePvids

`func (o *OsgPortConfigVO) HasAvailablePvids() bool`

HasAvailablePvids returns a boolean if a field has been set.

### GetBandCtrl

`func (o *OsgPortConfigVO) GetBandCtrl() OsgPortBandCtrlVO`

GetBandCtrl returns the BandCtrl field if non-nil, zero value otherwise.

### GetBandCtrlOk

`func (o *OsgPortConfigVO) GetBandCtrlOk() (*OsgPortBandCtrlVO, bool)`

GetBandCtrlOk returns a tuple with the BandCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandCtrl

`func (o *OsgPortConfigVO) SetBandCtrl(v OsgPortBandCtrlVO)`

SetBandCtrl sets BandCtrl field to given value.

### HasBandCtrl

`func (o *OsgPortConfigVO) HasBandCtrl() bool`

HasBandCtrl returns a boolean if a field has been set.

### GetBandWidthCtrlType

`func (o *OsgPortConfigVO) GetBandWidthCtrlType() int32`

GetBandWidthCtrlType returns the BandWidthCtrlType field if non-nil, zero value otherwise.

### GetBandWidthCtrlTypeOk

`func (o *OsgPortConfigVO) GetBandWidthCtrlTypeOk() (*int32, bool)`

GetBandWidthCtrlTypeOk returns a tuple with the BandWidthCtrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidthCtrlType

`func (o *OsgPortConfigVO) SetBandWidthCtrlType(v int32)`

SetBandWidthCtrlType sets BandWidthCtrlType field to given value.

### HasBandWidthCtrlType

`func (o *OsgPortConfigVO) HasBandWidthCtrlType() bool`

HasBandWidthCtrlType returns a boolean if a field has been set.

### GetDslSettings

`func (o *OsgPortConfigVO) GetDslSettings() DslSettingsVO`

GetDslSettings returns the DslSettings field if non-nil, zero value otherwise.

### GetDslSettingsOk

`func (o *OsgPortConfigVO) GetDslSettingsOk() (*DslSettingsVO, bool)`

GetDslSettingsOk returns a tuple with the DslSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslSettings

`func (o *OsgPortConfigVO) SetDslSettings(v DslSettingsVO)`

SetDslSettings sets DslSettings field to given value.

### HasDslSettings

`func (o *OsgPortConfigVO) HasDslSettings() bool`

HasDslSettings returns a boolean if a field has been set.

### GetDuplex

`func (o *OsgPortConfigVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OsgPortConfigVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OsgPortConfigVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OsgPortConfigVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetFlowControl

`func (o *OsgPortConfigVO) GetFlowControl() bool`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *OsgPortConfigVO) GetFlowControlOk() (*bool, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *OsgPortConfigVO) SetFlowControl(v bool)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *OsgPortConfigVO) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OsgPortConfigVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OsgPortConfigVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OsgPortConfigVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OsgPortConfigVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLoopbackControl

`func (o *OsgPortConfigVO) GetLoopbackControl() int32`

GetLoopbackControl returns the LoopbackControl field if non-nil, zero value otherwise.

### GetLoopbackControlOk

`func (o *OsgPortConfigVO) GetLoopbackControlOk() (*int32, bool)`

GetLoopbackControlOk returns a tuple with the LoopbackControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackControl

`func (o *OsgPortConfigVO) SetLoopbackControl(v int32)`

SetLoopbackControl sets LoopbackControl field to given value.

### HasLoopbackControl

`func (o *OsgPortConfigVO) HasLoopbackControl() bool`

HasLoopbackControl returns a boolean if a field has been set.

### GetMirrorEnable

`func (o *OsgPortConfigVO) GetMirrorEnable() bool`

GetMirrorEnable returns the MirrorEnable field if non-nil, zero value otherwise.

### GetMirrorEnableOk

`func (o *OsgPortConfigVO) GetMirrorEnableOk() (*bool, bool)`

GetMirrorEnableOk returns a tuple with the MirrorEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorEnable

`func (o *OsgPortConfigVO) SetMirrorEnable(v bool)`

SetMirrorEnable sets MirrorEnable field to given value.

### HasMirrorEnable

`func (o *OsgPortConfigVO) HasMirrorEnable() bool`

HasMirrorEnable returns a boolean if a field has been set.

### GetMirrorMode

`func (o *OsgPortConfigVO) GetMirrorMode() int32`

GetMirrorMode returns the MirrorMode field if non-nil, zero value otherwise.

### GetMirrorModeOk

`func (o *OsgPortConfigVO) GetMirrorModeOk() (*int32, bool)`

GetMirrorModeOk returns a tuple with the MirrorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorMode

`func (o *OsgPortConfigVO) SetMirrorMode(v int32)`

SetMirrorMode sets MirrorMode field to given value.

### HasMirrorMode

`func (o *OsgPortConfigVO) HasMirrorMode() bool`

HasMirrorMode returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *OsgPortConfigVO) GetMirroredPorts() []int32`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *OsgPortConfigVO) GetMirroredPortsOk() (*[]int32, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *OsgPortConfigVO) SetMirroredPorts(v []int32)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *OsgPortConfigVO) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetPoeMode

`func (o *OsgPortConfigVO) GetPoeMode() int32`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *OsgPortConfigVO) GetPoeModeOk() (*int32, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *OsgPortConfigVO) SetPoeMode(v int32)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *OsgPortConfigVO) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### GetPort

`func (o *OsgPortConfigVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OsgPortConfigVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OsgPortConfigVO) SetPort(v int32)`

SetPort sets Port field to given value.


### GetPortCap

`func (o *OsgPortConfigVO) GetPortCap() []OsgLinkVO`

GetPortCap returns the PortCap field if non-nil, zero value otherwise.

### GetPortCapOk

`func (o *OsgPortConfigVO) GetPortCapOk() (*[]OsgLinkVO, bool)`

GetPortCapOk returns a tuple with the PortCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCap

`func (o *OsgPortConfigVO) SetPortCap(v []OsgLinkVO)`

SetPortCap sets PortCap field to given value.

### HasPortCap

`func (o *OsgPortConfigVO) HasPortCap() bool`

HasPortCap returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *OsgPortConfigVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *OsgPortConfigVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *OsgPortConfigVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *OsgPortConfigVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetPortList

`func (o *OsgPortConfigVO) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *OsgPortConfigVO) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *OsgPortConfigVO) SetPortList(v []int32)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *OsgPortConfigVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetPortStat

`func (o *OsgPortConfigVO) GetPortStat() OsgPortStatVO`

GetPortStat returns the PortStat field if non-nil, zero value otherwise.

### GetPortStatOk

`func (o *OsgPortConfigVO) GetPortStatOk() (*OsgPortStatVO, bool)`

GetPortStatOk returns a tuple with the PortStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStat

`func (o *OsgPortConfigVO) SetPortStat(v OsgPortStatVO)`

SetPortStat sets PortStat field to given value.

### HasPortStat

`func (o *OsgPortConfigVO) HasPortStat() bool`

HasPortStat returns a boolean if a field has been set.

### GetPortSupportPoe

`func (o *OsgPortConfigVO) GetPortSupportPoe() bool`

GetPortSupportPoe returns the PortSupportPoe field if non-nil, zero value otherwise.

### GetPortSupportPoeOk

`func (o *OsgPortConfigVO) GetPortSupportPoeOk() (*bool, bool)`

GetPortSupportPoeOk returns a tuple with the PortSupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSupportPoe

`func (o *OsgPortConfigVO) SetPortSupportPoe(v bool)`

SetPortSupportPoe sets PortSupportPoe field to given value.

### HasPortSupportPoe

`func (o *OsgPortConfigVO) HasPortSupportPoe() bool`

HasPortSupportPoe returns a boolean if a field has been set.

### GetPvid

`func (o *OsgPortConfigVO) GetPvid() int32`

GetPvid returns the Pvid field if non-nil, zero value otherwise.

### GetPvidOk

`func (o *OsgPortConfigVO) GetPvidOk() (*int32, bool)`

GetPvidOk returns a tuple with the Pvid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvid

`func (o *OsgPortConfigVO) SetPvid(v int32)`

SetPvid sets Pvid field to given value.

### HasPvid

`func (o *OsgPortConfigVO) HasPvid() bool`

HasPvid returns a boolean if a field has been set.

### GetResource

`func (o *OsgPortConfigVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OsgPortConfigVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OsgPortConfigVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OsgPortConfigVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStatus

`func (o *OsgPortConfigVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsgPortConfigVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsgPortConfigVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsgPortConfigVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStormCtrl

`func (o *OsgPortConfigVO) GetStormCtrl() OsgPortStormCtrlVO`

GetStormCtrl returns the StormCtrl field if non-nil, zero value otherwise.

### GetStormCtrlOk

`func (o *OsgPortConfigVO) GetStormCtrlOk() (*OsgPortStormCtrlVO, bool)`

GetStormCtrlOk returns a tuple with the StormCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormCtrl

`func (o *OsgPortConfigVO) SetStormCtrl(v OsgPortStormCtrlVO)`

SetStormCtrl sets StormCtrl field to given value.

### HasStormCtrl

`func (o *OsgPortConfigVO) HasStormCtrl() bool`

HasStormCtrl returns a boolean if a field has been set.

### GetSupportBandWidthCtrl

`func (o *OsgPortConfigVO) GetSupportBandWidthCtrl() bool`

GetSupportBandWidthCtrl returns the SupportBandWidthCtrl field if non-nil, zero value otherwise.

### GetSupportBandWidthCtrlOk

`func (o *OsgPortConfigVO) GetSupportBandWidthCtrlOk() (*bool, bool)`

GetSupportBandWidthCtrlOk returns a tuple with the SupportBandWidthCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBandWidthCtrl

`func (o *OsgPortConfigVO) SetSupportBandWidthCtrl(v bool)`

SetSupportBandWidthCtrl sets SupportBandWidthCtrl field to given value.

### HasSupportBandWidthCtrl

`func (o *OsgPortConfigVO) HasSupportBandWidthCtrl() bool`

HasSupportBandWidthCtrl returns a boolean if a field has been set.

### GetSupportFlowControl

`func (o *OsgPortConfigVO) GetSupportFlowControl() bool`

GetSupportFlowControl returns the SupportFlowControl field if non-nil, zero value otherwise.

### GetSupportFlowControlOk

`func (o *OsgPortConfigVO) GetSupportFlowControlOk() (*bool, bool)`

GetSupportFlowControlOk returns a tuple with the SupportFlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportFlowControl

`func (o *OsgPortConfigVO) SetSupportFlowControl(v bool)`

SetSupportFlowControl sets SupportFlowControl field to given value.

### HasSupportFlowControl

`func (o *OsgPortConfigVO) HasSupportFlowControl() bool`

HasSupportFlowControl returns a boolean if a field has been set.

### GetSupportLoopbackControl

`func (o *OsgPortConfigVO) GetSupportLoopbackControl() bool`

GetSupportLoopbackControl returns the SupportLoopbackControl field if non-nil, zero value otherwise.

### GetSupportLoopbackControlOk

`func (o *OsgPortConfigVO) GetSupportLoopbackControlOk() (*bool, bool)`

GetSupportLoopbackControlOk returns a tuple with the SupportLoopbackControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLoopbackControl

`func (o *OsgPortConfigVO) SetSupportLoopbackControl(v bool)`

SetSupportLoopbackControl sets SupportLoopbackControl field to given value.

### HasSupportLoopbackControl

`func (o *OsgPortConfigVO) HasSupportLoopbackControl() bool`

HasSupportLoopbackControl returns a boolean if a field has been set.

### GetSupportMirror

`func (o *OsgPortConfigVO) GetSupportMirror() bool`

GetSupportMirror returns the SupportMirror field if non-nil, zero value otherwise.

### GetSupportMirrorOk

`func (o *OsgPortConfigVO) GetSupportMirrorOk() (*bool, bool)`

GetSupportMirrorOk returns a tuple with the SupportMirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMirror

`func (o *OsgPortConfigVO) SetSupportMirror(v bool)`

SetSupportMirror sets SupportMirror field to given value.

### HasSupportMirror

`func (o *OsgPortConfigVO) HasSupportMirror() bool`

HasSupportMirror returns a boolean if a field has been set.

### GetSupportPortControl

`func (o *OsgPortConfigVO) GetSupportPortControl() bool`

GetSupportPortControl returns the SupportPortControl field if non-nil, zero value otherwise.

### GetSupportPortControlOk

`func (o *OsgPortConfigVO) GetSupportPortControlOk() (*bool, bool)`

GetSupportPortControlOk returns a tuple with the SupportPortControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPortControl

`func (o *OsgPortConfigVO) SetSupportPortControl(v bool)`

SetSupportPortControl sets SupportPortControl field to given value.

### HasSupportPortControl

`func (o *OsgPortConfigVO) HasSupportPortControl() bool`

HasSupportPortControl returns a boolean if a field has been set.

### GetSupportPortIsolation

`func (o *OsgPortConfigVO) GetSupportPortIsolation() bool`

GetSupportPortIsolation returns the SupportPortIsolation field if non-nil, zero value otherwise.

### GetSupportPortIsolationOk

`func (o *OsgPortConfigVO) GetSupportPortIsolationOk() (*bool, bool)`

GetSupportPortIsolationOk returns a tuple with the SupportPortIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPortIsolation

`func (o *OsgPortConfigVO) SetSupportPortIsolation(v bool)`

SetSupportPortIsolation sets SupportPortIsolation field to given value.

### HasSupportPortIsolation

`func (o *OsgPortConfigVO) HasSupportPortIsolation() bool`

HasSupportPortIsolation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


