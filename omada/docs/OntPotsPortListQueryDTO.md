# OntPotsPortListQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchKey** | Pointer to **string** | Query conditions. | [optional] 
**ServiceId** | Pointer to **int32** | The associated Service Profile ID | [optional] 
**Sorts** | Pointer to [**[]QuerySort**](QuerySort.md) | Sorting fields. | [optional] 

## Methods

### NewOntPotsPortListQueryDTO

`func NewOntPotsPortListQueryDTO() *OntPotsPortListQueryDTO`

NewOntPotsPortListQueryDTO instantiates a new OntPotsPortListQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntPotsPortListQueryDTOWithDefaults

`func NewOntPotsPortListQueryDTOWithDefaults() *OntPotsPortListQueryDTO`

NewOntPotsPortListQueryDTOWithDefaults instantiates a new OntPotsPortListQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchKey

`func (o *OntPotsPortListQueryDTO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *OntPotsPortListQueryDTO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *OntPotsPortListQueryDTO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *OntPotsPortListQueryDTO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetServiceId

`func (o *OntPotsPortListQueryDTO) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *OntPotsPortListQueryDTO) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *OntPotsPortListQueryDTO) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *OntPotsPortListQueryDTO) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSorts

`func (o *OntPotsPortListQueryDTO) GetSorts() []QuerySort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *OntPotsPortListQueryDTO) GetSortsOk() (*[]QuerySort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *OntPotsPortListQueryDTO) SetSorts(v []QuerySort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *OntPotsPortListQueryDTO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


