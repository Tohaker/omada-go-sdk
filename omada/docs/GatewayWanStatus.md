# GatewayWanStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Port duplex, 1-Half，2-Full | [optional] 
**HealthLevel** | Pointer to **int32** | Wan health level. 0-GOOD, 1-FAIR, 2-POOR, 3-NO_DATA, 4-OFFLINE, 5-DISABLE, 6-ONLINE | [optional] 
**InternetState** | Pointer to **int32** | Wan internet state should be a value as follows: 0: disconnected; 1: connected | [optional] 
**Ip** | Pointer to **string** | WAN ip address. | [optional] 
**Ip2** | Pointer to **string** | Secondary ip address. | [optional] 
**Latency** | Pointer to **int32** | Wan latency, when mode is wan and device is connected, Unit: ms | [optional] 
**Loss** | Pointer to **float64** | Wan packet loss rate, Unit : % | [optional] 
**Mode** | Pointer to **int32** | Port mode, 0:WAN,1:LAN; | [optional] 
**Name** | Pointer to **string** | Port name | [optional] 
**OnlineDetection** | Pointer to **int32** | online status, 0-offline, 1-online | [optional] 
**PhysicalType** | Pointer to **int32** | SFP port type, 0:normal,1:SFP,2:SFP+; | [optional] 
**Port** | Pointer to **int32** | Port serial number | [optional] 
**Proto** | Pointer to **string** | Wan ipv4 proto type, use static，dhcp，pppoe，l2tp，pptp. | [optional] 
**Rx** | Pointer to **int64** | Port total rx bytes | [optional] 
**RxErrorPkts** | Pointer to **int64** | rx error pkts. | [optional] 
**RxPkt** | Pointer to **int64** | Port total rx packets | [optional] 
**RxPktRate** | Pointer to **int64** | Port rx Packet rate, Unit: Pkt/s; | [optional] 
**RxRate** | Pointer to **int64** | Port rx rate, Unit: KB/s; | [optional] 
**Speed** | Pointer to **int32** | Port speed, 1-10M，2-100M，3-1000M | [optional] 
**Status** | Pointer to **int32** | Port status should be a value as follows: 0: disconnected; 1: connected | [optional] 
**Tx** | Pointer to **int64** | Port total tx bytes | [optional] 
**TxErrorPkts** | Pointer to **int64** | tx error pkts. | [optional] 
**TxPkt** | Pointer to **int64** | Port total tx packets | [optional] 
**TxPktRate** | Pointer to **int64** | Port tx packet rate, Unit: Pkt/s; | [optional] 
**TxRate** | Pointer to **int64** | Port tx rate, Unit: KB/s; | [optional] 
**Type** | Pointer to **int32** | Port type, 0:WAN,1:WAN/LAN,2:LAN,3:SFP WAN; | [optional] 

## Methods

### NewGatewayWanStatus

`func NewGatewayWanStatus() *GatewayWanStatus`

NewGatewayWanStatus instantiates a new GatewayWanStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWanStatusWithDefaults

`func NewGatewayWanStatusWithDefaults() *GatewayWanStatus`

NewGatewayWanStatusWithDefaults instantiates a new GatewayWanStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *GatewayWanStatus) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *GatewayWanStatus) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *GatewayWanStatus) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *GatewayWanStatus) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetHealthLevel

`func (o *GatewayWanStatus) GetHealthLevel() int32`

GetHealthLevel returns the HealthLevel field if non-nil, zero value otherwise.

### GetHealthLevelOk

`func (o *GatewayWanStatus) GetHealthLevelOk() (*int32, bool)`

GetHealthLevelOk returns a tuple with the HealthLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthLevel

`func (o *GatewayWanStatus) SetHealthLevel(v int32)`

SetHealthLevel sets HealthLevel field to given value.

### HasHealthLevel

`func (o *GatewayWanStatus) HasHealthLevel() bool`

HasHealthLevel returns a boolean if a field has been set.

### GetInternetState

`func (o *GatewayWanStatus) GetInternetState() int32`

GetInternetState returns the InternetState field if non-nil, zero value otherwise.

### GetInternetStateOk

`func (o *GatewayWanStatus) GetInternetStateOk() (*int32, bool)`

GetInternetStateOk returns a tuple with the InternetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetState

`func (o *GatewayWanStatus) SetInternetState(v int32)`

SetInternetState sets InternetState field to given value.

### HasInternetState

`func (o *GatewayWanStatus) HasInternetState() bool`

HasInternetState returns a boolean if a field has been set.

### GetIp

`func (o *GatewayWanStatus) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayWanStatus) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayWanStatus) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GatewayWanStatus) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp2

`func (o *GatewayWanStatus) GetIp2() string`

GetIp2 returns the Ip2 field if non-nil, zero value otherwise.

### GetIp2Ok

`func (o *GatewayWanStatus) GetIp2Ok() (*string, bool)`

GetIp2Ok returns a tuple with the Ip2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp2

`func (o *GatewayWanStatus) SetIp2(v string)`

SetIp2 sets Ip2 field to given value.

### HasIp2

`func (o *GatewayWanStatus) HasIp2() bool`

HasIp2 returns a boolean if a field has been set.

### GetLatency

`func (o *GatewayWanStatus) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *GatewayWanStatus) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *GatewayWanStatus) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *GatewayWanStatus) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLoss

`func (o *GatewayWanStatus) GetLoss() float64`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *GatewayWanStatus) GetLossOk() (*float64, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *GatewayWanStatus) SetLoss(v float64)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *GatewayWanStatus) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetMode

`func (o *GatewayWanStatus) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GatewayWanStatus) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GatewayWanStatus) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GatewayWanStatus) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GatewayWanStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayWanStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayWanStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayWanStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnlineDetection

`func (o *GatewayWanStatus) GetOnlineDetection() int32`

GetOnlineDetection returns the OnlineDetection field if non-nil, zero value otherwise.

### GetOnlineDetectionOk

`func (o *GatewayWanStatus) GetOnlineDetectionOk() (*int32, bool)`

GetOnlineDetectionOk returns a tuple with the OnlineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDetection

`func (o *GatewayWanStatus) SetOnlineDetection(v int32)`

SetOnlineDetection sets OnlineDetection field to given value.

### HasOnlineDetection

`func (o *GatewayWanStatus) HasOnlineDetection() bool`

HasOnlineDetection returns a boolean if a field has been set.

### GetPhysicalType

`func (o *GatewayWanStatus) GetPhysicalType() int32`

GetPhysicalType returns the PhysicalType field if non-nil, zero value otherwise.

### GetPhysicalTypeOk

`func (o *GatewayWanStatus) GetPhysicalTypeOk() (*int32, bool)`

GetPhysicalTypeOk returns a tuple with the PhysicalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalType

`func (o *GatewayWanStatus) SetPhysicalType(v int32)`

SetPhysicalType sets PhysicalType field to given value.

### HasPhysicalType

`func (o *GatewayWanStatus) HasPhysicalType() bool`

HasPhysicalType returns a boolean if a field has been set.

### GetPort

`func (o *GatewayWanStatus) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayWanStatus) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayWanStatus) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GatewayWanStatus) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProto

`func (o *GatewayWanStatus) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *GatewayWanStatus) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *GatewayWanStatus) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *GatewayWanStatus) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetRx

`func (o *GatewayWanStatus) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *GatewayWanStatus) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *GatewayWanStatus) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *GatewayWanStatus) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxErrorPkts

`func (o *GatewayWanStatus) GetRxErrorPkts() int64`

GetRxErrorPkts returns the RxErrorPkts field if non-nil, zero value otherwise.

### GetRxErrorPktsOk

`func (o *GatewayWanStatus) GetRxErrorPktsOk() (*int64, bool)`

GetRxErrorPktsOk returns a tuple with the RxErrorPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrorPkts

`func (o *GatewayWanStatus) SetRxErrorPkts(v int64)`

SetRxErrorPkts sets RxErrorPkts field to given value.

### HasRxErrorPkts

`func (o *GatewayWanStatus) HasRxErrorPkts() bool`

HasRxErrorPkts returns a boolean if a field has been set.

### GetRxPkt

`func (o *GatewayWanStatus) GetRxPkt() int64`

GetRxPkt returns the RxPkt field if non-nil, zero value otherwise.

### GetRxPktOk

`func (o *GatewayWanStatus) GetRxPktOk() (*int64, bool)`

GetRxPktOk returns a tuple with the RxPkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkt

`func (o *GatewayWanStatus) SetRxPkt(v int64)`

SetRxPkt sets RxPkt field to given value.

### HasRxPkt

`func (o *GatewayWanStatus) HasRxPkt() bool`

HasRxPkt returns a boolean if a field has been set.

### GetRxPktRate

`func (o *GatewayWanStatus) GetRxPktRate() int64`

GetRxPktRate returns the RxPktRate field if non-nil, zero value otherwise.

### GetRxPktRateOk

`func (o *GatewayWanStatus) GetRxPktRateOk() (*int64, bool)`

GetRxPktRateOk returns a tuple with the RxPktRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPktRate

`func (o *GatewayWanStatus) SetRxPktRate(v int64)`

SetRxPktRate sets RxPktRate field to given value.

### HasRxPktRate

`func (o *GatewayWanStatus) HasRxPktRate() bool`

HasRxPktRate returns a boolean if a field has been set.

### GetRxRate

`func (o *GatewayWanStatus) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *GatewayWanStatus) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *GatewayWanStatus) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *GatewayWanStatus) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSpeed

`func (o *GatewayWanStatus) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GatewayWanStatus) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GatewayWanStatus) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GatewayWanStatus) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayWanStatus) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayWanStatus) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayWanStatus) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GatewayWanStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTx

`func (o *GatewayWanStatus) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *GatewayWanStatus) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *GatewayWanStatus) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *GatewayWanStatus) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxErrorPkts

`func (o *GatewayWanStatus) GetTxErrorPkts() int64`

GetTxErrorPkts returns the TxErrorPkts field if non-nil, zero value otherwise.

### GetTxErrorPktsOk

`func (o *GatewayWanStatus) GetTxErrorPktsOk() (*int64, bool)`

GetTxErrorPktsOk returns a tuple with the TxErrorPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrorPkts

`func (o *GatewayWanStatus) SetTxErrorPkts(v int64)`

SetTxErrorPkts sets TxErrorPkts field to given value.

### HasTxErrorPkts

`func (o *GatewayWanStatus) HasTxErrorPkts() bool`

HasTxErrorPkts returns a boolean if a field has been set.

### GetTxPkt

`func (o *GatewayWanStatus) GetTxPkt() int64`

GetTxPkt returns the TxPkt field if non-nil, zero value otherwise.

### GetTxPktOk

`func (o *GatewayWanStatus) GetTxPktOk() (*int64, bool)`

GetTxPktOk returns a tuple with the TxPkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkt

`func (o *GatewayWanStatus) SetTxPkt(v int64)`

SetTxPkt sets TxPkt field to given value.

### HasTxPkt

`func (o *GatewayWanStatus) HasTxPkt() bool`

HasTxPkt returns a boolean if a field has been set.

### GetTxPktRate

`func (o *GatewayWanStatus) GetTxPktRate() int64`

GetTxPktRate returns the TxPktRate field if non-nil, zero value otherwise.

### GetTxPktRateOk

`func (o *GatewayWanStatus) GetTxPktRateOk() (*int64, bool)`

GetTxPktRateOk returns a tuple with the TxPktRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPktRate

`func (o *GatewayWanStatus) SetTxPktRate(v int64)`

SetTxPktRate sets TxPktRate field to given value.

### HasTxPktRate

`func (o *GatewayWanStatus) HasTxPktRate() bool`

HasTxPktRate returns a boolean if a field has been set.

### GetTxRate

`func (o *GatewayWanStatus) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *GatewayWanStatus) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *GatewayWanStatus) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *GatewayWanStatus) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetType

`func (o *GatewayWanStatus) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayWanStatus) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayWanStatus) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayWanStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


