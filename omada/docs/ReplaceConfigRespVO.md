# ReplaceConfigRespVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureList** | Pointer to **[]string** | List of incompatible features. | [optional] 
**NeedPrompt** | Pointer to **bool** | Whether to show a prompt. | [optional] 
**NewModel** | Pointer to **string** | New model. | [optional] 
**OldModel** | Pointer to **string** | Old model. | [optional] 

## Methods

### NewReplaceConfigRespVO

`func NewReplaceConfigRespVO() *ReplaceConfigRespVO`

NewReplaceConfigRespVO instantiates a new ReplaceConfigRespVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceConfigRespVOWithDefaults

`func NewReplaceConfigRespVOWithDefaults() *ReplaceConfigRespVO`

NewReplaceConfigRespVOWithDefaults instantiates a new ReplaceConfigRespVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureList

`func (o *ReplaceConfigRespVO) GetFeatureList() []string`

GetFeatureList returns the FeatureList field if non-nil, zero value otherwise.

### GetFeatureListOk

`func (o *ReplaceConfigRespVO) GetFeatureListOk() (*[]string, bool)`

GetFeatureListOk returns a tuple with the FeatureList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureList

`func (o *ReplaceConfigRespVO) SetFeatureList(v []string)`

SetFeatureList sets FeatureList field to given value.

### HasFeatureList

`func (o *ReplaceConfigRespVO) HasFeatureList() bool`

HasFeatureList returns a boolean if a field has been set.

### GetNeedPrompt

`func (o *ReplaceConfigRespVO) GetNeedPrompt() bool`

GetNeedPrompt returns the NeedPrompt field if non-nil, zero value otherwise.

### GetNeedPromptOk

`func (o *ReplaceConfigRespVO) GetNeedPromptOk() (*bool, bool)`

GetNeedPromptOk returns a tuple with the NeedPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedPrompt

`func (o *ReplaceConfigRespVO) SetNeedPrompt(v bool)`

SetNeedPrompt sets NeedPrompt field to given value.

### HasNeedPrompt

`func (o *ReplaceConfigRespVO) HasNeedPrompt() bool`

HasNeedPrompt returns a boolean if a field has been set.

### GetNewModel

`func (o *ReplaceConfigRespVO) GetNewModel() string`

GetNewModel returns the NewModel field if non-nil, zero value otherwise.

### GetNewModelOk

`func (o *ReplaceConfigRespVO) GetNewModelOk() (*string, bool)`

GetNewModelOk returns a tuple with the NewModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewModel

`func (o *ReplaceConfigRespVO) SetNewModel(v string)`

SetNewModel sets NewModel field to given value.

### HasNewModel

`func (o *ReplaceConfigRespVO) HasNewModel() bool`

HasNewModel returns a boolean if a field has been set.

### GetOldModel

`func (o *ReplaceConfigRespVO) GetOldModel() string`

GetOldModel returns the OldModel field if non-nil, zero value otherwise.

### GetOldModelOk

`func (o *ReplaceConfigRespVO) GetOldModelOk() (*string, bool)`

GetOldModelOk returns a tuple with the OldModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldModel

`func (o *ReplaceConfigRespVO) SetOldModel(v string)`

SetOldModel sets OldModel field to given value.

### HasOldModel

`func (o *ReplaceConfigRespVO) HasOldModel() bool`

HasOldModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


