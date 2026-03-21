# ModelFirmwarePoolGridInfoModelLatestFwInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]ModelLatestFwInfo**](ModelLatestFwInfo.md) |  | [optional] 
**FwNum** | Pointer to **map[string]int32** | The firmware quantity map for each channel: The key is channel: (0: stable; 1: Release Candidate(RC); 2: Beta), The Value is the quantity | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewModelFirmwarePoolGridInfoModelLatestFwInfo

`func NewModelFirmwarePoolGridInfoModelLatestFwInfo() *ModelFirmwarePoolGridInfoModelLatestFwInfo`

NewModelFirmwarePoolGridInfoModelLatestFwInfo instantiates a new ModelFirmwarePoolGridInfoModelLatestFwInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelFirmwarePoolGridInfoModelLatestFwInfoWithDefaults

`func NewModelFirmwarePoolGridInfoModelLatestFwInfoWithDefaults() *ModelFirmwarePoolGridInfoModelLatestFwInfo`

NewModelFirmwarePoolGridInfoModelLatestFwInfoWithDefaults instantiates a new ModelFirmwarePoolGridInfoModelLatestFwInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetData() []ModelLatestFwInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetDataOk() (*[]ModelLatestFwInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) SetData(v []ModelLatestFwInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFwNum

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetFwNum() map[string]int32`

GetFwNum returns the FwNum field if non-nil, zero value otherwise.

### GetFwNumOk

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetFwNumOk() (*map[string]int32, bool)`

GetFwNumOk returns a tuple with the FwNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwNum

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) SetFwNum(v map[string]int32)`

SetFwNum sets FwNum field to given value.

### HasFwNum

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) HasFwNum() bool`

HasFwNum returns a boolean if a field has been set.

### GetTotalRows

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *ModelFirmwarePoolGridInfoModelLatestFwInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


