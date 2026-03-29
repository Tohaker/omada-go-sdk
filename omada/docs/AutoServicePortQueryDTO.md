# AutoServicePortQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Current page number must be greater or equal to 0, with a default value of 0. | [optional] 
**SearchField** | Pointer to **string** | Field to search | [optional] 
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**Size** | Pointer to **int32** | Page size must be greater or equal to 1,with the default value of 10. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewAutoServicePortQueryDTO

`func NewAutoServicePortQueryDTO() *AutoServicePortQueryDTO`

NewAutoServicePortQueryDTO instantiates a new AutoServicePortQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoServicePortQueryDTOWithDefaults

`func NewAutoServicePortQueryDTOWithDefaults() *AutoServicePortQueryDTO`

NewAutoServicePortQueryDTOWithDefaults instantiates a new AutoServicePortQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *AutoServicePortQueryDTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AutoServicePortQueryDTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AutoServicePortQueryDTO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AutoServicePortQueryDTO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSearchField

`func (o *AutoServicePortQueryDTO) GetSearchField() string`

GetSearchField returns the SearchField field if non-nil, zero value otherwise.

### GetSearchFieldOk

`func (o *AutoServicePortQueryDTO) GetSearchFieldOk() (*string, bool)`

GetSearchFieldOk returns a tuple with the SearchField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchField

`func (o *AutoServicePortQueryDTO) SetSearchField(v string)`

SetSearchField sets SearchField field to given value.

### HasSearchField

`func (o *AutoServicePortQueryDTO) HasSearchField() bool`

HasSearchField returns a boolean if a field has been set.

### GetSearchKey

`func (o *AutoServicePortQueryDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *AutoServicePortQueryDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *AutoServicePortQueryDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *AutoServicePortQueryDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSize

`func (o *AutoServicePortQueryDTO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AutoServicePortQueryDTO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AutoServicePortQueryDTO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AutoServicePortQueryDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSorts

`func (o *AutoServicePortQueryDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *AutoServicePortQueryDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *AutoServicePortQueryDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *AutoServicePortQueryDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


