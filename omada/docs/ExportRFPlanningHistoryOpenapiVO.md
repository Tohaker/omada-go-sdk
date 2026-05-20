# ExportRFPlanningHistoryOpenapiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** | End time of history to export (millisecond timestamp) | [optional] 
**ExportFormat** | **int32** | Export file format, should be a value as follows: 0: csv, 1: xlsx | 
**FilterMode** | Pointer to **int32** | filtermode should be a value as follows: 1: rfPlanningHistory mode is manual; 2:rfPlanningHistory mode is adaptive | [optional] 
**HistoryIds** | Pointer to **[]string** | RfPlanningHistoryId list | [optional] 
**Start** | Pointer to **int64** | Start time of history to export (millisecond timestamp) | [optional] 
**Type** | Pointer to **string** | Export policy, type should be a value as follows:, all: export all history, include: export history in ids, exclude: exclude history in ids. | [optional] 

## Methods

### NewExportRFPlanningHistoryOpenapiVO

`func NewExportRFPlanningHistoryOpenapiVO(exportFormat int32, ) *ExportRFPlanningHistoryOpenapiVO`

NewExportRFPlanningHistoryOpenapiVO instantiates a new ExportRFPlanningHistoryOpenapiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportRFPlanningHistoryOpenapiVOWithDefaults

`func NewExportRFPlanningHistoryOpenapiVOWithDefaults() *ExportRFPlanningHistoryOpenapiVO`

NewExportRFPlanningHistoryOpenapiVOWithDefaults instantiates a new ExportRFPlanningHistoryOpenapiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ExportRFPlanningHistoryOpenapiVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ExportRFPlanningHistoryOpenapiVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ExportRFPlanningHistoryOpenapiVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ExportRFPlanningHistoryOpenapiVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExportFormat

`func (o *ExportRFPlanningHistoryOpenapiVO) GetExportFormat() int32`

GetExportFormat returns the ExportFormat field if non-nil, zero value otherwise.

### GetExportFormatOk

`func (o *ExportRFPlanningHistoryOpenapiVO) GetExportFormatOk() (*int32, bool)`

GetExportFormatOk returns a tuple with the ExportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFormat

`func (o *ExportRFPlanningHistoryOpenapiVO) SetExportFormat(v int32)`

SetExportFormat sets ExportFormat field to given value.


### GetFilterMode

`func (o *ExportRFPlanningHistoryOpenapiVO) GetFilterMode() int32`

GetFilterMode returns the FilterMode field if non-nil, zero value otherwise.

### GetFilterModeOk

`func (o *ExportRFPlanningHistoryOpenapiVO) GetFilterModeOk() (*int32, bool)`

GetFilterModeOk returns a tuple with the FilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMode

`func (o *ExportRFPlanningHistoryOpenapiVO) SetFilterMode(v int32)`

SetFilterMode sets FilterMode field to given value.

### HasFilterMode

`func (o *ExportRFPlanningHistoryOpenapiVO) HasFilterMode() bool`

HasFilterMode returns a boolean if a field has been set.

### GetHistoryIds

`func (o *ExportRFPlanningHistoryOpenapiVO) GetHistoryIds() []string`

GetHistoryIds returns the HistoryIds field if non-nil, zero value otherwise.

### GetHistoryIdsOk

`func (o *ExportRFPlanningHistoryOpenapiVO) GetHistoryIdsOk() (*[]string, bool)`

GetHistoryIdsOk returns a tuple with the HistoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryIds

`func (o *ExportRFPlanningHistoryOpenapiVO) SetHistoryIds(v []string)`

SetHistoryIds sets HistoryIds field to given value.

### HasHistoryIds

`func (o *ExportRFPlanningHistoryOpenapiVO) HasHistoryIds() bool`

HasHistoryIds returns a boolean if a field has been set.

### GetStart

`func (o *ExportRFPlanningHistoryOpenapiVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ExportRFPlanningHistoryOpenapiVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ExportRFPlanningHistoryOpenapiVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ExportRFPlanningHistoryOpenapiVO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetType

`func (o *ExportRFPlanningHistoryOpenapiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportRFPlanningHistoryOpenapiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportRFPlanningHistoryOpenapiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExportRFPlanningHistoryOpenapiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


