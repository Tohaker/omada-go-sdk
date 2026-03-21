# ApplicationGridVOCategoryTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]CategoryTraffic**](CategoryTraffic.md) |  | [optional] 
**Stat** | Pointer to [**ApplicationStatVO**](ApplicationStatVO.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewApplicationGridVOCategoryTraffic

`func NewApplicationGridVOCategoryTraffic() *ApplicationGridVOCategoryTraffic`

NewApplicationGridVOCategoryTraffic instantiates a new ApplicationGridVOCategoryTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGridVOCategoryTrafficWithDefaults

`func NewApplicationGridVOCategoryTrafficWithDefaults() *ApplicationGridVOCategoryTraffic`

NewApplicationGridVOCategoryTrafficWithDefaults instantiates a new ApplicationGridVOCategoryTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ApplicationGridVOCategoryTraffic) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ApplicationGridVOCategoryTraffic) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ApplicationGridVOCategoryTraffic) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ApplicationGridVOCategoryTraffic) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *ApplicationGridVOCategoryTraffic) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *ApplicationGridVOCategoryTraffic) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *ApplicationGridVOCategoryTraffic) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *ApplicationGridVOCategoryTraffic) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *ApplicationGridVOCategoryTraffic) GetData() []CategoryTraffic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationGridVOCategoryTraffic) GetDataOk() (*[]CategoryTraffic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationGridVOCategoryTraffic) SetData(v []CategoryTraffic)`

SetData sets Data field to given value.

### HasData

`func (o *ApplicationGridVOCategoryTraffic) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStat

`func (o *ApplicationGridVOCategoryTraffic) GetStat() ApplicationStatVO`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *ApplicationGridVOCategoryTraffic) GetStatOk() (*ApplicationStatVO, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *ApplicationGridVOCategoryTraffic) SetStat(v ApplicationStatVO)`

SetStat sets Stat field to given value.

### HasStat

`func (o *ApplicationGridVOCategoryTraffic) HasStat() bool`

HasStat returns a boolean if a field has been set.

### GetTotalRows

`func (o *ApplicationGridVOCategoryTraffic) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *ApplicationGridVOCategoryTraffic) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *ApplicationGridVOCategoryTraffic) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *ApplicationGridVOCategoryTraffic) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


