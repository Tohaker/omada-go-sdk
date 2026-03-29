# UpdateTimeRangeProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDayMode** | Pointer to [**CustomDayModeOpenApiVO**](CustomDayModeOpenApiVO.md) |  | [optional] 
**DayMode** | **int32** | Time range profile day mode; 0: Every Day, 1: Weekday, 2: Weekend, 3: Customized | 
**Name** | **string** | Time range profile name should contain 1 to 64 characters. | 
**TimeList** | [**[]ScheduleTimeOpenApiVO**](ScheduleTimeOpenApiVO.md) | Time range profile schedule time config | 

## Methods

### NewUpdateTimeRangeProfileOpenApiVO

`func NewUpdateTimeRangeProfileOpenApiVO(dayMode int32, name string, timeList []ScheduleTimeOpenApiVO, ) *UpdateTimeRangeProfileOpenApiVO`

NewUpdateTimeRangeProfileOpenApiVO instantiates a new UpdateTimeRangeProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTimeRangeProfileOpenApiVOWithDefaults

`func NewUpdateTimeRangeProfileOpenApiVOWithDefaults() *UpdateTimeRangeProfileOpenApiVO`

NewUpdateTimeRangeProfileOpenApiVOWithDefaults instantiates a new UpdateTimeRangeProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDayMode

`func (o *UpdateTimeRangeProfileOpenApiVO) GetCustomDayMode() CustomDayModeOpenApiVO`

GetCustomDayMode returns the CustomDayMode field if non-nil, zero value otherwise.

### GetCustomDayModeOk

`func (o *UpdateTimeRangeProfileOpenApiVO) GetCustomDayModeOk() (*CustomDayModeOpenApiVO, bool)`

GetCustomDayModeOk returns a tuple with the CustomDayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDayMode

`func (o *UpdateTimeRangeProfileOpenApiVO) SetCustomDayMode(v CustomDayModeOpenApiVO)`

SetCustomDayMode sets CustomDayMode field to given value.

### HasCustomDayMode

`func (o *UpdateTimeRangeProfileOpenApiVO) HasCustomDayMode() bool`

HasCustomDayMode returns a boolean if a field has been set.

### GetDayMode

`func (o *UpdateTimeRangeProfileOpenApiVO) GetDayMode() int32`

GetDayMode returns the DayMode field if non-nil, zero value otherwise.

### GetDayModeOk

`func (o *UpdateTimeRangeProfileOpenApiVO) GetDayModeOk() (*int32, bool)`

GetDayModeOk returns a tuple with the DayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayMode

`func (o *UpdateTimeRangeProfileOpenApiVO) SetDayMode(v int32)`

SetDayMode sets DayMode field to given value.


### GetName

`func (o *UpdateTimeRangeProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTimeRangeProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTimeRangeProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetTimeList

`func (o *UpdateTimeRangeProfileOpenApiVO) GetTimeList() []ScheduleTimeOpenApiVO`

GetTimeList returns the TimeList field if non-nil, zero value otherwise.

### GetTimeListOk

`func (o *UpdateTimeRangeProfileOpenApiVO) GetTimeListOk() (*[]ScheduleTimeOpenApiVO, bool)`

GetTimeListOk returns a tuple with the TimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeList

`func (o *UpdateTimeRangeProfileOpenApiVO) SetTimeList(v []ScheduleTimeOpenApiVO)`

SetTimeList sets TimeList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


