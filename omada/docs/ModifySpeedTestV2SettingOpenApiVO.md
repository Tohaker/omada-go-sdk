# ModifySpeedTestV2SettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSpeedTest** | **bool** | Whether enable scheduled speed testing. | 
**DayOfMonth** | Pointer to **[]int32** | It is required when [timingType] is 3. The value should be within the range of 1~31. | [optional] 
**DayOfWeek** | Pointer to **[]int32** | It is required when [timingType] is 2. The value should be within the range of 0(Sunday)~6(Saturday). | [optional] 
**Hour** | Pointer to **int32** | Start time of speed test(unit: hour); It should be within the range of 0~23. | [optional] 
**Minute** | Pointer to **int32** | Start time of speed test(unit: minute); It should be within the range of 0~59. | [optional] 
**TimingType** | Pointer to **int32** | The timing type of speed test setting: 1:Daily, 2:Weekly, 3:Monthly. | [optional] 

## Methods

### NewModifySpeedTestV2SettingOpenApiVO

`func NewModifySpeedTestV2SettingOpenApiVO(autoSpeedTest bool, ) *ModifySpeedTestV2SettingOpenApiVO`

NewModifySpeedTestV2SettingOpenApiVO instantiates a new ModifySpeedTestV2SettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifySpeedTestV2SettingOpenApiVOWithDefaults

`func NewModifySpeedTestV2SettingOpenApiVOWithDefaults() *ModifySpeedTestV2SettingOpenApiVO`

NewModifySpeedTestV2SettingOpenApiVOWithDefaults instantiates a new ModifySpeedTestV2SettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSpeedTest

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetAutoSpeedTest() bool`

GetAutoSpeedTest returns the AutoSpeedTest field if non-nil, zero value otherwise.

### GetAutoSpeedTestOk

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetAutoSpeedTestOk() (*bool, bool)`

GetAutoSpeedTestOk returns a tuple with the AutoSpeedTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSpeedTest

`func (o *ModifySpeedTestV2SettingOpenApiVO) SetAutoSpeedTest(v bool)`

SetAutoSpeedTest sets AutoSpeedTest field to given value.


### GetDayOfMonth

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetDayOfMonth() []int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetDayOfMonthOk() (*[]int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ModifySpeedTestV2SettingOpenApiVO) SetDayOfMonth(v []int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *ModifySpeedTestV2SettingOpenApiVO) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetDayOfWeek() []int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetDayOfWeekOk() (*[]int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *ModifySpeedTestV2SettingOpenApiVO) SetDayOfWeek(v []int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *ModifySpeedTestV2SettingOpenApiVO) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *ModifySpeedTestV2SettingOpenApiVO) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *ModifySpeedTestV2SettingOpenApiVO) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *ModifySpeedTestV2SettingOpenApiVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *ModifySpeedTestV2SettingOpenApiVO) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetTimingType

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *ModifySpeedTestV2SettingOpenApiVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *ModifySpeedTestV2SettingOpenApiVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.

### HasTimingType

`func (o *ModifySpeedTestV2SettingOpenApiVO) HasTimingType() bool`

HasTimingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


