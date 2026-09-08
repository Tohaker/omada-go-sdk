# CondBroadcastCtrlVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **int32** | 0: uplink down, 1: Internet down. | [optional] 
**DownTime** | Pointer to **int32** | SSID Downlink Time | [optional] 
**Enable** | Pointer to **bool** | enable | [optional] 
**UpTime** | Pointer to **int32** | SSID Uplink Time | [optional] 

## Methods

### NewCondBroadcastCtrlVO

`func NewCondBroadcastCtrlVO() *CondBroadcastCtrlVO`

NewCondBroadcastCtrlVO instantiates a new CondBroadcastCtrlVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondBroadcastCtrlVOWithDefaults

`func NewCondBroadcastCtrlVOWithDefaults() *CondBroadcastCtrlVO`

NewCondBroadcastCtrlVOWithDefaults instantiates a new CondBroadcastCtrlVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *CondBroadcastCtrlVO) GetCondition() int32`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CondBroadcastCtrlVO) GetConditionOk() (*int32, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CondBroadcastCtrlVO) SetCondition(v int32)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CondBroadcastCtrlVO) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDownTime

`func (o *CondBroadcastCtrlVO) GetDownTime() int32`

GetDownTime returns the DownTime field if non-nil, zero value otherwise.

### GetDownTimeOk

`func (o *CondBroadcastCtrlVO) GetDownTimeOk() (*int32, bool)`

GetDownTimeOk returns a tuple with the DownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownTime

`func (o *CondBroadcastCtrlVO) SetDownTime(v int32)`

SetDownTime sets DownTime field to given value.

### HasDownTime

`func (o *CondBroadcastCtrlVO) HasDownTime() bool`

HasDownTime returns a boolean if a field has been set.

### GetEnable

`func (o *CondBroadcastCtrlVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CondBroadcastCtrlVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CondBroadcastCtrlVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CondBroadcastCtrlVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetUpTime

`func (o *CondBroadcastCtrlVO) GetUpTime() int32`

GetUpTime returns the UpTime field if non-nil, zero value otherwise.

### GetUpTimeOk

`func (o *CondBroadcastCtrlVO) GetUpTimeOk() (*int32, bool)`

GetUpTimeOk returns a tuple with the UpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTime

`func (o *CondBroadcastCtrlVO) SetUpTime(v int32)`

SetUpTime sets UpTime field to given value.

### HasUpTime

`func (o *CondBroadcastCtrlVO) HasUpTime() bool`

HasUpTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


