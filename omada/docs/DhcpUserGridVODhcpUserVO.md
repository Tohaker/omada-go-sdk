# DhcpUserGridVODhcpUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]DhcpUserVO**](DhcpUserVO.md) |  | [optional] 
**NetNameToIdMap** | Pointer to **map[string]string** | Mapping between lan network name and lan network ID | [optional] 
**SelectedNum** | Pointer to **int64** | The number of DHCP users that cannot be selected when adding from the user list in DHCP reservation. | [optional] 
**ServerNameToMacMap** | Pointer to **map[string][]string** | Mapping between server name and macs | [optional] 
**ServerNameToStackIdMap** | Pointer to **map[string][]string** | Mapping between server name and stack IDs | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewDhcpUserGridVODhcpUserVO

`func NewDhcpUserGridVODhcpUserVO() *DhcpUserGridVODhcpUserVO`

NewDhcpUserGridVODhcpUserVO instantiates a new DhcpUserGridVODhcpUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpUserGridVODhcpUserVOWithDefaults

`func NewDhcpUserGridVODhcpUserVOWithDefaults() *DhcpUserGridVODhcpUserVO`

NewDhcpUserGridVODhcpUserVOWithDefaults instantiates a new DhcpUserGridVODhcpUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *DhcpUserGridVODhcpUserVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *DhcpUserGridVODhcpUserVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *DhcpUserGridVODhcpUserVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *DhcpUserGridVODhcpUserVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *DhcpUserGridVODhcpUserVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *DhcpUserGridVODhcpUserVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *DhcpUserGridVODhcpUserVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *DhcpUserGridVODhcpUserVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *DhcpUserGridVODhcpUserVO) GetData() []DhcpUserVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpUserGridVODhcpUserVO) GetDataOk() (*[]DhcpUserVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpUserGridVODhcpUserVO) SetData(v []DhcpUserVO)`

SetData sets Data field to given value.

### HasData

`func (o *DhcpUserGridVODhcpUserVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNetNameToIdMap

`func (o *DhcpUserGridVODhcpUserVO) GetNetNameToIdMap() map[string]string`

GetNetNameToIdMap returns the NetNameToIdMap field if non-nil, zero value otherwise.

### GetNetNameToIdMapOk

`func (o *DhcpUserGridVODhcpUserVO) GetNetNameToIdMapOk() (*map[string]string, bool)`

GetNetNameToIdMapOk returns a tuple with the NetNameToIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetNameToIdMap

`func (o *DhcpUserGridVODhcpUserVO) SetNetNameToIdMap(v map[string]string)`

SetNetNameToIdMap sets NetNameToIdMap field to given value.

### HasNetNameToIdMap

`func (o *DhcpUserGridVODhcpUserVO) HasNetNameToIdMap() bool`

HasNetNameToIdMap returns a boolean if a field has been set.

### GetSelectedNum

`func (o *DhcpUserGridVODhcpUserVO) GetSelectedNum() int64`

GetSelectedNum returns the SelectedNum field if non-nil, zero value otherwise.

### GetSelectedNumOk

`func (o *DhcpUserGridVODhcpUserVO) GetSelectedNumOk() (*int64, bool)`

GetSelectedNumOk returns a tuple with the SelectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedNum

`func (o *DhcpUserGridVODhcpUserVO) SetSelectedNum(v int64)`

SetSelectedNum sets SelectedNum field to given value.

### HasSelectedNum

`func (o *DhcpUserGridVODhcpUserVO) HasSelectedNum() bool`

HasSelectedNum returns a boolean if a field has been set.

### GetServerNameToMacMap

`func (o *DhcpUserGridVODhcpUserVO) GetServerNameToMacMap() map[string][]string`

GetServerNameToMacMap returns the ServerNameToMacMap field if non-nil, zero value otherwise.

### GetServerNameToMacMapOk

`func (o *DhcpUserGridVODhcpUserVO) GetServerNameToMacMapOk() (*map[string][]string, bool)`

GetServerNameToMacMapOk returns a tuple with the ServerNameToMacMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNameToMacMap

`func (o *DhcpUserGridVODhcpUserVO) SetServerNameToMacMap(v map[string][]string)`

SetServerNameToMacMap sets ServerNameToMacMap field to given value.

### HasServerNameToMacMap

`func (o *DhcpUserGridVODhcpUserVO) HasServerNameToMacMap() bool`

HasServerNameToMacMap returns a boolean if a field has been set.

### GetServerNameToStackIdMap

`func (o *DhcpUserGridVODhcpUserVO) GetServerNameToStackIdMap() map[string][]string`

GetServerNameToStackIdMap returns the ServerNameToStackIdMap field if non-nil, zero value otherwise.

### GetServerNameToStackIdMapOk

`func (o *DhcpUserGridVODhcpUserVO) GetServerNameToStackIdMapOk() (*map[string][]string, bool)`

GetServerNameToStackIdMapOk returns a tuple with the ServerNameToStackIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNameToStackIdMap

`func (o *DhcpUserGridVODhcpUserVO) SetServerNameToStackIdMap(v map[string][]string)`

SetServerNameToStackIdMap sets ServerNameToStackIdMap field to given value.

### HasServerNameToStackIdMap

`func (o *DhcpUserGridVODhcpUserVO) HasServerNameToStackIdMap() bool`

HasServerNameToStackIdMap returns a boolean if a field has been set.

### GetTotalRows

`func (o *DhcpUserGridVODhcpUserVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *DhcpUserGridVODhcpUserVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *DhcpUserGridVODhcpUserVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *DhcpUserGridVODhcpUserVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


