# RadiusSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **int32** | Authtype should be a value as follows: 0：PAP; 1:CHAP | 
**DefaultGroup** | **string** | Default user group ID for the user on the radius server. User group can be created using &#39;Create SSL VPN user group&#39; interface, and user group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | 
**NasIp** | Pointer to **string** | NAS IP | [optional] 
**OverTime** | **int32** | Request Timeout, 1~60(s). | 
**RadiusProfile** | **string** | Radius profile ID. Radius profile can be created using &#39;Create a new Radius profile&#39; interface, and Radius profile ID can be obtained from &#39;Get Radius profile list&#39; interface. | 
**RepeatTime** | **int32** | Repeat time should be within the range of 1–10 | 

## Methods

### NewRadiusSettingOpenApiVO

`func NewRadiusSettingOpenApiVO(authType int32, defaultGroup string, overTime int32, radiusProfile string, repeatTime int32, ) *RadiusSettingOpenApiVO`

NewRadiusSettingOpenApiVO instantiates a new RadiusSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusSettingOpenApiVOWithDefaults

`func NewRadiusSettingOpenApiVOWithDefaults() *RadiusSettingOpenApiVO`

NewRadiusSettingOpenApiVOWithDefaults instantiates a new RadiusSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *RadiusSettingOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *RadiusSettingOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *RadiusSettingOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.


### GetDefaultGroup

`func (o *RadiusSettingOpenApiVO) GetDefaultGroup() string`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *RadiusSettingOpenApiVO) GetDefaultGroupOk() (*string, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *RadiusSettingOpenApiVO) SetDefaultGroup(v string)`

SetDefaultGroup sets DefaultGroup field to given value.


### GetNasIp

`func (o *RadiusSettingOpenApiVO) GetNasIp() string`

GetNasIp returns the NasIp field if non-nil, zero value otherwise.

### GetNasIpOk

`func (o *RadiusSettingOpenApiVO) GetNasIpOk() (*string, bool)`

GetNasIpOk returns a tuple with the NasIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasIp

`func (o *RadiusSettingOpenApiVO) SetNasIp(v string)`

SetNasIp sets NasIp field to given value.

### HasNasIp

`func (o *RadiusSettingOpenApiVO) HasNasIp() bool`

HasNasIp returns a boolean if a field has been set.

### GetOverTime

`func (o *RadiusSettingOpenApiVO) GetOverTime() int32`

GetOverTime returns the OverTime field if non-nil, zero value otherwise.

### GetOverTimeOk

`func (o *RadiusSettingOpenApiVO) GetOverTimeOk() (*int32, bool)`

GetOverTimeOk returns a tuple with the OverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTime

`func (o *RadiusSettingOpenApiVO) SetOverTime(v int32)`

SetOverTime sets OverTime field to given value.


### GetRadiusProfile

`func (o *RadiusSettingOpenApiVO) GetRadiusProfile() string`

GetRadiusProfile returns the RadiusProfile field if non-nil, zero value otherwise.

### GetRadiusProfileOk

`func (o *RadiusSettingOpenApiVO) GetRadiusProfileOk() (*string, bool)`

GetRadiusProfileOk returns a tuple with the RadiusProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfile

`func (o *RadiusSettingOpenApiVO) SetRadiusProfile(v string)`

SetRadiusProfile sets RadiusProfile field to given value.


### GetRepeatTime

`func (o *RadiusSettingOpenApiVO) GetRepeatTime() int32`

GetRepeatTime returns the RepeatTime field if non-nil, zero value otherwise.

### GetRepeatTimeOk

`func (o *RadiusSettingOpenApiVO) GetRepeatTimeOk() (*int32, bool)`

GetRepeatTimeOk returns a tuple with the RepeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatTime

`func (o *RadiusSettingOpenApiVO) SetRepeatTime(v int32)`

SetRepeatTime sets RepeatTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


