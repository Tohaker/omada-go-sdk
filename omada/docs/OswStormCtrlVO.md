# OswStormCtrlVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **int32** | Action should be a value as follows: 0: drop; 1: shutdown | 
**Broadcast** | Pointer to **int32** | Broadcast | [optional] 
**BroadcastEnable** | **bool** | Indicates whether broadcast is enabled | 
**Multicast** | Pointer to **int32** | Multicast | [optional] 
**MulticastEnable** | **bool** | Indicates whether multicast is enabled | 
**RateMode** | Pointer to **int32** | RateMode should be a value as follows: 0: ratio; 1: kbps | [optional] 
**RecoverTime** | Pointer to **int32** | Recover Time | [optional] 
**UnknownUnicast** | Pointer to **int32** | Unknown Unicast | [optional] 
**UnknownUnicastEnable** | **bool** | Indicates whether unknown unicast is enabled | 

## Methods

### NewOswStormCtrlVO

`func NewOswStormCtrlVO(action int32, broadcastEnable bool, multicastEnable bool, unknownUnicastEnable bool, ) *OswStormCtrlVO`

NewOswStormCtrlVO instantiates a new OswStormCtrlVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStormCtrlVOWithDefaults

`func NewOswStormCtrlVOWithDefaults() *OswStormCtrlVO`

NewOswStormCtrlVOWithDefaults instantiates a new OswStormCtrlVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OswStormCtrlVO) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OswStormCtrlVO) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OswStormCtrlVO) SetAction(v int32)`

SetAction sets Action field to given value.


### GetBroadcast

`func (o *OswStormCtrlVO) GetBroadcast() int32`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *OswStormCtrlVO) GetBroadcastOk() (*int32, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *OswStormCtrlVO) SetBroadcast(v int32)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *OswStormCtrlVO) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetBroadcastEnable

`func (o *OswStormCtrlVO) GetBroadcastEnable() bool`

GetBroadcastEnable returns the BroadcastEnable field if non-nil, zero value otherwise.

### GetBroadcastEnableOk

`func (o *OswStormCtrlVO) GetBroadcastEnableOk() (*bool, bool)`

GetBroadcastEnableOk returns a tuple with the BroadcastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastEnable

`func (o *OswStormCtrlVO) SetBroadcastEnable(v bool)`

SetBroadcastEnable sets BroadcastEnable field to given value.


### GetMulticast

`func (o *OswStormCtrlVO) GetMulticast() int32`

GetMulticast returns the Multicast field if non-nil, zero value otherwise.

### GetMulticastOk

`func (o *OswStormCtrlVO) GetMulticastOk() (*int32, bool)`

GetMulticastOk returns a tuple with the Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticast

`func (o *OswStormCtrlVO) SetMulticast(v int32)`

SetMulticast sets Multicast field to given value.

### HasMulticast

`func (o *OswStormCtrlVO) HasMulticast() bool`

HasMulticast returns a boolean if a field has been set.

### GetMulticastEnable

`func (o *OswStormCtrlVO) GetMulticastEnable() bool`

GetMulticastEnable returns the MulticastEnable field if non-nil, zero value otherwise.

### GetMulticastEnableOk

`func (o *OswStormCtrlVO) GetMulticastEnableOk() (*bool, bool)`

GetMulticastEnableOk returns a tuple with the MulticastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastEnable

`func (o *OswStormCtrlVO) SetMulticastEnable(v bool)`

SetMulticastEnable sets MulticastEnable field to given value.


### GetRateMode

`func (o *OswStormCtrlVO) GetRateMode() int32`

GetRateMode returns the RateMode field if non-nil, zero value otherwise.

### GetRateModeOk

`func (o *OswStormCtrlVO) GetRateModeOk() (*int32, bool)`

GetRateModeOk returns a tuple with the RateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateMode

`func (o *OswStormCtrlVO) SetRateMode(v int32)`

SetRateMode sets RateMode field to given value.

### HasRateMode

`func (o *OswStormCtrlVO) HasRateMode() bool`

HasRateMode returns a boolean if a field has been set.

### GetRecoverTime

`func (o *OswStormCtrlVO) GetRecoverTime() int32`

GetRecoverTime returns the RecoverTime field if non-nil, zero value otherwise.

### GetRecoverTimeOk

`func (o *OswStormCtrlVO) GetRecoverTimeOk() (*int32, bool)`

GetRecoverTimeOk returns a tuple with the RecoverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTime

`func (o *OswStormCtrlVO) SetRecoverTime(v int32)`

SetRecoverTime sets RecoverTime field to given value.

### HasRecoverTime

`func (o *OswStormCtrlVO) HasRecoverTime() bool`

HasRecoverTime returns a boolean if a field has been set.

### GetUnknownUnicast

`func (o *OswStormCtrlVO) GetUnknownUnicast() int32`

GetUnknownUnicast returns the UnknownUnicast field if non-nil, zero value otherwise.

### GetUnknownUnicastOk

`func (o *OswStormCtrlVO) GetUnknownUnicastOk() (*int32, bool)`

GetUnknownUnicastOk returns a tuple with the UnknownUnicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicast

`func (o *OswStormCtrlVO) SetUnknownUnicast(v int32)`

SetUnknownUnicast sets UnknownUnicast field to given value.

### HasUnknownUnicast

`func (o *OswStormCtrlVO) HasUnknownUnicast() bool`

HasUnknownUnicast returns a boolean if a field has been set.

### GetUnknownUnicastEnable

`func (o *OswStormCtrlVO) GetUnknownUnicastEnable() bool`

GetUnknownUnicastEnable returns the UnknownUnicastEnable field if non-nil, zero value otherwise.

### GetUnknownUnicastEnableOk

`func (o *OswStormCtrlVO) GetUnknownUnicastEnableOk() (*bool, bool)`

GetUnknownUnicastEnableOk returns a tuple with the UnknownUnicastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicastEnable

`func (o *OswStormCtrlVO) SetUnknownUnicastEnable(v bool)`

SetUnknownUnicastEnable sets UnknownUnicastEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


