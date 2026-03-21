# ControllerUserAppVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IncidentNotification** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OmadacId** | Pointer to **string** |  | [optional] 
**ParentUserId** | Pointer to **string** |  | [optional] 
**Privilege** | Pointer to [**PrivilegeResultVO**](PrivilegeResultVO.md) |  | [optional] 
**RoleDetail** | Pointer to [**RoleDetailVO**](RoleDetailVO.md) |  | [optional] 
**RoleType** | Pointer to **int32** |  | [optional] 
**ShowTree** | Pointer to **bool** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**TemporaryEnable** | Pointer to **bool** |  | [optional] 
**TemporaryValidity** | Pointer to **int32** |  | [optional] 
**Type** | **int32** |  | 
**UserLevel** | Pointer to **int32** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**ViewerParentUserId** | Pointer to **string** |  | [optional] 

## Methods

### NewControllerUserAppVO

`func NewControllerUserAppVO(type_ int32, ) *ControllerUserAppVO`

NewControllerUserAppVO instantiates a new ControllerUserAppVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerUserAppVOWithDefaults

`func NewControllerUserAppVOWithDefaults() *ControllerUserAppVO`

NewControllerUserAppVOWithDefaults instantiates a new ControllerUserAppVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *ControllerUserAppVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *ControllerUserAppVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *ControllerUserAppVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *ControllerUserAppVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetEmail

`func (o *ControllerUserAppVO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ControllerUserAppVO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ControllerUserAppVO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ControllerUserAppVO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *ControllerUserAppVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ControllerUserAppVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ControllerUserAppVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ControllerUserAppVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *ControllerUserAppVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerUserAppVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerUserAppVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerUserAppVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncidentNotification

`func (o *ControllerUserAppVO) GetIncidentNotification() bool`

GetIncidentNotification returns the IncidentNotification field if non-nil, zero value otherwise.

### GetIncidentNotificationOk

`func (o *ControllerUserAppVO) GetIncidentNotificationOk() (*bool, bool)`

GetIncidentNotificationOk returns a tuple with the IncidentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentNotification

`func (o *ControllerUserAppVO) SetIncidentNotification(v bool)`

SetIncidentNotification sets IncidentNotification field to given value.

### HasIncidentNotification

`func (o *ControllerUserAppVO) HasIncidentNotification() bool`

HasIncidentNotification returns a boolean if a field has been set.

### GetName

`func (o *ControllerUserAppVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerUserAppVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerUserAppVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllerUserAppVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *ControllerUserAppVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *ControllerUserAppVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *ControllerUserAppVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *ControllerUserAppVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetParentUserId

`func (o *ControllerUserAppVO) GetParentUserId() string`

GetParentUserId returns the ParentUserId field if non-nil, zero value otherwise.

### GetParentUserIdOk

`func (o *ControllerUserAppVO) GetParentUserIdOk() (*string, bool)`

GetParentUserIdOk returns a tuple with the ParentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUserId

`func (o *ControllerUserAppVO) SetParentUserId(v string)`

SetParentUserId sets ParentUserId field to given value.

### HasParentUserId

`func (o *ControllerUserAppVO) HasParentUserId() bool`

HasParentUserId returns a boolean if a field has been set.

### GetPrivilege

`func (o *ControllerUserAppVO) GetPrivilege() PrivilegeResultVO`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *ControllerUserAppVO) GetPrivilegeOk() (*PrivilegeResultVO, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *ControllerUserAppVO) SetPrivilege(v PrivilegeResultVO)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *ControllerUserAppVO) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetRoleDetail

`func (o *ControllerUserAppVO) GetRoleDetail() RoleDetailVO`

GetRoleDetail returns the RoleDetail field if non-nil, zero value otherwise.

### GetRoleDetailOk

`func (o *ControllerUserAppVO) GetRoleDetailOk() (*RoleDetailVO, bool)`

GetRoleDetailOk returns a tuple with the RoleDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDetail

`func (o *ControllerUserAppVO) SetRoleDetail(v RoleDetailVO)`

SetRoleDetail sets RoleDetail field to given value.

### HasRoleDetail

`func (o *ControllerUserAppVO) HasRoleDetail() bool`

HasRoleDetail returns a boolean if a field has been set.

### GetRoleType

`func (o *ControllerUserAppVO) GetRoleType() int32`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *ControllerUserAppVO) GetRoleTypeOk() (*int32, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *ControllerUserAppVO) SetRoleType(v int32)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *ControllerUserAppVO) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetShowTree

`func (o *ControllerUserAppVO) GetShowTree() bool`

GetShowTree returns the ShowTree field if non-nil, zero value otherwise.

### GetShowTreeOk

`func (o *ControllerUserAppVO) GetShowTreeOk() (*bool, bool)`

GetShowTreeOk returns a tuple with the ShowTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTree

`func (o *ControllerUserAppVO) SetShowTree(v bool)`

SetShowTree sets ShowTree field to given value.

### HasShowTree

`func (o *ControllerUserAppVO) HasShowTree() bool`

HasShowTree returns a boolean if a field has been set.

### GetStartTime

`func (o *ControllerUserAppVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ControllerUserAppVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ControllerUserAppVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ControllerUserAppVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *ControllerUserAppVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *ControllerUserAppVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *ControllerUserAppVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *ControllerUserAppVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetTemporaryValidity

`func (o *ControllerUserAppVO) GetTemporaryValidity() int32`

GetTemporaryValidity returns the TemporaryValidity field if non-nil, zero value otherwise.

### GetTemporaryValidityOk

`func (o *ControllerUserAppVO) GetTemporaryValidityOk() (*int32, bool)`

GetTemporaryValidityOk returns a tuple with the TemporaryValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryValidity

`func (o *ControllerUserAppVO) SetTemporaryValidity(v int32)`

SetTemporaryValidity sets TemporaryValidity field to given value.

### HasTemporaryValidity

`func (o *ControllerUserAppVO) HasTemporaryValidity() bool`

HasTemporaryValidity returns a boolean if a field has been set.

### GetType

`func (o *ControllerUserAppVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControllerUserAppVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControllerUserAppVO) SetType(v int32)`

SetType sets Type field to given value.


### GetUserLevel

`func (o *ControllerUserAppVO) GetUserLevel() int32`

GetUserLevel returns the UserLevel field if non-nil, zero value otherwise.

### GetUserLevelOk

`func (o *ControllerUserAppVO) GetUserLevelOk() (*int32, bool)`

GetUserLevelOk returns a tuple with the UserLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLevel

`func (o *ControllerUserAppVO) SetUserLevel(v int32)`

SetUserLevel sets UserLevel field to given value.

### HasUserLevel

`func (o *ControllerUserAppVO) HasUserLevel() bool`

HasUserLevel returns a boolean if a field has been set.

### GetVerified

`func (o *ControllerUserAppVO) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ControllerUserAppVO) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ControllerUserAppVO) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ControllerUserAppVO) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetViewerParentUserId

`func (o *ControllerUserAppVO) GetViewerParentUserId() string`

GetViewerParentUserId returns the ViewerParentUserId field if non-nil, zero value otherwise.

### GetViewerParentUserIdOk

`func (o *ControllerUserAppVO) GetViewerParentUserIdOk() (*string, bool)`

GetViewerParentUserIdOk returns a tuple with the ViewerParentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerParentUserId

`func (o *ControllerUserAppVO) SetViewerParentUserId(v string)`

SetViewerParentUserId sets ViewerParentUserId field to given value.

### HasViewerParentUserId

`func (o *ControllerUserAppVO) HasViewerParentUserId() bool`

HasViewerParentUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


