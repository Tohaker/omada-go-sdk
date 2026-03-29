# HealthStatisticsTimeScoreItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScoreDetail** | Pointer to [**HealthStatisticsScoreVO**](HealthStatisticsScoreVO.md) |  | [optional] 
**Time** | Pointer to **int64** | Time(unit:ms) | [optional] 

## Methods

### NewHealthStatisticsTimeScoreItemVO

`func NewHealthStatisticsTimeScoreItemVO() *HealthStatisticsTimeScoreItemVO`

NewHealthStatisticsTimeScoreItemVO instantiates a new HealthStatisticsTimeScoreItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStatisticsTimeScoreItemVOWithDefaults

`func NewHealthStatisticsTimeScoreItemVOWithDefaults() *HealthStatisticsTimeScoreItemVO`

NewHealthStatisticsTimeScoreItemVOWithDefaults instantiates a new HealthStatisticsTimeScoreItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScoreDetail

`func (o *HealthStatisticsTimeScoreItemVO) GetScoreDetail() HealthStatisticsScoreVO`

GetScoreDetail returns the ScoreDetail field if non-nil, zero value otherwise.

### GetScoreDetailOk

`func (o *HealthStatisticsTimeScoreItemVO) GetScoreDetailOk() (*HealthStatisticsScoreVO, bool)`

GetScoreDetailOk returns a tuple with the ScoreDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreDetail

`func (o *HealthStatisticsTimeScoreItemVO) SetScoreDetail(v HealthStatisticsScoreVO)`

SetScoreDetail sets ScoreDetail field to given value.

### HasScoreDetail

`func (o *HealthStatisticsTimeScoreItemVO) HasScoreDetail() bool`

HasScoreDetail returns a boolean if a field has been set.

### GetTime

`func (o *HealthStatisticsTimeScoreItemVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HealthStatisticsTimeScoreItemVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HealthStatisticsTimeScoreItemVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *HealthStatisticsTimeScoreItemVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


