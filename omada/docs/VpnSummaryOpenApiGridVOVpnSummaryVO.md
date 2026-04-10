# VpnSummaryOpenApiGridVOVpnSummaryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VpnSummaryVO**](VpnSummaryVO.md) |  | [optional] 
**PfsCap** | Pointer to **[]int32** | All PFS configuration types supported by the device, 1:dh1; 2:dh2; 3:dh3; 14:dh14; 15:dh15. Default: dh1、dh2、dh5 | [optional] 
**Phase1Proposal1Cap** | Pointer to **[]int32** | Proposal options for IKE negotiation phase-1. 0: MD5, 1: SHA1, 2:SHA256, 3:SHA384, 4:SHA512. Default: MD5, SHA1 and SHA256 | [optional] 
**Phase2Proposal2Cap** | Pointer to **[]int32** | Proposal options for IKE negotiation phase-2. 0: MD5, 1: SHA1, 2:SHA256, 3:SHA384, 4:SHA512. Default: MD5, SHA1 and SHA256 | [optional] 
**SslVpnNum** | Pointer to **int64** | The number of SSLVPN servers. | [optional] 
**SubnetsLimitSize** | Pointer to **int32** | The limit on the number of entries for remote subnets and local network. | [optional] 
**SupportAccountAuth** | Pointer to **bool** | Whether account auth configuration is supported of the VPN | [optional] 
**SupportAutoWireGuard** | Pointer to **bool** | Whether auto Wire Guard in site to site is supported. | [optional] 
**SupportByDsLiteAndMapE** | Pointer to **bool** | Whether this feature is supported for the DS-Lite or Map-E WAN connection types. | [optional] 
**SupportCustomDns** | Pointer to **bool** | Whether custom dns configuration is supported of the VPN | [optional] 
**SupportCustomNetwork** | Pointer to **bool** | Whether custom network configuration is supported of the VPN | [optional] 
**SupportCustomServer** | Pointer to **bool** | Whether custom server is supported. | [optional] 
**SupportDnsStatus** | Pointer to **bool** | Whether supports dns auto. | [optional] 
**SupportIPsec** | Pointer to **bool** | Whether IPsec VPN is supported. | [optional] 
**SupportIPsecFailover** | Pointer to **bool** | Whether supports IPSec failover in manual IPSec. | [optional] 
**SupportIPsecIpRange** | Pointer to **bool** | Whether IP range is supported of IPSec VPN Server. | [optional] 
**SupportIpRange** | Pointer to **bool** | Whether ip range configuration is supported of the VPN | [optional] 
**SupportL2TP** | Pointer to **bool** | Whether L2TP VPN is supported. | [optional] 
**SupportLdap** | Pointer to **bool** | Whether LDAP authentication mode is supported of the VPN server | [optional] 
**SupportManualWgLocalNetwork** | Pointer to **bool** | Whether manual Wire Guard in site to site supports local network. | [optional] 
**SupportOpenVpnDomain** | Pointer to **bool** | Whether OpenVPN supports inputting domain. | [optional] 
**SupportServerClientWireGuard** | Pointer to **bool** | Whether Server/Client Wire Guard is supported. | [optional] 
**SupportServerOpenVpnGoogleLdap** | Pointer to **bool** | Whether Open VPN Server supports Google LDAP. | [optional] 
**SupportSslVpnLdap** | Pointer to **bool** | Whether LDAP authentication mode is supported of the SSl VPN server | [optional] 
**SupportSslVpnRadius** | Pointer to **bool** | Whether radius authentication mode is supported of the SSl VPN server | [optional] 
**SupportTunnelMode** | Pointer to **bool** | Whether tunnel mode configuration is supported of the VPN | [optional] 
**SupportVpnUserTab** | Pointer to **bool** | Whether user tab configuration is supported of the VPN | [optional] 
**SupportWgDomain** | Pointer to **bool** | Whether peer supports inputting domain. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewVpnSummaryOpenApiGridVOVpnSummaryVO

`func NewVpnSummaryOpenApiGridVOVpnSummaryVO() *VpnSummaryOpenApiGridVOVpnSummaryVO`

NewVpnSummaryOpenApiGridVOVpnSummaryVO instantiates a new VpnSummaryOpenApiGridVOVpnSummaryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnSummaryOpenApiGridVOVpnSummaryVOWithDefaults

`func NewVpnSummaryOpenApiGridVOVpnSummaryVOWithDefaults() *VpnSummaryOpenApiGridVOVpnSummaryVO`

NewVpnSummaryOpenApiGridVOVpnSummaryVOWithDefaults instantiates a new VpnSummaryOpenApiGridVOVpnSummaryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetData() []VpnSummaryVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetDataOk() (*[]VpnSummaryVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetData(v []VpnSummaryVO)`

SetData sets Data field to given value.

### HasData

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPfsCap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetPfsCap() []int32`

GetPfsCap returns the PfsCap field if non-nil, zero value otherwise.

### GetPfsCapOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetPfsCapOk() (*[]int32, bool)`

GetPfsCapOk returns a tuple with the PfsCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsCap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetPfsCap(v []int32)`

SetPfsCap sets PfsCap field to given value.

### HasPfsCap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasPfsCap() bool`

HasPfsCap returns a boolean if a field has been set.

### GetPhase1Proposal1Cap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetPhase1Proposal1Cap() []int32`

GetPhase1Proposal1Cap returns the Phase1Proposal1Cap field if non-nil, zero value otherwise.

### GetPhase1Proposal1CapOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetPhase1Proposal1CapOk() (*[]int32, bool)`

GetPhase1Proposal1CapOk returns a tuple with the Phase1Proposal1Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Proposal1Cap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetPhase1Proposal1Cap(v []int32)`

SetPhase1Proposal1Cap sets Phase1Proposal1Cap field to given value.

### HasPhase1Proposal1Cap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasPhase1Proposal1Cap() bool`

HasPhase1Proposal1Cap returns a boolean if a field has been set.

### GetPhase2Proposal2Cap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetPhase2Proposal2Cap() []int32`

GetPhase2Proposal2Cap returns the Phase2Proposal2Cap field if non-nil, zero value otherwise.

### GetPhase2Proposal2CapOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetPhase2Proposal2CapOk() (*[]int32, bool)`

GetPhase2Proposal2CapOk returns a tuple with the Phase2Proposal2Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Proposal2Cap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetPhase2Proposal2Cap(v []int32)`

SetPhase2Proposal2Cap sets Phase2Proposal2Cap field to given value.

### HasPhase2Proposal2Cap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasPhase2Proposal2Cap() bool`

HasPhase2Proposal2Cap returns a boolean if a field has been set.

### GetSslVpnNum

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSslVpnNum() int64`

GetSslVpnNum returns the SslVpnNum field if non-nil, zero value otherwise.

### GetSslVpnNumOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSslVpnNumOk() (*int64, bool)`

GetSslVpnNumOk returns a tuple with the SslVpnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVpnNum

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSslVpnNum(v int64)`

SetSslVpnNum sets SslVpnNum field to given value.

### HasSslVpnNum

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSslVpnNum() bool`

HasSslVpnNum returns a boolean if a field has been set.

### GetSubnetsLimitSize

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSubnetsLimitSize() int32`

GetSubnetsLimitSize returns the SubnetsLimitSize field if non-nil, zero value otherwise.

### GetSubnetsLimitSizeOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSubnetsLimitSizeOk() (*int32, bool)`

GetSubnetsLimitSizeOk returns a tuple with the SubnetsLimitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsLimitSize

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSubnetsLimitSize(v int32)`

SetSubnetsLimitSize sets SubnetsLimitSize field to given value.

### HasSubnetsLimitSize

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSubnetsLimitSize() bool`

HasSubnetsLimitSize returns a boolean if a field has been set.

### GetSupportAccountAuth

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportAccountAuth() bool`

GetSupportAccountAuth returns the SupportAccountAuth field if non-nil, zero value otherwise.

### GetSupportAccountAuthOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportAccountAuthOk() (*bool, bool)`

GetSupportAccountAuthOk returns a tuple with the SupportAccountAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccountAuth

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportAccountAuth(v bool)`

SetSupportAccountAuth sets SupportAccountAuth field to given value.

### HasSupportAccountAuth

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportAccountAuth() bool`

HasSupportAccountAuth returns a boolean if a field has been set.

### GetSupportAutoWireGuard

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportAutoWireGuard() bool`

GetSupportAutoWireGuard returns the SupportAutoWireGuard field if non-nil, zero value otherwise.

### GetSupportAutoWireGuardOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportAutoWireGuardOk() (*bool, bool)`

GetSupportAutoWireGuardOk returns a tuple with the SupportAutoWireGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAutoWireGuard

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportAutoWireGuard(v bool)`

SetSupportAutoWireGuard sets SupportAutoWireGuard field to given value.

### HasSupportAutoWireGuard

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportAutoWireGuard() bool`

HasSupportAutoWireGuard returns a boolean if a field has been set.

### GetSupportByDsLiteAndMapE

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportByDsLiteAndMapE() bool`

GetSupportByDsLiteAndMapE returns the SupportByDsLiteAndMapE field if non-nil, zero value otherwise.

### GetSupportByDsLiteAndMapEOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportByDsLiteAndMapEOk() (*bool, bool)`

GetSupportByDsLiteAndMapEOk returns a tuple with the SupportByDsLiteAndMapE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportByDsLiteAndMapE

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportByDsLiteAndMapE(v bool)`

SetSupportByDsLiteAndMapE sets SupportByDsLiteAndMapE field to given value.

### HasSupportByDsLiteAndMapE

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportByDsLiteAndMapE() bool`

HasSupportByDsLiteAndMapE returns a boolean if a field has been set.

### GetSupportCustomDns

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportCustomDns() bool`

GetSupportCustomDns returns the SupportCustomDns field if non-nil, zero value otherwise.

### GetSupportCustomDnsOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportCustomDnsOk() (*bool, bool)`

GetSupportCustomDnsOk returns a tuple with the SupportCustomDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDns

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportCustomDns(v bool)`

SetSupportCustomDns sets SupportCustomDns field to given value.

### HasSupportCustomDns

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportCustomDns() bool`

HasSupportCustomDns returns a boolean if a field has been set.

### GetSupportCustomNetwork

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportCustomNetwork() bool`

GetSupportCustomNetwork returns the SupportCustomNetwork field if non-nil, zero value otherwise.

### GetSupportCustomNetworkOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportCustomNetworkOk() (*bool, bool)`

GetSupportCustomNetworkOk returns a tuple with the SupportCustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomNetwork

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportCustomNetwork(v bool)`

SetSupportCustomNetwork sets SupportCustomNetwork field to given value.

### HasSupportCustomNetwork

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportCustomNetwork() bool`

HasSupportCustomNetwork returns a boolean if a field has been set.

### GetSupportCustomServer

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportCustomServer() bool`

GetSupportCustomServer returns the SupportCustomServer field if non-nil, zero value otherwise.

### GetSupportCustomServerOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportCustomServerOk() (*bool, bool)`

GetSupportCustomServerOk returns a tuple with the SupportCustomServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomServer

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportCustomServer(v bool)`

SetSupportCustomServer sets SupportCustomServer field to given value.

### HasSupportCustomServer

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportCustomServer() bool`

HasSupportCustomServer returns a boolean if a field has been set.

### GetSupportDnsStatus

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportDnsStatus() bool`

GetSupportDnsStatus returns the SupportDnsStatus field if non-nil, zero value otherwise.

### GetSupportDnsStatusOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportDnsStatusOk() (*bool, bool)`

GetSupportDnsStatusOk returns a tuple with the SupportDnsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDnsStatus

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportDnsStatus(v bool)`

SetSupportDnsStatus sets SupportDnsStatus field to given value.

### HasSupportDnsStatus

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportDnsStatus() bool`

HasSupportDnsStatus returns a boolean if a field has been set.

### GetSupportIPsec

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportIPsec() bool`

GetSupportIPsec returns the SupportIPsec field if non-nil, zero value otherwise.

### GetSupportIPsecOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportIPsecOk() (*bool, bool)`

GetSupportIPsecOk returns a tuple with the SupportIPsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIPsec

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportIPsec(v bool)`

SetSupportIPsec sets SupportIPsec field to given value.

### HasSupportIPsec

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportIPsec() bool`

HasSupportIPsec returns a boolean if a field has been set.

### GetSupportIPsecFailover

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportIPsecFailover() bool`

GetSupportIPsecFailover returns the SupportIPsecFailover field if non-nil, zero value otherwise.

### GetSupportIPsecFailoverOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportIPsecFailoverOk() (*bool, bool)`

GetSupportIPsecFailoverOk returns a tuple with the SupportIPsecFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIPsecFailover

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportIPsecFailover(v bool)`

SetSupportIPsecFailover sets SupportIPsecFailover field to given value.

### HasSupportIPsecFailover

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportIPsecFailover() bool`

HasSupportIPsecFailover returns a boolean if a field has been set.

### GetSupportIPsecIpRange

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportIPsecIpRange() bool`

GetSupportIPsecIpRange returns the SupportIPsecIpRange field if non-nil, zero value otherwise.

### GetSupportIPsecIpRangeOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportIPsecIpRangeOk() (*bool, bool)`

GetSupportIPsecIpRangeOk returns a tuple with the SupportIPsecIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIPsecIpRange

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportIPsecIpRange(v bool)`

SetSupportIPsecIpRange sets SupportIPsecIpRange field to given value.

### HasSupportIPsecIpRange

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportIPsecIpRange() bool`

HasSupportIPsecIpRange returns a boolean if a field has been set.

### GetSupportIpRange

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportIpRange() bool`

GetSupportIpRange returns the SupportIpRange field if non-nil, zero value otherwise.

### GetSupportIpRangeOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportIpRangeOk() (*bool, bool)`

GetSupportIpRangeOk returns a tuple with the SupportIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpRange

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportIpRange(v bool)`

SetSupportIpRange sets SupportIpRange field to given value.

### HasSupportIpRange

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportIpRange() bool`

HasSupportIpRange returns a boolean if a field has been set.

### GetSupportL2TP

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportL2TP() bool`

GetSupportL2TP returns the SupportL2TP field if non-nil, zero value otherwise.

### GetSupportL2TPOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportL2TPOk() (*bool, bool)`

GetSupportL2TPOk returns a tuple with the SupportL2TP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2TP

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportL2TP(v bool)`

SetSupportL2TP sets SupportL2TP field to given value.

### HasSupportL2TP

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportL2TP() bool`

HasSupportL2TP returns a boolean if a field has been set.

### GetSupportLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportLdap() bool`

GetSupportLdap returns the SupportLdap field if non-nil, zero value otherwise.

### GetSupportLdapOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportLdapOk() (*bool, bool)`

GetSupportLdapOk returns a tuple with the SupportLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportLdap(v bool)`

SetSupportLdap sets SupportLdap field to given value.

### HasSupportLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportLdap() bool`

HasSupportLdap returns a boolean if a field has been set.

### GetSupportManualWgLocalNetwork

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportManualWgLocalNetwork() bool`

GetSupportManualWgLocalNetwork returns the SupportManualWgLocalNetwork field if non-nil, zero value otherwise.

### GetSupportManualWgLocalNetworkOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportManualWgLocalNetworkOk() (*bool, bool)`

GetSupportManualWgLocalNetworkOk returns a tuple with the SupportManualWgLocalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportManualWgLocalNetwork

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportManualWgLocalNetwork(v bool)`

SetSupportManualWgLocalNetwork sets SupportManualWgLocalNetwork field to given value.

### HasSupportManualWgLocalNetwork

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportManualWgLocalNetwork() bool`

HasSupportManualWgLocalNetwork returns a boolean if a field has been set.

### GetSupportOpenVpnDomain

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportOpenVpnDomain() bool`

GetSupportOpenVpnDomain returns the SupportOpenVpnDomain field if non-nil, zero value otherwise.

### GetSupportOpenVpnDomainOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportOpenVpnDomainOk() (*bool, bool)`

GetSupportOpenVpnDomainOk returns a tuple with the SupportOpenVpnDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOpenVpnDomain

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportOpenVpnDomain(v bool)`

SetSupportOpenVpnDomain sets SupportOpenVpnDomain field to given value.

### HasSupportOpenVpnDomain

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportOpenVpnDomain() bool`

HasSupportOpenVpnDomain returns a boolean if a field has been set.

### GetSupportServerClientWireGuard

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportServerClientWireGuard() bool`

GetSupportServerClientWireGuard returns the SupportServerClientWireGuard field if non-nil, zero value otherwise.

### GetSupportServerClientWireGuardOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportServerClientWireGuardOk() (*bool, bool)`

GetSupportServerClientWireGuardOk returns a tuple with the SupportServerClientWireGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportServerClientWireGuard

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportServerClientWireGuard(v bool)`

SetSupportServerClientWireGuard sets SupportServerClientWireGuard field to given value.

### HasSupportServerClientWireGuard

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportServerClientWireGuard() bool`

HasSupportServerClientWireGuard returns a boolean if a field has been set.

### GetSupportServerOpenVpnGoogleLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportServerOpenVpnGoogleLdap() bool`

GetSupportServerOpenVpnGoogleLdap returns the SupportServerOpenVpnGoogleLdap field if non-nil, zero value otherwise.

### GetSupportServerOpenVpnGoogleLdapOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportServerOpenVpnGoogleLdapOk() (*bool, bool)`

GetSupportServerOpenVpnGoogleLdapOk returns a tuple with the SupportServerOpenVpnGoogleLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportServerOpenVpnGoogleLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportServerOpenVpnGoogleLdap(v bool)`

SetSupportServerOpenVpnGoogleLdap sets SupportServerOpenVpnGoogleLdap field to given value.

### HasSupportServerOpenVpnGoogleLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportServerOpenVpnGoogleLdap() bool`

HasSupportServerOpenVpnGoogleLdap returns a boolean if a field has been set.

### GetSupportSslVpnLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportSslVpnLdap() bool`

GetSupportSslVpnLdap returns the SupportSslVpnLdap field if non-nil, zero value otherwise.

### GetSupportSslVpnLdapOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportSslVpnLdapOk() (*bool, bool)`

GetSupportSslVpnLdapOk returns a tuple with the SupportSslVpnLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSslVpnLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportSslVpnLdap(v bool)`

SetSupportSslVpnLdap sets SupportSslVpnLdap field to given value.

### HasSupportSslVpnLdap

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportSslVpnLdap() bool`

HasSupportSslVpnLdap returns a boolean if a field has been set.

### GetSupportSslVpnRadius

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportSslVpnRadius() bool`

GetSupportSslVpnRadius returns the SupportSslVpnRadius field if non-nil, zero value otherwise.

### GetSupportSslVpnRadiusOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportSslVpnRadiusOk() (*bool, bool)`

GetSupportSslVpnRadiusOk returns a tuple with the SupportSslVpnRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSslVpnRadius

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportSslVpnRadius(v bool)`

SetSupportSslVpnRadius sets SupportSslVpnRadius field to given value.

### HasSupportSslVpnRadius

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportSslVpnRadius() bool`

HasSupportSslVpnRadius returns a boolean if a field has been set.

### GetSupportTunnelMode

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportTunnelMode() bool`

GetSupportTunnelMode returns the SupportTunnelMode field if non-nil, zero value otherwise.

### GetSupportTunnelModeOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportTunnelModeOk() (*bool, bool)`

GetSupportTunnelModeOk returns a tuple with the SupportTunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTunnelMode

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportTunnelMode(v bool)`

SetSupportTunnelMode sets SupportTunnelMode field to given value.

### HasSupportTunnelMode

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportTunnelMode() bool`

HasSupportTunnelMode returns a boolean if a field has been set.

### GetSupportVpnUserTab

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportVpnUserTab() bool`

GetSupportVpnUserTab returns the SupportVpnUserTab field if non-nil, zero value otherwise.

### GetSupportVpnUserTabOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportVpnUserTabOk() (*bool, bool)`

GetSupportVpnUserTabOk returns a tuple with the SupportVpnUserTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVpnUserTab

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportVpnUserTab(v bool)`

SetSupportVpnUserTab sets SupportVpnUserTab field to given value.

### HasSupportVpnUserTab

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportVpnUserTab() bool`

HasSupportVpnUserTab returns a boolean if a field has been set.

### GetSupportWgDomain

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportWgDomain() bool`

GetSupportWgDomain returns the SupportWgDomain field if non-nil, zero value otherwise.

### GetSupportWgDomainOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetSupportWgDomainOk() (*bool, bool)`

GetSupportWgDomainOk returns a tuple with the SupportWgDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWgDomain

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetSupportWgDomain(v bool)`

SetSupportWgDomain sets SupportWgDomain field to given value.

### HasSupportWgDomain

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasSupportWgDomain() bool`

HasSupportWgDomain returns a boolean if a field has been set.

### GetTotalRows

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *VpnSummaryOpenApiGridVOVpnSummaryVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


