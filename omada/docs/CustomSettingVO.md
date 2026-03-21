# CustomSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSetting** | Pointer to [**PlanningAdvancedSettingVO**](PlanningAdvancedSettingVO.md) |  | [optional] 
**BandDeployEnable** | **bool** | Whether to enable band deployment. | 
**ChannelDeployEnable** | **bool** | Whether to enable channel deployment. | 
**ChannelWidthDeployEnable** | **bool** | Whether to enable channel width deployment. | 
**PowerAdjustEnable** | **bool** | Whether to enable power adjustment. | 

## Methods

### NewCustomSettingVO

`func NewCustomSettingVO(bandDeployEnable bool, channelDeployEnable bool, channelWidthDeployEnable bool, powerAdjustEnable bool, ) *CustomSettingVO`

NewCustomSettingVO instantiates a new CustomSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomSettingVOWithDefaults

`func NewCustomSettingVOWithDefaults() *CustomSettingVO`

NewCustomSettingVOWithDefaults instantiates a new CustomSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSetting

`func (o *CustomSettingVO) GetAdvancedSetting() PlanningAdvancedSettingVO`

GetAdvancedSetting returns the AdvancedSetting field if non-nil, zero value otherwise.

### GetAdvancedSettingOk

`func (o *CustomSettingVO) GetAdvancedSettingOk() (*PlanningAdvancedSettingVO, bool)`

GetAdvancedSettingOk returns a tuple with the AdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSetting

`func (o *CustomSettingVO) SetAdvancedSetting(v PlanningAdvancedSettingVO)`

SetAdvancedSetting sets AdvancedSetting field to given value.

### HasAdvancedSetting

`func (o *CustomSettingVO) HasAdvancedSetting() bool`

HasAdvancedSetting returns a boolean if a field has been set.

### GetBandDeployEnable

`func (o *CustomSettingVO) GetBandDeployEnable() bool`

GetBandDeployEnable returns the BandDeployEnable field if non-nil, zero value otherwise.

### GetBandDeployEnableOk

`func (o *CustomSettingVO) GetBandDeployEnableOk() (*bool, bool)`

GetBandDeployEnableOk returns a tuple with the BandDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandDeployEnable

`func (o *CustomSettingVO) SetBandDeployEnable(v bool)`

SetBandDeployEnable sets BandDeployEnable field to given value.


### GetChannelDeployEnable

`func (o *CustomSettingVO) GetChannelDeployEnable() bool`

GetChannelDeployEnable returns the ChannelDeployEnable field if non-nil, zero value otherwise.

### GetChannelDeployEnableOk

`func (o *CustomSettingVO) GetChannelDeployEnableOk() (*bool, bool)`

GetChannelDeployEnableOk returns a tuple with the ChannelDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDeployEnable

`func (o *CustomSettingVO) SetChannelDeployEnable(v bool)`

SetChannelDeployEnable sets ChannelDeployEnable field to given value.


### GetChannelWidthDeployEnable

`func (o *CustomSettingVO) GetChannelWidthDeployEnable() bool`

GetChannelWidthDeployEnable returns the ChannelWidthDeployEnable field if non-nil, zero value otherwise.

### GetChannelWidthDeployEnableOk

`func (o *CustomSettingVO) GetChannelWidthDeployEnableOk() (*bool, bool)`

GetChannelWidthDeployEnableOk returns a tuple with the ChannelWidthDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidthDeployEnable

`func (o *CustomSettingVO) SetChannelWidthDeployEnable(v bool)`

SetChannelWidthDeployEnable sets ChannelWidthDeployEnable field to given value.


### GetPowerAdjustEnable

`func (o *CustomSettingVO) GetPowerAdjustEnable() bool`

GetPowerAdjustEnable returns the PowerAdjustEnable field if non-nil, zero value otherwise.

### GetPowerAdjustEnableOk

`func (o *CustomSettingVO) GetPowerAdjustEnableOk() (*bool, bool)`

GetPowerAdjustEnableOk returns a tuple with the PowerAdjustEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerAdjustEnable

`func (o *CustomSettingVO) SetPowerAdjustEnable(v bool)`

SetPowerAdjustEnable sets PowerAdjustEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


