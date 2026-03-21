# BatchModifyVoipDeviceSettingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMacList** | **[]string** | Device MAC list. | 
**SkipConfirm** | **bool** | skipConfirm indicates whether to skip the query of \&quot;#&#39;s conflict\&quot; of voip device settings. false: Not to skip the query. true: Skip the query and the device(s) will give priority to the \&quot;end with #\&quot; configuration.  | 
**TelephoneNumberAdvancedSetting** | Pointer to [**TelephoneNumberAdvancedSettingApOpenApiVO**](TelephoneNumberAdvancedSettingApOpenApiVO.md) |  | [optional] 
**VoipDeviceConfiguration** | Pointer to [**BatchModifyVoipDeviceConfigurationEntity**](BatchModifyVoipDeviceConfigurationEntity.md) |  | [optional] 

## Methods

### NewBatchModifyVoipDeviceSettingEntity

`func NewBatchModifyVoipDeviceSettingEntity(deviceMacList []string, skipConfirm bool, ) *BatchModifyVoipDeviceSettingEntity`

NewBatchModifyVoipDeviceSettingEntity instantiates a new BatchModifyVoipDeviceSettingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchModifyVoipDeviceSettingEntityWithDefaults

`func NewBatchModifyVoipDeviceSettingEntityWithDefaults() *BatchModifyVoipDeviceSettingEntity`

NewBatchModifyVoipDeviceSettingEntityWithDefaults instantiates a new BatchModifyVoipDeviceSettingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMacList

`func (o *BatchModifyVoipDeviceSettingEntity) GetDeviceMacList() []string`

GetDeviceMacList returns the DeviceMacList field if non-nil, zero value otherwise.

### GetDeviceMacListOk

`func (o *BatchModifyVoipDeviceSettingEntity) GetDeviceMacListOk() (*[]string, bool)`

GetDeviceMacListOk returns a tuple with the DeviceMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacList

`func (o *BatchModifyVoipDeviceSettingEntity) SetDeviceMacList(v []string)`

SetDeviceMacList sets DeviceMacList field to given value.


### GetSkipConfirm

`func (o *BatchModifyVoipDeviceSettingEntity) GetSkipConfirm() bool`

GetSkipConfirm returns the SkipConfirm field if non-nil, zero value otherwise.

### GetSkipConfirmOk

`func (o *BatchModifyVoipDeviceSettingEntity) GetSkipConfirmOk() (*bool, bool)`

GetSkipConfirmOk returns a tuple with the SkipConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirm

`func (o *BatchModifyVoipDeviceSettingEntity) SetSkipConfirm(v bool)`

SetSkipConfirm sets SkipConfirm field to given value.


### GetTelephoneNumberAdvancedSetting

`func (o *BatchModifyVoipDeviceSettingEntity) GetTelephoneNumberAdvancedSetting() TelephoneNumberAdvancedSettingApOpenApiVO`

GetTelephoneNumberAdvancedSetting returns the TelephoneNumberAdvancedSetting field if non-nil, zero value otherwise.

### GetTelephoneNumberAdvancedSettingOk

`func (o *BatchModifyVoipDeviceSettingEntity) GetTelephoneNumberAdvancedSettingOk() (*TelephoneNumberAdvancedSettingApOpenApiVO, bool)`

GetTelephoneNumberAdvancedSettingOk returns a tuple with the TelephoneNumberAdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumberAdvancedSetting

`func (o *BatchModifyVoipDeviceSettingEntity) SetTelephoneNumberAdvancedSetting(v TelephoneNumberAdvancedSettingApOpenApiVO)`

SetTelephoneNumberAdvancedSetting sets TelephoneNumberAdvancedSetting field to given value.

### HasTelephoneNumberAdvancedSetting

`func (o *BatchModifyVoipDeviceSettingEntity) HasTelephoneNumberAdvancedSetting() bool`

HasTelephoneNumberAdvancedSetting returns a boolean if a field has been set.

### GetVoipDeviceConfiguration

`func (o *BatchModifyVoipDeviceSettingEntity) GetVoipDeviceConfiguration() BatchModifyVoipDeviceConfigurationEntity`

GetVoipDeviceConfiguration returns the VoipDeviceConfiguration field if non-nil, zero value otherwise.

### GetVoipDeviceConfigurationOk

`func (o *BatchModifyVoipDeviceSettingEntity) GetVoipDeviceConfigurationOk() (*BatchModifyVoipDeviceConfigurationEntity, bool)`

GetVoipDeviceConfigurationOk returns a tuple with the VoipDeviceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipDeviceConfiguration

`func (o *BatchModifyVoipDeviceSettingEntity) SetVoipDeviceConfiguration(v BatchModifyVoipDeviceConfigurationEntity)`

SetVoipDeviceConfiguration sets VoipDeviceConfiguration field to given value.

### HasVoipDeviceConfiguration

`func (o *BatchModifyVoipDeviceSettingEntity) HasVoipDeviceConfiguration() bool`

HasVoipDeviceConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


