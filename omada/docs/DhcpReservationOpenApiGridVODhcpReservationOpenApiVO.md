# DhcpReservationOpenApiGridVODhcpReservationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]DhcpReservationOpenApiVO**](DhcpReservationOpenApiVO.md) |  | [optional] 
**NetNameToIdMap** | Pointer to **map[string]string** | Mapping between lan network name and lan network ID | [optional] 
**SupportOptions** | Pointer to **[]int32** | Options supported by DHCP Reservation. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewDhcpReservationOpenApiGridVODhcpReservationOpenApiVO

`func NewDhcpReservationOpenApiGridVODhcpReservationOpenApiVO() *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO`

NewDhcpReservationOpenApiGridVODhcpReservationOpenApiVO instantiates a new DhcpReservationOpenApiGridVODhcpReservationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpReservationOpenApiGridVODhcpReservationOpenApiVOWithDefaults

`func NewDhcpReservationOpenApiGridVODhcpReservationOpenApiVOWithDefaults() *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO`

NewDhcpReservationOpenApiGridVODhcpReservationOpenApiVOWithDefaults instantiates a new DhcpReservationOpenApiGridVODhcpReservationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetData() []DhcpReservationOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetDataOk() (*[]DhcpReservationOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) SetData(v []DhcpReservationOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNetNameToIdMap

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetNetNameToIdMap() map[string]string`

GetNetNameToIdMap returns the NetNameToIdMap field if non-nil, zero value otherwise.

### GetNetNameToIdMapOk

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetNetNameToIdMapOk() (*map[string]string, bool)`

GetNetNameToIdMapOk returns a tuple with the NetNameToIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetNameToIdMap

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) SetNetNameToIdMap(v map[string]string)`

SetNetNameToIdMap sets NetNameToIdMap field to given value.

### HasNetNameToIdMap

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) HasNetNameToIdMap() bool`

HasNetNameToIdMap returns a boolean if a field has been set.

### GetSupportOptions

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetSupportOptions() []int32`

GetSupportOptions returns the SupportOptions field if non-nil, zero value otherwise.

### GetSupportOptionsOk

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetSupportOptionsOk() (*[]int32, bool)`

GetSupportOptionsOk returns a tuple with the SupportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOptions

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) SetSupportOptions(v []int32)`

SetSupportOptions sets SupportOptions field to given value.

### HasSupportOptions

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) HasSupportOptions() bool`

HasSupportOptions returns a boolean if a field has been set.

### GetTotalRows

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *DhcpReservationOpenApiGridVODhcpReservationOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


