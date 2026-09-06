# ConnectScoreSubHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectStatus** | Pointer to **int32** | Connect status of the wired client | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewConnectScoreSubHealthInfoDetailVO

`func NewConnectScoreSubHealthInfoDetailVO() *ConnectScoreSubHealthInfoDetailVO`

NewConnectScoreSubHealthInfoDetailVO instantiates a new ConnectScoreSubHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectScoreSubHealthInfoDetailVOWithDefaults

`func NewConnectScoreSubHealthInfoDetailVOWithDefaults() *ConnectScoreSubHealthInfoDetailVO`

NewConnectScoreSubHealthInfoDetailVOWithDefaults instantiates a new ConnectScoreSubHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectStatus

`func (o *ConnectScoreSubHealthInfoDetailVO) GetConnectStatus() int32`

GetConnectStatus returns the ConnectStatus field if non-nil, zero value otherwise.

### GetConnectStatusOk

`func (o *ConnectScoreSubHealthInfoDetailVO) GetConnectStatusOk() (*int32, bool)`

GetConnectStatusOk returns a tuple with the ConnectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectStatus

`func (o *ConnectScoreSubHealthInfoDetailVO) SetConnectStatus(v int32)`

SetConnectStatus sets ConnectStatus field to given value.

### HasConnectStatus

`func (o *ConnectScoreSubHealthInfoDetailVO) HasConnectStatus() bool`

HasConnectStatus returns a boolean if a field has been set.

### GetIncidents

`func (o *ConnectScoreSubHealthInfoDetailVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ConnectScoreSubHealthInfoDetailVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ConnectScoreSubHealthInfoDetailVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ConnectScoreSubHealthInfoDetailVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetSummaryScore

`func (o *ConnectScoreSubHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *ConnectScoreSubHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *ConnectScoreSubHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *ConnectScoreSubHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *ConnectScoreSubHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *ConnectScoreSubHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *ConnectScoreSubHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *ConnectScoreSubHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


