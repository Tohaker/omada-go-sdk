# RadiusAuthSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **int32** | Authtype should be a value as follows: 0：PAP; 1:CHAP. | 
**DefaultGroup** | **string** | Only for SSL VPN. Default user group ID for the user on the LDAP/Radius server. User group can be created using &#39;Create SSL VPN user group&#39; interface, and user group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | 
**NasIp** | Pointer to **string** | NAS IP | [optional] 
**OverTime** | **int32** | Request Timeout, 1~60(s). | 
**ProfileId** | **string** | Profile ID. LDAP profile can be created using &#39;Create a new LDAP profile&#39; interface, and LDAP profile ID can be obtained from &#39;Get LDAP profile list&#39; interface. Radius profile can be created using &#39;Create a new Radius profile&#39; interface, and Radius profile ID can be obtained from &#39;Get Radius profile list&#39; interface. | 
**RepeatTime** | **int32** | Repeat time should be within the range of 1–10. | 

## Methods

### NewRadiusAuthSettingOpenApiVO

`func NewRadiusAuthSettingOpenApiVO(authType int32, defaultGroup string, overTime int32, profileId string, repeatTime int32, ) *RadiusAuthSettingOpenApiVO`

NewRadiusAuthSettingOpenApiVO instantiates a new RadiusAuthSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusAuthSettingOpenApiVOWithDefaults

`func NewRadiusAuthSettingOpenApiVOWithDefaults() *RadiusAuthSettingOpenApiVO`

NewRadiusAuthSettingOpenApiVOWithDefaults instantiates a new RadiusAuthSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *RadiusAuthSettingOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *RadiusAuthSettingOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *RadiusAuthSettingOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.


### GetDefaultGroup

`func (o *RadiusAuthSettingOpenApiVO) GetDefaultGroup() string`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *RadiusAuthSettingOpenApiVO) GetDefaultGroupOk() (*string, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *RadiusAuthSettingOpenApiVO) SetDefaultGroup(v string)`

SetDefaultGroup sets DefaultGroup field to given value.


### GetNasIp

`func (o *RadiusAuthSettingOpenApiVO) GetNasIp() string`

GetNasIp returns the NasIp field if non-nil, zero value otherwise.

### GetNasIpOk

`func (o *RadiusAuthSettingOpenApiVO) GetNasIpOk() (*string, bool)`

GetNasIpOk returns a tuple with the NasIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasIp

`func (o *RadiusAuthSettingOpenApiVO) SetNasIp(v string)`

SetNasIp sets NasIp field to given value.

### HasNasIp

`func (o *RadiusAuthSettingOpenApiVO) HasNasIp() bool`

HasNasIp returns a boolean if a field has been set.

### GetOverTime

`func (o *RadiusAuthSettingOpenApiVO) GetOverTime() int32`

GetOverTime returns the OverTime field if non-nil, zero value otherwise.

### GetOverTimeOk

`func (o *RadiusAuthSettingOpenApiVO) GetOverTimeOk() (*int32, bool)`

GetOverTimeOk returns a tuple with the OverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTime

`func (o *RadiusAuthSettingOpenApiVO) SetOverTime(v int32)`

SetOverTime sets OverTime field to given value.


### GetProfileId

`func (o *RadiusAuthSettingOpenApiVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *RadiusAuthSettingOpenApiVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *RadiusAuthSettingOpenApiVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetRepeatTime

`func (o *RadiusAuthSettingOpenApiVO) GetRepeatTime() int32`

GetRepeatTime returns the RepeatTime field if non-nil, zero value otherwise.

### GetRepeatTimeOk

`func (o *RadiusAuthSettingOpenApiVO) GetRepeatTimeOk() (*int32, bool)`

GetRepeatTimeOk returns a tuple with the RepeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatTime

`func (o *RadiusAuthSettingOpenApiVO) SetRepeatTime(v int32)`

SetRepeatTime sets RepeatTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


