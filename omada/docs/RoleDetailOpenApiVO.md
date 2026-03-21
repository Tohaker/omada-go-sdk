# RoleDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRole** | Pointer to **bool** | Whether role is default role. | [optional] 
**HasAllPrivilege** | Pointer to **bool** | Whether role has all site privilege. | [optional] 
**Name** | **string** | Role name should contain 1 to 128 ASCII characters. | 
**Owner** | Pointer to **bool** | Whether role is owner. | [optional] 
**Privilege** | Pointer to [**RoleVO**](RoleVO.md) |  | [optional] 
**Source** | Pointer to **int32** | Role created resource. It should be a value as follows: 0: default; 1:create by standard controller or customer controller; 2: create by MSP | [optional] 
**Type** | Pointer to **int32** | Role Type should be a value as follows: 0: standard; 1: customer; 2: msp. | [optional] 

## Methods

### NewRoleDetailOpenApiVO

`func NewRoleDetailOpenApiVO(name string, ) *RoleDetailOpenApiVO`

NewRoleDetailOpenApiVO instantiates a new RoleDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDetailOpenApiVOWithDefaults

`func NewRoleDetailOpenApiVOWithDefaults() *RoleDetailOpenApiVO`

NewRoleDetailOpenApiVOWithDefaults instantiates a new RoleDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRole

`func (o *RoleDetailOpenApiVO) GetDefaultRole() bool`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *RoleDetailOpenApiVO) GetDefaultRoleOk() (*bool, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *RoleDetailOpenApiVO) SetDefaultRole(v bool)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *RoleDetailOpenApiVO) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetHasAllPrivilege

`func (o *RoleDetailOpenApiVO) GetHasAllPrivilege() bool`

GetHasAllPrivilege returns the HasAllPrivilege field if non-nil, zero value otherwise.

### GetHasAllPrivilegeOk

`func (o *RoleDetailOpenApiVO) GetHasAllPrivilegeOk() (*bool, bool)`

GetHasAllPrivilegeOk returns a tuple with the HasAllPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAllPrivilege

`func (o *RoleDetailOpenApiVO) SetHasAllPrivilege(v bool)`

SetHasAllPrivilege sets HasAllPrivilege field to given value.

### HasHasAllPrivilege

`func (o *RoleDetailOpenApiVO) HasHasAllPrivilege() bool`

HasHasAllPrivilege returns a boolean if a field has been set.

### GetName

`func (o *RoleDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *RoleDetailOpenApiVO) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RoleDetailOpenApiVO) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RoleDetailOpenApiVO) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *RoleDetailOpenApiVO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivilege

`func (o *RoleDetailOpenApiVO) GetPrivilege() RoleVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *RoleDetailOpenApiVO) GetPrivilegeOk() (*RoleVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *RoleDetailOpenApiVO) SetPrivilege(v RoleVO)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *RoleDetailOpenApiVO) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetSource

`func (o *RoleDetailOpenApiVO) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RoleDetailOpenApiVO) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RoleDetailOpenApiVO) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *RoleDetailOpenApiVO) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *RoleDetailOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleDetailOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleDetailOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RoleDetailOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


