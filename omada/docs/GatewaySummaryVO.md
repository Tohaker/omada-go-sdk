# GatewaySummaryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAlert** | Pointer to [**AlertVO**](AlertVO.md) |  | [optional] 
**GatewayUltilization** | Pointer to [**[]GatewayCpuMemUtilListVO**](GatewayCpuMemUtilListVO.md) | Gateway CPU and memory utilization | [optional] 

## Methods

### NewGatewaySummaryVO

`func NewGatewaySummaryVO() *GatewaySummaryVO`

NewGatewaySummaryVO instantiates a new GatewaySummaryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaySummaryVOWithDefaults

`func NewGatewaySummaryVOWithDefaults() *GatewaySummaryVO`

NewGatewaySummaryVOWithDefaults instantiates a new GatewaySummaryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAlert

`func (o *GatewaySummaryVO) GetGatewayAlert() AlertVO`

GetGatewayAlert returns the GatewayAlert field if non-nil, zero value otherwise.

### GetGatewayAlertOk

`func (o *GatewaySummaryVO) GetGatewayAlertOk() (*AlertVO, bool)`

GetGatewayAlertOk returns a tuple with the GatewayAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAlert

`func (o *GatewaySummaryVO) SetGatewayAlert(v AlertVO)`

SetGatewayAlert sets GatewayAlert field to given value.

### HasGatewayAlert

`func (o *GatewaySummaryVO) HasGatewayAlert() bool`

HasGatewayAlert returns a boolean if a field has been set.

### GetGatewayUltilization

`func (o *GatewaySummaryVO) GetGatewayUltilization() []GatewayCpuMemUtilListVO`

GetGatewayUltilization returns the GatewayUltilization field if non-nil, zero value otherwise.

### GetGatewayUltilizationOk

`func (o *GatewaySummaryVO) GetGatewayUltilizationOk() (*[]GatewayCpuMemUtilListVO, bool)`

GetGatewayUltilizationOk returns a tuple with the GatewayUltilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUltilization

`func (o *GatewaySummaryVO) SetGatewayUltilization(v []GatewayCpuMemUtilListVO)`

SetGatewayUltilization sets GatewayUltilization field to given value.

### HasGatewayUltilization

`func (o *GatewaySummaryVO) HasGatewayUltilization() bool`

HasGatewayUltilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


