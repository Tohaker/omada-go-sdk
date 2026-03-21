# AlertLogGridVOAlertLogOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertLogStat** | Pointer to [**AlertLogStatOpenApiVO**](AlertLogStatOpenApiVO.md) |  | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]AlertLogOpenApiVO**](AlertLogOpenApiVO.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewAlertLogGridVOAlertLogOpenApiVO

`func NewAlertLogGridVOAlertLogOpenApiVO() *AlertLogGridVOAlertLogOpenApiVO`

NewAlertLogGridVOAlertLogOpenApiVO instantiates a new AlertLogGridVOAlertLogOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertLogGridVOAlertLogOpenApiVOWithDefaults

`func NewAlertLogGridVOAlertLogOpenApiVOWithDefaults() *AlertLogGridVOAlertLogOpenApiVO`

NewAlertLogGridVOAlertLogOpenApiVOWithDefaults instantiates a new AlertLogGridVOAlertLogOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertLogStat

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetAlertLogStat() AlertLogStatOpenApiVO`

GetAlertLogStat returns the AlertLogStat field if non-nil, zero value otherwise.

### GetAlertLogStatOk

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetAlertLogStatOk() (*AlertLogStatOpenApiVO, bool)`

GetAlertLogStatOk returns a tuple with the AlertLogStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLogStat

`func (o *AlertLogGridVOAlertLogOpenApiVO) SetAlertLogStat(v AlertLogStatOpenApiVO)`

SetAlertLogStat sets AlertLogStat field to given value.

### HasAlertLogStat

`func (o *AlertLogGridVOAlertLogOpenApiVO) HasAlertLogStat() bool`

HasAlertLogStat returns a boolean if a field has been set.

### GetCurrentPage

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *AlertLogGridVOAlertLogOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *AlertLogGridVOAlertLogOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *AlertLogGridVOAlertLogOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *AlertLogGridVOAlertLogOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetData() []AlertLogOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetDataOk() (*[]AlertLogOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertLogGridVOAlertLogOpenApiVO) SetData(v []AlertLogOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *AlertLogGridVOAlertLogOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *AlertLogGridVOAlertLogOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *AlertLogGridVOAlertLogOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *AlertLogGridVOAlertLogOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


