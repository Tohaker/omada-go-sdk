# VpnClientConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedServerAddress** | Pointer to **[]string** | Wire Guard allowed server address. | [optional] 
**ClientPassword** | Pointer to **string** | Client password of the VPN. | [optional] 
**ClientPort** | Pointer to **int32** | Client port should be within the range of 1–65535. | [optional] 
**ClientUsername** | Pointer to **string** | Client username of the VPN. | [optional] 
**CustomNetwork** | Pointer to [**[]VpnIPSubnetsOpenApiVO**](VpnIPSubnetsOpenApiVO.md) | Custom networks of the VPN. | [optional] 
**Dns1** | Pointer to **string** | Primary DNS of the VPN. | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS of the VPN. | [optional] 
**Encryption** | Pointer to **int32** | Encryption should be a value as follows: 0: Encrypted; 1: Unencrypted; 2: Auto. | [optional] 
**KeepAlive** | Pointer to **int32** | The keepalive second of WireGuard peer should be within the range of 0-65535. | [optional] 
**Mtu** | Pointer to **int32** | The MTU of WireGuard VPN should be within the range of 576-1440. | [optional] 
**Name** | **string** | Name should contain 1 to 63 characters. | 
**NetworkList** | Pointer to **[]string** | Network list of the VPN. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**NetworkType** | Pointer to **int32** | Network type should be a value as follows: 0: network list; 1: custom networks. | [optional] 
**OpenVpnMode** | Pointer to **int32** | OpenVPN mode should be a value as follows: 0: certification; 1: certification+account. | [optional] 
**PreSharedKey** | Pointer to **string** | Pre-shared key of the VPN. | [optional] 
**PrivateKey** | Pointer to **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**RemoteIp** | **string** | Remote IP of the VPN | 
**RemoteSubnet** | Pointer to **[]string** | Remote subnet of the VPN. | [optional] 
**ServerPublicKey** | **string** | The public key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | 
**ServicePort** | **int32** | Service port should be within the range of 1–65535. | 
**Setup** | **int32** | Wire Guard setup should be a value as follows: 0: file; 1: manual. | 
**Status** | **bool** | Status of the VPN. | 
**TunnelIp** | Pointer to **string** | The local IP address of WireGuard VPN. | [optional] 
**VpnConfiguration** | [**VpnCertificateOpenApiVO**](VpnCertificateOpenApiVO.md) |  | 
**VpnType** | **int32** | Vpn type should be a value as follows: 0: L2TP; 1: PPTP; 3: OpenVPN; 4: WireGuard. | 
**Wans** | **[]string** | WAN list of the VPN. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 
**WorkingMode** | Pointer to **int32** | Working Mode should be a value as follow: 0:NAT 1:Routing. | [optional] 

## Methods

### NewVpnClientConfigOpenApiVO

`func NewVpnClientConfigOpenApiVO(name string, remoteIp string, serverPublicKey string, servicePort int32, setup int32, status bool, vpnConfiguration VpnCertificateOpenApiVO, vpnType int32, wans []string, ) *VpnClientConfigOpenApiVO`

NewVpnClientConfigOpenApiVO instantiates a new VpnClientConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnClientConfigOpenApiVOWithDefaults

`func NewVpnClientConfigOpenApiVOWithDefaults() *VpnClientConfigOpenApiVO`

NewVpnClientConfigOpenApiVOWithDefaults instantiates a new VpnClientConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedServerAddress

`func (o *VpnClientConfigOpenApiVO) GetAllowedServerAddress() []string`

GetAllowedServerAddress returns the AllowedServerAddress field if non-nil, zero value otherwise.

### GetAllowedServerAddressOk

`func (o *VpnClientConfigOpenApiVO) GetAllowedServerAddressOk() (*[]string, bool)`

GetAllowedServerAddressOk returns a tuple with the AllowedServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServerAddress

`func (o *VpnClientConfigOpenApiVO) SetAllowedServerAddress(v []string)`

SetAllowedServerAddress sets AllowedServerAddress field to given value.

### HasAllowedServerAddress

`func (o *VpnClientConfigOpenApiVO) HasAllowedServerAddress() bool`

HasAllowedServerAddress returns a boolean if a field has been set.

### GetClientPassword

`func (o *VpnClientConfigOpenApiVO) GetClientPassword() string`

GetClientPassword returns the ClientPassword field if non-nil, zero value otherwise.

### GetClientPasswordOk

`func (o *VpnClientConfigOpenApiVO) GetClientPasswordOk() (*string, bool)`

GetClientPasswordOk returns a tuple with the ClientPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPassword

`func (o *VpnClientConfigOpenApiVO) SetClientPassword(v string)`

SetClientPassword sets ClientPassword field to given value.

### HasClientPassword

`func (o *VpnClientConfigOpenApiVO) HasClientPassword() bool`

HasClientPassword returns a boolean if a field has been set.

### GetClientPort

`func (o *VpnClientConfigOpenApiVO) GetClientPort() int32`

GetClientPort returns the ClientPort field if non-nil, zero value otherwise.

### GetClientPortOk

`func (o *VpnClientConfigOpenApiVO) GetClientPortOk() (*int32, bool)`

GetClientPortOk returns a tuple with the ClientPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPort

`func (o *VpnClientConfigOpenApiVO) SetClientPort(v int32)`

SetClientPort sets ClientPort field to given value.

### HasClientPort

`func (o *VpnClientConfigOpenApiVO) HasClientPort() bool`

HasClientPort returns a boolean if a field has been set.

### GetClientUsername

`func (o *VpnClientConfigOpenApiVO) GetClientUsername() string`

GetClientUsername returns the ClientUsername field if non-nil, zero value otherwise.

### GetClientUsernameOk

`func (o *VpnClientConfigOpenApiVO) GetClientUsernameOk() (*string, bool)`

GetClientUsernameOk returns a tuple with the ClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsername

`func (o *VpnClientConfigOpenApiVO) SetClientUsername(v string)`

SetClientUsername sets ClientUsername field to given value.

### HasClientUsername

`func (o *VpnClientConfigOpenApiVO) HasClientUsername() bool`

HasClientUsername returns a boolean if a field has been set.

### GetCustomNetwork

`func (o *VpnClientConfigOpenApiVO) GetCustomNetwork() []VpnIPSubnetsOpenApiVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *VpnClientConfigOpenApiVO) GetCustomNetworkOk() (*[]VpnIPSubnetsOpenApiVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *VpnClientConfigOpenApiVO) SetCustomNetwork(v []VpnIPSubnetsOpenApiVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *VpnClientConfigOpenApiVO) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetDns1

`func (o *VpnClientConfigOpenApiVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VpnClientConfigOpenApiVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VpnClientConfigOpenApiVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VpnClientConfigOpenApiVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VpnClientConfigOpenApiVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VpnClientConfigOpenApiVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VpnClientConfigOpenApiVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VpnClientConfigOpenApiVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetEncryption

`func (o *VpnClientConfigOpenApiVO) GetEncryption() int32`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VpnClientConfigOpenApiVO) GetEncryptionOk() (*int32, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VpnClientConfigOpenApiVO) SetEncryption(v int32)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VpnClientConfigOpenApiVO) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetKeepAlive

`func (o *VpnClientConfigOpenApiVO) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *VpnClientConfigOpenApiVO) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *VpnClientConfigOpenApiVO) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *VpnClientConfigOpenApiVO) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetMtu

`func (o *VpnClientConfigOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VpnClientConfigOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VpnClientConfigOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VpnClientConfigOpenApiVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *VpnClientConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnClientConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnClientConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkList

`func (o *VpnClientConfigOpenApiVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VpnClientConfigOpenApiVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VpnClientConfigOpenApiVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VpnClientConfigOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *VpnClientConfigOpenApiVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VpnClientConfigOpenApiVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VpnClientConfigOpenApiVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VpnClientConfigOpenApiVO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOpenVpnMode

`func (o *VpnClientConfigOpenApiVO) GetOpenVpnMode() int32`

GetOpenVpnMode returns the OpenVpnMode field if non-nil, zero value otherwise.

### GetOpenVpnModeOk

`func (o *VpnClientConfigOpenApiVO) GetOpenVpnModeOk() (*int32, bool)`

GetOpenVpnModeOk returns a tuple with the OpenVpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVpnMode

`func (o *VpnClientConfigOpenApiVO) SetOpenVpnMode(v int32)`

SetOpenVpnMode sets OpenVpnMode field to given value.

### HasOpenVpnMode

`func (o *VpnClientConfigOpenApiVO) HasOpenVpnMode() bool`

HasOpenVpnMode returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *VpnClientConfigOpenApiVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnClientConfigOpenApiVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnClientConfigOpenApiVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *VpnClientConfigOpenApiVO) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *VpnClientConfigOpenApiVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VpnClientConfigOpenApiVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VpnClientConfigOpenApiVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *VpnClientConfigOpenApiVO) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetRemoteIp

`func (o *VpnClientConfigOpenApiVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnClientConfigOpenApiVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnClientConfigOpenApiVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetRemoteSubnet

`func (o *VpnClientConfigOpenApiVO) GetRemoteSubnet() []string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnClientConfigOpenApiVO) GetRemoteSubnetOk() (*[]string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnClientConfigOpenApiVO) SetRemoteSubnet(v []string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *VpnClientConfigOpenApiVO) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetServerPublicKey

`func (o *VpnClientConfigOpenApiVO) GetServerPublicKey() string`

GetServerPublicKey returns the ServerPublicKey field if non-nil, zero value otherwise.

### GetServerPublicKeyOk

`func (o *VpnClientConfigOpenApiVO) GetServerPublicKeyOk() (*string, bool)`

GetServerPublicKeyOk returns a tuple with the ServerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPublicKey

`func (o *VpnClientConfigOpenApiVO) SetServerPublicKey(v string)`

SetServerPublicKey sets ServerPublicKey field to given value.


### GetServicePort

`func (o *VpnClientConfigOpenApiVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VpnClientConfigOpenApiVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VpnClientConfigOpenApiVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.


### GetSetup

`func (o *VpnClientConfigOpenApiVO) GetSetup() int32`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *VpnClientConfigOpenApiVO) GetSetupOk() (*int32, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *VpnClientConfigOpenApiVO) SetSetup(v int32)`

SetSetup sets Setup field to given value.


### GetStatus

`func (o *VpnClientConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnClientConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnClientConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTunnelIp

`func (o *VpnClientConfigOpenApiVO) GetTunnelIp() string`

GetTunnelIp returns the TunnelIp field if non-nil, zero value otherwise.

### GetTunnelIpOk

`func (o *VpnClientConfigOpenApiVO) GetTunnelIpOk() (*string, bool)`

GetTunnelIpOk returns a tuple with the TunnelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIp

`func (o *VpnClientConfigOpenApiVO) SetTunnelIp(v string)`

SetTunnelIp sets TunnelIp field to given value.

### HasTunnelIp

`func (o *VpnClientConfigOpenApiVO) HasTunnelIp() bool`

HasTunnelIp returns a boolean if a field has been set.

### GetVpnConfiguration

`func (o *VpnClientConfigOpenApiVO) GetVpnConfiguration() VpnCertificateOpenApiVO`

GetVpnConfiguration returns the VpnConfiguration field if non-nil, zero value otherwise.

### GetVpnConfigurationOk

`func (o *VpnClientConfigOpenApiVO) GetVpnConfigurationOk() (*VpnCertificateOpenApiVO, bool)`

GetVpnConfigurationOk returns a tuple with the VpnConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConfiguration

`func (o *VpnClientConfigOpenApiVO) SetVpnConfiguration(v VpnCertificateOpenApiVO)`

SetVpnConfiguration sets VpnConfiguration field to given value.


### GetVpnType

`func (o *VpnClientConfigOpenApiVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnClientConfigOpenApiVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnClientConfigOpenApiVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.


### GetWans

`func (o *VpnClientConfigOpenApiVO) GetWans() []string`

GetWans returns the Wans field if non-nil, zero value otherwise.

### GetWansOk

`func (o *VpnClientConfigOpenApiVO) GetWansOk() (*[]string, bool)`

GetWansOk returns a tuple with the Wans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWans

`func (o *VpnClientConfigOpenApiVO) SetWans(v []string)`

SetWans sets Wans field to given value.


### GetWorkingMode

`func (o *VpnClientConfigOpenApiVO) GetWorkingMode() int32`

GetWorkingMode returns the WorkingMode field if non-nil, zero value otherwise.

### GetWorkingModeOk

`func (o *VpnClientConfigOpenApiVO) GetWorkingModeOk() (*int32, bool)`

GetWorkingModeOk returns a tuple with the WorkingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingMode

`func (o *VpnClientConfigOpenApiVO) SetWorkingMode(v int32)`

SetWorkingMode sets WorkingMode field to given value.

### HasWorkingMode

`func (o *VpnClientConfigOpenApiVO) HasWorkingMode() bool`

HasWorkingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


