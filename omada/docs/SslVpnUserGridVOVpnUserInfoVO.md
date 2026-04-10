# SslVpnUserGridVOVpnUserInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VpnUserInfoVO**](VpnUserInfoVO.md) |  | [optional] 
**Expired** | Pointer to **int32** |  | [optional] 
**SslVpnMaxConUserNum** | Pointer to **int32** |  | [optional] 
**SubnetsLimitSize** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewSslVpnUserGridVOVpnUserInfoVO

`func NewSslVpnUserGridVOVpnUserInfoVO() *SslVpnUserGridVOVpnUserInfoVO`

NewSslVpnUserGridVOVpnUserInfoVO instantiates a new SslVpnUserGridVOVpnUserInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserGridVOVpnUserInfoVOWithDefaults

`func NewSslVpnUserGridVOVpnUserInfoVOWithDefaults() *SslVpnUserGridVOVpnUserInfoVO`

NewSslVpnUserGridVOVpnUserInfoVOWithDefaults instantiates a new SslVpnUserGridVOVpnUserInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCurrentPage

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetData() []VpnUserInfoVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetDataOk() (*[]VpnUserInfoVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetData(v []VpnUserInfoVO)`

SetData sets Data field to given value.

### HasData

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExpired

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetExpired() int32`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetExpiredOk() (*int32, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetExpired(v int32)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetSslVpnMaxConUserNum

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetSslVpnMaxConUserNum() int32`

GetSslVpnMaxConUserNum returns the SslVpnMaxConUserNum field if non-nil, zero value otherwise.

### GetSslVpnMaxConUserNumOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetSslVpnMaxConUserNumOk() (*int32, bool)`

GetSslVpnMaxConUserNumOk returns a tuple with the SslVpnMaxConUserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVpnMaxConUserNum

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetSslVpnMaxConUserNum(v int32)`

SetSslVpnMaxConUserNum sets SslVpnMaxConUserNum field to given value.

### HasSslVpnMaxConUserNum

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasSslVpnMaxConUserNum() bool`

HasSslVpnMaxConUserNum returns a boolean if a field has been set.

### GetSubnetsLimitSize

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetSubnetsLimitSize() int32`

GetSubnetsLimitSize returns the SubnetsLimitSize field if non-nil, zero value otherwise.

### GetSubnetsLimitSizeOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetSubnetsLimitSizeOk() (*int32, bool)`

GetSubnetsLimitSizeOk returns a tuple with the SubnetsLimitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsLimitSize

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetSubnetsLimitSize(v int32)`

SetSubnetsLimitSize sets SubnetsLimitSize field to given value.

### HasSubnetsLimitSize

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasSubnetsLimitSize() bool`

HasSubnetsLimitSize returns a boolean if a field has been set.

### GetTotal

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalRows

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *SslVpnUserGridVOVpnUserInfoVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *SslVpnUserGridVOVpnUserInfoVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *SslVpnUserGridVOVpnUserInfoVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


