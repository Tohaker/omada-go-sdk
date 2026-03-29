# ModifyProviderProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileName** | Pointer to **string** | Provider profile name | [optional] 
**ProviderSettings** | Pointer to [**ModifyProviderSettingOpenApiVO**](ModifyProviderSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewModifyProviderProfileEntity

`func NewModifyProviderProfileEntity() *ModifyProviderProfileEntity`

NewModifyProviderProfileEntity instantiates a new ModifyProviderProfileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyProviderProfileEntityWithDefaults

`func NewModifyProviderProfileEntityWithDefaults() *ModifyProviderProfileEntity`

NewModifyProviderProfileEntityWithDefaults instantiates a new ModifyProviderProfileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileName

`func (o *ModifyProviderProfileEntity) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *ModifyProviderProfileEntity) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *ModifyProviderProfileEntity) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *ModifyProviderProfileEntity) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ModifyProviderProfileEntity) GetProviderSettings() ModifyProviderSettingOpenApiVO`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ModifyProviderProfileEntity) GetProviderSettingsOk() (*ModifyProviderSettingOpenApiVO, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ModifyProviderProfileEntity) SetProviderSettings(v ModifyProviderSettingOpenApiVO)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ModifyProviderProfileEntity) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


