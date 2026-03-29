# GridVOGetGridIpsThreat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]GetGridIpsThreat**](GetGridIpsThreat.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewGridVOGetGridIpsThreat

`func NewGridVOGetGridIpsThreat() *GridVOGetGridIpsThreat`

NewGridVOGetGridIpsThreat instantiates a new GridVOGetGridIpsThreat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridVOGetGridIpsThreatWithDefaults

`func NewGridVOGetGridIpsThreatWithDefaults() *GridVOGetGridIpsThreat`

NewGridVOGetGridIpsThreatWithDefaults instantiates a new GridVOGetGridIpsThreat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *GridVOGetGridIpsThreat) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *GridVOGetGridIpsThreat) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *GridVOGetGridIpsThreat) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *GridVOGetGridIpsThreat) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *GridVOGetGridIpsThreat) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *GridVOGetGridIpsThreat) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *GridVOGetGridIpsThreat) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *GridVOGetGridIpsThreat) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *GridVOGetGridIpsThreat) GetData() []GetGridIpsThreat`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GridVOGetGridIpsThreat) GetDataOk() (*[]GetGridIpsThreat, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GridVOGetGridIpsThreat) SetData(v []GetGridIpsThreat)`

SetData sets Data field to given value.

### HasData

`func (o *GridVOGetGridIpsThreat) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *GridVOGetGridIpsThreat) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *GridVOGetGridIpsThreat) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *GridVOGetGridIpsThreat) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *GridVOGetGridIpsThreat) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


