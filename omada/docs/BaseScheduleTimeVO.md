# BaseScheduleTimeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | Day of month when schedule execute, only for timingType: Monthly(3), value is between 1 to 31 | [optional] 
**DayOfWeek** | Pointer to **int32** | Day of week when schedule execute, only for timingType: Weekly(2), value is between 0 to 6 | [optional] 
**Hour** | **int32** | Hour when schedule execute, value is between 0 to 23 | 
**Minute** | **int32** | Minute when schedule execute, value is between 0 to 59 | 
**MonthOfYear** | Pointer to **int32** | Month of year when schedule execute, only for timingType: Yearly(4), value is between 1 to 12 | [optional] 
**TimingType** | **int32** | Time type for schedule task, values are as follows:  Daily(1), Weekly(2), Monthly(3), Yearly(4) | 

## Methods

### NewBaseScheduleTimeVO

`func NewBaseScheduleTimeVO(hour int32, minute int32, timingType int32, ) *BaseScheduleTimeVO`

NewBaseScheduleTimeVO instantiates a new BaseScheduleTimeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseScheduleTimeVOWithDefaults

`func NewBaseScheduleTimeVOWithDefaults() *BaseScheduleTimeVO`

NewBaseScheduleTimeVOWithDefaults instantiates a new BaseScheduleTimeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *BaseScheduleTimeVO) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *BaseScheduleTimeVO) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *BaseScheduleTimeVO) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *BaseScheduleTimeVO) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *BaseScheduleTimeVO) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *BaseScheduleTimeVO) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *BaseScheduleTimeVO) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *BaseScheduleTimeVO) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *BaseScheduleTimeVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *BaseScheduleTimeVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *BaseScheduleTimeVO) SetHour(v int32)`

SetHour sets Hour field to given value.


### GetMinute

`func (o *BaseScheduleTimeVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *BaseScheduleTimeVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *BaseScheduleTimeVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.


### GetMonthOfYear

`func (o *BaseScheduleTimeVO) GetMonthOfYear() int32`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *BaseScheduleTimeVO) GetMonthOfYearOk() (*int32, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *BaseScheduleTimeVO) SetMonthOfYear(v int32)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *BaseScheduleTimeVO) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### GetTimingType

`func (o *BaseScheduleTimeVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *BaseScheduleTimeVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *BaseScheduleTimeVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


