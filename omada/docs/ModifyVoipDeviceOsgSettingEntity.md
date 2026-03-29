# ModifyVoipDeviceOsgSettingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipConfirm** | **bool** | skipConfirm indicates whether to skip the query of \&quot;#&#39;s conflict\&quot; of voip device settings. false: Not to skip the query. true: Skip the query and the device will give priority to the \&quot;end with #\&quot; configuration.  | 
**TelephoneNumberAdvancedSetting** | Pointer to [**TelephoneNumberAdvancedSettingOsgOpenApiVO**](TelephoneNumberAdvancedSettingOsgOpenApiVO.md) |  | [optional] 
**VoipDeviceOsgConfiguration** | Pointer to [**VoipDeviceOsgConfigurationOpenApiVO**](VoipDeviceOsgConfigurationOpenApiVO.md) |  | [optional] 

## Methods

### NewModifyVoipDeviceOsgSettingEntity

`func NewModifyVoipDeviceOsgSettingEntity(skipConfirm bool, ) *ModifyVoipDeviceOsgSettingEntity`

NewModifyVoipDeviceOsgSettingEntity instantiates a new ModifyVoipDeviceOsgSettingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyVoipDeviceOsgSettingEntityWithDefaults

`func NewModifyVoipDeviceOsgSettingEntityWithDefaults() *ModifyVoipDeviceOsgSettingEntity`

NewModifyVoipDeviceOsgSettingEntityWithDefaults instantiates a new ModifyVoipDeviceOsgSettingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipConfirm

`func (o *ModifyVoipDeviceOsgSettingEntity) GetSkipConfirm() bool`

GetSkipConfirm returns the SkipConfirm field if non-nil, zero value otherwise.

### GetSkipConfirmOk

`func (o *ModifyVoipDeviceOsgSettingEntity) GetSkipConfirmOk() (*bool, bool)`

GetSkipConfirmOk returns a tuple with the SkipConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirm

`func (o *ModifyVoipDeviceOsgSettingEntity) SetSkipConfirm(v bool)`

SetSkipConfirm sets SkipConfirm field to given value.


### GetTelephoneNumberAdvancedSetting

`func (o *ModifyVoipDeviceOsgSettingEntity) GetTelephoneNumberAdvancedSetting() TelephoneNumberAdvancedSettingOsgOpenApiVO`

GetTelephoneNumberAdvancedSetting returns the TelephoneNumberAdvancedSetting field if non-nil, zero value otherwise.

### GetTelephoneNumberAdvancedSettingOk

`func (o *ModifyVoipDeviceOsgSettingEntity) GetTelephoneNumberAdvancedSettingOk() (*TelephoneNumberAdvancedSettingOsgOpenApiVO, bool)`

GetTelephoneNumberAdvancedSettingOk returns a tuple with the TelephoneNumberAdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumberAdvancedSetting

`func (o *ModifyVoipDeviceOsgSettingEntity) SetTelephoneNumberAdvancedSetting(v TelephoneNumberAdvancedSettingOsgOpenApiVO)`

SetTelephoneNumberAdvancedSetting sets TelephoneNumberAdvancedSetting field to given value.

### HasTelephoneNumberAdvancedSetting

`func (o *ModifyVoipDeviceOsgSettingEntity) HasTelephoneNumberAdvancedSetting() bool`

HasTelephoneNumberAdvancedSetting returns a boolean if a field has been set.

### GetVoipDeviceOsgConfiguration

`func (o *ModifyVoipDeviceOsgSettingEntity) GetVoipDeviceOsgConfiguration() VoipDeviceOsgConfigurationOpenApiVO`

GetVoipDeviceOsgConfiguration returns the VoipDeviceOsgConfiguration field if non-nil, zero value otherwise.

### GetVoipDeviceOsgConfigurationOk

`func (o *ModifyVoipDeviceOsgSettingEntity) GetVoipDeviceOsgConfigurationOk() (*VoipDeviceOsgConfigurationOpenApiVO, bool)`

GetVoipDeviceOsgConfigurationOk returns a tuple with the VoipDeviceOsgConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipDeviceOsgConfiguration

`func (o *ModifyVoipDeviceOsgSettingEntity) SetVoipDeviceOsgConfiguration(v VoipDeviceOsgConfigurationOpenApiVO)`

SetVoipDeviceOsgConfiguration sets VoipDeviceOsgConfiguration field to given value.

### HasVoipDeviceOsgConfiguration

`func (o *ModifyVoipDeviceOsgSettingEntity) HasVoipDeviceOsgConfiguration() bool`

HasVoipDeviceOsgConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


