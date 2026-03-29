# OperationResponseListOnuInformationRebootStatusConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]OnuInformationRebootStatusConfigDTO**](OnuInformationRebootStatusConfigDTO.md) |  | [optional] 

## Methods

### NewOperationResponseListOnuInformationRebootStatusConfigDTO

`func NewOperationResponseListOnuInformationRebootStatusConfigDTO() *OperationResponseListOnuInformationRebootStatusConfigDTO`

NewOperationResponseListOnuInformationRebootStatusConfigDTO instantiates a new OperationResponseListOnuInformationRebootStatusConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListOnuInformationRebootStatusConfigDTOWithDefaults

`func NewOperationResponseListOnuInformationRebootStatusConfigDTOWithDefaults() *OperationResponseListOnuInformationRebootStatusConfigDTO`

NewOperationResponseListOnuInformationRebootStatusConfigDTOWithDefaults instantiates a new OperationResponseListOnuInformationRebootStatusConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) GetResult() []OnuInformationRebootStatusConfigDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) GetResultOk() (*[]OnuInformationRebootStatusConfigDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) SetResult(v []OnuInformationRebootStatusConfigDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListOnuInformationRebootStatusConfigDTO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


