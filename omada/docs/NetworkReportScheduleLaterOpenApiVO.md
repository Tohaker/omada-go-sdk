# NetworkReportScheduleLaterOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | **string** | According to the Tab, fill in the cards parameter with the corresponding json string. The string is a json array whose elements represent a single card. The type and granularity parameters for the cards cannot be removed. Only the granularity value can be changed, which can be 0/1/2 [0:5min, 1:hourly, 2:daily].Summary:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;apSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;ssidSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;bandSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;switchSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;deviceSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;clientsSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;trafficSummary\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;trafficDistribution\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;alertSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;eventSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;appFlowSummary\\\&quot;}]\&quot;;Wireless Summary:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;onlineOfflineAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;apSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;ssidSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;bandSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;apAlert\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;apRebootTimes\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;wirelessOverview\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;wirelessTxRxTraffic\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;wirelessBandTraffic\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;wirelessClientsActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0}]\&quot;;Wired Summary:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;wiredOverview\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;ispLoad\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;onlineOfflineSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;wiredClientsActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;switchSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topPoeUtilSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topPoePowerSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;switchAlert\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;switchRebootTimes\\\&quot;}]\&quot;;Wireless Devices:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;onlineOfflineAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;apSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topCpuUtilAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topMemUtilAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;apAlert\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;apRebootTimes\\\&quot;}]\&quot;;Wired Devices:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;onlineOfflineSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;switchSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;gatewayCpuMemUtil\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topCpuUtilSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topMemUtilSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;switchAlert\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;switchRebootTimes\\\&quot;}]\&quot;;SSID:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;ssidSummary\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;topTrafficSsid\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumSsid\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;ssidNumActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0}]\&quot;;Clients:\&quot;[{\\\&quot;type\\\&quot;:\\\&quot;clientDistribution\\\&quot;},{\\\&quot;type\\\&quot;:\\\&quot;clientNumActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;wirelessClientsActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;guestNumActivities\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumAp\\\&quot;,\\\&quot;granularity\\\&quot;:0},{\\\&quot;type\\\&quot;:\\\&quot;topClientNumSwitch\\\&quot;,\\\&quot;granularity\\\&quot;:0}]\&quot;; | 
**EmailList** | **[]string** | Email List. Example:[\&quot;11@qq.com\&quot;,\&quot;11@qq.com\&quot;] | 
**Enable** | **bool** | Whether to enable periodic report generation and sending. | 
**End** | **int64** | The end timestamp of the query for the data needed to generate the report. | 
**Hour** | **int32** | The value ranges from 0 to 23. | 
**Minute** | **int32** | The value ranges from 0 to 59. | 
**ReportName** | **string** | Report Name. | 
**ReportType** | **int32** | Type of the file to be exported [0:PDF,1:CSV]. | 
**Start** | **int64** | The start timestamp of the query for the data needed to generate the report. | 
**Tab** | **int32** | Report tab, 0-summary, 1-wireless summary, 2-wired summary, 3-wireless devices,4-wired devices, 5-ssid, 6-client. | 
**Time** | **int64** | The timestamp of the year, month, and day is generated. | 

## Methods

### NewNetworkReportScheduleLaterOpenApiVO

`func NewNetworkReportScheduleLaterOpenApiVO(cards string, emailList []string, enable bool, end int64, hour int32, minute int32, reportName string, reportType int32, start int64, tab int32, time int64, ) *NetworkReportScheduleLaterOpenApiVO`

NewNetworkReportScheduleLaterOpenApiVO instantiates a new NetworkReportScheduleLaterOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkReportScheduleLaterOpenApiVOWithDefaults

`func NewNetworkReportScheduleLaterOpenApiVOWithDefaults() *NetworkReportScheduleLaterOpenApiVO`

NewNetworkReportScheduleLaterOpenApiVOWithDefaults instantiates a new NetworkReportScheduleLaterOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *NetworkReportScheduleLaterOpenApiVO) GetCards() string`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetCardsOk() (*string, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *NetworkReportScheduleLaterOpenApiVO) SetCards(v string)`

SetCards sets Cards field to given value.


### GetEmailList

`func (o *NetworkReportScheduleLaterOpenApiVO) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *NetworkReportScheduleLaterOpenApiVO) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.


### GetEnable

`func (o *NetworkReportScheduleLaterOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *NetworkReportScheduleLaterOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetEnd

`func (o *NetworkReportScheduleLaterOpenApiVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *NetworkReportScheduleLaterOpenApiVO) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetHour

`func (o *NetworkReportScheduleLaterOpenApiVO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *NetworkReportScheduleLaterOpenApiVO) SetHour(v int32)`

SetHour sets Hour field to given value.


### GetMinute

`func (o *NetworkReportScheduleLaterOpenApiVO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *NetworkReportScheduleLaterOpenApiVO) SetMinute(v int32)`

SetMinute sets Minute field to given value.


### GetReportName

`func (o *NetworkReportScheduleLaterOpenApiVO) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *NetworkReportScheduleLaterOpenApiVO) SetReportName(v string)`

SetReportName sets ReportName field to given value.


### GetReportType

`func (o *NetworkReportScheduleLaterOpenApiVO) GetReportType() int32`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetReportTypeOk() (*int32, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *NetworkReportScheduleLaterOpenApiVO) SetReportType(v int32)`

SetReportType sets ReportType field to given value.


### GetStart

`func (o *NetworkReportScheduleLaterOpenApiVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *NetworkReportScheduleLaterOpenApiVO) SetStart(v int64)`

SetStart sets Start field to given value.


### GetTab

`func (o *NetworkReportScheduleLaterOpenApiVO) GetTab() int32`

GetTab returns the Tab field if non-nil, zero value otherwise.

### GetTabOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetTabOk() (*int32, bool)`

GetTabOk returns a tuple with the Tab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTab

`func (o *NetworkReportScheduleLaterOpenApiVO) SetTab(v int32)`

SetTab sets Tab field to given value.


### GetTime

`func (o *NetworkReportScheduleLaterOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NetworkReportScheduleLaterOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NetworkReportScheduleLaterOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


