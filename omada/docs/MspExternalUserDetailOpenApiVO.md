# MspExternalUserDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerRoleId** | Pointer to **string** | The ID of customer role which is used by this user. | [optional] 
**CustomerRoleName** | Pointer to **string** | The name of customer role which is used by this user. | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Id** | Pointer to **string** | External user ID. | [optional] 
**IdpId** | Pointer to **string** | The ID of IdP which is used by this user. | [optional] 
**IdpName** | Pointer to **string** | The name of IdP which is used by this user. | [optional] 
**MspPrivilege** | Pointer to [**MspPrivilegeOpenApiVO**](MspPrivilegeOpenApiVO.md) |  | [optional] 
**Name** | Pointer to **string** | External user name. | [optional] 
**RoleId** | Pointer to **string** | The ID of role which is used by this user. | [optional] 
**RoleName** | Pointer to **string** | The name of role which is used by this user. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 
**TemporaryValidity** | Pointer to **int32** | Whether the temporary user is still valid | [optional] 

## Methods

### NewMspExternalUserDetailOpenApiVO

`func NewMspExternalUserDetailOpenApiVO() *MspExternalUserDetailOpenApiVO`

NewMspExternalUserDetailOpenApiVO instantiates a new MspExternalUserDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspExternalUserDetailOpenApiVOWithDefaults

`func NewMspExternalUserDetailOpenApiVOWithDefaults() *MspExternalUserDetailOpenApiVO`

NewMspExternalUserDetailOpenApiVOWithDefaults instantiates a new MspExternalUserDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerRoleId

`func (o *MspExternalUserDetailOpenApiVO) GetCustomerRoleId() string`

GetCustomerRoleId returns the CustomerRoleId field if non-nil, zero value otherwise.

### GetCustomerRoleIdOk

`func (o *MspExternalUserDetailOpenApiVO) GetCustomerRoleIdOk() (*string, bool)`

GetCustomerRoleIdOk returns a tuple with the CustomerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleId

`func (o *MspExternalUserDetailOpenApiVO) SetCustomerRoleId(v string)`

SetCustomerRoleId sets CustomerRoleId field to given value.

### HasCustomerRoleId

`func (o *MspExternalUserDetailOpenApiVO) HasCustomerRoleId() bool`

HasCustomerRoleId returns a boolean if a field has been set.

### GetCustomerRoleName

`func (o *MspExternalUserDetailOpenApiVO) GetCustomerRoleName() string`

GetCustomerRoleName returns the CustomerRoleName field if non-nil, zero value otherwise.

### GetCustomerRoleNameOk

`func (o *MspExternalUserDetailOpenApiVO) GetCustomerRoleNameOk() (*string, bool)`

GetCustomerRoleNameOk returns a tuple with the CustomerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleName

`func (o *MspExternalUserDetailOpenApiVO) SetCustomerRoleName(v string)`

SetCustomerRoleName sets CustomerRoleName field to given value.

### HasCustomerRoleName

`func (o *MspExternalUserDetailOpenApiVO) HasCustomerRoleName() bool`

HasCustomerRoleName returns a boolean if a field has been set.

### GetEndTime

`func (o *MspExternalUserDetailOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MspExternalUserDetailOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MspExternalUserDetailOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MspExternalUserDetailOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *MspExternalUserDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MspExternalUserDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MspExternalUserDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MspExternalUserDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpId

`func (o *MspExternalUserDetailOpenApiVO) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *MspExternalUserDetailOpenApiVO) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *MspExternalUserDetailOpenApiVO) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.

### HasIdpId

`func (o *MspExternalUserDetailOpenApiVO) HasIdpId() bool`

HasIdpId returns a boolean if a field has been set.

### GetIdpName

`func (o *MspExternalUserDetailOpenApiVO) GetIdpName() string`

GetIdpName returns the IdpName field if non-nil, zero value otherwise.

### GetIdpNameOk

`func (o *MspExternalUserDetailOpenApiVO) GetIdpNameOk() (*string, bool)`

GetIdpNameOk returns a tuple with the IdpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpName

`func (o *MspExternalUserDetailOpenApiVO) SetIdpName(v string)`

SetIdpName sets IdpName field to given value.

### HasIdpName

`func (o *MspExternalUserDetailOpenApiVO) HasIdpName() bool`

HasIdpName returns a boolean if a field has been set.

### GetMspPrivilege

`func (o *MspExternalUserDetailOpenApiVO) GetMspPrivilege() MspPrivilegeOpenApiVO`

GetMspPrivilege returns the MspPrivilege field if non-nil, zero value otherwise.

### GetMspPrivilegeOk

`func (o *MspExternalUserDetailOpenApiVO) GetMspPrivilegeOk() (*MspPrivilegeOpenApiVO, bool)`

GetMspPrivilegeOk returns a tuple with the MspPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspPrivilege

`func (o *MspExternalUserDetailOpenApiVO) SetMspPrivilege(v MspPrivilegeOpenApiVO)`

SetMspPrivilege sets MspPrivilege field to given value.

### HasMspPrivilege

`func (o *MspExternalUserDetailOpenApiVO) HasMspPrivilege() bool`

HasMspPrivilege returns a boolean if a field has been set.

### GetName

`func (o *MspExternalUserDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MspExternalUserDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MspExternalUserDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MspExternalUserDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleId

`func (o *MspExternalUserDetailOpenApiVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MspExternalUserDetailOpenApiVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MspExternalUserDetailOpenApiVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *MspExternalUserDetailOpenApiVO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *MspExternalUserDetailOpenApiVO) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *MspExternalUserDetailOpenApiVO) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *MspExternalUserDetailOpenApiVO) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *MspExternalUserDetailOpenApiVO) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetStartTime

`func (o *MspExternalUserDetailOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MspExternalUserDetailOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MspExternalUserDetailOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MspExternalUserDetailOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *MspExternalUserDetailOpenApiVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *MspExternalUserDetailOpenApiVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *MspExternalUserDetailOpenApiVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *MspExternalUserDetailOpenApiVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetTemporaryValidity

`func (o *MspExternalUserDetailOpenApiVO) GetTemporaryValidity() int32`

GetTemporaryValidity returns the TemporaryValidity field if non-nil, zero value otherwise.

### GetTemporaryValidityOk

`func (o *MspExternalUserDetailOpenApiVO) GetTemporaryValidityOk() (*int32, bool)`

GetTemporaryValidityOk returns a tuple with the TemporaryValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryValidity

`func (o *MspExternalUserDetailOpenApiVO) SetTemporaryValidity(v int32)`

SetTemporaryValidity sets TemporaryValidity field to given value.

### HasTemporaryValidity

`func (o *MspExternalUserDetailOpenApiVO) HasTemporaryValidity() bool`

HasTemporaryValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


