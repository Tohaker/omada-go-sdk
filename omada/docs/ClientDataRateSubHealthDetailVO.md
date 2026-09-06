# ClientDataRateSubHealthDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageRxRate** | Pointer to **int64** |  | [optional] 
**AverageTxRate** | Pointer to **int64** |  | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**PastRxRate** | Pointer to [**[]LongTimeValueItemVO**](LongTimeValueItemVO.md) |  | [optional] 
**PastTxRate** | Pointer to [**[]LongTimeValueItemVO**](LongTimeValueItemVO.md) |  | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewClientDataRateSubHealthDetailVO

`func NewClientDataRateSubHealthDetailVO() *ClientDataRateSubHealthDetailVO`

NewClientDataRateSubHealthDetailVO instantiates a new ClientDataRateSubHealthDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDataRateSubHealthDetailVOWithDefaults

`func NewClientDataRateSubHealthDetailVOWithDefaults() *ClientDataRateSubHealthDetailVO`

NewClientDataRateSubHealthDetailVOWithDefaults instantiates a new ClientDataRateSubHealthDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageRxRate

`func (o *ClientDataRateSubHealthDetailVO) GetAverageRxRate() int64`

GetAverageRxRate returns the AverageRxRate field if non-nil, zero value otherwise.

### GetAverageRxRateOk

`func (o *ClientDataRateSubHealthDetailVO) GetAverageRxRateOk() (*int64, bool)`

GetAverageRxRateOk returns a tuple with the AverageRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxRate

`func (o *ClientDataRateSubHealthDetailVO) SetAverageRxRate(v int64)`

SetAverageRxRate sets AverageRxRate field to given value.

### HasAverageRxRate

`func (o *ClientDataRateSubHealthDetailVO) HasAverageRxRate() bool`

HasAverageRxRate returns a boolean if a field has been set.

### GetAverageTxRate

`func (o *ClientDataRateSubHealthDetailVO) GetAverageTxRate() int64`

GetAverageTxRate returns the AverageTxRate field if non-nil, zero value otherwise.

### GetAverageTxRateOk

`func (o *ClientDataRateSubHealthDetailVO) GetAverageTxRateOk() (*int64, bool)`

GetAverageTxRateOk returns a tuple with the AverageTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxRate

`func (o *ClientDataRateSubHealthDetailVO) SetAverageTxRate(v int64)`

SetAverageTxRate sets AverageTxRate field to given value.

### HasAverageTxRate

`func (o *ClientDataRateSubHealthDetailVO) HasAverageTxRate() bool`

HasAverageTxRate returns a boolean if a field has been set.

### GetIncidents

`func (o *ClientDataRateSubHealthDetailVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ClientDataRateSubHealthDetailVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ClientDataRateSubHealthDetailVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ClientDataRateSubHealthDetailVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetPastRxRate

`func (o *ClientDataRateSubHealthDetailVO) GetPastRxRate() []LongTimeValueItemVO`

GetPastRxRate returns the PastRxRate field if non-nil, zero value otherwise.

### GetPastRxRateOk

`func (o *ClientDataRateSubHealthDetailVO) GetPastRxRateOk() (*[]LongTimeValueItemVO, bool)`

GetPastRxRateOk returns a tuple with the PastRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastRxRate

`func (o *ClientDataRateSubHealthDetailVO) SetPastRxRate(v []LongTimeValueItemVO)`

SetPastRxRate sets PastRxRate field to given value.

### HasPastRxRate

`func (o *ClientDataRateSubHealthDetailVO) HasPastRxRate() bool`

HasPastRxRate returns a boolean if a field has been set.

### GetPastTxRate

`func (o *ClientDataRateSubHealthDetailVO) GetPastTxRate() []LongTimeValueItemVO`

GetPastTxRate returns the PastTxRate field if non-nil, zero value otherwise.

### GetPastTxRateOk

`func (o *ClientDataRateSubHealthDetailVO) GetPastTxRateOk() (*[]LongTimeValueItemVO, bool)`

GetPastTxRateOk returns a tuple with the PastTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastTxRate

`func (o *ClientDataRateSubHealthDetailVO) SetPastTxRate(v []LongTimeValueItemVO)`

SetPastTxRate sets PastTxRate field to given value.

### HasPastTxRate

`func (o *ClientDataRateSubHealthDetailVO) HasPastTxRate() bool`

HasPastTxRate returns a boolean if a field has been set.

### GetSummaryScore

`func (o *ClientDataRateSubHealthDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *ClientDataRateSubHealthDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *ClientDataRateSubHealthDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *ClientDataRateSubHealthDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *ClientDataRateSubHealthDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *ClientDataRateSubHealthDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *ClientDataRateSubHealthDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *ClientDataRateSubHealthDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


