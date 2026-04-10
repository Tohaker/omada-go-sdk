# VpnOpenApiGridVOClientToSiteVpnServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]ClientToSiteVpnServer**](ClientToSiteVpnServer.md) |  | [optional] 
**PfsCap** | Pointer to **[]int32** | All PFS configuration types supported by the device, 1:dh1; 2:dh2; 3:dh3; 14:dh14; 15:dh15. Default: dh1、dh2、dh5 | [optional] 
**Phase1Proposal1Cap** | Pointer to **[]int32** | Proposal options for IKE negotiation phase-1. 0: MD5, 1: SHA1, 2:SHA256, 3:SHA384, 4:SHA512. Default: MD5, SHA1 and SHA256 | [optional] 
**Phase2Proposal2Cap** | Pointer to **[]int32** | Proposal options for IKE negotiation phase-2. 0: MD5, 1: SHA1, 2:SHA256, 3:SHA384, 4:SHA512.+ Default: MD5, SHA1 and SHA256 | [optional] 
**SubnetsLimitSize** | Pointer to **int32** | The limit on the number of entries for remote subnets and local network. | [optional] 
**SupportAccountAuth** | Pointer to **bool** | Whether account auth configuration is supported of the VPN | [optional] 
**SupportByDsLiteAndMapE** | Pointer to **bool** | Whether this feature is supported for the DS-Lite or Map-E WAN connection types. | [optional] 
**SupportCustomDns** | Pointer to **bool** | Whether custom dns configuration is supported of the VPN | [optional] 
**SupportCustomNetwork** | Pointer to **bool** | Whether custom network configuration is supported of the VPN | [optional] 
**SupportIPsec** | Pointer to **bool** | Whether IPsec VPN is supported. | [optional] 
**SupportIpRange** | Pointer to **bool** | Whether ip range configuration is supported of the VPN | [optional] 
**SupportL2TP** | Pointer to **bool** | Whether L2TP VPN is supported. | [optional] 
**SupportLdap** | Pointer to **bool** | Whether LDAP authentication mode is supported of the VPN server | [optional] 
**SupportOpenVpnDomain** | Pointer to **bool** | Whether OpenVPN supports inputting domain | [optional] 
**SupportTunnelMode** | Pointer to **bool** | Whether tunnel mode configuration is supported of the VPN | [optional] 
**SupportVpnUserTab** | Pointer to **bool** | Whether user tab configuration is supported of the VPN | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewVpnOpenApiGridVOClientToSiteVpnServer

`func NewVpnOpenApiGridVOClientToSiteVpnServer() *VpnOpenApiGridVOClientToSiteVpnServer`

NewVpnOpenApiGridVOClientToSiteVpnServer instantiates a new VpnOpenApiGridVOClientToSiteVpnServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnOpenApiGridVOClientToSiteVpnServerWithDefaults

`func NewVpnOpenApiGridVOClientToSiteVpnServerWithDefaults() *VpnOpenApiGridVOClientToSiteVpnServer`

NewVpnOpenApiGridVOClientToSiteVpnServerWithDefaults instantiates a new VpnOpenApiGridVOClientToSiteVpnServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetData() []ClientToSiteVpnServer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetDataOk() (*[]ClientToSiteVpnServer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetData(v []ClientToSiteVpnServer)`

SetData sets Data field to given value.

### HasData

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPfsCap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetPfsCap() []int32`

GetPfsCap returns the PfsCap field if non-nil, zero value otherwise.

### GetPfsCapOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetPfsCapOk() (*[]int32, bool)`

GetPfsCapOk returns a tuple with the PfsCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsCap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetPfsCap(v []int32)`

SetPfsCap sets PfsCap field to given value.

### HasPfsCap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasPfsCap() bool`

HasPfsCap returns a boolean if a field has been set.

### GetPhase1Proposal1Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetPhase1Proposal1Cap() []int32`

GetPhase1Proposal1Cap returns the Phase1Proposal1Cap field if non-nil, zero value otherwise.

### GetPhase1Proposal1CapOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetPhase1Proposal1CapOk() (*[]int32, bool)`

GetPhase1Proposal1CapOk returns a tuple with the Phase1Proposal1Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Proposal1Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetPhase1Proposal1Cap(v []int32)`

SetPhase1Proposal1Cap sets Phase1Proposal1Cap field to given value.

### HasPhase1Proposal1Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasPhase1Proposal1Cap() bool`

HasPhase1Proposal1Cap returns a boolean if a field has been set.

### GetPhase2Proposal2Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetPhase2Proposal2Cap() []int32`

GetPhase2Proposal2Cap returns the Phase2Proposal2Cap field if non-nil, zero value otherwise.

### GetPhase2Proposal2CapOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetPhase2Proposal2CapOk() (*[]int32, bool)`

GetPhase2Proposal2CapOk returns a tuple with the Phase2Proposal2Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Proposal2Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetPhase2Proposal2Cap(v []int32)`

SetPhase2Proposal2Cap sets Phase2Proposal2Cap field to given value.

### HasPhase2Proposal2Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasPhase2Proposal2Cap() bool`

HasPhase2Proposal2Cap returns a boolean if a field has been set.

### GetSubnetsLimitSize

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSubnetsLimitSize() int32`

GetSubnetsLimitSize returns the SubnetsLimitSize field if non-nil, zero value otherwise.

### GetSubnetsLimitSizeOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSubnetsLimitSizeOk() (*int32, bool)`

GetSubnetsLimitSizeOk returns a tuple with the SubnetsLimitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsLimitSize

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSubnetsLimitSize(v int32)`

SetSubnetsLimitSize sets SubnetsLimitSize field to given value.

### HasSubnetsLimitSize

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSubnetsLimitSize() bool`

HasSubnetsLimitSize returns a boolean if a field has been set.

### GetSupportAccountAuth

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportAccountAuth() bool`

GetSupportAccountAuth returns the SupportAccountAuth field if non-nil, zero value otherwise.

### GetSupportAccountAuthOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportAccountAuthOk() (*bool, bool)`

GetSupportAccountAuthOk returns a tuple with the SupportAccountAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccountAuth

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportAccountAuth(v bool)`

SetSupportAccountAuth sets SupportAccountAuth field to given value.

### HasSupportAccountAuth

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportAccountAuth() bool`

HasSupportAccountAuth returns a boolean if a field has been set.

### GetSupportByDsLiteAndMapE

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportByDsLiteAndMapE() bool`

GetSupportByDsLiteAndMapE returns the SupportByDsLiteAndMapE field if non-nil, zero value otherwise.

### GetSupportByDsLiteAndMapEOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportByDsLiteAndMapEOk() (*bool, bool)`

GetSupportByDsLiteAndMapEOk returns a tuple with the SupportByDsLiteAndMapE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportByDsLiteAndMapE

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportByDsLiteAndMapE(v bool)`

SetSupportByDsLiteAndMapE sets SupportByDsLiteAndMapE field to given value.

### HasSupportByDsLiteAndMapE

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportByDsLiteAndMapE() bool`

HasSupportByDsLiteAndMapE returns a boolean if a field has been set.

### GetSupportCustomDns

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportCustomDns() bool`

GetSupportCustomDns returns the SupportCustomDns field if non-nil, zero value otherwise.

### GetSupportCustomDnsOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportCustomDnsOk() (*bool, bool)`

GetSupportCustomDnsOk returns a tuple with the SupportCustomDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDns

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportCustomDns(v bool)`

SetSupportCustomDns sets SupportCustomDns field to given value.

### HasSupportCustomDns

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportCustomDns() bool`

HasSupportCustomDns returns a boolean if a field has been set.

### GetSupportCustomNetwork

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportCustomNetwork() bool`

GetSupportCustomNetwork returns the SupportCustomNetwork field if non-nil, zero value otherwise.

### GetSupportCustomNetworkOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportCustomNetworkOk() (*bool, bool)`

GetSupportCustomNetworkOk returns a tuple with the SupportCustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomNetwork

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportCustomNetwork(v bool)`

SetSupportCustomNetwork sets SupportCustomNetwork field to given value.

### HasSupportCustomNetwork

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportCustomNetwork() bool`

HasSupportCustomNetwork returns a boolean if a field has been set.

### GetSupportIPsec

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportIPsec() bool`

GetSupportIPsec returns the SupportIPsec field if non-nil, zero value otherwise.

### GetSupportIPsecOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportIPsecOk() (*bool, bool)`

GetSupportIPsecOk returns a tuple with the SupportIPsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIPsec

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportIPsec(v bool)`

SetSupportIPsec sets SupportIPsec field to given value.

### HasSupportIPsec

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportIPsec() bool`

HasSupportIPsec returns a boolean if a field has been set.

### GetSupportIpRange

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportIpRange() bool`

GetSupportIpRange returns the SupportIpRange field if non-nil, zero value otherwise.

### GetSupportIpRangeOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportIpRangeOk() (*bool, bool)`

GetSupportIpRangeOk returns a tuple with the SupportIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpRange

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportIpRange(v bool)`

SetSupportIpRange sets SupportIpRange field to given value.

### HasSupportIpRange

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportIpRange() bool`

HasSupportIpRange returns a boolean if a field has been set.

### GetSupportL2TP

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportL2TP() bool`

GetSupportL2TP returns the SupportL2TP field if non-nil, zero value otherwise.

### GetSupportL2TPOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportL2TPOk() (*bool, bool)`

GetSupportL2TPOk returns a tuple with the SupportL2TP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2TP

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportL2TP(v bool)`

SetSupportL2TP sets SupportL2TP field to given value.

### HasSupportL2TP

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportL2TP() bool`

HasSupportL2TP returns a boolean if a field has been set.

### GetSupportLdap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportLdap() bool`

GetSupportLdap returns the SupportLdap field if non-nil, zero value otherwise.

### GetSupportLdapOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportLdapOk() (*bool, bool)`

GetSupportLdapOk returns a tuple with the SupportLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLdap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportLdap(v bool)`

SetSupportLdap sets SupportLdap field to given value.

### HasSupportLdap

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportLdap() bool`

HasSupportLdap returns a boolean if a field has been set.

### GetSupportOpenVpnDomain

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportOpenVpnDomain() bool`

GetSupportOpenVpnDomain returns the SupportOpenVpnDomain field if non-nil, zero value otherwise.

### GetSupportOpenVpnDomainOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportOpenVpnDomainOk() (*bool, bool)`

GetSupportOpenVpnDomainOk returns a tuple with the SupportOpenVpnDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOpenVpnDomain

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportOpenVpnDomain(v bool)`

SetSupportOpenVpnDomain sets SupportOpenVpnDomain field to given value.

### HasSupportOpenVpnDomain

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportOpenVpnDomain() bool`

HasSupportOpenVpnDomain returns a boolean if a field has been set.

### GetSupportTunnelMode

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportTunnelMode() bool`

GetSupportTunnelMode returns the SupportTunnelMode field if non-nil, zero value otherwise.

### GetSupportTunnelModeOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportTunnelModeOk() (*bool, bool)`

GetSupportTunnelModeOk returns a tuple with the SupportTunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTunnelMode

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportTunnelMode(v bool)`

SetSupportTunnelMode sets SupportTunnelMode field to given value.

### HasSupportTunnelMode

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportTunnelMode() bool`

HasSupportTunnelMode returns a boolean if a field has been set.

### GetSupportVpnUserTab

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportVpnUserTab() bool`

GetSupportVpnUserTab returns the SupportVpnUserTab field if non-nil, zero value otherwise.

### GetSupportVpnUserTabOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetSupportVpnUserTabOk() (*bool, bool)`

GetSupportVpnUserTabOk returns a tuple with the SupportVpnUserTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVpnUserTab

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetSupportVpnUserTab(v bool)`

SetSupportVpnUserTab sets SupportVpnUserTab field to given value.

### HasSupportVpnUserTab

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasSupportVpnUserTab() bool`

HasSupportVpnUserTab returns a boolean if a field has been set.

### GetTotalRows

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *VpnOpenApiGridVOClientToSiteVpnServer) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


