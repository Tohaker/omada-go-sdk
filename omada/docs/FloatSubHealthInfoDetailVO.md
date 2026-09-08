# FloatSubHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageFloat** | Pointer to **float32** | Average float value of common dimension | [optional] 
**AverageNum** | Pointer to **int32** | Average value of common dimension, such as cpu、memory | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**PastFloats** | Pointer to [**[]TimeFloatValueItemVO**](TimeFloatValueItemVO.md) | List of common dimension float value | [optional] 
**PastNums** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of common dimension value , such as cpu、memory | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewFloatSubHealthInfoDetailVO

`func NewFloatSubHealthInfoDetailVO() *FloatSubHealthInfoDetailVO`

NewFloatSubHealthInfoDetailVO instantiates a new FloatSubHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatSubHealthInfoDetailVOWithDefaults

`func NewFloatSubHealthInfoDetailVOWithDefaults() *FloatSubHealthInfoDetailVO`

NewFloatSubHealthInfoDetailVOWithDefaults instantiates a new FloatSubHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageFloat

`func (o *FloatSubHealthInfoDetailVO) GetAverageFloat() float32`

GetAverageFloat returns the AverageFloat field if non-nil, zero value otherwise.

### GetAverageFloatOk

`func (o *FloatSubHealthInfoDetailVO) GetAverageFloatOk() (*float32, bool)`

GetAverageFloatOk returns a tuple with the AverageFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageFloat

`func (o *FloatSubHealthInfoDetailVO) SetAverageFloat(v float32)`

SetAverageFloat sets AverageFloat field to given value.

### HasAverageFloat

`func (o *FloatSubHealthInfoDetailVO) HasAverageFloat() bool`

HasAverageFloat returns a boolean if a field has been set.

### GetAverageNum

`func (o *FloatSubHealthInfoDetailVO) GetAverageNum() int32`

GetAverageNum returns the AverageNum field if non-nil, zero value otherwise.

### GetAverageNumOk

`func (o *FloatSubHealthInfoDetailVO) GetAverageNumOk() (*int32, bool)`

GetAverageNumOk returns a tuple with the AverageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum

`func (o *FloatSubHealthInfoDetailVO) SetAverageNum(v int32)`

SetAverageNum sets AverageNum field to given value.

### HasAverageNum

`func (o *FloatSubHealthInfoDetailVO) HasAverageNum() bool`

HasAverageNum returns a boolean if a field has been set.

### GetIncidents

`func (o *FloatSubHealthInfoDetailVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *FloatSubHealthInfoDetailVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *FloatSubHealthInfoDetailVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *FloatSubHealthInfoDetailVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetPastFloats

`func (o *FloatSubHealthInfoDetailVO) GetPastFloats() []TimeFloatValueItemVO`

GetPastFloats returns the PastFloats field if non-nil, zero value otherwise.

### GetPastFloatsOk

`func (o *FloatSubHealthInfoDetailVO) GetPastFloatsOk() (*[]TimeFloatValueItemVO, bool)`

GetPastFloatsOk returns a tuple with the PastFloats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastFloats

`func (o *FloatSubHealthInfoDetailVO) SetPastFloats(v []TimeFloatValueItemVO)`

SetPastFloats sets PastFloats field to given value.

### HasPastFloats

`func (o *FloatSubHealthInfoDetailVO) HasPastFloats() bool`

HasPastFloats returns a boolean if a field has been set.

### GetPastNums

`func (o *FloatSubHealthInfoDetailVO) GetPastNums() []TimeValueItemVO`

GetPastNums returns the PastNums field if non-nil, zero value otherwise.

### GetPastNumsOk

`func (o *FloatSubHealthInfoDetailVO) GetPastNumsOk() (*[]TimeValueItemVO, bool)`

GetPastNumsOk returns a tuple with the PastNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums

`func (o *FloatSubHealthInfoDetailVO) SetPastNums(v []TimeValueItemVO)`

SetPastNums sets PastNums field to given value.

### HasPastNums

`func (o *FloatSubHealthInfoDetailVO) HasPastNums() bool`

HasPastNums returns a boolean if a field has been set.

### GetSummaryScore

`func (o *FloatSubHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *FloatSubHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *FloatSubHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *FloatSubHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *FloatSubHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *FloatSubHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *FloatSubHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *FloatSubHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


