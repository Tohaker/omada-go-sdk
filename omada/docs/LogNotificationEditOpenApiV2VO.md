# LogNotificationEditOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **bool** | Log Notification Enable or Disable Email | 
**Enable** | **bool** | Log Notification Enable or Disable | 
**Key** | **string** | For the values of Log Notification Key, refer to section 5.6.1 of the Open API Access | 
**Webhook** | Pointer to **bool** | Log Notification Category Enable or Disable Webhook (This config applies to the Omada Pro Controller only and should not bu null) | [optional] 

## Methods

### NewLogNotificationEditOpenApiV2VO

`func NewLogNotificationEditOpenApiV2VO(email bool, enable bool, key string, ) *LogNotificationEditOpenApiV2VO`

NewLogNotificationEditOpenApiV2VO instantiates a new LogNotificationEditOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogNotificationEditOpenApiV2VOWithDefaults

`func NewLogNotificationEditOpenApiV2VOWithDefaults() *LogNotificationEditOpenApiV2VO`

NewLogNotificationEditOpenApiV2VOWithDefaults instantiates a new LogNotificationEditOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LogNotificationEditOpenApiV2VO) GetEmail() bool`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LogNotificationEditOpenApiV2VO) GetEmailOk() (*bool, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LogNotificationEditOpenApiV2VO) SetEmail(v bool)`

SetEmail sets Email field to given value.


### GetEnable

`func (o *LogNotificationEditOpenApiV2VO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LogNotificationEditOpenApiV2VO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LogNotificationEditOpenApiV2VO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetKey

`func (o *LogNotificationEditOpenApiV2VO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LogNotificationEditOpenApiV2VO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LogNotificationEditOpenApiV2VO) SetKey(v string)`

SetKey sets Key field to given value.


### GetWebhook

`func (o *LogNotificationEditOpenApiV2VO) GetWebhook() bool`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *LogNotificationEditOpenApiV2VO) GetWebhookOk() (*bool, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *LogNotificationEditOpenApiV2VO) SetWebhook(v bool)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *LogNotificationEditOpenApiV2VO) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


