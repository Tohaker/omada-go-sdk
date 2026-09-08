# AnomalyListSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | Pointer to **string** | Anomaly event code. Required when selectAll is true. | [optional] 
**End** | Pointer to **int64** | End time in seconds. If both start and end are omitted, the server defaults to the last 30 days. | [optional] 
**Incidents** | Pointer to **[]string** | List of anomaly event IDs to operate on. IDs are from &#39;Get grid incident List&#39; API. | [optional] 
**Object** | Pointer to **string** | Device or client MAC, multiple MACs separated by comma. Required when selectAll is true. | [optional] 
**SelectAll** | Pointer to **bool** | Whether to select all incidents. If true, &#39;incidents&#39; list items are excluded; otherwise only listed items are processed. | [optional] 
**Start** | Pointer to **int64** | Start time in seconds. If both start and end are omitted, the server defaults to the last 30 days. | [optional] 
**Status** | Pointer to **int32** | Current event status. 0: Unresolved, 1: Resolved, 2: Ignored. 3：Ongoing | [optional] 
**TargetStatus** | Pointer to **int32** | Target event status for modification. 0: Unresolved 1: Resolved, 2: Ignored. Not required for deletion. | [optional] 

## Methods

### NewAnomalyListSettingVO

`func NewAnomalyListSettingVO() *AnomalyListSettingVO`

NewAnomalyListSettingVO instantiates a new AnomalyListSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyListSettingVOWithDefaults

`func NewAnomalyListSettingVOWithDefaults() *AnomalyListSettingVO`

NewAnomalyListSettingVOWithDefaults instantiates a new AnomalyListSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *AnomalyListSettingVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *AnomalyListSettingVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *AnomalyListSettingVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.

### HasAnomalyCode

`func (o *AnomalyListSettingVO) HasAnomalyCode() bool`

HasAnomalyCode returns a boolean if a field has been set.

### GetEnd

`func (o *AnomalyListSettingVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AnomalyListSettingVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AnomalyListSettingVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AnomalyListSettingVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetIncidents

`func (o *AnomalyListSettingVO) GetIncidents() []string`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *AnomalyListSettingVO) GetIncidentsOk() (*[]string, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *AnomalyListSettingVO) SetIncidents(v []string)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *AnomalyListSettingVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetObject

`func (o *AnomalyListSettingVO) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AnomalyListSettingVO) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AnomalyListSettingVO) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AnomalyListSettingVO) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSelectAll

`func (o *AnomalyListSettingVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *AnomalyListSettingVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *AnomalyListSettingVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *AnomalyListSettingVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.

### GetStart

`func (o *AnomalyListSettingVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AnomalyListSettingVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AnomalyListSettingVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *AnomalyListSettingVO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *AnomalyListSettingVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnomalyListSettingVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnomalyListSettingVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnomalyListSettingVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetStatus

`func (o *AnomalyListSettingVO) GetTargetStatus() int32`

GetTargetStatus returns the TargetStatus field if non-nil, zero value otherwise.

### GetTargetStatusOk

`func (o *AnomalyListSettingVO) GetTargetStatusOk() (*int32, bool)`

GetTargetStatusOk returns a tuple with the TargetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatus

`func (o *AnomalyListSettingVO) SetTargetStatus(v int32)`

SetTargetStatus sets TargetStatus field to given value.

### HasTargetStatus

`func (o *AnomalyListSettingVO) HasTargetStatus() bool`

HasTargetStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


