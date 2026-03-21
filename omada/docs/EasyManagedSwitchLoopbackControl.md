# EasyManagedSwitchLoopbackControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoopbackDetectEnable** | Pointer to **bool** | LoopbackDetectEnable | [optional] 
**Priority** | Pointer to **int32** | priority | [optional] 
**Stp** | Pointer to **int32** | STP should be a value as follows: 0: OFF 1: STP 2: RSTP 3: MSTP  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


