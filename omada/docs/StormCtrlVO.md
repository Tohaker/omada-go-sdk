# StormCtrlVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **int32** | Action should be a value as follows: 0: drop(default), 1: shutdown | 
**Broadcast** | Pointer to **int32** | Broadcast | [optional] 
**BroadcastEnable** | **bool** | Broadcast enable status | 
**Multicast** | Pointer to **int32** | Multicast | [optional] 
**MulticastEnable** | **bool** | Multicast enable status | 
**RateMode** | Pointer to **int32** | RateMode should be a value as follows: 0: ratio; 1: kbps | [optional] 
**RecoverTime** | Pointer to **int32** | Recover Time should be within the range of 1-3600, default 3600 | [optional] 
**UnknownUnicast** | Pointer to **int32** | Unknown-Unicast | [optional] 
**UnknownUnicastEnable** | **bool** | Unknown-Unicast enable status | 

## Methods

### NewStormCtrlVO

`func NewStormCtrlVO(action int32, broadcastEnable bool, multicastEnable bool, unknownUnicastEnable bool, ) *StormCtrlVO`

NewStormCtrlVO instantiates a new StormCtrlVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStormCtrlVOWithDefaults

`func NewStormCtrlVOWithDefaults() *StormCtrlVO`

NewStormCtrlVOWithDefaults instantiates a new StormCtrlVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *StormCtrlVO) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StormCtrlVO) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StormCtrlVO) SetAction(v int32)`

SetAction sets Action field to given value.


### GetBroadcast

`func (o *StormCtrlVO) GetBroadcast() int32`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *StormCtrlVO) GetBroadcastOk() (*int32, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *StormCtrlVO) SetBroadcast(v int32)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *StormCtrlVO) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetBroadcastEnable

`func (o *StormCtrlVO) GetBroadcastEnable() bool`

GetBroadcastEnable returns the BroadcastEnable field if non-nil, zero value otherwise.

### GetBroadcastEnableOk

`func (o *StormCtrlVO) GetBroadcastEnableOk() (*bool, bool)`

GetBroadcastEnableOk returns a tuple with the BroadcastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastEnable

`func (o *StormCtrlVO) SetBroadcastEnable(v bool)`

SetBroadcastEnable sets BroadcastEnable field to given value.


### GetMulticast

`func (o *StormCtrlVO) GetMulticast() int32`

GetMulticast returns the Multicast field if non-nil, zero value otherwise.

### GetMulticastOk

`func (o *StormCtrlVO) GetMulticastOk() (*int32, bool)`

GetMulticastOk returns a tuple with the Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticast

`func (o *StormCtrlVO) SetMulticast(v int32)`

SetMulticast sets Multicast field to given value.

### HasMulticast

`func (o *StormCtrlVO) HasMulticast() bool`

HasMulticast returns a boolean if a field has been set.

### GetMulticastEnable

`func (o *StormCtrlVO) GetMulticastEnable() bool`

GetMulticastEnable returns the MulticastEnable field if non-nil, zero value otherwise.

### GetMulticastEnableOk

`func (o *StormCtrlVO) GetMulticastEnableOk() (*bool, bool)`

GetMulticastEnableOk returns a tuple with the MulticastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastEnable

`func (o *StormCtrlVO) SetMulticastEnable(v bool)`

SetMulticastEnable sets MulticastEnable field to given value.


### GetRateMode

`func (o *StormCtrlVO) GetRateMode() int32`

GetRateMode returns the RateMode field if non-nil, zero value otherwise.

### GetRateModeOk

`func (o *StormCtrlVO) GetRateModeOk() (*int32, bool)`

GetRateModeOk returns a tuple with the RateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateMode

`func (o *StormCtrlVO) SetRateMode(v int32)`

SetRateMode sets RateMode field to given value.

### HasRateMode

`func (o *StormCtrlVO) HasRateMode() bool`

HasRateMode returns a boolean if a field has been set.

### GetRecoverTime

`func (o *StormCtrlVO) GetRecoverTime() int32`

GetRecoverTime returns the RecoverTime field if non-nil, zero value otherwise.

### GetRecoverTimeOk

`func (o *StormCtrlVO) GetRecoverTimeOk() (*int32, bool)`

GetRecoverTimeOk returns a tuple with the RecoverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTime

`func (o *StormCtrlVO) SetRecoverTime(v int32)`

SetRecoverTime sets RecoverTime field to given value.

### HasRecoverTime

`func (o *StormCtrlVO) HasRecoverTime() bool`

HasRecoverTime returns a boolean if a field has been set.

### GetUnknownUnicast

`func (o *StormCtrlVO) GetUnknownUnicast() int32`

GetUnknownUnicast returns the UnknownUnicast field if non-nil, zero value otherwise.

### GetUnknownUnicastOk

`func (o *StormCtrlVO) GetUnknownUnicastOk() (*int32, bool)`

GetUnknownUnicastOk returns a tuple with the UnknownUnicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicast

`func (o *StormCtrlVO) SetUnknownUnicast(v int32)`

SetUnknownUnicast sets UnknownUnicast field to given value.

### HasUnknownUnicast

`func (o *StormCtrlVO) HasUnknownUnicast() bool`

HasUnknownUnicast returns a boolean if a field has been set.

### GetUnknownUnicastEnable

`func (o *StormCtrlVO) GetUnknownUnicastEnable() bool`

GetUnknownUnicastEnable returns the UnknownUnicastEnable field if non-nil, zero value otherwise.

### GetUnknownUnicastEnableOk

`func (o *StormCtrlVO) GetUnknownUnicastEnableOk() (*bool, bool)`

GetUnknownUnicastEnableOk returns a tuple with the UnknownUnicastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicastEnable

`func (o *StormCtrlVO) SetUnknownUnicastEnable(v bool)`

SetUnknownUnicastEnable sets UnknownUnicastEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


