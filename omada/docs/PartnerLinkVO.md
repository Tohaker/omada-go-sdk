# PartnerLinkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **int64** | Uplink AP activity: (change of(downBytes+upBytes))/time | [optional] 
**Channel** | Pointer to **int32** | Uplink AP channel | [optional] 
**DownBytes** | Pointer to **int64** | Uplink AP downBytes; Unit:Byte | [optional] 
**DownPackets** | Pointer to **int64** | Uplink AP downPackets | [optional] 
**DownRate** | Pointer to **int64** | Uplink AP downRate: the rxRate calculate by Controller | [optional] 
**RadioId** | Pointer to **int32** | radioId, 0:2.4G, 1:5G, 2:5G2, 3:6G | [optional] 
**Rssi** | Pointer to **int32** | Uplink AP rssi | [optional] 
**RxRate** | Pointer to **string** | Uplink AP rxRate | [optional] 
**RxRateInt** | Pointer to **int32** | Uplink AP rxRateInt; Unit: Mbps | [optional] 
**Snr** | Pointer to **int32** | Uplink AP Signal-noise ratio | [optional] 
**TxRate** | Pointer to **string** | Uplink AP txRate | [optional] 
**TxRateInt** | Pointer to **int32** | Uplink AP txRateInt; Unit: Mbps | [optional] 
**UpBytes** | Pointer to **int64** | Uplink AP upBytes; Unit: Byte | [optional] 
**UpPackets** | Pointer to **int64** | Uplink AP upPackets | [optional] 
**UpRate** | Pointer to **int64** | Uplink AP upRate: the txRate calculate by Controller | [optional] 

## Methods

### NewPartnerLinkVO

`func NewPartnerLinkVO() *PartnerLinkVO`

NewPartnerLinkVO instantiates a new PartnerLinkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerLinkVOWithDefaults

`func NewPartnerLinkVOWithDefaults() *PartnerLinkVO`

NewPartnerLinkVOWithDefaults instantiates a new PartnerLinkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *PartnerLinkVO) GetActivity() int64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *PartnerLinkVO) GetActivityOk() (*int64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *PartnerLinkVO) SetActivity(v int64)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *PartnerLinkVO) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetChannel

`func (o *PartnerLinkVO) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PartnerLinkVO) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PartnerLinkVO) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PartnerLinkVO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDownBytes

`func (o *PartnerLinkVO) GetDownBytes() int64`

GetDownBytes returns the DownBytes field if non-nil, zero value otherwise.

### GetDownBytesOk

`func (o *PartnerLinkVO) GetDownBytesOk() (*int64, bool)`

GetDownBytesOk returns a tuple with the DownBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBytes

`func (o *PartnerLinkVO) SetDownBytes(v int64)`

SetDownBytes sets DownBytes field to given value.

### HasDownBytes

`func (o *PartnerLinkVO) HasDownBytes() bool`

HasDownBytes returns a boolean if a field has been set.

### GetDownPackets

`func (o *PartnerLinkVO) GetDownPackets() int64`

GetDownPackets returns the DownPackets field if non-nil, zero value otherwise.

### GetDownPacketsOk

`func (o *PartnerLinkVO) GetDownPacketsOk() (*int64, bool)`

GetDownPacketsOk returns a tuple with the DownPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPackets

`func (o *PartnerLinkVO) SetDownPackets(v int64)`

SetDownPackets sets DownPackets field to given value.

### HasDownPackets

`func (o *PartnerLinkVO) HasDownPackets() bool`

HasDownPackets returns a boolean if a field has been set.

### GetDownRate

`func (o *PartnerLinkVO) GetDownRate() int64`

GetDownRate returns the DownRate field if non-nil, zero value otherwise.

### GetDownRateOk

`func (o *PartnerLinkVO) GetDownRateOk() (*int64, bool)`

GetDownRateOk returns a tuple with the DownRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownRate

`func (o *PartnerLinkVO) SetDownRate(v int64)`

SetDownRate sets DownRate field to given value.

### HasDownRate

`func (o *PartnerLinkVO) HasDownRate() bool`

HasDownRate returns a boolean if a field has been set.

### GetRadioId

`func (o *PartnerLinkVO) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *PartnerLinkVO) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *PartnerLinkVO) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *PartnerLinkVO) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetRssi

`func (o *PartnerLinkVO) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *PartnerLinkVO) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *PartnerLinkVO) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *PartnerLinkVO) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRxRate

`func (o *PartnerLinkVO) GetRxRate() string`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *PartnerLinkVO) GetRxRateOk() (*string, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *PartnerLinkVO) SetRxRate(v string)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *PartnerLinkVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetRxRateInt

`func (o *PartnerLinkVO) GetRxRateInt() int32`

GetRxRateInt returns the RxRateInt field if non-nil, zero value otherwise.

### GetRxRateIntOk

`func (o *PartnerLinkVO) GetRxRateIntOk() (*int32, bool)`

GetRxRateIntOk returns a tuple with the RxRateInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRateInt

`func (o *PartnerLinkVO) SetRxRateInt(v int32)`

SetRxRateInt sets RxRateInt field to given value.

### HasRxRateInt

`func (o *PartnerLinkVO) HasRxRateInt() bool`

HasRxRateInt returns a boolean if a field has been set.

### GetSnr

`func (o *PartnerLinkVO) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *PartnerLinkVO) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *PartnerLinkVO) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *PartnerLinkVO) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetTxRate

`func (o *PartnerLinkVO) GetTxRate() string`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *PartnerLinkVO) GetTxRateOk() (*string, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *PartnerLinkVO) SetTxRate(v string)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *PartnerLinkVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetTxRateInt

`func (o *PartnerLinkVO) GetTxRateInt() int32`

GetTxRateInt returns the TxRateInt field if non-nil, zero value otherwise.

### GetTxRateIntOk

`func (o *PartnerLinkVO) GetTxRateIntOk() (*int32, bool)`

GetTxRateIntOk returns a tuple with the TxRateInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRateInt

`func (o *PartnerLinkVO) SetTxRateInt(v int32)`

SetTxRateInt sets TxRateInt field to given value.

### HasTxRateInt

`func (o *PartnerLinkVO) HasTxRateInt() bool`

HasTxRateInt returns a boolean if a field has been set.

### GetUpBytes

`func (o *PartnerLinkVO) GetUpBytes() int64`

GetUpBytes returns the UpBytes field if non-nil, zero value otherwise.

### GetUpBytesOk

`func (o *PartnerLinkVO) GetUpBytesOk() (*int64, bool)`

GetUpBytesOk returns a tuple with the UpBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBytes

`func (o *PartnerLinkVO) SetUpBytes(v int64)`

SetUpBytes sets UpBytes field to given value.

### HasUpBytes

`func (o *PartnerLinkVO) HasUpBytes() bool`

HasUpBytes returns a boolean if a field has been set.

### GetUpPackets

`func (o *PartnerLinkVO) GetUpPackets() int64`

GetUpPackets returns the UpPackets field if non-nil, zero value otherwise.

### GetUpPacketsOk

`func (o *PartnerLinkVO) GetUpPacketsOk() (*int64, bool)`

GetUpPacketsOk returns a tuple with the UpPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPackets

`func (o *PartnerLinkVO) SetUpPackets(v int64)`

SetUpPackets sets UpPackets field to given value.

### HasUpPackets

`func (o *PartnerLinkVO) HasUpPackets() bool`

HasUpPackets returns a boolean if a field has been set.

### GetUpRate

`func (o *PartnerLinkVO) GetUpRate() int64`

GetUpRate returns the UpRate field if non-nil, zero value otherwise.

### GetUpRateOk

`func (o *PartnerLinkVO) GetUpRateOk() (*int64, bool)`

GetUpRateOk returns a tuple with the UpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpRate

`func (o *PartnerLinkVO) SetUpRate(v int64)`

SetUpRate sets UpRate field to given value.

### HasUpRate

`func (o *PartnerLinkVO) HasUpRate() bool`

HasUpRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


