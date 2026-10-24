# RFPlanningConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSetting** | Pointer to [**PlanningAdvancedSettingVO**](PlanningAdvancedSettingVO.md) |  | [optional] 
**BandDeployEnable** | Pointer to **bool** | Whether to enable band deployment. | [optional] 
**ChannelDeployEnable** | Pointer to **bool** | Whether to enable channel deployment. | [optional] 
**ChannelWidthDeployEnable** | Pointer to **bool** | Whether to enable channel width deployment. | [optional] 
**ExcludeAps** | Pointer to [**[]ExcludeApVO**](ExcludeApVO.md) | Info of Excluded APs. | [optional] 
**Mode** | Pointer to **int32** | 0: by WLAN Optimization schedule. 1: by one-click WLAN Optimization. | [optional] 
**PowerAdjustEnable** | Pointer to **bool** | Whether to enable power adjustment. | [optional] 
**Resource** | Pointer to **int32** | Site setting resource. 0: New created. 1: From template. 2: Override. | [optional] 

## Methods

### NewRFPlanningConfig

`func NewRFPlanningConfig() *RFPlanningConfig`

NewRFPlanningConfig instantiates a new RFPlanningConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRFPlanningConfigWithDefaults

`func NewRFPlanningConfigWithDefaults() *RFPlanningConfig`

NewRFPlanningConfigWithDefaults instantiates a new RFPlanningConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSetting

`func (o *RFPlanningConfig) GetAdvancedSetting() PlanningAdvancedSettingVO`

GetAdvancedSetting returns the AdvancedSetting field if non-nil, zero value otherwise.

### GetAdvancedSettingOk

`func (o *RFPlanningConfig) GetAdvancedSettingOk() (*PlanningAdvancedSettingVO, bool)`

GetAdvancedSettingOk returns a tuple with the AdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSetting

`func (o *RFPlanningConfig) SetAdvancedSetting(v PlanningAdvancedSettingVO)`

SetAdvancedSetting sets AdvancedSetting field to given value.

### HasAdvancedSetting

`func (o *RFPlanningConfig) HasAdvancedSetting() bool`

HasAdvancedSetting returns a boolean if a field has been set.

### GetBandDeployEnable

`func (o *RFPlanningConfig) GetBandDeployEnable() bool`

GetBandDeployEnable returns the BandDeployEnable field if non-nil, zero value otherwise.

### GetBandDeployEnableOk

`func (o *RFPlanningConfig) GetBandDeployEnableOk() (*bool, bool)`

GetBandDeployEnableOk returns a tuple with the BandDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandDeployEnable

`func (o *RFPlanningConfig) SetBandDeployEnable(v bool)`

SetBandDeployEnable sets BandDeployEnable field to given value.

### HasBandDeployEnable

`func (o *RFPlanningConfig) HasBandDeployEnable() bool`

HasBandDeployEnable returns a boolean if a field has been set.

### GetChannelDeployEnable

`func (o *RFPlanningConfig) GetChannelDeployEnable() bool`

GetChannelDeployEnable returns the ChannelDeployEnable field if non-nil, zero value otherwise.

### GetChannelDeployEnableOk

`func (o *RFPlanningConfig) GetChannelDeployEnableOk() (*bool, bool)`

GetChannelDeployEnableOk returns a tuple with the ChannelDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDeployEnable

`func (o *RFPlanningConfig) SetChannelDeployEnable(v bool)`

SetChannelDeployEnable sets ChannelDeployEnable field to given value.

### HasChannelDeployEnable

`func (o *RFPlanningConfig) HasChannelDeployEnable() bool`

HasChannelDeployEnable returns a boolean if a field has been set.

### GetChannelWidthDeployEnable

`func (o *RFPlanningConfig) GetChannelWidthDeployEnable() bool`

GetChannelWidthDeployEnable returns the ChannelWidthDeployEnable field if non-nil, zero value otherwise.

### GetChannelWidthDeployEnableOk

`func (o *RFPlanningConfig) GetChannelWidthDeployEnableOk() (*bool, bool)`

GetChannelWidthDeployEnableOk returns a tuple with the ChannelWidthDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidthDeployEnable

`func (o *RFPlanningConfig) SetChannelWidthDeployEnable(v bool)`

SetChannelWidthDeployEnable sets ChannelWidthDeployEnable field to given value.

### HasChannelWidthDeployEnable

`func (o *RFPlanningConfig) HasChannelWidthDeployEnable() bool`

HasChannelWidthDeployEnable returns a boolean if a field has been set.

### GetExcludeAps

`func (o *RFPlanningConfig) GetExcludeAps() []ExcludeApVO`

GetExcludeAps returns the ExcludeAps field if non-nil, zero value otherwise.

### GetExcludeApsOk

`func (o *RFPlanningConfig) GetExcludeApsOk() (*[]ExcludeApVO, bool)`

GetExcludeApsOk returns a tuple with the ExcludeAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAps

`func (o *RFPlanningConfig) SetExcludeAps(v []ExcludeApVO)`

SetExcludeAps sets ExcludeAps field to given value.

### HasExcludeAps

`func (o *RFPlanningConfig) HasExcludeAps() bool`

HasExcludeAps returns a boolean if a field has been set.

### GetMode

`func (o *RFPlanningConfig) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RFPlanningConfig) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RFPlanningConfig) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RFPlanningConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPowerAdjustEnable

`func (o *RFPlanningConfig) GetPowerAdjustEnable() bool`

GetPowerAdjustEnable returns the PowerAdjustEnable field if non-nil, zero value otherwise.

### GetPowerAdjustEnableOk

`func (o *RFPlanningConfig) GetPowerAdjustEnableOk() (*bool, bool)`

GetPowerAdjustEnableOk returns a tuple with the PowerAdjustEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerAdjustEnable

`func (o *RFPlanningConfig) SetPowerAdjustEnable(v bool)`

SetPowerAdjustEnable sets PowerAdjustEnable field to given value.

### HasPowerAdjustEnable

`func (o *RFPlanningConfig) HasPowerAdjustEnable() bool`

HasPowerAdjustEnable returns a boolean if a field has been set.

### GetResource

`func (o *RFPlanningConfig) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RFPlanningConfig) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RFPlanningConfig) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RFPlanningConfig) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


