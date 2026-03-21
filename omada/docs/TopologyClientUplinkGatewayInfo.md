# TopologyClientUplinkGatewayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Channel, only for wireless router. | [optional] 
**Duplex** | Pointer to **int32** | Duplex, it should be a value as follows: 1:Half Duplex. | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed, it should be a value as follows: 1:10Mbps, 2:100Mbps, 3:1000Mbps, 4:2.5Gbps, 5:10Gbps, 6:5Gbps, 7:25Gbps. | [optional] 
**Port** | Pointer to [**TopologyClientUplinkGatewayPortInfo**](TopologyClientUplinkGatewayPortInfo.md) |  | [optional] 
**Radio** | Pointer to **int32** | Radio, it should be a value as follows: 0:2G, 1:5G, 2:5G2, 3:6G, only for wireless router | [optional] 
**Ssid** | Pointer to **string** | Ssid, only for wireless router | [optional] 
**Support5g2** | Pointer to **bool** | Whether the device supports the 5G2 frequency band, only for wireless router. | [optional] 

## Methods

### NewTopologyClientUplinkGatewayInfo

`func NewTopologyClientUplinkGatewayInfo() *TopologyClientUplinkGatewayInfo`

NewTopologyClientUplinkGatewayInfo instantiates a new TopologyClientUplinkGatewayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientUplinkGatewayInfoWithDefaults

`func NewTopologyClientUplinkGatewayInfoWithDefaults() *TopologyClientUplinkGatewayInfo`

NewTopologyClientUplinkGatewayInfoWithDefaults instantiates a new TopologyClientUplinkGatewayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *TopologyClientUplinkGatewayInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TopologyClientUplinkGatewayInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TopologyClientUplinkGatewayInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TopologyClientUplinkGatewayInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDuplex

`func (o *TopologyClientUplinkGatewayInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *TopologyClientUplinkGatewayInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *TopologyClientUplinkGatewayInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *TopologyClientUplinkGatewayInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *TopologyClientUplinkGatewayInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *TopologyClientUplinkGatewayInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *TopologyClientUplinkGatewayInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *TopologyClientUplinkGatewayInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetPort

`func (o *TopologyClientUplinkGatewayInfo) GetPort() TopologyClientUplinkGatewayPortInfo`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TopologyClientUplinkGatewayInfo) GetPortOk() (*TopologyClientUplinkGatewayPortInfo, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TopologyClientUplinkGatewayInfo) SetPort(v TopologyClientUplinkGatewayPortInfo)`

SetPort sets Port field to given value.

### HasPort

`func (o *TopologyClientUplinkGatewayInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRadio

`func (o *TopologyClientUplinkGatewayInfo) GetRadio() int32`

GetRadio returns the Radio field if non-nil, zero value otherwise.

### GetRadioOk

`func (o *TopologyClientUplinkGatewayInfo) GetRadioOk() (*int32, bool)`

GetRadioOk returns a tuple with the Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio

`func (o *TopologyClientUplinkGatewayInfo) SetRadio(v int32)`

SetRadio sets Radio field to given value.

### HasRadio

`func (o *TopologyClientUplinkGatewayInfo) HasRadio() bool`

HasRadio returns a boolean if a field has been set.

### GetSsid

`func (o *TopologyClientUplinkGatewayInfo) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *TopologyClientUplinkGatewayInfo) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *TopologyClientUplinkGatewayInfo) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *TopologyClientUplinkGatewayInfo) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSupport5g2

`func (o *TopologyClientUplinkGatewayInfo) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *TopologyClientUplinkGatewayInfo) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *TopologyClientUplinkGatewayInfo) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *TopologyClientUplinkGatewayInfo) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


