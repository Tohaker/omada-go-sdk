# VoipDeviceOsgConfigurationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallBlockingEnable** | Pointer to **bool** | Whether to enable callBlocking. | [optional] 
**CallBlockingProfileId** | Pointer to **string** | The call blocking profile ID of voip device.When callBlockingEnable is true, it can not be null. | [optional] 
**CallBlockingProfileName** | Pointer to **string** | The call blocking profile name of voip device. | [optional] 
**PortSettings** | Pointer to [**[]VoipDevicePortSettingOpenApiVO**](VoipDevicePortSettingOpenApiVO.md) | VOIP device port setting. | [optional] 
**ViaIpv6** | Pointer to **bool** | Whether via IPv6. | [optional] 

## Methods

### NewVoipDeviceOsgConfigurationOpenApiVO

`func NewVoipDeviceOsgConfigurationOpenApiVO() *VoipDeviceOsgConfigurationOpenApiVO`

NewVoipDeviceOsgConfigurationOpenApiVO instantiates a new VoipDeviceOsgConfigurationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipDeviceOsgConfigurationOpenApiVOWithDefaults

`func NewVoipDeviceOsgConfigurationOpenApiVOWithDefaults() *VoipDeviceOsgConfigurationOpenApiVO`

NewVoipDeviceOsgConfigurationOpenApiVOWithDefaults instantiates a new VoipDeviceOsgConfigurationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallBlockingEnable

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetCallBlockingEnable() bool`

GetCallBlockingEnable returns the CallBlockingEnable field if non-nil, zero value otherwise.

### GetCallBlockingEnableOk

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetCallBlockingEnableOk() (*bool, bool)`

GetCallBlockingEnableOk returns a tuple with the CallBlockingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingEnable

`func (o *VoipDeviceOsgConfigurationOpenApiVO) SetCallBlockingEnable(v bool)`

SetCallBlockingEnable sets CallBlockingEnable field to given value.

### HasCallBlockingEnable

`func (o *VoipDeviceOsgConfigurationOpenApiVO) HasCallBlockingEnable() bool`

HasCallBlockingEnable returns a boolean if a field has been set.

### GetCallBlockingProfileId

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetCallBlockingProfileId() string`

GetCallBlockingProfileId returns the CallBlockingProfileId field if non-nil, zero value otherwise.

### GetCallBlockingProfileIdOk

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetCallBlockingProfileIdOk() (*string, bool)`

GetCallBlockingProfileIdOk returns a tuple with the CallBlockingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingProfileId

`func (o *VoipDeviceOsgConfigurationOpenApiVO) SetCallBlockingProfileId(v string)`

SetCallBlockingProfileId sets CallBlockingProfileId field to given value.

### HasCallBlockingProfileId

`func (o *VoipDeviceOsgConfigurationOpenApiVO) HasCallBlockingProfileId() bool`

HasCallBlockingProfileId returns a boolean if a field has been set.

### GetCallBlockingProfileName

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetCallBlockingProfileName() string`

GetCallBlockingProfileName returns the CallBlockingProfileName field if non-nil, zero value otherwise.

### GetCallBlockingProfileNameOk

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetCallBlockingProfileNameOk() (*string, bool)`

GetCallBlockingProfileNameOk returns a tuple with the CallBlockingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingProfileName

`func (o *VoipDeviceOsgConfigurationOpenApiVO) SetCallBlockingProfileName(v string)`

SetCallBlockingProfileName sets CallBlockingProfileName field to given value.

### HasCallBlockingProfileName

`func (o *VoipDeviceOsgConfigurationOpenApiVO) HasCallBlockingProfileName() bool`

HasCallBlockingProfileName returns a boolean if a field has been set.

### GetPortSettings

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetPortSettings() []VoipDevicePortSettingOpenApiVO`

GetPortSettings returns the PortSettings field if non-nil, zero value otherwise.

### GetPortSettingsOk

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetPortSettingsOk() (*[]VoipDevicePortSettingOpenApiVO, bool)`

GetPortSettingsOk returns a tuple with the PortSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSettings

`func (o *VoipDeviceOsgConfigurationOpenApiVO) SetPortSettings(v []VoipDevicePortSettingOpenApiVO)`

SetPortSettings sets PortSettings field to given value.

### HasPortSettings

`func (o *VoipDeviceOsgConfigurationOpenApiVO) HasPortSettings() bool`

HasPortSettings returns a boolean if a field has been set.

### GetViaIpv6

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetViaIpv6() bool`

GetViaIpv6 returns the ViaIpv6 field if non-nil, zero value otherwise.

### GetViaIpv6Ok

`func (o *VoipDeviceOsgConfigurationOpenApiVO) GetViaIpv6Ok() (*bool, bool)`

GetViaIpv6Ok returns a tuple with the ViaIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaIpv6

`func (o *VoipDeviceOsgConfigurationOpenApiVO) SetViaIpv6(v bool)`

SetViaIpv6 sets ViaIpv6 field to given value.

### HasViaIpv6

`func (o *VoipDeviceOsgConfigurationOpenApiVO) HasViaIpv6() bool`

HasViaIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


