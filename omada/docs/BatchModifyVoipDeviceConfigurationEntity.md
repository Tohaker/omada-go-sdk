# BatchModifyVoipDeviceConfigurationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallBlockingEnable** | Pointer to **bool** | Whether to enable callBlocking. | [optional] 
**CallBlockingProfileId** | Pointer to **string** | The call blocking profile ID of voip device.When callBlockingEnable is true, it can not be null. | [optional] 
**CallBlockingProfileName** | Pointer to **string** | The call blocking profile name of voip device. | [optional] 
**DigitMapProfileId** | Pointer to **string** | The digit map profile ID of voip device. | [optional] 
**DigitMapProfileName** | Pointer to **string** | The digit map profile name of voip device. | [optional] 
**MicGain** | Pointer to **int32** | The mic gain of voip device. | [optional] 
**SpeakerGain** | Pointer to **int32** | The speaker gain of voip device. | [optional] 
**VadSupportEnable** | Pointer to **bool** | Whether to enable vad support. | [optional] 

## Methods

### NewBatchModifyVoipDeviceConfigurationEntity

`func NewBatchModifyVoipDeviceConfigurationEntity() *BatchModifyVoipDeviceConfigurationEntity`

NewBatchModifyVoipDeviceConfigurationEntity instantiates a new BatchModifyVoipDeviceConfigurationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchModifyVoipDeviceConfigurationEntityWithDefaults

`func NewBatchModifyVoipDeviceConfigurationEntityWithDefaults() *BatchModifyVoipDeviceConfigurationEntity`

NewBatchModifyVoipDeviceConfigurationEntityWithDefaults instantiates a new BatchModifyVoipDeviceConfigurationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallBlockingEnable

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetCallBlockingEnable() bool`

GetCallBlockingEnable returns the CallBlockingEnable field if non-nil, zero value otherwise.

### GetCallBlockingEnableOk

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetCallBlockingEnableOk() (*bool, bool)`

GetCallBlockingEnableOk returns a tuple with the CallBlockingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingEnable

`func (o *BatchModifyVoipDeviceConfigurationEntity) SetCallBlockingEnable(v bool)`

SetCallBlockingEnable sets CallBlockingEnable field to given value.

### HasCallBlockingEnable

`func (o *BatchModifyVoipDeviceConfigurationEntity) HasCallBlockingEnable() bool`

HasCallBlockingEnable returns a boolean if a field has been set.

### GetCallBlockingProfileId

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetCallBlockingProfileId() string`

GetCallBlockingProfileId returns the CallBlockingProfileId field if non-nil, zero value otherwise.

### GetCallBlockingProfileIdOk

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetCallBlockingProfileIdOk() (*string, bool)`

GetCallBlockingProfileIdOk returns a tuple with the CallBlockingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingProfileId

`func (o *BatchModifyVoipDeviceConfigurationEntity) SetCallBlockingProfileId(v string)`

SetCallBlockingProfileId sets CallBlockingProfileId field to given value.

### HasCallBlockingProfileId

`func (o *BatchModifyVoipDeviceConfigurationEntity) HasCallBlockingProfileId() bool`

HasCallBlockingProfileId returns a boolean if a field has been set.

### GetCallBlockingProfileName

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetCallBlockingProfileName() string`

GetCallBlockingProfileName returns the CallBlockingProfileName field if non-nil, zero value otherwise.

### GetCallBlockingProfileNameOk

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetCallBlockingProfileNameOk() (*string, bool)`

GetCallBlockingProfileNameOk returns a tuple with the CallBlockingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBlockingProfileName

`func (o *BatchModifyVoipDeviceConfigurationEntity) SetCallBlockingProfileName(v string)`

SetCallBlockingProfileName sets CallBlockingProfileName field to given value.

### HasCallBlockingProfileName

`func (o *BatchModifyVoipDeviceConfigurationEntity) HasCallBlockingProfileName() bool`

HasCallBlockingProfileName returns a boolean if a field has been set.

### GetDigitMapProfileId

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetDigitMapProfileId() string`

GetDigitMapProfileId returns the DigitMapProfileId field if non-nil, zero value otherwise.

### GetDigitMapProfileIdOk

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetDigitMapProfileIdOk() (*string, bool)`

GetDigitMapProfileIdOk returns a tuple with the DigitMapProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitMapProfileId

`func (o *BatchModifyVoipDeviceConfigurationEntity) SetDigitMapProfileId(v string)`

SetDigitMapProfileId sets DigitMapProfileId field to given value.

### HasDigitMapProfileId

`func (o *BatchModifyVoipDeviceConfigurationEntity) HasDigitMapProfileId() bool`

HasDigitMapProfileId returns a boolean if a field has been set.

### GetDigitMapProfileName

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetDigitMapProfileName() string`

GetDigitMapProfileName returns the DigitMapProfileName field if non-nil, zero value otherwise.

### GetDigitMapProfileNameOk

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetDigitMapProfileNameOk() (*string, bool)`

GetDigitMapProfileNameOk returns a tuple with the DigitMapProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitMapProfileName

`func (o *BatchModifyVoipDeviceConfigurationEntity) SetDigitMapProfileName(v string)`

SetDigitMapProfileName sets DigitMapProfileName field to given value.

### HasDigitMapProfileName

`func (o *BatchModifyVoipDeviceConfigurationEntity) HasDigitMapProfileName() bool`

HasDigitMapProfileName returns a boolean if a field has been set.

### GetMicGain

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetMicGain() int32`

GetMicGain returns the MicGain field if non-nil, zero value otherwise.

### GetMicGainOk

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetMicGainOk() (*int32, bool)`

GetMicGainOk returns a tuple with the MicGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicGain

`func (o *BatchModifyVoipDeviceConfigurationEntity) SetMicGain(v int32)`

SetMicGain sets MicGain field to given value.

### HasMicGain

`func (o *BatchModifyVoipDeviceConfigurationEntity) HasMicGain() bool`

HasMicGain returns a boolean if a field has been set.

### GetSpeakerGain

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetSpeakerGain() int32`

GetSpeakerGain returns the SpeakerGain field if non-nil, zero value otherwise.

### GetSpeakerGainOk

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetSpeakerGainOk() (*int32, bool)`

GetSpeakerGainOk returns a tuple with the SpeakerGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerGain

`func (o *BatchModifyVoipDeviceConfigurationEntity) SetSpeakerGain(v int32)`

SetSpeakerGain sets SpeakerGain field to given value.

### HasSpeakerGain

`func (o *BatchModifyVoipDeviceConfigurationEntity) HasSpeakerGain() bool`

HasSpeakerGain returns a boolean if a field has been set.

### GetVadSupportEnable

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetVadSupportEnable() bool`

GetVadSupportEnable returns the VadSupportEnable field if non-nil, zero value otherwise.

### GetVadSupportEnableOk

`func (o *BatchModifyVoipDeviceConfigurationEntity) GetVadSupportEnableOk() (*bool, bool)`

GetVadSupportEnableOk returns a tuple with the VadSupportEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVadSupportEnable

`func (o *BatchModifyVoipDeviceConfigurationEntity) SetVadSupportEnable(v bool)`

SetVadSupportEnable sets VadSupportEnable field to given value.

### HasVadSupportEnable

`func (o *BatchModifyVoipDeviceConfigurationEntity) HasVadSupportEnable() bool`

HasVadSupportEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


