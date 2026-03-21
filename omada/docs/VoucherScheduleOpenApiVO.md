# VoucherScheduleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyEndHour** | Pointer to **int32** | The hour of end time, should be within the range of 0-23 When parameter [type] is 0, parameter [dailyEndHour] is required | [optional] 
**DailyEndMin** | Pointer to **int32** | The minute of end time, should be within the range of 0-59 When parameter [type] is 0, parameter [dailyEndMin] is required | [optional] 
**DailyStartHour** | Pointer to **int32** | The hour of start time, should be within the range of 0-23. When parameter [type] is 0, parameter [dailyStartHour] is required | [optional] 
**DailyStartMin** | Pointer to **int32** | The minute of start time, should be within the range of 0-59 When parameter [type] is 0, parameter [dailyStartMin] is required | [optional] 
**Type** | **int32** | The type of schedule. It should be a value as follows: 0: Limit time by daily, 1: Limit time by weekly | 
**WeeklyEnableDays** | Pointer to **[]int32** | The effective days of week, array number should be within the range of 1-7, 1 represents Monday, 2 represents Tuesday... 7 represents Sunday. When parameter [type] is 1, parameter [weeklyEnableDays] is required | [optional] 

## Methods

### NewVoucherScheduleOpenApiVO

`func NewVoucherScheduleOpenApiVO(type_ int32, ) *VoucherScheduleOpenApiVO`

NewVoucherScheduleOpenApiVO instantiates a new VoucherScheduleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherScheduleOpenApiVOWithDefaults

`func NewVoucherScheduleOpenApiVOWithDefaults() *VoucherScheduleOpenApiVO`

NewVoucherScheduleOpenApiVOWithDefaults instantiates a new VoucherScheduleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyEndHour

`func (o *VoucherScheduleOpenApiVO) GetDailyEndHour() int32`

GetDailyEndHour returns the DailyEndHour field if non-nil, zero value otherwise.

### GetDailyEndHourOk

`func (o *VoucherScheduleOpenApiVO) GetDailyEndHourOk() (*int32, bool)`

GetDailyEndHourOk returns a tuple with the DailyEndHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyEndHour

`func (o *VoucherScheduleOpenApiVO) SetDailyEndHour(v int32)`

SetDailyEndHour sets DailyEndHour field to given value.

### HasDailyEndHour

`func (o *VoucherScheduleOpenApiVO) HasDailyEndHour() bool`

HasDailyEndHour returns a boolean if a field has been set.

### GetDailyEndMin

`func (o *VoucherScheduleOpenApiVO) GetDailyEndMin() int32`

GetDailyEndMin returns the DailyEndMin field if non-nil, zero value otherwise.

### GetDailyEndMinOk

`func (o *VoucherScheduleOpenApiVO) GetDailyEndMinOk() (*int32, bool)`

GetDailyEndMinOk returns a tuple with the DailyEndMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyEndMin

`func (o *VoucherScheduleOpenApiVO) SetDailyEndMin(v int32)`

SetDailyEndMin sets DailyEndMin field to given value.

### HasDailyEndMin

`func (o *VoucherScheduleOpenApiVO) HasDailyEndMin() bool`

HasDailyEndMin returns a boolean if a field has been set.

### GetDailyStartHour

`func (o *VoucherScheduleOpenApiVO) GetDailyStartHour() int32`

GetDailyStartHour returns the DailyStartHour field if non-nil, zero value otherwise.

### GetDailyStartHourOk

`func (o *VoucherScheduleOpenApiVO) GetDailyStartHourOk() (*int32, bool)`

GetDailyStartHourOk returns a tuple with the DailyStartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyStartHour

`func (o *VoucherScheduleOpenApiVO) SetDailyStartHour(v int32)`

SetDailyStartHour sets DailyStartHour field to given value.

### HasDailyStartHour

`func (o *VoucherScheduleOpenApiVO) HasDailyStartHour() bool`

HasDailyStartHour returns a boolean if a field has been set.

### GetDailyStartMin

`func (o *VoucherScheduleOpenApiVO) GetDailyStartMin() int32`

GetDailyStartMin returns the DailyStartMin field if non-nil, zero value otherwise.

### GetDailyStartMinOk

`func (o *VoucherScheduleOpenApiVO) GetDailyStartMinOk() (*int32, bool)`

GetDailyStartMinOk returns a tuple with the DailyStartMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyStartMin

`func (o *VoucherScheduleOpenApiVO) SetDailyStartMin(v int32)`

SetDailyStartMin sets DailyStartMin field to given value.

### HasDailyStartMin

`func (o *VoucherScheduleOpenApiVO) HasDailyStartMin() bool`

HasDailyStartMin returns a boolean if a field has been set.

### GetType

`func (o *VoucherScheduleOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VoucherScheduleOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VoucherScheduleOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.


### GetWeeklyEnableDays

`func (o *VoucherScheduleOpenApiVO) GetWeeklyEnableDays() []int32`

GetWeeklyEnableDays returns the WeeklyEnableDays field if non-nil, zero value otherwise.

### GetWeeklyEnableDaysOk

`func (o *VoucherScheduleOpenApiVO) GetWeeklyEnableDaysOk() (*[]int32, bool)`

GetWeeklyEnableDaysOk returns a tuple with the WeeklyEnableDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyEnableDays

`func (o *VoucherScheduleOpenApiVO) SetWeeklyEnableDays(v []int32)`

SetWeeklyEnableDays sets WeeklyEnableDays field to given value.

### HasWeeklyEnableDays

`func (o *VoucherScheduleOpenApiVO) HasWeeklyEnableDays() bool`

HasWeeklyEnableDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


