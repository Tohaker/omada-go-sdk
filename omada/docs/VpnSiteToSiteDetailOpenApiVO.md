# VpnSiteToSiteDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSetting** | Pointer to [**VpnAdvancedSettingOpenApiVO**](VpnAdvancedSettingOpenApiVO.md) |  | [optional] 
**CustomNetwork** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Custom networks of the VPN. | [optional] 
**FailoverSetting** | Pointer to [**IPsecFailoverSettingOpenApiVO**](IPsecFailoverSettingOpenApiVO.md) |  | [optional] 
**FeatureDescription** | Pointer to [**[]FeatureInfoVO**](FeatureInfoVO.md) | Gateway Feature Description. | [optional] 
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**Mtu** | Pointer to **int32** | The MTU of WireGuard VPN should be within the range of 576-1440. | [optional] 
**Name** | Pointer to **string** | VPN name. | [optional] 
**NetworkList** | Pointer to **[]string** | Network list of the VPN. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**NetworkType** | Pointer to **int32** | Network type should be a value as follows: 0: network list; 1: custom networks. | [optional] 
**Peers** | Pointer to [**[]SiteToSiteManualWgPeerConfigVO**](SiteToSiteManualWgPeerConfigVO.md) | List of Site-To-Site manual WireGuard peer Configuration. | [optional] 
**PreSharedKey** | Pointer to **string** | Pre-shared key of the VPN. | [optional] 
**PrivateKey** | Pointer to **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PublicKey** | Pointer to **string** | The public key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**RemoteIp** | Pointer to **string** | Remote IP of the VPN. | [optional] 
**RemoteSite** | Pointer to **string** | Remote site of the VPN. | [optional] 
**RemoteSiteName** | Pointer to **string** | Remote site name of the VPN. | [optional] 
**RemoteSubnet** | Pointer to **[]string** | Remote subnet of the VPN. | [optional] 
**ServicePort** | Pointer to **int32** | Service port for VPN. | [optional] 
**SiteVpnType** | Pointer to **int32** | Site VPN type of the VPN. 0: Auto; 1: Manual. | [optional] 
**Status** | Pointer to **bool** | Status of the VPN. | [optional] 
**TunnelIp** | Pointer to **string** | IP address. | [optional] 
**VpnType** | Pointer to **int32** | Server Vpn type. 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. | [optional] 
**Wans** | Pointer to **[]string** | WAN list of the VPN, only for Manual IPSec type. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | [optional] 

## Methods

### NewVpnSiteToSiteDetailOpenApiVO

`func NewVpnSiteToSiteDetailOpenApiVO() *VpnSiteToSiteDetailOpenApiVO`

NewVpnSiteToSiteDetailOpenApiVO instantiates a new VpnSiteToSiteDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnSiteToSiteDetailOpenApiVOWithDefaults

`func NewVpnSiteToSiteDetailOpenApiVOWithDefaults() *VpnSiteToSiteDetailOpenApiVO`

NewVpnSiteToSiteDetailOpenApiVOWithDefaults instantiates a new VpnSiteToSiteDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSetting

`func (o *VpnSiteToSiteDetailOpenApiVO) GetAdvancedSetting() VpnAdvancedSettingOpenApiVO`

GetAdvancedSetting returns the AdvancedSetting field if non-nil, zero value otherwise.

### GetAdvancedSettingOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetAdvancedSettingOk() (*VpnAdvancedSettingOpenApiVO, bool)`

GetAdvancedSettingOk returns a tuple with the AdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSetting

`func (o *VpnSiteToSiteDetailOpenApiVO) SetAdvancedSetting(v VpnAdvancedSettingOpenApiVO)`

SetAdvancedSetting sets AdvancedSetting field to given value.

### HasAdvancedSetting

`func (o *VpnSiteToSiteDetailOpenApiVO) HasAdvancedSetting() bool`

HasAdvancedSetting returns a boolean if a field has been set.

### GetCustomNetwork

`func (o *VpnSiteToSiteDetailOpenApiVO) GetCustomNetwork() []IPSubnetsVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetCustomNetworkOk() (*[]IPSubnetsVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *VpnSiteToSiteDetailOpenApiVO) SetCustomNetwork(v []IPSubnetsVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *VpnSiteToSiteDetailOpenApiVO) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetFailoverSetting

`func (o *VpnSiteToSiteDetailOpenApiVO) GetFailoverSetting() IPsecFailoverSettingOpenApiVO`

GetFailoverSetting returns the FailoverSetting field if non-nil, zero value otherwise.

### GetFailoverSettingOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetFailoverSettingOk() (*IPsecFailoverSettingOpenApiVO, bool)`

GetFailoverSettingOk returns a tuple with the FailoverSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverSetting

`func (o *VpnSiteToSiteDetailOpenApiVO) SetFailoverSetting(v IPsecFailoverSettingOpenApiVO)`

SetFailoverSetting sets FailoverSetting field to given value.

### HasFailoverSetting

`func (o *VpnSiteToSiteDetailOpenApiVO) HasFailoverSetting() bool`

HasFailoverSetting returns a boolean if a field has been set.

### GetFeatureDescription

`func (o *VpnSiteToSiteDetailOpenApiVO) GetFeatureDescription() []FeatureInfoVO`

GetFeatureDescription returns the FeatureDescription field if non-nil, zero value otherwise.

### GetFeatureDescriptionOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetFeatureDescriptionOk() (*[]FeatureInfoVO, bool)`

GetFeatureDescriptionOk returns a tuple with the FeatureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureDescription

`func (o *VpnSiteToSiteDetailOpenApiVO) SetFeatureDescription(v []FeatureInfoVO)`

SetFeatureDescription sets FeatureDescription field to given value.

### HasFeatureDescription

`func (o *VpnSiteToSiteDetailOpenApiVO) HasFeatureDescription() bool`

HasFeatureDescription returns a boolean if a field has been set.

### GetId

`func (o *VpnSiteToSiteDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnSiteToSiteDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnSiteToSiteDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtu

`func (o *VpnSiteToSiteDetailOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VpnSiteToSiteDetailOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VpnSiteToSiteDetailOpenApiVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *VpnSiteToSiteDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnSiteToSiteDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnSiteToSiteDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkList

`func (o *VpnSiteToSiteDetailOpenApiVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VpnSiteToSiteDetailOpenApiVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VpnSiteToSiteDetailOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *VpnSiteToSiteDetailOpenApiVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VpnSiteToSiteDetailOpenApiVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VpnSiteToSiteDetailOpenApiVO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPeers

`func (o *VpnSiteToSiteDetailOpenApiVO) GetPeers() []SiteToSiteManualWgPeerConfigVO`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetPeersOk() (*[]SiteToSiteManualWgPeerConfigVO, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *VpnSiteToSiteDetailOpenApiVO) SetPeers(v []SiteToSiteManualWgPeerConfigVO)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *VpnSiteToSiteDetailOpenApiVO) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *VpnSiteToSiteDetailOpenApiVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnSiteToSiteDetailOpenApiVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *VpnSiteToSiteDetailOpenApiVO) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *VpnSiteToSiteDetailOpenApiVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VpnSiteToSiteDetailOpenApiVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *VpnSiteToSiteDetailOpenApiVO) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *VpnSiteToSiteDetailOpenApiVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *VpnSiteToSiteDetailOpenApiVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *VpnSiteToSiteDetailOpenApiVO) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRemoteIp

`func (o *VpnSiteToSiteDetailOpenApiVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnSiteToSiteDetailOpenApiVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *VpnSiteToSiteDetailOpenApiVO) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteSite

`func (o *VpnSiteToSiteDetailOpenApiVO) GetRemoteSite() string`

GetRemoteSite returns the RemoteSite field if non-nil, zero value otherwise.

### GetRemoteSiteOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetRemoteSiteOk() (*string, bool)`

GetRemoteSiteOk returns a tuple with the RemoteSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSite

`func (o *VpnSiteToSiteDetailOpenApiVO) SetRemoteSite(v string)`

SetRemoteSite sets RemoteSite field to given value.

### HasRemoteSite

`func (o *VpnSiteToSiteDetailOpenApiVO) HasRemoteSite() bool`

HasRemoteSite returns a boolean if a field has been set.

### GetRemoteSiteName

`func (o *VpnSiteToSiteDetailOpenApiVO) GetRemoteSiteName() string`

GetRemoteSiteName returns the RemoteSiteName field if non-nil, zero value otherwise.

### GetRemoteSiteNameOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetRemoteSiteNameOk() (*string, bool)`

GetRemoteSiteNameOk returns a tuple with the RemoteSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSiteName

`func (o *VpnSiteToSiteDetailOpenApiVO) SetRemoteSiteName(v string)`

SetRemoteSiteName sets RemoteSiteName field to given value.

### HasRemoteSiteName

`func (o *VpnSiteToSiteDetailOpenApiVO) HasRemoteSiteName() bool`

HasRemoteSiteName returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *VpnSiteToSiteDetailOpenApiVO) GetRemoteSubnet() []string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetRemoteSubnetOk() (*[]string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnSiteToSiteDetailOpenApiVO) SetRemoteSubnet(v []string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *VpnSiteToSiteDetailOpenApiVO) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetServicePort

`func (o *VpnSiteToSiteDetailOpenApiVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VpnSiteToSiteDetailOpenApiVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *VpnSiteToSiteDetailOpenApiVO) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetSiteVpnType

`func (o *VpnSiteToSiteDetailOpenApiVO) GetSiteVpnType() int32`

GetSiteVpnType returns the SiteVpnType field if non-nil, zero value otherwise.

### GetSiteVpnTypeOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetSiteVpnTypeOk() (*int32, bool)`

GetSiteVpnTypeOk returns a tuple with the SiteVpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteVpnType

`func (o *VpnSiteToSiteDetailOpenApiVO) SetSiteVpnType(v int32)`

SetSiteVpnType sets SiteVpnType field to given value.

### HasSiteVpnType

`func (o *VpnSiteToSiteDetailOpenApiVO) HasSiteVpnType() bool`

HasSiteVpnType returns a boolean if a field has been set.

### GetStatus

`func (o *VpnSiteToSiteDetailOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnSiteToSiteDetailOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnSiteToSiteDetailOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnelIp

`func (o *VpnSiteToSiteDetailOpenApiVO) GetTunnelIp() string`

GetTunnelIp returns the TunnelIp field if non-nil, zero value otherwise.

### GetTunnelIpOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetTunnelIpOk() (*string, bool)`

GetTunnelIpOk returns a tuple with the TunnelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIp

`func (o *VpnSiteToSiteDetailOpenApiVO) SetTunnelIp(v string)`

SetTunnelIp sets TunnelIp field to given value.

### HasTunnelIp

`func (o *VpnSiteToSiteDetailOpenApiVO) HasTunnelIp() bool`

HasTunnelIp returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnSiteToSiteDetailOpenApiVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnSiteToSiteDetailOpenApiVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnSiteToSiteDetailOpenApiVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetWans

`func (o *VpnSiteToSiteDetailOpenApiVO) GetWans() []string`

GetWans returns the Wans field if non-nil, zero value otherwise.

### GetWansOk

`func (o *VpnSiteToSiteDetailOpenApiVO) GetWansOk() (*[]string, bool)`

GetWansOk returns a tuple with the Wans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWans

`func (o *VpnSiteToSiteDetailOpenApiVO) SetWans(v []string)`

SetWans sets Wans field to given value.

### HasWans

`func (o *VpnSiteToSiteDetailOpenApiVO) HasWans() bool`

HasWans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


