# DeleteSiteEventLogListOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **int64** | The end timeStamp of the delete site event log, unit: MS. | 
**FilterModule** | Pointer to **string** | The module of the delete site event log; It is required when [selectType] is &#39;all&#39;, filterModule should be a value as follows: [&#39;System&#39;, &#39;Device&#39;, &#39;Client&#39;]. | [optional] 
**Logs** | Pointer to **[]string** | Select the logs to delete; Log ID list can be obtained from &#39;Get site event log list&#39; interface. | [optional] 
**SelectType** | **string** | Select type of logs. include: include selected logs, exclude: all but exclude selected logs, all: include all logs(Parameter [logs] need input &#39;[]&#39;). | 
**StartTime** | **int64** | The start timeStamp of the delete site event log, unit: MS. | 

## Methods

### NewDeleteSiteEventLogListOpenApiVO

`func NewDeleteSiteEventLogListOpenApiVO(endTime int64, selectType string, startTime int64, ) *DeleteSiteEventLogListOpenApiVO`

NewDeleteSiteEventLogListOpenApiVO instantiates a new DeleteSiteEventLogListOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSiteEventLogListOpenApiVOWithDefaults

`func NewDeleteSiteEventLogListOpenApiVOWithDefaults() *DeleteSiteEventLogListOpenApiVO`

NewDeleteSiteEventLogListOpenApiVOWithDefaults instantiates a new DeleteSiteEventLogListOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *DeleteSiteEventLogListOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DeleteSiteEventLogListOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DeleteSiteEventLogListOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetFilterModule

`func (o *DeleteSiteEventLogListOpenApiVO) GetFilterModule() string`

GetFilterModule returns the FilterModule field if non-nil, zero value otherwise.

### GetFilterModuleOk

`func (o *DeleteSiteEventLogListOpenApiVO) GetFilterModuleOk() (*string, bool)`

GetFilterModuleOk returns a tuple with the FilterModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterModule

`func (o *DeleteSiteEventLogListOpenApiVO) SetFilterModule(v string)`

SetFilterModule sets FilterModule field to given value.

### HasFilterModule

`func (o *DeleteSiteEventLogListOpenApiVO) HasFilterModule() bool`

HasFilterModule returns a boolean if a field has been set.

### GetLogs

`func (o *DeleteSiteEventLogListOpenApiVO) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DeleteSiteEventLogListOpenApiVO) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DeleteSiteEventLogListOpenApiVO) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *DeleteSiteEventLogListOpenApiVO) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetSelectType

`func (o *DeleteSiteEventLogListOpenApiVO) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *DeleteSiteEventLogListOpenApiVO) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *DeleteSiteEventLogListOpenApiVO) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.


### GetStartTime

`func (o *DeleteSiteEventLogListOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DeleteSiteEventLogListOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DeleteSiteEventLogListOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


