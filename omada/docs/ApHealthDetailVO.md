# ApHealthDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelInterf** | Pointer to [**ChannelSubHealthInfoDetailVO**](ChannelSubHealthInfoDetailVO.md) |  | [optional] 
**ChannelUtil** | Pointer to [**ChannelSubHealthInfoDetailVO**](ChannelSubHealthInfoDetailVO.md) |  | [optional] 
**Cpu** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**Memory** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**Score** | Pointer to **int32** | Device health score. | [optional] 
**SupportedChannelInterf** | Pointer to [**ChannelSubHealthInfoDetailVO**](ChannelSubHealthInfoDetailVO.md) |  | [optional] 
**SupportedChannelUtil** | Pointer to [**ChannelSubHealthInfoDetailVO**](ChannelSubHealthInfoDetailVO.md) |  | [optional] 
**Transmission** | Pointer to [**TransmissionSubHealthInfoDetailVO**](TransmissionSubHealthInfoDetailVO.md) |  | [optional] 

## Methods

### NewApHealthDetailVO

`func NewApHealthDetailVO() *ApHealthDetailVO`

NewApHealthDetailVO instantiates a new ApHealthDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApHealthDetailVOWithDefaults

`func NewApHealthDetailVOWithDefaults() *ApHealthDetailVO`

NewApHealthDetailVOWithDefaults instantiates a new ApHealthDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelInterf

`func (o *ApHealthDetailVO) GetChannelInterf() ChannelSubHealthInfoDetailVO`

GetChannelInterf returns the ChannelInterf field if non-nil, zero value otherwise.

### GetChannelInterfOk

`func (o *ApHealthDetailVO) GetChannelInterfOk() (*ChannelSubHealthInfoDetailVO, bool)`

GetChannelInterfOk returns a tuple with the ChannelInterf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelInterf

`func (o *ApHealthDetailVO) SetChannelInterf(v ChannelSubHealthInfoDetailVO)`

SetChannelInterf sets ChannelInterf field to given value.

### HasChannelInterf

`func (o *ApHealthDetailVO) HasChannelInterf() bool`

HasChannelInterf returns a boolean if a field has been set.

### GetChannelUtil

`func (o *ApHealthDetailVO) GetChannelUtil() ChannelSubHealthInfoDetailVO`

GetChannelUtil returns the ChannelUtil field if non-nil, zero value otherwise.

### GetChannelUtilOk

`func (o *ApHealthDetailVO) GetChannelUtilOk() (*ChannelSubHealthInfoDetailVO, bool)`

GetChannelUtilOk returns a tuple with the ChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtil

`func (o *ApHealthDetailVO) SetChannelUtil(v ChannelSubHealthInfoDetailVO)`

SetChannelUtil sets ChannelUtil field to given value.

### HasChannelUtil

`func (o *ApHealthDetailVO) HasChannelUtil() bool`

HasChannelUtil returns a boolean if a field has been set.

### GetCpu

`func (o *ApHealthDetailVO) GetCpu() CommonSubHealthInfoDetailVO`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApHealthDetailVO) GetCpuOk() (*CommonSubHealthInfoDetailVO, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApHealthDetailVO) SetCpu(v CommonSubHealthInfoDetailVO)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ApHealthDetailVO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ApHealthDetailVO) GetMemory() CommonSubHealthInfoDetailVO`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ApHealthDetailVO) GetMemoryOk() (*CommonSubHealthInfoDetailVO, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ApHealthDetailVO) SetMemory(v CommonSubHealthInfoDetailVO)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ApHealthDetailVO) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetScore

`func (o *ApHealthDetailVO) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ApHealthDetailVO) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ApHealthDetailVO) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ApHealthDetailVO) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSupportedChannelInterf

`func (o *ApHealthDetailVO) GetSupportedChannelInterf() ChannelSubHealthInfoDetailVO`

GetSupportedChannelInterf returns the SupportedChannelInterf field if non-nil, zero value otherwise.

### GetSupportedChannelInterfOk

`func (o *ApHealthDetailVO) GetSupportedChannelInterfOk() (*ChannelSubHealthInfoDetailVO, bool)`

GetSupportedChannelInterfOk returns a tuple with the SupportedChannelInterf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedChannelInterf

`func (o *ApHealthDetailVO) SetSupportedChannelInterf(v ChannelSubHealthInfoDetailVO)`

SetSupportedChannelInterf sets SupportedChannelInterf field to given value.

### HasSupportedChannelInterf

`func (o *ApHealthDetailVO) HasSupportedChannelInterf() bool`

HasSupportedChannelInterf returns a boolean if a field has been set.

### GetSupportedChannelUtil

`func (o *ApHealthDetailVO) GetSupportedChannelUtil() ChannelSubHealthInfoDetailVO`

GetSupportedChannelUtil returns the SupportedChannelUtil field if non-nil, zero value otherwise.

### GetSupportedChannelUtilOk

`func (o *ApHealthDetailVO) GetSupportedChannelUtilOk() (*ChannelSubHealthInfoDetailVO, bool)`

GetSupportedChannelUtilOk returns a tuple with the SupportedChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedChannelUtil

`func (o *ApHealthDetailVO) SetSupportedChannelUtil(v ChannelSubHealthInfoDetailVO)`

SetSupportedChannelUtil sets SupportedChannelUtil field to given value.

### HasSupportedChannelUtil

`func (o *ApHealthDetailVO) HasSupportedChannelUtil() bool`

HasSupportedChannelUtil returns a boolean if a field has been set.

### GetTransmission

`func (o *ApHealthDetailVO) GetTransmission() TransmissionSubHealthInfoDetailVO`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *ApHealthDetailVO) GetTransmissionOk() (*TransmissionSubHealthInfoDetailVO, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *ApHealthDetailVO) SetTransmission(v TransmissionSubHealthInfoDetailVO)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *ApHealthDetailVO) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


