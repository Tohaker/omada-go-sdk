# RFPlanningDeployConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomSetting** | Pointer to [**CustomSettingVO**](CustomSettingVO.md) |  | [optional] 
**Mode** | **int32** | 0: Auto configuration. 1: Custom configuration. | 

## Methods

### NewRFPlanningDeployConfig

`func NewRFPlanningDeployConfig(mode int32, ) *RFPlanningDeployConfig`

NewRFPlanningDeployConfig instantiates a new RFPlanningDeployConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRFPlanningDeployConfigWithDefaults

`func NewRFPlanningDeployConfigWithDefaults() *RFPlanningDeployConfig`

NewRFPlanningDeployConfigWithDefaults instantiates a new RFPlanningDeployConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomSetting

`func (o *RFPlanningDeployConfig) GetCustomSetting() CustomSettingVO`

GetCustomSetting returns the CustomSetting field if non-nil, zero value otherwise.

### GetCustomSettingOk

`func (o *RFPlanningDeployConfig) GetCustomSettingOk() (*CustomSettingVO, bool)`

GetCustomSettingOk returns a tuple with the CustomSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSetting

`func (o *RFPlanningDeployConfig) SetCustomSetting(v CustomSettingVO)`

SetCustomSetting sets CustomSetting field to given value.

### HasCustomSetting

`func (o *RFPlanningDeployConfig) HasCustomSetting() bool`

HasCustomSetting returns a boolean if a field has been set.

### GetMode

`func (o *RFPlanningDeployConfig) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RFPlanningDeployConfig) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RFPlanningDeployConfig) SetMode(v int32)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


