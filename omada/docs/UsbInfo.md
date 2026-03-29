# UsbInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomGreetingName** | Pointer to **string** | Custom greeting name. | [optional] 
**RemainAvailableCapacity** | Pointer to **int32** | Remain available capacity of voice mail box. | [optional] 
**Selected** | Pointer to **bool** | Whether selected the USB device. | [optional] 
**UsbName** | Pointer to **string** | USB name. | [optional] 
**UsbUuid** | Pointer to **string** | USB UUID. | [optional] 
**VoiceMailboxCapacity** | Pointer to **int32** | The capacity of voice mail box. | [optional] 

## Methods

### NewUsbInfo

`func NewUsbInfo() *UsbInfo`

NewUsbInfo instantiates a new UsbInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsbInfoWithDefaults

`func NewUsbInfoWithDefaults() *UsbInfo`

NewUsbInfoWithDefaults instantiates a new UsbInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomGreetingName

`func (o *UsbInfo) GetCustomGreetingName() string`

GetCustomGreetingName returns the CustomGreetingName field if non-nil, zero value otherwise.

### GetCustomGreetingNameOk

`func (o *UsbInfo) GetCustomGreetingNameOk() (*string, bool)`

GetCustomGreetingNameOk returns a tuple with the CustomGreetingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGreetingName

`func (o *UsbInfo) SetCustomGreetingName(v string)`

SetCustomGreetingName sets CustomGreetingName field to given value.

### HasCustomGreetingName

`func (o *UsbInfo) HasCustomGreetingName() bool`

HasCustomGreetingName returns a boolean if a field has been set.

### GetRemainAvailableCapacity

`func (o *UsbInfo) GetRemainAvailableCapacity() int32`

GetRemainAvailableCapacity returns the RemainAvailableCapacity field if non-nil, zero value otherwise.

### GetRemainAvailableCapacityOk

`func (o *UsbInfo) GetRemainAvailableCapacityOk() (*int32, bool)`

GetRemainAvailableCapacityOk returns a tuple with the RemainAvailableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainAvailableCapacity

`func (o *UsbInfo) SetRemainAvailableCapacity(v int32)`

SetRemainAvailableCapacity sets RemainAvailableCapacity field to given value.

### HasRemainAvailableCapacity

`func (o *UsbInfo) HasRemainAvailableCapacity() bool`

HasRemainAvailableCapacity returns a boolean if a field has been set.

### GetSelected

`func (o *UsbInfo) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *UsbInfo) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *UsbInfo) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *UsbInfo) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetUsbName

`func (o *UsbInfo) GetUsbName() string`

GetUsbName returns the UsbName field if non-nil, zero value otherwise.

### GetUsbNameOk

`func (o *UsbInfo) GetUsbNameOk() (*string, bool)`

GetUsbNameOk returns a tuple with the UsbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbName

`func (o *UsbInfo) SetUsbName(v string)`

SetUsbName sets UsbName field to given value.

### HasUsbName

`func (o *UsbInfo) HasUsbName() bool`

HasUsbName returns a boolean if a field has been set.

### GetUsbUuid

`func (o *UsbInfo) GetUsbUuid() string`

GetUsbUuid returns the UsbUuid field if non-nil, zero value otherwise.

### GetUsbUuidOk

`func (o *UsbInfo) GetUsbUuidOk() (*string, bool)`

GetUsbUuidOk returns a tuple with the UsbUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbUuid

`func (o *UsbInfo) SetUsbUuid(v string)`

SetUsbUuid sets UsbUuid field to given value.

### HasUsbUuid

`func (o *UsbInfo) HasUsbUuid() bool`

HasUsbUuid returns a boolean if a field has been set.

### GetVoiceMailboxCapacity

`func (o *UsbInfo) GetVoiceMailboxCapacity() int32`

GetVoiceMailboxCapacity returns the VoiceMailboxCapacity field if non-nil, zero value otherwise.

### GetVoiceMailboxCapacityOk

`func (o *UsbInfo) GetVoiceMailboxCapacityOk() (*int32, bool)`

GetVoiceMailboxCapacityOk returns a tuple with the VoiceMailboxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMailboxCapacity

`func (o *UsbInfo) SetVoiceMailboxCapacity(v int32)`

SetVoiceMailboxCapacity sets VoiceMailboxCapacity field to given value.

### HasVoiceMailboxCapacity

`func (o *UsbInfo) HasVoiceMailboxCapacity() bool`

HasVoiceMailboxCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


