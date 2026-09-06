# InsightAnomalyStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **int32** | Total number of incidents. | [optional] 
**AllLevel** | Pointer to **int32** | Total number of incidents in level. | [optional] 
**AllStatus** | Pointer to **int32** | Total number of incidents in status. | [optional] 
**Critical** | Pointer to **int32** | Number of critical level incidents. | [optional] 
**Error** | Pointer to **int32** | Number of error level incidents. | [optional] 
**Ignored** | Pointer to **int32** | Number of ignored incidents. | [optional] 
**Info** | Pointer to **int32** | Number of info level incidents. | [optional] 
**Ongoing** | Pointer to **int32** | Number of ongoing incidents. | [optional] 
**Resolved** | Pointer to **int32** | Number of resolved incidents. | [optional] 
**Unresolved** | Pointer to **int32** | Number of unresolved incidents. | [optional] 
**Warning** | Pointer to **int32** | Number of warning level incidents. | [optional] 

## Methods

### NewInsightAnomalyStatVO

`func NewInsightAnomalyStatVO() *InsightAnomalyStatVO`

NewInsightAnomalyStatVO instantiates a new InsightAnomalyStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightAnomalyStatVOWithDefaults

`func NewInsightAnomalyStatVOWithDefaults() *InsightAnomalyStatVO`

NewInsightAnomalyStatVOWithDefaults instantiates a new InsightAnomalyStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *InsightAnomalyStatVO) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *InsightAnomalyStatVO) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *InsightAnomalyStatVO) SetAll(v int32)`

SetAll sets All field to given value.

### HasAll

`func (o *InsightAnomalyStatVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAllLevel

`func (o *InsightAnomalyStatVO) GetAllLevel() int32`

GetAllLevel returns the AllLevel field if non-nil, zero value otherwise.

### GetAllLevelOk

`func (o *InsightAnomalyStatVO) GetAllLevelOk() (*int32, bool)`

GetAllLevelOk returns a tuple with the AllLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLevel

`func (o *InsightAnomalyStatVO) SetAllLevel(v int32)`

SetAllLevel sets AllLevel field to given value.

### HasAllLevel

`func (o *InsightAnomalyStatVO) HasAllLevel() bool`

HasAllLevel returns a boolean if a field has been set.

### GetAllStatus

`func (o *InsightAnomalyStatVO) GetAllStatus() int32`

GetAllStatus returns the AllStatus field if non-nil, zero value otherwise.

### GetAllStatusOk

`func (o *InsightAnomalyStatVO) GetAllStatusOk() (*int32, bool)`

GetAllStatusOk returns a tuple with the AllStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus

`func (o *InsightAnomalyStatVO) SetAllStatus(v int32)`

SetAllStatus sets AllStatus field to given value.

### HasAllStatus

`func (o *InsightAnomalyStatVO) HasAllStatus() bool`

HasAllStatus returns a boolean if a field has been set.

### GetCritical

`func (o *InsightAnomalyStatVO) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *InsightAnomalyStatVO) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *InsightAnomalyStatVO) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *InsightAnomalyStatVO) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetError

`func (o *InsightAnomalyStatVO) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InsightAnomalyStatVO) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InsightAnomalyStatVO) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *InsightAnomalyStatVO) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIgnored

`func (o *InsightAnomalyStatVO) GetIgnored() int32`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *InsightAnomalyStatVO) GetIgnoredOk() (*int32, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *InsightAnomalyStatVO) SetIgnored(v int32)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *InsightAnomalyStatVO) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetInfo

`func (o *InsightAnomalyStatVO) GetInfo() int32`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InsightAnomalyStatVO) GetInfoOk() (*int32, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InsightAnomalyStatVO) SetInfo(v int32)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InsightAnomalyStatVO) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetOngoing

`func (o *InsightAnomalyStatVO) GetOngoing() int32`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *InsightAnomalyStatVO) GetOngoingOk() (*int32, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *InsightAnomalyStatVO) SetOngoing(v int32)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *InsightAnomalyStatVO) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.

### GetResolved

`func (o *InsightAnomalyStatVO) GetResolved() int32`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *InsightAnomalyStatVO) GetResolvedOk() (*int32, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *InsightAnomalyStatVO) SetResolved(v int32)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *InsightAnomalyStatVO) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetUnresolved

`func (o *InsightAnomalyStatVO) GetUnresolved() int32`

GetUnresolved returns the Unresolved field if non-nil, zero value otherwise.

### GetUnresolvedOk

`func (o *InsightAnomalyStatVO) GetUnresolvedOk() (*int32, bool)`

GetUnresolvedOk returns a tuple with the Unresolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolved

`func (o *InsightAnomalyStatVO) SetUnresolved(v int32)`

SetUnresolved sets Unresolved field to given value.

### HasUnresolved

`func (o *InsightAnomalyStatVO) HasUnresolved() bool`

HasUnresolved returns a boolean if a field has been set.

### GetWarning

`func (o *InsightAnomalyStatVO) GetWarning() int32`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *InsightAnomalyStatVO) GetWarningOk() (*int32, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *InsightAnomalyStatVO) SetWarning(v int32)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *InsightAnomalyStatVO) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


