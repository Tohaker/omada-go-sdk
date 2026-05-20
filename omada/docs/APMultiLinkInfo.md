# APMultiLinkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | (Wireless)  Actual channel. | [optional] 
**RadioId** | Pointer to **int32** | Radio ID should be a value as follows: 0: 2.4GHz; 1: 5GHz-1; 2:5GHz-2; 3: 6GHz. | [optional] 
**Rssi** | Pointer to **int32** | Signal strength, unit: dBm. | [optional] 
**RxRate** | Pointer to **int64** | Uplink negotiation rate (Kbit/s) | [optional] 
**TrafficDown** | Pointer to **int64** | Downstream traffic (Byte). | [optional] 
**TrafficUp** | Pointer to **int64** | Upstream traffic (Byte). | [optional] 
**TxRate** | Pointer to **int64** | Downlink negotiation rate (Kbit/s) | [optional] 

## Methods

### NewAPMultiLinkInfo

`func NewAPMultiLinkInfo() *APMultiLinkInfo`

NewAPMultiLinkInfo instantiates a new APMultiLinkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPMultiLinkInfoWithDefaults

`func NewAPMultiLinkInfoWithDefaults() *APMultiLinkInfo`

NewAPMultiLinkInfoWithDefaults instantiates a new APMultiLinkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *APMultiLinkInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *APMultiLinkInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *APMultiLinkInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *APMultiLinkInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetRadioId

`func (o *APMultiLinkInfo) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *APMultiLinkInfo) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *APMultiLinkInfo) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *APMultiLinkInfo) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetRssi

`func (o *APMultiLinkInfo) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *APMultiLinkInfo) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *APMultiLinkInfo) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *APMultiLinkInfo) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRxRate

`func (o *APMultiLinkInfo) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *APMultiLinkInfo) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *APMultiLinkInfo) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *APMultiLinkInfo) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetTrafficDown

`func (o *APMultiLinkInfo) GetTrafficDown() int64`

GetTrafficDown returns the TrafficDown field if non-nil, zero value otherwise.

### GetTrafficDownOk

`func (o *APMultiLinkInfo) GetTrafficDownOk() (*int64, bool)`

GetTrafficDownOk returns a tuple with the TrafficDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDown

`func (o *APMultiLinkInfo) SetTrafficDown(v int64)`

SetTrafficDown sets TrafficDown field to given value.

### HasTrafficDown

`func (o *APMultiLinkInfo) HasTrafficDown() bool`

HasTrafficDown returns a boolean if a field has been set.

### GetTrafficUp

`func (o *APMultiLinkInfo) GetTrafficUp() int64`

GetTrafficUp returns the TrafficUp field if non-nil, zero value otherwise.

### GetTrafficUpOk

`func (o *APMultiLinkInfo) GetTrafficUpOk() (*int64, bool)`

GetTrafficUpOk returns a tuple with the TrafficUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUp

`func (o *APMultiLinkInfo) SetTrafficUp(v int64)`

SetTrafficUp sets TrafficUp field to given value.

### HasTrafficUp

`func (o *APMultiLinkInfo) HasTrafficUp() bool`

HasTrafficUp returns a boolean if a field has been set.

### GetTxRate

`func (o *APMultiLinkInfo) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *APMultiLinkInfo) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *APMultiLinkInfo) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *APMultiLinkInfo) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


