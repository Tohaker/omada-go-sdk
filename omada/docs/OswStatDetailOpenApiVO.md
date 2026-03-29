# OswStatDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int32** | Utilization of cpu | [optional] 
**Mem** | Pointer to **int32** | Utilization of memory | [optional] 
**Ports** | Pointer to [**[]OswPortStatOpenApiVO**](OswPortStatOpenApiVO.md) | Traffic information of ports | [optional] 
**Time** | Pointer to **int64** | sampling moment | [optional] 

## Methods

### NewOswStatDetailOpenApiVO

`func NewOswStatDetailOpenApiVO() *OswStatDetailOpenApiVO`

NewOswStatDetailOpenApiVO instantiates a new OswStatDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStatDetailOpenApiVOWithDefaults

`func NewOswStatDetailOpenApiVOWithDefaults() *OswStatDetailOpenApiVO`

NewOswStatDetailOpenApiVOWithDefaults instantiates a new OswStatDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *OswStatDetailOpenApiVO) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OswStatDetailOpenApiVO) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OswStatDetailOpenApiVO) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *OswStatDetailOpenApiVO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMem

`func (o *OswStatDetailOpenApiVO) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *OswStatDetailOpenApiVO) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *OswStatDetailOpenApiVO) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *OswStatDetailOpenApiVO) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetPorts

`func (o *OswStatDetailOpenApiVO) GetPorts() []OswPortStatOpenApiVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswStatDetailOpenApiVO) GetPortsOk() (*[]OswPortStatOpenApiVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswStatDetailOpenApiVO) SetPorts(v []OswPortStatOpenApiVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswStatDetailOpenApiVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTime

`func (o *OswStatDetailOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OswStatDetailOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OswStatDetailOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OswStatDetailOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


