# ApRadioSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSwitchOffWifi** | Pointer to **bool** | Enable auto Switch Off Wifi | [optional] 
**AutoSwitchOffWifiInterval** | Pointer to **int32** | Auto Switch Off Wifi Interval | [optional] 
**Channel** | Pointer to **string** | Channel Index; The channel index list supported by device can be obtained from interface : Get available channel list of ap; If select auto configuration need to enter 0. | [optional] 
**ChannelLimitEnable** | Pointer to **bool** | Enable channel limit(site level) | [optional] 
**ChannelRange** | Pointer to **[]int32** | Custom optional channel freq collection. | [optional] 
**ChannelWidth** | Pointer to **string** | RADIO_20 &#x3D; 2; RADIO_40 &#x3D; 3; RADIO_40_20 &#x3D; 4(corresponding 2G Auto); RADIO_80 &#x3D; 5; RADIO_80_40_20 &#x3D; 6(corresponding 5G Auto); RADIO_160&#x3D; 7; RADIO_160_80_40_20 &#x3D; 8; RADIO_240 &#x3D; 9; RADIO_320 &#x3D; 10 | [optional] 
**Freq** | Pointer to **int32** | Frequency;The frequency list supported by device can be obtained from interface : Get available channel list of ap.The freq and channel fields should be corresponding, otherwise the channel will be corrected based on the freq field. | [optional] 
**InterTimeStamp** | Pointer to **int64** |  | [optional] 
**RadioEnable** | Pointer to **bool** | Enable/Disable radio setting(if false, other params is not required) | [optional] 
**RrmType** | Pointer to **int32** |  | [optional] 
**TxPower** | Pointer to **int32** | TX Power | [optional] 
**TxPowerLevel** | Pointer to **int32** | It should be a value as follows: 0: Low; 1: Medium; 2: High; 3: Custom; 4: Auto | [optional] 
**WirelessMode** | Pointer to **int32** | Wireless mode config status of the device; -2 : Auto; 3 : 802.11b/g mixed(Only for 2.4G); 4 : 802.11b/g/n mixed(Only for 2.4G); 13 : 802.11b/g/n/ax mixed(Only for 2.4G); 17 : 802.11b/g/n/ax/be mixed(Only for 2.4G); 7 : 802.11a/n mixed(Only for 5G); 10 : 802.11a/n/ac mixed(Only for 5G); 16 : 802.11a/n/ac/ax mixed(Only for 5G); 18 : 802.11a/n/ac/ax/be mixed(Only for 5G); 11 : 802.11ax only(Only for 6G); 19 : 802.11ax/be mixed(Only for 6G). | [optional] 

## Methods

### NewApRadioSetting

`func NewApRadioSetting() *ApRadioSetting`

NewApRadioSetting instantiates a new ApRadioSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRadioSettingWithDefaults

`func NewApRadioSettingWithDefaults() *ApRadioSetting`

NewApRadioSettingWithDefaults instantiates a new ApRadioSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSwitchOffWifi

`func (o *ApRadioSetting) GetAutoSwitchOffWifi() bool`

GetAutoSwitchOffWifi returns the AutoSwitchOffWifi field if non-nil, zero value otherwise.

### GetAutoSwitchOffWifiOk

`func (o *ApRadioSetting) GetAutoSwitchOffWifiOk() (*bool, bool)`

GetAutoSwitchOffWifiOk returns a tuple with the AutoSwitchOffWifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSwitchOffWifi

`func (o *ApRadioSetting) SetAutoSwitchOffWifi(v bool)`

SetAutoSwitchOffWifi sets AutoSwitchOffWifi field to given value.

### HasAutoSwitchOffWifi

`func (o *ApRadioSetting) HasAutoSwitchOffWifi() bool`

HasAutoSwitchOffWifi returns a boolean if a field has been set.

### GetAutoSwitchOffWifiInterval

`func (o *ApRadioSetting) GetAutoSwitchOffWifiInterval() int32`

GetAutoSwitchOffWifiInterval returns the AutoSwitchOffWifiInterval field if non-nil, zero value otherwise.

### GetAutoSwitchOffWifiIntervalOk

`func (o *ApRadioSetting) GetAutoSwitchOffWifiIntervalOk() (*int32, bool)`

GetAutoSwitchOffWifiIntervalOk returns a tuple with the AutoSwitchOffWifiInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSwitchOffWifiInterval

`func (o *ApRadioSetting) SetAutoSwitchOffWifiInterval(v int32)`

SetAutoSwitchOffWifiInterval sets AutoSwitchOffWifiInterval field to given value.

### HasAutoSwitchOffWifiInterval

`func (o *ApRadioSetting) HasAutoSwitchOffWifiInterval() bool`

HasAutoSwitchOffWifiInterval returns a boolean if a field has been set.

### GetChannel

`func (o *ApRadioSetting) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApRadioSetting) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApRadioSetting) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApRadioSetting) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelLimitEnable

`func (o *ApRadioSetting) GetChannelLimitEnable() bool`

GetChannelLimitEnable returns the ChannelLimitEnable field if non-nil, zero value otherwise.

### GetChannelLimitEnableOk

`func (o *ApRadioSetting) GetChannelLimitEnableOk() (*bool, bool)`

GetChannelLimitEnableOk returns a tuple with the ChannelLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimitEnable

`func (o *ApRadioSetting) SetChannelLimitEnable(v bool)`

SetChannelLimitEnable sets ChannelLimitEnable field to given value.

### HasChannelLimitEnable

`func (o *ApRadioSetting) HasChannelLimitEnable() bool`

HasChannelLimitEnable returns a boolean if a field has been set.

### GetChannelRange

`func (o *ApRadioSetting) GetChannelRange() []int32`

GetChannelRange returns the ChannelRange field if non-nil, zero value otherwise.

### GetChannelRangeOk

`func (o *ApRadioSetting) GetChannelRangeOk() (*[]int32, bool)`

GetChannelRangeOk returns a tuple with the ChannelRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelRange

`func (o *ApRadioSetting) SetChannelRange(v []int32)`

SetChannelRange sets ChannelRange field to given value.

### HasChannelRange

`func (o *ApRadioSetting) HasChannelRange() bool`

HasChannelRange returns a boolean if a field has been set.

### GetChannelWidth

`func (o *ApRadioSetting) GetChannelWidth() string`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *ApRadioSetting) GetChannelWidthOk() (*string, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *ApRadioSetting) SetChannelWidth(v string)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *ApRadioSetting) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetFreq

`func (o *ApRadioSetting) GetFreq() int32`

GetFreq returns the Freq field if non-nil, zero value otherwise.

### GetFreqOk

`func (o *ApRadioSetting) GetFreqOk() (*int32, bool)`

GetFreqOk returns a tuple with the Freq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreq

`func (o *ApRadioSetting) SetFreq(v int32)`

SetFreq sets Freq field to given value.

### HasFreq

`func (o *ApRadioSetting) HasFreq() bool`

HasFreq returns a boolean if a field has been set.

### GetInterTimeStamp

`func (o *ApRadioSetting) GetInterTimeStamp() int64`

GetInterTimeStamp returns the InterTimeStamp field if non-nil, zero value otherwise.

### GetInterTimeStampOk

`func (o *ApRadioSetting) GetInterTimeStampOk() (*int64, bool)`

GetInterTimeStampOk returns a tuple with the InterTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterTimeStamp

`func (o *ApRadioSetting) SetInterTimeStamp(v int64)`

SetInterTimeStamp sets InterTimeStamp field to given value.

### HasInterTimeStamp

`func (o *ApRadioSetting) HasInterTimeStamp() bool`

HasInterTimeStamp returns a boolean if a field has been set.

### GetRadioEnable

`func (o *ApRadioSetting) GetRadioEnable() bool`

GetRadioEnable returns the RadioEnable field if non-nil, zero value otherwise.

### GetRadioEnableOk

`func (o *ApRadioSetting) GetRadioEnableOk() (*bool, bool)`

GetRadioEnableOk returns a tuple with the RadioEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioEnable

`func (o *ApRadioSetting) SetRadioEnable(v bool)`

SetRadioEnable sets RadioEnable field to given value.

### HasRadioEnable

`func (o *ApRadioSetting) HasRadioEnable() bool`

HasRadioEnable returns a boolean if a field has been set.

### GetRrmType

`func (o *ApRadioSetting) GetRrmType() int32`

GetRrmType returns the RrmType field if non-nil, zero value otherwise.

### GetRrmTypeOk

`func (o *ApRadioSetting) GetRrmTypeOk() (*int32, bool)`

GetRrmTypeOk returns a tuple with the RrmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmType

`func (o *ApRadioSetting) SetRrmType(v int32)`

SetRrmType sets RrmType field to given value.

### HasRrmType

`func (o *ApRadioSetting) HasRrmType() bool`

HasRrmType returns a boolean if a field has been set.

### GetTxPower

`func (o *ApRadioSetting) GetTxPower() int32`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *ApRadioSetting) GetTxPowerOk() (*int32, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *ApRadioSetting) SetTxPower(v int32)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *ApRadioSetting) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### GetTxPowerLevel

`func (o *ApRadioSetting) GetTxPowerLevel() int32`

GetTxPowerLevel returns the TxPowerLevel field if non-nil, zero value otherwise.

### GetTxPowerLevelOk

`func (o *ApRadioSetting) GetTxPowerLevelOk() (*int32, bool)`

GetTxPowerLevelOk returns a tuple with the TxPowerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPowerLevel

`func (o *ApRadioSetting) SetTxPowerLevel(v int32)`

SetTxPowerLevel sets TxPowerLevel field to given value.

### HasTxPowerLevel

`func (o *ApRadioSetting) HasTxPowerLevel() bool`

HasTxPowerLevel returns a boolean if a field has been set.

### GetWirelessMode

`func (o *ApRadioSetting) GetWirelessMode() int32`

GetWirelessMode returns the WirelessMode field if non-nil, zero value otherwise.

### GetWirelessModeOk

`func (o *ApRadioSetting) GetWirelessModeOk() (*int32, bool)`

GetWirelessModeOk returns a tuple with the WirelessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessMode

`func (o *ApRadioSetting) SetWirelessMode(v int32)`

SetWirelessMode sets WirelessMode field to given value.

### HasWirelessMode

`func (o *ApRadioSetting) HasWirelessMode() bool`

HasWirelessMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


