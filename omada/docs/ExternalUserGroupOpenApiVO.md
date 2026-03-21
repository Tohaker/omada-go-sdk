# ExternalUserGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSite** | **bool** | Whether having all site permissions. | 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Name** | **string** | External user group name should contain 1 to 128 characters. | 
**RoleId** | **string** | Role ID which can be obtained from &#39;Get role list&#39; interface. | 
**Sites** | Pointer to **[]string** | The site IDs that can be accessed. Effective when allSite is false. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 

## Methods

### NewExternalUserGroupOpenApiVO

`func NewExternalUserGroupOpenApiVO(allSite bool, name string, roleId string, ) *ExternalUserGroupOpenApiVO`

NewExternalUserGroupOpenApiVO instantiates a new ExternalUserGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalUserGroupOpenApiVOWithDefaults

`func NewExternalUserGroupOpenApiVOWithDefaults() *ExternalUserGroupOpenApiVO`

NewExternalUserGroupOpenApiVOWithDefaults instantiates a new ExternalUserGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSite

`func (o *ExternalUserGroupOpenApiVO) GetAllSite() bool`

GetAllSite returns the AllSite field if non-nil, zero value otherwise.

### GetAllSiteOk

`func (o *ExternalUserGroupOpenApiVO) GetAllSiteOk() (*bool, bool)`

GetAllSiteOk returns a tuple with the AllSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSite

`func (o *ExternalUserGroupOpenApiVO) SetAllSite(v bool)`

SetAllSite sets AllSite field to given value.


### GetEndTime

`func (o *ExternalUserGroupOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ExternalUserGroupOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ExternalUserGroupOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ExternalUserGroupOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetName

`func (o *ExternalUserGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalUserGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalUserGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRoleId

`func (o *ExternalUserGroupOpenApiVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ExternalUserGroupOpenApiVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ExternalUserGroupOpenApiVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetSites

`func (o *ExternalUserGroupOpenApiVO) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ExternalUserGroupOpenApiVO) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ExternalUserGroupOpenApiVO) SetSites(v []string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ExternalUserGroupOpenApiVO) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetStartTime

`func (o *ExternalUserGroupOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExternalUserGroupOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExternalUserGroupOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ExternalUserGroupOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *ExternalUserGroupOpenApiVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *ExternalUserGroupOpenApiVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *ExternalUserGroupOpenApiVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *ExternalUserGroupOpenApiVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


