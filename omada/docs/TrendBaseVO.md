# TrendBaseVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of trend | [optional] 
**Time** | Pointer to **int64** | timestamp | [optional] 

## Methods

### NewTrendBaseVO

`func NewTrendBaseVO() *TrendBaseVO`

NewTrendBaseVO instantiates a new TrendBaseVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendBaseVOWithDefaults

`func NewTrendBaseVOWithDefaults() *TrendBaseVO`

NewTrendBaseVOWithDefaults instantiates a new TrendBaseVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TrendBaseVO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TrendBaseVO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TrendBaseVO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TrendBaseVO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTime

`func (o *TrendBaseVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TrendBaseVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TrendBaseVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *TrendBaseVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


