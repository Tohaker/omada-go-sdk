# GridVOSSLVPNLockEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]SSLVPNLockEntity**](SSLVPNLockEntity.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewGridVOSSLVPNLockEntity

`func NewGridVOSSLVPNLockEntity() *GridVOSSLVPNLockEntity`

NewGridVOSSLVPNLockEntity instantiates a new GridVOSSLVPNLockEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridVOSSLVPNLockEntityWithDefaults

`func NewGridVOSSLVPNLockEntityWithDefaults() *GridVOSSLVPNLockEntity`

NewGridVOSSLVPNLockEntityWithDefaults instantiates a new GridVOSSLVPNLockEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *GridVOSSLVPNLockEntity) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *GridVOSSLVPNLockEntity) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *GridVOSSLVPNLockEntity) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *GridVOSSLVPNLockEntity) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *GridVOSSLVPNLockEntity) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *GridVOSSLVPNLockEntity) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *GridVOSSLVPNLockEntity) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *GridVOSSLVPNLockEntity) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *GridVOSSLVPNLockEntity) GetData() []SSLVPNLockEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GridVOSSLVPNLockEntity) GetDataOk() (*[]SSLVPNLockEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GridVOSSLVPNLockEntity) SetData(v []SSLVPNLockEntity)`

SetData sets Data field to given value.

### HasData

`func (o *GridVOSSLVPNLockEntity) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *GridVOSSLVPNLockEntity) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *GridVOSSLVPNLockEntity) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *GridVOSSLVPNLockEntity) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *GridVOSSLVPNLockEntity) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


