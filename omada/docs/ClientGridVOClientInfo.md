# ClientGridVOClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientStat** | Pointer to [**ClientStatVO**](ClientStatVO.md) |  | [optional] 
**ClientTypeStat** | Pointer to [**ClientTypeStatVO**](ClientTypeStatVO.md) |  | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]ClientInfo**](ClientInfo.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewClientGridVOClientInfo

`func NewClientGridVOClientInfo() *ClientGridVOClientInfo`

NewClientGridVOClientInfo instantiates a new ClientGridVOClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGridVOClientInfoWithDefaults

`func NewClientGridVOClientInfoWithDefaults() *ClientGridVOClientInfo`

NewClientGridVOClientInfoWithDefaults instantiates a new ClientGridVOClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientStat

`func (o *ClientGridVOClientInfo) GetClientStat() ClientStatVO`

GetClientStat returns the ClientStat field if non-nil, zero value otherwise.

### GetClientStatOk

`func (o *ClientGridVOClientInfo) GetClientStatOk() (*ClientStatVO, bool)`

GetClientStatOk returns a tuple with the ClientStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientStat

`func (o *ClientGridVOClientInfo) SetClientStat(v ClientStatVO)`

SetClientStat sets ClientStat field to given value.

### HasClientStat

`func (o *ClientGridVOClientInfo) HasClientStat() bool`

HasClientStat returns a boolean if a field has been set.

### GetClientTypeStat

`func (o *ClientGridVOClientInfo) GetClientTypeStat() ClientTypeStatVO`

GetClientTypeStat returns the ClientTypeStat field if non-nil, zero value otherwise.

### GetClientTypeStatOk

`func (o *ClientGridVOClientInfo) GetClientTypeStatOk() (*ClientTypeStatVO, bool)`

GetClientTypeStatOk returns a tuple with the ClientTypeStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTypeStat

`func (o *ClientGridVOClientInfo) SetClientTypeStat(v ClientTypeStatVO)`

SetClientTypeStat sets ClientTypeStat field to given value.

### HasClientTypeStat

`func (o *ClientGridVOClientInfo) HasClientTypeStat() bool`

HasClientTypeStat returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ClientGridVOClientInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ClientGridVOClientInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ClientGridVOClientInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ClientGridVOClientInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *ClientGridVOClientInfo) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *ClientGridVOClientInfo) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *ClientGridVOClientInfo) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *ClientGridVOClientInfo) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *ClientGridVOClientInfo) GetData() []ClientInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientGridVOClientInfo) GetDataOk() (*[]ClientInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientGridVOClientInfo) SetData(v []ClientInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ClientGridVOClientInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *ClientGridVOClientInfo) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *ClientGridVOClientInfo) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *ClientGridVOClientInfo) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *ClientGridVOClientInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


