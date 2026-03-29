# OpenApiWebhookMessageVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageCode** | Pointer to **int32** | Http Message Code of Webhook Push | [optional] 
**Status** | Pointer to **int32** | Webhook Push Status. It should be a value as follows: 0: Success; 1: Failed | [optional] 

## Methods

### NewOpenApiWebhookMessageVO

`func NewOpenApiWebhookMessageVO() *OpenApiWebhookMessageVO`

NewOpenApiWebhookMessageVO instantiates a new OpenApiWebhookMessageVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiWebhookMessageVOWithDefaults

`func NewOpenApiWebhookMessageVOWithDefaults() *OpenApiWebhookMessageVO`

NewOpenApiWebhookMessageVOWithDefaults instantiates a new OpenApiWebhookMessageVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageCode

`func (o *OpenApiWebhookMessageVO) GetMessageCode() int32`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *OpenApiWebhookMessageVO) GetMessageCodeOk() (*int32, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *OpenApiWebhookMessageVO) SetMessageCode(v int32)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *OpenApiWebhookMessageVO) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetStatus

`func (o *OpenApiWebhookMessageVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenApiWebhookMessageVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenApiWebhookMessageVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenApiWebhookMessageVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


