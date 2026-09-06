# TimeFloatValueItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PastNum** | Pointer to **float32** | Value corresponding to time | [optional] 
**Time** | Pointer to **int64** | Time(unit:ms) | [optional] 

## Methods

### NewTimeFloatValueItemVO

`func NewTimeFloatValueItemVO() *TimeFloatValueItemVO`

NewTimeFloatValueItemVO instantiates a new TimeFloatValueItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeFloatValueItemVOWithDefaults

`func NewTimeFloatValueItemVOWithDefaults() *TimeFloatValueItemVO`

NewTimeFloatValueItemVOWithDefaults instantiates a new TimeFloatValueItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPastNum

`func (o *TimeFloatValueItemVO) GetPastNum() float32`

GetPastNum returns the PastNum field if non-nil, zero value otherwise.

### GetPastNumOk

`func (o *TimeFloatValueItemVO) GetPastNumOk() (*float32, bool)`

GetPastNumOk returns a tuple with the PastNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNum

`func (o *TimeFloatValueItemVO) SetPastNum(v float32)`

SetPastNum sets PastNum field to given value.

### HasPastNum

`func (o *TimeFloatValueItemVO) HasPastNum() bool`

HasPastNum returns a boolean if a field has been set.

### GetTime

`func (o *TimeFloatValueItemVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TimeFloatValueItemVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TimeFloatValueItemVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *TimeFloatValueItemVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


