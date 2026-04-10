# OsgDhcpUserGridVOOsgDhcpUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]OsgDhcpUserVO**](OsgDhcpUserVO.md) |  | [optional] 
**NetNameToIdMap** | Pointer to **map[string]string** | Mapping between lan network name and lan network ID | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewOsgDhcpUserGridVOOsgDhcpUserVO

`func NewOsgDhcpUserGridVOOsgDhcpUserVO() *OsgDhcpUserGridVOOsgDhcpUserVO`

NewOsgDhcpUserGridVOOsgDhcpUserVO instantiates a new OsgDhcpUserGridVOOsgDhcpUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgDhcpUserGridVOOsgDhcpUserVOWithDefaults

`func NewOsgDhcpUserGridVOOsgDhcpUserVOWithDefaults() *OsgDhcpUserGridVOOsgDhcpUserVO`

NewOsgDhcpUserGridVOOsgDhcpUserVOWithDefaults instantiates a new OsgDhcpUserGridVOOsgDhcpUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetData() []OsgDhcpUserVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetDataOk() (*[]OsgDhcpUserVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) SetData(v []OsgDhcpUserVO)`

SetData sets Data field to given value.

### HasData

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNetNameToIdMap

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetNetNameToIdMap() map[string]string`

GetNetNameToIdMap returns the NetNameToIdMap field if non-nil, zero value otherwise.

### GetNetNameToIdMapOk

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetNetNameToIdMapOk() (*map[string]string, bool)`

GetNetNameToIdMapOk returns a tuple with the NetNameToIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetNameToIdMap

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) SetNetNameToIdMap(v map[string]string)`

SetNetNameToIdMap sets NetNameToIdMap field to given value.

### HasNetNameToIdMap

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) HasNetNameToIdMap() bool`

HasNetNameToIdMap returns a boolean if a field has been set.

### GetTotalRows

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *OsgDhcpUserGridVOOsgDhcpUserVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


