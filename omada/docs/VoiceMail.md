# VoiceMail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **int64** | The date or time of voice mail. | [optional] 
**Duration** | Pointer to **int32** | The duration of voice mail. | [optional] 
**Id** | Pointer to **string** | Voice mail ID. | [optional] 
**IncomingNumber** | Pointer to **string** | The incoming number of voice mail. | [optional] 
**Read** | Pointer to **bool** | Whether the voice mail has been read, false-not read, true-read | [optional] 
**TelephoneNumber** | Pointer to **string** | The telephone number of voice mail. | [optional] 

## Methods

### NewVoiceMail

`func NewVoiceMail() *VoiceMail`

NewVoiceMail instantiates a new VoiceMail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceMailWithDefaults

`func NewVoiceMailWithDefaults() *VoiceMail`

NewVoiceMailWithDefaults instantiates a new VoiceMail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *VoiceMail) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *VoiceMail) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *VoiceMail) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *VoiceMail) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDuration

`func (o *VoiceMail) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoiceMail) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoiceMail) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoiceMail) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *VoiceMail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoiceMail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoiceMail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoiceMail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncomingNumber

`func (o *VoiceMail) GetIncomingNumber() string`

GetIncomingNumber returns the IncomingNumber field if non-nil, zero value otherwise.

### GetIncomingNumberOk

`func (o *VoiceMail) GetIncomingNumberOk() (*string, bool)`

GetIncomingNumberOk returns a tuple with the IncomingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingNumber

`func (o *VoiceMail) SetIncomingNumber(v string)`

SetIncomingNumber sets IncomingNumber field to given value.

### HasIncomingNumber

`func (o *VoiceMail) HasIncomingNumber() bool`

HasIncomingNumber returns a boolean if a field has been set.

### GetRead

`func (o *VoiceMail) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *VoiceMail) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *VoiceMail) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *VoiceMail) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *VoiceMail) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *VoiceMail) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *VoiceMail) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *VoiceMail) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


