# DeviceTemplateEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoBind** | Pointer to **bool** | Devices will automatically bind the template and use template configurations. | [optional] 
**Status** | Pointer to **int32** | The status of device template. 0:completed ; 1:draft | [optional] 
**TemplateName** | Pointer to **string** | The name of device template. | [optional] 

## Methods

### NewDeviceTemplateEdit

`func NewDeviceTemplateEdit() *DeviceTemplateEdit`

NewDeviceTemplateEdit instantiates a new DeviceTemplateEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTemplateEditWithDefaults

`func NewDeviceTemplateEditWithDefaults() *DeviceTemplateEdit`

NewDeviceTemplateEditWithDefaults instantiates a new DeviceTemplateEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoBind

`func (o *DeviceTemplateEdit) GetAutoBind() bool`

GetAutoBind returns the AutoBind field if non-nil, zero value otherwise.

### GetAutoBindOk

`func (o *DeviceTemplateEdit) GetAutoBindOk() (*bool, bool)`

GetAutoBindOk returns a tuple with the AutoBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBind

`func (o *DeviceTemplateEdit) SetAutoBind(v bool)`

SetAutoBind sets AutoBind field to given value.

### HasAutoBind

`func (o *DeviceTemplateEdit) HasAutoBind() bool`

HasAutoBind returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceTemplateEdit) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceTemplateEdit) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceTemplateEdit) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceTemplateEdit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateName

`func (o *DeviceTemplateEdit) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *DeviceTemplateEdit) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *DeviceTemplateEdit) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *DeviceTemplateEdit) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


