# QueryOswDateByNetworkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsyncColumns** | Pointer to **string** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 
**CurrentPageSize** | Pointer to **int32** |  | [optional] 
**Filters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**LanNetworkIds** | Pointer to **[]string** |  | [optional] 
**MultiSearchMap** | Pointer to **map[string]string** |  | [optional] 
**SearchField** | Pointer to **string** |  | [optional] 
**SearchKey** | Pointer to **string** |  | [optional] 
**Sorts** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewQueryOswDateByNetworkVO

`func NewQueryOswDateByNetworkVO() *QueryOswDateByNetworkVO`

NewQueryOswDateByNetworkVO instantiates a new QueryOswDateByNetworkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOswDateByNetworkVOWithDefaults

`func NewQueryOswDateByNetworkVOWithDefaults() *QueryOswDateByNetworkVO`

NewQueryOswDateByNetworkVOWithDefaults instantiates a new QueryOswDateByNetworkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsyncColumns

`func (o *QueryOswDateByNetworkVO) GetAsyncColumns() string`

GetAsyncColumns returns the AsyncColumns field if non-nil, zero value otherwise.

### GetAsyncColumnsOk

`func (o *QueryOswDateByNetworkVO) GetAsyncColumnsOk() (*string, bool)`

GetAsyncColumnsOk returns a tuple with the AsyncColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncColumns

`func (o *QueryOswDateByNetworkVO) SetAsyncColumns(v string)`

SetAsyncColumns sets AsyncColumns field to given value.

### HasAsyncColumns

`func (o *QueryOswDateByNetworkVO) HasAsyncColumns() bool`

HasAsyncColumns returns a boolean if a field has been set.

### GetCurrentPage

`func (o *QueryOswDateByNetworkVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *QueryOswDateByNetworkVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *QueryOswDateByNetworkVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *QueryOswDateByNetworkVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentPageSize

`func (o *QueryOswDateByNetworkVO) GetCurrentPageSize() int32`

GetCurrentPageSize returns the CurrentPageSize field if non-nil, zero value otherwise.

### GetCurrentPageSizeOk

`func (o *QueryOswDateByNetworkVO) GetCurrentPageSizeOk() (*int32, bool)`

GetCurrentPageSizeOk returns a tuple with the CurrentPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageSize

`func (o *QueryOswDateByNetworkVO) SetCurrentPageSize(v int32)`

SetCurrentPageSize sets CurrentPageSize field to given value.

### HasCurrentPageSize

`func (o *QueryOswDateByNetworkVO) HasCurrentPageSize() bool`

HasCurrentPageSize returns a boolean if a field has been set.

### GetFilters

`func (o *QueryOswDateByNetworkVO) GetFilters() map[string]map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryOswDateByNetworkVO) GetFiltersOk() (*map[string]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryOswDateByNetworkVO) SetFilters(v map[string]map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryOswDateByNetworkVO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLanNetworkIds

`func (o *QueryOswDateByNetworkVO) GetLanNetworkIds() []string`

GetLanNetworkIds returns the LanNetworkIds field if non-nil, zero value otherwise.

### GetLanNetworkIdsOk

`func (o *QueryOswDateByNetworkVO) GetLanNetworkIdsOk() (*[]string, bool)`

GetLanNetworkIdsOk returns a tuple with the LanNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkIds

`func (o *QueryOswDateByNetworkVO) SetLanNetworkIds(v []string)`

SetLanNetworkIds sets LanNetworkIds field to given value.

### HasLanNetworkIds

`func (o *QueryOswDateByNetworkVO) HasLanNetworkIds() bool`

HasLanNetworkIds returns a boolean if a field has been set.

### GetMultiSearchMap

`func (o *QueryOswDateByNetworkVO) GetMultiSearchMap() map[string]string`

GetMultiSearchMap returns the MultiSearchMap field if non-nil, zero value otherwise.

### GetMultiSearchMapOk

`func (o *QueryOswDateByNetworkVO) GetMultiSearchMapOk() (*map[string]string, bool)`

GetMultiSearchMapOk returns a tuple with the MultiSearchMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSearchMap

`func (o *QueryOswDateByNetworkVO) SetMultiSearchMap(v map[string]string)`

SetMultiSearchMap sets MultiSearchMap field to given value.

### HasMultiSearchMap

`func (o *QueryOswDateByNetworkVO) HasMultiSearchMap() bool`

HasMultiSearchMap returns a boolean if a field has been set.

### GetSearchField

`func (o *QueryOswDateByNetworkVO) GetSearchField() string`

GetSearchField returns the SearchField field if non-nil, zero value otherwise.

### GetSearchFieldOk

`func (o *QueryOswDateByNetworkVO) GetSearchFieldOk() (*string, bool)`

GetSearchFieldOk returns a tuple with the SearchField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchField

`func (o *QueryOswDateByNetworkVO) SetSearchField(v string)`

SetSearchField sets SearchField field to given value.

### HasSearchField

`func (o *QueryOswDateByNetworkVO) HasSearchField() bool`

HasSearchField returns a boolean if a field has been set.

### GetSearchKey

`func (o *QueryOswDateByNetworkVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *QueryOswDateByNetworkVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *QueryOswDateByNetworkVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *QueryOswDateByNetworkVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *QueryOswDateByNetworkVO) GetSorts() map[string]string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *QueryOswDateByNetworkVO) GetSortsOk() (*map[string]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *QueryOswDateByNetworkVO) SetSorts(v map[string]string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *QueryOswDateByNetworkVO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


