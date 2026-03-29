# ModifyMspRoleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Role name should contain 1 to 128 ASCII characters. | 
**Privilege** | [**MspRoleVO**](MspRoleVO.md) |  | 

## Methods

### NewModifyMspRoleVO

`func NewModifyMspRoleVO(name string, privilege MspRoleVO, ) *ModifyMspRoleVO`

NewModifyMspRoleVO instantiates a new ModifyMspRoleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMspRoleVOWithDefaults

`func NewModifyMspRoleVOWithDefaults() *ModifyMspRoleVO`

NewModifyMspRoleVOWithDefaults instantiates a new ModifyMspRoleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModifyMspRoleVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyMspRoleVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyMspRoleVO) SetName(v string)`

SetName sets Name field to given value.


### GetPrivilege

`func (o *ModifyMspRoleVO) GetPrivilege() MspRoleVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *ModifyMspRoleVO) GetPrivilegeOk() (*MspRoleVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *ModifyMspRoleVO) SetPrivilege(v MspRoleVO)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


