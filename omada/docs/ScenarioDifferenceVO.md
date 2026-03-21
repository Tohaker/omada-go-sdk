# ScenarioDifferenceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** | Identify whether the current scene is the default scenario. | [optional] 
**Name** | Pointer to **string** | Scenario name should contain 1 to 128 visible ASCII characters. | [optional] 

## Methods

### NewScenarioDifferenceVO

`func NewScenarioDifferenceVO() *ScenarioDifferenceVO`

NewScenarioDifferenceVO instantiates a new ScenarioDifferenceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioDifferenceVOWithDefaults

`func NewScenarioDifferenceVOWithDefaults() *ScenarioDifferenceVO`

NewScenarioDifferenceVOWithDefaults instantiates a new ScenarioDifferenceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *ScenarioDifferenceVO) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ScenarioDifferenceVO) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ScenarioDifferenceVO) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ScenarioDifferenceVO) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *ScenarioDifferenceVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScenarioDifferenceVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScenarioDifferenceVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScenarioDifferenceVO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


