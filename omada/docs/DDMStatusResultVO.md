# DDMStatusResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiasCurrent** | Pointer to **string** | Port optical module bias current. | [optional] 
**BiasCurrentFlag** | Pointer to **int32** | Port optical module bias current alert information. | [optional] 
**DataReady** | Pointer to **bool** | Whether the port optical module is in an available state. | [optional] 
**LossOfSignal** | Pointer to **bool** | Whether the local optical module signal is distorted. | [optional] 
**Port** | Pointer to **string** | OLT optical ports, including Ethernet optical ports and PON ports. | [optional] 
**RxPower** | Pointer to **string** | Port optical module received power. | [optional] 
**RxPowerFlag** | Pointer to **int32** | Port optical module received power alert information. | [optional] 
**Temperature** | Pointer to **string** | Temperature of the port optical module. | [optional] 
**TemperatureFlag** | Pointer to **int32** | Port optical module temperature alert information. | [optional] 
**TransmitFault** | Pointer to **int32** | Weather the signal from the remote optical module is distorted | [optional] 
**TxPower** | Pointer to **string** | Port optical module transmission power. | [optional] 
**TxPowerFlag** | Pointer to **int32** | Port optical module transmission power alert information. | [optional] 
**Voltage** | Pointer to **string** | Port optical module voltage. | [optional] 
**VoltageFlag** | Pointer to **int32** | Port optical module voltage alert information. | [optional] 

## Methods

### NewDDMStatusResultVO

`func NewDDMStatusResultVO() *DDMStatusResultVO`

NewDDMStatusResultVO instantiates a new DDMStatusResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDDMStatusResultVOWithDefaults

`func NewDDMStatusResultVOWithDefaults() *DDMStatusResultVO`

NewDDMStatusResultVOWithDefaults instantiates a new DDMStatusResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiasCurrent

`func (o *DDMStatusResultVO) GetBiasCurrent() string`

GetBiasCurrent returns the BiasCurrent field if non-nil, zero value otherwise.

### GetBiasCurrentOk

`func (o *DDMStatusResultVO) GetBiasCurrentOk() (*string, bool)`

GetBiasCurrentOk returns a tuple with the BiasCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiasCurrent

`func (o *DDMStatusResultVO) SetBiasCurrent(v string)`

SetBiasCurrent sets BiasCurrent field to given value.

### HasBiasCurrent

`func (o *DDMStatusResultVO) HasBiasCurrent() bool`

HasBiasCurrent returns a boolean if a field has been set.

### GetBiasCurrentFlag

`func (o *DDMStatusResultVO) GetBiasCurrentFlag() int32`

GetBiasCurrentFlag returns the BiasCurrentFlag field if non-nil, zero value otherwise.

### GetBiasCurrentFlagOk

`func (o *DDMStatusResultVO) GetBiasCurrentFlagOk() (*int32, bool)`

GetBiasCurrentFlagOk returns a tuple with the BiasCurrentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiasCurrentFlag

`func (o *DDMStatusResultVO) SetBiasCurrentFlag(v int32)`

SetBiasCurrentFlag sets BiasCurrentFlag field to given value.

### HasBiasCurrentFlag

`func (o *DDMStatusResultVO) HasBiasCurrentFlag() bool`

HasBiasCurrentFlag returns a boolean if a field has been set.

### GetDataReady

`func (o *DDMStatusResultVO) GetDataReady() bool`

GetDataReady returns the DataReady field if non-nil, zero value otherwise.

### GetDataReadyOk

`func (o *DDMStatusResultVO) GetDataReadyOk() (*bool, bool)`

GetDataReadyOk returns a tuple with the DataReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReady

`func (o *DDMStatusResultVO) SetDataReady(v bool)`

SetDataReady sets DataReady field to given value.

### HasDataReady

`func (o *DDMStatusResultVO) HasDataReady() bool`

HasDataReady returns a boolean if a field has been set.

### GetLossOfSignal

`func (o *DDMStatusResultVO) GetLossOfSignal() bool`

GetLossOfSignal returns the LossOfSignal field if non-nil, zero value otherwise.

### GetLossOfSignalOk

`func (o *DDMStatusResultVO) GetLossOfSignalOk() (*bool, bool)`

GetLossOfSignalOk returns a tuple with the LossOfSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossOfSignal

`func (o *DDMStatusResultVO) SetLossOfSignal(v bool)`

SetLossOfSignal sets LossOfSignal field to given value.

### HasLossOfSignal

`func (o *DDMStatusResultVO) HasLossOfSignal() bool`

HasLossOfSignal returns a boolean if a field has been set.

### GetPort

`func (o *DDMStatusResultVO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DDMStatusResultVO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DDMStatusResultVO) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *DDMStatusResultVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRxPower

`func (o *DDMStatusResultVO) GetRxPower() string`

GetRxPower returns the RxPower field if non-nil, zero value otherwise.

### GetRxPowerOk

`func (o *DDMStatusResultVO) GetRxPowerOk() (*string, bool)`

GetRxPowerOk returns a tuple with the RxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPower

`func (o *DDMStatusResultVO) SetRxPower(v string)`

SetRxPower sets RxPower field to given value.

### HasRxPower

`func (o *DDMStatusResultVO) HasRxPower() bool`

HasRxPower returns a boolean if a field has been set.

### GetRxPowerFlag

`func (o *DDMStatusResultVO) GetRxPowerFlag() int32`

GetRxPowerFlag returns the RxPowerFlag field if non-nil, zero value otherwise.

### GetRxPowerFlagOk

`func (o *DDMStatusResultVO) GetRxPowerFlagOk() (*int32, bool)`

GetRxPowerFlagOk returns a tuple with the RxPowerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPowerFlag

`func (o *DDMStatusResultVO) SetRxPowerFlag(v int32)`

SetRxPowerFlag sets RxPowerFlag field to given value.

### HasRxPowerFlag

`func (o *DDMStatusResultVO) HasRxPowerFlag() bool`

HasRxPowerFlag returns a boolean if a field has been set.

### GetTemperature

`func (o *DDMStatusResultVO) GetTemperature() string`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *DDMStatusResultVO) GetTemperatureOk() (*string, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *DDMStatusResultVO) SetTemperature(v string)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *DDMStatusResultVO) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTemperatureFlag

`func (o *DDMStatusResultVO) GetTemperatureFlag() int32`

GetTemperatureFlag returns the TemperatureFlag field if non-nil, zero value otherwise.

### GetTemperatureFlagOk

`func (o *DDMStatusResultVO) GetTemperatureFlagOk() (*int32, bool)`

GetTemperatureFlagOk returns a tuple with the TemperatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureFlag

`func (o *DDMStatusResultVO) SetTemperatureFlag(v int32)`

SetTemperatureFlag sets TemperatureFlag field to given value.

### HasTemperatureFlag

`func (o *DDMStatusResultVO) HasTemperatureFlag() bool`

HasTemperatureFlag returns a boolean if a field has been set.

### GetTransmitFault

`func (o *DDMStatusResultVO) GetTransmitFault() int32`

GetTransmitFault returns the TransmitFault field if non-nil, zero value otherwise.

### GetTransmitFaultOk

`func (o *DDMStatusResultVO) GetTransmitFaultOk() (*int32, bool)`

GetTransmitFaultOk returns a tuple with the TransmitFault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitFault

`func (o *DDMStatusResultVO) SetTransmitFault(v int32)`

SetTransmitFault sets TransmitFault field to given value.

### HasTransmitFault

`func (o *DDMStatusResultVO) HasTransmitFault() bool`

HasTransmitFault returns a boolean if a field has been set.

### GetTxPower

`func (o *DDMStatusResultVO) GetTxPower() string`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *DDMStatusResultVO) GetTxPowerOk() (*string, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *DDMStatusResultVO) SetTxPower(v string)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *DDMStatusResultVO) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### GetTxPowerFlag

`func (o *DDMStatusResultVO) GetTxPowerFlag() int32`

GetTxPowerFlag returns the TxPowerFlag field if non-nil, zero value otherwise.

### GetTxPowerFlagOk

`func (o *DDMStatusResultVO) GetTxPowerFlagOk() (*int32, bool)`

GetTxPowerFlagOk returns a tuple with the TxPowerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPowerFlag

`func (o *DDMStatusResultVO) SetTxPowerFlag(v int32)`

SetTxPowerFlag sets TxPowerFlag field to given value.

### HasTxPowerFlag

`func (o *DDMStatusResultVO) HasTxPowerFlag() bool`

HasTxPowerFlag returns a boolean if a field has been set.

### GetVoltage

`func (o *DDMStatusResultVO) GetVoltage() string`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *DDMStatusResultVO) GetVoltageOk() (*string, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *DDMStatusResultVO) SetVoltage(v string)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *DDMStatusResultVO) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetVoltageFlag

`func (o *DDMStatusResultVO) GetVoltageFlag() int32`

GetVoltageFlag returns the VoltageFlag field if non-nil, zero value otherwise.

### GetVoltageFlagOk

`func (o *DDMStatusResultVO) GetVoltageFlagOk() (*int32, bool)`

GetVoltageFlagOk returns a tuple with the VoltageFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltageFlag

`func (o *DDMStatusResultVO) SetVoltageFlag(v int32)`

SetVoltageFlag sets VoltageFlag field to given value.

### HasVoltageFlag

`func (o *DDMStatusResultVO) HasVoltageFlag() bool`

HasVoltageFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


