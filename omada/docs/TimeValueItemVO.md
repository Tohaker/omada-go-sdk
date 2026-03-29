# TimeValueItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PastNum** | Pointer to **int32** | Value corresponding to time | [optional] 
**Time** | Pointer to **int64** | Time(unit:ms) | [optional] 

## Methods

### NewTimeValueItemVO

`func NewTimeValueItemVO() *TimeValueItemVO`

NewTimeValueItemVO instantiates a new TimeValueItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeValueItemVOWithDefaults

`func NewTimeValueItemVOWithDefaults() *TimeValueItemVO`

NewTimeValueItemVOWithDefaults instantiates a new TimeValueItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPastNum

`func (o *TimeValueItemVO) GetPastNum() int32`

GetPastNum returns the PastNum field if non-nil, zero value otherwise.

### GetPastNumOk

`func (o *TimeValueItemVO) GetPastNumOk() (*int32, bool)`

GetPastNumOk returns a tuple with the PastNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNum

`func (o *TimeValueItemVO) SetPastNum(v int32)`

SetPastNum sets PastNum field to given value.

### HasPastNum

`func (o *TimeValueItemVO) HasPastNum() bool`

HasPastNum returns a boolean if a field has been set.

### GetTime

`func (o *TimeValueItemVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TimeValueItemVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TimeValueItemVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *TimeValueItemVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


