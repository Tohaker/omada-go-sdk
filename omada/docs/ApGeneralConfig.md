# ApGeneralConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableHwReset** | Pointer to **bool** | When this value is [true], the device&#39;s hardware reset is invalid.(Only some specific devices support this feature) | [optional] 
**GpsEnable** | Pointer to **bool** | gps Enable only for ap which support gps function. | [optional] 
**LedSetting** | Pointer to **int32** | Led setting should be a value as follows: 0:off; 1:on; 2:Use Site Settings. | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Name** | Pointer to **string** | Device name should contain 1 to 128 characters. | [optional] 
**RememberDevice** | Pointer to **int32** | Parameter [rememberDevice] should be a value as follows: 0:off; 1:on; 2:Use Site Settings. | [optional] 
**RemoteReset** | Pointer to **bool** | Whether the remote reset function is enabled. | [optional] 
**RssiLedSettings** | Pointer to [**[]ApRssiLedSettingVO**](ApRssiLedSettingVO.md) | Rssi Led Setting. | [optional] 
**TagIds** | Pointer to **[]string** | Tag IDs. | [optional] 
**WifiControlEnable** | Pointer to **bool** | Whether the wifi Control function is enabled. | [optional] 

## Methods

### NewApGeneralConfig

`func NewApGeneralConfig() *ApGeneralConfig`

NewApGeneralConfig instantiates a new ApGeneralConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApGeneralConfigWithDefaults

`func NewApGeneralConfigWithDefaults() *ApGeneralConfig`

NewApGeneralConfigWithDefaults instantiates a new ApGeneralConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableHwReset

`func (o *ApGeneralConfig) GetDisableHwReset() bool`

GetDisableHwReset returns the DisableHwReset field if non-nil, zero value otherwise.

### GetDisableHwResetOk

`func (o *ApGeneralConfig) GetDisableHwResetOk() (*bool, bool)`

GetDisableHwResetOk returns a tuple with the DisableHwReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHwReset

`func (o *ApGeneralConfig) SetDisableHwReset(v bool)`

SetDisableHwReset sets DisableHwReset field to given value.

### HasDisableHwReset

`func (o *ApGeneralConfig) HasDisableHwReset() bool`

HasDisableHwReset returns a boolean if a field has been set.

### GetGpsEnable

`func (o *ApGeneralConfig) GetGpsEnable() bool`

GetGpsEnable returns the GpsEnable field if non-nil, zero value otherwise.

### GetGpsEnableOk

`func (o *ApGeneralConfig) GetGpsEnableOk() (*bool, bool)`

GetGpsEnableOk returns a tuple with the GpsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsEnable

`func (o *ApGeneralConfig) SetGpsEnable(v bool)`

SetGpsEnable sets GpsEnable field to given value.

### HasGpsEnable

`func (o *ApGeneralConfig) HasGpsEnable() bool`

HasGpsEnable returns a boolean if a field has been set.

### GetLedSetting

`func (o *ApGeneralConfig) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *ApGeneralConfig) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *ApGeneralConfig) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.

### HasLedSetting

`func (o *ApGeneralConfig) HasLedSetting() bool`

HasLedSetting returns a boolean if a field has been set.

### GetLocation

`func (o *ApGeneralConfig) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApGeneralConfig) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApGeneralConfig) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApGeneralConfig) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ApGeneralConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApGeneralConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApGeneralConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApGeneralConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRememberDevice

`func (o *ApGeneralConfig) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *ApGeneralConfig) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *ApGeneralConfig) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *ApGeneralConfig) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetRemoteReset

`func (o *ApGeneralConfig) GetRemoteReset() bool`

GetRemoteReset returns the RemoteReset field if non-nil, zero value otherwise.

### GetRemoteResetOk

`func (o *ApGeneralConfig) GetRemoteResetOk() (*bool, bool)`

GetRemoteResetOk returns a tuple with the RemoteReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteReset

`func (o *ApGeneralConfig) SetRemoteReset(v bool)`

SetRemoteReset sets RemoteReset field to given value.

### HasRemoteReset

`func (o *ApGeneralConfig) HasRemoteReset() bool`

HasRemoteReset returns a boolean if a field has been set.

### GetRssiLedSettings

`func (o *ApGeneralConfig) GetRssiLedSettings() []ApRssiLedSettingVO`

GetRssiLedSettings returns the RssiLedSettings field if non-nil, zero value otherwise.

### GetRssiLedSettingsOk

`func (o *ApGeneralConfig) GetRssiLedSettingsOk() (*[]ApRssiLedSettingVO, bool)`

GetRssiLedSettingsOk returns a tuple with the RssiLedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiLedSettings

`func (o *ApGeneralConfig) SetRssiLedSettings(v []ApRssiLedSettingVO)`

SetRssiLedSettings sets RssiLedSettings field to given value.

### HasRssiLedSettings

`func (o *ApGeneralConfig) HasRssiLedSettings() bool`

HasRssiLedSettings returns a boolean if a field has been set.

### GetTagIds

`func (o *ApGeneralConfig) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ApGeneralConfig) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ApGeneralConfig) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ApGeneralConfig) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetWifiControlEnable

`func (o *ApGeneralConfig) GetWifiControlEnable() bool`

GetWifiControlEnable returns the WifiControlEnable field if non-nil, zero value otherwise.

### GetWifiControlEnableOk

`func (o *ApGeneralConfig) GetWifiControlEnableOk() (*bool, bool)`

GetWifiControlEnableOk returns a tuple with the WifiControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiControlEnable

`func (o *ApGeneralConfig) SetWifiControlEnable(v bool)`

SetWifiControlEnable sets WifiControlEnable field to given value.

### HasWifiControlEnable

`func (o *ApGeneralConfig) HasWifiControlEnable() bool`

HasWifiControlEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


