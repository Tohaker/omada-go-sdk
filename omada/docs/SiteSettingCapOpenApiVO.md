# SiteSettingCapOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsolationSettings** | Pointer to **bool** | Whether support Network Isolation. | [optional] 
**ClientRateLimit** | Pointer to **bool** | false: there is a wireless route device that does not support ClientRateLimit.true or null: all devices support ClientRateLimit | [optional] 
**Cluster** | Pointer to **bool** | Whether support cluster mode. | [optional] 
**ClusterModeOn** | Pointer to **bool** | Whether the current software is running in cluster mode. | [optional] 
**CustomAcl** | Pointer to **bool** | Whether ACL support custom mode. | [optional] 
**DisableNat** | Pointer to **bool** | Whether support disable NAT configuration. | [optional] 
**DnsProxy** | Pointer to **bool** | Whether support dnsProxy configuration. | [optional] 
**Dpi** | Pointer to **bool** | Whether support DPI configuration. Deprecated. | [optional] 
**DpiStat** | Pointer to **bool** | Whether support dpi stat.Deprecated, same as dpi. Deprecated. | [optional] 
**Dsl** | Pointer to **bool** | Whether support DSL configuration. | [optional] 
**Firewall** | Pointer to **bool** | Whether support firewall configuration. | [optional] 
**GoogleLdap** | Pointer to **bool** | Whether support Google LDAP configuration. | [optional] 
**InitClusterReminder** | Pointer to **bool** | Whether the current cluster prompts that it has not been restarted after configuration initialization. | [optional] 
**IpMacBinding** | Pointer to **bool** | Whether support ipMacBinding configuration. | [optional] 
**IpPortGroup** | Pointer to **bool** | Whether support IP-Port Group configuration. | [optional] 
**IpsIds** | Pointer to **bool** | Whether support IPS configuration. | [optional] 
**Ipsec** | Pointer to **bool** | Whether support Ipsec VPN configuration. | [optional] 
**IpsecFailover** | Pointer to **bool** | Whether support IPsec failover configuration. | [optional] 
**L2TP** | Pointer to **bool** | Whether support L2TP VPN configuration. | [optional] 
**LanDns** | Pointer to **bool** | Whether support LAN DNS configuration. | [optional] 
**LdapSsl** | Pointer to **bool** | Whether SSL VPN support ldap. | [optional] 
**LdapVpn** | Pointer to **bool** | Whether support VPN configuration with LDAP. | [optional] 
**LockToAp** | Pointer to **bool** | false: there is a wireless route device that does not support LockToAp.true or null: all devices support LockToAp | [optional] 
**MacFilter** | Pointer to **bool** | Whether support network security-mac filter configuration. | [optional] 
**OltVlan** | Pointer to **bool** | An OLT device exists. | [optional] 
**OneToOneNat** | Pointer to **bool** | Whether support one-to-one NAT configuration. | [optional] 
**P2p** | Pointer to **bool** | An P2P device exists. | [optional] 
**PackageCaptureGateway** | Pointer to **bool** | Whether the gateway can be selected in the packet capture page. | [optional] 
**PeerEndpointDomain** | Pointer to **bool** | Whether peer-endpoint support entering domain names. | [optional] 
**PolicyRouting** | Pointer to **bool** | Whether support policyRouting. | [optional] 
**Qos** | Pointer to **bool** | Whether support QOS configuration. | [optional] 
**ServerOpenVpnGoogleLdap** | Pointer to **bool** | Whether server-openvpn support Google LDAP. | [optional] 
**ServiceIptv** | Pointer to **bool** | Whether support IGMP or IPTV. | [optional] 
**ServiceType** | Pointer to **bool** | Whether support serviceType configuration. | [optional] 
**Sim** | Pointer to **bool** | Whether support SIM configuration. | [optional] 
**SpeedTest** | Pointer to **bool** | Whether support speedTest. | [optional] 
**SslVpn** | Pointer to **bool** | Whether support SSL VPN configuration. | [optional] 
**SubVpn** | Pointer to **bool** | Whether support VPN menu in VPN module. | [optional] 
**SupportDpi** | Pointer to **bool** | Whether support Dpi. | [optional] 
**SupportES** | Pointer to **bool** | Whether the site supports adopting Agile Series Switches | [optional] 
**SupportL2** | Pointer to **bool** | Whether the site supports adopting Non-Agile Series Switches | [optional] 
**SupportShowServerInReservation** | Pointer to **bool** | Whether support show server in DHCP Reservation. | [optional] 
**UrlCategory** | Pointer to **bool** | Whether Url Filter support category mode. | [optional] 
**VirtualWan** | Pointer to **bool** | Whether support virtual WAN configuration. | [optional] 
**Voip** | Pointer to **bool** | Whether support voip configuration. | [optional] 
**VpnStatus** | Pointer to **bool** | Whether support vpnStatus. | [optional] 
**VpnUser** | Pointer to **bool** | Whether support VPN User configuration. | [optional] 
**Wireguard** | Pointer to **bool** | Whether support wireguard VPN configuration. | [optional] 

## Methods

### NewSiteSettingCapOpenApiVO

`func NewSiteSettingCapOpenApiVO() *SiteSettingCapOpenApiVO`

NewSiteSettingCapOpenApiVO instantiates a new SiteSettingCapOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingCapOpenApiVOWithDefaults

`func NewSiteSettingCapOpenApiVOWithDefaults() *SiteSettingCapOpenApiVO`

NewSiteSettingCapOpenApiVOWithDefaults instantiates a new SiteSettingCapOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsolationSettings

`func (o *SiteSettingCapOpenApiVO) GetIsolationSettings() bool`

GetIsolationSettings returns the IsolationSettings field if non-nil, zero value otherwise.

### GetIsolationSettingsOk

`func (o *SiteSettingCapOpenApiVO) GetIsolationSettingsOk() (*bool, bool)`

GetIsolationSettingsOk returns a tuple with the IsolationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationSettings

`func (o *SiteSettingCapOpenApiVO) SetIsolationSettings(v bool)`

SetIsolationSettings sets IsolationSettings field to given value.

### HasIsolationSettings

`func (o *SiteSettingCapOpenApiVO) HasIsolationSettings() bool`

HasIsolationSettings returns a boolean if a field has been set.

### GetClientRateLimit

`func (o *SiteSettingCapOpenApiVO) GetClientRateLimit() bool`

GetClientRateLimit returns the ClientRateLimit field if non-nil, zero value otherwise.

### GetClientRateLimitOk

`func (o *SiteSettingCapOpenApiVO) GetClientRateLimitOk() (*bool, bool)`

GetClientRateLimitOk returns a tuple with the ClientRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRateLimit

`func (o *SiteSettingCapOpenApiVO) SetClientRateLimit(v bool)`

SetClientRateLimit sets ClientRateLimit field to given value.

### HasClientRateLimit

`func (o *SiteSettingCapOpenApiVO) HasClientRateLimit() bool`

HasClientRateLimit returns a boolean if a field has been set.

### GetCluster

`func (o *SiteSettingCapOpenApiVO) GetCluster() bool`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SiteSettingCapOpenApiVO) GetClusterOk() (*bool, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SiteSettingCapOpenApiVO) SetCluster(v bool)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SiteSettingCapOpenApiVO) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterModeOn

`func (o *SiteSettingCapOpenApiVO) GetClusterModeOn() bool`

GetClusterModeOn returns the ClusterModeOn field if non-nil, zero value otherwise.

### GetClusterModeOnOk

`func (o *SiteSettingCapOpenApiVO) GetClusterModeOnOk() (*bool, bool)`

GetClusterModeOnOk returns a tuple with the ClusterModeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterModeOn

`func (o *SiteSettingCapOpenApiVO) SetClusterModeOn(v bool)`

SetClusterModeOn sets ClusterModeOn field to given value.

### HasClusterModeOn

`func (o *SiteSettingCapOpenApiVO) HasClusterModeOn() bool`

HasClusterModeOn returns a boolean if a field has been set.

### GetCustomAcl

`func (o *SiteSettingCapOpenApiVO) GetCustomAcl() bool`

GetCustomAcl returns the CustomAcl field if non-nil, zero value otherwise.

### GetCustomAclOk

`func (o *SiteSettingCapOpenApiVO) GetCustomAclOk() (*bool, bool)`

GetCustomAclOk returns a tuple with the CustomAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAcl

`func (o *SiteSettingCapOpenApiVO) SetCustomAcl(v bool)`

SetCustomAcl sets CustomAcl field to given value.

### HasCustomAcl

`func (o *SiteSettingCapOpenApiVO) HasCustomAcl() bool`

HasCustomAcl returns a boolean if a field has been set.

### GetDisableNat

`func (o *SiteSettingCapOpenApiVO) GetDisableNat() bool`

GetDisableNat returns the DisableNat field if non-nil, zero value otherwise.

### GetDisableNatOk

`func (o *SiteSettingCapOpenApiVO) GetDisableNatOk() (*bool, bool)`

GetDisableNatOk returns a tuple with the DisableNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNat

`func (o *SiteSettingCapOpenApiVO) SetDisableNat(v bool)`

SetDisableNat sets DisableNat field to given value.

### HasDisableNat

`func (o *SiteSettingCapOpenApiVO) HasDisableNat() bool`

HasDisableNat returns a boolean if a field has been set.

### GetDnsProxy

`func (o *SiteSettingCapOpenApiVO) GetDnsProxy() bool`

GetDnsProxy returns the DnsProxy field if non-nil, zero value otherwise.

### GetDnsProxyOk

`func (o *SiteSettingCapOpenApiVO) GetDnsProxyOk() (*bool, bool)`

GetDnsProxyOk returns a tuple with the DnsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProxy

`func (o *SiteSettingCapOpenApiVO) SetDnsProxy(v bool)`

SetDnsProxy sets DnsProxy field to given value.

### HasDnsProxy

`func (o *SiteSettingCapOpenApiVO) HasDnsProxy() bool`

HasDnsProxy returns a boolean if a field has been set.

### GetDpi

`func (o *SiteSettingCapOpenApiVO) GetDpi() bool`

GetDpi returns the Dpi field if non-nil, zero value otherwise.

### GetDpiOk

`func (o *SiteSettingCapOpenApiVO) GetDpiOk() (*bool, bool)`

GetDpiOk returns a tuple with the Dpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpi

`func (o *SiteSettingCapOpenApiVO) SetDpi(v bool)`

SetDpi sets Dpi field to given value.

### HasDpi

`func (o *SiteSettingCapOpenApiVO) HasDpi() bool`

HasDpi returns a boolean if a field has been set.

### GetDpiStat

`func (o *SiteSettingCapOpenApiVO) GetDpiStat() bool`

GetDpiStat returns the DpiStat field if non-nil, zero value otherwise.

### GetDpiStatOk

`func (o *SiteSettingCapOpenApiVO) GetDpiStatOk() (*bool, bool)`

GetDpiStatOk returns a tuple with the DpiStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpiStat

`func (o *SiteSettingCapOpenApiVO) SetDpiStat(v bool)`

SetDpiStat sets DpiStat field to given value.

### HasDpiStat

`func (o *SiteSettingCapOpenApiVO) HasDpiStat() bool`

HasDpiStat returns a boolean if a field has been set.

### GetDsl

`func (o *SiteSettingCapOpenApiVO) GetDsl() bool`

GetDsl returns the Dsl field if non-nil, zero value otherwise.

### GetDslOk

`func (o *SiteSettingCapOpenApiVO) GetDslOk() (*bool, bool)`

GetDslOk returns a tuple with the Dsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsl

`func (o *SiteSettingCapOpenApiVO) SetDsl(v bool)`

SetDsl sets Dsl field to given value.

### HasDsl

`func (o *SiteSettingCapOpenApiVO) HasDsl() bool`

HasDsl returns a boolean if a field has been set.

### GetFirewall

`func (o *SiteSettingCapOpenApiVO) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *SiteSettingCapOpenApiVO) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *SiteSettingCapOpenApiVO) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *SiteSettingCapOpenApiVO) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetGoogleLdap

`func (o *SiteSettingCapOpenApiVO) GetGoogleLdap() bool`

GetGoogleLdap returns the GoogleLdap field if non-nil, zero value otherwise.

### GetGoogleLdapOk

`func (o *SiteSettingCapOpenApiVO) GetGoogleLdapOk() (*bool, bool)`

GetGoogleLdapOk returns a tuple with the GoogleLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleLdap

`func (o *SiteSettingCapOpenApiVO) SetGoogleLdap(v bool)`

SetGoogleLdap sets GoogleLdap field to given value.

### HasGoogleLdap

`func (o *SiteSettingCapOpenApiVO) HasGoogleLdap() bool`

HasGoogleLdap returns a boolean if a field has been set.

### GetInitClusterReminder

`func (o *SiteSettingCapOpenApiVO) GetInitClusterReminder() bool`

GetInitClusterReminder returns the InitClusterReminder field if non-nil, zero value otherwise.

### GetInitClusterReminderOk

`func (o *SiteSettingCapOpenApiVO) GetInitClusterReminderOk() (*bool, bool)`

GetInitClusterReminderOk returns a tuple with the InitClusterReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitClusterReminder

`func (o *SiteSettingCapOpenApiVO) SetInitClusterReminder(v bool)`

SetInitClusterReminder sets InitClusterReminder field to given value.

### HasInitClusterReminder

`func (o *SiteSettingCapOpenApiVO) HasInitClusterReminder() bool`

HasInitClusterReminder returns a boolean if a field has been set.

### GetIpMacBinding

`func (o *SiteSettingCapOpenApiVO) GetIpMacBinding() bool`

GetIpMacBinding returns the IpMacBinding field if non-nil, zero value otherwise.

### GetIpMacBindingOk

`func (o *SiteSettingCapOpenApiVO) GetIpMacBindingOk() (*bool, bool)`

GetIpMacBindingOk returns a tuple with the IpMacBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMacBinding

`func (o *SiteSettingCapOpenApiVO) SetIpMacBinding(v bool)`

SetIpMacBinding sets IpMacBinding field to given value.

### HasIpMacBinding

`func (o *SiteSettingCapOpenApiVO) HasIpMacBinding() bool`

HasIpMacBinding returns a boolean if a field has been set.

### GetIpPortGroup

`func (o *SiteSettingCapOpenApiVO) GetIpPortGroup() bool`

GetIpPortGroup returns the IpPortGroup field if non-nil, zero value otherwise.

### GetIpPortGroupOk

`func (o *SiteSettingCapOpenApiVO) GetIpPortGroupOk() (*bool, bool)`

GetIpPortGroupOk returns a tuple with the IpPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPortGroup

`func (o *SiteSettingCapOpenApiVO) SetIpPortGroup(v bool)`

SetIpPortGroup sets IpPortGroup field to given value.

### HasIpPortGroup

`func (o *SiteSettingCapOpenApiVO) HasIpPortGroup() bool`

HasIpPortGroup returns a boolean if a field has been set.

### GetIpsIds

`func (o *SiteSettingCapOpenApiVO) GetIpsIds() bool`

GetIpsIds returns the IpsIds field if non-nil, zero value otherwise.

### GetIpsIdsOk

`func (o *SiteSettingCapOpenApiVO) GetIpsIdsOk() (*bool, bool)`

GetIpsIdsOk returns a tuple with the IpsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsIds

`func (o *SiteSettingCapOpenApiVO) SetIpsIds(v bool)`

SetIpsIds sets IpsIds field to given value.

### HasIpsIds

`func (o *SiteSettingCapOpenApiVO) HasIpsIds() bool`

HasIpsIds returns a boolean if a field has been set.

### GetIpsec

`func (o *SiteSettingCapOpenApiVO) GetIpsec() bool`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *SiteSettingCapOpenApiVO) GetIpsecOk() (*bool, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *SiteSettingCapOpenApiVO) SetIpsec(v bool)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *SiteSettingCapOpenApiVO) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetIpsecFailover

`func (o *SiteSettingCapOpenApiVO) GetIpsecFailover() bool`

GetIpsecFailover returns the IpsecFailover field if non-nil, zero value otherwise.

### GetIpsecFailoverOk

`func (o *SiteSettingCapOpenApiVO) GetIpsecFailoverOk() (*bool, bool)`

GetIpsecFailoverOk returns a tuple with the IpsecFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecFailover

`func (o *SiteSettingCapOpenApiVO) SetIpsecFailover(v bool)`

SetIpsecFailover sets IpsecFailover field to given value.

### HasIpsecFailover

`func (o *SiteSettingCapOpenApiVO) HasIpsecFailover() bool`

HasIpsecFailover returns a boolean if a field has been set.

### GetL2TP

`func (o *SiteSettingCapOpenApiVO) GetL2TP() bool`

GetL2TP returns the L2TP field if non-nil, zero value otherwise.

### GetL2TPOk

`func (o *SiteSettingCapOpenApiVO) GetL2TPOk() (*bool, bool)`

GetL2TPOk returns a tuple with the L2TP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2TP

`func (o *SiteSettingCapOpenApiVO) SetL2TP(v bool)`

SetL2TP sets L2TP field to given value.

### HasL2TP

`func (o *SiteSettingCapOpenApiVO) HasL2TP() bool`

HasL2TP returns a boolean if a field has been set.

### GetLanDns

`func (o *SiteSettingCapOpenApiVO) GetLanDns() bool`

GetLanDns returns the LanDns field if non-nil, zero value otherwise.

### GetLanDnsOk

`func (o *SiteSettingCapOpenApiVO) GetLanDnsOk() (*bool, bool)`

GetLanDnsOk returns a tuple with the LanDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanDns

`func (o *SiteSettingCapOpenApiVO) SetLanDns(v bool)`

SetLanDns sets LanDns field to given value.

### HasLanDns

`func (o *SiteSettingCapOpenApiVO) HasLanDns() bool`

HasLanDns returns a boolean if a field has been set.

### GetLdapSsl

`func (o *SiteSettingCapOpenApiVO) GetLdapSsl() bool`

GetLdapSsl returns the LdapSsl field if non-nil, zero value otherwise.

### GetLdapSslOk

`func (o *SiteSettingCapOpenApiVO) GetLdapSslOk() (*bool, bool)`

GetLdapSslOk returns a tuple with the LdapSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSsl

`func (o *SiteSettingCapOpenApiVO) SetLdapSsl(v bool)`

SetLdapSsl sets LdapSsl field to given value.

### HasLdapSsl

`func (o *SiteSettingCapOpenApiVO) HasLdapSsl() bool`

HasLdapSsl returns a boolean if a field has been set.

### GetLdapVpn

`func (o *SiteSettingCapOpenApiVO) GetLdapVpn() bool`

GetLdapVpn returns the LdapVpn field if non-nil, zero value otherwise.

### GetLdapVpnOk

`func (o *SiteSettingCapOpenApiVO) GetLdapVpnOk() (*bool, bool)`

GetLdapVpnOk returns a tuple with the LdapVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapVpn

`func (o *SiteSettingCapOpenApiVO) SetLdapVpn(v bool)`

SetLdapVpn sets LdapVpn field to given value.

### HasLdapVpn

`func (o *SiteSettingCapOpenApiVO) HasLdapVpn() bool`

HasLdapVpn returns a boolean if a field has been set.

### GetLockToAp

`func (o *SiteSettingCapOpenApiVO) GetLockToAp() bool`

GetLockToAp returns the LockToAp field if non-nil, zero value otherwise.

### GetLockToApOk

`func (o *SiteSettingCapOpenApiVO) GetLockToApOk() (*bool, bool)`

GetLockToApOk returns a tuple with the LockToAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockToAp

`func (o *SiteSettingCapOpenApiVO) SetLockToAp(v bool)`

SetLockToAp sets LockToAp field to given value.

### HasLockToAp

`func (o *SiteSettingCapOpenApiVO) HasLockToAp() bool`

HasLockToAp returns a boolean if a field has been set.

### GetMacFilter

`func (o *SiteSettingCapOpenApiVO) GetMacFilter() bool`

GetMacFilter returns the MacFilter field if non-nil, zero value otherwise.

### GetMacFilterOk

`func (o *SiteSettingCapOpenApiVO) GetMacFilterOk() (*bool, bool)`

GetMacFilterOk returns a tuple with the MacFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilter

`func (o *SiteSettingCapOpenApiVO) SetMacFilter(v bool)`

SetMacFilter sets MacFilter field to given value.

### HasMacFilter

`func (o *SiteSettingCapOpenApiVO) HasMacFilter() bool`

HasMacFilter returns a boolean if a field has been set.

### GetOltVlan

`func (o *SiteSettingCapOpenApiVO) GetOltVlan() bool`

GetOltVlan returns the OltVlan field if non-nil, zero value otherwise.

### GetOltVlanOk

`func (o *SiteSettingCapOpenApiVO) GetOltVlanOk() (*bool, bool)`

GetOltVlanOk returns a tuple with the OltVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOltVlan

`func (o *SiteSettingCapOpenApiVO) SetOltVlan(v bool)`

SetOltVlan sets OltVlan field to given value.

### HasOltVlan

`func (o *SiteSettingCapOpenApiVO) HasOltVlan() bool`

HasOltVlan returns a boolean if a field has been set.

### GetOneToOneNat

`func (o *SiteSettingCapOpenApiVO) GetOneToOneNat() bool`

GetOneToOneNat returns the OneToOneNat field if non-nil, zero value otherwise.

### GetOneToOneNatOk

`func (o *SiteSettingCapOpenApiVO) GetOneToOneNatOk() (*bool, bool)`

GetOneToOneNatOk returns a tuple with the OneToOneNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneToOneNat

`func (o *SiteSettingCapOpenApiVO) SetOneToOneNat(v bool)`

SetOneToOneNat sets OneToOneNat field to given value.

### HasOneToOneNat

`func (o *SiteSettingCapOpenApiVO) HasOneToOneNat() bool`

HasOneToOneNat returns a boolean if a field has been set.

### GetP2p

`func (o *SiteSettingCapOpenApiVO) GetP2p() bool`

GetP2p returns the P2p field if non-nil, zero value otherwise.

### GetP2pOk

`func (o *SiteSettingCapOpenApiVO) GetP2pOk() (*bool, bool)`

GetP2pOk returns a tuple with the P2p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2p

`func (o *SiteSettingCapOpenApiVO) SetP2p(v bool)`

SetP2p sets P2p field to given value.

### HasP2p

`func (o *SiteSettingCapOpenApiVO) HasP2p() bool`

HasP2p returns a boolean if a field has been set.

### GetPackageCaptureGateway

`func (o *SiteSettingCapOpenApiVO) GetPackageCaptureGateway() bool`

GetPackageCaptureGateway returns the PackageCaptureGateway field if non-nil, zero value otherwise.

### GetPackageCaptureGatewayOk

`func (o *SiteSettingCapOpenApiVO) GetPackageCaptureGatewayOk() (*bool, bool)`

GetPackageCaptureGatewayOk returns a tuple with the PackageCaptureGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCaptureGateway

`func (o *SiteSettingCapOpenApiVO) SetPackageCaptureGateway(v bool)`

SetPackageCaptureGateway sets PackageCaptureGateway field to given value.

### HasPackageCaptureGateway

`func (o *SiteSettingCapOpenApiVO) HasPackageCaptureGateway() bool`

HasPackageCaptureGateway returns a boolean if a field has been set.

### GetPeerEndpointDomain

`func (o *SiteSettingCapOpenApiVO) GetPeerEndpointDomain() bool`

GetPeerEndpointDomain returns the PeerEndpointDomain field if non-nil, zero value otherwise.

### GetPeerEndpointDomainOk

`func (o *SiteSettingCapOpenApiVO) GetPeerEndpointDomainOk() (*bool, bool)`

GetPeerEndpointDomainOk returns a tuple with the PeerEndpointDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerEndpointDomain

`func (o *SiteSettingCapOpenApiVO) SetPeerEndpointDomain(v bool)`

SetPeerEndpointDomain sets PeerEndpointDomain field to given value.

### HasPeerEndpointDomain

`func (o *SiteSettingCapOpenApiVO) HasPeerEndpointDomain() bool`

HasPeerEndpointDomain returns a boolean if a field has been set.

### GetPolicyRouting

`func (o *SiteSettingCapOpenApiVO) GetPolicyRouting() bool`

GetPolicyRouting returns the PolicyRouting field if non-nil, zero value otherwise.

### GetPolicyRoutingOk

`func (o *SiteSettingCapOpenApiVO) GetPolicyRoutingOk() (*bool, bool)`

GetPolicyRoutingOk returns a tuple with the PolicyRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRouting

`func (o *SiteSettingCapOpenApiVO) SetPolicyRouting(v bool)`

SetPolicyRouting sets PolicyRouting field to given value.

### HasPolicyRouting

`func (o *SiteSettingCapOpenApiVO) HasPolicyRouting() bool`

HasPolicyRouting returns a boolean if a field has been set.

### GetQos

`func (o *SiteSettingCapOpenApiVO) GetQos() bool`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *SiteSettingCapOpenApiVO) GetQosOk() (*bool, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *SiteSettingCapOpenApiVO) SetQos(v bool)`

SetQos sets Qos field to given value.

### HasQos

`func (o *SiteSettingCapOpenApiVO) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetServerOpenVpnGoogleLdap

`func (o *SiteSettingCapOpenApiVO) GetServerOpenVpnGoogleLdap() bool`

GetServerOpenVpnGoogleLdap returns the ServerOpenVpnGoogleLdap field if non-nil, zero value otherwise.

### GetServerOpenVpnGoogleLdapOk

`func (o *SiteSettingCapOpenApiVO) GetServerOpenVpnGoogleLdapOk() (*bool, bool)`

GetServerOpenVpnGoogleLdapOk returns a tuple with the ServerOpenVpnGoogleLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOpenVpnGoogleLdap

`func (o *SiteSettingCapOpenApiVO) SetServerOpenVpnGoogleLdap(v bool)`

SetServerOpenVpnGoogleLdap sets ServerOpenVpnGoogleLdap field to given value.

### HasServerOpenVpnGoogleLdap

`func (o *SiteSettingCapOpenApiVO) HasServerOpenVpnGoogleLdap() bool`

HasServerOpenVpnGoogleLdap returns a boolean if a field has been set.

### GetServiceIptv

`func (o *SiteSettingCapOpenApiVO) GetServiceIptv() bool`

GetServiceIptv returns the ServiceIptv field if non-nil, zero value otherwise.

### GetServiceIptvOk

`func (o *SiteSettingCapOpenApiVO) GetServiceIptvOk() (*bool, bool)`

GetServiceIptvOk returns a tuple with the ServiceIptv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIptv

`func (o *SiteSettingCapOpenApiVO) SetServiceIptv(v bool)`

SetServiceIptv sets ServiceIptv field to given value.

### HasServiceIptv

`func (o *SiteSettingCapOpenApiVO) HasServiceIptv() bool`

HasServiceIptv returns a boolean if a field has been set.

### GetServiceType

`func (o *SiteSettingCapOpenApiVO) GetServiceType() bool`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *SiteSettingCapOpenApiVO) GetServiceTypeOk() (*bool, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *SiteSettingCapOpenApiVO) SetServiceType(v bool)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *SiteSettingCapOpenApiVO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSim

`func (o *SiteSettingCapOpenApiVO) GetSim() bool`

GetSim returns the Sim field if non-nil, zero value otherwise.

### GetSimOk

`func (o *SiteSettingCapOpenApiVO) GetSimOk() (*bool, bool)`

GetSimOk returns a tuple with the Sim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSim

`func (o *SiteSettingCapOpenApiVO) SetSim(v bool)`

SetSim sets Sim field to given value.

### HasSim

`func (o *SiteSettingCapOpenApiVO) HasSim() bool`

HasSim returns a boolean if a field has been set.

### GetSpeedTest

`func (o *SiteSettingCapOpenApiVO) GetSpeedTest() bool`

GetSpeedTest returns the SpeedTest field if non-nil, zero value otherwise.

### GetSpeedTestOk

`func (o *SiteSettingCapOpenApiVO) GetSpeedTestOk() (*bool, bool)`

GetSpeedTestOk returns a tuple with the SpeedTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedTest

`func (o *SiteSettingCapOpenApiVO) SetSpeedTest(v bool)`

SetSpeedTest sets SpeedTest field to given value.

### HasSpeedTest

`func (o *SiteSettingCapOpenApiVO) HasSpeedTest() bool`

HasSpeedTest returns a boolean if a field has been set.

### GetSslVpn

`func (o *SiteSettingCapOpenApiVO) GetSslVpn() bool`

GetSslVpn returns the SslVpn field if non-nil, zero value otherwise.

### GetSslVpnOk

`func (o *SiteSettingCapOpenApiVO) GetSslVpnOk() (*bool, bool)`

GetSslVpnOk returns a tuple with the SslVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVpn

`func (o *SiteSettingCapOpenApiVO) SetSslVpn(v bool)`

SetSslVpn sets SslVpn field to given value.

### HasSslVpn

`func (o *SiteSettingCapOpenApiVO) HasSslVpn() bool`

HasSslVpn returns a boolean if a field has been set.

### GetSubVpn

`func (o *SiteSettingCapOpenApiVO) GetSubVpn() bool`

GetSubVpn returns the SubVpn field if non-nil, zero value otherwise.

### GetSubVpnOk

`func (o *SiteSettingCapOpenApiVO) GetSubVpnOk() (*bool, bool)`

GetSubVpnOk returns a tuple with the SubVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVpn

`func (o *SiteSettingCapOpenApiVO) SetSubVpn(v bool)`

SetSubVpn sets SubVpn field to given value.

### HasSubVpn

`func (o *SiteSettingCapOpenApiVO) HasSubVpn() bool`

HasSubVpn returns a boolean if a field has been set.

### GetSupportDpi

`func (o *SiteSettingCapOpenApiVO) GetSupportDpi() bool`

GetSupportDpi returns the SupportDpi field if non-nil, zero value otherwise.

### GetSupportDpiOk

`func (o *SiteSettingCapOpenApiVO) GetSupportDpiOk() (*bool, bool)`

GetSupportDpiOk returns a tuple with the SupportDpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDpi

`func (o *SiteSettingCapOpenApiVO) SetSupportDpi(v bool)`

SetSupportDpi sets SupportDpi field to given value.

### HasSupportDpi

`func (o *SiteSettingCapOpenApiVO) HasSupportDpi() bool`

HasSupportDpi returns a boolean if a field has been set.

### GetSupportES

`func (o *SiteSettingCapOpenApiVO) GetSupportES() bool`

GetSupportES returns the SupportES field if non-nil, zero value otherwise.

### GetSupportESOk

`func (o *SiteSettingCapOpenApiVO) GetSupportESOk() (*bool, bool)`

GetSupportESOk returns a tuple with the SupportES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportES

`func (o *SiteSettingCapOpenApiVO) SetSupportES(v bool)`

SetSupportES sets SupportES field to given value.

### HasSupportES

`func (o *SiteSettingCapOpenApiVO) HasSupportES() bool`

HasSupportES returns a boolean if a field has been set.

### GetSupportL2

`func (o *SiteSettingCapOpenApiVO) GetSupportL2() bool`

GetSupportL2 returns the SupportL2 field if non-nil, zero value otherwise.

### GetSupportL2Ok

`func (o *SiteSettingCapOpenApiVO) GetSupportL2Ok() (*bool, bool)`

GetSupportL2Ok returns a tuple with the SupportL2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2

`func (o *SiteSettingCapOpenApiVO) SetSupportL2(v bool)`

SetSupportL2 sets SupportL2 field to given value.

### HasSupportL2

`func (o *SiteSettingCapOpenApiVO) HasSupportL2() bool`

HasSupportL2 returns a boolean if a field has been set.

### GetSupportShowServerInReservation

`func (o *SiteSettingCapOpenApiVO) GetSupportShowServerInReservation() bool`

GetSupportShowServerInReservation returns the SupportShowServerInReservation field if non-nil, zero value otherwise.

### GetSupportShowServerInReservationOk

`func (o *SiteSettingCapOpenApiVO) GetSupportShowServerInReservationOk() (*bool, bool)`

GetSupportShowServerInReservationOk returns a tuple with the SupportShowServerInReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportShowServerInReservation

`func (o *SiteSettingCapOpenApiVO) SetSupportShowServerInReservation(v bool)`

SetSupportShowServerInReservation sets SupportShowServerInReservation field to given value.

### HasSupportShowServerInReservation

`func (o *SiteSettingCapOpenApiVO) HasSupportShowServerInReservation() bool`

HasSupportShowServerInReservation returns a boolean if a field has been set.

### GetUrlCategory

`func (o *SiteSettingCapOpenApiVO) GetUrlCategory() bool`

GetUrlCategory returns the UrlCategory field if non-nil, zero value otherwise.

### GetUrlCategoryOk

`func (o *SiteSettingCapOpenApiVO) GetUrlCategoryOk() (*bool, bool)`

GetUrlCategoryOk returns a tuple with the UrlCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCategory

`func (o *SiteSettingCapOpenApiVO) SetUrlCategory(v bool)`

SetUrlCategory sets UrlCategory field to given value.

### HasUrlCategory

`func (o *SiteSettingCapOpenApiVO) HasUrlCategory() bool`

HasUrlCategory returns a boolean if a field has been set.

### GetVirtualWan

`func (o *SiteSettingCapOpenApiVO) GetVirtualWan() bool`

GetVirtualWan returns the VirtualWan field if non-nil, zero value otherwise.

### GetVirtualWanOk

`func (o *SiteSettingCapOpenApiVO) GetVirtualWanOk() (*bool, bool)`

GetVirtualWanOk returns a tuple with the VirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWan

`func (o *SiteSettingCapOpenApiVO) SetVirtualWan(v bool)`

SetVirtualWan sets VirtualWan field to given value.

### HasVirtualWan

`func (o *SiteSettingCapOpenApiVO) HasVirtualWan() bool`

HasVirtualWan returns a boolean if a field has been set.

### GetVoip

`func (o *SiteSettingCapOpenApiVO) GetVoip() bool`

GetVoip returns the Voip field if non-nil, zero value otherwise.

### GetVoipOk

`func (o *SiteSettingCapOpenApiVO) GetVoipOk() (*bool, bool)`

GetVoipOk returns a tuple with the Voip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoip

`func (o *SiteSettingCapOpenApiVO) SetVoip(v bool)`

SetVoip sets Voip field to given value.

### HasVoip

`func (o *SiteSettingCapOpenApiVO) HasVoip() bool`

HasVoip returns a boolean if a field has been set.

### GetVpnStatus

`func (o *SiteSettingCapOpenApiVO) GetVpnStatus() bool`

GetVpnStatus returns the VpnStatus field if non-nil, zero value otherwise.

### GetVpnStatusOk

`func (o *SiteSettingCapOpenApiVO) GetVpnStatusOk() (*bool, bool)`

GetVpnStatusOk returns a tuple with the VpnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnStatus

`func (o *SiteSettingCapOpenApiVO) SetVpnStatus(v bool)`

SetVpnStatus sets VpnStatus field to given value.

### HasVpnStatus

`func (o *SiteSettingCapOpenApiVO) HasVpnStatus() bool`

HasVpnStatus returns a boolean if a field has been set.

### GetVpnUser

`func (o *SiteSettingCapOpenApiVO) GetVpnUser() bool`

GetVpnUser returns the VpnUser field if non-nil, zero value otherwise.

### GetVpnUserOk

`func (o *SiteSettingCapOpenApiVO) GetVpnUserOk() (*bool, bool)`

GetVpnUserOk returns a tuple with the VpnUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnUser

`func (o *SiteSettingCapOpenApiVO) SetVpnUser(v bool)`

SetVpnUser sets VpnUser field to given value.

### HasVpnUser

`func (o *SiteSettingCapOpenApiVO) HasVpnUser() bool`

HasVpnUser returns a boolean if a field has been set.

### GetWireguard

`func (o *SiteSettingCapOpenApiVO) GetWireguard() bool`

GetWireguard returns the Wireguard field if non-nil, zero value otherwise.

### GetWireguardOk

`func (o *SiteSettingCapOpenApiVO) GetWireguardOk() (*bool, bool)`

GetWireguardOk returns a tuple with the Wireguard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguard

`func (o *SiteSettingCapOpenApiVO) SetWireguard(v bool)`

SetWireguard sets Wireguard field to given value.

### HasWireguard

`func (o *SiteSettingCapOpenApiVO) HasWireguard() bool`

HasWireguard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


