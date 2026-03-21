# OltStatDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int32** | Utilization of cpu | [optional] 
**Down** | Pointer to **int64** | Downlink traffic | [optional] 
**Mem** | Pointer to **int32** | Utilization of memory | [optional] 
**OnuCount** | Pointer to **int32** | Number of ONU | [optional] 
**Ports** | Pointer to [**[]OltPortStatOpenApiVO**](OltPortStatOpenApiVO.md) | Traffic information of ports | [optional] 
**Time** | Pointer to **int64** | sampling moment | [optional] 
**Up** | Pointer to **int64** | Uplink traffic | [optional] 

## Methods

### NewOltStatDetailOpenApiVO

`func NewOltStatDetailOpenApiVO() *OltStatDetailOpenApiVO`

NewOltStatDetailOpenApiVO instantiates a new OltStatDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOltStatDetailOpenApiVOWithDefaults

`func NewOltStatDetailOpenApiVOWithDefaults() *OltStatDetailOpenApiVO`

NewOltStatDetailOpenApiVOWithDefaults instantiates a new OltStatDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *OltStatDetailOpenApiVO) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OltStatDetailOpenApiVO) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OltStatDetailOpenApiVO) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *OltStatDetailOpenApiVO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDown

`func (o *OltStatDetailOpenApiVO) GetDown() int64`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *OltStatDetailOpenApiVO) GetDownOk() (*int64, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *OltStatDetailOpenApiVO) SetDown(v int64)`

SetDown sets Down field to given value.

### HasDown

`func (o *OltStatDetailOpenApiVO) HasDown() bool`

HasDown returns a boolean if a field has been set.

### GetMem

`func (o *OltStatDetailOpenApiVO) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *OltStatDetailOpenApiVO) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *OltStatDetailOpenApiVO) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *OltStatDetailOpenApiVO) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetOnuCount

`func (o *OltStatDetailOpenApiVO) GetOnuCount() int32`

GetOnuCount returns the OnuCount field if non-nil, zero value otherwise.

### GetOnuCountOk

`func (o *OltStatDetailOpenApiVO) GetOnuCountOk() (*int32, bool)`

GetOnuCountOk returns a tuple with the OnuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuCount

`func (o *OltStatDetailOpenApiVO) SetOnuCount(v int32)`

SetOnuCount sets OnuCount field to given value.

### HasOnuCount

`func (o *OltStatDetailOpenApiVO) HasOnuCount() bool`

HasOnuCount returns a boolean if a field has been set.

### GetPorts

`func (o *OltStatDetailOpenApiVO) GetPorts() []OltPortStatOpenApiVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OltStatDetailOpenApiVO) GetPortsOk() (*[]OltPortStatOpenApiVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OltStatDetailOpenApiVO) SetPorts(v []OltPortStatOpenApiVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OltStatDetailOpenApiVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTime

`func (o *OltStatDetailOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OltStatDetailOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OltStatDetailOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OltStatDetailOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUp

`func (o *OltStatDetailOpenApiVO) GetUp() int64`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *OltStatDetailOpenApiVO) GetUpOk() (*int64, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *OltStatDetailOpenApiVO) SetUp(v int64)`

SetUp sets Up field to given value.

### HasUp

`func (o *OltStatDetailOpenApiVO) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


