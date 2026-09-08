# OsgHealthDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelUtil** | Pointer to [**ChannelSubHealthInfoDetailVO**](ChannelSubHealthInfoDetailVO.md) |  | [optional] 
**Cpu** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**Incident** | Pointer to [**IncidentSubHealthInfoDetailVO**](IncidentSubHealthInfoDetailVO.md) |  | [optional] 
**LinkError** | Pointer to [**LinkErrorHealthInfoDetailVO**](LinkErrorHealthInfoDetailVO.md) |  | [optional] 
**Memory** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**Score** | Pointer to **int32** | Device health score. | [optional] 
**SupportedChannelUtil** | Pointer to [**ChannelSubHealthInfoDetailVO**](ChannelSubHealthInfoDetailVO.md) |  | [optional] 
**Temperature** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**Transmission** | Pointer to [**TransmissionSubHealthInfoDetailVO**](TransmissionSubHealthInfoDetailVO.md) |  | [optional] 
**WanLatency** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**WanStatus** | Pointer to [**WanSubHealthInfoDetailVO**](WanSubHealthInfoDetailVO.md) |  | [optional] 

## Methods

### NewOsgHealthDetailVO

`func NewOsgHealthDetailVO() *OsgHealthDetailVO`

NewOsgHealthDetailVO instantiates a new OsgHealthDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgHealthDetailVOWithDefaults

`func NewOsgHealthDetailVOWithDefaults() *OsgHealthDetailVO`

NewOsgHealthDetailVOWithDefaults instantiates a new OsgHealthDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelUtil

`func (o *OsgHealthDetailVO) GetChannelUtil() ChannelSubHealthInfoDetailVO`

GetChannelUtil returns the ChannelUtil field if non-nil, zero value otherwise.

### GetChannelUtilOk

`func (o *OsgHealthDetailVO) GetChannelUtilOk() (*ChannelSubHealthInfoDetailVO, bool)`

GetChannelUtilOk returns a tuple with the ChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtil

`func (o *OsgHealthDetailVO) SetChannelUtil(v ChannelSubHealthInfoDetailVO)`

SetChannelUtil sets ChannelUtil field to given value.

### HasChannelUtil

`func (o *OsgHealthDetailVO) HasChannelUtil() bool`

HasChannelUtil returns a boolean if a field has been set.

### GetCpu

`func (o *OsgHealthDetailVO) GetCpu() CommonSubHealthInfoDetailVO`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OsgHealthDetailVO) GetCpuOk() (*CommonSubHealthInfoDetailVO, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OsgHealthDetailVO) SetCpu(v CommonSubHealthInfoDetailVO)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *OsgHealthDetailVO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetIncident

`func (o *OsgHealthDetailVO) GetIncident() IncidentSubHealthInfoDetailVO`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *OsgHealthDetailVO) GetIncidentOk() (*IncidentSubHealthInfoDetailVO, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *OsgHealthDetailVO) SetIncident(v IncidentSubHealthInfoDetailVO)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *OsgHealthDetailVO) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetLinkError

`func (o *OsgHealthDetailVO) GetLinkError() LinkErrorHealthInfoDetailVO`

GetLinkError returns the LinkError field if non-nil, zero value otherwise.

### GetLinkErrorOk

`func (o *OsgHealthDetailVO) GetLinkErrorOk() (*LinkErrorHealthInfoDetailVO, bool)`

GetLinkErrorOk returns a tuple with the LinkError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkError

`func (o *OsgHealthDetailVO) SetLinkError(v LinkErrorHealthInfoDetailVO)`

SetLinkError sets LinkError field to given value.

### HasLinkError

`func (o *OsgHealthDetailVO) HasLinkError() bool`

HasLinkError returns a boolean if a field has been set.

### GetMemory

`func (o *OsgHealthDetailVO) GetMemory() CommonSubHealthInfoDetailVO`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *OsgHealthDetailVO) GetMemoryOk() (*CommonSubHealthInfoDetailVO, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *OsgHealthDetailVO) SetMemory(v CommonSubHealthInfoDetailVO)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *OsgHealthDetailVO) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetScore

`func (o *OsgHealthDetailVO) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OsgHealthDetailVO) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OsgHealthDetailVO) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *OsgHealthDetailVO) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSupportedChannelUtil

`func (o *OsgHealthDetailVO) GetSupportedChannelUtil() ChannelSubHealthInfoDetailVO`

GetSupportedChannelUtil returns the SupportedChannelUtil field if non-nil, zero value otherwise.

### GetSupportedChannelUtilOk

`func (o *OsgHealthDetailVO) GetSupportedChannelUtilOk() (*ChannelSubHealthInfoDetailVO, bool)`

GetSupportedChannelUtilOk returns a tuple with the SupportedChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedChannelUtil

`func (o *OsgHealthDetailVO) SetSupportedChannelUtil(v ChannelSubHealthInfoDetailVO)`

SetSupportedChannelUtil sets SupportedChannelUtil field to given value.

### HasSupportedChannelUtil

`func (o *OsgHealthDetailVO) HasSupportedChannelUtil() bool`

HasSupportedChannelUtil returns a boolean if a field has been set.

### GetTemperature

`func (o *OsgHealthDetailVO) GetTemperature() CommonSubHealthInfoDetailVO`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *OsgHealthDetailVO) GetTemperatureOk() (*CommonSubHealthInfoDetailVO, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *OsgHealthDetailVO) SetTemperature(v CommonSubHealthInfoDetailVO)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *OsgHealthDetailVO) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTransmission

`func (o *OsgHealthDetailVO) GetTransmission() TransmissionSubHealthInfoDetailVO`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *OsgHealthDetailVO) GetTransmissionOk() (*TransmissionSubHealthInfoDetailVO, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *OsgHealthDetailVO) SetTransmission(v TransmissionSubHealthInfoDetailVO)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *OsgHealthDetailVO) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetWanLatency

`func (o *OsgHealthDetailVO) GetWanLatency() CommonSubHealthInfoDetailVO`

GetWanLatency returns the WanLatency field if non-nil, zero value otherwise.

### GetWanLatencyOk

`func (o *OsgHealthDetailVO) GetWanLatencyOk() (*CommonSubHealthInfoDetailVO, bool)`

GetWanLatencyOk returns a tuple with the WanLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanLatency

`func (o *OsgHealthDetailVO) SetWanLatency(v CommonSubHealthInfoDetailVO)`

SetWanLatency sets WanLatency field to given value.

### HasWanLatency

`func (o *OsgHealthDetailVO) HasWanLatency() bool`

HasWanLatency returns a boolean if a field has been set.

### GetWanStatus

`func (o *OsgHealthDetailVO) GetWanStatus() WanSubHealthInfoDetailVO`

GetWanStatus returns the WanStatus field if non-nil, zero value otherwise.

### GetWanStatusOk

`func (o *OsgHealthDetailVO) GetWanStatusOk() (*WanSubHealthInfoDetailVO, bool)`

GetWanStatusOk returns a tuple with the WanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanStatus

`func (o *OsgHealthDetailVO) SetWanStatus(v WanSubHealthInfoDetailVO)`

SetWanStatus sets WanStatus field to given value.

### HasWanStatus

`func (o *OsgHealthDetailVO) HasWanStatus() bool`

HasWanStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


