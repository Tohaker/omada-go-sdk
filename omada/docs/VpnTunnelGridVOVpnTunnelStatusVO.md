# VpnTunnelGridVOVpnTunnelStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VpnTunnelStatusVO**](VpnTunnelStatusVO.md) |  | [optional] 
**LockStatus** | Pointer to **bool** | Lock status of the SSL VPN tunnel | [optional] 
**LockType** | Pointer to **int32** | Lock type of the SSL VPN tunnel | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewVpnTunnelGridVOVpnTunnelStatusVO

`func NewVpnTunnelGridVOVpnTunnelStatusVO() *VpnTunnelGridVOVpnTunnelStatusVO`

NewVpnTunnelGridVOVpnTunnelStatusVO instantiates a new VpnTunnelGridVOVpnTunnelStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelGridVOVpnTunnelStatusVOWithDefaults

`func NewVpnTunnelGridVOVpnTunnelStatusVOWithDefaults() *VpnTunnelGridVOVpnTunnelStatusVO`

NewVpnTunnelGridVOVpnTunnelStatusVOWithDefaults instantiates a new VpnTunnelGridVOVpnTunnelStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetData() []VpnTunnelStatusVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetDataOk() (*[]VpnTunnelStatusVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) SetData(v []VpnTunnelStatusVO)`

SetData sets Data field to given value.

### HasData

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLockStatus

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetLockStatus() bool`

GetLockStatus returns the LockStatus field if non-nil, zero value otherwise.

### GetLockStatusOk

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetLockStatusOk() (*bool, bool)`

GetLockStatusOk returns a tuple with the LockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockStatus

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) SetLockStatus(v bool)`

SetLockStatus sets LockStatus field to given value.

### HasLockStatus

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) HasLockStatus() bool`

HasLockStatus returns a boolean if a field has been set.

### GetLockType

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetLockType() int32`

GetLockType returns the LockType field if non-nil, zero value otherwise.

### GetLockTypeOk

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetLockTypeOk() (*int32, bool)`

GetLockTypeOk returns a tuple with the LockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockType

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) SetLockType(v int32)`

SetLockType sets LockType field to given value.

### HasLockType

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) HasLockType() bool`

HasLockType returns a boolean if a field has been set.

### GetTotalRows

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *VpnTunnelGridVOVpnTunnelStatusVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


