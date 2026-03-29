# UserDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | Whether this user want to receive alert emails | [optional] 
**AllSite** | Pointer to **bool** | Whether user has all site permission, including new created site | [optional] 
**Email** | Pointer to **string** | User email | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Favorites** | Pointer to **[]string** | User favorite site list | [optional] 
**Id** | Pointer to **string** | User ID | [optional] 
**IncidentNotification** | Pointer to **bool** | Incident notification | [optional] 
**Name** | Pointer to **string** | User name | [optional] 
**OmadacId** | Pointer to **string** | Omada ID | [optional] 
**ParentUserId** | Pointer to **string** | User&#39;s parent user id | [optional] 
**RoleId** | Pointer to **string** | User role ID | [optional] 
**RoleName** | Pointer to **string** | User bind role name | [optional] 
**ShowTree** | Pointer to **bool** | Whether this user has Sub-users. | [optional] 
**SiteIds** | Pointer to **[]string** | User site privilege list | [optional] 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 
**TemporaryValidity** | Pointer to **int32** | Whether the temporary user is still valid | [optional] 
**Type** | Pointer to **int32** | Type of user, type should be a value as follows: 0:local user; 1: cloud user | [optional] 
**UserLevel** | Pointer to **int32** | User level, user level should be a value as follows: 0:standard user; 1:customer user; 2:msp user | [optional] 
**Verified** | Pointer to **bool** | Whether this cloud user has verified | [optional] 

## Methods

### NewUserDetailVO

`func NewUserDetailVO() *UserDetailVO`

NewUserDetailVO instantiates a new UserDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailVOWithDefaults

`func NewUserDetailVOWithDefaults() *UserDetailVO`

NewUserDetailVOWithDefaults instantiates a new UserDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *UserDetailVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *UserDetailVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *UserDetailVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *UserDetailVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAllSite

`func (o *UserDetailVO) GetAllSite() bool`

GetAllSite returns the AllSite field if non-nil, zero value otherwise.

### GetAllSiteOk

`func (o *UserDetailVO) GetAllSiteOk() (*bool, bool)`

GetAllSiteOk returns a tuple with the AllSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSite

`func (o *UserDetailVO) SetAllSite(v bool)`

SetAllSite sets AllSite field to given value.

### HasAllSite

`func (o *UserDetailVO) HasAllSite() bool`

HasAllSite returns a boolean if a field has been set.

### GetEmail

`func (o *UserDetailVO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDetailVO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDetailVO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserDetailVO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *UserDetailVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UserDetailVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UserDetailVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UserDetailVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFavorites

`func (o *UserDetailVO) GetFavorites() []string`

GetFavorites returns the Favorites field if non-nil, zero value otherwise.

### GetFavoritesOk

`func (o *UserDetailVO) GetFavoritesOk() (*[]string, bool)`

GetFavoritesOk returns a tuple with the Favorites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorites

`func (o *UserDetailVO) SetFavorites(v []string)`

SetFavorites sets Favorites field to given value.

### HasFavorites

`func (o *UserDetailVO) HasFavorites() bool`

HasFavorites returns a boolean if a field has been set.

### GetId

`func (o *UserDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncidentNotification

`func (o *UserDetailVO) GetIncidentNotification() bool`

GetIncidentNotification returns the IncidentNotification field if non-nil, zero value otherwise.

### GetIncidentNotificationOk

`func (o *UserDetailVO) GetIncidentNotificationOk() (*bool, bool)`

GetIncidentNotificationOk returns a tuple with the IncidentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentNotification

`func (o *UserDetailVO) SetIncidentNotification(v bool)`

SetIncidentNotification sets IncidentNotification field to given value.

### HasIncidentNotification

`func (o *UserDetailVO) HasIncidentNotification() bool`

HasIncidentNotification returns a boolean if a field has been set.

### GetName

`func (o *UserDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *UserDetailVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *UserDetailVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *UserDetailVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *UserDetailVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetParentUserId

`func (o *UserDetailVO) GetParentUserId() string`

GetParentUserId returns the ParentUserId field if non-nil, zero value otherwise.

### GetParentUserIdOk

`func (o *UserDetailVO) GetParentUserIdOk() (*string, bool)`

GetParentUserIdOk returns a tuple with the ParentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUserId

`func (o *UserDetailVO) SetParentUserId(v string)`

SetParentUserId sets ParentUserId field to given value.

### HasParentUserId

`func (o *UserDetailVO) HasParentUserId() bool`

HasParentUserId returns a boolean if a field has been set.

### GetRoleId

`func (o *UserDetailVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UserDetailVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UserDetailVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UserDetailVO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *UserDetailVO) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *UserDetailVO) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *UserDetailVO) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *UserDetailVO) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetShowTree

`func (o *UserDetailVO) GetShowTree() bool`

GetShowTree returns the ShowTree field if non-nil, zero value otherwise.

### GetShowTreeOk

`func (o *UserDetailVO) GetShowTreeOk() (*bool, bool)`

GetShowTreeOk returns a tuple with the ShowTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTree

`func (o *UserDetailVO) SetShowTree(v bool)`

SetShowTree sets ShowTree field to given value.

### HasShowTree

`func (o *UserDetailVO) HasShowTree() bool`

HasShowTree returns a boolean if a field has been set.

### GetSiteIds

`func (o *UserDetailVO) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *UserDetailVO) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *UserDetailVO) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *UserDetailVO) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### GetStartTime

`func (o *UserDetailVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UserDetailVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UserDetailVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UserDetailVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *UserDetailVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *UserDetailVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *UserDetailVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *UserDetailVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetTemporaryValidity

`func (o *UserDetailVO) GetTemporaryValidity() int32`

GetTemporaryValidity returns the TemporaryValidity field if non-nil, zero value otherwise.

### GetTemporaryValidityOk

`func (o *UserDetailVO) GetTemporaryValidityOk() (*int32, bool)`

GetTemporaryValidityOk returns a tuple with the TemporaryValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryValidity

`func (o *UserDetailVO) SetTemporaryValidity(v int32)`

SetTemporaryValidity sets TemporaryValidity field to given value.

### HasTemporaryValidity

`func (o *UserDetailVO) HasTemporaryValidity() bool`

HasTemporaryValidity returns a boolean if a field has been set.

### GetType

`func (o *UserDetailVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDetailVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDetailVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UserDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserLevel

`func (o *UserDetailVO) GetUserLevel() int32`

GetUserLevel returns the UserLevel field if non-nil, zero value otherwise.

### GetUserLevelOk

`func (o *UserDetailVO) GetUserLevelOk() (*int32, bool)`

GetUserLevelOk returns a tuple with the UserLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLevel

`func (o *UserDetailVO) SetUserLevel(v int32)`

SetUserLevel sets UserLevel field to given value.

### HasUserLevel

`func (o *UserDetailVO) HasUserLevel() bool`

HasUserLevel returns a boolean if a field has been set.

### GetVerified

`func (o *UserDetailVO) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserDetailVO) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserDetailVO) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserDetailVO) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


