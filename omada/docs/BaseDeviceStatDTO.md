# BaseDeviceStatDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int32** | Utilization of cpu | [optional] 
**Mac** | Pointer to **string** | the mac of the device | [optional] 
**Mem** | Pointer to **int32** | Utilization of memory | [optional] 
**RebootTimes** | Pointer to **int32** | reboot times | [optional] 
**Time** | Pointer to **int64** | sampling moment | [optional] 

## Methods

### NewBaseDeviceStatDTO

`func NewBaseDeviceStatDTO() *BaseDeviceStatDTO`

NewBaseDeviceStatDTO instantiates a new BaseDeviceStatDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDeviceStatDTOWithDefaults

`func NewBaseDeviceStatDTOWithDefaults() *BaseDeviceStatDTO`

NewBaseDeviceStatDTOWithDefaults instantiates a new BaseDeviceStatDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *BaseDeviceStatDTO) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *BaseDeviceStatDTO) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *BaseDeviceStatDTO) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *BaseDeviceStatDTO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMac

`func (o *BaseDeviceStatDTO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *BaseDeviceStatDTO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *BaseDeviceStatDTO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *BaseDeviceStatDTO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMem

`func (o *BaseDeviceStatDTO) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *BaseDeviceStatDTO) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *BaseDeviceStatDTO) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *BaseDeviceStatDTO) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetRebootTimes

`func (o *BaseDeviceStatDTO) GetRebootTimes() int32`

GetRebootTimes returns the RebootTimes field if non-nil, zero value otherwise.

### GetRebootTimesOk

`func (o *BaseDeviceStatDTO) GetRebootTimesOk() (*int32, bool)`

GetRebootTimesOk returns a tuple with the RebootTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootTimes

`func (o *BaseDeviceStatDTO) SetRebootTimes(v int32)`

SetRebootTimes sets RebootTimes field to given value.

### HasRebootTimes

`func (o *BaseDeviceStatDTO) HasRebootTimes() bool`

HasRebootTimes returns a boolean if a field has been set.

### GetTime

`func (o *BaseDeviceStatDTO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *BaseDeviceStatDTO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *BaseDeviceStatDTO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *BaseDeviceStatDTO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


