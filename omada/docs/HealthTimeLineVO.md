# HealthTimeLineVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeline** | Pointer to [**[]MacTimeScoreListVO**](MacTimeScoreListVO.md) | Health timeline | [optional] 

## Methods

### NewHealthTimeLineVO

`func NewHealthTimeLineVO() *HealthTimeLineVO`

NewHealthTimeLineVO instantiates a new HealthTimeLineVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthTimeLineVOWithDefaults

`func NewHealthTimeLineVOWithDefaults() *HealthTimeLineVO`

NewHealthTimeLineVOWithDefaults instantiates a new HealthTimeLineVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeline

`func (o *HealthTimeLineVO) GetTimeline() []MacTimeScoreListVO`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *HealthTimeLineVO) GetTimelineOk() (*[]MacTimeScoreListVO, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *HealthTimeLineVO) SetTimeline(v []MacTimeScoreListVO)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *HealthTimeLineVO) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


