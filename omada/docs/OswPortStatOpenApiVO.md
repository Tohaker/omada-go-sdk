# OswPortStatOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropPkts** | Pointer to **int64** | Drop packets | [optional] 
**Port** | Pointer to **int32** | Port of Switch | [optional] 
**Rx** | Pointer to **int64** | Receive traffic of the port, in bytes | [optional] 
**RxBroadPkts** | Pointer to **int64** | Receive broadcast packets | [optional] 
**RxErrPkts** | Pointer to **int64** | Receive error packets | [optional] 
**RxMultiPkts** | Pointer to **int64** | Receive multicast packets | [optional] 
**RxPkts** | Pointer to **int64** | Receive receive packets of the port | [optional] 
**RxRate** | Pointer to **int64** | Receive rate, byte/s | [optional] 
**Tx** | Pointer to **int64** | Transmit traffic of the port, in bytes | [optional] 
**TxBroadPkts** | Pointer to **int64** | Transmit broadcast packets | [optional] 
**TxErrPkts** | Pointer to **int64** | Transmit error packets | [optional] 
**TxMultiPkts** | Pointer to **int64** | Transmit multicast packets | [optional] 
**TxPkts** | Pointer to **int64** | Transmit packets of the port | [optional] 
**TxRate** | Pointer to **int64** | Transmit rate, byte/s | [optional] 

## Methods

### NewOswPortStatOpenApiVO

`func NewOswPortStatOpenApiVO() *OswPortStatOpenApiVO`

NewOswPortStatOpenApiVO instantiates a new OswPortStatOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortStatOpenApiVOWithDefaults

`func NewOswPortStatOpenApiVOWithDefaults() *OswPortStatOpenApiVO`

NewOswPortStatOpenApiVOWithDefaults instantiates a new OswPortStatOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropPkts

`func (o *OswPortStatOpenApiVO) GetDropPkts() int64`

GetDropPkts returns the DropPkts field if non-nil, zero value otherwise.

### GetDropPktsOk

`func (o *OswPortStatOpenApiVO) GetDropPktsOk() (*int64, bool)`

GetDropPktsOk returns a tuple with the DropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropPkts

`func (o *OswPortStatOpenApiVO) SetDropPkts(v int64)`

SetDropPkts sets DropPkts field to given value.

### HasDropPkts

`func (o *OswPortStatOpenApiVO) HasDropPkts() bool`

HasDropPkts returns a boolean if a field has been set.

### GetPort

`func (o *OswPortStatOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswPortStatOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswPortStatOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswPortStatOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRx

`func (o *OswPortStatOpenApiVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OswPortStatOpenApiVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OswPortStatOpenApiVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OswPortStatOpenApiVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxBroadPkts

`func (o *OswPortStatOpenApiVO) GetRxBroadPkts() int64`

GetRxBroadPkts returns the RxBroadPkts field if non-nil, zero value otherwise.

### GetRxBroadPktsOk

`func (o *OswPortStatOpenApiVO) GetRxBroadPktsOk() (*int64, bool)`

GetRxBroadPktsOk returns a tuple with the RxBroadPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBroadPkts

`func (o *OswPortStatOpenApiVO) SetRxBroadPkts(v int64)`

SetRxBroadPkts sets RxBroadPkts field to given value.

### HasRxBroadPkts

`func (o *OswPortStatOpenApiVO) HasRxBroadPkts() bool`

HasRxBroadPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *OswPortStatOpenApiVO) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *OswPortStatOpenApiVO) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *OswPortStatOpenApiVO) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *OswPortStatOpenApiVO) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetRxMultiPkts

`func (o *OswPortStatOpenApiVO) GetRxMultiPkts() int64`

GetRxMultiPkts returns the RxMultiPkts field if non-nil, zero value otherwise.

### GetRxMultiPktsOk

`func (o *OswPortStatOpenApiVO) GetRxMultiPktsOk() (*int64, bool)`

GetRxMultiPktsOk returns a tuple with the RxMultiPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMultiPkts

`func (o *OswPortStatOpenApiVO) SetRxMultiPkts(v int64)`

SetRxMultiPkts sets RxMultiPkts field to given value.

### HasRxMultiPkts

`func (o *OswPortStatOpenApiVO) HasRxMultiPkts() bool`

HasRxMultiPkts returns a boolean if a field has been set.

### GetRxPkts

`func (o *OswPortStatOpenApiVO) GetRxPkts() int64`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *OswPortStatOpenApiVO) GetRxPktsOk() (*int64, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *OswPortStatOpenApiVO) SetRxPkts(v int64)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *OswPortStatOpenApiVO) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetRxRate

`func (o *OswPortStatOpenApiVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswPortStatOpenApiVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswPortStatOpenApiVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswPortStatOpenApiVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetTx

`func (o *OswPortStatOpenApiVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OswPortStatOpenApiVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OswPortStatOpenApiVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OswPortStatOpenApiVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxBroadPkts

`func (o *OswPortStatOpenApiVO) GetTxBroadPkts() int64`

GetTxBroadPkts returns the TxBroadPkts field if non-nil, zero value otherwise.

### GetTxBroadPktsOk

`func (o *OswPortStatOpenApiVO) GetTxBroadPktsOk() (*int64, bool)`

GetTxBroadPktsOk returns a tuple with the TxBroadPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBroadPkts

`func (o *OswPortStatOpenApiVO) SetTxBroadPkts(v int64)`

SetTxBroadPkts sets TxBroadPkts field to given value.

### HasTxBroadPkts

`func (o *OswPortStatOpenApiVO) HasTxBroadPkts() bool`

HasTxBroadPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *OswPortStatOpenApiVO) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *OswPortStatOpenApiVO) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *OswPortStatOpenApiVO) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *OswPortStatOpenApiVO) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetTxMultiPkts

`func (o *OswPortStatOpenApiVO) GetTxMultiPkts() int64`

GetTxMultiPkts returns the TxMultiPkts field if non-nil, zero value otherwise.

### GetTxMultiPktsOk

`func (o *OswPortStatOpenApiVO) GetTxMultiPktsOk() (*int64, bool)`

GetTxMultiPktsOk returns a tuple with the TxMultiPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMultiPkts

`func (o *OswPortStatOpenApiVO) SetTxMultiPkts(v int64)`

SetTxMultiPkts sets TxMultiPkts field to given value.

### HasTxMultiPkts

`func (o *OswPortStatOpenApiVO) HasTxMultiPkts() bool`

HasTxMultiPkts returns a boolean if a field has been set.

### GetTxPkts

`func (o *OswPortStatOpenApiVO) GetTxPkts() int64`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *OswPortStatOpenApiVO) GetTxPktsOk() (*int64, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *OswPortStatOpenApiVO) SetTxPkts(v int64)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *OswPortStatOpenApiVO) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetTxRate

`func (o *OswPortStatOpenApiVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswPortStatOpenApiVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswPortStatOpenApiVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswPortStatOpenApiVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


