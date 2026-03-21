# VoiceMailSettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGreetingName** | Pointer to **string** | Greeting name. | [optional] 
**Duration** | Pointer to **int32** | Voice mail duration. | [optional] 
**Enable** | Pointer to **bool** | Whether voice mail is enable. | [optional] 
**GreetingForVoiceMailMode** | Pointer to **int32** | Parameter [greetingForVoiceMailMode] will be 0 or 1. 0: Default, 1: Custom. | [optional] 
**NoAnswerTime** | Pointer to **int32** | The no answer time of telephone number. | [optional] 
**OmadacId** | Pointer to **string** | Omadac ID. | [optional] 
**RemoteAccessEnable** | Pointer to **bool** | Whether voice mail allows remote access. | [optional] 
**RemoteAccessPin** | Pointer to **string** | Remote access pin. | [optional] 
**SiteId** | Pointer to **string** | Site ID. | [optional] 
**UsbList** | Pointer to [**[]UsbInfo**](UsbInfo.md) | USB info list. | [optional] 
**VoiceMailInUsb** | Pointer to **bool** | Whether voice mail is in USB. | [optional] 

## Methods

### NewVoiceMailSettingResponse

`func NewVoiceMailSettingResponse() *VoiceMailSettingResponse`

NewVoiceMailSettingResponse instantiates a new VoiceMailSettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceMailSettingResponseWithDefaults

`func NewVoiceMailSettingResponseWithDefaults() *VoiceMailSettingResponse`

NewVoiceMailSettingResponseWithDefaults instantiates a new VoiceMailSettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGreetingName

`func (o *VoiceMailSettingResponse) GetDefaultGreetingName() string`

GetDefaultGreetingName returns the DefaultGreetingName field if non-nil, zero value otherwise.

### GetDefaultGreetingNameOk

`func (o *VoiceMailSettingResponse) GetDefaultGreetingNameOk() (*string, bool)`

GetDefaultGreetingNameOk returns a tuple with the DefaultGreetingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGreetingName

`func (o *VoiceMailSettingResponse) SetDefaultGreetingName(v string)`

SetDefaultGreetingName sets DefaultGreetingName field to given value.

### HasDefaultGreetingName

`func (o *VoiceMailSettingResponse) HasDefaultGreetingName() bool`

HasDefaultGreetingName returns a boolean if a field has been set.

### GetDuration

`func (o *VoiceMailSettingResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoiceMailSettingResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoiceMailSettingResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoiceMailSettingResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnable

`func (o *VoiceMailSettingResponse) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *VoiceMailSettingResponse) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *VoiceMailSettingResponse) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *VoiceMailSettingResponse) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGreetingForVoiceMailMode

`func (o *VoiceMailSettingResponse) GetGreetingForVoiceMailMode() int32`

GetGreetingForVoiceMailMode returns the GreetingForVoiceMailMode field if non-nil, zero value otherwise.

### GetGreetingForVoiceMailModeOk

`func (o *VoiceMailSettingResponse) GetGreetingForVoiceMailModeOk() (*int32, bool)`

GetGreetingForVoiceMailModeOk returns a tuple with the GreetingForVoiceMailMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingForVoiceMailMode

`func (o *VoiceMailSettingResponse) SetGreetingForVoiceMailMode(v int32)`

SetGreetingForVoiceMailMode sets GreetingForVoiceMailMode field to given value.

### HasGreetingForVoiceMailMode

`func (o *VoiceMailSettingResponse) HasGreetingForVoiceMailMode() bool`

HasGreetingForVoiceMailMode returns a boolean if a field has been set.

### GetNoAnswerTime

`func (o *VoiceMailSettingResponse) GetNoAnswerTime() int32`

GetNoAnswerTime returns the NoAnswerTime field if non-nil, zero value otherwise.

### GetNoAnswerTimeOk

`func (o *VoiceMailSettingResponse) GetNoAnswerTimeOk() (*int32, bool)`

GetNoAnswerTimeOk returns a tuple with the NoAnswerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAnswerTime

`func (o *VoiceMailSettingResponse) SetNoAnswerTime(v int32)`

SetNoAnswerTime sets NoAnswerTime field to given value.

### HasNoAnswerTime

`func (o *VoiceMailSettingResponse) HasNoAnswerTime() bool`

HasNoAnswerTime returns a boolean if a field has been set.

### GetOmadacId

`func (o *VoiceMailSettingResponse) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *VoiceMailSettingResponse) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *VoiceMailSettingResponse) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *VoiceMailSettingResponse) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetRemoteAccessEnable

`func (o *VoiceMailSettingResponse) GetRemoteAccessEnable() bool`

GetRemoteAccessEnable returns the RemoteAccessEnable field if non-nil, zero value otherwise.

### GetRemoteAccessEnableOk

`func (o *VoiceMailSettingResponse) GetRemoteAccessEnableOk() (*bool, bool)`

GetRemoteAccessEnableOk returns a tuple with the RemoteAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessEnable

`func (o *VoiceMailSettingResponse) SetRemoteAccessEnable(v bool)`

SetRemoteAccessEnable sets RemoteAccessEnable field to given value.

### HasRemoteAccessEnable

`func (o *VoiceMailSettingResponse) HasRemoteAccessEnable() bool`

HasRemoteAccessEnable returns a boolean if a field has been set.

### GetRemoteAccessPin

`func (o *VoiceMailSettingResponse) GetRemoteAccessPin() string`

GetRemoteAccessPin returns the RemoteAccessPin field if non-nil, zero value otherwise.

### GetRemoteAccessPinOk

`func (o *VoiceMailSettingResponse) GetRemoteAccessPinOk() (*string, bool)`

GetRemoteAccessPinOk returns a tuple with the RemoteAccessPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessPin

`func (o *VoiceMailSettingResponse) SetRemoteAccessPin(v string)`

SetRemoteAccessPin sets RemoteAccessPin field to given value.

### HasRemoteAccessPin

`func (o *VoiceMailSettingResponse) HasRemoteAccessPin() bool`

HasRemoteAccessPin returns a boolean if a field has been set.

### GetSiteId

`func (o *VoiceMailSettingResponse) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VoiceMailSettingResponse) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VoiceMailSettingResponse) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VoiceMailSettingResponse) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUsbList

`func (o *VoiceMailSettingResponse) GetUsbList() []UsbInfo`

GetUsbList returns the UsbList field if non-nil, zero value otherwise.

### GetUsbListOk

`func (o *VoiceMailSettingResponse) GetUsbListOk() (*[]UsbInfo, bool)`

GetUsbListOk returns a tuple with the UsbList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbList

`func (o *VoiceMailSettingResponse) SetUsbList(v []UsbInfo)`

SetUsbList sets UsbList field to given value.

### HasUsbList

`func (o *VoiceMailSettingResponse) HasUsbList() bool`

HasUsbList returns a boolean if a field has been set.

### GetVoiceMailInUsb

`func (o *VoiceMailSettingResponse) GetVoiceMailInUsb() bool`

GetVoiceMailInUsb returns the VoiceMailInUsb field if non-nil, zero value otherwise.

### GetVoiceMailInUsbOk

`func (o *VoiceMailSettingResponse) GetVoiceMailInUsbOk() (*bool, bool)`

GetVoiceMailInUsbOk returns a tuple with the VoiceMailInUsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMailInUsb

`func (o *VoiceMailSettingResponse) SetVoiceMailInUsb(v bool)`

SetVoiceMailInUsb sets VoiceMailInUsb field to given value.

### HasVoiceMailInUsb

`func (o *VoiceMailSettingResponse) HasVoiceMailInUsb() bool`

HasVoiceMailInUsb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


