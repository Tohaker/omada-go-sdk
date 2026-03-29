# LogNotificationEditOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | **bool** | Log Notification Enable or Disable Alert | 
**Email** | **bool** | Log Notification Enable or Disable Email | 
**Event** | **bool** | Log Notification Enable or Disable Event | 
**Key** | **string** | For the values of Log Notification Key, refer to section 5.6.1 of the Open API Access | 
**Webhook** | Pointer to **bool** | Log Notification Category Enable or Disable Webhook (This config applies to the Omada Pro Controller only and should not bu null) | [optional] 

## Methods

### NewLogNotificationEditOpenApiVO

`func NewLogNotificationEditOpenApiVO(alert bool, email bool, event bool, key string, ) *LogNotificationEditOpenApiVO`

NewLogNotificationEditOpenApiVO instantiates a new LogNotificationEditOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogNotificationEditOpenApiVOWithDefaults

`func NewLogNotificationEditOpenApiVOWithDefaults() *LogNotificationEditOpenApiVO`

NewLogNotificationEditOpenApiVOWithDefaults instantiates a new LogNotificationEditOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *LogNotificationEditOpenApiVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *LogNotificationEditOpenApiVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *LogNotificationEditOpenApiVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.


### GetEmail

`func (o *LogNotificationEditOpenApiVO) GetEmail() bool`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LogNotificationEditOpenApiVO) GetEmailOk() (*bool, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LogNotificationEditOpenApiVO) SetEmail(v bool)`

SetEmail sets Email field to given value.


### GetEvent

`func (o *LogNotificationEditOpenApiVO) GetEvent() bool`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *LogNotificationEditOpenApiVO) GetEventOk() (*bool, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *LogNotificationEditOpenApiVO) SetEvent(v bool)`

SetEvent sets Event field to given value.


### GetKey

`func (o *LogNotificationEditOpenApiVO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LogNotificationEditOpenApiVO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LogNotificationEditOpenApiVO) SetKey(v string)`

SetKey sets Key field to given value.


### GetWebhook

`func (o *LogNotificationEditOpenApiVO) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *LogNotificationEditOpenApiVO) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *LogNotificationEditOpenApiVO) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *LogNotificationEditOpenApiVO) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


