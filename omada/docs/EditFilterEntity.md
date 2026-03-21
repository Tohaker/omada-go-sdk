# EditFilterEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of filter. It should be 1 - 128 characters | [optional] 
**FilterName** | **string** | Filter name. It should be 1 - 128 characters | 
**Rules** | **[]int32** | Rule ID list can be obtained from &#39;Get rule list&#39; interface. | 

## Methods

### NewEditFilterEntity

`func NewEditFilterEntity(filterName string, rules []int32, ) *EditFilterEntity`

NewEditFilterEntity instantiates a new EditFilterEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditFilterEntityWithDefaults

`func NewEditFilterEntityWithDefaults() *EditFilterEntity`

NewEditFilterEntityWithDefaults instantiates a new EditFilterEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EditFilterEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditFilterEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditFilterEntity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditFilterEntity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterName

`func (o *EditFilterEntity) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *EditFilterEntity) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *EditFilterEntity) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.


### GetRules

`func (o *EditFilterEntity) GetRules() []int32`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *EditFilterEntity) GetRulesOk() (*[]int32, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *EditFilterEntity) SetRules(v []int32)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


