# LinkErrorHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageNum** | Pointer to **int32** | Average value of common dimension, such as cpu、memory | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**PastNums** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of common dimension value , such as cpu、memory | [optional] 
**PortNum** | Pointer to **int32** | Total number of ports. | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewLinkErrorHealthInfoDetailVO

`func NewLinkErrorHealthInfoDetailVO() *LinkErrorHealthInfoDetailVO`

NewLinkErrorHealthInfoDetailVO instantiates a new LinkErrorHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkErrorHealthInfoDetailVOWithDefaults

`func NewLinkErrorHealthInfoDetailVOWithDefaults() *LinkErrorHealthInfoDetailVO`

NewLinkErrorHealthInfoDetailVOWithDefaults instantiates a new LinkErrorHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageNum

`func (o *LinkErrorHealthInfoDetailVO) GetAverageNum() int32`

GetAverageNum returns the AverageNum field if non-nil, zero value otherwise.

### GetAverageNumOk

`func (o *LinkErrorHealthInfoDetailVO) GetAverageNumOk() (*int32, bool)`

GetAverageNumOk returns a tuple with the AverageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum

`func (o *LinkErrorHealthInfoDetailVO) SetAverageNum(v int32)`

SetAverageNum sets AverageNum field to given value.

### HasAverageNum

`func (o *LinkErrorHealthInfoDetailVO) HasAverageNum() bool`

HasAverageNum returns a boolean if a field has been set.

### GetIncidents

`func (o *LinkErrorHealthInfoDetailVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *LinkErrorHealthInfoDetailVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *LinkErrorHealthInfoDetailVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *LinkErrorHealthInfoDetailVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetPastNums

`func (o *LinkErrorHealthInfoDetailVO) GetPastNums() []TimeValueItemVO`

GetPastNums returns the PastNums field if non-nil, zero value otherwise.

### GetPastNumsOk

`func (o *LinkErrorHealthInfoDetailVO) GetPastNumsOk() (*[]TimeValueItemVO, bool)`

GetPastNumsOk returns a tuple with the PastNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums

`func (o *LinkErrorHealthInfoDetailVO) SetPastNums(v []TimeValueItemVO)`

SetPastNums sets PastNums field to given value.

### HasPastNums

`func (o *LinkErrorHealthInfoDetailVO) HasPastNums() bool`

HasPastNums returns a boolean if a field has been set.

### GetPortNum

`func (o *LinkErrorHealthInfoDetailVO) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *LinkErrorHealthInfoDetailVO) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *LinkErrorHealthInfoDetailVO) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *LinkErrorHealthInfoDetailVO) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.

### GetSummaryScore

`func (o *LinkErrorHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *LinkErrorHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *LinkErrorHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *LinkErrorHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *LinkErrorHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *LinkErrorHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *LinkErrorHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *LinkErrorHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


