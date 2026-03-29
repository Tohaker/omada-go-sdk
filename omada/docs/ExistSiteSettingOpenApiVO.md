# ExistSiteSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableNat** | Pointer to **bool** | Whether the current site has Disable NAT configured. | [optional] 
**DnsCache** | Pointer to **bool** | Whether the current site has DNS Cache configured. | [optional] 
**DnsProxy** | Pointer to **bool** | Whether the current site has DNS proxy configured. | [optional] 
**GoogleLdap** | Pointer to **bool** | Whether the current site has Google LDAP configured. | [optional] 
**IpMacBinding** | Pointer to **bool** | Whether the current site has IP-MAC Binding configured. | [optional] 
**IpsIds** | Pointer to **bool** | Whether the current site has IDS/IPS configured. | [optional] 
**IpsecFailover** | Pointer to **bool** | Whether the current site has IPsec failover configured. | [optional] 
**LanDns** | Pointer to **bool** | Whether the current site has LAN DNS configured. | [optional] 
**MacFilter** | Pointer to **bool** | Whether the current site has MAC Filtering configured. | [optional] 
**OneToOneNat** | Pointer to **bool** | Whether the current site has One-to-One NAT configured. | [optional] 
**PolicyRouting** | Pointer to **bool** | Whether the current site has Policy Routing configured. | [optional] 
**Qos** | Pointer to **bool** | Whether the current site has QoS configured. | [optional] 
**ServiceType** | Pointer to **bool** | Whether the current site has Gateway QoS service configured. | [optional] 
**SslVpn** | Pointer to **bool** | Whether the current site has SSL VPN configured. | [optional] 
**UrlCategory** | Pointer to **bool** | Whether the General Config of URL Filtering has been configured. | [optional] 
**VirtualWan** | Pointer to **bool** | Whether the current site has Virtual WAN configured. | [optional] 
**VpnUser** | Pointer to **bool** | Whether the current site has VPN User configured. | [optional] 
**Wireguard** | Pointer to **bool** | Whether the current site has WireGuard VPN configured. | [optional] 

## Methods

### NewExistSiteSettingOpenApiVO

`func NewExistSiteSettingOpenApiVO() *ExistSiteSettingOpenApiVO`

NewExistSiteSettingOpenApiVO instantiates a new ExistSiteSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExistSiteSettingOpenApiVOWithDefaults

`func NewExistSiteSettingOpenApiVOWithDefaults() *ExistSiteSettingOpenApiVO`

NewExistSiteSettingOpenApiVOWithDefaults instantiates a new ExistSiteSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableNat

`func (o *ExistSiteSettingOpenApiVO) GetDisableNat() bool`

GetDisableNat returns the DisableNat field if non-nil, zero value otherwise.

### GetDisableNatOk

`func (o *ExistSiteSettingOpenApiVO) GetDisableNatOk() (*bool, bool)`

GetDisableNatOk returns a tuple with the DisableNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNat

`func (o *ExistSiteSettingOpenApiVO) SetDisableNat(v bool)`

SetDisableNat sets DisableNat field to given value.

### HasDisableNat

`func (o *ExistSiteSettingOpenApiVO) HasDisableNat() bool`

HasDisableNat returns a boolean if a field has been set.

### GetDnsCache

`func (o *ExistSiteSettingOpenApiVO) GetDnsCache() bool`

GetDnsCache returns the DnsCache field if non-nil, zero value otherwise.

### GetDnsCacheOk

`func (o *ExistSiteSettingOpenApiVO) GetDnsCacheOk() (*bool, bool)`

GetDnsCacheOk returns a tuple with the DnsCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCache

`func (o *ExistSiteSettingOpenApiVO) SetDnsCache(v bool)`

SetDnsCache sets DnsCache field to given value.

### HasDnsCache

`func (o *ExistSiteSettingOpenApiVO) HasDnsCache() bool`

HasDnsCache returns a boolean if a field has been set.

### GetDnsProxy

`func (o *ExistSiteSettingOpenApiVO) GetDnsProxy() bool`

GetDnsProxy returns the DnsProxy field if non-nil, zero value otherwise.

### GetDnsProxyOk

`func (o *ExistSiteSettingOpenApiVO) GetDnsProxyOk() (*bool, bool)`

GetDnsProxyOk returns a tuple with the DnsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProxy

`func (o *ExistSiteSettingOpenApiVO) SetDnsProxy(v bool)`

SetDnsProxy sets DnsProxy field to given value.

### HasDnsProxy

`func (o *ExistSiteSettingOpenApiVO) HasDnsProxy() bool`

HasDnsProxy returns a boolean if a field has been set.

### GetGoogleLdap

`func (o *ExistSiteSettingOpenApiVO) GetGoogleLdap() bool`

GetGoogleLdap returns the GoogleLdap field if non-nil, zero value otherwise.

### GetGoogleLdapOk

`func (o *ExistSiteSettingOpenApiVO) GetGoogleLdapOk() (*bool, bool)`

GetGoogleLdapOk returns a tuple with the GoogleLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleLdap

`func (o *ExistSiteSettingOpenApiVO) SetGoogleLdap(v bool)`

SetGoogleLdap sets GoogleLdap field to given value.

### HasGoogleLdap

`func (o *ExistSiteSettingOpenApiVO) HasGoogleLdap() bool`

HasGoogleLdap returns a boolean if a field has been set.

### GetIpMacBinding

`func (o *ExistSiteSettingOpenApiVO) GetIpMacBinding() bool`

GetIpMacBinding returns the IpMacBinding field if non-nil, zero value otherwise.

### GetIpMacBindingOk

`func (o *ExistSiteSettingOpenApiVO) GetIpMacBindingOk() (*bool, bool)`

GetIpMacBindingOk returns a tuple with the IpMacBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMacBinding

`func (o *ExistSiteSettingOpenApiVO) SetIpMacBinding(v bool)`

SetIpMacBinding sets IpMacBinding field to given value.

### HasIpMacBinding

`func (o *ExistSiteSettingOpenApiVO) HasIpMacBinding() bool`

HasIpMacBinding returns a boolean if a field has been set.

### GetIpsIds

`func (o *ExistSiteSettingOpenApiVO) GetIpsIds() bool`

GetIpsIds returns the IpsIds field if non-nil, zero value otherwise.

### GetIpsIdsOk

`func (o *ExistSiteSettingOpenApiVO) GetIpsIdsOk() (*bool, bool)`

GetIpsIdsOk returns a tuple with the IpsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsIds

`func (o *ExistSiteSettingOpenApiVO) SetIpsIds(v bool)`

SetIpsIds sets IpsIds field to given value.

### HasIpsIds

`func (o *ExistSiteSettingOpenApiVO) HasIpsIds() bool`

HasIpsIds returns a boolean if a field has been set.

### GetIpsecFailover

`func (o *ExistSiteSettingOpenApiVO) GetIpsecFailover() bool`

GetIpsecFailover returns the IpsecFailover field if non-nil, zero value otherwise.

### GetIpsecFailoverOk

`func (o *ExistSiteSettingOpenApiVO) GetIpsecFailoverOk() (*bool, bool)`

GetIpsecFailoverOk returns a tuple with the IpsecFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecFailover

`func (o *ExistSiteSettingOpenApiVO) SetIpsecFailover(v bool)`

SetIpsecFailover sets IpsecFailover field to given value.

### HasIpsecFailover

`func (o *ExistSiteSettingOpenApiVO) HasIpsecFailover() bool`

HasIpsecFailover returns a boolean if a field has been set.

### GetLanDns

`func (o *ExistSiteSettingOpenApiVO) GetLanDns() bool`

GetLanDns returns the LanDns field if non-nil, zero value otherwise.

### GetLanDnsOk

`func (o *ExistSiteSettingOpenApiVO) GetLanDnsOk() (*bool, bool)`

GetLanDnsOk returns a tuple with the LanDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanDns

`func (o *ExistSiteSettingOpenApiVO) SetLanDns(v bool)`

SetLanDns sets LanDns field to given value.

### HasLanDns

`func (o *ExistSiteSettingOpenApiVO) HasLanDns() bool`

HasLanDns returns a boolean if a field has been set.

### GetMacFilter

`func (o *ExistSiteSettingOpenApiVO) GetMacFilter() bool`

GetMacFilter returns the MacFilter field if non-nil, zero value otherwise.

### GetMacFilterOk

`func (o *ExistSiteSettingOpenApiVO) GetMacFilterOk() (*bool, bool)`

GetMacFilterOk returns a tuple with the MacFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilter

`func (o *ExistSiteSettingOpenApiVO) SetMacFilter(v bool)`

SetMacFilter sets MacFilter field to given value.

### HasMacFilter

`func (o *ExistSiteSettingOpenApiVO) HasMacFilter() bool`

HasMacFilter returns a boolean if a field has been set.

### GetOneToOneNat

`func (o *ExistSiteSettingOpenApiVO) GetOneToOneNat() bool`

GetOneToOneNat returns the OneToOneNat field if non-nil, zero value otherwise.

### GetOneToOneNatOk

`func (o *ExistSiteSettingOpenApiVO) GetOneToOneNatOk() (*bool, bool)`

GetOneToOneNatOk returns a tuple with the OneToOneNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneToOneNat

`func (o *ExistSiteSettingOpenApiVO) SetOneToOneNat(v bool)`

SetOneToOneNat sets OneToOneNat field to given value.

### HasOneToOneNat

`func (o *ExistSiteSettingOpenApiVO) HasOneToOneNat() bool`

HasOneToOneNat returns a boolean if a field has been set.

### GetPolicyRouting

`func (o *ExistSiteSettingOpenApiVO) GetPolicyRouting() bool`

GetPolicyRouting returns the PolicyRouting field if non-nil, zero value otherwise.

### GetPolicyRoutingOk

`func (o *ExistSiteSettingOpenApiVO) GetPolicyRoutingOk() (*bool, bool)`

GetPolicyRoutingOk returns a tuple with the PolicyRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRouting

`func (o *ExistSiteSettingOpenApiVO) SetPolicyRouting(v bool)`

SetPolicyRouting sets PolicyRouting field to given value.

### HasPolicyRouting

`func (o *ExistSiteSettingOpenApiVO) HasPolicyRouting() bool`

HasPolicyRouting returns a boolean if a field has been set.

### GetQos

`func (o *ExistSiteSettingOpenApiVO) GetQos() bool`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *ExistSiteSettingOpenApiVO) GetQosOk() (*bool, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *ExistSiteSettingOpenApiVO) SetQos(v bool)`

SetQos sets Qos field to given value.

### HasQos

`func (o *ExistSiteSettingOpenApiVO) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetServiceType

`func (o *ExistSiteSettingOpenApiVO) GetServiceType() bool`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ExistSiteSettingOpenApiVO) GetServiceTypeOk() (*bool, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ExistSiteSettingOpenApiVO) SetServiceType(v bool)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ExistSiteSettingOpenApiVO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSslVpn

`func (o *ExistSiteSettingOpenApiVO) GetSslVpn() bool`

GetSslVpn returns the SslVpn field if non-nil, zero value otherwise.

### GetSslVpnOk

`func (o *ExistSiteSettingOpenApiVO) GetSslVpnOk() (*bool, bool)`

GetSslVpnOk returns a tuple with the SslVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVpn

`func (o *ExistSiteSettingOpenApiVO) SetSslVpn(v bool)`

SetSslVpn sets SslVpn field to given value.

### HasSslVpn

`func (o *ExistSiteSettingOpenApiVO) HasSslVpn() bool`

HasSslVpn returns a boolean if a field has been set.

### GetUrlCategory

`func (o *ExistSiteSettingOpenApiVO) GetUrlCategory() bool`

GetUrlCategory returns the UrlCategory field if non-nil, zero value otherwise.

### GetUrlCategoryOk

`func (o *ExistSiteSettingOpenApiVO) GetUrlCategoryOk() (*bool, bool)`

GetUrlCategoryOk returns a tuple with the UrlCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCategory

`func (o *ExistSiteSettingOpenApiVO) SetUrlCategory(v bool)`

SetUrlCategory sets UrlCategory field to given value.

### HasUrlCategory

`func (o *ExistSiteSettingOpenApiVO) HasUrlCategory() bool`

HasUrlCategory returns a boolean if a field has been set.

### GetVirtualWan

`func (o *ExistSiteSettingOpenApiVO) GetVirtualWan() bool`

GetVirtualWan returns the VirtualWan field if non-nil, zero value otherwise.

### GetVirtualWanOk

`func (o *ExistSiteSettingOpenApiVO) GetVirtualWanOk() (*bool, bool)`

GetVirtualWanOk returns a tuple with the VirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWan

`func (o *ExistSiteSettingOpenApiVO) SetVirtualWan(v bool)`

SetVirtualWan sets VirtualWan field to given value.

### HasVirtualWan

`func (o *ExistSiteSettingOpenApiVO) HasVirtualWan() bool`

HasVirtualWan returns a boolean if a field has been set.

### GetVpnUser

`func (o *ExistSiteSettingOpenApiVO) GetVpnUser() bool`

GetVpnUser returns the VpnUser field if non-nil, zero value otherwise.

### GetVpnUserOk

`func (o *ExistSiteSettingOpenApiVO) GetVpnUserOk() (*bool, bool)`

GetVpnUserOk returns a tuple with the VpnUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnUser

`func (o *ExistSiteSettingOpenApiVO) SetVpnUser(v bool)`

SetVpnUser sets VpnUser field to given value.

### HasVpnUser

`func (o *ExistSiteSettingOpenApiVO) HasVpnUser() bool`

HasVpnUser returns a boolean if a field has been set.

### GetWireguard

`func (o *ExistSiteSettingOpenApiVO) GetWireguard() bool`

GetWireguard returns the Wireguard field if non-nil, zero value otherwise.

### GetWireguardOk

`func (o *ExistSiteSettingOpenApiVO) GetWireguardOk() (*bool, bool)`

GetWireguardOk returns a tuple with the Wireguard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguard

`func (o *ExistSiteSettingOpenApiVO) SetWireguard(v bool)`

SetWireguard sets Wireguard field to given value.

### HasWireguard

`func (o *ExistSiteSettingOpenApiVO) HasWireguard() bool`

HasWireguard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


