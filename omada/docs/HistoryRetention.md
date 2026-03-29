# HistoryRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDataTrendDaily** | Pointer to **int32** | Retention Configuration of client data trend records(only effective in local controller), clientStatDaily should be a value as follows: 1: 1day; 7: 7days; 31: 31days; 60: 60days, 90: 90days. | [optional] 
**ClientDataTrendEnable** | Pointer to **bool** | Whether the client data trend records is recorded. | [optional] 
**ClientHistory** | Pointer to **int32** | Retention configuration of client History(only effective in local controller), clientHistory should be a value as follows: -1: Disabled; 0: All Time(Windows, Linux Only); 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**ClientsDataEnable** | **bool** | Whether the clients&#39; history data is recorded. | 
**Daily** | Pointer to **int32** | Retention configuration of time series with daily granularity, daily should be a value as follows: 90: 90days; 180: 180days; 365: 365days(Fixed value in Cloud Based Controller as 365 days). | [optional] 
**FiveMin** | Pointer to **int32** | Retention configuration of Time Series with 5 Minutes Granularity. It is fixed to 2days and cannot be changed. | [optional] 
**Hourly** | Pointer to **int32** | Retention configuration of time series with hourly granularity, hourly should be a value as follows: 7: 7days. | [optional] 
**KnownClient** | Pointer to **int32** | Retention configuration of known client Data, knownClient should be a value as follows: -1: Disabled; 0: All Time(Windows, Linux Only); 1: 1day; 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**Log** | Pointer to **int32** | Retention Configuration of log data(only effective in local controller), log should be a value as follows: 0: All Time(Windows, Linux Only); 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**Override** | Pointer to **bool** | Whether the customer overrides the retention configuration of MSP. | [optional] 
**PortalAuth** | Pointer to **int32** | Retention configuration of portal authentication records, portalAuth should be a value as follows: 0: All Time(Windows, Linux Only); 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**RogueAp** | Pointer to **int32** | Retention Configuration of rogue ap data, rogueAp should be a value as follows: 0: All Time(Windows, Linux Only); 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**TenMin** | Pointer to **int32** | Retention configuration of Time Series with 10 Minutes Granularity, only for Lite Cloud-Based Controller. It is fixed to 2days and cannot be changed. | [optional] 
**Weekly** | Pointer to **int32** | Retention configuration of time series with weekly granularity, weekly should be a value as follows: 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
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


