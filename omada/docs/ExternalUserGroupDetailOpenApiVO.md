# ExternalUserGroupDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSite** | Pointer to **bool** | Whether having all site permissions. | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Id** | Pointer to **string** | External user group ID. | [optional] 
**Name** | Pointer to **string** | External user group name. | [optional] 
**RoleId** | Pointer to **string** | Role ID. | [optional] 
**RoleName** | Pointer to **string** | Role name. | [optional] 
**RoleType** | Pointer to **int32** | Role type. It should be a value as follows: 0: standard, 1: customer, 2: msp. | [optional] 
**Sites** | Pointer to [**[]SiteInfoOpenApiVO**](SiteInfoOpenApiVO.md) | The sites which can be accessed. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permissione | [optional] 
**TemporaryValidity** | Pointer to **int32** | Whether the temporary user is still valid | [optional] 

## Methods

### NewExternalUserGroupDetailOpenApiVO

`func NewExternalUserGroupDetailOpenApiVO() *ExternalUserGroupDetailOpenApiVO`

NewExternalUserGroupDetailOpenApiVO instantiates a new ExternalUserGroupDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalUserGroupDetailOpenApiVOWithDefaults

`func NewExternalUserGroupDetailOpenApiVOWithDefaults() *ExternalUserGroupDetailOpenApiVO`

NewExternalUserGroupDetailOpenApiVOWithDefaults instantiates a new ExternalUserGroupDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSite

`func (o *ExternalUserGroupDetailOpenApiVO) GetAllSite() bool`

GetAllSite returns the AllSite field if non-nil, zero value otherwise.

### GetAllSiteOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetAllSiteOk() (*bool, bool)`

GetAllSiteOk returns a tuple with the AllSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSite

`func (o *ExternalUserGroupDetailOpenApiVO) SetAllSite(v bool)`

SetAllSite sets AllSite field to given value.

### HasAllSite

`func (o *ExternalUserGroupDetailOpenApiVO) HasAllSite() bool`

HasAllSite returns a boolean if a field has been set.

### GetEndTime

`func (o *ExternalUserGroupDetailOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ExternalUserGroupDetailOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ExternalUserGroupDetailOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *ExternalUserGroupDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalUserGroupDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalUserGroupDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExternalUserGroupDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalUserGroupDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalUserGroupDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleId

`func (o *ExternalUserGroupDetailOpenApiVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ExternalUserGroupDetailOpenApiVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ExternalUserGroupDetailOpenApiVO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *ExternalUserGroupDetailOpenApiVO) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ExternalUserGroupDetailOpenApiVO) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *ExternalUserGroupDetailOpenApiVO) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRoleType

`func (o *ExternalUserGroupDetailOpenApiVO) GetRoleType() int32`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetRoleTypeOk() (*int32, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *ExternalUserGroupDetailOpenApiVO) SetRoleType(v int32)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *ExternalUserGroupDetailOpenApiVO) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetSites

`func (o *ExternalUserGroupDetailOpenApiVO) GetSites() []SiteInfoOpenApiVO`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetSitesOk() (*[]SiteInfoOpenApiVO, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ExternalUserGroupDetailOpenApiVO) SetSites(v []SiteInfoOpenApiVO)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ExternalUserGroupDetailOpenApiVO) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetStartTime

`func (o *ExternalUserGroupDetailOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExternalUserGroupDetailOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ExternalUserGroupDetailOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *ExternalUserGroupDetailOpenApiVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *ExternalUserGroupDetailOpenApiVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *ExternalUserGroupDetailOpenApiVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetTemporaryValidity

`func (o *ExternalUserGroupDetailOpenApiVO) GetTemporaryValidity() int32`

GetTemporaryValidity returns the TemporaryValidity field if non-nil, zero value otherwise.

### GetTemporaryValidityOk

`func (o *ExternalUserGroupDetailOpenApiVO) GetTemporaryValidityOk() (*int32, bool)`

GetTemporaryValidityOk returns a tuple with the TemporaryValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryValidity

`func (o *ExternalUserGroupDetailOpenApiVO) SetTemporaryValidity(v int32)`

SetTemporaryValidity sets TemporaryValidity field to given value.

### HasTemporaryValidity

`func (o *ExternalUserGroupDetailOpenApiVO) HasTemporaryValidity() bool`

HasTemporaryValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


