# AttackDefenseSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IcmpConnEnable** | **bool** | Multi-connections ICMP flood enable of the attack defense setting. | 
**IcmpConnLimit** | Pointer to **int32** | Multi-connections ICMP flood limit should be within the range of 100–99999. | [optional] 
**IcmpSrcEnable** | **bool** | Stationary source ICMP flood enable of the attack defense setting. | 
**IcmpSrcLimit** | Pointer to **int32** | Stationary source ICMP flood limit should be within the range of 100–99999. | [optional] 
**IcmpTimestampRequestReject** | Pointer to **bool** | ICMP Timestamp Request setting of the attack defense setting. | [optional] 
**LargePingEnable** | **bool** | Block large ping of the attack defense setting. | 
**LargePingThreshold** | Pointer to **int32** | Block large ping threshold of the attack defense setting, Value is between 28 and 65535. | [optional] 
**PingDeathEnable** | **bool** | Block ping of death of the attack defense setting. | 
**PingWanEnable** | **bool** | Block ping from WAN of the attack defense setting. | 
**SpecifiedOption** | Pointer to [**SpecifiedOptionOpenApiVO**](SpecifiedOptionOpenApiVO.md) |  | [optional] 
**SpecifiedOptionEnable** | **bool** | Block packets with specified options of the attack defense setting. | 
**TcpConnEnable** | **bool** | Multi-connections TCP SYN flood enable of the attack defense setting. | 
**TcpConnLimit** | Pointer to **int32** | Multi-connections TCP SYN flood limit should be within the range of 100–99999. | [optional] 
**TcpFinNoAckEnable** | **bool** | Block TCP packets with FIN Bit set but no ACK Bit set of the attack defense setting. | 
**TcpScanEnable** | **bool** | Block TCP scan enable of the attack defense setting. | 
**TcpScanReject** | Pointer to **bool** | Block TCP scan with reject of the attack defense setting. | [optional] 
**TcpSrcEnable** | **bool** | Stationary source TCP SYN flood enable of the attack defense setting. | 
**TcpSrcLimit** | Pointer to **int32** | Stationary source TCP SYN flood limit should be within the range of 100–99999. | [optional] 
**TcpSynFinEnable** | **bool** | Block TCP packets with SYN and FIN Bits set of the attack defense setting. | 
**UdpConnEnable** | **bool** | Multi-connections UDP flood enable of the attack defense setting. | 
**UdpConnLimit** | Pointer to **int32** | Multi-connections UDP flood limit should be within the range of 100–99999. | [optional] 
**UdpSrcEnable** | **bool** | Stationary source UDP flood enable of the attack defense setting. | 
**UdpSrcLimit** | Pointer to **int32** | Stationary source UDP flood limit should be within the range of 100–99999. | [optional] 
**WinNukeAttackEnable** | **bool** | Block WinNuke attack of the attack defense setting. | 

## Methods

### NewAttackDefenseSetting

`func NewAttackDefenseSetting(icmpConnEnable bool, icmpSrcEnable bool, largePingEnable bool, pingDeathEnable bool, pingWanEnable bool, specifiedOptionEnable bool, tcpConnEnable bool, tcpFinNoAckEnable bool, tcpScanEnable bool, tcpSrcEnable bool, tcpSynFinEnable bool, udpConnEnable bool, udpSrcEnable bool, winNukeAttackEnable bool, ) *AttackDefenseSetting`

NewAttackDefenseSetting instantiates a new AttackDefenseSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackDefenseSettingWithDefaults

`func NewAttackDefenseSettingWithDefaults() *AttackDefenseSetting`

NewAttackDefenseSettingWithDefaults instantiates a new AttackDefenseSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcmpConnEnable

`func (o *AttackDefenseSetting) GetIcmpConnEnable() bool`

GetIcmpConnEnable returns the IcmpConnEnable field if non-nil, zero value otherwise.

### GetIcmpConnEnableOk

`func (o *AttackDefenseSetting) GetIcmpConnEnableOk() (*bool, bool)`

GetIcmpConnEnableOk returns a tuple with the IcmpConnEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpConnEnable

`func (o *AttackDefenseSetting) SetIcmpConnEnable(v bool)`

SetIcmpConnEnable sets IcmpConnEnable field to given value.


### GetIcmpConnLimit

`func (o *AttackDefenseSetting) GetIcmpConnLimit() int32`

GetIcmpConnLimit returns the IcmpConnLimit field if non-nil, zero value otherwise.

### GetIcmpConnLimitOk

`func (o *AttackDefenseSetting) GetIcmpConnLimitOk() (*int32, bool)`

GetIcmpConnLimitOk returns a tuple with the IcmpConnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpConnLimit

`func (o *AttackDefenseSetting) SetIcmpConnLimit(v int32)`

SetIcmpConnLimit sets IcmpConnLimit field to given value.

### HasIcmpConnLimit

`func (o *AttackDefenseSetting) HasIcmpConnLimit() bool`

HasIcmpConnLimit returns a boolean if a field has been set.

### GetIcmpSrcEnable

`func (o *AttackDefenseSetting) GetIcmpSrcEnable() bool`

GetIcmpSrcEnable returns the IcmpSrcEnable field if non-nil, zero value otherwise.

### GetIcmpSrcEnableOk

`func (o *AttackDefenseSetting) GetIcmpSrcEnableOk() (*bool, bool)`

GetIcmpSrcEnableOk returns a tuple with the IcmpSrcEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpSrcEnable

`func (o *AttackDefenseSetting) SetIcmpSrcEnable(v bool)`

SetIcmpSrcEnable sets IcmpSrcEnable field to given value.


### GetIcmpSrcLimit

`func (o *AttackDefenseSetting) GetIcmpSrcLimit() int32`

GetIcmpSrcLimit returns the IcmpSrcLimit field if non-nil, zero value otherwise.

### GetIcmpSrcLimitOk

`func (o *AttackDefenseSetting) GetIcmpSrcLimitOk() (*int32, bool)`

GetIcmpSrcLimitOk returns a tuple with the IcmpSrcLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpSrcLimit

`func (o *AttackDefenseSetting) SetIcmpSrcLimit(v int32)`

SetIcmpSrcLimit sets IcmpSrcLimit field to given value.

### HasIcmpSrcLimit

`func (o *AttackDefenseSetting) HasIcmpSrcLimit() bool`

HasIcmpSrcLimit returns a boolean if a field has been set.

### GetIcmpTimestampRequestReject

`func (o *AttackDefenseSetting) GetIcmpTimestampRequestReject() bool`

GetIcmpTimestampRequestReject returns the IcmpTimestampRequestReject field if non-nil, zero value otherwise.

### GetIcmpTimestampRequestRejectOk

`func (o *AttackDefenseSetting) GetIcmpTimestampRequestRejectOk() (*bool, bool)`

GetIcmpTimestampRequestRejectOk returns a tuple with the IcmpTimestampRequestReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpTimestampRequestReject

`func (o *AttackDefenseSetting) SetIcmpTimestampRequestReject(v bool)`

SetIcmpTimestampRequestReject sets IcmpTimestampRequestReject field to given value.

### HasIcmpTimestampRequestReject

`func (o *AttackDefenseSetting) HasIcmpTimestampRequestReject() bool`

HasIcmpTimestampRequestReject returns a boolean if a field has been set.

### GetLargePingEnable

`func (o *AttackDefenseSetting) GetLargePingEnable() bool`

GetLargePingEnable returns the LargePingEnable field if non-nil, zero value otherwise.

### GetLargePingEnableOk

`func (o *AttackDefenseSetting) GetLargePingEnableOk() (*bool, bool)`

GetLargePingEnableOk returns a tuple with the LargePingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargePingEnable

`func (o *AttackDefenseSetting) SetLargePingEnable(v bool)`

SetLargePingEnable sets LargePingEnable field to given value.


### GetLargePingThreshold

`func (o *AttackDefenseSetting) GetLargePingThreshold() int32`

GetLargePingThreshold returns the LargePingThreshold field if non-nil, zero value otherwise.

### GetLargePingThresholdOk

`func (o *AttackDefenseSetting) GetLargePingThresholdOk() (*int32, bool)`

GetLargePingThresholdOk returns a tuple with the LargePingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargePingThreshold

`func (o *AttackDefenseSetting) SetLargePingThreshold(v int32)`

SetLargePingThreshold sets LargePingThreshold field to given value.

### HasLargePingThreshold

`func (o *AttackDefenseSetting) HasLargePingThreshold() bool`

HasLargePingThreshold returns a boolean if a field has been set.

### GetPingDeathEnable

`func (o *AttackDefenseSetting) GetPingDeathEnable() bool`

GetPingDeathEnable returns the PingDeathEnable field if non-nil, zero value otherwise.

### GetPingDeathEnableOk

`func (o *AttackDefenseSetting) GetPingDeathEnableOk() (*bool, bool)`

GetPingDeathEnableOk returns a tuple with the PingDeathEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingDeathEnable

`func (o *AttackDefenseSetting) SetPingDeathEnable(v bool)`

SetPingDeathEnable sets PingDeathEnable field to given value.


### GetPingWanEnable

`func (o *AttackDefenseSetting) GetPingWanEnable() bool`

GetPingWanEnable returns the PingWanEnable field if non-nil, zero value otherwise.

### GetPingWanEnableOk

`func (o *AttackDefenseSetting) GetPingWanEnableOk() (*bool, bool)`

GetPingWanEnableOk returns a tuple with the PingWanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingWanEnable

`func (o *AttackDefenseSetting) SetPingWanEnable(v bool)`

SetPingWanEnable sets PingWanEnable field to given value.


### GetSpecifiedOption

`func (o *AttackDefenseSetting) GetSpecifiedOption() SpecifiedOptionOpenApiVO`

GetSpecifiedOption returns the SpecifiedOption field if non-nil, zero value otherwise.

### GetSpecifiedOptionOk

`func (o *AttackDefenseSetting) GetSpecifiedOptionOk() (*SpecifiedOptionOpenApiVO, bool)`

GetSpecifiedOptionOk returns a tuple with the SpecifiedOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedOption

`func (o *AttackDefenseSetting) SetSpecifiedOption(v SpecifiedOptionOpenApiVO)`

SetSpecifiedOption sets SpecifiedOption field to given value.

### HasSpecifiedOption

`func (o *AttackDefenseSetting) HasSpecifiedOption() bool`

HasSpecifiedOption returns a boolean if a field has been set.

### GetSpecifiedOptionEnable

`func (o *AttackDefenseSetting) GetSpecifiedOptionEnable() bool`

GetSpecifiedOptionEnable returns the SpecifiedOptionEnable field if non-nil, zero value otherwise.

### GetSpecifiedOptionEnableOk

`func (o *AttackDefenseSetting) GetSpecifiedOptionEnableOk() (*bool, bool)`

GetSpecifiedOptionEnableOk returns a tuple with the SpecifiedOptionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedOptionEnable

`func (o *AttackDefenseSetting) SetSpecifiedOptionEnable(v bool)`

SetSpecifiedOptionEnable sets SpecifiedOptionEnable field to given value.


### GetTcpConnEnable

`func (o *AttackDefenseSetting) GetTcpConnEnable() bool`

GetTcpConnEnable returns the TcpConnEnable field if non-nil, zero value otherwise.

### GetTcpConnEnableOk

`func (o *AttackDefenseSetting) GetTcpConnEnableOk() (*bool, bool)`

GetTcpConnEnableOk returns a tuple with the TcpConnEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpConnEnable

`func (o *AttackDefenseSetting) SetTcpConnEnable(v bool)`

SetTcpConnEnable sets TcpConnEnable field to given value.


### GetTcpConnLimit

`func (o *AttackDefenseSetting) GetTcpConnLimit() int32`

GetTcpConnLimit returns the TcpConnLimit field if non-nil, zero value otherwise.

### GetTcpConnLimitOk

`func (o *AttackDefenseSetting) GetTcpConnLimitOk() (*int32, bool)`

GetTcpConnLimitOk returns a tuple with the TcpConnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpConnLimit

`func (o *AttackDefenseSetting) SetTcpConnLimit(v int32)`

SetTcpConnLimit sets TcpConnLimit field to given value.

### HasTcpConnLimit

`func (o *AttackDefenseSetting) HasTcpConnLimit() bool`

HasTcpConnLimit returns a boolean if a field has been set.

### GetTcpFinNoAckEnable

`func (o *AttackDefenseSetting) GetTcpFinNoAckEnable() bool`

GetTcpFinNoAckEnable returns the TcpFinNoAckEnable field if non-nil, zero value otherwise.

### GetTcpFinNoAckEnableOk

`func (o *AttackDefenseSetting) GetTcpFinNoAckEnableOk() (*bool, bool)`

GetTcpFinNoAckEnableOk returns a tuple with the TcpFinNoAckEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpFinNoAckEnable

`func (o *AttackDefenseSetting) SetTcpFinNoAckEnable(v bool)`

SetTcpFinNoAckEnable sets TcpFinNoAckEnable field to given value.


### GetTcpScanEnable

`func (o *AttackDefenseSetting) GetTcpScanEnable() bool`

GetTcpScanEnable returns the TcpScanEnable field if non-nil, zero value otherwise.

### GetTcpScanEnableOk

`func (o *AttackDefenseSetting) GetTcpScanEnableOk() (*bool, bool)`

GetTcpScanEnableOk returns a tuple with the TcpScanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanEnable

`func (o *AttackDefenseSetting) SetTcpScanEnable(v bool)`

SetTcpScanEnable sets TcpScanEnable field to given value.


### GetTcpScanReject

`func (o *AttackDefenseSetting) GetTcpScanReject() bool`

GetTcpScanReject returns the TcpScanReject field if non-nil, zero value otherwise.

### GetTcpScanRejectOk

`func (o *AttackDefenseSetting) GetTcpScanRejectOk() (*bool, bool)`

GetTcpScanRejectOk returns a tuple with the TcpScanReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanReject

`func (o *AttackDefenseSetting) SetTcpScanReject(v bool)`

SetTcpScanReject sets TcpScanReject field to given value.

### HasTcpScanReject

`func (o *AttackDefenseSetting) HasTcpScanReject() bool`

HasTcpScanReject returns a boolean if a field has been set.

### GetTcpSrcEnable

`func (o *AttackDefenseSetting) GetTcpSrcEnable() bool`

GetTcpSrcEnable returns the TcpSrcEnable field if non-nil, zero value otherwise.

### GetTcpSrcEnableOk

`func (o *AttackDefenseSetting) GetTcpSrcEnableOk() (*bool, bool)`

GetTcpSrcEnableOk returns a tuple with the TcpSrcEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSrcEnable

`func (o *AttackDefenseSetting) SetTcpSrcEnable(v bool)`

SetTcpSrcEnable sets TcpSrcEnable field to given value.


### GetTcpSrcLimit

`func (o *AttackDefenseSetting) GetTcpSrcLimit() int32`

GetTcpSrcLimit returns the TcpSrcLimit field if non-nil, zero value otherwise.

### GetTcpSrcLimitOk

`func (o *AttackDefenseSetting) GetTcpSrcLimitOk() (*int32, bool)`

GetTcpSrcLimitOk returns a tuple with the TcpSrcLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSrcLimit

`func (o *AttackDefenseSetting) SetTcpSrcLimit(v int32)`

SetTcpSrcLimit sets TcpSrcLimit field to given value.

### HasTcpSrcLimit

`func (o *AttackDefenseSetting) HasTcpSrcLimit() bool`

HasTcpSrcLimit returns a boolean if a field has been set.

### GetTcpSynFinEnable

`func (o *AttackDefenseSetting) GetTcpSynFinEnable() bool`

GetTcpSynFinEnable returns the TcpSynFinEnable field if non-nil, zero value otherwise.

### GetTcpSynFinEnableOk

`func (o *AttackDefenseSetting) GetTcpSynFinEnableOk() (*bool, bool)`

GetTcpSynFinEnableOk returns a tuple with the TcpSynFinEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSynFinEnable

`func (o *AttackDefenseSetting) SetTcpSynFinEnable(v bool)`

SetTcpSynFinEnable sets TcpSynFinEnable field to given value.


### GetUdpConnEnable

`func (o *AttackDefenseSetting) GetUdpConnEnable() bool`

GetUdpConnEnable returns the UdpConnEnable field if non-nil, zero value otherwise.

### GetUdpConnEnableOk

`func (o *AttackDefenseSetting) GetUdpConnEnableOk() (*bool, bool)`

GetUdpConnEnableOk returns a tuple with the UdpConnEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpConnEnable

`func (o *AttackDefenseSetting) SetUdpConnEnable(v bool)`

SetUdpConnEnable sets UdpConnEnable field to given value.


### GetUdpConnLimit

`func (o *AttackDefenseSetting) GetUdpConnLimit() int32`

GetUdpConnLimit returns the UdpConnLimit field if non-nil, zero value otherwise.

### GetUdpConnLimitOk

`func (o *AttackDefenseSetting) GetUdpConnLimitOk() (*int32, bool)`

GetUdpConnLimitOk returns a tuple with the UdpConnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpConnLimit

`func (o *AttackDefenseSetting) SetUdpConnLimit(v int32)`

SetUdpConnLimit sets UdpConnLimit field to given value.

### HasUdpConnLimit

`func (o *AttackDefenseSetting) HasUdpConnLimit() bool`

HasUdpConnLimit returns a boolean if a field has been set.

### GetUdpSrcEnable

`func (o *AttackDefenseSetting) GetUdpSrcEnable() bool`

GetUdpSrcEnable returns the UdpSrcEnable field if non-nil, zero value otherwise.

### GetUdpSrcEnableOk

`func (o *AttackDefenseSetting) GetUdpSrcEnableOk() (*bool, bool)`

GetUdpSrcEnableOk returns a tuple with the UdpSrcEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpSrcEnable

`func (o *AttackDefenseSetting) SetUdpSrcEnable(v bool)`

SetUdpSrcEnable sets UdpSrcEnable field to given value.


### GetUdpSrcLimit

`func (o *AttackDefenseSetting) GetUdpSrcLimit() int32`

GetUdpSrcLimit returns the UdpSrcLimit field if non-nil, zero value otherwise.

### GetUdpSrcLimitOk

`func (o *AttackDefenseSetting) GetUdpSrcLimitOk() (*int32, bool)`

GetUdpSrcLimitOk returns a tuple with the UdpSrcLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpSrcLimit

`func (o *AttackDefenseSetting) SetUdpSrcLimit(v int32)`

SetUdpSrcLimit sets UdpSrcLimit field to given value.

### HasUdpSrcLimit

`func (o *AttackDefenseSetting) HasUdpSrcLimit() bool`

HasUdpSrcLimit returns a boolean if a field has been set.

### GetWinNukeAttackEnable

`func (o *AttackDefenseSetting) GetWinNukeAttackEnable() bool`

GetWinNukeAttackEnable returns the WinNukeAttackEnable field if non-nil, zero value otherwise.

### GetWinNukeAttackEnableOk

`func (o *AttackDefenseSetting) GetWinNukeAttackEnableOk() (*bool, bool)`

GetWinNukeAttackEnableOk returns a tuple with the WinNukeAttackEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinNukeAttackEnable

`func (o *AttackDefenseSetting) SetWinNukeAttackEnable(v bool)`

SetWinNukeAttackEnable sets WinNukeAttackEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


