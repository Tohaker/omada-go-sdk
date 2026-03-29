# TopologyClientWirelessUpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Channel. | [optional] 
**MultiLink** | Pointer to [**[]TopologyClientWirelessLinkInfo**](TopologyClientWirelessLinkInfo.md) | Client link information while client connects to multiple frequency bands. | [optional] 
**Radio** | Pointer to **int32** | Radio, it should be a value as follows: 0:2G, 1:5G, 2:5G2, 3:6G. | [optional] 
**Ssid** | Pointer to **string** | Ssid. | [optional] 
**Support5g2** | Pointer to **bool** | Whether the device supports the 5G2 frequency band. | [optional] 

## Methods

### NewTopologyClientWirelessUpInfo

`func NewTopologyClientWirelessUpInfo() *TopologyClientWirelessUpInfo`

NewTopologyClientWirelessUpInfo instantiates a new TopologyClientWirelessUpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientWirelessUpInfoWithDefaults

`func NewTopologyClientWirelessUpInfoWithDefaults() *TopologyClientWirelessUpInfo`

NewTopologyClientWirelessUpInfoWithDefaults instantiates a new TopologyClientWirelessUpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *TopologyClientWirelessUpInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TopologyClientWirelessUpInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TopologyClientWirelessUpInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TopologyClientWirelessUpInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetMultiLink

`func (o *TopologyClientWirelessUpInfo) GetMultiLink() []TopologyClientWirelessLinkInfo`

GetMultiLink returns the MultiLink field if non-nil, zero value otherwise.

### GetMultiLinkOk

`func (o *TopologyClientWirelessUpInfo) GetMultiLinkOk() (*[]TopologyClientWirelessLinkInfo, bool)`

GetMultiLinkOk returns a tuple with the MultiLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiLink

`func (o *TopologyClientWirelessUpInfo) SetMultiLink(v []TopologyClientWirelessLinkInfo)`

SetMultiLink sets MultiLink field to given value.

### HasMultiLink

`func (o *TopologyClientWirelessUpInfo) HasMultiLink() bool`

HasMultiLink returns a boolean if a field has been set.

### GetRadio

`func (o *TopologyClientWirelessUpInfo) GetRadio() int32`

GetRadio returns the Radio field if non-nil, zero value otherwise.

### GetRadioOk

`func (o *TopologyClientWirelessUpInfo) GetRadioOk() (*int32, bool)`

GetRadioOk returns a tuple with the Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio

`func (o *TopologyClientWirelessUpInfo) SetRadio(v int32)`

SetRadio sets Radio field to given value.

### HasRadio

`func (o *TopologyClientWirelessUpInfo) HasRadio() bool`

HasRadio returns a boolean if a field has been set.

### GetSsid

`func (o *TopologyClientWirelessUpInfo) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *TopologyClientWirelessUpInfo) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *TopologyClientWirelessUpInfo) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *TopologyClientWirelessUpInfo) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSupport5g2

`func (o *TopologyClientWirelessUpInfo) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *TopologyClientWirelessUpInfo) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *TopologyClientWirelessUpInfo) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *TopologyClientWirelessUpInfo) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


