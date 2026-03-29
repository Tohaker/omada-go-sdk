# StatQueryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attrs** | **[]string** | Attributes to be queried. Attributes not included in attrs will return a value of 0.Item of attrs should be a value as follows: mem, cpu, tx, rx, txRate, rxRate, txPkts, rxPkts, txBroadPkts, rxBroadPkts, txMultiPkts, rxMultiPkts, dropPkts, txErrPkts, rxErrPkts | 
**End** | **int64** | end time, number of seconds from UTC0 1970/01/01 | 
**OltPorts** | Pointer to **[]string** | Statistics of the selected ports of olt would be queried. | [optional] 
**Ports** | Pointer to **[]int32** | The ports in the lag. Each item is Integer, for example: [1, 2]. | [optional] 
**StandardPorts** | Pointer to **[]string** | Statistics of the selected ports would be queried , the param is valid when stacking | [optional] 
**Start** | **int64** | Start time, number of seconds from UTC0 1970/01/01 | 

## Methods

### NewStatQueryVO

`func NewStatQueryVO(attrs []string, end int64, start int64, ) *StatQueryVO`

NewStatQueryVO instantiates a new StatQueryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatQueryVOWithDefaults

`func NewStatQueryVOWithDefaults() *StatQueryVO`

NewStatQueryVOWithDefaults instantiates a new StatQueryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrs

`func (o *StatQueryVO) GetAttrs() []string`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *StatQueryVO) GetAttrsOk() (*[]string, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *StatQueryVO) SetAttrs(v []string)`

SetAttrs sets Attrs field to given value.


### GetEnd

`func (o *StatQueryVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *StatQueryVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *StatQueryVO) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetOltPorts

`func (o *StatQueryVO) GetOltPorts() []string`

GetOltPorts returns the OltPorts field if non-nil, zero value otherwise.

### GetOltPortsOk

`func (o *StatQueryVO) GetOltPortsOk() (*[]string, bool)`

GetOltPortsOk returns a tuple with the OltPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOltPorts

`func (o *StatQueryVO) SetOltPorts(v []string)`

SetOltPorts sets OltPorts field to given value.

### HasOltPorts

`func (o *StatQueryVO) HasOltPorts() bool`

HasOltPorts returns a boolean if a field has been set.

### GetPorts

`func (o *StatQueryVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *StatQueryVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *StatQueryVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *StatQueryVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStandardPorts

`func (o *StatQueryVO) GetStandardPorts() []string`

GetStandardPorts returns the StandardPorts field if non-nil, zero value otherwise.

### GetStandardPortsOk

`func (o *StatQueryVO) GetStandardPortsOk() (*[]string, bool)`

GetStandardPortsOk returns a tuple with the StandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPorts

`func (o *StatQueryVO) SetStandardPorts(v []string)`

SetStandardPorts sets StandardPorts field to given value.

### HasStandardPorts

`func (o *StatQueryVO) HasStandardPorts() bool`

HasStandardPorts returns a boolean if a field has been set.

### GetStart

`func (o *StatQueryVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *StatQueryVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *StatQueryVO) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


