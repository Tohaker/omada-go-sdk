# OperationResponseDhcpServerForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**DhcpServerForVlanVO**](DhcpServerForVlanVO.md) |  | [optional] 

## Methods

### NewOperationResponseDhcpServerForVlanVO

`func NewOperationResponseDhcpServerForVlanVO() *OperationResponseDhcpServerForVlanVO`

NewOperationResponseDhcpServerForVlanVO instantiates a new OperationResponseDhcpServerForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseDhcpServerForVlanVOWithDefaults

`func NewOperationResponseDhcpServerForVlanVOWithDefaults() *OperationResponseDhcpServerForVlanVO`

NewOperationResponseDhcpServerForVlanVOWithDefaults instantiates a new OperationResponseDhcpServerForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseDhcpServerForVlanVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseDhcpServerForVlanVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseDhcpServerForVlanVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseDhcpServerForVlanVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseDhcpServerForVlanVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseDhcpServerForVlanVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseDhcpServerForVlanVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseDhcpServerForVlanVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseDhcpServerForVlanVO) GetResult() DhcpServerForVlanVO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseDhcpServerForVlanVO) GetResultOk() (*DhcpServerForVlanVO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseDhcpServerForVlanVO) SetResult(v DhcpServerForVlanVO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseDhcpServerForVlanVO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


