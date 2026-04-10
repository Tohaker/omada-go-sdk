# SslVpnUserGroupGridVOSslVpnUserGroupEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]SslVpnUserGroupEntity**](SslVpnUserGroupEntity.md) |  | [optional] 
**SupportLdap** | Pointer to **bool** |  | [optional] 
**SupportRadius** | Pointer to **bool** |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewSslVpnUserGroupGridVOSslVpnUserGroupEntity

`func NewSslVpnUserGroupGridVOSslVpnUserGroupEntity() *SslVpnUserGroupGridVOSslVpnUserGroupEntity`

NewSslVpnUserGroupGridVOSslVpnUserGroupEntity instantiates a new SslVpnUserGroupGridVOSslVpnUserGroupEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserGroupGridVOSslVpnUserGroupEntityWithDefaults

`func NewSslVpnUserGroupGridVOSslVpnUserGroupEntityWithDefaults() *SslVpnUserGroupGridVOSslVpnUserGroupEntity`

NewSslVpnUserGroupGridVOSslVpnUserGroupEntityWithDefaults instantiates a new SslVpnUserGroupGridVOSslVpnUserGroupEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetData() []SslVpnUserGroupEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetDataOk() (*[]SslVpnUserGroupEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) SetData(v []SslVpnUserGroupEntity)`

SetData sets Data field to given value.

### HasData

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportLdap

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetSupportLdap() bool`

GetSupportLdap returns the SupportLdap field if non-nil, zero value otherwise.

### GetSupportLdapOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetSupportLdapOk() (*bool, bool)`

GetSupportLdapOk returns a tuple with the SupportLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLdap

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) SetSupportLdap(v bool)`

SetSupportLdap sets SupportLdap field to given value.

### HasSupportLdap

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) HasSupportLdap() bool`

HasSupportLdap returns a boolean if a field has been set.

### GetSupportRadius

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetSupportRadius() bool`

GetSupportRadius returns the SupportRadius field if non-nil, zero value otherwise.

### GetSupportRadiusOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetSupportRadiusOk() (*bool, bool)`

GetSupportRadiusOk returns a tuple with the SupportRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRadius

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) SetSupportRadius(v bool)`

SetSupportRadius sets SupportRadius field to given value.

### HasSupportRadius

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) HasSupportRadius() bool`

HasSupportRadius returns a boolean if a field has been set.

### GetTotalRows

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupEntity) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


