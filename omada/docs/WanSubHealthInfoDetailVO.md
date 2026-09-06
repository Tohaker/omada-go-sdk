# WanSubHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageJitter** | Pointer to **int32** | Average jitter | [optional] 
**AverageLatency** | Pointer to **int32** | Average latency | [optional] 
**AverageMos** | Pointer to **float32** | Average mean opinion score | [optional] 
**AveragePktLoss** | Pointer to **float32** | Average packet loss | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**Jitter** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of jitter | [optional] 
**Latency** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of latency | [optional] 
**Mos** | Pointer to [**[]TimeFloatValueItemVO**](TimeFloatValueItemVO.md) | List of mos | [optional] 
**PktLoss** | Pointer to [**[]TimeFloatValueItemVO**](TimeFloatValueItemVO.md) | List of packet loss | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewWanSubHealthInfoDetailVO

`func NewWanSubHealthInfoDetailVO() *WanSubHealthInfoDetailVO`

NewWanSubHealthInfoDetailVO instantiates a new WanSubHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanSubHealthInfoDetailVOWithDefaults

`func NewWanSubHealthInfoDetailVOWithDefaults() *WanSubHealthInfoDetailVO`

NewWanSubHealthInfoDetailVOWithDefaults instantiates a new WanSubHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageJitter

`func (o *WanSubHealthInfoDetailVO) GetAverageJitter() int32`

GetAverageJitter returns the AverageJitter field if non-nil, zero value otherwise.

### GetAverageJitterOk

`func (o *WanSubHealthInfoDetailVO) GetAverageJitterOk() (*int32, bool)`

GetAverageJitterOk returns a tuple with the AverageJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageJitter

`func (o *WanSubHealthInfoDetailVO) SetAverageJitter(v int32)`

SetAverageJitter sets AverageJitter field to given value.

### HasAverageJitter

`func (o *WanSubHealthInfoDetailVO) HasAverageJitter() bool`

HasAverageJitter returns a boolean if a field has been set.

### GetAverageLatency

`func (o *WanSubHealthInfoDetailVO) GetAverageLatency() int32`

GetAverageLatency returns the AverageLatency field if non-nil, zero value otherwise.

### GetAverageLatencyOk

`func (o *WanSubHealthInfoDetailVO) GetAverageLatencyOk() (*int32, bool)`

GetAverageLatencyOk returns a tuple with the AverageLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageLatency

`func (o *WanSubHealthInfoDetailVO) SetAverageLatency(v int32)`

SetAverageLatency sets AverageLatency field to given value.

### HasAverageLatency

`func (o *WanSubHealthInfoDetailVO) HasAverageLatency() bool`

HasAverageLatency returns a boolean if a field has been set.

### GetAverageMos

`func (o *WanSubHealthInfoDetailVO) GetAverageMos() float32`

GetAverageMos returns the AverageMos field if non-nil, zero value otherwise.

### GetAverageMosOk

`func (o *WanSubHealthInfoDetailVO) GetAverageMosOk() (*float32, bool)`

GetAverageMosOk returns a tuple with the AverageMos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageMos

`func (o *WanSubHealthInfoDetailVO) SetAverageMos(v float32)`

SetAverageMos sets AverageMos field to given value.

### HasAverageMos

`func (o *WanSubHealthInfoDetailVO) HasAverageMos() bool`

HasAverageMos returns a boolean if a field has been set.

### GetAveragePktLoss

`func (o *WanSubHealthInfoDetailVO) GetAveragePktLoss() float32`

GetAveragePktLoss returns the AveragePktLoss field if non-nil, zero value otherwise.

### GetAveragePktLossOk

`func (o *WanSubHealthInfoDetailVO) GetAveragePktLossOk() (*float32, bool)`

GetAveragePktLossOk returns a tuple with the AveragePktLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePktLoss

`func (o *WanSubHealthInfoDetailVO) SetAveragePktLoss(v float32)`

SetAveragePktLoss sets AveragePktLoss field to given value.

### HasAveragePktLoss

`func (o *WanSubHealthInfoDetailVO) HasAveragePktLoss() bool`

HasAveragePktLoss returns a boolean if a field has been set.

### GetIncidents

`func (o *WanSubHealthInfoDetailVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *WanSubHealthInfoDetailVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *WanSubHealthInfoDetailVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *WanSubHealthInfoDetailVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetJitter

`func (o *WanSubHealthInfoDetailVO) GetJitter() []TimeValueItemVO`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *WanSubHealthInfoDetailVO) GetJitterOk() (*[]TimeValueItemVO, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *WanSubHealthInfoDetailVO) SetJitter(v []TimeValueItemVO)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *WanSubHealthInfoDetailVO) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### GetLatency

`func (o *WanSubHealthInfoDetailVO) GetLatency() []TimeValueItemVO`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *WanSubHealthInfoDetailVO) GetLatencyOk() (*[]TimeValueItemVO, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *WanSubHealthInfoDetailVO) SetLatency(v []TimeValueItemVO)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *WanSubHealthInfoDetailVO) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMos

`func (o *WanSubHealthInfoDetailVO) GetMos() []TimeFloatValueItemVO`

GetMos returns the Mos field if non-nil, zero value otherwise.

### GetMosOk

`func (o *WanSubHealthInfoDetailVO) GetMosOk() (*[]TimeFloatValueItemVO, bool)`

GetMosOk returns a tuple with the Mos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMos

`func (o *WanSubHealthInfoDetailVO) SetMos(v []TimeFloatValueItemVO)`

SetMos sets Mos field to given value.

### HasMos

`func (o *WanSubHealthInfoDetailVO) HasMos() bool`

HasMos returns a boolean if a field has been set.

### GetPktLoss

`func (o *WanSubHealthInfoDetailVO) GetPktLoss() []TimeFloatValueItemVO`

GetPktLoss returns the PktLoss field if non-nil, zero value otherwise.

### GetPktLossOk

`func (o *WanSubHealthInfoDetailVO) GetPktLossOk() (*[]TimeFloatValueItemVO, bool)`

GetPktLossOk returns a tuple with the PktLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPktLoss

`func (o *WanSubHealthInfoDetailVO) SetPktLoss(v []TimeFloatValueItemVO)`

SetPktLoss sets PktLoss field to given value.

### HasPktLoss

`func (o *WanSubHealthInfoDetailVO) HasPktLoss() bool`

HasPktLoss returns a boolean if a field has been set.

### GetSummaryScore

`func (o *WanSubHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *WanSubHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *WanSubHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *WanSubHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *WanSubHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *WanSubHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *WanSubHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *WanSubHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


