# VoipCallLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTime** | Pointer to **int32** | The dateTime of callLog. | [optional] 
**DeviceNumber** | Pointer to **string** | The deviceNumber of callLog. | [optional] 
**Duration** | Pointer to **int32** | The duration of callLog. | [optional] 
**EntryId** | Pointer to **string** | The entry ID of callLog. | [optional] 
**NumberOrContact** | Pointer to **string** | The number or contact person of callLog. | [optional] 
**Port** | Pointer to **int32** | FXS port id of this call log. | [optional] 
**Status** | Pointer to **int32** | The status of callLog. | [optional] 
**TelephonyDevice** | Pointer to [**CallLogDeviceOpenApiVO**](CallLogDeviceOpenApiVO.md) |  | [optional] 

## Methods

### NewVoipCallLog

`func NewVoipCallLog() *VoipCallLog`

NewVoipCallLog instantiates a new VoipCallLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipCallLogWithDefaults

`func NewVoipCallLogWithDefaults() *VoipCallLog`

NewVoipCallLogWithDefaults instantiates a new VoipCallLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTime

`func (o *VoipCallLog) GetDateTime() int32`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *VoipCallLog) GetDateTimeOk() (*int32, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *VoipCallLog) SetDateTime(v int32)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *VoipCallLog) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetDeviceNumber

`func (o *VoipCallLog) GetDeviceNumber() string`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *VoipCallLog) GetDeviceNumberOk() (*string, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNumber

`func (o *VoipCallLog) SetDeviceNumber(v string)`

SetDeviceNumber sets DeviceNumber field to given value.

### HasDeviceNumber

`func (o *VoipCallLog) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### GetDuration

`func (o *VoipCallLog) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoipCallLog) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoipCallLog) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoipCallLog) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEntryId

`func (o *VoipCallLog) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *VoipCallLog) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *VoipCallLog) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *VoipCallLog) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetNumberOrContact

`func (o *VoipCallLog) GetNumberOrContact() string`

GetNumberOrContact returns the NumberOrContact field if non-nil, zero value otherwise.

### GetNumberOrContactOk

`func (o *VoipCallLog) GetNumberOrContactOk() (*string, bool)`

GetNumberOrContactOk returns a tuple with the NumberOrContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOrContact

`func (o *VoipCallLog) SetNumberOrContact(v string)`

SetNumberOrContact sets NumberOrContact field to given value.

### HasNumberOrContact

`func (o *VoipCallLog) HasNumberOrContact() bool`

HasNumberOrContact returns a boolean if a field has been set.

### GetPort

`func (o *VoipCallLog) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VoipCallLog) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VoipCallLog) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VoipCallLog) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *VoipCallLog) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoipCallLog) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoipCallLog) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VoipCallLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTelephonyDevice

`func (o *VoipCallLog) GetTelephonyDevice() CallLogDeviceOpenApiVO`

GetTelephonyDevice returns the TelephonyDevice field if non-nil, zero value otherwise.

### GetTelephonyDeviceOk

`func (o *VoipCallLog) GetTelephonyDeviceOk() (*CallLogDeviceOpenApiVO, bool)`

GetTelephonyDeviceOk returns a tuple with the TelephonyDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephonyDevice

`func (o *VoipCallLog) SetTelephonyDevice(v CallLogDeviceOpenApiVO)`

SetTelephonyDevice sets TelephonyDevice field to given value.

### HasTelephonyDevice

`func (o *VoipCallLog) HasTelephonyDevice() bool`

HasTelephonyDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


