# PlanningAdvancedSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude5gCh** | Pointer to [**[]ExcludeChannelVO**](ExcludeChannelVO.md) | Excluded 5 GHz channel list. | [optional] 
**Exclude5gChEn** | Pointer to **bool** | Whether to enable the Excluded 5GHz Channels function. | [optional] 
**PowerMode** | Pointer to **int32** | Power mode. 0: Auto. 1: Custom. | [optional] 
**PowerRange** | Pointer to [**PowerRangeVO**](PowerRangeVO.md) |  | [optional] 
**PowerThreshold** | Pointer to [**PowerThresholdVO**](PowerThresholdVO.md) |  | [optional] 
**WidthRange** | Pointer to [**WidthRangeVO**](WidthRangeVO.md) |  | [optional] 
**WidthSelectEn** | Pointer to **bool** | Whether to enable channel width selection. | [optional] 

## Methods

### NewPlanningAdvancedSettingVO

`func NewPlanningAdvancedSettingVO() *PlanningAdvancedSettingVO`

NewPlanningAdvancedSettingVO instantiates a new PlanningAdvancedSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanningAdvancedSettingVOWithDefaults

`func NewPlanningAdvancedSettingVOWithDefaults() *PlanningAdvancedSettingVO`

NewPlanningAdvancedSettingVOWithDefaults instantiates a new PlanningAdvancedSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude5gCh

`func (o *PlanningAdvancedSettingVO) GetExclude5gCh() []ExcludeChannelVO`

GetExclude5gCh returns the Exclude5gCh field if non-nil, zero value otherwise.

### GetExclude5gChOk

`func (o *PlanningAdvancedSettingVO) GetExclude5gChOk() (*[]ExcludeChannelVO, bool)`

GetExclude5gChOk returns a tuple with the Exclude5gCh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude5gCh

`func (o *PlanningAdvancedSettingVO) SetExclude5gCh(v []ExcludeChannelVO)`

SetExclude5gCh sets Exclude5gCh field to given value.

### HasExclude5gCh

`func (o *PlanningAdvancedSettingVO) HasExclude5gCh() bool`

HasExclude5gCh returns a boolean if a field has been set.

### GetExclude5gChEn

`func (o *PlanningAdvancedSettingVO) GetExclude5gChEn() bool`

GetExclude5gChEn returns the Exclude5gChEn field if non-nil, zero value otherwise.

### GetExclude5gChEnOk

`func (o *PlanningAdvancedSettingVO) GetExclude5gChEnOk() (*bool, bool)`

GetExclude5gChEnOk returns a tuple with the Exclude5gChEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude5gChEn

`func (o *PlanningAdvancedSettingVO) SetExclude5gChEn(v bool)`

SetExclude5gChEn sets Exclude5gChEn field to given value.

### HasExclude5gChEn

`func (o *PlanningAdvancedSettingVO) HasExclude5gChEn() bool`

HasExclude5gChEn returns a boolean if a field has been set.

### GetPowerMode

`func (o *PlanningAdvancedSettingVO) GetPowerMode() int32`

GetPowerMode returns the PowerMode field if non-nil, zero value otherwise.

### GetPowerModeOk

`func (o *PlanningAdvancedSettingVO) GetPowerModeOk() (*int32, bool)`

GetPowerModeOk returns a tuple with the PowerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMode

`func (o *PlanningAdvancedSettingVO) SetPowerMode(v int32)`

SetPowerMode sets PowerMode field to given value.

### HasPowerMode

`func (o *PlanningAdvancedSettingVO) HasPowerMode() bool`

HasPowerMode returns a boolean if a field has been set.

### GetPowerRange

`func (o *PlanningAdvancedSettingVO) GetPowerRange() PowerRangeVO`

GetPowerRange returns the PowerRange field if non-nil, zero value otherwise.

### GetPowerRangeOk

`func (o *PlanningAdvancedSettingVO) GetPowerRangeOk() (*PowerRangeVO, bool)`

GetPowerRangeOk returns a tuple with the PowerRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerRange

`func (o *PlanningAdvancedSettingVO) SetPowerRange(v PowerRangeVO)`

SetPowerRange sets PowerRange field to given value.

### HasPowerRange

`func (o *PlanningAdvancedSettingVO) HasPowerRange() bool`

HasPowerRange returns a boolean if a field has been set.

### GetPowerThreshold

`func (o *PlanningAdvancedSettingVO) GetPowerThreshold() PowerThresholdVO`

GetPowerThreshold returns the PowerThreshold field if non-nil, zero value otherwise.

### GetPowerThresholdOk

`func (o *PlanningAdvancedSettingVO) GetPowerThresholdOk() (*PowerThresholdVO, bool)`

GetPowerThresholdOk returns a tuple with the PowerThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerThreshold

`func (o *PlanningAdvancedSettingVO) SetPowerThreshold(v PowerThresholdVO)`

SetPowerThreshold sets PowerThreshold field to given value.

### HasPowerThreshold

`func (o *PlanningAdvancedSettingVO) HasPowerThreshold() bool`

HasPowerThreshold returns a boolean if a field has been set.

### GetWidthRange

`func (o *PlanningAdvancedSettingVO) GetWidthRange() WidthRangeVO`

GetWidthRange returns the WidthRange field if non-nil, zero value otherwise.

### GetWidthRangeOk

`func (o *PlanningAdvancedSettingVO) GetWidthRangeOk() (*WidthRangeVO, bool)`

GetWidthRangeOk returns a tuple with the WidthRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthRange

`func (o *PlanningAdvancedSettingVO) SetWidthRange(v WidthRangeVO)`

SetWidthRange sets WidthRange field to given value.

### HasWidthRange

`func (o *PlanningAdvancedSettingVO) HasWidthRange() bool`

HasWidthRange returns a boolean if a field has been set.

### GetWidthSelectEn

`func (o *PlanningAdvancedSettingVO) GetWidthSelectEn() bool`

GetWidthSelectEn returns the WidthSelectEn field if non-nil, zero value otherwise.

### GetWidthSelectEnOk

`func (o *PlanningAdvancedSettingVO) GetWidthSelectEnOk() (*bool, bool)`

GetWidthSelectEnOk returns a tuple with the WidthSelectEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthSelectEn

`func (o *PlanningAdvancedSettingVO) SetWidthSelectEn(v bool)`

SetWidthSelectEn sets WidthSelectEn field to given value.

### HasWidthSelectEn

`func (o *PlanningAdvancedSettingVO) HasWidthSelectEn() bool`

HasWidthSelectEn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


