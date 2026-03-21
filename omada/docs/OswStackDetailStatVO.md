# OswStackDetailStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**[]OswStackDetailIntPropertyDTO**](OswStackDetailIntPropertyDTO.md) | Cpu Utilization information | [optional] 
**DropPkts** | Pointer to [**[]OswStackDetailLongPropertyDTO**](OswStackDetailLongPropertyDTO.md) | Drop Packets information | [optional] 
**Mem** | Pointer to [**[]OswStackDetailIntPropertyDTO**](OswStackDetailIntPropertyDTO.md) | Memory Utilization information | [optional] 
**RxErrPkts** | Pointer to [**[]OswStackDetailLongPropertyDTO**](OswStackDetailLongPropertyDTO.md) | Receive error packets information | [optional] 
**Time** | Pointer to **int64** | Time | [optional] 
**TxErrPkts** | Pointer to [**[]OswStackDetailLongPropertyDTO**](OswStackDetailLongPropertyDTO.md) | Transmit error packets information | [optional] 

## Methods

### NewOswStackDetailStatVO

`func NewOswStackDetailStatVO() *OswStackDetailStatVO`

NewOswStackDetailStatVO instantiates a new OswStackDetailStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackDetailStatVOWithDefaults

`func NewOswStackDetailStatVOWithDefaults() *OswStackDetailStatVO`

NewOswStackDetailStatVOWithDefaults instantiates a new OswStackDetailStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *OswStackDetailStatVO) GetCpu() []OswStackDetailIntPropertyDTO`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OswStackDetailStatVO) GetCpuOk() (*[]OswStackDetailIntPropertyDTO, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OswStackDetailStatVO) SetCpu(v []OswStackDetailIntPropertyDTO)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *OswStackDetailStatVO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDropPkts

`func (o *OswStackDetailStatVO) GetDropPkts() []OswStackDetailLongPropertyDTO`

GetDropPkts returns the DropPkts field if non-nil, zero value otherwise.

### GetDropPktsOk

`func (o *OswStackDetailStatVO) GetDropPktsOk() (*[]OswStackDetailLongPropertyDTO, bool)`

GetDropPktsOk returns a tuple with the DropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropPkts

`func (o *OswStackDetailStatVO) SetDropPkts(v []OswStackDetailLongPropertyDTO)`

SetDropPkts sets DropPkts field to given value.

### HasDropPkts

`func (o *OswStackDetailStatVO) HasDropPkts() bool`

HasDropPkts returns a boolean if a field has been set.

### GetMem

`func (o *OswStackDetailStatVO) GetMem() []OswStackDetailIntPropertyDTO`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *OswStackDetailStatVO) GetMemOk() (*[]OswStackDetailIntPropertyDTO, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *OswStackDetailStatVO) SetMem(v []OswStackDetailIntPropertyDTO)`

SetMem sets Mem field to given value.

### HasMem

`func (o *OswStackDetailStatVO) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *OswStackDetailStatVO) GetRxErrPkts() []OswStackDetailLongPropertyDTO`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *OswStackDetailStatVO) GetRxErrPktsOk() (*[]OswStackDetailLongPropertyDTO, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *OswStackDetailStatVO) SetRxErrPkts(v []OswStackDetailLongPropertyDTO)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *OswStackDetailStatVO) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetTime

`func (o *OswStackDetailStatVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OswStackDetailStatVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OswStackDetailStatVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OswStackDetailStatVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *OswStackDetailStatVO) GetTxErrPkts() []OswStackDetailLongPropertyDTO`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *OswStackDetailStatVO) GetTxErrPktsOk() (*[]OswStackDetailLongPropertyDTO, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *OswStackDetailStatVO) SetTxErrPkts(v []OswStackDetailLongPropertyDTO)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *OswStackDetailStatVO) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


