# PlanFirmwareUpgradeCreateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | The day of the month selected by the user, It should be within the range of 1-31 | [optional] 
**Hour** | Pointer to **int32** | The hour of the day selected by the user, It should be within the range of 0-23 | [optional] 
**Minute** | Pointer to **int32** | The minute in the hour selected by the user, It should be within the range of 0-59 | [optional] 
**MonthOfYear** | Pointer to **int32** | The month of the year selected by the user, It should be within the range of 1-12 | [optional] 
**ScheduleType** | **int32** | The type of upgrade execution time, where 0 represents now and 1 represents the specified time | 
**Sites** | **[]string** | List of site ID selected by the user | 
**Year** | Pointer to **int32** | The year selected by the user | [optional] 

## Methods

### NewPlanFirmwareUpgradeCreateInfo

`func NewPlanFirmwareUpgradeCreateInfo(scheduleType int32, sites []string, ) *PlanFirmwareUpgradeCreateInfo`

NewPlanFirmwareUpgradeCreateInfo instantiates a new PlanFirmwareUpgradeCreateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanFirmwareUpgradeCreateInfoWithDefaults

`func NewPlanFirmwareUpgradeCreateInfoWithDefaults() *PlanFirmwareUpgradeCreateInfo`

NewPlanFirmwareUpgradeCreateInfoWithDefaults instantiates a new PlanFirmwareUpgradeCreateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *PlanFirmwareUpgradeCreateInfo) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *PlanFirmwareUpgradeCreateInfo) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *PlanFirmwareUpgradeCreateInfo) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *PlanFirmwareUpgradeCreateInfo) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetHour

`func (o *PlanFirmwareUpgradeCreateInfo) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *PlanFirmwareUpgradeCreateInfo) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *PlanFirmwareUpgradeCreateInfo) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *PlanFirmwareUpgradeCreateInfo) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *PlanFirmwareUpgradeCreateInfo) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *PlanFirmwareUpgradeCreateInfo) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *PlanFirmwareUpgradeCreateInfo) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *PlanFirmwareUpgradeCreateInfo) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetMonthOfYear

`func (o *PlanFirmwareUpgradeCreateInfo) GetMonthOfYear() int32`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *PlanFirmwareUpgradeCreateInfo) GetMonthOfYearOk() (*int32, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *PlanFirmwareUpgradeCreateInfo) SetMonthOfYear(v int32)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *PlanFirmwareUpgradeCreateInfo) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### GetScheduleType

`func (o *PlanFirmwareUpgradeCreateInfo) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *PlanFirmwareUpgradeCreateInfo) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *PlanFirmwareUpgradeCreateInfo) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.


### GetSites

`func (o *PlanFirmwareUpgradeCreateInfo) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PlanFirmwareUpgradeCreateInfo) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PlanFirmwareUpgradeCreateInfo) SetSites(v []string)`

SetSites sets Sites field to given value.


### GetYear

`func (o *PlanFirmwareUpgradeCreateInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *PlanFirmwareUpgradeCreateInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *PlanFirmwareUpgradeCreateInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *PlanFirmwareUpgradeCreateInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


