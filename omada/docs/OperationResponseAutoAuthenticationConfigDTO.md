# OperationResponseAutoAuthenticationConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**AutoAuthenticationConfigDTO**](AutoAuthenticationConfigDTO.md) |  | [optional] 

## Methods

### NewOperationResponseAutoAuthenticationConfigDTO

`func NewOperationResponseAutoAuthenticationConfigDTO() *OperationResponseAutoAuthenticationConfigDTO`

NewOperationResponseAutoAuthenticationConfigDTO instantiates a new OperationResponseAutoAuthenticationConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseAutoAuthenticationConfigDTOWithDefaults

`func NewOperationResponseAutoAuthenticationConfigDTOWithDefaults() *OperationResponseAutoAuthenticationConfigDTO`

NewOperationResponseAutoAuthenticationConfigDTOWithDefaults instantiates a new OperationResponseAutoAuthenticationConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseAutoAuthenticationConfigDTO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseAutoAuthenticationConfigDTO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseAutoAuthenticationConfigDTO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseAutoAuthenticationConfigDTO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseAutoAuthenticationConfigDTO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseAutoAuthenticationConfigDTO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseAutoAuthenticationConfigDTO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseAutoAuthenticationConfigDTO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseAutoAuthenticationConfigDTO) GetResult() AutoAuthenticationConfigDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseAutoAuthenticationConfigDTO) GetResultOk() (*AutoAuthenticationConfigDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseAutoAuthenticationConfigDTO) SetResult(v AutoAuthenticationConfigDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseAutoAuthenticationConfigDTO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


