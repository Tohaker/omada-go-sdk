# ClientToSiteVpnServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPassword** | Pointer to **bool** | Account password of the VPN, only for server OpenVPN. | [optional] 
**AdvancedSetting** | Pointer to [**VpnAdvancedSettingOpenApiVO**](VpnAdvancedSettingOpenApiVO.md) |  | [optional] 
**AuthMode** | Pointer to **int32** | Authentication mode should be a value as follows: 0: Local; 1: LDAP. | [optional] 
**ClientVpnType** | **int32** | Client Vpn type should be a value as follows: 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN. | 
**CustomNetwork** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Custom networks of the VPN. | [optional] 
**Encryption** | Pointer to **int32** | Encryption should be a value as follows: 0: Encrypted; 1: Unencrypted. | [optional] 
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**IpPool** | [**IPSubnetsVO**](IPSubnetsVO.md) |  | 
**IpPoolEnd** | Pointer to **string** | The end IP of the IP pool. | [optional] 
**IpPoolStart** | Pointer to **string** | The start IP of the IP pool. | [optional] 
**IpPoolType** | Pointer to **int32** | IP pool type should be a value as follows: 0: Ip Address/Mask; 1: Ip Address Range. | [optional] 
**Name** | **string** | Name should contain 1 to 63 characters. | 
**NetworkList** | Pointer to **[]string** | Network list of the VPN. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**NetworkType** | Pointer to **int32** | Network type should be a value as follows: 0: network list; 1: custom networks. | [optional] 
**OpenVpnMode** | Pointer to **int32** | OpenVPN mode should be a value as follows: 0: certification; 1: certification+account. | [optional] 
**PreSharedKey** | Pointer to **string** | Pre-shared key of the VPN. | [optional] 
**PrimaryDns** | Pointer to **string** | Primary DNS of the VPN. | [optional] 
**RemoteIp** | Pointer to **string** | Remote IP of the VPN | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS of the VPN. | [optional] 
**ServicePort** | Pointer to **int32** | Service port should be within the range of 1–65535. | [optional] 
**ServiceType** | Pointer to **int32** | Service type of the VPN. | [optional] 
**Status** | Pointer to **bool** | Status of the VPN. | [optional] 
**TunnelMode** | Pointer to **int32** | Tunnel mode(only for server OpenVPN)should be a value as follows: 0: split; 1: full. | [optional] 
**Wan** | **[]string** | WAN list of the VPN. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 

## Methods

### NewClientToSiteVpnServer

`func NewClientToSiteVpnServer(clientVpnType int32, ipPool IPSubnetsVO, name string, wan []string, ) *ClientToSiteVpnServer`

NewClientToSiteVpnServer instantiates a new ClientToSiteVpnServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientToSiteVpnServerWithDefaults

`func NewClientToSiteVpnServerWithDefaults() *ClientToSiteVpnServer`

NewClientToSiteVpnServerWithDefaults instantiates a new ClientToSiteVpnServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPassword

`func (o *ClientToSiteVpnServer) GetAccountPassword() bool`

GetAccountPassword returns the AccountPassword field if non-nil, zero value otherwise.

### GetAccountPasswordOk

`func (o *ClientToSiteVpnServer) GetAccountPasswordOk() (*bool, bool)`

GetAccountPasswordOk returns a tuple with the AccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPassword

`func (o *ClientToSiteVpnServer) SetAccountPassword(v bool)`

SetAccountPassword sets AccountPassword field to given value.

### HasAccountPassword

`func (o *ClientToSiteVpnServer) HasAccountPassword() bool`

HasAccountPassword returns a boolean if a field has been set.

### GetAdvancedSetting

`func (o *ClientToSiteVpnServer) GetAdvancedSetting() VpnAdvancedSettingOpenApiVO`

GetAdvancedSetting returns the AdvancedSetting field if non-nil, zero value otherwise.

### GetAdvancedSettingOk

`func (o *ClientToSiteVpnServer) GetAdvancedSettingOk() (*VpnAdvancedSettingOpenApiVO, bool)`

GetAdvancedSettingOk returns a tuple with the AdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSetting

`func (o *ClientToSiteVpnServer) SetAdvancedSetting(v VpnAdvancedSettingOpenApiVO)`

SetAdvancedSetting sets AdvancedSetting field to given value.

### HasAdvancedSetting

`func (o *ClientToSiteVpnServer) HasAdvancedSetting() bool`

HasAdvancedSetting returns a boolean if a field has been set.

### GetAuthMode

`func (o *ClientToSiteVpnServer) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *ClientToSiteVpnServer) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *ClientToSiteVpnServer) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *ClientToSiteVpnServer) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetClientVpnType

`func (o *ClientToSiteVpnServer) GetClientVpnType() int32`

GetClientVpnType returns the ClientVpnType field if non-nil, zero value otherwise.

### GetClientVpnTypeOk

`func (o *ClientToSiteVpnServer) GetClientVpnTypeOk() (*int32, bool)`

GetClientVpnTypeOk returns a tuple with the ClientVpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVpnType

`func (o *ClientToSiteVpnServer) SetClientVpnType(v int32)`

SetClientVpnType sets ClientVpnType field to given value.


### GetCustomNetwork

`func (o *ClientToSiteVpnServer) GetCustomNetwork() []IPSubnetsVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *ClientToSiteVpnServer) GetCustomNetworkOk() (*[]IPSubnetsVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *ClientToSiteVpnServer) SetCustomNetwork(v []IPSubnetsVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *ClientToSiteVpnServer) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetEncryption

`func (o *ClientToSiteVpnServer) GetEncryption() int32`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ClientToSiteVpnServer) GetEncryptionOk() (*int32, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ClientToSiteVpnServer) SetEncryption(v int32)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *ClientToSiteVpnServer) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetId

`func (o *ClientToSiteVpnServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientToSiteVpnServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientToSiteVpnServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientToSiteVpnServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpPool

`func (o *ClientToSiteVpnServer) GetIpPool() IPSubnetsVO`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *ClientToSiteVpnServer) GetIpPoolOk() (*IPSubnetsVO, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *ClientToSiteVpnServer) SetIpPool(v IPSubnetsVO)`

SetIpPool sets IpPool field to given value.


### GetIpPoolEnd

`func (o *ClientToSiteVpnServer) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *ClientToSiteVpnServer) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *ClientToSiteVpnServer) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *ClientToSiteVpnServer) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *ClientToSiteVpnServer) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *ClientToSiteVpnServer) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *ClientToSiteVpnServer) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *ClientToSiteVpnServer) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetIpPoolType

`func (o *ClientToSiteVpnServer) GetIpPoolType() int32`

GetIpPoolType returns the IpPoolType field if non-nil, zero value otherwise.

### GetIpPoolTypeOk

`func (o *ClientToSiteVpnServer) GetIpPoolTypeOk() (*int32, bool)`

GetIpPoolTypeOk returns a tuple with the IpPoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolType

`func (o *ClientToSiteVpnServer) SetIpPoolType(v int32)`

SetIpPoolType sets IpPoolType field to given value.

### HasIpPoolType

`func (o *ClientToSiteVpnServer) HasIpPoolType() bool`

HasIpPoolType returns a boolean if a field has been set.

### GetName

`func (o *ClientToSiteVpnServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientToSiteVpnServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientToSiteVpnServer) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkList

`func (o *ClientToSiteVpnServer) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *ClientToSiteVpnServer) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *ClientToSiteVpnServer) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *ClientToSiteVpnServer) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *ClientToSiteVpnServer) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *ClientToSiteVpnServer) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *ClientToSiteVpnServer) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *ClientToSiteVpnServer) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOpenVpnMode

`func (o *ClientToSiteVpnServer) GetOpenVpnMode() int32`

GetOpenVpnMode returns the OpenVpnMode field if non-nil, zero value otherwise.

### GetOpenVpnModeOk

`func (o *ClientToSiteVpnServer) GetOpenVpnModeOk() (*int32, bool)`

GetOpenVpnModeOk returns a tuple with the OpenVpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVpnMode

`func (o *ClientToSiteVpnServer) SetOpenVpnMode(v int32)`

SetOpenVpnMode sets OpenVpnMode field to given value.

### HasOpenVpnMode

`func (o *ClientToSiteVpnServer) HasOpenVpnMode() bool`

HasOpenVpnMode returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *ClientToSiteVpnServer) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *ClientToSiteVpnServer) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *ClientToSiteVpnServer) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *ClientToSiteVpnServer) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *ClientToSiteVpnServer) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *ClientToSiteVpnServer) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *ClientToSiteVpnServer) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *ClientToSiteVpnServer) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetRemoteIp

`func (o *ClientToSiteVpnServer) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *ClientToSiteVpnServer) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *ClientToSiteVpnServer) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *ClientToSiteVpnServer) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *ClientToSiteVpnServer) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *ClientToSiteVpnServer) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *ClientToSiteVpnServer) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *ClientToSiteVpnServer) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetServicePort

`func (o *ClientToSiteVpnServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ClientToSiteVpnServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ClientToSiteVpnServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *ClientToSiteVpnServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceType

`func (o *ClientToSiteVpnServer) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ClientToSiteVpnServer) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ClientToSiteVpnServer) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ClientToSiteVpnServer) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *ClientToSiteVpnServer) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientToSiteVpnServer) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientToSiteVpnServer) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientToSiteVpnServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnelMode

`func (o *ClientToSiteVpnServer) GetTunnelMode() int32`

GetTunnelMode returns the TunnelMode field if non-nil, zero value otherwise.

### GetTunnelModeOk

`func (o *ClientToSiteVpnServer) GetTunnelModeOk() (*int32, bool)`

GetTunnelModeOk returns a tuple with the TunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelMode

`func (o *ClientToSiteVpnServer) SetTunnelMode(v int32)`

SetTunnelMode sets TunnelMode field to given value.

### HasTunnelMode

`func (o *ClientToSiteVpnServer) HasTunnelMode() bool`

HasTunnelMode returns a boolean if a field has been set.

### GetWan

`func (o *ClientToSiteVpnServer) GetWan() []string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *ClientToSiteVpnServer) GetWanOk() (*[]string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *ClientToSiteVpnServer) SetWan(v []string)`

SetWan sets Wan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


