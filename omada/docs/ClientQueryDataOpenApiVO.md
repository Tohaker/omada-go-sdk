# ClientQueryDataOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**ClientQueryFiltersOpenApiVO**](ClientQueryFiltersOpenApiVO.md) |  | [optional] 
**Page** | **int32** | Start from 1. | 
**PageSize** | **int32** | It should be within the range of 1–1000. | 
**Scope** | Pointer to **int32** | Scope of clients to query, 0: all, 1: online(default), 2:offline, 3:blocked. | [optional] 
**SearchKey** | Pointer to **string** | Fuzzy query parameters, support field name, mac, ip. | [optional] 
**Sorts** | Pointer to **map[string]string** | Sort rule, key: sort field, value: sort direction, value parameter may be one of asc or desc.Optional parameter. | [optional] 

## Methods

### NewClientQueryDataOpenApiVO

`func NewClientQueryDataOpenApiVO(page int32, pageSize int32, ) *ClientQueryDataOpenApiVO`

NewClientQueryDataOpenApiVO instantiates a new ClientQueryDataOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientQueryDataOpenApiVOWithDefaults

`func NewClientQueryDataOpenApiVOWithDefaults() *ClientQueryDataOpenApiVO`

NewClientQueryDataOpenApiVOWithDefaults instantiates a new ClientQueryDataOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ClientQueryDataOpenApiVO) GetFilters() ClientQueryFiltersOpenApiVO`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ClientQueryDataOpenApiVO) GetFiltersOk() (*ClientQueryFiltersOpenApiVO, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ClientQueryDataOpenApiVO) SetFilters(v ClientQueryFiltersOpenApiVO)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ClientQueryDataOpenApiVO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPage

`func (o *ClientQueryDataOpenApiVO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ClientQueryDataOpenApiVO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ClientQueryDataOpenApiVO) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *ClientQueryDataOpenApiVO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ClientQueryDataOpenApiVO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ClientQueryDataOpenApiVO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetScope

`func (o *ClientQueryDataOpenApiVO) GetScope() int32`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ClientQueryDataOpenApiVO) GetScopeOk() (*int32, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ClientQueryDataOpenApiVO) SetScope(v int32)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ClientQueryDataOpenApiVO) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSearchKey

`func (o *ClientQueryDataOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *ClientQueryDataOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *ClientQueryDataOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *ClientQueryDataOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *ClientQueryDataOpenApiVO) GetSorts() map[string]string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ClientQueryDataOpenApiVO) GetSortsOk() (*map[string]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ClientQueryDataOpenApiVO) SetSorts(v map[string]string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ClientQueryDataOpenApiVO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


