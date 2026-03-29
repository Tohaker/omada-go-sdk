# FailedCreateModelPlanUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | The day of the month selected by the user, It should be within the range of 1-31 | [optional] 
**Hour** | Pointer to **int32** | The hour of the day selected by the user, It should be within the range of 0-23 | [optional] 
**Minute** | Pointer to **int32** | The minute in the hour selected by the user, It should be within the range of 0-59 | [optional] 
**MonthOfYear** | Pointer to **int32** | The month of the year selected by the user, It should be within the range of 1-12 | [optional] 
**ScheduleType** | **int32** | The type of upgrade execution time, where 0 represents now and 1 represents the specified time | 
**Year** | Pointer to **int32** | The year selected by the user | [optional] 

## Methods

### NewFailedCreateModelPlanUpgradeInfo

`func NewFailedCreateModelPlanUpgradeInfo(scheduleType int32, ) *FailedCreateModelPlanUpgradeInfo`

NewFailedCreateModelPlanUpgradeInfo instantiates a new FailedCreateModelPlanUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedCreateModelPlanUpgradeInfoWithDefaults

`func NewFailedCreateModelPlanUpgradeInfoWithDefaults() *FailedCreateModelPlanUpgradeInfo`

NewFailedCreateModelPlanUpgradeInfoWithDefaults instantiates a new FailedCreateModelPlanUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *FailedCreateModelPlanUpgradeInfo) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *FailedCreateModelPlanUpgradeInfo) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *FailedCreateModelPlanUpgradeInfo) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *FailedCreateModelPlanUpgradeInfo) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetHour

`func (o *FailedCreateModelPlanUpgradeInfo) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *FailedCreateModelPlanUpgradeInfo) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *FailedCreateModelPlanUpgradeInfo) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *FailedCreateModelPlanUpgradeInfo) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *FailedCreateModelPlanUpgradeInfo) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *FailedCreateModelPlanUpgradeInfo) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *FailedCreateModelPlanUpgradeInfo) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *FailedCreateModelPlanUpgradeInfo) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetMonthOfYear

`func (o *FailedCreateModelPlanUpgradeInfo) GetMonthOfYear() int32`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *FailedCreateModelPlanUpgradeInfo) GetMonthOfYearOk() (*int32, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *FailedCreateModelPlanUpgradeInfo) SetMonthOfYear(v int32)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *FailedCreateModelPlanUpgradeInfo) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### GetScheduleType

`func (o *FailedCreateModelPlanUpgradeInfo) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *FailedCreateModelPlanUpgradeInfo) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *FailedCreateModelPlanUpgradeInfo) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.


### GetYear

`func (o *FailedCreateModelPlanUpgradeInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *FailedCreateModelPlanUpgradeInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *FailedCreateModelPlanUpgradeInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *FailedCreateModelPlanUpgradeInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


