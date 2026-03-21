# SelectStacksWithVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackIds** | **[]string** | Stack ID List. | 
**Vlan** | Pointer to **int32** | vlan, only valid when vlanType is 0 (single vlan) | [optional] 
**VlanType** | **int32** | Network type, It should be a value as follows : 0:single vlan 1:multi vlan | 
**Vlans** | Pointer to **string** | vlans, only valid when vlanType is 1 (multi vlan). VLAN format: 200, 1-100. | [optional] 

## Methods

### NewSelectStacksWithVlanVO

`func NewSelectStacksWithVlanVO(stackIds []string, vlanType int32, ) *SelectStacksWithVlanVO`

NewSelectStacksWithVlanVO instantiates a new SelectStacksWithVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectStacksWithVlanVOWithDefaults

`func NewSelectStacksWithVlanVOWithDefaults() *SelectStacksWithVlanVO`

NewSelectStacksWithVlanVOWithDefaults instantiates a new SelectStacksWithVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackIds

`func (o *SelectStacksWithVlanVO) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *SelectStacksWithVlanVO) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *SelectStacksWithVlanVO) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.


### GetVlan

`func (o *SelectStacksWithVlanVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *SelectStacksWithVlanVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *SelectStacksWithVlanVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *SelectStacksWithVlanVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *SelectStacksWithVlanVO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *SelectStacksWithVlanVO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *SelectStacksWithVlanVO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.


### GetVlans

`func (o *SelectStacksWithVlanVO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *SelectStacksWithVlanVO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *SelectStacksWithVlanVO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *SelectStacksWithVlanVO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


