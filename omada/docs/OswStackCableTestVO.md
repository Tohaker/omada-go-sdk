# OswStackCableTestVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to [**[]OswStackMemberCableTestVO**](OswStackMemberCableTestVO.md) | stack member list | [optional] 
**Name** | Pointer to **string** | stack name | [optional] 
**StackId** | Pointer to **string** | stack ID | [optional] 

## Methods

### NewOswStackCableTestVO

`func NewOswStackCableTestVO() *OswStackCableTestVO`

NewOswStackCableTestVO instantiates a new OswStackCableTestVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackCableTestVOWithDefaults

`func NewOswStackCableTestVOWithDefaults() *OswStackCableTestVO`

NewOswStackCableTestVOWithDefaults instantiates a new OswStackCableTestVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *OswStackCableTestVO) GetMember() []OswStackMemberCableTestVO`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OswStackCableTestVO) GetMemberOk() (*[]OswStackMemberCableTestVO, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OswStackCableTestVO) SetMember(v []OswStackMemberCableTestVO)`

SetMember sets Member field to given value.

### HasMember

`func (o *OswStackCableTestVO) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *OswStackCableTestVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackCableTestVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackCableTestVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStackCableTestVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStackId

`func (o *OswStackCableTestVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswStackCableTestVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswStackCableTestVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswStackCableTestVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


