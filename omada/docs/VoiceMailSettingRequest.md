# VoiceMailSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | Parameter [duration] should be from 20 to 120. | [optional] 
**Enable** | **bool** | Whether to enable voice mail settings. | 
**GreetingForVoiceMailMode** | Pointer to **int32** | Parameter [greetingForVoiceMailMode] should be 0 or 1. 0: Default, 1: Custom. | [optional] 
**GreetingName** | Pointer to **string** | Greeting name. | [optional] 
**NoAnswerTime** | Pointer to **int32** | The no answer time of telephone number. Parameter [noAnswerTime] should be from 5 to 60 | [optional] 
**OmadacId** | Pointer to **string** | Omadac ID. | [optional] 
**RemoteAccessEnable** | Pointer to **bool** | Whether voice mail allows remote access. | [optional] 
**RemoteAccessPin** | Pointer to **string** | Remote access pin. | [optional] 
**SiteId** | Pointer to **string** | Site ID. | [optional] 
**UsbUuid** | Pointer to **string** | USB UUID. | [optional] 
**VoiceMailCapacity** | Pointer to **int32** | The capacity of voice mail. | [optional] 
**VoiceMailInUsb** | Pointer to **bool** | Whether voice mail is in USB. | [optional] 

## Methods

### NewVoiceMailSettingRequest

`func NewVoiceMailSettingRequest(enable bool, ) *VoiceMailSettingRequest`

NewVoiceMailSettingRequest instantiates a new VoiceMailSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceMailSettingRequestWithDefaults

`func NewVoiceMailSettingRequestWithDefaults() *VoiceMailSettingRequest`

NewVoiceMailSettingRequestWithDefaults instantiates a new VoiceMailSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *VoiceMailSettingRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoiceMailSettingRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoiceMailSettingRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoiceMailSettingRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnable

`func (o *VoiceMailSettingRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *VoiceMailSettingRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *VoiceMailSettingRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetGreetingForVoiceMailMode

`func (o *VoiceMailSettingRequest) GetGreetingForVoiceMailMode() int32`

GetGreetingForVoiceMailMode returns the GreetingForVoiceMailMode field if non-nil, zero value otherwise.

### GetGreetingForVoiceMailModeOk

`func (o *VoiceMailSettingRequest) GetGreetingForVoiceMailModeOk() (*int32, bool)`

GetGreetingForVoiceMailModeOk returns a tuple with the GreetingForVoiceMailMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingForVoiceMailMode

`func (o *VoiceMailSettingRequest) SetGreetingForVoiceMailMode(v int32)`

SetGreetingForVoiceMailMode sets GreetingForVoiceMailMode field to given value.

### HasGreetingForVoiceMailMode

`func (o *VoiceMailSettingRequest) HasGreetingForVoiceMailMode() bool`

HasGreetingForVoiceMailMode returns a boolean if a field has been set.

### GetGreetingName

`func (o *VoiceMailSettingRequest) GetGreetingName() string`

GetGreetingName returns the GreetingName field if non-nil, zero value otherwise.

### GetGreetingNameOk

`func (o *VoiceMailSettingRequest) GetGreetingNameOk() (*string, bool)`

GetGreetingNameOk returns a tuple with the GreetingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingName

`func (o *VoiceMailSettingRequest) SetGreetingName(v string)`

SetGreetingName sets GreetingName field to given value.

### HasGreetingName

`func (o *VoiceMailSettingRequest) HasGreetingName() bool`

HasGreetingName returns a boolean if a field has been set.

### GetNoAnswerTime

`func (o *VoiceMailSettingRequest) GetNoAnswerTime() int32`

GetNoAnswerTime returns the NoAnswerTime field if non-nil, zero value otherwise.

### GetNoAnswerTimeOk

`func (o *VoiceMailSettingRequest) GetNoAnswerTimeOk() (*int32, bool)`

GetNoAnswerTimeOk returns a tuple with the NoAnswerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAnswerTime

`func (o *VoiceMailSettingRequest) SetNoAnswerTime(v int32)`

SetNoAnswerTime sets NoAnswerTime field to given value.

### HasNoAnswerTime

`func (o *VoiceMailSettingRequest) HasNoAnswerTime() bool`

HasNoAnswerTime returns a boolean if a field has been set.

### GetOmadacId

`func (o *VoiceMailSettingRequest) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *VoiceMailSettingRequest) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *VoiceMailSettingRequest) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *VoiceMailSettingRequest) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetRemoteAccessEnable

`func (o *VoiceMailSettingRequest) GetRemoteAccessEnable() bool`

GetRemoteAccessEnable returns the RemoteAccessEnable field if non-nil, zero value otherwise.

### GetRemoteAccessEnableOk

`func (o *VoiceMailSettingRequest) GetRemoteAccessEnableOk() (*bool, bool)`

GetRemoteAccessEnableOk returns a tuple with the RemoteAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessEnable

`func (o *VoiceMailSettingRequest) SetRemoteAccessEnable(v bool)`

SetRemoteAccessEnable sets RemoteAccessEnable field to given value.

### HasRemoteAccessEnable

`func (o *VoiceMailSettingRequest) HasRemoteAccessEnable() bool`

HasRemoteAccessEnable returns a boolean if a field has been set.

### GetRemoteAccessPin

`func (o *VoiceMailSettingRequest) GetRemoteAccessPin() string`

GetRemoteAccessPin returns the RemoteAccessPin field if non-nil, zero value otherwise.

### GetRemoteAccessPinOk

`func (o *VoiceMailSettingRequest) GetRemoteAccessPinOk() (*string, bool)`

GetRemoteAccessPinOk returns a tuple with the RemoteAccessPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessPin

`func (o *VoiceMailSettingRequest) SetRemoteAccessPin(v string)`

SetRemoteAccessPin sets RemoteAccessPin field to given value.

### HasRemoteAccessPin

`func (o *VoiceMailSettingRequest) HasRemoteAccessPin() bool`

HasRemoteAccessPin returns a boolean if a field has been set.

### GetSiteId

`func (o *VoiceMailSettingRequest) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VoiceMailSettingRequest) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VoiceMailSettingRequest) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VoiceMailSettingRequest) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUsbUuid

`func (o *VoiceMailSettingRequest) GetUsbUuid() string`

GetUsbUuid returns the UsbUuid field if non-nil, zero value otherwise.

### GetUsbUuidOk

`func (o *VoiceMailSettingRequest) GetUsbUuidOk() (*string, bool)`

GetUsbUuidOk returns a tuple with the UsbUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbUuid

`func (o *VoiceMailSettingRequest) SetUsbUuid(v string)`

SetUsbUuid sets UsbUuid field to given value.

### HasUsbUuid

`func (o *VoiceMailSettingRequest) HasUsbUuid() bool`

HasUsbUuid returns a boolean if a field has been set.

### GetVoiceMailCapacity

`func (o *VoiceMailSettingRequest) GetVoiceMailCapacity() int32`

GetVoiceMailCapacity returns the VoiceMailCapacity field if non-nil, zero value otherwise.

### GetVoiceMailCapacityOk

`func (o *VoiceMailSettingRequest) GetVoiceMailCapacityOk() (*int32, bool)`

GetVoiceMailCapacityOk returns a tuple with the VoiceMailCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMailCapacity

`func (o *VoiceMailSettingRequest) SetVoiceMailCapacity(v int32)`

SetVoiceMailCapacity sets VoiceMailCapacity field to given value.

### HasVoiceMailCapacity

`func (o *VoiceMailSettingRequest) HasVoiceMailCapacity() bool`

HasVoiceMailCapacity returns a boolean if a field has been set.

### GetVoiceMailInUsb

`func (o *VoiceMailSettingRequest) GetVoiceMailInUsb() bool`

GetVoiceMailInUsb returns the VoiceMailInUsb field if non-nil, zero value otherwise.

### GetVoiceMailInUsbOk

`func (o *VoiceMailSettingRequest) GetVoiceMailInUsbOk() (*bool, bool)`

GetVoiceMailInUsbOk returns a tuple with the VoiceMailInUsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMailInUsb

`func (o *VoiceMailSettingRequest) SetVoiceMailInUsb(v bool)`

SetVoiceMailInUsb sets VoiceMailInUsb field to given value.

### HasVoiceMailInUsb

`func (o *VoiceMailSettingRequest) HasVoiceMailInUsb() bool`

HasVoiceMailInUsb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


