# GridVOUpgradeLogOpenApiInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]UpgradeLogOpenApiInfo**](UpgradeLogOpenApiInfo.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewGridVOUpgradeLogOpenApiInfo

`func NewGridVOUpgradeLogOpenApiInfo() *GridVOUpgradeLogOpenApiInfo`

NewGridVOUpgradeLogOpenApiInfo instantiates a new GridVOUpgradeLogOpenApiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridVOUpgradeLogOpenApiInfoWithDefaults

`func NewGridVOUpgradeLogOpenApiInfoWithDefaults() *GridVOUpgradeLogOpenApiInfo`

NewGridVOUpgradeLogOpenApiInfoWithDefaults instantiates a new GridVOUpgradeLogOpenApiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *GridVOUpgradeLogOpenApiInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *GridVOUpgradeLogOpenApiInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *GridVOUpgradeLogOpenApiInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *GridVOUpgradeLogOpenApiInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *GridVOUpgradeLogOpenApiInfo) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *GridVOUpgradeLogOpenApiInfo) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *GridVOUpgradeLogOpenApiInfo) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *GridVOUpgradeLogOpenApiInfo) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *GridVOUpgradeLogOpenApiInfo) GetData() []UpgradeLogOpenApiInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GridVOUpgradeLogOpenApiInfo) GetDataOk() (*[]UpgradeLogOpenApiInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GridVOUpgradeLogOpenApiInfo) SetData(v []UpgradeLogOpenApiInfo)`

SetData sets Data field to given value.

### HasData

`func (o *GridVOUpgradeLogOpenApiInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *GridVOUpgradeLogOpenApiInfo) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *GridVOUpgradeLogOpenApiInfo) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *GridVOUpgradeLogOpenApiInfo) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *GridVOUpgradeLogOpenApiInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


