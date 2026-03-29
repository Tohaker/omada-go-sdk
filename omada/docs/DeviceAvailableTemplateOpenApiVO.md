# DeviceAvailableTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The mac of device. | [optional] 
**Model** | Pointer to **string** | The model name of device. | [optional] 
**ModelVersion** | Pointer to **string** | The model version of device. | [optional] 
**Name** | Pointer to **string** | The name of device, default value is mac. | [optional] 
**ValidTemplates** | Pointer to [**[]AvailableTemplateOpenApiVO**](AvailableTemplateOpenApiVO.md) | The available templates of device. | [optional] 

## Methods

### NewDeviceAvailableTemplateOpenApiVO

`func NewDeviceAvailableTemplateOpenApiVO() *DeviceAvailableTemplateOpenApiVO`

NewDeviceAvailableTemplateOpenApiVO instantiates a new DeviceAvailableTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAvailableTemplateOpenApiVOWithDefaults

`func NewDeviceAvailableTemplateOpenApiVOWithDefaults() *DeviceAvailableTemplateOpenApiVO`

NewDeviceAvailableTemplateOpenApiVOWithDefaults instantiates a new DeviceAvailableTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *DeviceAvailableTemplateOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceAvailableTemplateOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceAvailableTemplateOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceAvailableTemplateOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *DeviceAvailableTemplateOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceAvailableTemplateOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceAvailableTemplateOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceAvailableTemplateOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceAvailableTemplateOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceAvailableTemplateOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceAvailableTemplateOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceAvailableTemplateOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DeviceAvailableTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAvailableTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAvailableTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceAvailableTemplateOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValidTemplates

`func (o *DeviceAvailableTemplateOpenApiVO) GetValidTemplates() []AvailableTemplateOpenApiVO`

GetValidTemplates returns the ValidTemplates field if non-nil, zero value otherwise.

### GetValidTemplatesOk

`func (o *DeviceAvailableTemplateOpenApiVO) GetValidTemplatesOk() (*[]AvailableTemplateOpenApiVO, bool)`

GetValidTemplatesOk returns a tuple with the ValidTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTemplates

`func (o *DeviceAvailableTemplateOpenApiVO) SetValidTemplates(v []AvailableTemplateOpenApiVO)`

SetValidTemplates sets ValidTemplates field to given value.

### HasValidTemplates

`func (o *DeviceAvailableTemplateOpenApiVO) HasValidTemplates() bool`

HasValidTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


