# OperationResponseClientConnectionHistories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**ClientConnectionHistories**](ClientConnectionHistories.md) |  | [optional] 

## Methods

### NewOperationResponseClientConnectionHistories

`func NewOperationResponseClientConnectionHistories() *OperationResponseClientConnectionHistories`

NewOperationResponseClientConnectionHistories instantiates a new OperationResponseClientConnectionHistories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseClientConnectionHistoriesWithDefaults

`func NewOperationResponseClientConnectionHistoriesWithDefaults() *OperationResponseClientConnectionHistories`

NewOperationResponseClientConnectionHistoriesWithDefaults instantiates a new OperationResponseClientConnectionHistories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseClientConnectionHistories) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseClientConnectionHistories) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseClientConnectionHistories) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseClientConnectionHistories) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseClientConnectionHistories) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseClientConnectionHistories) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseClientConnectionHistories) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseClientConnectionHistories) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseClientConnectionHistories) GetResult() ClientConnectionHistories`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseClientConnectionHistories) GetResultOk() (*ClientConnectionHistories, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseClientConnectionHistories) SetResult(v ClientConnectionHistories)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseClientConnectionHistories) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


