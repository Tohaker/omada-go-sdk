# OperationResponseAddCallForwardingRuleResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**AddCallForwardingRuleResp**](AddCallForwardingRuleResp.md) |  | [optional] 

## Methods

### NewOperationResponseAddCallForwardingRuleResp

`func NewOperationResponseAddCallForwardingRuleResp() *OperationResponseAddCallForwardingRuleResp`

NewOperationResponseAddCallForwardingRuleResp instantiates a new OperationResponseAddCallForwardingRuleResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseAddCallForwardingRuleRespWithDefaults

`func NewOperationResponseAddCallForwardingRuleRespWithDefaults() *OperationResponseAddCallForwardingRuleResp`

NewOperationResponseAddCallForwardingRuleRespWithDefaults instantiates a new OperationResponseAddCallForwardingRuleResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseAddCallForwardingRuleResp) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseAddCallForwardingRuleResp) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseAddCallForwardingRuleResp) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseAddCallForwardingRuleResp) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseAddCallForwardingRuleResp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseAddCallForwardingRuleResp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseAddCallForwardingRuleResp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseAddCallForwardingRuleResp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseAddCallForwardingRuleResp) GetResult() AddCallForwardingRuleResp`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseAddCallForwardingRuleResp) GetResultOk() (*AddCallForwardingRuleResp, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseAddCallForwardingRuleResp) SetResult(v AddCallForwardingRuleResp)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseAddCallForwardingRuleResp) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


