# SwitchBatchLoopbackControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardDelay** | Pointer to **int32** | forwardDelay should be between 4 and 30. | [optional] 
**HelloTime** | Pointer to **int32** | helloTime should be between 1 and 10. | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | LoopbackDetectEnable | [optional] 
**MaxAge** | Pointer to **int32** | maxAge should be between 6 and 40. | [optional] 
**MaxHops** | Pointer to **int32** | maxHops should be between 1 and 40. | [optional] 
**Mstp** | Pointer to [**OswStpMstpConfig**](OswStpMstpConfig.md) |  | [optional] 
**Priority** | Pointer to **int32** | Parameter [priority] should be an integer from 0 to 61440 and divisible by 4096. | [optional] 
**Stp** | Pointer to **int32** | STP should be a value as follows: 0: OFF 1: STP 2: RSTP 3: MSTP  | [optional] 
**SwitchMacList** | **[]string** | Parameter [switchMacList] should contain at least one switch MAC | 
**TxHoldCount** | Pointer to **int32** | txHoldCount should be between 1 and 20. | [optional] 

## Methods

### NewSwitchBatchLoopbackControl

`func NewSwitchBatchLoopbackControl(switchMacList []string, ) *SwitchBatchLoopbackControl`

NewSwitchBatchLoopbackControl instantiates a new SwitchBatchLoopbackControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchBatchLoopbackControlWithDefaults

`func NewSwitchBatchLoopbackControlWithDefaults() *SwitchBatchLoopbackControl`

NewSwitchBatchLoopbackControlWithDefaults instantiates a new SwitchBatchLoopbackControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardDelay

`func (o *SwitchBatchLoopbackControl) GetForwardDelay() int32`

GetForwardDelay returns the ForwardDelay field if non-nil, zero value otherwise.

### GetForwardDelayOk

`func (o *SwitchBatchLoopbackControl) GetForwardDelayOk() (*int32, bool)`

GetForwardDelayOk returns a tuple with the ForwardDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDelay

`func (o *SwitchBatchLoopbackControl) SetForwardDelay(v int32)`

SetForwardDelay sets ForwardDelay field to given value.

### HasForwardDelay

`func (o *SwitchBatchLoopbackControl) HasForwardDelay() bool`

HasForwardDelay returns a boolean if a field has been set.

### GetHelloTime

`func (o *SwitchBatchLoopbackControl) GetHelloTime() int32`

GetHelloTime returns the HelloTime field if non-nil, zero value otherwise.

### GetHelloTimeOk

`func (o *SwitchBatchLoopbackControl) GetHelloTimeOk() (*int32, bool)`

GetHelloTimeOk returns a tuple with the HelloTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTime

`func (o *SwitchBatchLoopbackControl) SetHelloTime(v int32)`

SetHelloTime sets HelloTime field to given value.

### HasHelloTime

`func (o *SwitchBatchLoopbackControl) HasHelloTime() bool`

HasHelloTime returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *SwitchBatchLoopbackControl) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *SwitchBatchLoopbackControl) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *SwitchBatchLoopbackControl) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *SwitchBatchLoopbackControl) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetMaxAge

`func (o *SwitchBatchLoopbackControl) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *SwitchBatchLoopbackControl) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *SwitchBatchLoopbackControl) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *SwitchBatchLoopbackControl) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxHops

`func (o *SwitchBatchLoopbackControl) GetMaxHops() int32`

GetMaxHops returns the MaxHops field if non-nil, zero value otherwise.

### GetMaxHopsOk

`func (o *SwitchBatchLoopbackControl) GetMaxHopsOk() (*int32, bool)`

GetMaxHopsOk returns a tuple with the MaxHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHops

`func (o *SwitchBatchLoopbackControl) SetMaxHops(v int32)`

SetMaxHops sets MaxHops field to given value.

### HasMaxHops

`func (o *SwitchBatchLoopbackControl) HasMaxHops() bool`

HasMaxHops returns a boolean if a field has been set.

### GetMstp

`func (o *SwitchBatchLoopbackControl) GetMstp() OswStpMstpConfig`

GetMstp returns the Mstp field if non-nil, zero value otherwise.

### GetMstpOk

`func (o *SwitchBatchLoopbackControl) GetMstpOk() (*OswStpMstpConfig, bool)`

GetMstpOk returns a tuple with the Mstp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstp

`func (o *SwitchBatchLoopbackControl) SetMstp(v OswStpMstpConfig)`

SetMstp sets Mstp field to given value.

### HasMstp

`func (o *SwitchBatchLoopbackControl) HasMstp() bool`

HasMstp returns a boolean if a field has been set.

### GetPriority

`func (o *SwitchBatchLoopbackControl) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SwitchBatchLoopbackControl) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SwitchBatchLoopbackControl) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SwitchBatchLoopbackControl) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStp

`func (o *SwitchBatchLoopbackControl) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *SwitchBatchLoopbackControl) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *SwitchBatchLoopbackControl) SetStp(v int32)`

SetStp sets Stp field to given value.

### HasStp

`func (o *SwitchBatchLoopbackControl) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetSwitchMacList

`func (o *SwitchBatchLoopbackControl) GetSwitchMacList() []string`

GetSwitchMacList returns the SwitchMacList field if non-nil, zero value otherwise.

### GetSwitchMacListOk

`func (o *SwitchBatchLoopbackControl) GetSwitchMacListOk() (*[]string, bool)`

GetSwitchMacListOk returns a tuple with the SwitchMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMacList

`func (o *SwitchBatchLoopbackControl) SetSwitchMacList(v []string)`

SetSwitchMacList sets SwitchMacList field to given value.


### GetTxHoldCount

`func (o *SwitchBatchLoopbackControl) GetTxHoldCount() int32`

GetTxHoldCount returns the TxHoldCount field if non-nil, zero value otherwise.

### GetTxHoldCountOk

`func (o *SwitchBatchLoopbackControl) GetTxHoldCountOk() (*int32, bool)`

GetTxHoldCountOk returns a tuple with the TxHoldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHoldCount

`func (o *SwitchBatchLoopbackControl) SetTxHoldCount(v int32)`

SetTxHoldCount sets TxHoldCount field to given value.

### HasTxHoldCount

`func (o *SwitchBatchLoopbackControl) HasTxHoldCount() bool`

HasTxHoldCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


