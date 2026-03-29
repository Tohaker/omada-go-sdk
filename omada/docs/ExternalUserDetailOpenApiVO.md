# ExternalUserDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Id** | Pointer to **string** | External user ID. | [optional] 
**IdpId** | Pointer to **string** | The ID of IdP which is used by this user. | [optional] 
**IdpName** | Pointer to **string** | The name of IdP which is used by this user. | [optional] 
**Name** | Pointer to **string** | External user name. | [optional] 
**Privilege** | Pointer to [**PrivilegeOpenApiVO**](PrivilegeOpenApiVO.md) |  | [optional] 
**RoleId** | Pointer to **string** | The ID of role which is used by this user. | [optional] 
**RoleName** | Pointer to **string** | The name of role which is used by this user. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 
**TemporaryValidity** | Pointer to **int32** | Whether the temporary user is still valid | [optional] 

## Methods

### NewExternalUserDetailOpenApiVO

`func NewExternalUserDetailOpenApiVO() *ExternalUserDetailOpenApiVO`

NewExternalUserDetailOpenApiVO instantiates a new ExternalUserDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalUserDetailOpenApiVOWithDefaults

`func NewExternalUserDetailOpenApiVOWithDefaults() *ExternalUserDetailOpenApiVO`

NewExternalUserDetailOpenApiVOWithDefaults instantiates a new ExternalUserDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *ExternalUserDetailOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ExternalUserDetailOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ExternalUserDetailOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ExternalUserDetailOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *ExternalUserDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalUserDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalUserDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalUserDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpId

`func (o *ExternalUserDetailOpenApiVO) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *ExternalUserDetailOpenApiVO) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *ExternalUserDetailOpenApiVO) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.

### HasIdpId

`func (o *ExternalUserDetailOpenApiVO) HasIdpId() bool`

HasIdpId returns a boolean if a field has been set.

### GetIdpName

`func (o *ExternalUserDetailOpenApiVO) GetIdpName() string`

GetIdpName returns the IdpName field if non-nil, zero value otherwise.

### GetIdpNameOk

`func (o *ExternalUserDetailOpenApiVO) GetIdpNameOk() (*string, bool)`

GetIdpNameOk returns a tuple with the IdpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpName

`func (o *ExternalUserDetailOpenApiVO) SetIdpName(v string)`

SetIdpName sets IdpName field to given value.

### HasIdpName

`func (o *ExternalUserDetailOpenApiVO) HasIdpName() bool`

HasIdpName returns a boolean if a field has been set.

### GetName

`func (o *ExternalUserDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalUserDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalUserDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalUserDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivilege

`func (o *ExternalUserDetailOpenApiVO) GetPrivilege() PrivilegeOpenApiVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *ExternalUserDetailOpenApiVO) GetPrivilegeOk() (*PrivilegeOpenApiVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *ExternalUserDetailOpenApiVO) SetPrivilege(v PrivilegeOpenApiVO)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *ExternalUserDetailOpenApiVO) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetRoleId

`func (o *ExternalUserDetailOpenApiVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ExternalUserDetailOpenApiVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ExternalUserDetailOpenApiVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ExternalUserDetailOpenApiVO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *ExternalUserDetailOpenApiVO) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ExternalUserDetailOpenApiVO) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ExternalUserDetailOpenApiVO) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *ExternalUserDetailOpenApiVO) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetStartTime

`func (o *ExternalUserDetailOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExternalUserDetailOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExternalUserDetailOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ExternalUserDetailOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *ExternalUserDetailOpenApiVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *ExternalUserDetailOpenApiVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *ExternalUserDetailOpenApiVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *ExternalUserDetailOpenApiVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetTemporaryValidity

`func (o *ExternalUserDetailOpenApiVO) GetTemporaryValidity() int32`

GetTemporaryValidity returns the TemporaryValidity field if non-nil, zero value otherwise.

### GetTemporaryValidityOk

`func (o *ExternalUserDetailOpenApiVO) GetTemporaryValidityOk() (*int32, bool)`

GetTemporaryValidityOk returns a tuple with the TemporaryValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryValidity

`func (o *ExternalUserDetailOpenApiVO) SetTemporaryValidity(v int32)`

SetTemporaryValidity sets TemporaryValidity field to given value.

### HasTemporaryValidity

`func (o *ExternalUserDetailOpenApiVO) HasTemporaryValidity() bool`

HasTemporaryValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


