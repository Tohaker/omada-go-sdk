# NetworkReportScheduleLaterVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to **string** |  | [optional] 
**EmailList** | Pointer to **[]string** |  | [optional] 
**Enable** | **bool** |  | 
**End** | Pointer to **int64** |  | [optional] 
**Hour** | Pointer to **int32** |  | [optional] 
**Minute** | Pointer to **int32** |  | [optional] 
**ReportName** | Pointer to **string** |  | [optional] 
**ReportType** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**Tab** | Pointer to **int32** |  | [optional] 
**TabIdList** | Pointer to **[]string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewNetworkReportScheduleLaterVO

`func NewNetworkReportScheduleLaterVO(enable bool, ) *NetworkReportScheduleLaterVO`

NewNetworkReportScheduleLaterVO instantiates a new NetworkReportScheduleLaterVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkReportScheduleLaterVOWithDefaults

`func NewNetworkReportScheduleLaterVOWithDefaults() *NetworkReportScheduleLaterVO`

NewNetworkReportScheduleLaterVOWithDefaults instantiates a new NetworkReportScheduleLaterVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *NetworkReportScheduleLaterVO) GetCards() string`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *NetworkReportScheduleLaterVO) GetCardsOk() (*string, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *NetworkReportScheduleLaterVO) SetCards(v string)`

SetCards sets Cards field to given value.

### HasCards

`func (o *NetworkReportScheduleLaterVO) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetEmailList

`func (o *NetworkReportScheduleLaterVO) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *NetworkReportScheduleLaterVO) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *NetworkReportScheduleLaterVO) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *NetworkReportScheduleLaterVO) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnable

`func (o *NetworkReportScheduleLaterVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *NetworkReportScheduleLaterVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *NetworkReportScheduleLaterVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetEnd

`func (o *NetworkReportScheduleLaterVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *NetworkReportScheduleLaterVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *NetworkReportScheduleLaterVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *NetworkReportScheduleLaterVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetHour

`func (o *NetworkReportScheduleLaterVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *NetworkReportScheduleLaterVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *NetworkReportScheduleLaterVO) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *NetworkReportScheduleLaterVO) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *NetworkReportScheduleLaterVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *NetworkReportScheduleLaterVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *NetworkReportScheduleLaterVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *NetworkReportScheduleLaterVO) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetReportName

`func (o *NetworkReportScheduleLaterVO) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *NetworkReportScheduleLaterVO) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *NetworkReportScheduleLaterVO) SetReportName(v string)`

SetReportName sets ReportName field to given value.

### HasReportName

`func (o *NetworkReportScheduleLaterVO) HasReportName() bool`

HasReportName returns a boolean if a field has been set.

### GetReportType

`func (o *NetworkReportScheduleLaterVO) GetReportType() int32`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *NetworkReportScheduleLaterVO) GetReportTypeOk() (*int32, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *NetworkReportScheduleLaterVO) SetReportType(v int32)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *NetworkReportScheduleLaterVO) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetStart

`func (o *NetworkReportScheduleLaterVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *NetworkReportScheduleLaterVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *NetworkReportScheduleLaterVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *NetworkReportScheduleLaterVO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTab

`func (o *NetworkReportScheduleLaterVO) GetTab() int32`

GetTab returns the Tab field if non-nil, zero value otherwise.

### GetTabOk

`func (o *NetworkReportScheduleLaterVO) GetTabOk() (*int32, bool)`

GetTabOk returns a tuple with the Tab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTab

`func (o *NetworkReportScheduleLaterVO) SetTab(v int32)`

SetTab sets Tab field to given value.

### HasTab

`func (o *NetworkReportScheduleLaterVO) HasTab() bool`

HasTab returns a boolean if a field has been set.

### GetTabIdList

`func (o *NetworkReportScheduleLaterVO) GetTabIdList() []string`

GetTabIdList returns the TabIdList field if non-nil, zero value otherwise.

### GetTabIdListOk

`func (o *NetworkReportScheduleLaterVO) GetTabIdListOk() (*[]string, bool)`

GetTabIdListOk returns a tuple with the TabIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabIdList

`func (o *NetworkReportScheduleLaterVO) SetTabIdList(v []string)`

SetTabIdList sets TabIdList field to given value.

### HasTabIdList

`func (o *NetworkReportScheduleLaterVO) HasTabIdList() bool`

HasTabIdList returns a boolean if a field has been set.

### GetTime

`func (o *NetworkReportScheduleLaterVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NetworkReportScheduleLaterVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NetworkReportScheduleLaterVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *NetworkReportScheduleLaterVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


