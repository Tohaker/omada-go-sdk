# NetworkReportScheduleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | **string** | According to the Tab, fill in the cards parameter with the corresponding string. The string is a json array whose elements represent a single card. The type and granularity parameters for the cards cannot be removed. Only the granularity value can be changed, which can be 0/1/2 [0:5min, 1:hourly, 2:daily].Summary:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;apSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;ssidSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;bandSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;switchSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;deviceSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;clientsSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;trafficSummary\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;trafficDistribution\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;alertSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;eventSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;appFlowSummary\\\&quot;}]\&quot;;Wireless Summary:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;onlineOfflineAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;apSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;ssidSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;bandSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;apAlert\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;apRebootTimes\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;wirelessOverview\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;wirelessTxRxTraffic\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;wirelessBandTraffic\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;wirelessClientsActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0}]\&quot;;Wired Summary:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;wiredOverview\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;ispLoad\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;onlineOfflineSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;wiredClientsActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;switchSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topPoeUtilSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topPoePowerSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;switchAlert\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;switchRebootTimes\\\&quot;}]\&quot;;Wireless Devices:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;onlineOfflineAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;apSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topCpuUtilAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topMemUtilAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;apAlert\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;apRebootTimes\\\&quot;}]\&quot;;Wired Devices:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;onlineOfflineSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;switchSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;gatewayCpuMemUtil\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topCpuUtilSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topMemUtilSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;switchAlert\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;switchRebootTimes\\\&quot;}]\&quot;;SSID:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;ssidSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficSsid\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumSsid\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;ssidNumActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0}]\&quot;;Clients:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;clientDistribution\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;clientNumActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;wirelessClientsActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;guestNumActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0}]\&quot;; | 
**DayOfMonth** | Pointer to **int32** | If timing type is monthly, enter and only enter [dayOfMonth].The value ranges from 1 to 31. Schedule will fail on the day of the monthif you select 29th,30th,31th but the month doesn&#39;t have these days. | [optional] 
**DayOfWeek** | Pointer to **int32** | If timing type is weekly, enter and only enter [dayOfWeek].The value ranges from 0 to 6.(&#39;0&#39; indicate Sunday) | [optional] 
**EmailList** | **[]string** | Email List. Example:[\&quot;11@qq.com\&quot;,\&quot;11@qq.com\&quot;]. If &#39;enable&#39; is false, the parameter must be []. | 
**Enable** | **bool** | Whether to enable periodic report generation and sending. | 
**Hour** | **int32** | The value ranges from 0 to 23. | 
**Minute** | **int32** | The value ranges from 0 to 59. | 
**ReportName** | **string** | Parameter [reportName] should be a string of 1 to 64 characters without any space at the beginning or end. | 
**ReportType** | **int32** | Type of the file to be exported [0:PDF,1:CSV]. | 
**Tab** | **int32** | Report tab, 0-summary, 1-wireless summary, 2-wired summary, 3-wireless devices,4-wired devices, 5-ssid, 6-client. | 
**TimingType** | **int32** | Frequency of executing schedule task Daily(1), Weekly(2), Monthly(3). | 

## Methods

### NewNetworkReportScheduleOpenApiVO

`func NewNetworkReportScheduleOpenApiVO(cards string, emailList []string, enable bool, hour int32, minute int32, reportName string, reportType int32, tab int32, timingType int32, ) *NetworkReportScheduleOpenApiVO`

NewNetworkReportScheduleOpenApiVO instantiates a new NetworkReportScheduleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkReportScheduleOpenApiVOWithDefaults

`func NewNetworkReportScheduleOpenApiVOWithDefaults() *NetworkReportScheduleOpenApiVO`

NewNetworkReportScheduleOpenApiVOWithDefaults instantiates a new NetworkReportScheduleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *NetworkReportScheduleOpenApiVO) GetCards() string`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *NetworkReportScheduleOpenApiVO) GetCardsOk() (*string, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *NetworkReportScheduleOpenApiVO) SetCards(v string)`

SetCards sets Cards field to given value.


### GetDayOfMonth

`func (o *NetworkReportScheduleOpenApiVO) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *NetworkReportScheduleOpenApiVO) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *NetworkReportScheduleOpenApiVO) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *NetworkReportScheduleOpenApiVO) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *NetworkReportScheduleOpenApiVO) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *NetworkReportScheduleOpenApiVO) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *NetworkReportScheduleOpenApiVO) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *NetworkReportScheduleOpenApiVO) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetEmailList

`func (o *NetworkReportScheduleOpenApiVO) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *NetworkReportScheduleOpenApiVO) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *NetworkReportScheduleOpenApiVO) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.


### GetEnable

`func (o *NetworkReportScheduleOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *NetworkReportScheduleOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *NetworkReportScheduleOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetHour

`func (o *NetworkReportScheduleOpenApiVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *NetworkReportScheduleOpenApiVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *NetworkReportScheduleOpenApiVO) SetHour(v int32)`

SetHour sets Hour field to given value.


### GetMinute

`func (o *NetworkReportScheduleOpenApiVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *NetworkReportScheduleOpenApiVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *NetworkReportScheduleOpenApiVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.


### GetReportName

`func (o *NetworkReportScheduleOpenApiVO) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *NetworkReportScheduleOpenApiVO) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *NetworkReportScheduleOpenApiVO) SetReportName(v string)`

SetReportName sets ReportName field to given value.


### GetReportType

`func (o *NetworkReportScheduleOpenApiVO) GetReportType() int32`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *NetworkReportScheduleOpenApiVO) GetReportTypeOk() (*int32, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *NetworkReportScheduleOpenApiVO) SetReportType(v int32)`

SetReportType sets ReportType field to given value.


### GetTab

`func (o *NetworkReportScheduleOpenApiVO) GetTab() int32`

GetTab returns the Tab field if non-nil, zero value otherwise.

### GetTabOk

`func (o *NetworkReportScheduleOpenApiVO) GetTabOk() (*int32, bool)`

GetTabOk returns a tuple with the Tab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTab

`func (o *NetworkReportScheduleOpenApiVO) SetTab(v int32)`

SetTab sets Tab field to given value.


### GetTimingType

`func (o *NetworkReportScheduleOpenApiVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *NetworkReportScheduleOpenApiVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *NetworkReportScheduleOpenApiVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


