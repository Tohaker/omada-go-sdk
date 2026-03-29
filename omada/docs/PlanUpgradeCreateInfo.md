# PlanUpgradeCreateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | The day of the month selected by the user, It should be within the range of 1-31 | [optional] 
**Hour** | Pointer to **int32** | The hour of the day selected by the user, It should be within the range of 0-23 | [optional] 
**Minute** | Pointer to **int32** | The minute in the hour selected by the user, It should be within the range of 0-59 | [optional] 
**ModelList** | [**[]PlanUpgradeSelectedModel**](PlanUpgradeSelectedModel.md) | List of model selected for planned upgrade | 
**MonthOfYear** | Pointer to **int32** | The month of the year selected by the user, It should be within the range of 1-12 | [optional] 
**ScheduleType** | **int32** | The type of upgrade execution time, where 0 represents now and 1 represents the specified time | 
**Sites** | **[]string** | List of site ID selected by the user | 
**Year** | Pointer to **int32** | The year selected by the user | [optional] 

## Methods

### NewPlanUpgradeCreateInfo

`func NewPlanUpgradeCreateInfo(modelList []PlanUpgradeSelectedModel, scheduleType int32, sites []string, ) *PlanUpgradeCreateInfo`

NewPlanUpgradeCreateInfo instantiates a new PlanUpgradeCreateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanUpgradeCreateInfoWithDefaults

`func NewPlanUpgradeCreateInfoWithDefaults() *PlanUpgradeCreateInfo`

NewPlanUpgradeCreateInfoWithDefaults instantiates a new PlanUpgradeCreateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *PlanUpgradeCreateInfo) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *PlanUpgradeCreateInfo) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *PlanUpgradeCreateInfo) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *PlanUpgradeCreateInfo) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetHour

`func (o *PlanUpgradeCreateInfo) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *PlanUpgradeCreateInfo) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *PlanUpgradeCreateInfo) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *PlanUpgradeCreateInfo) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *PlanUpgradeCreateInfo) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *PlanUpgradeCreateInfo) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *PlanUpgradeCreateInfo) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *PlanUpgradeCreateInfo) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetModelList

`func (o *PlanUpgradeCreateInfo) GetModelList() []PlanUpgradeSelectedModel`

GetModelList returns the ModelList field if non-nil, zero value otherwise.

### GetModelListOk

`func (o *PlanUpgradeCreateInfo) GetModelListOk() (*[]PlanUpgradeSelectedModel, bool)`

GetModelListOk returns a tuple with the ModelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelList

`func (o *PlanUpgradeCreateInfo) SetModelList(v []PlanUpgradeSelectedModel)`

SetModelList sets ModelList field to given value.


### GetMonthOfYear

`func (o *PlanUpgradeCreateInfo) GetMonthOfYear() int32`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *PlanUpgradeCreateInfo) GetMonthOfYearOk() (*int32, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *PlanUpgradeCreateInfo) SetMonthOfYear(v int32)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *PlanUpgradeCreateInfo) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### GetScheduleType

`func (o *PlanUpgradeCreateInfo) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *PlanUpgradeCreateInfo) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *PlanUpgradeCreateInfo) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.


### GetSites

`func (o *PlanUpgradeCreateInfo) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PlanUpgradeCreateInfo) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PlanUpgradeCreateInfo) SetSites(v []string)`

SetSites sets Sites field to given value.


### GetYear

`func (o *PlanUpgradeCreateInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *PlanUpgradeCreateInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *PlanUpgradeCreateInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *PlanUpgradeCreateInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


