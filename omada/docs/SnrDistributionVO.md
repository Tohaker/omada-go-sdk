# SnrDistributionVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnrDistribution** | Pointer to **[]int32** | SNR distribution: [&gt;&#x3D;15, [10,15), &lt;10] | [optional] 
**Time** | Pointer to **int64** | Timestamp in seconds | [optional] 

## Methods

### NewSnrDistributionVO

`func NewSnrDistributionVO() *SnrDistributionVO`

NewSnrDistributionVO instantiates a new SnrDistributionVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnrDistributionVOWithDefaults

`func NewSnrDistributionVOWithDefaults() *SnrDistributionVO`

NewSnrDistributionVOWithDefaults instantiates a new SnrDistributionVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnrDistribution

`func (o *SnrDistributionVO) GetSnrDistribution() []int32`

GetSnrDistribution returns the SnrDistribution field if non-nil, zero value otherwise.

### GetSnrDistributionOk

`func (o *SnrDistributionVO) GetSnrDistributionOk() (*[]int32, bool)`

GetSnrDistributionOk returns a tuple with the SnrDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnrDistribution

`func (o *SnrDistributionVO) SetSnrDistribution(v []int32)`

SetSnrDistribution sets SnrDistribution field to given value.

### HasSnrDistribution

`func (o *SnrDistributionVO) HasSnrDistribution() bool`

HasSnrDistribution returns a boolean if a field has been set.

### GetTime

`func (o *SnrDistributionVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SnrDistributionVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SnrDistributionVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *SnrDistributionVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


