# ChannelInterferenceSubHealthDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageNoiseFloor** | Pointer to **int32** | Average value of noise floor | [optional] 
**AverageNum2g** | Pointer to **int32** | Average value of 2g channel | [optional] 
**AverageNum5g** | Pointer to **int32** | Average value of 5g channel | [optional] 
**AverageNum6g** | Pointer to **int32** | Average value of 6g channel | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**PastNoiseFloor** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of noise floor | [optional] 
**PastNums2g** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of 2g channel value | [optional] 
**PastNums5g** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of 5g channel value | [optional] 
**PastNums6g** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of 6g channel value | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewChannelInterferenceSubHealthDetailVO

`func NewChannelInterferenceSubHealthDetailVO() *ChannelInterferenceSubHealthDetailVO`

NewChannelInterferenceSubHealthDetailVO instantiates a new ChannelInterferenceSubHealthDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelInterferenceSubHealthDetailVOWithDefaults

`func NewChannelInterferenceSubHealthDetailVOWithDefaults() *ChannelInterferenceSubHealthDetailVO`

NewChannelInterferenceSubHealthDetailVOWithDefaults instantiates a new ChannelInterferenceSubHealthDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageNoiseFloor

`func (o *ChannelInterferenceSubHealthDetailVO) GetAverageNoiseFloor() int32`

GetAverageNoiseFloor returns the AverageNoiseFloor field if non-nil, zero value otherwise.

### GetAverageNoiseFloorOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetAverageNoiseFloorOk() (*int32, bool)`

GetAverageNoiseFloorOk returns a tuple with the AverageNoiseFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNoiseFloor

`func (o *ChannelInterferenceSubHealthDetailVO) SetAverageNoiseFloor(v int32)`

SetAverageNoiseFloor sets AverageNoiseFloor field to given value.

### HasAverageNoiseFloor

`func (o *ChannelInterferenceSubHealthDetailVO) HasAverageNoiseFloor() bool`

HasAverageNoiseFloor returns a boolean if a field has been set.

### GetAverageNum2g

`func (o *ChannelInterferenceSubHealthDetailVO) GetAverageNum2g() int32`

GetAverageNum2g returns the AverageNum2g field if non-nil, zero value otherwise.

### GetAverageNum2gOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetAverageNum2gOk() (*int32, bool)`

GetAverageNum2gOk returns a tuple with the AverageNum2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum2g

`func (o *ChannelInterferenceSubHealthDetailVO) SetAverageNum2g(v int32)`

SetAverageNum2g sets AverageNum2g field to given value.

### HasAverageNum2g

`func (o *ChannelInterferenceSubHealthDetailVO) HasAverageNum2g() bool`

HasAverageNum2g returns a boolean if a field has been set.

### GetAverageNum5g

`func (o *ChannelInterferenceSubHealthDetailVO) GetAverageNum5g() int32`

GetAverageNum5g returns the AverageNum5g field if non-nil, zero value otherwise.

### GetAverageNum5gOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetAverageNum5gOk() (*int32, bool)`

GetAverageNum5gOk returns a tuple with the AverageNum5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum5g

`func (o *ChannelInterferenceSubHealthDetailVO) SetAverageNum5g(v int32)`

SetAverageNum5g sets AverageNum5g field to given value.

### HasAverageNum5g

`func (o *ChannelInterferenceSubHealthDetailVO) HasAverageNum5g() bool`

HasAverageNum5g returns a boolean if a field has been set.

### GetAverageNum6g

`func (o *ChannelInterferenceSubHealthDetailVO) GetAverageNum6g() int32`

GetAverageNum6g returns the AverageNum6g field if non-nil, zero value otherwise.

### GetAverageNum6gOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetAverageNum6gOk() (*int32, bool)`

GetAverageNum6gOk returns a tuple with the AverageNum6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum6g

`func (o *ChannelInterferenceSubHealthDetailVO) SetAverageNum6g(v int32)`

SetAverageNum6g sets AverageNum6g field to given value.

### HasAverageNum6g

`func (o *ChannelInterferenceSubHealthDetailVO) HasAverageNum6g() bool`

HasAverageNum6g returns a boolean if a field has been set.

### GetIncidents

`func (o *ChannelInterferenceSubHealthDetailVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ChannelInterferenceSubHealthDetailVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ChannelInterferenceSubHealthDetailVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetPastNoiseFloor

`func (o *ChannelInterferenceSubHealthDetailVO) GetPastNoiseFloor() []TimeValueItemVO`

GetPastNoiseFloor returns the PastNoiseFloor field if non-nil, zero value otherwise.

### GetPastNoiseFloorOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetPastNoiseFloorOk() (*[]TimeValueItemVO, bool)`

GetPastNoiseFloorOk returns a tuple with the PastNoiseFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNoiseFloor

`func (o *ChannelInterferenceSubHealthDetailVO) SetPastNoiseFloor(v []TimeValueItemVO)`

SetPastNoiseFloor sets PastNoiseFloor field to given value.

### HasPastNoiseFloor

`func (o *ChannelInterferenceSubHealthDetailVO) HasPastNoiseFloor() bool`

HasPastNoiseFloor returns a boolean if a field has been set.

### GetPastNums2g

`func (o *ChannelInterferenceSubHealthDetailVO) GetPastNums2g() []TimeValueItemVO`

GetPastNums2g returns the PastNums2g field if non-nil, zero value otherwise.

### GetPastNums2gOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetPastNums2gOk() (*[]TimeValueItemVO, bool)`

GetPastNums2gOk returns a tuple with the PastNums2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums2g

`func (o *ChannelInterferenceSubHealthDetailVO) SetPastNums2g(v []TimeValueItemVO)`

SetPastNums2g sets PastNums2g field to given value.

### HasPastNums2g

`func (o *ChannelInterferenceSubHealthDetailVO) HasPastNums2g() bool`

HasPastNums2g returns a boolean if a field has been set.

### GetPastNums5g

`func (o *ChannelInterferenceSubHealthDetailVO) GetPastNums5g() []TimeValueItemVO`

GetPastNums5g returns the PastNums5g field if non-nil, zero value otherwise.

### GetPastNums5gOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetPastNums5gOk() (*[]TimeValueItemVO, bool)`

GetPastNums5gOk returns a tuple with the PastNums5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums5g

`func (o *ChannelInterferenceSubHealthDetailVO) SetPastNums5g(v []TimeValueItemVO)`

SetPastNums5g sets PastNums5g field to given value.

### HasPastNums5g

`func (o *ChannelInterferenceSubHealthDetailVO) HasPastNums5g() bool`

HasPastNums5g returns a boolean if a field has been set.

### GetPastNums6g

`func (o *ChannelInterferenceSubHealthDetailVO) GetPastNums6g() []TimeValueItemVO`

GetPastNums6g returns the PastNums6g field if non-nil, zero value otherwise.

### GetPastNums6gOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetPastNums6gOk() (*[]TimeValueItemVO, bool)`

GetPastNums6gOk returns a tuple with the PastNums6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums6g

`func (o *ChannelInterferenceSubHealthDetailVO) SetPastNums6g(v []TimeValueItemVO)`

SetPastNums6g sets PastNums6g field to given value.

### HasPastNums6g

`func (o *ChannelInterferenceSubHealthDetailVO) HasPastNums6g() bool`

HasPastNums6g returns a boolean if a field has been set.

### GetSummaryScore

`func (o *ChannelInterferenceSubHealthDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *ChannelInterferenceSubHealthDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *ChannelInterferenceSubHealthDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *ChannelInterferenceSubHealthDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *ChannelInterferenceSubHealthDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *ChannelInterferenceSubHealthDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *ChannelInterferenceSubHealthDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


