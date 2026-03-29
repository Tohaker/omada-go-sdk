# OswStackInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterMac** | Pointer to **string** | The mac of the master device of the stack device. | [optional] 
**Members** | Pointer to [**[]StackMemberVO**](StackMemberVO.md) | The members of the stack device. | [optional] 
**StackId** | Pointer to **string** | The stack id of the stack device. | [optional] 
**StackName** | Pointer to **string** | The stack name of the stack device. | [optional] 

## Methods

### NewOswStackInfoVO

`func NewOswStackInfoVO() *OswStackInfoVO`

NewOswStackInfoVO instantiates a new OswStackInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackInfoVOWithDefaults

`func NewOswStackInfoVOWithDefaults() *OswStackInfoVO`

NewOswStackInfoVOWithDefaults instantiates a new OswStackInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterMac

`func (o *OswStackInfoVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *OswStackInfoVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *OswStackInfoVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *OswStackInfoVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetMembers

`func (o *OswStackInfoVO) GetMembers() []StackMemberVO`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OswStackInfoVO) GetMembersOk() (*[]StackMemberVO, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OswStackInfoVO) SetMembers(v []StackMemberVO)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OswStackInfoVO) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetStackId

`func (o *OswStackInfoVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswStackInfoVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswStackInfoVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswStackInfoVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswStackInfoVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswStackInfoVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswStackInfoVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswStackInfoVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


