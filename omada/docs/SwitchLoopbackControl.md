# SwitchLoopbackControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardDelay** | Pointer to **int32** | forwardDelay should be between 4 and 30. | [optional] 
**HelloTime** | Pointer to **int32** | helloTime should be between 1 and 10. | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | LoopbackDetectEnable | [optional] 
**MaxAge** | Pointer to **int32** | maxAge should be between 6 and 40. | [optional] 
**MaxHops** | Pointer to **int32** | maxHops should be between 1 and 40. | [optional] 
**Mstp** | Pointer to [**OswStpMstpConfigOpenApiVO**](OswStpMstpConfigOpenApiVO.md) |  | [optional] 
**Priority** | Pointer to **int32** | Parameter [priority] should be an integer from 0 to 61440 and divisible by 4096. | [optional] 
**Stp** | Pointer to **int32** | STP should be a value as follows: 0: OFF 1: STP 2: RSTP 3: MSTP  | [optional] 
**TxHoldCount** | Pointer to **int32** | txHoldCount should be between 1 and 20. | [optional] 

## Methods

### NewSwitchLoopbackControl

`func NewSwitchLoopbackControl() *SwitchLoopbackControl`

NewSwitchLoopbackControl instantiates a new SwitchLoopbackControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchLoopbackControlWithDefaults

`func NewSwitchLoopbackControlWithDefaults() *SwitchLoopbackControl`

NewSwitchLoopbackControlWithDefaults instantiates a new SwitchLoopbackControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardDelay

`func (o *SwitchLoopbackControl) GetForwardDelay() int32`

GetForwardDelay returns the ForwardDelay field if non-nil, zero value otherwise.

### GetForwardDelayOk

`func (o *SwitchLoopbackControl) GetForwardDelayOk() (*int32, bool)`

GetForwardDelayOk returns a tuple with the ForwardDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDelay

`func (o *SwitchLoopbackControl) SetForwardDelay(v int32)`

SetForwardDelay sets ForwardDelay field to given value.

### HasForwardDelay

`func (o *SwitchLoopbackControl) HasForwardDelay() bool`

HasForwardDelay returns a boolean if a field has been set.

### GetHelloTime

`func (o *SwitchLoopbackControl) GetHelloTime() int32`

GetHelloTime returns the HelloTime field if non-nil, zero value otherwise.

### GetHelloTimeOk

`func (o *SwitchLoopbackControl) GetHelloTimeOk() (*int32, bool)`

GetHelloTimeOk returns a tuple with the HelloTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTime

`func (o *SwitchLoopbackControl) SetHelloTime(v int32)`

SetHelloTime sets HelloTime field to given value.

### HasHelloTime

`func (o *SwitchLoopbackControl) HasHelloTime() bool`

HasHelloTime returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *SwitchLoopbackControl) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *SwitchLoopbackControl) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *SwitchLoopbackControl) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *SwitchLoopbackControl) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetMaxAge

`func (o *SwitchLoopbackControl) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *SwitchLoopbackControl) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *SwitchLoopbackControl) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *SwitchLoopbackControl) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxHops

`func (o *SwitchLoopbackControl) GetMaxHops() int32`

GetMaxHops returns the MaxHops field if non-nil, zero value otherwise.

### GetMaxHopsOk

`func (o *SwitchLoopbackControl) GetMaxHopsOk() (*int32, bool)`

GetMaxHopsOk returns a tuple with the MaxHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHops

`func (o *SwitchLoopbackControl) SetMaxHops(v int32)`

SetMaxHops sets MaxHops field to given value.

### HasMaxHops

`func (o *SwitchLoopbackControl) HasMaxHops() bool`

HasMaxHops returns a boolean if a field has been set.

### GetMstp

`func (o *SwitchLoopbackControl) GetMstp() OswStpMstpConfigOpenApiVO`

GetMstp returns the Mstp field if non-nil, zero value otherwise.

### GetMstpOk

`func (o *SwitchLoopbackControl) GetMstpOk() (*OswStpMstpConfigOpenApiVO, bool)`

GetMstpOk returns a tuple with the Mstp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstp

`func (o *SwitchLoopbackControl) SetMstp(v OswStpMstpConfigOpenApiVO)`

SetMstp sets Mstp field to given value.

### HasMstp

`func (o *SwitchLoopbackControl) HasMstp() bool`

HasMstp returns a boolean if a field has been set.

### GetPriority

`func (o *SwitchLoopbackControl) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SwitchLoopbackControl) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SwitchLoopbackControl) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SwitchLoopbackControl) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStp

`func (o *SwitchLoopbackControl) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *SwitchLoopbackControl) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *SwitchLoopbackControl) SetStp(v int32)`

SetStp sets Stp field to given value.

### HasStp

`func (o *SwitchLoopbackControl) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetTxHoldCount

`func (o *SwitchLoopbackControl) GetTxHoldCount() int32`

GetTxHoldCount returns the TxHoldCount field if non-nil, zero value otherwise.

### GetTxHoldCountOk

`func (o *SwitchLoopbackControl) GetTxHoldCountOk() (*int32, bool)`

GetTxHoldCountOk returns a tuple with the TxHoldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHoldCount

`func (o *SwitchLoopbackControl) SetTxHoldCount(v int32)`

SetTxHoldCount sets TxHoldCount field to given value.

### HasTxHoldCount

`func (o *SwitchLoopbackControl) HasTxHoldCount() bool`

HasTxHoldCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


