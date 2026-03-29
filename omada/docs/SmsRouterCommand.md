# SmsRouterCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**AccessEnable** | **bool** | Whether to enable the allow list of the above functions, and only allow users in the list to interact with the device. | 
**Reboot** | Pointer to [**Reboot**](Reboot.md) |  | [optional] 
**RebootEnable** | **bool** | Whether to enable the control of device restart via SMS. | 
**Resource** | Pointer to **int32** | The SMS router command setting creation resource, such as: 0: new created, 1: from template, 2: override. | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**StatusEnable** | **bool** | Whether to enable Viewing device-related information and WAN port-related information via SMS. | 

## Methods

### NewSmsRouterCommand

`func NewSmsRouterCommand(accessEnable bool, rebootEnable bool, statusEnable bool, ) *SmsRouterCommand`

NewSmsRouterCommand instantiates a new SmsRouterCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsRouterCommandWithDefaults

`func NewSmsRouterCommandWithDefaults() *SmsRouterCommand`

NewSmsRouterCommandWithDefaults instantiates a new SmsRouterCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *SmsRouterCommand) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SmsRouterCommand) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SmsRouterCommand) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SmsRouterCommand) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAccessEnable

`func (o *SmsRouterCommand) GetAccessEnable() bool`

GetAccessEnable returns the AccessEnable field if non-nil, zero value otherwise.

### GetAccessEnableOk

`func (o *SmsRouterCommand) GetAccessEnableOk() (*bool, bool)`

GetAccessEnableOk returns a tuple with the AccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessEnable

`func (o *SmsRouterCommand) SetAccessEnable(v bool)`

SetAccessEnable sets AccessEnable field to given value.


### GetReboot

`func (o *SmsRouterCommand) GetReboot() Reboot`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *SmsRouterCommand) GetRebootOk() (*Reboot, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *SmsRouterCommand) SetReboot(v Reboot)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *SmsRouterCommand) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRebootEnable

`func (o *SmsRouterCommand) GetRebootEnable() bool`

GetRebootEnable returns the RebootEnable field if non-nil, zero value otherwise.

### GetRebootEnableOk

`func (o *SmsRouterCommand) GetRebootEnableOk() (*bool, bool)`

GetRebootEnableOk returns a tuple with the RebootEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootEnable

`func (o *SmsRouterCommand) SetRebootEnable(v bool)`

SetRebootEnable sets RebootEnable field to given value.


### GetResource

`func (o *SmsRouterCommand) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SmsRouterCommand) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SmsRouterCommand) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *SmsRouterCommand) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStatus

`func (o *SmsRouterCommand) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsRouterCommand) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsRouterCommand) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsRouterCommand) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusEnable

`func (o *SmsRouterCommand) GetStatusEnable() bool`

GetStatusEnable returns the StatusEnable field if non-nil, zero value otherwise.

### GetStatusEnableOk

`func (o *SmsRouterCommand) GetStatusEnableOk() (*bool, bool)`

GetStatusEnableOk returns a tuple with the StatusEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEnable

`func (o *SmsRouterCommand) SetStatusEnable(v bool)`

SetStatusEnable sets StatusEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


