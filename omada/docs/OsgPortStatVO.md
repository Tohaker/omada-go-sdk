# OsgPortStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnexType** | Pointer to **int32** |  | [optional] 
**Band** | Pointer to **string** |  | [optional] 
**CardStatus** | Pointer to **int32** |  | [optional] 
**ChargeMode** | Pointer to **int32** |  | [optional] 
**DataUsage** | Pointer to **float32** |  | [optional] 
**DslModulationType** | Pointer to **int32** |  | [optional] 
**Duplex** | Pointer to **int32** | Port duplex, 1-”Half”，2-“Full” | [optional] 
**HealthLevel** | Pointer to **int32** | Wan health level | [optional] 
**InternetState** | Pointer to **int32** | Wan internet state, 0-disconnected; 1-connected; | [optional] 
**InternetStatus** | Pointer to **int32** |  | [optional] 
**Ip** | Pointer to **string** | Ip | [optional] 
**Ip2** | Pointer to **string** | Ip2 | [optional] 
**Isp** | Pointer to **string** |  | [optional] 
**IspVersion** | Pointer to **string** |  | [optional] 
**Latency** | Pointer to **int32** | Wan latency, when mode is wan and device is connected, Unit: ms | [optional] 
**LineStatus** | Pointer to **int32** |  | [optional] 
**Loss** | Pointer to **float64** | Wan packet loss rate, Unit : % | [optional] 
**Mac** | Pointer to **string** | Port mac | [optional] 
**MaxRxRate** | Pointer to **int64** |  | [optional] 
**MaxTxRate** | Pointer to **int64** |  | [optional] 
**MirroredPorts** | Pointer to [**[]MirroredPort**](MirroredPort.md) | Mirrored ports | [optional] 
**Mode** | Pointer to **int32** | Port mode, 0:WAN,1:LAN; | [optional] 
**Name** | Pointer to **string** | Port name | [optional] 
**NetType** | Pointer to **int32** |  | [optional] 
**OnlineDetection** | Pointer to **int32** |  | [optional] 
**PhysicalType** | Pointer to **int32** |  | [optional] 
**Poe** | Pointer to **bool** | Port poe power supply, 1 for active, 0 or null for not | [optional] 
**PoePower** | Pointer to **float64** | Port poe power | [optional] 
**Port** | Pointer to **int32** | Port serial number | [optional] 
**PortDesc** | Pointer to **string** | Port description | [optional] 
**Proto** | Pointer to **string** | WAN IPv4 connection type, it supports Static IP, DHCP, PPPoE, L2TP, PPTP, DS-Lite, and MAP-E. | [optional] 
**RoamingStatus** | Pointer to **int32** |  | [optional] 
**Rsrp** | Pointer to **int32** |  | [optional] 
**Rsrq** | Pointer to **int32** |  | [optional] 
**Rx** | Pointer to **int64** | Port total rx bytes | [optional] 
**RxErrorPkts** | Pointer to **int64** |  | [optional] 
**RxLineAttenuation** | Pointer to **int64** |  | [optional] 
**RxPkt** | Pointer to **int64** | Port total rx packets | [optional] 
**RxPktRate** | Pointer to **int64** | Port rx Packet rate, Unit: Pkt/s; | [optional] 
**RxRate** | Pointer to **int64** | Port rx rate, Unit: KB/s; | [optional] 
**RxSnrMargin** | Pointer to **int64** |  | [optional] 
**RxTrainingRate** | Pointer to **int64** |  | [optional] 
**Signal** | Pointer to **int32** |  | [optional] 
**SimCardUsed** | Pointer to **int32** |  | [optional] 
**SmsOperator** | Pointer to **string** |  | [optional] 
**Snr** | Pointer to **int32** |  | [optional] 
**Speed** | Pointer to **int32** | Port speed, 1-“10M”，2-“100M”，3-“1000M” | [optional] 
**Status** | Pointer to **int32** | Port status, 0-disconnected; 1-connected; | [optional] 
**SupportSms** | Pointer to **bool** |  | [optional] 
**TotalData** | Pointer to **float32** |  | [optional] 
**TrafficStatus** | Pointer to **int32** |  | [optional] 
**Tx** | Pointer to **int64** | Port total tx bytes | [optional] 
**TxErrorPkts** | Pointer to **int64** |  | [optional] 
**TxLineAttenuation** | Pointer to **int64** |  | [optional] 
**TxPkt** | Pointer to **int64** | Port total tx packets | [optional] 
**TxPktRate** | Pointer to **int64** | Port tx packet rate, Unit: Pkt/s; | [optional] 
**TxRate** | Pointer to **int64** | Port tx rate, Unit: KB/s; | [optional] 
**TxSnrMargin** | Pointer to **int64** |  | [optional] 
**TxTrainingRate** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **int32** | Port type, 0:WAN,1:WAN/LAN,2:LAN; | [optional] 
**WanIpv6Comptent** | Pointer to **int32** | Gateway wan ipv6 component version | [optional] 
**WanPortIpv4Config** | Pointer to [**OsgWanPortIpv4ConfigVO**](OsgWanPortIpv4ConfigVO.md) |  | [optional] 
**WanPortIpv6Config** | Pointer to [**OsgWanPortIpv6ConfigVO**](OsgWanPortIpv6ConfigVO.md) |  | [optional] 

## Methods

### NewOsgPortStatVO

`func NewOsgPortStatVO() *OsgPortStatVO`

NewOsgPortStatVO instantiates a new OsgPortStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgPortStatVOWithDefaults

`func NewOsgPortStatVOWithDefaults() *OsgPortStatVO`

NewOsgPortStatVOWithDefaults instantiates a new OsgPortStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnexType

`func (o *OsgPortStatVO) GetAnnexType() int32`

GetAnnexType returns the AnnexType field if non-nil, zero value otherwise.

### GetAnnexTypeOk

`func (o *OsgPortStatVO) GetAnnexTypeOk() (*int32, bool)`

GetAnnexTypeOk returns a tuple with the AnnexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnexType

`func (o *OsgPortStatVO) SetAnnexType(v int32)`

SetAnnexType sets AnnexType field to given value.

### HasAnnexType

`func (o *OsgPortStatVO) HasAnnexType() bool`

HasAnnexType returns a boolean if a field has been set.

### GetBand

`func (o *OsgPortStatVO) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *OsgPortStatVO) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *OsgPortStatVO) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *OsgPortStatVO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetCardStatus

`func (o *OsgPortStatVO) GetCardStatus() int32`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *OsgPortStatVO) GetCardStatusOk() (*int32, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *OsgPortStatVO) SetCardStatus(v int32)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *OsgPortStatVO) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetChargeMode

`func (o *OsgPortStatVO) GetChargeMode() int32`

GetChargeMode returns the ChargeMode field if non-nil, zero value otherwise.

### GetChargeModeOk

`func (o *OsgPortStatVO) GetChargeModeOk() (*int32, bool)`

GetChargeModeOk returns a tuple with the ChargeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeMode

`func (o *OsgPortStatVO) SetChargeMode(v int32)`

SetChargeMode sets ChargeMode field to given value.

### HasChargeMode

`func (o *OsgPortStatVO) HasChargeMode() bool`

HasChargeMode returns a boolean if a field has been set.

### GetDataUsage

`func (o *OsgPortStatVO) GetDataUsage() float32`

GetDataUsage returns the DataUsage field if non-nil, zero value otherwise.

### GetDataUsageOk

`func (o *OsgPortStatVO) GetDataUsageOk() (*float32, bool)`

GetDataUsageOk returns a tuple with the DataUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUsage

`func (o *OsgPortStatVO) SetDataUsage(v float32)`

SetDataUsage sets DataUsage field to given value.

### HasDataUsage

`func (o *OsgPortStatVO) HasDataUsage() bool`

HasDataUsage returns a boolean if a field has been set.

### GetDslModulationType

`func (o *OsgPortStatVO) GetDslModulationType() int32`

GetDslModulationType returns the DslModulationType field if non-nil, zero value otherwise.

### GetDslModulationTypeOk

`func (o *OsgPortStatVO) GetDslModulationTypeOk() (*int32, bool)`

GetDslModulationTypeOk returns a tuple with the DslModulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslModulationType

`func (o *OsgPortStatVO) SetDslModulationType(v int32)`

SetDslModulationType sets DslModulationType field to given value.

### HasDslModulationType

`func (o *OsgPortStatVO) HasDslModulationType() bool`

HasDslModulationType returns a boolean if a field has been set.

### GetDuplex

`func (o *OsgPortStatVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OsgPortStatVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OsgPortStatVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OsgPortStatVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetHealthLevel

`func (o *OsgPortStatVO) GetHealthLevel() int32`

GetHealthLevel returns the HealthLevel field if non-nil, zero value otherwise.

### GetHealthLevelOk

`func (o *OsgPortStatVO) GetHealthLevelOk() (*int32, bool)`

GetHealthLevelOk returns a tuple with the HealthLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthLevel

`func (o *OsgPortStatVO) SetHealthLevel(v int32)`

SetHealthLevel sets HealthLevel field to given value.

### HasHealthLevel

`func (o *OsgPortStatVO) HasHealthLevel() bool`

HasHealthLevel returns a boolean if a field has been set.

### GetInternetState

`func (o *OsgPortStatVO) GetInternetState() int32`

GetInternetState returns the InternetState field if non-nil, zero value otherwise.

### GetInternetStateOk

`func (o *OsgPortStatVO) GetInternetStateOk() (*int32, bool)`

GetInternetStateOk returns a tuple with the InternetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetState

`func (o *OsgPortStatVO) SetInternetState(v int32)`

SetInternetState sets InternetState field to given value.

### HasInternetState

`func (o *OsgPortStatVO) HasInternetState() bool`

HasInternetState returns a boolean if a field has been set.

### GetInternetStatus

`func (o *OsgPortStatVO) GetInternetStatus() int32`

GetInternetStatus returns the InternetStatus field if non-nil, zero value otherwise.

### GetInternetStatusOk

`func (o *OsgPortStatVO) GetInternetStatusOk() (*int32, bool)`

GetInternetStatusOk returns a tuple with the InternetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetStatus

`func (o *OsgPortStatVO) SetInternetStatus(v int32)`

SetInternetStatus sets InternetStatus field to given value.

### HasInternetStatus

`func (o *OsgPortStatVO) HasInternetStatus() bool`

HasInternetStatus returns a boolean if a field has been set.

### GetIp

`func (o *OsgPortStatVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OsgPortStatVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OsgPortStatVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OsgPortStatVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp2

`func (o *OsgPortStatVO) GetIp2() string`

GetIp2 returns the Ip2 field if non-nil, zero value otherwise.

### GetIp2Ok

`func (o *OsgPortStatVO) GetIp2Ok() (*string, bool)`

GetIp2Ok returns a tuple with the Ip2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp2

`func (o *OsgPortStatVO) SetIp2(v string)`

SetIp2 sets Ip2 field to given value.

### HasIp2

`func (o *OsgPortStatVO) HasIp2() bool`

HasIp2 returns a boolean if a field has been set.

### GetIsp

`func (o *OsgPortStatVO) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *OsgPortStatVO) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *OsgPortStatVO) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *OsgPortStatVO) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetIspVersion

`func (o *OsgPortStatVO) GetIspVersion() string`

GetIspVersion returns the IspVersion field if non-nil, zero value otherwise.

### GetIspVersionOk

`func (o *OsgPortStatVO) GetIspVersionOk() (*string, bool)`

GetIspVersionOk returns a tuple with the IspVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspVersion

`func (o *OsgPortStatVO) SetIspVersion(v string)`

SetIspVersion sets IspVersion field to given value.

### HasIspVersion

`func (o *OsgPortStatVO) HasIspVersion() bool`

HasIspVersion returns a boolean if a field has been set.

### GetLatency

`func (o *OsgPortStatVO) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *OsgPortStatVO) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *OsgPortStatVO) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *OsgPortStatVO) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLineStatus

`func (o *OsgPortStatVO) GetLineStatus() int32`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *OsgPortStatVO) GetLineStatusOk() (*int32, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *OsgPortStatVO) SetLineStatus(v int32)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *OsgPortStatVO) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.

### GetLoss

`func (o *OsgPortStatVO) GetLoss() float64`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *OsgPortStatVO) GetLossOk() (*float64, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *OsgPortStatVO) SetLoss(v float64)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *OsgPortStatVO) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetMac

`func (o *OsgPortStatVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OsgPortStatVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OsgPortStatVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OsgPortStatVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMaxRxRate

`func (o *OsgPortStatVO) GetMaxRxRate() int64`

GetMaxRxRate returns the MaxRxRate field if non-nil, zero value otherwise.

### GetMaxRxRateOk

`func (o *OsgPortStatVO) GetMaxRxRateOk() (*int64, bool)`

GetMaxRxRateOk returns a tuple with the MaxRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRxRate

`func (o *OsgPortStatVO) SetMaxRxRate(v int64)`

SetMaxRxRate sets MaxRxRate field to given value.

### HasMaxRxRate

`func (o *OsgPortStatVO) HasMaxRxRate() bool`

HasMaxRxRate returns a boolean if a field has been set.

### GetMaxTxRate

`func (o *OsgPortStatVO) GetMaxTxRate() int64`

GetMaxTxRate returns the MaxTxRate field if non-nil, zero value otherwise.

### GetMaxTxRateOk

`func (o *OsgPortStatVO) GetMaxTxRateOk() (*int64, bool)`

GetMaxTxRateOk returns a tuple with the MaxTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTxRate

`func (o *OsgPortStatVO) SetMaxTxRate(v int64)`

SetMaxTxRate sets MaxTxRate field to given value.

### HasMaxTxRate

`func (o *OsgPortStatVO) HasMaxTxRate() bool`

HasMaxTxRate returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *OsgPortStatVO) GetMirroredPorts() []MirroredPort`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *OsgPortStatVO) GetMirroredPortsOk() (*[]MirroredPort, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *OsgPortStatVO) SetMirroredPorts(v []MirroredPort)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *OsgPortStatVO) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetMode

`func (o *OsgPortStatVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OsgPortStatVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OsgPortStatVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OsgPortStatVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *OsgPortStatVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsgPortStatVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsgPortStatVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsgPortStatVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetType

`func (o *OsgPortStatVO) GetNetType() int32`

GetNetType returns the NetType field if non-nil, zero value otherwise.

### GetNetTypeOk

`func (o *OsgPortStatVO) GetNetTypeOk() (*int32, bool)`

GetNetTypeOk returns a tuple with the NetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetType

`func (o *OsgPortStatVO) SetNetType(v int32)`

SetNetType sets NetType field to given value.

### HasNetType

`func (o *OsgPortStatVO) HasNetType() bool`

HasNetType returns a boolean if a field has been set.

### GetOnlineDetection

`func (o *OsgPortStatVO) GetOnlineDetection() int32`

GetOnlineDetection returns the OnlineDetection field if non-nil, zero value otherwise.

### GetOnlineDetectionOk

`func (o *OsgPortStatVO) GetOnlineDetectionOk() (*int32, bool)`

GetOnlineDetectionOk returns a tuple with the OnlineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDetection

`func (o *OsgPortStatVO) SetOnlineDetection(v int32)`

SetOnlineDetection sets OnlineDetection field to given value.

### HasOnlineDetection

`func (o *OsgPortStatVO) HasOnlineDetection() bool`

HasOnlineDetection returns a boolean if a field has been set.

### GetPhysicalType

`func (o *OsgPortStatVO) GetPhysicalType() int32`

GetPhysicalType returns the PhysicalType field if non-nil, zero value otherwise.

### GetPhysicalTypeOk

`func (o *OsgPortStatVO) GetPhysicalTypeOk() (*int32, bool)`

GetPhysicalTypeOk returns a tuple with the PhysicalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalType

`func (o *OsgPortStatVO) SetPhysicalType(v int32)`

SetPhysicalType sets PhysicalType field to given value.

### HasPhysicalType

`func (o *OsgPortStatVO) HasPhysicalType() bool`

HasPhysicalType returns a boolean if a field has been set.

### GetPoe

`func (o *OsgPortStatVO) GetPoe() bool`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OsgPortStatVO) GetPoeOk() (*bool, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OsgPortStatVO) SetPoe(v bool)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OsgPortStatVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPoePower

`func (o *OsgPortStatVO) GetPoePower() float64`

GetPoePower returns the PoePower field if non-nil, zero value otherwise.

### GetPoePowerOk

`func (o *OsgPortStatVO) GetPoePowerOk() (*float64, bool)`

GetPoePowerOk returns a tuple with the PoePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePower

`func (o *OsgPortStatVO) SetPoePower(v float64)`

SetPoePower sets PoePower field to given value.

### HasPoePower

`func (o *OsgPortStatVO) HasPoePower() bool`

HasPoePower returns a boolean if a field has been set.

### GetPort

`func (o *OsgPortStatVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OsgPortStatVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OsgPortStatVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OsgPortStatVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortDesc

`func (o *OsgPortStatVO) GetPortDesc() string`

GetPortDesc returns the PortDesc field if non-nil, zero value otherwise.

### GetPortDescOk

`func (o *OsgPortStatVO) GetPortDescOk() (*string, bool)`

GetPortDescOk returns a tuple with the PortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDesc

`func (o *OsgPortStatVO) SetPortDesc(v string)`

SetPortDesc sets PortDesc field to given value.

### HasPortDesc

`func (o *OsgPortStatVO) HasPortDesc() bool`

HasPortDesc returns a boolean if a field has been set.

### GetProto

`func (o *OsgPortStatVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *OsgPortStatVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *OsgPortStatVO) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *OsgPortStatVO) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetRoamingStatus

`func (o *OsgPortStatVO) GetRoamingStatus() int32`

GetRoamingStatus returns the RoamingStatus field if non-nil, zero value otherwise.

### GetRoamingStatusOk

`func (o *OsgPortStatVO) GetRoamingStatusOk() (*int32, bool)`

GetRoamingStatusOk returns a tuple with the RoamingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingStatus

`func (o *OsgPortStatVO) SetRoamingStatus(v int32)`

SetRoamingStatus sets RoamingStatus field to given value.

### HasRoamingStatus

`func (o *OsgPortStatVO) HasRoamingStatus() bool`

HasRoamingStatus returns a boolean if a field has been set.

### GetRsrp

`func (o *OsgPortStatVO) GetRsrp() int32`

GetRsrp returns the Rsrp field if non-nil, zero value otherwise.

### GetRsrpOk

`func (o *OsgPortStatVO) GetRsrpOk() (*int32, bool)`

GetRsrpOk returns a tuple with the Rsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsrp

`func (o *OsgPortStatVO) SetRsrp(v int32)`

SetRsrp sets Rsrp field to given value.

### HasRsrp

`func (o *OsgPortStatVO) HasRsrp() bool`

HasRsrp returns a boolean if a field has been set.

### GetRsrq

`func (o *OsgPortStatVO) GetRsrq() int32`

GetRsrq returns the Rsrq field if non-nil, zero value otherwise.

### GetRsrqOk

`func (o *OsgPortStatVO) GetRsrqOk() (*int32, bool)`

GetRsrqOk returns a tuple with the Rsrq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsrq

`func (o *OsgPortStatVO) SetRsrq(v int32)`

SetRsrq sets Rsrq field to given value.

### HasRsrq

`func (o *OsgPortStatVO) HasRsrq() bool`

HasRsrq returns a boolean if a field has been set.

### GetRx

`func (o *OsgPortStatVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OsgPortStatVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OsgPortStatVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OsgPortStatVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxErrorPkts

`func (o *OsgPortStatVO) GetRxErrorPkts() int64`

GetRxErrorPkts returns the RxErrorPkts field if non-nil, zero value otherwise.

### GetRxErrorPktsOk

`func (o *OsgPortStatVO) GetRxErrorPktsOk() (*int64, bool)`

GetRxErrorPktsOk returns a tuple with the RxErrorPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrorPkts

`func (o *OsgPortStatVO) SetRxErrorPkts(v int64)`

SetRxErrorPkts sets RxErrorPkts field to given value.

### HasRxErrorPkts

`func (o *OsgPortStatVO) HasRxErrorPkts() bool`

HasRxErrorPkts returns a boolean if a field has been set.

### GetRxLineAttenuation

`func (o *OsgPortStatVO) GetRxLineAttenuation() int64`

GetRxLineAttenuation returns the RxLineAttenuation field if non-nil, zero value otherwise.

### GetRxLineAttenuationOk

`func (o *OsgPortStatVO) GetRxLineAttenuationOk() (*int64, bool)`

GetRxLineAttenuationOk returns a tuple with the RxLineAttenuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxLineAttenuation

`func (o *OsgPortStatVO) SetRxLineAttenuation(v int64)`

SetRxLineAttenuation sets RxLineAttenuation field to given value.

### HasRxLineAttenuation

`func (o *OsgPortStatVO) HasRxLineAttenuation() bool`

HasRxLineAttenuation returns a boolean if a field has been set.

### GetRxPkt

`func (o *OsgPortStatVO) GetRxPkt() int64`

GetRxPkt returns the RxPkt field if non-nil, zero value otherwise.

### GetRxPktOk

`func (o *OsgPortStatVO) GetRxPktOk() (*int64, bool)`

GetRxPktOk returns a tuple with the RxPkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkt

`func (o *OsgPortStatVO) SetRxPkt(v int64)`

SetRxPkt sets RxPkt field to given value.

### HasRxPkt

`func (o *OsgPortStatVO) HasRxPkt() bool`

HasRxPkt returns a boolean if a field has been set.

### GetRxPktRate

`func (o *OsgPortStatVO) GetRxPktRate() int64`

GetRxPktRate returns the RxPktRate field if non-nil, zero value otherwise.

### GetRxPktRateOk

`func (o *OsgPortStatVO) GetRxPktRateOk() (*int64, bool)`

GetRxPktRateOk returns a tuple with the RxPktRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPktRate

`func (o *OsgPortStatVO) SetRxPktRate(v int64)`

SetRxPktRate sets RxPktRate field to given value.

### HasRxPktRate

`func (o *OsgPortStatVO) HasRxPktRate() bool`

HasRxPktRate returns a boolean if a field has been set.

### GetRxRate

`func (o *OsgPortStatVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OsgPortStatVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OsgPortStatVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OsgPortStatVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetRxSnrMargin

`func (o *OsgPortStatVO) GetRxSnrMargin() int64`

GetRxSnrMargin returns the RxSnrMargin field if non-nil, zero value otherwise.

### GetRxSnrMarginOk

`func (o *OsgPortStatVO) GetRxSnrMarginOk() (*int64, bool)`

GetRxSnrMarginOk returns a tuple with the RxSnrMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxSnrMargin

`func (o *OsgPortStatVO) SetRxSnrMargin(v int64)`

SetRxSnrMargin sets RxSnrMargin field to given value.

### HasRxSnrMargin

`func (o *OsgPortStatVO) HasRxSnrMargin() bool`

HasRxSnrMargin returns a boolean if a field has been set.

### GetRxTrainingRate

`func (o *OsgPortStatVO) GetRxTrainingRate() int64`

GetRxTrainingRate returns the RxTrainingRate field if non-nil, zero value otherwise.

### GetRxTrainingRateOk

`func (o *OsgPortStatVO) GetRxTrainingRateOk() (*int64, bool)`

GetRxTrainingRateOk returns a tuple with the RxTrainingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxTrainingRate

`func (o *OsgPortStatVO) SetRxTrainingRate(v int64)`

SetRxTrainingRate sets RxTrainingRate field to given value.

### HasRxTrainingRate

`func (o *OsgPortStatVO) HasRxTrainingRate() bool`

HasRxTrainingRate returns a boolean if a field has been set.

### GetSignal

`func (o *OsgPortStatVO) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *OsgPortStatVO) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *OsgPortStatVO) SetSignal(v int32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *OsgPortStatVO) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetSimCardUsed

`func (o *OsgPortStatVO) GetSimCardUsed() int32`

GetSimCardUsed returns the SimCardUsed field if non-nil, zero value otherwise.

### GetSimCardUsedOk

`func (o *OsgPortStatVO) GetSimCardUsedOk() (*int32, bool)`

GetSimCardUsedOk returns a tuple with the SimCardUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardUsed

`func (o *OsgPortStatVO) SetSimCardUsed(v int32)`

SetSimCardUsed sets SimCardUsed field to given value.

### HasSimCardUsed

`func (o *OsgPortStatVO) HasSimCardUsed() bool`

HasSimCardUsed returns a boolean if a field has been set.

### GetSmsOperator

`func (o *OsgPortStatVO) GetSmsOperator() string`

GetSmsOperator returns the SmsOperator field if non-nil, zero value otherwise.

### GetSmsOperatorOk

`func (o *OsgPortStatVO) GetSmsOperatorOk() (*string, bool)`

GetSmsOperatorOk returns a tuple with the SmsOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsOperator

`func (o *OsgPortStatVO) SetSmsOperator(v string)`

SetSmsOperator sets SmsOperator field to given value.

### HasSmsOperator

`func (o *OsgPortStatVO) HasSmsOperator() bool`

HasSmsOperator returns a boolean if a field has been set.

### GetSnr

`func (o *OsgPortStatVO) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *OsgPortStatVO) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *OsgPortStatVO) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *OsgPortStatVO) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetSpeed

`func (o *OsgPortStatVO) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OsgPortStatVO) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OsgPortStatVO) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *OsgPortStatVO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *OsgPortStatVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsgPortStatVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsgPortStatVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsgPortStatVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportSms

`func (o *OsgPortStatVO) GetSupportSms() bool`

GetSupportSms returns the SupportSms field if non-nil, zero value otherwise.

### GetSupportSmsOk

`func (o *OsgPortStatVO) GetSupportSmsOk() (*bool, bool)`

GetSupportSmsOk returns a tuple with the SupportSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSms

`func (o *OsgPortStatVO) SetSupportSms(v bool)`

SetSupportSms sets SupportSms field to given value.

### HasSupportSms

`func (o *OsgPortStatVO) HasSupportSms() bool`

HasSupportSms returns a boolean if a field has been set.

### GetTotalData

`func (o *OsgPortStatVO) GetTotalData() float32`

GetTotalData returns the TotalData field if non-nil, zero value otherwise.

### GetTotalDataOk

`func (o *OsgPortStatVO) GetTotalDataOk() (*float32, bool)`

GetTotalDataOk returns a tuple with the TotalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalData

`func (o *OsgPortStatVO) SetTotalData(v float32)`

SetTotalData sets TotalData field to given value.

### HasTotalData

`func (o *OsgPortStatVO) HasTotalData() bool`

HasTotalData returns a boolean if a field has been set.

### GetTrafficStatus

`func (o *OsgPortStatVO) GetTrafficStatus() int32`

GetTrafficStatus returns the TrafficStatus field if non-nil, zero value otherwise.

### GetTrafficStatusOk

`func (o *OsgPortStatVO) GetTrafficStatusOk() (*int32, bool)`

GetTrafficStatusOk returns a tuple with the TrafficStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficStatus

`func (o *OsgPortStatVO) SetTrafficStatus(v int32)`

SetTrafficStatus sets TrafficStatus field to given value.

### HasTrafficStatus

`func (o *OsgPortStatVO) HasTrafficStatus() bool`

HasTrafficStatus returns a boolean if a field has been set.

### GetTx

`func (o *OsgPortStatVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OsgPortStatVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OsgPortStatVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OsgPortStatVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxErrorPkts

`func (o *OsgPortStatVO) GetTxErrorPkts() int64`

GetTxErrorPkts returns the TxErrorPkts field if non-nil, zero value otherwise.

### GetTxErrorPktsOk

`func (o *OsgPortStatVO) GetTxErrorPktsOk() (*int64, bool)`

GetTxErrorPktsOk returns a tuple with the TxErrorPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrorPkts

`func (o *OsgPortStatVO) SetTxErrorPkts(v int64)`

SetTxErrorPkts sets TxErrorPkts field to given value.

### HasTxErrorPkts

`func (o *OsgPortStatVO) HasTxErrorPkts() bool`

HasTxErrorPkts returns a boolean if a field has been set.

### GetTxLineAttenuation

`func (o *OsgPortStatVO) GetTxLineAttenuation() int64`

GetTxLineAttenuation returns the TxLineAttenuation field if non-nil, zero value otherwise.

### GetTxLineAttenuationOk

`func (o *OsgPortStatVO) GetTxLineAttenuationOk() (*int64, bool)`

GetTxLineAttenuationOk returns a tuple with the TxLineAttenuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxLineAttenuation

`func (o *OsgPortStatVO) SetTxLineAttenuation(v int64)`

SetTxLineAttenuation sets TxLineAttenuation field to given value.

### HasTxLineAttenuation

`func (o *OsgPortStatVO) HasTxLineAttenuation() bool`

HasTxLineAttenuation returns a boolean if a field has been set.

### GetTxPkt

`func (o *OsgPortStatVO) GetTxPkt() int64`

GetTxPkt returns the TxPkt field if non-nil, zero value otherwise.

### GetTxPktOk

`func (o *OsgPortStatVO) GetTxPktOk() (*int64, bool)`

GetTxPktOk returns a tuple with the TxPkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkt

`func (o *OsgPortStatVO) SetTxPkt(v int64)`

SetTxPkt sets TxPkt field to given value.

### HasTxPkt

`func (o *OsgPortStatVO) HasTxPkt() bool`

HasTxPkt returns a boolean if a field has been set.

### GetTxPktRate

`func (o *OsgPortStatVO) GetTxPktRate() int64`

GetTxPktRate returns the TxPktRate field if non-nil, zero value otherwise.

### GetTxPktRateOk

`func (o *OsgPortStatVO) GetTxPktRateOk() (*int64, bool)`

GetTxPktRateOk returns a tuple with the TxPktRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPktRate

`func (o *OsgPortStatVO) SetTxPktRate(v int64)`

SetTxPktRate sets TxPktRate field to given value.

### HasTxPktRate

`func (o *OsgPortStatVO) HasTxPktRate() bool`

HasTxPktRate returns a boolean if a field has been set.

### GetTxRate

`func (o *OsgPortStatVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OsgPortStatVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OsgPortStatVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OsgPortStatVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetTxSnrMargin

`func (o *OsgPortStatVO) GetTxSnrMargin() int64`

GetTxSnrMargin returns the TxSnrMargin field if non-nil, zero value otherwise.

### GetTxSnrMarginOk

`func (o *OsgPortStatVO) GetTxSnrMarginOk() (*int64, bool)`

GetTxSnrMarginOk returns a tuple with the TxSnrMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSnrMargin

`func (o *OsgPortStatVO) SetTxSnrMargin(v int64)`

SetTxSnrMargin sets TxSnrMargin field to given value.

### HasTxSnrMargin

`func (o *OsgPortStatVO) HasTxSnrMargin() bool`

HasTxSnrMargin returns a boolean if a field has been set.

### GetTxTrainingRate

`func (o *OsgPortStatVO) GetTxTrainingRate() int64`

GetTxTrainingRate returns the TxTrainingRate field if non-nil, zero value otherwise.

### GetTxTrainingRateOk

`func (o *OsgPortStatVO) GetTxTrainingRateOk() (*int64, bool)`

GetTxTrainingRateOk returns a tuple with the TxTrainingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxTrainingRate

`func (o *OsgPortStatVO) SetTxTrainingRate(v int64)`

SetTxTrainingRate sets TxTrainingRate field to given value.

### HasTxTrainingRate

`func (o *OsgPortStatVO) HasTxTrainingRate() bool`

HasTxTrainingRate returns a boolean if a field has been set.

### GetType

`func (o *OsgPortStatVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsgPortStatVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsgPortStatVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OsgPortStatVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWanIpv6Comptent

`func (o *OsgPortStatVO) GetWanIpv6Comptent() int32`

GetWanIpv6Comptent returns the WanIpv6Comptent field if non-nil, zero value otherwise.

### GetWanIpv6ComptentOk

`func (o *OsgPortStatVO) GetWanIpv6ComptentOk() (*int32, bool)`

GetWanIpv6ComptentOk returns a tuple with the WanIpv6Comptent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIpv6Comptent

`func (o *OsgPortStatVO) SetWanIpv6Comptent(v int32)`

SetWanIpv6Comptent sets WanIpv6Comptent field to given value.

### HasWanIpv6Comptent

`func (o *OsgPortStatVO) HasWanIpv6Comptent() bool`

HasWanIpv6Comptent returns a boolean if a field has been set.

### GetWanPortIpv4Config

`func (o *OsgPortStatVO) GetWanPortIpv4Config() OsgWanPortIpv4ConfigVO`

GetWanPortIpv4Config returns the WanPortIpv4Config field if non-nil, zero value otherwise.

### GetWanPortIpv4ConfigOk

`func (o *OsgPortStatVO) GetWanPortIpv4ConfigOk() (*OsgWanPortIpv4ConfigVO, bool)`

GetWanPortIpv4ConfigOk returns a tuple with the WanPortIpv4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIpv4Config

`func (o *OsgPortStatVO) SetWanPortIpv4Config(v OsgWanPortIpv4ConfigVO)`

SetWanPortIpv4Config sets WanPortIpv4Config field to given value.

### HasWanPortIpv4Config

`func (o *OsgPortStatVO) HasWanPortIpv4Config() bool`

HasWanPortIpv4Config returns a boolean if a field has been set.

### GetWanPortIpv6Config

`func (o *OsgPortStatVO) GetWanPortIpv6Config() OsgWanPortIpv6ConfigVO`

GetWanPortIpv6Config returns the WanPortIpv6Config field if non-nil, zero value otherwise.

### GetWanPortIpv6ConfigOk

`func (o *OsgPortStatVO) GetWanPortIpv6ConfigOk() (*OsgWanPortIpv6ConfigVO, bool)`

GetWanPortIpv6ConfigOk returns a tuple with the WanPortIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIpv6Config

`func (o *OsgPortStatVO) SetWanPortIpv6Config(v OsgWanPortIpv6ConfigVO)`

SetWanPortIpv6Config sets WanPortIpv6Config field to given value.

### HasWanPortIpv6Config

`func (o *OsgPortStatVO) HasWanPortIpv6Config() bool`

HasWanPortIpv6Config returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


