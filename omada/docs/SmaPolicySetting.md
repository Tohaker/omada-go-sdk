# SmaPolicySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboxPolicy** | **int32** | Policy parameter [inboxPolicy] should be from 0 to 2. 0: If SMS inbox/outbox is full, delete the oldest read SMS; 1: If SMS inbox/outbox is full, send e-mail alert to Administrator;2: If SMS inbox/outbox is full, Forward new SMS with e-mail to Administrator. | 
**MailServer** | Pointer to **string** | Mail server ID. | [optional] 
**Resource** | Pointer to **int32** | SMS policy setting creation resource, such as: 0: new created, 1: from template, 2: override. | [optional] 

## Methods

### NewSmaPolicySetting

`func NewSmaPolicySetting(inboxPolicy int32, ) *SmaPolicySetting`

NewSmaPolicySetting instantiates a new SmaPolicySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmaPolicySettingWithDefaults

`func NewSmaPolicySettingWithDefaults() *SmaPolicySetting`

NewSmaPolicySettingWithDefaults instantiates a new SmaPolicySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInboxPolicy

`func (o *SmaPolicySetting) GetInboxPolicy() int32`

GetInboxPolicy returns the InboxPolicy field if non-nil, zero value otherwise.

### GetInboxPolicyOk

`func (o *SmaPolicySetting) GetInboxPolicyOk() (*int32, bool)`

GetInboxPolicyOk returns a tuple with the InboxPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxPolicy

`func (o *SmaPolicySetting) SetInboxPolicy(v int32)`

SetInboxPolicy sets InboxPolicy field to given value.


### GetMailServer

`func (o *SmaPolicySetting) GetMailServer() string`

GetMailServer returns the MailServer field if non-nil, zero value otherwise.

### GetMailServerOk

`func (o *SmaPolicySetting) GetMailServerOk() (*string, bool)`

GetMailServerOk returns a tuple with the MailServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailServer

`func (o *SmaPolicySetting) SetMailServer(v string)`

SetMailServer sets MailServer field to given value.

### HasMailServer

`func (o *SmaPolicySetting) HasMailServer() bool`

HasMailServer returns a boolean if a field has been set.

### GetResource

`func (o *SmaPolicySetting) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SmaPolicySetting) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SmaPolicySetting) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *SmaPolicySetting) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


