# AttackDefenseSettingForQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistIcmpTimestampRequestReject** | Pointer to **bool** | Whether ICMP Timestamp Request of the attack defense setting is ON. | [optional] 
**ExistLargePingThreshold** | Pointer to **bool** | Whether custom Block Large Ping Threshold of attack defense is configured. | [optional] 
**ExistTcpScanReject** | Pointer to **bool** | Whether Block TCP scan with reject of the attack defense setting is ON. | [optional] 
**IcmpConnEnable** | Pointer to **bool** | Multi-connections ICMP flood enable of the attack defense setting. | [optional] 
**IcmpConnLimit** | Pointer to **int32** | Multi-connections ICMP flood limit should be within the range of 100–99999. | [optional] 
**IcmpSrcEnable** | Pointer to **bool** | Stationary source ICMP flood enable of the attack defense setting. | [optional] 
**IcmpSrcLimit** | Pointer to **int32** | Stationary source ICMP flood limit should be within the range of 100–99999. | [optional] 
**IcmpTimestampRequestReject** | Pointer to **bool** | ICMP Timestamp Request setting of the attack defense setting. | [optional] 
**LargePingEnable** | Pointer to **bool** | Block Large Ping Threshold of the attack defense setting. | [optional] 
**LargePingThreshold** | Pointer to **int32** | Block Large Ping Threshold of the attack defense setting, Value is between 28 and 65535. | [optional] 
**PingDeathEnable** | Pointer to **bool** | Block ping of death of the attack defense setting. | [optional] 
**PingWanEnable** | Pointer to **bool** | Block ping from WAN of the attack defense setting. | [optional] 
**SpecifiedOption** | Pointer to [**SpecifiedOptionOpenApiVO**](SpecifiedOptionOpenApiVO.md) |  | [optional] 
**SpecifiedOptionEnable** | Pointer to **bool** | Block packets with specified options of the attack defense setting. | [optional] 
**SupportIcmpTimestampRequestReject** | Pointer to **bool** | Whether ICMP Timestamp Request is supported of the attack defense setting. | [optional] 
**SupportLargePingThreshold** | Pointer to **bool** | Whether custom Block Large Ping Threshold of the attack defense is supported. | [optional] 
**SupportTcpScanReject** | Pointer to **bool** | Whether Block TCP scan with reject of the attack defense setting is supported. | [optional] 
**TcpConnEnable** | Pointer to **bool** | Multi-connections TCP SYN flood enable of the attack defense setting. | [optional] 
**TcpConnLimit** | Pointer to **int32** | Multi-connections TCP SYN flood limit should be within the range of 100–99999. | [optional] 
**TcpFinNoAckEnable** | Pointer to **bool** | Block TCP packets with FIN Bit set but no ACK Bit set of the attack defense setting. | [optional] 
**TcpScanEnable** | Pointer to **bool** | Block TCP scan enable of the attack defense setting. | [optional] 
**TcpScanReject** | Pointer to **bool** | Block TCP scan with reject of the attack defense setting. | [optional] 
**TcpSrcEnable** | Pointer to **bool** | Stationary source TCP SYN flood enable of the attack defense setting. | [optional] 
**TcpSrcLimit** | Pointer to **int32** | Stationary source TCP SYN flood limit should be within the range of 100–99999. | [optional] 
**TcpSynFinEnable** | Pointer to **bool** | Block TCP packets with SYN and FIN Bits set of the attack defense setting. | [optional] 
**UdpConnEnable** | Pointer to **bool** | Multi-connections UDP flood enable of the attack defense setting. | [optional] 
**UdpConnLimit** | Pointer to **int32** | Multi-connections UDP flood limit should be within the range of 100–99999. | [optional] 
**UdpSrcEnable** | Pointer to **bool** | Stationary source UDP flood enable of the attack defense setting. | [optional] 
**UdpSrcLimit** | Pointer to **int32** | Stationary source UDP flood limit should be within the range of 100–99999. | [optional] 
**WinNukeAttackEnable** | Pointer to **bool** | Block WinNuke attack of the attack defense setting. | [optional] 

## Methods

### NewAttackDefenseSettingForQuery

`func NewAttackDefenseSettingForQuery() *AttackDefenseSettingForQuery`

NewAttackDefenseSettingForQuery instantiates a new AttackDefenseSettingForQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackDefenseSettingForQueryWithDefaults

`func NewAttackDefenseSettingForQueryWithDefaults() *AttackDefenseSettingForQuery`

NewAttackDefenseSettingForQueryWithDefaults instantiates a new AttackDefenseSettingForQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) GetExistIcmpTimestampRequestReject() bool`

GetExistIcmpTimestampRequestReject returns the ExistIcmpTimestampRequestReject field if non-nil, zero value otherwise.

### GetExistIcmpTimestampRequestRejectOk

`func (o *AttackDefenseSettingForQuery) GetExistIcmpTimestampRequestRejectOk() (*bool, bool)`

GetExistIcmpTimestampRequestRejectOk returns a tuple with the ExistIcmpTimestampRequestReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) SetExistIcmpTimestampRequestReject(v bool)`

SetExistIcmpTimestampRequestReject sets ExistIcmpTimestampRequestReject field to given value.

### HasExistIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) HasExistIcmpTimestampRequestReject() bool`

HasExistIcmpTimestampRequestReject returns a boolean if a field has been set.

### GetExistLargePingThreshold

`func (o *AttackDefenseSettingForQuery) GetExistLargePingThreshold() bool`

GetExistLargePingThreshold returns the ExistLargePingThreshold field if non-nil, zero value otherwise.

### GetExistLargePingThresholdOk

`func (o *AttackDefenseSettingForQuery) GetExistLargePingThresholdOk() (*bool, bool)`

GetExistLargePingThresholdOk returns a tuple with the ExistLargePingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistLargePingThreshold

`func (o *AttackDefenseSettingForQuery) SetExistLargePingThreshold(v bool)`

SetExistLargePingThreshold sets ExistLargePingThreshold field to given value.

### HasExistLargePingThreshold

`func (o *AttackDefenseSettingForQuery) HasExistLargePingThreshold() bool`

HasExistLargePingThreshold returns a boolean if a field has been set.

### GetExistTcpScanReject

`func (o *AttackDefenseSettingForQuery) GetExistTcpScanReject() bool`

GetExistTcpScanReject returns the ExistTcpScanReject field if non-nil, zero value otherwise.

### GetExistTcpScanRejectOk

`func (o *AttackDefenseSettingForQuery) GetExistTcpScanRejectOk() (*bool, bool)`

GetExistTcpScanRejectOk returns a tuple with the ExistTcpScanReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistTcpScanReject

`func (o *AttackDefenseSettingForQuery) SetExistTcpScanReject(v bool)`

SetExistTcpScanReject sets ExistTcpScanReject field to given value.

### HasExistTcpScanReject

`func (o *AttackDefenseSettingForQuery) HasExistTcpScanReject() bool`

HasExistTcpScanReject returns a boolean if a field has been set.

### GetIcmpConnEnable

`func (o *AttackDefenseSettingForQuery) GetIcmpConnEnable() bool`

GetIcmpConnEnable returns the IcmpConnEnable field if non-nil, zero value otherwise.

### GetIcmpConnEnableOk

`func (o *AttackDefenseSettingForQuery) GetIcmpConnEnableOk() (*bool, bool)`

GetIcmpConnEnableOk returns a tuple with the IcmpConnEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpConnEnable

`func (o *AttackDefenseSettingForQuery) SetIcmpConnEnable(v bool)`

SetIcmpConnEnable sets IcmpConnEnable field to given value.

### HasIcmpConnEnable

`func (o *AttackDefenseSettingForQuery) HasIcmpConnEnable() bool`

HasIcmpConnEnable returns a boolean if a field has been set.

### GetIcmpConnLimit

`func (o *AttackDefenseSettingForQuery) GetIcmpConnLimit() int32`

GetIcmpConnLimit returns the IcmpConnLimit field if non-nil, zero value otherwise.

### GetIcmpConnLimitOk

`func (o *AttackDefenseSettingForQuery) GetIcmpConnLimitOk() (*int32, bool)`

GetIcmpConnLimitOk returns a tuple with the IcmpConnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpConnLimit

`func (o *AttackDefenseSettingForQuery) SetIcmpConnLimit(v int32)`

SetIcmpConnLimit sets IcmpConnLimit field to given value.

### HasIcmpConnLimit

`func (o *AttackDefenseSettingForQuery) HasIcmpConnLimit() bool`

HasIcmpConnLimit returns a boolean if a field has been set.

### GetIcmpSrcEnable

`func (o *AttackDefenseSettingForQuery) GetIcmpSrcEnable() bool`

GetIcmpSrcEnable returns the IcmpSrcEnable field if non-nil, zero value otherwise.

### GetIcmpSrcEnableOk

`func (o *AttackDefenseSettingForQuery) GetIcmpSrcEnableOk() (*bool, bool)`

GetIcmpSrcEnableOk returns a tuple with the IcmpSrcEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpSrcEnable

`func (o *AttackDefenseSettingForQuery) SetIcmpSrcEnable(v bool)`

SetIcmpSrcEnable sets IcmpSrcEnable field to given value.

### HasIcmpSrcEnable

`func (o *AttackDefenseSettingForQuery) HasIcmpSrcEnable() bool`

HasIcmpSrcEnable returns a boolean if a field has been set.

### GetIcmpSrcLimit

`func (o *AttackDefenseSettingForQuery) GetIcmpSrcLimit() int32`

GetIcmpSrcLimit returns the IcmpSrcLimit field if non-nil, zero value otherwise.

### GetIcmpSrcLimitOk

`func (o *AttackDefenseSettingForQuery) GetIcmpSrcLimitOk() (*int32, bool)`

GetIcmpSrcLimitOk returns a tuple with the IcmpSrcLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpSrcLimit

`func (o *AttackDefenseSettingForQuery) SetIcmpSrcLimit(v int32)`

SetIcmpSrcLimit sets IcmpSrcLimit field to given value.

### HasIcmpSrcLimit

`func (o *AttackDefenseSettingForQuery) HasIcmpSrcLimit() bool`

HasIcmpSrcLimit returns a boolean if a field has been set.

### GetIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) GetIcmpTimestampRequestReject() bool`

GetIcmpTimestampRequestReject returns the IcmpTimestampRequestReject field if non-nil, zero value otherwise.

### GetIcmpTimestampRequestRejectOk

`func (o *AttackDefenseSettingForQuery) GetIcmpTimestampRequestRejectOk() (*bool, bool)`

GetIcmpTimestampRequestRejectOk returns a tuple with the IcmpTimestampRequestReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) SetIcmpTimestampRequestReject(v bool)`

SetIcmpTimestampRequestReject sets IcmpTimestampRequestReject field to given value.

### HasIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) HasIcmpTimestampRequestReject() bool`

HasIcmpTimestampRequestReject returns a boolean if a field has been set.

### GetLargePingEnable

`func (o *AttackDefenseSettingForQuery) GetLargePingEnable() bool`

GetLargePingEnable returns the LargePingEnable field if non-nil, zero value otherwise.

### GetLargePingEnableOk

`func (o *AttackDefenseSettingForQuery) GetLargePingEnableOk() (*bool, bool)`

GetLargePingEnableOk returns a tuple with the LargePingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargePingEnable

`func (o *AttackDefenseSettingForQuery) SetLargePingEnable(v bool)`

SetLargePingEnable sets LargePingEnable field to given value.

### HasLargePingEnable

`func (o *AttackDefenseSettingForQuery) HasLargePingEnable() bool`

HasLargePingEnable returns a boolean if a field has been set.

### GetLargePingThreshold

`func (o *AttackDefenseSettingForQuery) GetLargePingThreshold() int32`

GetLargePingThreshold returns the LargePingThreshold field if non-nil, zero value otherwise.

### GetLargePingThresholdOk

`func (o *AttackDefenseSettingForQuery) GetLargePingThresholdOk() (*int32, bool)`

GetLargePingThresholdOk returns a tuple with the LargePingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargePingThreshold

`func (o *AttackDefenseSettingForQuery) SetLargePingThreshold(v int32)`

SetLargePingThreshold sets LargePingThreshold field to given value.

### HasLargePingThreshold

`func (o *AttackDefenseSettingForQuery) HasLargePingThreshold() bool`

HasLargePingThreshold returns a boolean if a field has been set.

### GetPingDeathEnable

`func (o *AttackDefenseSettingForQuery) GetPingDeathEnable() bool`

GetPingDeathEnable returns the PingDeathEnable field if non-nil, zero value otherwise.

### GetPingDeathEnableOk

`func (o *AttackDefenseSettingForQuery) GetPingDeathEnableOk() (*bool, bool)`

GetPingDeathEnableOk returns a tuple with the PingDeathEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingDeathEnable

`func (o *AttackDefenseSettingForQuery) SetPingDeathEnable(v bool)`

SetPingDeathEnable sets PingDeathEnable field to given value.

### HasPingDeathEnable

`func (o *AttackDefenseSettingForQuery) HasPingDeathEnable() bool`

HasPingDeathEnable returns a boolean if a field has been set.

### GetPingWanEnable

`func (o *AttackDefenseSettingForQuery) GetPingWanEnable() bool`

GetPingWanEnable returns the PingWanEnable field if non-nil, zero value otherwise.

### GetPingWanEnableOk

`func (o *AttackDefenseSettingForQuery) GetPingWanEnableOk() (*bool, bool)`

GetPingWanEnableOk returns a tuple with the PingWanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingWanEnable

`func (o *AttackDefenseSettingForQuery) SetPingWanEnable(v bool)`

SetPingWanEnable sets PingWanEnable field to given value.

### HasPingWanEnable

`func (o *AttackDefenseSettingForQuery) HasPingWanEnable() bool`

HasPingWanEnable returns a boolean if a field has been set.

### GetSpecifiedOption

`func (o *AttackDefenseSettingForQuery) GetSpecifiedOption() SpecifiedOptionOpenApiVO`

GetSpecifiedOption returns the SpecifiedOption field if non-nil, zero value otherwise.

### GetSpecifiedOptionOk

`func (o *AttackDefenseSettingForQuery) GetSpecifiedOptionOk() (*SpecifiedOptionOpenApiVO, bool)`

GetSpecifiedOptionOk returns a tuple with the SpecifiedOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedOption

`func (o *AttackDefenseSettingForQuery) SetSpecifiedOption(v SpecifiedOptionOpenApiVO)`

SetSpecifiedOption sets SpecifiedOption field to given value.

### HasSpecifiedOption

`func (o *AttackDefenseSettingForQuery) HasSpecifiedOption() bool`

HasSpecifiedOption returns a boolean if a field has been set.

### GetSpecifiedOptionEnable

`func (o *AttackDefenseSettingForQuery) GetSpecifiedOptionEnable() bool`

GetSpecifiedOptionEnable returns the SpecifiedOptionEnable field if non-nil, zero value otherwise.

### GetSpecifiedOptionEnableOk

`func (o *AttackDefenseSettingForQuery) GetSpecifiedOptionEnableOk() (*bool, bool)`

GetSpecifiedOptionEnableOk returns a tuple with the SpecifiedOptionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedOptionEnable

`func (o *AttackDefenseSettingForQuery) SetSpecifiedOptionEnable(v bool)`

SetSpecifiedOptionEnable sets SpecifiedOptionEnable field to given value.

### HasSpecifiedOptionEnable

`func (o *AttackDefenseSettingForQuery) HasSpecifiedOptionEnable() bool`

HasSpecifiedOptionEnable returns a boolean if a field has been set.

### GetSupportIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) GetSupportIcmpTimestampRequestReject() bool`

GetSupportIcmpTimestampRequestReject returns the SupportIcmpTimestampRequestReject field if non-nil, zero value otherwise.

### GetSupportIcmpTimestampRequestRejectOk

`func (o *AttackDefenseSettingForQuery) GetSupportIcmpTimestampRequestRejectOk() (*bool, bool)`

GetSupportIcmpTimestampRequestRejectOk returns a tuple with the SupportIcmpTimestampRequestReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) SetSupportIcmpTimestampRequestReject(v bool)`

SetSupportIcmpTimestampRequestReject sets SupportIcmpTimestampRequestReject field to given value.

### HasSupportIcmpTimestampRequestReject

`func (o *AttackDefenseSettingForQuery) HasSupportIcmpTimestampRequestReject() bool`

HasSupportIcmpTimestampRequestReject returns a boolean if a field has been set.

### GetSupportLargePingThreshold

`func (o *AttackDefenseSettingForQuery) GetSupportLargePingThreshold() bool`

GetSupportLargePingThreshold returns the SupportLargePingThreshold field if non-nil, zero value otherwise.

### GetSupportLargePingThresholdOk

`func (o *AttackDefenseSettingForQuery) GetSupportLargePingThresholdOk() (*bool, bool)`

GetSupportLargePingThresholdOk returns a tuple with the SupportLargePingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLargePingThreshold

`func (o *AttackDefenseSettingForQuery) SetSupportLargePingThreshold(v bool)`

SetSupportLargePingThreshold sets SupportLargePingThreshold field to given value.

### HasSupportLargePingThreshold

`func (o *AttackDefenseSettingForQuery) HasSupportLargePingThreshold() bool`

HasSupportLargePingThreshold returns a boolean if a field has been set.

### GetSupportTcpScanReject

`func (o *AttackDefenseSettingForQuery) GetSupportTcpScanReject() bool`

GetSupportTcpScanReject returns the SupportTcpScanReject field if non-nil, zero value otherwise.

### GetSupportTcpScanRejectOk

`func (o *AttackDefenseSettingForQuery) GetSupportTcpScanRejectOk() (*bool, bool)`

GetSupportTcpScanRejectOk returns a tuple with the SupportTcpScanReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTcpScanReject

`func (o *AttackDefenseSettingForQuery) SetSupportTcpScanReject(v bool)`

SetSupportTcpScanReject sets SupportTcpScanReject field to given value.

### HasSupportTcpScanReject

`func (o *AttackDefenseSettingForQuery) HasSupportTcpScanReject() bool`

HasSupportTcpScanReject returns a boolean if a field has been set.

### GetTcpConnEnable

`func (o *AttackDefenseSettingForQuery) GetTcpConnEnable() bool`

GetTcpConnEnable returns the TcpConnEnable field if non-nil, zero value otherwise.

### GetTcpConnEnableOk

`func (o *AttackDefenseSettingForQuery) GetTcpConnEnableOk() (*bool, bool)`

GetTcpConnEnableOk returns a tuple with the TcpConnEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpConnEnable

`func (o *AttackDefenseSettingForQuery) SetTcpConnEnable(v bool)`

SetTcpConnEnable sets TcpConnEnable field to given value.

### HasTcpConnEnable

`func (o *AttackDefenseSettingForQuery) HasTcpConnEnable() bool`

HasTcpConnEnable returns a boolean if a field has been set.

### GetTcpConnLimit

`func (o *AttackDefenseSettingForQuery) GetTcpConnLimit() int32`

GetTcpConnLimit returns the TcpConnLimit field if non-nil, zero value otherwise.

### GetTcpConnLimitOk

`func (o *AttackDefenseSettingForQuery) GetTcpConnLimitOk() (*int32, bool)`

GetTcpConnLimitOk returns a tuple with the TcpConnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpConnLimit

`func (o *AttackDefenseSettingForQuery) SetTcpConnLimit(v int32)`

SetTcpConnLimit sets TcpConnLimit field to given value.

### HasTcpConnLimit

`func (o *AttackDefenseSettingForQuery) HasTcpConnLimit() bool`

HasTcpConnLimit returns a boolean if a field has been set.

### GetTcpFinNoAckEnable

`func (o *AttackDefenseSettingForQuery) GetTcpFinNoAckEnable() bool`

GetTcpFinNoAckEnable returns the TcpFinNoAckEnable field if non-nil, zero value otherwise.

### GetTcpFinNoAckEnableOk

`func (o *AttackDefenseSettingForQuery) GetTcpFinNoAckEnableOk() (*bool, bool)`

GetTcpFinNoAckEnableOk returns a tuple with the TcpFinNoAckEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpFinNoAckEnable

`func (o *AttackDefenseSettingForQuery) SetTcpFinNoAckEnable(v bool)`

SetTcpFinNoAckEnable sets TcpFinNoAckEnable field to given value.

### HasTcpFinNoAckEnable

`func (o *AttackDefenseSettingForQuery) HasTcpFinNoAckEnable() bool`

HasTcpFinNoAckEnable returns a boolean if a field has been set.

### GetTcpScanEnable

`func (o *AttackDefenseSettingForQuery) GetTcpScanEnable() bool`

GetTcpScanEnable returns the TcpScanEnable field if non-nil, zero value otherwise.

### GetTcpScanEnableOk

`func (o *AttackDefenseSettingForQuery) GetTcpScanEnableOk() (*bool, bool)`

GetTcpScanEnableOk returns a tuple with the TcpScanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanEnable

`func (o *AttackDefenseSettingForQuery) SetTcpScanEnable(v bool)`

SetTcpScanEnable sets TcpScanEnable field to given value.

### HasTcpScanEnable

`func (o *AttackDefenseSettingForQuery) HasTcpScanEnable() bool`

HasTcpScanEnable returns a boolean if a field has been set.

### GetTcpScanReject

`func (o *AttackDefenseSettingForQuery) GetTcpScanReject() bool`

GetTcpScanReject returns the TcpScanReject field if non-nil, zero value otherwise.

### GetTcpScanRejectOk

`func (o *AttackDefenseSettingForQuery) GetTcpScanRejectOk() (*bool, bool)`

GetTcpScanRejectOk returns a tuple with the TcpScanReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanReject

`func (o *AttackDefenseSettingForQuery) SetTcpScanReject(v bool)`

SetTcpScanReject sets TcpScanReject field to given value.

### HasTcpScanReject

`func (o *AttackDefenseSettingForQuery) HasTcpScanReject() bool`

HasTcpScanReject returns a boolean if a field has been set.

### GetTcpSrcEnable

`func (o *AttackDefenseSettingForQuery) GetTcpSrcEnable() bool`

GetTcpSrcEnable returns the TcpSrcEnable field if non-nil, zero value otherwise.

### GetTcpSrcEnableOk

`func (o *AttackDefenseSettingForQuery) GetTcpSrcEnableOk() (*bool, bool)`

GetTcpSrcEnableOk returns a tuple with the TcpSrcEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSrcEnable

`func (o *AttackDefenseSettingForQuery) SetTcpSrcEnable(v bool)`

SetTcpSrcEnable sets TcpSrcEnable field to given value.

### HasTcpSrcEnable

`func (o *AttackDefenseSettingForQuery) HasTcpSrcEnable() bool`

HasTcpSrcEnable returns a boolean if a field has been set.

### GetTcpSrcLimit

`func (o *AttackDefenseSettingForQuery) GetTcpSrcLimit() int32`

GetTcpSrcLimit returns the TcpSrcLimit field if non-nil, zero value otherwise.

### GetTcpSrcLimitOk

`func (o *AttackDefenseSettingForQuery) GetTcpSrcLimitOk() (*int32, bool)`

GetTcpSrcLimitOk returns a tuple with the TcpSrcLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSrcLimit

`func (o *AttackDefenseSettingForQuery) SetTcpSrcLimit(v int32)`

SetTcpSrcLimit sets TcpSrcLimit field to given value.

### HasTcpSrcLimit

`func (o *AttackDefenseSettingForQuery) HasTcpSrcLimit() bool`

HasTcpSrcLimit returns a boolean if a field has been set.

### GetTcpSynFinEnable

`func (o *AttackDefenseSettingForQuery) GetTcpSynFinEnable() bool`

GetTcpSynFinEnable returns the TcpSynFinEnable field if non-nil, zero value otherwise.

### GetTcpSynFinEnableOk

`func (o *AttackDefenseSettingForQuery) GetTcpSynFinEnableOk() (*bool, bool)`

GetTcpSynFinEnableOk returns a tuple with the TcpSynFinEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSynFinEnable

`func (o *AttackDefenseSettingForQuery) SetTcpSynFinEnable(v bool)`

SetTcpSynFinEnable sets TcpSynFinEnable field to given value.

### HasTcpSynFinEnable

`func (o *AttackDefenseSettingForQuery) HasTcpSynFinEnable() bool`

HasTcpSynFinEnable returns a boolean if a field has been set.

### GetUdpConnEnable

`func (o *AttackDefenseSettingForQuery) GetUdpConnEnable() bool`

GetUdpConnEnable returns the UdpConnEnable field if non-nil, zero value otherwise.

### GetUdpConnEnableOk

`func (o *AttackDefenseSettingForQuery) GetUdpConnEnableOk() (*bool, bool)`

GetUdpConnEnableOk returns a tuple with the UdpConnEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpConnEnable

`func (o *AttackDefenseSettingForQuery) SetUdpConnEnable(v bool)`

SetUdpConnEnable sets UdpConnEnable field to given value.

### HasUdpConnEnable

`func (o *AttackDefenseSettingForQuery) HasUdpConnEnable() bool`

HasUdpConnEnable returns a boolean if a field has been set.

### GetUdpConnLimit

`func (o *AttackDefenseSettingForQuery) GetUdpConnLimit() int32`

GetUdpConnLimit returns the UdpConnLimit field if non-nil, zero value otherwise.

### GetUdpConnLimitOk

`func (o *AttackDefenseSettingForQuery) GetUdpConnLimitOk() (*int32, bool)`

GetUdpConnLimitOk returns a tuple with the UdpConnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpConnLimit

`func (o *AttackDefenseSettingForQuery) SetUdpConnLimit(v int32)`

SetUdpConnLimit sets UdpConnLimit field to given value.

### HasUdpConnLimit

`func (o *AttackDefenseSettingForQuery) HasUdpConnLimit() bool`

HasUdpConnLimit returns a boolean if a field has been set.

### GetUdpSrcEnable

`func (o *AttackDefenseSettingForQuery) GetUdpSrcEnable() bool`

GetUdpSrcEnable returns the UdpSrcEnable field if non-nil, zero value otherwise.

### GetUdpSrcEnableOk

`func (o *AttackDefenseSettingForQuery) GetUdpSrcEnableOk() (*bool, bool)`

GetUdpSrcEnableOk returns a tuple with the UdpSrcEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpSrcEnable

`func (o *AttackDefenseSettingForQuery) SetUdpSrcEnable(v bool)`

SetUdpSrcEnable sets UdpSrcEnable field to given value.

### HasUdpSrcEnable

`func (o *AttackDefenseSettingForQuery) HasUdpSrcEnable() bool`

HasUdpSrcEnable returns a boolean if a field has been set.

### GetUdpSrcLimit

`func (o *AttackDefenseSettingForQuery) GetUdpSrcLimit() int32`

GetUdpSrcLimit returns the UdpSrcLimit field if non-nil, zero value otherwise.

### GetUdpSrcLimitOk

`func (o *AttackDefenseSettingForQuery) GetUdpSrcLimitOk() (*int32, bool)`

GetUdpSrcLimitOk returns a tuple with the UdpSrcLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpSrcLimit

`func (o *AttackDefenseSettingForQuery) SetUdpSrcLimit(v int32)`

SetUdpSrcLimit sets UdpSrcLimit field to given value.

### HasUdpSrcLimit

`func (o *AttackDefenseSettingForQuery) HasUdpSrcLimit() bool`

HasUdpSrcLimit returns a boolean if a field has been set.

### GetWinNukeAttackEnable

`func (o *AttackDefenseSettingForQuery) GetWinNukeAttackEnable() bool`

GetWinNukeAttackEnable returns the WinNukeAttackEnable field if non-nil, zero value otherwise.

### GetWinNukeAttackEnableOk

`func (o *AttackDefenseSettingForQuery) GetWinNukeAttackEnableOk() (*bool, bool)`

GetWinNukeAttackEnableOk returns a tuple with the WinNukeAttackEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinNukeAttackEnable

`func (o *AttackDefenseSettingForQuery) SetWinNukeAttackEnable(v bool)`

SetWinNukeAttackEnable sets WinNukeAttackEnable field to given value.

### HasWinNukeAttackEnable

`func (o *AttackDefenseSettingForQuery) HasWinNukeAttackEnable() bool`

HasWinNukeAttackEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


