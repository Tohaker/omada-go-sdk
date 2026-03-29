# OpenApiQueryDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Page** | **int32** | Start from 1. | 
**PageSize** | **int32** | It should be within the range of 1–1000. | 
**SearchField** | Pointer to **string** |  | [optional] 
**SearchKey** | Pointer to **string** | Look for a specific piece of data. | [optional] 
**Sorts** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOpenApiQueryDataVO

`func NewOpenApiQueryDataVO(page int32, pageSize int32, ) *OpenApiQueryDataVO`

NewOpenApiQueryDataVO instantiates a new OpenApiQueryDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiQueryDataVOWithDefaults

`func NewOpenApiQueryDataVOWithDefaults() *OpenApiQueryDataVO`

NewOpenApiQueryDataVOWithDefaults instantiates a new OpenApiQueryDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *OpenApiQueryDataVO) GetFilters() map[string]map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OpenApiQueryDataVO) GetFiltersOk() (*map[string]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OpenApiQueryDataVO) SetFilters(v map[string]map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OpenApiQueryDataVO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPage

`func (o *OpenApiQueryDataVO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OpenApiQueryDataVO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OpenApiQueryDataVO) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *OpenApiQueryDataVO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OpenApiQueryDataVO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OpenApiQueryDataVO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetSearchField

`func (o *OpenApiQueryDataVO) GetSearchField() string`

GetSearchField returns the SearchField field if non-nil, zero value otherwise.

### GetSearchFieldOk

`func (o *OpenApiQueryDataVO) GetSearchFieldOk() (*string, bool)`

GetSearchFieldOk returns a tuple with the SearchField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchField

`func (o *OpenApiQueryDataVO) SetSearchField(v string)`

SetSearchField sets SearchField field to given value.

### HasSearchField

`func (o *OpenApiQueryDataVO) HasSearchField() bool`

HasSearchField returns a boolean if a field has been set.

### GetSearchKey

`func (o *OpenApiQueryDataVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *OpenApiQueryDataVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *OpenApiQueryDataVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *OpenApiQueryDataVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *OpenApiQueryDataVO) GetSorts() map[string]string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *OpenApiQueryDataVO) GetSortsOk() (*map[string]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *OpenApiQueryDataVO) SetSorts(v map[string]string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *OpenApiQueryDataVO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


