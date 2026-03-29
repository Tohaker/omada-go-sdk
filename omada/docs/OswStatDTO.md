# OswStatDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int32** | Utilization of cpu | [optional] 
**Mac** | Pointer to **string** | the mac of the device | [optional] 
**Mem** | Pointer to **int32** | Utilization of memory | [optional] 
**PoePower** | Pointer to **int32** |  | [optional] 
**PoeUtil** | Pointer to **int32** |  | [optional] 
**Ports** | Pointer to [**[]OswPortStatDTO**](OswPortStatDTO.md) |  | [optional] 
**RebootTimes** | Pointer to **int32** | reboot times | [optional] 
**Time** | Pointer to **int64** | sampling moment | [optional] 

## Methods

### NewOswStatDTO

`func NewOswStatDTO() *OswStatDTO`

NewOswStatDTO instantiates a new OswStatDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStatDTOWithDefaults

`func NewOswStatDTOWithDefaults() *OswStatDTO`

NewOswStatDTOWithDefaults instantiates a new OswStatDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *OswStatDTO) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OswStatDTO) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OswStatDTO) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *OswStatDTO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMac

`func (o *OswStatDTO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStatDTO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStatDTO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStatDTO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMem

`func (o *OswStatDTO) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *OswStatDTO) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *OswStatDTO) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *OswStatDTO) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetPoePower

`func (o *OswStatDTO) GetPoePower() int32`

GetPoePower returns the PoePower field if non-nil, zero value otherwise.

### GetPoePowerOk

`func (o *OswStatDTO) GetPoePowerOk() (*int32, bool)`

GetPoePowerOk returns a tuple with the PoePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePower

`func (o *OswStatDTO) SetPoePower(v int32)`

SetPoePower sets PoePower field to given value.

### HasPoePower

`func (o *OswStatDTO) HasPoePower() bool`

HasPoePower returns a boolean if a field has been set.

### GetPoeUtil

`func (o *OswStatDTO) GetPoeUtil() int32`

GetPoeUtil returns the PoeUtil field if non-nil, zero value otherwise.

### GetPoeUtilOk

`func (o *OswStatDTO) GetPoeUtilOk() (*int32, bool)`

GetPoeUtilOk returns a tuple with the PoeUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeUtil

`func (o *OswStatDTO) SetPoeUtil(v int32)`

SetPoeUtil sets PoeUtil field to given value.

### HasPoeUtil

`func (o *OswStatDTO) HasPoeUtil() bool`

HasPoeUtil returns a boolean if a field has been set.

### GetPorts

`func (o *OswStatDTO) GetPorts() []OswPortStatDTO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswStatDTO) GetPortsOk() (*[]OswPortStatDTO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswStatDTO) SetPorts(v []OswPortStatDTO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswStatDTO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetRebootTimes

`func (o *OswStatDTO) GetRebootTimes() int32`

GetRebootTimes returns the RebootTimes field if non-nil, zero value otherwise.

### GetRebootTimesOk

`func (o *OswStatDTO) GetRebootTimesOk() (*int32, bool)`

GetRebootTimesOk returns a tuple with the RebootTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootTimes

`func (o *OswStatDTO) SetRebootTimes(v int32)`

SetRebootTimes sets RebootTimes field to given value.

### HasRebootTimes

`func (o *OswStatDTO) HasRebootTimes() bool`

HasRebootTimes returns a boolean if a field has been set.

### GetTime

`func (o *OswStatDTO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OswStatDTO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OswStatDTO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OswStatDTO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


