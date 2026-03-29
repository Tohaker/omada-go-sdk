# LogNotificationSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmailSetting** | Pointer to [**LogAlertEmailOpenApiVO**](LogAlertEmailOpenApiVO.md) |  | [optional] 
**LogNotifications** | Pointer to [**[]LogNotificationOpenApiVO**](LogNotificationOpenApiVO.md) | Log Notification List | [optional] 
**WebhookConfig** | Pointer to [**WebhookConfigOpenApiVO**](WebhookConfigOpenApiVO.md) |  | [optional] 

## Methods

### NewLogNotificationSettingOpenApiVO

`func NewLogNotificationSettingOpenApiVO() *LogNotificationSettingOpenApiVO`

NewLogNotificationSettingOpenApiVO instantiates a new LogNotificationSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogNotificationSettingOpenApiVOWithDefaults

`func NewLogNotificationSettingOpenApiVOWithDefaults() *LogNotificationSettingOpenApiVO`

NewLogNotificationSettingOpenApiVOWithDefaults instantiates a new LogNotificationSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmailSetting

`func (o *LogNotificationSettingOpenApiVO) GetAlertEmailSetting() LogAlertEmailOpenApiVO`

GetAlertEmailSetting returns the AlertEmailSetting field if non-nil, zero value otherwise.

### GetAlertEmailSettingOk

`func (o *LogNotificationSettingOpenApiVO) GetAlertEmailSettingOk() (*LogAlertEmailOpenApiVO, bool)`

GetAlertEmailSettingOk returns a tuple with the AlertEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmailSetting

`func (o *LogNotificationSettingOpenApiVO) SetAlertEmailSetting(v LogAlertEmailOpenApiVO)`

SetAlertEmailSetting sets AlertEmailSetting field to given value.

### HasAlertEmailSetting

`func (o *LogNotificationSettingOpenApiVO) HasAlertEmailSetting() bool`

HasAlertEmailSetting returns a boolean if a field has been set.

### GetLogNotifications

`func (o *LogNotificationSettingOpenApiVO) GetLogNotifications() []LogNotificationOpenApiVO`

GetLogNotifications returns the LogNotifications field if non-nil, zero value otherwise.

### GetLogNotificationsOk

`func (o *LogNotificationSettingOpenApiVO) GetLogNotificationsOk() (*[]LogNotificationOpenApiVO, bool)`

GetLogNotificationsOk returns a tuple with the LogNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotifications

`func (o *LogNotificationSettingOpenApiVO) SetLogNotifications(v []LogNotificationOpenApiVO)`

SetLogNotifications sets LogNotifications field to given value.

### HasLogNotifications

`func (o *LogNotificationSettingOpenApiVO) HasLogNotifications() bool`

HasLogNotifications returns a boolean if a field has been set.

### GetWebhookConfig

`func (o *LogNotificationSettingOpenApiVO) GetWebhookConfig() WebhookConfigOpenApiVO`

GetWebhookConfig returns the WebhookConfig field if non-nil, zero value otherwise.

### GetWebhookConfigOk

`func (o *LogNotificationSettingOpenApiVO) GetWebhookConfigOk() (*WebhookConfigOpenApiVO, bool)`

GetWebhookConfigOk returns a tuple with the WebhookConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfig

`func (o *LogNotificationSettingOpenApiVO) SetWebhookConfig(v WebhookConfigOpenApiVO)`

SetWebhookConfig sets WebhookConfig field to given value.

### HasWebhookConfig

`func (o *LogNotificationSettingOpenApiVO) HasWebhookConfig() bool`

HasWebhookConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


