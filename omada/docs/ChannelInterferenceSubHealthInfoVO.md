# ChannelInterferenceSubHealthInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageInterf2g** | Pointer to **int32** | Average 2.4G interference | [optional] 
**AverageInterf5g** | Pointer to **int32** | Average 5G interference | [optional] 
**AverageInterf6g** | Pointer to **int32** | Average 6G interference | [optional] 
**AverageNoiseFloor** | Pointer to **int32** | Average noise floor | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**Scores** | Pointer to [**[]TimeScoreItemVO**](TimeScoreItemVO.md) | List of time score items | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewChannelInterferenceSubHealthInfoVO

`func NewChannelInterferenceSubHealthInfoVO() *ChannelInterferenceSubHealthInfoVO`

NewChannelInterferenceSubHealthInfoVO instantiates a new ChannelInterferenceSubHealthInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelInterferenceSubHealthInfoVOWithDefaults

`func NewChannelInterferenceSubHealthInfoVOWithDefaults() *ChannelInterferenceSubHealthInfoVO`

NewChannelInterferenceSubHealthInfoVOWithDefaults instantiates a new ChannelInterferenceSubHealthInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageInterf2g

`func (o *ChannelInterferenceSubHealthInfoVO) GetAverageInterf2g() int32`

GetAverageInterf2g returns the AverageInterf2g field if non-nil, zero value otherwise.

### GetAverageInterf2gOk

`func (o *ChannelInterferenceSubHealthInfoVO) GetAverageInterf2gOk() (*int32, bool)`

GetAverageInterf2gOk returns a tuple with the AverageInterf2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageInterf2g

`func (o *ChannelInterferenceSubHealthInfoVO) SetAverageInterf2g(v int32)`

SetAverageInterf2g sets AverageInterf2g field to given value.

### HasAverageInterf2g

`func (o *ChannelInterferenceSubHealthInfoVO) HasAverageInterf2g() bool`

HasAverageInterf2g returns a boolean if a field has been set.

### GetAverageInterf5g

`func (o *ChannelInterferenceSubHealthInfoVO) GetAverageInterf5g() int32`

GetAverageInterf5g returns the AverageInterf5g field if non-nil, zero value otherwise.

### GetAverageInterf5gOk

`func (o *ChannelInterferenceSubHealthInfoVO) GetAverageInterf5gOk() (*int32, bool)`

GetAverageInterf5gOk returns a tuple with the AverageInterf5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageInterf5g

`func (o *ChannelInterferenceSubHealthInfoVO) SetAverageInterf5g(v int32)`

SetAverageInterf5g sets AverageInterf5g field to given value.

### HasAverageInterf5g

`func (o *ChannelInterferenceSubHealthInfoVO) HasAverageInterf5g() bool`

HasAverageInterf5g returns a boolean if a field has been set.

### GetAverageInterf6g

`func (o *ChannelInterferenceSubHealthInfoVO) GetAverageInterf6g() int32`

GetAverageInterf6g returns the AverageInterf6g field if non-nil, zero value otherwise.

### GetAverageInterf6gOk

`func (o *ChannelInterferenceSubHealthInfoVO) GetAverageInterf6gOk() (*int32, bool)`

GetAverageInterf6gOk returns a tuple with the AverageInterf6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageInterf6g

`func (o *ChannelInterferenceSubHealthInfoVO) SetAverageInterf6g(v int32)`

SetAverageInterf6g sets AverageInterf6g field to given value.

### HasAverageInterf6g

`func (o *ChannelInterferenceSubHealthInfoVO) HasAverageInterf6g() bool`

HasAverageInterf6g returns a boolean if a field has been set.

### GetAverageNoiseFloor

`func (o *ChannelInterferenceSubHealthInfoVO) GetAverageNoiseFloor() int32`

GetAverageNoiseFloor returns the AverageNoiseFloor field if non-nil, zero value otherwise.

### GetAverageNoiseFloorOk

`func (o *ChannelInterferenceSubHealthInfoVO) GetAverageNoiseFloorOk() (*int32, bool)`

GetAverageNoiseFloorOk returns a tuple with the AverageNoiseFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNoiseFloor

`func (o *ChannelInterferenceSubHealthInfoVO) SetAverageNoiseFloor(v int32)`

SetAverageNoiseFloor sets AverageNoiseFloor field to given value.

### HasAverageNoiseFloor

`func (o *ChannelInterferenceSubHealthInfoVO) HasAverageNoiseFloor() bool`

HasAverageNoiseFloor returns a boolean if a field has been set.

### GetIncidents

`func (o *ChannelInterferenceSubHealthInfoVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ChannelInterferenceSubHealthInfoVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ChannelInterferenceSubHealthInfoVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ChannelInterferenceSubHealthInfoVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetScores

`func (o *ChannelInterferenceSubHealthInfoVO) GetScores() []TimeScoreItemVO`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *ChannelInterferenceSubHealthInfoVO) GetScoresOk() (*[]TimeScoreItemVO, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *ChannelInterferenceSubHealthInfoVO) SetScores(v []TimeScoreItemVO)`

SetScores sets Scores field to given value.

### HasScores

`func (o *ChannelInterferenceSubHealthInfoVO) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetSummaryScore

`func (o *ChannelInterferenceSubHealthInfoVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *ChannelInterferenceSubHealthInfoVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *ChannelInterferenceSubHealthInfoVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *ChannelInterferenceSubHealthInfoVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *ChannelInterferenceSubHealthInfoVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *ChannelInterferenceSubHealthInfoVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *ChannelInterferenceSubHealthInfoVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *ChannelInterferenceSubHealthInfoVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


