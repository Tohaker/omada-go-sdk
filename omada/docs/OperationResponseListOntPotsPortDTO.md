# OperationResponseListOntPotsPortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]OntPotsPortDTO**](OntPotsPortDTO.md) |  | [optional] 

## Methods

### NewOperationResponseListOntPotsPortDTO

`func NewOperationResponseListOntPotsPortDTO() *OperationResponseListOntPotsPortDTO`

NewOperationResponseListOntPotsPortDTO instantiates a new OperationResponseListOntPotsPortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListOntPotsPortDTOWithDefaults

`func NewOperationResponseListOntPotsPortDTOWithDefaults() *OperationResponseListOntPotsPortDTO`

NewOperationResponseListOntPotsPortDTOWithDefaults instantiates a new OperationResponseListOntPotsPortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListOntPotsPortDTO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListOntPotsPortDTO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListOntPotsPortDTO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListOntPotsPortDTO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListOntPotsPortDTO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListOntPotsPortDTO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListOntPotsPortDTO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListOntPotsPortDTO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListOntPotsPortDTO) GetResult() []OntPotsPortDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListOntPotsPortDTO) GetResultOk() (*[]OntPotsPortDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListOntPotsPortDTO) SetResult(v []OntPotsPortDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListOntPotsPortDTO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


