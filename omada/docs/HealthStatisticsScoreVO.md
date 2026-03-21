# HealthStatisticsScoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageCount** | Pointer to **int32** | The number of average level device or client, whose health score is between 4 and 7 | [optional] 
**GoodCount** | Pointer to **int32** | The number of good level device or client, whose health score is between 8 and 10 | [optional] 
**NoDataCount** | Pointer to **int32** | The number of noData level device or client, whose health score is 0 | [optional] 
**PoorCount** | Pointer to **int32** | The number of good poor device or client, whose health score is between 1 and 3 | [optional] 

## Methods

### NewHealthStatisticsScoreVO

`func NewHealthStatisticsScoreVO() *HealthStatisticsScoreVO`

NewHealthStatisticsScoreVO instantiates a new HealthStatisticsScoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStatisticsScoreVOWithDefaults

`func NewHealthStatisticsScoreVOWithDefaults() *HealthStatisticsScoreVO`

NewHealthStatisticsScoreVOWithDefaults instantiates a new HealthStatisticsScoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageCount

`func (o *HealthStatisticsScoreVO) GetAverageCount() int32`

GetAverageCount returns the AverageCount field if non-nil, zero value otherwise.

### GetAverageCountOk

`func (o *HealthStatisticsScoreVO) GetAverageCountOk() (*int32, bool)`

GetAverageCountOk returns a tuple with the AverageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCount

`func (o *HealthStatisticsScoreVO) SetAverageCount(v int32)`

SetAverageCount sets AverageCount field to given value.

### HasAverageCount

`func (o *HealthStatisticsScoreVO) HasAverageCount() bool`

HasAverageCount returns a boolean if a field has been set.

### GetGoodCount

`func (o *HealthStatisticsScoreVO) GetGoodCount() int32`

GetGoodCount returns the GoodCount field if non-nil, zero value otherwise.

### GetGoodCountOk

`func (o *HealthStatisticsScoreVO) GetGoodCountOk() (*int32, bool)`

GetGoodCountOk returns a tuple with the GoodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodCount

`func (o *HealthStatisticsScoreVO) SetGoodCount(v int32)`

SetGoodCount sets GoodCount field to given value.

### HasGoodCount

`func (o *HealthStatisticsScoreVO) HasGoodCount() bool`

HasGoodCount returns a boolean if a field has been set.

### GetNoDataCount

`func (o *HealthStatisticsScoreVO) GetNoDataCount() int32`

GetNoDataCount returns the NoDataCount field if non-nil, zero value otherwise.

### GetNoDataCountOk

`func (o *HealthStatisticsScoreVO) GetNoDataCountOk() (*int32, bool)`

GetNoDataCountOk returns a tuple with the NoDataCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataCount

`func (o *HealthStatisticsScoreVO) SetNoDataCount(v int32)`

SetNoDataCount sets NoDataCount field to given value.

### HasNoDataCount

`func (o *HealthStatisticsScoreVO) HasNoDataCount() bool`

HasNoDataCount returns a boolean if a field has been set.

### GetPoorCount

`func (o *HealthStatisticsScoreVO) GetPoorCount() int32`

GetPoorCount returns the PoorCount field if non-nil, zero value otherwise.

### GetPoorCountOk

`func (o *HealthStatisticsScoreVO) GetPoorCountOk() (*int32, bool)`

GetPoorCountOk returns a tuple with the PoorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoorCount

`func (o *HealthStatisticsScoreVO) SetPoorCount(v int32)`

SetPoorCount sets PoorCount field to given value.

### HasPoorCount

`func (o *HealthStatisticsScoreVO) HasPoorCount() bool`

HasPoorCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


