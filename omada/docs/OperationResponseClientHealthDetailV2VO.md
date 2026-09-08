# OperationResponseClientHealthDetailV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**ClientHealthDetailV2VO**](ClientHealthDetailV2VO.md) |  | [optional] 

## Methods

### NewOperationResponseClientHealthDetailV2VO

`func NewOperationResponseClientHealthDetailV2VO() *OperationResponseClientHealthDetailV2VO`

NewOperationResponseClientHealthDetailV2VO instantiates a new OperationResponseClientHealthDetailV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseClientHealthDetailV2VOWithDefaults

`func NewOperationResponseClientHealthDetailV2VOWithDefaults() *OperationResponseClientHealthDetailV2VO`

NewOperationResponseClientHealthDetailV2VOWithDefaults instantiates a new OperationResponseClientHealthDetailV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseClientHealthDetailV2VO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseClientHealthDetailV2VO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseClientHealthDetailV2VO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseClientHealthDetailV2VO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseClientHealthDetailV2VO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseClientHealthDetailV2VO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseClientHealthDetailV2VO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseClientHealthDetailV2VO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseClientHealthDetailV2VO) GetResult() ClientHealthDetailV2VO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseClientHealthDetailV2VO) GetResultOk() (*ClientHealthDetailV2VO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseClientHealthDetailV2VO) SetResult(v ClientHealthDetailV2VO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseClientHealthDetailV2VO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


