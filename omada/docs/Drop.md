# Drop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropRate** | Pointer to **float64** | AP drop packet rate within one hour | [optional] 
**DropTimes** | Pointer to **int32** | Number of drop packets within one hour, such as 60 : 60% | [optional] 
**Time** | Pointer to **int64** | Sampling time second | [optional] 

## Methods

### NewDrop

`func NewDrop() *Drop`

NewDrop instantiates a new Drop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropWithDefaults

`func NewDropWithDefaults() *Drop`

NewDropWithDefaults instantiates a new Drop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropRate

`func (o *Drop) GetDropRate() float64`

GetDropRate returns the DropRate field if non-nil, zero value otherwise.

### GetDropRateOk

`func (o *Drop) GetDropRateOk() (*float64, bool)`

GetDropRateOk returns a tuple with the DropRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropRate

`func (o *Drop) SetDropRate(v float64)`

SetDropRate sets DropRate field to given value.

### HasDropRate

`func (o *Drop) HasDropRate() bool`

HasDropRate returns a boolean if a field has been set.

### GetDropTimes

`func (o *Drop) GetDropTimes() int32`

GetDropTimes returns the DropTimes field if non-nil, zero value otherwise.

### GetDropTimesOk

`func (o *Drop) GetDropTimesOk() (*int32, bool)`

GetDropTimesOk returns a tuple with the DropTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropTimes

`func (o *Drop) SetDropTimes(v int32)`

SetDropTimes sets DropTimes field to given value.

### HasDropTimes

`func (o *Drop) HasDropTimes() bool`

HasDropTimes returns a boolean if a field has been set.

### GetTime

`func (o *Drop) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Drop) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Drop) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *Drop) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


