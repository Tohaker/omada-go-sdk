# AuditLogNotificationEditOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | For the values of Audit Log Notification Category Key, refer to section 5.2.1 of the Open API Access | 
**Webhook** | **bool** | Audit Log Notification Category Enable or Disable Webhook | 

## Methods

### NewAuditLogNotificationEditOpenApiVO

`func NewAuditLogNotificationEditOpenApiVO(key string, webhook bool, ) *AuditLogNotificationEditOpenApiVO`

NewAuditLogNotificationEditOpenApiVO instantiates a new AuditLogNotificationEditOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogNotificationEditOpenApiVOWithDefaults

`func NewAuditLogNotificationEditOpenApiVOWithDefaults() *AuditLogNotificationEditOpenApiVO`

NewAuditLogNotificationEditOpenApiVOWithDefaults instantiates a new AuditLogNotificationEditOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AuditLogNotificationEditOpenApiVO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuditLogNotificationEditOpenApiVO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuditLogNotificationEditOpenApiVO) SetKey(v string)`

SetKey sets Key field to given value.


### GetWebhook

`func (o *AuditLogNotificationEditOpenApiVO) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *AuditLogNotificationEditOpenApiVO) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *AuditLogNotificationEditOpenApiVO) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


