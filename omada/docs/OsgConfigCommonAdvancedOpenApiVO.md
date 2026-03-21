# OsgConfigCommonAdvancedOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EchoServer** | Pointer to **string** |  | [optional] 
**HwOffloadEnable** | Pointer to **bool** |  | [optional] 
**LldpEnable** | Pointer to **bool** | Deprecated, filed lldpSetting is recommended | [optional] 
**LldpSetting** | Pointer to **int32** |  | [optional] 
**PoeSettings** | Pointer to [**[]OsgPortPoeVO**](OsgPortPoeVO.md) |  | [optional] 

## Methods

### NewOsgConfigCommonAdvancedOpenApiVO

`func NewOsgConfigCommonAdvancedOpenApiVO() *OsgConfigCommonAdvancedOpenApiVO`

NewOsgConfigCommonAdvancedOpenApiVO instantiates a new OsgConfigCommonAdvancedOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgConfigCommonAdvancedOpenApiVOWithDefaults

`func NewOsgConfigCommonAdvancedOpenApiVOWithDefaults() *OsgConfigCommonAdvancedOpenApiVO`

NewOsgConfigCommonAdvancedOpenApiVOWithDefaults instantiates a new OsgConfigCommonAdvancedOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEchoServer

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetEchoServer() string`

GetEchoServer returns the EchoServer field if non-nil, zero value otherwise.

### GetEchoServerOk

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetEchoServerOk() (*string, bool)`

GetEchoServerOk returns a tuple with the EchoServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEchoServer

`func (o *OsgConfigCommonAdvancedOpenApiVO) SetEchoServer(v string)`

SetEchoServer sets EchoServer field to given value.

### HasEchoServer

`func (o *OsgConfigCommonAdvancedOpenApiVO) HasEchoServer() bool`

HasEchoServer returns a boolean if a field has been set.

### GetHwOffloadEnable

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetHwOffloadEnable() bool`

GetHwOffloadEnable returns the HwOffloadEnable field if non-nil, zero value otherwise.

### GetHwOffloadEnableOk

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetHwOffloadEnableOk() (*bool, bool)`

GetHwOffloadEnableOk returns a tuple with the HwOffloadEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwOffloadEnable

`func (o *OsgConfigCommonAdvancedOpenApiVO) SetHwOffloadEnable(v bool)`

SetHwOffloadEnable sets HwOffloadEnable field to given value.

### HasHwOffloadEnable

`func (o *OsgConfigCommonAdvancedOpenApiVO) HasHwOffloadEnable() bool`

HasHwOffloadEnable returns a boolean if a field has been set.

### GetLldpEnable

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetLldpEnable() bool`

GetLldpEnable returns the LldpEnable field if non-nil, zero value otherwise.

### GetLldpEnableOk

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetLldpEnableOk() (*bool, bool)`

GetLldpEnableOk returns a tuple with the LldpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnable

`func (o *OsgConfigCommonAdvancedOpenApiVO) SetLldpEnable(v bool)`

SetLldpEnable sets LldpEnable field to given value.

### HasLldpEnable

`func (o *OsgConfigCommonAdvancedOpenApiVO) HasLldpEnable() bool`

HasLldpEnable returns a boolean if a field has been set.

### GetLldpSetting

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetLldpSetting() int32`

GetLldpSetting returns the LldpSetting field if non-nil, zero value otherwise.

### GetLldpSettingOk

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetLldpSettingOk() (*int32, bool)`

GetLldpSettingOk returns a tuple with the LldpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSetting

`func (o *OsgConfigCommonAdvancedOpenApiVO) SetLldpSetting(v int32)`

SetLldpSetting sets LldpSetting field to given value.

### HasLldpSetting

`func (o *OsgConfigCommonAdvancedOpenApiVO) HasLldpSetting() bool`

HasLldpSetting returns a boolean if a field has been set.

### GetPoeSettings

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetPoeSettings() []OsgPortPoeVO`

GetPoeSettings returns the PoeSettings field if non-nil, zero value otherwise.

### GetPoeSettingsOk

`func (o *OsgConfigCommonAdvancedOpenApiVO) GetPoeSettingsOk() (*[]OsgPortPoeVO, bool)`

GetPoeSettingsOk returns a tuple with the PoeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeSettings

`func (o *OsgConfigCommonAdvancedOpenApiVO) SetPoeSettings(v []OsgPortPoeVO)`

SetPoeSettings sets PoeSettings field to given value.

### HasPoeSettings

`func (o *OsgConfigCommonAdvancedOpenApiVO) HasPoeSettings() bool`

HasPoeSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


