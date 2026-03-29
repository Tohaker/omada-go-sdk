# GemPortModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypt** | Pointer to **string** | Whether to enable data encryption feature. Encrypt should be a value as follows:ENABLE;DISABLE. Default value: DISABLE. | [optional] 
**GemPortId** | **int32** | Gem port ID should be within the range of 1 to 1023 | 
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfile should be within the range of 1 to 512. | [optional] 
**TcontId** | Pointer to **int32** | The ID of the T-cont to which the Gem Port is bound. TcontId should be a current existing T-cont ID. | [optional] 

## Methods

### NewGemPortModifyDTO

`func NewGemPortModifyDTO(gemPortId int32, ) *GemPortModifyDTO`

NewGemPortModifyDTO instantiates a new GemPortModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemPortModifyDTOWithDefaults

`func NewGemPortModifyDTOWithDefaults() *GemPortModifyDTO`

NewGemPortModifyDTOWithDefaults instantiates a new GemPortModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypt

`func (o *GemPortModifyDTO) GetEncrypt() string`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *GemPortModifyDTO) GetEncryptOk() (*string, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *GemPortModifyDTO) SetEncrypt(v string)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *GemPortModifyDTO) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetGemPortId

`func (o *GemPortModifyDTO) GetGemPortId() int32`

GetGemPortId returns the GemPortId field if non-nil, zero value otherwise.

### GetGemPortIdOk

`func (o *GemPortModifyDTO) GetGemPortIdOk() (*int32, bool)`

GetGemPortIdOk returns a tuple with the GemPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortId

`func (o *GemPortModifyDTO) SetGemPortId(v int32)`

SetGemPortId sets GemPortId field to given value.


### GetLineProfileId

`func (o *GemPortModifyDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *GemPortModifyDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *GemPortModifyDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *GemPortModifyDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.

### GetTcontId

`func (o *GemPortModifyDTO) GetTcontId() int32`

GetTcontId returns the TcontId field if non-nil, zero value otherwise.

### GetTcontIdOk

`func (o *GemPortModifyDTO) GetTcontIdOk() (*int32, bool)`

GetTcontIdOk returns a tuple with the TcontId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcontId

`func (o *GemPortModifyDTO) SetTcontId(v int32)`

SetTcontId sets TcontId field to given value.

### HasTcontId

`func (o *GemPortModifyDTO) HasTcontId() bool`

HasTcontId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


