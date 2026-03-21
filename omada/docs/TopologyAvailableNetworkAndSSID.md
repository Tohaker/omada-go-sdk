# TopologyAvailableNetworkAndSSID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SplitNetworks** | Pointer to [**[]TopologySplitWiredNetwork**](TopologySplitWiredNetwork.md) | Networks split by vlan. | [optional] 
**Ssids** | Pointer to [**[]TopologySSID**](TopologySSID.md) | SSID List. | [optional] 

## Methods

### NewTopologyAvailableNetworkAndSSID

`func NewTopologyAvailableNetworkAndSSID() *TopologyAvailableNetworkAndSSID`

NewTopologyAvailableNetworkAndSSID instantiates a new TopologyAvailableNetworkAndSSID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyAvailableNetworkAndSSIDWithDefaults

`func NewTopologyAvailableNetworkAndSSIDWithDefaults() *TopologyAvailableNetworkAndSSID`

NewTopologyAvailableNetworkAndSSIDWithDefaults instantiates a new TopologyAvailableNetworkAndSSID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSplitNetworks

`func (o *TopologyAvailableNetworkAndSSID) GetSplitNetworks() []TopologySplitWiredNetwork`

GetSplitNetworks returns the SplitNetworks field if non-nil, zero value otherwise.

### GetSplitNetworksOk

`func (o *TopologyAvailableNetworkAndSSID) GetSplitNetworksOk() (*[]TopologySplitWiredNetwork, bool)`

GetSplitNetworksOk returns a tuple with the SplitNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitNetworks

`func (o *TopologyAvailableNetworkAndSSID) SetSplitNetworks(v []TopologySplitWiredNetwork)`

SetSplitNetworks sets SplitNetworks field to given value.

### HasSplitNetworks

`func (o *TopologyAvailableNetworkAndSSID) HasSplitNetworks() bool`

HasSplitNetworks returns a boolean if a field has been set.

### GetSsids

`func (o *TopologyAvailableNetworkAndSSID) GetSsids() []TopologySSID`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *TopologyAvailableNetworkAndSSID) GetSsidsOk() (*[]TopologySSID, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *TopologyAvailableNetworkAndSSID) SetSsids(v []TopologySSID)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *TopologyAvailableNetworkAndSSID) HasSsids() bool`

HasSsids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


