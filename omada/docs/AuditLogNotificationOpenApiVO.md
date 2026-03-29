# AuditLogNotificationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Audit Log Notification Category Key | [optional] 
**ShortMsg** | Pointer to **string** | Audit Log Notification Category Short Message | [optional] 
**Webhook** | Pointer to **bool** | Audit Log Notification Category Enable or Disable Webhook | [optional] 

## Methods

### NewAuditLogNotificationOpenApiVO

`func NewAuditLogNotificationOpenApiVO() *AuditLogNotificationOpenApiVO`

NewAuditLogNotificationOpenApiVO instantiates a new AuditLogNotificationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogNotificationOpenApiVOWithDefaults

`func NewAuditLogNotificationOpenApiVOWithDefaults() *AuditLogNotificationOpenApiVO`

NewAuditLogNotificationOpenApiVOWithDefaults instantiates a new AuditLogNotificationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AuditLogNotificationOpenApiVO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuditLogNotificationOpenApiVO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuditLogNotificationOpenApiVO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuditLogNotificationOpenApiVO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetShortMsg

`func (o *AuditLogNotificationOpenApiVO) GetShortMsg() string`

GetShortMsg returns the ShortMsg field if non-nil, zero value otherwise.

### GetShortMsgOk

`func (o *AuditLogNotificationOpenApiVO) GetShortMsgOk() (*string, bool)`

GetShortMsgOk returns a tuple with the ShortMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMsg

`func (o *AuditLogNotificationOpenApiVO) SetShortMsg(v string)`

SetShortMsg sets ShortMsg field to given value.

### HasShortMsg

`func (o *AuditLogNotificationOpenApiVO) HasShortMsg() bool`

HasShortMsg returns a boolean if a field has been set.

### GetWebhook

`func (o *AuditLogNotificationOpenApiVO) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *AuditLogNotificationOpenApiVO) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *AuditLogNotificationOpenApiVO) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *AuditLogNotificationOpenApiVO) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


