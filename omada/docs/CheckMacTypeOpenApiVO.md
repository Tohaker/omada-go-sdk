# CheckMacTypeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **bool** | Whether the mac is device | [optional] 
**DeviceType** | Pointer to **string** | Device type:ap、switch、gateway、olt、IPC、NVR. | [optional] 

## Methods

### NewCheckMacTypeOpenApiVO

`func NewCheckMacTypeOpenApiVO() *CheckMacTypeOpenApiVO`

NewCheckMacTypeOpenApiVO instantiates a new CheckMacTypeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckMacTypeOpenApiVOWithDefaults

`func NewCheckMacTypeOpenApiVOWithDefaults() *CheckMacTypeOpenApiVO`

NewCheckMacTypeOpenApiVOWithDefaults instantiates a new CheckMacTypeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *CheckMacTypeOpenApiVO) GetDevice() bool`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CheckMacTypeOpenApiVO) GetDeviceOk() (*bool, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CheckMacTypeOpenApiVO) SetDevice(v bool)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CheckMacTypeOpenApiVO) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDeviceType

`func (o *CheckMacTypeOpenApiVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CheckMacTypeOpenApiVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CheckMacTypeOpenApiVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *CheckMacTypeOpenApiVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


