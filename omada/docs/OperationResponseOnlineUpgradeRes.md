# OperationResponseOnlineUpgradeRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**OnlineUpgradeRes**](OnlineUpgradeRes.md) |  | [optional] 

## Methods

### NewOperationResponseOnlineUpgradeRes

`func NewOperationResponseOnlineUpgradeRes() *OperationResponseOnlineUpgradeRes`

NewOperationResponseOnlineUpgradeRes instantiates a new OperationResponseOnlineUpgradeRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseOnlineUpgradeResWithDefaults

`func NewOperationResponseOnlineUpgradeResWithDefaults() *OperationResponseOnlineUpgradeRes`

NewOperationResponseOnlineUpgradeResWithDefaults instantiates a new OperationResponseOnlineUpgradeRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseOnlineUpgradeRes) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseOnlineUpgradeRes) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseOnlineUpgradeRes) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseOnlineUpgradeRes) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseOnlineUpgradeRes) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseOnlineUpgradeRes) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseOnlineUpgradeRes) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseOnlineUpgradeRes) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseOnlineUpgradeRes) GetResult() OnlineUpgradeRes`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseOnlineUpgradeRes) GetResultOk() (*OnlineUpgradeRes, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseOnlineUpgradeRes) SetResult(v OnlineUpgradeRes)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseOnlineUpgradeRes) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


