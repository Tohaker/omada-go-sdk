# ModifyRoleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Role name should contain 1 to 128 ASCII characters. | 
**Privilege** | Pointer to [**RoleVO**](RoleVO.md) |  | [optional] 

## Methods

### NewModifyRoleVO

`func NewModifyRoleVO(name string, ) *ModifyRoleVO`

NewModifyRoleVO instantiates a new ModifyRoleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyRoleVOWithDefaults

`func NewModifyRoleVOWithDefaults() *ModifyRoleVO`

NewModifyRoleVOWithDefaults instantiates a new ModifyRoleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModifyRoleVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyRoleVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyRoleVO) SetName(v string)`

SetName sets Name field to given value.


### GetPrivilege

`func (o *ModifyRoleVO) GetPrivilege() RoleVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *ModifyRoleVO) GetPrivilegeOk() (*RoleVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *ModifyRoleVO) SetPrivilege(v RoleVO)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *ModifyRoleVO) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


