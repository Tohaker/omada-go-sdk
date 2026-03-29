# GridVORogueAPScanResultEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]RogueAPScanResultEntry**](RogueAPScanResultEntry.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewGridVORogueAPScanResultEntry

`func NewGridVORogueAPScanResultEntry() *GridVORogueAPScanResultEntry`

NewGridVORogueAPScanResultEntry instantiates a new GridVORogueAPScanResultEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridVORogueAPScanResultEntryWithDefaults

`func NewGridVORogueAPScanResultEntryWithDefaults() *GridVORogueAPScanResultEntry`

NewGridVORogueAPScanResultEntryWithDefaults instantiates a new GridVORogueAPScanResultEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *GridVORogueAPScanResultEntry) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *GridVORogueAPScanResultEntry) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *GridVORogueAPScanResultEntry) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *GridVORogueAPScanResultEntry) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *GridVORogueAPScanResultEntry) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *GridVORogueAPScanResultEntry) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *GridVORogueAPScanResultEntry) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *GridVORogueAPScanResultEntry) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *GridVORogueAPScanResultEntry) GetData() []RogueAPScanResultEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GridVORogueAPScanResultEntry) GetDataOk() (*[]RogueAPScanResultEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GridVORogueAPScanResultEntry) SetData(v []RogueAPScanResultEntry)`

SetData sets Data field to given value.

### HasData

`func (o *GridVORogueAPScanResultEntry) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *GridVORogueAPScanResultEntry) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *GridVORogueAPScanResultEntry) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *GridVORogueAPScanResultEntry) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *GridVORogueAPScanResultEntry) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


