# OswCableTestResultWithStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]OswCableTestResultVO**](OswCableTestResultVO.md) | Test results. | [optional] 
**StackData** | Pointer to [**[]OswStackMemberCableTestResultVO**](OswStackMemberCableTestResultVO.md) | Stack Test results | [optional] 
**Status** | Pointer to **int32** | Test status.It should be a value as follows: 0:free. 1:testing. 2:test done. 3:time out. | [optional] 

## Methods

### NewOswCableTestResultWithStatusVO

`func NewOswCableTestResultWithStatusVO() *OswCableTestResultWithStatusVO`

NewOswCableTestResultWithStatusVO instantiates a new OswCableTestResultWithStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswCableTestResultWithStatusVOWithDefaults

`func NewOswCableTestResultWithStatusVOWithDefaults() *OswCableTestResultWithStatusVO`

NewOswCableTestResultWithStatusVOWithDefaults instantiates a new OswCableTestResultWithStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OswCableTestResultWithStatusVO) GetData() []OswCableTestResultVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OswCableTestResultWithStatusVO) GetDataOk() (*[]OswCableTestResultVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OswCableTestResultWithStatusVO) SetData(v []OswCableTestResultVO)`

SetData sets Data field to given value.

### HasData

`func (o *OswCableTestResultWithStatusVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStackData

`func (o *OswCableTestResultWithStatusVO) GetStackData() []OswStackMemberCableTestResultVO`

GetStackData returns the StackData field if non-nil, zero value otherwise.

### GetStackDataOk

`func (o *OswCableTestResultWithStatusVO) GetStackDataOk() (*[]OswStackMemberCableTestResultVO, bool)`

GetStackDataOk returns a tuple with the StackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackData

`func (o *OswCableTestResultWithStatusVO) SetStackData(v []OswStackMemberCableTestResultVO)`

SetStackData sets StackData field to given value.

### HasStackData

`func (o *OswCableTestResultWithStatusVO) HasStackData() bool`

HasStackData returns a boolean if a field has been set.

### GetStatus

`func (o *OswCableTestResultWithStatusVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswCableTestResultWithStatusVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswCableTestResultWithStatusVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswCableTestResultWithStatusVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


