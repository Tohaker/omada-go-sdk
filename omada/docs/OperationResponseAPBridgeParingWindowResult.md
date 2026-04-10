# OperationResponseAPBridgeParingWindowResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**APBridgeParingWindowResult**](APBridgeParingWindowResult.md) |  | [optional] 

## Methods

### NewOperationResponseAPBridgeParingWindowResult

`func NewOperationResponseAPBridgeParingWindowResult() *OperationResponseAPBridgeParingWindowResult`

NewOperationResponseAPBridgeParingWindowResult instantiates a new OperationResponseAPBridgeParingWindowResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseAPBridgeParingWindowResultWithDefaults

`func NewOperationResponseAPBridgeParingWindowResultWithDefaults() *OperationResponseAPBridgeParingWindowResult`

NewOperationResponseAPBridgeParingWindowResultWithDefaults instantiates a new OperationResponseAPBridgeParingWindowResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseAPBridgeParingWindowResult) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseAPBridgeParingWindowResult) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseAPBridgeParingWindowResult) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseAPBridgeParingWindowResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseAPBridgeParingWindowResult) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseAPBridgeParingWindowResult) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseAPBridgeParingWindowResult) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseAPBridgeParingWindowResult) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseAPBridgeParingWindowResult) GetResult() APBridgeParingWindowResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseAPBridgeParingWindowResult) GetResultOk() (*APBridgeParingWindowResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseAPBridgeParingWindowResult) SetResult(v APBridgeParingWindowResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseAPBridgeParingWindowResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


