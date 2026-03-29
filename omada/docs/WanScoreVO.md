# WanScoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **int32** | WAN health score | [optional] 
**Time** | Pointer to **int64** | timestamp | [optional] 

## Methods

### NewWanScoreVO

`func NewWanScoreVO() *WanScoreVO`

NewWanScoreVO instantiates a new WanScoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanScoreVOWithDefaults

`func NewWanScoreVOWithDefaults() *WanScoreVO`

NewWanScoreVOWithDefaults instantiates a new WanScoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *WanScoreVO) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *WanScoreVO) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *WanScoreVO) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *WanScoreVO) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTime

`func (o *WanScoreVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WanScoreVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WanScoreVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *WanScoreVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


