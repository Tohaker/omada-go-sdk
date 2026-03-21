# CreateProviderProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileName** | **string** | Provider profile name | 
**ProviderSettings** | [**ProviderSettingVO**](ProviderSettingVO.md) |  | 

## Methods

### NewCreateProviderProfileEntity

`func NewCreateProviderProfileEntity(profileName string, providerSettings ProviderSettingVO, ) *CreateProviderProfileEntity`

NewCreateProviderProfileEntity instantiates a new CreateProviderProfileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProviderProfileEntityWithDefaults

`func NewCreateProviderProfileEntityWithDefaults() *CreateProviderProfileEntity`

NewCreateProviderProfileEntityWithDefaults instantiates a new CreateProviderProfileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileName

`func (o *CreateProviderProfileEntity) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *CreateProviderProfileEntity) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *CreateProviderProfileEntity) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.


### GetProviderSettings

`func (o *CreateProviderProfileEntity) GetProviderSettings() ProviderSettingVO`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *CreateProviderProfileEntity) GetProviderSettingsOk() (*ProviderSettingVO, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *CreateProviderProfileEntity) SetProviderSettings(v ProviderSettingVO)`

SetProviderSettings sets ProviderSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


