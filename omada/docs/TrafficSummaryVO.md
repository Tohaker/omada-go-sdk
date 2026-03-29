# TrafficSummaryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RxTraffic** | Pointer to **int64** | rx traffic | [optional] 
**TotalTraffic** | Pointer to **int64** | total traffic | [optional] 
**TrafficSummary** | Pointer to [**[]TrafficSummaryListVO**](TrafficSummaryListVO.md) | traffic data point | [optional] 
**TxTraffic** | Pointer to **int64** | tx traffic | [optional] 

## Methods

### NewTrafficSummaryVO

`func NewTrafficSummaryVO() *TrafficSummaryVO`

NewTrafficSummaryVO instantiates a new TrafficSummaryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficSummaryVOWithDefaults

`func NewTrafficSummaryVOWithDefaults() *TrafficSummaryVO`

NewTrafficSummaryVOWithDefaults instantiates a new TrafficSummaryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRxTraffic

`func (o *TrafficSummaryVO) GetRxTraffic() int64`

GetRxTraffic returns the RxTraffic field if non-nil, zero value otherwise.

### GetRxTrafficOk

`func (o *TrafficSummaryVO) GetRxTrafficOk() (*int64, bool)`

GetRxTrafficOk returns a tuple with the RxTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxTraffic

`func (o *TrafficSummaryVO) SetRxTraffic(v int64)`

SetRxTraffic sets RxTraffic field to given value.

### HasRxTraffic

`func (o *TrafficSummaryVO) HasRxTraffic() bool`

HasRxTraffic returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *TrafficSummaryVO) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *TrafficSummaryVO) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *TrafficSummaryVO) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *TrafficSummaryVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetTrafficSummary

`func (o *TrafficSummaryVO) GetTrafficSummary() []TrafficSummaryListVO`

GetTrafficSummary returns the TrafficSummary field if non-nil, zero value otherwise.

### GetTrafficSummaryOk

`func (o *TrafficSummaryVO) GetTrafficSummaryOk() (*[]TrafficSummaryListVO, bool)`

GetTrafficSummaryOk returns a tuple with the TrafficSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSummary

`func (o *TrafficSummaryVO) SetTrafficSummary(v []TrafficSummaryListVO)`

SetTrafficSummary sets TrafficSummary field to given value.

### HasTrafficSummary

`func (o *TrafficSummaryVO) HasTrafficSummary() bool`

HasTrafficSummary returns a boolean if a field has been set.

### GetTxTraffic

`func (o *TrafficSummaryVO) GetTxTraffic() int64`

GetTxTraffic returns the TxTraffic field if non-nil, zero value otherwise.

### GetTxTrafficOk

`func (o *TrafficSummaryVO) GetTxTrafficOk() (*int64, bool)`

GetTxTrafficOk returns a tuple with the TxTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxTraffic

`func (o *TrafficSummaryVO) SetTxTraffic(v int64)`

SetTxTraffic sets TxTraffic field to given value.

### HasTxTraffic

`func (o *TrafficSummaryVO) HasTxTraffic() bool`

HasTxTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


