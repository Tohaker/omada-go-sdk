# OperationResponseApIPv6Setting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**ApIPv6Setting**](ApIPv6Setting.md) |  | [optional] 

## Methods

### NewOperationResponseApIPv6Setting

`func NewOperationResponseApIPv6Setting() *OperationResponseApIPv6Setting`

NewOperationResponseApIPv6Setting instantiates a new OperationResponseApIPv6Setting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseApIPv6SettingWithDefaults

`func NewOperationResponseApIPv6SettingWithDefaults() *OperationResponseApIPv6Setting`

NewOperationResponseApIPv6SettingWithDefaults instantiates a new OperationResponseApIPv6Setting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseApIPv6Setting) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseApIPv6Setting) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseApIPv6Setting) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseApIPv6Setting) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseApIPv6Setting) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseApIPv6Setting) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseApIPv6Setting) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseApIPv6Setting) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseApIPv6Setting) GetResult() ApIPv6Setting`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseApIPv6Setting) GetResultOk() (*ApIPv6Setting, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseApIPv6Setting) SetResult(v ApIPv6Setting)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseApIPv6Setting) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


