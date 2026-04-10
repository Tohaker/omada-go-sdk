# SpeedTestV2SettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSpeedTest** | **bool** | Whether enable scheduled speed testing. | 
**DayOfMonth** | Pointer to **[]int32** | It is required when [timingType] is 3. The value should be within the range of 1~31. | [optional] 
**DayOfWeek** | Pointer to **[]int32** | It is required when [timingType] is 2. The value should be within the range of 0(Sunday)~6(Saturday). | [optional] 
**Hour** | Pointer to **int32** | Start time of speed test(unit: hour); It should be within the range of 0~23. | [optional] 
**Minute** | Pointer to **int32** | Start time of speed test(unit: minute); It should be within the range of 0~59. | [optional] 
**Resource** | Pointer to **int32** | Data Source. Resource should be a value as follows: 0: new created; 1: from template; 2: override | [optional] 
**TimingType** | Pointer to **int32** | The timing type of speed test setting: 1:Daily, 2:Weekly, 3:Monthly. | [optional] 

## Methods

### NewSpeedTestV2SettingVO

`func NewSpeedTestV2SettingVO(autoSpeedTest bool, ) *SpeedTestV2SettingVO`

NewSpeedTestV2SettingVO instantiates a new SpeedTestV2SettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeedTestV2SettingVOWithDefaults

`func NewSpeedTestV2SettingVOWithDefaults() *SpeedTestV2SettingVO`

NewSpeedTestV2SettingVOWithDefaults instantiates a new SpeedTestV2SettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSpeedTest

`func (o *SpeedTestV2SettingVO) GetAutoSpeedTest() bool`

GetAutoSpeedTest returns the AutoSpeedTest field if non-nil, zero value otherwise.

### GetAutoSpeedTestOk

`func (o *SpeedTestV2SettingVO) GetAutoSpeedTestOk() (*bool, bool)`

GetAutoSpeedTestOk returns a tuple with the AutoSpeedTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSpeedTest

`func (o *SpeedTestV2SettingVO) SetAutoSpeedTest(v bool)`

SetAutoSpeedTest sets AutoSpeedTest field to given value.


### GetDayOfMonth

`func (o *SpeedTestV2SettingVO) GetDayOfMonth() []int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *SpeedTestV2SettingVO) GetDayOfMonthOk() (*[]int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *SpeedTestV2SettingVO) SetDayOfMonth(v []int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *SpeedTestV2SettingVO) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *SpeedTestV2SettingVO) GetDayOfWeek() []int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SpeedTestV2SettingVO) GetDayOfWeekOk() (*[]int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SpeedTestV2SettingVO) SetDayOfWeek(v []int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SpeedTestV2SettingVO) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *SpeedTestV2SettingVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *SpeedTestV2SettingVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *SpeedTestV2SettingVO) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *SpeedTestV2SettingVO) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *SpeedTestV2SettingVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *SpeedTestV2SettingVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *SpeedTestV2SettingVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *SpeedTestV2SettingVO) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetResource

`func (o *SpeedTestV2SettingVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SpeedTestV2SettingVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SpeedTestV2SettingVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *SpeedTestV2SettingVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetTimingType

`func (o *SpeedTestV2SettingVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *SpeedTestV2SettingVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *SpeedTestV2SettingVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.

### HasTimingType

`func (o *SpeedTestV2SettingVO) HasTimingType() bool`

HasTimingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


