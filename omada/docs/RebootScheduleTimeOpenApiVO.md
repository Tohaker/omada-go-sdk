# RebootScheduleTimeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | It should be within the range of 1~31. Required when parameter [timingType] is 3. | [optional] 
**DayOfWeek** | Pointer to **int32** | It should be within the range of 0(Sunday)~6(Saturday). Required when parameter [timingType] is 2. | [optional] 
**Hour** | **int32** | It should be within the range of 0~23 | 
**Minute** | **int32** | It should be within the range of 0~59 | 
**TimingType** | **int32** | Time type should be a value as follows: 1:Daily; 2:Weekly; 3:Monthly | 

## Methods

### NewRebootScheduleTimeOpenApiVO

`func NewRebootScheduleTimeOpenApiVO(hour int32, minute int32, timingType int32, ) *RebootScheduleTimeOpenApiVO`

NewRebootScheduleTimeOpenApiVO instantiates a new RebootScheduleTimeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootScheduleTimeOpenApiVOWithDefaults

`func NewRebootScheduleTimeOpenApiVOWithDefaults() *RebootScheduleTimeOpenApiVO`

NewRebootScheduleTimeOpenApiVOWithDefaults instantiates a new RebootScheduleTimeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *RebootScheduleTimeOpenApiVO) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *RebootScheduleTimeOpenApiVO) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *RebootScheduleTimeOpenApiVO) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *RebootScheduleTimeOpenApiVO) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *RebootScheduleTimeOpenApiVO) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *RebootScheduleTimeOpenApiVO) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *RebootScheduleTimeOpenApiVO) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *RebootScheduleTimeOpenApiVO) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *RebootScheduleTimeOpenApiVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *RebootScheduleTimeOpenApiVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *RebootScheduleTimeOpenApiVO) SetHour(v int32)`

SetHour sets Hour field to given value.


### GetMinute

`func (o *RebootScheduleTimeOpenApiVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *RebootScheduleTimeOpenApiVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *RebootScheduleTimeOpenApiVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.


### GetTimingType

`func (o *RebootScheduleTimeOpenApiVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *RebootScheduleTimeOpenApiVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *RebootScheduleTimeOpenApiVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


