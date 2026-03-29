# ControllerRoleDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRole** | Pointer to **bool** | Whether role is default role | [optional] 
**Id** | Pointer to **string** | Role ID | [optional] 
**Name** | Pointer to **string** | Role Name | [optional] 
**Privilege** | Pointer to [**ControllerRoleVO**](ControllerRoleVO.md) |  | [optional] 
**Source** | Pointer to **int32** | Role created resource. It should be a value as follows: 0: default; 1:create by standard controller or customer controller; 2: create by MSP | [optional] 
**Type** | Pointer to **int32** | Role Type should be a value as follows: 0: standard; 1: customer; 2: msp. | [optional] 

## Methods

### NewControllerRoleDetailVO

`func NewControllerRoleDetailVO() *ControllerRoleDetailVO`

NewControllerRoleDetailVO instantiates a new ControllerRoleDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerRoleDetailVOWithDefaults

`func NewControllerRoleDetailVOWithDefaults() *ControllerRoleDetailVO`

NewControllerRoleDetailVOWithDefaults instantiates a new ControllerRoleDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRole

`func (o *ControllerRoleDetailVO) GetDefaultRole() bool`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *ControllerRoleDetailVO) GetDefaultRoleOk() (*bool, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *ControllerRoleDetailVO) SetDefaultRole(v bool)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *ControllerRoleDetailVO) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetId

`func (o *ControllerRoleDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerRoleDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerRoleDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerRoleDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ControllerRoleDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerRoleDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerRoleDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllerRoleDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivilege

`func (o *ControllerRoleDetailVO) GetPrivilege() ControllerRoleVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *ControllerRoleDetailVO) GetPrivilegeOk() (*ControllerRoleVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *ControllerRoleDetailVO) SetPrivilege(v ControllerRoleVO)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *ControllerRoleDetailVO) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetSource

`func (o *ControllerRoleDetailVO) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ControllerRoleDetailVO) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ControllerRoleDetailVO) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *ControllerRoleDetailVO) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *ControllerRoleDetailVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControllerRoleDetailVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControllerRoleDetailVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ControllerRoleDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


