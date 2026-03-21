# BaseScheduleTimeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | It should be within the range of 1~31. Required when parameter [timingType] is 3. | [optional] 
**DayOfWeek** | Pointer to **int32** | It should be within the range of 0(Sunday)~6(Saturday). Required when parameter [timingType] is 2. | [optional] 
**Hour** | **int32** | It should be within the range of 0~23 | 
**Minute** | **int32** | It should be within the range of 0~59 | 
**MonthOfYear** | Pointer to **int32** | It should be within the range of 1~12. Required when parameter [timingType] is 4. | [optional] 
**TimingType** | **int32** | Time type should be a value as follows: 1:Daily; 2:Weekly; 3:Monthly | 

## Methods

### NewBaseScheduleTimeOpenApiVO

`func NewBaseScheduleTimeOpenApiVO(hour int32, minute int32, timingType int32, ) *BaseScheduleTimeOpenApiVO`

NewBaseScheduleTimeOpenApiVO instantiates a new BaseScheduleTimeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseScheduleTimeOpenApiVOWithDefaults

`func NewBaseScheduleTimeOpenApiVOWithDefaults() *BaseScheduleTimeOpenApiVO`

NewBaseScheduleTimeOpenApiVOWithDefaults instantiates a new BaseScheduleTimeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *BaseScheduleTimeOpenApiVO) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *BaseScheduleTimeOpenApiVO) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *BaseScheduleTimeOpenApiVO) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *BaseScheduleTimeOpenApiVO) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *BaseScheduleTimeOpenApiVO) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *BaseScheduleTimeOpenApiVO) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *BaseScheduleTimeOpenApiVO) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *BaseScheduleTimeOpenApiVO) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *BaseScheduleTimeOpenApiVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *BaseScheduleTimeOpenApiVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *BaseScheduleTimeOpenApiVO) SetHour(v int32)`

SetHour sets Hour field to given value.


### GetMinute

`func (o *BaseScheduleTimeOpenApiVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *BaseScheduleTimeOpenApiVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *BaseScheduleTimeOpenApiVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.


### GetMonthOfYear

`func (o *BaseScheduleTimeOpenApiVO) GetMonthOfYear() int32`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *BaseScheduleTimeOpenApiVO) GetMonthOfYearOk() (*int32, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *BaseScheduleTimeOpenApiVO) SetMonthOfYear(v int32)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *BaseScheduleTimeOpenApiVO) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### GetTimingType

`func (o *BaseScheduleTimeOpenApiVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *BaseScheduleTimeOpenApiVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *BaseScheduleTimeOpenApiVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


