# SiteSettingCapVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppControl** | Pointer to **bool** |  | [optional] 
**ApplicationControl** | Pointer to **bool** |  | [optional] 
**ArpTableSupport** | Pointer to **bool** |  | [optional] 
**AttackDefense** | Pointer to **bool** |  | [optional] 
**BeaconControl** | Pointer to **bool** |  | [optional] 
**CertProfile** | Pointer to **bool** |  | [optional] 
**ClientRateLimit** | Pointer to **bool** |  | [optional] 
**Cluster** | Pointer to **bool** |  | [optional] 
**ClusterModeOn** | Pointer to **bool** |  | [optional] 
**CustomAcl** | Pointer to **bool** |  | [optional] 
**Ddns** | Pointer to **bool** |  | [optional] 
**DhcpReservation** | Pointer to **bool** |  | [optional] 
**DisableNat** | Pointer to **bool** |  | [optional] 
**DnsCache** | Pointer to **bool** |  | [optional] 
**DnsLoopUpSupport** | Pointer to **bool** |  | [optional] 
**DnsProxy** | Pointer to **bool** |  | [optional] 
**Dpi** | Pointer to **bool** |  | [optional] 
**DpiStat** | Pointer to **bool** |  | [optional] 
**Dsl** | Pointer to **bool** |  | [optional] 
**Firewall** | Pointer to **bool** |  | [optional] 
**GoogleLdap** | Pointer to **bool** |  | [optional] 
**InitClusterReminder** | Pointer to **bool** |  | [optional] 
**Internet** | Pointer to **bool** |  | [optional] 
**IpMacBinding** | Pointer to **bool** |  | [optional] 
**IpPortGroup** | Pointer to **bool** |  | [optional] 
**IpsIds** | Pointer to **bool** |  | [optional] 
**Ipsec** | Pointer to **bool** |  | [optional] 
**IpsecFailover** | Pointer to **bool** |  | [optional] 
**Iptv** | Pointer to **bool** |  | [optional] 
**IsolationSettings** | Pointer to **bool** |  | [optional] 
**L2TP** | Pointer to **bool** |  | [optional] 
**LanDns** | Pointer to **bool** |  | [optional] 
**LdapSsl** | Pointer to **bool** |  | [optional] 
**LdapVpn** | Pointer to **bool** |  | [optional] 
**LockToAp** | Pointer to **bool** |  | [optional] 
**MacFilter** | Pointer to **bool** |  | [optional] 
**NetworkCheckSupport** | Pointer to **bool** |  | [optional] 
**NetworkSecurity** | Pointer to **bool** |  | [optional] 
**OltVlan** | Pointer to **bool** |  | [optional] 
**OneToOneNat** | Pointer to **bool** |  | [optional] 
**P2p** | Pointer to **bool** |  | [optional] 
**PackageCaptureGateway** | Pointer to **bool** |  | [optional] 
**PacketCaptureSupport** | Pointer to **bool** |  | [optional] 
**PeerEndpointDomain** | Pointer to **bool** |  | [optional] 
**PingSupport** | Pointer to **bool** |  | [optional] 
**PolicyRouting** | Pointer to **bool** |  | [optional] 
**Qos** | Pointer to **bool** |  | [optional] 
**Radios** | Pointer to **bool** |  | [optional] 
**ServerOpenVpnGoogleLdap** | Pointer to **bool** |  | [optional] 
**ServiceIptv** | Pointer to **bool** |  | [optional] 
**ServiceType** | Pointer to **bool** |  | [optional] 
**Sim** | Pointer to **bool** |  | [optional] 
**SnmpServiceSetting** | Pointer to **bool** |  | [optional] 
**Ssh** | Pointer to **bool** |  | [optional] 
**SslVpn** | Pointer to **bool** |  | [optional] 
**Statistics** | Pointer to **bool** |  | [optional] 
**SubVpn** | Pointer to **bool** |  | [optional] 
**SupportDpi** | Pointer to **bool** |  | [optional] 
**SupportES** | Pointer to **bool** |  | [optional] 
**SupportL2** | Pointer to **bool** |  | [optional] 
**SupportL3** | Pointer to **bool** |  | [optional] 
**SupportShowServerInReservation** | Pointer to **bool** |  | [optional] 
**TerminalSupport** | Pointer to **bool** |  | [optional] 
**TracerouteSupport** | Pointer to **bool** |  | [optional] 
**Transmission** | Pointer to **bool** |  | [optional] 
**Upnp** | Pointer to **bool** |  | [optional] 
**UrlCategory** | Pointer to **bool** |  | [optional] 
**VirtualWan** | Pointer to **bool** |  | [optional] 
**Voip** | Pointer to **bool** |  | [optional] 
**VpnStatus** | Pointer to **bool** |  | [optional] 
**VpnUser** | Pointer to **bool** |  | [optional] 
**Wireguard** | Pointer to **bool** |  | [optional] 
**Wireless** | Pointer to **bool** |  | [optional] 
**Wlans** | Pointer to **bool** |  | [optional] 

## Methods

### NewSiteSettingCapVO

`func NewSiteSettingCapVO() *SiteSettingCapVO`

NewSiteSettingCapVO instantiates a new SiteSettingCapVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingCapVOWithDefaults

`func NewSiteSettingCapVOWithDefaults() *SiteSettingCapVO`

NewSiteSettingCapVOWithDefaults instantiates a new SiteSettingCapVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppControl

`func (o *SiteSettingCapVO) GetAppControl() bool`

GetAppControl returns the AppControl field if non-nil, zero value otherwise.

### GetAppControlOk

`func (o *SiteSettingCapVO) GetAppControlOk() (*bool, bool)`

GetAppControlOk returns a tuple with the AppControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppControl

`func (o *SiteSettingCapVO) SetAppControl(v bool)`

SetAppControl sets AppControl field to given value.

### HasAppControl

`func (o *SiteSettingCapVO) HasAppControl() bool`

HasAppControl returns a boolean if a field has been set.

### GetApplicationControl

`func (o *SiteSettingCapVO) GetApplicationControl() bool`

GetApplicationControl returns the ApplicationControl field if non-nil, zero value otherwise.

### GetApplicationControlOk

`func (o *SiteSettingCapVO) GetApplicationControlOk() (*bool, bool)`

GetApplicationControlOk returns a tuple with the ApplicationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationControl

`func (o *SiteSettingCapVO) SetApplicationControl(v bool)`

SetApplicationControl sets ApplicationControl field to given value.

### HasApplicationControl

`func (o *SiteSettingCapVO) HasApplicationControl() bool`

HasApplicationControl returns a boolean if a field has been set.

### GetArpTableSupport

`func (o *SiteSettingCapVO) GetArpTableSupport() bool`

GetArpTableSupport returns the ArpTableSupport field if non-nil, zero value otherwise.

### GetArpTableSupportOk

`func (o *SiteSettingCapVO) GetArpTableSupportOk() (*bool, bool)`

GetArpTableSupportOk returns a tuple with the ArpTableSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpTableSupport

`func (o *SiteSettingCapVO) SetArpTableSupport(v bool)`

SetArpTableSupport sets ArpTableSupport field to given value.

### HasArpTableSupport

`func (o *SiteSettingCapVO) HasArpTableSupport() bool`

HasArpTableSupport returns a boolean if a field has been set.

### GetAttackDefense

`func (o *SiteSettingCapVO) GetAttackDefense() bool`

GetAttackDefense returns the AttackDefense field if non-nil, zero value otherwise.

### GetAttackDefenseOk

`func (o *SiteSettingCapVO) GetAttackDefenseOk() (*bool, bool)`

GetAttackDefenseOk returns a tuple with the AttackDefense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackDefense

`func (o *SiteSettingCapVO) SetAttackDefense(v bool)`

SetAttackDefense sets AttackDefense field to given value.

### HasAttackDefense

`func (o *SiteSettingCapVO) HasAttackDefense() bool`

HasAttackDefense returns a boolean if a field has been set.

### GetBeaconControl

`func (o *SiteSettingCapVO) GetBeaconControl() bool`

GetBeaconControl returns the BeaconControl field if non-nil, zero value otherwise.

### GetBeaconControlOk

`func (o *SiteSettingCapVO) GetBeaconControlOk() (*bool, bool)`

GetBeaconControlOk returns a tuple with the BeaconControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconControl

`func (o *SiteSettingCapVO) SetBeaconControl(v bool)`

SetBeaconControl sets BeaconControl field to given value.

### HasBeaconControl

`func (o *SiteSettingCapVO) HasBeaconControl() bool`

HasBeaconControl returns a boolean if a field has been set.

### GetCertProfile

`func (o *SiteSettingCapVO) GetCertProfile() bool`

GetCertProfile returns the CertProfile field if non-nil, zero value otherwise.

### GetCertProfileOk

`func (o *SiteSettingCapVO) GetCertProfileOk() (*bool, bool)`

GetCertProfileOk returns a tuple with the CertProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertProfile

`func (o *SiteSettingCapVO) SetCertProfile(v bool)`

SetCertProfile sets CertProfile field to given value.

### HasCertProfile

`func (o *SiteSettingCapVO) HasCertProfile() bool`

HasCertProfile returns a boolean if a field has been set.

### GetClientRateLimit

`func (o *SiteSettingCapVO) GetClientRateLimit() bool`

GetClientRateLimit returns the ClientRateLimit field if non-nil, zero value otherwise.

### GetClientRateLimitOk

`func (o *SiteSettingCapVO) GetClientRateLimitOk() (*bool, bool)`

GetClientRateLimitOk returns a tuple with the ClientRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRateLimit

`func (o *SiteSettingCapVO) SetClientRateLimit(v bool)`

SetClientRateLimit sets ClientRateLimit field to given value.

### HasClientRateLimit

`func (o *SiteSettingCapVO) HasClientRateLimit() bool`

HasClientRateLimit returns a boolean if a field has been set.

### GetCluster

`func (o *SiteSettingCapVO) GetCluster() bool`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SiteSettingCapVO) GetClusterOk() (*bool, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SiteSettingCapVO) SetCluster(v bool)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SiteSettingCapVO) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterModeOn

`func (o *SiteSettingCapVO) GetClusterModeOn() bool`

GetClusterModeOn returns the ClusterModeOn field if non-nil, zero value otherwise.

### GetClusterModeOnOk

`func (o *SiteSettingCapVO) GetClusterModeOnOk() (*bool, bool)`

GetClusterModeOnOk returns a tuple with the ClusterModeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterModeOn

`func (o *SiteSettingCapVO) SetClusterModeOn(v bool)`

SetClusterModeOn sets ClusterModeOn field to given value.

### HasClusterModeOn

`func (o *SiteSettingCapVO) HasClusterModeOn() bool`

HasClusterModeOn returns a boolean if a field has been set.

### GetCustomAcl

`func (o *SiteSettingCapVO) GetCustomAcl() bool`

GetCustomAcl returns the CustomAcl field if non-nil, zero value otherwise.

### GetCustomAclOk

`func (o *SiteSettingCapVO) GetCustomAclOk() (*bool, bool)`

GetCustomAclOk returns a tuple with the CustomAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAcl

`func (o *SiteSettingCapVO) SetCustomAcl(v bool)`

SetCustomAcl sets CustomAcl field to given value.

### HasCustomAcl

`func (o *SiteSettingCapVO) HasCustomAcl() bool`

HasCustomAcl returns a boolean if a field has been set.

### GetDdns

`func (o *SiteSettingCapVO) GetDdns() bool`

GetDdns returns the Ddns field if non-nil, zero value otherwise.

### GetDdnsOk

`func (o *SiteSettingCapVO) GetDdnsOk() (*bool, bool)`

GetDdnsOk returns a tuple with the Ddns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdns

`func (o *SiteSettingCapVO) SetDdns(v bool)`

SetDdns sets Ddns field to given value.

### HasDdns

`func (o *SiteSettingCapVO) HasDdns() bool`

HasDdns returns a boolean if a field has been set.

### GetDhcpReservation

`func (o *SiteSettingCapVO) GetDhcpReservation() bool`

GetDhcpReservation returns the DhcpReservation field if non-nil, zero value otherwise.

### GetDhcpReservationOk

`func (o *SiteSettingCapVO) GetDhcpReservationOk() (*bool, bool)`

GetDhcpReservationOk returns a tuple with the DhcpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpReservation

`func (o *SiteSettingCapVO) SetDhcpReservation(v bool)`

SetDhcpReservation sets DhcpReservation field to given value.

### HasDhcpReservation

`func (o *SiteSettingCapVO) HasDhcpReservation() bool`

HasDhcpReservation returns a boolean if a field has been set.

### GetDisableNat

`func (o *SiteSettingCapVO) GetDisableNat() bool`

GetDisableNat returns the DisableNat field if non-nil, zero value otherwise.

### GetDisableNatOk

`func (o *SiteSettingCapVO) GetDisableNatOk() (*bool, bool)`

GetDisableNatOk returns a tuple with the DisableNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNat

`func (o *SiteSettingCapVO) SetDisableNat(v bool)`

SetDisableNat sets DisableNat field to given value.

### HasDisableNat

`func (o *SiteSettingCapVO) HasDisableNat() bool`

HasDisableNat returns a boolean if a field has been set.

### GetDnsCache

`func (o *SiteSettingCapVO) GetDnsCache() bool`

GetDnsCache returns the DnsCache field if non-nil, zero value otherwise.

### GetDnsCacheOk

`func (o *SiteSettingCapVO) GetDnsCacheOk() (*bool, bool)`

GetDnsCacheOk returns a tuple with the DnsCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCache

`func (o *SiteSettingCapVO) SetDnsCache(v bool)`

SetDnsCache sets DnsCache field to given value.

### HasDnsCache

`func (o *SiteSettingCapVO) HasDnsCache() bool`

HasDnsCache returns a boolean if a field has been set.

### GetDnsLoopUpSupport

`func (o *SiteSettingCapVO) GetDnsLoopUpSupport() bool`

GetDnsLoopUpSupport returns the DnsLoopUpSupport field if non-nil, zero value otherwise.

### GetDnsLoopUpSupportOk

`func (o *SiteSettingCapVO) GetDnsLoopUpSupportOk() (*bool, bool)`

GetDnsLoopUpSupportOk returns a tuple with the DnsLoopUpSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLoopUpSupport

`func (o *SiteSettingCapVO) SetDnsLoopUpSupport(v bool)`

SetDnsLoopUpSupport sets DnsLoopUpSupport field to given value.

### HasDnsLoopUpSupport

`func (o *SiteSettingCapVO) HasDnsLoopUpSupport() bool`

HasDnsLoopUpSupport returns a boolean if a field has been set.

### GetDnsProxy

`func (o *SiteSettingCapVO) GetDnsProxy() bool`

GetDnsProxy returns the DnsProxy field if non-nil, zero value otherwise.

### GetDnsProxyOk

`func (o *SiteSettingCapVO) GetDnsProxyOk() (*bool, bool)`

GetDnsProxyOk returns a tuple with the DnsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProxy

`func (o *SiteSettingCapVO) SetDnsProxy(v bool)`

SetDnsProxy sets DnsProxy field to given value.

### HasDnsProxy

`func (o *SiteSettingCapVO) HasDnsProxy() bool`

HasDnsProxy returns a boolean if a field has been set.

### GetDpi

`func (o *SiteSettingCapVO) GetDpi() bool`

GetDpi returns the Dpi field if non-nil, zero value otherwise.

### GetDpiOk

`func (o *SiteSettingCapVO) GetDpiOk() (*bool, bool)`

GetDpiOk returns a tuple with the Dpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpi

`func (o *SiteSettingCapVO) SetDpi(v bool)`

SetDpi sets Dpi field to given value.

### HasDpi

`func (o *SiteSettingCapVO) HasDpi() bool`

HasDpi returns a boolean if a field has been set.

### GetDpiStat

`func (o *SiteSettingCapVO) GetDpiStat() bool`

GetDpiStat returns the DpiStat field if non-nil, zero value otherwise.

### GetDpiStatOk

`func (o *SiteSettingCapVO) GetDpiStatOk() (*bool, bool)`

GetDpiStatOk returns a tuple with the DpiStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpiStat

`func (o *SiteSettingCapVO) SetDpiStat(v bool)`

SetDpiStat sets DpiStat field to given value.

### HasDpiStat

`func (o *SiteSettingCapVO) HasDpiStat() bool`

HasDpiStat returns a boolean if a field has been set.

### GetDsl

`func (o *SiteSettingCapVO) GetDsl() bool`

GetDsl returns the Dsl field if non-nil, zero value otherwise.

### GetDslOk

`func (o *SiteSettingCapVO) GetDslOk() (*bool, bool)`

GetDslOk returns a tuple with the Dsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsl

`func (o *SiteSettingCapVO) SetDsl(v bool)`

SetDsl sets Dsl field to given value.

### HasDsl

`func (o *SiteSettingCapVO) HasDsl() bool`

HasDsl returns a boolean if a field has been set.

### GetFirewall

`func (o *SiteSettingCapVO) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *SiteSettingCapVO) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *SiteSettingCapVO) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *SiteSettingCapVO) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetGoogleLdap

`func (o *SiteSettingCapVO) GetGoogleLdap() bool`

GetGoogleLdap returns the GoogleLdap field if non-nil, zero value otherwise.

### GetGoogleLdapOk

`func (o *SiteSettingCapVO) GetGoogleLdapOk() (*bool, bool)`

GetGoogleLdapOk returns a tuple with the GoogleLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleLdap

`func (o *SiteSettingCapVO) SetGoogleLdap(v bool)`

SetGoogleLdap sets GoogleLdap field to given value.

### HasGoogleLdap

`func (o *SiteSettingCapVO) HasGoogleLdap() bool`

HasGoogleLdap returns a boolean if a field has been set.

### GetInitClusterReminder

`func (o *SiteSettingCapVO) GetInitClusterReminder() bool`

GetInitClusterReminder returns the InitClusterReminder field if non-nil, zero value otherwise.

### GetInitClusterReminderOk

`func (o *SiteSettingCapVO) GetInitClusterReminderOk() (*bool, bool)`

GetInitClusterReminderOk returns a tuple with the InitClusterReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitClusterReminder

`func (o *SiteSettingCapVO) SetInitClusterReminder(v bool)`

SetInitClusterReminder sets InitClusterReminder field to given value.

### HasInitClusterReminder

`func (o *SiteSettingCapVO) HasInitClusterReminder() bool`

HasInitClusterReminder returns a boolean if a field has been set.

### GetInternet

`func (o *SiteSettingCapVO) GetInternet() bool`

GetInternet returns the Internet field if non-nil, zero value otherwise.

### GetInternetOk

`func (o *SiteSettingCapVO) GetInternetOk() (*bool, bool)`

GetInternetOk returns a tuple with the Internet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternet

`func (o *SiteSettingCapVO) SetInternet(v bool)`

SetInternet sets Internet field to given value.

### HasInternet

`func (o *SiteSettingCapVO) HasInternet() bool`

HasInternet returns a boolean if a field has been set.

### GetIpMacBinding

`func (o *SiteSettingCapVO) GetIpMacBinding() bool`

GetIpMacBinding returns the IpMacBinding field if non-nil, zero value otherwise.

### GetIpMacBindingOk

`func (o *SiteSettingCapVO) GetIpMacBindingOk() (*bool, bool)`

GetIpMacBindingOk returns a tuple with the IpMacBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMacBinding

`func (o *SiteSettingCapVO) SetIpMacBinding(v bool)`

SetIpMacBinding sets IpMacBinding field to given value.

### HasIpMacBinding

`func (o *SiteSettingCapVO) HasIpMacBinding() bool`

HasIpMacBinding returns a boolean if a field has been set.

### GetIpPortGroup

`func (o *SiteSettingCapVO) GetIpPortGroup() bool`

GetIpPortGroup returns the IpPortGroup field if non-nil, zero value otherwise.

### GetIpPortGroupOk

`func (o *SiteSettingCapVO) GetIpPortGroupOk() (*bool, bool)`

GetIpPortGroupOk returns a tuple with the IpPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPortGroup

`func (o *SiteSettingCapVO) SetIpPortGroup(v bool)`

SetIpPortGroup sets IpPortGroup field to given value.

### HasIpPortGroup

`func (o *SiteSettingCapVO) HasIpPortGroup() bool`

HasIpPortGroup returns a boolean if a field has been set.

### GetIpsIds

`func (o *SiteSettingCapVO) GetIpsIds() bool`

GetIpsIds returns the IpsIds field if non-nil, zero value otherwise.

### GetIpsIdsOk

`func (o *SiteSettingCapVO) GetIpsIdsOk() (*bool, bool)`

GetIpsIdsOk returns a tuple with the IpsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsIds

`func (o *SiteSettingCapVO) SetIpsIds(v bool)`

SetIpsIds sets IpsIds field to given value.

### HasIpsIds

`func (o *SiteSettingCapVO) HasIpsIds() bool`

HasIpsIds returns a boolean if a field has been set.

### GetIpsec

`func (o *SiteSettingCapVO) GetIpsec() bool`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *SiteSettingCapVO) GetIpsecOk() (*bool, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *SiteSettingCapVO) SetIpsec(v bool)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *SiteSettingCapVO) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetIpsecFailover

`func (o *SiteSettingCapVO) GetIpsecFailover() bool`

GetIpsecFailover returns the IpsecFailover field if non-nil, zero value otherwise.

### GetIpsecFailoverOk

`func (o *SiteSettingCapVO) GetIpsecFailoverOk() (*bool, bool)`

GetIpsecFailoverOk returns a tuple with the IpsecFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecFailover

`func (o *SiteSettingCapVO) SetIpsecFailover(v bool)`

SetIpsecFailover sets IpsecFailover field to given value.

### HasIpsecFailover

`func (o *SiteSettingCapVO) HasIpsecFailover() bool`

HasIpsecFailover returns a boolean if a field has been set.

### GetIptv

`func (o *SiteSettingCapVO) GetIptv() bool`

GetIptv returns the Iptv field if non-nil, zero value otherwise.

### GetIptvOk

`func (o *SiteSettingCapVO) GetIptvOk() (*bool, bool)`

GetIptvOk returns a tuple with the Iptv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptv

`func (o *SiteSettingCapVO) SetIptv(v bool)`

SetIptv sets Iptv field to given value.

### HasIptv

`func (o *SiteSettingCapVO) HasIptv() bool`

HasIptv returns a boolean if a field has been set.

### GetIsolationSettings

`func (o *SiteSettingCapVO) GetIsolationSettings() bool`

GetIsolationSettings returns the IsolationSettings field if non-nil, zero value otherwise.

### GetIsolationSettingsOk

`func (o *SiteSettingCapVO) GetIsolationSettingsOk() (*bool, bool)`

GetIsolationSettingsOk returns a tuple with the IsolationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationSettings

`func (o *SiteSettingCapVO) SetIsolationSettings(v bool)`

SetIsolationSettings sets IsolationSettings field to given value.

### HasIsolationSettings

`func (o *SiteSettingCapVO) HasIsolationSettings() bool`

HasIsolationSettings returns a boolean if a field has been set.

### GetL2TP

`func (o *SiteSettingCapVO) GetL2TP() bool`

GetL2TP returns the L2TP field if non-nil, zero value otherwise.

### GetL2TPOk

`func (o *SiteSettingCapVO) GetL2TPOk() (*bool, bool)`

GetL2TPOk returns a tuple with the L2TP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2TP

`func (o *SiteSettingCapVO) SetL2TP(v bool)`

SetL2TP sets L2TP field to given value.

### HasL2TP

`func (o *SiteSettingCapVO) HasL2TP() bool`

HasL2TP returns a boolean if a field has been set.

### GetLanDns

`func (o *SiteSettingCapVO) GetLanDns() bool`

GetLanDns returns the LanDns field if non-nil, zero value otherwise.

### GetLanDnsOk

`func (o *SiteSettingCapVO) GetLanDnsOk() (*bool, bool)`

GetLanDnsOk returns a tuple with the LanDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanDns

`func (o *SiteSettingCapVO) SetLanDns(v bool)`

SetLanDns sets LanDns field to given value.

### HasLanDns

`func (o *SiteSettingCapVO) HasLanDns() bool`

HasLanDns returns a boolean if a field has been set.

### GetLdapSsl

`func (o *SiteSettingCapVO) GetLdapSsl() bool`

GetLdapSsl returns the LdapSsl field if non-nil, zero value otherwise.

### GetLdapSslOk

`func (o *SiteSettingCapVO) GetLdapSslOk() (*bool, bool)`

GetLdapSslOk returns a tuple with the LdapSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSsl

`func (o *SiteSettingCapVO) SetLdapSsl(v bool)`

SetLdapSsl sets LdapSsl field to given value.

### HasLdapSsl

`func (o *SiteSettingCapVO) HasLdapSsl() bool`

HasLdapSsl returns a boolean if a field has been set.

### GetLdapVpn

`func (o *SiteSettingCapVO) GetLdapVpn() bool`

GetLdapVpn returns the LdapVpn field if non-nil, zero value otherwise.

### GetLdapVpnOk

`func (o *SiteSettingCapVO) GetLdapVpnOk() (*bool, bool)`

GetLdapVpnOk returns a tuple with the LdapVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapVpn

`func (o *SiteSettingCapVO) SetLdapVpn(v bool)`

SetLdapVpn sets LdapVpn field to given value.

### HasLdapVpn

`func (o *SiteSettingCapVO) HasLdapVpn() bool`

HasLdapVpn returns a boolean if a field has been set.

### GetLockToAp

`func (o *SiteSettingCapVO) GetLockToAp() bool`

GetLockToAp returns the LockToAp field if non-nil, zero value otherwise.

### GetLockToApOk

`func (o *SiteSettingCapVO) GetLockToApOk() (*bool, bool)`

GetLockToApOk returns a tuple with the LockToAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockToAp

`func (o *SiteSettingCapVO) SetLockToAp(v bool)`

SetLockToAp sets LockToAp field to given value.

### HasLockToAp

`func (o *SiteSettingCapVO) HasLockToAp() bool`

HasLockToAp returns a boolean if a field has been set.

### GetMacFilter

`func (o *SiteSettingCapVO) GetMacFilter() bool`

GetMacFilter returns the MacFilter field if non-nil, zero value otherwise.

### GetMacFilterOk

`func (o *SiteSettingCapVO) GetMacFilterOk() (*bool, bool)`

GetMacFilterOk returns a tuple with the MacFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilter

`func (o *SiteSettingCapVO) SetMacFilter(v bool)`

SetMacFilter sets MacFilter field to given value.

### HasMacFilter

`func (o *SiteSettingCapVO) HasMacFilter() bool`

HasMacFilter returns a boolean if a field has been set.

### GetNetworkCheckSupport

`func (o *SiteSettingCapVO) GetNetworkCheckSupport() bool`

GetNetworkCheckSupport returns the NetworkCheckSupport field if non-nil, zero value otherwise.

### GetNetworkCheckSupportOk

`func (o *SiteSettingCapVO) GetNetworkCheckSupportOk() (*bool, bool)`

GetNetworkCheckSupportOk returns a tuple with the NetworkCheckSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCheckSupport

`func (o *SiteSettingCapVO) SetNetworkCheckSupport(v bool)`

SetNetworkCheckSupport sets NetworkCheckSupport field to given value.

### HasNetworkCheckSupport

`func (o *SiteSettingCapVO) HasNetworkCheckSupport() bool`

HasNetworkCheckSupport returns a boolean if a field has been set.

### GetNetworkSecurity

`func (o *SiteSettingCapVO) GetNetworkSecurity() bool`

GetNetworkSecurity returns the NetworkSecurity field if non-nil, zero value otherwise.

### GetNetworkSecurityOk

`func (o *SiteSettingCapVO) GetNetworkSecurityOk() (*bool, bool)`

GetNetworkSecurityOk returns a tuple with the NetworkSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecurity

`func (o *SiteSettingCapVO) SetNetworkSecurity(v bool)`

SetNetworkSecurity sets NetworkSecurity field to given value.

### HasNetworkSecurity

`func (o *SiteSettingCapVO) HasNetworkSecurity() bool`

HasNetworkSecurity returns a boolean if a field has been set.

### GetOltVlan

`func (o *SiteSettingCapVO) GetOltVlan() bool`

GetOltVlan returns the OltVlan field if non-nil, zero value otherwise.

### GetOltVlanOk

`func (o *SiteSettingCapVO) GetOltVlanOk() (*bool, bool)`

GetOltVlanOk returns a tuple with the OltVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOltVlan

`func (o *SiteSettingCapVO) SetOltVlan(v bool)`

SetOltVlan sets OltVlan field to given value.

### HasOltVlan

`func (o *SiteSettingCapVO) HasOltVlan() bool`

HasOltVlan returns a boolean if a field has been set.

### GetOneToOneNat

`func (o *SiteSettingCapVO) GetOneToOneNat() bool`

GetOneToOneNat returns the OneToOneNat field if non-nil, zero value otherwise.

### GetOneToOneNatOk

`func (o *SiteSettingCapVO) GetOneToOneNatOk() (*bool, bool)`

GetOneToOneNatOk returns a tuple with the OneToOneNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneToOneNat

`func (o *SiteSettingCapVO) SetOneToOneNat(v bool)`

SetOneToOneNat sets OneToOneNat field to given value.

### HasOneToOneNat

`func (o *SiteSettingCapVO) HasOneToOneNat() bool`

HasOneToOneNat returns a boolean if a field has been set.

### GetP2p

`func (o *SiteSettingCapVO) GetP2p() bool`

GetP2p returns the P2p field if non-nil, zero value otherwise.

### GetP2pOk

`func (o *SiteSettingCapVO) GetP2pOk() (*bool, bool)`

GetP2pOk returns a tuple with the P2p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2p

`func (o *SiteSettingCapVO) SetP2p(v bool)`

SetP2p sets P2p field to given value.

### HasP2p

`func (o *SiteSettingCapVO) HasP2p() bool`

HasP2p returns a boolean if a field has been set.

### GetPackageCaptureGateway

`func (o *SiteSettingCapVO) GetPackageCaptureGateway() bool`

GetPackageCaptureGateway returns the PackageCaptureGateway field if non-nil, zero value otherwise.

### GetPackageCaptureGatewayOk

`func (o *SiteSettingCapVO) GetPackageCaptureGatewayOk() (*bool, bool)`

GetPackageCaptureGatewayOk returns a tuple with the PackageCaptureGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCaptureGateway

`func (o *SiteSettingCapVO) SetPackageCaptureGateway(v bool)`

SetPackageCaptureGateway sets PackageCaptureGateway field to given value.

### HasPackageCaptureGateway

`func (o *SiteSettingCapVO) HasPackageCaptureGateway() bool`

HasPackageCaptureGateway returns a boolean if a field has been set.

### GetPacketCaptureSupport

`func (o *SiteSettingCapVO) GetPacketCaptureSupport() bool`

GetPacketCaptureSupport returns the PacketCaptureSupport field if non-nil, zero value otherwise.

### GetPacketCaptureSupportOk

`func (o *SiteSettingCapVO) GetPacketCaptureSupportOk() (*bool, bool)`

GetPacketCaptureSupportOk returns a tuple with the PacketCaptureSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketCaptureSupport

`func (o *SiteSettingCapVO) SetPacketCaptureSupport(v bool)`

SetPacketCaptureSupport sets PacketCaptureSupport field to given value.

### HasPacketCaptureSupport

`func (o *SiteSettingCapVO) HasPacketCaptureSupport() bool`

HasPacketCaptureSupport returns a boolean if a field has been set.

### GetPeerEndpointDomain

`func (o *SiteSettingCapVO) GetPeerEndpointDomain() bool`

GetPeerEndpointDomain returns the PeerEndpointDomain field if non-nil, zero value otherwise.

### GetPeerEndpointDomainOk

`func (o *SiteSettingCapVO) GetPeerEndpointDomainOk() (*bool, bool)`

GetPeerEndpointDomainOk returns a tuple with the PeerEndpointDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerEndpointDomain

`func (o *SiteSettingCapVO) SetPeerEndpointDomain(v bool)`

SetPeerEndpointDomain sets PeerEndpointDomain field to given value.

### HasPeerEndpointDomain

`func (o *SiteSettingCapVO) HasPeerEndpointDomain() bool`

HasPeerEndpointDomain returns a boolean if a field has been set.

### GetPingSupport

`func (o *SiteSettingCapVO) GetPingSupport() bool`

GetPingSupport returns the PingSupport field if non-nil, zero value otherwise.

### GetPingSupportOk

`func (o *SiteSettingCapVO) GetPingSupportOk() (*bool, bool)`

GetPingSupportOk returns a tuple with the PingSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSupport

`func (o *SiteSettingCapVO) SetPingSupport(v bool)`

SetPingSupport sets PingSupport field to given value.

### HasPingSupport

`func (o *SiteSettingCapVO) HasPingSupport() bool`

HasPingSupport returns a boolean if a field has been set.

### GetPolicyRouting

`func (o *SiteSettingCapVO) GetPolicyRouting() bool`

GetPolicyRouting returns the PolicyRouting field if non-nil, zero value otherwise.

### GetPolicyRoutingOk

`func (o *SiteSettingCapVO) GetPolicyRoutingOk() (*bool, bool)`

GetPolicyRoutingOk returns a tuple with the PolicyRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRouting

`func (o *SiteSettingCapVO) SetPolicyRouting(v bool)`

SetPolicyRouting sets PolicyRouting field to given value.

### HasPolicyRouting

`func (o *SiteSettingCapVO) HasPolicyRouting() bool`

HasPolicyRouting returns a boolean if a field has been set.

### GetQos

`func (o *SiteSettingCapVO) GetQos() bool`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *SiteSettingCapVO) GetQosOk() (*bool, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *SiteSettingCapVO) SetQos(v bool)`

SetQos sets Qos field to given value.

### HasQos

`func (o *SiteSettingCapVO) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetRadios

`func (o *SiteSettingCapVO) GetRadios() bool`

GetRadios returns the Radios field if non-nil, zero value otherwise.

### GetRadiosOk

`func (o *SiteSettingCapVO) GetRadiosOk() (*bool, bool)`

GetRadiosOk returns a tuple with the Radios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadios

`func (o *SiteSettingCapVO) SetRadios(v bool)`

SetRadios sets Radios field to given value.

### HasRadios

`func (o *SiteSettingCapVO) HasRadios() bool`

HasRadios returns a boolean if a field has been set.

### GetServerOpenVpnGoogleLdap

`func (o *SiteSettingCapVO) GetServerOpenVpnGoogleLdap() bool`

GetServerOpenVpnGoogleLdap returns the ServerOpenVpnGoogleLdap field if non-nil, zero value otherwise.

### GetServerOpenVpnGoogleLdapOk

`func (o *SiteSettingCapVO) GetServerOpenVpnGoogleLdapOk() (*bool, bool)`

GetServerOpenVpnGoogleLdapOk returns a tuple with the ServerOpenVpnGoogleLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOpenVpnGoogleLdap

`func (o *SiteSettingCapVO) SetServerOpenVpnGoogleLdap(v bool)`

SetServerOpenVpnGoogleLdap sets ServerOpenVpnGoogleLdap field to given value.

### HasServerOpenVpnGoogleLdap

`func (o *SiteSettingCapVO) HasServerOpenVpnGoogleLdap() bool`

HasServerOpenVpnGoogleLdap returns a boolean if a field has been set.

### GetServiceIptv

`func (o *SiteSettingCapVO) GetServiceIptv() bool`

GetServiceIptv returns the ServiceIptv field if non-nil, zero value otherwise.

### GetServiceIptvOk

`func (o *SiteSettingCapVO) GetServiceIptvOk() (*bool, bool)`

GetServiceIptvOk returns a tuple with the ServiceIptv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIptv

`func (o *SiteSettingCapVO) SetServiceIptv(v bool)`

SetServiceIptv sets ServiceIptv field to given value.

### HasServiceIptv

`func (o *SiteSettingCapVO) HasServiceIptv() bool`

HasServiceIptv returns a boolean if a field has been set.

### GetServiceType

`func (o *SiteSettingCapVO) GetServiceType() bool`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *SiteSettingCapVO) GetServiceTypeOk() (*bool, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *SiteSettingCapVO) SetServiceType(v bool)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *SiteSettingCapVO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSim

`func (o *SiteSettingCapVO) GetSim() bool`

GetSim returns the Sim field if non-nil, zero value otherwise.

### GetSimOk

`func (o *SiteSettingCapVO) GetSimOk() (*bool, bool)`

GetSimOk returns a tuple with the Sim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSim

`func (o *SiteSettingCapVO) SetSim(v bool)`

SetSim sets Sim field to given value.

### HasSim

`func (o *SiteSettingCapVO) HasSim() bool`

HasSim returns a boolean if a field has been set.

### GetSnmpServiceSetting

`func (o *SiteSettingCapVO) GetSnmpServiceSetting() bool`

GetSnmpServiceSetting returns the SnmpServiceSetting field if non-nil, zero value otherwise.

### GetSnmpServiceSettingOk

`func (o *SiteSettingCapVO) GetSnmpServiceSettingOk() (*bool, bool)`

GetSnmpServiceSettingOk returns a tuple with the SnmpServiceSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpServiceSetting

`func (o *SiteSettingCapVO) SetSnmpServiceSetting(v bool)`

SetSnmpServiceSetting sets SnmpServiceSetting field to given value.

### HasSnmpServiceSetting

`func (o *SiteSettingCapVO) HasSnmpServiceSetting() bool`

HasSnmpServiceSetting returns a boolean if a field has been set.

### GetSsh

`func (o *SiteSettingCapVO) GetSsh() bool`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *SiteSettingCapVO) GetSshOk() (*bool, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *SiteSettingCapVO) SetSsh(v bool)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *SiteSettingCapVO) HasSsh() bool`

HasSsh returns a boolean if a field has been set.

### GetSslVpn

`func (o *SiteSettingCapVO) GetSslVpn() bool`

GetSslVpn returns the SslVpn field if non-nil, zero value otherwise.

### GetSslVpnOk

`func (o *SiteSettingCapVO) GetSslVpnOk() (*bool, bool)`

GetSslVpnOk returns a tuple with the SslVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVpn

`func (o *SiteSettingCapVO) SetSslVpn(v bool)`

SetSslVpn sets SslVpn field to given value.

### HasSslVpn

`func (o *SiteSettingCapVO) HasSslVpn() bool`

HasSslVpn returns a boolean if a field has been set.

### GetStatistics

`func (o *SiteSettingCapVO) GetStatistics() bool`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *SiteSettingCapVO) GetStatisticsOk() (*bool, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *SiteSettingCapVO) SetStatistics(v bool)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *SiteSettingCapVO) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetSubVpn

`func (o *SiteSettingCapVO) GetSubVpn() bool`

GetSubVpn returns the SubVpn field if non-nil, zero value otherwise.

### GetSubVpnOk

`func (o *SiteSettingCapVO) GetSubVpnOk() (*bool, bool)`

GetSubVpnOk returns a tuple with the SubVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVpn

`func (o *SiteSettingCapVO) SetSubVpn(v bool)`

SetSubVpn sets SubVpn field to given value.

### HasSubVpn

`func (o *SiteSettingCapVO) HasSubVpn() bool`

HasSubVpn returns a boolean if a field has been set.

### GetSupportDpi

`func (o *SiteSettingCapVO) GetSupportDpi() bool`

GetSupportDpi returns the SupportDpi field if non-nil, zero value otherwise.

### GetSupportDpiOk

`func (o *SiteSettingCapVO) GetSupportDpiOk() (*bool, bool)`

GetSupportDpiOk returns a tuple with the SupportDpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDpi

`func (o *SiteSettingCapVO) SetSupportDpi(v bool)`

SetSupportDpi sets SupportDpi field to given value.

### HasSupportDpi

`func (o *SiteSettingCapVO) HasSupportDpi() bool`

HasSupportDpi returns a boolean if a field has been set.

### GetSupportES

`func (o *SiteSettingCapVO) GetSupportES() bool`

GetSupportES returns the SupportES field if non-nil, zero value otherwise.

### GetSupportESOk

`func (o *SiteSettingCapVO) GetSupportESOk() (*bool, bool)`

GetSupportESOk returns a tuple with the SupportES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportES

`func (o *SiteSettingCapVO) SetSupportES(v bool)`

SetSupportES sets SupportES field to given value.

### HasSupportES

`func (o *SiteSettingCapVO) HasSupportES() bool`

HasSupportES returns a boolean if a field has been set.

### GetSupportL2

`func (o *SiteSettingCapVO) GetSupportL2() bool`

GetSupportL2 returns the SupportL2 field if non-nil, zero value otherwise.

### GetSupportL2Ok

`func (o *SiteSettingCapVO) GetSupportL2Ok() (*bool, bool)`

GetSupportL2Ok returns a tuple with the SupportL2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2

`func (o *SiteSettingCapVO) SetSupportL2(v bool)`

SetSupportL2 sets SupportL2 field to given value.

### HasSupportL2

`func (o *SiteSettingCapVO) HasSupportL2() bool`

HasSupportL2 returns a boolean if a field has been set.

### GetSupportL3

`func (o *SiteSettingCapVO) GetSupportL3() bool`

GetSupportL3 returns the SupportL3 field if non-nil, zero value otherwise.

### GetSupportL3Ok

`func (o *SiteSettingCapVO) GetSupportL3Ok() (*bool, bool)`

GetSupportL3Ok returns a tuple with the SupportL3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL3

`func (o *SiteSettingCapVO) SetSupportL3(v bool)`

SetSupportL3 sets SupportL3 field to given value.

### HasSupportL3

`func (o *SiteSettingCapVO) HasSupportL3() bool`

HasSupportL3 returns a boolean if a field has been set.

### GetSupportShowServerInReservation

`func (o *SiteSettingCapVO) GetSupportShowServerInReservation() bool`

GetSupportShowServerInReservation returns the SupportShowServerInReservation field if non-nil, zero value otherwise.

### GetSupportShowServerInReservationOk

`func (o *SiteSettingCapVO) GetSupportShowServerInReservationOk() (*bool, bool)`

GetSupportShowServerInReservationOk returns a tuple with the SupportShowServerInReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportShowServerInReservation

`func (o *SiteSettingCapVO) SetSupportShowServerInReservation(v bool)`

SetSupportShowServerInReservation sets SupportShowServerInReservation field to given value.

### HasSupportShowServerInReservation

`func (o *SiteSettingCapVO) HasSupportShowServerInReservation() bool`

HasSupportShowServerInReservation returns a boolean if a field has been set.

### GetTerminalSupport

`func (o *SiteSettingCapVO) GetTerminalSupport() bool`

GetTerminalSupport returns the TerminalSupport field if non-nil, zero value otherwise.

### GetTerminalSupportOk

`func (o *SiteSettingCapVO) GetTerminalSupportOk() (*bool, bool)`

GetTerminalSupportOk returns a tuple with the TerminalSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalSupport

`func (o *SiteSettingCapVO) SetTerminalSupport(v bool)`

SetTerminalSupport sets TerminalSupport field to given value.

### HasTerminalSupport

`func (o *SiteSettingCapVO) HasTerminalSupport() bool`

HasTerminalSupport returns a boolean if a field has been set.

### GetTracerouteSupport

`func (o *SiteSettingCapVO) GetTracerouteSupport() bool`

GetTracerouteSupport returns the TracerouteSupport field if non-nil, zero value otherwise.

### GetTracerouteSupportOk

`func (o *SiteSettingCapVO) GetTracerouteSupportOk() (*bool, bool)`

GetTracerouteSupportOk returns a tuple with the TracerouteSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracerouteSupport

`func (o *SiteSettingCapVO) SetTracerouteSupport(v bool)`

SetTracerouteSupport sets TracerouteSupport field to given value.

### HasTracerouteSupport

`func (o *SiteSettingCapVO) HasTracerouteSupport() bool`

HasTracerouteSupport returns a boolean if a field has been set.

### GetTransmission

`func (o *SiteSettingCapVO) GetTransmission() bool`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *SiteSettingCapVO) GetTransmissionOk() (*bool, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *SiteSettingCapVO) SetTransmission(v bool)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *SiteSettingCapVO) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetUpnp

`func (o *SiteSettingCapVO) GetUpnp() bool`

GetUpnp returns the Upnp field if non-nil, zero value otherwise.

### GetUpnpOk

`func (o *SiteSettingCapVO) GetUpnpOk() (*bool, bool)`

GetUpnpOk returns a tuple with the Upnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpnp

`func (o *SiteSettingCapVO) SetUpnp(v bool)`

SetUpnp sets Upnp field to given value.

### HasUpnp

`func (o *SiteSettingCapVO) HasUpnp() bool`

HasUpnp returns a boolean if a field has been set.

### GetUrlCategory

`func (o *SiteSettingCapVO) GetUrlCategory() bool`

GetUrlCategory returns the UrlCategory field if non-nil, zero value otherwise.

### GetUrlCategoryOk

`func (o *SiteSettingCapVO) GetUrlCategoryOk() (*bool, bool)`

GetUrlCategoryOk returns a tuple with the UrlCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCategory

`func (o *SiteSettingCapVO) SetUrlCategory(v bool)`

SetUrlCategory sets UrlCategory field to given value.

### HasUrlCategory

`func (o *SiteSettingCapVO) HasUrlCategory() bool`

HasUrlCategory returns a boolean if a field has been set.

### GetVirtualWan

`func (o *SiteSettingCapVO) GetVirtualWan() bool`

GetVirtualWan returns the VirtualWan field if non-nil, zero value otherwise.

### GetVirtualWanOk

`func (o *SiteSettingCapVO) GetVirtualWanOk() (*bool, bool)`

GetVirtualWanOk returns a tuple with the VirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWan

`func (o *SiteSettingCapVO) SetVirtualWan(v bool)`

SetVirtualWan sets VirtualWan field to given value.

### HasVirtualWan

`func (o *SiteSettingCapVO) HasVirtualWan() bool`

HasVirtualWan returns a boolean if a field has been set.

### GetVoip

`func (o *SiteSettingCapVO) GetVoip() bool`

GetVoip returns the Voip field if non-nil, zero value otherwise.

### GetVoipOk

`func (o *SiteSettingCapVO) GetVoipOk() (*bool, bool)`

GetVoipOk returns a tuple with the Voip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoip

`func (o *SiteSettingCapVO) SetVoip(v bool)`

SetVoip sets Voip field to given value.

### HasVoip

`func (o *SiteSettingCapVO) HasVoip() bool`

HasVoip returns a boolean if a field has been set.

### GetVpnStatus

`func (o *SiteSettingCapVO) GetVpnStatus() bool`

GetVpnStatus returns the VpnStatus field if non-nil, zero value otherwise.

### GetVpnStatusOk

`func (o *SiteSettingCapVO) GetVpnStatusOk() (*bool, bool)`

GetVpnStatusOk returns a tuple with the VpnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnStatus

`func (o *SiteSettingCapVO) SetVpnStatus(v bool)`

SetVpnStatus sets VpnStatus field to given value.

### HasVpnStatus

`func (o *SiteSettingCapVO) HasVpnStatus() bool`

HasVpnStatus returns a boolean if a field has been set.

### GetVpnUser

`func (o *SiteSettingCapVO) GetVpnUser() bool`

GetVpnUser returns the VpnUser field if non-nil, zero value otherwise.

### GetVpnUserOk

`func (o *SiteSettingCapVO) GetVpnUserOk() (*bool, bool)`

GetVpnUserOk returns a tuple with the VpnUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnUser

`func (o *SiteSettingCapVO) SetVpnUser(v bool)`

SetVpnUser sets VpnUser field to given value.

### HasVpnUser

`func (o *SiteSettingCapVO) HasVpnUser() bool`

HasVpnUser returns a boolean if a field has been set.

### GetWireguard

`func (o *SiteSettingCapVO) GetWireguard() bool`

GetWireguard returns the Wireguard field if non-nil, zero value otherwise.

### GetWireguardOk

`func (o *SiteSettingCapVO) GetWireguardOk() (*bool, bool)`

GetWireguardOk returns a tuple with the Wireguard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguard

`func (o *SiteSettingCapVO) SetWireguard(v bool)`

SetWireguard sets Wireguard field to given value.

### HasWireguard

`func (o *SiteSettingCapVO) HasWireguard() bool`

HasWireguard returns a boolean if a field has been set.

### GetWireless

`func (o *SiteSettingCapVO) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *SiteSettingCapVO) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *SiteSettingCapVO) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *SiteSettingCapVO) HasWireless() bool`

HasWireless returns a boolean if a field has been set.

### GetWlans

`func (o *SiteSettingCapVO) GetWlans() bool`

GetWlans returns the Wlans field if non-nil, zero value otherwise.

### GetWlansOk

`func (o *SiteSettingCapVO) GetWlansOk() (*bool, bool)`

GetWlansOk returns a tuple with the Wlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlans

`func (o *SiteSettingCapVO) SetWlans(v bool)`

SetWlans sets Wlans field to given value.

### HasWlans

`func (o *SiteSettingCapVO) HasWlans() bool`

HasWlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


