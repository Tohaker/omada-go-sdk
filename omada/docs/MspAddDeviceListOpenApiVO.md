# MspAddDeviceListOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddType** | Pointer to **int32** | 0/null indicates SN, and 1 indicates deviceKey | [optional] 
**AllowOffline** | Pointer to **bool** | Whether to allow offline addition. The default value is False | [optional] 
**Devices** | [**[]AddDeviceVO**](AddDeviceVO.md) | Add devices list | 
**KeepResult** | Pointer to **bool** | Save the last result | [optional] 
**MspId** | Pointer to **string** | MSP ID | [optional] 
**VerifyCode** | Pointer to **string** |  | [optional] 

## Methods

### NewMspAddDeviceListOpenApiVO

`func NewMspAddDeviceListOpenApiVO(devices []AddDeviceVO, ) *MspAddDeviceListOpenApiVO`

NewMspAddDeviceListOpenApiVO instantiates a new MspAddDeviceListOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspAddDeviceListOpenApiVOWithDefaults

`func NewMspAddDeviceListOpenApiVOWithDefaults() *MspAddDeviceListOpenApiVO`

NewMspAddDeviceListOpenApiVOWithDefaults instantiates a new MspAddDeviceListOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddType

`func (o *MspAddDeviceListOpenApiVO) GetAddType() int32`

GetAddType returns the AddType field if non-nil, zero value otherwise.

### GetAddTypeOk

`func (o *MspAddDeviceListOpenApiVO) GetAddTypeOk() (*int32, bool)`

GetAddTypeOk returns a tuple with the AddType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddType

`func (o *MspAddDeviceListOpenApiVO) SetAddType(v int32)`

SetAddType sets AddType field to given value.

### HasAddType

`func (o *MspAddDeviceListOpenApiVO) HasAddType() bool`

HasAddType returns a boolean if a field has been set.

### GetAllowOffline

`func (o *MspAddDeviceListOpenApiVO) GetAllowOffline() bool`

GetAllowOffline returns the AllowOffline field if non-nil, zero value otherwise.

### GetAllowOfflineOk

`func (o *MspAddDeviceListOpenApiVO) GetAllowOfflineOk() (*bool, bool)`

GetAllowOfflineOk returns a tuple with the AllowOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOffline

`func (o *MspAddDeviceListOpenApiVO) SetAllowOffline(v bool)`

SetAllowOffline sets AllowOffline field to given value.

### HasAllowOffline

`func (o *MspAddDeviceListOpenApiVO) HasAllowOffline() bool`

HasAllowOffline returns a boolean if a field has been set.

### GetDevices

`func (o *MspAddDeviceListOpenApiVO) GetDevices() []AddDeviceVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *MspAddDeviceListOpenApiVO) GetDevicesOk() (*[]AddDeviceVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *MspAddDeviceListOpenApiVO) SetDevices(v []AddDeviceVO)`

SetDevices sets Devices field to given value.


### GetKeepResult

`func (o *MspAddDeviceListOpenApiVO) GetKeepResult() bool`

GetKeepResult returns the KeepResult field if non-nil, zero value otherwise.

### GetKeepResultOk

`func (o *MspAddDeviceListOpenApiVO) GetKeepResultOk() (*bool, bool)`

GetKeepResultOk returns a tuple with the KeepResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResult

`func (o *MspAddDeviceListOpenApiVO) SetKeepResult(v bool)`

SetKeepResult sets KeepResult field to given value.

### HasKeepResult

`func (o *MspAddDeviceListOpenApiVO) HasKeepResult() bool`

HasKeepResult returns a boolean if a field has been set.

### GetMspId

`func (o *MspAddDeviceListOpenApiVO) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *MspAddDeviceListOpenApiVO) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *MspAddDeviceListOpenApiVO) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *MspAddDeviceListOpenApiVO) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetVerifyCode

`func (o *MspAddDeviceListOpenApiVO) GetVerifyCode() string`

GetVerifyCode returns the VerifyCode field if non-nil, zero value otherwise.

### GetVerifyCodeOk

`func (o *MspAddDeviceListOpenApiVO) GetVerifyCodeOk() (*string, bool)`

GetVerifyCodeOk returns a tuple with the VerifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCode

`func (o *MspAddDeviceListOpenApiVO) SetVerifyCode(v string)`

SetVerifyCode sets VerifyCode field to given value.

### HasVerifyCode

`func (o *MspAddDeviceListOpenApiVO) HasVerifyCode() bool`

HasVerifyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


