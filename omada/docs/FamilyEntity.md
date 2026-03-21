# FamilyEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of family | [optional] 
**FamilyId** | Pointer to **int32** | Family ID | [optional] 
**FamilyName** | Pointer to **string** | Family name | [optional] 

## Methods

### NewFamilyEntity

`func NewFamilyEntity() *FamilyEntity`

NewFamilyEntity instantiates a new FamilyEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFamilyEntityWithDefaults

`func NewFamilyEntityWithDefaults() *FamilyEntity`

NewFamilyEntityWithDefaults instantiates a new FamilyEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FamilyEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FamilyEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FamilyEntity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FamilyEntity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFamilyId

`func (o *FamilyEntity) GetFamilyId() int32`

GetFamilyId returns the FamilyId field if non-nil, zero value otherwise.

### GetFamilyIdOk

`func (o *FamilyEntity) GetFamilyIdOk() (*int32, bool)`

GetFamilyIdOk returns a tuple with the FamilyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyId

`func (o *FamilyEntity) SetFamilyId(v int32)`

SetFamilyId sets FamilyId field to given value.

### HasFamilyId

`func (o *FamilyEntity) HasFamilyId() bool`

HasFamilyId returns a boolean if a field has been set.

### GetFamilyName

`func (o *FamilyEntity) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *FamilyEntity) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *FamilyEntity) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *FamilyEntity) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


