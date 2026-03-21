# PlanningHistoryDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandDeployEnable** | Pointer to **bool** | Whether to enable band deployment. | [optional] 
**ChannelDeployEnable** | Pointer to **bool** | Whether to enable channel deployment. | [optional] 
**ChannelWidthDeployEnable** | Pointer to **bool** | Whether to enable channel width deployment. | [optional] 
**CurrentConfig** | Pointer to **int32** | 0: This entry does not correspond to the current configuration. 1: The recommended values are the current configuration. 2: The previous/origin values are the current configuration. | [optional] 
**Data** | Pointer to [**[]ApPlanningHistoryDetailVO**](ApPlanningHistoryDetailVO.md) | Optimization history details. | [optional] 
**FailDevices** | Pointer to **int32** | The number of devices which have been optimized unsuccessfully. | [optional] 
**PowerAdjustEnable** | Pointer to **bool** | Whether to enable power adjustment. | [optional] 
**SuccessDevices** | Pointer to **int32** | The number of devices which have been successfully optimized. | [optional] 
**Time** | Pointer to **int64** | Timestamp corresponding to the start of optimization. | [optional] 

## Methods

### NewPlanningHistoryDetail

`func NewPlanningHistoryDetail() *PlanningHistoryDetail`

NewPlanningHistoryDetail instantiates a new PlanningHistoryDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanningHistoryDetailWithDefaults

`func NewPlanningHistoryDetailWithDefaults() *PlanningHistoryDetail`

NewPlanningHistoryDetailWithDefaults instantiates a new PlanningHistoryDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandDeployEnable

`func (o *PlanningHistoryDetail) GetBandDeployEnable() bool`

GetBandDeployEnable returns the BandDeployEnable field if non-nil, zero value otherwise.

### GetBandDeployEnableOk

`func (o *PlanningHistoryDetail) GetBandDeployEnableOk() (*bool, bool)`

GetBandDeployEnableOk returns a tuple with the BandDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandDeployEnable

`func (o *PlanningHistoryDetail) SetBandDeployEnable(v bool)`

SetBandDeployEnable sets BandDeployEnable field to given value.

### HasBandDeployEnable

`func (o *PlanningHistoryDetail) HasBandDeployEnable() bool`

HasBandDeployEnable returns a boolean if a field has been set.

### GetChannelDeployEnable

`func (o *PlanningHistoryDetail) GetChannelDeployEnable() bool`

GetChannelDeployEnable returns the ChannelDeployEnable field if non-nil, zero value otherwise.

### GetChannelDeployEnableOk

`func (o *PlanningHistoryDetail) GetChannelDeployEnableOk() (*bool, bool)`

GetChannelDeployEnableOk returns a tuple with the ChannelDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDeployEnable

`func (o *PlanningHistoryDetail) SetChannelDeployEnable(v bool)`

SetChannelDeployEnable sets ChannelDeployEnable field to given value.

### HasChannelDeployEnable

`func (o *PlanningHistoryDetail) HasChannelDeployEnable() bool`

HasChannelDeployEnable returns a boolean if a field has been set.

### GetChannelWidthDeployEnable

`func (o *PlanningHistoryDetail) GetChannelWidthDeployEnable() bool`

GetChannelWidthDeployEnable returns the ChannelWidthDeployEnable field if non-nil, zero value otherwise.

### GetChannelWidthDeployEnableOk

`func (o *PlanningHistoryDetail) GetChannelWidthDeployEnableOk() (*bool, bool)`

GetChannelWidthDeployEnableOk returns a tuple with the ChannelWidthDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidthDeployEnable

`func (o *PlanningHistoryDetail) SetChannelWidthDeployEnable(v bool)`

SetChannelWidthDeployEnable sets ChannelWidthDeployEnable field to given value.

### HasChannelWidthDeployEnable

`func (o *PlanningHistoryDetail) HasChannelWidthDeployEnable() bool`

HasChannelWidthDeployEnable returns a boolean if a field has been set.

### GetCurrentConfig

`func (o *PlanningHistoryDetail) GetCurrentConfig() int32`

GetCurrentConfig returns the CurrentConfig field if non-nil, zero value otherwise.

### GetCurrentConfigOk

`func (o *PlanningHistoryDetail) GetCurrentConfigOk() (*int32, bool)`

GetCurrentConfigOk returns a tuple with the CurrentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfig

`func (o *PlanningHistoryDetail) SetCurrentConfig(v int32)`

SetCurrentConfig sets CurrentConfig field to given value.

### HasCurrentConfig

`func (o *PlanningHistoryDetail) HasCurrentConfig() bool`

HasCurrentConfig returns a boolean if a field has been set.

### GetData

`func (o *PlanningHistoryDetail) GetData() []ApPlanningHistoryDetailVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PlanningHistoryDetail) GetDataOk() (*[]ApPlanningHistoryDetailVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PlanningHistoryDetail) SetData(v []ApPlanningHistoryDetailVO)`

SetData sets Data field to given value.

### HasData

`func (o *PlanningHistoryDetail) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFailDevices

`func (o *PlanningHistoryDetail) GetFailDevices() int32`

GetFailDevices returns the FailDevices field if non-nil, zero value otherwise.

### GetFailDevicesOk

`func (o *PlanningHistoryDetail) GetFailDevicesOk() (*int32, bool)`

GetFailDevicesOk returns a tuple with the FailDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDevices

`func (o *PlanningHistoryDetail) SetFailDevices(v int32)`

SetFailDevices sets FailDevices field to given value.

### HasFailDevices

`func (o *PlanningHistoryDetail) HasFailDevices() bool`

HasFailDevices returns a boolean if a field has been set.

### GetPowerAdjustEnable

`func (o *PlanningHistoryDetail) GetPowerAdjustEnable() bool`

GetPowerAdjustEnable returns the PowerAdjustEnable field if non-nil, zero value otherwise.

### GetPowerAdjustEnableOk

`func (o *PlanningHistoryDetail) GetPowerAdjustEnableOk() (*bool, bool)`

GetPowerAdjustEnableOk returns a tuple with the PowerAdjustEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerAdjustEnable

`func (o *PlanningHistoryDetail) SetPowerAdjustEnable(v bool)`

SetPowerAdjustEnable sets PowerAdjustEnable field to given value.

### HasPowerAdjustEnable

`func (o *PlanningHistoryDetail) HasPowerAdjustEnable() bool`

HasPowerAdjustEnable returns a boolean if a field has been set.

### GetSuccessDevices

`func (o *PlanningHistoryDetail) GetSuccessDevices() int32`

GetSuccessDevices returns the SuccessDevices field if non-nil, zero value otherwise.

### GetSuccessDevicesOk

`func (o *PlanningHistoryDetail) GetSuccessDevicesOk() (*int32, bool)`

GetSuccessDevicesOk returns a tuple with the SuccessDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessDevices

`func (o *PlanningHistoryDetail) SetSuccessDevices(v int32)`

SetSuccessDevices sets SuccessDevices field to given value.

### HasSuccessDevices

`func (o *PlanningHistoryDetail) HasSuccessDevices() bool`

HasSuccessDevices returns a boolean if a field has been set.

### GetTime

`func (o *PlanningHistoryDetail) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PlanningHistoryDetail) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PlanningHistoryDetail) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *PlanningHistoryDetail) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


