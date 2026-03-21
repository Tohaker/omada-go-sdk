# OperationResponseGetUserListResponseVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**GetUserListResponseVO**](GetUserListResponseVO.md) |  | [optional] 

## Methods

### NewOperationResponseGetUserListResponseVO

`func NewOperationResponseGetUserListResponseVO() *OperationResponseGetUserListResponseVO`

NewOperationResponseGetUserListResponseVO instantiates a new OperationResponseGetUserListResponseVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseGetUserListResponseVOWithDefaults

`func NewOperationResponseGetUserListResponseVOWithDefaults() *OperationResponseGetUserListResponseVO`

NewOperationResponseGetUserListResponseVOWithDefaults instantiates a new OperationResponseGetUserListResponseVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseGetUserListResponseVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseGetUserListResponseVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseGetUserListResponseVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseGetUserListResponseVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseGetUserListResponseVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseGetUserListResponseVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseGetUserListResponseVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseGetUserListResponseVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseGetUserListResponseVO) GetResult() GetUserListResponseVO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseGetUserListResponseVO) GetResultOk() (*GetUserListResponseVO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseGetUserListResponseVO) SetResult(v GetUserListResponseVO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseGetUserListResponseVO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


