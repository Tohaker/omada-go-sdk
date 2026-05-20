# OswStackSdmApplicationVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurTmplName** | Pointer to **string** | Current used sdm template. | [optional] 
**StackMemberSdmList** | Pointer to [**[]OswStackMemberSdmVO**](OswStackMemberSdmVO.md) | The sdm resources usage detail of stack members | [optional] 
**TcamUtilization** | Pointer to **int32** | The overall TCAM utilization, calculated as usedTcam / totalTcam. | [optional] 

## Methods

### NewOswStackSdmApplicationVO

`func NewOswStackSdmApplicationVO() *OswStackSdmApplicationVO`

NewOswStackSdmApplicationVO instantiates a new OswStackSdmApplicationVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackSdmApplicationVOWithDefaults

`func NewOswStackSdmApplicationVOWithDefaults() *OswStackSdmApplicationVO`

NewOswStackSdmApplicationVOWithDefaults instantiates a new OswStackSdmApplicationVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurTmplName

`func (o *OswStackSdmApplicationVO) GetCurTmplName() string`

GetCurTmplName returns the CurTmplName field if non-nil, zero value otherwise.

### GetCurTmplNameOk

`func (o *OswStackSdmApplicationVO) GetCurTmplNameOk() (*string, bool)`

GetCurTmplNameOk returns a tuple with the CurTmplName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurTmplName

`func (o *OswStackSdmApplicationVO) SetCurTmplName(v string)`

SetCurTmplName sets CurTmplName field to given value.

### HasCurTmplName

`func (o *OswStackSdmApplicationVO) HasCurTmplName() bool`

HasCurTmplName returns a boolean if a field has been set.

### GetStackMemberSdmList

`func (o *OswStackSdmApplicationVO) GetStackMemberSdmList() []OswStackMemberSdmVO`

GetStackMemberSdmList returns the StackMemberSdmList field if non-nil, zero value otherwise.

### GetStackMemberSdmListOk

`func (o *OswStackSdmApplicationVO) GetStackMemberSdmListOk() (*[]OswStackMemberSdmVO, bool)`

GetStackMemberSdmListOk returns a tuple with the StackMemberSdmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackMemberSdmList

`func (o *OswStackSdmApplicationVO) SetStackMemberSdmList(v []OswStackMemberSdmVO)`

SetStackMemberSdmList sets StackMemberSdmList field to given value.

### HasStackMemberSdmList

`func (o *OswStackSdmApplicationVO) HasStackMemberSdmList() bool`

HasStackMemberSdmList returns a boolean if a field has been set.

### GetTcamUtilization

`func (o *OswStackSdmApplicationVO) GetTcamUtilization() int32`

GetTcamUtilization returns the TcamUtilization field if non-nil, zero value otherwise.

### GetTcamUtilizationOk

`func (o *OswStackSdmApplicationVO) GetTcamUtilizationOk() (*int32, bool)`

GetTcamUtilizationOk returns a tuple with the TcamUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcamUtilization

`func (o *OswStackSdmApplicationVO) SetTcamUtilization(v int32)`

SetTcamUtilization sets TcamUtilization field to given value.

### HasTcamUtilization

`func (o *OswStackSdmApplicationVO) HasTcamUtilization() bool`

HasTcamUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


