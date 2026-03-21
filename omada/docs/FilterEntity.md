# FilterEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of filter | [optional] 
**FilterId** | Pointer to **int32** | Filter ID | [optional] 
**FilterName** | Pointer to **string** | Filter name | [optional] 
**Rules** | Pointer to [**[]RuleEntity**](RuleEntity.md) | Rule list | [optional] 

## Methods

### NewFilterEntity

`func NewFilterEntity() *FilterEntity`

NewFilterEntity instantiates a new FilterEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterEntityWithDefaults

`func NewFilterEntityWithDefaults() *FilterEntity`

NewFilterEntityWithDefaults instantiates a new FilterEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FilterEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterEntity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterEntity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterId

`func (o *FilterEntity) GetFilterId() int32`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *FilterEntity) GetFilterIdOk() (*int32, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *FilterEntity) SetFilterId(v int32)`

SetFilterId sets FilterId field to given value.

### HasFilterId

`func (o *FilterEntity) HasFilterId() bool`

HasFilterId returns a boolean if a field has been set.

### GetFilterName

`func (o *FilterEntity) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *FilterEntity) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *FilterEntity) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.

### HasFilterName

`func (o *FilterEntity) HasFilterName() bool`

HasFilterName returns a boolean if a field has been set.

### GetRules

`func (o *FilterEntity) GetRules() []RuleEntity`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FilterEntity) GetRulesOk() (*[]RuleEntity, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FilterEntity) SetRules(v []RuleEntity)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FilterEntity) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


