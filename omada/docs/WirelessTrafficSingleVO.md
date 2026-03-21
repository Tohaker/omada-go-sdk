# WirelessTrafficSingleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RxTraffic** | Pointer to **int64** | rx traffic | [optional] 
**TotalTraffic** | Pointer to **int64** | total traffic | [optional] 
**TrafficTrend** | Pointer to [**[]WirelessTrafficTrendVO**](WirelessTrafficTrendVO.md) |  | [optional] 
**TxTraffic** | Pointer to **int64** | tx traffic | [optional] 

## Methods

### NewWirelessTrafficSingleVO

`func NewWirelessTrafficSingleVO() *WirelessTrafficSingleVO`

NewWirelessTrafficSingleVO instantiates a new WirelessTrafficSingleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessTrafficSingleVOWithDefaults

`func NewWirelessTrafficSingleVOWithDefaults() *WirelessTrafficSingleVO`

NewWirelessTrafficSingleVOWithDefaults instantiates a new WirelessTrafficSingleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRxTraffic

`func (o *WirelessTrafficSingleVO) GetRxTraffic() int64`

GetRxTraffic returns the RxTraffic field if non-nil, zero value otherwise.

### GetRxTrafficOk

`func (o *WirelessTrafficSingleVO) GetRxTrafficOk() (*int64, bool)`

GetRxTrafficOk returns a tuple with the RxTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxTraffic

`func (o *WirelessTrafficSingleVO) SetRxTraffic(v int64)`

SetRxTraffic sets RxTraffic field to given value.

### HasRxTraffic

`func (o *WirelessTrafficSingleVO) HasRxTraffic() bool`

HasRxTraffic returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *WirelessTrafficSingleVO) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *WirelessTrafficSingleVO) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *WirelessTrafficSingleVO) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *WirelessTrafficSingleVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetTrafficTrend

`func (o *WirelessTrafficSingleVO) GetTrafficTrend() []WirelessTrafficTrendVO`

GetTrafficTrend returns the TrafficTrend field if non-nil, zero value otherwise.

### GetTrafficTrendOk

`func (o *WirelessTrafficSingleVO) GetTrafficTrendOk() (*[]WirelessTrafficTrendVO, bool)`

GetTrafficTrendOk returns a tuple with the TrafficTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficTrend

`func (o *WirelessTrafficSingleVO) SetTrafficTrend(v []WirelessTrafficTrendVO)`

SetTrafficTrend sets TrafficTrend field to given value.

### HasTrafficTrend

`func (o *WirelessTrafficSingleVO) HasTrafficTrend() bool`

HasTrafficTrend returns a boolean if a field has been set.

### GetTxTraffic

`func (o *WirelessTrafficSingleVO) GetTxTraffic() int64`

GetTxTraffic returns the TxTraffic field if non-nil, zero value otherwise.

### GetTxTrafficOk

`func (o *WirelessTrafficSingleVO) GetTxTrafficOk() (*int64, bool)`

GetTxTrafficOk returns a tuple with the TxTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxTraffic

`func (o *WirelessTrafficSingleVO) SetTxTraffic(v int64)`

SetTxTraffic sets TxTraffic field to given value.

### HasTxTraffic

`func (o *WirelessTrafficSingleVO) HasTxTraffic() bool`

HasTxTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


