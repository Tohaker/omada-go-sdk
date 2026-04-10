# LanDnsGridVOLanDnsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]LanDnsOpenApiVO**](LanDnsOpenApiVO.md) |  | [optional] 
**SupportTtl** | Pointer to **bool** | Whether the current gateway support custom TTL config. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewLanDnsGridVOLanDnsOpenApiVO

`func NewLanDnsGridVOLanDnsOpenApiVO() *LanDnsGridVOLanDnsOpenApiVO`

NewLanDnsGridVOLanDnsOpenApiVO instantiates a new LanDnsGridVOLanDnsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanDnsGridVOLanDnsOpenApiVOWithDefaults

`func NewLanDnsGridVOLanDnsOpenApiVOWithDefaults() *LanDnsGridVOLanDnsOpenApiVO`

NewLanDnsGridVOLanDnsOpenApiVOWithDefaults instantiates a new LanDnsGridVOLanDnsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *LanDnsGridVOLanDnsOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *LanDnsGridVOLanDnsOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *LanDnsGridVOLanDnsOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *LanDnsGridVOLanDnsOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetData() []LanDnsOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetDataOk() (*[]LanDnsOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LanDnsGridVOLanDnsOpenApiVO) SetData(v []LanDnsOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *LanDnsGridVOLanDnsOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportTtl

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetSupportTtl() bool`

GetSupportTtl returns the SupportTtl field if non-nil, zero value otherwise.

### GetSupportTtlOk

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetSupportTtlOk() (*bool, bool)`

GetSupportTtlOk returns a tuple with the SupportTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTtl

`func (o *LanDnsGridVOLanDnsOpenApiVO) SetSupportTtl(v bool)`

SetSupportTtl sets SupportTtl field to given value.

### HasSupportTtl

`func (o *LanDnsGridVOLanDnsOpenApiVO) HasSupportTtl() bool`

HasSupportTtl returns a boolean if a field has been set.

### GetTotalRows

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *LanDnsGridVOLanDnsOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *LanDnsGridVOLanDnsOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *LanDnsGridVOLanDnsOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


