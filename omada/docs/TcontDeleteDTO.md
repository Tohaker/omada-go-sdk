# TcontDeleteDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]int32** | ID list | 
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfile should be within the range of 1 to 512. | [optional] 

## Methods

### NewTcontDeleteDTO

`func NewTcontDeleteDTO(ids []int32, ) *TcontDeleteDTO`

NewTcontDeleteDTO instantiates a new TcontDeleteDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcontDeleteDTOWithDefaults

`func NewTcontDeleteDTOWithDefaults() *TcontDeleteDTO`

NewTcontDeleteDTOWithDefaults instantiates a new TcontDeleteDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *TcontDeleteDTO) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *TcontDeleteDTO) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *TcontDeleteDTO) SetIds(v []int32)`

SetIds sets Ids field to given value.


### GetLineProfileId

`func (o *TcontDeleteDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *TcontDeleteDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *TcontDeleteDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *TcontDeleteDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


