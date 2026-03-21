# AnomalySettingGridVOAnomalyEventSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]AnomalyEventSettingOpenApiVO**](AnomalyEventSettingOpenApiVO.md) |  | [optional] 
**Resource** | Pointer to **int32** | The anomaly event setting creation resource, such as: 0: new created, 1: from template, 2: override | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewAnomalySettingGridVOAnomalyEventSettingOpenApiVO

`func NewAnomalySettingGridVOAnomalyEventSettingOpenApiVO() *AnomalySettingGridVOAnomalyEventSettingOpenApiVO`

NewAnomalySettingGridVOAnomalyEventSettingOpenApiVO instantiates a new AnomalySettingGridVOAnomalyEventSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalySettingGridVOAnomalyEventSettingOpenApiVOWithDefaults

`func NewAnomalySettingGridVOAnomalyEventSettingOpenApiVOWithDefaults() *AnomalySettingGridVOAnomalyEventSettingOpenApiVO`

NewAnomalySettingGridVOAnomalyEventSettingOpenApiVOWithDefaults instantiates a new AnomalySettingGridVOAnomalyEventSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetData() []AnomalyEventSettingOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetDataOk() (*[]AnomalyEventSettingOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) SetData(v []AnomalyEventSettingOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetResource

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetTotalRows

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *AnomalySettingGridVOAnomalyEventSettingOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


