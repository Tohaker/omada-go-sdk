# GemPortPageQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfile should be within the range of 1 to 512. | [optional] 
**Number** | Pointer to **int32** | Current page number must be greater or equal to 0, with a default value of 0. | [optional] 
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**Size** | Pointer to **int32** | Page size must be greater or equal to 1,with the default value of 10. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewGemPortPageQueryDTO

`func NewGemPortPageQueryDTO() *GemPortPageQueryDTO`

NewGemPortPageQueryDTO instantiates a new GemPortPageQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemPortPageQueryDTOWithDefaults

`func NewGemPortPageQueryDTOWithDefaults() *GemPortPageQueryDTO`

NewGemPortPageQueryDTOWithDefaults instantiates a new GemPortPageQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineProfileId

`func (o *GemPortPageQueryDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *GemPortPageQueryDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *GemPortPageQueryDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *GemPortPageQueryDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.

### GetNumber

`func (o *GemPortPageQueryDTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GemPortPageQueryDTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GemPortPageQueryDTO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GemPortPageQueryDTO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSearchKey

`func (o *GemPortPageQueryDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *GemPortPageQueryDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *GemPortPageQueryDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *GemPortPageQueryDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSize

`func (o *GemPortPageQueryDTO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GemPortPageQueryDTO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GemPortPageQueryDTO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GemPortPageQueryDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSorts

`func (o *GemPortPageQueryDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *GemPortPageQueryDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *GemPortPageQueryDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *GemPortPageQueryDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


