# PortTagOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Port label name should contain 1 to 128 ASCII characters. | 
**Resource** | Pointer to **int32** | Data Source. Resource should be a value as follows: 0: new created; 1: from template; 2: override | [optional] 
**TagId** | **string** | Port label ID | 

## Methods

### NewPortTagOpenApiVO

`func NewPortTagOpenApiVO(name string, tagId string, ) *PortTagOpenApiVO`

NewPortTagOpenApiVO instantiates a new PortTagOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortTagOpenApiVOWithDefaults

`func NewPortTagOpenApiVOWithDefaults() *PortTagOpenApiVO`

NewPortTagOpenApiVOWithDefaults instantiates a new PortTagOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PortTagOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortTagOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortTagOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetResource

`func (o *PortTagOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PortTagOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PortTagOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PortTagOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetTagId

`func (o *PortTagOpenApiVO) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *PortTagOpenApiVO) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *PortTagOpenApiVO) SetTagId(v string)`

SetTagId sets TagId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


