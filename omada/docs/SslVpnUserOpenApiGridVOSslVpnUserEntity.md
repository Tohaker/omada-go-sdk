# SslVpnUserOpenApiGridVOSslVpnUserEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **int32** | Available number of the SSL VPN user | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]SslVpnUserEntity**](SslVpnUserEntity.md) |  | [optional] 
**Expired** | Pointer to **int32** | Expired number of the SSL VPN user | [optional] 
**Total** | Pointer to **int32** | Total number of the SSL VPN user | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewSslVpnUserOpenApiGridVOSslVpnUserEntity

`func NewSslVpnUserOpenApiGridVOSslVpnUserEntity() *SslVpnUserOpenApiGridVOSslVpnUserEntity`

NewSslVpnUserOpenApiGridVOSslVpnUserEntity instantiates a new SslVpnUserOpenApiGridVOSslVpnUserEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserOpenApiGridVOSslVpnUserEntityWithDefaults

`func NewSslVpnUserOpenApiGridVOSslVpnUserEntityWithDefaults() *SslVpnUserOpenApiGridVOSslVpnUserEntity`

NewSslVpnUserOpenApiGridVOSslVpnUserEntityWithDefaults instantiates a new SslVpnUserOpenApiGridVOSslVpnUserEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCurrentPage

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetData() []SslVpnUserEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetDataOk() (*[]SslVpnUserEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) SetData(v []SslVpnUserEntity)`

SetData sets Data field to given value.

### HasData

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExpired

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetExpired() int32`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetExpiredOk() (*int32, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) SetExpired(v int32)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetTotal

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalRows

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *SslVpnUserOpenApiGridVOSslVpnUserEntity) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


