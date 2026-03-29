# ApWirelessUplink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **int64** | Uplink AP activity: (change of(downBytes+upBytes))/time | [optional] 
**Channel** | Pointer to **int32** | Uplink AP channel | [optional] 
**DownBytes** | Pointer to **int64** | Uplink AP downBytes; Unit:Byte | [optional] 
**DownPackets** | Pointer to **int64** | Uplink AP downPackets | [optional] 
**Ip** | Pointer to **string** | Device IP. | [optional] 
**Model** | Pointer to **string** | Device model.(Used to display the device model diagram) | [optional] 
**ModelVersion** | Pointer to **string** | Device model version.(Used to display the device model diagram) | [optional] 
**Name** | Pointer to **string** | Uplink AP name | [optional] 
**Rssi** | Pointer to **int32** | Uplink AP rssi | [optional] 
**RxRate** | Pointer to **string** | Uplink AP rxRate | [optional] 
**RxRateInt** | Pointer to **int32** | Uplink AP rxRateInt; Unit: Mbps | [optional] 
**Snr** | Pointer to **int32** | Uplink AP Signal-noise ratio | [optional] 
**SupportSpeedTest** | Pointer to **bool** | Whether speed measurement is supported. | [optional] 
**TxRate** | Pointer to **string** | Uplink AP txRate | [optional] 
**TxRateInt** | Pointer to **int32** | Uplink AP txRateInt; Unit: Mbps | [optional] 
**Type** | Pointer to **string** | Device type. | [optional] 
**UpBytes** | Pointer to **int64** | Uplink AP upBytes; Unit: Byte | [optional] 
**UpPackets** | Pointer to **int64** | Uplink AP upPackets | [optional] 
**UplinkMac** | Pointer to **string** | Uplink AP MAC | [optional] 
**UplinkPort** | Pointer to **string** | Uplink port of current device. To mark that this port cannot be configured. | [optional] 

## Methods

### NewApWirelessUplink

`func NewApWirelessUplink() *ApWirelessUplink`

NewApWirelessUplink instantiates a new ApWirelessUplink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApWirelessUplinkWithDefaults

`func NewApWirelessUplinkWithDefaults() *ApWirelessUplink`

NewApWirelessUplinkWithDefaults instantiates a new ApWirelessUplink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *ApWirelessUplink) GetActivity() int64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ApWirelessUplink) GetActivityOk() (*int64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ApWirelessUplink) SetActivity(v int64)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *ApWirelessUplink) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetChannel

`func (o *ApWirelessUplink) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApWirelessUplink) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApWirelessUplink) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApWirelessUplink) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDownBytes

`func (o *ApWirelessUplink) GetDownBytes() int64`

GetDownBytes returns the DownBytes field if non-nil, zero value otherwise.

### GetDownBytesOk

`func (o *ApWirelessUplink) GetDownBytesOk() (*int64, bool)`

GetDownBytesOk returns a tuple with the DownBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBytes

`func (o *ApWirelessUplink) SetDownBytes(v int64)`

SetDownBytes sets DownBytes field to given value.

### HasDownBytes

`func (o *ApWirelessUplink) HasDownBytes() bool`

HasDownBytes returns a boolean if a field has been set.

### GetDownPackets

`func (o *ApWirelessUplink) GetDownPackets() int64`

GetDownPackets returns the DownPackets field if non-nil, zero value otherwise.

### GetDownPacketsOk

`func (o *ApWirelessUplink) GetDownPacketsOk() (*int64, bool)`

GetDownPacketsOk returns a tuple with the DownPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPackets

`func (o *ApWirelessUplink) SetDownPackets(v int64)`

SetDownPackets sets DownPackets field to given value.

### HasDownPackets

`func (o *ApWirelessUplink) HasDownPackets() bool`

HasDownPackets returns a boolean if a field has been set.

### GetIp

`func (o *ApWirelessUplink) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApWirelessUplink) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApWirelessUplink) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApWirelessUplink) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetModel

`func (o *ApWirelessUplink) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApWirelessUplink) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApWirelessUplink) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApWirelessUplink) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApWirelessUplink) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApWirelessUplink) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApWirelessUplink) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApWirelessUplink) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ApWirelessUplink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApWirelessUplink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApWirelessUplink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApWirelessUplink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRssi

`func (o *ApWirelessUplink) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *ApWirelessUplink) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *ApWirelessUplink) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *ApWirelessUplink) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRxRate

`func (o *ApWirelessUplink) GetRxRate() string`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *ApWirelessUplink) GetRxRateOk() (*string, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *ApWirelessUplink) SetRxRate(v string)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *ApWirelessUplink) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetRxRateInt

`func (o *ApWirelessUplink) GetRxRateInt() int32`

GetRxRateInt returns the RxRateInt field if non-nil, zero value otherwise.

### GetRxRateIntOk

`func (o *ApWirelessUplink) GetRxRateIntOk() (*int32, bool)`

GetRxRateIntOk returns a tuple with the RxRateInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRateInt

`func (o *ApWirelessUplink) SetRxRateInt(v int32)`

SetRxRateInt sets RxRateInt field to given value.

### HasRxRateInt

`func (o *ApWirelessUplink) HasRxRateInt() bool`

HasRxRateInt returns a boolean if a field has been set.

### GetSnr

`func (o *ApWirelessUplink) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *ApWirelessUplink) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *ApWirelessUplink) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *ApWirelessUplink) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetSupportSpeedTest

`func (o *ApWirelessUplink) GetSupportSpeedTest() bool`

GetSupportSpeedTest returns the SupportSpeedTest field if non-nil, zero value otherwise.

### GetSupportSpeedTestOk

`func (o *ApWirelessUplink) GetSupportSpeedTestOk() (*bool, bool)`

GetSupportSpeedTestOk returns a tuple with the SupportSpeedTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSpeedTest

`func (o *ApWirelessUplink) SetSupportSpeedTest(v bool)`

SetSupportSpeedTest sets SupportSpeedTest field to given value.

### HasSupportSpeedTest

`func (o *ApWirelessUplink) HasSupportSpeedTest() bool`

HasSupportSpeedTest returns a boolean if a field has been set.

### GetTxRate

`func (o *ApWirelessUplink) GetTxRate() string`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *ApWirelessUplink) GetTxRateOk() (*string, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *ApWirelessUplink) SetTxRate(v string)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *ApWirelessUplink) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetTxRateInt

`func (o *ApWirelessUplink) GetTxRateInt() int32`

GetTxRateInt returns the TxRateInt field if non-nil, zero value otherwise.

### GetTxRateIntOk

`func (o *ApWirelessUplink) GetTxRateIntOk() (*int32, bool)`

GetTxRateIntOk returns a tuple with the TxRateInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRateInt

`func (o *ApWirelessUplink) SetTxRateInt(v int32)`

SetTxRateInt sets TxRateInt field to given value.

### HasTxRateInt

`func (o *ApWirelessUplink) HasTxRateInt() bool`

HasTxRateInt returns a boolean if a field has been set.

### GetType

`func (o *ApWirelessUplink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApWirelessUplink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApWirelessUplink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApWirelessUplink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpBytes

`func (o *ApWirelessUplink) GetUpBytes() int64`

GetUpBytes returns the UpBytes field if non-nil, zero value otherwise.

### GetUpBytesOk

`func (o *ApWirelessUplink) GetUpBytesOk() (*int64, bool)`

GetUpBytesOk returns a tuple with the UpBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBytes

`func (o *ApWirelessUplink) SetUpBytes(v int64)`

SetUpBytes sets UpBytes field to given value.

### HasUpBytes

`func (o *ApWirelessUplink) HasUpBytes() bool`

HasUpBytes returns a boolean if a field has been set.

### GetUpPackets

`func (o *ApWirelessUplink) GetUpPackets() int64`

GetUpPackets returns the UpPackets field if non-nil, zero value otherwise.

### GetUpPacketsOk

`func (o *ApWirelessUplink) GetUpPacketsOk() (*int64, bool)`

GetUpPacketsOk returns a tuple with the UpPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPackets

`func (o *ApWirelessUplink) SetUpPackets(v int64)`

SetUpPackets sets UpPackets field to given value.

### HasUpPackets

`func (o *ApWirelessUplink) HasUpPackets() bool`

HasUpPackets returns a boolean if a field has been set.

### GetUplinkMac

`func (o *ApWirelessUplink) GetUplinkMac() string`

GetUplinkMac returns the UplinkMac field if non-nil, zero value otherwise.

### GetUplinkMacOk

`func (o *ApWirelessUplink) GetUplinkMacOk() (*string, bool)`

GetUplinkMacOk returns a tuple with the UplinkMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkMac

`func (o *ApWirelessUplink) SetUplinkMac(v string)`

SetUplinkMac sets UplinkMac field to given value.

### HasUplinkMac

`func (o *ApWirelessUplink) HasUplinkMac() bool`

HasUplinkMac returns a boolean if a field has been set.

### GetUplinkPort

`func (o *ApWirelessUplink) GetUplinkPort() string`

GetUplinkPort returns the UplinkPort field if non-nil, zero value otherwise.

### GetUplinkPortOk

`func (o *ApWirelessUplink) GetUplinkPortOk() (*string, bool)`

GetUplinkPortOk returns a tuple with the UplinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPort

`func (o *ApWirelessUplink) SetUplinkPort(v string)`

SetUplinkPort sets UplinkPort field to given value.

### HasUplinkPort

`func (o *ApWirelessUplink) HasUplinkPort() bool`

HasUplinkPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


