# WirelessUpInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rssi** | Pointer to **int32** | Rssi | [optional] 
**RssiPercent** | Pointer to **float32** | Rssi Percent | [optional] 
**Rx** | Pointer to **int32** | Rx | [optional] 
**RxDropPkts** | Pointer to **int64** | Rx Dropped Packets | [optional] 
**RxErrPkts** | Pointer to **int64** | Rx Error Packets | [optional] 
**RxRate** | Pointer to **string** | Rx Rate | [optional] 
**Snr** | Pointer to **int32** | Wireless P2P Ap Snr | [optional] 
**Tx** | Pointer to **int32** | Tx | [optional] 
**TxDropPkts** | Pointer to **int64** | Tx Dropped Packets | [optional] 
**TxErrPkts** | Pointer to **int64** | Tx Error Packets | [optional] 
**TxRate** | Pointer to **string** | Tx Rate | [optional] 

## Methods

### NewWirelessUpInfoDTO

`func NewWirelessUpInfoDTO() *WirelessUpInfoDTO`

NewWirelessUpInfoDTO instantiates a new WirelessUpInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessUpInfoDTOWithDefaults

`func NewWirelessUpInfoDTOWithDefaults() *WirelessUpInfoDTO`

NewWirelessUpInfoDTOWithDefaults instantiates a new WirelessUpInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRssi

`func (o *WirelessUpInfoDTO) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *WirelessUpInfoDTO) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *WirelessUpInfoDTO) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *WirelessUpInfoDTO) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRssiPercent

`func (o *WirelessUpInfoDTO) GetRssiPercent() float32`

GetRssiPercent returns the RssiPercent field if non-nil, zero value otherwise.

### GetRssiPercentOk

`func (o *WirelessUpInfoDTO) GetRssiPercentOk() (*float32, bool)`

GetRssiPercentOk returns a tuple with the RssiPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiPercent

`func (o *WirelessUpInfoDTO) SetRssiPercent(v float32)`

SetRssiPercent sets RssiPercent field to given value.

### HasRssiPercent

`func (o *WirelessUpInfoDTO) HasRssiPercent() bool`

HasRssiPercent returns a boolean if a field has been set.

### GetRx

`func (o *WirelessUpInfoDTO) GetRx() int32`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *WirelessUpInfoDTO) GetRxOk() (*int32, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *WirelessUpInfoDTO) SetRx(v int32)`

SetRx sets Rx field to given value.

### HasRx

`func (o *WirelessUpInfoDTO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *WirelessUpInfoDTO) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *WirelessUpInfoDTO) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *WirelessUpInfoDTO) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *WirelessUpInfoDTO) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *WirelessUpInfoDTO) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *WirelessUpInfoDTO) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *WirelessUpInfoDTO) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *WirelessUpInfoDTO) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetRxRate

`func (o *WirelessUpInfoDTO) GetRxRate() string`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *WirelessUpInfoDTO) GetRxRateOk() (*string, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *WirelessUpInfoDTO) SetRxRate(v string)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *WirelessUpInfoDTO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSnr

`func (o *WirelessUpInfoDTO) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *WirelessUpInfoDTO) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *WirelessUpInfoDTO) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *WirelessUpInfoDTO) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetTx

`func (o *WirelessUpInfoDTO) GetTx() int32`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *WirelessUpInfoDTO) GetTxOk() (*int32, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *WirelessUpInfoDTO) SetTx(v int32)`

SetTx sets Tx field to given value.

### HasTx

`func (o *WirelessUpInfoDTO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *WirelessUpInfoDTO) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *WirelessUpInfoDTO) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *WirelessUpInfoDTO) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *WirelessUpInfoDTO) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *WirelessUpInfoDTO) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *WirelessUpInfoDTO) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *WirelessUpInfoDTO) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *WirelessUpInfoDTO) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetTxRate

`func (o *WirelessUpInfoDTO) GetTxRate() string`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *WirelessUpInfoDTO) GetTxRateOk() (*string, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *WirelessUpInfoDTO) SetTxRate(v string)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *WirelessUpInfoDTO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


