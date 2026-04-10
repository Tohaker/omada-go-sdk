# SslVpnUserGroupGridVOSslVpnUserGroupBriefVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]SslVpnUserGroupBriefVO**](SslVpnUserGroupBriefVO.md) |  | [optional] 
**SupportLdap** | Pointer to **bool** |  | [optional] 
**SupportRadius** | Pointer to **bool** |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewSslVpnUserGroupGridVOSslVpnUserGroupBriefVO

`func NewSslVpnUserGroupGridVOSslVpnUserGroupBriefVO() *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO`

NewSslVpnUserGroupGridVOSslVpnUserGroupBriefVO instantiates a new SslVpnUserGroupGridVOSslVpnUserGroupBriefVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserGroupGridVOSslVpnUserGroupBriefVOWithDefaults

`func NewSslVpnUserGroupGridVOSslVpnUserGroupBriefVOWithDefaults() *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO`

NewSslVpnUserGroupGridVOSslVpnUserGroupBriefVOWithDefaults instantiates a new SslVpnUserGroupGridVOSslVpnUserGroupBriefVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetData() []SslVpnUserGroupBriefVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetDataOk() (*[]SslVpnUserGroupBriefVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) SetData(v []SslVpnUserGroupBriefVO)`

SetData sets Data field to given value.

### HasData

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportLdap

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetSupportLdap() bool`

GetSupportLdap returns the SupportLdap field if non-nil, zero value otherwise.

### GetSupportLdapOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetSupportLdapOk() (*bool, bool)`

GetSupportLdapOk returns a tuple with the SupportLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLdap

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) SetSupportLdap(v bool)`

SetSupportLdap sets SupportLdap field to given value.

### HasSupportLdap

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) HasSupportLdap() bool`

HasSupportLdap returns a boolean if a field has been set.

### GetSupportRadius

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetSupportRadius() bool`

GetSupportRadius returns the SupportRadius field if non-nil, zero value otherwise.

### GetSupportRadiusOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetSupportRadiusOk() (*bool, bool)`

GetSupportRadiusOk returns a tuple with the SupportRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRadius

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) SetSupportRadius(v bool)`

SetSupportRadius sets SupportRadius field to given value.

### HasSupportRadius

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) HasSupportRadius() bool`

HasSupportRadius returns a boolean if a field has been set.

### GetTotalRows

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *SslVpnUserGroupGridVOSslVpnUserGroupBriefVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


