# ApDownLinkStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **float64** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **bool** |  | [optional] 
**DownLinkPort** | Pointer to **string** |  | [optional] 
**Duplex** | Pointer to **int32** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **int32** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelVersion** | Pointer to **string** |  | [optional] 
**PoeState** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**PortType** | Pointer to **int32** |  | [optional] 
**Rx** | Pointer to **int64** |  | [optional] 
**RxDropPkts** | Pointer to **int64** |  | [optional] 
**RxErrPkts** | Pointer to **int64** |  | [optional] 
**RxPkts** | Pointer to **int64** |  | [optional] 
**RxPower** | Pointer to **float64** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**Temp** | Pointer to **float64** |  | [optional] 
**Tx** | Pointer to **int64** |  | [optional] 
**TxDropPkts** | Pointer to **int64** |  | [optional] 
**TxErrPkts** | Pointer to **int64** |  | [optional] 
**TxPkts** | Pointer to **int64** |  | [optional] 
**TxPower** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VoipState** | Pointer to **int32** |  | [optional] 
**Voltage** | Pointer to **float64** |  | [optional] 

## Methods

### NewApDownLinkStatusVO

`func NewApDownLinkStatusVO() *ApDownLinkStatusVO`

NewApDownLinkStatusVO instantiates a new ApDownLinkStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApDownLinkStatusVOWithDefaults

`func NewApDownLinkStatusVOWithDefaults() *ApDownLinkStatusVO`

NewApDownLinkStatusVOWithDefaults instantiates a new ApDownLinkStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *ApDownLinkStatusVO) GetCurrent() float64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ApDownLinkStatusVO) GetCurrentOk() (*float64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ApDownLinkStatusVO) SetCurrent(v float64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *ApDownLinkStatusVO) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetDeviceName

`func (o *ApDownLinkStatusVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ApDownLinkStatusVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ApDownLinkStatusVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ApDownLinkStatusVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDisplay

`func (o *ApDownLinkStatusVO) GetDisplay() bool`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ApDownLinkStatusVO) GetDisplayOk() (*bool, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ApDownLinkStatusVO) SetDisplay(v bool)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ApDownLinkStatusVO) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDownLinkPort

`func (o *ApDownLinkStatusVO) GetDownLinkPort() string`

GetDownLinkPort returns the DownLinkPort field if non-nil, zero value otherwise.

### GetDownLinkPortOk

`func (o *ApDownLinkStatusVO) GetDownLinkPortOk() (*string, bool)`

GetDownLinkPortOk returns a tuple with the DownLinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLinkPort

`func (o *ApDownLinkStatusVO) SetDownLinkPort(v string)`

SetDownLinkPort sets DownLinkPort field to given value.

### HasDownLinkPort

`func (o *ApDownLinkStatusVO) HasDownLinkPort() bool`

HasDownLinkPort returns a boolean if a field has been set.

### GetDuplex

`func (o *ApDownLinkStatusVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *ApDownLinkStatusVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *ApDownLinkStatusVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *ApDownLinkStatusVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetIp

`func (o *ApDownLinkStatusVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApDownLinkStatusVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApDownLinkStatusVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApDownLinkStatusVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLink

`func (o *ApDownLinkStatusVO) GetLink() int32`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ApDownLinkStatusVO) GetLinkOk() (*int32, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ApDownLinkStatusVO) SetLink(v int32)`

SetLink sets Link field to given value.

### HasLink

`func (o *ApDownLinkStatusVO) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMac

`func (o *ApDownLinkStatusVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApDownLinkStatusVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApDownLinkStatusVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ApDownLinkStatusVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ApDownLinkStatusVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApDownLinkStatusVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApDownLinkStatusVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApDownLinkStatusVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApDownLinkStatusVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApDownLinkStatusVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApDownLinkStatusVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApDownLinkStatusVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetPoeState

`func (o *ApDownLinkStatusVO) GetPoeState() int32`

GetPoeState returns the PoeState field if non-nil, zero value otherwise.

### GetPoeStateOk

`func (o *ApDownLinkStatusVO) GetPoeStateOk() (*int32, bool)`

GetPoeStateOk returns a tuple with the PoeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeState

`func (o *ApDownLinkStatusVO) SetPoeState(v int32)`

SetPoeState sets PoeState field to given value.

### HasPoeState

`func (o *ApDownLinkStatusVO) HasPoeState() bool`

HasPoeState returns a boolean if a field has been set.

### GetPort

`func (o *ApDownLinkStatusVO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApDownLinkStatusVO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApDownLinkStatusVO) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApDownLinkStatusVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortType

`func (o *ApDownLinkStatusVO) GetPortType() int32`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *ApDownLinkStatusVO) GetPortTypeOk() (*int32, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *ApDownLinkStatusVO) SetPortType(v int32)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *ApDownLinkStatusVO) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetRx

`func (o *ApDownLinkStatusVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *ApDownLinkStatusVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *ApDownLinkStatusVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *ApDownLinkStatusVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *ApDownLinkStatusVO) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *ApDownLinkStatusVO) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *ApDownLinkStatusVO) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *ApDownLinkStatusVO) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *ApDownLinkStatusVO) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *ApDownLinkStatusVO) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *ApDownLinkStatusVO) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *ApDownLinkStatusVO) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetRxPkts

`func (o *ApDownLinkStatusVO) GetRxPkts() int64`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *ApDownLinkStatusVO) GetRxPktsOk() (*int64, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *ApDownLinkStatusVO) SetRxPkts(v int64)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *ApDownLinkStatusVO) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetRxPower

`func (o *ApDownLinkStatusVO) GetRxPower() float64`

GetRxPower returns the RxPower field if non-nil, zero value otherwise.

### GetRxPowerOk

`func (o *ApDownLinkStatusVO) GetRxPowerOk() (*float64, bool)`

GetRxPowerOk returns a tuple with the RxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPower

`func (o *ApDownLinkStatusVO) SetRxPower(v float64)`

SetRxPower sets RxPower field to given value.

### HasRxPower

`func (o *ApDownLinkStatusVO) HasRxPower() bool`

HasRxPower returns a boolean if a field has been set.

### GetSpeed

`func (o *ApDownLinkStatusVO) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ApDownLinkStatusVO) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ApDownLinkStatusVO) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ApDownLinkStatusVO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTemp

`func (o *ApDownLinkStatusVO) GetTemp() float64`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *ApDownLinkStatusVO) GetTempOk() (*float64, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *ApDownLinkStatusVO) SetTemp(v float64)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *ApDownLinkStatusVO) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetTx

`func (o *ApDownLinkStatusVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ApDownLinkStatusVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ApDownLinkStatusVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *ApDownLinkStatusVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *ApDownLinkStatusVO) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *ApDownLinkStatusVO) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *ApDownLinkStatusVO) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *ApDownLinkStatusVO) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *ApDownLinkStatusVO) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *ApDownLinkStatusVO) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *ApDownLinkStatusVO) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *ApDownLinkStatusVO) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetTxPkts

`func (o *ApDownLinkStatusVO) GetTxPkts() int64`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *ApDownLinkStatusVO) GetTxPktsOk() (*int64, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *ApDownLinkStatusVO) SetTxPkts(v int64)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *ApDownLinkStatusVO) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetTxPower

`func (o *ApDownLinkStatusVO) GetTxPower() float64`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *ApDownLinkStatusVO) GetTxPowerOk() (*float64, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *ApDownLinkStatusVO) SetTxPower(v float64)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *ApDownLinkStatusVO) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### GetType

`func (o *ApDownLinkStatusVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApDownLinkStatusVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApDownLinkStatusVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApDownLinkStatusVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVoipState

`func (o *ApDownLinkStatusVO) GetVoipState() int32`

GetVoipState returns the VoipState field if non-nil, zero value otherwise.

### GetVoipStateOk

`func (o *ApDownLinkStatusVO) GetVoipStateOk() (*int32, bool)`

GetVoipStateOk returns a tuple with the VoipState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipState

`func (o *ApDownLinkStatusVO) SetVoipState(v int32)`

SetVoipState sets VoipState field to given value.

### HasVoipState

`func (o *ApDownLinkStatusVO) HasVoipState() bool`

HasVoipState returns a boolean if a field has been set.

### GetVoltage

`func (o *ApDownLinkStatusVO) GetVoltage() float64`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *ApDownLinkStatusVO) GetVoltageOk() (*float64, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *ApDownLinkStatusVO) SetVoltage(v float64)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *ApDownLinkStatusVO) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


