# InstancesVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **int32** | cost | [optional] 
**CostMode** | Pointer to **int32** | costMode, 0: auto / 1: custom | [optional] 
**Id** | Pointer to **int32** | mstp instanceId | [optional] 
**Priority** | Pointer to **int32** | priority | [optional] 
**Stp** | Pointer to **int32** | stp, 0: MSTP / 1: RPVST | [optional] 
**Vlan** | Pointer to **string** | rpvst vlanId | [optional] 

## Methods

### NewInstancesVO

`func NewInstancesVO() *InstancesVO`

NewInstancesVO instantiates a new InstancesVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesVOWithDefaults

`func NewInstancesVOWithDefaults() *InstancesVO`

NewInstancesVOWithDefaults instantiates a new InstancesVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *InstancesVO) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *InstancesVO) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *InstancesVO) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *InstancesVO) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCostMode

`func (o *InstancesVO) GetCostMode() int32`

GetCostMode returns the CostMode field if non-nil, zero value otherwise.

### GetCostModeOk

`func (o *InstancesVO) GetCostModeOk() (*int32, bool)`

GetCostModeOk returns a tuple with the CostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostMode

`func (o *InstancesVO) SetCostMode(v int32)`

SetCostMode sets CostMode field to given value.

### HasCostMode

`func (o *InstancesVO) HasCostMode() bool`

HasCostMode returns a boolean if a field has been set.

### GetId

`func (o *InstancesVO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesVO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesVO) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InstancesVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *InstancesVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InstancesVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InstancesVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InstancesVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStp

`func (o *InstancesVO) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *InstancesVO) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *InstancesVO) SetStp(v int32)`

SetStp sets Stp field to given value.

### HasStp

`func (o *InstancesVO) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetVlan

`func (o *InstancesVO) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InstancesVO) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InstancesVO) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InstancesVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


