# PageResponseServicePortProfileDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]ServicePortProfileDTO**](ServicePortProfileDTO.md) | Content | [optional] 
**First** | Pointer to **bool** | Whether is in the first page | [optional] 
**Last** | Pointer to **bool** | Whether is in the last page | [optional] 
**Number** | Pointer to **int32** | Number of the current page,starting from 0. | [optional] 
**NumberOfElements** | Pointer to **int32** | Actual number of current page | [optional] 
**Size** | Pointer to **int32** | Page size | [optional] 
**TotalElements** | Pointer to **int64** | Total elements | [optional] 
**TotalPages** | Pointer to **int32** | Total page number | [optional] 

## Methods

### NewPageResponseServicePortProfileDTO

`func NewPageResponseServicePortProfileDTO() *PageResponseServicePortProfileDTO`

NewPageResponseServicePortProfileDTO instantiates a new PageResponseServicePortProfileDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponseServicePortProfileDTOWithDefaults

`func NewPageResponseServicePortProfileDTOWithDefaults() *PageResponseServicePortProfileDTO`

NewPageResponseServicePortProfileDTOWithDefaults instantiates a new PageResponseServicePortProfileDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PageResponseServicePortProfileDTO) GetContent() []ServicePortProfileDTO`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageResponseServicePortProfileDTO) GetContentOk() (*[]ServicePortProfileDTO, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageResponseServicePortProfileDTO) SetContent(v []ServicePortProfileDTO)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageResponseServicePortProfileDTO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFirst

`func (o *PageResponseServicePortProfileDTO) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PageResponseServicePortProfileDTO) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PageResponseServicePortProfileDTO) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PageResponseServicePortProfileDTO) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PageResponseServicePortProfileDTO) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PageResponseServicePortProfileDTO) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PageResponseServicePortProfileDTO) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PageResponseServicePortProfileDTO) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNumber

`func (o *PageResponseServicePortProfileDTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageResponseServicePortProfileDTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageResponseServicePortProfileDTO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageResponseServicePortProfileDTO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PageResponseServicePortProfileDTO) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PageResponseServicePortProfileDTO) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PageResponseServicePortProfileDTO) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PageResponseServicePortProfileDTO) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PageResponseServicePortProfileDTO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageResponseServicePortProfileDTO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageResponseServicePortProfileDTO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageResponseServicePortProfileDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotalElements

`func (o *PageResponseServicePortProfileDTO) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PageResponseServicePortProfileDTO) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PageResponseServicePortProfileDTO) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PageResponseServicePortProfileDTO) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PageResponseServicePortProfileDTO) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PageResponseServicePortProfileDTO) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PageResponseServicePortProfileDTO) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PageResponseServicePortProfileDTO) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


