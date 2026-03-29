# ModifyControllerRoleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Role name should contain 1 to 128 characters. | 
**Privilege** | Pointer to [**ControllerRoleVO**](ControllerRoleVO.md) |  | [optional] 

## Methods

### NewModifyControllerRoleVO

`func NewModifyControllerRoleVO(name string, ) *ModifyControllerRoleVO`

NewModifyControllerRoleVO instantiates a new ModifyControllerRoleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyControllerRoleVOWithDefaults

`func NewModifyControllerRoleVOWithDefaults() *ModifyControllerRoleVO`

NewModifyControllerRoleVOWithDefaults instantiates a new ModifyControllerRoleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModifyControllerRoleVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyControllerRoleVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyControllerRoleVO) SetName(v string)`

SetName sets Name field to given value.


### GetPrivilege

`func (o *ModifyControllerRoleVO) GetPrivilege() ControllerRoleVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *ModifyControllerRoleVO) GetPrivilegeOk() (*ControllerRoleVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *ModifyControllerRoleVO) SetPrivilege(v ControllerRoleVO)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *ModifyControllerRoleVO) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


