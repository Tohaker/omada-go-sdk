# OperationResponseGatewayPortsConfigEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**GatewayPortsConfigEntity**](GatewayPortsConfigEntity.md) |  | [optional] 

## Methods

### NewOperationResponseGatewayPortsConfigEntity

`func NewOperationResponseGatewayPortsConfigEntity() *OperationResponseGatewayPortsConfigEntity`

NewOperationResponseGatewayPortsConfigEntity instantiates a new OperationResponseGatewayPortsConfigEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseGatewayPortsConfigEntityWithDefaults

`func NewOperationResponseGatewayPortsConfigEntityWithDefaults() *OperationResponseGatewayPortsConfigEntity`

NewOperationResponseGatewayPortsConfigEntityWithDefaults instantiates a new OperationResponseGatewayPortsConfigEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseGatewayPortsConfigEntity) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseGatewayPortsConfigEntity) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseGatewayPortsConfigEntity) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseGatewayPortsConfigEntity) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseGatewayPortsConfigEntity) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseGatewayPortsConfigEntity) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseGatewayPortsConfigEntity) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseGatewayPortsConfigEntity) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseGatewayPortsConfigEntity) GetResult() GatewayPortsConfigEntity`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseGatewayPortsConfigEntity) GetResultOk() (*GatewayPortsConfigEntity, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseGatewayPortsConfigEntity) SetResult(v GatewayPortsConfigEntity)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseGatewayPortsConfigEntity) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


