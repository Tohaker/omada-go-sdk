# Reboot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | To reboot the router via SMS, send a message starting with \&quot;LTE Router Reboot\&quot;, followed by Password/PIN(e.g. LTE Router Reboot 1234). Password is required for the control of device restart via SMS. | [optional] 

## Methods

### NewReboot

`func NewReboot() *Reboot`

NewReboot instantiates a new Reboot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootWithDefaults

`func NewRebootWithDefaults() *Reboot`

NewRebootWithDefaults instantiates a new Reboot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *Reboot) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Reboot) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Reboot) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Reboot) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


