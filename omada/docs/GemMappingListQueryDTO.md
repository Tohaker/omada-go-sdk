# GemMappingListQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfileId should be within the range of 1 to 512. | [optional] 
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewGemMappingListQueryDTO

`func NewGemMappingListQueryDTO() *GemMappingListQueryDTO`

NewGemMappingListQueryDTO instantiates a new GemMappingListQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemMappingListQueryDTOWithDefaults

`func NewGemMappingListQueryDTOWithDefaults() *GemMappingListQueryDTO`

NewGemMappingListQueryDTOWithDefaults instantiates a new GemMappingListQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineProfileId

`func (o *GemMappingListQueryDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *GemMappingListQueryDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *GemMappingListQueryDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *GemMappingListQueryDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.

### GetSearchKey

`func (o *GemMappingListQueryDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *GemMappingListQueryDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *GemMappingListQueryDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *GemMappingListQueryDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *GemMappingListQueryDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *GemMappingListQueryDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *GemMappingListQueryDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *GemMappingListQueryDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


