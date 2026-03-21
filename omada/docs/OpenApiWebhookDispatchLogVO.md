# OpenApiWebhookDispatchLogVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptMessages** | Pointer to [**[]OpenApiWebhookAttemptMessageVO**](OpenApiWebhookAttemptMessageVO.md) | Attempt Message of Dispatch Log Detail | [optional] 
**DispatchDetail** | Pointer to [**OpenApiDispatchLogDetailVO**](OpenApiDispatchLogDetailVO.md) |  | [optional] 
**DispatchMessage** | Pointer to [**OpenApiWebhookLogMessageVO**](OpenApiWebhookLogMessageVO.md) |  | [optional] 

## Methods

### NewOpenApiWebhookDispatchLogVO

`func NewOpenApiWebhookDispatchLogVO() *OpenApiWebhookDispatchLogVO`

NewOpenApiWebhookDispatchLogVO instantiates a new OpenApiWebhookDispatchLogVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiWebhookDispatchLogVOWithDefaults

`func NewOpenApiWebhookDispatchLogVOWithDefaults() *OpenApiWebhookDispatchLogVO`

NewOpenApiWebhookDispatchLogVOWithDefaults instantiates a new OpenApiWebhookDispatchLogVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptMessages

`func (o *OpenApiWebhookDispatchLogVO) GetAttemptMessages() []OpenApiWebhookAttemptMessageVO`

GetAttemptMessages returns the AttemptMessages field if non-nil, zero value otherwise.

### GetAttemptMessagesOk

`func (o *OpenApiWebhookDispatchLogVO) GetAttemptMessagesOk() (*[]OpenApiWebhookAttemptMessageVO, bool)`

GetAttemptMessagesOk returns a tuple with the AttemptMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptMessages

`func (o *OpenApiWebhookDispatchLogVO) SetAttemptMessages(v []OpenApiWebhookAttemptMessageVO)`

SetAttemptMessages sets AttemptMessages field to given value.

### HasAttemptMessages

`func (o *OpenApiWebhookDispatchLogVO) HasAttemptMessages() bool`

HasAttemptMessages returns a boolean if a field has been set.

### GetDispatchDetail

`func (o *OpenApiWebhookDispatchLogVO) GetDispatchDetail() OpenApiDispatchLogDetailVO`

GetDispatchDetail returns the DispatchDetail field if non-nil, zero value otherwise.

### GetDispatchDetailOk

`func (o *OpenApiWebhookDispatchLogVO) GetDispatchDetailOk() (*OpenApiDispatchLogDetailVO, bool)`

GetDispatchDetailOk returns a tuple with the DispatchDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchDetail

`func (o *OpenApiWebhookDispatchLogVO) SetDispatchDetail(v OpenApiDispatchLogDetailVO)`

SetDispatchDetail sets DispatchDetail field to given value.

### HasDispatchDetail

`func (o *OpenApiWebhookDispatchLogVO) HasDispatchDetail() bool`

HasDispatchDetail returns a boolean if a field has been set.

### GetDispatchMessage

`func (o *OpenApiWebhookDispatchLogVO) GetDispatchMessage() OpenApiWebhookLogMessageVO`

GetDispatchMessage returns the DispatchMessage field if non-nil, zero value otherwise.

### GetDispatchMessageOk

`func (o *OpenApiWebhookDispatchLogVO) GetDispatchMessageOk() (*OpenApiWebhookLogMessageVO, bool)`

GetDispatchMessageOk returns a tuple with the DispatchMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchMessage

`func (o *OpenApiWebhookDispatchLogVO) SetDispatchMessage(v OpenApiWebhookLogMessageVO)`

SetDispatchMessage sets DispatchMessage field to given value.

### HasDispatchMessage

`func (o *OpenApiWebhookDispatchLogVO) HasDispatchMessage() bool`

HasDispatchMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


