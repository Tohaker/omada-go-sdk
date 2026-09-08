# ApGroupGridVOApGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]ApGroupOpenApiVO**](ApGroupOpenApiVO.md) |  | [optional] 
**MaxSsids2G** | Pointer to **int32** | 2G radio max Ssid number in group | [optional] 
**MaxSsids5G** | Pointer to **int32** | 5G radio max Ssid number in group | [optional] 
**MaxSsids6G** | Pointer to **int32** | 6G radio max Ssid number in group | [optional] 
**MaxSsidsMlo** | Pointer to **int32** | max Mlo Ssid number in group | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewApGroupGridVOApGroupOpenApiVO

`func NewApGroupGridVOApGroupOpenApiVO() *ApGroupGridVOApGroupOpenApiVO`

NewApGroupGridVOApGroupOpenApiVO instantiates a new ApGroupGridVOApGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApGroupGridVOApGroupOpenApiVOWithDefaults

`func NewApGroupGridVOApGroupOpenApiVOWithDefaults() *ApGroupGridVOApGroupOpenApiVO`

NewApGroupGridVOApGroupOpenApiVOWithDefaults instantiates a new ApGroupGridVOApGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ApGroupGridVOApGroupOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ApGroupGridVOApGroupOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ApGroupGridVOApGroupOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ApGroupGridVOApGroupOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *ApGroupGridVOApGroupOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *ApGroupGridVOApGroupOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *ApGroupGridVOApGroupOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *ApGroupGridVOApGroupOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *ApGroupGridVOApGroupOpenApiVO) GetData() []ApGroupOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApGroupGridVOApGroupOpenApiVO) GetDataOk() (*[]ApGroupOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApGroupGridVOApGroupOpenApiVO) SetData(v []ApGroupOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *ApGroupGridVOApGroupOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMaxSsids2G

`func (o *ApGroupGridVOApGroupOpenApiVO) GetMaxSsids2G() int32`

GetMaxSsids2G returns the MaxSsids2G field if non-nil, zero value otherwise.

### GetMaxSsids2GOk

`func (o *ApGroupGridVOApGroupOpenApiVO) GetMaxSsids2GOk() (*int32, bool)`

GetMaxSsids2GOk returns a tuple with the MaxSsids2G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSsids2G

`func (o *ApGroupGridVOApGroupOpenApiVO) SetMaxSsids2G(v int32)`

SetMaxSsids2G sets MaxSsids2G field to given value.

### HasMaxSsids2G

`func (o *ApGroupGridVOApGroupOpenApiVO) HasMaxSsids2G() bool`

HasMaxSsids2G returns a boolean if a field has been set.

### GetMaxSsids5G

`func (o *ApGroupGridVOApGroupOpenApiVO) GetMaxSsids5G() int32`

GetMaxSsids5G returns the MaxSsids5G field if non-nil, zero value otherwise.

### GetMaxSsids5GOk

`func (o *ApGroupGridVOApGroupOpenApiVO) GetMaxSsids5GOk() (*int32, bool)`

GetMaxSsids5GOk returns a tuple with the MaxSsids5G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSsids5G

`func (o *ApGroupGridVOApGroupOpenApiVO) SetMaxSsids5G(v int32)`

SetMaxSsids5G sets MaxSsids5G field to given value.

### HasMaxSsids5G

`func (o *ApGroupGridVOApGroupOpenApiVO) HasMaxSsids5G() bool`

HasMaxSsids5G returns a boolean if a field has been set.

### GetMaxSsids6G

`func (o *ApGroupGridVOApGroupOpenApiVO) GetMaxSsids6G() int32`

GetMaxSsids6G returns the MaxSsids6G field if non-nil, zero value otherwise.

### GetMaxSsids6GOk

`func (o *ApGroupGridVOApGroupOpenApiVO) GetMaxSsids6GOk() (*int32, bool)`

GetMaxSsids6GOk returns a tuple with the MaxSsids6G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSsids6G

`func (o *ApGroupGridVOApGroupOpenApiVO) SetMaxSsids6G(v int32)`

SetMaxSsids6G sets MaxSsids6G field to given value.

### HasMaxSsids6G

`func (o *ApGroupGridVOApGroupOpenApiVO) HasMaxSsids6G() bool`

HasMaxSsids6G returns a boolean if a field has been set.

### GetMaxSsidsMlo

`func (o *ApGroupGridVOApGroupOpenApiVO) GetMaxSsidsMlo() int32`

GetMaxSsidsMlo returns the MaxSsidsMlo field if non-nil, zero value otherwise.

### GetMaxSsidsMloOk

`func (o *ApGroupGridVOApGroupOpenApiVO) GetMaxSsidsMloOk() (*int32, bool)`

GetMaxSsidsMloOk returns a tuple with the MaxSsidsMlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSsidsMlo

`func (o *ApGroupGridVOApGroupOpenApiVO) SetMaxSsidsMlo(v int32)`

SetMaxSsidsMlo sets MaxSsidsMlo field to given value.

### HasMaxSsidsMlo

`func (o *ApGroupGridVOApGroupOpenApiVO) HasMaxSsidsMlo() bool`

HasMaxSsidsMlo returns a boolean if a field has been set.

### GetTotalRows

`func (o *ApGroupGridVOApGroupOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *ApGroupGridVOApGroupOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *ApGroupGridVOApGroupOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *ApGroupGridVOApGroupOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


