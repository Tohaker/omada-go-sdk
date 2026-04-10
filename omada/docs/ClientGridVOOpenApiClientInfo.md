# ClientGridVOOpenApiClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientStat** | Pointer to [**ClientStatVO**](ClientStatVO.md) |  | [optional] 
**ClientTypeStat** | Pointer to [**ClientTypeStatVO**](ClientTypeStatVO.md) |  | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]OpenApiClientInfo**](OpenApiClientInfo.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewClientGridVOOpenApiClientInfo

`func NewClientGridVOOpenApiClientInfo() *ClientGridVOOpenApiClientInfo`

NewClientGridVOOpenApiClientInfo instantiates a new ClientGridVOOpenApiClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGridVOOpenApiClientInfoWithDefaults

`func NewClientGridVOOpenApiClientInfoWithDefaults() *ClientGridVOOpenApiClientInfo`

NewClientGridVOOpenApiClientInfoWithDefaults instantiates a new ClientGridVOOpenApiClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientStat

`func (o *ClientGridVOOpenApiClientInfo) GetClientStat() ClientStatVO`

GetClientStat returns the ClientStat field if non-nil, zero value otherwise.

### GetClientStatOk

`func (o *ClientGridVOOpenApiClientInfo) GetClientStatOk() (*ClientStatVO, bool)`

GetClientStatOk returns a tuple with the ClientStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientStat

`func (o *ClientGridVOOpenApiClientInfo) SetClientStat(v ClientStatVO)`

SetClientStat sets ClientStat field to given value.

### HasClientStat

`func (o *ClientGridVOOpenApiClientInfo) HasClientStat() bool`

HasClientStat returns a boolean if a field has been set.

### GetClientTypeStat

`func (o *ClientGridVOOpenApiClientInfo) GetClientTypeStat() ClientTypeStatVO`

GetClientTypeStat returns the ClientTypeStat field if non-nil, zero value otherwise.

### GetClientTypeStatOk

`func (o *ClientGridVOOpenApiClientInfo) GetClientTypeStatOk() (*ClientTypeStatVO, bool)`

GetClientTypeStatOk returns a tuple with the ClientTypeStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTypeStat

`func (o *ClientGridVOOpenApiClientInfo) SetClientTypeStat(v ClientTypeStatVO)`

SetClientTypeStat sets ClientTypeStat field to given value.

### HasClientTypeStat

`func (o *ClientGridVOOpenApiClientInfo) HasClientTypeStat() bool`

HasClientTypeStat returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ClientGridVOOpenApiClientInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ClientGridVOOpenApiClientInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ClientGridVOOpenApiClientInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ClientGridVOOpenApiClientInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *ClientGridVOOpenApiClientInfo) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *ClientGridVOOpenApiClientInfo) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *ClientGridVOOpenApiClientInfo) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *ClientGridVOOpenApiClientInfo) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *ClientGridVOOpenApiClientInfo) GetData() []OpenApiClientInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientGridVOOpenApiClientInfo) GetDataOk() (*[]OpenApiClientInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientGridVOOpenApiClientInfo) SetData(v []OpenApiClientInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ClientGridVOOpenApiClientInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *ClientGridVOOpenApiClientInfo) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *ClientGridVOOpenApiClientInfo) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *ClientGridVOOpenApiClientInfo) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *ClientGridVOOpenApiClientInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


