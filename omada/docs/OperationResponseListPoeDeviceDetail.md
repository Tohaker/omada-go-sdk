# OperationResponseListPoeDeviceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]PoeDeviceDetail**](PoeDeviceDetail.md) |  | [optional] 

## Methods

### NewOperationResponseListPoeDeviceDetail

`func NewOperationResponseListPoeDeviceDetail() *OperationResponseListPoeDeviceDetail`

NewOperationResponseListPoeDeviceDetail instantiates a new OperationResponseListPoeDeviceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListPoeDeviceDetailWithDefaults

`func NewOperationResponseListPoeDeviceDetailWithDefaults() *OperationResponseListPoeDeviceDetail`

NewOperationResponseListPoeDeviceDetailWithDefaults instantiates a new OperationResponseListPoeDeviceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListPoeDeviceDetail) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListPoeDeviceDetail) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListPoeDeviceDetail) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListPoeDeviceDetail) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListPoeDeviceDetail) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListPoeDeviceDetail) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListPoeDeviceDetail) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListPoeDeviceDetail) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListPoeDeviceDetail) GetResult() []PoeDeviceDetail`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListPoeDeviceDetail) GetResultOk() (*[]PoeDeviceDetail, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListPoeDeviceDetail) SetResult(v []PoeDeviceDetail)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListPoeDeviceDetail) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


