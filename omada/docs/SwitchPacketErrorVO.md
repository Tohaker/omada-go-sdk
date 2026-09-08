# SwitchPacketErrorVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** | Device name | [optional] 
**ErrPkts** | Pointer to **int64** | Total number of error packets | [optional] 
**ErrPortCnt** | Pointer to **int32** | Number of ports with error packets | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**PortCnt** | Pointer to **int32** | Total number of ports | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewSwitchPacketErrorVO

`func NewSwitchPacketErrorVO() *SwitchPacketErrorVO`

NewSwitchPacketErrorVO instantiates a new SwitchPacketErrorVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPacketErrorVOWithDefaults

`func NewSwitchPacketErrorVOWithDefaults() *SwitchPacketErrorVO`

NewSwitchPacketErrorVOWithDefaults instantiates a new SwitchPacketErrorVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *SwitchPacketErrorVO) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SwitchPacketErrorVO) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SwitchPacketErrorVO) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SwitchPacketErrorVO) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetErrPkts

`func (o *SwitchPacketErrorVO) GetErrPkts() int64`

GetErrPkts returns the ErrPkts field if non-nil, zero value otherwise.

### GetErrPktsOk

`func (o *SwitchPacketErrorVO) GetErrPktsOk() (*int64, bool)`

GetErrPktsOk returns a tuple with the ErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrPkts

`func (o *SwitchPacketErrorVO) SetErrPkts(v int64)`

SetErrPkts sets ErrPkts field to given value.

### HasErrPkts

`func (o *SwitchPacketErrorVO) HasErrPkts() bool`

HasErrPkts returns a boolean if a field has been set.

### GetErrPortCnt

`func (o *SwitchPacketErrorVO) GetErrPortCnt() int32`

GetErrPortCnt returns the ErrPortCnt field if non-nil, zero value otherwise.

### GetErrPortCntOk

`func (o *SwitchPacketErrorVO) GetErrPortCntOk() (*int32, bool)`

GetErrPortCntOk returns a tuple with the ErrPortCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrPortCnt

`func (o *SwitchPacketErrorVO) SetErrPortCnt(v int32)`

SetErrPortCnt sets ErrPortCnt field to given value.

### HasErrPortCnt

`func (o *SwitchPacketErrorVO) HasErrPortCnt() bool`

HasErrPortCnt returns a boolean if a field has been set.

### GetMac

`func (o *SwitchPacketErrorVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SwitchPacketErrorVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SwitchPacketErrorVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SwitchPacketErrorVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *SwitchPacketErrorVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SwitchPacketErrorVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SwitchPacketErrorVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SwitchPacketErrorVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *SwitchPacketErrorVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SwitchPacketErrorVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SwitchPacketErrorVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SwitchPacketErrorVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetPortCnt

`func (o *SwitchPacketErrorVO) GetPortCnt() int32`

GetPortCnt returns the PortCnt field if non-nil, zero value otherwise.

### GetPortCntOk

`func (o *SwitchPacketErrorVO) GetPortCntOk() (*int32, bool)`

GetPortCntOk returns a tuple with the PortCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCnt

`func (o *SwitchPacketErrorVO) SetPortCnt(v int32)`

SetPortCnt sets PortCnt field to given value.

### HasPortCnt

`func (o *SwitchPacketErrorVO) HasPortCnt() bool`

HasPortCnt returns a boolean if a field has been set.

### GetType

`func (o *SwitchPacketErrorVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchPacketErrorVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchPacketErrorVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SwitchPacketErrorVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


