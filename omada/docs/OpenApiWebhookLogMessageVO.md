# OpenApiWebhookLogMessageVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DispatchTime** | Pointer to **int64** | Webhook Dispatch Log Time, Unit (ms) | [optional] 
**Url** | Pointer to **string** | Webhook URL | [optional] 
**WebhookMessage** | Pointer to [**OpenApiWebhookMessageVO**](OpenApiWebhookMessageVO.md) |  | [optional] 

## Methods

### NewOpenApiWebhookLogMessageVO

`func NewOpenApiWebhookLogMessageVO() *OpenApiWebhookLogMessageVO`

NewOpenApiWebhookLogMessageVO instantiates a new OpenApiWebhookLogMessageVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiWebhookLogMessageVOWithDefaults

`func NewOpenApiWebhookLogMessageVOWithDefaults() *OpenApiWebhookLogMessageVO`

NewOpenApiWebhookLogMessageVOWithDefaults instantiates a new OpenApiWebhookLogMessageVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDispatchTime

`func (o *OpenApiWebhookLogMessageVO) GetDispatchTime() int64`

GetDispatchTime returns the DispatchTime field if non-nil, zero value otherwise.

### GetDispatchTimeOk

`func (o *OpenApiWebhookLogMessageVO) GetDispatchTimeOk() (*int64, bool)`

GetDispatchTimeOk returns a tuple with the DispatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchTime

`func (o *OpenApiWebhookLogMessageVO) SetDispatchTime(v int64)`

SetDispatchTime sets DispatchTime field to given value.

### HasDispatchTime

`func (o *OpenApiWebhookLogMessageVO) HasDispatchTime() bool`

HasDispatchTime returns a boolean if a field has been set.

### GetUrl

`func (o *OpenApiWebhookLogMessageVO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OpenApiWebhookLogMessageVO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OpenApiWebhookLogMessageVO) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OpenApiWebhookLogMessageVO) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWebhookMessage

`func (o *OpenApiWebhookLogMessageVO) GetWebhookMessage() OpenApiWebhookMessageVO`

GetWebhookMessage returns the WebhookMessage field if non-nil, zero value otherwise.

### GetWebhookMessageOk

`func (o *OpenApiWebhookLogMessageVO) GetWebhookMessageOk() (*OpenApiWebhookMessageVO, bool)`

GetWebhookMessageOk returns a tuple with the WebhookMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMessage

`func (o *OpenApiWebhookLogMessageVO) SetWebhookMessage(v OpenApiWebhookMessageVO)`

SetWebhookMessage sets WebhookMessage field to given value.

### HasWebhookMessage

`func (o *OpenApiWebhookLogMessageVO) HasWebhookMessage() bool`

HasWebhookMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


