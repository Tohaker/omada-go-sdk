# LanTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rx** | Pointer to **int64** | Total RX bytes, Unit: Byte | [optional] 
**RxDropPkts** | Pointer to **int64** | RX dropped packets | [optional] 
**RxErrPkts** | Pointer to **int64** | RX error packets | [optional] 
**RxPkts** | Pointer to **int64** | Total RX packets | [optional] 
**Tx** | Pointer to **int64** | Total TX bytes, Unit: Byte | [optional] 
**TxDropPkts** | Pointer to **int64** | TX dropped packets | [optional] 
**TxErrPkts** | Pointer to **int64** | TX error packets | [optional] 
**TxPkts** | Pointer to **int64** | Total TX packets | [optional] 

## Methods

### NewLanTraffic

`func NewLanTraffic() *LanTraffic`

NewLanTraffic instantiates a new LanTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanTrafficWithDefaults

`func NewLanTrafficWithDefaults() *LanTraffic`

NewLanTrafficWithDefaults instantiates a new LanTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRx

`func (o *LanTraffic) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *LanTraffic) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *LanTraffic) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *LanTraffic) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *LanTraffic) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *LanTraffic) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *LanTraffic) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *LanTraffic) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *LanTraffic) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *LanTraffic) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *LanTraffic) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *LanTraffic) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetRxPkts

`func (o *LanTraffic) GetRxPkts() int64`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *LanTraffic) GetRxPktsOk() (*int64, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *LanTraffic) SetRxPkts(v int64)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *LanTraffic) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetTx

`func (o *LanTraffic) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *LanTraffic) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *LanTraffic) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *LanTraffic) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *LanTraffic) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *LanTraffic) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *LanTraffic) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *LanTraffic) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *LanTraffic) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *LanTraffic) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *LanTraffic) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *LanTraffic) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetTxPkts

`func (o *LanTraffic) GetTxPkts() int64`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *LanTraffic) GetTxPktsOk() (*int64, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *LanTraffic) SetTxPkts(v int64)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *LanTraffic) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


