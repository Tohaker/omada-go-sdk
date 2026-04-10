# OperationResponseSpeedTestV2ResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**SpeedTestV2ResultVO**](SpeedTestV2ResultVO.md) |  | [optional] 

## Methods

### NewOperationResponseSpeedTestV2ResultVO

`func NewOperationResponseSpeedTestV2ResultVO() *OperationResponseSpeedTestV2ResultVO`

NewOperationResponseSpeedTestV2ResultVO instantiates a new OperationResponseSpeedTestV2ResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseSpeedTestV2ResultVOWithDefaults

`func NewOperationResponseSpeedTestV2ResultVOWithDefaults() *OperationResponseSpeedTestV2ResultVO`

NewOperationResponseSpeedTestV2ResultVOWithDefaults instantiates a new OperationResponseSpeedTestV2ResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseSpeedTestV2ResultVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseSpeedTestV2ResultVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseSpeedTestV2ResultVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseSpeedTestV2ResultVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseSpeedTestV2ResultVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseSpeedTestV2ResultVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseSpeedTestV2ResultVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseSpeedTestV2ResultVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseSpeedTestV2ResultVO) GetResult() SpeedTestV2ResultVO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseSpeedTestV2ResultVO) GetResultOk() (*SpeedTestV2ResultVO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseSpeedTestV2ResultVO) SetResult(v SpeedTestV2ResultVO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseSpeedTestV2ResultVO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


