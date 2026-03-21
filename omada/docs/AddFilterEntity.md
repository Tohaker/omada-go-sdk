# AddFilterEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of filter. It should be 1 - 128 characters | 
**FilterName** | **string** | Filter name. It should be 1 - 128 characters | 
**Rules** | **[]int32** | Rule ID list can be obtained from &#39;Get rule list&#39; interface. | 

## Methods

### NewAddFilterEntity

`func NewAddFilterEntity(description string, filterName string, rules []int32, ) *AddFilterEntity`

NewAddFilterEntity instantiates a new AddFilterEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFilterEntityWithDefaults

`func NewAddFilterEntityWithDefaults() *AddFilterEntity`

NewAddFilterEntityWithDefaults instantiates a new AddFilterEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AddFilterEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFilterEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFilterEntity) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFilterName

`func (o *AddFilterEntity) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *AddFilterEntity) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *AddFilterEntity) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.


### GetRules

`func (o *AddFilterEntity) GetRules() []int32`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AddFilterEntity) GetRulesOk() (*[]int32, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AddFilterEntity) SetRules(v []int32)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


