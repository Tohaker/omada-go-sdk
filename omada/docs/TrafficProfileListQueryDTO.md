# TrafficProfileListQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFields** | Pointer to **[]string** | Fields to search | [optional] 
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewTrafficProfileListQueryDTO

`func NewTrafficProfileListQueryDTO() *TrafficProfileListQueryDTO`

NewTrafficProfileListQueryDTO instantiates a new TrafficProfileListQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficProfileListQueryDTOWithDefaults

`func NewTrafficProfileListQueryDTOWithDefaults() *TrafficProfileListQueryDTO`

NewTrafficProfileListQueryDTOWithDefaults instantiates a new TrafficProfileListQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFields

`func (o *TrafficProfileListQueryDTO) GetSearchFields() []string`

GetSearchFields returns the SearchFields field if non-nil, zero value otherwise.

### GetSearchFieldsOk

`func (o *TrafficProfileListQueryDTO) GetSearchFieldsOk() (*[]string, bool)`

GetSearchFieldsOk returns a tuple with the SearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFields

`func (o *TrafficProfileListQueryDTO) SetSearchFields(v []string)`

SetSearchFields sets SearchFields field to given value.

### HasSearchFields

`func (o *TrafficProfileListQueryDTO) HasSearchFields() bool`

HasSearchFields returns a boolean if a field has been set.

### GetSearchKey

`func (o *TrafficProfileListQueryDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *TrafficProfileListQueryDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *TrafficProfileListQueryDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *TrafficProfileListQueryDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *TrafficProfileListQueryDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *TrafficProfileListQueryDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *TrafficProfileListQueryDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *TrafficProfileListQueryDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


