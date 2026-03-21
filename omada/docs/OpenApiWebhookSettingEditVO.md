# OpenApiWebhookSettingEditVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryPolicy** | **int32** | Webhook retry policy. It should be a value as follows: 0:None, 1:Important (Up to 5 retries over 60 minutes), 2:Critical (Up to 5 retries over 24 hours) | 
**ShardedSecret** | Pointer to **string** | Webhook Sharded Secret. It should contain 0 to 128 characters. | [optional] 
**Template** | **int32** | Webhook template, it should be a value as follow: 0:Omada template, 1:Google chat template. Example: 0.  | 
**UrlList** | **[]string** | Webhook URL List. Up to 3 entries are allowed for the URL list  | 

## Methods

### NewOpenApiWebhookSettingEditVO

`func NewOpenApiWebhookSettingEditVO(retryPolicy int32, template int32, urlList []string, ) *OpenApiWebhookSettingEditVO`

NewOpenApiWebhookSettingEditVO instantiates a new OpenApiWebhookSettingEditVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiWebhookSettingEditVOWithDefaults

`func NewOpenApiWebhookSettingEditVOWithDefaults() *OpenApiWebhookSettingEditVO`

NewOpenApiWebhookSettingEditVOWithDefaults instantiates a new OpenApiWebhookSettingEditVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryPolicy

`func (o *OpenApiWebhookSettingEditVO) GetRetryPolicy() int32`

GetRetryPolicy returns the RetryPolicy field if non-nil, zero value otherwise.

### GetRetryPolicyOk

`func (o *OpenApiWebhookSettingEditVO) GetRetryPolicyOk() (*int32, bool)`

GetRetryPolicyOk returns a tuple with the RetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryPolicy

`func (o *OpenApiWebhookSettingEditVO) SetRetryPolicy(v int32)`

SetRetryPolicy sets RetryPolicy field to given value.


### GetShardedSecret

`func (o *OpenApiWebhookSettingEditVO) GetShardedSecret() string`

GetShardedSecret returns the ShardedSecret field if non-nil, zero value otherwise.

### GetShardedSecretOk

`func (o *OpenApiWebhookSettingEditVO) GetShardedSecretOk() (*string, bool)`

GetShardedSecretOk returns a tuple with the ShardedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardedSecret

`func (o *OpenApiWebhookSettingEditVO) SetShardedSecret(v string)`

SetShardedSecret sets ShardedSecret field to given value.

### HasShardedSecret

`func (o *OpenApiWebhookSettingEditVO) HasShardedSecret() bool`

HasShardedSecret returns a boolean if a field has been set.

### GetTemplate

`func (o *OpenApiWebhookSettingEditVO) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *OpenApiWebhookSettingEditVO) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *OpenApiWebhookSettingEditVO) SetTemplate(v int32)`

SetTemplate sets Template field to given value.


### GetUrlList

`func (o *OpenApiWebhookSettingEditVO) GetUrlList() []string`

GetUrlList returns the UrlList field if non-nil, zero value otherwise.

### GetUrlListOk

`func (o *OpenApiWebhookSettingEditVO) GetUrlListOk() (*[]string, bool)`

GetUrlListOk returns a tuple with the UrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlList

`func (o *OpenApiWebhookSettingEditVO) SetUrlList(v []string)`

SetUrlList sets UrlList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


