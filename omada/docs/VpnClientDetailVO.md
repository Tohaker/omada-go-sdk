# VpnClientDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedServerAddress** | Pointer to **[]string** | Wire Guard allowed server address. | [optional] 
**ClientPassword** | Pointer to **string** | Client password should contain 1 to 64 characters. | [optional] 
**ClientPort** | Pointer to **int32** | Client port for VPN client. | [optional] 
**ClientUsername** | Pointer to **string** | Client username should contain 1 to 64 characters. | [optional] 
**CustomNetwork** | Pointer to [**[]VpnIPSubnetsOpenApiVO**](VpnIPSubnetsOpenApiVO.md) | Custom networks of the VPN. | [optional] 
**Dns1** | Pointer to **string** | Primary DNS of the VPN. | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS of the VPN. | [optional] 
**Encryption** | Pointer to **int32** | Encryption should be a value as follows: 0: Encrypted; 1: Unencrypted. | [optional] 
**FeatureDescription** | Pointer to [**[]FeatureInfoVO**](FeatureInfoVO.md) | Gateway Feature Description | [optional] 
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**KeepAlive** | Pointer to **int32** | The keepalive second of WireGuard peer should be within the range of 0-65535. | [optional] 
**Mtu** | Pointer to **int32** | The MTU of WireGuard VPN should be within the range of 576-1440. | [optional] 
**Name** | Pointer to **string** | VPN name. | [optional] 
**NetworkList** | Pointer to **[]string** | Network list of the VPN.  | [optional] 
**NetworkType** | Pointer to **int32** | Network type. 0: network list; 1: custom networks. | [optional] 
**OpenVpnMode** | Pointer to **int32** | OpenVPN mode should be a value as follows: 0: certification; 1: certification+account. | [optional] 
**PreSharedKey** | Pointer to **string** | Pre-shared key of the VPN. | [optional] 
**PrivateKey** | Pointer to **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PublicKey** | Pointer to **string** | The public key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**RemoteIp** | Pointer to **string** | Remote IP of the VPN. | [optional] 
**RemoteSubnet** | Pointer to **[]string** | Remote subnet of the VPN. | [optional] 
**ServerPublicKey** | Pointer to **string** | The public key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**ServicePort** | Pointer to **int32** | Service port for VPN server. | [optional] 
**Setup** | Pointer to **int32** | Wire Guard setup should be a value as follows: 0: file; 1: manual. | [optional] 
**Status** | Pointer to **bool** | Status of the VPN. | [optional] 
**TunnelIp** | Pointer to **string** | The local IP address of WireGuard VPN. | [optional] 
**VpnConfiguration** | Pointer to [**VpnCertificateOpenApiVO**](VpnCertificateOpenApiVO.md) |  | [optional] 
**VpnType** | Pointer to **int32** | Client Vpn type. 0: L2TP; 1: PPTP; 3: OpenVPN; 4: WireGuard. | [optional] 
**Wans** | Pointer to **[]string** | WAN port ID. | [optional] 
**WorkingMode** | Pointer to **int32** | Working mode should be a value as follows: 0: NAT; 1: Routing. | [optional] 

## Methods

### NewVpnClientDetailVO

`func NewVpnClientDetailVO() *VpnClientDetailVO`

NewVpnClientDetailVO instantiates a new VpnClientDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnClientDetailVOWithDefaults

`func NewVpnClientDetailVOWithDefaults() *VpnClientDetailVO`

NewVpnClientDetailVOWithDefaults instantiates a new VpnClientDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedServerAddress

`func (o *VpnClientDetailVO) GetAllowedServerAddress() []string`

GetAllowedServerAddress returns the AllowedServerAddress field if non-nil, zero value otherwise.

### GetAllowedServerAddressOk

`func (o *VpnClientDetailVO) GetAllowedServerAddressOk() (*[]string, bool)`

GetAllowedServerAddressOk returns a tuple with the AllowedServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServerAddress

`func (o *VpnClientDetailVO) SetAllowedServerAddress(v []string)`

SetAllowedServerAddress sets AllowedServerAddress field to given value.

### HasAllowedServerAddress

`func (o *VpnClientDetailVO) HasAllowedServerAddress() bool`

HasAllowedServerAddress returns a boolean if a field has been set.

### GetClientPassword

`func (o *VpnClientDetailVO) GetClientPassword() string`

GetClientPassword returns the ClientPassword field if non-nil, zero value otherwise.

### GetClientPasswordOk

`func (o *VpnClientDetailVO) GetClientPasswordOk() (*string, bool)`

GetClientPasswordOk returns a tuple with the ClientPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPassword

`func (o *VpnClientDetailVO) SetClientPassword(v string)`

SetClientPassword sets ClientPassword field to given value.

### HasClientPassword

`func (o *VpnClientDetailVO) HasClientPassword() bool`

HasClientPassword returns a boolean if a field has been set.

### GetClientPort

`func (o *VpnClientDetailVO) GetClientPort() int32`

GetClientPort returns the ClientPort field if non-nil, zero value otherwise.

### GetClientPortOk

`func (o *VpnClientDetailVO) GetClientPortOk() (*int32, bool)`

GetClientPortOk returns a tuple with the ClientPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPort

`func (o *VpnClientDetailVO) SetClientPort(v int32)`

SetClientPort sets ClientPort field to given value.

### HasClientPort

`func (o *VpnClientDetailVO) HasClientPort() bool`

HasClientPort returns a boolean if a field has been set.

### GetClientUsername

`func (o *VpnClientDetailVO) GetClientUsername() string`

GetClientUsername returns the ClientUsername field if non-nil, zero value otherwise.

### GetClientUsernameOk

`func (o *VpnClientDetailVO) GetClientUsernameOk() (*string, bool)`

GetClientUsernameOk returns a tuple with the ClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsername

`func (o *VpnClientDetailVO) SetClientUsername(v string)`

SetClientUsername sets ClientUsername field to given value.

### HasClientUsername

`func (o *VpnClientDetailVO) HasClientUsername() bool`

HasClientUsername returns a boolean if a field has been set.

### GetCustomNetwork

`func (o *VpnClientDetailVO) GetCustomNetwork() []VpnIPSubnetsOpenApiVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *VpnClientDetailVO) GetCustomNetworkOk() (*[]VpnIPSubnetsOpenApiVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *VpnClientDetailVO) SetCustomNetwork(v []VpnIPSubnetsOpenApiVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *VpnClientDetailVO) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetDns1

`func (o *VpnClientDetailVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VpnClientDetailVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VpnClientDetailVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VpnClientDetailVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VpnClientDetailVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VpnClientDetailVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VpnClientDetailVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VpnClientDetailVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetEncryption

`func (o *VpnClientDetailVO) GetEncryption() int32`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VpnClientDetailVO) GetEncryptionOk() (*int32, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VpnClientDetailVO) SetEncryption(v int32)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VpnClientDetailVO) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFeatureDescription

`func (o *VpnClientDetailVO) GetFeatureDescription() []FeatureInfoVO`

GetFeatureDescription returns the FeatureDescription field if non-nil, zero value otherwise.

### GetFeatureDescriptionOk

`func (o *VpnClientDetailVO) GetFeatureDescriptionOk() (*[]FeatureInfoVO, bool)`

GetFeatureDescriptionOk returns a tuple with the FeatureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureDescription

`func (o *VpnClientDetailVO) SetFeatureDescription(v []FeatureInfoVO)`

SetFeatureDescription sets FeatureDescription field to given value.

### HasFeatureDescription

`func (o *VpnClientDetailVO) HasFeatureDescription() bool`

HasFeatureDescription returns a boolean if a field has been set.

### GetId

`func (o *VpnClientDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnClientDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnClientDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnClientDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeepAlive

`func (o *VpnClientDetailVO) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *VpnClientDetailVO) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *VpnClientDetailVO) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *VpnClientDetailVO) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetMtu

`func (o *VpnClientDetailVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VpnClientDetailVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VpnClientDetailVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VpnClientDetailVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *VpnClientDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnClientDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnClientDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnClientDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkList

`func (o *VpnClientDetailVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VpnClientDetailVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VpnClientDetailVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VpnClientDetailVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *VpnClientDetailVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VpnClientDetailVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VpnClientDetailVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VpnClientDetailVO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOpenVpnMode

`func (o *VpnClientDetailVO) GetOpenVpnMode() int32`

GetOpenVpnMode returns the OpenVpnMode field if non-nil, zero value otherwise.

### GetOpenVpnModeOk

`func (o *VpnClientDetailVO) GetOpenVpnModeOk() (*int32, bool)`

GetOpenVpnModeOk returns a tuple with the OpenVpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVpnMode

`func (o *VpnClientDetailVO) SetOpenVpnMode(v int32)`

SetOpenVpnMode sets OpenVpnMode field to given value.

### HasOpenVpnMode

`func (o *VpnClientDetailVO) HasOpenVpnMode() bool`

HasOpenVpnMode returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *VpnClientDetailVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnClientDetailVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnClientDetailVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *VpnClientDetailVO) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *VpnClientDetailVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VpnClientDetailVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VpnClientDetailVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *VpnClientDetailVO) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *VpnClientDetailVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *VpnClientDetailVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *VpnClientDetailVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *VpnClientDetailVO) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRemoteIp

`func (o *VpnClientDetailVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnClientDetailVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnClientDetailVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *VpnClientDetailVO) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *VpnClientDetailVO) GetRemoteSubnet() []string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnClientDetailVO) GetRemoteSubnetOk() (*[]string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnClientDetailVO) SetRemoteSubnet(v []string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *VpnClientDetailVO) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetServerPublicKey

`func (o *VpnClientDetailVO) GetServerPublicKey() string`

GetServerPublicKey returns the ServerPublicKey field if non-nil, zero value otherwise.

### GetServerPublicKeyOk

`func (o *VpnClientDetailVO) GetServerPublicKeyOk() (*string, bool)`

GetServerPublicKeyOk returns a tuple with the ServerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPublicKey

`func (o *VpnClientDetailVO) SetServerPublicKey(v string)`

SetServerPublicKey sets ServerPublicKey field to given value.

### HasServerPublicKey

`func (o *VpnClientDetailVO) HasServerPublicKey() bool`

HasServerPublicKey returns a boolean if a field has been set.

### GetServicePort

`func (o *VpnClientDetailVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VpnClientDetailVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VpnClientDetailVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *VpnClientDetailVO) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetSetup

`func (o *VpnClientDetailVO) GetSetup() int32`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *VpnClientDetailVO) GetSetupOk() (*int32, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *VpnClientDetailVO) SetSetup(v int32)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *VpnClientDetailVO) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetStatus

`func (o *VpnClientDetailVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnClientDetailVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnClientDetailVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnClientDetailVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnelIp

`func (o *VpnClientDetailVO) GetTunnelIp() string`

GetTunnelIp returns the TunnelIp field if non-nil, zero value otherwise.

### GetTunnelIpOk

`func (o *VpnClientDetailVO) GetTunnelIpOk() (*string, bool)`

GetTunnelIpOk returns a tuple with the TunnelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIp

`func (o *VpnClientDetailVO) SetTunnelIp(v string)`

SetTunnelIp sets TunnelIp field to given value.

### HasTunnelIp

`func (o *VpnClientDetailVO) HasTunnelIp() bool`

HasTunnelIp returns a boolean if a field has been set.

### GetVpnConfiguration

`func (o *VpnClientDetailVO) GetVpnConfiguration() VpnCertificateOpenApiVO`

GetVpnConfiguration returns the VpnConfiguration field if non-nil, zero value otherwise.

### GetVpnConfigurationOk

`func (o *VpnClientDetailVO) GetVpnConfigurationOk() (*VpnCertificateOpenApiVO, bool)`

GetVpnConfigurationOk returns a tuple with the VpnConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConfiguration

`func (o *VpnClientDetailVO) SetVpnConfiguration(v VpnCertificateOpenApiVO)`

SetVpnConfiguration sets VpnConfiguration field to given value.

### HasVpnConfiguration

`func (o *VpnClientDetailVO) HasVpnConfiguration() bool`

HasVpnConfiguration returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnClientDetailVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnClientDetailVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnClientDetailVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnClientDetailVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetWans

`func (o *VpnClientDetailVO) GetWans() []string`

GetWans returns the Wans field if non-nil, zero value otherwise.

### GetWansOk

`func (o *VpnClientDetailVO) GetWansOk() (*[]string, bool)`

GetWansOk returns a tuple with the Wans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWans

`func (o *VpnClientDetailVO) SetWans(v []string)`

SetWans sets Wans field to given value.

### HasWans

`func (o *VpnClientDetailVO) HasWans() bool`

HasWans returns a boolean if a field has been set.

### GetWorkingMode

`func (o *VpnClientDetailVO) GetWorkingMode() int32`

GetWorkingMode returns the WorkingMode field if non-nil, zero value otherwise.

### GetWorkingModeOk

`func (o *VpnClientDetailVO) GetWorkingModeOk() (*int32, bool)`

GetWorkingModeOk returns a tuple with the WorkingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingMode

`func (o *VpnClientDetailVO) SetWorkingMode(v int32)`

SetWorkingMode sets WorkingMode field to given value.

### HasWorkingMode

`func (o *VpnClientDetailVO) HasWorkingMode() bool`

HasWorkingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


