# PageResponsePonAutoAuthenticationRuleConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]PonAutoAuthenticationRuleConfigDTO**](PonAutoAuthenticationRuleConfigDTO.md) | Content | [optional] 
**First** | Pointer to **bool** | Whether is in the first page | [optional] 
**Last** | Pointer to **bool** | Whether is in the last page | [optional] 
**Number** | Pointer to **int32** | Number of the current page,starting from 0. | [optional] 
**NumberOfElements** | Pointer to **int32** | Actual number of current page | [optional] 
**Size** | Pointer to **int32** | Page size | [optional] 
**TotalElements** | Pointer to **int64** | Total elements | [optional] 
**TotalPages** | Pointer to **int32** | Total page number | [optional] 

## Methods

### NewPageResponsePonAutoAuthenticationRuleConfigDTO

`func NewPageResponsePonAutoAuthenticationRuleConfigDTO() *PageResponsePonAutoAuthenticationRuleConfigDTO`

NewPageResponsePonAutoAuthenticationRuleConfigDTO instantiates a new PageResponsePonAutoAuthenticationRuleConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponsePonAutoAuthenticationRuleConfigDTOWithDefaults

`func NewPageResponsePonAutoAuthenticationRuleConfigDTOWithDefaults() *PageResponsePonAutoAuthenticationRuleConfigDTO`

NewPageResponsePonAutoAuthenticationRuleConfigDTOWithDefaults instantiates a new PageResponsePonAutoAuthenticationRuleConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetContent() []PonAutoAuthenticationRuleConfigDTO`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetContentOk() (*[]PonAutoAuthenticationRuleConfigDTO, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) SetContent(v []PonAutoAuthenticationRuleConfigDTO)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFirst

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNumber

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetSize

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotalElements

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetTotalElements() int64`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetTotalElementsOk() (*int64, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) SetTotalElements(v int64)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PageResponsePonAutoAuthenticationRuleConfigDTO) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


