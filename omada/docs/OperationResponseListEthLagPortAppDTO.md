# OperationResponseListEthLagPortAppDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]EthLagPortAppDTO**](EthLagPortAppDTO.md) |  | [optional] 

## Methods

### NewOperationResponseListEthLagPortAppDTO

`func NewOperationResponseListEthLagPortAppDTO() *OperationResponseListEthLagPortAppDTO`

NewOperationResponseListEthLagPortAppDTO instantiates a new OperationResponseListEthLagPortAppDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListEthLagPortAppDTOWithDefaults

`func NewOperationResponseListEthLagPortAppDTOWithDefaults() *OperationResponseListEthLagPortAppDTO`

NewOperationResponseListEthLagPortAppDTOWithDefaults instantiates a new OperationResponseListEthLagPortAppDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListEthLagPortAppDTO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListEthLagPortAppDTO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListEthLagPortAppDTO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListEthLagPortAppDTO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListEthLagPortAppDTO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListEthLagPortAppDTO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListEthLagPortAppDTO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListEthLagPortAppDTO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListEthLagPortAppDTO) GetResult() []EthLagPortAppDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListEthLagPortAppDTO) GetResultOk() (*[]EthLagPortAppDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListEthLagPortAppDTO) SetResult(v []EthLagPortAppDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListEthLagPortAppDTO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


