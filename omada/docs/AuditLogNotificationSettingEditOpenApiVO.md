# AuditLogNotificationSettingEditOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogNotifications** | [**[]AuditLogNotificationEditOpenApiVO**](AuditLogNotificationEditOpenApiVO.md) | Audit Log Notification List | 
**WebhookConfig** | [**WebhookConfigEditOpenApiVO**](WebhookConfigEditOpenApiVO.md) |  | 

## Methods

### NewAuditLogNotificationSettingEditOpenApiVO

`func NewAuditLogNotificationSettingEditOpenApiVO(auditLogNotifications []AuditLogNotificationEditOpenApiVO, webhookConfig WebhookConfigEditOpenApiVO, ) *AuditLogNotificationSettingEditOpenApiVO`

NewAuditLogNotificationSettingEditOpenApiVO instantiates a new AuditLogNotificationSettingEditOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogNotificationSettingEditOpenApiVOWithDefaults

`func NewAuditLogNotificationSettingEditOpenApiVOWithDefaults() *AuditLogNotificationSettingEditOpenApiVO`

NewAuditLogNotificationSettingEditOpenApiVOWithDefaults instantiates a new AuditLogNotificationSettingEditOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogNotifications

`func (o *AuditLogNotificationSettingEditOpenApiVO) GetAuditLogNotifications() []AuditLogNotificationEditOpenApiVO`

GetAuditLogNotifications returns the AuditLogNotifications field if non-nil, zero value otherwise.

### GetAuditLogNotificationsOk

`func (o *AuditLogNotificationSettingEditOpenApiVO) GetAuditLogNotificationsOk() (*[]AuditLogNotificationEditOpenApiVO, bool)`

GetAuditLogNotificationsOk returns a tuple with the AuditLogNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogNotifications

`func (o *AuditLogNotificationSettingEditOpenApiVO) SetAuditLogNotifications(v []AuditLogNotificationEditOpenApiVO)`

SetAuditLogNotifications sets AuditLogNotifications field to given value.


### GetWebhookConfig

`func (o *AuditLogNotificationSettingEditOpenApiVO) GetWebhookConfig() WebhookConfigEditOpenApiVO`

GetWebhookConfig returns the WebhookConfig field if non-nil, zero value otherwise.

### GetWebhookConfigOk

`func (o *AuditLogNotificationSettingEditOpenApiVO) GetWebhookConfigOk() (*WebhookConfigEditOpenApiVO, bool)`

GetWebhookConfigOk returns a tuple with the WebhookConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfig

`func (o *AuditLogNotificationSettingEditOpenApiVO) SetWebhookConfig(v WebhookConfigEditOpenApiVO)`

SetWebhookConfig sets WebhookConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


