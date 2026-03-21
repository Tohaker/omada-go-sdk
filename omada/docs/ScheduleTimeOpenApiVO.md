# ScheduleTimeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayType** | **int32** | Time range schedule time day type, 1 to 7 indicates Monday to Sunday, if parameter [dayMode] is 0 to 2, the input value for this field should be 0. | 
**EndTimeH** | **int32** | Time range schedule end time (unit: hour); It should be within the range of 0–24. | 
**EndTimeM** | **int32** | Time range schedule end time (unit: minute); It should be a value as follows: [0, 15, 30, 45]. | 
**StartTimeH** | **int32** | Time range schedule start time (unit: hour); It should be within the range of 0–24. | 
**StartTimeM** | **int32** | Time range schedule start time (unit: minute); It should be a value as follows: [0, 15, 30, 45]. | 

## Methods

### NewScheduleTimeOpenApiVO

`func NewScheduleTimeOpenApiVO(dayType int32, endTimeH int32, endTimeM int32, startTimeH int32, startTimeM int32, ) *ScheduleTimeOpenApiVO`

NewScheduleTimeOpenApiVO instantiates a new ScheduleTimeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTimeOpenApiVOWithDefaults

`func NewScheduleTimeOpenApiVOWithDefaults() *ScheduleTimeOpenApiVO`

NewScheduleTimeOpenApiVOWithDefaults instantiates a new ScheduleTimeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayType

`func (o *ScheduleTimeOpenApiVO) GetDayType() int32`

GetDayType returns the DayType field if non-nil, zero value otherwise.

### GetDayTypeOk

`func (o *ScheduleTimeOpenApiVO) GetDayTypeOk() (*int32, bool)`

GetDayTypeOk returns a tuple with the DayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayType

`func (o *ScheduleTimeOpenApiVO) SetDayType(v int32)`

SetDayType sets DayType field to given value.


### GetEndTimeH

`func (o *ScheduleTimeOpenApiVO) GetEndTimeH() int32`

GetEndTimeH returns the EndTimeH field if non-nil, zero value otherwise.

### GetEndTimeHOk

`func (o *ScheduleTimeOpenApiVO) GetEndTimeHOk() (*int32, bool)`

GetEndTimeHOk returns a tuple with the EndTimeH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeH

`func (o *ScheduleTimeOpenApiVO) SetEndTimeH(v int32)`

SetEndTimeH sets EndTimeH field to given value.


### GetEndTimeM

`func (o *ScheduleTimeOpenApiVO) GetEndTimeM() int32`

GetEndTimeM returns the EndTimeM field if non-nil, zero value otherwise.

### GetEndTimeMOk

`func (o *ScheduleTimeOpenApiVO) GetEndTimeMOk() (*int32, bool)`

GetEndTimeMOk returns a tuple with the EndTimeM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeM

`func (o *ScheduleTimeOpenApiVO) SetEndTimeM(v int32)`

SetEndTimeM sets EndTimeM field to given value.


### GetStartTimeH

`func (o *ScheduleTimeOpenApiVO) GetStartTimeH() int32`

GetStartTimeH returns the StartTimeH field if non-nil, zero value otherwise.

### GetStartTimeHOk

`func (o *ScheduleTimeOpenApiVO) GetStartTimeHOk() (*int32, bool)`

GetStartTimeHOk returns a tuple with the StartTimeH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeH

`func (o *ScheduleTimeOpenApiVO) SetStartTimeH(v int32)`

SetStartTimeH sets StartTimeH field to given value.


### GetStartTimeM

`func (o *ScheduleTimeOpenApiVO) GetStartTimeM() int32`

GetStartTimeM returns the StartTimeM field if non-nil, zero value otherwise.

### GetStartTimeMOk

`func (o *ScheduleTimeOpenApiVO) GetStartTimeMOk() (*int32, bool)`

GetStartTimeMOk returns a tuple with the StartTimeM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeM

`func (o *ScheduleTimeOpenApiVO) SetStartTimeM(v int32)`

SetStartTimeM sets StartTimeM field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


