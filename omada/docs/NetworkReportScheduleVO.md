# NetworkReportScheduleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to **string** |  | [optional] 
**DayOfMonth** | Pointer to **int32** |  | [optional] 
**DayOfWeek** | Pointer to **int32** |  | [optional] 
**EmailList** | Pointer to **[]string** |  | [optional] 
**Enable** | **bool** |  | 
**Hour** | Pointer to **int32** |  | [optional] 
**Minute** | Pointer to **int32** |  | [optional] 
**MonthOfYear** | Pointer to **int32** |  | [optional] 
**ReportName** | Pointer to **string** |  | [optional] 
**ReportType** | Pointer to **int32** |  | [optional] 
**Tab** | Pointer to **int32** |  | [optional] 
**TabIdList** | Pointer to **[]string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimingType** | Pointer to **int32** |  | [optional] 

## Methods

### NewNetworkReportScheduleVO

`func NewNetworkReportScheduleVO(enable bool, ) *NetworkReportScheduleVO`

NewNetworkReportScheduleVO instantiates a new NetworkReportScheduleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkReportScheduleVOWithDefaults

`func NewNetworkReportScheduleVOWithDefaults() *NetworkReportScheduleVO`

NewNetworkReportScheduleVOWithDefaults instantiates a new NetworkReportScheduleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *NetworkReportScheduleVO) GetCards() string`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *NetworkReportScheduleVO) GetCardsOk() (*string, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *NetworkReportScheduleVO) SetCards(v string)`

SetCards sets Cards field to given value.

### HasCards

`func (o *NetworkReportScheduleVO) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *NetworkReportScheduleVO) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *NetworkReportScheduleVO) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *NetworkReportScheduleVO) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *NetworkReportScheduleVO) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *NetworkReportScheduleVO) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *NetworkReportScheduleVO) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *NetworkReportScheduleVO) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *NetworkReportScheduleVO) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetEmailList

`func (o *NetworkReportScheduleVO) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *NetworkReportScheduleVO) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *NetworkReportScheduleVO) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *NetworkReportScheduleVO) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnable

`func (o *NetworkReportScheduleVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *NetworkReportScheduleVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *NetworkReportScheduleVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetHour

`func (o *NetworkReportScheduleVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *NetworkReportScheduleVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *NetworkReportScheduleVO) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *NetworkReportScheduleVO) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *NetworkReportScheduleVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *NetworkReportScheduleVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *NetworkReportScheduleVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *NetworkReportScheduleVO) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetMonthOfYear

`func (o *NetworkReportScheduleVO) GetMonthOfYear() int32`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *NetworkReportScheduleVO) GetMonthOfYearOk() (*int32, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *NetworkReportScheduleVO) SetMonthOfYear(v int32)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *NetworkReportScheduleVO) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### GetReportName

`func (o *NetworkReportScheduleVO) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *NetworkReportScheduleVO) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *NetworkReportScheduleVO) SetReportName(v string)`

SetReportName sets ReportName field to given value.

### HasReportName

`func (o *NetworkReportScheduleVO) HasReportName() bool`

HasReportName returns a boolean if a field has been set.

### GetReportType

`func (o *NetworkReportScheduleVO) GetReportType() int32`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *NetworkReportScheduleVO) GetReportTypeOk() (*int32, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *NetworkReportScheduleVO) SetReportType(v int32)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *NetworkReportScheduleVO) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetTab

`func (o *NetworkReportScheduleVO) GetTab() int32`

GetTab returns the Tab field if non-nil, zero value otherwise.

### GetTabOk

`func (o *NetworkReportScheduleVO) GetTabOk() (*int32, bool)`

GetTabOk returns a tuple with the Tab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTab

`func (o *NetworkReportScheduleVO) SetTab(v int32)`

SetTab sets Tab field to given value.

### HasTab

`func (o *NetworkReportScheduleVO) HasTab() bool`

HasTab returns a boolean if a field has been set.

### GetTabIdList

`func (o *NetworkReportScheduleVO) GetTabIdList() []string`

GetTabIdList returns the TabIdList field if non-nil, zero value otherwise.

### GetTabIdListOk

`func (o *NetworkReportScheduleVO) GetTabIdListOk() (*[]string, bool)`

GetTabIdListOk returns a tuple with the TabIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabIdList

`func (o *NetworkReportScheduleVO) SetTabIdList(v []string)`

SetTabIdList sets TabIdList field to given value.

### HasTabIdList

`func (o *NetworkReportScheduleVO) HasTabIdList() bool`

HasTabIdList returns a boolean if a field has been set.

### GetTime

`func (o *NetworkReportScheduleVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NetworkReportScheduleVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NetworkReportScheduleVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *NetworkReportScheduleVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimingType

`func (o *NetworkReportScheduleVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *NetworkReportScheduleVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *NetworkReportScheduleVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.

### HasTimingType

`func (o *NetworkReportScheduleVO) HasTimingType() bool`

HasTimingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


