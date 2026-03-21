# OperationResponseListClientActivitiesVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]ClientActivitiesVO**](ClientActivitiesVO.md) |  | [optional] 

## Methods

### NewOperationResponseListClientActivitiesVO

`func NewOperationResponseListClientActivitiesVO() *OperationResponseListClientActivitiesVO`

NewOperationResponseListClientActivitiesVO instantiates a new OperationResponseListClientActivitiesVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListClientActivitiesVOWithDefaults

`func NewOperationResponseListClientActivitiesVOWithDefaults() *OperationResponseListClientActivitiesVO`

NewOperationResponseListClientActivitiesVOWithDefaults instantiates a new OperationResponseListClientActivitiesVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListClientActivitiesVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListClientActivitiesVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListClientActivitiesVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListClientActivitiesVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListClientActivitiesVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListClientActivitiesVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListClientActivitiesVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListClientActivitiesVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListClientActivitiesVO) GetResult() []ClientActivitiesVO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListClientActivitiesVO) GetResultOk() (*[]ClientActivitiesVO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListClientActivitiesVO) SetResult(v []ClientActivitiesVO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListClientActivitiesVO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


