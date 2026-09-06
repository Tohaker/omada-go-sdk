# SwitchPacketLossVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** | Device name | [optional] 
**LossPkts** | Pointer to **int64** | Total number of lost packets | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**PktsLossPortCnt** | Pointer to **int32** | Number of ports with packet loss | [optional] 
**PortCnt** | Pointer to **int32** | Total number of ports | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewSwitchPacketLossVO

`func NewSwitchPacketLossVO() *SwitchPacketLossVO`

NewSwitchPacketLossVO instantiates a new SwitchPacketLossVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPacketLossVOWithDefaults

`func NewSwitchPacketLossVOWithDefaults() *SwitchPacketLossVO`

NewSwitchPacketLossVOWithDefaults instantiates a new SwitchPacketLossVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *SwitchPacketLossVO) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SwitchPacketLossVO) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SwitchPacketLossVO) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SwitchPacketLossVO) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLossPkts

`func (o *SwitchPacketLossVO) GetLossPkts() int64`

GetLossPkts returns the LossPkts field if non-nil, zero value otherwise.

### GetLossPktsOk

`func (o *SwitchPacketLossVO) GetLossPktsOk() (*int64, bool)`

GetLossPktsOk returns a tuple with the LossPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossPkts

`func (o *SwitchPacketLossVO) SetLossPkts(v int64)`

SetLossPkts sets LossPkts field to given value.

### HasLossPkts

`func (o *SwitchPacketLossVO) HasLossPkts() bool`

HasLossPkts returns a boolean if a field has been set.

### GetMac

`func (o *SwitchPacketLossVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SwitchPacketLossVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SwitchPacketLossVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SwitchPacketLossVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *SwitchPacketLossVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SwitchPacketLossVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SwitchPacketLossVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SwitchPacketLossVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *SwitchPacketLossVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SwitchPacketLossVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SwitchPacketLossVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SwitchPacketLossVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetPktsLossPortCnt

`func (o *SwitchPacketLossVO) GetPktsLossPortCnt() int32`

GetPktsLossPortCnt returns the PktsLossPortCnt field if non-nil, zero value otherwise.

### GetPktsLossPortCntOk

`func (o *SwitchPacketLossVO) GetPktsLossPortCntOk() (*int32, bool)`

GetPktsLossPortCntOk returns a tuple with the PktsLossPortCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPktsLossPortCnt

`func (o *SwitchPacketLossVO) SetPktsLossPortCnt(v int32)`

SetPktsLossPortCnt sets PktsLossPortCnt field to given value.

### HasPktsLossPortCnt

`func (o *SwitchPacketLossVO) HasPktsLossPortCnt() bool`

HasPktsLossPortCnt returns a boolean if a field has been set.

### GetPortCnt

`func (o *SwitchPacketLossVO) GetPortCnt() int32`

GetPortCnt returns the PortCnt field if non-nil, zero value otherwise.

### GetPortCntOk

`func (o *SwitchPacketLossVO) GetPortCntOk() (*int32, bool)`

GetPortCntOk returns a tuple with the PortCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCnt

`func (o *SwitchPacketLossVO) SetPortCnt(v int32)`

SetPortCnt sets PortCnt field to given value.

### HasPortCnt

`func (o *SwitchPacketLossVO) HasPortCnt() bool`

HasPortCnt returns a boolean if a field has been set.

### GetType

`func (o *SwitchPacketLossVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchPacketLossVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchPacketLossVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SwitchPacketLossVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


