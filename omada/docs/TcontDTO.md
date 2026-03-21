# TcontDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbaId** | **int32** | The DBA template ID bound to the T-cont. DbaId should be a within the range of 0 to 512. Currently existing DBA templates, including DBA system templates, can also be referenced by the T-cont. | 
**GemPortIds** | Pointer to **[]int32** | A list of GEM Port IDs bound to this T-Cont. | [optional] 
**IsInUse** | Pointer to **bool** | Whether the T-cont has been used. If it has been used, it cannot be deleted. True indicates it has been used, while false indicates it has not been used.  | [optional] 
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfile should be within the range of 1 to 512. | [optional] 
**TcontId** | **int32** | T-cont ID should be within the range of 1 to 127 and should not be null | 

## Methods

### NewTcontDTO

`func NewTcontDTO(dbaId int32, tcontId int32, ) *TcontDTO`

NewTcontDTO instantiates a new TcontDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcontDTOWithDefaults

`func NewTcontDTOWithDefaults() *TcontDTO`

NewTcontDTOWithDefaults instantiates a new TcontDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbaId

`func (o *TcontDTO) GetDbaId() int32`

GetDbaId returns the DbaId field if non-nil, zero value otherwise.

### GetDbaIdOk

`func (o *TcontDTO) GetDbaIdOk() (*int32, bool)`

GetDbaIdOk returns a tuple with the DbaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaId

`func (o *TcontDTO) SetDbaId(v int32)`

SetDbaId sets DbaId field to given value.


### GetGemPortIds

`func (o *TcontDTO) GetGemPortIds() []int32`

GetGemPortIds returns the GemPortIds field if non-nil, zero value otherwise.

### GetGemPortIdsOk

`func (o *TcontDTO) GetGemPortIdsOk() (*[]int32, bool)`

GetGemPortIdsOk returns a tuple with the GemPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortIds

`func (o *TcontDTO) SetGemPortIds(v []int32)`

SetGemPortIds sets GemPortIds field to given value.

### HasGemPortIds

`func (o *TcontDTO) HasGemPortIds() bool`

HasGemPortIds returns a boolean if a field has been set.

### GetIsInUse

`func (o *TcontDTO) GetIsInUse() bool`

GetIsInUse returns the IsInUse field if non-nil, zero value otherwise.

### GetIsInUseOk

`func (o *TcontDTO) GetIsInUseOk() (*bool, bool)`

GetIsInUseOk returns a tuple with the IsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInUse

`func (o *TcontDTO) SetIsInUse(v bool)`

SetIsInUse sets IsInUse field to given value.

### HasIsInUse

`func (o *TcontDTO) HasIsInUse() bool`

HasIsInUse returns a boolean if a field has been set.

### GetLineProfileId

`func (o *TcontDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *TcontDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *TcontDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *TcontDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.

### GetTcontId

`func (o *TcontDTO) GetTcontId() int32`

GetTcontId returns the TcontId field if non-nil, zero value otherwise.

### GetTcontIdOk

`func (o *TcontDTO) GetTcontIdOk() (*int32, bool)`

GetTcontIdOk returns a tuple with the TcontId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcontId

`func (o *TcontDTO) SetTcontId(v int32)`

SetTcontId sets TcontId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


