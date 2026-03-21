# OpenApiWebhookSettingTestVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Webhook name. It should contain 1 to 128 characters. | 
**ShardedSecret** | Pointer to **string** | Webhook Sharded Secret. It should contain 0 to 128 characters. | [optional] 
**Template** | **int32** | Webhook template, it should be a value as follow: 0:Omada template, 1:Google chat template. Example: 0.  | 
**UrlList** | **[]string** | Webhook URL List. Up to 3 entries are allowed for the URL list | 

## Methods

### NewOpenApiWebhookSettingTestVO

`func NewOpenApiWebhookSettingTestVO(name string, template int32, urlList []string, ) *OpenApiWebhookSettingTestVO`

NewOpenApiWebhookSettingTestVO instantiates a new OpenApiWebhookSettingTestVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiWebhookSettingTestVOWithDefaults

`func NewOpenApiWebhookSettingTestVOWithDefaults() *OpenApiWebhookSettingTestVO`

NewOpenApiWebhookSettingTestVOWithDefaults instantiates a new OpenApiWebhookSettingTestVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OpenApiWebhookSettingTestVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenApiWebhookSettingTestVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenApiWebhookSettingTestVO) SetName(v string)`

SetName sets Name field to given value.


### GetShardedSecret

`func (o *OpenApiWebhookSettingTestVO) GetShardedSecret() string`

GetShardedSecret returns the ShardedSecret field if non-nil, zero value otherwise.

### GetShardedSecretOk

`func (o *OpenApiWebhookSettingTestVO) GetShardedSecretOk() (*string, bool)`

GetShardedSecretOk returns a tuple with the ShardedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardedSecret

`func (o *OpenApiWebhookSettingTestVO) SetShardedSecret(v string)`

SetShardedSecret sets ShardedSecret field to given value.

### HasShardedSecret

`func (o *OpenApiWebhookSettingTestVO) HasShardedSecret() bool`

HasShardedSecret returns a boolean if a field has been set.

### GetTemplate

`func (o *OpenApiWebhookSettingTestVO) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *OpenApiWebhookSettingTestVO) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *OpenApiWebhookSettingTestVO) SetTemplate(v int32)`

SetTemplate sets Template field to given value.


### GetUrlList

`func (o *OpenApiWebhookSettingTestVO) GetUrlList() []string`

GetUrlList returns the UrlList field if non-nil, zero value otherwise.

### GetUrlListOk

`func (o *OpenApiWebhookSettingTestVO) GetUrlListOk() (*[]string, bool)`

GetUrlListOk returns a tuple with the UrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlList

`func (o *OpenApiWebhookSettingTestVO) SetUrlList(v []string)`

SetUrlList sets UrlList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


