# APRadioTrafficEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rx** | Pointer to **int64** | Total RX bytes | [optional] 
**RxDropPkts** | Pointer to **int64** | RX dropped packets | [optional] 
**RxErrPkts** | Pointer to **int64** | RX error packets | [optional] 
**RxPkts** | Pointer to **int64** | Total RX packets | [optional] 
**RxRetryPkts** | Pointer to **int64** | RX retried packets | [optional] 
**Tx** | Pointer to **int64** | Total TX bytes | [optional] 
**TxDropPkts** | Pointer to **int64** | TX dropped Packets | [optional] 
**TxErrPkts** | Pointer to **int64** | TX error packets | [optional] 
**TxPkts** | Pointer to **int64** | Total TX packets | [optional] 
**TxRetryPkts** | Pointer to **int64** | TX retried packets | [optional] 

## Methods

### NewAPRadioTrafficEntity

`func NewAPRadioTrafficEntity() *APRadioTrafficEntity`

NewAPRadioTrafficEntity instantiates a new APRadioTrafficEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPRadioTrafficEntityWithDefaults

`func NewAPRadioTrafficEntityWithDefaults() *APRadioTrafficEntity`

NewAPRadioTrafficEntityWithDefaults instantiates a new APRadioTrafficEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRx

`func (o *APRadioTrafficEntity) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *APRadioTrafficEntity) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *APRadioTrafficEntity) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *APRadioTrafficEntity) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *APRadioTrafficEntity) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *APRadioTrafficEntity) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *APRadioTrafficEntity) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *APRadioTrafficEntity) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *APRadioTrafficEntity) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *APRadioTrafficEntity) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *APRadioTrafficEntity) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *APRadioTrafficEntity) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetRxPkts

`func (o *APRadioTrafficEntity) GetRxPkts() int64`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *APRadioTrafficEntity) GetRxPktsOk() (*int64, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *APRadioTrafficEntity) SetRxPkts(v int64)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *APRadioTrafficEntity) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetRxRetryPkts

`func (o *APRadioTrafficEntity) GetRxRetryPkts() int64`

GetRxRetryPkts returns the RxRetryPkts field if non-nil, zero value otherwise.

### GetRxRetryPktsOk

`func (o *APRadioTrafficEntity) GetRxRetryPktsOk() (*int64, bool)`

GetRxRetryPktsOk returns a tuple with the RxRetryPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRetryPkts

`func (o *APRadioTrafficEntity) SetRxRetryPkts(v int64)`

SetRxRetryPkts sets RxRetryPkts field to given value.

### HasRxRetryPkts

`func (o *APRadioTrafficEntity) HasRxRetryPkts() bool`

HasRxRetryPkts returns a boolean if a field has been set.

### GetTx

`func (o *APRadioTrafficEntity) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *APRadioTrafficEntity) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *APRadioTrafficEntity) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *APRadioTrafficEntity) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *APRadioTrafficEntity) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *APRadioTrafficEntity) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *APRadioTrafficEntity) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *APRadioTrafficEntity) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *APRadioTrafficEntity) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *APRadioTrafficEntity) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *APRadioTrafficEntity) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *APRadioTrafficEntity) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetTxPkts

`func (o *APRadioTrafficEntity) GetTxPkts() int64`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *APRadioTrafficEntity) GetTxPktsOk() (*int64, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *APRadioTrafficEntity) SetTxPkts(v int64)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *APRadioTrafficEntity) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetTxRetryPkts

`func (o *APRadioTrafficEntity) GetTxRetryPkts() int64`

GetTxRetryPkts returns the TxRetryPkts field if non-nil, zero value otherwise.

### GetTxRetryPktsOk

`func (o *APRadioTrafficEntity) GetTxRetryPktsOk() (*int64, bool)`

GetTxRetryPktsOk returns a tuple with the TxRetryPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRetryPkts

`func (o *APRadioTrafficEntity) SetTxRetryPkts(v int64)`

SetTxRetryPkts sets TxRetryPkts field to given value.

### HasTxRetryPkts

`func (o *APRadioTrafficEntity) HasTxRetryPkts() bool`

HasTxRetryPkts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


