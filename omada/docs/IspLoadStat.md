# IspLoadStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latency** | Pointer to **int32** | WAN port delay should be within the range of 0 or 1000. (in ms) | [optional] 
**Time** | Pointer to **int64** | Timestamp of statistical data (in seconds) | [optional] 
**TotalRate** | Pointer to **int64** | WAN port rate (rxR+txR) | [optional] 

## Methods

### NewIspLoadStat

`func NewIspLoadStat() *IspLoadStat`

NewIspLoadStat instantiates a new IspLoadStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspLoadStatWithDefaults

`func NewIspLoadStatWithDefaults() *IspLoadStat`

NewIspLoadStatWithDefaults instantiates a new IspLoadStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatency

`func (o *IspLoadStat) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *IspLoadStat) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *IspLoadStat) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *IspLoadStat) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetTime

`func (o *IspLoadStat) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IspLoadStat) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IspLoadStat) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *IspLoadStat) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotalRate

`func (o *IspLoadStat) GetTotalRate() int64`

GetTotalRate returns the TotalRate field if non-nil, zero value otherwise.

### GetTotalRateOk

`func (o *IspLoadStat) GetTotalRateOk() (*int64, bool)`

GetTotalRateOk returns a tuple with the TotalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRate

`func (o *IspLoadStat) SetTotalRate(v int64)`

SetTotalRate sets TotalRate field to given value.

### HasTotalRate

`func (o *IspLoadStat) HasTotalRate() bool`

HasTotalRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


