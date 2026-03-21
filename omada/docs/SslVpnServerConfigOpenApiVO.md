# SslVpnServerConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LDAPSetting** | Pointer to [**LdapSettingOpenApiVO**](LdapSettingOpenApiVO.md) |  | [optional] 
**AuthType** | Pointer to **int32** | Authentication type of the SSL VPN server should be a value as follows: 0: local; 1: radius; 2: LDAP. It is required when parameter [status] is true. | [optional] 
**ExitAtIdle** | Pointer to **bool** | Whether to exit when idle.  It is required when parameter [status] is true. | [optional] 
**ExitTime** | Pointer to **int32** | Exit time should be within the range of 5–3600(s). It is required when parameter [exitAtIdle] is true. | [optional] 
**IpLockSetting** | Pointer to [**LockSettingOpenApiVO**](LockSettingOpenApiVO.md) |  | [optional] 
**IpPoolEnd** | Pointer to **string** | The end IP of the IP pool. It is required when parameter [status] is true. | [optional] 
**IpPoolStart** | Pointer to **string** | The start IP of the IP pool. It is required when parameter [status] is true. | [optional] 
**NameLockSetting** | Pointer to [**LockSettingOpenApiVO**](LockSettingOpenApiVO.md) |  | [optional] 
**PrimaryDns** | Pointer to **string** | Primary DNS Server of the SSL VPN server. It is required when parameter [status] is true. | [optional] 
**RadiusSetting** | Pointer to [**RadiusSettingOpenApiVO**](RadiusSettingOpenApiVO.md) |  | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS Server of the SSL VPN server. | [optional] 
**ServicePort** | Pointer to **int32** | Service port of the SSL VPN server. Default value is 1194. It is required when parameter [status] is true, and it should be within the range of 1–65535 | [optional] 
**Status** | **bool** | Status of the SSL VPN server. | 
**TotalTraffic** | Pointer to **bool** | Whether to proxy all traffic. It is required when parameter [status] is true. | [optional] 
**WanPort** | Pointer to **string** | WAN port ID of the SSL VPN server. It is required when parameter [status] is true. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface | [optional] 

## Methods

### NewSslVpnServerConfigOpenApiVO

`func NewSslVpnServerConfigOpenApiVO(status bool, ) *SslVpnServerConfigOpenApiVO`

NewSslVpnServerConfigOpenApiVO instantiates a new SslVpnServerConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnServerConfigOpenApiVOWithDefaults

`func NewSslVpnServerConfigOpenApiVOWithDefaults() *SslVpnServerConfigOpenApiVO`

NewSslVpnServerConfigOpenApiVOWithDefaults instantiates a new SslVpnServerConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLDAPSetting

`func (o *SslVpnServerConfigOpenApiVO) GetLDAPSetting() LdapSettingOpenApiVO`

GetLDAPSetting returns the LDAPSetting field if non-nil, zero value otherwise.

### GetLDAPSettingOk

`func (o *SslVpnServerConfigOpenApiVO) GetLDAPSettingOk() (*LdapSettingOpenApiVO, bool)`

GetLDAPSettingOk returns a tuple with the LDAPSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPSetting

`func (o *SslVpnServerConfigOpenApiVO) SetLDAPSetting(v LdapSettingOpenApiVO)`

SetLDAPSetting sets LDAPSetting field to given value.

### HasLDAPSetting

`func (o *SslVpnServerConfigOpenApiVO) HasLDAPSetting() bool`

HasLDAPSetting returns a boolean if a field has been set.

### GetAuthType

`func (o *SslVpnServerConfigOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *SslVpnServerConfigOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *SslVpnServerConfigOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *SslVpnServerConfigOpenApiVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetExitAtIdle

`func (o *SslVpnServerConfigOpenApiVO) GetExitAtIdle() bool`

GetExitAtIdle returns the ExitAtIdle field if non-nil, zero value otherwise.

### GetExitAtIdleOk

`func (o *SslVpnServerConfigOpenApiVO) GetExitAtIdleOk() (*bool, bool)`

GetExitAtIdleOk returns a tuple with the ExitAtIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitAtIdle

`func (o *SslVpnServerConfigOpenApiVO) SetExitAtIdle(v bool)`

SetExitAtIdle sets ExitAtIdle field to given value.

### HasExitAtIdle

`func (o *SslVpnServerConfigOpenApiVO) HasExitAtIdle() bool`

HasExitAtIdle returns a boolean if a field has been set.

### GetExitTime

`func (o *SslVpnServerConfigOpenApiVO) GetExitTime() int32`

GetExitTime returns the ExitTime field if non-nil, zero value otherwise.

### GetExitTimeOk

`func (o *SslVpnServerConfigOpenApiVO) GetExitTimeOk() (*int32, bool)`

GetExitTimeOk returns a tuple with the ExitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitTime

`func (o *SslVpnServerConfigOpenApiVO) SetExitTime(v int32)`

SetExitTime sets ExitTime field to given value.

### HasExitTime

`func (o *SslVpnServerConfigOpenApiVO) HasExitTime() bool`

HasExitTime returns a boolean if a field has been set.

### GetIpLockSetting

`func (o *SslVpnServerConfigOpenApiVO) GetIpLockSetting() LockSettingOpenApiVO`

GetIpLockSetting returns the IpLockSetting field if non-nil, zero value otherwise.

### GetIpLockSettingOk

`func (o *SslVpnServerConfigOpenApiVO) GetIpLockSettingOk() (*LockSettingOpenApiVO, bool)`

GetIpLockSettingOk returns a tuple with the IpLockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLockSetting

`func (o *SslVpnServerConfigOpenApiVO) SetIpLockSetting(v LockSettingOpenApiVO)`

SetIpLockSetting sets IpLockSetting field to given value.

### HasIpLockSetting

`func (o *SslVpnServerConfigOpenApiVO) HasIpLockSetting() bool`

HasIpLockSetting returns a boolean if a field has been set.

### GetIpPoolEnd

`func (o *SslVpnServerConfigOpenApiVO) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *SslVpnServerConfigOpenApiVO) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *SslVpnServerConfigOpenApiVO) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *SslVpnServerConfigOpenApiVO) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *SslVpnServerConfigOpenApiVO) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *SslVpnServerConfigOpenApiVO) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *SslVpnServerConfigOpenApiVO) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *SslVpnServerConfigOpenApiVO) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetNameLockSetting

`func (o *SslVpnServerConfigOpenApiVO) GetNameLockSetting() LockSettingOpenApiVO`

GetNameLockSetting returns the NameLockSetting field if non-nil, zero value otherwise.

### GetNameLockSettingOk

`func (o *SslVpnServerConfigOpenApiVO) GetNameLockSettingOk() (*LockSettingOpenApiVO, bool)`

GetNameLockSettingOk returns a tuple with the NameLockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLockSetting

`func (o *SslVpnServerConfigOpenApiVO) SetNameLockSetting(v LockSettingOpenApiVO)`

SetNameLockSetting sets NameLockSetting field to given value.

### HasNameLockSetting

`func (o *SslVpnServerConfigOpenApiVO) HasNameLockSetting() bool`

HasNameLockSetting returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *SslVpnServerConfigOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *SslVpnServerConfigOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *SslVpnServerConfigOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *SslVpnServerConfigOpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetRadiusSetting

`func (o *SslVpnServerConfigOpenApiVO) GetRadiusSetting() RadiusSettingOpenApiVO`

GetRadiusSetting returns the RadiusSetting field if non-nil, zero value otherwise.

### GetRadiusSettingOk

`func (o *SslVpnServerConfigOpenApiVO) GetRadiusSettingOk() (*RadiusSettingOpenApiVO, bool)`

GetRadiusSettingOk returns a tuple with the RadiusSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusSetting

`func (o *SslVpnServerConfigOpenApiVO) SetRadiusSetting(v RadiusSettingOpenApiVO)`

SetRadiusSetting sets RadiusSetting field to given value.

### HasRadiusSetting

`func (o *SslVpnServerConfigOpenApiVO) HasRadiusSetting() bool`

HasRadiusSetting returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *SslVpnServerConfigOpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *SslVpnServerConfigOpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *SslVpnServerConfigOpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *SslVpnServerConfigOpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetServicePort

`func (o *SslVpnServerConfigOpenApiVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *SslVpnServerConfigOpenApiVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *SslVpnServerConfigOpenApiVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *SslVpnServerConfigOpenApiVO) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetStatus

`func (o *SslVpnServerConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SslVpnServerConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SslVpnServerConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTotalTraffic

`func (o *SslVpnServerConfigOpenApiVO) GetTotalTraffic() bool`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *SslVpnServerConfigOpenApiVO) GetTotalTrafficOk() (*bool, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *SslVpnServerConfigOpenApiVO) SetTotalTraffic(v bool)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *SslVpnServerConfigOpenApiVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetWanPort

`func (o *SslVpnServerConfigOpenApiVO) GetWanPort() string`

GetWanPort returns the WanPort field if non-nil, zero value otherwise.

### GetWanPortOk

`func (o *SslVpnServerConfigOpenApiVO) GetWanPortOk() (*string, bool)`

GetWanPortOk returns a tuple with the WanPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPort

`func (o *SslVpnServerConfigOpenApiVO) SetWanPort(v string)`

SetWanPort sets WanPort field to given value.

### HasWanPort

`func (o *SslVpnServerConfigOpenApiVO) HasWanPort() bool`

HasWanPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


