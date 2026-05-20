# OperationResponseApBridgeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**ApBridgeInfo**](ApBridgeInfo.md) |  | [optional] 

## Methods

### NewOperationResponseApBridgeInfo

`func NewOperationResponseApBridgeInfo() *OperationResponseApBridgeInfo`

NewOperationResponseApBridgeInfo instantiates a new OperationResponseApBridgeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseApBridgeInfoWithDefaults

`func NewOperationResponseApBridgeInfoWithDefaults() *OperationResponseApBridgeInfo`

NewOperationResponseApBridgeInfoWithDefaults instantiates a new OperationResponseApBridgeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseApBridgeInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseApBridgeInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseApBridgeInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseApBridgeInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseApBridgeInfo) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseApBridgeInfo) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseApBridgeInfo) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseApBridgeInfo) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseApBridgeInfo) GetResult() ApBridgeInfo`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseApBridgeInfo) GetResultOk() (*ApBridgeInfo, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseApBridgeInfo) SetResult(v ApBridgeInfo)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseApBridgeInfo) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


