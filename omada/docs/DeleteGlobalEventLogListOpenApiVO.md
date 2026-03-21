# DeleteGlobalEventLogListOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **int64** | The end timeStamp of the delete global event log, unit: MS. | 
**FilterModule** | Pointer to **string** | The module of the delete global log; It is required when [selectType] is &#39;all&#39;, filterModule should be a value as follows: &#39;System&#39; or &#39;Device&#39;. | [optional] 
**Logs** | Pointer to **[]string** | Select the logs to delete; Log ID list can be obtained from &#39;Get global event log list&#39; interface. | [optional] 
**SelectType** | **string** | Select type of logs. include: include selected logs, exclude: all but exclude selected logs, all: include all logs(Parameter [logs] need input &#39;[]&#39;). | 
**StartTime** | **int64** | The start timeStamp of the delete global event log, unit: MS. | 

## Methods

### NewDeleteGlobalEventLogListOpenApiVO

`func NewDeleteGlobalEventLogListOpenApiVO(endTime int64, selectType string, startTime int64, ) *DeleteGlobalEventLogListOpenApiVO`

NewDeleteGlobalEventLogListOpenApiVO instantiates a new DeleteGlobalEventLogListOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGlobalEventLogListOpenApiVOWithDefaults

`func NewDeleteGlobalEventLogListOpenApiVOWithDefaults() *DeleteGlobalEventLogListOpenApiVO`

NewDeleteGlobalEventLogListOpenApiVOWithDefaults instantiates a new DeleteGlobalEventLogListOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *DeleteGlobalEventLogListOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DeleteGlobalEventLogListOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DeleteGlobalEventLogListOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetFilterModule

`func (o *DeleteGlobalEventLogListOpenApiVO) GetFilterModule() string`

GetFilterModule returns the FilterModule field if non-nil, zero value otherwise.

### GetFilterModuleOk

`func (o *DeleteGlobalEventLogListOpenApiVO) GetFilterModuleOk() (*string, bool)`

GetFilterModuleOk returns a tuple with the FilterModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterModule

`func (o *DeleteGlobalEventLogListOpenApiVO) SetFilterModule(v string)`

SetFilterModule sets FilterModule field to given value.

### HasFilterModule

`func (o *DeleteGlobalEventLogListOpenApiVO) HasFilterModule() bool`

HasFilterModule returns a boolean if a field has been set.

### GetLogs

`func (o *DeleteGlobalEventLogListOpenApiVO) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DeleteGlobalEventLogListOpenApiVO) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DeleteGlobalEventLogListOpenApiVO) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *DeleteGlobalEventLogListOpenApiVO) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetSelectType

`func (o *DeleteGlobalEventLogListOpenApiVO) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *DeleteGlobalEventLogListOpenApiVO) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *DeleteGlobalEventLogListOpenApiVO) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.


### GetStartTime

`func (o *DeleteGlobalEventLogListOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DeleteGlobalEventLogListOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DeleteGlobalEventLogListOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


