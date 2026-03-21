# TimeScoreItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **int32** | Health score corresponding to time | [optional] 
**Time** | Pointer to **int64** | Time(unit:ms) | [optional] 

## Methods

### NewTimeScoreItemVO

`func NewTimeScoreItemVO() *TimeScoreItemVO`

NewTimeScoreItemVO instantiates a new TimeScoreItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeScoreItemVOWithDefaults

`func NewTimeScoreItemVOWithDefaults() *TimeScoreItemVO`

NewTimeScoreItemVOWithDefaults instantiates a new TimeScoreItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *TimeScoreItemVO) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TimeScoreItemVO) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TimeScoreItemVO) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *TimeScoreItemVO) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTime

`func (o *TimeScoreItemVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TimeScoreItemVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TimeScoreItemVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *TimeScoreItemVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


