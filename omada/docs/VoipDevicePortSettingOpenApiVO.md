# VoipDevicePortSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auto** | Pointer to **bool** | Whether to enable auto. | [optional] 
**MicGain** | Pointer to **int32** | The mic gain of voip device. | [optional] 
**NumberIdForOutgoingCalls** | Pointer to **string** | The outgoing calls number ID of voip device. | [optional] 
**NumberIdListForIncomingCalls** | Pointer to **[]string** | The list of incoming calls number ID. | [optional] 
**Port** | Pointer to **int32** | Port ID. | [optional] 
**SpeakerGain** | Pointer to **int32** | The speaker gain of voip device. | [optional] 
**VadSupportEnable** | Pointer to **bool** | Whether to enable vad support. | [optional] 

## Methods

### NewVoipDevicePortSettingOpenApiVO

`func NewVoipDevicePortSettingOpenApiVO() *VoipDevicePortSettingOpenApiVO`

NewVoipDevicePortSettingOpenApiVO instantiates a new VoipDevicePortSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipDevicePortSettingOpenApiVOWithDefaults

`func NewVoipDevicePortSettingOpenApiVOWithDefaults() *VoipDevicePortSettingOpenApiVO`

NewVoipDevicePortSettingOpenApiVOWithDefaults instantiates a new VoipDevicePortSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuto

`func (o *VoipDevicePortSettingOpenApiVO) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *VoipDevicePortSettingOpenApiVO) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *VoipDevicePortSettingOpenApiVO) SetAuto(v bool)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *VoipDevicePortSettingOpenApiVO) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### GetMicGain

`func (o *VoipDevicePortSettingOpenApiVO) GetMicGain() int32`

GetMicGain returns the MicGain field if non-nil, zero value otherwise.

### GetMicGainOk

`func (o *VoipDevicePortSettingOpenApiVO) GetMicGainOk() (*int32, bool)`

GetMicGainOk returns a tuple with the MicGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicGain

`func (o *VoipDevicePortSettingOpenApiVO) SetMicGain(v int32)`

SetMicGain sets MicGain field to given value.

### HasMicGain

`func (o *VoipDevicePortSettingOpenApiVO) HasMicGain() bool`

HasMicGain returns a boolean if a field has been set.

### GetNumberIdForOutgoingCalls

`func (o *VoipDevicePortSettingOpenApiVO) GetNumberIdForOutgoingCalls() string`

GetNumberIdForOutgoingCalls returns the NumberIdForOutgoingCalls field if non-nil, zero value otherwise.

### GetNumberIdForOutgoingCallsOk

`func (o *VoipDevicePortSettingOpenApiVO) GetNumberIdForOutgoingCallsOk() (*string, bool)`

GetNumberIdForOutgoingCallsOk returns a tuple with the NumberIdForOutgoingCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberIdForOutgoingCalls

`func (o *VoipDevicePortSettingOpenApiVO) SetNumberIdForOutgoingCalls(v string)`

SetNumberIdForOutgoingCalls sets NumberIdForOutgoingCalls field to given value.

### HasNumberIdForOutgoingCalls

`func (o *VoipDevicePortSettingOpenApiVO) HasNumberIdForOutgoingCalls() bool`

HasNumberIdForOutgoingCalls returns a boolean if a field has been set.

### GetNumberIdListForIncomingCalls

`func (o *VoipDevicePortSettingOpenApiVO) GetNumberIdListForIncomingCalls() []string`

GetNumberIdListForIncomingCalls returns the NumberIdListForIncomingCalls field if non-nil, zero value otherwise.

### GetNumberIdListForIncomingCallsOk

`func (o *VoipDevicePortSettingOpenApiVO) GetNumberIdListForIncomingCallsOk() (*[]string, bool)`

GetNumberIdListForIncomingCallsOk returns a tuple with the NumberIdListForIncomingCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberIdListForIncomingCalls

`func (o *VoipDevicePortSettingOpenApiVO) SetNumberIdListForIncomingCalls(v []string)`

SetNumberIdListForIncomingCalls sets NumberIdListForIncomingCalls field to given value.

### HasNumberIdListForIncomingCalls

`func (o *VoipDevicePortSettingOpenApiVO) HasNumberIdListForIncomingCalls() bool`

HasNumberIdListForIncomingCalls returns a boolean if a field has been set.

### GetPort

`func (o *VoipDevicePortSettingOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VoipDevicePortSettingOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VoipDevicePortSettingOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VoipDevicePortSettingOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSpeakerGain

`func (o *VoipDevicePortSettingOpenApiVO) GetSpeakerGain() int32`

GetSpeakerGain returns the SpeakerGain field if non-nil, zero value otherwise.

### GetSpeakerGainOk

`func (o *VoipDevicePortSettingOpenApiVO) GetSpeakerGainOk() (*int32, bool)`

GetSpeakerGainOk returns a tuple with the SpeakerGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerGain

`func (o *VoipDevicePortSettingOpenApiVO) SetSpeakerGain(v int32)`

SetSpeakerGain sets SpeakerGain field to given value.

### HasSpeakerGain

`func (o *VoipDevicePortSettingOpenApiVO) HasSpeakerGain() bool`

HasSpeakerGain returns a boolean if a field has been set.

### GetVadSupportEnable

`func (o *VoipDevicePortSettingOpenApiVO) GetVadSupportEnable() bool`

GetVadSupportEnable returns the VadSupportEnable field if non-nil, zero value otherwise.

### GetVadSupportEnableOk

`func (o *VoipDevicePortSettingOpenApiVO) GetVadSupportEnableOk() (*bool, bool)`

GetVadSupportEnableOk returns a tuple with the VadSupportEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVadSupportEnable

`func (o *VoipDevicePortSettingOpenApiVO) SetVadSupportEnable(v bool)`

SetVadSupportEnable sets VadSupportEnable field to given value.

### HasVadSupportEnable

`func (o *VoipDevicePortSettingOpenApiVO) HasVadSupportEnable() bool`

HasVadSupportEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


