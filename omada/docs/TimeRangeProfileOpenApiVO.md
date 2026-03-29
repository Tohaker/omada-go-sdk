# TimeRangeProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDayMode** | Pointer to [**CustomDayModeOpenApiVO**](CustomDayModeOpenApiVO.md) |  | [optional] 
**DayMode** | Pointer to **int32** | Time range profile day mode; It should be a value as follows: 0: Every Day, 1: Weekday, 2: Weekend, 3: Customized | [optional] 
**Name** | Pointer to **string** | Time range profile name should contain 1 to 64 characters. | [optional] 
**ProfileId** | Pointer to **string** | Time range profile ID | [optional] 
**TimeList** | Pointer to [**[]ScheduleTimeOpenApiVO**](ScheduleTimeOpenApiVO.md) | Time range profile schedule time config | [optional] 

## Methods

### NewTimeRangeProfileOpenApiVO

`func NewTimeRangeProfileOpenApiVO() *TimeRangeProfileOpenApiVO`

NewTimeRangeProfileOpenApiVO instantiates a new TimeRangeProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeRangeProfileOpenApiVOWithDefaults

`func NewTimeRangeProfileOpenApiVOWithDefaults() *TimeRangeProfileOpenApiVO`

NewTimeRangeProfileOpenApiVOWithDefaults instantiates a new TimeRangeProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDayMode

`func (o *TimeRangeProfileOpenApiVO) GetCustomDayMode() CustomDayModeOpenApiVO`

GetCustomDayMode returns the CustomDayMode field if non-nil, zero value otherwise.

### GetCustomDayModeOk

`func (o *TimeRangeProfileOpenApiVO) GetCustomDayModeOk() (*CustomDayModeOpenApiVO, bool)`

GetCustomDayModeOk returns a tuple with the CustomDayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDayMode

`func (o *TimeRangeProfileOpenApiVO) SetCustomDayMode(v CustomDayModeOpenApiVO)`

SetCustomDayMode sets CustomDayMode field to given value.

### HasCustomDayMode

`func (o *TimeRangeProfileOpenApiVO) HasCustomDayMode() bool`

HasCustomDayMode returns a boolean if a field has been set.

### GetDayMode

`func (o *TimeRangeProfileOpenApiVO) GetDayMode() int32`

GetDayMode returns the DayMode field if non-nil, zero value otherwise.

### GetDayModeOk

`func (o *TimeRangeProfileOpenApiVO) GetDayModeOk() (*int32, bool)`

GetDayModeOk returns a tuple with the DayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayMode

`func (o *TimeRangeProfileOpenApiVO) SetDayMode(v int32)`

SetDayMode sets DayMode field to given value.

### HasDayMode

`func (o *TimeRangeProfileOpenApiVO) HasDayMode() bool`

HasDayMode returns a boolean if a field has been set.

### GetName

`func (o *TimeRangeProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeRangeProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeRangeProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimeRangeProfileOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileId

`func (o *TimeRangeProfileOpenApiVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *TimeRangeProfileOpenApiVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *TimeRangeProfileOpenApiVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *TimeRangeProfileOpenApiVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetTimeList

`func (o *TimeRangeProfileOpenApiVO) GetTimeList() []ScheduleTimeOpenApiVO`

GetTimeList returns the TimeList field if non-nil, zero value otherwise.

### GetTimeListOk

`func (o *TimeRangeProfileOpenApiVO) GetTimeListOk() (*[]ScheduleTimeOpenApiVO, bool)`

GetTimeListOk returns a tuple with the TimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeList

`func (o *TimeRangeProfileOpenApiVO) SetTimeList(v []ScheduleTimeOpenApiVO)`

SetTimeList sets TimeList field to given value.

### HasTimeList

`func (o *TimeRangeProfileOpenApiVO) HasTimeList() bool`

HasTimeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


