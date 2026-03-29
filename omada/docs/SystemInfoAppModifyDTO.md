# SystemInfoAppModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactInformation** | Pointer to **string** | Contact information should contain 1-32 bits numbers, Upper and lower letters, -@_:/. . | [optional] 
**DeviceLocation** | Pointer to **string** | Device location should contain 1-32 bits numbers, Upper and lower letters, -@_:/. . | [optional] 
**DeviceName** | Pointer to **string** | Device name should contain 1-32 bits numbers, Upper and lower letters, -@_:/. . | [optional] 

## Methods

### NewSystemInfoAppModifyDTO

`func NewSystemInfoAppModifyDTO() *SystemInfoAppModifyDTO`

NewSystemInfoAppModifyDTO instantiates a new SystemInfoAppModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInfoAppModifyDTOWithDefaults

`func NewSystemInfoAppModifyDTOWithDefaults() *SystemInfoAppModifyDTO`

NewSystemInfoAppModifyDTOWithDefaults instantiates a new SystemInfoAppModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactInformation

`func (o *SystemInfoAppModifyDTO) GetContactInformation() string`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *SystemInfoAppModifyDTO) GetContactInformationOk() (*string, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *SystemInfoAppModifyDTO) SetContactInformation(v string)`

SetContactInformation sets ContactInformation field to given value.

### HasContactInformation

`func (o *SystemInfoAppModifyDTO) HasContactInformation() bool`

HasContactInformation returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *SystemInfoAppModifyDTO) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *SystemInfoAppModifyDTO) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *SystemInfoAppModifyDTO) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *SystemInfoAppModifyDTO) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceName

`func (o *SystemInfoAppModifyDTO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SystemInfoAppModifyDTO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SystemInfoAppModifyDTO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SystemInfoAppModifyDTO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


