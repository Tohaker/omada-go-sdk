# TopologyClientUplinkApInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Channel, only for wireless connection. | [optional] 
**Duplex** | Pointer to **int32** | Duplex, it should be a value as follows: 1:Half Duplex, 2:Full Duplex, only for wired connection. | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed, it should be a value as follows: 1:10Mbps, 2:100Mbps, 3:1000Mbps, 4:2.5Gbps, 5:10Gbps, only for wired connection. | [optional] 
**Radio** | Pointer to **int32** | Radio, it should be a value as follows: 0:2G, 1:5G, 2:5G2, 3:6G, only for wireless connection. | [optional] 
**Ssid** | Pointer to **string** | Ssid, only for wireless connection. | [optional] 
**Support5g2** | Pointer to **bool** | Whether the device supports the 5G2 frequency band. | [optional] 

## Methods

### NewTopologyClientUplinkApInfo

`func NewTopologyClientUplinkApInfo() *TopologyClientUplinkApInfo`

NewTopologyClientUplinkApInfo instantiates a new TopologyClientUplinkApInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientUplinkApInfoWithDefaults

`func NewTopologyClientUplinkApInfoWithDefaults() *TopologyClientUplinkApInfo`

NewTopologyClientUplinkApInfoWithDefaults instantiates a new TopologyClientUplinkApInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *TopologyClientUplinkApInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TopologyClientUplinkApInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TopologyClientUplinkApInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TopologyClientUplinkApInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDuplex

`func (o *TopologyClientUplinkApInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *TopologyClientUplinkApInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *TopologyClientUplinkApInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *TopologyClientUplinkApInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *TopologyClientUplinkApInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *TopologyClientUplinkApInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *TopologyClientUplinkApInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *TopologyClientUplinkApInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetRadio

`func (o *TopologyClientUplinkApInfo) GetRadio() int32`

GetRadio returns the Radio field if non-nil, zero value otherwise.

### GetRadioOk

`func (o *TopologyClientUplinkApInfo) GetRadioOk() (*int32, bool)`

GetRadioOk returns a tuple with the Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio

`func (o *TopologyClientUplinkApInfo) SetRadio(v int32)`

SetRadio sets Radio field to given value.

### HasRadio

`func (o *TopologyClientUplinkApInfo) HasRadio() bool`

HasRadio returns a boolean if a field has been set.

### GetSsid

`func (o *TopologyClientUplinkApInfo) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *TopologyClientUplinkApInfo) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *TopologyClientUplinkApInfo) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *TopologyClientUplinkApInfo) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSupport5g2

`func (o *TopologyClientUplinkApInfo) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *TopologyClientUplinkApInfo) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *TopologyClientUplinkApInfo) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *TopologyClientUplinkApInfo) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


