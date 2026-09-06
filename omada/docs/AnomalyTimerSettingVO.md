# AnomalyTimerSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Event category filter, comma-separated. Required when selectAll is true. | [optional] 
**End** | Pointer to **int64** | End time in seconds (Unix timestamp). | [optional] 
**EndTime** | **int64** | End time in milliseconds (Unix timestamp). | 
**Incidents** | Pointer to [**[]AnomalyBriefQueryVO**](AnomalyBriefQueryVO.md) | List of anomaly event objects to operate on, identified by anomaly code and MAC. | [optional] 
**Level** | Pointer to **int32** | Event severity level filter. 0: Critical, 1: Error, 2: Warning, 3: Info. | [optional] 
**ObjectType** | Pointer to **string** | Object type filter, comma-separated (e.g. &#39;ap,wirelessClient&#39;). Required when selectAll is true. | [optional] 
**SearchKey** | Pointer to **string** | Search keyword for event content, device name/MAC/IP. Required when selectAll is true. | [optional] 
**SelectAll** | Pointer to **bool** | Whether to select all incidents. If true, &#39;incidents&#39; list items are excluded; otherwise only listed items are processed. | [optional] 
**Start** | Pointer to **int64** | Start time in seconds (Unix timestamp). | [optional] 
**StartTime** | **int64** | Start time in milliseconds (Unix timestamp). | 
**Status** | **int32** | Current event status. 0: Unresolved, 1: Resolved, 2: Ignored. 3: Ongoing | 
**TargetStatus** | Pointer to **int32** | Target event status for modification. 0: Unresolved 1: Resolved, 2: Ignored. Not required for deletion. | [optional] 

## Methods

### NewAnomalyTimerSettingVO

`func NewAnomalyTimerSettingVO(endTime int64, startTime int64, status int32, ) *AnomalyTimerSettingVO`

NewAnomalyTimerSettingVO instantiates a new AnomalyTimerSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyTimerSettingVOWithDefaults

`func NewAnomalyTimerSettingVOWithDefaults() *AnomalyTimerSettingVO`

NewAnomalyTimerSettingVOWithDefaults instantiates a new AnomalyTimerSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AnomalyTimerSettingVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AnomalyTimerSettingVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AnomalyTimerSettingVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AnomalyTimerSettingVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnd

`func (o *AnomalyTimerSettingVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AnomalyTimerSettingVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AnomalyTimerSettingVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AnomalyTimerSettingVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEndTime

`func (o *AnomalyTimerSettingVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AnomalyTimerSettingVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AnomalyTimerSettingVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetIncidents

`func (o *AnomalyTimerSettingVO) GetIncidents() []AnomalyBriefQueryVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *AnomalyTimerSettingVO) GetIncidentsOk() (*[]AnomalyBriefQueryVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *AnomalyTimerSettingVO) SetIncidents(v []AnomalyBriefQueryVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *AnomalyTimerSettingVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetLevel

`func (o *AnomalyTimerSettingVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnomalyTimerSettingVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnomalyTimerSettingVO) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AnomalyTimerSettingVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetObjectType

`func (o *AnomalyTimerSettingVO) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AnomalyTimerSettingVO) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AnomalyTimerSettingVO) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AnomalyTimerSettingVO) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetSearchKey

`func (o *AnomalyTimerSettingVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *AnomalyTimerSettingVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *AnomalyTimerSettingVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *AnomalyTimerSettingVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectAll

`func (o *AnomalyTimerSettingVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *AnomalyTimerSettingVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *AnomalyTimerSettingVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *AnomalyTimerSettingVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.

### GetStart

`func (o *AnomalyTimerSettingVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AnomalyTimerSettingVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AnomalyTimerSettingVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *AnomalyTimerSettingVO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStartTime

`func (o *AnomalyTimerSettingVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AnomalyTimerSettingVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AnomalyTimerSettingVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetStatus

`func (o *AnomalyTimerSettingVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnomalyTimerSettingVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnomalyTimerSettingVO) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTargetStatus

`func (o *AnomalyTimerSettingVO) GetTargetStatus() int32`

GetTargetStatus returns the TargetStatus field if non-nil, zero value otherwise.

### GetTargetStatusOk

`func (o *AnomalyTimerSettingVO) GetTargetStatusOk() (*int32, bool)`

GetTargetStatusOk returns a tuple with the TargetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatus

`func (o *AnomalyTimerSettingVO) SetTargetStatus(v int32)`

SetTargetStatus sets TargetStatus field to given value.

### HasTargetStatus

`func (o *AnomalyTimerSettingVO) HasTargetStatus() bool`

HasTargetStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


