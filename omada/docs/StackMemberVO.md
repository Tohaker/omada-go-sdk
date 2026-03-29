# StackMemberVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceModel** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**DeviceModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. | [optional] 
**DeviceType** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Mac** | Pointer to **string** | The unique identification of the device. | [optional] 

## Methods

### NewStackMemberVO

`func NewStackMemberVO() *StackMemberVO`

NewStackMemberVO instantiates a new StackMemberVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackMemberVOWithDefaults

`func NewStackMemberVOWithDefaults() *StackMemberVO`

NewStackMemberVOWithDefaults instantiates a new StackMemberVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *StackMemberVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *StackMemberVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *StackMemberVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *StackMemberVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceModelVersion

`func (o *StackMemberVO) GetDeviceModelVersion() string`

GetDeviceModelVersion returns the DeviceModelVersion field if non-nil, zero value otherwise.

### GetDeviceModelVersionOk

`func (o *StackMemberVO) GetDeviceModelVersionOk() (*string, bool)`

GetDeviceModelVersionOk returns a tuple with the DeviceModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModelVersion

`func (o *StackMemberVO) SetDeviceModelVersion(v string)`

SetDeviceModelVersion sets DeviceModelVersion field to given value.

### HasDeviceModelVersion

`func (o *StackMemberVO) HasDeviceModelVersion() bool`

HasDeviceModelVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *StackMemberVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *StackMemberVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *StackMemberVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *StackMemberVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *StackMemberVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *StackMemberVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *StackMemberVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *StackMemberVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetMac

`func (o *StackMemberVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *StackMemberVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *StackMemberVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *StackMemberVO) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


