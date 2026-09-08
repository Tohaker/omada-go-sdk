# PartnerLinkDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** |  | [optional] 
**DownRate** | Pointer to **int64** |  | [optional] 
**RadioId** | Pointer to **int32** |  | [optional] 
**Rssi** | Pointer to **int32** |  | [optional] 
**RssiPercent** | Pointer to **float32** |  | [optional] 
**Rx** | Pointer to **int32** |  | [optional] 
**RxDropPkts** | Pointer to **int64** |  | [optional] 
**RxErrPkts** | Pointer to **int64** |  | [optional] 
**RxRate** | Pointer to **string** |  | [optional] 
**Snr** | Pointer to **int32** |  | [optional] 
**Tx** | Pointer to **int32** |  | [optional] 
**TxDropPkts** | Pointer to **int64** |  | [optional] 
**TxErrPkts** | Pointer to **int64** |  | [optional] 
**TxRate** | Pointer to **string** |  | [optional] 
**UpRate** | Pointer to **int64** |  | [optional] 

## Methods

### NewPartnerLinkDTO

`func NewPartnerLinkDTO() *PartnerLinkDTO`

NewPartnerLinkDTO instantiates a new PartnerLinkDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerLinkDTOWithDefaults

`func NewPartnerLinkDTOWithDefaults() *PartnerLinkDTO`

NewPartnerLinkDTOWithDefaults instantiates a new PartnerLinkDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *PartnerLinkDTO) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PartnerLinkDTO) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PartnerLinkDTO) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PartnerLinkDTO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDownRate

`func (o *PartnerLinkDTO) GetDownRate() int64`

GetDownRate returns the DownRate field if non-nil, zero value otherwise.

### GetDownRateOk

`func (o *PartnerLinkDTO) GetDownRateOk() (*int64, bool)`

GetDownRateOk returns a tuple with the DownRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownRate

`func (o *PartnerLinkDTO) SetDownRate(v int64)`

SetDownRate sets DownRate field to given value.

### HasDownRate

`func (o *PartnerLinkDTO) HasDownRate() bool`

HasDownRate returns a boolean if a field has been set.

### GetRadioId

`func (o *PartnerLinkDTO) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *PartnerLinkDTO) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *PartnerLinkDTO) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *PartnerLinkDTO) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetRssi

`func (o *PartnerLinkDTO) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *PartnerLinkDTO) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *PartnerLinkDTO) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *PartnerLinkDTO) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRssiPercent

`func (o *PartnerLinkDTO) GetRssiPercent() float32`

GetRssiPercent returns the RssiPercent field if non-nil, zero value otherwise.

### GetRssiPercentOk

`func (o *PartnerLinkDTO) GetRssiPercentOk() (*float32, bool)`

GetRssiPercentOk returns a tuple with the RssiPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiPercent

`func (o *PartnerLinkDTO) SetRssiPercent(v float32)`

SetRssiPercent sets RssiPercent field to given value.

### HasRssiPercent

`func (o *PartnerLinkDTO) HasRssiPercent() bool`

HasRssiPercent returns a boolean if a field has been set.

### GetRx

`func (o *PartnerLinkDTO) GetRx() int32`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *PartnerLinkDTO) GetRxOk() (*int32, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *PartnerLinkDTO) SetRx(v int32)`

SetRx sets Rx field to given value.

### HasRx

`func (o *PartnerLinkDTO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *PartnerLinkDTO) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *PartnerLinkDTO) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *PartnerLinkDTO) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *PartnerLinkDTO) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *PartnerLinkDTO) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *PartnerLinkDTO) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *PartnerLinkDTO) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *PartnerLinkDTO) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetRxRate

`func (o *PartnerLinkDTO) GetRxRate() string`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *PartnerLinkDTO) GetRxRateOk() (*string, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *PartnerLinkDTO) SetRxRate(v string)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *PartnerLinkDTO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSnr

`func (o *PartnerLinkDTO) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *PartnerLinkDTO) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *PartnerLinkDTO) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *PartnerLinkDTO) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetTx

`func (o *PartnerLinkDTO) GetTx() int32`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *PartnerLinkDTO) GetTxOk() (*int32, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *PartnerLinkDTO) SetTx(v int32)`

SetTx sets Tx field to given value.

### HasTx

`func (o *PartnerLinkDTO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *PartnerLinkDTO) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *PartnerLinkDTO) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *PartnerLinkDTO) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *PartnerLinkDTO) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *PartnerLinkDTO) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *PartnerLinkDTO) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *PartnerLinkDTO) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *PartnerLinkDTO) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetTxRate

`func (o *PartnerLinkDTO) GetTxRate() string`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *PartnerLinkDTO) GetTxRateOk() (*string, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *PartnerLinkDTO) SetTxRate(v string)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *PartnerLinkDTO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetUpRate

`func (o *PartnerLinkDTO) GetUpRate() int64`

GetUpRate returns the UpRate field if non-nil, zero value otherwise.

### GetUpRateOk

`func (o *PartnerLinkDTO) GetUpRateOk() (*int64, bool)`

GetUpRateOk returns a tuple with the UpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpRate

`func (o *PartnerLinkDTO) SetUpRate(v int64)`

SetUpRate sets UpRate field to given value.

### HasUpRate

`func (o *PartnerLinkDTO) HasUpRate() bool`

HasUpRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


