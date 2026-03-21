# OswHealthDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**ErrorFrame** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**Memory** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**PacketLoss** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**Score** | Pointer to **int32** | Device health score. | [optional] 
**Temperature** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 

## Methods

### NewOswHealthDetailVO

`func NewOswHealthDetailVO() *OswHealthDetailVO`

NewOswHealthDetailVO instantiates a new OswHealthDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswHealthDetailVOWithDefaults

`func NewOswHealthDetailVOWithDefaults() *OswHealthDetailVO`

NewOswHealthDetailVOWithDefaults instantiates a new OswHealthDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *OswHealthDetailVO) GetCpu() CommonSubHealthInfoDetailVO`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OswHealthDetailVO) GetCpuOk() (*CommonSubHealthInfoDetailVO, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OswHealthDetailVO) SetCpu(v CommonSubHealthInfoDetailVO)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *OswHealthDetailVO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetErrorFrame

`func (o *OswHealthDetailVO) GetErrorFrame() CommonSubHealthInfoDetailVO`

GetErrorFrame returns the ErrorFrame field if non-nil, zero value otherwise.

### GetErrorFrameOk

`func (o *OswHealthDetailVO) GetErrorFrameOk() (*CommonSubHealthInfoDetailVO, bool)`

GetErrorFrameOk returns a tuple with the ErrorFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFrame

`func (o *OswHealthDetailVO) SetErrorFrame(v CommonSubHealthInfoDetailVO)`

SetErrorFrame sets ErrorFrame field to given value.

### HasErrorFrame

`func (o *OswHealthDetailVO) HasErrorFrame() bool`

HasErrorFrame returns a boolean if a field has been set.

### GetMemory

`func (o *OswHealthDetailVO) GetMemory() CommonSubHealthInfoDetailVO`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *OswHealthDetailVO) GetMemoryOk() (*CommonSubHealthInfoDetailVO, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *OswHealthDetailVO) SetMemory(v CommonSubHealthInfoDetailVO)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *OswHealthDetailVO) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetPacketLoss

`func (o *OswHealthDetailVO) GetPacketLoss() CommonSubHealthInfoDetailVO`

GetPacketLoss returns the PacketLoss field if non-nil, zero value otherwise.

### GetPacketLossOk

`func (o *OswHealthDetailVO) GetPacketLossOk() (*CommonSubHealthInfoDetailVO, bool)`

GetPacketLossOk returns a tuple with the PacketLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLoss

`func (o *OswHealthDetailVO) SetPacketLoss(v CommonSubHealthInfoDetailVO)`

SetPacketLoss sets PacketLoss field to given value.

### HasPacketLoss

`func (o *OswHealthDetailVO) HasPacketLoss() bool`

HasPacketLoss returns a boolean if a field has been set.

### GetScore

`func (o *OswHealthDetailVO) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OswHealthDetailVO) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OswHealthDetailVO) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *OswHealthDetailVO) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTemperature

`func (o *OswHealthDetailVO) GetTemperature() CommonSubHealthInfoDetailVO`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *OswHealthDetailVO) GetTemperatureOk() (*CommonSubHealthInfoDetailVO, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *OswHealthDetailVO) SetTemperature(v CommonSubHealthInfoDetailVO)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *OswHealthDetailVO) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


