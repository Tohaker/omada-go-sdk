# VoipDeviceApConfigurationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallBlockingEnable** | Pointer to **bool** | Whether to enable callBlocking. | [optional] 
**CallBlockingProfileId** | Pointer to **string** | The call blocking profile ID of voip device.When callBlockingEnable is true, it can not be null. | [optional] 
**CallBlockingProfileName** | Pointer to **string** | The call blocking profile name of voip device. | [optional] 
**DigitMapProfileId** | Pointer to **string** | The digit map profile ID of voip device. | [optional] 
**DigitMapProfileName** | Pointer to **string** | The digit map profile name of voip device. | [optional] 
**MicGain** | Pointer to **int32** | The mic gain of voip device. | [optional] 
**NumberIdForOutgoingCalls** | Pointer to **string** | The outgoing calls number ID of voip device. | [optional] 
**NumberIdListForIncomingCalls** | Pointer to **[]string** | The list of incoming calls number ID. | [optional] 
**SpeakerGain** | Pointer to **int32** | The speaker gain of voip device. | [optional] 
**VadSupportEnable** | Pointer to **bool** | Whether to enable vad support. | [optional] 

## Methods

### NewVoipDeviceApConfigurationOpenApiVO

`func NewVoipDeviceApConfigurationOpenApiVO() *VoipDeviceApConfigurationOpenApiVO`

NewVoipDeviceApConfigurationOpenApiVO instantiates a new VoipDeviceApConfigurationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipDeviceApConfigurationOpenApiVOWithDefaults

`func NewVoipDeviceApConfigurationOpenApiVOWithDefaults() *VoipDeviceApConfigurationOpenApiVO`

NewVoipDeviceApConfigurationOpenApiVOWithDefaults instantiates a new VoipDeviceApConfigurationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallBlockingEnable

`func (o *VoipDeviceApConfigurationOpenApiVO) GetCallBlockingEnable() bool`

GetCallBlockingEnable returns the CallBlockingEnable field if non-nil, zero value otherwise.

### GetCallBlockingEnableOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetCallBlockingEnableOk() (*bool, bool)`

GetCallBlockingEnableOk returns a tuple with the CallBlockingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingEnable

`func (o *VoipDeviceApConfigurationOpenApiVO) SetCallBlockingEnable(v bool)`

SetCallBlockingEnable sets CallBlockingEnable field to given value.

### HasCallBlockingEnable

`func (o *VoipDeviceApConfigurationOpenApiVO) HasCallBlockingEnable() bool`

HasCallBlockingEnable returns a boolean if a field has been set.

### GetCallBlockingProfileId

`func (o *VoipDeviceApConfigurationOpenApiVO) GetCallBlockingProfileId() string`

GetCallBlockingProfileId returns the CallBlockingProfileId field if non-nil, zero value otherwise.

### GetCallBlockingProfileIdOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetCallBlockingProfileIdOk() (*string, bool)`

GetCallBlockingProfileIdOk returns a tuple with the CallBlockingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingProfileId

`func (o *VoipDeviceApConfigurationOpenApiVO) SetCallBlockingProfileId(v string)`

SetCallBlockingProfileId sets CallBlockingProfileId field to given value.

### HasCallBlockingProfileId

`func (o *VoipDeviceApConfigurationOpenApiVO) HasCallBlockingProfileId() bool`

HasCallBlockingProfileId returns a boolean if a field has been set.

### GetCallBlockingProfileName

`func (o *VoipDeviceApConfigurationOpenApiVO) GetCallBlockingProfileName() string`

GetCallBlockingProfileName returns the CallBlockingProfileName field if non-nil, zero value otherwise.

### GetCallBlockingProfileNameOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetCallBlockingProfileNameOk() (*string, bool)`

GetCallBlockingProfileNameOk returns a tuple with the CallBlockingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingProfileName

`func (o *VoipDeviceApConfigurationOpenApiVO) SetCallBlockingProfileName(v string)`

SetCallBlockingProfileName sets CallBlockingProfileName field to given value.

### HasCallBlockingProfileName

`func (o *VoipDeviceApConfigurationOpenApiVO) HasCallBlockingProfileName() bool`

HasCallBlockingProfileName returns a boolean if a field has been set.

### GetDigitMapProfileId

`func (o *VoipDeviceApConfigurationOpenApiVO) GetDigitMapProfileId() string`

GetDigitMapProfileId returns the DigitMapProfileId field if non-nil, zero value otherwise.

### GetDigitMapProfileIdOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetDigitMapProfileIdOk() (*string, bool)`

GetDigitMapProfileIdOk returns a tuple with the DigitMapProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitMapProfileId

`func (o *VoipDeviceApConfigurationOpenApiVO) SetDigitMapProfileId(v string)`

SetDigitMapProfileId sets DigitMapProfileId field to given value.

### HasDigitMapProfileId

`func (o *VoipDeviceApConfigurationOpenApiVO) HasDigitMapProfileId() bool`

HasDigitMapProfileId returns a boolean if a field has been set.

### GetDigitMapProfileName

`func (o *VoipDeviceApConfigurationOpenApiVO) GetDigitMapProfileName() string`

GetDigitMapProfileName returns the DigitMapProfileName field if non-nil, zero value otherwise.

### GetDigitMapProfileNameOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetDigitMapProfileNameOk() (*string, bool)`

GetDigitMapProfileNameOk returns a tuple with the DigitMapProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitMapProfileName

`func (o *VoipDeviceApConfigurationOpenApiVO) SetDigitMapProfileName(v string)`

SetDigitMapProfileName sets DigitMapProfileName field to given value.

### HasDigitMapProfileName

`func (o *VoipDeviceApConfigurationOpenApiVO) HasDigitMapProfileName() bool`

HasDigitMapProfileName returns a boolean if a field has been set.

### GetMicGain

`func (o *VoipDeviceApConfigurationOpenApiVO) GetMicGain() int32`

GetMicGain returns the MicGain field if non-nil, zero value otherwise.

### GetMicGainOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetMicGainOk() (*int32, bool)`

GetMicGainOk returns a tuple with the MicGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicGain

`func (o *VoipDeviceApConfigurationOpenApiVO) SetMicGain(v int32)`

SetMicGain sets MicGain field to given value.

### HasMicGain

`func (o *VoipDeviceApConfigurationOpenApiVO) HasMicGain() bool`

HasMicGain returns a boolean if a field has been set.

### GetNumberIdForOutgoingCalls

`func (o *VoipDeviceApConfigurationOpenApiVO) GetNumberIdForOutgoingCalls() string`

GetNumberIdForOutgoingCalls returns the NumberIdForOutgoingCalls field if non-nil, zero value otherwise.

### GetNumberIdForOutgoingCallsOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetNumberIdForOutgoingCallsOk() (*string, bool)`

GetNumberIdForOutgoingCallsOk returns a tuple with the NumberIdForOutgoingCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberIdForOutgoingCalls

`func (o *VoipDeviceApConfigurationOpenApiVO) SetNumberIdForOutgoingCalls(v string)`

SetNumberIdForOutgoingCalls sets NumberIdForOutgoingCalls field to given value.

### HasNumberIdForOutgoingCalls

`func (o *VoipDeviceApConfigurationOpenApiVO) HasNumberIdForOutgoingCalls() bool`

HasNumberIdForOutgoingCalls returns a boolean if a field has been set.

### GetNumberIdListForIncomingCalls

`func (o *VoipDeviceApConfigurationOpenApiVO) GetNumberIdListForIncomingCalls() []string`

GetNumberIdListForIncomingCalls returns the NumberIdListForIncomingCalls field if non-nil, zero value otherwise.

### GetNumberIdListForIncomingCallsOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetNumberIdListForIncomingCallsOk() (*[]string, bool)`

GetNumberIdListForIncomingCallsOk returns a tuple with the NumberIdListForIncomingCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberIdListForIncomingCalls

`func (o *VoipDeviceApConfigurationOpenApiVO) SetNumberIdListForIncomingCalls(v []string)`

SetNumberIdListForIncomingCalls sets NumberIdListForIncomingCalls field to given value.

### HasNumberIdListForIncomingCalls

`func (o *VoipDeviceApConfigurationOpenApiVO) HasNumberIdListForIncomingCalls() bool`

HasNumberIdListForIncomingCalls returns a boolean if a field has been set.

### GetSpeakerGain

`func (o *VoipDeviceApConfigurationOpenApiVO) GetSpeakerGain() int32`

GetSpeakerGain returns the SpeakerGain field if non-nil, zero value otherwise.

### GetSpeakerGainOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetSpeakerGainOk() (*int32, bool)`

GetSpeakerGainOk returns a tuple with the SpeakerGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerGain

`func (o *VoipDeviceApConfigurationOpenApiVO) SetSpeakerGain(v int32)`

SetSpeakerGain sets SpeakerGain field to given value.

### HasSpeakerGain

`func (o *VoipDeviceApConfigurationOpenApiVO) HasSpeakerGain() bool`

HasSpeakerGain returns a boolean if a field has been set.

### GetVadSupportEnable

`func (o *VoipDeviceApConfigurationOpenApiVO) GetVadSupportEnable() bool`

GetVadSupportEnable returns the VadSupportEnable field if non-nil, zero value otherwise.

### GetVadSupportEnableOk

`func (o *VoipDeviceApConfigurationOpenApiVO) GetVadSupportEnableOk() (*bool, bool)`

GetVadSupportEnableOk returns a tuple with the VadSupportEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVadSupportEnable

`func (o *VoipDeviceApConfigurationOpenApiVO) SetVadSupportEnable(v bool)`

SetVadSupportEnable sets VadSupportEnable field to given value.

### HasVadSupportEnable

`func (o *VoipDeviceApConfigurationOpenApiVO) HasVadSupportEnable() bool`

HasVadSupportEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


