# GatewayWanStatusEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnexType** | Pointer to **int32** |  | [optional] 
**Band** | Pointer to **string** |  | [optional] 
**CardStatus** | Pointer to **int32** |  | [optional] 
**DataUsage** | Pointer to **float32** |  | [optional] 
**DslModulationType** | Pointer to **int32** |  | [optional] 
**Duplex** | Pointer to **int32** | Port duplex, 1-Half，2-Full | [optional] 
**HealthLevel** | Pointer to **int32** | Wan health level | [optional] 
**InternetState** | Pointer to **int32** | Wan internet state should be a value as follows: 0: disconnected; 1: connected | [optional] 
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
**Signal** | Pointer to **int32** |  | [optional] 
**SimCardUsed** | Pointer to **int32** |  | [optional] 
**SmsOperator** | Pointer to **string** |  | [optional] 
**Snr** | Pointer to **int32** |  | [optional] 
**Speed** | Pointer to **int32** | Port speed, 1-10M，2-100M，3-1000M | [optional] 
**Status** | Pointer to **int32** | Port status should be a value as follows: 0: disconnected; 1: connected | [optional] 
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
**Type** | Pointer to **int32** | Port type, 0:WAN,1:WAN/LAN,2:LAN; | [optional] 
**WanIpv6Comptent** | Pointer to **int32** | Gateway wan ipv6 component version | [optional] 
**WanPortIpv4Config** | Pointer to [**OsgWanPortIpv4ConfigVO**](OsgWanPortIpv4ConfigVO.md) |  | [optional] 
**WanPortIpv6Config** | Pointer to [**OsgWanPortIpv6ConfigVO**](OsgWanPortIpv6ConfigVO.md) |  | [optional] 

## Methods

### NewGatewayWanStatusEntity

`func NewGatewayWanStatusEntity() *GatewayWanStatusEntity`

NewGatewayWanStatusEntity instantiates a new GatewayWanStatusEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWanStatusEntityWithDefaults

`func NewGatewayWanStatusEntityWithDefaults() *GatewayWanStatusEntity`

NewGatewayWanStatusEntityWithDefaults instantiates a new GatewayWanStatusEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnexType

`func (o *GatewayWanStatusEntity) GetAnnexType() int32`

GetAnnexType returns the AnnexType field if non-nil, zero value otherwise.

### GetAnnexTypeOk

`func (o *GatewayWanStatusEntity) GetAnnexTypeOk() (*int32, bool)`

GetAnnexTypeOk returns a tuple with the AnnexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnexType

`func (o *GatewayWanStatusEntity) SetAnnexType(v int32)`

SetAnnexType sets AnnexType field to given value.

### HasAnnexType

`func (o *GatewayWanStatusEntity) HasAnnexType() bool`

HasAnnexType returns a boolean if a field has been set.

### GetBand

`func (o *GatewayWanStatusEntity) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *GatewayWanStatusEntity) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *GatewayWanStatusEntity) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *GatewayWanStatusEntity) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetCardStatus

`func (o *GatewayWanStatusEntity) GetCardStatus() int32`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *GatewayWanStatusEntity) GetCardStatusOk() (*int32, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *GatewayWanStatusEntity) SetCardStatus(v int32)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *GatewayWanStatusEntity) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetDataUsage

`func (o *GatewayWanStatusEntity) GetDataUsage() float32`

GetDataUsage returns the DataUsage field if non-nil, zero value otherwise.

### GetDataUsageOk

`func (o *GatewayWanStatusEntity) GetDataUsageOk() (*float32, bool)`

GetDataUsageOk returns a tuple with the DataUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUsage

`func (o *GatewayWanStatusEntity) SetDataUsage(v float32)`

SetDataUsage sets DataUsage field to given value.

### HasDataUsage

`func (o *GatewayWanStatusEntity) HasDataUsage() bool`

HasDataUsage returns a boolean if a field has been set.

### GetDslModulationType

`func (o *GatewayWanStatusEntity) GetDslModulationType() int32`

GetDslModulationType returns the DslModulationType field if non-nil, zero value otherwise.

### GetDslModulationTypeOk

`func (o *GatewayWanStatusEntity) GetDslModulationTypeOk() (*int32, bool)`

GetDslModulationTypeOk returns a tuple with the DslModulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslModulationType

`func (o *GatewayWanStatusEntity) SetDslModulationType(v int32)`

SetDslModulationType sets DslModulationType field to given value.

### HasDslModulationType

`func (o *GatewayWanStatusEntity) HasDslModulationType() bool`

HasDslModulationType returns a boolean if a field has been set.

### GetDuplex

`func (o *GatewayWanStatusEntity) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *GatewayWanStatusEntity) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *GatewayWanStatusEntity) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *GatewayWanStatusEntity) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetHealthLevel

`func (o *GatewayWanStatusEntity) GetHealthLevel() int32`

GetHealthLevel returns the HealthLevel field if non-nil, zero value otherwise.

### GetHealthLevelOk

`func (o *GatewayWanStatusEntity) GetHealthLevelOk() (*int32, bool)`

GetHealthLevelOk returns a tuple with the HealthLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthLevel

`func (o *GatewayWanStatusEntity) SetHealthLevel(v int32)`

SetHealthLevel sets HealthLevel field to given value.

### HasHealthLevel

`func (o *GatewayWanStatusEntity) HasHealthLevel() bool`

HasHealthLevel returns a boolean if a field has been set.

### GetInternetState

`func (o *GatewayWanStatusEntity) GetInternetState() int32`

GetInternetState returns the InternetState field if non-nil, zero value otherwise.

### GetInternetStateOk

`func (o *GatewayWanStatusEntity) GetInternetStateOk() (*int32, bool)`

GetInternetStateOk returns a tuple with the InternetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetState

`func (o *GatewayWanStatusEntity) SetInternetState(v int32)`

SetInternetState sets InternetState field to given value.

### HasInternetState

`func (o *GatewayWanStatusEntity) HasInternetState() bool`

HasInternetState returns a boolean if a field has been set.

### GetInternetStatus

`func (o *GatewayWanStatusEntity) GetInternetStatus() int32`

GetInternetStatus returns the InternetStatus field if non-nil, zero value otherwise.

### GetInternetStatusOk

`func (o *GatewayWanStatusEntity) GetInternetStatusOk() (*int32, bool)`

GetInternetStatusOk returns a tuple with the InternetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetStatus

`func (o *GatewayWanStatusEntity) SetInternetStatus(v int32)`

SetInternetStatus sets InternetStatus field to given value.

### HasInternetStatus

`func (o *GatewayWanStatusEntity) HasInternetStatus() bool`

HasInternetStatus returns a boolean if a field has been set.

### GetIp

`func (o *GatewayWanStatusEntity) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayWanStatusEntity) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayWanStatusEntity) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GatewayWanStatusEntity) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp2

`func (o *GatewayWanStatusEntity) GetIp2() string`

GetIp2 returns the Ip2 field if non-nil, zero value otherwise.

### GetIp2Ok

`func (o *GatewayWanStatusEntity) GetIp2Ok() (*string, bool)`

GetIp2Ok returns a tuple with the Ip2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp2

`func (o *GatewayWanStatusEntity) SetIp2(v string)`

SetIp2 sets Ip2 field to given value.

### HasIp2

`func (o *GatewayWanStatusEntity) HasIp2() bool`

HasIp2 returns a boolean if a field has been set.

### GetIsp

`func (o *GatewayWanStatusEntity) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *GatewayWanStatusEntity) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *GatewayWanStatusEntity) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *GatewayWanStatusEntity) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetIspVersion

`func (o *GatewayWanStatusEntity) GetIspVersion() string`

GetIspVersion returns the IspVersion field if non-nil, zero value otherwise.

### GetIspVersionOk

`func (o *GatewayWanStatusEntity) GetIspVersionOk() (*string, bool)`

GetIspVersionOk returns a tuple with the IspVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspVersion

`func (o *GatewayWanStatusEntity) SetIspVersion(v string)`

SetIspVersion sets IspVersion field to given value.

### HasIspVersion

`func (o *GatewayWanStatusEntity) HasIspVersion() bool`

HasIspVersion returns a boolean if a field has been set.

### GetLatency

`func (o *GatewayWanStatusEntity) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *GatewayWanStatusEntity) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *GatewayWanStatusEntity) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *GatewayWanStatusEntity) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLineStatus

`func (o *GatewayWanStatusEntity) GetLineStatus() int32`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *GatewayWanStatusEntity) GetLineStatusOk() (*int32, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *GatewayWanStatusEntity) SetLineStatus(v int32)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *GatewayWanStatusEntity) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.

### GetLoss

`func (o *GatewayWanStatusEntity) GetLoss() float64`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *GatewayWanStatusEntity) GetLossOk() (*float64, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *GatewayWanStatusEntity) SetLoss(v float64)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *GatewayWanStatusEntity) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetMac

`func (o *GatewayWanStatusEntity) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GatewayWanStatusEntity) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GatewayWanStatusEntity) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GatewayWanStatusEntity) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMaxRxRate

`func (o *GatewayWanStatusEntity) GetMaxRxRate() int64`

GetMaxRxRate returns the MaxRxRate field if non-nil, zero value otherwise.

### GetMaxRxRateOk

`func (o *GatewayWanStatusEntity) GetMaxRxRateOk() (*int64, bool)`

GetMaxRxRateOk returns a tuple with the MaxRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRxRate

`func (o *GatewayWanStatusEntity) SetMaxRxRate(v int64)`

SetMaxRxRate sets MaxRxRate field to given value.

### HasMaxRxRate

`func (o *GatewayWanStatusEntity) HasMaxRxRate() bool`

HasMaxRxRate returns a boolean if a field has been set.

### GetMaxTxRate

`func (o *GatewayWanStatusEntity) GetMaxTxRate() int64`

GetMaxTxRate returns the MaxTxRate field if non-nil, zero value otherwise.

### GetMaxTxRateOk

`func (o *GatewayWanStatusEntity) GetMaxTxRateOk() (*int64, bool)`

GetMaxTxRateOk returns a tuple with the MaxTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTxRate

`func (o *GatewayWanStatusEntity) SetMaxTxRate(v int64)`

SetMaxTxRate sets MaxTxRate field to given value.

### HasMaxTxRate

`func (o *GatewayWanStatusEntity) HasMaxTxRate() bool`

HasMaxTxRate returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *GatewayWanStatusEntity) GetMirroredPorts() []MirroredPort`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *GatewayWanStatusEntity) GetMirroredPortsOk() (*[]MirroredPort, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *GatewayWanStatusEntity) SetMirroredPorts(v []MirroredPort)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *GatewayWanStatusEntity) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetMode

`func (o *GatewayWanStatusEntity) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GatewayWanStatusEntity) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GatewayWanStatusEntity) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GatewayWanStatusEntity) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GatewayWanStatusEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayWanStatusEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayWanStatusEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayWanStatusEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetType

`func (o *GatewayWanStatusEntity) GetNetType() int32`

GetNetType returns the NetType field if non-nil, zero value otherwise.

### GetNetTypeOk

`func (o *GatewayWanStatusEntity) GetNetTypeOk() (*int32, bool)`

GetNetTypeOk returns a tuple with the NetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetType

`func (o *GatewayWanStatusEntity) SetNetType(v int32)`

SetNetType sets NetType field to given value.

### HasNetType

`func (o *GatewayWanStatusEntity) HasNetType() bool`

HasNetType returns a boolean if a field has been set.

### GetOnlineDetection

`func (o *GatewayWanStatusEntity) GetOnlineDetection() int32`

GetOnlineDetection returns the OnlineDetection field if non-nil, zero value otherwise.

### GetOnlineDetectionOk

`func (o *GatewayWanStatusEntity) GetOnlineDetectionOk() (*int32, bool)`

GetOnlineDetectionOk returns a tuple with the OnlineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDetection

`func (o *GatewayWanStatusEntity) SetOnlineDetection(v int32)`

SetOnlineDetection sets OnlineDetection field to given value.

### HasOnlineDetection

`func (o *GatewayWanStatusEntity) HasOnlineDetection() bool`

HasOnlineDetection returns a boolean if a field has been set.

### GetPhysicalType

`func (o *GatewayWanStatusEntity) GetPhysicalType() int32`

GetPhysicalType returns the PhysicalType field if non-nil, zero value otherwise.

### GetPhysicalTypeOk

`func (o *GatewayWanStatusEntity) GetPhysicalTypeOk() (*int32, bool)`

GetPhysicalTypeOk returns a tuple with the PhysicalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalType

`func (o *GatewayWanStatusEntity) SetPhysicalType(v int32)`

SetPhysicalType sets PhysicalType field to given value.

### HasPhysicalType

`func (o *GatewayWanStatusEntity) HasPhysicalType() bool`

HasPhysicalType returns a boolean if a field has been set.

### GetPoe

`func (o *GatewayWanStatusEntity) GetPoe() bool`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *GatewayWanStatusEntity) GetPoeOk() (*bool, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *GatewayWanStatusEntity) SetPoe(v bool)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *GatewayWanStatusEntity) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPoePower

`func (o *GatewayWanStatusEntity) GetPoePower() float64`

GetPoePower returns the PoePower field if non-nil, zero value otherwise.

### GetPoePowerOk

`func (o *GatewayWanStatusEntity) GetPoePowerOk() (*float64, bool)`

GetPoePowerOk returns a tuple with the PoePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePower

`func (o *GatewayWanStatusEntity) SetPoePower(v float64)`

SetPoePower sets PoePower field to given value.

### HasPoePower

`func (o *GatewayWanStatusEntity) HasPoePower() bool`

HasPoePower returns a boolean if a field has been set.

### GetPort

`func (o *GatewayWanStatusEntity) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayWanStatusEntity) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayWanStatusEntity) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GatewayWanStatusEntity) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortDesc

`func (o *GatewayWanStatusEntity) GetPortDesc() string`

GetPortDesc returns the PortDesc field if non-nil, zero value otherwise.

### GetPortDescOk

`func (o *GatewayWanStatusEntity) GetPortDescOk() (*string, bool)`

GetPortDescOk returns a tuple with the PortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDesc

`func (o *GatewayWanStatusEntity) SetPortDesc(v string)`

SetPortDesc sets PortDesc field to given value.

### HasPortDesc

`func (o *GatewayWanStatusEntity) HasPortDesc() bool`

HasPortDesc returns a boolean if a field has been set.

### GetProto

`func (o *GatewayWanStatusEntity) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *GatewayWanStatusEntity) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *GatewayWanStatusEntity) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *GatewayWanStatusEntity) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetRoamingStatus

`func (o *GatewayWanStatusEntity) GetRoamingStatus() int32`

GetRoamingStatus returns the RoamingStatus field if non-nil, zero value otherwise.

### GetRoamingStatusOk

`func (o *GatewayWanStatusEntity) GetRoamingStatusOk() (*int32, bool)`

GetRoamingStatusOk returns a tuple with the RoamingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingStatus

`func (o *GatewayWanStatusEntity) SetRoamingStatus(v int32)`

SetRoamingStatus sets RoamingStatus field to given value.

### HasRoamingStatus

`func (o *GatewayWanStatusEntity) HasRoamingStatus() bool`

HasRoamingStatus returns a boolean if a field has been set.

### GetRsrp

`func (o *GatewayWanStatusEntity) GetRsrp() int32`

GetRsrp returns the Rsrp field if non-nil, zero value otherwise.

### GetRsrpOk

`func (o *GatewayWanStatusEntity) GetRsrpOk() (*int32, bool)`

GetRsrpOk returns a tuple with the Rsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsrp

`func (o *GatewayWanStatusEntity) SetRsrp(v int32)`

SetRsrp sets Rsrp field to given value.

### HasRsrp

`func (o *GatewayWanStatusEntity) HasRsrp() bool`

HasRsrp returns a boolean if a field has been set.

### GetRsrq

`func (o *GatewayWanStatusEntity) GetRsrq() int32`

GetRsrq returns the Rsrq field if non-nil, zero value otherwise.

### GetRsrqOk

`func (o *GatewayWanStatusEntity) GetRsrqOk() (*int32, bool)`

GetRsrqOk returns a tuple with the Rsrq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsrq

`func (o *GatewayWanStatusEntity) SetRsrq(v int32)`

SetRsrq sets Rsrq field to given value.

### HasRsrq

`func (o *GatewayWanStatusEntity) HasRsrq() bool`

HasRsrq returns a boolean if a field has been set.

### GetRx

`func (o *GatewayWanStatusEntity) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *GatewayWanStatusEntity) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *GatewayWanStatusEntity) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *GatewayWanStatusEntity) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxErrorPkts

`func (o *GatewayWanStatusEntity) GetRxErrorPkts() int64`

GetRxErrorPkts returns the RxErrorPkts field if non-nil, zero value otherwise.

### GetRxErrorPktsOk

`func (o *GatewayWanStatusEntity) GetRxErrorPktsOk() (*int64, bool)`

GetRxErrorPktsOk returns a tuple with the RxErrorPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrorPkts

`func (o *GatewayWanStatusEntity) SetRxErrorPkts(v int64)`

SetRxErrorPkts sets RxErrorPkts field to given value.

### HasRxErrorPkts

`func (o *GatewayWanStatusEntity) HasRxErrorPkts() bool`

HasRxErrorPkts returns a boolean if a field has been set.

### GetRxLineAttenuation

`func (o *GatewayWanStatusEntity) GetRxLineAttenuation() int64`

GetRxLineAttenuation returns the RxLineAttenuation field if non-nil, zero value otherwise.

### GetRxLineAttenuationOk

`func (o *GatewayWanStatusEntity) GetRxLineAttenuationOk() (*int64, bool)`

GetRxLineAttenuationOk returns a tuple with the RxLineAttenuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxLineAttenuation

`func (o *GatewayWanStatusEntity) SetRxLineAttenuation(v int64)`

SetRxLineAttenuation sets RxLineAttenuation field to given value.

### HasRxLineAttenuation

`func (o *GatewayWanStatusEntity) HasRxLineAttenuation() bool`

HasRxLineAttenuation returns a boolean if a field has been set.

### GetRxPkt

`func (o *GatewayWanStatusEntity) GetRxPkt() int64`

GetRxPkt returns the RxPkt field if non-nil, zero value otherwise.

### GetRxPktOk

`func (o *GatewayWanStatusEntity) GetRxPktOk() (*int64, bool)`

GetRxPktOk returns a tuple with the RxPkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkt

`func (o *GatewayWanStatusEntity) SetRxPkt(v int64)`

SetRxPkt sets RxPkt field to given value.

### HasRxPkt

`func (o *GatewayWanStatusEntity) HasRxPkt() bool`

HasRxPkt returns a boolean if a field has been set.

### GetRxPktRate

`func (o *GatewayWanStatusEntity) GetRxPktRate() int64`

GetRxPktRate returns the RxPktRate field if non-nil, zero value otherwise.

### GetRxPktRateOk

`func (o *GatewayWanStatusEntity) GetRxPktRateOk() (*int64, bool)`

GetRxPktRateOk returns a tuple with the RxPktRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPktRate

`func (o *GatewayWanStatusEntity) SetRxPktRate(v int64)`

SetRxPktRate sets RxPktRate field to given value.

### HasRxPktRate

`func (o *GatewayWanStatusEntity) HasRxPktRate() bool`

HasRxPktRate returns a boolean if a field has been set.

### GetRxRate

`func (o *GatewayWanStatusEntity) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *GatewayWanStatusEntity) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *GatewayWanStatusEntity) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *GatewayWanStatusEntity) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetRxSnrMargin

`func (o *GatewayWanStatusEntity) GetRxSnrMargin() int64`

GetRxSnrMargin returns the RxSnrMargin field if non-nil, zero value otherwise.

### GetRxSnrMarginOk

`func (o *GatewayWanStatusEntity) GetRxSnrMarginOk() (*int64, bool)`

GetRxSnrMarginOk returns a tuple with the RxSnrMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxSnrMargin

`func (o *GatewayWanStatusEntity) SetRxSnrMargin(v int64)`

SetRxSnrMargin sets RxSnrMargin field to given value.

### HasRxSnrMargin

`func (o *GatewayWanStatusEntity) HasRxSnrMargin() bool`

HasRxSnrMargin returns a boolean if a field has been set.

### GetSignal

`func (o *GatewayWanStatusEntity) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *GatewayWanStatusEntity) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *GatewayWanStatusEntity) SetSignal(v int32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *GatewayWanStatusEntity) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetSimCardUsed

`func (o *GatewayWanStatusEntity) GetSimCardUsed() int32`

GetSimCardUsed returns the SimCardUsed field if non-nil, zero value otherwise.

### GetSimCardUsedOk

`func (o *GatewayWanStatusEntity) GetSimCardUsedOk() (*int32, bool)`

GetSimCardUsedOk returns a tuple with the SimCardUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardUsed

`func (o *GatewayWanStatusEntity) SetSimCardUsed(v int32)`

SetSimCardUsed sets SimCardUsed field to given value.

### HasSimCardUsed

`func (o *GatewayWanStatusEntity) HasSimCardUsed() bool`

HasSimCardUsed returns a boolean if a field has been set.

### GetSmsOperator

`func (o *GatewayWanStatusEntity) GetSmsOperator() string`

GetSmsOperator returns the SmsOperator field if non-nil, zero value otherwise.

### GetSmsOperatorOk

`func (o *GatewayWanStatusEntity) GetSmsOperatorOk() (*string, bool)`

GetSmsOperatorOk returns a tuple with the SmsOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsOperator

`func (o *GatewayWanStatusEntity) SetSmsOperator(v string)`

SetSmsOperator sets SmsOperator field to given value.

### HasSmsOperator

`func (o *GatewayWanStatusEntity) HasSmsOperator() bool`

HasSmsOperator returns a boolean if a field has been set.

### GetSnr

`func (o *GatewayWanStatusEntity) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *GatewayWanStatusEntity) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *GatewayWanStatusEntity) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *GatewayWanStatusEntity) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetSpeed

`func (o *GatewayWanStatusEntity) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GatewayWanStatusEntity) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GatewayWanStatusEntity) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GatewayWanStatusEntity) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayWanStatusEntity) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayWanStatusEntity) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayWanStatusEntity) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GatewayWanStatusEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportSms

`func (o *GatewayWanStatusEntity) GetSupportSms() bool`

GetSupportSms returns the SupportSms field if non-nil, zero value otherwise.

### GetSupportSmsOk

`func (o *GatewayWanStatusEntity) GetSupportSmsOk() (*bool, bool)`

GetSupportSmsOk returns a tuple with the SupportSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSms

`func (o *GatewayWanStatusEntity) SetSupportSms(v bool)`

SetSupportSms sets SupportSms field to given value.

### HasSupportSms

`func (o *GatewayWanStatusEntity) HasSupportSms() bool`

HasSupportSms returns a boolean if a field has been set.

### GetTotalData

`func (o *GatewayWanStatusEntity) GetTotalData() float32`

GetTotalData returns the TotalData field if non-nil, zero value otherwise.

### GetTotalDataOk

`func (o *GatewayWanStatusEntity) GetTotalDataOk() (*float32, bool)`

GetTotalDataOk returns a tuple with the TotalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalData

`func (o *GatewayWanStatusEntity) SetTotalData(v float32)`

SetTotalData sets TotalData field to given value.

### HasTotalData

`func (o *GatewayWanStatusEntity) HasTotalData() bool`

HasTotalData returns a boolean if a field has been set.

### GetTrafficStatus

`func (o *GatewayWanStatusEntity) GetTrafficStatus() int32`

GetTrafficStatus returns the TrafficStatus field if non-nil, zero value otherwise.

### GetTrafficStatusOk

`func (o *GatewayWanStatusEntity) GetTrafficStatusOk() (*int32, bool)`

GetTrafficStatusOk returns a tuple with the TrafficStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficStatus

`func (o *GatewayWanStatusEntity) SetTrafficStatus(v int32)`

SetTrafficStatus sets TrafficStatus field to given value.

### HasTrafficStatus

`func (o *GatewayWanStatusEntity) HasTrafficStatus() bool`

HasTrafficStatus returns a boolean if a field has been set.

### GetTx

`func (o *GatewayWanStatusEntity) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *GatewayWanStatusEntity) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *GatewayWanStatusEntity) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *GatewayWanStatusEntity) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxErrorPkts

`func (o *GatewayWanStatusEntity) GetTxErrorPkts() int64`

GetTxErrorPkts returns the TxErrorPkts field if non-nil, zero value otherwise.

### GetTxErrorPktsOk

`func (o *GatewayWanStatusEntity) GetTxErrorPktsOk() (*int64, bool)`

GetTxErrorPktsOk returns a tuple with the TxErrorPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrorPkts

`func (o *GatewayWanStatusEntity) SetTxErrorPkts(v int64)`

SetTxErrorPkts sets TxErrorPkts field to given value.

### HasTxErrorPkts

`func (o *GatewayWanStatusEntity) HasTxErrorPkts() bool`

HasTxErrorPkts returns a boolean if a field has been set.

### GetTxLineAttenuation

`func (o *GatewayWanStatusEntity) GetTxLineAttenuation() int64`

GetTxLineAttenuation returns the TxLineAttenuation field if non-nil, zero value otherwise.

### GetTxLineAttenuationOk

`func (o *GatewayWanStatusEntity) GetTxLineAttenuationOk() (*int64, bool)`

GetTxLineAttenuationOk returns a tuple with the TxLineAttenuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxLineAttenuation

`func (o *GatewayWanStatusEntity) SetTxLineAttenuation(v int64)`

SetTxLineAttenuation sets TxLineAttenuation field to given value.

### HasTxLineAttenuation

`func (o *GatewayWanStatusEntity) HasTxLineAttenuation() bool`

HasTxLineAttenuation returns a boolean if a field has been set.

### GetTxPkt

`func (o *GatewayWanStatusEntity) GetTxPkt() int64`

GetTxPkt returns the TxPkt field if non-nil, zero value otherwise.

### GetTxPktOk

`func (o *GatewayWanStatusEntity) GetTxPktOk() (*int64, bool)`

GetTxPktOk returns a tuple with the TxPkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkt

`func (o *GatewayWanStatusEntity) SetTxPkt(v int64)`

SetTxPkt sets TxPkt field to given value.

### HasTxPkt

`func (o *GatewayWanStatusEntity) HasTxPkt() bool`

HasTxPkt returns a boolean if a field has been set.

### GetTxPktRate

`func (o *GatewayWanStatusEntity) GetTxPktRate() int64`

GetTxPktRate returns the TxPktRate field if non-nil, zero value otherwise.

### GetTxPktRateOk

`func (o *GatewayWanStatusEntity) GetTxPktRateOk() (*int64, bool)`

GetTxPktRateOk returns a tuple with the TxPktRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPktRate

`func (o *GatewayWanStatusEntity) SetTxPktRate(v int64)`

SetTxPktRate sets TxPktRate field to given value.

### HasTxPktRate

`func (o *GatewayWanStatusEntity) HasTxPktRate() bool`

HasTxPktRate returns a boolean if a field has been set.

### GetTxRate

`func (o *GatewayWanStatusEntity) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *GatewayWanStatusEntity) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *GatewayWanStatusEntity) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *GatewayWanStatusEntity) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetTxSnrMargin

`func (o *GatewayWanStatusEntity) GetTxSnrMargin() int64`

GetTxSnrMargin returns the TxSnrMargin field if non-nil, zero value otherwise.

### GetTxSnrMarginOk

`func (o *GatewayWanStatusEntity) GetTxSnrMarginOk() (*int64, bool)`

GetTxSnrMarginOk returns a tuple with the TxSnrMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSnrMargin

`func (o *GatewayWanStatusEntity) SetTxSnrMargin(v int64)`

SetTxSnrMargin sets TxSnrMargin field to given value.

### HasTxSnrMargin

`func (o *GatewayWanStatusEntity) HasTxSnrMargin() bool`

HasTxSnrMargin returns a boolean if a field has been set.

### GetType

`func (o *GatewayWanStatusEntity) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayWanStatusEntity) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayWanStatusEntity) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayWanStatusEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWanIpv6Comptent

`func (o *GatewayWanStatusEntity) GetWanIpv6Comptent() int32`

GetWanIpv6Comptent returns the WanIpv6Comptent field if non-nil, zero value otherwise.

### GetWanIpv6ComptentOk

`func (o *GatewayWanStatusEntity) GetWanIpv6ComptentOk() (*int32, bool)`

GetWanIpv6ComptentOk returns a tuple with the WanIpv6Comptent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIpv6Comptent

`func (o *GatewayWanStatusEntity) SetWanIpv6Comptent(v int32)`

SetWanIpv6Comptent sets WanIpv6Comptent field to given value.

### HasWanIpv6Comptent

`func (o *GatewayWanStatusEntity) HasWanIpv6Comptent() bool`

HasWanIpv6Comptent returns a boolean if a field has been set.

### GetWanPortIpv4Config

`func (o *GatewayWanStatusEntity) GetWanPortIpv4Config() OsgWanPortIpv4ConfigVO`

GetWanPortIpv4Config returns the WanPortIpv4Config field if non-nil, zero value otherwise.

### GetWanPortIpv4ConfigOk

`func (o *GatewayWanStatusEntity) GetWanPortIpv4ConfigOk() (*OsgWanPortIpv4ConfigVO, bool)`

GetWanPortIpv4ConfigOk returns a tuple with the WanPortIpv4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIpv4Config

`func (o *GatewayWanStatusEntity) SetWanPortIpv4Config(v OsgWanPortIpv4ConfigVO)`

SetWanPortIpv4Config sets WanPortIpv4Config field to given value.

### HasWanPortIpv4Config

`func (o *GatewayWanStatusEntity) HasWanPortIpv4Config() bool`

HasWanPortIpv4Config returns a boolean if a field has been set.

### GetWanPortIpv6Config

`func (o *GatewayWanStatusEntity) GetWanPortIpv6Config() OsgWanPortIpv6ConfigVO`

GetWanPortIpv6Config returns the WanPortIpv6Config field if non-nil, zero value otherwise.

### GetWanPortIpv6ConfigOk

`func (o *GatewayWanStatusEntity) GetWanPortIpv6ConfigOk() (*OsgWanPortIpv6ConfigVO, bool)`

GetWanPortIpv6ConfigOk returns a tuple with the WanPortIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIpv6Config

`func (o *GatewayWanStatusEntity) SetWanPortIpv6Config(v OsgWanPortIpv6ConfigVO)`

SetWanPortIpv6Config sets WanPortIpv6Config field to given value.

### HasWanPortIpv6Config

`func (o *GatewayWanStatusEntity) HasWanPortIpv6Config() bool`

HasWanPortIpv6Config returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


