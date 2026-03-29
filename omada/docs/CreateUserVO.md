# CreateUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | Whether this user want to receive alert emails. Make sure your email is not null | [optional] 
**AllSite** | **bool** | Whether user has all site permission, including new created site. | 
**Email** | Pointer to **string** | Email of user. | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**IncidentNotification** | Pointer to **bool** | Incident notification | [optional] 
**Name** | **string** | User name should contain 1 to 128 characters and start with letters, numbers, and underscores. When creating cloud user, you should set TP-LINK ID. | 
**Password** | Pointer to **string** | Password of local user should contain 8 to 128 characters. And password must be a combination of uppercase letters, lowercase letters, numbers, and special symbols. Symbols such as ! # $ % &amp; * @ ^ are supported. | [optional] 
**RoleId** | **string** | This field represents Role ID. Role can be created using &#39;Create new role&#39; interface, and Role ID can be obtained from &#39;Get role list&#39; interface. | 
**Sites** | Pointer to **[]string** | User site privilege list | [optional] 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 
**Type** | **int32** | Type of user should be a value as follows: 0: local user; 1: cloud user. | 

## Methods

### NewCreateUserVO

`func NewCreateUserVO(allSite bool, name string, roleId string, type_ int32, ) *CreateUserVO`

NewCreateUserVO instantiates a new CreateUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserVOWithDefaults

`func NewCreateUserVOWithDefaults() *CreateUserVO`

NewCreateUserVOWithDefaults instantiates a new CreateUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *CreateUserVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *CreateUserVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *CreateUserVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *CreateUserVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAllSite

`func (o *CreateUserVO) GetAllSite() bool`

GetAllSite returns the AllSite field if non-nil, zero value otherwise.

### GetAllSiteOk

`func (o *CreateUserVO) GetAllSiteOk() (*bool, bool)`

GetAllSiteOk returns a tuple with the AllSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSite

`func (o *CreateUserVO) SetAllSite(v bool)`

SetAllSite sets AllSite field to given value.


### GetEmail

`func (o *CreateUserVO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserVO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserVO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserVO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *CreateUserVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateUserVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateUserVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CreateUserVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIncidentNotification

`func (o *CreateUserVO) GetIncidentNotification() bool`

GetIncidentNotification returns the IncidentNotification field if non-nil, zero value otherwise.

### GetIncidentNotificationOk

`func (o *CreateUserVO) GetIncidentNotificationOk() (*bool, bool)`

GetIncidentNotificationOk returns a tuple with the IncidentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentNotification

`func (o *CreateUserVO) SetIncidentNotification(v bool)`

SetIncidentNotification sets IncidentNotification field to given value.

### HasIncidentNotification

`func (o *CreateUserVO) HasIncidentNotification() bool`

HasIncidentNotification returns a boolean if a field has been set.

### GetName

`func (o *CreateUserVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserVO) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateUserVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoleId

`func (o *CreateUserVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateUserVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateUserVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetSites

`func (o *CreateUserVO) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *CreateUserVO) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *CreateUserVO) SetSites(v []string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *CreateUserVO) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetStartTime

`func (o *CreateUserVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateUserVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateUserVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CreateUserVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *CreateUserVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *CreateUserVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *CreateUserVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *CreateUserVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetType

`func (o *CreateUserVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUserVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUserVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


