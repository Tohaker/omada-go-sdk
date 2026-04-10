# HistoryRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDataTrendDaily** | Pointer to **int32** | Retention Configuration of client data trend records(only effective in local controller), clientStatDaily should be a value as follows: 1: 1day; 7: 7days; 31: 31days; 60: 60days, 90: 90days. | [optional] 
**ClientDataTrendEnable** | Pointer to **bool** | Whether the client data trend records is recorded. | [optional] 
**ClientHealthEnable** | Pointer to **bool** | Whether client health is enabled. When enabled, client health data will be recorded, which may consume a significant amount of storage space. | [optional] 
**ClientHistory** | Pointer to **int32** | Retention configuration of client History(only effective in local controller), clientHistory should be a value as follows: -1: Disabled; 0: All Time(Windows, Linux Only); 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**ClientHistoryAvailableRetentionDays** | Pointer to **[]int32** | Provide optional clientHistory data Retention Configuration list. | [optional] 
**ClientRecognitionEnable** | Pointer to **bool** | Whether client recognition is enabled. With the feature enabled, network devices will report client information in real time to ensure the accuracy of client recognition. Cloud Access is required for client recognition. This feature is not supported in the MSP view. | [optional] 
**ClientTrendAvailableRetentionDays** | Pointer to **[]int32** | Provide optional clientDataTrendDaily data Retention Configuration list. | [optional] 
**ClientsDataEnable** | **bool** | Whether the clients&#39; history data is recorded. | 
**Daily** | Pointer to **int32** | Retention configuration of time series with daily granularity, daily should be a value as follows: 90: 90days; 180: 180days; 365: 365days(Fixed value in Cloud Based Controller as 365 days). | [optional] 
**DailyAvailableRetentionDays** | Pointer to **[]int32** | Provide optional daily data Retention Configuration list. | [optional] 
**FiveMin** | Pointer to **int32** | Retention configuration of Time Series with 5 Minutes Granularity. It is fixed to 2days and cannot be changed. | [optional] 
**Hourly** | Pointer to **int32** | Retention configuration of time series with hourly granularity, hourly should be a value as follows: 7: 7days. | [optional] 
**KnownClient** | Pointer to **int32** | Retention configuration of known client Data, knownClient should be a value as follows: -1: Disabled; 0: All Time(Windows, Linux Only); 1: 1day; 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**KnownClientAvailableRetentionDays** | Pointer to **[]int32** | Provide optional knownClient data Retention Configuration list. | [optional] 
**Log** | Pointer to **int32** | Retention Configuration of log data(only effective in local controller), log should be a value as follows: 0: All Time(Windows, Linux Only); 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**LogAvailableRetentionDays** | Pointer to **[]int32** | Provide optional log data Retention Configuration list. | [optional] 
**Override** | Pointer to **bool** | Whether the customer overrides the retention configuration of MSP. | [optional] 
**PortalAuth** | Pointer to **int32** | Retention configuration of portal authentication records, portalAuth should be a value as follows: 0: All Time(Windows, Linux Only); 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**PortalAuthAvailableRetentionDays** | Pointer to **[]int32** | Provide optional portalAuth data Retention Configuration list. | [optional] 
**RogueAp** | Pointer to **int32** | Retention Configuration of rogue ap data, rogueAp should be a value as follows: 0: All Time(Windows, Linux Only); 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**TenMin** | Pointer to **int32** | Retention configuration of Time Series with 10 Minutes Granularity, only for Lite Cloud-Based Controller. It is fixed to 2days and cannot be changed. | [optional] 
**Weekly** | Pointer to **int32** | Retention configuration of time series with weekly granularity, weekly should be a value as follows: 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**WeeklyAvailableRetentionDays** | Pointer to **[]int32** | Provide optional weekly data Retention Configuration list. | [optional] 
**WidsData** | Pointer to **int32** | Retention Configuration of wids data(only effective in local pro controller), widsData should be a value as follows: 0: All Time(Windows, Linux Only); 90: 90days; 180: 180days; 365: 365days. | [optional] 

## Methods

### NewHistoryRetention

`func NewHistoryRetention(clientsDataEnable bool, ) *HistoryRetention`

NewHistoryRetention instantiates a new HistoryRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryRetentionWithDefaults

`func NewHistoryRetentionWithDefaults() *HistoryRetention`

NewHistoryRetentionWithDefaults instantiates a new HistoryRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientDataTrendDaily

`func (o *HistoryRetention) GetClientDataTrendDaily() int32`

GetClientDataTrendDaily returns the ClientDataTrendDaily field if non-nil, zero value otherwise.

### GetClientDataTrendDailyOk

`func (o *HistoryRetention) GetClientDataTrendDailyOk() (*int32, bool)`

GetClientDataTrendDailyOk returns a tuple with the ClientDataTrendDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDataTrendDaily

`func (o *HistoryRetention) SetClientDataTrendDaily(v int32)`

SetClientDataTrendDaily sets ClientDataTrendDaily field to given value.

### HasClientDataTrendDaily

`func (o *HistoryRetention) HasClientDataTrendDaily() bool`

HasClientDataTrendDaily returns a boolean if a field has been set.

### GetClientDataTrendEnable

`func (o *HistoryRetention) GetClientDataTrendEnable() bool`

GetClientDataTrendEnable returns the ClientDataTrendEnable field if non-nil, zero value otherwise.

### GetClientDataTrendEnableOk

`func (o *HistoryRetention) GetClientDataTrendEnableOk() (*bool, bool)`

GetClientDataTrendEnableOk returns a tuple with the ClientDataTrendEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDataTrendEnable

`func (o *HistoryRetention) SetClientDataTrendEnable(v bool)`

SetClientDataTrendEnable sets ClientDataTrendEnable field to given value.

### HasClientDataTrendEnable

`func (o *HistoryRetention) HasClientDataTrendEnable() bool`

HasClientDataTrendEnable returns a boolean if a field has been set.

### GetClientHealthEnable

`func (o *HistoryRetention) GetClientHealthEnable() bool`

GetClientHealthEnable returns the ClientHealthEnable field if non-nil, zero value otherwise.

### GetClientHealthEnableOk

`func (o *HistoryRetention) GetClientHealthEnableOk() (*bool, bool)`

GetClientHealthEnableOk returns a tuple with the ClientHealthEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealthEnable

`func (o *HistoryRetention) SetClientHealthEnable(v bool)`

SetClientHealthEnable sets ClientHealthEnable field to given value.

### HasClientHealthEnable

`func (o *HistoryRetention) HasClientHealthEnable() bool`

HasClientHealthEnable returns a boolean if a field has been set.

### GetClientHistory

`func (o *HistoryRetention) GetClientHistory() int32`

GetClientHistory returns the ClientHistory field if non-nil, zero value otherwise.

### GetClientHistoryOk

`func (o *HistoryRetention) GetClientHistoryOk() (*int32, bool)`

GetClientHistoryOk returns a tuple with the ClientHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHistory

`func (o *HistoryRetention) SetClientHistory(v int32)`

SetClientHistory sets ClientHistory field to given value.

### HasClientHistory

`func (o *HistoryRetention) HasClientHistory() bool`

HasClientHistory returns a boolean if a field has been set.

### GetClientHistoryAvailableRetentionDays

`func (o *HistoryRetention) GetClientHistoryAvailableRetentionDays() []int32`

GetClientHistoryAvailableRetentionDays returns the ClientHistoryAvailableRetentionDays field if non-nil, zero value otherwise.

### GetClientHistoryAvailableRetentionDaysOk

`func (o *HistoryRetention) GetClientHistoryAvailableRetentionDaysOk() (*[]int32, bool)`

GetClientHistoryAvailableRetentionDaysOk returns a tuple with the ClientHistoryAvailableRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHistoryAvailableRetentionDays

`func (o *HistoryRetention) SetClientHistoryAvailableRetentionDays(v []int32)`

SetClientHistoryAvailableRetentionDays sets ClientHistoryAvailableRetentionDays field to given value.

### HasClientHistoryAvailableRetentionDays

`func (o *HistoryRetention) HasClientHistoryAvailableRetentionDays() bool`

HasClientHistoryAvailableRetentionDays returns a boolean if a field has been set.

### GetClientRecognitionEnable

`func (o *HistoryRetention) GetClientRecognitionEnable() bool`

GetClientRecognitionEnable returns the ClientRecognitionEnable field if non-nil, zero value otherwise.

### GetClientRecognitionEnableOk

`func (o *HistoryRetention) GetClientRecognitionEnableOk() (*bool, bool)`

GetClientRecognitionEnableOk returns a tuple with the ClientRecognitionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRecognitionEnable

`func (o *HistoryRetention) SetClientRecognitionEnable(v bool)`

SetClientRecognitionEnable sets ClientRecognitionEnable field to given value.

### HasClientRecognitionEnable

`func (o *HistoryRetention) HasClientRecognitionEnable() bool`

HasClientRecognitionEnable returns a boolean if a field has been set.

### GetClientTrendAvailableRetentionDays

`func (o *HistoryRetention) GetClientTrendAvailableRetentionDays() []int32`

GetClientTrendAvailableRetentionDays returns the ClientTrendAvailableRetentionDays field if non-nil, zero value otherwise.

### GetClientTrendAvailableRetentionDaysOk

`func (o *HistoryRetention) GetClientTrendAvailableRetentionDaysOk() (*[]int32, bool)`

GetClientTrendAvailableRetentionDaysOk returns a tuple with the ClientTrendAvailableRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTrendAvailableRetentionDays

`func (o *HistoryRetention) SetClientTrendAvailableRetentionDays(v []int32)`

SetClientTrendAvailableRetentionDays sets ClientTrendAvailableRetentionDays field to given value.

### HasClientTrendAvailableRetentionDays

`func (o *HistoryRetention) HasClientTrendAvailableRetentionDays() bool`

HasClientTrendAvailableRetentionDays returns a boolean if a field has been set.

### GetClientsDataEnable

`func (o *HistoryRetention) GetClientsDataEnable() bool`

GetClientsDataEnable returns the ClientsDataEnable field if non-nil, zero value otherwise.

### GetClientsDataEnableOk

`func (o *HistoryRetention) GetClientsDataEnableOk() (*bool, bool)`

GetClientsDataEnableOk returns a tuple with the ClientsDataEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsDataEnable

`func (o *HistoryRetention) SetClientsDataEnable(v bool)`

SetClientsDataEnable sets ClientsDataEnable field to given value.


### GetDaily

`func (o *HistoryRetention) GetDaily() int32`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *HistoryRetention) GetDailyOk() (*int32, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *HistoryRetention) SetDaily(v int32)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *HistoryRetention) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetDailyAvailableRetentionDays

`func (o *HistoryRetention) GetDailyAvailableRetentionDays() []int32`

GetDailyAvailableRetentionDays returns the DailyAvailableRetentionDays field if non-nil, zero value otherwise.

### GetDailyAvailableRetentionDaysOk

`func (o *HistoryRetention) GetDailyAvailableRetentionDaysOk() (*[]int32, bool)`

GetDailyAvailableRetentionDaysOk returns a tuple with the DailyAvailableRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyAvailableRetentionDays

`func (o *HistoryRetention) SetDailyAvailableRetentionDays(v []int32)`

SetDailyAvailableRetentionDays sets DailyAvailableRetentionDays field to given value.

### HasDailyAvailableRetentionDays

`func (o *HistoryRetention) HasDailyAvailableRetentionDays() bool`

HasDailyAvailableRetentionDays returns a boolean if a field has been set.

### GetFiveMin

`func (o *HistoryRetention) GetFiveMin() int32`

GetFiveMin returns the FiveMin field if non-nil, zero value otherwise.

### GetFiveMinOk

`func (o *HistoryRetention) GetFiveMinOk() (*int32, bool)`

GetFiveMinOk returns a tuple with the FiveMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveMin

`func (o *HistoryRetention) SetFiveMin(v int32)`

SetFiveMin sets FiveMin field to given value.

### HasFiveMin

`func (o *HistoryRetention) HasFiveMin() bool`

HasFiveMin returns a boolean if a field has been set.

### GetHourly

`func (o *HistoryRetention) GetHourly() int32`

GetHourly returns the Hourly field if non-nil, zero value otherwise.

### GetHourlyOk

`func (o *HistoryRetention) GetHourlyOk() (*int32, bool)`

GetHourlyOk returns a tuple with the Hourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourly

`func (o *HistoryRetention) SetHourly(v int32)`

SetHourly sets Hourly field to given value.

### HasHourly

`func (o *HistoryRetention) HasHourly() bool`

HasHourly returns a boolean if a field has been set.

### GetKnownClient

`func (o *HistoryRetention) GetKnownClient() int32`

GetKnownClient returns the KnownClient field if non-nil, zero value otherwise.

### GetKnownClientOk

`func (o *HistoryRetention) GetKnownClientOk() (*int32, bool)`

GetKnownClientOk returns a tuple with the KnownClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownClient

`func (o *HistoryRetention) SetKnownClient(v int32)`

SetKnownClient sets KnownClient field to given value.

### HasKnownClient

`func (o *HistoryRetention) HasKnownClient() bool`

HasKnownClient returns a boolean if a field has been set.

### GetKnownClientAvailableRetentionDays

`func (o *HistoryRetention) GetKnownClientAvailableRetentionDays() []int32`

GetKnownClientAvailableRetentionDays returns the KnownClientAvailableRetentionDays field if non-nil, zero value otherwise.

### GetKnownClientAvailableRetentionDaysOk

`func (o *HistoryRetention) GetKnownClientAvailableRetentionDaysOk() (*[]int32, bool)`

GetKnownClientAvailableRetentionDaysOk returns a tuple with the KnownClientAvailableRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownClientAvailableRetentionDays

`func (o *HistoryRetention) SetKnownClientAvailableRetentionDays(v []int32)`

SetKnownClientAvailableRetentionDays sets KnownClientAvailableRetentionDays field to given value.

### HasKnownClientAvailableRetentionDays

`func (o *HistoryRetention) HasKnownClientAvailableRetentionDays() bool`

HasKnownClientAvailableRetentionDays returns a boolean if a field has been set.

### GetLog

`func (o *HistoryRetention) GetLog() int32`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *HistoryRetention) GetLogOk() (*int32, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *HistoryRetention) SetLog(v int32)`

SetLog sets Log field to given value.

### HasLog

`func (o *HistoryRetention) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetLogAvailableRetentionDays

`func (o *HistoryRetention) GetLogAvailableRetentionDays() []int32`

GetLogAvailableRetentionDays returns the LogAvailableRetentionDays field if non-nil, zero value otherwise.

### GetLogAvailableRetentionDaysOk

`func (o *HistoryRetention) GetLogAvailableRetentionDaysOk() (*[]int32, bool)`

GetLogAvailableRetentionDaysOk returns a tuple with the LogAvailableRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAvailableRetentionDays

`func (o *HistoryRetention) SetLogAvailableRetentionDays(v []int32)`

SetLogAvailableRetentionDays sets LogAvailableRetentionDays field to given value.

### HasLogAvailableRetentionDays

`func (o *HistoryRetention) HasLogAvailableRetentionDays() bool`

HasLogAvailableRetentionDays returns a boolean if a field has been set.

### GetOverride

`func (o *HistoryRetention) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *HistoryRetention) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *HistoryRetention) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *HistoryRetention) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetPortalAuth

`func (o *HistoryRetention) GetPortalAuth() int32`

GetPortalAuth returns the PortalAuth field if non-nil, zero value otherwise.

### GetPortalAuthOk

`func (o *HistoryRetention) GetPortalAuthOk() (*int32, bool)`

GetPortalAuthOk returns a tuple with the PortalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalAuth

`func (o *HistoryRetention) SetPortalAuth(v int32)`

SetPortalAuth sets PortalAuth field to given value.

### HasPortalAuth

`func (o *HistoryRetention) HasPortalAuth() bool`

HasPortalAuth returns a boolean if a field has been set.

### GetPortalAuthAvailableRetentionDays

`func (o *HistoryRetention) GetPortalAuthAvailableRetentionDays() []int32`

GetPortalAuthAvailableRetentionDays returns the PortalAuthAvailableRetentionDays field if non-nil, zero value otherwise.

### GetPortalAuthAvailableRetentionDaysOk

`func (o *HistoryRetention) GetPortalAuthAvailableRetentionDaysOk() (*[]int32, bool)`

GetPortalAuthAvailableRetentionDaysOk returns a tuple with the PortalAuthAvailableRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalAuthAvailableRetentionDays

`func (o *HistoryRetention) SetPortalAuthAvailableRetentionDays(v []int32)`

SetPortalAuthAvailableRetentionDays sets PortalAuthAvailableRetentionDays field to given value.

### HasPortalAuthAvailableRetentionDays

`func (o *HistoryRetention) HasPortalAuthAvailableRetentionDays() bool`

HasPortalAuthAvailableRetentionDays returns a boolean if a field has been set.

### GetRogueAp

`func (o *HistoryRetention) GetRogueAp() int32`

GetRogueAp returns the RogueAp field if non-nil, zero value otherwise.

### GetRogueApOk

`func (o *HistoryRetention) GetRogueApOk() (*int32, bool)`

GetRogueApOk returns a tuple with the RogueAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRogueAp

`func (o *HistoryRetention) SetRogueAp(v int32)`

SetRogueAp sets RogueAp field to given value.

### HasRogueAp

`func (o *HistoryRetention) HasRogueAp() bool`

HasRogueAp returns a boolean if a field has been set.

### GetTenMin

`func (o *HistoryRetention) GetTenMin() int32`

GetTenMin returns the TenMin field if non-nil, zero value otherwise.

### GetTenMinOk

`func (o *HistoryRetention) GetTenMinOk() (*int32, bool)`

GetTenMinOk returns a tuple with the TenMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenMin

`func (o *HistoryRetention) SetTenMin(v int32)`

SetTenMin sets TenMin field to given value.

### HasTenMin

`func (o *HistoryRetention) HasTenMin() bool`

HasTenMin returns a boolean if a field has been set.

### GetWeekly

`func (o *HistoryRetention) GetWeekly() int32`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *HistoryRetention) GetWeeklyOk() (*int32, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *HistoryRetention) SetWeekly(v int32)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *HistoryRetention) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetWeeklyAvailableRetentionDays

`func (o *HistoryRetention) GetWeeklyAvailableRetentionDays() []int32`

GetWeeklyAvailableRetentionDays returns the WeeklyAvailableRetentionDays field if non-nil, zero value otherwise.

### GetWeeklyAvailableRetentionDaysOk

`func (o *HistoryRetention) GetWeeklyAvailableRetentionDaysOk() (*[]int32, bool)`

GetWeeklyAvailableRetentionDaysOk returns a tuple with the WeeklyAvailableRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyAvailableRetentionDays

`func (o *HistoryRetention) SetWeeklyAvailableRetentionDays(v []int32)`

SetWeeklyAvailableRetentionDays sets WeeklyAvailableRetentionDays field to given value.

### HasWeeklyAvailableRetentionDays

`func (o *HistoryRetention) HasWeeklyAvailableRetentionDays() bool`

HasWeeklyAvailableRetentionDays returns a boolean if a field has been set.

### GetWidsData

`func (o *HistoryRetention) GetWidsData() int32`

GetWidsData returns the WidsData field if non-nil, zero value otherwise.

### GetWidsDataOk

`func (o *HistoryRetention) GetWidsDataOk() (*int32, bool)`

GetWidsDataOk returns a tuple with the WidsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidsData

`func (o *HistoryRetention) SetWidsData(v int32)`

SetWidsData sets WidsData field to given value.

### HasWidsData

`func (o *HistoryRetention) HasWidsData() bool`

HasWidsData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


