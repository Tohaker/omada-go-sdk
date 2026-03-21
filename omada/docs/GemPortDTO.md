# GemPortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypt** | **string** | Whether to enable data encryption feature. Encrypt should be a value as follows: ENABLE;DISABLE. Default value: DISABLE. | 
**GemMappingId** | Pointer to **[]int32** | Gem Mapping Id using the gem port should be null | [optional] 
**GemPortId** | **int32** | Gem port ID should be within the range of 1 to 1023 | 
**IsInUse** | Pointer to **bool** | Whether the Gem Port is in use: if it is in use, deletion is not allowed. true indicates it is in use, while false indicates it is not in use. This field should not be provided. | [optional] 
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfile should be within the range of 1 to 512. | [optional] 
**TcontId** | **int32** | The ID of the T-cont to which the Gem Port is bound. TcontId should be s current existing T-cont ID. | 

## Methods

### NewGemPortDTO

`func NewGemPortDTO(encrypt string, gemPortId int32, tcontId int32, ) *GemPortDTO`

NewGemPortDTO instantiates a new GemPortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemPortDTOWithDefaults

`func NewGemPortDTOWithDefaults() *GemPortDTO`

NewGemPortDTOWithDefaults instantiates a new GemPortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypt

`func (o *GemPortDTO) GetEncrypt() string`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *GemPortDTO) GetEncryptOk() (*string, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *GemPortDTO) SetEncrypt(v string)`

SetEncrypt sets Encrypt field to given value.


### GetGemMappingId

`func (o *GemPortDTO) GetGemMappingId() []int32`

GetGemMappingId returns the GemMappingId field if non-nil, zero value otherwise.

### GetGemMappingIdOk

`func (o *GemPortDTO) GetGemMappingIdOk() (*[]int32, bool)`

GetGemMappingIdOk returns a tuple with the GemMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemMappingId

`func (o *GemPortDTO) SetGemMappingId(v []int32)`

SetGemMappingId sets GemMappingId field to given value.

### HasGemMappingId

`func (o *GemPortDTO) HasGemMappingId() bool`

HasGemMappingId returns a boolean if a field has been set.

### GetGemPortId

`func (o *GemPortDTO) GetGemPortId() int32`

GetGemPortId returns the GemPortId field if non-nil, zero value otherwise.

### GetGemPortIdOk

`func (o *GemPortDTO) GetGemPortIdOk() (*int32, bool)`

GetGemPortIdOk returns a tuple with the GemPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortId

`func (o *GemPortDTO) SetGemPortId(v int32)`

SetGemPortId sets GemPortId field to given value.


### GetIsInUse

`func (o *GemPortDTO) GetIsInUse() bool`

GetIsInUse returns the IsInUse field if non-nil, zero value otherwise.

### GetIsInUseOk

`func (o *GemPortDTO) GetIsInUseOk() (*bool, bool)`

GetIsInUseOk returns a tuple with the IsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInUse

`func (o *GemPortDTO) SetIsInUse(v bool)`

SetIsInUse sets IsInUse field to given value.

### HasIsInUse

`func (o *GemPortDTO) HasIsInUse() bool`

HasIsInUse returns a boolean if a field has been set.

### GetLineProfileId

`func (o *GemPortDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *GemPortDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *GemPortDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *GemPortDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.

### GetTcontId

`func (o *GemPortDTO) GetTcontId() int32`

GetTcontId returns the TcontId field if non-nil, zero value otherwise.

### GetTcontIdOk

`func (o *GemPortDTO) GetTcontIdOk() (*int32, bool)`

GetTcontIdOk returns a tuple with the TcontId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcontId

`func (o *GemPortDTO) SetTcontId(v int32)`

SetTcontId sets TcontId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


