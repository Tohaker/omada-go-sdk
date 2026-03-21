# OnuRegisterAuthenticationPageQueryRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Current page number must be greater or equal to 0, with a default value of 0. | [optional] 
**PonPort** | **string** | Pon port.e.g.GPON 1/0/1 | 
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**Size** | Pointer to **int32** | Page size must be greater or equal to 1,with the default value of 10. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewOnuRegisterAuthenticationPageQueryRequestDTO

`func NewOnuRegisterAuthenticationPageQueryRequestDTO(ponPort string, ) *OnuRegisterAuthenticationPageQueryRequestDTO`

NewOnuRegisterAuthenticationPageQueryRequestDTO instantiates a new OnuRegisterAuthenticationPageQueryRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuRegisterAuthenticationPageQueryRequestDTOWithDefaults

`func NewOnuRegisterAuthenticationPageQueryRequestDTOWithDefaults() *OnuRegisterAuthenticationPageQueryRequestDTO`

NewOnuRegisterAuthenticationPageQueryRequestDTOWithDefaults instantiates a new OnuRegisterAuthenticationPageQueryRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPonPort

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.


### GetSearchKey

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSize

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSorts

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *OnuRegisterAuthenticationPageQueryRequestDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


