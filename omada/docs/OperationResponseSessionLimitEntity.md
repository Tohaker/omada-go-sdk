# OperationResponseSessionLimitEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**SessionLimitEntity**](SessionLimitEntity.md) |  | [optional] 

## Methods

### NewOperationResponseSessionLimitEntity

`func NewOperationResponseSessionLimitEntity() *OperationResponseSessionLimitEntity`

NewOperationResponseSessionLimitEntity instantiates a new OperationResponseSessionLimitEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseSessionLimitEntityWithDefaults

`func NewOperationResponseSessionLimitEntityWithDefaults() *OperationResponseSessionLimitEntity`

NewOperationResponseSessionLimitEntityWithDefaults instantiates a new OperationResponseSessionLimitEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseSessionLimitEntity) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseSessionLimitEntity) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseSessionLimitEntity) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseSessionLimitEntity) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseSessionLimitEntity) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseSessionLimitEntity) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseSessionLimitEntity) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseSessionLimitEntity) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseSessionLimitEntity) GetResult() SessionLimitEntity`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseSessionLimitEntity) GetResultOk() (*SessionLimitEntity, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseSessionLimitEntity) SetResult(v SessionLimitEntity)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseSessionLimitEntity) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


