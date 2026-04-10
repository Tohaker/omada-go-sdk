# ClientTrafficOpenApiQueryDataV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientTrafficFilters** | Pointer to [**[]ClientTrafficFilterOpenApiVO**](ClientTrafficFilterOpenApiVO.md) | Client traffic filtering condition | [optional] 
**End** | Pointer to **int64** | End timestamp, in seconds, such as 1682000000 | [optional] 
**Filters** | Pointer to **map[string]map[string]interface{}** | Filter [family] in the form of Map. When the value of the [family] key is empty, return an empty result. | [optional] 
**Page** | **int32** | Start from 1. | 
**PageSize** | **int32** | It should be within the range of 1–100. | 
**Sorts** | Pointer to **map[string]string** | Sort rule, key: sort field, value: sort direction, value parameter may be one of asc or desc. | [optional] 
**Start** | Pointer to **int64** | Start timestamp, in seconds, such as 1682000000 | [optional] 

## Methods

### NewClientTrafficOpenApiQueryDataV2VO

`func NewClientTrafficOpenApiQueryDataV2VO(page int32, pageSize int32, ) *ClientTrafficOpenApiQueryDataV2VO`

NewClientTrafficOpenApiQueryDataV2VO instantiates a new ClientTrafficOpenApiQueryDataV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTrafficOpenApiQueryDataV2VOWithDefaults

`func NewClientTrafficOpenApiQueryDataV2VOWithDefaults() *ClientTrafficOpenApiQueryDataV2VO`

NewClientTrafficOpenApiQueryDataV2VOWithDefaults instantiates a new ClientTrafficOpenApiQueryDataV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTrafficFilters

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetClientTrafficFilters() []ClientTrafficFilterOpenApiVO`

GetClientTrafficFilters returns the ClientTrafficFilters field if non-nil, zero value otherwise.

### GetClientTrafficFiltersOk

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetClientTrafficFiltersOk() (*[]ClientTrafficFilterOpenApiVO, bool)`

GetClientTrafficFiltersOk returns a tuple with the ClientTrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTrafficFilters

`func (o *ClientTrafficOpenApiQueryDataV2VO) SetClientTrafficFilters(v []ClientTrafficFilterOpenApiVO)`

SetClientTrafficFilters sets ClientTrafficFilters field to given value.

### HasClientTrafficFilters

`func (o *ClientTrafficOpenApiQueryDataV2VO) HasClientTrafficFilters() bool`

HasClientTrafficFilters returns a boolean if a field has been set.

### GetEnd

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ClientTrafficOpenApiQueryDataV2VO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ClientTrafficOpenApiQueryDataV2VO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFilters

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetFilters() map[string]map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetFiltersOk() (*map[string]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ClientTrafficOpenApiQueryDataV2VO) SetFilters(v map[string]map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ClientTrafficOpenApiQueryDataV2VO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPage

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ClientTrafficOpenApiQueryDataV2VO) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ClientTrafficOpenApiQueryDataV2VO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetSorts

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetSorts() map[string]string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetSortsOk() (*map[string]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ClientTrafficOpenApiQueryDataV2VO) SetSorts(v map[string]string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ClientTrafficOpenApiQueryDataV2VO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.

### GetStart

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientTrafficOpenApiQueryDataV2VO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientTrafficOpenApiQueryDataV2VO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ClientTrafficOpenApiQueryDataV2VO) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


