# ChannelUtilInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelUtil** | Pointer to **int32** | Channel utilization rate | [optional] 
**ChannelUtil2g** | Pointer to **int32** | Channel 2g utilization rate | [optional] 
**ChannelUtil5g** | Pointer to **int32** | Channel 5g utilization rate | [optional] 
**ChannelUtil6g** | Pointer to **int32** | Channel 6g utilization rate | [optional] 
**DeviceName** | Pointer to **string** | Device name | [optional] 
**DeviceType** | Pointer to **string** | Device type | [optional] 
**Ip** | Pointer to **string** | ip | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**Model** | Pointer to **string** | Device model name with version | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device, for example:3.0 | [optional] 

## Methods

### NewChannelUtilInfoVO

`func NewChannelUtilInfoVO() *ChannelUtilInfoVO`

NewChannelUtilInfoVO instantiates a new ChannelUtilInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelUtilInfoVOWithDefaults

`func NewChannelUtilInfoVOWithDefaults() *ChannelUtilInfoVO`

NewChannelUtilInfoVOWithDefaults instantiates a new ChannelUtilInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelUtil

`func (o *ChannelUtilInfoVO) GetChannelUtil() int32`

GetChannelUtil returns the ChannelUtil field if non-nil, zero value otherwise.

### GetChannelUtilOk

`func (o *ChannelUtilInfoVO) GetChannelUtilOk() (*int32, bool)`

GetChannelUtilOk returns a tuple with the ChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtil

`func (o *ChannelUtilInfoVO) SetChannelUtil(v int32)`

SetChannelUtil sets ChannelUtil field to given value.

### HasChannelUtil

`func (o *ChannelUtilInfoVO) HasChannelUtil() bool`

HasChannelUtil returns a boolean if a field has been set.

### GetChannelUtil2g

`func (o *ChannelUtilInfoVO) GetChannelUtil2g() int32`

GetChannelUtil2g returns the ChannelUtil2g field if non-nil, zero value otherwise.

### GetChannelUtil2gOk

`func (o *ChannelUtilInfoVO) GetChannelUtil2gOk() (*int32, bool)`

GetChannelUtil2gOk returns a tuple with the ChannelUtil2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtil2g

`func (o *ChannelUtilInfoVO) SetChannelUtil2g(v int32)`

SetChannelUtil2g sets ChannelUtil2g field to given value.

### HasChannelUtil2g

`func (o *ChannelUtilInfoVO) HasChannelUtil2g() bool`

HasChannelUtil2g returns a boolean if a field has been set.

### GetChannelUtil5g

`func (o *ChannelUtilInfoVO) GetChannelUtil5g() int32`

GetChannelUtil5g returns the ChannelUtil5g field if non-nil, zero value otherwise.

### GetChannelUtil5gOk

`func (o *ChannelUtilInfoVO) GetChannelUtil5gOk() (*int32, bool)`

GetChannelUtil5gOk returns a tuple with the ChannelUtil5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtil5g

`func (o *ChannelUtilInfoVO) SetChannelUtil5g(v int32)`

SetChannelUtil5g sets ChannelUtil5g field to given value.

### HasChannelUtil5g

`func (o *ChannelUtilInfoVO) HasChannelUtil5g() bool`

HasChannelUtil5g returns a boolean if a field has been set.

### GetChannelUtil6g

`func (o *ChannelUtilInfoVO) GetChannelUtil6g() int32`

GetChannelUtil6g returns the ChannelUtil6g field if non-nil, zero value otherwise.

### GetChannelUtil6gOk

`func (o *ChannelUtilInfoVO) GetChannelUtil6gOk() (*int32, bool)`

GetChannelUtil6gOk returns a tuple with the ChannelUtil6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtil6g

`func (o *ChannelUtilInfoVO) SetChannelUtil6g(v int32)`

SetChannelUtil6g sets ChannelUtil6g field to given value.

### HasChannelUtil6g

`func (o *ChannelUtilInfoVO) HasChannelUtil6g() bool`

HasChannelUtil6g returns a boolean if a field has been set.

### GetDeviceName

`func (o *ChannelUtilInfoVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ChannelUtilInfoVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ChannelUtilInfoVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ChannelUtilInfoVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *ChannelUtilInfoVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ChannelUtilInfoVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ChannelUtilInfoVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ChannelUtilInfoVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIp

`func (o *ChannelUtilInfoVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ChannelUtilInfoVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ChannelUtilInfoVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ChannelUtilInfoVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *ChannelUtilInfoVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ChannelUtilInfoVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ChannelUtilInfoVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ChannelUtilInfoVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ChannelUtilInfoVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChannelUtilInfoVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChannelUtilInfoVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChannelUtilInfoVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ChannelUtilInfoVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ChannelUtilInfoVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ChannelUtilInfoVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ChannelUtilInfoVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


