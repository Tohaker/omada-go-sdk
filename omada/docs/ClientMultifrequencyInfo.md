# ClientMultifrequencyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **int64** | Real-time downlink rate (Byte/s). | [optional] 
**Channel** | Pointer to **int32** | (Wireless)  Actual channel. | [optional] 
**DownPacket** | Pointer to **int64** | Number of downstream packets. | [optional] 
**PowerSave** | Pointer to **bool** | (Wireless)  true: Power save mode enabled. | [optional] 
**RadioId** | Pointer to **int32** | Radio ID should be a value as follows: 0: 2.4GHz; 1: 5GHz-1; 2:5GHz-2; 3: 6GHz. | [optional] 
**Rssi** | Pointer to **int32** | (Wireless) Signal strength, unit: dBm. | [optional] 
**RxRate** | Pointer to **int64** | (Wireless) Uplink negotiation rate (Kbit/s). | [optional] 
**SignalLevel** | Pointer to **int32** | (Wireless) Signal strength percentage should be within the range of 0-100. | [optional] 
**SignalLevelAndRank** | Pointer to **int32** |  | [optional] 
**SignalRank** | Pointer to **int32** | (Wireless) Signal strength level should be within the range of 0-5. | [optional] 
**Snr** | Pointer to **int32** | (Wireless) Signal Noise Ratio. | [optional] 
**TrafficDown** | Pointer to **int64** | Downstream traffic (Byte). | [optional] 
**TrafficUp** | Pointer to **int64** | Upstream traffic (Byte). | [optional] 
**TxRate** | Pointer to **int64** | (Wireless) Downlink negotiation rate (Kbit/s). | [optional] 
**UpPacket** | Pointer to **int64** | Number of upstream packets. | [optional] 
**WifiMode** | Pointer to **int32** | Wi-Fi mode should be a value as follows: 0: 11a; 1: 11b; 2: 11g; 3: 11na; 4: 11ng; 5: 11ac; 6: 11axa; 7: 11axg. | [optional] 

## Methods

### NewClientMultifrequencyInfo

`func NewClientMultifrequencyInfo() *ClientMultifrequencyInfo`

NewClientMultifrequencyInfo instantiates a new ClientMultifrequencyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMultifrequencyInfoWithDefaults

`func NewClientMultifrequencyInfoWithDefaults() *ClientMultifrequencyInfo`

NewClientMultifrequencyInfoWithDefaults instantiates a new ClientMultifrequencyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *ClientMultifrequencyInfo) GetActivity() int64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ClientMultifrequencyInfo) GetActivityOk() (*int64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ClientMultifrequencyInfo) SetActivity(v int64)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *ClientMultifrequencyInfo) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetChannel

`func (o *ClientMultifrequencyInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ClientMultifrequencyInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ClientMultifrequencyInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ClientMultifrequencyInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDownPacket

`func (o *ClientMultifrequencyInfo) GetDownPacket() int64`

GetDownPacket returns the DownPacket field if non-nil, zero value otherwise.

### GetDownPacketOk

`func (o *ClientMultifrequencyInfo) GetDownPacketOk() (*int64, bool)`

GetDownPacketOk returns a tuple with the DownPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPacket

`func (o *ClientMultifrequencyInfo) SetDownPacket(v int64)`

SetDownPacket sets DownPacket field to given value.

### HasDownPacket

`func (o *ClientMultifrequencyInfo) HasDownPacket() bool`

HasDownPacket returns a boolean if a field has been set.

### GetPowerSave

`func (o *ClientMultifrequencyInfo) GetPowerSave() bool`

GetPowerSave returns the PowerSave field if non-nil, zero value otherwise.

### GetPowerSaveOk

`func (o *ClientMultifrequencyInfo) GetPowerSaveOk() (*bool, bool)`

GetPowerSaveOk returns a tuple with the PowerSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSave

`func (o *ClientMultifrequencyInfo) SetPowerSave(v bool)`

SetPowerSave sets PowerSave field to given value.

### HasPowerSave

`func (o *ClientMultifrequencyInfo) HasPowerSave() bool`

HasPowerSave returns a boolean if a field has been set.

### GetRadioId

`func (o *ClientMultifrequencyInfo) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *ClientMultifrequencyInfo) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *ClientMultifrequencyInfo) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *ClientMultifrequencyInfo) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetRssi

`func (o *ClientMultifrequencyInfo) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *ClientMultifrequencyInfo) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *ClientMultifrequencyInfo) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *ClientMultifrequencyInfo) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRxRate

`func (o *ClientMultifrequencyInfo) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *ClientMultifrequencyInfo) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *ClientMultifrequencyInfo) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *ClientMultifrequencyInfo) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSignalLevel

`func (o *ClientMultifrequencyInfo) GetSignalLevel() int32`

GetSignalLevel returns the SignalLevel field if non-nil, zero value otherwise.

### GetSignalLevelOk

`func (o *ClientMultifrequencyInfo) GetSignalLevelOk() (*int32, bool)`

GetSignalLevelOk returns a tuple with the SignalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalLevel

`func (o *ClientMultifrequencyInfo) SetSignalLevel(v int32)`

SetSignalLevel sets SignalLevel field to given value.

### HasSignalLevel

`func (o *ClientMultifrequencyInfo) HasSignalLevel() bool`

HasSignalLevel returns a boolean if a field has been set.

### GetSignalLevelAndRank

`func (o *ClientMultifrequencyInfo) GetSignalLevelAndRank() int32`

GetSignalLevelAndRank returns the SignalLevelAndRank field if non-nil, zero value otherwise.

### GetSignalLevelAndRankOk

`func (o *ClientMultifrequencyInfo) GetSignalLevelAndRankOk() (*int32, bool)`

GetSignalLevelAndRankOk returns a tuple with the SignalLevelAndRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalLevelAndRank

`func (o *ClientMultifrequencyInfo) SetSignalLevelAndRank(v int32)`

SetSignalLevelAndRank sets SignalLevelAndRank field to given value.

### HasSignalLevelAndRank

`func (o *ClientMultifrequencyInfo) HasSignalLevelAndRank() bool`

HasSignalLevelAndRank returns a boolean if a field has been set.

### GetSignalRank

`func (o *ClientMultifrequencyInfo) GetSignalRank() int32`

GetSignalRank returns the SignalRank field if non-nil, zero value otherwise.

### GetSignalRankOk

`func (o *ClientMultifrequencyInfo) GetSignalRankOk() (*int32, bool)`

GetSignalRankOk returns a tuple with the SignalRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalRank

`func (o *ClientMultifrequencyInfo) SetSignalRank(v int32)`

SetSignalRank sets SignalRank field to given value.

### HasSignalRank

`func (o *ClientMultifrequencyInfo) HasSignalRank() bool`

HasSignalRank returns a boolean if a field has been set.

### GetSnr

`func (o *ClientMultifrequencyInfo) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *ClientMultifrequencyInfo) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *ClientMultifrequencyInfo) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *ClientMultifrequencyInfo) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetTrafficDown

`func (o *ClientMultifrequencyInfo) GetTrafficDown() int64`

GetTrafficDown returns the TrafficDown field if non-nil, zero value otherwise.

### GetTrafficDownOk

`func (o *ClientMultifrequencyInfo) GetTrafficDownOk() (*int64, bool)`

GetTrafficDownOk returns a tuple with the TrafficDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDown

`func (o *ClientMultifrequencyInfo) SetTrafficDown(v int64)`

SetTrafficDown sets TrafficDown field to given value.

### HasTrafficDown

`func (o *ClientMultifrequencyInfo) HasTrafficDown() bool`

HasTrafficDown returns a boolean if a field has been set.

### GetTrafficUp

`func (o *ClientMultifrequencyInfo) GetTrafficUp() int64`

GetTrafficUp returns the TrafficUp field if non-nil, zero value otherwise.

### GetTrafficUpOk

`func (o *ClientMultifrequencyInfo) GetTrafficUpOk() (*int64, bool)`

GetTrafficUpOk returns a tuple with the TrafficUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUp

`func (o *ClientMultifrequencyInfo) SetTrafficUp(v int64)`

SetTrafficUp sets TrafficUp field to given value.

### HasTrafficUp

`func (o *ClientMultifrequencyInfo) HasTrafficUp() bool`

HasTrafficUp returns a boolean if a field has been set.

### GetTxRate

`func (o *ClientMultifrequencyInfo) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *ClientMultifrequencyInfo) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *ClientMultifrequencyInfo) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *ClientMultifrequencyInfo) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetUpPacket

`func (o *ClientMultifrequencyInfo) GetUpPacket() int64`

GetUpPacket returns the UpPacket field if non-nil, zero value otherwise.

### GetUpPacketOk

`func (o *ClientMultifrequencyInfo) GetUpPacketOk() (*int64, bool)`

GetUpPacketOk returns a tuple with the UpPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPacket

`func (o *ClientMultifrequencyInfo) SetUpPacket(v int64)`

SetUpPacket sets UpPacket field to given value.

### HasUpPacket

`func (o *ClientMultifrequencyInfo) HasUpPacket() bool`

HasUpPacket returns a boolean if a field has been set.

### GetWifiMode

`func (o *ClientMultifrequencyInfo) GetWifiMode() int32`

GetWifiMode returns the WifiMode field if non-nil, zero value otherwise.

### GetWifiModeOk

`func (o *ClientMultifrequencyInfo) GetWifiModeOk() (*int32, bool)`

GetWifiModeOk returns a tuple with the WifiMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMode

`func (o *ClientMultifrequencyInfo) SetWifiMode(v int32)`

SetWifiMode sets WifiMode field to given value.

### HasWifiMode

`func (o *ClientMultifrequencyInfo) HasWifiMode() bool`

HasWifiMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


