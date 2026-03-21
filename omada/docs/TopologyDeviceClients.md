# TopologyDeviceClients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGroup** | Pointer to [**[]TopologyClientNode**](TopologyClientNode.md) | Client Group. | [optional] 
**DevMac** | Pointer to **string** | Device Mac. | [optional] 

## Methods

### NewTopologyDeviceClients

`func NewTopologyDeviceClients() *TopologyDeviceClients`

NewTopologyDeviceClients instantiates a new TopologyDeviceClients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyDeviceClientsWithDefaults

`func NewTopologyDeviceClientsWithDefaults() *TopologyDeviceClients`

NewTopologyDeviceClientsWithDefaults instantiates a new TopologyDeviceClients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientGroup

`func (o *TopologyDeviceClients) GetClientGroup() []TopologyClientNode`

GetClientGroup returns the ClientGroup field if non-nil, zero value otherwise.

### GetClientGroupOk

`func (o *TopologyDeviceClients) GetClientGroupOk() (*[]TopologyClientNode, bool)`

GetClientGroupOk returns a tuple with the ClientGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGroup

`func (o *TopologyDeviceClients) SetClientGroup(v []TopologyClientNode)`

SetClientGroup sets ClientGroup field to given value.

### HasClientGroup

`func (o *TopologyDeviceClients) HasClientGroup() bool`

HasClientGroup returns a boolean if a field has been set.

### GetDevMac

`func (o *TopologyDeviceClients) GetDevMac() string`

GetDevMac returns the DevMac field if non-nil, zero value otherwise.

### GetDevMacOk

`func (o *TopologyDeviceClients) GetDevMacOk() (*string, bool)`

GetDevMacOk returns a tuple with the DevMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevMac

`func (o *TopologyDeviceClients) SetDevMac(v string)`

SetDevMac sets DevMac field to given value.

### HasDevMac

`func (o *TopologyDeviceClients) HasDevMac() bool`

HasDevMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


