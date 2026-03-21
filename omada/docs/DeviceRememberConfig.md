# DeviceRememberConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remember** | Pointer to **int32** | The remember configuration of the device, 0 means disable, 1 means enable, 2 means follow site. | [optional] 

## Methods

### NewDeviceRememberConfig

`func NewDeviceRememberConfig() *DeviceRememberConfig`

NewDeviceRememberConfig instantiates a new DeviceRememberConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRememberConfigWithDefaults

`func NewDeviceRememberConfigWithDefaults() *DeviceRememberConfig`

NewDeviceRememberConfigWithDefaults instantiates a new DeviceRememberConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemember

`func (o *DeviceRememberConfig) GetRemember() int32`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *DeviceRememberConfig) GetRememberOk() (*int32, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *DeviceRememberConfig) SetRemember(v int32)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *DeviceRememberConfig) HasRemember() bool`

HasRemember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


