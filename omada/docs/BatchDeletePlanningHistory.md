# BatchDeletePlanningHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** | End time of the history to be deleted(millisecond timestamp) | [optional] 
**FilterMode** | Pointer to **int32** | filtermode should be a value as follows: 1: rfPlanningHistory mode is manual; 2:rfPlanningHistory mode is adaptive | [optional] 
**Ids** | Pointer to **[]string** | HistoryId list | [optional] 
**Start** | Pointer to **int64** | Start time of the history to be deleted(millisecond timestamp) | [optional] 
**Type** | Pointer to **string** | Delete policy, all: delete all history, include: delete history in ids, exclude: exclude history in ids. | [optional] 

## Methods

### NewBatchDeletePlanningHistory

`func NewBatchDeletePlanningHistory() *BatchDeletePlanningHistory`

NewBatchDeletePlanningHistory instantiates a new BatchDeletePlanningHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchDeletePlanningHistoryWithDefaults

`func NewBatchDeletePlanningHistoryWithDefaults() *BatchDeletePlanningHistory`

NewBatchDeletePlanningHistoryWithDefaults instantiates a new BatchDeletePlanningHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *BatchDeletePlanningHistory) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *BatchDeletePlanningHistory) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *BatchDeletePlanningHistory) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *BatchDeletePlanningHistory) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFilterMode

`func (o *BatchDeletePlanningHistory) GetFilterMode() int32`

GetFilterMode returns the FilterMode field if non-nil, zero value otherwise.

### GetFilterModeOk

`func (o *BatchDeletePlanningHistory) GetFilterModeOk() (*int32, bool)`

GetFilterModeOk returns a tuple with the FilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMode

`func (o *BatchDeletePlanningHistory) SetFilterMode(v int32)`

SetFilterMode sets FilterMode field to given value.

### HasFilterMode

`func (o *BatchDeletePlanningHistory) HasFilterMode() bool`

HasFilterMode returns a boolean if a field has been set.

### GetIds

`func (o *BatchDeletePlanningHistory) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchDeletePlanningHistory) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchDeletePlanningHistory) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BatchDeletePlanningHistory) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetStart

`func (o *BatchDeletePlanningHistory) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BatchDeletePlanningHistory) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BatchDeletePlanningHistory) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *BatchDeletePlanningHistory) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetType

`func (o *BatchDeletePlanningHistory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BatchDeletePlanningHistory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BatchDeletePlanningHistory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BatchDeletePlanningHistory) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


