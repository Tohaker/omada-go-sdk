# CliVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apply** | Pointer to **bool** | Whether the CLI configuration can be applied. Only the CLI configuration in inactive state can be applied. | [optional] 
**Description** | Pointer to **string** | CLI configuration description, it should be within the range of 0 - 256 characters. | [optional] 
**Devices** | Pointer to [**[]DeviceCliVO**](DeviceCliVO.md) | List of devices bound to the CLI configuration, only device CLI has this field. | [optional] 
**Id** | Pointer to **string** | CLI configuration ID | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | CLI configuration name | [optional] 
**Stacks** | Pointer to [**[]StackCliVO**](StackCliVO.md) | List of stacks bound to the CLI configuration, only device CLI has this field. | [optional] 
**Status** | Pointer to **int32** | CLI configuration status, it should be a value as follows: 0: active, 1: inactive | [optional] 

## Methods

### NewCliVO

`func NewCliVO() *CliVO`

NewCliVO instantiates a new CliVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCliVOWithDefaults

`func NewCliVOWithDefaults() *CliVO`

NewCliVOWithDefaults instantiates a new CliVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApply

`func (o *CliVO) GetApply() bool`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *CliVO) GetApplyOk() (*bool, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *CliVO) SetApply(v bool)`

SetApply sets Apply field to given value.

### HasApply

`func (o *CliVO) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetDescription

`func (o *CliVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CliVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CliVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CliVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevices

`func (o *CliVO) GetDevices() []DeviceCliVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CliVO) GetDevicesOk() (*[]DeviceCliVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CliVO) SetDevices(v []DeviceCliVO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *CliVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetId

`func (o *CliVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CliVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CliVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CliVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *CliVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CliVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CliVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CliVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *CliVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *CliVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *CliVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *CliVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *CliVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CliVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CliVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CliVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStacks

`func (o *CliVO) GetStacks() []StackCliVO`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *CliVO) GetStacksOk() (*[]StackCliVO, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *CliVO) SetStacks(v []StackCliVO)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *CliVO) HasStacks() bool`

HasStacks returns a boolean if a field has been set.

### GetStatus

`func (o *CliVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CliVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CliVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CliVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


