# MulticastExceptDeviceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceModel** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**DeviceModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**DeviceName** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**DeviceType** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**StackId** | Pointer to **string** | Stack Id | [optional] 
**StackName** | Pointer to **string** | Stack name | [optional] 

## Methods

### NewMulticastExceptDeviceVO

`func NewMulticastExceptDeviceVO() *MulticastExceptDeviceVO`

NewMulticastExceptDeviceVO instantiates a new MulticastExceptDeviceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticastExceptDeviceVOWithDefaults

`func NewMulticastExceptDeviceVOWithDefaults() *MulticastExceptDeviceVO`

NewMulticastExceptDeviceVOWithDefaults instantiates a new MulticastExceptDeviceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *MulticastExceptDeviceVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *MulticastExceptDeviceVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *MulticastExceptDeviceVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *MulticastExceptDeviceVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceModelVersion

`func (o *MulticastExceptDeviceVO) GetDeviceModelVersion() string`

GetDeviceModelVersion returns the DeviceModelVersion field if non-nil, zero value otherwise.

### GetDeviceModelVersionOk

`func (o *MulticastExceptDeviceVO) GetDeviceModelVersionOk() (*string, bool)`

GetDeviceModelVersionOk returns a tuple with the DeviceModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModelVersion

`func (o *MulticastExceptDeviceVO) SetDeviceModelVersion(v string)`

SetDeviceModelVersion sets DeviceModelVersion field to given value.

### HasDeviceModelVersion

`func (o *MulticastExceptDeviceVO) HasDeviceModelVersion() bool`

HasDeviceModelVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *MulticastExceptDeviceVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MulticastExceptDeviceVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MulticastExceptDeviceVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *MulticastExceptDeviceVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *MulticastExceptDeviceVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MulticastExceptDeviceVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MulticastExceptDeviceVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *MulticastExceptDeviceVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetMac

`func (o *MulticastExceptDeviceVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MulticastExceptDeviceVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MulticastExceptDeviceVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MulticastExceptDeviceVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStackId

`func (o *MulticastExceptDeviceVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *MulticastExceptDeviceVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *MulticastExceptDeviceVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *MulticastExceptDeviceVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *MulticastExceptDeviceVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *MulticastExceptDeviceVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *MulticastExceptDeviceVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *MulticastExceptDeviceVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


