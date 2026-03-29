# WirelessUpLinkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rssi** | Pointer to **int32** | Rssi | [optional] 
**RssiPercent** | Pointer to **int32** | Rssi percent | [optional] 
**RxDropPkts** | Pointer to **int64** | Rx Dropped Packets | [optional] 
**RxErrPkts** | Pointer to **int64** | Rx Error Packets | [optional] 
**RxRate** | Pointer to **string** | Rx Rate | [optional] 
**Snr** | Pointer to **int32** | Wireless P2P Ap Snr | [optional] 
**TxDropPkts** | Pointer to **int64** | Tx Dropped Packets | [optional] 
**TxErrPkts** | Pointer to **int64** | Tx Error Packets | [optional] 
**TxRate** | Pointer to **string** | Tx Rate | [optional] 

## Methods

### NewWirelessUpLinkInfo

`func NewWirelessUpLinkInfo() *WirelessUpLinkInfo`

NewWirelessUpLinkInfo instantiates a new WirelessUpLinkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessUpLinkInfoWithDefaults

`func NewWirelessUpLinkInfoWithDefaults() *WirelessUpLinkInfo`

NewWirelessUpLinkInfoWithDefaults instantiates a new WirelessUpLinkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRssi

`func (o *WirelessUpLinkInfo) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *WirelessUpLinkInfo) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *WirelessUpLinkInfo) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *WirelessUpLinkInfo) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRssiPercent

`func (o *WirelessUpLinkInfo) GetRssiPercent() int32`

GetRssiPercent returns the RssiPercent field if non-nil, zero value otherwise.

### GetRssiPercentOk

`func (o *WirelessUpLinkInfo) GetRssiPercentOk() (*int32, bool)`

GetRssiPercentOk returns a tuple with the RssiPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiPercent

`func (o *WirelessUpLinkInfo) SetRssiPercent(v int32)`

SetRssiPercent sets RssiPercent field to given value.

### HasRssiPercent

`func (o *WirelessUpLinkInfo) HasRssiPercent() bool`

HasRssiPercent returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *WirelessUpLinkInfo) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *WirelessUpLinkInfo) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *WirelessUpLinkInfo) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *WirelessUpLinkInfo) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *WirelessUpLinkInfo) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *WirelessUpLinkInfo) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *WirelessUpLinkInfo) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *WirelessUpLinkInfo) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetRxRate

`func (o *WirelessUpLinkInfo) GetRxRate() string`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *WirelessUpLinkInfo) GetRxRateOk() (*string, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *WirelessUpLinkInfo) SetRxRate(v string)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *WirelessUpLinkInfo) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSnr

`func (o *WirelessUpLinkInfo) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *WirelessUpLinkInfo) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *WirelessUpLinkInfo) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *WirelessUpLinkInfo) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *WirelessUpLinkInfo) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *WirelessUpLinkInfo) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *WirelessUpLinkInfo) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *WirelessUpLinkInfo) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *WirelessUpLinkInfo) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *WirelessUpLinkInfo) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *WirelessUpLinkInfo) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *WirelessUpLinkInfo) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetTxRate

`func (o *WirelessUpLinkInfo) GetTxRate() string`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *WirelessUpLinkInfo) GetTxRateOk() (*string, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *WirelessUpLinkInfo) SetTxRate(v string)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *WirelessUpLinkInfo) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


