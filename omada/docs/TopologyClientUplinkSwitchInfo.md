# TopologyClientUplinkSwitchInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Duplex, it should be a value as follows: 1:Half Duplex. | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed, it should be a value as follows: 1:10Mbps, 2:100Mbps, 3:1000Mbps, 4:2.5Gbps, 5:10Gbps, 6:5Gbps, 7:25Gbps. | [optional] 
**Port** | Pointer to [**TopologyClientUplinkSwitchPortInfo**](TopologyClientUplinkSwitchPortInfo.md) |  | [optional] 

## Methods

### NewTopologyClientUplinkSwitchInfo

`func NewTopologyClientUplinkSwitchInfo() *TopologyClientUplinkSwitchInfo`

NewTopologyClientUplinkSwitchInfo instantiates a new TopologyClientUplinkSwitchInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientUplinkSwitchInfoWithDefaults

`func NewTopologyClientUplinkSwitchInfoWithDefaults() *TopologyClientUplinkSwitchInfo`

NewTopologyClientUplinkSwitchInfoWithDefaults instantiates a new TopologyClientUplinkSwitchInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *TopologyClientUplinkSwitchInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *TopologyClientUplinkSwitchInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *TopologyClientUplinkSwitchInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *TopologyClientUplinkSwitchInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *TopologyClientUplinkSwitchInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *TopologyClientUplinkSwitchInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *TopologyClientUplinkSwitchInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *TopologyClientUplinkSwitchInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetPort

`func (o *TopologyClientUplinkSwitchInfo) GetPort() TopologyClientUplinkSwitchPortInfo`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TopologyClientUplinkSwitchInfo) GetPortOk() (*TopologyClientUplinkSwitchPortInfo, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TopologyClientUplinkSwitchInfo) SetPort(v TopologyClientUplinkSwitchPortInfo)`

SetPort sets Port field to given value.

### HasPort

`func (o *TopologyClientUplinkSwitchInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


