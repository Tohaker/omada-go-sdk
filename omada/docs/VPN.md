# VPN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAuth** | Pointer to **bool** | Account auth of the VPN, only for server OpenVPN. | [optional] 
**AuthMode** | Pointer to **int32** | Authentication mode should be a value as follows: 0: Local; 1: LDAP. | [optional] 
**ClientPassword** | Pointer to **string** | Client password of the VPN. | [optional] 
**ClientUserName** | Pointer to **string** | Client username of the VPN. | [optional] 
**ClientVpnType1** | **int32** | Client to Site VPN : 0:VPN Server 1:VPN Client. | 
**ClientVpnType2** | **int32** | Client Vpn type should be a value as follows: 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN. | 
**CustomNetwork** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Custom networks of the VPN, only for Manual IPSec type. | [optional] 
**Encryption** | Pointer to **int32** | Encryption should be a value as follows: 0: Encrypted; 1: Unencrypted. | [optional] 
**ExistAccountAuth** | Pointer to **bool** | Whether Account Auth is configured. | [optional] 
**ExistCustomDns** | Pointer to **bool** | Whether custom DNS Server has been configured in current VPN. | [optional] 
**ExistCustomNetwork** | Pointer to **bool** | Whether Local Network Type is Custom. | [optional] 
**ExistIPsec** | Pointer to **bool** | Whether current VPN is IPsec type. | [optional] 
**ExistIpRange** | Pointer to **bool** | Whether IP pool type is Ip Address Range. | [optional] 
**ExistL2TP** | Pointer to **bool** | Whether current VPN is L2TP type. | [optional] 
**ExistLdap** | Pointer to **bool** | Whether Auth Mode is LDAP. | [optional] 
**ExistOpenVpnDomain** | Pointer to **bool** | Whether Open VPN Client Remote Server is Domain Name. | [optional] 
**ExistPfs** | Pointer to **bool** | Whether current VPN has configured none-default option: dh14 or dh15. | [optional] 
**ExistPhase1Proposal1** | Pointer to **bool** | Whether current VPN has configured none-default Proposal option for IKE negotiation phase-1: SHA384 or SHA512. | [optional] 
**ExistPhase2Proposal2** | Pointer to **bool** | Whether current VPN has configured none-default Proposal option for IKE negotiation phase-2: SHA384 or SHA512. | [optional] 
**ExistServerOpenVpnGoogleLdap** | Pointer to **bool** | Whether Auth Mode is Google LDAP. | [optional] 
**ExistTunnelMode** | Pointer to **bool** | OpenVPN tunnel mode is Full. | [optional] 
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**IpPool** | [**IPSubnetsVO**](IPSubnetsVO.md) |  | 
**IpPoolEnd** | Pointer to **string** | The end IP of the IP pool. | [optional] 
**IpPoolStart** | Pointer to **string** | The start IP of the IP pool. | [optional] 
**IpPoolType** | Pointer to **int32** | IP pool type should be a value as follows: 0: Ip Address/Mask; 1: Ip Address Range. | [optional] 
**LdapProfile** | Pointer to **string** | Ldap Profile. | [optional] 
**Name** | **string** | Name should contain 1 to 63 characters. | 
**NetworkList** | Pointer to **[]string** | Network list of the VPN, only for Manual IPSec type. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**NetworkType** | Pointer to **int32** | Network type should be a value as follows: 0: network list; 1: custom networks. | [optional] 
**OpenVpnMode** | Pointer to **int32** | OpenVPN mode should be a value as follows: 0: certification; 1: certification+account. | [optional] 
**OpenVpnTunnelMode** | Pointer to **int32** | OpenVPN tunnel mode should be a value as follows: 0: split; 1: full. | [optional] 
**PreSharedKey** | Pointer to **string** | Pre-shared key of the VPN. | [optional] 
**PrimaryDns** | Pointer to **string** | Primary DNS of the VPN. | [optional] 
**Purpose** | **int32** | Purpose of the VPN. | 
**RemoteIp** | Pointer to **string** | Remote IP of the VPN. | [optional] 
**RemoteSite** | Pointer to **string** | Remote site of the VPN. | [optional] 
**RemoteSubnet** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Remote subnet of the VPN, only for Manual IPSec type. | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS of the VPN. | [optional] 
**ServicePort** | Pointer to **int32** | Service port should be within the range of 1–65535. | [optional] 
**ServiceType** | Pointer to **int32** | Service type of the VPN. | [optional] 
**SiteVpnType** | Pointer to **int32** | Site VPN type of the VPN. | [optional] 
**Status** | Pointer to **bool** | Status of the VPN. | [optional] 
**VpnConfiguration** | Pointer to [**VpnCertificateOpenApiVO**](VpnCertificateOpenApiVO.md) |  | [optional] 
**Wans** | Pointer to **[]string** | WAN list of the VPN. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | [optional] 
**WorkingMode** | Pointer to **int32** | Working Mode should be a value as follow: 0:NAT 1:Routing. | [optional] 

## Methods

### NewVPN

`func NewVPN(clientVpnType1 int32, clientVpnType2 int32, ipPool IPSubnetsVO, name string, purpose int32, ) *VPN`

NewVPN instantiates a new VPN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNWithDefaults

`func NewVPNWithDefaults() *VPN`

NewVPNWithDefaults instantiates a new VPN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAuth

`func (o *VPN) GetAccountAuth() bool`

GetAccountAuth returns the AccountAuth field if non-nil, zero value otherwise.

### GetAccountAuthOk

`func (o *VPN) GetAccountAuthOk() (*bool, bool)`

GetAccountAuthOk returns a tuple with the AccountAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuth

`func (o *VPN) SetAccountAuth(v bool)`

SetAccountAuth sets AccountAuth field to given value.

### HasAccountAuth

`func (o *VPN) HasAccountAuth() bool`

HasAccountAuth returns a boolean if a field has been set.

### GetAuthMode

`func (o *VPN) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *VPN) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *VPN) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *VPN) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetClientPassword

`func (o *VPN) GetClientPassword() string`

GetClientPassword returns the ClientPassword field if non-nil, zero value otherwise.

### GetClientPasswordOk

`func (o *VPN) GetClientPasswordOk() (*string, bool)`

GetClientPasswordOk returns a tuple with the ClientPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPassword

`func (o *VPN) SetClientPassword(v string)`

SetClientPassword sets ClientPassword field to given value.

### HasClientPassword

`func (o *VPN) HasClientPassword() bool`

HasClientPassword returns a boolean if a field has been set.

### GetClientUserName

`func (o *VPN) GetClientUserName() string`

GetClientUserName returns the ClientUserName field if non-nil, zero value otherwise.

### GetClientUserNameOk

`func (o *VPN) GetClientUserNameOk() (*string, bool)`

GetClientUserNameOk returns a tuple with the ClientUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUserName

`func (o *VPN) SetClientUserName(v string)`

SetClientUserName sets ClientUserName field to given value.

### HasClientUserName

`func (o *VPN) HasClientUserName() bool`

HasClientUserName returns a boolean if a field has been set.

### GetClientVpnType1

`func (o *VPN) GetClientVpnType1() int32`

GetClientVpnType1 returns the ClientVpnType1 field if non-nil, zero value otherwise.

### GetClientVpnType1Ok

`func (o *VPN) GetClientVpnType1Ok() (*int32, bool)`

GetClientVpnType1Ok returns a tuple with the ClientVpnType1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVpnType1

`func (o *VPN) SetClientVpnType1(v int32)`

SetClientVpnType1 sets ClientVpnType1 field to given value.


### GetClientVpnType2

`func (o *VPN) GetClientVpnType2() int32`

GetClientVpnType2 returns the ClientVpnType2 field if non-nil, zero value otherwise.

### GetClientVpnType2Ok

`func (o *VPN) GetClientVpnType2Ok() (*int32, bool)`

GetClientVpnType2Ok returns a tuple with the ClientVpnType2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVpnType2

`func (o *VPN) SetClientVpnType2(v int32)`

SetClientVpnType2 sets ClientVpnType2 field to given value.


### GetCustomNetwork

`func (o *VPN) GetCustomNetwork() []IPSubnetsVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *VPN) GetCustomNetworkOk() (*[]IPSubnetsVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *VPN) SetCustomNetwork(v []IPSubnetsVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *VPN) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetEncryption

`func (o *VPN) GetEncryption() int32`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VPN) GetEncryptionOk() (*int32, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VPN) SetEncryption(v int32)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VPN) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetExistAccountAuth

`func (o *VPN) GetExistAccountAuth() bool`

GetExistAccountAuth returns the ExistAccountAuth field if non-nil, zero value otherwise.

### GetExistAccountAuthOk

`func (o *VPN) GetExistAccountAuthOk() (*bool, bool)`

GetExistAccountAuthOk returns a tuple with the ExistAccountAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistAccountAuth

`func (o *VPN) SetExistAccountAuth(v bool)`

SetExistAccountAuth sets ExistAccountAuth field to given value.

### HasExistAccountAuth

`func (o *VPN) HasExistAccountAuth() bool`

HasExistAccountAuth returns a boolean if a field has been set.

### GetExistCustomDns

`func (o *VPN) GetExistCustomDns() bool`

GetExistCustomDns returns the ExistCustomDns field if non-nil, zero value otherwise.

### GetExistCustomDnsOk

`func (o *VPN) GetExistCustomDnsOk() (*bool, bool)`

GetExistCustomDnsOk returns a tuple with the ExistCustomDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCustomDns

`func (o *VPN) SetExistCustomDns(v bool)`

SetExistCustomDns sets ExistCustomDns field to given value.

### HasExistCustomDns

`func (o *VPN) HasExistCustomDns() bool`

HasExistCustomDns returns a boolean if a field has been set.

### GetExistCustomNetwork

`func (o *VPN) GetExistCustomNetwork() bool`

GetExistCustomNetwork returns the ExistCustomNetwork field if non-nil, zero value otherwise.

### GetExistCustomNetworkOk

`func (o *VPN) GetExistCustomNetworkOk() (*bool, bool)`

GetExistCustomNetworkOk returns a tuple with the ExistCustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCustomNetwork

`func (o *VPN) SetExistCustomNetwork(v bool)`

SetExistCustomNetwork sets ExistCustomNetwork field to given value.

### HasExistCustomNetwork

`func (o *VPN) HasExistCustomNetwork() bool`

HasExistCustomNetwork returns a boolean if a field has been set.

### GetExistIPsec

`func (o *VPN) GetExistIPsec() bool`

GetExistIPsec returns the ExistIPsec field if non-nil, zero value otherwise.

### GetExistIPsecOk

`func (o *VPN) GetExistIPsecOk() (*bool, bool)`

GetExistIPsecOk returns a tuple with the ExistIPsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistIPsec

`func (o *VPN) SetExistIPsec(v bool)`

SetExistIPsec sets ExistIPsec field to given value.

### HasExistIPsec

`func (o *VPN) HasExistIPsec() bool`

HasExistIPsec returns a boolean if a field has been set.

### GetExistIpRange

`func (o *VPN) GetExistIpRange() bool`

GetExistIpRange returns the ExistIpRange field if non-nil, zero value otherwise.

### GetExistIpRangeOk

`func (o *VPN) GetExistIpRangeOk() (*bool, bool)`

GetExistIpRangeOk returns a tuple with the ExistIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistIpRange

`func (o *VPN) SetExistIpRange(v bool)`

SetExistIpRange sets ExistIpRange field to given value.

### HasExistIpRange

`func (o *VPN) HasExistIpRange() bool`

HasExistIpRange returns a boolean if a field has been set.

### GetExistL2TP

`func (o *VPN) GetExistL2TP() bool`

GetExistL2TP returns the ExistL2TP field if non-nil, zero value otherwise.

### GetExistL2TPOk

`func (o *VPN) GetExistL2TPOk() (*bool, bool)`

GetExistL2TPOk returns a tuple with the ExistL2TP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistL2TP

`func (o *VPN) SetExistL2TP(v bool)`

SetExistL2TP sets ExistL2TP field to given value.

### HasExistL2TP

`func (o *VPN) HasExistL2TP() bool`

HasExistL2TP returns a boolean if a field has been set.

### GetExistLdap

`func (o *VPN) GetExistLdap() bool`

GetExistLdap returns the ExistLdap field if non-nil, zero value otherwise.

### GetExistLdapOk

`func (o *VPN) GetExistLdapOk() (*bool, bool)`

GetExistLdapOk returns a tuple with the ExistLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistLdap

`func (o *VPN) SetExistLdap(v bool)`

SetExistLdap sets ExistLdap field to given value.

### HasExistLdap

`func (o *VPN) HasExistLdap() bool`

HasExistLdap returns a boolean if a field has been set.

### GetExistOpenVpnDomain

`func (o *VPN) GetExistOpenVpnDomain() bool`

GetExistOpenVpnDomain returns the ExistOpenVpnDomain field if non-nil, zero value otherwise.

### GetExistOpenVpnDomainOk

`func (o *VPN) GetExistOpenVpnDomainOk() (*bool, bool)`

GetExistOpenVpnDomainOk returns a tuple with the ExistOpenVpnDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistOpenVpnDomain

`func (o *VPN) SetExistOpenVpnDomain(v bool)`

SetExistOpenVpnDomain sets ExistOpenVpnDomain field to given value.

### HasExistOpenVpnDomain

`func (o *VPN) HasExistOpenVpnDomain() bool`

HasExistOpenVpnDomain returns a boolean if a field has been set.

### GetExistPfs

`func (o *VPN) GetExistPfs() bool`

GetExistPfs returns the ExistPfs field if non-nil, zero value otherwise.

### GetExistPfsOk

`func (o *VPN) GetExistPfsOk() (*bool, bool)`

GetExistPfsOk returns a tuple with the ExistPfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistPfs

`func (o *VPN) SetExistPfs(v bool)`

SetExistPfs sets ExistPfs field to given value.

### HasExistPfs

`func (o *VPN) HasExistPfs() bool`

HasExistPfs returns a boolean if a field has been set.

### GetExistPhase1Proposal1

`func (o *VPN) GetExistPhase1Proposal1() bool`

GetExistPhase1Proposal1 returns the ExistPhase1Proposal1 field if non-nil, zero value otherwise.

### GetExistPhase1Proposal1Ok

`func (o *VPN) GetExistPhase1Proposal1Ok() (*bool, bool)`

GetExistPhase1Proposal1Ok returns a tuple with the ExistPhase1Proposal1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistPhase1Proposal1

`func (o *VPN) SetExistPhase1Proposal1(v bool)`

SetExistPhase1Proposal1 sets ExistPhase1Proposal1 field to given value.

### HasExistPhase1Proposal1

`func (o *VPN) HasExistPhase1Proposal1() bool`

HasExistPhase1Proposal1 returns a boolean if a field has been set.

### GetExistPhase2Proposal2

`func (o *VPN) GetExistPhase2Proposal2() bool`

GetExistPhase2Proposal2 returns the ExistPhase2Proposal2 field if non-nil, zero value otherwise.

### GetExistPhase2Proposal2Ok

`func (o *VPN) GetExistPhase2Proposal2Ok() (*bool, bool)`

GetExistPhase2Proposal2Ok returns a tuple with the ExistPhase2Proposal2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistPhase2Proposal2

`func (o *VPN) SetExistPhase2Proposal2(v bool)`

SetExistPhase2Proposal2 sets ExistPhase2Proposal2 field to given value.

### HasExistPhase2Proposal2

`func (o *VPN) HasExistPhase2Proposal2() bool`

HasExistPhase2Proposal2 returns a boolean if a field has been set.

### GetExistServerOpenVpnGoogleLdap

`func (o *VPN) GetExistServerOpenVpnGoogleLdap() bool`

GetExistServerOpenVpnGoogleLdap returns the ExistServerOpenVpnGoogleLdap field if non-nil, zero value otherwise.

### GetExistServerOpenVpnGoogleLdapOk

`func (o *VPN) GetExistServerOpenVpnGoogleLdapOk() (*bool, bool)`

GetExistServerOpenVpnGoogleLdapOk returns a tuple with the ExistServerOpenVpnGoogleLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistServerOpenVpnGoogleLdap

`func (o *VPN) SetExistServerOpenVpnGoogleLdap(v bool)`

SetExistServerOpenVpnGoogleLdap sets ExistServerOpenVpnGoogleLdap field to given value.

### HasExistServerOpenVpnGoogleLdap

`func (o *VPN) HasExistServerOpenVpnGoogleLdap() bool`

HasExistServerOpenVpnGoogleLdap returns a boolean if a field has been set.

### GetExistTunnelMode

`func (o *VPN) GetExistTunnelMode() bool`

GetExistTunnelMode returns the ExistTunnelMode field if non-nil, zero value otherwise.

### GetExistTunnelModeOk

`func (o *VPN) GetExistTunnelModeOk() (*bool, bool)`

GetExistTunnelModeOk returns a tuple with the ExistTunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistTunnelMode

`func (o *VPN) SetExistTunnelMode(v bool)`

SetExistTunnelMode sets ExistTunnelMode field to given value.

### HasExistTunnelMode

`func (o *VPN) HasExistTunnelMode() bool`

HasExistTunnelMode returns a boolean if a field has been set.

### GetId

`func (o *VPN) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPN) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPN) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPN) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpPool

`func (o *VPN) GetIpPool() IPSubnetsVO`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *VPN) GetIpPoolOk() (*IPSubnetsVO, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *VPN) SetIpPool(v IPSubnetsVO)`

SetIpPool sets IpPool field to given value.


### GetIpPoolEnd

`func (o *VPN) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *VPN) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *VPN) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *VPN) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *VPN) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *VPN) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *VPN) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *VPN) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetIpPoolType

`func (o *VPN) GetIpPoolType() int32`

GetIpPoolType returns the IpPoolType field if non-nil, zero value otherwise.

### GetIpPoolTypeOk

`func (o *VPN) GetIpPoolTypeOk() (*int32, bool)`

GetIpPoolTypeOk returns a tuple with the IpPoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolType

`func (o *VPN) SetIpPoolType(v int32)`

SetIpPoolType sets IpPoolType field to given value.

### HasIpPoolType

`func (o *VPN) HasIpPoolType() bool`

HasIpPoolType returns a boolean if a field has been set.

### GetLdapProfile

`func (o *VPN) GetLdapProfile() string`

GetLdapProfile returns the LdapProfile field if non-nil, zero value otherwise.

### GetLdapProfileOk

`func (o *VPN) GetLdapProfileOk() (*string, bool)`

GetLdapProfileOk returns a tuple with the LdapProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProfile

`func (o *VPN) SetLdapProfile(v string)`

SetLdapProfile sets LdapProfile field to given value.

### HasLdapProfile

`func (o *VPN) HasLdapProfile() bool`

HasLdapProfile returns a boolean if a field has been set.

### GetName

`func (o *VPN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPN) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkList

`func (o *VPN) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VPN) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VPN) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VPN) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *VPN) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VPN) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VPN) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VPN) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOpenVpnMode

`func (o *VPN) GetOpenVpnMode() int32`

GetOpenVpnMode returns the OpenVpnMode field if non-nil, zero value otherwise.

### GetOpenVpnModeOk

`func (o *VPN) GetOpenVpnModeOk() (*int32, bool)`

GetOpenVpnModeOk returns a tuple with the OpenVpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVpnMode

`func (o *VPN) SetOpenVpnMode(v int32)`

SetOpenVpnMode sets OpenVpnMode field to given value.

### HasOpenVpnMode

`func (o *VPN) HasOpenVpnMode() bool`

HasOpenVpnMode returns a boolean if a field has been set.

### GetOpenVpnTunnelMode

`func (o *VPN) GetOpenVpnTunnelMode() int32`

GetOpenVpnTunnelMode returns the OpenVpnTunnelMode field if non-nil, zero value otherwise.

### GetOpenVpnTunnelModeOk

`func (o *VPN) GetOpenVpnTunnelModeOk() (*int32, bool)`

GetOpenVpnTunnelModeOk returns a tuple with the OpenVpnTunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVpnTunnelMode

`func (o *VPN) SetOpenVpnTunnelMode(v int32)`

SetOpenVpnTunnelMode sets OpenVpnTunnelMode field to given value.

### HasOpenVpnTunnelMode

`func (o *VPN) HasOpenVpnTunnelMode() bool`

HasOpenVpnTunnelMode returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *VPN) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VPN) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VPN) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *VPN) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *VPN) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *VPN) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *VPN) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *VPN) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetPurpose

`func (o *VPN) GetPurpose() int32`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *VPN) GetPurposeOk() (*int32, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *VPN) SetPurpose(v int32)`

SetPurpose sets Purpose field to given value.


### GetRemoteIp

`func (o *VPN) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VPN) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VPN) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *VPN) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteSite

`func (o *VPN) GetRemoteSite() string`

GetRemoteSite returns the RemoteSite field if non-nil, zero value otherwise.

### GetRemoteSiteOk

`func (o *VPN) GetRemoteSiteOk() (*string, bool)`

GetRemoteSiteOk returns a tuple with the RemoteSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSite

`func (o *VPN) SetRemoteSite(v string)`

SetRemoteSite sets RemoteSite field to given value.

### HasRemoteSite

`func (o *VPN) HasRemoteSite() bool`

HasRemoteSite returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *VPN) GetRemoteSubnet() []IPSubnetsVO`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VPN) GetRemoteSubnetOk() (*[]IPSubnetsVO, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VPN) SetRemoteSubnet(v []IPSubnetsVO)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *VPN) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *VPN) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *VPN) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *VPN) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *VPN) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetServicePort

`func (o *VPN) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VPN) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VPN) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *VPN) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceType

`func (o *VPN) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *VPN) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *VPN) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *VPN) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSiteVpnType

`func (o *VPN) GetSiteVpnType() int32`

GetSiteVpnType returns the SiteVpnType field if non-nil, zero value otherwise.

### GetSiteVpnTypeOk

`func (o *VPN) GetSiteVpnTypeOk() (*int32, bool)`

GetSiteVpnTypeOk returns a tuple with the SiteVpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteVpnType

`func (o *VPN) SetSiteVpnType(v int32)`

SetSiteVpnType sets SiteVpnType field to given value.

### HasSiteVpnType

`func (o *VPN) HasSiteVpnType() bool`

HasSiteVpnType returns a boolean if a field has been set.

### GetStatus

`func (o *VPN) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VPN) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VPN) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VPN) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVpnConfiguration

`func (o *VPN) GetVpnConfiguration() VpnCertificateOpenApiVO`

GetVpnConfiguration returns the VpnConfiguration field if non-nil, zero value otherwise.

### GetVpnConfigurationOk

`func (o *VPN) GetVpnConfigurationOk() (*VpnCertificateOpenApiVO, bool)`

GetVpnConfigurationOk returns a tuple with the VpnConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConfiguration

`func (o *VPN) SetVpnConfiguration(v VpnCertificateOpenApiVO)`

SetVpnConfiguration sets VpnConfiguration field to given value.

### HasVpnConfiguration

`func (o *VPN) HasVpnConfiguration() bool`

HasVpnConfiguration returns a boolean if a field has been set.

### GetWans

`func (o *VPN) GetWans() []string`

GetWans returns the Wans field if non-nil, zero value otherwise.

### GetWansOk

`func (o *VPN) GetWansOk() (*[]string, bool)`

GetWansOk returns a tuple with the Wans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWans

`func (o *VPN) SetWans(v []string)`

SetWans sets Wans field to given value.

### HasWans

`func (o *VPN) HasWans() bool`

HasWans returns a boolean if a field has been set.

### GetWorkingMode

`func (o *VPN) GetWorkingMode() int32`

GetWorkingMode returns the WorkingMode field if non-nil, zero value otherwise.

### GetWorkingModeOk

`func (o *VPN) GetWorkingModeOk() (*int32, bool)`

GetWorkingModeOk returns a tuple with the WorkingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingMode

`func (o *VPN) SetWorkingMode(v int32)`

SetWorkingMode sets WorkingMode field to given value.

### HasWorkingMode

`func (o *VPN) HasWorkingMode() bool`

HasWorkingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


