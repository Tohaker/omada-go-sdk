# OperationResponseVpnPreSharedKeyVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**VpnPreSharedKeyVO**](VpnPreSharedKeyVO.md) |  | [optional] 

## Methods

### NewOperationResponseVpnPreSharedKeyVO

`func NewOperationResponseVpnPreSharedKeyVO() *OperationResponseVpnPreSharedKeyVO`

NewOperationResponseVpnPreSharedKeyVO instantiates a new OperationResponseVpnPreSharedKeyVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseVpnPreSharedKeyVOWithDefaults

`func NewOperationResponseVpnPreSharedKeyVOWithDefaults() *OperationResponseVpnPreSharedKeyVO`

NewOperationResponseVpnPreSharedKeyVOWithDefaults instantiates a new OperationResponseVpnPreSharedKeyVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseVpnPreSharedKeyVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseVpnPreSharedKeyVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseVpnPreSharedKeyVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseVpnPreSharedKeyVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseVpnPreSharedKeyVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseVpnPreSharedKeyVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseVpnPreSharedKeyVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseVpnPreSharedKeyVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseVpnPreSharedKeyVO) GetResult() VpnPreSharedKeyVO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseVpnPreSharedKeyVO) GetResultOk() (*VpnPreSharedKeyVO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseVpnPreSharedKeyVO) SetResult(v VpnPreSharedKeyVO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseVpnPreSharedKeyVO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


