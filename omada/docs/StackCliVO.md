# StackCliVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterMac** | Pointer to **string** | Master mac | [optional] 
**StackId** | Pointer to **string** | Stack ID | [optional] 
**StackName** | Pointer to **string** | Stack name | [optional] 
**VariableMap** | Pointer to **map[string]string** | The values of different user-defined variables. | [optional] 

## Methods

### NewStackCliVO

`func NewStackCliVO() *StackCliVO`

NewStackCliVO instantiates a new StackCliVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackCliVOWithDefaults

`func NewStackCliVOWithDefaults() *StackCliVO`

NewStackCliVOWithDefaults instantiates a new StackCliVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterMac

`func (o *StackCliVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *StackCliVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *StackCliVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *StackCliVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetStackId

`func (o *StackCliVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *StackCliVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *StackCliVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *StackCliVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *StackCliVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *StackCliVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *StackCliVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *StackCliVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetVariableMap

`func (o *StackCliVO) GetVariableMap() map[string]string`

GetVariableMap returns the VariableMap field if non-nil, zero value otherwise.

### GetVariableMapOk

`func (o *StackCliVO) GetVariableMapOk() (*map[string]string, bool)`

GetVariableMapOk returns a tuple with the VariableMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMap

`func (o *StackCliVO) SetVariableMap(v map[string]string)`

SetVariableMap sets VariableMap field to given value.

### HasVariableMap

`func (o *StackCliVO) HasVariableMap() bool`

HasVariableMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


