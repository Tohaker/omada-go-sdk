# VirtualWanWeightOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | virtual wan name | [optional] 
**VirtualWanId** | **string** | virtual wan id | 
**Weight** | Pointer to **int32** | virtual wan weight | [optional] 

## Methods

### NewVirtualWanWeightOpenApiVO

`func NewVirtualWanWeightOpenApiVO(virtualWanId string, ) *VirtualWanWeightOpenApiVO`

NewVirtualWanWeightOpenApiVO instantiates a new VirtualWanWeightOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanWeightOpenApiVOWithDefaults

`func NewVirtualWanWeightOpenApiVOWithDefaults() *VirtualWanWeightOpenApiVO`

NewVirtualWanWeightOpenApiVOWithDefaults instantiates a new VirtualWanWeightOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualWanWeightOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualWanWeightOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualWanWeightOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualWanWeightOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVirtualWanId

`func (o *VirtualWanWeightOpenApiVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *VirtualWanWeightOpenApiVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *VirtualWanWeightOpenApiVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.


### GetWeight

`func (o *VirtualWanWeightOpenApiVO) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *VirtualWanWeightOpenApiVO) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *VirtualWanWeightOpenApiVO) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *VirtualWanWeightOpenApiVO) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


