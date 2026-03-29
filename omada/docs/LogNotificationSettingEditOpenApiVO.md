# LogNotificationSettingEditOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmailSetting** | Pointer to [**LogAlertEmailOpenApiVO**](LogAlertEmailOpenApiVO.md) |  | [optional] 
**LogNotifications** | [**[]LogNotificationEditOpenApiVO**](LogNotificationEditOpenApiVO.md) | Log Notification List | 
**WebhookConfig** | Pointer to [**WebhookConfigEditOpenApiVO**](WebhookConfigEditOpenApiVO.md) |  | [optional] 

## Methods

### NewLogNotificationSettingEditOpenApiVO

`func NewLogNotificationSettingEditOpenApiVO(logNotifications []LogNotificationEditOpenApiVO, ) *LogNotificationSettingEditOpenApiVO`

NewLogNotificationSettingEditOpenApiVO instantiates a new LogNotificationSettingEditOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogNotificationSettingEditOpenApiVOWithDefaults

`func NewLogNotificationSettingEditOpenApiVOWithDefaults() *LogNotificationSettingEditOpenApiVO`

NewLogNotificationSettingEditOpenApiVOWithDefaults instantiates a new LogNotificationSettingEditOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmailSetting

`func (o *LogNotificationSettingEditOpenApiVO) GetAlertEmailSetting() LogAlertEmailOpenApiVO`

GetAlertEmailSetting returns the AlertEmailSetting field if non-nil, zero value otherwise.

### GetAlertEmailSettingOk

`func (o *LogNotificationSettingEditOpenApiVO) GetAlertEmailSettingOk() (*LogAlertEmailOpenApiVO, bool)`

GetAlertEmailSettingOk returns a tuple with the AlertEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmailSetting

`func (o *LogNotificationSettingEditOpenApiVO) SetAlertEmailSetting(v LogAlertEmailOpenApiVO)`

SetAlertEmailSetting sets AlertEmailSetting field to given value.

### HasAlertEmailSetting

`func (o *LogNotificationSettingEditOpenApiVO) HasAlertEmailSetting() bool`

HasAlertEmailSetting returns a boolean if a field has been set.

### GetLogNotifications

`func (o *LogNotificationSettingEditOpenApiVO) GetLogNotifications() []LogNotificationEditOpenApiVO`

GetLogNotifications returns the LogNotifications field if non-nil, zero value otherwise.

### GetLogNotificationsOk

`func (o *LogNotificationSettingEditOpenApiVO) GetLogNotificationsOk() (*[]LogNotificationEditOpenApiVO, bool)`

GetLogNotificationsOk returns a tuple with the LogNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotifications

`func (o *LogNotificationSettingEditOpenApiVO) SetLogNotifications(v []LogNotificationEditOpenApiVO)`

SetLogNotifications sets LogNotifications field to given value.


### GetWebhookConfig

`func (o *LogNotificationSettingEditOpenApiVO) GetWebhookConfig() WebhookConfigEditOpenApiVO`

GetWebhookConfig returns the WebhookConfig field if non-nil, zero value otherwise.

### GetWebhookConfigOk

`func (o *LogNotificationSettingEditOpenApiVO) GetWebhookConfigOk() (*WebhookConfigEditOpenApiVO, bool)`

GetWebhookConfigOk returns a tuple with the WebhookConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfig

`func (o *LogNotificationSettingEditOpenApiVO) SetWebhookConfig(v WebhookConfigEditOpenApiVO)`

SetWebhookConfig sets WebhookConfig field to given value.

### HasWebhookConfig

`func (o *LogNotificationSettingEditOpenApiVO) HasWebhookConfig() bool`

HasWebhookConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


