# ClientScoreTimelineListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scores** | Pointer to [**[]HealthStatisticsTimeScoreItemVO**](HealthStatisticsTimeScoreItemVO.md) |  | [optional] 

## Methods

### NewClientScoreTimelineListVO

`func NewClientScoreTimelineListVO() *ClientScoreTimelineListVO`

NewClientScoreTimelineListVO instantiates a new ClientScoreTimelineListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientScoreTimelineListVOWithDefaults

`func NewClientScoreTimelineListVOWithDefaults() *ClientScoreTimelineListVO`

NewClientScoreTimelineListVOWithDefaults instantiates a new ClientScoreTimelineListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScores

`func (o *ClientScoreTimelineListVO) GetScores() []HealthStatisticsTimeScoreItemVO`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *ClientScoreTimelineListVO) GetScoresOk() (*[]HealthStatisticsTimeScoreItemVO, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *ClientScoreTimelineListVO) SetScores(v []HealthStatisticsTimeScoreItemVO)`

SetScores sets Scores field to given value.

### HasScores

`func (o *ClientScoreTimelineListVO) HasScores() bool`

HasScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


