# WebhookSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTime** | Pointer to **int64** | Webhook setting last update time (ms) | [optional] 
**Name** | Pointer to **string** | Webhook Name | [optional] 
**RetryPolicy** | Pointer to **int32** | Webhook retry policy. It should be a value as follows: 0:None, 1:Important (Up to 5 retries over 60 minutes), 2:Critical (Up to 5 retries over 24 hours) | [optional] 
**ShardedSecret** | Pointer to **string** | Webhook Sharded Secret (old token) | [optional] 
**Template** | Pointer to **int32** | Webhook template, it should be a value as follow: 0:Omada template, 1:Google chat template. Example: 0.  | [optional] 
**UrlList** | Pointer to **[]string** | Webhook URL List | [optional] 
**UrlNum** | Pointer to **int32** | Webhook URL Number | [optional] 
**WebhookId** | Pointer to **string** | Webhook ID | [optional] 

## Methods

### NewWebhookSetting

`func NewWebhookSetting() *WebhookSetting`

NewWebhookSetting instantiates a new WebhookSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSettingWithDefaults

`func NewWebhookSettingWithDefaults() *WebhookSetting`

NewWebhookSettingWithDefaults instantiates a new WebhookSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTime

`func (o *WebhookSetting) GetLastTime() int64`

GetLastTime returns the LastTime field if non-nil, zero value otherwise.

### GetLastTimeOk

`func (o *WebhookSetting) GetLastTimeOk() (*int64, bool)`

GetLastTimeOk returns a tuple with the LastTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTime

`func (o *WebhookSetting) SetLastTime(v int64)`

SetLastTime sets LastTime field to given value.

### HasLastTime

`func (o *WebhookSetting) HasLastTime() bool`

HasLastTime returns a boolean if a field has been set.

### GetName

`func (o *WebhookSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetryPolicy

`func (o *WebhookSetting) GetRetryPolicy() int32`

GetRetryPolicy returns the RetryPolicy field if non-nil, zero value otherwise.

### GetRetryPolicyOk

`func (o *WebhookSetting) GetRetryPolicyOk() (*int32, bool)`

GetRetryPolicyOk returns a tuple with the RetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryPolicy

`func (o *WebhookSetting) SetRetryPolicy(v int32)`

SetRetryPolicy sets RetryPolicy field to given value.

### HasRetryPolicy

`func (o *WebhookSetting) HasRetryPolicy() bool`

HasRetryPolicy returns a boolean if a field has been set.

### GetShardedSecret

`func (o *WebhookSetting) GetShardedSecret() string`

GetShardedSecret returns the ShardedSecret field if non-nil, zero value otherwise.

### GetShardedSecretOk

`func (o *WebhookSetting) GetShardedSecretOk() (*string, bool)`

GetShardedSecretOk returns a tuple with the ShardedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardedSecret

`func (o *WebhookSetting) SetShardedSecret(v string)`

SetShardedSecret sets ShardedSecret field to given value.

### HasShardedSecret

`func (o *WebhookSetting) HasShardedSecret() bool`

HasShardedSecret returns a boolean if a field has been set.

### GetTemplate

`func (o *WebhookSetting) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebhookSetting) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebhookSetting) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebhookSetting) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUrlList

`func (o *WebhookSetting) GetUrlList() []string`

GetUrlList returns the UrlList field if non-nil, zero value otherwise.

### GetUrlListOk

`func (o *WebhookSetting) GetUrlListOk() (*[]string, bool)`

GetUrlListOk returns a tuple with the UrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlList

`func (o *WebhookSetting) SetUrlList(v []string)`

SetUrlList sets UrlList field to given value.

### HasUrlList

`func (o *WebhookSetting) HasUrlList() bool`

HasUrlList returns a boolean if a field has been set.

### GetUrlNum

`func (o *WebhookSetting) GetUrlNum() int32`

GetUrlNum returns the UrlNum field if non-nil, zero value otherwise.

### GetUrlNumOk

`func (o *WebhookSetting) GetUrlNumOk() (*int32, bool)`

GetUrlNumOk returns a tuple with the UrlNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlNum

`func (o *WebhookSetting) SetUrlNum(v int32)`

SetUrlNum sets UrlNum field to given value.

### HasUrlNum

`func (o *WebhookSetting) HasUrlNum() bool`

HasUrlNum returns a boolean if a field has been set.

### GetWebhookId

`func (o *WebhookSetting) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookSetting) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookSetting) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebhookSetting) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


