# ModifyCliTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliConfig** | **string** | CLI configuration content | 
**Description** | Pointer to **string** | CLI configuration description, it should be within the range of 0 - 256 characters. | [optional] 
**DeviceType** | **string** | Device type for CLI configuration application, it should be a value as follows: switch. | 
**Name** | **string** | CLI configuration name, it should be within the range of 1 - 64 characters | 

## Methods

### NewModifyCliTemplateOpenApiVO

`func NewModifyCliTemplateOpenApiVO(cliConfig string, deviceType string, name string, ) *ModifyCliTemplateOpenApiVO`

NewModifyCliTemplateOpenApiVO instantiates a new ModifyCliTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCliTemplateOpenApiVOWithDefaults

`func NewModifyCliTemplateOpenApiVOWithDefaults() *ModifyCliTemplateOpenApiVO`

NewModifyCliTemplateOpenApiVOWithDefaults instantiates a new ModifyCliTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliConfig

`func (o *ModifyCliTemplateOpenApiVO) GetCliConfig() string`

GetCliConfig returns the CliConfig field if non-nil, zero value otherwise.

### GetCliConfigOk

`func (o *ModifyCliTemplateOpenApiVO) GetCliConfigOk() (*string, bool)`

GetCliConfigOk returns a tuple with the CliConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliConfig

`func (o *ModifyCliTemplateOpenApiVO) SetCliConfig(v string)`

SetCliConfig sets CliConfig field to given value.


### GetDescription

`func (o *ModifyCliTemplateOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyCliTemplateOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyCliTemplateOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModifyCliTemplateOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceType

`func (o *ModifyCliTemplateOpenApiVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ModifyCliTemplateOpenApiVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ModifyCliTemplateOpenApiVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetName

`func (o *ModifyCliTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyCliTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyCliTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


