# OnuManagementPageQueryRequestDTO

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

### NewOnuManagementPageQueryRequestDTO

`func NewOnuManagementPageQueryRequestDTO(ponPort string, ) *OnuManagementPageQueryRequestDTO`

NewOnuManagementPageQueryRequestDTO instantiates a new OnuManagementPageQueryRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuManagementPageQueryRequestDTOWithDefaults

`func NewOnuManagementPageQueryRequestDTOWithDefaults() *OnuManagementPageQueryRequestDTO`

NewOnuManagementPageQueryRequestDTOWithDefaults instantiates a new OnuManagementPageQueryRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *OnuManagementPageQueryRequestDTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OnuManagementPageQueryRequestDTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OnuManagementPageQueryRequestDTO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OnuManagementPageQueryRequestDTO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPonPort

`func (o *OnuManagementPageQueryRequestDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *OnuManagementPageQueryRequestDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *OnuManagementPageQueryRequestDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.


### GetSearchFields

`func (o *OnuManagementPageQueryRequestDTO) GetSearchFields() []string`

GetSearchFields returns the SearchFields field if non-nil, zero value otherwise.

### GetSearchFieldsOk

`func (o *OnuManagementPageQueryRequestDTO) GetSearchFieldsOk() (*[]string, bool)`

GetSearchFieldsOk returns a tuple with the SearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFields

`func (o *OnuManagementPageQueryRequestDTO) SetSearchFields(v []string)`

SetSearchFields sets SearchFields field to given value.

### HasSearchFields

`func (o *OnuManagementPageQueryRequestDTO) HasSearchFields() bool`

HasSearchFields returns a boolean if a field has been set.

### GetSearchKey

`func (o *OnuManagementPageQueryRequestDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *OnuManagementPageQueryRequestDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *OnuManagementPageQueryRequestDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *OnuManagementPageQueryRequestDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSize

`func (o *OnuManagementPageQueryRequestDTO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OnuManagementPageQueryRequestDTO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OnuManagementPageQueryRequestDTO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *OnuManagementPageQueryRequestDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSorts

`func (o *OnuManagementPageQueryRequestDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *OnuManagementPageQueryRequestDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *OnuManagementPageQueryRequestDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *OnuManagementPageQueryRequestDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


