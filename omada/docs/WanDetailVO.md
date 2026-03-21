# WanDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Latency** | Pointer to **int32** |  | [optional] 
**Loss** | Pointer to **float64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Proto** | Pointer to **int32** |  | [optional] 
**Rx** | Pointer to **int64** |  | [optional] 
**RxRate** | Pointer to **int64** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**Tx** | Pointer to **int64** |  | [optional] 
**TxRate** | Pointer to **int64** |  | [optional] 

## Methods

### NewWanDetailVO

`func NewWanDetailVO() *WanDetailVO`

NewWanDetailVO instantiates a new WanDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanDetailVOWithDefaults

`func NewWanDetailVOWithDefaults() *WanDetailVO`

NewWanDetailVOWithDefaults instantiates a new WanDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *WanDetailVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *WanDetailVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *WanDetailVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *WanDetailVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetIp

`func (o *WanDetailVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *WanDetailVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *WanDetailVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *WanDetailVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLatency

`func (o *WanDetailVO) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *WanDetailVO) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *WanDetailVO) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *WanDetailVO) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLoss

`func (o *WanDetailVO) GetLoss() float64`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *WanDetailVO) GetLossOk() (*float64, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *WanDetailVO) SetLoss(v float64)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *WanDetailVO) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetName

`func (o *WanDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WanDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WanDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WanDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *WanDetailVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WanDetailVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WanDetailVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *WanDetailVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProto

`func (o *WanDetailVO) GetProto() int32`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *WanDetailVO) GetProtoOk() (*int32, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *WanDetailVO) SetProto(v int32)`

SetProto sets Proto field to given value.

### HasProto

`func (o *WanDetailVO) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetRx

`func (o *WanDetailVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *WanDetailVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *WanDetailVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *WanDetailVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxRate

`func (o *WanDetailVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *WanDetailVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *WanDetailVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *WanDetailVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSpeed

`func (o *WanDetailVO) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *WanDetailVO) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *WanDetailVO) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *WanDetailVO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTx

`func (o *WanDetailVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *WanDetailVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *WanDetailVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *WanDetailVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxRate

`func (o *WanDetailVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *WanDetailVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *WanDetailVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *WanDetailVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


