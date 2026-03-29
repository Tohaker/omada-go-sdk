# OperationResponseListApChannelLoadResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]ApChannelLoadResult**](ApChannelLoadResult.md) |  | [optional] 

## Methods

### NewOperationResponseListApChannelLoadResult

`func NewOperationResponseListApChannelLoadResult() *OperationResponseListApChannelLoadResult`

NewOperationResponseListApChannelLoadResult instantiates a new OperationResponseListApChannelLoadResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListApChannelLoadResultWithDefaults

`func NewOperationResponseListApChannelLoadResultWithDefaults() *OperationResponseListApChannelLoadResult`

NewOperationResponseListApChannelLoadResultWithDefaults instantiates a new OperationResponseListApChannelLoadResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListApChannelLoadResult) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListApChannelLoadResult) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListApChannelLoadResult) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListApChannelLoadResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListApChannelLoadResult) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListApChannelLoadResult) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListApChannelLoadResult) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListApChannelLoadResult) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListApChannelLoadResult) GetResult() []ApChannelLoadResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListApChannelLoadResult) GetResultOk() (*[]ApChannelLoadResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListApChannelLoadResult) SetResult(v []ApChannelLoadResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListApChannelLoadResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


