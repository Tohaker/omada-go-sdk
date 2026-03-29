# SslVpnServerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LDAPSetting** | Pointer to [**LdapSettingOpenApiVO**](LdapSettingOpenApiVO.md) |  | [optional] 
**AuthType** | Pointer to **int32** | Authentication type of the SSL VPN server should be a value as follows: 0: local; 1: radius; 2: LDAP. | [optional] 
**ExistLdap** | Pointer to **bool** | Whether LDAP Authentication has been configured in SSL VPN Server. | [optional] 
**ExistRadius** | Pointer to **bool** | Whether Radius Authentication has been configured in SSL VPN Server. | [optional] 
**ExitAtIdle** | Pointer to **bool** | Whether to exit when idle | [optional] 
**ExitTime** | Pointer to **int32** | Exit time should be within the range of 5–3600(s). It is required when parameter [exitAtIdle] is true. | [optional] 
**Id** | Pointer to **string** | ID of the SSL VPN server. | [optional] 
**IpLockSetting** | Pointer to [**LockSettingOpenApiVO**](LockSettingOpenApiVO.md) |  | [optional] 
**IpPoolEnd** | Pointer to **string** | The end IP of the IP pool. | [optional] 
**IpPoolStart** | Pointer to **string** | The start IP of the IP pool. | [optional] 
**NameLockSetting** | Pointer to [**LockSettingOpenApiVO**](LockSettingOpenApiVO.md) |  | [optional] 
**PrimaryDns** | Pointer to **string** | Primary DNS Server of the SSL VPN server. | [optional] 
**RadiusSetting** | Pointer to [**RadiusSettingOpenApiVO**](RadiusSettingOpenApiVO.md) |  | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS Server of the SSL VPN server. | [optional] 
**ServicePort** | Pointer to **int32** | Service port of the SSL VPN server should be within the range of 1–65535 | [optional] 
**Status** | **bool** | Status of the SSL VPN server. | 
**SupportLDAP** | Pointer to **bool** | Whether the adopted gateway supports LDAP. | [optional] 
**SupportRadius** | Pointer to **bool** | Whether the adopted gateway supports Radius. | [optional] 
**TotalTraffic** | Pointer to **bool** | Whether to proxy all traffic. | [optional] 
**WanIp** | Pointer to **string** | WAP IP of the SSL VPN server. | [optional] 
**WanPort** | Pointer to **string** | WAN port of the SSL VPN server. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface | [optional] 

## Methods

### NewSslVpnServerSetting

`func NewSslVpnServerSetting(status bool, ) *SslVpnServerSetting`

NewSslVpnServerSetting instantiates a new SslVpnServerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnServerSettingWithDefaults

`func NewSslVpnServerSettingWithDefaults() *SslVpnServerSetting`

NewSslVpnServerSettingWithDefaults instantiates a new SslVpnServerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLDAPSetting

`func (o *SslVpnServerSetting) GetLDAPSetting() LdapSettingOpenApiVO`

GetLDAPSetting returns the LDAPSetting field if non-nil, zero value otherwise.

### GetLDAPSettingOk

`func (o *SslVpnServerSetting) GetLDAPSettingOk() (*LdapSettingOpenApiVO, bool)`

GetLDAPSettingOk returns a tuple with the LDAPSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPSetting

`func (o *SslVpnServerSetting) SetLDAPSetting(v LdapSettingOpenApiVO)`

SetLDAPSetting sets LDAPSetting field to given value.

### HasLDAPSetting

`func (o *SslVpnServerSetting) HasLDAPSetting() bool`

HasLDAPSetting returns a boolean if a field has been set.

### GetAuthType

`func (o *SslVpnServerSetting) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *SslVpnServerSetting) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *SslVpnServerSetting) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *SslVpnServerSetting) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetExistLdap

`func (o *SslVpnServerSetting) GetExistLdap() bool`

GetExistLdap returns the ExistLdap field if non-nil, zero value otherwise.

### GetExistLdapOk

`func (o *SslVpnServerSetting) GetExistLdapOk() (*bool, bool)`

GetExistLdapOk returns a tuple with the ExistLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistLdap

`func (o *SslVpnServerSetting) SetExistLdap(v bool)`

SetExistLdap sets ExistLdap field to given value.

### HasExistLdap

`func (o *SslVpnServerSetting) HasExistLdap() bool`

HasExistLdap returns a boolean if a field has been set.

### GetExistRadius

`func (o *SslVpnServerSetting) GetExistRadius() bool`

GetExistRadius returns the ExistRadius field if non-nil, zero value otherwise.

### GetExistRadiusOk

`func (o *SslVpnServerSetting) GetExistRadiusOk() (*bool, bool)`

GetExistRadiusOk returns a tuple with the ExistRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistRadius

`func (o *SslVpnServerSetting) SetExistRadius(v bool)`

SetExistRadius sets ExistRadius field to given value.

### HasExistRadius

`func (o *SslVpnServerSetting) HasExistRadius() bool`

HasExistRadius returns a boolean if a field has been set.

### GetExitAtIdle

`func (o *SslVpnServerSetting) GetExitAtIdle() bool`

GetExitAtIdle returns the ExitAtIdle field if non-nil, zero value otherwise.

### GetExitAtIdleOk

`func (o *SslVpnServerSetting) GetExitAtIdleOk() (*bool, bool)`

GetExitAtIdleOk returns a tuple with the ExitAtIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitAtIdle

`func (o *SslVpnServerSetting) SetExitAtIdle(v bool)`

SetExitAtIdle sets ExitAtIdle field to given value.

### HasExitAtIdle

`func (o *SslVpnServerSetting) HasExitAtIdle() bool`

HasExitAtIdle returns a boolean if a field has been set.

### GetExitTime

`func (o *SslVpnServerSetting) GetExitTime() int32`

GetExitTime returns the ExitTime field if non-nil, zero value otherwise.

### GetExitTimeOk

`func (o *SslVpnServerSetting) GetExitTimeOk() (*int32, bool)`

GetExitTimeOk returns a tuple with the ExitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitTime

`func (o *SslVpnServerSetting) SetExitTime(v int32)`

SetExitTime sets ExitTime field to given value.

### HasExitTime

`func (o *SslVpnServerSetting) HasExitTime() bool`

HasExitTime returns a boolean if a field has been set.

### GetId

`func (o *SslVpnServerSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SslVpnServerSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SslVpnServerSetting) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SslVpnServerSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpLockSetting

`func (o *SslVpnServerSetting) GetIpLockSetting() LockSettingOpenApiVO`

GetIpLockSetting returns the IpLockSetting field if non-nil, zero value otherwise.

### GetIpLockSettingOk

`func (o *SslVpnServerSetting) GetIpLockSettingOk() (*LockSettingOpenApiVO, bool)`

GetIpLockSettingOk returns a tuple with the IpLockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLockSetting

`func (o *SslVpnServerSetting) SetIpLockSetting(v LockSettingOpenApiVO)`

SetIpLockSetting sets IpLockSetting field to given value.

### HasIpLockSetting

`func (o *SslVpnServerSetting) HasIpLockSetting() bool`

HasIpLockSetting returns a boolean if a field has been set.

### GetIpPoolEnd

`func (o *SslVpnServerSetting) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *SslVpnServerSetting) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *SslVpnServerSetting) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *SslVpnServerSetting) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *SslVpnServerSetting) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *SslVpnServerSetting) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *SslVpnServerSetting) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *SslVpnServerSetting) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetNameLockSetting

`func (o *SslVpnServerSetting) GetNameLockSetting() LockSettingOpenApiVO`

GetNameLockSetting returns the NameLockSetting field if non-nil, zero value otherwise.

### GetNameLockSettingOk

`func (o *SslVpnServerSetting) GetNameLockSettingOk() (*LockSettingOpenApiVO, bool)`

GetNameLockSettingOk returns a tuple with the NameLockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLockSetting

`func (o *SslVpnServerSetting) SetNameLockSetting(v LockSettingOpenApiVO)`

SetNameLockSetting sets NameLockSetting field to given value.

### HasNameLockSetting

`func (o *SslVpnServerSetting) HasNameLockSetting() bool`

HasNameLockSetting returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *SslVpnServerSetting) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *SslVpnServerSetting) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *SslVpnServerSetting) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *SslVpnServerSetting) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetRadiusSetting

`func (o *SslVpnServerSetting) GetRadiusSetting() RadiusSettingOpenApiVO`

GetRadiusSetting returns the RadiusSetting field if non-nil, zero value otherwise.

### GetRadiusSettingOk

`func (o *SslVpnServerSetting) GetRadiusSettingOk() (*RadiusSettingOpenApiVO, bool)`

GetRadiusSettingOk returns a tuple with the RadiusSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusSetting

`func (o *SslVpnServerSetting) SetRadiusSetting(v RadiusSettingOpenApiVO)`

SetRadiusSetting sets RadiusSetting field to given value.

### HasRadiusSetting

`func (o *SslVpnServerSetting) HasRadiusSetting() bool`

HasRadiusSetting returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *SslVpnServerSetting) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *SslVpnServerSetting) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *SslVpnServerSetting) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *SslVpnServerSetting) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetServicePort

`func (o *SslVpnServerSetting) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *SslVpnServerSetting) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *SslVpnServerSetting) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *SslVpnServerSetting) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetStatus

`func (o *SslVpnServerSetting) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SslVpnServerSetting) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SslVpnServerSetting) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetSupportLDAP

`func (o *SslVpnServerSetting) GetSupportLDAP() bool`

GetSupportLDAP returns the SupportLDAP field if non-nil, zero value otherwise.

### GetSupportLDAPOk

`func (o *SslVpnServerSetting) GetSupportLDAPOk() (*bool, bool)`

GetSupportLDAPOk returns a tuple with the SupportLDAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLDAP

`func (o *SslVpnServerSetting) SetSupportLDAP(v bool)`

SetSupportLDAP sets SupportLDAP field to given value.

### HasSupportLDAP

`func (o *SslVpnServerSetting) HasSupportLDAP() bool`

HasSupportLDAP returns a boolean if a field has been set.

### GetSupportRadius

`func (o *SslVpnServerSetting) GetSupportRadius() bool`

GetSupportRadius returns the SupportRadius field if non-nil, zero value otherwise.

### GetSupportRadiusOk

`func (o *SslVpnServerSetting) GetSupportRadiusOk() (*bool, bool)`

GetSupportRadiusOk returns a tuple with the SupportRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRadius

`func (o *SslVpnServerSetting) SetSupportRadius(v bool)`

SetSupportRadius sets SupportRadius field to given value.

### HasSupportRadius

`func (o *SslVpnServerSetting) HasSupportRadius() bool`

HasSupportRadius returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *SslVpnServerSetting) GetTotalTraffic() bool`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *SslVpnServerSetting) GetTotalTrafficOk() (*bool, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *SslVpnServerSetting) SetTotalTraffic(v bool)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *SslVpnServerSetting) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetWanIp

`func (o *SslVpnServerSetting) GetWanIp() string`

GetWanIp returns the WanIp field if non-nil, zero value otherwise.

### GetWanIpOk

`func (o *SslVpnServerSetting) GetWanIpOk() (*string, bool)`

GetWanIpOk returns a tuple with the WanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIp

`func (o *SslVpnServerSetting) SetWanIp(v string)`

SetWanIp sets WanIp field to given value.

### HasWanIp

`func (o *SslVpnServerSetting) HasWanIp() bool`

HasWanIp returns a boolean if a field has been set.

### GetWanPort

`func (o *SslVpnServerSetting) GetWanPort() string`

GetWanPort returns the WanPort field if non-nil, zero value otherwise.

### GetWanPortOk

`func (o *SslVpnServerSetting) GetWanPortOk() (*string, bool)`

GetWanPortOk returns a tuple with the WanPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPort

`func (o *SslVpnServerSetting) SetWanPort(v string)`

SetWanPort sets WanPort field to given value.

### HasWanPort

`func (o *SslVpnServerSetting) HasWanPort() bool`

HasWanPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


