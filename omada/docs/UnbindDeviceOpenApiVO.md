# UnbindDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMacs** | **[]string** | Unbind device mac list, like AA-BB-CC-DD-EE-FF | 

## Methods

### NewUnbindDeviceOpenApiVO

`func NewUnbindDeviceOpenApiVO(deviceMacs []string, ) *UnbindDeviceOpenApiVO`

NewUnbindDeviceOpenApiVO instantiates a new UnbindDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnbindDeviceOpenApiVOWithDefaults

`func NewUnbindDeviceOpenApiVOWithDefaults() *UnbindDeviceOpenApiVO`

NewUnbindDeviceOpenApiVOWithDefaults instantiates a new UnbindDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMacs

`func (o *UnbindDeviceOpenApiVO) GetDeviceMacs() []string`

GetDeviceMacs returns the DeviceMacs field if non-nil, zero value otherwise.

### GetDeviceMacsOk

`func (o *UnbindDeviceOpenApiVO) GetDeviceMacsOk() (*[]string, bool)`

GetDeviceMacsOk returns a tuple with the DeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacs

`func (o *UnbindDeviceOpenApiVO) SetDeviceMacs(v []string)`

SetDeviceMacs sets DeviceMacs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


