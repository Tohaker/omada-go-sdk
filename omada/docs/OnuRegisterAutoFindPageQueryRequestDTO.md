# OnuRegisterAutoFindPageQueryRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Current page number must be greater or equal to 0, with a default value of 0. | [optional] 
**PonPort** | **string** | Pon port collection.e.g.,&#39;GPON 1/1/1&#39;、&#39;GPON 1/1/1-3&#39;、&#39;GPON 1/1/1-3,GPON 1/1/5,GPON 1/1/7-8&#39; | 
**SearchFields** | Pointer to **[]string** | Fields to search | [optional] 
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**Size** | Pointer to **int32** | Page size must be greater or equal to 1,with the default value of 10. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewOnuRegisterAutoFindPageQueryRequestDTO

`func NewOnuRegisterAutoFindPageQueryRequestDTO(ponPort string, ) *OnuRegisterAutoFindPageQueryRequestDTO`

NewOnuRegisterAutoFindPageQueryRequestDTO instantiates a new OnuRegisterAutoFindPageQueryRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuRegisterAutoFindPageQueryRequestDTOWithDefaults

`func NewOnuRegisterAutoFindPageQueryRequestDTOWithDefaults() *OnuRegisterAutoFindPageQueryRequestDTO`

NewOnuRegisterAutoFindPageQueryRequestDTOWithDefaults instantiates a new OnuRegisterAutoFindPageQueryRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPonPort

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.


### GetSearchFields

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetSearchFields() []string`

GetSearchFields returns the SearchFields field if non-nil, zero value otherwise.

### GetSearchFieldsOk

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetSearchFieldsOk() (*[]string, bool)`

GetSearchFieldsOk returns a tuple with the SearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFields

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) SetSearchFields(v []string)`

SetSearchFields sets SearchFields field to given value.

### HasSearchFields

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) HasSearchFields() bool`

HasSearchFields returns a boolean if a field has been set.

### GetSearchKey

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSize

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSorts

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *OnuRegisterAutoFindPageQueryRequestDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


