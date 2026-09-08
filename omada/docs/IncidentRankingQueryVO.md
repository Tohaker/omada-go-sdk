# IncidentRankingQueryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | **int32** | For the values of Anomaly event code, refer to section 5.7.2.1 of the Open API Access | 
**EndTime** | **int64** | End time in milliseconds | 
**StartTime** | **int64** | Start time in milliseconds | 
**TopK** | Pointer to **int32** | Number of top items to return. Defaults to 5 if not provided, maximum to 20. | [optional] 

## Methods

### NewIncidentRankingQueryVO

`func NewIncidentRankingQueryVO(anomalyCode int32, endTime int64, startTime int64, ) *IncidentRankingQueryVO`

NewIncidentRankingQueryVO instantiates a new IncidentRankingQueryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentRankingQueryVOWithDefaults

`func NewIncidentRankingQueryVOWithDefaults() *IncidentRankingQueryVO`

NewIncidentRankingQueryVOWithDefaults instantiates a new IncidentRankingQueryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *IncidentRankingQueryVO) GetAnomalyCode() int32`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *IncidentRankingQueryVO) GetAnomalyCodeOk() (*int32, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *IncidentRankingQueryVO) SetAnomalyCode(v int32)`

SetAnomalyCode sets AnomalyCode field to given value.


### GetEndTime

`func (o *IncidentRankingQueryVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *IncidentRankingQueryVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *IncidentRankingQueryVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetStartTime

`func (o *IncidentRankingQueryVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *IncidentRankingQueryVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *IncidentRankingQueryVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetTopK

`func (o *IncidentRankingQueryVO) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *IncidentRankingQueryVO) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *IncidentRankingQueryVO) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *IncidentRankingQueryVO) HasTopK() bool`

HasTopK returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


