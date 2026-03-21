# OnuManagementListQueryRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PonPort** | **string** | Pon port collection.e.g.,&#39;GPON 1/1/1&#39;、&#39;GPON 1/1/1-3&#39;、&#39;GPON 1/1/1-3,GPON 1/1/5,GPON 1/1/7-8&#39; | 
**SearchFields** | Pointer to **[]string** | Query filter fields. | [optional] 
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewOnuManagementListQueryRequestDTO

`func NewOnuManagementListQueryRequestDTO(ponPort string, ) *OnuManagementListQueryRequestDTO`

NewOnuManagementListQueryRequestDTO instantiates a new OnuManagementListQueryRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuManagementListQueryRequestDTOWithDefaults

`func NewOnuManagementListQueryRequestDTOWithDefaults() *OnuManagementListQueryRequestDTO`

NewOnuManagementListQueryRequestDTOWithDefaults instantiates a new OnuManagementListQueryRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPonPort

`func (o *OnuManagementListQueryRequestDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *OnuManagementListQueryRequestDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *OnuManagementListQueryRequestDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.


### GetSearchFields

`func (o *OnuManagementListQueryRequestDTO) GetSearchFields() []string`

GetSearchFields returns the SearchFields field if non-nil, zero value otherwise.

### GetSearchFieldsOk

`func (o *OnuManagementListQueryRequestDTO) GetSearchFieldsOk() (*[]string, bool)`

GetSearchFieldsOk returns a tuple with the SearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFields

`func (o *OnuManagementListQueryRequestDTO) SetSearchFields(v []string)`

SetSearchFields sets SearchFields field to given value.

### HasSearchFields

`func (o *OnuManagementListQueryRequestDTO) HasSearchFields() bool`

HasSearchFields returns a boolean if a field has been set.

### GetSearchKey

`func (o *OnuManagementListQueryRequestDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *OnuManagementListQueryRequestDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *OnuManagementListQueryRequestDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *OnuManagementListQueryRequestDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *OnuManagementListQueryRequestDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *OnuManagementListQueryRequestDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *OnuManagementListQueryRequestDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *OnuManagementListQueryRequestDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


