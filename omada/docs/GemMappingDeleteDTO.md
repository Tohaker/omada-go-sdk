# GemMappingDeleteDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GemMappingList** | [**[]GemMappingDeleteItem**](GemMappingDeleteItem.md) | gemMappingList | 
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfileId should be within the range of 1 to 512. | [optional] 

## Methods

### NewGemMappingDeleteDTO

`func NewGemMappingDeleteDTO(gemMappingList []GemMappingDeleteItem, ) *GemMappingDeleteDTO`

NewGemMappingDeleteDTO instantiates a new GemMappingDeleteDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemMappingDeleteDTOWithDefaults

`func NewGemMappingDeleteDTOWithDefaults() *GemMappingDeleteDTO`

NewGemMappingDeleteDTOWithDefaults instantiates a new GemMappingDeleteDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGemMappingList

`func (o *GemMappingDeleteDTO) GetGemMappingList() []GemMappingDeleteItem`

GetGemMappingList returns the GemMappingList field if non-nil, zero value otherwise.

### GetGemMappingListOk

`func (o *GemMappingDeleteDTO) GetGemMappingListOk() (*[]GemMappingDeleteItem, bool)`

GetGemMappingListOk returns a tuple with the GemMappingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemMappingList

`func (o *GemMappingDeleteDTO) SetGemMappingList(v []GemMappingDeleteItem)`

SetGemMappingList sets GemMappingList field to given value.


### GetLineProfileId

`func (o *GemMappingDeleteDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *GemMappingDeleteDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *GemMappingDeleteDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *GemMappingDeleteDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


