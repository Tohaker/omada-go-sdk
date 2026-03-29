# TcontListQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfile should be within the range of 1 to 512. | [optional] 
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewTcontListQueryDTO

`func NewTcontListQueryDTO() *TcontListQueryDTO`

NewTcontListQueryDTO instantiates a new TcontListQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcontListQueryDTOWithDefaults

`func NewTcontListQueryDTOWithDefaults() *TcontListQueryDTO`

NewTcontListQueryDTOWithDefaults instantiates a new TcontListQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineProfileId

`func (o *TcontListQueryDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *TcontListQueryDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *TcontListQueryDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *TcontListQueryDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.

### GetSearchKey

`func (o *TcontListQueryDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *TcontListQueryDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *TcontListQueryDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *TcontListQueryDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *TcontListQueryDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *TcontListQueryDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *TcontListQueryDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *TcontListQueryDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


