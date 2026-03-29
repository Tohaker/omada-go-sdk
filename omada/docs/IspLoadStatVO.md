# IspLoadStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownTraffic** | Pointer to **int64** | Downstream traffic | [optional] 
**Latency** | Pointer to **int32** | WAN latency, when mode is WAN and device is connected, Unit: ms | [optional] 
**Time** | Pointer to **int64** | Timestamp, in seconds, such as 1682000000 | [optional] 
**TotalRate** | Pointer to **int64** | WAN port rate (rxR+txR) | [optional] 
**TotalTraffic** | Pointer to **int64** | Upstream and downstream traffic | [optional] 
**UpTraffic** | Pointer to **int64** | Upstream traffic | [optional] 

## Methods

### NewIspLoadStatVO

`func NewIspLoadStatVO() *IspLoadStatVO`

NewIspLoadStatVO instantiates a new IspLoadStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspLoadStatVOWithDefaults

`func NewIspLoadStatVOWithDefaults() *IspLoadStatVO`

NewIspLoadStatVOWithDefaults instantiates a new IspLoadStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownTraffic

`func (o *IspLoadStatVO) GetDownTraffic() int64`

GetDownTraffic returns the DownTraffic field if non-nil, zero value otherwise.

### GetDownTrafficOk

`func (o *IspLoadStatVO) GetDownTrafficOk() (*int64, bool)`

GetDownTrafficOk returns a tuple with the DownTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownTraffic

`func (o *IspLoadStatVO) SetDownTraffic(v int64)`

SetDownTraffic sets DownTraffic field to given value.

### HasDownTraffic

`func (o *IspLoadStatVO) HasDownTraffic() bool`

HasDownTraffic returns a boolean if a field has been set.

### GetLatency

`func (o *IspLoadStatVO) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *IspLoadStatVO) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *IspLoadStatVO) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *IspLoadStatVO) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetTime

`func (o *IspLoadStatVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IspLoadStatVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IspLoadStatVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *IspLoadStatVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotalRate

`func (o *IspLoadStatVO) GetTotalRate() int64`

GetTotalRate returns the TotalRate field if non-nil, zero value otherwise.

### GetTotalRateOk

`func (o *IspLoadStatVO) GetTotalRateOk() (*int64, bool)`

GetTotalRateOk returns a tuple with the TotalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRate

`func (o *IspLoadStatVO) SetTotalRate(v int64)`

SetTotalRate sets TotalRate field to given value.

### HasTotalRate

`func (o *IspLoadStatVO) HasTotalRate() bool`

HasTotalRate returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *IspLoadStatVO) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *IspLoadStatVO) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *IspLoadStatVO) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *IspLoadStatVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetUpTraffic

`func (o *IspLoadStatVO) GetUpTraffic() int64`

GetUpTraffic returns the UpTraffic field if non-nil, zero value otherwise.

### GetUpTrafficOk

`func (o *IspLoadStatVO) GetUpTrafficOk() (*int64, bool)`

GetUpTrafficOk returns a tuple with the UpTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTraffic

`func (o *IspLoadStatVO) SetUpTraffic(v int64)`

SetUpTraffic sets UpTraffic field to given value.

### HasUpTraffic

`func (o *IspLoadStatVO) HasUpTraffic() bool`

HasUpTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


