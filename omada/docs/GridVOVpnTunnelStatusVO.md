# GridVOVpnTunnelStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VpnTunnelStatusVO**](VpnTunnelStatusVO.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewGridVOVpnTunnelStatusVO

`func NewGridVOVpnTunnelStatusVO() *GridVOVpnTunnelStatusVO`

NewGridVOVpnTunnelStatusVO instantiates a new GridVOVpnTunnelStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridVOVpnTunnelStatusVOWithDefaults

`func NewGridVOVpnTunnelStatusVOWithDefaults() *GridVOVpnTunnelStatusVO`

NewGridVOVpnTunnelStatusVOWithDefaults instantiates a new GridVOVpnTunnelStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *GridVOVpnTunnelStatusVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *GridVOVpnTunnelStatusVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *GridVOVpnTunnelStatusVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *GridVOVpnTunnelStatusVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *GridVOVpnTunnelStatusVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *GridVOVpnTunnelStatusVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *GridVOVpnTunnelStatusVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *GridVOVpnTunnelStatusVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *GridVOVpnTunnelStatusVO) GetData() []VpnTunnelStatusVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GridVOVpnTunnelStatusVO) GetDataOk() (*[]VpnTunnelStatusVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GridVOVpnTunnelStatusVO) SetData(v []VpnTunnelStatusVO)`

SetData sets Data field to given value.

### HasData

`func (o *GridVOVpnTunnelStatusVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *GridVOVpnTunnelStatusVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *GridVOVpnTunnelStatusVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *GridVOVpnTunnelStatusVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *GridVOVpnTunnelStatusVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


