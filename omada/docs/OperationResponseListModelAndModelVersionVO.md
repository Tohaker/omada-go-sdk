# OperationResponseListModelAndModelVersionVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]ModelAndModelVersionVO**](ModelAndModelVersionVO.md) |  | [optional] 

## Methods

### NewOperationResponseListModelAndModelVersionVO

`func NewOperationResponseListModelAndModelVersionVO() *OperationResponseListModelAndModelVersionVO`

NewOperationResponseListModelAndModelVersionVO instantiates a new OperationResponseListModelAndModelVersionVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListModelAndModelVersionVOWithDefaults

`func NewOperationResponseListModelAndModelVersionVOWithDefaults() *OperationResponseListModelAndModelVersionVO`

NewOperationResponseListModelAndModelVersionVOWithDefaults instantiates a new OperationResponseListModelAndModelVersionVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListModelAndModelVersionVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListModelAndModelVersionVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListModelAndModelVersionVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListModelAndModelVersionVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListModelAndModelVersionVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListModelAndModelVersionVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListModelAndModelVersionVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListModelAndModelVersionVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListModelAndModelVersionVO) GetResult() []ModelAndModelVersionVO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListModelAndModelVersionVO) GetResultOk() (*[]ModelAndModelVersionVO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListModelAndModelVersionVO) SetResult(v []ModelAndModelVersionVO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListModelAndModelVersionVO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


