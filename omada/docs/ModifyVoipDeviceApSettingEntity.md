# ModifyVoipDeviceApSettingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipConfirm** | **bool** | skipConfirm indicates whether to skip the query of \&quot;#&#39;s conflict\&quot; of voip device settings. false: Not to skip the query. true: Skip the query and the device will give priority to the \&quot;end with #\&quot; configuration.  | 
**TelephoneNumberAdvancedSetting** | Pointer to [**TelephoneNumberAdvancedSettingApOpenApiVO**](TelephoneNumberAdvancedSettingApOpenApiVO.md) |  | [optional] 
**VoipDeviceApConfiguration** | Pointer to [**VoipDeviceApConfigurationOpenApiVO**](VoipDeviceApConfigurationOpenApiVO.md) |  | [optional] 

## Methods

### NewModifyVoipDeviceApSettingEntity

`func NewModifyVoipDeviceApSettingEntity(skipConfirm bool, ) *ModifyVoipDeviceApSettingEntity`

NewModifyVoipDeviceApSettingEntity instantiates a new ModifyVoipDeviceApSettingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyVoipDeviceApSettingEntityWithDefaults

`func NewModifyVoipDeviceApSettingEntityWithDefaults() *ModifyVoipDeviceApSettingEntity`

NewModifyVoipDeviceApSettingEntityWithDefaults instantiates a new ModifyVoipDeviceApSettingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipConfirm

`func (o *ModifyVoipDeviceApSettingEntity) GetSkipConfirm() bool`

GetSkipConfirm returns the SkipConfirm field if non-nil, zero value otherwise.

### GetSkipConfirmOk

`func (o *ModifyVoipDeviceApSettingEntity) GetSkipConfirmOk() (*bool, bool)`

GetSkipConfirmOk returns a tuple with the SkipConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirm

`func (o *ModifyVoipDeviceApSettingEntity) SetSkipConfirm(v bool)`

SetSkipConfirm sets SkipConfirm field to given value.


### GetTelephoneNumberAdvancedSetting

`func (o *ModifyVoipDeviceApSettingEntity) GetTelephoneNumberAdvancedSetting() TelephoneNumberAdvancedSettingApOpenApiVO`

GetTelephoneNumberAdvancedSetting returns the TelephoneNumberAdvancedSetting field if non-nil, zero value otherwise.

### GetTelephoneNumberAdvancedSettingOk

`func (o *ModifyVoipDeviceApSettingEntity) GetTelephoneNumberAdvancedSettingOk() (*TelephoneNumberAdvancedSettingApOpenApiVO, bool)`

GetTelephoneNumberAdvancedSettingOk returns a tuple with the TelephoneNumberAdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumberAdvancedSetting

`func (o *ModifyVoipDeviceApSettingEntity) SetTelephoneNumberAdvancedSetting(v TelephoneNumberAdvancedSettingApOpenApiVO)`

SetTelephoneNumberAdvancedSetting sets TelephoneNumberAdvancedSetting field to given value.

### HasTelephoneNumberAdvancedSetting

`func (o *ModifyVoipDeviceApSettingEntity) HasTelephoneNumberAdvancedSetting() bool`

HasTelephoneNumberAdvancedSetting returns a boolean if a field has been set.

### GetVoipDeviceApConfiguration

`func (o *ModifyVoipDeviceApSettingEntity) GetVoipDeviceApConfiguration() VoipDeviceApConfigurationOpenApiVO`

GetVoipDeviceApConfiguration returns the VoipDeviceApConfiguration field if non-nil, zero value otherwise.

### GetVoipDeviceApConfigurationOk

`func (o *ModifyVoipDeviceApSettingEntity) GetVoipDeviceApConfigurationOk() (*VoipDeviceApConfigurationOpenApiVO, bool)`

GetVoipDeviceApConfigurationOk returns a tuple with the VoipDeviceApConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipDeviceApConfiguration

`func (o *ModifyVoipDeviceApSettingEntity) SetVoipDeviceApConfiguration(v VoipDeviceApConfigurationOpenApiVO)`

SetVoipDeviceApConfiguration sets VoipDeviceApConfiguration field to given value.

### HasVoipDeviceApConfiguration

`func (o *ModifyVoipDeviceApSettingEntity) HasVoipDeviceApConfiguration() bool`

HasVoipDeviceApConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


