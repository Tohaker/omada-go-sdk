# OltPortStatOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | Port of OLT | [optional] 
**Rx** | Pointer to **int64** | Receive traffic of the port, in bytes | [optional] 
**RxBroadcastPackets** | Pointer to **int64** | Receive broadcast packets | [optional] 
**RxMulticastPackets** | Pointer to **int64** | Receive multicast packets | [optional] 
**RxPackets** | Pointer to **int64** | Receive receive packets of the port | [optional] 
**RxRate** | Pointer to **float64** | Receive rate, byte/s | [optional] 
**Tx** | Pointer to **int64** | Transmit traffic of the port, in bytes | [optional] 
**TxBroadcastPackets** | Pointer to **int64** | Transmit broadcast packets | [optional] 
**TxMulticastPackets** | Pointer to **int64** | Transmit multicast packets | [optional] 
**TxPackets** | Pointer to **int64** | Transmit packets of the port | [optional] 
**TxRate** | Pointer to **float64** | Transmit rate, byte/s | [optional] 

## Methods

### NewOltPortStatOpenApiVO

`func NewOltPortStatOpenApiVO() *OltPortStatOpenApiVO`

NewOltPortStatOpenApiVO instantiates a new OltPortStatOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOltPortStatOpenApiVOWithDefaults

`func NewOltPortStatOpenApiVOWithDefaults() *OltPortStatOpenApiVO`

NewOltPortStatOpenApiVOWithDefaults instantiates a new OltPortStatOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *OltPortStatOpenApiVO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OltPortStatOpenApiVO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OltPortStatOpenApiVO) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *OltPortStatOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRx

`func (o *OltPortStatOpenApiVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OltPortStatOpenApiVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OltPortStatOpenApiVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OltPortStatOpenApiVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxBroadcastPackets

`func (o *OltPortStatOpenApiVO) GetRxBroadcastPackets() int64`

GetRxBroadcastPackets returns the RxBroadcastPackets field if non-nil, zero value otherwise.

### GetRxBroadcastPacketsOk

`func (o *OltPortStatOpenApiVO) GetRxBroadcastPacketsOk() (*int64, bool)`

GetRxBroadcastPacketsOk returns a tuple with the RxBroadcastPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBroadcastPackets

`func (o *OltPortStatOpenApiVO) SetRxBroadcastPackets(v int64)`

SetRxBroadcastPackets sets RxBroadcastPackets field to given value.

### HasRxBroadcastPackets

`func (o *OltPortStatOpenApiVO) HasRxBroadcastPackets() bool`

HasRxBroadcastPackets returns a boolean if a field has been set.

### GetRxMulticastPackets

`func (o *OltPortStatOpenApiVO) GetRxMulticastPackets() int64`

GetRxMulticastPackets returns the RxMulticastPackets field if non-nil, zero value otherwise.

### GetRxMulticastPacketsOk

`func (o *OltPortStatOpenApiVO) GetRxMulticastPacketsOk() (*int64, bool)`

GetRxMulticastPacketsOk returns a tuple with the RxMulticastPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMulticastPackets

`func (o *OltPortStatOpenApiVO) SetRxMulticastPackets(v int64)`

SetRxMulticastPackets sets RxMulticastPackets field to given value.

### HasRxMulticastPackets

`func (o *OltPortStatOpenApiVO) HasRxMulticastPackets() bool`

HasRxMulticastPackets returns a boolean if a field has been set.

### GetRxPackets

`func (o *OltPortStatOpenApiVO) GetRxPackets() int64`

GetRxPackets returns the RxPackets field if non-nil, zero value otherwise.

### GetRxPacketsOk

`func (o *OltPortStatOpenApiVO) GetRxPacketsOk() (*int64, bool)`

GetRxPacketsOk returns a tuple with the RxPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPackets

`func (o *OltPortStatOpenApiVO) SetRxPackets(v int64)`

SetRxPackets sets RxPackets field to given value.

### HasRxPackets

`func (o *OltPortStatOpenApiVO) HasRxPackets() bool`

HasRxPackets returns a boolean if a field has been set.

### GetRxRate

`func (o *OltPortStatOpenApiVO) GetRxRate() float64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OltPortStatOpenApiVO) GetRxRateOk() (*float64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OltPortStatOpenApiVO) SetRxRate(v float64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OltPortStatOpenApiVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetTx

`func (o *OltPortStatOpenApiVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OltPortStatOpenApiVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OltPortStatOpenApiVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OltPortStatOpenApiVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxBroadcastPackets

`func (o *OltPortStatOpenApiVO) GetTxBroadcastPackets() int64`

GetTxBroadcastPackets returns the TxBroadcastPackets field if non-nil, zero value otherwise.

### GetTxBroadcastPacketsOk

`func (o *OltPortStatOpenApiVO) GetTxBroadcastPacketsOk() (*int64, bool)`

GetTxBroadcastPacketsOk returns a tuple with the TxBroadcastPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBroadcastPackets

`func (o *OltPortStatOpenApiVO) SetTxBroadcastPackets(v int64)`

SetTxBroadcastPackets sets TxBroadcastPackets field to given value.

### HasTxBroadcastPackets

`func (o *OltPortStatOpenApiVO) HasTxBroadcastPackets() bool`

HasTxBroadcastPackets returns a boolean if a field has been set.

### GetTxMulticastPackets

`func (o *OltPortStatOpenApiVO) GetTxMulticastPackets() int64`

GetTxMulticastPackets returns the TxMulticastPackets field if non-nil, zero value otherwise.

### GetTxMulticastPacketsOk

`func (o *OltPortStatOpenApiVO) GetTxMulticastPacketsOk() (*int64, bool)`

GetTxMulticastPacketsOk returns a tuple with the TxMulticastPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMulticastPackets

`func (o *OltPortStatOpenApiVO) SetTxMulticastPackets(v int64)`

SetTxMulticastPackets sets TxMulticastPackets field to given value.

### HasTxMulticastPackets

`func (o *OltPortStatOpenApiVO) HasTxMulticastPackets() bool`

HasTxMulticastPackets returns a boolean if a field has been set.

### GetTxPackets

`func (o *OltPortStatOpenApiVO) GetTxPackets() int64`

GetTxPackets returns the TxPackets field if non-nil, zero value otherwise.

### GetTxPacketsOk

`func (o *OltPortStatOpenApiVO) GetTxPacketsOk() (*int64, bool)`

GetTxPacketsOk returns a tuple with the TxPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPackets

`func (o *OltPortStatOpenApiVO) SetTxPackets(v int64)`

SetTxPackets sets TxPackets field to given value.

### HasTxPackets

`func (o *OltPortStatOpenApiVO) HasTxPackets() bool`

HasTxPackets returns a boolean if a field has been set.

### GetTxRate

`func (o *OltPortStatOpenApiVO) GetTxRate() float64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OltPortStatOpenApiVO) GetTxRateOk() (*float64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OltPortStatOpenApiVO) SetTxRate(v float64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OltPortStatOpenApiVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


