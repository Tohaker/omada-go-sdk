# WebhookConfigEditOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookEnable** | **bool** | Audit or Omada Log Notification Enable for Disable Webhook | 
**WebhookId** | Pointer to **string** | Webhook ID (Webhook ID should be configured in webhook setting, when Webhook Enable is true) | [optional] 

## Methods

### NewWebhookConfigEditOpenApiVO

`func NewWebhookConfigEditOpenApiVO(webhookEnable bool, ) *WebhookConfigEditOpenApiVO`

NewWebhookConfigEditOpenApiVO instantiates a new WebhookConfigEditOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookConfigEditOpenApiVOWithDefaults

`func NewWebhookConfigEditOpenApiVOWithDefaults() *WebhookConfigEditOpenApiVO`

NewWebhookConfigEditOpenApiVOWithDefaults instantiates a new WebhookConfigEditOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookEnable

`func (o *WebhookConfigEditOpenApiVO) GetWebhookEnable() bool`

GetWebhookEnable returns the WebhookEnable field if non-nil, zero value otherwise.

### GetWebhookEnableOk

`func (o *WebhookConfigEditOpenApiVO) GetWebhookEnableOk() (*bool, bool)`

GetWebhookEnableOk returns a tuple with the WebhookEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEnable

`func (o *WebhookConfigEditOpenApiVO) SetWebhookEnable(v bool)`

SetWebhookEnable sets WebhookEnable field to given value.


### GetWebhookId

`func (o *WebhookConfigEditOpenApiVO) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookConfigEditOpenApiVO) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookConfigEditOpenApiVO) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebhookConfigEditOpenApiVO) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


