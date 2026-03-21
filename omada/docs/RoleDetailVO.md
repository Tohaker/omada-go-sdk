# RoleDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRole** | Pointer to **bool** | Whether role is default role. | [optional] 
**HasAllPrivilege** | Pointer to **bool** | Whether role has all site privilege. | [optional] 
**Id** | Pointer to **string** | Role Detail ID | [optional] 
**Name** | **string** | Role name should contain 1 to 128 ASCII characters. | 
**Owner** | Pointer to **bool** | Whether role is owner. | [optional] 
**Privilege** | Pointer to [**RoleVO**](RoleVO.md) |  | [optional] 
**Source** | Pointer to **int32** | Role created resource. It should be a value as follows: 0: default; 1:create by standard controller or customer controller; 2: create by MSP | [optional] 
**Type** | Pointer to **int32** | Role Type should be a value as follows: 0: standard; 1: customer; 2: msp. | [optional] 

## Methods

### NewRoleDetailVO

`func NewRoleDetailVO(name string, ) *RoleDetailVO`

NewRoleDetailVO instantiates a new RoleDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDetailVOWithDefaults

`func NewRoleDetailVOWithDefaults() *RoleDetailVO`

NewRoleDetailVOWithDefaults instantiates a new RoleDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRole

`func (o *RoleDetailVO) GetDefaultRole() bool`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *RoleDetailVO) GetDefaultRoleOk() (*bool, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *RoleDetailVO) SetDefaultRole(v bool)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *RoleDetailVO) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetHasAllPrivilege

`func (o *RoleDetailVO) GetHasAllPrivilege() bool`

GetHasAllPrivilege returns the HasAllPrivilege field if non-nil, zero value otherwise.

### GetHasAllPrivilegeOk

`func (o *RoleDetailVO) GetHasAllPrivilegeOk() (*bool, bool)`

GetHasAllPrivilegeOk returns a tuple with the HasAllPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAllPrivilege

`func (o *RoleDetailVO) SetHasAllPrivilege(v bool)`

SetHasAllPrivilege sets HasAllPrivilege field to given value.

### HasHasAllPrivilege

`func (o *RoleDetailVO) HasHasAllPrivilege() bool`

HasHasAllPrivilege returns a boolean if a field has been set.

### GetId

`func (o *RoleDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoleDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleDetailVO) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *RoleDetailVO) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RoleDetailVO) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RoleDetailVO) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *RoleDetailVO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivilege

`func (o *RoleDetailVO) GetPrivilege() RoleVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *RoleDetailVO) GetPrivilegeOk() (*RoleVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *RoleDetailVO) SetPrivilege(v RoleVO)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *RoleDetailVO) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetSource

`func (o *RoleDetailVO) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RoleDetailVO) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RoleDetailVO) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *RoleDetailVO) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *RoleDetailVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleDetailVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleDetailVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RoleDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


