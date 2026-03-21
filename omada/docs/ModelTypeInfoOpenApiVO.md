# ModelTypeInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompoundModel** | **string** | Model complex used on the backend, you can also get this field throw: \&quot;Get the model of the specified site\&quot; | 
**ShowModel** | **string** | Model complex displayed on the front end, you can also get this field throw: \&quot;Get the model of the specified site\&quot; | 

## Methods

### NewModelTypeInfoOpenApiVO

`func NewModelTypeInfoOpenApiVO(compoundModel string, showModel string, ) *ModelTypeInfoOpenApiVO`

NewModelTypeInfoOpenApiVO instantiates a new ModelTypeInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTypeInfoOpenApiVOWithDefaults

`func NewModelTypeInfoOpenApiVOWithDefaults() *ModelTypeInfoOpenApiVO`

NewModelTypeInfoOpenApiVOWithDefaults instantiates a new ModelTypeInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompoundModel

`func (o *ModelTypeInfoOpenApiVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *ModelTypeInfoOpenApiVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *ModelTypeInfoOpenApiVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.


### GetShowModel

`func (o *ModelTypeInfoOpenApiVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *ModelTypeInfoOpenApiVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *ModelTypeInfoOpenApiVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


