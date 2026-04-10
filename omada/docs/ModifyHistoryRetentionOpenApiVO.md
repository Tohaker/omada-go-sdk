# ModifyHistoryRetentionOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDataTrendDaily** | Pointer to **int32** | Retention Configuration of client data trend records(only effective in local controller), clientDataTrendDaily should be a value as follows: 1: 1day; 7: 7days; 31: 31days; 60: 60days, 90: 90days. | [optional] 
**ClientDataTrendEnable** | Pointer to **bool** | Whether the client data trend records is recorded. | [optional] 
**ClientHealthEnable** | Pointer to **bool** | Whether client health is enabled. When enabled, client health data will be recorded, which may consume a significant amount of storage space. | [optional] 
**ClientHistory** | Pointer to **int32** | Retention configuration of client History(only effective in local controller), clientHistory should be a value as follows: -1: Disabled; 0: All Time(Windows, Linux Only); 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**ClientRecognitionEnable** | Pointer to **bool** | Whether client recognition is enabled. With the feature enabled, network devices will report client information in real time to ensure the accuracy of client recognition. Cloud Access is required for client recognition. | [optional] 
**ClientsDataEnable** | **bool** | Whether the clients&#39; history data is recorded. | 
**Daily** | Pointer to **int32** | Retention configuration of time series with daily granularity, daily should be a value as follows: 90: 90days; 180: 180days; 365: 365days(Fixed value in Cloud Based Controller as 365 days). | [optional] 
**KnownClient** | Pointer to **int32** | Retention configuration of known client Data, knownClient should be a value as follows: -1: Disabled; 0: All Time(Windows, Linux Only); 1: 1day; 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**Log** | Pointer to **int32** | Retention Configuration of log data(only effective in local controller), log should be a value as follows: 0: All Time(Windows, Linux Only); 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**Override** | Pointer to **bool** | Whether the customer overrides the retention configuration of MSP. This configuration applies to the customer in MSP mode only. If true, customer retention configuration will be used (excluding daily, clientHistory, log and widsData). | [optional] 
**PortalAuth** | Pointer to **int32** | Retention configuration of portal authentication records, portalAuth should be a value as follows: 0: All Time(Windows, Linux Only); 7: 7days; 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**RogueAp** | Pointer to **int32** | Retention Configuration of rogue ap data, rogueAp should be a value as follows: 0: All Time(Windows, Linux Only); 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**Weekly** | Pointer to **int32** | Retention configuration of time series with weekly granularity, weekly should be a value as follows: 31: 31days; 90: 90days; 180: 180days; 365: 365days. | [optional] 
**WidsData** | Pointer to **int32** | Retention Configuration of wids data(only effective in local pro controller), widsData should be a value as follows: 0: All Time(Windows, Linux Only); 90: 90days; 180: 180days; 365: 365days. | [optional] 

## Methods

### NewModifyHistoryRetentionOpenApiVO

`func NewModifyHistoryRetentionOpenApiVO(clientsDataEnable bool, ) *ModifyHistoryRetentionOpenApiVO`

NewModifyHistoryRetentionOpenApiVO instantiates a new ModifyHistoryRetentionOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyHistoryRetentionOpenApiVOWithDefaults

`func NewModifyHistoryRetentionOpenApiVOWithDefaults() *ModifyHistoryRetentionOpenApiVO`

NewModifyHistoryRetentionOpenApiVOWithDefaults instantiates a new ModifyHistoryRetentionOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientDataTrendDaily

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientDataTrendDaily() int32`

GetClientDataTrendDaily returns the ClientDataTrendDaily field if non-nil, zero value otherwise.

### GetClientDataTrendDailyOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientDataTrendDailyOk() (*int32, bool)`

GetClientDataTrendDailyOk returns a tuple with the ClientDataTrendDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDataTrendDaily

`func (o *ModifyHistoryRetentionOpenApiVO) SetClientDataTrendDaily(v int32)`

SetClientDataTrendDaily sets ClientDataTrendDaily field to given value.

### HasClientDataTrendDaily

`func (o *ModifyHistoryRetentionOpenApiVO) HasClientDataTrendDaily() bool`

HasClientDataTrendDaily returns a boolean if a field has been set.

### GetClientDataTrendEnable

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientDataTrendEnable() bool`

GetClientDataTrendEnable returns the ClientDataTrendEnable field if non-nil, zero value otherwise.

### GetClientDataTrendEnableOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientDataTrendEnableOk() (*bool, bool)`

GetClientDataTrendEnableOk returns a tuple with the ClientDataTrendEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDataTrendEnable

`func (o *ModifyHistoryRetentionOpenApiVO) SetClientDataTrendEnable(v bool)`

SetClientDataTrendEnable sets ClientDataTrendEnable field to given value.

### HasClientDataTrendEnable

`func (o *ModifyHistoryRetentionOpenApiVO) HasClientDataTrendEnable() bool`

HasClientDataTrendEnable returns a boolean if a field has been set.

### GetClientHealthEnable

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientHealthEnable() bool`

GetClientHealthEnable returns the ClientHealthEnable field if non-nil, zero value otherwise.

### GetClientHealthEnableOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientHealthEnableOk() (*bool, bool)`

GetClientHealthEnableOk returns a tuple with the ClientHealthEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealthEnable

`func (o *ModifyHistoryRetentionOpenApiVO) SetClientHealthEnable(v bool)`

SetClientHealthEnable sets ClientHealthEnable field to given value.

### HasClientHealthEnable

`func (o *ModifyHistoryRetentionOpenApiVO) HasClientHealthEnable() bool`

HasClientHealthEnable returns a boolean if a field has been set.

### GetClientHistory

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientHistory() int32`

GetClientHistory returns the ClientHistory field if non-nil, zero value otherwise.

### GetClientHistoryOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientHistoryOk() (*int32, bool)`

GetClientHistoryOk returns a tuple with the ClientHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHistory

`func (o *ModifyHistoryRetentionOpenApiVO) SetClientHistory(v int32)`

SetClientHistory sets ClientHistory field to given value.

### HasClientHistory

`func (o *ModifyHistoryRetentionOpenApiVO) HasClientHistory() bool`

HasClientHistory returns a boolean if a field has been set.

### GetClientRecognitionEnable

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientRecognitionEnable() bool`

GetClientRecognitionEnable returns the ClientRecognitionEnable field if non-nil, zero value otherwise.

### GetClientRecognitionEnableOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientRecognitionEnableOk() (*bool, bool)`

GetClientRecognitionEnableOk returns a tuple with the ClientRecognitionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRecognitionEnable

`func (o *ModifyHistoryRetentionOpenApiVO) SetClientRecognitionEnable(v bool)`

SetClientRecognitionEnable sets ClientRecognitionEnable field to given value.

### HasClientRecognitionEnable

`func (o *ModifyHistoryRetentionOpenApiVO) HasClientRecognitionEnable() bool`

HasClientRecognitionEnable returns a boolean if a field has been set.

### GetClientsDataEnable

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientsDataEnable() bool`

GetClientsDataEnable returns the ClientsDataEnable field if non-nil, zero value otherwise.

### GetClientsDataEnableOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetClientsDataEnableOk() (*bool, bool)`

GetClientsDataEnableOk returns a tuple with the ClientsDataEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsDataEnable

`func (o *ModifyHistoryRetentionOpenApiVO) SetClientsDataEnable(v bool)`

SetClientsDataEnable sets ClientsDataEnable field to given value.


### GetDaily

`func (o *ModifyHistoryRetentionOpenApiVO) GetDaily() int32`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetDailyOk() (*int32, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *ModifyHistoryRetentionOpenApiVO) SetDaily(v int32)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *ModifyHistoryRetentionOpenApiVO) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetKnownClient

`func (o *ModifyHistoryRetentionOpenApiVO) GetKnownClient() int32`

GetKnownClient returns the KnownClient field if non-nil, zero value otherwise.

### GetKnownClientOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetKnownClientOk() (*int32, bool)`

GetKnownClientOk returns a tuple with the KnownClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownClient

`func (o *ModifyHistoryRetentionOpenApiVO) SetKnownClient(v int32)`

SetKnownClient sets KnownClient field to given value.

### HasKnownClient

`func (o *ModifyHistoryRetentionOpenApiVO) HasKnownClient() bool`

HasKnownClient returns a boolean if a field has been set.

### GetLog

`func (o *ModifyHistoryRetentionOpenApiVO) GetLog() int32`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetLogOk() (*int32, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ModifyHistoryRetentionOpenApiVO) SetLog(v int32)`

SetLog sets Log field to given value.

### HasLog

`func (o *ModifyHistoryRetentionOpenApiVO) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetOverride

`func (o *ModifyHistoryRetentionOpenApiVO) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *ModifyHistoryRetentionOpenApiVO) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *ModifyHistoryRetentionOpenApiVO) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetPortalAuth

`func (o *ModifyHistoryRetentionOpenApiVO) GetPortalAuth() int32`

GetPortalAuth returns the PortalAuth field if non-nil, zero value otherwise.

### GetPortalAuthOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetPortalAuthOk() (*int32, bool)`

GetPortalAuthOk returns a tuple with the PortalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalAuth

`func (o *ModifyHistoryRetentionOpenApiVO) SetPortalAuth(v int32)`

SetPortalAuth sets PortalAuth field to given value.

### HasPortalAuth

`func (o *ModifyHistoryRetentionOpenApiVO) HasPortalAuth() bool`

HasPortalAuth returns a boolean if a field has been set.

### GetRogueAp

`func (o *ModifyHistoryRetentionOpenApiVO) GetRogueAp() int32`

GetRogueAp returns the RogueAp field if non-nil, zero value otherwise.

### GetRogueApOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetRogueApOk() (*int32, bool)`

GetRogueApOk returns a tuple with the RogueAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRogueAp

`func (o *ModifyHistoryRetentionOpenApiVO) SetRogueAp(v int32)`

SetRogueAp sets RogueAp field to given value.

### HasRogueAp

`func (o *ModifyHistoryRetentionOpenApiVO) HasRogueAp() bool`

HasRogueAp returns a boolean if a field has been set.

### GetWeekly

`func (o *ModifyHistoryRetentionOpenApiVO) GetWeekly() int32`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetWeeklyOk() (*int32, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *ModifyHistoryRetentionOpenApiVO) SetWeekly(v int32)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *ModifyHistoryRetentionOpenApiVO) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetWidsData

`func (o *ModifyHistoryRetentionOpenApiVO) GetWidsData() int32`

GetWidsData returns the WidsData field if non-nil, zero value otherwise.

### GetWidsDataOk

`func (o *ModifyHistoryRetentionOpenApiVO) GetWidsDataOk() (*int32, bool)`

GetWidsDataOk returns a tuple with the WidsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidsData

`func (o *ModifyHistoryRetentionOpenApiVO) SetWidsData(v int32)`

SetWidsData sets WidsData field to given value.

### HasWidsData

`func (o *ModifyHistoryRetentionOpenApiVO) HasWidsData() bool`

HasWidsData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


