# UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]QueryUrlFilteringOpenApiVO**](QueryUrlFilteringOpenApiVO.md) |  | [optional] 
**SupportCategory** | Pointer to **bool** | Whether category is supported of the URL filtering. | [optional] 
**SupportKeyword** | Pointer to **bool** | Whether keyword is supported of the URL filtering. | [optional] 
**SupportTimeSchedule** | Pointer to **bool** | Whether time schedule is supported of the URL filtering. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO

`func NewUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO() *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO`

NewUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO instantiates a new UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVOWithDefaults

`func NewUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVOWithDefaults() *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO`

NewUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVOWithDefaults instantiates a new UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetData() []QueryUrlFilteringOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetDataOk() (*[]QueryUrlFilteringOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) SetData(v []QueryUrlFilteringOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportCategory

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetSupportCategory() bool`

GetSupportCategory returns the SupportCategory field if non-nil, zero value otherwise.

### GetSupportCategoryOk

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetSupportCategoryOk() (*bool, bool)`

GetSupportCategoryOk returns a tuple with the SupportCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCategory

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) SetSupportCategory(v bool)`

SetSupportCategory sets SupportCategory field to given value.

### HasSupportCategory

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) HasSupportCategory() bool`

HasSupportCategory returns a boolean if a field has been set.

### GetSupportKeyword

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetSupportKeyword() bool`

GetSupportKeyword returns the SupportKeyword field if non-nil, zero value otherwise.

### GetSupportKeywordOk

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetSupportKeywordOk() (*bool, bool)`

GetSupportKeywordOk returns a tuple with the SupportKeyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportKeyword

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) SetSupportKeyword(v bool)`

SetSupportKeyword sets SupportKeyword field to given value.

### HasSupportKeyword

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) HasSupportKeyword() bool`

HasSupportKeyword returns a boolean if a field has been set.

### GetSupportTimeSchedule

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetSupportTimeSchedule() bool`

GetSupportTimeSchedule returns the SupportTimeSchedule field if non-nil, zero value otherwise.

### GetSupportTimeScheduleOk

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetSupportTimeScheduleOk() (*bool, bool)`

GetSupportTimeScheduleOk returns a tuple with the SupportTimeSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTimeSchedule

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) SetSupportTimeSchedule(v bool)`

SetSupportTimeSchedule sets SupportTimeSchedule field to given value.

### HasSupportTimeSchedule

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) HasSupportTimeSchedule() bool`

HasSupportTimeSchedule returns a boolean if a field has been set.

### GetTotalRows

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *UrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


