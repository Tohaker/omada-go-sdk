# GatewayPortConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DslSettings** | Pointer to [**DslSettings**](DslSettings.md) |  | [optional] 
**Duplex** | Pointer to **int32** | Port duplex mode should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**FlowControl** | Pointer to **bool** | Enable flow control or not.(When the port supports flow control.) | [optional] 
**LinkSpeed** | Pointer to **int32** | Port link speed should be a value as follows: 0: Auto; 1: 10M; 2: 100M; 3: 1000M; 4: 2500M; 5: 10G; 6: 5G | [optional] 
**MirrorEnable** | Pointer to **bool** | Port enabled mirror or not | [optional] 
**MirrorMode** | Pointer to **int32** | Port mirror mode should be a value as follows: 0: ingress; 1: egress; 2: ingress and egress | [optional] 
**MirroredPorts** | Pointer to **[]int32** | Mirrored Ports Set | [optional] 
**Port** | **int32** | Port serial number | 
**Pvid** | Pointer to **int32** | Pvid(only for lan port.) | [optional] 
**Status** | Pointer to **int32** | Enable port or not, status should be a value as follows: 0: disable; 1: enable.(When the port supports status.) | [optional] 

## Methods

### NewGatewayPortConfig

`func NewGatewayPortConfig(port int32, ) *GatewayPortConfig`

NewGatewayPortConfig instantiates a new GatewayPortConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortConfigWithDefaults

`func NewGatewayPortConfigWithDefaults() *GatewayPortConfig`

NewGatewayPortConfigWithDefaults instantiates a new GatewayPortConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDslSettings

`func (o *GatewayPortConfig) GetDslSettings() DslSettings`

GetDslSettings returns the DslSettings field if non-nil, zero value otherwise.

### GetDslSettingsOk

`func (o *GatewayPortConfig) GetDslSettingsOk() (*DslSettings, bool)`

GetDslSettingsOk returns a tuple with the DslSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslSettings

`func (o *GatewayPortConfig) SetDslSettings(v DslSettings)`

SetDslSettings sets DslSettings field to given value.

### HasDslSettings

`func (o *GatewayPortConfig) HasDslSettings() bool`

HasDslSettings returns a boolean if a field has been set.

### GetDuplex

`func (o *GatewayPortConfig) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *GatewayPortConfig) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *GatewayPortConfig) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *GatewayPortConfig) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetFlowControl

`func (o *GatewayPortConfig) GetFlowControl() bool`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *GatewayPortConfig) GetFlowControlOk() (*bool, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *GatewayPortConfig) SetFlowControl(v bool)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *GatewayPortConfig) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *GatewayPortConfig) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *GatewayPortConfig) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *GatewayPortConfig) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *GatewayPortConfig) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMirrorEnable

`func (o *GatewayPortConfig) GetMirrorEnable() bool`

GetMirrorEnable returns the MirrorEnable field if non-nil, zero value otherwise.

### GetMirrorEnableOk

`func (o *GatewayPortConfig) GetMirrorEnableOk() (*bool, bool)`

GetMirrorEnableOk returns a tuple with the MirrorEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorEnable

`func (o *GatewayPortConfig) SetMirrorEnable(v bool)`

SetMirrorEnable sets MirrorEnable field to given value.

### HasMirrorEnable

`func (o *GatewayPortConfig) HasMirrorEnable() bool`

HasMirrorEnable returns a boolean if a field has been set.

### GetMirrorMode

`func (o *GatewayPortConfig) GetMirrorMode() int32`

GetMirrorMode returns the MirrorMode field if non-nil, zero value otherwise.

### GetMirrorModeOk

`func (o *GatewayPortConfig) GetMirrorModeOk() (*int32, bool)`

GetMirrorModeOk returns a tuple with the MirrorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorMode

`func (o *GatewayPortConfig) SetMirrorMode(v int32)`

SetMirrorMode sets MirrorMode field to given value.

### HasMirrorMode

`func (o *GatewayPortConfig) HasMirrorMode() bool`

HasMirrorMode returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *GatewayPortConfig) GetMirroredPorts() []int32`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *GatewayPortConfig) GetMirroredPortsOk() (*[]int32, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *GatewayPortConfig) SetMirroredPorts(v []int32)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *GatewayPortConfig) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetPort

`func (o *GatewayPortConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayPortConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayPortConfig) SetPort(v int32)`

SetPort sets Port field to given value.


### GetPvid

`func (o *GatewayPortConfig) GetPvid() int32`

GetPvid returns the Pvid field if non-nil, zero value otherwise.

### GetPvidOk

`func (o *GatewayPortConfig) GetPvidOk() (*int32, bool)`

GetPvidOk returns a tuple with the Pvid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvid

`func (o *GatewayPortConfig) SetPvid(v int32)`

SetPvid sets Pvid field to given value.

### HasPvid

`func (o *GatewayPortConfig) HasPvid() bool`

HasPvid returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayPortConfig) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayPortConfig) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayPortConfig) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GatewayPortConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


