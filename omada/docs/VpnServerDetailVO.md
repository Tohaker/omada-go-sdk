# VpnServerDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LDAPSetting** | Pointer to [**LdapSettingOpenApiVO**](LdapSettingOpenApiVO.md) |  | [optional] 
**AccountAuth** | Pointer to **bool** | Account auth of the VPN, only for server OpenVPN. | [optional] 
**AdvancedSetting** | Pointer to [**VpnAdvancedSettingOpenApiVO**](VpnAdvancedSettingOpenApiVO.md) |  | [optional] 
**AuthMode** | Pointer to **int32** | Authentication mode. 0: Local; 1: LDAP. | [optional] 
**Clients** | Pointer to [**[]ServerWireGuardClientsVO**](ServerWireGuardClientsVO.md) | WireGuard clients. | [optional] 
**CustomNetwork** | Pointer to [**[]VpnIPSubnetsOpenApiVO**](VpnIPSubnetsOpenApiVO.md) | Custom networks of the VPN. | [optional] 
**CustomServer** | Pointer to **bool** | Whether enable custom server. | [optional] 
**CustomServerAddress** | Pointer to **string** | IPv4 Address or FQDN. | [optional] 
**Dns1** | Pointer to **string** | Primary DNS of the VPN. | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS of the VPN. | [optional] 
**DnsStatus** | Pointer to **bool** | Dns auto status. | [optional] 
**Encryption** | Pointer to **int32** | Encryption type. Only L2TP can configure auto type. 0: Encrypted; 1: Unencrypted; 2: Auto. | [optional] 
**ExitAtIdle** | Pointer to **bool** | Whether to exit when idle. | [optional] 
**ExitTime** | Pointer to **int32** | Exit time should be within the range of 5–3600(s). It is required when parameter [exitAtIdle] is true. | [optional] 
**FeatureDescription** | Pointer to [**[]FeatureInfoVO**](FeatureInfoVO.md) | Gateway Feature Description. | [optional] 
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**IpLockSetting** | Pointer to [**LockSettingOpenApiVO**](LockSettingOpenApiVO.md) |  | [optional] 
**IpPool** | Pointer to [**IPSubnetsVO**](IPSubnetsVO.md) |  | [optional] 
**IpPoolEnd** | Pointer to **string** | The end IP of the IP pool. | [optional] 
**IpPoolStart** | Pointer to **string** | The start IP of the IP pool. | [optional] 
**IpPoolType** | Pointer to **int32** | IP pool type should be a value as follows: 0: Ip Address/Mask; 1: Ip Address Range. | [optional] 
**KeepAlive** | Pointer to **int32** | The keepalive second of WireGuard peer should be within the range of 0-65535. | [optional] 
**LdapProfile** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **int32** | The MTU of WireGuard VPN should be within the range of 576-1440. | [optional] 
**Name** | Pointer to **string** | VPN name. | [optional] 
**NameLockSetting** | Pointer to [**LockSettingOpenApiVO**](LockSettingOpenApiVO.md) |  | [optional] 
**NetworkList** | Pointer to **[]string** | Network list of the VPN.  | [optional] 
**NetworkType** | Pointer to **int32** | Network type. 0: network list; 1: custom networks. | [optional] 
**PreSharedKey** | Pointer to **string** | Pre-shared key of the VPN. | [optional] 
**PrivateKey** | Pointer to **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PublicKey** | Pointer to **string** | The public key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**RadiusSetting** | Pointer to [**RadiusAuthSettingOpenApiVO**](RadiusAuthSettingOpenApiVO.md) |  | [optional] 
**RemoteIp** | Pointer to **string** | Remote IP of the VPN. | [optional] 
**ServicePort** | Pointer to **int32** | Service port for VPN server. | [optional] 
**ServiceType** | Pointer to **int32** | Service type of the VPN. | [optional] 
**Status** | Pointer to **bool** | Status of the VPN. | [optional] 
**TotalTraffic** | Pointer to **bool** | Whether to proxy all traffic. | [optional] 
**TunnelMode** | Pointer to **int32** | OpenVPN tunnel mode should be a value as follows: 0: split; 1: full. | [optional] 
**VpnType** | Pointer to **int32** | Server Vpn type. 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. | [optional] 
**VpnUserList** | Pointer to **[]string** | VPN user id list. | [optional] 
**Wans** | Pointer to **[]string** | WAN port ID. | [optional] 

## Methods

### NewVpnServerDetailVO

`func NewVpnServerDetailVO() *VpnServerDetailVO`

NewVpnServerDetailVO instantiates a new VpnServerDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnServerDetailVOWithDefaults

`func NewVpnServerDetailVOWithDefaults() *VpnServerDetailVO`

NewVpnServerDetailVOWithDefaults instantiates a new VpnServerDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLDAPSetting

`func (o *VpnServerDetailVO) GetLDAPSetting() LdapSettingOpenApiVO`

GetLDAPSetting returns the LDAPSetting field if non-nil, zero value otherwise.

### GetLDAPSettingOk

`func (o *VpnServerDetailVO) GetLDAPSettingOk() (*LdapSettingOpenApiVO, bool)`

GetLDAPSettingOk returns a tuple with the LDAPSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPSetting

`func (o *VpnServerDetailVO) SetLDAPSetting(v LdapSettingOpenApiVO)`

SetLDAPSetting sets LDAPSetting field to given value.

### HasLDAPSetting

`func (o *VpnServerDetailVO) HasLDAPSetting() bool`

HasLDAPSetting returns a boolean if a field has been set.

### GetAccountAuth

`func (o *VpnServerDetailVO) GetAccountAuth() bool`

GetAccountAuth returns the AccountAuth field if non-nil, zero value otherwise.

### GetAccountAuthOk

`func (o *VpnServerDetailVO) GetAccountAuthOk() (*bool, bool)`

GetAccountAuthOk returns a tuple with the AccountAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuth

`func (o *VpnServerDetailVO) SetAccountAuth(v bool)`

SetAccountAuth sets AccountAuth field to given value.

### HasAccountAuth

`func (o *VpnServerDetailVO) HasAccountAuth() bool`

HasAccountAuth returns a boolean if a field has been set.

### GetAdvancedSetting

`func (o *VpnServerDetailVO) GetAdvancedSetting() VpnAdvancedSettingOpenApiVO`

GetAdvancedSetting returns the AdvancedSetting field if non-nil, zero value otherwise.

### GetAdvancedSettingOk

`func (o *VpnServerDetailVO) GetAdvancedSettingOk() (*VpnAdvancedSettingOpenApiVO, bool)`

GetAdvancedSettingOk returns a tuple with the AdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSetting

`func (o *VpnServerDetailVO) SetAdvancedSetting(v VpnAdvancedSettingOpenApiVO)`

SetAdvancedSetting sets AdvancedSetting field to given value.

### HasAdvancedSetting

`func (o *VpnServerDetailVO) HasAdvancedSetting() bool`

HasAdvancedSetting returns a boolean if a field has been set.

### GetAuthMode

`func (o *VpnServerDetailVO) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *VpnServerDetailVO) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *VpnServerDetailVO) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *VpnServerDetailVO) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetClients

`func (o *VpnServerDetailVO) GetClients() []ServerWireGuardClientsVO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *VpnServerDetailVO) GetClientsOk() (*[]ServerWireGuardClientsVO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *VpnServerDetailVO) SetClients(v []ServerWireGuardClientsVO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *VpnServerDetailVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCustomNetwork

`func (o *VpnServerDetailVO) GetCustomNetwork() []VpnIPSubnetsOpenApiVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *VpnServerDetailVO) GetCustomNetworkOk() (*[]VpnIPSubnetsOpenApiVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *VpnServerDetailVO) SetCustomNetwork(v []VpnIPSubnetsOpenApiVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *VpnServerDetailVO) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetCustomServer

`func (o *VpnServerDetailVO) GetCustomServer() bool`

GetCustomServer returns the CustomServer field if non-nil, zero value otherwise.

### GetCustomServerOk

`func (o *VpnServerDetailVO) GetCustomServerOk() (*bool, bool)`

GetCustomServerOk returns a tuple with the CustomServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomServer

`func (o *VpnServerDetailVO) SetCustomServer(v bool)`

SetCustomServer sets CustomServer field to given value.

### HasCustomServer

`func (o *VpnServerDetailVO) HasCustomServer() bool`

HasCustomServer returns a boolean if a field has been set.

### GetCustomServerAddress

`func (o *VpnServerDetailVO) GetCustomServerAddress() string`

GetCustomServerAddress returns the CustomServerAddress field if non-nil, zero value otherwise.

### GetCustomServerAddressOk

`func (o *VpnServerDetailVO) GetCustomServerAddressOk() (*string, bool)`

GetCustomServerAddressOk returns a tuple with the CustomServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomServerAddress

`func (o *VpnServerDetailVO) SetCustomServerAddress(v string)`

SetCustomServerAddress sets CustomServerAddress field to given value.

### HasCustomServerAddress

`func (o *VpnServerDetailVO) HasCustomServerAddress() bool`

HasCustomServerAddress returns a boolean if a field has been set.

### GetDns1

`func (o *VpnServerDetailVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VpnServerDetailVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VpnServerDetailVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VpnServerDetailVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VpnServerDetailVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VpnServerDetailVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VpnServerDetailVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VpnServerDetailVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetDnsStatus

`func (o *VpnServerDetailVO) GetDnsStatus() bool`

GetDnsStatus returns the DnsStatus field if non-nil, zero value otherwise.

### GetDnsStatusOk

`func (o *VpnServerDetailVO) GetDnsStatusOk() (*bool, bool)`

GetDnsStatusOk returns a tuple with the DnsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsStatus

`func (o *VpnServerDetailVO) SetDnsStatus(v bool)`

SetDnsStatus sets DnsStatus field to given value.

### HasDnsStatus

`func (o *VpnServerDetailVO) HasDnsStatus() bool`

HasDnsStatus returns a boolean if a field has been set.

### GetEncryption

`func (o *VpnServerDetailVO) GetEncryption() int32`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VpnServerDetailVO) GetEncryptionOk() (*int32, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VpnServerDetailVO) SetEncryption(v int32)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VpnServerDetailVO) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetExitAtIdle

`func (o *VpnServerDetailVO) GetExitAtIdle() bool`

GetExitAtIdle returns the ExitAtIdle field if non-nil, zero value otherwise.

### GetExitAtIdleOk

`func (o *VpnServerDetailVO) GetExitAtIdleOk() (*bool, bool)`

GetExitAtIdleOk returns a tuple with the ExitAtIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitAtIdle

`func (o *VpnServerDetailVO) SetExitAtIdle(v bool)`

SetExitAtIdle sets ExitAtIdle field to given value.

### HasExitAtIdle

`func (o *VpnServerDetailVO) HasExitAtIdle() bool`

HasExitAtIdle returns a boolean if a field has been set.

### GetExitTime

`func (o *VpnServerDetailVO) GetExitTime() int32`

GetExitTime returns the ExitTime field if non-nil, zero value otherwise.

### GetExitTimeOk

`func (o *VpnServerDetailVO) GetExitTimeOk() (*int32, bool)`

GetExitTimeOk returns a tuple with the ExitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitTime

`func (o *VpnServerDetailVO) SetExitTime(v int32)`

SetExitTime sets ExitTime field to given value.

### HasExitTime

`func (o *VpnServerDetailVO) HasExitTime() bool`

HasExitTime returns a boolean if a field has been set.

### GetFeatureDescription

`func (o *VpnServerDetailVO) GetFeatureDescription() []FeatureInfoVO`

GetFeatureDescription returns the FeatureDescription field if non-nil, zero value otherwise.

### GetFeatureDescriptionOk

`func (o *VpnServerDetailVO) GetFeatureDescriptionOk() (*[]FeatureInfoVO, bool)`

GetFeatureDescriptionOk returns a tuple with the FeatureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureDescription

`func (o *VpnServerDetailVO) SetFeatureDescription(v []FeatureInfoVO)`

SetFeatureDescription sets FeatureDescription field to given value.

### HasFeatureDescription

`func (o *VpnServerDetailVO) HasFeatureDescription() bool`

HasFeatureDescription returns a boolean if a field has been set.

### GetId

`func (o *VpnServerDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnServerDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnServerDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnServerDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpLockSetting

`func (o *VpnServerDetailVO) GetIpLockSetting() LockSettingOpenApiVO`

GetIpLockSetting returns the IpLockSetting field if non-nil, zero value otherwise.

### GetIpLockSettingOk

`func (o *VpnServerDetailVO) GetIpLockSettingOk() (*LockSettingOpenApiVO, bool)`

GetIpLockSettingOk returns a tuple with the IpLockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLockSetting

`func (o *VpnServerDetailVO) SetIpLockSetting(v LockSettingOpenApiVO)`

SetIpLockSetting sets IpLockSetting field to given value.

### HasIpLockSetting

`func (o *VpnServerDetailVO) HasIpLockSetting() bool`

HasIpLockSetting returns a boolean if a field has been set.

### GetIpPool

`func (o *VpnServerDetailVO) GetIpPool() IPSubnetsVO`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *VpnServerDetailVO) GetIpPoolOk() (*IPSubnetsVO, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *VpnServerDetailVO) SetIpPool(v IPSubnetsVO)`

SetIpPool sets IpPool field to given value.

### HasIpPool

`func (o *VpnServerDetailVO) HasIpPool() bool`

HasIpPool returns a boolean if a field has been set.

### GetIpPoolEnd

`func (o *VpnServerDetailVO) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *VpnServerDetailVO) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *VpnServerDetailVO) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *VpnServerDetailVO) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *VpnServerDetailVO) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *VpnServerDetailVO) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *VpnServerDetailVO) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *VpnServerDetailVO) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetIpPoolType

`func (o *VpnServerDetailVO) GetIpPoolType() int32`

GetIpPoolType returns the IpPoolType field if non-nil, zero value otherwise.

### GetIpPoolTypeOk

`func (o *VpnServerDetailVO) GetIpPoolTypeOk() (*int32, bool)`

GetIpPoolTypeOk returns a tuple with the IpPoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolType

`func (o *VpnServerDetailVO) SetIpPoolType(v int32)`

SetIpPoolType sets IpPoolType field to given value.

### HasIpPoolType

`func (o *VpnServerDetailVO) HasIpPoolType() bool`

HasIpPoolType returns a boolean if a field has been set.

### GetKeepAlive

`func (o *VpnServerDetailVO) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *VpnServerDetailVO) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *VpnServerDetailVO) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *VpnServerDetailVO) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetLdapProfile

`func (o *VpnServerDetailVO) GetLdapProfile() string`

GetLdapProfile returns the LdapProfile field if non-nil, zero value otherwise.

### GetLdapProfileOk

`func (o *VpnServerDetailVO) GetLdapProfileOk() (*string, bool)`

GetLdapProfileOk returns a tuple with the LdapProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProfile

`func (o *VpnServerDetailVO) SetLdapProfile(v string)`

SetLdapProfile sets LdapProfile field to given value.

### HasLdapProfile

`func (o *VpnServerDetailVO) HasLdapProfile() bool`

HasLdapProfile returns a boolean if a field has been set.

### GetMtu

`func (o *VpnServerDetailVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VpnServerDetailVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VpnServerDetailVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VpnServerDetailVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *VpnServerDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnServerDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnServerDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnServerDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameLockSetting

`func (o *VpnServerDetailVO) GetNameLockSetting() LockSettingOpenApiVO`

GetNameLockSetting returns the NameLockSetting field if non-nil, zero value otherwise.

### GetNameLockSettingOk

`func (o *VpnServerDetailVO) GetNameLockSettingOk() (*LockSettingOpenApiVO, bool)`

GetNameLockSettingOk returns a tuple with the NameLockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLockSetting

`func (o *VpnServerDetailVO) SetNameLockSetting(v LockSettingOpenApiVO)`

SetNameLockSetting sets NameLockSetting field to given value.

### HasNameLockSetting

`func (o *VpnServerDetailVO) HasNameLockSetting() bool`

HasNameLockSetting returns a boolean if a field has been set.

### GetNetworkList

`func (o *VpnServerDetailVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VpnServerDetailVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VpnServerDetailVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VpnServerDetailVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *VpnServerDetailVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VpnServerDetailVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VpnServerDetailVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VpnServerDetailVO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *VpnServerDetailVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnServerDetailVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnServerDetailVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *VpnServerDetailVO) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *VpnServerDetailVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VpnServerDetailVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VpnServerDetailVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *VpnServerDetailVO) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *VpnServerDetailVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *VpnServerDetailVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *VpnServerDetailVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *VpnServerDetailVO) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRadiusSetting

`func (o *VpnServerDetailVO) GetRadiusSetting() RadiusAuthSettingOpenApiVO`

GetRadiusSetting returns the RadiusSetting field if non-nil, zero value otherwise.

### GetRadiusSettingOk

`func (o *VpnServerDetailVO) GetRadiusSettingOk() (*RadiusAuthSettingOpenApiVO, bool)`

GetRadiusSettingOk returns a tuple with the RadiusSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusSetting

`func (o *VpnServerDetailVO) SetRadiusSetting(v RadiusAuthSettingOpenApiVO)`

SetRadiusSetting sets RadiusSetting field to given value.

### HasRadiusSetting

`func (o *VpnServerDetailVO) HasRadiusSetting() bool`

HasRadiusSetting returns a boolean if a field has been set.

### GetRemoteIp

`func (o *VpnServerDetailVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnServerDetailVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnServerDetailVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *VpnServerDetailVO) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetServicePort

`func (o *VpnServerDetailVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VpnServerDetailVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VpnServerDetailVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *VpnServerDetailVO) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceType

`func (o *VpnServerDetailVO) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *VpnServerDetailVO) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *VpnServerDetailVO) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *VpnServerDetailVO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *VpnServerDetailVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnServerDetailVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnServerDetailVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnServerDetailVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *VpnServerDetailVO) GetTotalTraffic() bool`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *VpnServerDetailVO) GetTotalTrafficOk() (*bool, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *VpnServerDetailVO) SetTotalTraffic(v bool)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *VpnServerDetailVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetTunnelMode

`func (o *VpnServerDetailVO) GetTunnelMode() int32`

GetTunnelMode returns the TunnelMode field if non-nil, zero value otherwise.

### GetTunnelModeOk

`func (o *VpnServerDetailVO) GetTunnelModeOk() (*int32, bool)`

GetTunnelModeOk returns a tuple with the TunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelMode

`func (o *VpnServerDetailVO) SetTunnelMode(v int32)`

SetTunnelMode sets TunnelMode field to given value.

### HasTunnelMode

`func (o *VpnServerDetailVO) HasTunnelMode() bool`

HasTunnelMode returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnServerDetailVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnServerDetailVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnServerDetailVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnServerDetailVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetVpnUserList

`func (o *VpnServerDetailVO) GetVpnUserList() []string`

GetVpnUserList returns the VpnUserList field if non-nil, zero value otherwise.

### GetVpnUserListOk

`func (o *VpnServerDetailVO) GetVpnUserListOk() (*[]string, bool)`

GetVpnUserListOk returns a tuple with the VpnUserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnUserList

`func (o *VpnServerDetailVO) SetVpnUserList(v []string)`

SetVpnUserList sets VpnUserList field to given value.

### HasVpnUserList

`func (o *VpnServerDetailVO) HasVpnUserList() bool`

HasVpnUserList returns a boolean if a field has been set.

### GetWans

`func (o *VpnServerDetailVO) GetWans() []string`

GetWans returns the Wans field if non-nil, zero value otherwise.

### GetWansOk

`func (o *VpnServerDetailVO) GetWansOk() (*[]string, bool)`

GetWansOk returns a tuple with the Wans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWans

`func (o *VpnServerDetailVO) SetWans(v []string)`

SetWans sets Wans field to given value.

### HasWans

`func (o *VpnServerDetailVO) HasWans() bool`

HasWans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


