# DeleteAnomalyListSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | Pointer to **string** | Anomaly event code. Required when selectAll is true. | [optional] 
**End** | Pointer to **int64** | End time in seconds. If both start and end are omitted, the server defaults to the last 30 days. | [optional] 
**Incidents** | Pointer to **[]string** | List of anomaly event IDs to operate on. | [optional] 
**Object** | Pointer to **string** | Device or client MAC, multiple MACs separated by comma. Required when selectAll is true. | [optional] 
**SelectAll** | Pointer to **bool** | Whether to select all incidents. If true, &#39;incidents&#39; list items are excluded; otherwise only listed items are processed. | [optional] 
**Start** | Pointer to **int64** | Start time in seconds. If both start and end are omitted, the server defaults to the last 30 days. | [optional] 
**Status** | Pointer to **int32** | Current event status. 0: Unresolved, 1: Resolved, 2: Ignored. 3：Ongoing | [optional] 

## Methods

### NewDeleteAnomalyListSettingVO

`func NewDeleteAnomalyListSettingVO() *DeleteAnomalyListSettingVO`

NewDeleteAnomalyListSettingVO instantiates a new DeleteAnomalyListSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAnomalyListSettingVOWithDefaults

`func NewDeleteAnomalyListSettingVOWithDefaults() *DeleteAnomalyListSettingVO`

NewDeleteAnomalyListSettingVOWithDefaults instantiates a new DeleteAnomalyListSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *DeleteAnomalyListSettingVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *DeleteAnomalyListSettingVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *DeleteAnomalyListSettingVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.

### HasAnomalyCode

`func (o *DeleteAnomalyListSettingVO) HasAnomalyCode() bool`

HasAnomalyCode returns a boolean if a field has been set.

### GetEnd

`func (o *DeleteAnomalyListSettingVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DeleteAnomalyListSettingVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DeleteAnomalyListSettingVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DeleteAnomalyListSettingVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetIncidents

`func (o *DeleteAnomalyListSettingVO) GetIncidents() []string`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *DeleteAnomalyListSettingVO) GetIncidentsOk() (*[]string, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *DeleteAnomalyListSettingVO) SetIncidents(v []string)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *DeleteAnomalyListSettingVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetObject

`func (o *DeleteAnomalyListSettingVO) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DeleteAnomalyListSettingVO) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DeleteAnomalyListSettingVO) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *DeleteAnomalyListSettingVO) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSelectAll

`func (o *DeleteAnomalyListSettingVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *DeleteAnomalyListSettingVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *DeleteAnomalyListSettingVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *DeleteAnomalyListSettingVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.

### GetStart

`func (o *DeleteAnomalyListSettingVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DeleteAnomalyListSettingVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DeleteAnomalyListSettingVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *DeleteAnomalyListSettingVO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *DeleteAnomalyListSettingVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteAnomalyListSettingVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteAnomalyListSettingVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeleteAnomalyListSettingVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


