# OpenApiQueryDataV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Page** | **int32** | Start from 1. | 
**PageSize** | **int32** | It should be within the range of 1–100. | 
**SearchField** | Pointer to **string** |  | [optional] 
**SearchKey** | Pointer to **string** | Look for a specific piece of data. | [optional] 
**Sorts** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOpenApiQueryDataV2VO

`func NewOpenApiQueryDataV2VO(page int32, pageSize int32, ) *OpenApiQueryDataV2VO`

NewOpenApiQueryDataV2VO instantiates a new OpenApiQueryDataV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiQueryDataV2VOWithDefaults

`func NewOpenApiQueryDataV2VOWithDefaults() *OpenApiQueryDataV2VO`

NewOpenApiQueryDataV2VOWithDefaults instantiates a new OpenApiQueryDataV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *OpenApiQueryDataV2VO) GetFilters() map[string]map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OpenApiQueryDataV2VO) GetFiltersOk() (*map[string]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OpenApiQueryDataV2VO) SetFilters(v map[string]map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OpenApiQueryDataV2VO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPage

`func (o *OpenApiQueryDataV2VO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OpenApiQueryDataV2VO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OpenApiQueryDataV2VO) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *OpenApiQueryDataV2VO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OpenApiQueryDataV2VO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OpenApiQueryDataV2VO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetSearchField

`func (o *OpenApiQueryDataV2VO) GetSearchField() string`

GetSearchField returns the SearchField field if non-nil, zero value otherwise.

### GetSearchFieldOk

`func (o *OpenApiQueryDataV2VO) GetSearchFieldOk() (*string, bool)`

GetSearchFieldOk returns a tuple with the SearchField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchField

`func (o *OpenApiQueryDataV2VO) SetSearchField(v string)`

SetSearchField sets SearchField field to given value.

### HasSearchField

`func (o *OpenApiQueryDataV2VO) HasSearchField() bool`

HasSearchField returns a boolean if a field has been set.

### GetSearchKey

`func (o *OpenApiQueryDataV2VO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *OpenApiQueryDataV2VO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *OpenApiQueryDataV2VO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *OpenApiQueryDataV2VO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *OpenApiQueryDataV2VO) GetSorts() map[string]string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *OpenApiQueryDataV2VO) GetSortsOk() (*map[string]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *OpenApiQueryDataV2VO) SetSorts(v map[string]string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *OpenApiQueryDataV2VO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


