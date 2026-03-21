# LogNotificationSettingEditOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmailSetting** | Pointer to [**LogAlertEmailOpenApiVO**](LogAlertEmailOpenApiVO.md) |  | [optional] 
**AlertNotifications** | Pointer to [**[]LogNotificationEditOpenApiV2VO**](LogNotificationEditOpenApiV2VO.md) | Alert Notification List | [optional] 
**EventEmailSetting** | Pointer to [**LogAlertEmailOpenApiVO**](LogAlertEmailOpenApiVO.md) |  | [optional] 
**EventNotifications** | Pointer to [**[]LogNotificationEditOpenApiV2VO**](LogNotificationEditOpenApiV2VO.md) | Event Notification List | [optional] 
**WebhookConfig** | Pointer to [**WebhookConfigEditOpenApiVO**](WebhookConfigEditOpenApiVO.md) |  | [optional] 

## Methods

### NewLogNotificationSettingEditOpenApiV2VO

`func NewLogNotificationSettingEditOpenApiV2VO() *LogNotificationSettingEditOpenApiV2VO`

NewLogNotificationSettingEditOpenApiV2VO instantiates a new LogNotificationSettingEditOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogNotificationSettingEditOpenApiV2VOWithDefaults

`func NewLogNotificationSettingEditOpenApiV2VOWithDefaults() *LogNotificationSettingEditOpenApiV2VO`

NewLogNotificationSettingEditOpenApiV2VOWithDefaults instantiates a new LogNotificationSettingEditOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmailSetting

`func (o *LogNotificationSettingEditOpenApiV2VO) GetAlertEmailSetting() LogAlertEmailOpenApiVO`

GetAlertEmailSetting returns the AlertEmailSetting field if non-nil, zero value otherwise.

### GetAlertEmailSettingOk

`func (o *LogNotificationSettingEditOpenApiV2VO) GetAlertEmailSettingOk() (*LogAlertEmailOpenApiVO, bool)`

GetAlertEmailSettingOk returns a tuple with the AlertEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmailSetting

`func (o *LogNotificationSettingEditOpenApiV2VO) SetAlertEmailSetting(v LogAlertEmailOpenApiVO)`

SetAlertEmailSetting sets AlertEmailSetting field to given value.

### HasAlertEmailSetting

`func (o *LogNotificationSettingEditOpenApiV2VO) HasAlertEmailSetting() bool`

HasAlertEmailSetting returns a boolean if a field has been set.

### GetAlertNotifications

`func (o *LogNotificationSettingEditOpenApiV2VO) GetAlertNotifications() []LogNotificationEditOpenApiV2VO`

GetAlertNotifications returns the AlertNotifications field if non-nil, zero value otherwise.

### GetAlertNotificationsOk

`func (o *LogNotificationSettingEditOpenApiV2VO) GetAlertNotificationsOk() (*[]LogNotificationEditOpenApiV2VO, bool)`

GetAlertNotificationsOk returns a tuple with the AlertNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertNotifications

`func (o *LogNotificationSettingEditOpenApiV2VO) SetAlertNotifications(v []LogNotificationEditOpenApiV2VO)`

SetAlertNotifications sets AlertNotifications field to given value.

### HasAlertNotifications

`func (o *LogNotificationSettingEditOpenApiV2VO) HasAlertNotifications() bool`

HasAlertNotifications returns a boolean if a field has been set.

### GetEventEmailSetting

`func (o *LogNotificationSettingEditOpenApiV2VO) GetEventEmailSetting() LogAlertEmailOpenApiVO`

GetEventEmailSetting returns the EventEmailSetting field if non-nil, zero value otherwise.

### GetEventEmailSettingOk

`func (o *LogNotificationSettingEditOpenApiV2VO) GetEventEmailSettingOk() (*LogAlertEmailOpenApiVO, bool)`

GetEventEmailSettingOk returns a tuple with the EventEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEmailSetting

`func (o *LogNotificationSettingEditOpenApiV2VO) SetEventEmailSetting(v LogAlertEmailOpenApiVO)`

SetEventEmailSetting sets EventEmailSetting field to given value.

### HasEventEmailSetting

`func (o *LogNotificationSettingEditOpenApiV2VO) HasEventEmailSetting() bool`

HasEventEmailSetting returns a boolean if a field has been set.

### GetEventNotifications

`func (o *LogNotificationSettingEditOpenApiV2VO) GetEventNotifications() []LogNotificationEditOpenApiV2VO`

GetEventNotifications returns the EventNotifications field if non-nil, zero value otherwise.

### GetEventNotificationsOk

`func (o *LogNotificationSettingEditOpenApiV2VO) GetEventNotificationsOk() (*[]LogNotificationEditOpenApiV2VO, bool)`

GetEventNotificationsOk returns a tuple with the EventNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifications

`func (o *LogNotificationSettingEditOpenApiV2VO) SetEventNotifications(v []LogNotificationEditOpenApiV2VO)`

SetEventNotifications sets EventNotifications field to given value.

### HasEventNotifications

`func (o *LogNotificationSettingEditOpenApiV2VO) HasEventNotifications() bool`

HasEventNotifications returns a boolean if a field has been set.

### GetWebhookConfig

`func (o *LogNotificationSettingEditOpenApiV2VO) GetWebhookConfig() WebhookConfigEditOpenApiVO`

GetWebhookConfig returns the WebhookConfig field if non-nil, zero value otherwise.

### GetWebhookConfigOk

`func (o *LogNotificationSettingEditOpenApiV2VO) GetWebhookConfigOk() (*WebhookConfigEditOpenApiVO, bool)`

GetWebhookConfigOk returns a tuple with the WebhookConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfig

`func (o *LogNotificationSettingEditOpenApiV2VO) SetWebhookConfig(v WebhookConfigEditOpenApiVO)`

SetWebhookConfig sets WebhookConfig field to given value.

### HasWebhookConfig

`func (o *LogNotificationSettingEditOpenApiV2VO) HasWebhookConfig() bool`

HasWebhookConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


