# EasyManagedSwitchGeneralConfigDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JumboEnable** | Pointer to **bool** | Parameter [jumboEnable] should be true or false. | [optional] 
**LedSetting** | Pointer to **int32** | LED setting should be a value as follows: 0:off; 1:on; 2:Use Site Settings | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Name** | Pointer to **string** | Device name should contain 1 to 32 characters. | [optional] 
**PowerAlertEnable** | Pointer to **bool** | Power alert status. | [optional] 
**Remember** | Pointer to **bool** | Parameter [remember] should be true or false. | [optional] 
**SupportPowerAlert** | Pointer to **bool** | Parameter [supportPowerAlert] should be true or false. | [optional] 
**TagIds** | Pointer to **[]string** | Tag IDs | [optional] 

## Methods

### NewEasyManagedSwitchGeneralConfigDetail

`func NewEasyManagedSwitchGeneralConfigDetail() *EasyManagedSwitchGeneralConfigDetail`

NewEasyManagedSwitchGeneralConfigDetail instantiates a new EasyManagedSwitchGeneralConfigDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasyManagedSwitchGeneralConfigDetailWithDefaults

`func NewEasyManagedSwitchGeneralConfigDetailWithDefaults() *EasyManagedSwitchGeneralConfigDetail`

NewEasyManagedSwitchGeneralConfigDetailWithDefaults instantiates a new EasyManagedSwitchGeneralConfigDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJumboEnable

`func (o *EasyManagedSwitchGeneralConfigDetail) GetJumboEnable() bool`

GetJumboEnable returns the JumboEnable field if non-nil, zero value otherwise.

### GetJumboEnableOk

`func (o *EasyManagedSwitchGeneralConfigDetail) GetJumboEnableOk() (*bool, bool)`

GetJumboEnableOk returns a tuple with the JumboEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboEnable

`func (o *EasyManagedSwitchGeneralConfigDetail) SetJumboEnable(v bool)`

SetJumboEnable sets JumboEnable field to given value.

### HasJumboEnable

`func (o *EasyManagedSwitchGeneralConfigDetail) HasJumboEnable() bool`

HasJumboEnable returns a boolean if a field has been set.

### GetLedSetting

`func (o *EasyManagedSwitchGeneralConfigDetail) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *EasyManagedSwitchGeneralConfigDetail) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *EasyManagedSwitchGeneralConfigDetail) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.

### HasLedSetting

`func (o *EasyManagedSwitchGeneralConfigDetail) HasLedSetting() bool`

HasLedSetting returns a boolean if a field has been set.

### GetLocation

`func (o *EasyManagedSwitchGeneralConfigDetail) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EasyManagedSwitchGeneralConfigDetail) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EasyManagedSwitchGeneralConfigDetail) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EasyManagedSwitchGeneralConfigDetail) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *EasyManagedSwitchGeneralConfigDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EasyManagedSwitchGeneralConfigDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EasyManagedSwitchGeneralConfigDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EasyManagedSwitchGeneralConfigDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPowerAlertEnable

`func (o *EasyManagedSwitchGeneralConfigDetail) GetPowerAlertEnable() bool`

GetPowerAlertEnable returns the PowerAlertEnable field if non-nil, zero value otherwise.

### GetPowerAlertEnableOk

`func (o *EasyManagedSwitchGeneralConfigDetail) GetPowerAlertEnableOk() (*bool, bool)`

GetPowerAlertEnableOk returns a tuple with the PowerAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerAlertEnable

`func (o *EasyManagedSwitchGeneralConfigDetail) SetPowerAlertEnable(v bool)`

SetPowerAlertEnable sets PowerAlertEnable field to given value.

### HasPowerAlertEnable

`func (o *EasyManagedSwitchGeneralConfigDetail) HasPowerAlertEnable() bool`

HasPowerAlertEnable returns a boolean if a field has been set.

### GetRemember

`func (o *EasyManagedSwitchGeneralConfigDetail) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *EasyManagedSwitchGeneralConfigDetail) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *EasyManagedSwitchGeneralConfigDetail) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *EasyManagedSwitchGeneralConfigDetail) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetSupportPowerAlert

`func (o *EasyManagedSwitchGeneralConfigDetail) GetSupportPowerAlert() bool`

GetSupportPowerAlert returns the SupportPowerAlert field if non-nil, zero value otherwise.

### GetSupportPowerAlertOk

`func (o *EasyManagedSwitchGeneralConfigDetail) GetSupportPowerAlertOk() (*bool, bool)`

GetSupportPowerAlertOk returns a tuple with the SupportPowerAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPowerAlert

`func (o *EasyManagedSwitchGeneralConfigDetail) SetSupportPowerAlert(v bool)`

SetSupportPowerAlert sets SupportPowerAlert field to given value.

### HasSupportPowerAlert

`func (o *EasyManagedSwitchGeneralConfigDetail) HasSupportPowerAlert() bool`

HasSupportPowerAlert returns a boolean if a field has been set.

### GetTagIds

`func (o *EasyManagedSwitchGeneralConfigDetail) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *EasyManagedSwitchGeneralConfigDetail) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *EasyManagedSwitchGeneralConfigDetail) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *EasyManagedSwitchGeneralConfigDetail) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


