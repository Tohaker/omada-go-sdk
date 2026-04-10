# VpnSiteToSiteManualConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSetting** | [**VpnAdvancedSettingOpenApiVO**](VpnAdvancedSettingOpenApiVO.md) |  | 
**CustomNetwork** | Pointer to [**[]VpnIPSubnetsOpenApiVO**](VpnIPSubnetsOpenApiVO.md) | Custom networks of the VPN. | [optional] 
**ExistCustomNetwork** | Pointer to **bool** | Whether Local Network Type is Custom. | [optional] 
**FailoverSetting** | Pointer to [**IPsecFailoverSettingOpenApiVO**](IPsecFailoverSettingOpenApiVO.md) |  | [optional] 
**Mtu** | **int32** | The MTU of WireGuard VPN should be within the range of 576-1440. | 
**Name** | **string** | Name should contain 1 to 63 characters. | 
**NetworkList** | Pointer to **[]string** | Network list of the VPN. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**NetworkType** | Pointer to **int32** | Network type should be a value as follows: 0: network list; 1: custom networks. | [optional] 
**Peers** | Pointer to [**[]SiteToSiteManualWgPeerConfigVO**](SiteToSiteManualWgPeerConfigVO.md) | List of Site-To-Site manual WireGuard peer Configuration. | [optional] 
**PreSharedKey** | **string** | Pre-shared key of the VPN. | 
**PrivateKey** | **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | 
**RemoteIp** | **string** | Remote IP of the VPN. | 
**RemoteSubnet** | **[]string** | Remote subnet of the VPN. | 
**ServicePort** | **int32** | Service port should be within the range of 1–65535. | 
**Status** | **bool** | Status of the VPN. | 
**TunnelIp** | **string** | IP address. | 
**VpnType** | **int32** | Vpn type should be a value as follows: 2: IPSec; 4: WireGuard. | 
**Wans** | **[]string** | WAN list of the VPN, only for Manual IPSec type. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 

## Methods

### NewVpnSiteToSiteManualConfigOpenApiVO

`func NewVpnSiteToSiteManualConfigOpenApiVO(advancedSetting VpnAdvancedSettingOpenApiVO, mtu int32, name string, preSharedKey string, privateKey string, remoteIp string, remoteSubnet []string, servicePort int32, status bool, tunnelIp string, vpnType int32, wans []string, ) *VpnSiteToSiteManualConfigOpenApiVO`

NewVpnSiteToSiteManualConfigOpenApiVO instantiates a new VpnSiteToSiteManualConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnSiteToSiteManualConfigOpenApiVOWithDefaults

`func NewVpnSiteToSiteManualConfigOpenApiVOWithDefaults() *VpnSiteToSiteManualConfigOpenApiVO`

NewVpnSiteToSiteManualConfigOpenApiVOWithDefaults instantiates a new VpnSiteToSiteManualConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSetting

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetAdvancedSetting() VpnAdvancedSettingOpenApiVO`

GetAdvancedSetting returns the AdvancedSetting field if non-nil, zero value otherwise.

### GetAdvancedSettingOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetAdvancedSettingOk() (*VpnAdvancedSettingOpenApiVO, bool)`

GetAdvancedSettingOk returns a tuple with the AdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSetting

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetAdvancedSetting(v VpnAdvancedSettingOpenApiVO)`

SetAdvancedSetting sets AdvancedSetting field to given value.


### GetCustomNetwork

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetCustomNetwork() []VpnIPSubnetsOpenApiVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetCustomNetworkOk() (*[]VpnIPSubnetsOpenApiVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetCustomNetwork(v []VpnIPSubnetsOpenApiVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *VpnSiteToSiteManualConfigOpenApiVO) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetExistCustomNetwork

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetExistCustomNetwork() bool`

GetExistCustomNetwork returns the ExistCustomNetwork field if non-nil, zero value otherwise.

### GetExistCustomNetworkOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetExistCustomNetworkOk() (*bool, bool)`

GetExistCustomNetworkOk returns a tuple with the ExistCustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCustomNetwork

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetExistCustomNetwork(v bool)`

SetExistCustomNetwork sets ExistCustomNetwork field to given value.

### HasExistCustomNetwork

`func (o *VpnSiteToSiteManualConfigOpenApiVO) HasExistCustomNetwork() bool`

HasExistCustomNetwork returns a boolean if a field has been set.

### GetFailoverSetting

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetFailoverSetting() IPsecFailoverSettingOpenApiVO`

GetFailoverSetting returns the FailoverSetting field if non-nil, zero value otherwise.

### GetFailoverSettingOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetFailoverSettingOk() (*IPsecFailoverSettingOpenApiVO, bool)`

GetFailoverSettingOk returns a tuple with the FailoverSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverSetting

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetFailoverSetting(v IPsecFailoverSettingOpenApiVO)`

SetFailoverSetting sets FailoverSetting field to given value.

### HasFailoverSetting

`func (o *VpnSiteToSiteManualConfigOpenApiVO) HasFailoverSetting() bool`

HasFailoverSetting returns a boolean if a field has been set.

### GetMtu

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### GetName

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkList

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VpnSiteToSiteManualConfigOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VpnSiteToSiteManualConfigOpenApiVO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPeers

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetPeers() []SiteToSiteManualWgPeerConfigVO`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetPeersOk() (*[]SiteToSiteManualWgPeerConfigVO, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetPeers(v []SiteToSiteManualWgPeerConfigVO)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *VpnSiteToSiteManualConfigOpenApiVO) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.


### GetPrivateKey

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetRemoteIp

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetRemoteSubnet

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetRemoteSubnet() []string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetRemoteSubnetOk() (*[]string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetRemoteSubnet(v []string)`

SetRemoteSubnet sets RemoteSubnet field to given value.


### GetServicePort

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.


### GetStatus

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTunnelIp

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetTunnelIp() string`

GetTunnelIp returns the TunnelIp field if non-nil, zero value otherwise.

### GetTunnelIpOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetTunnelIpOk() (*string, bool)`

GetTunnelIpOk returns a tuple with the TunnelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIp

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetTunnelIp(v string)`

SetTunnelIp sets TunnelIp field to given value.


### GetVpnType

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.


### GetWans

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetWans() []string`

GetWans returns the Wans field if non-nil, zero value otherwise.

### GetWansOk

`func (o *VpnSiteToSiteManualConfigOpenApiVO) GetWansOk() (*[]string, bool)`

GetWansOk returns a tuple with the Wans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWans

`func (o *VpnSiteToSiteManualConfigOpenApiVO) SetWans(v []string)`

SetWans sets Wans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


