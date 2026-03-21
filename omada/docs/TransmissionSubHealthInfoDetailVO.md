# TransmissionSubHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageNumTxDrop** | Pointer to **int32** | Average value of tx drop | [optional] 
**AverageNumTxRetry** | Pointer to **int32** | Average value of tx retry | [optional] 
**PastNumsTxDrop** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of tx drop | [optional] 
**PastNumsTxRetry** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of tx retry | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewTransmissionSubHealthInfoDetailVO

`func NewTransmissionSubHealthInfoDetailVO() *TransmissionSubHealthInfoDetailVO`

NewTransmissionSubHealthInfoDetailVO instantiates a new TransmissionSubHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransmissionSubHealthInfoDetailVOWithDefaults

`func NewTransmissionSubHealthInfoDetailVOWithDefaults() *TransmissionSubHealthInfoDetailVO`

NewTransmissionSubHealthInfoDetailVOWithDefaults instantiates a new TransmissionSubHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageNumTxDrop

`func (o *TransmissionSubHealthInfoDetailVO) GetAverageNumTxDrop() int32`

GetAverageNumTxDrop returns the AverageNumTxDrop field if non-nil, zero value otherwise.

### GetAverageNumTxDropOk

`func (o *TransmissionSubHealthInfoDetailVO) GetAverageNumTxDropOk() (*int32, bool)`

GetAverageNumTxDropOk returns a tuple with the AverageNumTxDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNumTxDrop

`func (o *TransmissionSubHealthInfoDetailVO) SetAverageNumTxDrop(v int32)`

SetAverageNumTxDrop sets AverageNumTxDrop field to given value.

### HasAverageNumTxDrop

`func (o *TransmissionSubHealthInfoDetailVO) HasAverageNumTxDrop() bool`

HasAverageNumTxDrop returns a boolean if a field has been set.

### GetAverageNumTxRetry

`func (o *TransmissionSubHealthInfoDetailVO) GetAverageNumTxRetry() int32`

GetAverageNumTxRetry returns the AverageNumTxRetry field if non-nil, zero value otherwise.

### GetAverageNumTxRetryOk

`func (o *TransmissionSubHealthInfoDetailVO) GetAverageNumTxRetryOk() (*int32, bool)`

GetAverageNumTxRetryOk returns a tuple with the AverageNumTxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNumTxRetry

`func (o *TransmissionSubHealthInfoDetailVO) SetAverageNumTxRetry(v int32)`

SetAverageNumTxRetry sets AverageNumTxRetry field to given value.

### HasAverageNumTxRetry

`func (o *TransmissionSubHealthInfoDetailVO) HasAverageNumTxRetry() bool`

HasAverageNumTxRetry returns a boolean if a field has been set.

### GetPastNumsTxDrop

`func (o *TransmissionSubHealthInfoDetailVO) GetPastNumsTxDrop() []TimeValueItemVO`

GetPastNumsTxDrop returns the PastNumsTxDrop field if non-nil, zero value otherwise.

### GetPastNumsTxDropOk

`func (o *TransmissionSubHealthInfoDetailVO) GetPastNumsTxDropOk() (*[]TimeValueItemVO, bool)`

GetPastNumsTxDropOk returns a tuple with the PastNumsTxDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNumsTxDrop

`func (o *TransmissionSubHealthInfoDetailVO) SetPastNumsTxDrop(v []TimeValueItemVO)`

SetPastNumsTxDrop sets PastNumsTxDrop field to given value.

### HasPastNumsTxDrop

`func (o *TransmissionSubHealthInfoDetailVO) HasPastNumsTxDrop() bool`

HasPastNumsTxDrop returns a boolean if a field has been set.

### GetPastNumsTxRetry

`func (o *TransmissionSubHealthInfoDetailVO) GetPastNumsTxRetry() []TimeValueItemVO`

GetPastNumsTxRetry returns the PastNumsTxRetry field if non-nil, zero value otherwise.

### GetPastNumsTxRetryOk

`func (o *TransmissionSubHealthInfoDetailVO) GetPastNumsTxRetryOk() (*[]TimeValueItemVO, bool)`

GetPastNumsTxRetryOk returns a tuple with the PastNumsTxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNumsTxRetry

`func (o *TransmissionSubHealthInfoDetailVO) SetPastNumsTxRetry(v []TimeValueItemVO)`

SetPastNumsTxRetry sets PastNumsTxRetry field to given value.

### HasPastNumsTxRetry

`func (o *TransmissionSubHealthInfoDetailVO) HasPastNumsTxRetry() bool`

HasPastNumsTxRetry returns a boolean if a field has been set.

### GetSummaryScore

`func (o *TransmissionSubHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *TransmissionSubHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *TransmissionSubHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *TransmissionSubHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *TransmissionSubHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *TransmissionSubHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *TransmissionSubHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *TransmissionSubHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


