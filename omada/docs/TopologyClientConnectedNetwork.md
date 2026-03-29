# TopologyClientConnectedNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Network name. | [optional] 
**NetworkId** | Pointer to **string** | Network ID. | [optional] 
**VlanId** | Pointer to **int32** | Vlan Id, for vlan or single vlan interface, ranges from 1 to 4090. | [optional] 
**VlanIds** | Pointer to **[]int32** | Vlan Ids, for multiple vlan interface, for network . | [optional] 

## Methods

### NewTopologyClientConnectedNetwork

`func NewTopologyClientConnectedNetwork() *TopologyClientConnectedNetwork`

NewTopologyClientConnectedNetwork instantiates a new TopologyClientConnectedNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientConnectedNetworkWithDefaults

`func NewTopologyClientConnectedNetworkWithDefaults() *TopologyClientConnectedNetwork`

NewTopologyClientConnectedNetworkWithDefaults instantiates a new TopologyClientConnectedNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TopologyClientConnectedNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyClientConnectedNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyClientConnectedNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyClientConnectedNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkId

`func (o *TopologyClientConnectedNetwork) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *TopologyClientConnectedNetwork) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *TopologyClientConnectedNetwork) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *TopologyClientConnectedNetwork) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetVlanId

`func (o *TopologyClientConnectedNetwork) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *TopologyClientConnectedNetwork) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *TopologyClientConnectedNetwork) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *TopologyClientConnectedNetwork) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanIds

`func (o *TopologyClientConnectedNetwork) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *TopologyClientConnectedNetwork) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *TopologyClientConnectedNetwork) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *TopologyClientConnectedNetwork) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


