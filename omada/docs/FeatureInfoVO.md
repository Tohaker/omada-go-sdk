# FeatureInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changeable** | Pointer to **bool** | If the current gateway does not support this feature, then the [changeable] is false. | [optional] 
**Feature** | Pointer to **string** | Gateway feature name. | [optional] 
**FeatureState** | Pointer to **int32** | Gateway feature status: 0: The current gateway supports this feature. 1: The current gateway does not support this feature, but it is pre-configured. 2: The current gateway does not support this feature. | [optional] 

## Methods

### NewFeatureInfoVO

`func NewFeatureInfoVO() *FeatureInfoVO`

NewFeatureInfoVO instantiates a new FeatureInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureInfoVOWithDefaults

`func NewFeatureInfoVOWithDefaults() *FeatureInfoVO`

NewFeatureInfoVOWithDefaults instantiates a new FeatureInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeable

`func (o *FeatureInfoVO) GetChangeable() bool`

GetChangeable returns the Changeable field if non-nil, zero value otherwise.

### GetChangeableOk

`func (o *FeatureInfoVO) GetChangeableOk() (*bool, bool)`

GetChangeableOk returns a tuple with the Changeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeable

`func (o *FeatureInfoVO) SetChangeable(v bool)`

SetChangeable sets Changeable field to given value.

### HasChangeable

`func (o *FeatureInfoVO) HasChangeable() bool`

HasChangeable returns a boolean if a field has been set.

### GetFeature

`func (o *FeatureInfoVO) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *FeatureInfoVO) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *FeatureInfoVO) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *FeatureInfoVO) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetFeatureState

`func (o *FeatureInfoVO) GetFeatureState() int32`

GetFeatureState returns the FeatureState field if non-nil, zero value otherwise.

### GetFeatureStateOk

`func (o *FeatureInfoVO) GetFeatureStateOk() (*int32, bool)`

GetFeatureStateOk returns a tuple with the FeatureState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureState

`func (o *FeatureInfoVO) SetFeatureState(v int32)`

SetFeatureState sets FeatureState field to given value.

### HasFeatureState

`func (o *FeatureInfoVO) HasFeatureState() bool`

HasFeatureState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


