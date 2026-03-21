# GemMappingDeleteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GemMappingId** | **int32** | GemMappingId should be within the range of 1 to 8,and should not be null | 
**GemPortId** | **int32** | Gem port ID should be within the range of 1 to 1023 | 

## Methods

### NewGemMappingDeleteItem

`func NewGemMappingDeleteItem(gemMappingId int32, gemPortId int32, ) *GemMappingDeleteItem`

NewGemMappingDeleteItem instantiates a new GemMappingDeleteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemMappingDeleteItemWithDefaults

`func NewGemMappingDeleteItemWithDefaults() *GemMappingDeleteItem`

NewGemMappingDeleteItemWithDefaults instantiates a new GemMappingDeleteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGemMappingId

`func (o *GemMappingDeleteItem) GetGemMappingId() int32`

GetGemMappingId returns the GemMappingId field if non-nil, zero value otherwise.

### GetGemMappingIdOk

`func (o *GemMappingDeleteItem) GetGemMappingIdOk() (*int32, bool)`

GetGemMappingIdOk returns a tuple with the GemMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemMappingId

`func (o *GemMappingDeleteItem) SetGemMappingId(v int32)`

SetGemMappingId sets GemMappingId field to given value.


### GetGemPortId

`func (o *GemMappingDeleteItem) GetGemPortId() int32`

GetGemPortId returns the GemPortId field if non-nil, zero value otherwise.

### GetGemPortIdOk

`func (o *GemMappingDeleteItem) GetGemPortIdOk() (*int32, bool)`

GetGemPortIdOk returns a tuple with the GemPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortId

`func (o *GemMappingDeleteItem) SetGemPortId(v int32)`

SetGemPortId sets GemPortId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


