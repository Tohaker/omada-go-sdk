# DeviceBindOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceTemplateId** | Pointer to **string** | Switch device template ID. | [optional] 
**Mac** | Pointer to **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF. | [optional] 

## Methods

### NewDeviceBindOpenApiVO

`func NewDeviceBindOpenApiVO() *DeviceBindOpenApiVO`

NewDeviceBindOpenApiVO instantiates a new DeviceBindOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBindOpenApiVOWithDefaults

`func NewDeviceBindOpenApiVOWithDefaults() *DeviceBindOpenApiVO`

NewDeviceBindOpenApiVOWithDefaults instantiates a new DeviceBindOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceTemplateId

`func (o *DeviceBindOpenApiVO) GetDeviceTemplateId() string`

GetDeviceTemplateId returns the DeviceTemplateId field if non-nil, zero value otherwise.

### GetDeviceTemplateIdOk

`func (o *DeviceBindOpenApiVO) GetDeviceTemplateIdOk() (*string, bool)`

GetDeviceTemplateIdOk returns a tuple with the DeviceTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplateId

`func (o *DeviceBindOpenApiVO) SetDeviceTemplateId(v string)`

SetDeviceTemplateId sets DeviceTemplateId field to given value.

### HasDeviceTemplateId

`func (o *DeviceBindOpenApiVO) HasDeviceTemplateId() bool`

HasDeviceTemplateId returns a boolean if a field has been set.

### GetMac

`func (o *DeviceBindOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceBindOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceBindOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceBindOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


