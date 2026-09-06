# ChannelInterInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelInterf** | Pointer to **int32** | Channel interference rate | [optional] 
**ChannelInterf2g** | Pointer to **int32** | Channel interference rate | [optional] 
**ChannelInterf5g** | Pointer to **int32** | Channel interference rate | [optional] 
**ChannelInterf6g** | Pointer to **int32** | Channel interference rate | [optional] 
**DeviceName** | Pointer to **string** | Device name | [optional] 
**DeviceType** | Pointer to **string** | Device type | [optional] 
**Ip** | Pointer to **string** | ip | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**Model** | Pointer to **string** | Device model name with version | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device, for example:3.0 | [optional] 
**NoiseFloor2g** | Pointer to **int32** | Noise floor of 2.4g | [optional] 
**NoiseFloor5g** | Pointer to **int32** | Noise floor of 5g | [optional] 
**NoiseFloor6g** | Pointer to **int32** | Noise floor of 6g | [optional] 

## Methods

### NewChannelInterInfoVO

`func NewChannelInterInfoVO() *ChannelInterInfoVO`

NewChannelInterInfoVO instantiates a new ChannelInterInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelInterInfoVOWithDefaults

`func NewChannelInterInfoVOWithDefaults() *ChannelInterInfoVO`

NewChannelInterInfoVOWithDefaults instantiates a new ChannelInterInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelInterf

`func (o *ChannelInterInfoVO) GetChannelInterf() int32`

GetChannelInterf returns the ChannelInterf field if non-nil, zero value otherwise.

### GetChannelInterfOk

`func (o *ChannelInterInfoVO) GetChannelInterfOk() (*int32, bool)`

GetChannelInterfOk returns a tuple with the ChannelInterf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelInterf

`func (o *ChannelInterInfoVO) SetChannelInterf(v int32)`

SetChannelInterf sets ChannelInterf field to given value.

### HasChannelInterf

`func (o *ChannelInterInfoVO) HasChannelInterf() bool`

HasChannelInterf returns a boolean if a field has been set.

### GetChannelInterf2g

`func (o *ChannelInterInfoVO) GetChannelInterf2g() int32`

GetChannelInterf2g returns the ChannelInterf2g field if non-nil, zero value otherwise.

### GetChannelInterf2gOk

`func (o *ChannelInterInfoVO) GetChannelInterf2gOk() (*int32, bool)`

GetChannelInterf2gOk returns a tuple with the ChannelInterf2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelInterf2g

`func (o *ChannelInterInfoVO) SetChannelInterf2g(v int32)`

SetChannelInterf2g sets ChannelInterf2g field to given value.

### HasChannelInterf2g

`func (o *ChannelInterInfoVO) HasChannelInterf2g() bool`

HasChannelInterf2g returns a boolean if a field has been set.

### GetChannelInterf5g

`func (o *ChannelInterInfoVO) GetChannelInterf5g() int32`

GetChannelInterf5g returns the ChannelInterf5g field if non-nil, zero value otherwise.

### GetChannelInterf5gOk

`func (o *ChannelInterInfoVO) GetChannelInterf5gOk() (*int32, bool)`

GetChannelInterf5gOk returns a tuple with the ChannelInterf5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelInterf5g

`func (o *ChannelInterInfoVO) SetChannelInterf5g(v int32)`

SetChannelInterf5g sets ChannelInterf5g field to given value.

### HasChannelInterf5g

`func (o *ChannelInterInfoVO) HasChannelInterf5g() bool`

HasChannelInterf5g returns a boolean if a field has been set.

### GetChannelInterf6g

`func (o *ChannelInterInfoVO) GetChannelInterf6g() int32`

GetChannelInterf6g returns the ChannelInterf6g field if non-nil, zero value otherwise.

### GetChannelInterf6gOk

`func (o *ChannelInterInfoVO) GetChannelInterf6gOk() (*int32, bool)`

GetChannelInterf6gOk returns a tuple with the ChannelInterf6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelInterf6g

`func (o *ChannelInterInfoVO) SetChannelInterf6g(v int32)`

SetChannelInterf6g sets ChannelInterf6g field to given value.

### HasChannelInterf6g

`func (o *ChannelInterInfoVO) HasChannelInterf6g() bool`

HasChannelInterf6g returns a boolean if a field has been set.

### GetDeviceName

`func (o *ChannelInterInfoVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ChannelInterInfoVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ChannelInterInfoVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ChannelInterInfoVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *ChannelInterInfoVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ChannelInterInfoVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ChannelInterInfoVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ChannelInterInfoVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIp

`func (o *ChannelInterInfoVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ChannelInterInfoVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ChannelInterInfoVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ChannelInterInfoVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *ChannelInterInfoVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ChannelInterInfoVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ChannelInterInfoVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ChannelInterInfoVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ChannelInterInfoVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChannelInterInfoVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChannelInterInfoVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChannelInterInfoVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ChannelInterInfoVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ChannelInterInfoVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ChannelInterInfoVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ChannelInterInfoVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetNoiseFloor2g

`func (o *ChannelInterInfoVO) GetNoiseFloor2g() int32`

GetNoiseFloor2g returns the NoiseFloor2g field if non-nil, zero value otherwise.

### GetNoiseFloor2gOk

`func (o *ChannelInterInfoVO) GetNoiseFloor2gOk() (*int32, bool)`

GetNoiseFloor2gOk returns a tuple with the NoiseFloor2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoiseFloor2g

`func (o *ChannelInterInfoVO) SetNoiseFloor2g(v int32)`

SetNoiseFloor2g sets NoiseFloor2g field to given value.

### HasNoiseFloor2g

`func (o *ChannelInterInfoVO) HasNoiseFloor2g() bool`

HasNoiseFloor2g returns a boolean if a field has been set.

### GetNoiseFloor5g

`func (o *ChannelInterInfoVO) GetNoiseFloor5g() int32`

GetNoiseFloor5g returns the NoiseFloor5g field if non-nil, zero value otherwise.

### GetNoiseFloor5gOk

`func (o *ChannelInterInfoVO) GetNoiseFloor5gOk() (*int32, bool)`

GetNoiseFloor5gOk returns a tuple with the NoiseFloor5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoiseFloor5g

`func (o *ChannelInterInfoVO) SetNoiseFloor5g(v int32)`

SetNoiseFloor5g sets NoiseFloor5g field to given value.

### HasNoiseFloor5g

`func (o *ChannelInterInfoVO) HasNoiseFloor5g() bool`

HasNoiseFloor5g returns a boolean if a field has been set.

### GetNoiseFloor6g

`func (o *ChannelInterInfoVO) GetNoiseFloor6g() int32`

GetNoiseFloor6g returns the NoiseFloor6g field if non-nil, zero value otherwise.

### GetNoiseFloor6gOk

`func (o *ChannelInterInfoVO) GetNoiseFloor6gOk() (*int32, bool)`

GetNoiseFloor6gOk returns a tuple with the NoiseFloor6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoiseFloor6g

`func (o *ChannelInterInfoVO) SetNoiseFloor6g(v int32)`

SetNoiseFloor6g sets NoiseFloor6g field to given value.

### HasNoiseFloor6g

`func (o *ChannelInterInfoVO) HasNoiseFloor6g() bool`

HasNoiseFloor6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


