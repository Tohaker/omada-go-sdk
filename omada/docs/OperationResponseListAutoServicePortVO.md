# OperationResponseListAutoServicePortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]AutoServicePortVO**](AutoServicePortVO.md) |  | [optional] 

## Methods

### NewOperationResponseListAutoServicePortVO

`func NewOperationResponseListAutoServicePortVO() *OperationResponseListAutoServicePortVO`

NewOperationResponseListAutoServicePortVO instantiates a new OperationResponseListAutoServicePortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListAutoServicePortVOWithDefaults

`func NewOperationResponseListAutoServicePortVOWithDefaults() *OperationResponseListAutoServicePortVO`

NewOperationResponseListAutoServicePortVOWithDefaults instantiates a new OperationResponseListAutoServicePortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListAutoServicePortVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListAutoServicePortVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListAutoServicePortVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListAutoServicePortVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListAutoServicePortVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListAutoServicePortVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListAutoServicePortVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListAutoServicePortVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListAutoServicePortVO) GetResult() []AutoServicePortVO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListAutoServicePortVO) GetResultOk() (*[]AutoServicePortVO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListAutoServicePortVO) SetResult(v []AutoServicePortVO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListAutoServicePortVO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


