# TcontModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbaId** | Pointer to **int32** | The DBA template ID bound to the T-cont. DbaId should be within the range of 0 to 512. Currently existing DBA templates, including DBA system templates, can also be referenced by the T-cont. | [optional] 
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile,lineProfile should be within the range of 1 to 512. | [optional] 
**TcontId** | **int32** | T-cont ID should be within the range of 1 to 127 and should not be null | 

## Methods

### NewTcontModifyDTO

`func NewTcontModifyDTO(tcontId int32, ) *TcontModifyDTO`

NewTcontModifyDTO instantiates a new TcontModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcontModifyDTOWithDefaults

`func NewTcontModifyDTOWithDefaults() *TcontModifyDTO`

NewTcontModifyDTOWithDefaults instantiates a new TcontModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbaId

`func (o *TcontModifyDTO) GetDbaId() int32`

GetDbaId returns the DbaId field if non-nil, zero value otherwise.

### GetDbaIdOk

`func (o *TcontModifyDTO) GetDbaIdOk() (*int32, bool)`

GetDbaIdOk returns a tuple with the DbaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaId

`func (o *TcontModifyDTO) SetDbaId(v int32)`

SetDbaId sets DbaId field to given value.

### HasDbaId

`func (o *TcontModifyDTO) HasDbaId() bool`

HasDbaId returns a boolean if a field has been set.

### GetLineProfileId

`func (o *TcontModifyDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *TcontModifyDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *TcontModifyDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *TcontModifyDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.

### GetTcontId

`func (o *TcontModifyDTO) GetTcontId() int32`

GetTcontId returns the TcontId field if non-nil, zero value otherwise.

### GetTcontIdOk

`func (o *TcontModifyDTO) GetTcontIdOk() (*int32, bool)`

GetTcontIdOk returns a tuple with the TcontId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcontId

`func (o *TcontModifyDTO) SetTcontId(v int32)`

SetTcontId sets TcontId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


