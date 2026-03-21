# OswStackDetailConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardDelay** | Pointer to **int32** | STP should be should be within the range of 4-30 | [optional] 
**HelloTime** | Pointer to **int32** | STP helloTime should be should be within the range of 1-10 | [optional] 
**Jumbo** | Pointer to **int32** | Jumbo should be should be within the range of 1518-9216 | [optional] 
**LagHashAlg** | Pointer to **int32** | It should be a value as follows: 0: SRC MAC; 1: DST MAC; 2: SRC MAC + DST MAC; 3: SRC IP; 4: DST IP; 5: SRC IP + DST IP | [optional] 
**LedSetting** | Pointer to **int32** | LED setting should be a value as follows: 0:off; 1:on; 2:Use Site Settings | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | LoopbackDetect enable status | [optional] 
**MaxAge** | Pointer to **int32** | STP maxAge should be should be within the range of 6-40 | [optional] 
**MaxHops** | Pointer to **int32** | STP maxHops should be should be within the range of 1-40 | [optional] 
**Mstp** | Pointer to [**OswStpMstpConfig**](OswStpMstpConfig.md) |  | [optional] 
**MvlanBridgeVlan** | Pointer to **int32** | Only valid when mvlanNetworkId is bridge vlan | [optional] 
**MvlanNetworkId** | Pointer to **string** | Management VLAN network ID | [optional] 
**Name** | Pointer to **string** | Stack Name | [optional] 
**Priority** | Pointer to **int32** | STP priority should be an integer from 0 to 61440 and divisible by 4096 | [optional] 
**QosConfig** | Pointer to [**OswQosConfigVO**](OswQosConfigVO.md) |  | [optional] 
**RememberDevice** | Pointer to **int32** | Whether to remember the device.RememberDevice should be a value as follows: 0:off, 1:on, 2: follow site | [optional] 
**Snmp** | Pointer to [**OswSnmpVO**](OswSnmpVO.md) |  | [optional] 
**Stp** | Pointer to **int32** | Spanning Tree Protocol should be a value as follows: 1: STP; 2: RSTP; 3: MSTP; 0: OFF | [optional] 
**TagIds** | Pointer to **[]string** | Tag ID List | [optional] 
**TxHoldCount** | Pointer to **int32** | STP txHoldCount should be should be within the range of 1-20 | [optional] 

## Methods

### NewOswStackDetailConfigOpenApiVO

`func NewOswStackDetailConfigOpenApiVO() *OswStackDetailConfigOpenApiVO`

NewOswStackDetailConfigOpenApiVO instantiates a new OswStackDetailConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackDetailConfigOpenApiVOWithDefaults

`func NewOswStackDetailConfigOpenApiVOWithDefaults() *OswStackDetailConfigOpenApiVO`

NewOswStackDetailConfigOpenApiVOWithDefaults instantiates a new OswStackDetailConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardDelay

`func (o *OswStackDetailConfigOpenApiVO) GetForwardDelay() int32`

GetForwardDelay returns the ForwardDelay field if non-nil, zero value otherwise.

### GetForwardDelayOk

`func (o *OswStackDetailConfigOpenApiVO) GetForwardDelayOk() (*int32, bool)`

GetForwardDelayOk returns a tuple with the ForwardDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDelay

`func (o *OswStackDetailConfigOpenApiVO) SetForwardDelay(v int32)`

SetForwardDelay sets ForwardDelay field to given value.

### HasForwardDelay

`func (o *OswStackDetailConfigOpenApiVO) HasForwardDelay() bool`

HasForwardDelay returns a boolean if a field has been set.

### GetHelloTime

`func (o *OswStackDetailConfigOpenApiVO) GetHelloTime() int32`

GetHelloTime returns the HelloTime field if non-nil, zero value otherwise.

### GetHelloTimeOk

`func (o *OswStackDetailConfigOpenApiVO) GetHelloTimeOk() (*int32, bool)`

GetHelloTimeOk returns a tuple with the HelloTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTime

`func (o *OswStackDetailConfigOpenApiVO) SetHelloTime(v int32)`

SetHelloTime sets HelloTime field to given value.

### HasHelloTime

`func (o *OswStackDetailConfigOpenApiVO) HasHelloTime() bool`

HasHelloTime returns a boolean if a field has been set.

### GetJumbo

`func (o *OswStackDetailConfigOpenApiVO) GetJumbo() int32`

GetJumbo returns the Jumbo field if non-nil, zero value otherwise.

### GetJumboOk

`func (o *OswStackDetailConfigOpenApiVO) GetJumboOk() (*int32, bool)`

GetJumboOk returns a tuple with the Jumbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumbo

`func (o *OswStackDetailConfigOpenApiVO) SetJumbo(v int32)`

SetJumbo sets Jumbo field to given value.

### HasJumbo

`func (o *OswStackDetailConfigOpenApiVO) HasJumbo() bool`

HasJumbo returns a boolean if a field has been set.

### GetLagHashAlg

`func (o *OswStackDetailConfigOpenApiVO) GetLagHashAlg() int32`

GetLagHashAlg returns the LagHashAlg field if non-nil, zero value otherwise.

### GetLagHashAlgOk

`func (o *OswStackDetailConfigOpenApiVO) GetLagHashAlgOk() (*int32, bool)`

GetLagHashAlgOk returns a tuple with the LagHashAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagHashAlg

`func (o *OswStackDetailConfigOpenApiVO) SetLagHashAlg(v int32)`

SetLagHashAlg sets LagHashAlg field to given value.

### HasLagHashAlg

`func (o *OswStackDetailConfigOpenApiVO) HasLagHashAlg() bool`

HasLagHashAlg returns a boolean if a field has been set.

### GetLedSetting

`func (o *OswStackDetailConfigOpenApiVO) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *OswStackDetailConfigOpenApiVO) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *OswStackDetailConfigOpenApiVO) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.

### HasLedSetting

`func (o *OswStackDetailConfigOpenApiVO) HasLedSetting() bool`

HasLedSetting returns a boolean if a field has been set.

### GetLocation

`func (o *OswStackDetailConfigOpenApiVO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OswStackDetailConfigOpenApiVO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OswStackDetailConfigOpenApiVO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OswStackDetailConfigOpenApiVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswStackDetailConfigOpenApiVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswStackDetailConfigOpenApiVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswStackDetailConfigOpenApiVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswStackDetailConfigOpenApiVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetMaxAge

`func (o *OswStackDetailConfigOpenApiVO) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *OswStackDetailConfigOpenApiVO) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *OswStackDetailConfigOpenApiVO) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *OswStackDetailConfigOpenApiVO) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxHops

`func (o *OswStackDetailConfigOpenApiVO) GetMaxHops() int32`

GetMaxHops returns the MaxHops field if non-nil, zero value otherwise.

### GetMaxHopsOk

`func (o *OswStackDetailConfigOpenApiVO) GetMaxHopsOk() (*int32, bool)`

GetMaxHopsOk returns a tuple with the MaxHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHops

`func (o *OswStackDetailConfigOpenApiVO) SetMaxHops(v int32)`

SetMaxHops sets MaxHops field to given value.

### HasMaxHops

`func (o *OswStackDetailConfigOpenApiVO) HasMaxHops() bool`

HasMaxHops returns a boolean if a field has been set.

### GetMstp

`func (o *OswStackDetailConfigOpenApiVO) GetMstp() OswStpMstpConfig`

GetMstp returns the Mstp field if non-nil, zero value otherwise.

### GetMstpOk

`func (o *OswStackDetailConfigOpenApiVO) GetMstpOk() (*OswStpMstpConfig, bool)`

GetMstpOk returns a tuple with the Mstp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstp

`func (o *OswStackDetailConfigOpenApiVO) SetMstp(v OswStpMstpConfig)`

SetMstp sets Mstp field to given value.

### HasMstp

`func (o *OswStackDetailConfigOpenApiVO) HasMstp() bool`

HasMstp returns a boolean if a field has been set.

### GetMvlanBridgeVlan

`func (o *OswStackDetailConfigOpenApiVO) GetMvlanBridgeVlan() int32`

GetMvlanBridgeVlan returns the MvlanBridgeVlan field if non-nil, zero value otherwise.

### GetMvlanBridgeVlanOk

`func (o *OswStackDetailConfigOpenApiVO) GetMvlanBridgeVlanOk() (*int32, bool)`

GetMvlanBridgeVlanOk returns a tuple with the MvlanBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanBridgeVlan

`func (o *OswStackDetailConfigOpenApiVO) SetMvlanBridgeVlan(v int32)`

SetMvlanBridgeVlan sets MvlanBridgeVlan field to given value.

### HasMvlanBridgeVlan

`func (o *OswStackDetailConfigOpenApiVO) HasMvlanBridgeVlan() bool`

HasMvlanBridgeVlan returns a boolean if a field has been set.

### GetMvlanNetworkId

`func (o *OswStackDetailConfigOpenApiVO) GetMvlanNetworkId() string`

GetMvlanNetworkId returns the MvlanNetworkId field if non-nil, zero value otherwise.

### GetMvlanNetworkIdOk

`func (o *OswStackDetailConfigOpenApiVO) GetMvlanNetworkIdOk() (*string, bool)`

GetMvlanNetworkIdOk returns a tuple with the MvlanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanNetworkId

`func (o *OswStackDetailConfigOpenApiVO) SetMvlanNetworkId(v string)`

SetMvlanNetworkId sets MvlanNetworkId field to given value.

### HasMvlanNetworkId

`func (o *OswStackDetailConfigOpenApiVO) HasMvlanNetworkId() bool`

HasMvlanNetworkId returns a boolean if a field has been set.

### GetName

`func (o *OswStackDetailConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackDetailConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackDetailConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStackDetailConfigOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *OswStackDetailConfigOpenApiVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswStackDetailConfigOpenApiVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswStackDetailConfigOpenApiVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OswStackDetailConfigOpenApiVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQosConfig

`func (o *OswStackDetailConfigOpenApiVO) GetQosConfig() OswQosConfigVO`

GetQosConfig returns the QosConfig field if non-nil, zero value otherwise.

### GetQosConfigOk

`func (o *OswStackDetailConfigOpenApiVO) GetQosConfigOk() (*OswQosConfigVO, bool)`

GetQosConfigOk returns a tuple with the QosConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosConfig

`func (o *OswStackDetailConfigOpenApiVO) SetQosConfig(v OswQosConfigVO)`

SetQosConfig sets QosConfig field to given value.

### HasQosConfig

`func (o *OswStackDetailConfigOpenApiVO) HasQosConfig() bool`

HasQosConfig returns a boolean if a field has been set.

### GetRememberDevice

`func (o *OswStackDetailConfigOpenApiVO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *OswStackDetailConfigOpenApiVO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *OswStackDetailConfigOpenApiVO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *OswStackDetailConfigOpenApiVO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetSnmp

`func (o *OswStackDetailConfigOpenApiVO) GetSnmp() OswSnmpVO`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *OswStackDetailConfigOpenApiVO) GetSnmpOk() (*OswSnmpVO, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *OswStackDetailConfigOpenApiVO) SetSnmp(v OswSnmpVO)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *OswStackDetailConfigOpenApiVO) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetStp

`func (o *OswStackDetailConfigOpenApiVO) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *OswStackDetailConfigOpenApiVO) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *OswStackDetailConfigOpenApiVO) SetStp(v int32)`

SetStp sets Stp field to given value.

### HasStp

`func (o *OswStackDetailConfigOpenApiVO) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetTagIds

`func (o *OswStackDetailConfigOpenApiVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswStackDetailConfigOpenApiVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswStackDetailConfigOpenApiVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswStackDetailConfigOpenApiVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTxHoldCount

`func (o *OswStackDetailConfigOpenApiVO) GetTxHoldCount() int32`

GetTxHoldCount returns the TxHoldCount field if non-nil, zero value otherwise.

### GetTxHoldCountOk

`func (o *OswStackDetailConfigOpenApiVO) GetTxHoldCountOk() (*int32, bool)`

GetTxHoldCountOk returns a tuple with the TxHoldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHoldCount

`func (o *OswStackDetailConfigOpenApiVO) SetTxHoldCount(v int32)`

SetTxHoldCount sets TxHoldCount field to given value.

### HasTxHoldCount

`func (o *OswStackDetailConfigOpenApiVO) HasTxHoldCount() bool`

HasTxHoldCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


