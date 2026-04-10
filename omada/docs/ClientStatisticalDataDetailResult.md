# ClientStatisticalDataDetailResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgDownRate** | Pointer to **int64** | Average downlink rate (Byte/s). | [optional] 
**AvgRxR** | Pointer to **int64** | (Wireless) Average uplink negotiation rate (bit/s). | [optional] 
**AvgSignal** | Pointer to **int32** | (Wireless) Average signal strength, unit: dBm. | [optional] 
**AvgTxR** | Pointer to **int64** | (Wireless) Average downlink negotiation rate (bit/s). | [optional] 
**AvgUpRate** | Pointer to **int64** | Average uplink rate (Byte/s). | [optional] 
**Stats** | Pointer to [**[]ClientStatisticalDataDetail**](ClientStatisticalDataDetail.md) | Client Statistical Data Detail list. | [optional] 
**TotalDown** | Pointer to **int64** | Client total downstream traffic (Byte). | [optional] 
**TotalTxFP** | Pointer to **int64** | Total number of downstream failed packets. | [optional] 
**TotalUp** | Pointer to **int64** | Client total upstream traffic (Byte). | [optional] 

## Methods

### NewClientStatisticalDataDetailResult

`func NewClientStatisticalDataDetailResult() *ClientStatisticalDataDetailResult`

NewClientStatisticalDataDetailResult instantiates a new ClientStatisticalDataDetailResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientStatisticalDataDetailResultWithDefaults

`func NewClientStatisticalDataDetailResultWithDefaults() *ClientStatisticalDataDetailResult`

NewClientStatisticalDataDetailResultWithDefaults instantiates a new ClientStatisticalDataDetailResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgDownRate

`func (o *ClientStatisticalDataDetailResult) GetAvgDownRate() int64`

GetAvgDownRate returns the AvgDownRate field if non-nil, zero value otherwise.

### GetAvgDownRateOk

`func (o *ClientStatisticalDataDetailResult) GetAvgDownRateOk() (*int64, bool)`

GetAvgDownRateOk returns a tuple with the AvgDownRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDownRate

`func (o *ClientStatisticalDataDetailResult) SetAvgDownRate(v int64)`

SetAvgDownRate sets AvgDownRate field to given value.

### HasAvgDownRate

`func (o *ClientStatisticalDataDetailResult) HasAvgDownRate() bool`

HasAvgDownRate returns a boolean if a field has been set.

### GetAvgRxR

`func (o *ClientStatisticalDataDetailResult) GetAvgRxR() int64`

GetAvgRxR returns the AvgRxR field if non-nil, zero value otherwise.

### GetAvgRxROk

`func (o *ClientStatisticalDataDetailResult) GetAvgRxROk() (*int64, bool)`

GetAvgRxROk returns a tuple with the AvgRxR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRxR

`func (o *ClientStatisticalDataDetailResult) SetAvgRxR(v int64)`

SetAvgRxR sets AvgRxR field to given value.

### HasAvgRxR

`func (o *ClientStatisticalDataDetailResult) HasAvgRxR() bool`

HasAvgRxR returns a boolean if a field has been set.

### GetAvgSignal

`func (o *ClientStatisticalDataDetailResult) GetAvgSignal() int32`

GetAvgSignal returns the AvgSignal field if non-nil, zero value otherwise.

### GetAvgSignalOk

`func (o *ClientStatisticalDataDetailResult) GetAvgSignalOk() (*int32, bool)`

GetAvgSignalOk returns a tuple with the AvgSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSignal

`func (o *ClientStatisticalDataDetailResult) SetAvgSignal(v int32)`

SetAvgSignal sets AvgSignal field to given value.

### HasAvgSignal

`func (o *ClientStatisticalDataDetailResult) HasAvgSignal() bool`

HasAvgSignal returns a boolean if a field has been set.

### GetAvgTxR

`func (o *ClientStatisticalDataDetailResult) GetAvgTxR() int64`

GetAvgTxR returns the AvgTxR field if non-nil, zero value otherwise.

### GetAvgTxROk

`func (o *ClientStatisticalDataDetailResult) GetAvgTxROk() (*int64, bool)`

GetAvgTxROk returns a tuple with the AvgTxR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTxR

`func (o *ClientStatisticalDataDetailResult) SetAvgTxR(v int64)`

SetAvgTxR sets AvgTxR field to given value.

### HasAvgTxR

`func (o *ClientStatisticalDataDetailResult) HasAvgTxR() bool`

HasAvgTxR returns a boolean if a field has been set.

### GetAvgUpRate

`func (o *ClientStatisticalDataDetailResult) GetAvgUpRate() int64`

GetAvgUpRate returns the AvgUpRate field if non-nil, zero value otherwise.

### GetAvgUpRateOk

`func (o *ClientStatisticalDataDetailResult) GetAvgUpRateOk() (*int64, bool)`

GetAvgUpRateOk returns a tuple with the AvgUpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgUpRate

`func (o *ClientStatisticalDataDetailResult) SetAvgUpRate(v int64)`

SetAvgUpRate sets AvgUpRate field to given value.

### HasAvgUpRate

`func (o *ClientStatisticalDataDetailResult) HasAvgUpRate() bool`

HasAvgUpRate returns a boolean if a field has been set.

### GetStats

`func (o *ClientStatisticalDataDetailResult) GetStats() []ClientStatisticalDataDetail`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ClientStatisticalDataDetailResult) GetStatsOk() (*[]ClientStatisticalDataDetail, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ClientStatisticalDataDetailResult) SetStats(v []ClientStatisticalDataDetail)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ClientStatisticalDataDetailResult) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTotalDown

`func (o *ClientStatisticalDataDetailResult) GetTotalDown() int64`

GetTotalDown returns the TotalDown field if non-nil, zero value otherwise.

### GetTotalDownOk

`func (o *ClientStatisticalDataDetailResult) GetTotalDownOk() (*int64, bool)`

GetTotalDownOk returns a tuple with the TotalDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDown

`func (o *ClientStatisticalDataDetailResult) SetTotalDown(v int64)`

SetTotalDown sets TotalDown field to given value.

### HasTotalDown

`func (o *ClientStatisticalDataDetailResult) HasTotalDown() bool`

HasTotalDown returns a boolean if a field has been set.

### GetTotalTxFP

`func (o *ClientStatisticalDataDetailResult) GetTotalTxFP() int64`

GetTotalTxFP returns the TotalTxFP field if non-nil, zero value otherwise.

### GetTotalTxFPOk

`func (o *ClientStatisticalDataDetailResult) GetTotalTxFPOk() (*int64, bool)`

GetTotalTxFPOk returns a tuple with the TotalTxFP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxFP

`func (o *ClientStatisticalDataDetailResult) SetTotalTxFP(v int64)`

SetTotalTxFP sets TotalTxFP field to given value.

### HasTotalTxFP

`func (o *ClientStatisticalDataDetailResult) HasTotalTxFP() bool`

HasTotalTxFP returns a boolean if a field has been set.

### GetTotalUp

`func (o *ClientStatisticalDataDetailResult) GetTotalUp() int64`

GetTotalUp returns the TotalUp field if non-nil, zero value otherwise.

### GetTotalUpOk

`func (o *ClientStatisticalDataDetailResult) GetTotalUpOk() (*int64, bool)`

GetTotalUpOk returns a tuple with the TotalUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUp

`func (o *ClientStatisticalDataDetailResult) SetTotalUp(v int64)`

SetTotalUp sets TotalUp field to given value.

### HasTotalUp

`func (o *ClientStatisticalDataDetailResult) HasTotalUp() bool`

HasTotalUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


