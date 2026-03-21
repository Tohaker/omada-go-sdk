# ModifyDstDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | DST config status; If false, other parameters are not required. | [optional] 
**End** | Pointer to [**DstTimeOpenApiDTO**](DstTimeOpenApiDTO.md) |  | [optional] 
**Mode** | Pointer to **int32** | DST config mode; If disable, other parameters are not required. 0: disable, 1: auto, 2: manually | [optional] 
**Offset** | Pointer to **int64** | DST offset config(Unit: ms); It should be a value as follows: [1800000, 3600000, 5400000, 7200000]. | [optional] 
**Start** | Pointer to [**DstTimeOpenApiDTO**](DstTimeOpenApiDTO.md) |  | [optional] 

## Methods

### NewModifyDstDTO

`func NewModifyDstDTO() *ModifyDstDTO`

NewModifyDstDTO instantiates a new ModifyDstDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDstDTOWithDefaults

`func NewModifyDstDTOWithDefaults() *ModifyDstDTO`

NewModifyDstDTOWithDefaults instantiates a new ModifyDstDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ModifyDstDTO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ModifyDstDTO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ModifyDstDTO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ModifyDstDTO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnd

`func (o *ModifyDstDTO) GetEnd() DstTimeOpenApiDTO`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ModifyDstDTO) GetEndOk() (*DstTimeOpenApiDTO, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ModifyDstDTO) SetEnd(v DstTimeOpenApiDTO)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ModifyDstDTO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetMode

`func (o *ModifyDstDTO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ModifyDstDTO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ModifyDstDTO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ModifyDstDTO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOffset

`func (o *ModifyDstDTO) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ModifyDstDTO) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ModifyDstDTO) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ModifyDstDTO) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStart

`func (o *ModifyDstDTO) GetStart() DstTimeOpenApiDTO`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ModifyDstDTO) GetStartOk() (*DstTimeOpenApiDTO, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ModifyDstDTO) SetStart(v DstTimeOpenApiDTO)`

SetStart sets Start field to given value.

### HasStart

`func (o *ModifyDstDTO) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


