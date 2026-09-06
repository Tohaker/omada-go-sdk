# ClientGridVOSsidClientVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientStat** | Pointer to [**ClientStatVO**](ClientStatVO.md) |  | [optional] 
**ClientTypeStat** | Pointer to [**ClientTypeStatVO**](ClientTypeStatVO.md) |  | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]SsidClientVO**](SsidClientVO.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewClientGridVOSsidClientVO

`func NewClientGridVOSsidClientVO() *ClientGridVOSsidClientVO`

NewClientGridVOSsidClientVO instantiates a new ClientGridVOSsidClientVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGridVOSsidClientVOWithDefaults

`func NewClientGridVOSsidClientVOWithDefaults() *ClientGridVOSsidClientVO`

NewClientGridVOSsidClientVOWithDefaults instantiates a new ClientGridVOSsidClientVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientStat

`func (o *ClientGridVOSsidClientVO) GetClientStat() ClientStatVO`

GetClientStat returns the ClientStat field if non-nil, zero value otherwise.

### GetClientStatOk

`func (o *ClientGridVOSsidClientVO) GetClientStatOk() (*ClientStatVO, bool)`

GetClientStatOk returns a tuple with the ClientStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientStat

`func (o *ClientGridVOSsidClientVO) SetClientStat(v ClientStatVO)`

SetClientStat sets ClientStat field to given value.

### HasClientStat

`func (o *ClientGridVOSsidClientVO) HasClientStat() bool`

HasClientStat returns a boolean if a field has been set.

### GetClientTypeStat

`func (o *ClientGridVOSsidClientVO) GetClientTypeStat() ClientTypeStatVO`

GetClientTypeStat returns the ClientTypeStat field if non-nil, zero value otherwise.

### GetClientTypeStatOk

`func (o *ClientGridVOSsidClientVO) GetClientTypeStatOk() (*ClientTypeStatVO, bool)`

GetClientTypeStatOk returns a tuple with the ClientTypeStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTypeStat

`func (o *ClientGridVOSsidClientVO) SetClientTypeStat(v ClientTypeStatVO)`

SetClientTypeStat sets ClientTypeStat field to given value.

### HasClientTypeStat

`func (o *ClientGridVOSsidClientVO) HasClientTypeStat() bool`

HasClientTypeStat returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ClientGridVOSsidClientVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ClientGridVOSsidClientVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ClientGridVOSsidClientVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ClientGridVOSsidClientVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *ClientGridVOSsidClientVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *ClientGridVOSsidClientVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *ClientGridVOSsidClientVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *ClientGridVOSsidClientVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *ClientGridVOSsidClientVO) GetData() []SsidClientVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientGridVOSsidClientVO) GetDataOk() (*[]SsidClientVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientGridVOSsidClientVO) SetData(v []SsidClientVO)`

SetData sets Data field to given value.

### HasData

`func (o *ClientGridVOSsidClientVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *ClientGridVOSsidClientVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *ClientGridVOSsidClientVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *ClientGridVOSsidClientVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *ClientGridVOSsidClientVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


