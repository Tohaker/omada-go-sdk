# DeviceTemplateAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** | The type of device. | [optional] 
**Model** | Pointer to **string** | The model name of device. | [optional] 
**ModelVersion** | Pointer to **string** | The model version of device.For example: 1.0 | [optional] 
**TemplateName** | Pointer to **string** | The name of device template. | [optional] 
**TemplateSettings** | Pointer to **[]int32** | The configurable modules of device. For switch:  1:port; 2:vlanInterface; 3:staticRoute; 4:services; 5:ipSetting. For gateway: 1:port; 2:radios; 3:wlans; 4:services; 5:advanced. | [optional] 

## Methods

### NewDeviceTemplateAdd

`func NewDeviceTemplateAdd() *DeviceTemplateAdd`

NewDeviceTemplateAdd instantiates a new DeviceTemplateAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTemplateAddWithDefaults

`func NewDeviceTemplateAddWithDefaults() *DeviceTemplateAdd`

NewDeviceTemplateAddWithDefaults instantiates a new DeviceTemplateAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *DeviceTemplateAdd) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceTemplateAdd) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceTemplateAdd) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceTemplateAdd) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetModel

`func (o *DeviceTemplateAdd) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceTemplateAdd) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceTemplateAdd) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceTemplateAdd) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceTemplateAdd) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceTemplateAdd) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceTemplateAdd) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceTemplateAdd) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetTemplateName

`func (o *DeviceTemplateAdd) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *DeviceTemplateAdd) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *DeviceTemplateAdd) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *DeviceTemplateAdd) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateSettings

`func (o *DeviceTemplateAdd) GetTemplateSettings() []int32`

GetTemplateSettings returns the TemplateSettings field if non-nil, zero value otherwise.

### GetTemplateSettingsOk

`func (o *DeviceTemplateAdd) GetTemplateSettingsOk() (*[]int32, bool)`

GetTemplateSettingsOk returns a tuple with the TemplateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSettings

`func (o *DeviceTemplateAdd) SetTemplateSettings(v []int32)`

SetTemplateSettings sets TemplateSettings field to given value.

### HasTemplateSettings

`func (o *DeviceTemplateAdd) HasTemplateSettings() bool`

HasTemplateSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


