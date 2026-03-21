# OpenApiWebhookAttemptMessageVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptTime** | Pointer to **int64** | Webhook Attempt Push Time, Unit (ms) | [optional] 
**MessageCode** | Pointer to **int32** | Http Message Code of Webhook Attempt Push | [optional] 
**Status** | Pointer to **int32** | Webhook Attempt Push Status: 0/1 (Success/Failed) | [optional] 

## Methods

### NewOpenApiWebhookAttemptMessageVO

`func NewOpenApiWebhookAttemptMessageVO() *OpenApiWebhookAttemptMessageVO`

NewOpenApiWebhookAttemptMessageVO instantiates a new OpenApiWebhookAttemptMessageVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiWebhookAttemptMessageVOWithDefaults

`func NewOpenApiWebhookAttemptMessageVOWithDefaults() *OpenApiWebhookAttemptMessageVO`

NewOpenApiWebhookAttemptMessageVOWithDefaults instantiates a new OpenApiWebhookAttemptMessageVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptTime

`func (o *OpenApiWebhookAttemptMessageVO) GetAttemptTime() int64`

GetAttemptTime returns the AttemptTime field if non-nil, zero value otherwise.

### GetAttemptTimeOk

`func (o *OpenApiWebhookAttemptMessageVO) GetAttemptTimeOk() (*int64, bool)`

GetAttemptTimeOk returns a tuple with the AttemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptTime

`func (o *OpenApiWebhookAttemptMessageVO) SetAttemptTime(v int64)`

SetAttemptTime sets AttemptTime field to given value.

### HasAttemptTime

`func (o *OpenApiWebhookAttemptMessageVO) HasAttemptTime() bool`

HasAttemptTime returns a boolean if a field has been set.

### GetMessageCode

`func (o *OpenApiWebhookAttemptMessageVO) GetMessageCode() int32`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *OpenApiWebhookAttemptMessageVO) GetMessageCodeOk() (*int32, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *OpenApiWebhookAttemptMessageVO) SetMessageCode(v int32)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *OpenApiWebhookAttemptMessageVO) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetStatus

`func (o *OpenApiWebhookAttemptMessageVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenApiWebhookAttemptMessageVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenApiWebhookAttemptMessageVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenApiWebhookAttemptMessageVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


