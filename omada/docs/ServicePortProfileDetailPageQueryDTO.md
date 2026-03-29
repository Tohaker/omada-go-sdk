# ServicePortProfileDetailPageQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Current page number must be greater or equal to 0, with a default value of 0. | [optional] 
**SearchFields** | Pointer to **[]string** | Query Filter Fields. | [optional] 
**SearchKey** | Pointer to **string** | Fuzzy query parameters | [optional] 
**ServicePortProfileId** | **string** | Service Port Profile ID. The servicePortProfileId should be within the range of 1 to 127. | 
**Size** | Pointer to **int32** | Page size must be greater or equal to 1,with the default value of 10. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewServicePortProfileDetailPageQueryDTO

`func NewServicePortProfileDetailPageQueryDTO(servicePortProfileId string, ) *ServicePortProfileDetailPageQueryDTO`

NewServicePortProfileDetailPageQueryDTO instantiates a new ServicePortProfileDetailPageQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortProfileDetailPageQueryDTOWithDefaults

`func NewServicePortProfileDetailPageQueryDTOWithDefaults() *ServicePortProfileDetailPageQueryDTO`

NewServicePortProfileDetailPageQueryDTOWithDefaults instantiates a new ServicePortProfileDetailPageQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ServicePortProfileDetailPageQueryDTO) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ServicePortProfileDetailPageQueryDTO) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ServicePortProfileDetailPageQueryDTO) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ServicePortProfileDetailPageQueryDTO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSearchFields

`func (o *ServicePortProfileDetailPageQueryDTO) GetSearchFields() []string`

GetSearchFields returns the SearchFields field if non-nil, zero value otherwise.

### GetSearchFieldsOk

`func (o *ServicePortProfileDetailPageQueryDTO) GetSearchFieldsOk() (*[]string, bool)`

GetSearchFieldsOk returns a tuple with the SearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFields

`func (o *ServicePortProfileDetailPageQueryDTO) SetSearchFields(v []string)`

SetSearchFields sets SearchFields field to given value.

### HasSearchFields

`func (o *ServicePortProfileDetailPageQueryDTO) HasSearchFields() bool`

HasSearchFields returns a boolean if a field has been set.

### GetSearchKey

`func (o *ServicePortProfileDetailPageQueryDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *ServicePortProfileDetailPageQueryDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *ServicePortProfileDetailPageQueryDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *ServicePortProfileDetailPageQueryDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetServicePortProfileId

`func (o *ServicePortProfileDetailPageQueryDTO) GetServicePortProfileId() string`

GetServicePortProfileId returns the ServicePortProfileId field if non-nil, zero value otherwise.

### GetServicePortProfileIdOk

`func (o *ServicePortProfileDetailPageQueryDTO) GetServicePortProfileIdOk() (*string, bool)`

GetServicePortProfileIdOk returns a tuple with the ServicePortProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfileId

`func (o *ServicePortProfileDetailPageQueryDTO) SetServicePortProfileId(v string)`

SetServicePortProfileId sets ServicePortProfileId field to given value.


### GetSize

`func (o *ServicePortProfileDetailPageQueryDTO) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ServicePortProfileDetailPageQueryDTO) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ServicePortProfileDetailPageQueryDTO) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ServicePortProfileDetailPageQueryDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSorts

`func (o *ServicePortProfileDetailPageQueryDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ServicePortProfileDetailPageQueryDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ServicePortProfileDetailPageQueryDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ServicePortProfileDetailPageQueryDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


