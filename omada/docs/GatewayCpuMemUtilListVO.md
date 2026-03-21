# GatewayCpuMemUtilListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **int32** | Device CPU utilization rate | [optional] 
**MemUtil** | Pointer to **int32** | Device memory utilization rate | [optional] 
**Time** | Pointer to **int64** | Timestamp, in seconds, such as 1682000000 | [optional] 

## Methods

### NewGatewayCpuMemUtilListVO

`func NewGatewayCpuMemUtilListVO() *GatewayCpuMemUtilListVO`

NewGatewayCpuMemUtilListVO instantiates a new GatewayCpuMemUtilListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCpuMemUtilListVOWithDefaults

`func NewGatewayCpuMemUtilListVOWithDefaults() *GatewayCpuMemUtilListVO`

NewGatewayCpuMemUtilListVOWithDefaults instantiates a new GatewayCpuMemUtilListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *GatewayCpuMemUtilListVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *GatewayCpuMemUtilListVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *GatewayCpuMemUtilListVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *GatewayCpuMemUtilListVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetMemUtil

`func (o *GatewayCpuMemUtilListVO) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *GatewayCpuMemUtilListVO) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *GatewayCpuMemUtilListVO) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *GatewayCpuMemUtilListVO) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetTime

`func (o *GatewayCpuMemUtilListVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GatewayCpuMemUtilListVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GatewayCpuMemUtilListVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GatewayCpuMemUtilListVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


