# OswDDMInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiasCurrent** | Pointer to **float64** | The biasCurrent of the optical module. | [optional] 
**DataReady** | Pointer to **int32** | Whether the DDM data is valid. | [optional] 
**LossOfSignal** | Pointer to **int32** | The Loss Of Signal(LOS) of the optical module. | [optional] 
**Port** | Pointer to **int32** | Switch portId. | [optional] 
**RxPower** | Pointer to **float64** | The received optical power of the optical module. | [optional] 
**StandardPort** | Pointer to **string** | Switch stack portId(unit/slot/port). | [optional] 
**Temperature** | Pointer to **float64** | The temperature of the optical module. | [optional] 
**TransmitFault** | Pointer to **int32** | The transmission fault of the optical module. | [optional] 
**TxPower** | Pointer to **float64** | The transmitted optical power of the optical module. | [optional] 
**Voltage** | Pointer to **float64** | The voltage of the optical module. | [optional] 

## Methods

### NewOswDDMInfoOpenApiVO

`func NewOswDDMInfoOpenApiVO() *OswDDMInfoOpenApiVO`

NewOswDDMInfoOpenApiVO instantiates a new OswDDMInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDDMInfoOpenApiVOWithDefaults

`func NewOswDDMInfoOpenApiVOWithDefaults() *OswDDMInfoOpenApiVO`

NewOswDDMInfoOpenApiVOWithDefaults instantiates a new OswDDMInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiasCurrent

`func (o *OswDDMInfoOpenApiVO) GetBiasCurrent() float64`

GetBiasCurrent returns the BiasCurrent field if non-nil, zero value otherwise.

### GetBiasCurrentOk

`func (o *OswDDMInfoOpenApiVO) GetBiasCurrentOk() (*float64, bool)`

GetBiasCurrentOk returns a tuple with the BiasCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiasCurrent

`func (o *OswDDMInfoOpenApiVO) SetBiasCurrent(v float64)`

SetBiasCurrent sets BiasCurrent field to given value.

### HasBiasCurrent

`func (o *OswDDMInfoOpenApiVO) HasBiasCurrent() bool`

HasBiasCurrent returns a boolean if a field has been set.

### GetDataReady

`func (o *OswDDMInfoOpenApiVO) GetDataReady() int32`

GetDataReady returns the DataReady field if non-nil, zero value otherwise.

### GetDataReadyOk

`func (o *OswDDMInfoOpenApiVO) GetDataReadyOk() (*int32, bool)`

GetDataReadyOk returns a tuple with the DataReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReady

`func (o *OswDDMInfoOpenApiVO) SetDataReady(v int32)`

SetDataReady sets DataReady field to given value.

### HasDataReady

`func (o *OswDDMInfoOpenApiVO) HasDataReady() bool`

HasDataReady returns a boolean if a field has been set.

### GetLossOfSignal

`func (o *OswDDMInfoOpenApiVO) GetLossOfSignal() int32`

GetLossOfSignal returns the LossOfSignal field if non-nil, zero value otherwise.

### GetLossOfSignalOk

`func (o *OswDDMInfoOpenApiVO) GetLossOfSignalOk() (*int32, bool)`

GetLossOfSignalOk returns a tuple with the LossOfSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossOfSignal

`func (o *OswDDMInfoOpenApiVO) SetLossOfSignal(v int32)`

SetLossOfSignal sets LossOfSignal field to given value.

### HasLossOfSignal

`func (o *OswDDMInfoOpenApiVO) HasLossOfSignal() bool`

HasLossOfSignal returns a boolean if a field has been set.

### GetPort

`func (o *OswDDMInfoOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswDDMInfoOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswDDMInfoOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswDDMInfoOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRxPower

`func (o *OswDDMInfoOpenApiVO) GetRxPower() float64`

GetRxPower returns the RxPower field if non-nil, zero value otherwise.

### GetRxPowerOk

`func (o *OswDDMInfoOpenApiVO) GetRxPowerOk() (*float64, bool)`

GetRxPowerOk returns a tuple with the RxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPower

`func (o *OswDDMInfoOpenApiVO) SetRxPower(v float64)`

SetRxPower sets RxPower field to given value.

### HasRxPower

`func (o *OswDDMInfoOpenApiVO) HasRxPower() bool`

HasRxPower returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswDDMInfoOpenApiVO) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswDDMInfoOpenApiVO) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswDDMInfoOpenApiVO) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswDDMInfoOpenApiVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetTemperature

`func (o *OswDDMInfoOpenApiVO) GetTemperature() float64`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *OswDDMInfoOpenApiVO) GetTemperatureOk() (*float64, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *OswDDMInfoOpenApiVO) SetTemperature(v float64)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *OswDDMInfoOpenApiVO) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTransmitFault

`func (o *OswDDMInfoOpenApiVO) GetTransmitFault() int32`

GetTransmitFault returns the TransmitFault field if non-nil, zero value otherwise.

### GetTransmitFaultOk

`func (o *OswDDMInfoOpenApiVO) GetTransmitFaultOk() (*int32, bool)`

GetTransmitFaultOk returns a tuple with the TransmitFault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitFault

`func (o *OswDDMInfoOpenApiVO) SetTransmitFault(v int32)`

SetTransmitFault sets TransmitFault field to given value.

### HasTransmitFault

`func (o *OswDDMInfoOpenApiVO) HasTransmitFault() bool`

HasTransmitFault returns a boolean if a field has been set.

### GetTxPower

`func (o *OswDDMInfoOpenApiVO) GetTxPower() float64`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *OswDDMInfoOpenApiVO) GetTxPowerOk() (*float64, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *OswDDMInfoOpenApiVO) SetTxPower(v float64)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *OswDDMInfoOpenApiVO) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### GetVoltage

`func (o *OswDDMInfoOpenApiVO) GetVoltage() float64`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *OswDDMInfoOpenApiVO) GetVoltageOk() (*float64, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *OswDDMInfoOpenApiVO) SetVoltage(v float64)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *OswDDMInfoOpenApiVO) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


