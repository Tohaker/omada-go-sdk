# OswPoeResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceResult** | Pointer to [**PoeRecoverDeviceOpenApiVO**](PoeRecoverDeviceOpenApiVO.md) |  | [optional] 
**Process** | Pointer to **int32** | PoE recovery progress. | [optional] 

## Methods

### NewOswPoeResultOpenApiVO

`func NewOswPoeResultOpenApiVO() *OswPoeResultOpenApiVO`

NewOswPoeResultOpenApiVO instantiates a new OswPoeResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPoeResultOpenApiVOWithDefaults

`func NewOswPoeResultOpenApiVOWithDefaults() *OswPoeResultOpenApiVO`

NewOswPoeResultOpenApiVOWithDefaults instantiates a new OswPoeResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceResult

`func (o *OswPoeResultOpenApiVO) GetDeviceResult() PoeRecoverDeviceOpenApiVO`

GetDeviceResult returns the DeviceResult field if non-nil, zero value otherwise.

### GetDeviceResultOk

`func (o *OswPoeResultOpenApiVO) GetDeviceResultOk() (*PoeRecoverDeviceOpenApiVO, bool)`

GetDeviceResultOk returns a tuple with the DeviceResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceResult

`func (o *OswPoeResultOpenApiVO) SetDeviceResult(v PoeRecoverDeviceOpenApiVO)`

SetDeviceResult sets DeviceResult field to given value.

### HasDeviceResult

`func (o *OswPoeResultOpenApiVO) HasDeviceResult() bool`

HasDeviceResult returns a boolean if a field has been set.

### GetProcess

`func (o *OswPoeResultOpenApiVO) GetProcess() int32`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *OswPoeResultOpenApiVO) GetProcessOk() (*int32, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *OswPoeResultOpenApiVO) SetProcess(v int32)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *OswPoeResultOpenApiVO) HasProcess() bool`

HasProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


