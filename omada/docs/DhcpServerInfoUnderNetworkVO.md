# DhcpServerInfoUnderNetworkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]NetworkWithServerVO**](NetworkWithServerVO.md) | Networks With Dhcp Servers | [optional] 

## Methods

### NewDhcpServerInfoUnderNetworkVO

`func NewDhcpServerInfoUnderNetworkVO() *DhcpServerInfoUnderNetworkVO`

NewDhcpServerInfoUnderNetworkVO instantiates a new DhcpServerInfoUnderNetworkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpServerInfoUnderNetworkVOWithDefaults

`func NewDhcpServerInfoUnderNetworkVOWithDefaults() *DhcpServerInfoUnderNetworkVO`

NewDhcpServerInfoUnderNetworkVOWithDefaults instantiates a new DhcpServerInfoUnderNetworkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *DhcpServerInfoUnderNetworkVO) GetNetworks() []NetworkWithServerVO`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *DhcpServerInfoUnderNetworkVO) GetNetworksOk() (*[]NetworkWithServerVO, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *DhcpServerInfoUnderNetworkVO) SetNetworks(v []NetworkWithServerVO)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *DhcpServerInfoUnderNetworkVO) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


