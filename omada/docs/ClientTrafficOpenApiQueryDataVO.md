# ClientTrafficOpenApiQueryDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientTrafficFilters** | Pointer to [**[]ClientTrafficFilterOpenApiVO**](ClientTrafficFilterOpenApiVO.md) | Client traffic filtering condition | [optional] 
**End** | Pointer to **int64** | End timestamp, in seconds, such as 1682000000 | [optional] 
**Filters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Page** | **int32** | Start from 1. | 
**PageSize** | **int32** | It should be within the range of 1–100. | 
**SearchField** | Pointer to **string** |  | [optional] 
**SearchKey** | Pointer to **string** | Look for a specific piece of data. | [optional] 
**Sorts** | Pointer to **map[string]string** |  | [optional] 
**Start** | Pointer to **int64** | Start timestamp, in seconds, such as 1682000000 | [optional] 

## Methods

### NewClientTrafficOpenApiQueryDataVO

`func NewClientTrafficOpenApiQueryDataVO(page int32, pageSize int32, ) *ClientTrafficOpenApiQueryDataVO`

NewClientTrafficOpenApiQueryDataVO instantiates a new ClientTrafficOpenApiQueryDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTrafficOpenApiQueryDataVOWithDefaults

`func NewClientTrafficOpenApiQueryDataVOWithDefaults() *ClientTrafficOpenApiQueryDataVO`

NewClientTrafficOpenApiQueryDataVOWithDefaults instantiates a new ClientTrafficOpenApiQueryDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTrafficFilters

`func (o *ClientTrafficOpenApiQueryDataVO) GetClientTrafficFilters() []ClientTrafficFilterOpenApiVO`

GetClientTrafficFilters returns the ClientTrafficFilters field if non-nil, zero value otherwise.

### GetClientTrafficFiltersOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetClientTrafficFiltersOk() (*[]ClientTrafficFilterOpenApiVO, bool)`

GetClientTrafficFiltersOk returns a tuple with the ClientTrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTrafficFilters

`func (o *ClientTrafficOpenApiQueryDataVO) SetClientTrafficFilters(v []ClientTrafficFilterOpenApiVO)`

SetClientTrafficFilters sets ClientTrafficFilters field to given value.

### HasClientTrafficFilters

`func (o *ClientTrafficOpenApiQueryDataVO) HasClientTrafficFilters() bool`

HasClientTrafficFilters returns a boolean if a field has been set.

### GetEnd

`func (o *ClientTrafficOpenApiQueryDataVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ClientTrafficOpenApiQueryDataVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ClientTrafficOpenApiQueryDataVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFilters

`func (o *ClientTrafficOpenApiQueryDataVO) GetFilters() map[string]map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetFiltersOk() (*map[string]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ClientTrafficOpenApiQueryDataVO) SetFilters(v map[string]map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ClientTrafficOpenApiQueryDataVO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPage

`func (o *ClientTrafficOpenApiQueryDataVO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ClientTrafficOpenApiQueryDataVO) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *ClientTrafficOpenApiQueryDataVO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ClientTrafficOpenApiQueryDataVO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetSearchField

`func (o *ClientTrafficOpenApiQueryDataVO) GetSearchField() string`

GetSearchField returns the SearchField field if non-nil, zero value otherwise.

### GetSearchFieldOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetSearchFieldOk() (*string, bool)`

GetSearchFieldOk returns a tuple with the SearchField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchField

`func (o *ClientTrafficOpenApiQueryDataVO) SetSearchField(v string)`

SetSearchField sets SearchField field to given value.

### HasSearchField

`func (o *ClientTrafficOpenApiQueryDataVO) HasSearchField() bool`

HasSearchField returns a boolean if a field has been set.

### GetSearchKey

`func (o *ClientTrafficOpenApiQueryDataVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *ClientTrafficOpenApiQueryDataVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *ClientTrafficOpenApiQueryDataVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *ClientTrafficOpenApiQueryDataVO) GetSorts() map[string]string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetSortsOk() (*map[string]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ClientTrafficOpenApiQueryDataVO) SetSorts(v map[string]string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ClientTrafficOpenApiQueryDataVO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.

### GetStart

`func (o *ClientTrafficOpenApiQueryDataVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientTrafficOpenApiQueryDataVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientTrafficOpenApiQueryDataVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ClientTrafficOpenApiQueryDataVO) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


