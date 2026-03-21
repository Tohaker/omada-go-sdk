# ApWiredUplinkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **int64** | (Change of( downBytes+upBytes))/ time, Unit: Bytes/s | [optional] 
**Current** | Pointer to **float64** | Current, only supported by some devices. | [optional] 
**DownBytes** | Pointer to **int64** | Unit: Byte | [optional] 
**DownDropPackets** | Pointer to **int64** | Uplink Device downDropPackets | [optional] 
**DownErrorsPackets** | Pointer to **int64** | Uplink Device downErrorsPackets | [optional] 
**DownPackets** | Pointer to **int64** | Uplink Device downPackets | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: LAN disconnected; 1: half; 2: full | [optional] 
**Ip** | Pointer to **string** | Uplink device ip. | [optional] 
**LinkSpeed** | Pointer to **int32** | Uplink port link speed | [optional] 
**LinkStatus** | Pointer to **int32** | Uplink port link status | [optional] 
**Model** | Pointer to **string** | Uplink device model | [optional] 
**ModelVersion** | Pointer to **string** | Uplink device modelVersion | [optional] 
**Name** | Pointer to **string** | Uplink device name | [optional] 
**Port** | Pointer to **string** | Uplink port ID, only supported by some devices. | [optional] 
**PortType** | Pointer to **int32** | Port Type, 0:ETH, 1:POTS, 2:SFP | [optional] 
**Rate** | Pointer to **string** | Negotiation rate, LAN(connected): 10,100,1000,2500,10000, LAN(disconnected):0, Unit:Mbps | [optional] 
**RxPower** | Pointer to **float64** | Rx Power, only supported by some devices. | [optional] 
**StackId** | Pointer to **string** |  | [optional] 
**Temp** | Pointer to **float64** | Temperature, only supported by some devices. | [optional] 
**TxPower** | Pointer to **float64** | Tx power, only supported by some devices. | [optional] 
**Type** | Pointer to **string** | Uplink device type | [optional] 
**UpBytes** | Pointer to **int64** | Unit: Byte | [optional] 
**UpDropPackets** | Pointer to **int64** | Uplink Device upDropPackets | [optional] 
**UpErrorsPackets** | Pointer to **int64** | Uplink Device upErrorsPackets | [optional] 
**UpPackets** | Pointer to **int64** | Uplink Device upPackets | [optional] 
**UplinkMac** | Pointer to **string** | Uplink device MAC address | [optional] 
**UplinkPort** | Pointer to **string** | Uplink device port | [optional] 
**Voltage** | Pointer to **float64** | Voltage, only supported by some devices. | [optional] 

## Methods

### NewApWiredUplinkInfo

`func NewApWiredUplinkInfo() *ApWiredUplinkInfo`

NewApWiredUplinkInfo instantiates a new ApWiredUplinkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApWiredUplinkInfoWithDefaults

`func NewApWiredUplinkInfoWithDefaults() *ApWiredUplinkInfo`

NewApWiredUplinkInfoWithDefaults instantiates a new ApWiredUplinkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *ApWiredUplinkInfo) GetActivity() int64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ApWiredUplinkInfo) GetActivityOk() (*int64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ApWiredUplinkInfo) SetActivity(v int64)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *ApWiredUplinkInfo) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetCurrent

`func (o *ApWiredUplinkInfo) GetCurrent() float64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ApWiredUplinkInfo) GetCurrentOk() (*float64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ApWiredUplinkInfo) SetCurrent(v float64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *ApWiredUplinkInfo) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetDownBytes

`func (o *ApWiredUplinkInfo) GetDownBytes() int64`

GetDownBytes returns the DownBytes field if non-nil, zero value otherwise.

### GetDownBytesOk

`func (o *ApWiredUplinkInfo) GetDownBytesOk() (*int64, bool)`

GetDownBytesOk returns a tuple with the DownBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBytes

`func (o *ApWiredUplinkInfo) SetDownBytes(v int64)`

SetDownBytes sets DownBytes field to given value.

### HasDownBytes

`func (o *ApWiredUplinkInfo) HasDownBytes() bool`

HasDownBytes returns a boolean if a field has been set.

### GetDownDropPackets

`func (o *ApWiredUplinkInfo) GetDownDropPackets() int64`

GetDownDropPackets returns the DownDropPackets field if non-nil, zero value otherwise.

### GetDownDropPacketsOk

`func (o *ApWiredUplinkInfo) GetDownDropPacketsOk() (*int64, bool)`

GetDownDropPacketsOk returns a tuple with the DownDropPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownDropPackets

`func (o *ApWiredUplinkInfo) SetDownDropPackets(v int64)`

SetDownDropPackets sets DownDropPackets field to given value.

### HasDownDropPackets

`func (o *ApWiredUplinkInfo) HasDownDropPackets() bool`

HasDownDropPackets returns a boolean if a field has been set.

### GetDownErrorsPackets

`func (o *ApWiredUplinkInfo) GetDownErrorsPackets() int64`

GetDownErrorsPackets returns the DownErrorsPackets field if non-nil, zero value otherwise.

### GetDownErrorsPacketsOk

`func (o *ApWiredUplinkInfo) GetDownErrorsPacketsOk() (*int64, bool)`

GetDownErrorsPacketsOk returns a tuple with the DownErrorsPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownErrorsPackets

`func (o *ApWiredUplinkInfo) SetDownErrorsPackets(v int64)`

SetDownErrorsPackets sets DownErrorsPackets field to given value.

### HasDownErrorsPackets

`func (o *ApWiredUplinkInfo) HasDownErrorsPackets() bool`

HasDownErrorsPackets returns a boolean if a field has been set.

### GetDownPackets

`func (o *ApWiredUplinkInfo) GetDownPackets() int64`

GetDownPackets returns the DownPackets field if non-nil, zero value otherwise.

### GetDownPacketsOk

`func (o *ApWiredUplinkInfo) GetDownPacketsOk() (*int64, bool)`

GetDownPacketsOk returns a tuple with the DownPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPackets

`func (o *ApWiredUplinkInfo) SetDownPackets(v int64)`

SetDownPackets sets DownPackets field to given value.

### HasDownPackets

`func (o *ApWiredUplinkInfo) HasDownPackets() bool`

HasDownPackets returns a boolean if a field has been set.

### GetDuplex

`func (o *ApWiredUplinkInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *ApWiredUplinkInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *ApWiredUplinkInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *ApWiredUplinkInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetIp

`func (o *ApWiredUplinkInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApWiredUplinkInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApWiredUplinkInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApWiredUplinkInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *ApWiredUplinkInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *ApWiredUplinkInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *ApWiredUplinkInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *ApWiredUplinkInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkStatus

`func (o *ApWiredUplinkInfo) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *ApWiredUplinkInfo) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *ApWiredUplinkInfo) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *ApWiredUplinkInfo) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetModel

`func (o *ApWiredUplinkInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApWiredUplinkInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApWiredUplinkInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApWiredUplinkInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApWiredUplinkInfo) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApWiredUplinkInfo) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApWiredUplinkInfo) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApWiredUplinkInfo) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ApWiredUplinkInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApWiredUplinkInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApWiredUplinkInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApWiredUplinkInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ApWiredUplinkInfo) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApWiredUplinkInfo) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApWiredUplinkInfo) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApWiredUplinkInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortType

`func (o *ApWiredUplinkInfo) GetPortType() int32`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *ApWiredUplinkInfo) GetPortTypeOk() (*int32, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *ApWiredUplinkInfo) SetPortType(v int32)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *ApWiredUplinkInfo) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetRate

`func (o *ApWiredUplinkInfo) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ApWiredUplinkInfo) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ApWiredUplinkInfo) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ApWiredUplinkInfo) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRxPower

`func (o *ApWiredUplinkInfo) GetRxPower() float64`

GetRxPower returns the RxPower field if non-nil, zero value otherwise.

### GetRxPowerOk

`func (o *ApWiredUplinkInfo) GetRxPowerOk() (*float64, bool)`

GetRxPowerOk returns a tuple with the RxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPower

`func (o *ApWiredUplinkInfo) SetRxPower(v float64)`

SetRxPower sets RxPower field to given value.

### HasRxPower

`func (o *ApWiredUplinkInfo) HasRxPower() bool`

HasRxPower returns a boolean if a field has been set.

### GetStackId

`func (o *ApWiredUplinkInfo) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *ApWiredUplinkInfo) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *ApWiredUplinkInfo) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *ApWiredUplinkInfo) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetTemp

`func (o *ApWiredUplinkInfo) GetTemp() float64`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *ApWiredUplinkInfo) GetTempOk() (*float64, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *ApWiredUplinkInfo) SetTemp(v float64)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *ApWiredUplinkInfo) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetTxPower

`func (o *ApWiredUplinkInfo) GetTxPower() float64`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *ApWiredUplinkInfo) GetTxPowerOk() (*float64, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *ApWiredUplinkInfo) SetTxPower(v float64)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *ApWiredUplinkInfo) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### GetType

`func (o *ApWiredUplinkInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApWiredUplinkInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApWiredUplinkInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApWiredUplinkInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpBytes

`func (o *ApWiredUplinkInfo) GetUpBytes() int64`

GetUpBytes returns the UpBytes field if non-nil, zero value otherwise.

### GetUpBytesOk

`func (o *ApWiredUplinkInfo) GetUpBytesOk() (*int64, bool)`

GetUpBytesOk returns a tuple with the UpBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBytes

`func (o *ApWiredUplinkInfo) SetUpBytes(v int64)`

SetUpBytes sets UpBytes field to given value.

### HasUpBytes

`func (o *ApWiredUplinkInfo) HasUpBytes() bool`

HasUpBytes returns a boolean if a field has been set.

### GetUpDropPackets

`func (o *ApWiredUplinkInfo) GetUpDropPackets() int64`

GetUpDropPackets returns the UpDropPackets field if non-nil, zero value otherwise.

### GetUpDropPacketsOk

`func (o *ApWiredUplinkInfo) GetUpDropPacketsOk() (*int64, bool)`

GetUpDropPacketsOk returns a tuple with the UpDropPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpDropPackets

`func (o *ApWiredUplinkInfo) SetUpDropPackets(v int64)`

SetUpDropPackets sets UpDropPackets field to given value.

### HasUpDropPackets

`func (o *ApWiredUplinkInfo) HasUpDropPackets() bool`

HasUpDropPackets returns a boolean if a field has been set.

### GetUpErrorsPackets

`func (o *ApWiredUplinkInfo) GetUpErrorsPackets() int64`

GetUpErrorsPackets returns the UpErrorsPackets field if non-nil, zero value otherwise.

### GetUpErrorsPacketsOk

`func (o *ApWiredUplinkInfo) GetUpErrorsPacketsOk() (*int64, bool)`

GetUpErrorsPacketsOk returns a tuple with the UpErrorsPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpErrorsPackets

`func (o *ApWiredUplinkInfo) SetUpErrorsPackets(v int64)`

SetUpErrorsPackets sets UpErrorsPackets field to given value.

### HasUpErrorsPackets

`func (o *ApWiredUplinkInfo) HasUpErrorsPackets() bool`

HasUpErrorsPackets returns a boolean if a field has been set.

### GetUpPackets

`func (o *ApWiredUplinkInfo) GetUpPackets() int64`

GetUpPackets returns the UpPackets field if non-nil, zero value otherwise.

### GetUpPacketsOk

`func (o *ApWiredUplinkInfo) GetUpPacketsOk() (*int64, bool)`

GetUpPacketsOk returns a tuple with the UpPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPackets

`func (o *ApWiredUplinkInfo) SetUpPackets(v int64)`

SetUpPackets sets UpPackets field to given value.

### HasUpPackets

`func (o *ApWiredUplinkInfo) HasUpPackets() bool`

HasUpPackets returns a boolean if a field has been set.

### GetUplinkMac

`func (o *ApWiredUplinkInfo) GetUplinkMac() string`

GetUplinkMac returns the UplinkMac field if non-nil, zero value otherwise.

### GetUplinkMacOk

`func (o *ApWiredUplinkInfo) GetUplinkMacOk() (*string, bool)`

GetUplinkMacOk returns a tuple with the UplinkMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkMac

`func (o *ApWiredUplinkInfo) SetUplinkMac(v string)`

SetUplinkMac sets UplinkMac field to given value.

### HasUplinkMac

`func (o *ApWiredUplinkInfo) HasUplinkMac() bool`

HasUplinkMac returns a boolean if a field has been set.

### GetUplinkPort

`func (o *ApWiredUplinkInfo) GetUplinkPort() string`

GetUplinkPort returns the UplinkPort field if non-nil, zero value otherwise.

### GetUplinkPortOk

`func (o *ApWiredUplinkInfo) GetUplinkPortOk() (*string, bool)`

GetUplinkPortOk returns a tuple with the UplinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPort

`func (o *ApWiredUplinkInfo) SetUplinkPort(v string)`

SetUplinkPort sets UplinkPort field to given value.

### HasUplinkPort

`func (o *ApWiredUplinkInfo) HasUplinkPort() bool`

HasUplinkPort returns a boolean if a field has been set.

### GetVoltage

`func (o *ApWiredUplinkInfo) GetVoltage() float64`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *ApWiredUplinkInfo) GetVoltageOk() (*float64, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *ApWiredUplinkInfo) SetVoltage(v float64)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *ApWiredUplinkInfo) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


