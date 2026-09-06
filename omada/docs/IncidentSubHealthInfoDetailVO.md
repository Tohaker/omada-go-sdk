# IncidentSubHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentNum** | Pointer to **int32** | Number of incident | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewIncidentSubHealthInfoDetailVO

`func NewIncidentSubHealthInfoDetailVO() *IncidentSubHealthInfoDetailVO`

NewIncidentSubHealthInfoDetailVO instantiates a new IncidentSubHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSubHealthInfoDetailVOWithDefaults

`func NewIncidentSubHealthInfoDetailVOWithDefaults() *IncidentSubHealthInfoDetailVO`

NewIncidentSubHealthInfoDetailVOWithDefaults instantiates a new IncidentSubHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentNum

`func (o *IncidentSubHealthInfoDetailVO) GetIncidentNum() int32`

GetIncidentNum returns the IncidentNum field if non-nil, zero value otherwise.

### GetIncidentNumOk

`func (o *IncidentSubHealthInfoDetailVO) GetIncidentNumOk() (*int32, bool)`

GetIncidentNumOk returns a tuple with the IncidentNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentNum

`func (o *IncidentSubHealthInfoDetailVO) SetIncidentNum(v int32)`

SetIncidentNum sets IncidentNum field to given value.

### HasIncidentNum

`func (o *IncidentSubHealthInfoDetailVO) HasIncidentNum() bool`

HasIncidentNum returns a boolean if a field has been set.

### GetIncidents

`func (o *IncidentSubHealthInfoDetailVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *IncidentSubHealthInfoDetailVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *IncidentSubHealthInfoDetailVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *IncidentSubHealthInfoDetailVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetSummaryScore

`func (o *IncidentSubHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *IncidentSubHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *IncidentSubHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *IncidentSubHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *IncidentSubHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *IncidentSubHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *IncidentSubHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *IncidentSubHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


