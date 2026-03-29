# SpanningTreeSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BpduFilter** | Pointer to **bool** | Indicates whether bpdu filter is enabled | [optional] 
**BpduForward** | Pointer to **bool** | Indicates whether bpdu forward is enabled | [optional] 
**BpduProtect** | Pointer to **bool** | Indicates whether bpdu protect is enabled | [optional] 
**EdgePort** | **bool** | Indicates whether edge port is enabled | 
**ExtPathCost** | **int32** | ExtPathCost should be within the range of 0–2000000 | 
**IntPathCost** | **int32** | IntPathCost should be within the range of 0–2000000 | 
**LoopProtect** | Pointer to **bool** | Indicates whether loop protect is enabled | [optional] 
**Mcheck** | Pointer to **bool** | Indicates whether mcheck is enabled | [optional] 
**P2pLink** | **int32** | P2pLink should be within the range of 0-2 | 
**Priority** | **int32** | Priority should be within the range of 0–240 | 
**RootProtect** | Pointer to **bool** | Indicates whether root protect is enabled | [optional] 
**TcGuard** | Pointer to **bool** | Indicates whether tcGuard is enabled | [optional] 

## Methods

### NewSpanningTreeSettingVO

`func NewSpanningTreeSettingVO(edgePort bool, extPathCost int32, intPathCost int32, p2pLink int32, priority int32, ) *SpanningTreeSettingVO`

NewSpanningTreeSettingVO instantiates a new SpanningTreeSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanningTreeSettingVOWithDefaults

`func NewSpanningTreeSettingVOWithDefaults() *SpanningTreeSettingVO`

NewSpanningTreeSettingVOWithDefaults instantiates a new SpanningTreeSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBpduFilter

`func (o *SpanningTreeSettingVO) GetBpduFilter() bool`

GetBpduFilter returns the BpduFilter field if non-nil, zero value otherwise.

### GetBpduFilterOk

`func (o *SpanningTreeSettingVO) GetBpduFilterOk() (*bool, bool)`

GetBpduFilterOk returns a tuple with the BpduFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpduFilter

`func (o *SpanningTreeSettingVO) SetBpduFilter(v bool)`

SetBpduFilter sets BpduFilter field to given value.

### HasBpduFilter

`func (o *SpanningTreeSettingVO) HasBpduFilter() bool`

HasBpduFilter returns a boolean if a field has been set.

### GetBpduForward

`func (o *SpanningTreeSettingVO) GetBpduForward() bool`

GetBpduForward returns the BpduForward field if non-nil, zero value otherwise.

### GetBpduForwardOk

`func (o *SpanningTreeSettingVO) GetBpduForwardOk() (*bool, bool)`

GetBpduForwardOk returns a tuple with the BpduForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpduForward

`func (o *SpanningTreeSettingVO) SetBpduForward(v bool)`

SetBpduForward sets BpduForward field to given value.

### HasBpduForward

`func (o *SpanningTreeSettingVO) HasBpduForward() bool`

HasBpduForward returns a boolean if a field has been set.

### GetBpduProtect

`func (o *SpanningTreeSettingVO) GetBpduProtect() bool`

GetBpduProtect returns the BpduProtect field if non-nil, zero value otherwise.

### GetBpduProtectOk

`func (o *SpanningTreeSettingVO) GetBpduProtectOk() (*bool, bool)`

GetBpduProtectOk returns a tuple with the BpduProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpduProtect

`func (o *SpanningTreeSettingVO) SetBpduProtect(v bool)`

SetBpduProtect sets BpduProtect field to given value.

### HasBpduProtect

`func (o *SpanningTreeSettingVO) HasBpduProtect() bool`

HasBpduProtect returns a boolean if a field has been set.

### GetEdgePort

`func (o *SpanningTreeSettingVO) GetEdgePort() bool`

GetEdgePort returns the EdgePort field if non-nil, zero value otherwise.

### GetEdgePortOk

`func (o *SpanningTreeSettingVO) GetEdgePortOk() (*bool, bool)`

GetEdgePortOk returns a tuple with the EdgePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgePort

`func (o *SpanningTreeSettingVO) SetEdgePort(v bool)`

SetEdgePort sets EdgePort field to given value.


### GetExtPathCost

`func (o *SpanningTreeSettingVO) GetExtPathCost() int32`

GetExtPathCost returns the ExtPathCost field if non-nil, zero value otherwise.

### GetExtPathCostOk

`func (o *SpanningTreeSettingVO) GetExtPathCostOk() (*int32, bool)`

GetExtPathCostOk returns a tuple with the ExtPathCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtPathCost

`func (o *SpanningTreeSettingVO) SetExtPathCost(v int32)`

SetExtPathCost sets ExtPathCost field to given value.


### GetIntPathCost

`func (o *SpanningTreeSettingVO) GetIntPathCost() int32`

GetIntPathCost returns the IntPathCost field if non-nil, zero value otherwise.

### GetIntPathCostOk

`func (o *SpanningTreeSettingVO) GetIntPathCostOk() (*int32, bool)`

GetIntPathCostOk returns a tuple with the IntPathCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntPathCost

`func (o *SpanningTreeSettingVO) SetIntPathCost(v int32)`

SetIntPathCost sets IntPathCost field to given value.


### GetLoopProtect

`func (o *SpanningTreeSettingVO) GetLoopProtect() bool`

GetLoopProtect returns the LoopProtect field if non-nil, zero value otherwise.

### GetLoopProtectOk

`func (o *SpanningTreeSettingVO) GetLoopProtectOk() (*bool, bool)`

GetLoopProtectOk returns a tuple with the LoopProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopProtect

`func (o *SpanningTreeSettingVO) SetLoopProtect(v bool)`

SetLoopProtect sets LoopProtect field to given value.

### HasLoopProtect

`func (o *SpanningTreeSettingVO) HasLoopProtect() bool`

HasLoopProtect returns a boolean if a field has been set.

### GetMcheck

`func (o *SpanningTreeSettingVO) GetMcheck() bool`

GetMcheck returns the Mcheck field if non-nil, zero value otherwise.

### GetMcheckOk

`func (o *SpanningTreeSettingVO) GetMcheckOk() (*bool, bool)`

GetMcheckOk returns a tuple with the Mcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcheck

`func (o *SpanningTreeSettingVO) SetMcheck(v bool)`

SetMcheck sets Mcheck field to given value.

### HasMcheck

`func (o *SpanningTreeSettingVO) HasMcheck() bool`

HasMcheck returns a boolean if a field has been set.

### GetP2pLink

`func (o *SpanningTreeSettingVO) GetP2pLink() int32`

GetP2pLink returns the P2pLink field if non-nil, zero value otherwise.

### GetP2pLinkOk

`func (o *SpanningTreeSettingVO) GetP2pLinkOk() (*int32, bool)`

GetP2pLinkOk returns a tuple with the P2pLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pLink

`func (o *SpanningTreeSettingVO) SetP2pLink(v int32)`

SetP2pLink sets P2pLink field to given value.


### GetPriority

`func (o *SpanningTreeSettingVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SpanningTreeSettingVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SpanningTreeSettingVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetRootProtect

`func (o *SpanningTreeSettingVO) GetRootProtect() bool`

GetRootProtect returns the RootProtect field if non-nil, zero value otherwise.

### GetRootProtectOk

`func (o *SpanningTreeSettingVO) GetRootProtectOk() (*bool, bool)`

GetRootProtectOk returns a tuple with the RootProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProtect

`func (o *SpanningTreeSettingVO) SetRootProtect(v bool)`

SetRootProtect sets RootProtect field to given value.

### HasRootProtect

`func (o *SpanningTreeSettingVO) HasRootProtect() bool`

HasRootProtect returns a boolean if a field has been set.

### GetTcGuard

`func (o *SpanningTreeSettingVO) GetTcGuard() bool`

GetTcGuard returns the TcGuard field if non-nil, zero value otherwise.

### GetTcGuardOk

`func (o *SpanningTreeSettingVO) GetTcGuardOk() (*bool, bool)`

GetTcGuardOk returns a tuple with the TcGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcGuard

`func (o *SpanningTreeSettingVO) SetTcGuard(v bool)`

SetTcGuard sets TcGuard field to given value.

### HasTcGuard

`func (o *SpanningTreeSettingVO) HasTcGuard() bool`

HasTcGuard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


