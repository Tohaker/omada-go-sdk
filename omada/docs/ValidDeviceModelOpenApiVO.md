# ValidDeviceModelOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateways** | Pointer to [**[]DeviceGatewayModelOpenApiVO**](DeviceGatewayModelOpenApiVO.md) |  | [optional] 
**Switches** | Pointer to [**[]DeviceSwitchModelOpenApiVO**](DeviceSwitchModelOpenApiVO.md) |  | [optional] 

## Methods

### NewValidDeviceModelOpenApiVO

`func NewValidDeviceModelOpenApiVO() *ValidDeviceModelOpenApiVO`

NewValidDeviceModelOpenApiVO instantiates a new ValidDeviceModelOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidDeviceModelOpenApiVOWithDefaults

`func NewValidDeviceModelOpenApiVOWithDefaults() *ValidDeviceModelOpenApiVO`

NewValidDeviceModelOpenApiVOWithDefaults instantiates a new ValidDeviceModelOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateways

`func (o *ValidDeviceModelOpenApiVO) GetGateways() []DeviceGatewayModelOpenApiVO`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *ValidDeviceModelOpenApiVO) GetGatewaysOk() (*[]DeviceGatewayModelOpenApiVO, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *ValidDeviceModelOpenApiVO) SetGateways(v []DeviceGatewayModelOpenApiVO)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *ValidDeviceModelOpenApiVO) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetSwitches

`func (o *ValidDeviceModelOpenApiVO) GetSwitches() []DeviceSwitchModelOpenApiVO`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *ValidDeviceModelOpenApiVO) GetSwitchesOk() (*[]DeviceSwitchModelOpenApiVO, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *ValidDeviceModelOpenApiVO) SetSwitches(v []DeviceSwitchModelOpenApiVO)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *ValidDeviceModelOpenApiVO) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


