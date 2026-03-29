# PageResponseLineProfileVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]LineProfileVO**](LineProfileVO.md) | Content | [optional] 
**First** | Pointer to **bool** | Whether is in the first page | [optional] 
**Last** | Pointer to **bool** | Whether is in the last page | [optional] 
**Number** | Pointer to **int32** | Number of the current page,starting from 0. | [optional] 
**NumberOfElements** | Pointer to **int32** | Actual number of current page | [optional] 
**Size** | Pointer to **int32** | Page size | [optional] 
**TotalElements** | Pointer to **int64** | Total elements | [optional] 
**TotalPages** | Pointer to **int32** | Total page number | [optional] 

## Methods

### NewPageResponseLineProfileVO

`func NewPageResponseLineProfileVO() *PageResponseLineProfileVO`

NewPageResponseLineProfileVO instantiates a new PageResponseLineProfileVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponseLineProfileVOWithDefaults

`func NewPageResponseLineProfileVOWithDefaults() *PageResponseLineProfileVO`

NewPageResponseLineProfileVOWithDefaults instantiates a new PageResponseLineProfileVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PageResponseLineProfileVO) GetContent() []LineProfileVO`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageResponseLineProfileVO) GetContentOk() (*[]LineProfileVO, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageResponseLineProfileVO) SetContent(v []LineProfileVO)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageResponseLineProfileVO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFirst

`func (o *PageResponseLineProfileVO) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PageResponseLineProfileVO) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PageResponseLineProfileVO) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PageResponseLineProfileVO) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PageResponseLineProfileVO) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PageResponseLineProfileVO) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PageResponseLineProfileVO) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PageResponseLineProfileVO) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNumber

`func (o *PageResponseLineProfileVO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageResponseLineProfileVO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageResponseLineProfileVO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageResponseLineProfileVO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PageResponseLineProfileVO) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PageResponseLineProfileVO) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PageResponseLineProfileVO) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PageResponseLineProfileVO) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PageResponseLineProfileVO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageResponseLineProfileVO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageResponseLineProfileVO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageResponseLineProfileVO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotalElements

`func (o *PageResponseLineProfileVO) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PageResponseLineProfileVO) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PageResponseLineProfileVO) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PageResponseLineProfileVO) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PageResponseLineProfileVO) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PageResponseLineProfileVO) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PageResponseLineProfileVO) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PageResponseLineProfileVO) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


