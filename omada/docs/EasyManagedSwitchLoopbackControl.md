# EasyManagedSwitchLoopbackControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardDelay** | Pointer to **int32** | forwardDelay should be between 4 and 30. | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | LoopbackDetectEnable | [optional] 
**MaxAge** | Pointer to **int32** | maxAge should be between 6 and 40. | [optional] 
**Priority** | Pointer to **int32** | priority | [optional] 
**Stp** | Pointer to **int32** | STP should be a value as follows: 0: OFF 1: STP 2: RSTP | [optional] 
**TxHoldCount** | Pointer to **int32** | txHoldCount should be between 1 and 10. | [optional] 

## Methods

### NewEasyManagedSwitchLoopbackControl

`func NewEasyManagedSwitchLoopbackControl() *EasyManagedSwitchLoopbackControl`

NewEasyManagedSwitchLoopbackControl instantiates a new EasyManagedSwitchLoopbackControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasyManagedSwitchLoopbackControlWithDefaults

`func NewEasyManagedSwitchLoopbackControlWithDefaults() *EasyManagedSwitchLoopbackControl`

NewEasyManagedSwitchLoopbackControlWithDefaults instantiates a new EasyManagedSwitchLoopbackControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardDelay

`func (o *EasyManagedSwitchLoopbackControl) GetForwardDelay() int32`

GetForwardDelay returns the ForwardDelay field if non-nil, zero value otherwise.

### GetForwardDelayOk

`func (o *EasyManagedSwitchLoopbackControl) GetForwardDelayOk() (*int32, bool)`

GetForwardDelayOk returns a tuple with the ForwardDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDelay

`func (o *EasyManagedSwitchLoopbackControl) SetForwardDelay(v int32)`

SetForwardDelay sets ForwardDelay field to given value.

### HasForwardDelay

`func (o *EasyManagedSwitchLoopbackControl) HasForwardDelay() bool`

HasForwardDelay returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *EasyManagedSwitchLoopbackControl) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *EasyManagedSwitchLoopbackControl) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *EasyManagedSwitchLoopbackControl) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *EasyManagedSwitchLoopbackControl) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetMaxAge

`func (o *EasyManagedSwitchLoopbackControl) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *EasyManagedSwitchLoopbackControl) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *EasyManagedSwitchLoopbackControl) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *EasyManagedSwitchLoopbackControl) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetPriority

`func (o *EasyManagedSwitchLoopbackControl) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EasyManagedSwitchLoopbackControl) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EasyManagedSwitchLoopbackControl) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EasyManagedSwitchLoopbackControl) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStp

`func (o *EasyManagedSwitchLoopbackControl) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *EasyManagedSwitchLoopbackControl) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *EasyManagedSwitchLoopbackControl) SetStp(v int32)`

SetStp sets Stp field to given value.

### HasStp

`func (o *EasyManagedSwitchLoopbackControl) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetTxHoldCount

`func (o *EasyManagedSwitchLoopbackControl) GetTxHoldCount() int32`

GetTxHoldCount returns the TxHoldCount field if non-nil, zero value otherwise.

### GetTxHoldCountOk

`func (o *EasyManagedSwitchLoopbackControl) GetTxHoldCountOk() (*int32, bool)`

GetTxHoldCountOk returns a tuple with the TxHoldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHoldCount

`func (o *EasyManagedSwitchLoopbackControl) SetTxHoldCount(v int32)`

SetTxHoldCount sets TxHoldCount field to given value.

### HasTxHoldCount

`func (o *EasyManagedSwitchLoopbackControl) HasTxHoldCount() bool`

HasTxHoldCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


