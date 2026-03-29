# OswStackStatQueryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attrs** | **[]string** | Attributes to be queried. Attributes not included in attrs will return a value of 0.Item of attrs should be a value as follows: mem, cpu, tx, rx, txRate, rxRate, txPkts, rxPkts, txBroadPkts, rxBroadPkts, txMultiPkts, rxMultiPkts, dropPkts, txErrPkts, rxErrPkts | 
**End** | **int64** | End time, number of seconds from UTC0 1970/01/01 | 
**Interval** | **int32** | Interval should be a value as follows: 0:5min, 1:hourly, 2:daily | 
**Start** | **int64** | Start time, number of seconds from UTC0 1970/01/01 | 

## Methods

### NewOswStackStatQueryVO

`func NewOswStackStatQueryVO(attrs []string, end int64, interval int32, start int64, ) *OswStackStatQueryVO`

NewOswStackStatQueryVO instantiates a new OswStackStatQueryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackStatQueryVOWithDefaults

`func NewOswStackStatQueryVOWithDefaults() *OswStackStatQueryVO`

NewOswStackStatQueryVOWithDefaults instantiates a new OswStackStatQueryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrs

`func (o *OswStackStatQueryVO) GetAttrs() []string`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *OswStackStatQueryVO) GetAttrsOk() (*[]string, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *OswStackStatQueryVO) SetAttrs(v []string)`

SetAttrs sets Attrs field to given value.


### GetEnd

`func (o *OswStackStatQueryVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OswStackStatQueryVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OswStackStatQueryVO) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetInterval

`func (o *OswStackStatQueryVO) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *OswStackStatQueryVO) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *OswStackStatQueryVO) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetStart

`func (o *OswStackStatQueryVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OswStackStatQueryVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OswStackStatQueryVO) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


