# UplinkAPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Connected actual channel. | [optional] 
**LagId** | Pointer to **int32** | Lag Id | [optional] 
**Name** | Pointer to **string** | Client connected port name. | [optional] 
**Port** | Pointer to **int32** | Client connected port. | [optional] 
**Radio** | Pointer to **int32** | Radio ID, 0: 2.4GHz; 1: 5GHz-1; 2:5GHz-2; 3: 6GHz. | [optional] 
**Rssi** | Pointer to **int32** | Signal strength, unit: dBm. | [optional] 
**RxRate** | Pointer to **int64** | Uplink negotiation rate (Kbit/s) | [optional] 
**Ssid** | Pointer to **string** | Connected SSID name. | [optional] 
**Support5g2** | Pointer to **bool** | Whether the AP support 5G2 radio. | [optional] 
**TrafficDown** | Pointer to **int64** | Downstream traffic (Byte). | [optional] 
**TrafficUp** | Pointer to **int64** | Upstream traffic (Byte). | [optional] 
**TxRate** | Pointer to **int64** | Downlink negotiation rate (Kbit/s) | [optional] 

## Methods

### NewUplinkAPInfo

`func NewUplinkAPInfo() *UplinkAPInfo`

NewUplinkAPInfo instantiates a new UplinkAPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUplinkAPInfoWithDefaults

`func NewUplinkAPInfoWithDefaults() *UplinkAPInfo`

NewUplinkAPInfoWithDefaults instantiates a new UplinkAPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *UplinkAPInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UplinkAPInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UplinkAPInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UplinkAPInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetLagId

`func (o *UplinkAPInfo) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *UplinkAPInfo) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *UplinkAPInfo) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *UplinkAPInfo) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetName

`func (o *UplinkAPInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UplinkAPInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UplinkAPInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UplinkAPInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *UplinkAPInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UplinkAPInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UplinkAPInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UplinkAPInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRadio

`func (o *UplinkAPInfo) GetRadio() int32`

GetRadio returns the Radio field if non-nil, zero value otherwise.

### GetRadioOk

`func (o *UplinkAPInfo) GetRadioOk() (*int32, bool)`

GetRadioOk returns a tuple with the Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio

`func (o *UplinkAPInfo) SetRadio(v int32)`

SetRadio sets Radio field to given value.

### HasRadio

`func (o *UplinkAPInfo) HasRadio() bool`

HasRadio returns a boolean if a field has been set.

### GetRssi

`func (o *UplinkAPInfo) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *UplinkAPInfo) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *UplinkAPInfo) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *UplinkAPInfo) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRxRate

`func (o *UplinkAPInfo) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *UplinkAPInfo) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *UplinkAPInfo) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *UplinkAPInfo) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSsid

`func (o *UplinkAPInfo) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *UplinkAPInfo) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *UplinkAPInfo) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *UplinkAPInfo) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSupport5g2

`func (o *UplinkAPInfo) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *UplinkAPInfo) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *UplinkAPInfo) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *UplinkAPInfo) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.

### GetTrafficDown

`func (o *UplinkAPInfo) GetTrafficDown() int64`

GetTrafficDown returns the TrafficDown field if non-nil, zero value otherwise.

### GetTrafficDownOk

`func (o *UplinkAPInfo) GetTrafficDownOk() (*int64, bool)`

GetTrafficDownOk returns a tuple with the TrafficDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDown

`func (o *UplinkAPInfo) SetTrafficDown(v int64)`

SetTrafficDown sets TrafficDown field to given value.

### HasTrafficDown

`func (o *UplinkAPInfo) HasTrafficDown() bool`

HasTrafficDown returns a boolean if a field has been set.

### GetTrafficUp

`func (o *UplinkAPInfo) GetTrafficUp() int64`

GetTrafficUp returns the TrafficUp field if non-nil, zero value otherwise.

### GetTrafficUpOk

`func (o *UplinkAPInfo) GetTrafficUpOk() (*int64, bool)`

GetTrafficUpOk returns a tuple with the TrafficUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUp

`func (o *UplinkAPInfo) SetTrafficUp(v int64)`

SetTrafficUp sets TrafficUp field to given value.

### HasTrafficUp

`func (o *UplinkAPInfo) HasTrafficUp() bool`

HasTrafficUp returns a boolean if a field has been set.

### GetTxRate

`func (o *UplinkAPInfo) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *UplinkAPInfo) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *UplinkAPInfo) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *UplinkAPInfo) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


