# RssiDistributionVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RssiDistribution** | Pointer to **[]int32** | RSSI distribution: [&gt;&#x3D;−65dBm, [−72,−65), &lt;−72dBm] | [optional] 
**Time** | Pointer to **int64** | Timestamp in seconds | [optional] 

## Methods

### NewRssiDistributionVO

`func NewRssiDistributionVO() *RssiDistributionVO`

NewRssiDistributionVO instantiates a new RssiDistributionVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRssiDistributionVOWithDefaults

`func NewRssiDistributionVOWithDefaults() *RssiDistributionVO`

NewRssiDistributionVOWithDefaults instantiates a new RssiDistributionVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRssiDistribution

`func (o *RssiDistributionVO) GetRssiDistribution() []int32`

GetRssiDistribution returns the RssiDistribution field if non-nil, zero value otherwise.

### GetRssiDistributionOk

`func (o *RssiDistributionVO) GetRssiDistributionOk() (*[]int32, bool)`

GetRssiDistributionOk returns a tuple with the RssiDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiDistribution

`func (o *RssiDistributionVO) SetRssiDistribution(v []int32)`

SetRssiDistribution sets RssiDistribution field to given value.

### HasRssiDistribution

`func (o *RssiDistributionVO) HasRssiDistribution() bool`

HasRssiDistribution returns a boolean if a field has been set.

### GetTime

`func (o *RssiDistributionVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RssiDistributionVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RssiDistributionVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *RssiDistributionVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


