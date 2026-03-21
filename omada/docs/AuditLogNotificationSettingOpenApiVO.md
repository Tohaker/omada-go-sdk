# AuditLogNotificationSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogNotifications** | Pointer to [**[]AuditLogNotificationOpenApiVO**](AuditLogNotificationOpenApiVO.md) | Audit Log Notification List | [optional] 
**WebhookConfig** | Pointer to [**WebhookConfigOpenApiVO**](WebhookConfigOpenApiVO.md) |  | [optional] 

## Methods

### NewAuditLogNotificationSettingOpenApiVO

`func NewAuditLogNotificationSettingOpenApiVO() *AuditLogNotificationSettingOpenApiVO`

NewAuditLogNotificationSettingOpenApiVO instantiates a new AuditLogNotificationSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogNotificationSettingOpenApiVOWithDefaults

`func NewAuditLogNotificationSettingOpenApiVOWithDefaults() *AuditLogNotificationSettingOpenApiVO`

NewAuditLogNotificationSettingOpenApiVOWithDefaults instantiates a new AuditLogNotificationSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogNotifications

`func (o *AuditLogNotificationSettingOpenApiVO) GetAuditLogNotifications() []AuditLogNotificationOpenApiVO`

GetAuditLogNotifications returns the AuditLogNotifications field if non-nil, zero value otherwise.

### GetAuditLogNotificationsOk

`func (o *AuditLogNotificationSettingOpenApiVO) GetAuditLogNotificationsOk() (*[]AuditLogNotificationOpenApiVO, bool)`

GetAuditLogNotificationsOk returns a tuple with the AuditLogNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogNotifications

`func (o *AuditLogNotificationSettingOpenApiVO) SetAuditLogNotifications(v []AuditLogNotificationOpenApiVO)`

SetAuditLogNotifications sets AuditLogNotifications field to given value.

### HasAuditLogNotifications

`func (o *AuditLogNotificationSettingOpenApiVO) HasAuditLogNotifications() bool`

HasAuditLogNotifications returns a boolean if a field has been set.

### GetWebhookConfig

`func (o *AuditLogNotificationSettingOpenApiVO) GetWebhookConfig() WebhookConfigOpenApiVO`

GetWebhookConfig returns the WebhookConfig field if non-nil, zero value otherwise.

### GetWebhookConfigOk

`func (o *AuditLogNotificationSettingOpenApiVO) GetWebhookConfigOk() (*WebhookConfigOpenApiVO, bool)`

GetWebhookConfigOk returns a tuple with the WebhookConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfig

`func (o *AuditLogNotificationSettingOpenApiVO) SetWebhookConfig(v WebhookConfigOpenApiVO)`

SetWebhookConfig sets WebhookConfig field to given value.

### HasWebhookConfig

`func (o *AuditLogNotificationSettingOpenApiVO) HasWebhookConfig() bool`

HasWebhookConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


