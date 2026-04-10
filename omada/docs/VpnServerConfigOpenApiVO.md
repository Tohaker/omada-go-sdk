# VpnServerConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAuth** | Pointer to **bool** | Whether enable account password of the VPN, only for server OpenVPN. | [optional] 
**AdvancedSetting** | [**VpnAdvancedSettingOpenApiVO**](VpnAdvancedSettingOpenApiVO.md) |  | 
**AuthMode** | Pointer to **int32** | Authentication mode should be a value as follows: 0: Local; 1: LDAP; 2:RADIUS. | [optional] 
**Clients** | Pointer to [**[]ServerWireGuardClientsConfigVO**](ServerWireGuardClientsConfigVO.md) | WireGuard clients. | [optional] 
**CustomNetwork** | Pointer to [**[]VpnIPSubnetsOpenApiVO**](VpnIPSubnetsOpenApiVO.md) | Custom networks of the VPN. | [optional] 
**CustomServer** | Pointer to **bool** | Enable custom server. | [optional] 
**CustomServerAddress** | Pointer to **string** | When parameter [customServer] is enable, parameter [customServerAddress] should be IPv4 address or FQDN. | [optional] 
**Dns1** | Pointer to **string** | Primary DNS of the VPN. | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS of the VPN. | [optional] 
**DnsStatus** | Pointer to **bool** | Dns auto status. | [optional] 
**Encryption** | Pointer to **int32** | Encryption should be a value as follows: 0: Encrypted; 1: Unencrypted; 2: Auto. | [optional] 
**ExitAtIdle** | **bool** | Whether to exit when idle.  It is required when parameter [status] is true. | 
**ExitTime** | Pointer to **int32** | Exit time should be within the range of 5–3600(s). It is required when parameter [exitAtIdle] is true. | [optional] 
**IpLockSetting** | [**LockSettingOpenApiVO**](LockSettingOpenApiVO.md) |  | 
**IpPool** | Pointer to [**IPSubnetsVO**](IPSubnetsVO.md) |  | [optional] 
**IpPoolEnd** | Pointer to **string** | The end IP of the IP pool. | [optional] 
**IpPoolStart** | Pointer to **string** | The start IP of the IP pool. | [optional] 
**IpPoolType** | Pointer to **int32** | IP pool type should be a value as follows: 0: Ip Address/Mask; 1: Ip Address Range. | [optional] 
**KeepAlive** | **int32** | The keepalive second of WireGuard peer should be within the range of 0-65535. | 
**LdapProfile** | Pointer to **string** | It is required when parameter [authType] is 1. | [optional] 
**LdapSetting** | Pointer to [**VpnBaseAuthSettingOpenApiVO**](VpnBaseAuthSettingOpenApiVO.md) |  | [optional] 
**Mtu** | **int32** | The MTU of WireGuard VPN should be within the range of 576-1440. | 
**Name** | **string** | Name should contain 1 to 63 characters. | 
**NameLockSetting** | [**LockSettingOpenApiVO**](LockSettingOpenApiVO.md) |  | 
**NetworkList** | Pointer to **[]string** | Network list of the VPN. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**NetworkType** | Pointer to **int32** | Network type should be a value as follows: 0: network list; 1: custom networks. | [optional] 
**PreSharedKey** | **string** | Pre-shared key of the VPN. | 
**PrivateKey** | **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | 
**RadiusSetting** | Pointer to [**RadiusAuthSettingOpenApiVO**](RadiusAuthSettingOpenApiVO.md) |  | [optional] 
**RemoteIp** | **string** | Remote IP of the VPN. | 
**ServicePort** | **int32** | Service port should be within the range of 1–65535. | 
**ServiceType** | **int32** | Service type of the Open VPN should be a value as follows: 0: UDP; 1: TCP. | 
**Status** | **bool** | Status of the VPN. | 
**TotalTraffic** | Pointer to **bool** | Whether to proxy all traffic. It is required when parameter [status] is true. | [optional] 
**TunnelMode** | Pointer to **int32** | Tunnel mode(only for server OpenVPN) should be a value as follows: 0: split; 1: full. | [optional] 
**VpnType** | **int32** | Vpn type should be a value as follows: 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. | 
**VpnUserList** | **[]string** | VPN user ID list. | 
**Wans** | **[]string** | WAN list of the VPN. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 

## Methods

### NewVpnServerConfigOpenApiVO

`func NewVpnServerConfigOpenApiVO(advancedSetting VpnAdvancedSettingOpenApiVO, exitAtIdle bool, ipLockSetting LockSettingOpenApiVO, keepAlive int32, mtu int32, name string, nameLockSetting LockSettingOpenApiVO, preSharedKey string, privateKey string, remoteIp string, servicePort int32, serviceType int32, status bool, vpnType int32, vpnUserList []string, wans []string, ) *VpnServerConfigOpenApiVO`

NewVpnServerConfigOpenApiVO instantiates a new VpnServerConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnServerConfigOpenApiVOWithDefaults

`func NewVpnServerConfigOpenApiVOWithDefaults() *VpnServerConfigOpenApiVO`

NewVpnServerConfigOpenApiVOWithDefaults instantiates a new VpnServerConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAuth

`func (o *VpnServerConfigOpenApiVO) GetAccountAuth() bool`

GetAccountAuth returns the AccountAuth field if non-nil, zero value otherwise.

### GetAccountAuthOk

`func (o *VpnServerConfigOpenApiVO) GetAccountAuthOk() (*bool, bool)`

GetAccountAuthOk returns a tuple with the AccountAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuth

`func (o *VpnServerConfigOpenApiVO) SetAccountAuth(v bool)`

SetAccountAuth sets AccountAuth field to given value.

### HasAccountAuth

`func (o *VpnServerConfigOpenApiVO) HasAccountAuth() bool`

HasAccountAuth returns a boolean if a field has been set.

### GetAdvancedSetting

`func (o *VpnServerConfigOpenApiVO) GetAdvancedSetting() VpnAdvancedSettingOpenApiVO`

GetAdvancedSetting returns the AdvancedSetting field if non-nil, zero value otherwise.

### GetAdvancedSettingOk

`func (o *VpnServerConfigOpenApiVO) GetAdvancedSettingOk() (*VpnAdvancedSettingOpenApiVO, bool)`

GetAdvancedSettingOk returns a tuple with the AdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSetting

`func (o *VpnServerConfigOpenApiVO) SetAdvancedSetting(v VpnAdvancedSettingOpenApiVO)`

SetAdvancedSetting sets AdvancedSetting field to given value.


### GetAuthMode

`func (o *VpnServerConfigOpenApiVO) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *VpnServerConfigOpenApiVO) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *VpnServerConfigOpenApiVO) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *VpnServerConfigOpenApiVO) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetClients

`func (o *VpnServerConfigOpenApiVO) GetClients() []ServerWireGuardClientsConfigVO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *VpnServerConfigOpenApiVO) GetClientsOk() (*[]ServerWireGuardClientsConfigVO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *VpnServerConfigOpenApiVO) SetClients(v []ServerWireGuardClientsConfigVO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *VpnServerConfigOpenApiVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCustomNetwork

`func (o *VpnServerConfigOpenApiVO) GetCustomNetwork() []VpnIPSubnetsOpenApiVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *VpnServerConfigOpenApiVO) GetCustomNetworkOk() (*[]VpnIPSubnetsOpenApiVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *VpnServerConfigOpenApiVO) SetCustomNetwork(v []VpnIPSubnetsOpenApiVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *VpnServerConfigOpenApiVO) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetCustomServer

`func (o *VpnServerConfigOpenApiVO) GetCustomServer() bool`

GetCustomServer returns the CustomServer field if non-nil, zero value otherwise.

### GetCustomServerOk

`func (o *VpnServerConfigOpenApiVO) GetCustomServerOk() (*bool, bool)`

GetCustomServerOk returns a tuple with the CustomServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomServer

`func (o *VpnServerConfigOpenApiVO) SetCustomServer(v bool)`

SetCustomServer sets CustomServer field to given value.

### HasCustomServer

`func (o *VpnServerConfigOpenApiVO) HasCustomServer() bool`

HasCustomServer returns a boolean if a field has been set.

### GetCustomServerAddress

`func (o *VpnServerConfigOpenApiVO) GetCustomServerAddress() string`

GetCustomServerAddress returns the CustomServerAddress field if non-nil, zero value otherwise.

### GetCustomServerAddressOk

`func (o *VpnServerConfigOpenApiVO) GetCustomServerAddressOk() (*string, bool)`

GetCustomServerAddressOk returns a tuple with the CustomServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomServerAddress

`func (o *VpnServerConfigOpenApiVO) SetCustomServerAddress(v string)`

SetCustomServerAddress sets CustomServerAddress field to given value.

### HasCustomServerAddress

`func (o *VpnServerConfigOpenApiVO) HasCustomServerAddress() bool`

HasCustomServerAddress returns a boolean if a field has been set.

### GetDns1

`func (o *VpnServerConfigOpenApiVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VpnServerConfigOpenApiVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VpnServerConfigOpenApiVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VpnServerConfigOpenApiVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VpnServerConfigOpenApiVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VpnServerConfigOpenApiVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VpnServerConfigOpenApiVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VpnServerConfigOpenApiVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetDnsStatus

`func (o *VpnServerConfigOpenApiVO) GetDnsStatus() bool`

GetDnsStatus returns the DnsStatus field if non-nil, zero value otherwise.

### GetDnsStatusOk

`func (o *VpnServerConfigOpenApiVO) GetDnsStatusOk() (*bool, bool)`

GetDnsStatusOk returns a tuple with the DnsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsStatus

`func (o *VpnServerConfigOpenApiVO) SetDnsStatus(v bool)`

SetDnsStatus sets DnsStatus field to given value.

### HasDnsStatus

`func (o *VpnServerConfigOpenApiVO) HasDnsStatus() bool`

HasDnsStatus returns a boolean if a field has been set.

### GetEncryption

`func (o *VpnServerConfigOpenApiVO) GetEncryption() int32`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VpnServerConfigOpenApiVO) GetEncryptionOk() (*int32, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VpnServerConfigOpenApiVO) SetEncryption(v int32)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VpnServerConfigOpenApiVO) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetExitAtIdle

`func (o *VpnServerConfigOpenApiVO) GetExitAtIdle() bool`

GetExitAtIdle returns the ExitAtIdle field if non-nil, zero value otherwise.

### GetExitAtIdleOk

`func (o *VpnServerConfigOpenApiVO) GetExitAtIdleOk() (*bool, bool)`

GetExitAtIdleOk returns a tuple with the ExitAtIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitAtIdle

`func (o *VpnServerConfigOpenApiVO) SetExitAtIdle(v bool)`

SetExitAtIdle sets ExitAtIdle field to given value.


### GetExitTime

`func (o *VpnServerConfigOpenApiVO) GetExitTime() int32`

GetExitTime returns the ExitTime field if non-nil, zero value otherwise.

### GetExitTimeOk

`func (o *VpnServerConfigOpenApiVO) GetExitTimeOk() (*int32, bool)`

GetExitTimeOk returns a tuple with the ExitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitTime

`func (o *VpnServerConfigOpenApiVO) SetExitTime(v int32)`

SetExitTime sets ExitTime field to given value.

### HasExitTime

`func (o *VpnServerConfigOpenApiVO) HasExitTime() bool`

HasExitTime returns a boolean if a field has been set.

### GetIpLockSetting

`func (o *VpnServerConfigOpenApiVO) GetIpLockSetting() LockSettingOpenApiVO`

GetIpLockSetting returns the IpLockSetting field if non-nil, zero value otherwise.

### GetIpLockSettingOk

`func (o *VpnServerConfigOpenApiVO) GetIpLockSettingOk() (*LockSettingOpenApiVO, bool)`

GetIpLockSettingOk returns a tuple with the IpLockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLockSetting

`func (o *VpnServerConfigOpenApiVO) SetIpLockSetting(v LockSettingOpenApiVO)`

SetIpLockSetting sets IpLockSetting field to given value.


### GetIpPool

`func (o *VpnServerConfigOpenApiVO) GetIpPool() IPSubnetsVO`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *VpnServerConfigOpenApiVO) GetIpPoolOk() (*IPSubnetsVO, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *VpnServerConfigOpenApiVO) SetIpPool(v IPSubnetsVO)`

SetIpPool sets IpPool field to given value.

### HasIpPool

`func (o *VpnServerConfigOpenApiVO) HasIpPool() bool`

HasIpPool returns a boolean if a field has been set.

### GetIpPoolEnd

`func (o *VpnServerConfigOpenApiVO) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *VpnServerConfigOpenApiVO) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *VpnServerConfigOpenApiVO) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *VpnServerConfigOpenApiVO) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *VpnServerConfigOpenApiVO) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *VpnServerConfigOpenApiVO) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *VpnServerConfigOpenApiVO) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *VpnServerConfigOpenApiVO) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetIpPoolType

`func (o *VpnServerConfigOpenApiVO) GetIpPoolType() int32`

GetIpPoolType returns the IpPoolType field if non-nil, zero value otherwise.

### GetIpPoolTypeOk

`func (o *VpnServerConfigOpenApiVO) GetIpPoolTypeOk() (*int32, bool)`

GetIpPoolTypeOk returns a tuple with the IpPoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolType

`func (o *VpnServerConfigOpenApiVO) SetIpPoolType(v int32)`

SetIpPoolType sets IpPoolType field to given value.

### HasIpPoolType

`func (o *VpnServerConfigOpenApiVO) HasIpPoolType() bool`

HasIpPoolType returns a boolean if a field has been set.

### GetKeepAlive

`func (o *VpnServerConfigOpenApiVO) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *VpnServerConfigOpenApiVO) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *VpnServerConfigOpenApiVO) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.


### GetLdapProfile

`func (o *VpnServerConfigOpenApiVO) GetLdapProfile() string`

GetLdapProfile returns the LdapProfile field if non-nil, zero value otherwise.

### GetLdapProfileOk

`func (o *VpnServerConfigOpenApiVO) GetLdapProfileOk() (*string, bool)`

GetLdapProfileOk returns a tuple with the LdapProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProfile

`func (o *VpnServerConfigOpenApiVO) SetLdapProfile(v string)`

SetLdapProfile sets LdapProfile field to given value.

### HasLdapProfile

`func (o *VpnServerConfigOpenApiVO) HasLdapProfile() bool`

HasLdapProfile returns a boolean if a field has been set.

### GetLdapSetting

`func (o *VpnServerConfigOpenApiVO) GetLdapSetting() VpnBaseAuthSettingOpenApiVO`

GetLdapSetting returns the LdapSetting field if non-nil, zero value otherwise.

### GetLdapSettingOk

`func (o *VpnServerConfigOpenApiVO) GetLdapSettingOk() (*VpnBaseAuthSettingOpenApiVO, bool)`

GetLdapSettingOk returns a tuple with the LdapSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSetting

`func (o *VpnServerConfigOpenApiVO) SetLdapSetting(v VpnBaseAuthSettingOpenApiVO)`

SetLdapSetting sets LdapSetting field to given value.

### HasLdapSetting

`func (o *VpnServerConfigOpenApiVO) HasLdapSetting() bool`

HasLdapSetting returns a boolean if a field has been set.

### GetMtu

`func (o *VpnServerConfigOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VpnServerConfigOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VpnServerConfigOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### GetName

`func (o *VpnServerConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnServerConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnServerConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNameLockSetting

`func (o *VpnServerConfigOpenApiVO) GetNameLockSetting() LockSettingOpenApiVO`

GetNameLockSetting returns the NameLockSetting field if non-nil, zero value otherwise.

### GetNameLockSettingOk

`func (o *VpnServerConfigOpenApiVO) GetNameLockSettingOk() (*LockSettingOpenApiVO, bool)`

GetNameLockSettingOk returns a tuple with the NameLockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLockSetting

`func (o *VpnServerConfigOpenApiVO) SetNameLockSetting(v LockSettingOpenApiVO)`

SetNameLockSetting sets NameLockSetting field to given value.


### GetNetworkList

`func (o *VpnServerConfigOpenApiVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VpnServerConfigOpenApiVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VpnServerConfigOpenApiVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VpnServerConfigOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *VpnServerConfigOpenApiVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VpnServerConfigOpenApiVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VpnServerConfigOpenApiVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VpnServerConfigOpenApiVO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *VpnServerConfigOpenApiVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnServerConfigOpenApiVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnServerConfigOpenApiVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.


### GetPrivateKey

`func (o *VpnServerConfigOpenApiVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VpnServerConfigOpenApiVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VpnServerConfigOpenApiVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetRadiusSetting

`func (o *VpnServerConfigOpenApiVO) GetRadiusSetting() RadiusAuthSettingOpenApiVO`

GetRadiusSetting returns the RadiusSetting field if non-nil, zero value otherwise.

### GetRadiusSettingOk

`func (o *VpnServerConfigOpenApiVO) GetRadiusSettingOk() (*RadiusAuthSettingOpenApiVO, bool)`

GetRadiusSettingOk returns a tuple with the RadiusSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusSetting

`func (o *VpnServerConfigOpenApiVO) SetRadiusSetting(v RadiusAuthSettingOpenApiVO)`

SetRadiusSetting sets RadiusSetting field to given value.

### HasRadiusSetting

`func (o *VpnServerConfigOpenApiVO) HasRadiusSetting() bool`

HasRadiusSetting returns a boolean if a field has been set.

### GetRemoteIp

`func (o *VpnServerConfigOpenApiVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnServerConfigOpenApiVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnServerConfigOpenApiVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetServicePort

`func (o *VpnServerConfigOpenApiVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VpnServerConfigOpenApiVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VpnServerConfigOpenApiVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.


### GetServiceType

`func (o *VpnServerConfigOpenApiVO) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *VpnServerConfigOpenApiVO) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *VpnServerConfigOpenApiVO) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.


### GetStatus

`func (o *VpnServerConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnServerConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnServerConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTotalTraffic

`func (o *VpnServerConfigOpenApiVO) GetTotalTraffic() bool`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *VpnServerConfigOpenApiVO) GetTotalTrafficOk() (*bool, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *VpnServerConfigOpenApiVO) SetTotalTraffic(v bool)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *VpnServerConfigOpenApiVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetTunnelMode

`func (o *VpnServerConfigOpenApiVO) GetTunnelMode() int32`

GetTunnelMode returns the TunnelMode field if non-nil, zero value otherwise.

### GetTunnelModeOk

`func (o *VpnServerConfigOpenApiVO) GetTunnelModeOk() (*int32, bool)`

GetTunnelModeOk returns a tuple with the TunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelMode

`func (o *VpnServerConfigOpenApiVO) SetTunnelMode(v int32)`

SetTunnelMode sets TunnelMode field to given value.

### HasTunnelMode

`func (o *VpnServerConfigOpenApiVO) HasTunnelMode() bool`

HasTunnelMode returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnServerConfigOpenApiVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnServerConfigOpenApiVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnServerConfigOpenApiVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.


### GetVpnUserList

`func (o *VpnServerConfigOpenApiVO) GetVpnUserList() []string`

GetVpnUserList returns the VpnUserList field if non-nil, zero value otherwise.

### GetVpnUserListOk

`func (o *VpnServerConfigOpenApiVO) GetVpnUserListOk() (*[]string, bool)`

GetVpnUserListOk returns a tuple with the VpnUserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnUserList

`func (o *VpnServerConfigOpenApiVO) SetVpnUserList(v []string)`

SetVpnUserList sets VpnUserList field to given value.


### GetWans

`func (o *VpnServerConfigOpenApiVO) GetWans() []string`

GetWans returns the Wans field if non-nil, zero value otherwise.

### GetWansOk

`func (o *VpnServerConfigOpenApiVO) GetWansOk() (*[]string, bool)`

GetWansOk returns a tuple with the Wans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWans

`func (o *VpnServerConfigOpenApiVO) SetWans(v []string)`

SetWans sets Wans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


