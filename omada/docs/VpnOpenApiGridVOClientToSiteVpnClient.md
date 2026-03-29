# VpnOpenApiGridVOClientToSiteVpnClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]ClientToSiteVpnClient**](ClientToSiteVpnClient.md) |  | [optional] 
**PfsCap** | Pointer to **[]int32** | All PFS configuration types supported by the device, 1:dh1; 2:dh2; 3:dh3; 14:dh14; 15:dh15. Default: dh1、dh2、dh5 | [optional] 
**Phase1Proposal1Cap** | Pointer to **[]int32** | Proposal options for IKE negotiation phase-1. 0: MD5, 1: SHA1, 2:SHA256, 3:SHA384, 4:SHA512. Default: MD5, SHA1 and SHA256 | [optional] 
**Phase2Proposal2Cap** | Pointer to **[]int32** | Proposal options for IKE negotiation phase-2. 0: MD5, 1: SHA1, 2:SHA256, 3:SHA384, 4:SHA512.+ Default: MD5, SHA1 and SHA256 | [optional] 
**SupportAccountAuth** | Pointer to **bool** | Whether account auth configuration is supported of the VPN | [optional] 
**SupportCustomDns** | Pointer to **bool** | Whether custom dns configuration is supported of the VPN | [optional] 
**SupportCustomNetwork** | Pointer to **bool** | Whether custom network configuration is supported of the VPN | [optional] 
**SupportIPsec** | Pointer to **bool** | Whether IPsec VPN is supported. | [optional] 
**SupportIpRange** | Pointer to **bool** | Whether ip range configuration is supported of the VPN | [optional] 
**SupportIpSNat** | Pointer to **bool** | Whether IPSec Source NAT is supported of the VPN | [optional] 
**SupportL2TP** | Pointer to **bool** | Whether L2TP VPN is supported. | [optional] 
**SupportLdap** | Pointer to **bool** | Whether LDAP authentication mode is supported of the VPN client | [optional] 
**SupportOpenVpnDomain** | Pointer to **bool** | Whether OpenVPN supports inputting domain | [optional] 
**SupportTunnelMode** | Pointer to **bool** | Whether tunnel mode configuration is supported of the VPN | [optional] 
**SupportVpnUserTab** | Pointer to **bool** | Whether user tab configuration is supported of the VPN | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewVpnOpenApiGridVOClientToSiteVpnClient

`func NewVpnOpenApiGridVOClientToSiteVpnClient() *VpnOpenApiGridVOClientToSiteVpnClient`

NewVpnOpenApiGridVOClientToSiteVpnClient instantiates a new VpnOpenApiGridVOClientToSiteVpnClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnOpenApiGridVOClientToSiteVpnClientWithDefaults

`func NewVpnOpenApiGridVOClientToSiteVpnClientWithDefaults() *VpnOpenApiGridVOClientToSiteVpnClient`

NewVpnOpenApiGridVOClientToSiteVpnClientWithDefaults instantiates a new VpnOpenApiGridVOClientToSiteVpnClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetData() []ClientToSiteVpnClient`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetDataOk() (*[]ClientToSiteVpnClient, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetData(v []ClientToSiteVpnClient)`

SetData sets Data field to given value.

### HasData

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPfsCap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetPfsCap() []int32`

GetPfsCap returns the PfsCap field if non-nil, zero value otherwise.

### GetPfsCapOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetPfsCapOk() (*[]int32, bool)`

GetPfsCapOk returns a tuple with the PfsCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsCap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetPfsCap(v []int32)`

SetPfsCap sets PfsCap field to given value.

### HasPfsCap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasPfsCap() bool`

HasPfsCap returns a boolean if a field has been set.

### GetPhase1Proposal1Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetPhase1Proposal1Cap() []int32`

GetPhase1Proposal1Cap returns the Phase1Proposal1Cap field if non-nil, zero value otherwise.

### GetPhase1Proposal1CapOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetPhase1Proposal1CapOk() (*[]int32, bool)`

GetPhase1Proposal1CapOk returns a tuple with the Phase1Proposal1Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Proposal1Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetPhase1Proposal1Cap(v []int32)`

SetPhase1Proposal1Cap sets Phase1Proposal1Cap field to given value.

### HasPhase1Proposal1Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasPhase1Proposal1Cap() bool`

HasPhase1Proposal1Cap returns a boolean if a field has been set.

### GetPhase2Proposal2Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetPhase2Proposal2Cap() []int32`

GetPhase2Proposal2Cap returns the Phase2Proposal2Cap field if non-nil, zero value otherwise.

### GetPhase2Proposal2CapOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetPhase2Proposal2CapOk() (*[]int32, bool)`

GetPhase2Proposal2CapOk returns a tuple with the Phase2Proposal2Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Proposal2Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetPhase2Proposal2Cap(v []int32)`

SetPhase2Proposal2Cap sets Phase2Proposal2Cap field to given value.

### HasPhase2Proposal2Cap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasPhase2Proposal2Cap() bool`

HasPhase2Proposal2Cap returns a boolean if a field has been set.

### GetSupportAccountAuth

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportAccountAuth() bool`

GetSupportAccountAuth returns the SupportAccountAuth field if non-nil, zero value otherwise.

### GetSupportAccountAuthOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportAccountAuthOk() (*bool, bool)`

GetSupportAccountAuthOk returns a tuple with the SupportAccountAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccountAuth

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportAccountAuth(v bool)`

SetSupportAccountAuth sets SupportAccountAuth field to given value.

### HasSupportAccountAuth

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportAccountAuth() bool`

HasSupportAccountAuth returns a boolean if a field has been set.

### GetSupportCustomDns

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportCustomDns() bool`

GetSupportCustomDns returns the SupportCustomDns field if non-nil, zero value otherwise.

### GetSupportCustomDnsOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportCustomDnsOk() (*bool, bool)`

GetSupportCustomDnsOk returns a tuple with the SupportCustomDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDns

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportCustomDns(v bool)`

SetSupportCustomDns sets SupportCustomDns field to given value.

### HasSupportCustomDns

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportCustomDns() bool`

HasSupportCustomDns returns a boolean if a field has been set.

### GetSupportCustomNetwork

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportCustomNetwork() bool`

GetSupportCustomNetwork returns the SupportCustomNetwork field if non-nil, zero value otherwise.

### GetSupportCustomNetworkOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportCustomNetworkOk() (*bool, bool)`

GetSupportCustomNetworkOk returns a tuple with the SupportCustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomNetwork

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportCustomNetwork(v bool)`

SetSupportCustomNetwork sets SupportCustomNetwork field to given value.

### HasSupportCustomNetwork

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportCustomNetwork() bool`

HasSupportCustomNetwork returns a boolean if a field has been set.

### GetSupportIPsec

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportIPsec() bool`

GetSupportIPsec returns the SupportIPsec field if non-nil, zero value otherwise.

### GetSupportIPsecOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportIPsecOk() (*bool, bool)`

GetSupportIPsecOk returns a tuple with the SupportIPsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIPsec

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportIPsec(v bool)`

SetSupportIPsec sets SupportIPsec field to given value.

### HasSupportIPsec

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportIPsec() bool`

HasSupportIPsec returns a boolean if a field has been set.

### GetSupportIpRange

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportIpRange() bool`

GetSupportIpRange returns the SupportIpRange field if non-nil, zero value otherwise.

### GetSupportIpRangeOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportIpRangeOk() (*bool, bool)`

GetSupportIpRangeOk returns a tuple with the SupportIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpRange

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportIpRange(v bool)`

SetSupportIpRange sets SupportIpRange field to given value.

### HasSupportIpRange

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportIpRange() bool`

HasSupportIpRange returns a boolean if a field has been set.

### GetSupportIpSNat

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportIpSNat() bool`

GetSupportIpSNat returns the SupportIpSNat field if non-nil, zero value otherwise.

### GetSupportIpSNatOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportIpSNatOk() (*bool, bool)`

GetSupportIpSNatOk returns a tuple with the SupportIpSNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpSNat

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportIpSNat(v bool)`

SetSupportIpSNat sets SupportIpSNat field to given value.

### HasSupportIpSNat

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportIpSNat() bool`

HasSupportIpSNat returns a boolean if a field has been set.

### GetSupportL2TP

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportL2TP() bool`

GetSupportL2TP returns the SupportL2TP field if non-nil, zero value otherwise.

### GetSupportL2TPOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportL2TPOk() (*bool, bool)`

GetSupportL2TPOk returns a tuple with the SupportL2TP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2TP

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportL2TP(v bool)`

SetSupportL2TP sets SupportL2TP field to given value.

### HasSupportL2TP

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportL2TP() bool`

HasSupportL2TP returns a boolean if a field has been set.

### GetSupportLdap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportLdap() bool`

GetSupportLdap returns the SupportLdap field if non-nil, zero value otherwise.

### GetSupportLdapOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportLdapOk() (*bool, bool)`

GetSupportLdapOk returns a tuple with the SupportLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLdap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportLdap(v bool)`

SetSupportLdap sets SupportLdap field to given value.

### HasSupportLdap

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportLdap() bool`

HasSupportLdap returns a boolean if a field has been set.

### GetSupportOpenVpnDomain

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportOpenVpnDomain() bool`

GetSupportOpenVpnDomain returns the SupportOpenVpnDomain field if non-nil, zero value otherwise.

### GetSupportOpenVpnDomainOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportOpenVpnDomainOk() (*bool, bool)`

GetSupportOpenVpnDomainOk returns a tuple with the SupportOpenVpnDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOpenVpnDomain

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportOpenVpnDomain(v bool)`

SetSupportOpenVpnDomain sets SupportOpenVpnDomain field to given value.

### HasSupportOpenVpnDomain

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportOpenVpnDomain() bool`

HasSupportOpenVpnDomain returns a boolean if a field has been set.

### GetSupportTunnelMode

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportTunnelMode() bool`

GetSupportTunnelMode returns the SupportTunnelMode field if non-nil, zero value otherwise.

### GetSupportTunnelModeOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportTunnelModeOk() (*bool, bool)`

GetSupportTunnelModeOk returns a tuple with the SupportTunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTunnelMode

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportTunnelMode(v bool)`

SetSupportTunnelMode sets SupportTunnelMode field to given value.

### HasSupportTunnelMode

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportTunnelMode() bool`

HasSupportTunnelMode returns a boolean if a field has been set.

### GetSupportVpnUserTab

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportVpnUserTab() bool`

GetSupportVpnUserTab returns the SupportVpnUserTab field if non-nil, zero value otherwise.

### GetSupportVpnUserTabOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetSupportVpnUserTabOk() (*bool, bool)`

GetSupportVpnUserTabOk returns a tuple with the SupportVpnUserTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVpnUserTab

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetSupportVpnUserTab(v bool)`

SetSupportVpnUserTab sets SupportVpnUserTab field to given value.

### HasSupportVpnUserTab

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasSupportVpnUserTab() bool`

HasSupportVpnUserTab returns a boolean if a field has been set.

### GetTotalRows

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *VpnOpenApiGridVOClientToSiteVpnClient) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


