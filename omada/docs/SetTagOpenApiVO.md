# SetTagOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | Device MAC list, like AA-BB-CC-DD-EE-FF | [optional] 
**TagIds** | **[]string** | Tag ID list | 

## Methods

### NewSetTagOpenApiVO

`func NewSetTagOpenApiVO(tagIds []string, ) *SetTagOpenApiVO`

NewSetTagOpenApiVO instantiates a new SetTagOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetTagOpenApiVOWithDefaults

`func NewSetTagOpenApiVOWithDefaults() *SetTagOpenApiVO`

NewSetTagOpenApiVOWithDefaults instantiates a new SetTagOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *SetTagOpenApiVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *SetTagOpenApiVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *SetTagOpenApiVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *SetTagOpenApiVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetTagIds

`func (o *SetTagOpenApiVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *SetTagOpenApiVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *SetTagOpenApiVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


