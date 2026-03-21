# MspUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** |  | [optional] 
**CustomerRoleId** | Pointer to **string** |  | [optional] 
**CustomerRoleName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**ForceModify** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OmadacId** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **bool** |  | [optional] 
**ParentUserId** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Privilege** | Pointer to [**MspPrivilegeVO**](MspPrivilegeVO.md) |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**ShowTree** | Pointer to **bool** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**TemporaryEnable** | Pointer to **bool** |  | [optional] 
**TemporaryValidity** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Type** | **int32** |  | 
**Verified** | Pointer to **bool** |  | [optional] 
**ViewerParentUserId** | Pointer to **string** |  | [optional] 

## Methods

### NewMspUserVO

`func NewMspUserVO(name string, type_ int32, ) *MspUserVO`

NewMspUserVO instantiates a new MspUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspUserVOWithDefaults

`func NewMspUserVOWithDefaults() *MspUserVO`

NewMspUserVOWithDefaults instantiates a new MspUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *MspUserVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *MspUserVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *MspUserVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *MspUserVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetCustomerRoleId

`func (o *MspUserVO) GetCustomerRoleId() string`

GetCustomerRoleId returns the CustomerRoleId field if non-nil, zero value otherwise.

### GetCustomerRoleIdOk

`func (o *MspUserVO) GetCustomerRoleIdOk() (*string, bool)`

GetCustomerRoleIdOk returns a tuple with the CustomerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleId

`func (o *MspUserVO) SetCustomerRoleId(v string)`

SetCustomerRoleId sets CustomerRoleId field to given value.

### HasCustomerRoleId

`func (o *MspUserVO) HasCustomerRoleId() bool`

HasCustomerRoleId returns a boolean if a field has been set.

### GetCustomerRoleName

`func (o *MspUserVO) GetCustomerRoleName() string`

GetCustomerRoleName returns the CustomerRoleName field if non-nil, zero value otherwise.

### GetCustomerRoleNameOk

`func (o *MspUserVO) GetCustomerRoleNameOk() (*string, bool)`

GetCustomerRoleNameOk returns a tuple with the CustomerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleName

`func (o *MspUserVO) SetCustomerRoleName(v string)`

SetCustomerRoleName sets CustomerRoleName field to given value.

### HasCustomerRoleName

`func (o *MspUserVO) HasCustomerRoleName() bool`

HasCustomerRoleName returns a boolean if a field has been set.

### GetEmail

`func (o *MspUserVO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MspUserVO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MspUserVO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MspUserVO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *MspUserVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MspUserVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MspUserVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MspUserVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetForceModify

`func (o *MspUserVO) GetForceModify() bool`

GetForceModify returns the ForceModify field if non-nil, zero value otherwise.

### GetForceModifyOk

`func (o *MspUserVO) GetForceModifyOk() (*bool, bool)`

GetForceModifyOk returns a tuple with the ForceModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceModify

`func (o *MspUserVO) SetForceModify(v bool)`

SetForceModify sets ForceModify field to given value.

### HasForceModify

`func (o *MspUserVO) HasForceModify() bool`

HasForceModify returns a boolean if a field has been set.

### GetId

`func (o *MspUserVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MspUserVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MspUserVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MspUserVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MspUserVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MspUserVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MspUserVO) SetName(v string)`

SetName sets Name field to given value.


### GetOmadacId

`func (o *MspUserVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *MspUserVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *MspUserVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *MspUserVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOwner

`func (o *MspUserVO) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MspUserVO) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MspUserVO) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MspUserVO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentUserId

`func (o *MspUserVO) GetParentUserId() string`

GetParentUserId returns the ParentUserId field if non-nil, zero value otherwise.

### GetParentUserIdOk

`func (o *MspUserVO) GetParentUserIdOk() (*string, bool)`

GetParentUserIdOk returns a tuple with the ParentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUserId

`func (o *MspUserVO) SetParentUserId(v string)`

SetParentUserId sets ParentUserId field to given value.

### HasParentUserId

`func (o *MspUserVO) HasParentUserId() bool`

HasParentUserId returns a boolean if a field has been set.

### GetPassword

`func (o *MspUserVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MspUserVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MspUserVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MspUserVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivilege

`func (o *MspUserVO) GetPrivilege() MspPrivilegeVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *MspUserVO) GetPrivilegeOk() (*MspPrivilegeVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *MspUserVO) SetPrivilege(v MspPrivilegeVO)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *MspUserVO) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetRoleId

`func (o *MspUserVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MspUserVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MspUserVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *MspUserVO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *MspUserVO) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *MspUserVO) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *MspUserVO) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *MspUserVO) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetShowTree

`func (o *MspUserVO) GetShowTree() bool`

GetShowTree returns the ShowTree field if non-nil, zero value otherwise.

### GetShowTreeOk

`func (o *MspUserVO) GetShowTreeOk() (*bool, bool)`

GetShowTreeOk returns a tuple with the ShowTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTree

`func (o *MspUserVO) SetShowTree(v bool)`

SetShowTree sets ShowTree field to given value.

### HasShowTree

`func (o *MspUserVO) HasShowTree() bool`

HasShowTree returns a boolean if a field has been set.

### GetStartTime

`func (o *MspUserVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MspUserVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MspUserVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MspUserVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *MspUserVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *MspUserVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *MspUserVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *MspUserVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetTemporaryValidity

`func (o *MspUserVO) GetTemporaryValidity() int32`

GetTemporaryValidity returns the TemporaryValidity field if non-nil, zero value otherwise.

### GetTemporaryValidityOk

`func (o *MspUserVO) GetTemporaryValidityOk() (*int32, bool)`

GetTemporaryValidityOk returns a tuple with the TemporaryValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryValidity

`func (o *MspUserVO) SetTemporaryValidity(v int32)`

SetTemporaryValidity sets TemporaryValidity field to given value.

### HasTemporaryValidity

`func (o *MspUserVO) HasTemporaryValidity() bool`

HasTemporaryValidity returns a boolean if a field has been set.

### GetToken

`func (o *MspUserVO) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MspUserVO) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MspUserVO) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MspUserVO) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *MspUserVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MspUserVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MspUserVO) SetType(v int32)`

SetType sets Type field to given value.


### GetVerified

`func (o *MspUserVO) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *MspUserVO) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *MspUserVO) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *MspUserVO) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetViewerParentUserId

`func (o *MspUserVO) GetViewerParentUserId() string`

GetViewerParentUserId returns the ViewerParentUserId field if non-nil, zero value otherwise.

### GetViewerParentUserIdOk

`func (o *MspUserVO) GetViewerParentUserIdOk() (*string, bool)`

GetViewerParentUserIdOk returns a tuple with the ViewerParentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerParentUserId

`func (o *MspUserVO) SetViewerParentUserId(v string)`

SetViewerParentUserId sets ViewerParentUserId field to given value.

### HasViewerParentUserId

`func (o *MspUserVO) HasViewerParentUserId() bool`

HasViewerParentUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


