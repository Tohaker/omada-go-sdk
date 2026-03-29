# OpenApiDispatchLogDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptNumber** | Pointer to **int32** | Webhook Attempt Number | [optional] 
**Content** | Pointer to **string** | Webhook Push Content | [optional] 
**Headers** | Pointer to **string** | Http Headers of Webhook Push | [optional] 

## Methods

### NewOpenApiDispatchLogDetailVO

`func NewOpenApiDispatchLogDetailVO() *OpenApiDispatchLogDetailVO`

NewOpenApiDispatchLogDetailVO instantiates a new OpenApiDispatchLogDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiDispatchLogDetailVOWithDefaults

`func NewOpenApiDispatchLogDetailVOWithDefaults() *OpenApiDispatchLogDetailVO`

NewOpenApiDispatchLogDetailVOWithDefaults instantiates a new OpenApiDispatchLogDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptNumber

`func (o *OpenApiDispatchLogDetailVO) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *OpenApiDispatchLogDetailVO) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *OpenApiDispatchLogDetailVO) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *OpenApiDispatchLogDetailVO) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetContent

`func (o *OpenApiDispatchLogDetailVO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OpenApiDispatchLogDetailVO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OpenApiDispatchLogDetailVO) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OpenApiDispatchLogDetailVO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetHeaders

`func (o *OpenApiDispatchLogDetailVO) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OpenApiDispatchLogDetailVO) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OpenApiDispatchLogDetailVO) SetHeaders(v string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *OpenApiDispatchLogDetailVO) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


