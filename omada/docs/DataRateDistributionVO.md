# DataRateDistributionVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRateAboveWifi4** | Pointer to **[]int32** | 802.11n/ac/ax speed distribution: [&gt;&#x3D;72, [12,72), &lt;12] Mbps | [optional] 
**DataRateBelowWifi4** | Pointer to **[]int32** | 802.11a/b/g speed distribution: [&gt;&#x3D;21, [12,21), &lt;12] Mbps | [optional] 
**Time** | Pointer to **int64** | Timestamp in seconds | [optional] 

## Methods

### NewDataRateDistributionVO

`func NewDataRateDistributionVO() *DataRateDistributionVO`

NewDataRateDistributionVO instantiates a new DataRateDistributionVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataRateDistributionVOWithDefaults

`func NewDataRateDistributionVOWithDefaults() *DataRateDistributionVO`

NewDataRateDistributionVOWithDefaults instantiates a new DataRateDistributionVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRateAboveWifi4

`func (o *DataRateDistributionVO) GetDataRateAboveWifi4() []int32`

GetDataRateAboveWifi4 returns the DataRateAboveWifi4 field if non-nil, zero value otherwise.

### GetDataRateAboveWifi4Ok

`func (o *DataRateDistributionVO) GetDataRateAboveWifi4Ok() (*[]int32, bool)`

GetDataRateAboveWifi4Ok returns a tuple with the DataRateAboveWifi4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRateAboveWifi4

`func (o *DataRateDistributionVO) SetDataRateAboveWifi4(v []int32)`

SetDataRateAboveWifi4 sets DataRateAboveWifi4 field to given value.

### HasDataRateAboveWifi4

`func (o *DataRateDistributionVO) HasDataRateAboveWifi4() bool`

HasDataRateAboveWifi4 returns a boolean if a field has been set.

### GetDataRateBelowWifi4

`func (o *DataRateDistributionVO) GetDataRateBelowWifi4() []int32`

GetDataRateBelowWifi4 returns the DataRateBelowWifi4 field if non-nil, zero value otherwise.

### GetDataRateBelowWifi4Ok

`func (o *DataRateDistributionVO) GetDataRateBelowWifi4Ok() (*[]int32, bool)`

GetDataRateBelowWifi4Ok returns a tuple with the DataRateBelowWifi4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRateBelowWifi4

`func (o *DataRateDistributionVO) SetDataRateBelowWifi4(v []int32)`

SetDataRateBelowWifi4 sets DataRateBelowWifi4 field to given value.

### HasDataRateBelowWifi4

`func (o *DataRateDistributionVO) HasDataRateBelowWifi4() bool`

HasDataRateBelowWifi4 returns a boolean if a field has been set.

### GetTime

`func (o *DataRateDistributionVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DataRateDistributionVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DataRateDistributionVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *DataRateDistributionVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


