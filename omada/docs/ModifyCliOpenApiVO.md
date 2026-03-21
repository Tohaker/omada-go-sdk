# ModifyCliOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliConfig** | **string** | CLI configuration content | 
**Description** | Pointer to **string** | CLI configuration description, it should be within the range of 0 - 256 characters. | [optional] 
**DeviceType** | **string** | Device type for CLI configuration application, it should be a value as follows: switch. | 
**Devices** | Pointer to [**[]DeviceCliVO**](DeviceCliVO.md) | List of devices bound to the CLI configuration. | [optional] 
**Model** | Pointer to **string** | Device model of model cli | [optional] 
**ModelVersion** | Pointer to **string** | Device model version of model cli | [optional] 
**Name** | **string** | CLI configuration name, it should be within the range of 1 - 64 characters | 
**Stacks** | Pointer to [**[]StackCliVO**](StackCliVO.md) | List of stacks bound to the CLI configuration, only device CLI has this field. | [optional] 

## Methods

### NewModifyCliOpenApiVO

`func NewModifyCliOpenApiVO(cliConfig string, deviceType string, name string, ) *ModifyCliOpenApiVO`

NewModifyCliOpenApiVO instantiates a new ModifyCliOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCliOpenApiVOWithDefaults

`func NewModifyCliOpenApiVOWithDefaults() *ModifyCliOpenApiVO`

NewModifyCliOpenApiVOWithDefaults instantiates a new ModifyCliOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliConfig

`func (o *ModifyCliOpenApiVO) GetCliConfig() string`

GetCliConfig returns the CliConfig field if non-nil, zero value otherwise.

### GetCliConfigOk

`func (o *ModifyCliOpenApiVO) GetCliConfigOk() (*string, bool)`

GetCliConfigOk returns a tuple with the CliConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliConfig

`func (o *ModifyCliOpenApiVO) SetCliConfig(v string)`

SetCliConfig sets CliConfig field to given value.


### GetDescription

`func (o *ModifyCliOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyCliOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyCliOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModifyCliOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceType

`func (o *ModifyCliOpenApiVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ModifyCliOpenApiVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ModifyCliOpenApiVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetDevices

`func (o *ModifyCliOpenApiVO) GetDevices() []DeviceCliVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ModifyCliOpenApiVO) GetDevicesOk() (*[]DeviceCliVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ModifyCliOpenApiVO) SetDevices(v []DeviceCliVO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *ModifyCliOpenApiVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetModel

`func (o *ModifyCliOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModifyCliOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModifyCliOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModifyCliOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ModifyCliOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ModifyCliOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ModifyCliOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ModifyCliOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ModifyCliOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyCliOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyCliOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetStacks

`func (o *ModifyCliOpenApiVO) GetStacks() []StackCliVO`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *ModifyCliOpenApiVO) GetStacksOk() (*[]StackCliVO, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *ModifyCliOpenApiVO) SetStacks(v []StackCliVO)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *ModifyCliOpenApiVO) HasStacks() bool`

HasStacks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


