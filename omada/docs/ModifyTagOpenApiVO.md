# ModifyTagOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tag name should contain 1 to 128 ASCII characters. | 
**TagId** | **string** | Tag ID | 

## Methods

### NewModifyTagOpenApiVO

`func NewModifyTagOpenApiVO(name string, tagId string, ) *ModifyTagOpenApiVO`

NewModifyTagOpenApiVO instantiates a new ModifyTagOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyTagOpenApiVOWithDefaults

`func NewModifyTagOpenApiVOWithDefaults() *ModifyTagOpenApiVO`

NewModifyTagOpenApiVOWithDefaults instantiates a new ModifyTagOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModifyTagOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyTagOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyTagOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetTagId

`func (o *ModifyTagOpenApiVO) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *ModifyTagOpenApiVO) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *ModifyTagOpenApiVO) SetTagId(v string)`

SetTagId sets TagId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


