# LogNotificationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | Log Notification Enable or Disable Alert | [optional] 
**Email** | Pointer to **bool** | Log Notification Enable or Disable Email | [optional] 
**Event** | Pointer to **bool** | Log Notification Enable or Disable Event | [optional] 
**Key** | Pointer to **string** | Log Notification Key | [optional] 
**ShortMsg** | Pointer to **string** | Log Notification Short Message | [optional] 
**Webhook** | Pointer to **bool** | Log Notification Enable or Disable Webhook (This config applies to the Omada Pro Controller only) | [optional] 

## Methods

### NewLogNotificationOpenApiVO

`func NewLogNotificationOpenApiVO() *LogNotificationOpenApiVO`

NewLogNotificationOpenApiVO instantiates a new LogNotificationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogNotificationOpenApiVOWithDefaults

`func NewLogNotificationOpenApiVOWithDefaults() *LogNotificationOpenApiVO`

NewLogNotificationOpenApiVOWithDefaults instantiates a new LogNotificationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *LogNotificationOpenApiVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *LogNotificationOpenApiVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *LogNotificationOpenApiVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *LogNotificationOpenApiVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetEmail

`func (o *LogNotificationOpenApiVO) GetEmail() bool`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LogNotificationOpenApiVO) GetEmailOk() (*bool, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LogNotificationOpenApiVO) SetEmail(v bool)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LogNotificationOpenApiVO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEvent

`func (o *LogNotificationOpenApiVO) GetEvent() bool`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *LogNotificationOpenApiVO) GetEventOk() (*bool, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *LogNotificationOpenApiVO) SetEvent(v bool)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *LogNotificationOpenApiVO) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetKey

`func (o *LogNotificationOpenApiVO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LogNotificationOpenApiVO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LogNotificationOpenApiVO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LogNotificationOpenApiVO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetShortMsg

`func (o *LogNotificationOpenApiVO) GetShortMsg() string`

GetShortMsg returns the ShortMsg field if non-nil, zero value otherwise.

### GetShortMsgOk

`func (o *LogNotificationOpenApiVO) GetShortMsgOk() (*string, bool)`

GetShortMsgOk returns a tuple with the ShortMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMsg

`func (o *LogNotificationOpenApiVO) SetShortMsg(v string)`

SetShortMsg sets ShortMsg field to given value.

### HasShortMsg

`func (o *LogNotificationOpenApiVO) HasShortMsg() bool`

HasShortMsg returns a boolean if a field has been set.

### GetWebhook

`func (o *LogNotificationOpenApiVO) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *LogNotificationOpenApiVO) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *LogNotificationOpenApiVO) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *LogNotificationOpenApiVO) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


